//go:build ignore
// +build ignore

package network

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/ZingYao/autogo_scriptengine/common/email"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"
	"github.com/gorilla/websocket"

	lua "github.com/yuin/gopher-lua"
)

// NetworkModule 网络模块（懒人精灵兼容）
type NetworkModule struct {
	wsConnections map[int]*websocket.Conn
	nextHandle    int
	wsMu          sync.Mutex
	emailClients  map[string]*email.Client
}

// Name 返回模块名称
func (m *NetworkModule) Name() string {
	return "network"
}

// IsAvailable 检查模块是否可用
func (m *NetworkModule) IsAvailable() bool {
	return true
}

// Register 向引擎注册方法
func (m *NetworkModule) Register(engine model.Engine) error {
	state := engine.GetState()

	// 初始化管理器
	m.wsConnections = make(map[int]*websocket.Conn)
	m.nextHandle = 1
	m.emailClients = make(map[string]*email.Client)

	// 注册 import 函数（空操作，用于兼容懒人精灵代码）
	state.SetGlobal("import", state.NewFunction(func(L *lua.LState) int {
		// 在 Go 环境中不需要真的 import java 包
		// 这里只是空操作，保持与懒人精灵代码的兼容性
		return 0
	}))

	// 注册 httpGet - HTTP GET请求
	state.SetGlobal("httpGet", state.NewFunction(func(L *lua.LState) int {
		urlStr := L.CheckString(1)
		timeout := L.OptInt(2, 30)

		// 使用 Go 标准库发送 HTTP GET 请求
		client := &http.Client{Timeout: time.Duration(timeout) * time.Second}
		resp, err := client.Get(urlStr)
		if err != nil {
			L.Push(lua.LString(""))
			L.Push(lua.LNumber(0))
			return 2
		}
		defer resp.Body.Close()

		// 读取响应
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			L.Push(lua.LString(""))
			L.Push(lua.LNumber(0))
			return 2
		}

		L.Push(lua.LString(string(body)))
		L.Push(lua.LNumber(resp.StatusCode))
		return 2
	}))

	// 注册 httpPost - HTTP POST请求
	state.SetGlobal("httpPost", state.NewFunction(func(L *lua.LState) int {
		urlStr := L.CheckString(1)
		postData := L.CheckString(2)
		timeout := L.OptInt(3, 30)

		// 使用 Go 标准库发送 HTTP POST 请求
		client := &http.Client{Timeout: time.Duration(timeout) * time.Second}
		resp, err := client.Post(urlStr, "application/x-www-form-urlencoded", strings.NewReader(postData))
		if err != nil {
			L.Push(lua.LString(""))
			L.Push(lua.LNumber(0))
			return 2
		}
		defer resp.Body.Close()

		// 读取响应
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			L.Push(lua.LString(""))
			L.Push(lua.LNumber(0))
			return 2
		}

		L.Push(lua.LString(string(body)))
		L.Push(lua.LNumber(resp.StatusCode))
		return 2
	}))

	// 注册 asynHttpPost - HTTP POST异步请求
	state.SetGlobal("asynHttpPost", state.NewFunction(func(L *lua.LState) int {
		callback := L.CheckFunction(1)
		urlStr := L.CheckString(2)
		postData := L.CheckString(3)
		timeout := L.OptInt(4, 30)

		// 创建 channel 来传递结果
		resultChan := make(chan struct {
			body string
			code int
			err  error
		}, 1)

		// 在 goroutine 中异步执行
		go func() {
			client := &http.Client{Timeout: time.Duration(timeout) * time.Second}
			resp, err := client.Post(urlStr, "application/x-www-form-urlencoded", strings.NewReader(postData))
			if err != nil {
				resultChan <- struct {
					body string
					code int
					err  error
				}{"", 0, err}
				return
			}
			defer resp.Body.Close()

			// 读取响应
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				resultChan <- struct {
					body string
					code int
					err  error
				}{"", 0, err}
				return
			}

			resultChan <- struct {
				body string
				code int
				err  error
			}{string(body), resp.StatusCode, nil}
		}()

		// 在另一个 goroutine 中等待结果并调用回调
		go func() {
			result := <-resultChan
			// 调用回调函数
			L.Push(callback)
			L.Push(lua.LString(result.body))
			L.Push(lua.LNumber(result.code))
			L.Call(2, 0)
		}()

		return 0
	}))

	// 注册 asynHttpGet - HTTP GET异步请求
	state.SetGlobal("asynHttpGet", state.NewFunction(func(L *lua.LState) int {
		callback := L.CheckFunction(1)
		urlStr := L.CheckString(2)
		timeout := L.OptInt(3, 30)

		// 创建 channel 来传递结果
		resultChan := make(chan struct {
			body string
			code int
			err  error
		}, 1)

		// 在 goroutine 中异步执行
		go func() {
			client := &http.Client{Timeout: time.Duration(timeout) * time.Second}
			resp, err := client.Get(urlStr)
			if err != nil {
				resultChan <- struct {
					body string
					code int
					err  error
				}{"", 0, err}
				return
			}
			defer resp.Body.Close()

			// 读取响应
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				resultChan <- struct {
					body string
					code int
					err  error
				}{"", 0, err}
				return
			}

			resultChan <- struct {
				body string
				code int
				err  error
			}{string(body), resp.StatusCode, nil}
		}()

		// 在另一个 goroutine 中等待结果并调用回调
		go func() {
			result := <-resultChan
			// 调用回调函数
			L.Push(callback)
			L.Push(lua.LString(result.body))
			L.Push(lua.LNumber(result.code))
			L.Call(2, 0)
		}()

		return 0
	}))

	// 注册 downloadFile - 下载文件
	state.SetGlobal("downloadFile", state.NewFunction(func(L *lua.LState) int {
		urlStr := L.CheckString(1)
		savePath := L.CheckString(2)

		// 使用 Go 标准库下载文件
		client := &http.Client{Timeout: 30 * time.Second}
		resp, err := client.Get(urlStr)
		if err != nil {
			L.Push(lua.LBool(false))
			return 1
		}
		defer resp.Body.Close()

		// 创建目标目录
		dir := filepath.Dir(savePath)
		if err := os.MkdirAll(dir, 0755); err != nil {
			L.Push(lua.LBool(false))
			return 1
		}

		// 创建目标文件
		file, err := os.Create(savePath)
		if err != nil {
			L.Push(lua.LBool(false))
			return 1
		}
		defer file.Close()

		// 复制数据
		_, err = io.Copy(file, resp.Body)
		if err != nil {
			L.Push(lua.LBool(false))
			return 1
		}

		L.Push(lua.LBool(true))
		return 1
	}))

	// 注册 uploadFile - 上传文件
	state.SetGlobal("uploadFile", state.NewFunction(func(L *lua.LState) int {
		urlStr := L.CheckString(1)
		uploadFile := L.CheckString(2)
		timeout := L.OptInt(3, 30)

		// 打开要上传的文件
		file, err := os.Open(uploadFile)
		if err != nil {
			L.Push(lua.LString(""))
			return 1
		}
		defer file.Close()

		// 创建 multipart form
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		// 添加文件
		part, err := writer.CreateFormFile("file", filepath.Base(uploadFile))
		if err != nil {
			L.Push(lua.LString(""))
			return 1
		}

		io.Copy(part, file)
		writer.Close()

		// 发送请求
		req, _ := http.NewRequest("POST", urlStr, body)
		req.Header.Set("Content-Type", writer.FormDataContentType())

		client := &http.Client{Timeout: time.Duration(timeout) * time.Second}
		resp, err := client.Do(req)
		if err != nil {
			L.Push(lua.LString(""))
			return 1
		}
		defer resp.Body.Close()

		// 读取响应
		respBody, _ := io.ReadAll(resp.Body)
		L.Push(lua.LString(string(respBody)))
		return 1
	}))

	// 注册 httpUpload - HTTP 文件上传（别名）
	state.SetGlobal("httpUpload", state.NewFunction(func(L *lua.LState) int {
		urlStr := L.CheckString(1)
		fieldName := L.CheckString(2)
		uploadFile := L.CheckString(3)
		headers := L.OptTable(4, L.NewTable())
		timeout := L.OptInt(5, 30)

		// 打开要上传的文件
		file, err := os.Open(uploadFile)
		if err != nil {
			L.Push(lua.LString(""))
			L.Push(lua.LNumber(0))
			return 2
		}
		defer file.Close()

		// 创建 multipart form
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		// 添加文件
		part, err := writer.CreateFormFile(fieldName, filepath.Base(uploadFile))
		if err != nil {
			L.Push(lua.LString(""))
			L.Push(lua.LNumber(0))
			return 2
		}

		io.Copy(part, file)
		writer.Close()

		// 发送请求
		req, _ := http.NewRequest("POST", urlStr, body)
		req.Header.Set("Content-Type", writer.FormDataContentType())

		// 添加自定义请求头
		headers.ForEach(func(key, value lua.LValue) {
			if keyStr, ok := key.(lua.LString); ok {
				if valueStr, ok := value.(lua.LString); ok {
					req.Header.Set(string(keyStr), string(valueStr))
				}
			}
		})

		client := &http.Client{Timeout: time.Duration(timeout) * time.Second}
		resp, err := client.Do(req)
		if err != nil {
			L.Push(lua.LString(""))
			L.Push(lua.LNumber(0))
			return 2
		}
		defer resp.Body.Close()

		// 读取响应
		respBody, _ := io.ReadAll(resp.Body)
		L.Push(lua.LString(string(respBody)))
		L.Push(lua.LNumber(resp.StatusCode))
		return 2
	}))

	// 注册 httpDownload - HTTP 文件下载（别名）
	state.SetGlobal("httpDownload", state.NewFunction(func(L *lua.LState) int {
		urlStr := L.CheckString(1)
		savePath := L.CheckString(2)
		timeout := L.OptInt(3, 30)

		// 使用 Go 标准库下载文件
		client := &http.Client{Timeout: time.Duration(timeout) * time.Second}
		resp, err := client.Get(urlStr)
		if err != nil {
			L.Push(lua.LBool(false))
			L.Push(lua.LString(err.Error()))
			return 2
		}
		defer resp.Body.Close()

		// 创建目标目录
		dir := filepath.Dir(savePath)
		if err := os.MkdirAll(dir, 0755); err != nil {
			L.Push(lua.LBool(false))
			L.Push(lua.LString(err.Error()))
			return 2
		}

		// 创建目标文件
		file, err := os.Create(savePath)
		if err != nil {
			L.Push(lua.LBool(false))
			L.Push(lua.LString(err.Error()))
			return 2
		}
		defer file.Close()

		// 复制数据
		_, err = io.Copy(file, resp.Body)
		if err != nil {
			L.Push(lua.LBool(false))
			L.Push(lua.LString(err.Error()))
			return 2
		}

		L.Push(lua.LBool(true))
		return 1
	}))

	// 注册 startWebSocket - 打开一个websocket连接
	state.SetGlobal("startWebSocket", state.NewFunction(func(L *lua.LState) int {
		urlStr := L.CheckString(1)
		onOpened := L.CheckFunction(2)
		onClosed := L.CheckFunction(3)
		onError := L.CheckFunction(4)
		onRecv := L.CheckFunction(5)

		// 使用 gorilla/websocket 连接
		conn, _, err := websocket.DefaultDialer.Dial(urlStr, nil)
		if err != nil {
			L.Push(onError)
			L.Push(lua.LNumber(0))
			L.Push(lua.LString(err.Error()))
			L.Call(2, 0)
			L.Push(lua.LNumber(0))
			return 1
		}

		// 获取句柄并在闭包中捕获
		m.wsMu.Lock()
		handle := m.nextHandle
		m.nextHandle++
		m.wsConnections[handle] = conn
		m.wsMu.Unlock()

		// 启动读取循环
		go func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Printf("panic in readLoop: %v\n", r)
				}
			}()

			for {
				// 读取消息
				_, message, err := conn.ReadMessage()
				if err != nil {
					// 调用错误回调
					if onError != nil {
						thread, _ := L.NewThread()
						thread.Push(onError)
						thread.Push(lua.LNumber(handle))
						thread.Push(lua.LString(err.Error()))
						if err := thread.PCall(2, 0, nil); err != nil {
							fmt.Printf("Error calling onError callback: %v\n", err)
						}
					}

					// 调用关闭回调
					if onClosed != nil {
						thread, _ := L.NewThread()
						thread.Push(onClosed)
						thread.Push(lua.LNumber(handle))
						if err := thread.PCall(1, 0, nil); err != nil {
							fmt.Printf("Error calling onClosed callback: %v\n", err)
						}
					}
					break
				}

				// 调用消息回调
				if onRecv != nil {
					thread, _ := L.NewThread()
					thread.Push(onRecv)
					thread.Push(lua.LNumber(handle))
					thread.Push(lua.LString(string(message)))
					if err := thread.PCall(2, 0, nil); err != nil {
						fmt.Printf("Error calling onRecv callback: %v\n", err)
					}
				}
			}
		}()

		// 调用 onOpened 回调
		if onOpened != nil {
			L.Push(onOpened)
			L.Push(lua.LNumber(handle))
			if err := L.PCall(1, 0, nil); err != nil {
				fmt.Printf("Error calling onOpened callback: %v\n", err)
			}
		}

		// 返回句柄
		L.Push(lua.LNumber(handle))
		return 1
	}))

	// 注册 closeWebSocket - 关闭一个websocket连接
	state.SetGlobal("closeWebSocket", state.NewFunction(func(L *lua.LState) int {
		handle := L.CheckInt(1)
		m.wsMu.Lock()
		defer m.wsMu.Unlock()

		if conn, ok := m.wsConnections[handle]; ok {
			conn.Close()
			delete(m.wsConnections, handle)
		}
		return 0
	}))

	// 注册 sendWebSocket - 向服务器发送数据
	state.SetGlobal("sendWebSocket", state.NewFunction(func(L *lua.LState) int {
		handle := L.CheckInt(1)
		text := L.CheckString(2)

		m.wsMu.Lock()
		conn, ok := m.wsConnections[handle]
		m.wsMu.Unlock()

		if !ok {
			L.Push(lua.LBool(false))
			return 1
		}

		err := conn.WriteMessage(websocket.TextMessage, []byte(text))
		if err != nil {
			L.Push(lua.LBool(false))
			return 1
		}

		L.Push(lua.LBool(true))
		return 1
	}))

	// 注册 LuaEngine.httpPost - java层http post
	luaEngineTable := state.NewTable()
	state.SetGlobal("LuaEngine", luaEngineTable)
	luaEngineTable.RawSetString("httpPost", state.NewFunction(func(L *lua.LState) int {
		urlStr := L.CheckString(1)
		paramsTable := L.CheckTable(2)
		headersTable := L.OptTable(3, L.NewTable())
		timeout := L.OptInt(4, 30)

		// 将 Lua 表转换为 URL 参数
		values := url.Values{}
		paramsTable.ForEach(func(key, value lua.LValue) {
			if keyStr, ok := key.(lua.LString); ok {
				if valueStr, ok := value.(lua.LString); ok {
					values.Set(string(keyStr), string(valueStr))
				}
			}
		})

		// 创建请求
		req, err := http.NewRequest("POST", urlStr, strings.NewReader(values.Encode()))
		if err != nil {
			L.Push(lua.LString(""))
			return 1
		}

		// 设置请求头
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		headersTable.ForEach(func(key, value lua.LValue) {
			if keyStr, ok := key.(lua.LString); ok {
				if valueStr, ok := value.(lua.LString); ok {
					req.Header.Set(string(keyStr), string(valueStr))
				}
			}
		})

		// 发送请求
		client := &http.Client{Timeout: time.Duration(timeout) * time.Second}
		resp, err := client.Do(req)
		if err != nil {
			L.Push(lua.LString(""))
			return 1
		}
		defer resp.Body.Close()

		// 读取响应
		body, _ := io.ReadAll(resp.Body)
		L.Push(lua.LString(string(body)))
		return 1
	}))

	// 注册 LuaEngine.httpPostData - java层http post 任意数据
	luaEngineTable.RawSetString("httpPostData", state.NewFunction(func(L *lua.LState) int {
		urlStr := L.CheckString(1)
		params := L.CheckString(2)
		headers := L.CheckString(3)
		timeout := L.OptInt(4, 30)

		// 创建请求
		req, err := http.NewRequest("POST", urlStr, strings.NewReader(params))
		if err != nil {
			L.Push(lua.LString(""))
			return 1
		}

		// 设置请求头
		req.Header.Set("Content-Type", headers)

		// 发送请求
		client := &http.Client{Timeout: time.Duration(timeout) * time.Second}
		resp, err := client.Do(req)
		if err != nil {
			L.Push(lua.LString(""))
			return 1
		}
		defer resp.Body.Close()

		// 读取响应
		body, _ := io.ReadAll(resp.Body)
		L.Push(lua.LString(string(body)))
		return 1
	}))

	// 注册 LuaEngine.httpGet - java层http get
	luaEngineTable.RawSetString("httpGet", state.NewFunction(func(L *lua.LState) int {
		urlStr := L.CheckString(1)
		paramsTable := L.OptTable(2, L.NewTable())
		timeout := L.OptInt(3, 30)

		// 构建完整 URL
		fullURL := urlStr
		if paramsTable != nil {
			values := url.Values{}
			paramsTable.ForEach(func(key, value lua.LValue) {
				if keyStr, ok := key.(lua.LString); ok {
					if valueStr, ok := value.(lua.LString); ok {
						values.Set(string(keyStr), string(valueStr))
					}
				}
			})
			if len(values) > 0 {
				fullURL = urlStr + "?" + values.Encode()
			}
		}

		// 发送请求
		client := &http.Client{Timeout: time.Duration(timeout) * time.Second}
		resp, err := client.Get(fullURL)
		if err != nil {
			L.Push(lua.LString(""))
			return 1
		}
		defer resp.Body.Close()

		// 读取响应
		body, _ := io.ReadAll(resp.Body)
		L.Push(lua.LString(string(body)))
		return 1
	}))

	// 注册 LuaEngine.sendMail - 发送一个普通文本邮件
	luaEngineTable.RawSetString("sendMail", state.NewFunction(func(L *lua.LState) int {
		username := L.CheckString(1)
		password := L.CheckString(2)
		to := L.CheckString(3)
		smtpServer := L.CheckString(4)
		useAuth := L.CheckBool(5)
		subject := L.CheckString(6)
		body := L.CheckString(7)

		// 创建邮件客户端
		client := email.NewClient(smtpServer, username, password, useAuth)

		// 发送邮件
		err := client.Send(to, subject, body)
		if err != nil {
			L.Push(lua.LBool(false))
			return 1
		}

		L.Push(lua.LBool(true))
		return 1
	}))

	// 注册 LuaEngine.sendMailWithFile - 发送一个带单个附件的邮件
	luaEngineTable.RawSetString("sendMailWithFile", state.NewFunction(func(L *lua.LState) int {
		username := L.CheckString(1)
		password := L.CheckString(2)
		to := L.CheckString(3)
		smtpServer := L.CheckString(4)
		useAuth := L.CheckBool(5)
		subject := L.CheckString(6)
		body := L.CheckString(7)
		filePath := L.CheckString(8)

		// 创建邮件客户端
		client := email.NewClient(smtpServer, username, password, useAuth)

		// 发送带附件的邮件
		err := client.SendWithAttachment(to, subject, body, []string{filePath})
		if err != nil {
			L.Push(lua.LBool(false))
			return 1
		}

		L.Push(lua.LBool(true))
		return 1
	}))

	// 注册 LuaEngine.sendMailWithMultiFile - 发送一个带多个附件的邮件
	luaEngineTable.RawSetString("sendMailWithMultiFile", state.NewFunction(func(L *lua.LState) int {
		username := L.CheckString(1)
		password := L.CheckString(2)
		to := L.CheckString(3)
		smtpServer := L.CheckString(4)
		useAuth := L.CheckBool(5)
		subject := L.CheckString(6)
		body := L.CheckString(7)
		filesTable := L.CheckTable(8)

		// 提取文件路径
		var files []string
		filesTable.ForEach(func(key, value lua.LValue) {
			if filePath, ok := value.(lua.LString); ok {
				files = append(files, string(filePath))
			}
		})

		// 创建邮件客户端
		client := email.NewClient(smtpServer, username, password, useAuth)

		// 发送带多个附件的邮件
		err := client.SendWithAttachment(to, subject, body, files)
		if err != nil {
			L.Push(lua.LBool(false))
			return 1
		}

		L.Push(lua.LBool(true))
		return 1
	}))

	engine.RegisterMethod("network.httpGet", "HTTP GET请求", networkHTTPGet, true)
	engine.RegisterMethod("network.httpPost", "HTTP POST请求", networkHTTPPost, true)
	engine.RegisterMethod("network.httpPostData", "HTTP POST任意数据", networkHTTPPostData, true)
	engine.RegisterMethod("network.downloadFile", "下载文件", networkDownloadFile, true)
	engine.RegisterMethod("network.uploadFile", "上传文件", networkUploadFile, true)
	engine.RegisterMethod("network.httpUpload", "HTTP 文件上传", networkHTTPUpload, true)
	engine.RegisterMethod("network.httpDownload", "HTTP 文件下载", networkHTTPDownload, true)

	return nil
}

