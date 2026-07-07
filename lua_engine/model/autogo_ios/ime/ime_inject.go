//go:build ignore
// +build ignore

package ime

import (
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"

	autogoime "github.com/Dasongzi1366/AutoGo/ime"
	lua "github.com/yuin/gopher-lua"
)

type ImeModule struct{}

func (m *ImeModule) Name() string      { return "ime" }
func (m *ImeModule) IsAvailable() bool { return true }

func (m *ImeModule) Register(engine model.Engine) error {
	state := engine.GetState()
	imeObj := state.NewTable()
	state.SetGlobal("ime", imeObj)
	imeObj.RawSetString("getClipText", state.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LString(autogoime.GetClipText()))
		return 1
	}))
	imeObj.RawSetString("setClipText", state.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LBool(autogoime.SetClipText(L.CheckString(1))))
		return 1
	}))
	imeObj.RawSetString("inputText", state.NewFunction(func(L *lua.LState) int {
		autogoime.InputText(L.CheckString(1))
		return 0
	}))
	engine.RegisterMethod("ime.getClipText", "获取剪切板内容", autogoime.GetClipText, true)
	engine.RegisterMethod("ime.setClipText", "设置剪切板内容", autogoime.SetClipText, true)
	engine.RegisterMethod("ime.inputText", "输入文本", func(text string) { autogoime.InputText(text) }, true)
	return nil
}
