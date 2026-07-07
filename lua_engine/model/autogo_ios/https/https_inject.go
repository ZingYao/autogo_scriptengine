//go:build ignore
// +build ignore

package https

import (
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"

	autogohttps "github.com/Dasongzi1366/AutoGo/https"
	lua "github.com/yuin/gopher-lua"
)

type HttpsModule struct{}

func (m *HttpsModule) Name() string      { return "https" }
func (m *HttpsModule) IsAvailable() bool { return true }

func pushResponse(L *lua.LState, code int, data []byte) {
	result := L.NewTable()
	result.RawSetString("code", lua.LNumber(code))
	if data == nil {
		result.RawSetString("data", lua.LNil)
	} else {
		result.RawSetString("data", lua.LString(string(data)))
	}
	L.Push(result)
}

func (m *HttpsModule) Register(engine model.Engine) error {
	state := engine.GetState()
	httpsObj := state.NewTable()
	state.SetGlobal("https", httpsObj)
	httpsObj.RawSetString("get", state.NewFunction(func(L *lua.LState) int {
		timeout := 5000
		if L.GetTop() >= 2 {
			timeout = L.CheckInt(2)
		}
		code, data := autogohttps.Get(L.CheckString(1), timeout)
		pushResponse(L, code, data)
		return 1
	}))
	httpsObj.RawSetString("post", state.NewFunction(func(L *lua.LState) int {
		headers := map[string]string{}
		if L.GetTop() >= 3 && L.CheckAny(3).Type() == lua.LTTable {
			L.CheckTable(3).ForEach(func(key lua.LValue, value lua.LValue) {
				headers[key.String()] = value.String()
			})
		}
		timeout := 5000
		if L.GetTop() >= 4 {
			timeout = L.CheckInt(4)
		}
		code, data := autogohttps.Post(L.CheckString(1), []byte(L.OptString(2, "")), headers, timeout)
		pushResponse(L, code, data)
		return 1
	}))
	httpsObj.RawSetString("postMultipart", state.NewFunction(func(L *lua.LState) int {
		timeout := 5000
		if L.GetTop() >= 4 {
			timeout = L.CheckInt(4)
		}
		code, data := autogohttps.PostMultipart(L.CheckString(1), L.CheckString(2), []byte(L.OptString(3, "")), timeout)
		pushResponse(L, code, data)
		return 1
	}))
	engine.RegisterMethod("https.get", "发送 GET 请求", autogohttps.Get, true)
	engine.RegisterMethod("https.post", "发送 POST 请求", autogohttps.Post, true)
	engine.RegisterMethod("https.postMultipart", "发送 multipart POST 请求", autogohttps.PostMultipart, true)
	return nil
}