func networkHTTPGet(urlStr string, timeout ...int) (map[string]interface{}, error) {
	client := &http.Client{Timeout: networkTimeout(timeout...)}
	resp, err := client.Get(urlStr)
	if err != nil {
		return map[string]interface{}{"body": "", "statusCode": 0}, nil
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return map[string]interface{}{"body": "", "statusCode": 0}, nil
	}
	return map[string]interface{}{
		"body":       string(body),
		"statusCode": resp.StatusCode,
	}, nil
}

func networkHTTPPost(urlStr, postData string, timeout ...int) (map[string]interface{}, error) {
	client := &http.Client{Timeout: networkTimeout(timeout...)}
	resp, err := client.Post(urlStr, "application/x-www-form-urlencoded", strings.NewReader(postData))
	if err != nil {
		return map[string]interface{}{"body": "", "statusCode": 0}, nil
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return map[string]interface{}{"body": "", "statusCode": 0}, nil
	}
	return map[string]interface{}{
		"body":       string(body),
		"statusCode": resp.StatusCode,
	}, nil
}

func networkHTTPPostData(urlStr, params, contentType string, timeout ...int) string {
	req, err := http.NewRequest("POST", urlStr, strings.NewReader(params))
	if err != nil {
		return ""
	}
	req.Header.Set("Content-Type", contentType)

	client := &http.Client{Timeout: networkTimeout(timeout...)}
	resp, err := client.Do(req)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	return string(body)
}

