package plugin

import (
	"app/js_engine/model"

	"github.com/Dasongzi1366/AutoGo/plugin"
	"github.com/dop251/goja"
)

// PluginModule plugin 模块
type PluginModule struct{}

// Name 返回模块名称
func (m *PluginModule) Name() string {
	return "plugin"
}

// IsAvailable 检查模块是否可用
func (m *PluginModule) IsAvailable() bool {
	return true
}

// Register 向引擎注册方法
func (m *PluginModule) Register(engine model.Engine) error {
	vm := engine.GetVM()

	pluginObj := vm.NewObject()
	vm.Set("plugin", pluginObj)

	pluginObj.Set("loadApk", func(call goja.FunctionCall) goja.Value {
		apkPath := call.Argument(0).String()
		cl := plugin.LoadApk(apkPath)
		return vm.ToValue(cl)
	})

	engine.RegisterMethod("plugin.loadApk", "加载外部APK", plugin.LoadApk, true)

	return nil
}
