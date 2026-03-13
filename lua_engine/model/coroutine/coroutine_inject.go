package coroutine

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"

	lua "github.com/yuin/gopher-lua"
)

// CoroutineModule 协程模块
type CoroutineModule struct{}

// Name 返回模块名称
func (m *CoroutineModule) Name() string {
	return "coroutine"
}

// IsAvailable 检查模块是否可用
func (m *CoroutineModule) IsAvailable() bool {
	return true
}

// CoroutineState 协程状态
type CoroutineState int

const (
	StatePending   CoroutineState = iota // 等待中
	StateRunning                         // 运行中
	StateCompleted                       // 已完成
	StateCancelled                       // 已取消
	StateError                           // 错误
)

// String 返回状态的字符串表示
func (s CoroutineState) String() string {
	switch s {
	case StatePending:
		return "pending"
	case StateRunning:
		return "running"
	case StateCompleted:
		return "completed"
	case StateCancelled:
		return "cancelled"
	case StateError:
		return "error"
	default:
		return "unknown"
	}
}

// Coroutine 协程结构体
type Coroutine struct {
	ID          string
	Name        string
	Ctx         context.Context
	Cancel      context.CancelFunc
	StartedAt   time.Time
	CompletedAt time.Time
	State       CoroutineState
	Priority    int
	Error       error
	Result      lua.LValue
}

// CoroutinePool 协程池
type CoroutinePool struct {
	mu         sync.RWMutex
	name       string
	workerChan chan *Task
	workers    []*Worker
	taskQueue  []*Task
	maxWorkers int
	maxTasks   int
	active     int32
	closed     bool
}

// Task 任务结构体
type Task struct {
	ID        string
	Function  *lua.LFunction
	Args      []lua.LValue
	Priority  int
	CreatedAt time.Time
	Ctx       context.Context
	Cancel    context.CancelFunc
}

// Worker 工作协程
type Worker struct {
	ID      int
	Pool    *CoroutinePool
	Active  bool
	Current *Task
}

// Scheduler 调度器
type Scheduler struct {
	mu          sync.RWMutex
	strategy    string
	priorityMap map[string]int
}

// CoroutineManager 协程管理器
type CoroutineManager struct {
	mu         sync.RWMutex
	coroutines map[string]*Coroutine
	pools      map[string]*CoroutinePool
	scheduler  *Scheduler
	counter    int64
	totalTasks int64
	completed  int64
	failed     int64
	cancelled  int64
}

var manager = &CoroutineManager{
	coroutines: make(map[string]*Coroutine),
	pools:      make(map[string]*CoroutinePool),
	scheduler: &Scheduler{
		strategy:    "fifo",
		priorityMap: make(map[string]int),
	},
}

// NewCoroutinePool 创建协程池
func NewCoroutinePool(name string, maxWorkers, maxTasks int) *CoroutinePool {
	pool := &CoroutinePool{
		name:       name,
		workerChan: make(chan *Task, maxTasks),
		workers:    make([]*Worker, 0, maxWorkers),
		taskQueue:  make([]*Task, 0, maxTasks),
		maxWorkers: maxWorkers,
		maxTasks:   maxTasks,
		closed:     false,
	}

	for i := 0; i < maxWorkers; i++ {
		worker := &Worker{
			ID:   i,
			Pool: pool,
		}
		pool.workers = append(pool.workers, worker)
		go worker.run()
	}

	return pool
}

// run 工作协程运行
func (w *Worker) run() {
	for task := range w.Pool.workerChan {
		if w.Pool.closed {
			break
		}

		w.Active = true
		w.Current = task

		atomic.AddInt32(&w.Pool.active, 1)

		func() {
			defer func() {
				if r := recover(); r != nil {
					atomic.AddInt64(&manager.failed, 1)
				}
				w.Active = false
				w.Current = nil
				atomic.AddInt32(&w.Pool.active, -1)
			}()

			// 创建一个新的 Lua 状态来执行任务
			L := lua.NewState()
			defer L.Close()

			L.Push(task.Function)
			for _, arg := range task.Args {
				L.Push(arg)
			}
			L.Call(len(task.Args), 0)
		}()
	}
}

