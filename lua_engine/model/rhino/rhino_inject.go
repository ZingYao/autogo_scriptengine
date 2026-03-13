package rhino

import (
	"app/lua_engine/model"

	"github.com/Dasongzi1366/AutoGo/rhino"
	lua "github.com/yuin/gopher-lua"
)

// RhinoModule rhino 模块
type RhinoModule struct{}

// Name 返回模块名称
func (m *RhinoModule) Name() string {
	return "rhino"
}

// IsAvailable 检查模块是否可用
func (m *RhinoModule) IsAvailable() bool {
	return true
}

// Register 向引擎注册方法
func (m *RhinoModule) Register(engine model.Engine) error {
	state := engine.GetState()

	rhinoObj := state.NewTable()
	state.SetGlobal("rhino", rhinoObj)

	rhinoObj.RawSetString("eval", state.NewFunction(func(L *lua.LState) int {
		contextId := L.CheckString(1)
		script := L.CheckString(2)
		result := rhino.Eval(contextId, script)
		L.Push(lua.LString(result))
		return 1
	}))

	engine.RegisterMethod("rhino.eval", "执行指定的JavaScript脚本并返回结果", rhino.Eval, true)

	return nil
}
