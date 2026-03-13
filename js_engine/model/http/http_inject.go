package http

import (
	"app/js_engine/model"

	"github.com/Dasongzi1366/AutoGo/https"
	"github.com/dop251/goja"
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
	vm := engine.GetVM()

	httpObj := vm.NewObject()
	vm.Set("http", httpObj)

	httpObj.Set("get", func(call goja.FunctionCall) goja.Value {
		url := call.Argument(0).String()
		timeout := int(call.Argument(1).ToInteger())
		code, data := https.Get(url, timeout)
		result := vm.NewObject()
		result.Set("code", code)
		if data != nil {
			result.Set("data", string(data))
		} else {
			result.Set("data", goja.Null())
		}
		return result
	})

	httpObj.Set("post", func(call goja.FunctionCall) goja.Value {
		url := call.Argument(0).String()
		
		// 处理第二个参数，支持字符串或字节数组
		var data []byte
		arg1 := call.Argument(1)
		if str, ok := arg1.Export().(string); ok {
			data = []byte(str)
		} else if bytes, ok := arg1.Export().([]byte); ok {
			data = bytes
		}
		
		headers := make(map[string]string)
		if len(call.Arguments) > 2 {
			headersVal := call.Argument(2).Export()
			if headersMap, ok := headersVal.(map[string]interface{}); ok {
				for k, v := range headersMap {
					if str, ok := v.(string); ok {
						headers[k] = str
					}
				}
			}
		}
		
		timeout := 5000
		if len(call.Arguments) > 3 {
			timeout = int(call.Argument(3).ToInteger())
		}
		
		code, body := https.Post(url, data, headers, timeout)
		result := vm.NewObject()
		result.Set("code", code)
		if body != nil {
			result.Set("data", string(body))
		} else {
			result.Set("data", goja.Null())
		}
		return result
	})

	httpObj.Set("postMultipart", func(call goja.FunctionCall) goja.Value {
		url := call.Argument(0).String()
		fileName := call.Argument(1).String()
		timeout := int(call.Argument(3).ToInteger())

		// 处理第三个参数，支持字符串或字节数组
		var fileData []byte
		arg2 := call.Argument(2)
		if str, ok := arg2.Export().(string); ok {
			fileData = []byte(str)
		} else if bytes, ok := arg2.Export().([]byte); ok {
			fileData = bytes
		}

		code, responseData := https.PostMultipart(url, fileName, fileData, timeout)
		result := vm.NewObject()
		result.Set("code", code)
		if responseData != nil {
			result.Set("data", string(responseData))
		} else {
			result.Set("data", goja.Null())
		}
		return result
	})

	engine.RegisterMethod("http.get", "发送GET请求", func(url string, timeout int) (int, []byte) { return https.Get(url, timeout) }, true)
	engine.RegisterMethod("http.post", "发送POST请求", func(url string, data []byte, headers map[string]string, timeout int) (int, []byte) {
		return https.Post(url, data, headers, timeout)
	}, true)
	engine.RegisterMethod("http.postMultipart", "发送Multipart POST请求", func(url, fileName string, fileData []byte, timeout int) (int, []byte) {
		return https.PostMultipart(url, fileName, fileData, timeout)
	}, true)

	return nil
}
