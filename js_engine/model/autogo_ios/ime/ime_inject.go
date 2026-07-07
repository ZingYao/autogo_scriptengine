package ime

import (
	"sync"

	"github.com/ZingYao/autogo_scriptengine/js_engine/model"

	"github.com/dop251/goja"
)

var (
	clipboardMu   sync.Mutex
	clipboardText string
)

// ImeModule iOS ime 模块。
type ImeModule struct{}

// Name 返回模块名称。
func (m *ImeModule) Name() string {
	return "ime"
}

// IsAvailable 检查模块是否可用。
func (m *ImeModule) IsAvailable() bool {
	return true
}

// Register 向引擎注册 iOS ime 方法。
func (m *ImeModule) Register(engine model.Engine) error {
	vm := engine.GetVM()
	imeObj := vm.NewObject()
	vm.Set("ime", imeObj)

	imeObj.Set("getClipText", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(getClipText())
	})
	imeObj.Set("setClipText", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(setClipText(call.Argument(0).String()))
	})
	imeObj.Set("inputText", func(call goja.FunctionCall) goja.Value {
		inputText(call.Argument(0).String())
		return goja.Undefined()
	})

	engine.RegisterMethod("ime.getClipText", "获取剪切板内容", getClipText, true)
	engine.RegisterMethod("ime.setClipText", "设置剪切板内容", setClipText, true)
	engine.RegisterMethod("ime.inputText", "输入文本", inputText, true)
	return nil
}

func getClipText() string {
	clipboardMu.Lock()
	defer clipboardMu.Unlock()
	return clipboardText
}

func setClipText(text string) bool {
	clipboardMu.Lock()
	defer clipboardMu.Unlock()
	clipboardText = text
	return true
}

func inputText(text string) {
	_ = setClipText(text)
}
