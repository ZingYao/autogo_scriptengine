package system

import (
	"app/js_engine/model"

	"github.com/Dasongzi1366/AutoGo/system"
	"github.com/dop251/goja"
)

// SystemModule system 模块
type SystemModule struct{}

// Name 返回模块名称
func (m *SystemModule) Name() string {
	return "system"
}

// IsAvailable 检查模块是否可用
func (m *SystemModule) IsAvailable() bool {
	return true
}

// Register 向引擎注册方法
func (m *SystemModule) Register(engine model.Engine) error {
	vm := engine.GetVM()

	systemObj := vm.NewObject()
	vm.Set("system", systemObj)

	systemObj.Set("getPid", func(call goja.FunctionCall) goja.Value {
		processName := call.Argument(0).String()
		result := system.GetPid(processName)
		return vm.ToValue(result)
	})

	systemObj.Set("getMemoryUsage", func(call goja.FunctionCall) goja.Value {
		pid := int(call.Argument(0).ToInteger())
		result := system.GetMemoryUsage(pid)
		return vm.ToValue(result)
	})

	systemObj.Set("getCpuUsage", func(call goja.FunctionCall) goja.Value {
		pid := int(call.Argument(0).ToInteger())
		result := system.GetCpuUsage(pid)
		return vm.ToValue(result)
	})

	systemObj.Set("restartSelf", func(call goja.FunctionCall) goja.Value {
		system.RestartSelf()
		return goja.Undefined()
	})

	systemObj.Set("setBootStart", func(call goja.FunctionCall) goja.Value {
		enable := call.Argument(0).ToBoolean()
		system.SetBootStart(enable)
		return goja.Undefined()
	})

	engine.RegisterMethod("system.getPid", "获取进程ID", func(processName string) int { return system.GetPid(processName) }, true)
	engine.RegisterMethod("system.getMemoryUsage", "获取内存使用", func(pid int) int { return system.GetMemoryUsage(pid) }, true)
	engine.RegisterMethod("system.getCpuUsage", "获取CPU使用率", func(pid int) float64 { return system.GetCpuUsage(pid) }, true)
	engine.RegisterMethod("system.restartSelf", "重启自身", system.RestartSelf, true)
	engine.RegisterMethod("system.setBootStart", "设置开机自启", func(enable bool) { system.SetBootStart(enable) }, true)

	return nil
}
