package dotocr

import (
	"fmt"

	"github.com/ZingYao/autogo_scriptengine/js_engine/model"
	"github.com/ZingYao/goja"
)

// DotocrModule 在远程 AutoGo 缺少 dotocr 包时保留 iOS 同名入口并返回明确错误。
type DotocrModule struct{}

func (m *DotocrModule) Name() string { return "dotocr" }

func (m *DotocrModule) IsAvailable() bool { return false }

func (m *DotocrModule) Register(engine model.Engine) error {
	vm := engine.GetVM()
	dotocrObj := vm.NewObject()
	vm.Set("dotocr", dotocrObj)

	dotocrObj.Set("setDict", func(call goja.FunctionCall) goja.Value {
		panic(vm.NewGoError(errUnavailable()))
	})
	dotocrObj.Set("ocr", func(call goja.FunctionCall) goja.Value {
		panic(vm.NewGoError(errUnavailable()))
	})
	dotocrObj.Set("ocrFromImage", func(call goja.FunctionCall) goja.Value {
		panic(vm.NewGoError(errUnavailable()))
	})
	dotocrObj.Set("ocrFromBase64", func(call goja.FunctionCall) goja.Value {
		panic(vm.NewGoError(errUnavailable()))
	})
	dotocrObj.Set("ocrFromPath", func(call goja.FunctionCall) goja.Value {
		panic(vm.NewGoError(errUnavailable()))
	})
	dotocrObj.Set("findStr", func(call goja.FunctionCall) goja.Value {
		panic(vm.NewGoError(errUnavailable()))
	})
	dotocrObj.Set("findStrFromImage", func(call goja.FunctionCall) goja.Value {
		panic(vm.NewGoError(errUnavailable()))
	})
	dotocrObj.Set("findStrFromBase64", func(call goja.FunctionCall) goja.Value {
		panic(vm.NewGoError(errUnavailable()))
	})
	dotocrObj.Set("findStrFromPath", func(call goja.FunctionCall) goja.Value {
		panic(vm.NewGoError(errUnavailable()))
	})

	engine.RegisterMethod("dotocr.setDict", "AutoGo/dotocr remote package unavailable", func(args ...interface{}) (interface{}, error) {
		return nil, errUnavailable()
	}, true)
	engine.RegisterMethod("dotocr.ocr", "AutoGo/dotocr remote package unavailable", func(args ...interface{}) (interface{}, error) {
		return nil, errUnavailable()
	}, true)
	engine.RegisterMethod("dotocr.ocrFromImage", "AutoGo/dotocr remote package unavailable", func(args ...interface{}) (interface{}, error) {
		return nil, errUnavailable()
	}, true)
	engine.RegisterMethod("dotocr.ocrFromBase64", "AutoGo/dotocr remote package unavailable", func(args ...interface{}) (interface{}, error) {
		return nil, errUnavailable()
	}, true)
	engine.RegisterMethod("dotocr.ocrFromPath", "AutoGo/dotocr remote package unavailable", func(args ...interface{}) (interface{}, error) {
		return nil, errUnavailable()
	}, true)
	engine.RegisterMethod("dotocr.findStr", "AutoGo/dotocr remote package unavailable", func(args ...interface{}) (interface{}, error) {
		return nil, errUnavailable()
	}, true)
	engine.RegisterMethod("dotocr.findStrFromImage", "AutoGo/dotocr remote package unavailable", func(args ...interface{}) (interface{}, error) {
		return nil, errUnavailable()
	}, true)
	engine.RegisterMethod("dotocr.findStrFromBase64", "AutoGo/dotocr remote package unavailable", func(args ...interface{}) (interface{}, error) {
		return nil, errUnavailable()
	}, true)
	engine.RegisterMethod("dotocr.findStrFromPath", "AutoGo/dotocr remote package unavailable", func(args ...interface{}) (interface{}, error) {
		return nil, errUnavailable()
	}, true)
	return nil
}

func errUnavailable() error {
	return fmt.Errorf("AutoGo/%s package is unavailable in the remote github.com/Dasongzi1366/AutoGo module", "dotocr")
}
