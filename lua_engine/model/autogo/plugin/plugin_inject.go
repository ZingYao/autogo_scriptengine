package plugin

import (
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"

	"github.com/Dasongzi1366/AutoGo/plugin"
	lua "github.com/yuin/gopher-lua"
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
	state := engine.GetState()

	pluginObj := state.NewTable()
	state.SetGlobal("plugin", pluginObj)

	pluginObj.RawSetString("loadApk", state.NewFunction(func(L *lua.LState) int {
		path := L.CheckString(1)
		result := plugin.LoadApk(path)
		if result != nil {
			ud := L.NewUserData()
			ud.Value = result
			L.Push(ud)
		} else {
			L.Push(lua.LNil)
		}
		return 1
	}))

	engine.RegisterMethod("plugin.loadApk", "加载外部APK", plugin.LoadApk, true)

	return nil
}
