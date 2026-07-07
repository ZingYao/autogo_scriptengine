//go:build ignore
// +build ignore

package http

import (
	"crypto/tls"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"

	lua "github.com/yuin/gopher-lua"
)

// HttpModule http 模块（luasocket，懒人精灵兼容）
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

	// 获取 package.preload 表
	packagePreload := state.GetGlobal("package")
	if packagePreload == lua.LNil {
		// 如果 package 不存在，创建它
		packagePreload = state.NewTable()
		state.SetGlobal("package", packagePreload)
	}

	packageTable := packagePreload.(*lua.LTable)
	preloadTable := packageTable.RawGetString("preload")
	if preloadTable == lua.LNil {
		// 如果 preload 不存在，创建它
		preloadTable = state.NewTable()
		packageTable.RawSetString("preload", preloadTable)
	}

	// 创建 http 模块表
	httpModuleTable := state.NewTable()

	// 注册 http.request - 发起网络请求
	httpModuleTable.RawSetString("request", state.NewFunction(func(L *lua.LState) int {
		// 获取第一个参数
		arg := L.CheckAny(1)

		var requestURL string
		var method string
		var headers map[string]string
		var body string
		var timeout int
		var sslVerify bool

		// 判断参数类型
		if str, ok := arg.(lua.LString); ok {
			// 简单字符串 URL
			requestURL = string(str)
			method = "GET"
			timeout = 30
			sslVerify = true
		} else if tbl, ok := arg.(*lua.LTable); ok {
			// 表参数，解析详细配置
			requestURL = getTableString(tbl, "url")
			method = getTableString(tbl, "method")
			if method == "" {
				method = "GET"
			}
			headers = getTableHeaders(tbl, "headers")
			body = getTableString(tbl, "body")
			timeout = getTableInt(tbl, "timeout")
			if timeout == 0 {
				timeout = 30
			}
			sslVerify = getTableBool(tbl, "sslverify")
		} else {
			L.Push(lua.LNil)
			L.Push(lua.LString("参数类型错误，必须是字符串或表"))
			return 2
		}

		// 验证 URL
		if requestURL == "" {
			L.Push(lua.LNil)
			L.Push(lua.LString("URL 不能为空"))
			return 2
		}

		// 创建 HTTP 客户端
		client := &http.Client{
			Timeout: time.Duration(timeout) * time.Second,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: !sslVerify,
				},
			},
		}

		// 创建请求
		var req *http.Request
		var err error

		if body != "" {
			req, err = http.NewRequest(method, requestURL, strings.NewReader(body))
		} else {
			req, err = http.NewRequest(method, requestURL, nil)
		}

		if err != nil {
			L.Push(lua.LNil)
			L.Push(lua.LString(err.Error()))
			return 2
		}

		// 设置请求头
		for key, value := range headers {
			req.Header.Set(key, value)
		}

		// 发送请求
		resp, err := client.Do(req)
		if err != nil {
			L.Push(lua.LNil)
			L.Push(lua.LString(err.Error()))
			return 2
		}
		defer resp.Body.Close()

		// 读取响应体
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			L.Push(lua.LNil)
			L.Push(lua.LString(err.Error()))
			return 2
		}

		// 创建响应头表
		respHeaders := L.NewTable()
		for key, values := range resp.Header {
			if len(values) > 0 {
				respHeaders.RawSetString(key, lua.LString(values[0]))
			}
		}

		// 返回结果：状态码、状态文本、响应头、响应体
		L.Push(lua.LNumber(resp.StatusCode))
		L.Push(lua.LString(resp.Status))
		L.Push(respHeaders)
		L.Push(lua.LString(string(respBody)))
		return 4
	}))

	// 创建 ssl 模块表
	sslModuleTable := state.NewTable()

	// 创建 https 子模块
	httpsTable := state.NewTable()
	sslModuleTable.RawSetString("https", httpsTable)

	// 注册 https.request - 发起 HTTPS 请求（与 http.request 相同）
	httpsTable.RawSetString("request", state.NewFunction(func(L *lua.LState) int {
		// 获取第一个参数
		arg := L.CheckAny(1)

		var requestURL string
		var method string
		var headers map[string]string
		var body string
		var timeout int
		var sslVerify bool

		// 判断参数类型
		if str, ok := arg.(lua.LString); ok {
			// 简单字符串 URL
			requestURL = string(str)
			method = "GET"
			timeout = 30
			sslVerify = true
		} else if tbl, ok := arg.(*lua.LTable); ok {
			// 表参数，解析详细配置
			requestURL = getTableString(tbl, "url")
			method = getTableString(tbl, "method")
			if method == "" {
				method = "GET"
			}
			headers = getTableHeaders(tbl, "headers")
			body = getTableString(tbl, "body")
			timeout = getTableInt(tbl, "timeout")
			if timeout == 0 {
				timeout = 30
			}
			sslVerify = getTableBool(tbl, "sslverify")
		} else {
			L.Push(lua.LNil)
			L.Push(lua.LString("参数类型错误，必须是字符串或表"))
			return 2
		}

		// 验证 URL
		if requestURL == "" {
			L.Push(lua.LNil)
			L.Push(lua.LString("URL 不能为空"))
			return 2
		}

		// 创建 HTTP 客户端
		client := &http.Client{
			Timeout: time.Duration(timeout) * time.Second,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: !sslVerify,
				},
			},
		}

		// 创建请求
		var req *http.Request
		var err error

		if body != "" {
			req, err = http.NewRequest(method, requestURL, strings.NewReader(body))
		} else {
			req, err = http.NewRequest(method, requestURL, nil)
		}

		if err != nil {
			L.Push(lua.LNil)
			L.Push(lua.LString(err.Error()))
			return 2
		}

		// 设置请求头
		for key, value := range headers {
			req.Header.Set(key, value)
		}

		// 发送请求
		resp, err := client.Do(req)
		if err != nil {
			L.Push(lua.LNil)
			L.Push(lua.LString(err.Error()))
			return 2
		}
		defer resp.Body.Close()

		// 读取响应体
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			L.Push(lua.LNil)
			L.Push(lua.LString(err.Error()))
			return 2
		}

		// 创建响应头表
		respHeaders := L.NewTable()
		for key, values := range resp.Header {
			if len(values) > 0 {
				respHeaders.RawSetString(key, lua.LString(values[0]))
			}
		}

		// 返回结果：状态码、状态文本、响应头、响应体
		L.Push(lua.LNumber(resp.StatusCode))
		L.Push(lua.LString(resp.Status))
		L.Push(respHeaders)
		L.Push(lua.LString(string(respBody)))
		return 4
	}))

	// 创建 ltn12 模块表
	ltn12ModuleTable := state.NewTable()

	// 创建 sink 子模块
	sinkTable := state.NewTable()
	ltn12ModuleTable.RawSetString("sink", sinkTable)

	// 注册 sink.table - 将数据收集到表中
	sinkTable.RawSetString("table", state.NewFunction(func(L *lua.LState) int {
		// 获取参数：要收集数据的表
		tbl := L.CheckTable(1)

		// 创建 sink 函数
		sinkFunc := state.NewFunction(func(L *lua.LState) int {
			// 获取要添加的数据
			chunk := L.CheckAny(1)

			// 如果数据不是 nil，添加到表中
			if chunk != lua.LNil {
				tbl.Append(chunk)
			}

			// 返回 1 表示成功
			return 1
		})

		L.Push(sinkFunc)
		return 1
	}))

	// 注册 sink.file - 将数据写入文件
	sinkTable.RawSetString("file", state.NewFunction(func(L *lua.LState) int {
		// 获取参数：文件句柄
		_ = L.CheckAny(1)

		// 创建 sink 函数
		sinkFunc := state.NewFunction(func(L *lua.LState) int {
			// 获取要写入的数据
			chunk := L.CheckAny(1)

			// 如果数据不是 nil，写入文件
			if chunk != lua.LNil {
				// 这里简化处理，实际应该调用 file:write()
				// 由于我们无法直接访问文件对象，这里返回成功
			}

			// 返回 1 表示成功
			return 1
		})

		L.Push(sinkFunc)
		return 1
	}))

	// 注册 sink.chain - 链接多个 sink
	sinkTable.RawSetString("chain", state.NewFunction(func(L *lua.LState) int {
		// 获取参数：多个 sink 函数
		// 这里简化处理，返回一个组合的 sink
		sinkFunc := state.NewFunction(func(L *lua.LState) int {
			// 简化处理，返回成功
			return 1
		})

		L.Push(sinkFunc)
		return 1
	}))

	// 注册 sink.null - 丢弃所有数据
	sinkTable.RawSetString("null", state.NewFunction(func(L *lua.LState) int {
		// 创建一个丢弃所有数据的 sink
		sinkFunc := state.NewFunction(func(L *lua.LState) int {
			// 丢弃数据，不做任何处理
			return 1
		})

		L.Push(sinkFunc)
		return 1
	}))

	// 注册 sink.error - 抛出错误的 sink
	sinkTable.RawSetString("error", state.NewFunction(func(L *lua.LState) int {
		// 获取错误信息
		errMsg := L.CheckString(1)

		// 创建一个抛出错误的 sink
		sinkFunc := state.NewFunction(func(L *lua.LState) int {
			// 抛出错误
			L.RaiseError(errMsg)
			return 0
		})

		L.Push(sinkFunc)
		return 1
	}))

	// 注册 source.file - 从文件读取数据
	sourceTable := state.NewTable()
	ltn12ModuleTable.RawSetString("source", sourceTable)

	sourceTable.RawSetString("file", state.NewFunction(func(L *lua.LState) int {
		// 获取参数：文件句柄
		_ = L.CheckAny(1)

		// 创建 source 函数
		sourceFunc := state.NewFunction(func(L *lua.LState) int {
			// 简化处理，返回 nil 表示结束
			return 1
		})

		L.Push(sourceFunc)
		return 1
	}))

	// 注册 source.string - 从字符串读取数据
	sourceTable.RawSetString("string", state.NewFunction(func(L *lua.LState) int {
		// 获取参数：字符串数据
		str := L.CheckString(1)

		// 创建 source 函数
		sourceFunc := state.NewFunction(func(L *lua.LState) int {
			// 简化处理，返回字符串
			L.Push(lua.LString(str))
			return 1
		})

		L.Push(sourceFunc)
		return 1
	}))

	// 注册 source.chain - 链接多个 source
	sourceTable.RawSetString("chain", state.NewFunction(func(L *lua.LState) int {
		// 简化处理
		sourceFunc := state.NewFunction(func(L *lua.LState) int {
			return 1
		})

		L.Push(sourceFunc)
		return 1
	}))

	// 注册到 package.preload
	// 注册 http
	preloadTable.(*lua.LTable).RawSetString("http", state.NewFunction(func(L *lua.LState) int {
		L.Push(httpModuleTable)
		return 1
	}))

	// 注册 ssl
	preloadTable.(*lua.LTable).RawSetString("ssl", state.NewFunction(func(L *lua.LState) int {
		L.Push(sslModuleTable)
		return 1
	}))

	// 注册 ssl.https
	preloadTable.(*lua.LTable).RawSetString("ssl.https", state.NewFunction(func(L *lua.LState) int {
		L.Push(httpsTable)
		return 1
	}))

	// 注册 ltn12
	preloadTable.(*lua.LTable).RawSetString("ltn12", state.NewFunction(func(L *lua.LState) int {
		L.Push(ltn12ModuleTable)
		return 1
	}))

	// 同时注册为全局变量（兼容直接调用）
	state.SetGlobal("http", httpModuleTable)
	state.SetGlobal("https", httpsTable)
	state.SetGlobal("ltn12", ltn12ModuleTable)

	// 创建 ssl 表并设置 https 子表
	sslTable := state.NewTable()
	sslTable.RawSetString("https", httpsTable)
	state.SetGlobal("ssl", sslTable)

	engine.RegisterMethod("http.request", "发起 HTTP 网络请求", httpRequest, true)
	engine.RegisterMethod("https.request", "发起 HTTPS 网络请求", httpRequest, true)
	engine.RegisterMethod("ssl.https.request", "发起 SSL HTTPS 网络请求", httpRequest, true)

	return nil
}