func networkDownloadFile(urlStr, savePath string) bool {
	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Get(urlStr)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	dir := filepath.Dir(savePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return false
	}
	file, err := os.Create(savePath)
	if err != nil {
		return false
	}
	defer file.Close()
	_, err = io.Copy(file, resp.Body)
	return err == nil
}

func networkUploadFile(urlStr, uploadFile string, timeout ...int) string {
	options := make([]interface{}, 0, len(timeout))
	for _, timeoutValue := range timeout {
		options = append(options, timeoutValue)
	}
	body, statusCode := networkHTTPUpload(urlStr, "file", uploadFile, options...)
	if statusCode == 0 {
		return ""
	}
	return body
}

func networkHTTPUpload(urlStr, fieldName, uploadFile string, options ...interface{}) (string, int) {
	headers, timeout := networkUploadOptions(options...)
	file, err := os.Open(uploadFile)
	if err != nil {
		return "", 0
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(fieldName, filepath.Base(uploadFile))
	if err != nil {
		return "", 0
	}
	if _, err := io.Copy(part, file); err != nil {
		return "", 0
	}
	if err := writer.Close(); err != nil {
		return "", 0
	}

	req, err := http.NewRequest("POST", urlStr, body)
	if err != nil {
		return "", 0
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	for key, value := range headers {
		req.Header.Set(key, fmt.Sprint(value))
	}

	client := &http.Client{Timeout: networkTimeout(timeout)}
	resp, err := client.Do(req)
	if err != nil {
		return "", 0
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", 0
	}
	return string(respBody), resp.StatusCode
}

func networkUploadOptions(options ...interface{}) (map[string]interface{}, int) {
	headers := map[string]interface{}{}
	timeout := 30
	for _, option := range options {
		switch value := option.(type) {
		case map[string]interface{}:
			headers = value
		case map[string]string:
			for key, item := range value {
				headers[key] = item
			}
		case int:
			timeout = value
		case int64:
			timeout = int(value)
		case float64:
			timeout = int(value)
		}
	}
	return headers, timeout
}

func networkHTTPDownload(urlStr, savePath string, timeout ...int) bool {
	client := &http.Client{Timeout: networkTimeout(timeout...)}
	resp, err := client.Get(urlStr)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	dir := filepath.Dir(savePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return false
	}
	file, err := os.Create(savePath)
	if err != nil {
		return false
	}
	defer file.Close()
	_, err = io.Copy(file, resp.Body)
	return err == nil
}

func networkTimeout(timeout ...int) time.Duration {
	seconds := 30
	if len(timeout) > 0 && timeout[0] > 0 {
		seconds = timeout[0]
	}
	return time.Duration(seconds) * time.Second
}
