package ime

import (
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"

	"github.com/Dasongzi1366/AutoGo/ime"
	lua "github.com/yuin/gopher-lua"
)

// ImeModule ime 模块
type ImeModule struct{}

// Name 返回模块名称
func (m *ImeModule) Name() string {
	return "ime"
}

// IsAvailable 检查模块是否可用
func (m *ImeModule) IsAvailable() bool {
	return true
}

// Register 向引擎注册方法
func (m *ImeModule) Register(engine model.Engine) error {
	state := engine.GetState()

	imeObj := state.NewTable()
	state.SetGlobal("ime", imeObj)

	imeObj.RawSetString("getClipText", state.NewFunction(func(L *lua.LState) int {
		result := ime.GetClipText()
		L.Push(lua.LString(result))
		return 1
	}))

	imeObj.RawSetString("setClipText", state.NewFunction(func(L *lua.LState) int {
		text := L.CheckString(1)
		ime.SetClipText(text)
		return 0
	}))

	imeObj.RawSetString("keyAction", state.NewFunction(func(L *lua.LState) int {
		code := L.CheckInt(1)
		ime.KeyAction(code)
		return 0
	}))

	imeObj.RawSetString("inputText", state.NewFunction(func(L *lua.LState) int {
		text := L.CheckString(1)
		displayId := 0
		if L.GetTop() > 1 {
			displayId = L.CheckInt(2)
		}
		ime.InputText(text, displayId)
		return 0
	}))

	imeObj.RawSetString("getIMEList", state.NewFunction(func(L *lua.LState) int {
		result := ime.GetIMEList()
		resultTable := L.NewTable()
		for i, imeName := range result {
			resultTable.RawSetInt(i+1, lua.LString(imeName))
		}
		L.Push(resultTable)
		return 1
	}))

	imeObj.RawSetString("setCurrentIME", state.NewFunction(func(L *lua.LState) int {
		imeName := L.CheckString(1)
		ime.SetCurrentIME(imeName)
		return 0
	}))

	engine.RegisterMethod("ime.getClipText", "获取剪切板内容", ime.GetClipText, true)
	engine.RegisterMethod("ime.setClipText", "设置剪切板内容", ime.SetClipText, true)
	engine.RegisterMethod("ime.keyAction", "模拟按键", ime.KeyAction, true)
	engine.RegisterMethod("ime.inputText", "输入文本", ime.InputText, true)
	engine.RegisterMethod("ime.getIMEList", "获取输入法列表", ime.GetIMEList, true)
	engine.RegisterMethod("ime.setCurrentIME", "设置当前输入法", ime.SetCurrentIME, true)

	return nil
}
