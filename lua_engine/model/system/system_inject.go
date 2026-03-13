package system

import (
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"

	"github.com/Dasongzi1366/AutoGo/system"
	lua "github.com/yuin/gopher-lua"
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
	state := engine.GetState()

	systemObj := state.NewTable()
	state.SetGlobal("system", systemObj)

	systemObj.RawSetString("getPid", state.NewFunction(func(L *lua.LState) int {
		processName := L.CheckString(1)
		result := system.GetPid(processName)
		L.Push(lua.LNumber(result))
		return 1
	}))

	systemObj.RawSetString("getMemoryUsage", state.NewFunction(func(L *lua.LState) int {
		pid := L.CheckInt(1)
		result := system.GetMemoryUsage(pid)
		L.Push(lua.LNumber(result))
		return 1
	}))

	systemObj.RawSetString("getCpuUsage", state.NewFunction(func(L *lua.LState) int {
		pid := L.CheckInt(1)
		result := system.GetCpuUsage(pid)
		L.Push(lua.LNumber(result))
		return 1
	}))

	systemObj.RawSetString("restartSelf", state.NewFunction(func(L *lua.LState) int {
		system.RestartSelf()
		return 0
	}))

	systemObj.RawSetString("setBootStart", state.NewFunction(func(L *lua.LState) int {
		enable := L.CheckBool(1)
		system.SetBootStart(enable)
		return 0
	}))

	engine.RegisterMethod("system.getPid", "获取进程ID", func(processName string) int { return system.GetPid(processName) }, true)
	engine.RegisterMethod("system.getMemoryUsage", "获取内存使用", func(pid int) int { return system.GetMemoryUsage(pid) }, true)
	engine.RegisterMethod("system.getCpuUsage", "获取CPU使用率", func(pid int) float64 { return system.GetCpuUsage(pid) }, true)
	engine.RegisterMethod("system.restartSelf", "重启自身", system.RestartSelf, true)
	engine.RegisterMethod("system.setBootStart", "设置开机自启", func(enable bool) { system.SetBootStart(enable) }, true)

	return nil
}
