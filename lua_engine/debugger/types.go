package debugger

// StopReason 描述调试会话暂停原因。
type StopReason string

const (
	StopReasonEntrypoint StopReason = "entry"
	StopReasonBreakpoint StopReason = "breakpoint"
	StopReasonStep       StopReason = "step"
	StopReasonPause      StopReason = "pause"
	StopReasonException  StopReason = "exception"
)

// StepMode 描述下一次运行的单步策略。
type StepMode string

const (
	StepModeNone StepMode = ""
	StepModeInto StepMode = "into"
)

// Config 控制 Lua 调试器是否启用以及变量采集范围。
type Config struct {
	Enabled          bool
	BreakOnStart     bool
	BreakOnError     bool
	CollectGlobals   bool
	CollectLocals    bool
	MaxVariableDepth int
	EventBufferSize  int
}

// Breakpoint 表示一个源码断点。
type Breakpoint struct {
	ID       int    `json:"id"`
	File     string `json:"file"`
	Line     int    `json:"line"`
	Verified bool   `json:"verified"`
}

// Position 表示当前暂停所在源码位置。
type Position struct {
	File string `json:"file"`
	Line int    `json:"line"`
}

// Variable 表示调试变量快照。
type Variable struct {
	Name  string      `json:"name"`
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}

// Frame 表示调用栈中的一个栈帧。
type Frame struct {
	ID       int        `json:"id"`
	Name     string     `json:"name"`
	Position Position   `json:"position"`
	Locals   []Variable `json:"locals,omitempty"`
	Globals  []Variable `json:"globals,omitempty"`
}

// Event 表示 Debug Core 对外发出的状态事件。
type Event struct {
	Type     string     `json:"type"`
	Reason   StopReason `json:"reason,omitempty"`
	ThreadID int        `json:"threadId"`
	Position Position   `json:"position,omitempty"`
	Message  string     `json:"message,omitempty"`
	Frame    *Frame     `json:"frame,omitempty"`
}
