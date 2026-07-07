package app

import (
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"

	autogoapp "github.com/Dasongzi1366/AutoGo/app"
)

// AppModule 是 go-lua-vm 迁移后的模块壳。
type AppModule struct{}

func New() *AppModule { return &AppModule{} }

func (m *AppModule) Name() string { return "app" }

func (m *AppModule) IsAvailable() bool { return true }

func (m *AppModule) Register(engine model.Engine) error {
	engine.RegisterMethod("app.currentPackage", "获取当前页面应用包名", autogoapp.CurrentPackage, true)
	engine.RegisterMethod("app.currentActivity", "获取当前页面应用类名", autogoapp.CurrentActivity, true)
	engine.RegisterMethod("app.launch", "通过应用包名启动应用", func(packageName string, displayID ...int) bool {
		targetDisplayID := 0
		if len(displayID) > 0 {
			targetDisplayID = displayID[0]
		}
		return autogoapp.Launch(packageName, targetDisplayID)
	}, true)
	engine.RegisterMethod("app.getList", "获取应用列表", func(includeSystemApps ...bool) []autogoapp.AppInfo {
		include := true
		if len(includeSystemApps) > 0 {
			include = includeSystemApps[0]
		}
		return autogoapp.GetList(include)
	}, true)
	engine.RegisterMethod("app.getName", "获取应用名称", autogoapp.GetName, true)
	engine.RegisterMethod("app.getIcon", "获取应用图标", autogoapp.GetIcon, true)
	engine.RegisterMethod("app.getVersion", "获取应用版本", autogoapp.GetVersion, true)
	engine.RegisterMethod("app.openSetting", "打开应用详情页", autogoapp.OpenSetting, true)
	engine.RegisterMethod("app.viewFile", "用其他应用查看文件", autogoapp.ViewFile, true)
	engine.RegisterMethod("app.editFile", "用其他应用编辑文件", autogoapp.EditFile, true)
	engine.RegisterMethod("app.uninstall", "卸载应用", autogoapp.Uninstall, true)
	engine.RegisterMethod("app.install", "安装应用", autogoapp.Install, true)
	engine.RegisterMethod("app.isInstalled", "判断是否已经安装某个应用", autogoapp.IsInstalled, true)
	engine.RegisterMethod("app.clear", "清除应用数据", autogoapp.Clear, true)
	engine.RegisterMethod("app.forceStop", "强制停止应用", autogoapp.ForceStop, true)
	engine.RegisterMethod("app.disable", "禁用应用", autogoapp.Disable, true)
	engine.RegisterMethod("app.enable", "启用应用", autogoapp.Enable, true)
	engine.RegisterMethod("app.enableAccessibility", "启用无障碍服务", autogoapp.EnableAccessibility, true)
	engine.RegisterMethod("app.disableAccessibility", "禁用无障碍服务", autogoapp.DisableAccessibility, true)
	engine.RegisterMethod("app.ignoreBattOpt", "忽略电池优化", autogoapp.IgnoreBattOpt, true)
	engine.RegisterMethod("app.getBrowserPackage", "获取系统默认浏览器包名", autogoapp.GetBrowserPackage, true)
	engine.RegisterMethod("app.openUrl", "用浏览器打开网站 URL", autogoapp.OpenUrl, true)
	engine.RegisterMethod("app.startActivity", "启动 Activity", autogoapp.StartActivity, true)
	engine.RegisterMethod("app.sendBroadcast", "发送广播", autogoapp.SendBroadcast, true)
	engine.RegisterMethod("app.startService", "启动服务", autogoapp.StartService, true)
	return nil
}

func GetModule() model.Module { return &AppModule{} }
