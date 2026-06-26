package app

import (
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"

	autogoapp "github.com/Dasongzi1366/AutoGo/app"
	lua "github.com/yuin/gopher-lua"
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
	state := engine.GetState()
	appObj := state.NewTable()
	state.SetGlobal("app", appObj)

	appObj.RawSetString("currentPackage", state.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LString(autogoapp.CurrentPackage()))
		return 1
	}))

	appObj.RawSetString("launch", state.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LBool(autogoapp.Launch(L.CheckString(1))))
		return 1
	}))

	appObj.RawSetString("forceStop", state.NewFunction(func(L *lua.LState) int {
		autogoapp.ForceStop(L.CheckString(1))
		return 0
	}))

	appObj.RawSetString("getList", state.NewFunction(func(L *lua.LState) int {
		includeSystemApps := true
		if L.GetTop() >= 1 {
			includeSystemApps = L.CheckBool(1)
		}
		resultTable := L.NewTable()
		for index, appInfo := range autogoapp.GetList(includeSystemApps) {
			appTable := L.NewTable()
			appTable.RawSetString("packageName", lua.LString(appInfo.PackageName))
			appTable.RawSetString("appName", lua.LString(appInfo.AppName))
			appTable.RawSetString("versionName", lua.LString(appInfo.VersionName))
			appTable.RawSetString("versionCode", lua.LString(appInfo.VersionCode))
			appTable.RawSetString("isSystemApp", lua.LBool(appInfo.IsSystemApp))
			resultTable.RawSetInt(index+1, appTable)
		}
		L.Push(resultTable)
		return 1
	}))

	appObj.RawSetString("selfPackage", state.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LString(autogoapp.SelfPackage()))
		return 1
	}))

	appObj.RawSetString("getName", state.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LString(autogoapp.GetName(L.CheckString(1))))
		return 1
	}))

	appObj.RawSetString("getVersion", state.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LString(autogoapp.GetVersion(L.CheckString(1))))
		return 1
	}))

	appObj.RawSetString("getIcon", state.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LString(string(autogoapp.GetIcon(L.CheckString(1)))))
		return 1
	}))

	appObj.RawSetString("isInstalled", state.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LBool(autogoapp.IsInstalled(L.CheckString(1))))
		return 1
	}))

	appObj.RawSetString("uninstall", state.NewFunction(func(L *lua.LState) int {
		autogoapp.Uninstall(L.CheckString(1))
		return 0
	}))

	appObj.RawSetString("install", state.NewFunction(func(L *lua.LState) int {
		autogoapp.Install(L.CheckString(1))
		return 0
	}))

	appObj.RawSetString("clear", state.NewFunction(func(L *lua.LState) int {
		autogoapp.Clear(L.CheckString(1))
		return 0
	}))

	appObj.RawSetString("openUrl", state.NewFunction(func(L *lua.LState) int {
		autogoapp.OpenUrl(L.CheckString(1))
		return 0
	}))

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
