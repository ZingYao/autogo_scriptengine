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
	vm := engine.GetVM()

	imeObj := vm.NewObject()
	vm.Set("ime", imeObj)

	imeObj.Set("getClipText", func(call goja.FunctionCall) goja.Value {
		result := getClipText()
		return vm.ToValue(result)
	})

	imeObj.Set("setClipText", func(call goja.FunctionCall) goja.Value {
		text := call.Argument(0).String()
		result := setClipText(text)
		return vm.ToValue(result)
	})

	imeObj.Set("keyAction", func(call goja.FunctionCall) goja.Value {
		code := int(call.Argument(0).ToInteger())
		keyAction(code)
		return goja.Undefined()
	})

	imeObj.Set("inputText", func(call goja.FunctionCall) goja.Value {
		text := call.Argument(0).String()
		displayId := 0
		if len(call.Arguments) >= 2 {
			displayId = int(call.Argument(1).ToInteger())
		}
		inputText(text, displayId)
		return goja.Undefined()
	})

	imeObj.Set("getIMEList", func(call goja.FunctionCall) goja.Value {
		result := getIMEList()
		return vm.ToValue(result)
	})

	imeObj.Set("setCurrentIME", func(call goja.FunctionCall) goja.Value {
		packageName := call.Argument(0).String()
		setCurrentIME(packageName)
		return goja.Undefined()
	})

	engine.RegisterMethod("ime.getClipText", "获取剪切板内容", getClipText, true)
	engine.RegisterMethod("ime.setClipText", "设置剪切板内容", setClipText, true)
	engine.RegisterMethod("ime.keyAction", "模拟按键", keyAction, true)
	engine.RegisterMethod("ime.inputText", "输入文本", inputText, true)
	engine.RegisterMethod("ime.getIMEList", "获取输入法列表", getIMEList, true)
	engine.RegisterMethod("ime.setCurrentIME", "设置当前输入法", setCurrentIME, true)

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

func keyAction(_ int) {
}

func inputText(text string, _ int) {
	_ = setClipText(text)
}

func getIMEList() []string {
	return []string{}
}

func setCurrentIME(_ string) {
}
