package app

import (
	"github.com/ZingYao/autogo_scriptengine/js_engine/model"

	"github.com/Dasongzi1366/AutoGo/app"
	"github.com/dop251/goja"
)

// AppModule app 模块
type AppModule struct{}

// Name 返回模块名称
func (m *AppModule) Name() string {
	return "app"
}

// IsAvailable 检查模块是否可用
func (m *AppModule) IsAvailable() bool {
	return true
}

// Register 向引擎注册方法
func (m *AppModule) Register(engine model.Engine) error {
	vm := engine.GetVM()

	appObj := vm.NewObject()
	vm.Set("app", appObj)

	appObj.Set("currentPackage", func(call goja.FunctionCall) goja.Value {
		result := app.CurrentPackage()
		return vm.ToValue(result)
	})

	appObj.Set("currentActivity", func(call goja.FunctionCall) goja.Value {
		result := app.CurrentActivity()
		return vm.ToValue(result)
	})

	appObj.Set("launch", func(call goja.FunctionCall) goja.Value {
		packageName := call.Argument(0).String()
		displayId := int(call.Argument(1).ToInteger())
		result := app.Launch(packageName, displayId)
		return vm.ToValue(result)
	})

	appObj.Set("getList", func(call goja.FunctionCall) goja.Value {
		includeSystemApps := true
		if len(call.Arguments) >= 1 {
			includeSystemApps = call.Argument(0).ToBoolean()
		}
		result := app.GetList(includeSystemApps)
		return vm.ToValue(result)
	})

	appObj.Set("getName", func(call goja.FunctionCall) goja.Value {
		packageName := call.Argument(0).String()
		result := app.GetName(packageName)
		return vm.ToValue(result)
	})

	appObj.Set("getIcon", func(call goja.FunctionCall) goja.Value {
		packageName := call.Argument(0).String()
		result := app.GetIcon(packageName)
		return vm.ToValue(result)
	})

	appObj.Set("getVersion", func(call goja.FunctionCall) goja.Value {
		packageName := call.Argument(0).String()
		result := app.GetVersion(packageName)
		return vm.ToValue(result)
	})

	appObj.Set("openAppSetting", func(call goja.FunctionCall) goja.Value {
		packageName := call.Argument(0).String()
		result := app.OpenSetting(packageName)
		return vm.ToValue(result)
	})

	appObj.Set("viewFile", func(call goja.FunctionCall) goja.Value {
		path := call.Argument(0).String()
		app.ViewFile(path)
		return goja.Undefined()
	})

	appObj.Set("editFile", func(call goja.FunctionCall) goja.Value {
		path := call.Argument(0).String()
		app.EditFile(path)
		return goja.Undefined()
	})

	appObj.Set("uninstall", func(call goja.FunctionCall) goja.Value {
		packageName := call.Argument(0).String()
		app.Uninstall(packageName)
		return goja.Undefined()
	})

	appObj.Set("install", func(call goja.FunctionCall) goja.Value {
		path := call.Argument(0).String()
		app.Install(path)
		return goja.Undefined()
	})

	appObj.Set("isInstalled", func(call goja.FunctionCall) goja.Value {
		packageName := call.Argument(0).String()
		result := app.IsInstalled(packageName)
		return vm.ToValue(result)
	})

	appObj.Set("clear", func(call goja.FunctionCall) goja.Value {
		packageName := call.Argument(0).String()
		app.Clear(packageName)
		return goja.Undefined()
	})

	appObj.Set("forceStop", func(call goja.FunctionCall) goja.Value {
		packageName := call.Argument(0).String()
		app.ForceStop(packageName)
		return goja.Undefined()
	})

	appObj.Set("disable", func(call goja.FunctionCall) goja.Value {
		packageName := call.Argument(0).String()
		app.Disable(packageName)
		return goja.Undefined()
	})

	appObj.Set("ignoreBattOpt", func(call goja.FunctionCall) goja.Value {
		packageName := call.Argument(0).String()
		app.IgnoreBattOpt(packageName)
		return goja.Undefined()
	})

	appObj.Set("enableAccessibility", func(call goja.FunctionCall) goja.Value {
		packageName := call.Argument(0).String()
		app.EnableAccessibility(packageName)
		return goja.Undefined()
	})

	appObj.Set("disableAccessibility", func(call goja.FunctionCall) goja.Value {
		packageName := call.Argument(0).String()
		app.DisableAccessibility(packageName)
		return goja.Undefined()
	})

	appObj.Set("enable", func(call goja.FunctionCall) goja.Value {
		packageName := call.Argument(0).String()
		app.Enable(packageName)
		return goja.Undefined()
	})

	appObj.Set("getBrowserPackage", func(call goja.FunctionCall) goja.Value {
		result := app.GetBrowserPackage()
		return vm.ToValue(result)
	})

	appObj.Set("openUrl", func(call goja.FunctionCall) goja.Value {
		url := call.Argument(0).String()
		app.OpenUrl(url)
		return goja.Undefined()
	})

	appObj.Set("startActivity", func(call goja.FunctionCall) goja.Value {
		options := call.Argument(0).Export()
		if opts, ok := options.(map[string]interface{}); ok {
			intentOpts := app.IntentOptions{}
			if action, ok := opts["action"].(string); ok {
				intentOpts.Action = action
			}
			if type_, ok := opts["type"].(string); ok {
				intentOpts.Type = type_
			}
			if data, ok := opts["data"].(string); ok {
				intentOpts.Data = data
			}
			if packageName, ok := opts["packageName"].(string); ok {
				intentOpts.PackageName = packageName
			}
			app.StartActivity(intentOpts)
		}
		return goja.Undefined()
	})

	appObj.Set("sendBroadcast", func(call goja.FunctionCall) goja.Value {
		options := call.Argument(0).Export()
		if opts, ok := options.(map[string]interface{}); ok {
			intentOpts := app.IntentOptions{}
			if action, ok := opts["action"].(string); ok {
				intentOpts.Action = action
			}
			if type_, ok := opts["type"].(string); ok {
				intentOpts.Type = type_
			}
			if data, ok := opts["data"].(string); ok {
				intentOpts.Data = data
			}
			if packageName, ok := opts["packageName"].(string); ok {
				intentOpts.PackageName = packageName
			}
			app.SendBroadcast(intentOpts)
		}
		return goja.Undefined()
	})

	appObj.Set("startService", func(call goja.FunctionCall) goja.Value {
		options := call.Argument(0).Export()
		if opts, ok := options.(map[string]interface{}); ok {
			intentOpts := app.IntentOptions{}
			if action, ok := opts["action"].(string); ok {
				intentOpts.Action = action
			}
			if type_, ok := opts["type"].(string); ok {
				intentOpts.Type = type_
			}
			if data, ok := opts["data"].(string); ok {
				intentOpts.Data = data
			}
			if packageName, ok := opts["packageName"].(string); ok {
				intentOpts.PackageName = packageName
			}
			app.StartService(intentOpts)
		}
		return goja.Undefined()
	})

	engine.RegisterMethod("app.currentPackage", "获取当前页面应用包名", app.CurrentPackage, true)
	engine.RegisterMethod("app.currentActivity", "获取当前页面应用类名", app.CurrentActivity, true)
	engine.RegisterMethod("app.launch", "通过应用包名启动应用", func(packageName string, displayId int) bool {
		return app.Launch(packageName, displayId)
	}, true)
	engine.RegisterMethod("app.getList", "获取应用列表", func(includeSystemApps bool) []app.AppInfo {
		return app.GetList(includeSystemApps)
	}, true)
	engine.RegisterMethod("app.getName", "获取应用名称", func(packageName string) string {
		return app.GetName(packageName)
	}, true)
	engine.RegisterMethod("app.getIcon", "获取应用图标", func(packageName string) []byte {
		return app.GetIcon(packageName)
	}, true)
	engine.RegisterMethod("app.getVersion", "获取应用版本", func(packageName string) string {
		return app.GetVersion(packageName)
	}, true)
	engine.RegisterMethod("app.openAppSetting", "打开应用的详情页(设置页)", func(packageName string) bool {
		return app.OpenSetting(packageName)
	}, true)
	engine.RegisterMethod("app.viewFile", "用其他应用查看文件", func(path string) {
		app.ViewFile(path)
	}, true)
	engine.RegisterMethod("app.editFile", "用其他应用编辑文件", func(path string) {
		app.EditFile(path)
	}, true)
	engine.RegisterMethod("app.uninstall", "卸载应用", func(packageName string) {
		app.Uninstall(packageName)
	}, true)
	engine.RegisterMethod("app.install", "安装应用", func(path string) {
		app.Install(path)
	}, true)
	engine.RegisterMethod("app.isInstalled", "判断是否已经安装某个应用", func(packageName string) bool {
		return app.IsInstalled(packageName)
	}, true)
	engine.RegisterMethod("app.clear", "清除应用数据", func(packageName string) {
		app.Clear(packageName)
	}, true)
	engine.RegisterMethod("app.forceStop", "强制停止应用", func(packageName string) {
		app.ForceStop(packageName)
	}, true)
	engine.RegisterMethod("app.disable", "禁用应用", func(packageName string) {
		app.Disable(packageName)
	}, true)
	engine.RegisterMethod("app.enableAccessibility", "启用无障碍服务", func(packageName string) {
		app.EnableAccessibility(packageName)
	}, true)
	engine.RegisterMethod("app.disableAccessibility", "禁用无障碍服务", func(packageName string) {
		app.DisableAccessibility(packageName)
	}, true)
	engine.RegisterMethod("app.ignoreBattOpt", "忽略电池优化", func(packageName string) {
		app.IgnoreBattOpt(packageName)
	}, true)
	engine.RegisterMethod("app.enable", "启用应用", func(packageName string) {
		app.Enable(packageName)
	}, true)
	engine.RegisterMethod("app.getBrowserPackage", "获取系统默认浏览器包名", app.GetBrowserPackage, true)
	engine.RegisterMethod("app.openUrl", "用浏览器打开网站url", func(url string) {
		app.OpenUrl(url)
	}, true)

	return nil
}
