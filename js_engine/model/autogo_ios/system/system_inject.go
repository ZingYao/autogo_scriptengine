package system

import (
	"github.com/ZingYao/autogo_scriptengine/js_engine/model"

	autogosystem "github.com/Dasongzi1366/AutoGo/system"
	"github.com/dop251/goja"
)

// SystemModule iOS system 模块。
type SystemModule struct{}

// Name 返回模块名称。
func (m *SystemModule) Name() string {
	return "system"
}

// IsAvailable 检查模块是否可用。
func (m *SystemModule) IsAvailable() bool {
	return true
}

// Register 向引擎注册 iOS system 方法。
func (m *SystemModule) Register(engine model.Engine) error {
	vm := engine.GetVM()
	systemObj := vm.NewObject()
	vm.Set("system", systemObj)

	systemObj.Set("getPid", func(call goja.FunctionCall) goja.Value {
		processName := ""
		if len(call.Arguments) >= 1 {
			processName = call.Argument(0).String()
		}
		return vm.ToValue(autogosystem.GetPid(processName))
	})
	systemObj.Set("getMemoryUsage", func(call goja.FunctionCall) goja.Value {
		pid := 0
		if len(call.Arguments) >= 1 {
			pid = int(call.Argument(0).ToInteger())
		}
		return vm.ToValue(autogosystem.GetMemoryUsage(pid))
	})
	systemObj.Set("getCpuUsage", func(call goja.FunctionCall) goja.Value {
		pid := 0
		if len(call.Arguments) >= 1 {
			pid = int(call.Argument(0).ToInteger())
		}
		return vm.ToValue(autogosystem.GetCpuUsage(pid))
	})
	systemObj.Set("restartSelf", func(call goja.FunctionCall) goja.Value {
		autogosystem.RestartSelf()
		return goja.Undefined()
	})

	engine.RegisterMethod("system.getPid", "获取当前进程 PID", autogosystem.GetPid, true)
	engine.RegisterMethod("system.getMemoryUsage", "获取当前进程内存占用", autogosystem.GetMemoryUsage, true)
	engine.RegisterMethod("system.getCpuUsage", "获取当前进程 CPU 占用", autogosystem.GetCpuUsage, true)
	engine.RegisterMethod("system.restartSelf", "重启当前进程", autogosystem.RestartSelf, true)
	return nil
}
