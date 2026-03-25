package http

import (
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"

	"github.com/Dasongzi1366/AutoGo/https"
	lua "github.com/yuin/gopher-lua"
)

// HttpModule http 模块
type HttpModule struct{}

// Name 返回模块名称
func (m *HttpModule) Name() string {
	return "http"
}

// IsAvailable 检查模块是否可用
func (m *HttpModule) IsAvailable() bool {
	return true
}

// Register 向引擎注册方法
func (m *HttpModule) Register(engine model.Engine) error {
	state := engine.GetState()

	httpObj := state.NewTable()
	state.SetGlobal("http", httpObj)

	httpObj.RawSetString("get", state.NewFunction(func(L *lua.LState) int {
		url := L.CheckString(1)
		timeout := 5000
		if L.GetTop() > 1 {
			timeout = L.CheckInt(2)
		}
		code, data := https.Get(url, timeout)
		result := L.NewTable()
		L.SetField(result, "code", lua.LNumber(code))
		if data != nil {
			L.SetField(result, "data", lua.LString(string(data)))
		} else {
			L.SetField(result, "data", lua.LNil)
		}
		L.Push(result)
		return 1
	}))

	httpObj.RawSetString("post", state.NewFunction(func(L *lua.LState) int {
		url := L.CheckString(1)
		
		var data []byte
		if L.GetTop() > 1 {
			dataStr := L.CheckString(2)
			data = []byte(dataStr)
		}
		
		headers := make(map[string]string)
		if L.GetTop() > 2 {
			if L.GetTop() >= 3 && L.CheckAny(3).Type() == lua.LTTable {
				headersTable := L.CheckTable(3)
				headersTable.ForEach(func(key lua.LValue, value lua.LValue) {
					if keyStr, ok := key.(lua.LString); ok {
						if valueStr, ok := value.(lua.LString); ok {
							headers[string(keyStr)] = string(valueStr)
						}
					}
				})
			}
		}
		
		timeout := 5000
		if L.GetTop() > 3 {
			timeout = L.CheckInt(4)
		}
		
		code, body := https.Post(url, data, headers, timeout)
		result := L.NewTable()
		L.SetField(result, "code", lua.LNumber(code))
		if body != nil {
			L.SetField(result, "data", lua.LString(string(body)))
		} else {
			L.SetField(result, "data", lua.LNil)
		}
		L.Push(result)
		return 1
	}))

	httpObj.RawSetString("postMultipart", state.NewFunction(func(L *lua.LState) int {
		url := L.CheckString(1)
		fileName := L.CheckString(2)
		
		var fileData []byte
		if L.GetTop() > 2 {
			dataStr := L.CheckString(3)
			fileData = []byte(dataStr)
		}
		
		timeout := 5000
		if L.GetTop() > 3 {
			timeout = L.CheckInt(4)
		}
		
		code, responseData := https.PostMultipart(url, fileName, fileData, timeout)
		result := L.NewTable()
		L.SetField(result, "code", lua.LNumber(code))
		if responseData != nil {
			L.SetField(result, "data", lua.LString(string(responseData)))
		} else {
			L.SetField(result, "data", lua.LNil)
		}
		L.Push(result)
		return 1
	}))

	engine.RegisterMethod("http.get", "发送GET请求", func(url string, timeout int) (int, []byte) { return https.Get(url, timeout) }, true)
	engine.RegisterMethod("http.post", "发送POST请求", func(url string, data []byte, headers map[string]string, timeout int) (int, []byte) {
		return https.Post(url, data, headers, timeout)
	}, true)
	engine.RegisterMethod("http.postMultipart", "发送Multipart POST请求", func(url, fileName string, fileData []byte, timeout int) (int, []byte) {
		return https.PostMultipart(url, fileName, fileData, timeout)
	}, true)

	return nil
}