// Submit 提交任务到协程池
func (p *CoroutinePool) Submit(task *Task) bool {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.closed {
		return false
	}

	if len(p.taskQueue) >= p.maxTasks {
		return false
	}

	task.CreatedAt = time.Now()
	p.taskQueue = append(p.taskQueue, task)
	p.scheduleTasks()

	return true
}

// scheduleTasks 调度任务
func (p *CoroutinePool) scheduleTasks() {
	if len(p.taskQueue) == 0 {
		return
	}

	for len(p.taskQueue) > 0 && atomic.LoadInt32(&p.active) < int32(p.maxWorkers) {
		task := p.taskQueue[0]
		p.taskQueue = p.taskQueue[1:]

		select {
		case p.workerChan <- task:
			atomic.AddInt64(&manager.totalTasks, 1)
		default:
			p.taskQueue = append([]*Task{task}, p.taskQueue...)
			break
		}
	}
}

// Close 关闭协程池
func (p *CoroutinePool) Close() {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.closed {
		return
	}

	p.closed = true
	close(p.workerChan)

	for _, task := range p.taskQueue {
		if task.Cancel != nil {
			task.Cancel()
		}
	}
	p.taskQueue = nil
}

// GetStats 获取协程池统计信息
func (p *CoroutinePool) GetStats() map[string]interface{} {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return map[string]interface{}{
		"name":       p.name,
		"maxWorkers": p.maxWorkers,
		"maxTasks":   p.maxTasks,
		"active":     atomic.LoadInt32(&p.active),
		"queued":     len(p.taskQueue),
		"workers":    len(p.workers),
		"closed":     p.closed,
	}
}

// SetStrategy 设置调度策略
func (s *Scheduler) SetStrategy(strategy string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.strategy = strategy
}

// GetStrategy 获取调度策略
func (s *Scheduler) GetStrategy() string {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.strategy
}

// SetPriority 设置优先级
func (s *Scheduler) SetPriority(name string, priority int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.priorityMap[name] = priority
}

// GetPriority 获取优先级
func (s *Scheduler) GetPriority(name string) int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if priority, exists := s.priorityMap[name]; exists {
		return priority
	}
	return 0
}

