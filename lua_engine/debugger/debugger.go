package debugger

import (
	"path/filepath"
	"sync"

	lua "github.com/yuin/gopher-lua"
)

const (
	defaultThreadID = 1
	hitFunction     = "__autogo_debug_hit"
)

// Debugger 提供 Lua 源码插桩后的断点、暂停和单步控制。
type Debugger struct {
	mu          sync.Mutex
	config      Config
	breakpoints map[string]map[int]Breakpoint
	nextBPID    int
	events      chan Event
	resume      chan struct{}
	paused      bool
	stopped     bool
	pauseNext   bool
	stepMode    StepMode
	lastFrame   *Frame
	lastPos     Position
}

// New 创建一个 Lua Debug Core。
func New(config Config) *Debugger {
	if config.EventBufferSize <= 0 {
		config.EventBufferSize = 32
	}
	if config.MaxVariableDepth <= 0 {
		config.MaxVariableDepth = 2
	}
	d := &Debugger{
		config:      config,
		breakpoints: make(map[string]map[int]Breakpoint),
		nextBPID:    1,
		events:      make(chan Event, config.EventBufferSize),
		resume:      make(chan struct{}),
	}
	if config.BreakOnStart {
		d.pauseNext = true
	}
	return d
}

// Enabled 返回调试器是否启用。
func (d *Debugger) Enabled() bool {
	return d != nil && d.config.Enabled
}

// Install 将调试命中函数注入到 Lua 状态机。
func (d *Debugger) Install(L *lua.LState) {
	if !d.Enabled() || L == nil {
		return
	}
	L.SetGlobal(hitFunction, L.NewFunction(func(L *lua.LState) int {
		file := L.CheckString(1)
		line := L.CheckInt(2)
		d.Hit(L, file, line)
		return 0
	}))
}

// Events 返回调试事件通道。
func (d *Debugger) Events() <-chan Event {
	return d.events
}

// SetBreakpoints 替换指定文件的断点集合。
func (d *Debugger) SetBreakpoints(file string, lines []int) []Breakpoint {
	d.mu.Lock()
	defer d.mu.Unlock()

	file = normalizeFile(file)
	points := make(map[int]Breakpoint, len(lines))
	result := make([]Breakpoint, 0, len(lines))
	for _, line := range lines {
		if line <= 0 {
			continue
		}
		bp := Breakpoint{
			ID:       d.nextBPID,
			File:     file,
			Line:     line,
			Verified: true,
		}
		d.nextBPID++
		points[line] = bp
		result = append(result, bp)
	}
	d.breakpoints[file] = points
	return result
}

// Continue 恢复暂停中的脚本。
func (d *Debugger) Continue() {
	d.resumeWith(StepModeNone)
}

// StepInto 恢复脚本并在下一条插桩语句处暂停。
func (d *Debugger) StepInto() {
	d.resumeWith(StepModeInto)
}

// Pause 请求脚本在下一条插桩语句处暂停。
func (d *Debugger) Pause() {
	d.mu.Lock()
	d.pauseNext = true
	d.mu.Unlock()
}

// Stop 请求脚本在下一条插桩语句处停止。
func (d *Debugger) Stop() {
	d.mu.Lock()
	d.stopped = true
	paused := d.paused
	d.mu.Unlock()
	if paused {
		d.signalResume()
	}
}

// LastFrame 返回最近一次暂停时的栈帧快照。
func (d *Debugger) LastFrame() *Frame {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.lastFrame
}

// Hit 由插桩后的 Lua 源码在语句执行前调用。
func (d *Debugger) Hit(L *lua.LState, file string, line int) {
	if !d.Enabled() {
		return
	}

	if d.isStopped() {
		d.publish(Event{Type: "terminated", ThreadID: defaultThreadID})
		L.RaiseError("stopped by debugger")
		return
	}

	pos := Position{File: normalizeFile(file), Line: line}
	reason, shouldStop := d.shouldPause(pos)
	if !shouldStop {
		return
	}

	frame := d.captureFrame(L, pos)
	d.enterPaused(reason, pos, frame)
	d.waitResume()
	if d.isStopped() {
		d.publish(Event{Type: "terminated", ThreadID: defaultThreadID})
		L.RaiseError("stopped by debugger")
	}
}

// NotifyError 在脚本执行异常时向调试客户端发送异常事件。
func (d *Debugger) NotifyError(file string, err error) {
	if !d.Enabled() || err == nil || !d.config.BreakOnError {
		return
	}
	pos := Position{File: normalizeFile(file)}
	d.publish(Event{
		Type:     "stopped",
		Reason:   StopReasonException,
		ThreadID: defaultThreadID,
		Position: pos,
		Message:  err.Error(),
	})
}

func (d *Debugger) shouldPause(pos Position) (StopReason, bool) {
	d.mu.Lock()
	defer d.mu.Unlock()

	if d.stopped {
		return StopReasonPause, true
	}
	if d.pauseNext {
		d.pauseNext = false
		return StopReasonPause, true
	}
	if d.stepMode == StepModeInto {
		d.stepMode = StepModeNone
		return StopReasonStep, true
	}
	if byLine, ok := d.breakpoints[pos.File]; ok {
		if _, ok := byLine[pos.Line]; ok {
			return StopReasonBreakpoint, true
		}
	}
	return "", false
}

func (d *Debugger) enterPaused(reason StopReason, pos Position, frame *Frame) {
	d.mu.Lock()
	d.paused = true
	d.lastPos = pos
	d.lastFrame = frame
	d.mu.Unlock()

	d.publish(Event{
		Type:     "stopped",
		Reason:   reason,
		ThreadID: defaultThreadID,
		Position: pos,
		Frame:    frame,
	})
}

func (d *Debugger) waitResume() {
	<-d.resume
	d.mu.Lock()
	d.paused = false
	d.resume = make(chan struct{})
	d.mu.Unlock()
}

func (d *Debugger) resumeWith(mode StepMode) {
	d.mu.Lock()
	d.stepMode = mode
	paused := d.paused
	d.mu.Unlock()
	if paused {
		d.signalResume()
	}
}

func (d *Debugger) signalResume() {
	d.mu.Lock()
	ch := d.resume
	d.mu.Unlock()
	select {
	case ch <- struct{}{}:
	default:
	}
}

func (d *Debugger) isStopped() bool {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.stopped
}

func (d *Debugger) publish(event Event) {
	select {
	case d.events <- event:
	default:
	}
}

func normalizeFile(file string) string {
	if file == "" {
		return "<string>"
	}
	cleaned := filepath.Clean(file)
	if abs, err := filepath.Abs(cleaned); err == nil {
		return abs
	}
	return cleaned
}
