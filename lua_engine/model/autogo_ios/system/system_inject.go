package system

import (
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"

	autogosystem "github.com/Dasongzi1366/AutoGo/system"
	lua "github.com/yuin/gopher-lua"
)

type SystemModule struct{}

func (m *SystemModule) Name() string      { return "system" }
func (m *SystemModule) IsAvailable() bool { return true }

func (m *SystemModule) Register(engine model.Engine) error {
	state := engine.GetState()
	systemObj := state.NewTable()
	state.SetGlobal("system", systemObj)
	systemObj.RawSetString("getPid", state.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LNumber(autogosystem.GetPid(L.OptString(1, ""))))
		return 1
	}))
	systemObj.RawSetString("getMemoryUsage", state.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LNumber(autogosystem.GetMemoryUsage(L.OptInt(1, 0))))
		return 1
	}))
	systemObj.RawSetString("getCpuUsage", state.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LNumber(autogosystem.GetCpuUsage(L.OptInt(1, 0))))
		return 1
	}))
	systemObj.RawSetString("restartSelf", state.NewFunction(func(L *lua.LState) int {
		autogosystem.RestartSelf()
		return 0
	}))
	engine.RegisterMethod("system.getPid", "获取进程 ID", autogosystem.GetPid, true)
	engine.RegisterMethod("system.getMemoryUsage", "获取内存使用", autogosystem.GetMemoryUsage, true)
	engine.RegisterMethod("system.getCpuUsage", "获取 CPU 使用率", autogosystem.GetCpuUsage, true)
	engine.RegisterMethod("system.restartSelf", "重启自身", autogosystem.RestartSelf, true)
	return nil
}