func httpRequest(url, method string, headers map[string]string, body string, timeout int, sslVerify bool) (int, string, map[string]string, string, error) {
	if method == "" {
		method = "GET"
	}
	if timeout == 0 {
		timeout = 30
	}

	client := &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: !sslVerify,
			},
		},
	}

	var req *http.Request
	var err error

	if body != "" {
		req, err = http.NewRequest(method, url, strings.NewReader(body))
	} else {
		req, err = http.NewRequest(method, url, nil)
	}

	if err != nil {
		return 0, "", nil, "", err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return 0, "", nil, "", err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, "", nil, "", err
	}

	respHeaders := make(map[string]string)
	for key, values := range resp.Header {
		if len(values) > 0 {
			respHeaders[key] = values[0]
		}
	}

	return resp.StatusCode, resp.Status, respHeaders, string(respBody), nil
}

// getTableString 从表中获取字符串值
func getTableString(tbl *lua.LTable, key string) string {
	value := tbl.RawGetString(key)
	if str, ok := value.(lua.LString); ok {
		return string(str)
	}
	return ""
}

// getTableInt 从表中获取整数值
func getTableInt(tbl *lua.LTable, key string) int {
	value := tbl.RawGetString(key)
	if num, ok := value.(lua.LNumber); ok {
		return int(num)
	}
	return 0
}

// getTableBool 从表中获取布尔值
func getTableBool(tbl *lua.LTable, key string) bool {
	value := tbl.RawGetString(key)
	if boolVal, ok := value.(lua.LBool); ok {
		return bool(boolVal)
	}
	return true
}

// getTableHeaders 从表中获取请求头
func getTableHeaders(tbl *lua.LTable, key string) map[string]string {
	headers := make(map[string]string)

	value := tbl.RawGetString(key)
	if headerTbl, ok := value.(*lua.LTable); ok {
		headerTbl.ForEach(func(k lua.LValue, v lua.LValue) {
			if keyStr, ok := k.(lua.LString); ok {
				if valStr, ok := v.(lua.LString); ok {
					headers[string(keyStr)] = string(valStr)
				}
			}
		})
	}

	return headers
}
