package js_engine

import (
	"github.com/Dasongzi1366/AutoGo/https"
	"github.com/dop251/goja"
)

func injectHttpsMethods(engine *JSEngine) {
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

	httpObj.Set("postMultipart", func(call goja.FunctionCall) goja.Value {
		url := call.Argument(0).String()
		fileName := call.Argument(1).String()
		fileData := call.Argument(2).Export().([]byte)
		timeout := int(call.Argument(3).ToInteger())
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
	engine.RegisterMethod("http.postMultipart", "发送Multipart POST请求", func(url, fileName string, fileData []byte, timeout int) (int, []byte) {
		return https.PostMultipart(url, fileName, fileData, timeout)
	}, true)
}
