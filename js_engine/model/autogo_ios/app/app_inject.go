package app

import (
	"github.com/ZingYao/autogo_scriptengine/js_engine/model"

	autogoapp "github.com/Dasongzi1366/AutoGo/app"
	"github.com/ZingYao/goja"
)

// AppModule iOS app 模块。
type AppModule struct{}

// Name 返回模块名称。
func (m *AppModule) Name() string {
	return "app"
}

// IsAvailable 检查模块是否可用。
func (m *AppModule) IsAvailable() bool {
	return true
}

// Register 向引擎注册 iOS app 方法。
func (m *AppModule) Register(engine model.Engine) error {
	vm := engine.GetVM()
	appObj := vm.NewObject()
	vm.Set("app", appObj)

	appObj.Set("currentPackage", func(call goja.FunctionCall) goja.Value { return vm.ToValue(autogoapp.CurrentPackage()) })
	appObj.Set("launch", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(autogoapp.Launch(call.Argument(0).String()))
	})
	appObj.Set("forceStop", func(call goja.FunctionCall) goja.Value {
		autogoapp.ForceStop(call.Argument(0).String())
		return goja.Undefined()
	})
	appObj.Set("getList", func(call goja.FunctionCall) goja.Value {
		includeSystemApps := true
		if len(call.Arguments) >= 1 {
			includeSystemApps = call.Argument(0).ToBoolean()
		}
		return vm.ToValue(autogoapp.GetList(includeSystemApps))
	})
	appObj.Set("selfPackage", func(call goja.FunctionCall) goja.Value { return vm.ToValue(autogoapp.SelfPackage()) })
	appObj.Set("getName", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(autogoapp.GetName(call.Argument(0).String()))
	})
	appObj.Set("getVersion", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(autogoapp.GetVersion(call.Argument(0).String()))
	})
	appObj.Set("getIcon", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(autogoapp.GetIcon(call.Argument(0).String()))
	})
	appObj.Set("isInstalled", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(autogoapp.IsInstalled(call.Argument(0).String()))
	})
	appObj.Set("uninstall", func(call goja.FunctionCall) goja.Value {
		autogoapp.Uninstall(call.Argument(0).String())
		return goja.Undefined()
	})
	appObj.Set("install", func(call goja.FunctionCall) goja.Value {
		autogoapp.Install(call.Argument(0).String())
		return goja.Undefined()
	})
	appObj.Set("clear", func(call goja.FunctionCall) goja.Value {
		autogoapp.Clear(call.Argument(0).String())
		return goja.Undefined()
	})
	appObj.Set("openUrl", func(call goja.FunctionCall) goja.Value {
		autogoapp.OpenUrl(call.Argument(0).String())
		return goja.Undefined()
	})

	engine.RegisterMethod("app.currentPackage", "获取当前前台应用包名", autogoapp.CurrentPackage, true)
	engine.RegisterMethod("app.launch", "通过应用包名启动应用", autogoapp.Launch, true)
	engine.RegisterMethod("app.forceStop", "关闭应用", func(packageName string) { autogoapp.ForceStop(packageName) }, true)
	engine.RegisterMethod("app.getList", "获取已安装应用列表", autogoapp.GetList, true)
	engine.RegisterMethod("app.selfPackage", "获取自身应用包名", autogoapp.SelfPackage, true)
	engine.RegisterMethod("app.getName", "通过包名获取应用显示名称", autogoapp.GetName, true)
	engine.RegisterMethod("app.getVersion", "获取应用版本号", autogoapp.GetVersion, true)
	engine.RegisterMethod("app.getIcon", "获取应用图标", autogoapp.GetIcon, true)
	engine.RegisterMethod("app.isInstalled", "判断是否已安装某个应用", autogoapp.IsInstalled, true)
	engine.RegisterMethod("app.uninstall", "卸载应用", func(packageName string) { autogoapp.Uninstall(packageName) }, true)
	engine.RegisterMethod("app.install", "安装应用", func(path string) { autogoapp.Install(path) }, true)
	engine.RegisterMethod("app.clear", "清除应用数据", func(packageName string) { autogoapp.Clear(packageName) }, true)
	engine.RegisterMethod("app.openUrl", "用系统默认方式打开 URL", func(url string) { autogoapp.OpenUrl(url) }, true)

	return nil
}
