package https

import (
	"fmt"

	"github.com/ZingYao/autogo_scriptengine/js_engine/model"

	autogohttps "github.com/Dasongzi1366/AutoGo/https"
	"github.com/dop251/goja"
)

// HttpsModule iOS https 模块。
type HttpsModule struct{}

// Name 返回模块名称。
func (m *HttpsModule) Name() string {
	return "https"
}

// IsAvailable 检查模块是否可用。
func (m *HttpsModule) IsAvailable() bool {
	return true
}

// bytesFromValue 将脚本侧字符串或字节数组参数转换为 Go 字节切片。
func bytesFromValue(value goja.Value) []byte {
	switch data := value.Export().(type) {
	case []byte:
		return data
	case string:
		return []byte(data)
	case []interface{}:
		result := make([]byte, 0, len(data))
		for _, item := range data {
			switch number := item.(type) {
			case int:
				result = append(result, byte(number))
			case int64:
				result = append(result, byte(number))
			case float64:
				result = append(result, byte(number))
			}
		}
		return result
	default:
		return nil
	}
}

// headersFromValue 将脚本侧对象参数转换为 HTTP 请求头。
func headersFromValue(value goja.Value) map[string]string {
	headers := map[string]string{}
	if goja.IsUndefined(value) || goja.IsNull(value) {
		return headers
	}
	exported := value.Export()
	switch typed := exported.(type) {
	case map[string]string:
		return typed
	case map[string]interface{}:
		for key, item := range typed {
			headers[key] = vmString(item)
		}
	}
	return headers
}

// vmString 将请求头值转换为字符串，避免非字符串值丢失。
func vmString(value interface{}) string {
	switch typed := value.(type) {
	case string:
		return typed
	case []byte:
		return string(typed)
	default:
		return fmt.Sprint(typed)
	}
}

// responseObject 统一返回 HTTP 状态码和响应体。
func responseObject(vm *goja.Runtime, code int, data []byte) *goja.Object {
	result := vm.NewObject()
	result.Set("code", code)
	if data == nil {
		result.Set("data", nil)
		return result
	}
	result.Set("data", string(data))
	return result
}

// Register 向引擎注册 iOS https 方法。
func (m *HttpsModule) Register(engine model.Engine) error {
	vm := engine.GetVM()
	httpsObj := vm.NewObject()
	vm.Set("https", httpsObj)

	httpsObj.Set("get", func(call goja.FunctionCall) goja.Value {
		timeout := 5000
		if len(call.Arguments) >= 2 {
			timeout = int(call.Argument(1).ToInteger())
		}
		code, data := autogohttps.Get(call.Argument(0).String(), timeout)
		return responseObject(vm, code, data)
	})
	httpsObj.Set("post", func(call goja.FunctionCall) goja.Value {
		timeout := 5000
		if len(call.Arguments) >= 4 {
			timeout = int(call.Argument(3).ToInteger())
		}
		code, data := autogohttps.Post(call.Argument(0).String(), bytesFromValue(call.Argument(1)), headersFromValue(call.Argument(2)), timeout)
		return responseObject(vm, code, data)
	})
	httpsObj.Set("postMultipart", func(call goja.FunctionCall) goja.Value {
		timeout := 5000
		if len(call.Arguments) >= 4 {
			timeout = int(call.Argument(3).ToInteger())
		}
		code, data := autogohttps.PostMultipart(call.Argument(0).String(), call.Argument(1).String(), bytesFromValue(call.Argument(2)), timeout)
		return responseObject(vm, code, data)
	})

	engine.RegisterMethod("https.get", "发送 GET 请求", autogohttps.Get, true)
	engine.RegisterMethod("https.post", "发送 POST 请求", autogohttps.Post, true)
	engine.RegisterMethod("https.postMultipart", "发送 multipart POST 请求", autogohttps.PostMultipart, true)
	return nil
}
