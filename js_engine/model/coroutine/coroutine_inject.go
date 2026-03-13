package coroutine

import (
	"app/js_engine/model"
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/dop251/goja"
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
	StatePending CoroutineState = iota // 等待中
	StateRunning                       // 运行中
	StateCompleted                     // 已完成
	StateCancelled                     // 已取消
	StateError                         // 错误
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
	ID         string
	Name       string
	Ctx        context.Context
	Cancel     context.CancelFunc
	VM         *goja.Runtime
	StartedAt  time.Time
	CompletedAt time.Time
	State      CoroutineState
	Priority   int
	Error      error
	Result     interface{}
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
	Function  func(goja.FunctionCall) goja.Value
	Args      goja.FunctionCall
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
	mu          sync.RWMutex
	coroutines  map[string]*Coroutine
	pools       map[string]*CoroutinePool
	scheduler   *Scheduler
	counter     int64
	totalTasks  int64
	completed   int64
	failed      int64
	cancelled   int64
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

			task.Function(task.Args)
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
		"name":        p.name,
		"maxWorkers":  p.maxWorkers,
		"maxTasks":    p.maxTasks,
		"active":      atomic.LoadInt32(&p.active),
		"queued":      len(p.taskQueue),
		"workers":     len(p.workers),
		"closed":      p.closed,
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
	vm := engine.GetVM()

	coroutineObj := vm.NewObject()
	vm.Set("coroutine", coroutineObj)

	// ========== 基础协程方法 ==========

	// launch - 启动一个新的协程
	coroutineObj.Set("launch", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) == 0 {
			return goja.Undefined()
		}

		fn, ok := call.Argument(0).Export().(func(goja.FunctionCall) goja.Value)
		if !ok {
			return goja.Undefined()
		}

		ctx, cancel := context.WithCancel(context.Background())
		coroutineID := fmt.Sprintf("coro_%d", atomic.AddInt64(&manager.counter, 1))

		name := ""
		if len(call.Arguments) > 1 {
			name = call.Argument(1).String()
		}

		priority := 0
		if len(call.Arguments) > 2 {
			priority = int(call.Argument(2).ToInteger())
		}

		coroutine := &Coroutine{
			ID:        coroutineID,
			Name:      name,
			Ctx:       ctx,
			Cancel:    cancel,
			VM:        vm,
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
				result := fn(call)
				coroutine.Result = result
			}
		}()

		return vm.ToValue(coroutineID)
	})

	// delay - 延迟执行
	coroutineObj.Set("delay", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 2 {
			return goja.Undefined()
		}

		delayMs := int(call.Argument(0).ToInteger())
		fn, ok := call.Argument(1).Export().(func(goja.FunctionCall) goja.Value)
		if !ok {
			return goja.Undefined()
		}

		ctx, cancel := context.WithCancel(context.Background())
		coroutineID := fmt.Sprintf("delay_%d", atomic.AddInt64(&manager.counter, 1))

		coroutine := &Coroutine{
			ID:        coroutineID,
			Name:      "delay",
			Ctx:       ctx,
			Cancel:    cancel,
			VM:        vm,
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
				result := fn(call)
				coroutine.Result = result
			}
		}()

		return vm.ToValue(coroutineID)
	})

	// async - 异步执行函数并返回结果（同步等待）
	coroutineObj.Set("async", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) == 0 {
			return goja.Undefined()
		}

		fn, ok := call.Argument(0).Export().(func(goja.FunctionCall) goja.Value)
		if !ok {
			return goja.Undefined()
		}

		// 创建一个 channel 来同步等待结果
		resultChan := make(chan goja.Value, 1)
		errorChan := make(chan error, 1)

		go func() {
			defer func() {
				if r := recover(); r != nil {
					errorChan <- fmt.Errorf("协程执行出错: %v", r)
					atomic.AddInt64(&manager.failed, 1)
				}
			}()

			result := fn(call)
			resultChan <- result
			atomic.AddInt64(&manager.completed, 1)
		}()

		// 等待结果或错误
		select {
		case result := <-resultChan:
			return result
		case err := <-errorChan:
			return vm.ToValue(err.Error())
		}
	})

	// await - 等待并返回传入的值（简化版本）
	coroutineObj.Set("await", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) == 0 {
			return goja.Undefined()
		}

		value := call.Argument(0)
		
		// 如果是 undefined 或 null，直接返回
		if goja.IsUndefined(value) || goja.IsNull(value) {
			return goja.Undefined()
		}
		
		// 简化版本：直接返回传入的值
		// 不尝试调用 then 方法，避免空指针错误
		return value
	})

	// sleep - 在协程中睡眠
	coroutineObj.Set("sleep", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) == 0 {
			return goja.Undefined()
		}

		delayMs := int(call.Argument(0).ToInteger())
		time.Sleep(time.Duration(delayMs) * time.Millisecond)
		return goja.Undefined()
	})

	// ========== 协程管理方法 ==========

	// cancel - 取消协程
	coroutineObj.Set("cancel", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) == 0 {
			return goja.Undefined()
		}

		coroutineID := call.Argument(0).String()

		manager.mu.RLock()
		coroutine, exists := manager.coroutines[coroutineID]
		manager.mu.RUnlock()

		if exists && coroutine.Cancel != nil {
			coroutine.Cancel()
			return vm.ToValue(true)
		}

		return vm.ToValue(false)
	})

	// getActiveCoroutines - 获取活跃的协程数量
	coroutineObj.Set("getActiveCoroutines", func(call goja.FunctionCall) goja.Value {
		manager.mu.RLock()
		count := len(manager.coroutines)
		manager.mu.RUnlock()
		return vm.ToValue(count)
	})

	// getCoroutineList - 获取协程列表
	coroutineObj.Set("getCoroutineList", func(call goja.FunctionCall) goja.Value {
		manager.mu.RLock()
		defer manager.mu.RUnlock()

		list := vm.NewArray()
		i := 0
		for id, coro := range manager.coroutines {
			info := vm.NewObject()
			info.Set("id", id)
			info.Set("name", coro.Name)
			info.Set("state", coro.State.String())
			info.Set("priority", coro.Priority)
			info.Set("duration", time.Since(coro.StartedAt).Milliseconds())
			list.Set(fmt.Sprintf("%d", i), info)
			i++
		}

		return list
	})

	// getCoroutineInfo - 获取指定协程的详细信息
	coroutineObj.Set("getCoroutineInfo", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) == 0 {
			return goja.Undefined()
		}

		coroutineID := call.Argument(0).String()

		manager.mu.RLock()
		coroutine, exists := manager.coroutines[coroutineID]
		manager.mu.RUnlock()

		if !exists {
			return goja.Undefined()
		}

		info := vm.NewObject()
		info.Set("id", coroutine.ID)
		info.Set("name", coroutine.Name)
		info.Set("state", coroutine.State.String())
		info.Set("priority", coroutine.Priority)
		info.Set("startedAt", coroutine.StartedAt.Format("2006-01-02 15:04:05"))

		if !coroutine.CompletedAt.IsZero() {
			info.Set("completedAt", coroutine.CompletedAt.Format("2006-01-02 15:04:05"))
			info.Set("duration", coroutine.CompletedAt.Sub(coroutine.StartedAt).Milliseconds())
		} else {
			info.Set("duration", time.Since(coroutine.StartedAt).Milliseconds())
		}

		if coroutine.Error != nil {
			info.Set("error", coroutine.Error.Error())
		}

		if coroutine.Result != nil {
			info.Set("result", fmt.Sprintf("%v", coroutine.Result))
		}

		return info
	})

	// cancelAll - 取消所有协程
	coroutineObj.Set("cancelAll", func(call goja.FunctionCall) goja.Value {
		manager.mu.Lock()
		defer manager.mu.Unlock()

		count := 0
		for _, coro := range manager.coroutines {
			if coro.Cancel != nil {
				coro.Cancel()
				count++
			}
		}

		return vm.ToValue(count)
	})

	// getStats - 获取全局统计信息
	coroutineObj.Set("getStats", func(call goja.FunctionCall) goja.Value {
		stats := vm.NewObject()
		stats.Set("totalTasks", atomic.LoadInt64(&manager.totalTasks))
		stats.Set("completed", atomic.LoadInt64(&manager.completed))
		stats.Set("failed", atomic.LoadInt64(&manager.failed))
		stats.Set("cancelled", atomic.LoadInt64(&manager.cancelled))
		stats.Set("active", len(manager.coroutines))
		stats.Set("pools", len(manager.pools))
		return stats
	})

	// ========== 协程池方法 ==========

	// createPool - 创建协程池
	coroutineObj.Set("createPool", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 2 {
			return goja.Undefined()
		}

		name := call.Argument(0).String()
		maxWorkers := int(call.Argument(1).ToInteger())

		maxTasks := 100
		if len(call.Arguments) > 2 {
			maxTasks = int(call.Argument(2).ToInteger())
		}

		pool := NewCoroutinePool(name, maxWorkers, maxTasks)

		manager.mu.Lock()
		manager.pools[name] = pool
		manager.mu.Unlock()

		return vm.ToValue(name)
	})

	// submitToPool - 提交任务到协程池
	coroutineObj.Set("submitToPool", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 2 {
			return goja.Undefined()
		}

		poolName := call.Argument(0).String()
		fn, ok := call.Argument(1).Export().(func(goja.FunctionCall) goja.Value)
		if !ok {
			return goja.Undefined()
		}

		manager.mu.RLock()
		pool, exists := manager.pools[poolName]
		manager.mu.RUnlock()

		if !exists {
			return goja.Undefined()
		}

		ctx, cancel := context.WithCancel(context.Background())
		taskID := fmt.Sprintf("task_%d", atomic.AddInt64(&manager.counter, 1))

		priority := 0
		if len(call.Arguments) > 2 {
			priority = int(call.Argument(2).ToInteger())
		}

		task := &Task{
			ID:       taskID,
			Function: fn,
			Args:     call,
			Priority: priority,
			Ctx:      ctx,
			Cancel:   cancel,
		}

		success := pool.Submit(task)
		return vm.ToValue(success)
	})

	// getPoolStats - 获取协程池统计信息
	coroutineObj.Set("getPoolStats", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) == 0 {
			return goja.Undefined()
		}

		poolName := call.Argument(0).String()

		manager.mu.RLock()
		pool, exists := manager.pools[poolName]
		manager.mu.RUnlock()

		if !exists {
			return goja.Undefined()
		}

		stats := pool.GetStats()
		statsObj := vm.NewObject()
		statsObj.Set("name", stats["name"])
		statsObj.Set("maxWorkers", stats["maxWorkers"])
		statsObj.Set("maxTasks", stats["maxTasks"])
		statsObj.Set("active", stats["active"])
		statsObj.Set("queued", stats["queued"])
		statsObj.Set("workers", stats["workers"])
		statsObj.Set("closed", stats["closed"])

		return statsObj
	})

	// closePool - 关闭协程池
	coroutineObj.Set("closePool", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) == 0 {
			return goja.Undefined()
		}

		poolName := call.Argument(0).String()

		manager.mu.Lock()
		pool, exists := manager.pools[poolName]
		if exists {
			delete(manager.pools, poolName)
		}
		manager.mu.Unlock()

		if exists {
			pool.Close()
			return vm.ToValue(true)
		}

		return vm.ToValue(false)
	})

	// listPools - 列出所有协程池
	coroutineObj.Set("listPools", func(call goja.FunctionCall) goja.Value {
		manager.mu.RLock()
		defer manager.mu.RUnlock()

		list := vm.NewArray()
		i := 0
		for name := range manager.pools {
			list.Set(fmt.Sprintf("%d", i), name)
			i++
		}

		return list
	})

	// ========== 调度器方法 ==========

	// setScheduleStrategy - 设置调度策略
	coroutineObj.Set("setScheduleStrategy", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) == 0 {
			return goja.Undefined()
		}

		strategy := call.Argument(0).String()
		manager.scheduler.SetStrategy(strategy)
		return goja.Undefined()
	})

	// getScheduleStrategy - 获取当前调度策略
	coroutineObj.Set("getScheduleStrategy", func(call goja.FunctionCall) goja.Value {
		strategy := manager.scheduler.GetStrategy()
		return vm.ToValue(strategy)
	})

	// setPriority - 设置协程优先级
	coroutineObj.Set("setPriority", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) < 2 {
			return goja.Undefined()
		}

		name := call.Argument(0).String()
		priority := int(call.Argument(1).ToInteger())
		manager.scheduler.SetPriority(name, priority)
		return goja.Undefined()
	})

	// getPriority - 获取协程优先级
	coroutineObj.Set("getPriority", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) == 0 {
			return goja.Undefined()
		}

		name := call.Argument(0).String()
		priority := manager.scheduler.GetPriority(name)
		return vm.ToValue(priority)
	})

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
