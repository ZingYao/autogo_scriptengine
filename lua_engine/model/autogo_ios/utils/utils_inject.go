//go:build ignore
// +build ignore

package utils

import (
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"

	autogoutils "github.com/Dasongzi1366/AutoGo/utils"
	lua "github.com/yuin/gopher-lua"
)

type UtilsModule struct{}

func (m *UtilsModule) Name() string      { return "utils" }
func (m *UtilsModule) IsAvailable() bool { return true }

func (m *UtilsModule) Register(engine model.Engine) error {
	state := engine.GetState()
	utilsObj := state.NewTable()
	state.SetGlobal("utils", utilsObj)
	utilsObj.RawSetString("toast", state.NewFunction(func(L *lua.LState) int {
		autogoutils.Toast(L.CheckString(1), L.OptInt(2, -1), L.OptInt(3, -1), L.OptInt(4, -1))
		return 0
	}))
	utilsObj.RawSetString("alert", state.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LNumber(autogoutils.Alert(L.CheckString(1), L.CheckString(2), L.OptString(3, ""), L.OptString(4, ""))))
		return 1
	}))
	utilsObj.RawSetString("inputAlert", state.NewFunction(func(L *lua.LState) int {
		text, ok := autogoutils.InputAlert(L.CheckString(1), L.CheckString(2), L.OptString(3, ""), L.OptString(4, ""), L.OptString(5, ""), L.OptString(6, ""))
		L.Push(lua.LString(text))
		L.Push(lua.LBool(ok))
		return 2
	}))
	utilsObj.RawSetString("random", state.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LNumber(autogoutils.Random(L.CheckInt(1), L.CheckInt(2))))
		return 1
	}))
	utilsObj.RawSetString("sleep", state.NewFunction(func(L *lua.LState) int { autogoutils.Sleep(L.CheckInt(1)); return 0 }))
	utilsObj.RawSetString("i2s", state.NewFunction(func(L *lua.LState) int { L.Push(lua.LString(autogoutils.I2s(L.CheckInt(1)))); return 1 }))
	utilsObj.RawSetString("s2i", state.NewFunction(func(L *lua.LState) int { L.Push(lua.LNumber(autogoutils.S2i(L.CheckString(1)))); return 1 }))
	utilsObj.RawSetString("f2s", state.NewFunction(func(L *lua.LState) int { L.Push(lua.LString(autogoutils.F2s(float64(L.CheckNumber(1))))); return 1 }))
	utilsObj.RawSetString("s2f", state.NewFunction(func(L *lua.LState) int { L.Push(lua.LNumber(autogoutils.S2f(L.CheckString(1)))); return 1 }))
	utilsObj.RawSetString("b2s", state.NewFunction(func(L *lua.LState) int { L.Push(lua.LString(autogoutils.B2s(L.CheckBool(1)))); return 1 }))
	utilsObj.RawSetString("s2b", state.NewFunction(func(L *lua.LState) int { L.Push(lua.LBool(autogoutils.S2b(L.CheckString(1)))); return 1 }))
	engine.RegisterMethod("utils.toast", "显示 Toast 提示", autogoutils.Toast, true)
	engine.RegisterMethod("utils.alert", "显示 Alert 对话框", autogoutils.Alert, true)
	engine.RegisterMethod("utils.inputAlert", "显示输入对话框", autogoutils.InputAlert, true)
	engine.RegisterMethod("utils.random", "返回指定范围内的随机整数", autogoutils.Random, true)
	engine.RegisterMethod("utils.sleep", "暂停当前线程指定毫秒数", autogoutils.Sleep, true)
	engine.RegisterMethod("utils.i2s", "整数转字符串", autogoutils.I2s, true)
	engine.RegisterMethod("utils.s2i", "字符串转整数", autogoutils.S2i, true)
	engine.RegisterMethod("utils.f2s", "浮点数转字符串", autogoutils.F2s, true)
	engine.RegisterMethod("utils.s2f", "字符串转浮点数", autogoutils.S2f, true)
	engine.RegisterMethod("utils.b2s", "布尔值转字符串", autogoutils.B2s, true)
	engine.RegisterMethod("utils.s2b", "字符串转布尔值", autogoutils.S2b, true)
	return nil
}
