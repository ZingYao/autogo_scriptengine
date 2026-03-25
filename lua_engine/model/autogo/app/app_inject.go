package app

import (
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"

	"github.com/Dasongzi1366/AutoGo/app"
	lua "github.com/yuin/gopher-lua"
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
	state := engine.GetState()

	appObj := state.NewTable()
	state.SetGlobal("app", appObj)

	appObj.RawSetString("currentPackage", state.NewFunction(func(L *lua.LState) int {
		result := app.CurrentPackage()
		L.Push(lua.LString(result))
		return 1
	}))

	appObj.RawSetString("currentActivity", state.NewFunction(func(L *lua.LState) int {
		result := app.CurrentActivity()
		L.Push(lua.LString(result))
		return 1
	}))

	appObj.RawSetString("launch", state.NewFunction(func(L *lua.LState) int {
		packageName := L.CheckString(1)
		displayId := 0
		if L.GetTop() >= 2 {
			displayId = L.CheckInt(2)
		}
		result := app.Launch(packageName, displayId)
		L.Push(lua.LBool(result))
		return 1
	}))

	appObj.RawSetString("getList", state.NewFunction(func(L *lua.LState) int {
		includeSystemApps := true
		if L.GetTop() >= 1 {
			includeSystemApps = L.CheckBool(1)
		}
		result := app.GetList(includeSystemApps)
		resultTable := L.NewTable()
		for i, appInfo := range result {
			appTable := L.NewTable()
			appTable.RawSetString("packageName", lua.LString(appInfo.PackageName))
			appTable.RawSetString("appName", lua.LString(appInfo.AppName))
			appTable.RawSetString("versionName", lua.LString(appInfo.VersionName))
			appTable.RawSetString("versionCode", lua.LString(appInfo.VersionCode))
			appTable.RawSetString("isSystemApp", lua.LBool(appInfo.IsSystemApp))
			appTable.RawSetString("enabled", lua.LBool(appInfo.Enabled))
			resultTable.RawSetInt(i+1, appTable)
		}
		L.Push(resultTable)
		return 1
	}))

	appObj.RawSetString("openAppSetting", state.NewFunction(func(L *lua.LState) int {
		packageName := L.CheckString(1)
		result := app.OpenSetting(packageName)
		L.Push(lua.LBool(result))
		return 1
	}))

	appObj.RawSetString("viewFile", state.NewFunction(func(L *lua.LState) int {
		path := L.CheckString(1)
		app.ViewFile(path)
		return 0
	}))

	appObj.RawSetString("editFile", state.NewFunction(func(L *lua.LState) int {
		path := L.CheckString(1)
		app.EditFile(path)
		return 0
	}))

	appObj.RawSetString("uninstall", state.NewFunction(func(L *lua.LState) int {
		packageName := L.CheckString(1)
		app.Uninstall(packageName)
		return 0
	}))

	appObj.RawSetString("install", state.NewFunction(func(L *lua.LState) int {
		path := L.CheckString(1)
		app.Install(path)
		return 0
	}))

	appObj.RawSetString("isInstalled", state.NewFunction(func(L *lua.LState) int {
		packageName := L.CheckString(1)
		result := app.IsInstalled(packageName)
		L.Push(lua.LBool(result))
		return 1
	}))

	appObj.RawSetString("clear", state.NewFunction(func(L *lua.LState) int {
		packageName := L.CheckString(1)
		app.Clear(packageName)
		return 0
	}))

	appObj.RawSetString("forceStop", state.NewFunction(func(L *lua.LState) int {
		packageName := L.CheckString(1)
		app.ForceStop(packageName)
		return 0
	}))

	appObj.RawSetString("disable", state.NewFunction(func(L *lua.LState) int {
		packageName := L.CheckString(1)
		app.Disable(packageName)
		return 0
	}))

	appObj.RawSetString("ignoreBattOpt", state.NewFunction(func(L *lua.LState) int {
		packageName := L.CheckString(1)
		app.IgnoreBattOpt(packageName)
		return 0
	}))

	appObj.RawSetString("enableAccessibility", state.NewFunction(func(L *lua.LState) int {
		packageName := L.CheckString(1)
		app.EnableAccessibility(packageName)
		return 0
	}))

	appObj.RawSetString("disableAccessibility", state.NewFunction(func(L *lua.LState) int {
		packageName := L.CheckString(1)
		app.DisableAccessibility(packageName)
		return 0
	}))

	appObj.RawSetString("enable", state.NewFunction(func(L *lua.LState) int {
		packageName := L.CheckString(1)
		app.Enable(packageName)
		return 0
	}))

	appObj.RawSetString("getBrowserPackage", state.NewFunction(func(L *lua.LState) int {
		result := app.GetBrowserPackage()
		L.Push(lua.LString(result))
		return 1
	}))

	appObj.RawSetString("openUrl", state.NewFunction(func(L *lua.LState) int {
		url := L.CheckString(1)
		app.OpenUrl(url)
		return 0
	}))

	appObj.RawSetString("startActivity", state.NewFunction(func(L *lua.LState) int {
		optsTable := L.CheckTable(1)
		intentOpts := app.IntentOptions{}
		if action := optsTable.RawGetString("action"); action.Type() == lua.LTString {
			intentOpts.Action = string(action.(lua.LString))
		}
		if type_ := optsTable.RawGetString("type"); type_.Type() == lua.LTString {
			intentOpts.Type = string(type_.(lua.LString))
		}
		if data := optsTable.RawGetString("data"); data.Type() == lua.LTString {
			intentOpts.Data = string(data.(lua.LString))
		}
		if packageName := optsTable.RawGetString("packageName"); packageName.Type() == lua.LTString {
			intentOpts.PackageName = string(packageName.(lua.LString))
		}
		app.StartActivity(intentOpts)
		return 0
	}))

	appObj.RawSetString("sendBroadcast", state.NewFunction(func(L *lua.LState) int {
		optsTable := L.CheckTable(1)
		intentOpts := app.IntentOptions{}
		if action := optsTable.RawGetString("action"); action.Type() == lua.LTString {
			intentOpts.Action = string(action.(lua.LString))
		}
		if type_ := optsTable.RawGetString("type"); type_.Type() == lua.LTString {
			intentOpts.Type = string(type_.(lua.LString))
		}
		if data := optsTable.RawGetString("data"); data.Type() == lua.LTString {
			intentOpts.Data = string(data.(lua.LString))
		}
		if packageName := optsTable.RawGetString("packageName"); packageName.Type() == lua.LTString {
			intentOpts.PackageName = string(packageName.(lua.LString))
		}
		app.SendBroadcast(intentOpts)
		return 0
	}))

	appObj.RawSetString("startService", state.NewFunction(func(L *lua.LState) int {
		optsTable := L.CheckTable(1)
		intentOpts := app.IntentOptions{}
		if action := optsTable.RawGetString("action"); action.Type() == lua.LTString {
			intentOpts.Action = string(action.(lua.LString))
		}
		if type_ := optsTable.RawGetString("type"); type_.Type() == lua.LTString {
			intentOpts.Type = string(type_.(lua.LString))
		}
		if data := optsTable.RawGetString("data"); data.Type() == lua.LTString {
			intentOpts.Data = string(data.(lua.LString))
		}
		if packageName := optsTable.RawGetString("packageName"); packageName.Type() == lua.LTString {
			intentOpts.PackageName = string(packageName.(lua.LString))
		}
		app.StartService(intentOpts)
		return 0
	}))

	appObj.RawSetString("getName", state.NewFunction(func(L *lua.LState) int {
		packageName := L.CheckString(1)
		result := app.GetName(packageName)
		L.Push(lua.LString(result))
		return 1
	}))

	appObj.RawSetString("getVersion", state.NewFunction(func(L *lua.LState) int {
		packageName := L.CheckString(1)
		result := app.GetVersion(packageName)
		L.Push(lua.LString(result))
		return 1
	}))

	appObj.RawSetString("getIcon", state.NewFunction(func(L *lua.LState) int {
		packageName := L.CheckString(1)
		result := app.GetIcon(packageName)
		L.Push(lua.LString(string(result)))
		return 1
	}))

	engine.RegisterMethod("app.currentPackage", "获取当前页面应用包名", app.CurrentPackage, true)
	engine.RegisterMethod("app.currentActivity", "获取当前页面应用类名", app.CurrentActivity, true)
	engine.RegisterMethod("app.launch", "通过应用包名启动应用", func(packageName string, displayId int) bool {
		return app.Launch(packageName, displayId)
	}, true)
	engine.RegisterMethod("app.getList", "获取应用列表", func(includeSystemApps bool) []app.AppInfo {
		return app.GetList(includeSystemApps)
	}, true)
	engine.RegisterMethod("app.getName", "获取应用名称", app.GetName, true)
	engine.RegisterMethod("app.getIcon", "获取应用图标", app.GetIcon, true)
	engine.RegisterMethod("app.getVersion", "获取应用版本", app.GetVersion, true)
	engine.RegisterMethod("app.startActivity", "启动Activity", func(options app.IntentOptions) {
		app.StartActivity(options)
	}, true)
	engine.RegisterMethod("app.sendBroadcast", "发送广播", func(options app.IntentOptions) {
		app.SendBroadcast(options)
	}, true)
	engine.RegisterMethod("app.startService", "启动服务", func(options app.IntentOptions) {
		app.StartService(options)
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
	engine.RegisterMethod("app.getName", "获取应用名称", app.GetName, true)
	engine.RegisterMethod("app.getVersion", "获取应用版本", app.GetVersion, true)

	return nil
}

// GetModule 获取模块实例
func GetModule() model.Module {
	return &AppModule{}
}