// Register 向引擎注册方法
func (m *CoroutineModule) Register(engine model.Engine) error {
	state := engine.GetState()

	coroutineObj := state.NewTable()
	state.SetGlobal("coroutine", coroutineObj)

	// ========== 基础协程方法 ==========

	// launch - 启动一个新的协程
	coroutineObj.RawSetString("launch", state.NewFunction(func(L *lua.LState) int {
		if L.GetTop() == 0 {
			return 0
		}

		fn := L.CheckFunction(1)

		ctx, cancel := context.WithCancel(context.Background())
		coroutineID := fmt.Sprintf("coro_%d", atomic.AddInt64(&manager.counter, 1))

		name := ""
		if L.GetTop() >= 2 {
			name = L.CheckString(2)
		}

		priority := 0
		if L.GetTop() >= 3 {
			priority = L.CheckInt(3)
		}

		coroutine := &Coroutine{
			ID:        coroutineID,
			Name:      name,
			Ctx:       ctx,
			Cancel:    cancel,
			StartedAt: time.Now(),
			State:     StateRunning,
			Priority:  priority,
		}

		manager.mu.Lock()
		manager.coroutines[coroutineID] = coroutine
		manager.mu.Unlock()

		go func() {
			defer func() {
				if r := recover(); r != nil {
					coroutine.State = StateError
					coroutine.Error = fmt.Errorf("协程执行出错: %v", r)
					atomic.AddInt64(&manager.failed, 1)
				} else {
					coroutine.State = StateCompleted
					atomic.AddInt64(&manager.completed, 1)
				}
				coroutine.CompletedAt = time.Now()

				manager.mu.Lock()
				delete(manager.coroutines, coroutineID)
				manager.mu.Unlock()
			}()

			select {
			case <-ctx.Done():
				coroutine.State = StateCancelled
				atomic.AddInt64(&manager.cancelled, 1)
				return
			default:
				// 创建一个新的 Lua 状态来执行协程
				coroL := lua.NewState()
				defer coroL.Close()

				coroL.Push(fn)
				coroL.Call(0, 0)
			}
		}()

		L.Push(lua.LString(coroutineID))
		return 1
	}))

	// delay - 延迟执行
	coroutineObj.RawSetString("delay", state.NewFunction(func(L *lua.LState) int {
		if L.GetTop() < 2 {
			return 0
		}

		delayMs := L.CheckInt(1)
		fn := L.CheckFunction(2)

		ctx, cancel := context.WithCancel(context.Background())
		coroutineID := fmt.Sprintf("delay_%d", atomic.AddInt64(&manager.counter, 1))

		coroutine := &Coroutine{
			ID:        coroutineID,
			Name:      "delay",
			Ctx:       ctx,
			Cancel:    cancel,
			StartedAt: time.Now(),
			State:     StatePending,
		}

		manager.mu.Lock()
		manager.coroutines[coroutineID] = coroutine
		manager.mu.Unlock()

		go func() {
			defer func() {
				if r := recover(); r != nil {
					coroutine.State = StateError
					coroutine.Error = fmt.Errorf("延迟执行出错: %v", r)
					atomic.AddInt64(&manager.failed, 1)
				} else {
					coroutine.State = StateCompleted
					atomic.AddInt64(&manager.completed, 1)
				}
				coroutine.CompletedAt = time.Now()

				manager.mu.Lock()
				delete(manager.coroutines, coroutineID)
				manager.mu.Unlock()
			}()

			select {
			case <-ctx.Done():
				coroutine.State = StateCancelled
				atomic.AddInt64(&manager.cancelled, 1)
				return
			case <-time.After(time.Duration(delayMs) * time.Millisecond):
				coroutine.State = StateRunning
				// 创建一个新的 Lua 状态来执行延迟函数
				delayL := lua.NewState()
				defer delayL.Close()

				delayL.Push(fn)
				delayL.Call(0, 0)
			}
		}()

		L.Push(lua.LString(coroutineID))
		return 1
	}))

	// async - 异步执行函数并返回结果（同步等待）
	coroutineObj.RawSetString("async", state.NewFunction(func(L *lua.LState) int {
		if L.GetTop() == 0 {
			return 0
		}

		fn := L.CheckFunction(1)

		// 创建一个 channel 来同步等待结果
		resultChan := make(chan lua.LValue, 1)
		errorChan := make(chan error, 1)

		go func() {
			defer func() {
				if r := recover(); r != nil {
					errorChan <- fmt.Errorf("协程执行出错: %v", r)
					atomic.AddInt64(&manager.failed, 1)
				}
			}()

			// 创建一个新的 Lua 状态来执行异步函数
			asyncL := lua.NewState()
			defer asyncL.Close()

			asyncL.Push(fn)
			asyncL.Call(0, 1)
			result := asyncL.Get(-1)
			resultChan <- result
			atomic.AddInt64(&manager.completed, 1)
		}()

		// 等待结果或错误
		select {
		case result := <-resultChan:
			L.Push(result)
			return 1
		case err := <-errorChan:
			L.Push(lua.LString(err.Error()))
			return 1
		}
	}))

	// await - 等待并返回传入的值（简化版本）
	coroutineObj.RawSetString("await", state.NewFunction(func(L *lua.LState) int {
		if L.GetTop() == 0 {
			return 0
		}

		value := L.CheckAny(1)

		// 如果是 nil，直接返回 nil
		if value == lua.LNil {
			return 0
		}

		// 简化版本：直接返回传入的值
		L.Push(value)
		return 1
	}))

	// sleep - 在协程中睡眠
	coroutineObj.RawSetString("sleep", state.NewFunction(func(L *lua.LState) int {
		if L.GetTop() == 0 {
			return 0
		}

		delayMs := L.CheckInt(1)
		time.Sleep(time.Duration(delayMs) * time.Millisecond)
		return 0
	}))

	// ========== 协程管理方法 ==========

	// cancel - 取消协程
	coroutineObj.RawSetString("cancel", state.NewFunction(func(L *lua.LState) int {
		if L.GetTop() == 0 {
			L.Push(lua.LBool(false))
			return 1
		}

		coroutineID := L.CheckString(1)

		manager.mu.RLock()
		coroutine, exists := manager.coroutines[coroutineID]
		manager.mu.RUnlock()

		if exists && coroutine.Cancel != nil {
			coroutine.Cancel()
			L.Push(lua.LBool(true))
			return 1
		}

		L.Push(lua.LBool(false))
		return 1
	}))

	// getActiveCoroutines - 获取活跃的协程数量
	coroutineObj.RawSetString("getActiveCoroutines", state.NewFunction(func(L *lua.LState) int {
		manager.mu.RLock()
		count := len(manager.coroutines)
		manager.mu.RUnlock()
		L.Push(lua.LNumber(count))
		return 1
	}))

	// getCoroutineList - 获取协程列表
	coroutineObj.RawSetString("getCoroutineList", state.NewFunction(func(L *lua.LState) int {
		manager.mu.RLock()
		defer manager.mu.RUnlock()

		table := L.NewTable()
		i := 1
		for id, coro := range manager.coroutines {
			info := L.NewTable()
			info.RawSetString("id", lua.LString(id))
			info.RawSetString("name", lua.LString(coro.Name))
			info.RawSetString("state", lua.LString(coro.State.String()))
			info.RawSetString("priority", lua.LNumber(coro.Priority))
			info.RawSetString("duration", lua.LNumber(time.Since(coro.StartedAt).Milliseconds()))
			table.RawSetInt(i, info)
			i++
		}

		L.Push(table)
		return 1
	}))

	// getCoroutineInfo - 获取指定协程的详细信息
	coroutineObj.RawSetString("getCoroutineInfo", state.NewFunction(func(L *lua.LState) int {
		if L.GetTop() == 0 {
			return 0
		}

		coroutineID := L.CheckString(1)

		manager.mu.RLock()
		coroutine, exists := manager.coroutines[coroutineID]
		manager.mu.RUnlock()

		if !exists {
			return 0
		}

		info := L.NewTable()
		info.RawSetString("id", lua.LString(coroutine.ID))
		info.RawSetString("name", lua.LString(coroutine.Name))
		info.RawSetString("state", lua.LString(coroutine.State.String()))
		info.RawSetString("priority", lua.LNumber(coroutine.Priority))
		info.RawSetString("startedAt", lua.LString(coroutine.StartedAt.Format("2006-01-02 15:04:05")))

		if !coroutine.CompletedAt.IsZero() {
			info.RawSetString("completedAt", lua.LString(coroutine.CompletedAt.Format("2006-01-02 15:04:05")))
			info.RawSetString("duration", lua.LNumber(coroutine.CompletedAt.Sub(coroutine.StartedAt).Milliseconds()))
		} else {
			info.RawSetString("duration", lua.LNumber(time.Since(coroutine.StartedAt).Milliseconds()))
		}

		if coroutine.Error != nil {
			info.RawSetString("error", lua.LString(coroutine.Error.Error()))
		}

		if coroutine.Result != nil {
			info.RawSetString("result", coroutine.Result)
		}

		L.Push(info)
		return 1
	}))

	// cancelAll - 取消所有协程
	coroutineObj.RawSetString("cancelAll", state.NewFunction(func(L *lua.LState) int {
		manager.mu.Lock()
		defer manager.mu.Unlock()

		count := 0
		for _, coro := range manager.coroutines {
			if coro.Cancel != nil {
				coro.Cancel()
				count++
			}
		}

		L.Push(lua.LNumber(count))
		return 1
	}))

	// getStats - 获取全局统计信息
	coroutineObj.RawSetString("getStats", state.NewFunction(func(L *lua.LState) int {
		stats := L.NewTable()
		stats.RawSetString("totalTasks", lua.LNumber(atomic.LoadInt64(&manager.totalTasks)))
		stats.RawSetString("completed", lua.LNumber(atomic.LoadInt64(&manager.completed)))
		stats.RawSetString("failed", lua.LNumber(atomic.LoadInt64(&manager.failed)))
		stats.RawSetString("cancelled", lua.LNumber(atomic.LoadInt64(&manager.cancelled)))
		stats.RawSetString("active", lua.LNumber(len(manager.coroutines)))
		stats.RawSetString("pools", lua.LNumber(len(manager.pools)))
		L.Push(stats)
		return 1
	}))

	// ========== 协程池方法 ==========

	// createPool - 创建协程池
	coroutineObj.RawSetString("createPool", state.NewFunction(func(L *lua.LState) int {
		if L.GetTop() < 2 {
			return 0
		}

		name := L.CheckString(1)
		maxWorkers := L.CheckInt(2)

		maxTasks := 100
		if L.GetTop() >= 3 {
			maxTasks = L.CheckInt(3)
		}

		pool := NewCoroutinePool(name, maxWorkers, maxTasks)

		manager.mu.Lock()
		manager.pools[name] = pool
		manager.mu.Unlock()

		L.Push(lua.LString(name))
		return 1
	}))

	// submitToPool - 提交任务到协程池
	coroutineObj.RawSetString("submitToPool", state.NewFunction(func(L *lua.LState) int {
		if L.GetTop() < 2 {
			L.Push(lua.LBool(false))
			return 1
		}

		poolName := L.CheckString(1)
		fn := L.CheckFunction(2)

		manager.mu.RLock()
		pool, exists := manager.pools[poolName]
		manager.mu.RUnlock()

		if !exists {
			L.Push(lua.LBool(false))
			return 1
		}

		ctx, cancel := context.WithCancel(context.Background())
		taskID := fmt.Sprintf("task_%d", atomic.AddInt64(&manager.counter, 1))

		priority := 0
		if L.GetTop() >= 3 {
			priority = L.CheckInt(3)
		}

		// 获取参数
		args := make([]lua.LValue, 0)
		if L.GetTop() >= 4 {
			table := L.CheckTable(4)
			table.ForEach(func(key, value lua.LValue) {
				args = append(args, value)
			})
		}

		task := &Task{
			ID:       taskID,
			Function: fn,
			Args:     args,
			Priority: priority,
			Ctx:      ctx,
			Cancel:   cancel,
		}

		success := pool.Submit(task)
		L.Push(lua.LBool(success))
		return 1
	}))

	// getPoolStats - 获取协程池统计信息
	coroutineObj.RawSetString("getPoolStats", state.NewFunction(func(L *lua.LState) int {
		if L.GetTop() == 0 {
			return 0
		}

		poolName := L.CheckString(1)

		manager.mu.RLock()
		pool, exists := manager.pools[poolName]
		manager.mu.RUnlock()

		if !exists {
			return 0
		}

		stats := pool.GetStats()
		statsObj := L.NewTable()
		statsObj.RawSetString("name", lua.LString(stats["name"].(string)))
		statsObj.RawSetString("maxWorkers", lua.LNumber(stats["maxWorkers"].(int)))
		statsObj.RawSetString("maxTasks", lua.LNumber(stats["maxTasks"].(int)))
		statsObj.RawSetString("active", lua.LNumber(stats["active"].(int32)))
		statsObj.RawSetString("queued", lua.LNumber(stats["queued"].(int)))
		statsObj.RawSetString("workers", lua.LNumber(stats["workers"].(int)))
		statsObj.RawSetString("closed", lua.LBool(stats["closed"].(bool)))

		L.Push(statsObj)
		return 1
	}))

	// closePool - 关闭协程池
	coroutineObj.RawSetString("closePool", state.NewFunction(func(L *lua.LState) int {
		if L.GetTop() == 0 {
			L.Push(lua.LBool(false))
			return 1
		}

		poolName := L.CheckString(1)

		manager.mu.Lock()
		pool, exists := manager.pools[poolName]
		if exists {
			delete(manager.pools, poolName)
		}
		manager.mu.Unlock()

		if exists {
			pool.Close()
			L.Push(lua.LBool(true))
			return 1
		}

		L.Push(lua.LBool(false))
		return 1
	}))

	// listPools - 列出所有协程池
	coroutineObj.RawSetString("listPools", state.NewFunction(func(L *lua.LState) int {
		manager.mu.RLock()
		defer manager.mu.RUnlock()

		table := L.NewTable()
		i := 1
		for name := range manager.pools {
			table.RawSetInt(i, lua.LString(name))
			i++
		}

		L.Push(table)
		return 1
	}))

	// ========== 调度器方法 ==========

	// setScheduleStrategy - 设置调度策略
	coroutineObj.RawSetString("setScheduleStrategy", state.NewFunction(func(L *lua.LState) int {
		if L.GetTop() == 0 {
			return 0
		}

		strategy := L.CheckString(1)
		manager.scheduler.SetStrategy(strategy)
		return 0
	}))

	// getScheduleStrategy - 获取当前调度策略
	coroutineObj.RawSetString("getScheduleStrategy", state.NewFunction(func(L *lua.LState) int {
		strategy := manager.scheduler.GetStrategy()
		L.Push(lua.LString(strategy))
		return 1
	}))

	// setPriority - 设置协程优先级
	coroutineObj.RawSetString("setPriority", state.NewFunction(func(L *lua.LState) int {
		if L.GetTop() < 2 {
			return 0
		}

		name := L.CheckString(1)
		priority := L.CheckInt(2)
		manager.scheduler.SetPriority(name, priority)
		return 0
	}))

	// getPriority - 获取协程优先级
	coroutineObj.RawSetString("getPriority", state.NewFunction(func(L *lua.LState) int {
		if L.GetTop() == 0 {
			return 0
		}

		name := L.CheckString(1)
		priority := manager.scheduler.GetPriority(name)
		L.Push(lua.LNumber(priority))
		return 1
	}))

	// 注册方法文档
	engine.RegisterMethod("coroutine.launch", "启动一个新的协程", nil, true)
	engine.RegisterMethod("coroutine.delay", "延迟执行函数", nil, true)
	engine.RegisterMethod("coroutine.async", "异步执行函数并返回结果（同步等待）", nil, true)
	engine.RegisterMethod("coroutine.await", "等待并返回传入的值（简化版本）", nil, true)
	engine.RegisterMethod("coroutine.cancel", "取消指定的协程", nil, true)
	engine.RegisterMethod("coroutine.sleep", "在协程中睡眠指定的毫秒数", nil, true)
	engine.RegisterMethod("coroutine.getActiveCoroutines", "获取活跃的协程数量", nil, true)
	engine.RegisterMethod("coroutine.getCoroutineList", "获取协程列表", nil, true)
	engine.RegisterMethod("coroutine.getCoroutineInfo", "获取指定协程的详细信息", nil, true)
	engine.RegisterMethod("coroutine.cancelAll", "取消所有协程", nil, true)
	engine.RegisterMethod("coroutine.getStats", "获取全局统计信息", nil, true)
	engine.RegisterMethod("coroutine.createPool", "创建协程池", nil, true)
	engine.RegisterMethod("coroutine.submitToPool", "提交任务到协程池", nil, true)
	engine.RegisterMethod("coroutine.getPoolStats", "获取协程池统计信息", nil, true)
	engine.RegisterMethod("coroutine.closePool", "关闭协程池", nil, true)
	engine.RegisterMethod("coroutine.listPools", "列出所有协程池", nil, true)
	engine.RegisterMethod("coroutine.setScheduleStrategy", "设置调度策略", nil, true)
	engine.RegisterMethod("coroutine.getScheduleStrategy", "获取当前调度策略", nil, true)
	engine.RegisterMethod("coroutine.setPriority", "设置协程优先级", nil, true)
	engine.RegisterMethod("coroutine.getPriority", "获取协程优先级", nil, true)

	return nil
}
