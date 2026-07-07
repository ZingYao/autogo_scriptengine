package coroutine

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"
)

var errNilCoroutineTask = errors.New("coroutine task is nil")

// CoroutineModule 是 go-lua-vm 迁移后的模块壳。
type CoroutineModule struct {
	mu               sync.Mutex
	counter          int64
	cancellations    map[string]context.CancelFunc
	pools            map[string]map[string]interface{}
	priorities       map[string]int
	scheduleStrategy string
}

func New() *CoroutineModule { return &CoroutineModule{} }

func (m *CoroutineModule) Name() string { return "coroutine" }

func (m *CoroutineModule) IsAvailable() bool { return true }

func (m *CoroutineModule) Register(engine model.Engine) error {
	if m.cancellations == nil {
		m.cancellations = make(map[string]context.CancelFunc)
	}
	if m.pools == nil {
		m.pools = make(map[string]map[string]interface{})
	}
	if m.priorities == nil {
		m.priorities = make(map[string]int)
	}
	engine.RegisterMethod("coroutine.sleep", "暂停当前协程", func(milliseconds int) {
		if milliseconds > 0 {
			time.Sleep(time.Duration(milliseconds) * time.Millisecond)
		}
	}, true)
	engine.RegisterMethod("coroutine.launch", "启动一个新的协程", func(task func() (interface{}, error)) string {
		id, _ := m.launch(task)
		return id
	}, true)
	engine.RegisterMethod("coroutine.delay", "延迟执行函数", func(milliseconds int, task func() (interface{}, error)) string {
		id, _ := m.delay(milliseconds, task)
		return id
	}, true)
	engine.RegisterMethod("coroutine.cancel", "取消指定协程", func(id string) bool {
		return m.cancel(id)
	}, true)
	engine.RegisterMethod("coroutine.cancelAll", "取消所有协程", func() int {
		return m.cancelAll()
	}, true)
	engine.RegisterMethod("coroutine.async", "同步执行 Lua 回调并返回结果", func(task func() (interface{}, error)) interface{} {
		result, err := task()
		if err != nil {
			return err.Error()
		}
		return result
	}, true)
	engine.RegisterMethod("coroutine.await", "返回等待值", func(value interface{}) interface{} {
		return value
	}, true)
	engine.RegisterMethod("coroutine.getActiveCoroutines", "返回活跃协程数", func() int {
		return 0
	}, true)
	engine.RegisterMethod("coroutine.getCoroutineList", "返回协程列表", func() []interface{} {
		return []interface{}{}
	}, true)
	engine.RegisterMethod("coroutine.getCoroutineInfo", "返回协程信息", func(_ string) interface{} {
		return nil
	}, true)
	engine.RegisterMethod("coroutine.getStats", "返回协程统计", func() map[string]interface{} {
		return map[string]interface{}{"active": 0, "total": 0}
	}, true)
	engine.RegisterMethod("coroutine.createPool", "创建协程池", func(name string, maxWorkers int, maxTasks int) string {
		m.pools[name] = map[string]interface{}{"name": name, "maxWorkers": maxWorkers, "maxTasks": maxTasks}
		return name
	}, true)
	engine.RegisterMethod("coroutine.listPools", "返回协程池列表", func() []map[string]interface{} {
		pools := make([]map[string]interface{}, 0, len(m.pools))
		for _, pool := range m.pools {
			pools = append(pools, pool)
		}
		return pools
	}, true)
	engine.RegisterMethod("coroutine.getPoolStats", "返回协程池状态", func(name string) interface{} {
		return m.pools[name]
	}, true)
	engine.RegisterMethod("coroutine.closePool", "关闭协程池", func(name string) bool {
		if _, ok := m.pools[name]; !ok {
			return false
		}
		delete(m.pools, name)
		return true
	}, true)
	engine.RegisterMethod("coroutine.submitToPool", "提交任务到协程池", func(name string, task func() (interface{}, error), priority ...int) bool {
		if _, ok := m.pools[name]; !ok {
			return false
		}
		id, err := m.launch(task)
		if err != nil {
			return false
		}
		if len(priority) > 0 {
			m.priorities[id] = priority[0]
		}
		return true
	}, true)
	engine.RegisterMethod("coroutine.setScheduleStrategy", "设置调度策略", func(strategy string) {
		m.scheduleStrategy = strategy
	}, true)
	engine.RegisterMethod("coroutine.getScheduleStrategy", "获取调度策略", func() string {
		return m.scheduleStrategy
	}, true)
	engine.RegisterMethod("coroutine.setPriority", "设置任务优先级", func(name string, priority int) {
		m.priorities[name] = priority
	}, true)
	engine.RegisterMethod("coroutine.getPriority", "获取任务优先级", func(name string) int {
		return m.priorities[name]
	}, true)
	return nil
}

func GetModule() model.Module { return &CoroutineModule{} }

func (m *CoroutineModule) launch(task func() (interface{}, error)) (string, error) {
	if task == nil {
		return "", errNilCoroutineTask
	}
	id := fmt.Sprintf("coroutine_%d", atomic.AddInt64(&m.counter, 1))
	ctx, cancel := context.WithCancel(context.Background())
	m.remember(id, cancel)
	go func() {
		defer m.forget(id)
		select {
		case <-ctx.Done():
			return
		default:
			_, _ = task()
		}
	}()
	return id, nil
}

func (m *CoroutineModule) delay(milliseconds int, task func() (interface{}, error)) (string, error) {
	if task == nil {
		return "", errNilCoroutineTask
	}
	id := fmt.Sprintf("delay_%d", atomic.AddInt64(&m.counter, 1))
	ctx, cancel := context.WithCancel(context.Background())
	m.remember(id, cancel)
	go func() {
		defer m.forget(id)
		timer := time.NewTimer(time.Duration(milliseconds) * time.Millisecond)
		defer timer.Stop()
		select {
		case <-ctx.Done():
			return
		case <-timer.C:
			_, _ = task()
		}
	}()
	return id, nil
}

func (m *CoroutineModule) cancel(id string) bool {
	m.mu.Lock()
	cancel, ok := m.cancellations[id]
	if ok {
		delete(m.cancellations, id)
	}
	m.mu.Unlock()
	if ok {
		cancel()
	}
	return ok
}

func (m *CoroutineModule) cancelAll() int {
	m.mu.Lock()
	cancellations := make([]context.CancelFunc, 0, len(m.cancellations))
	for id, cancel := range m.cancellations {
		cancellations = append(cancellations, cancel)
		delete(m.cancellations, id)
	}
	m.mu.Unlock()
	for _, cancel := range cancellations {
		cancel()
	}
	return len(cancellations)
}

func (m *CoroutineModule) remember(id string, cancel context.CancelFunc) {
	m.mu.Lock()
	m.cancellations[id] = cancel
	m.mu.Unlock()
}

func (m *CoroutineModule) forget(id string) {
	m.mu.Lock()
	delete(m.cancellations, id)
	m.mu.Unlock()
}
