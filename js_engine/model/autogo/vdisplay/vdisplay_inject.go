package vdisplay

import (
	"fmt"

	"github.com/ZingYao/autogo_scriptengine/js_engine/model"
	"github.com/ZingYao/goja"
)

// VdisplayModule 在远程 AutoGo 缺少 vdisplay 包时保留同名入口并返回明确错误。
type VdisplayModule struct{}

func (m *VdisplayModule) Name() string { return "vdisplay" }

func (m *VdisplayModule) IsAvailable() bool { return false }

func (m *VdisplayModule) Register(engine model.Engine) error {
	vm := engine.GetVM()
	vdisplayObj := vm.NewObject()
	vm.Set("vdisplay", vdisplayObj)

	vdisplayObj.Set("create", func(call goja.FunctionCall) goja.Value {
		panic(vm.NewGoError(errUnavailable()))
	})
	vdisplayObj.Set("getDisplayId", func(call goja.FunctionCall) goja.Value {
		panic(vm.NewGoError(errUnavailable()))
	})
	vdisplayObj.Set("launchApp", func(call goja.FunctionCall) goja.Value {
		panic(vm.NewGoError(errUnavailable()))
	})
	vdisplayObj.Set("setTitle", func(call goja.FunctionCall) goja.Value {
		panic(vm.NewGoError(errUnavailable()))
	})
	vdisplayObj.Set("setTouchCallback", func(call goja.FunctionCall) goja.Value {
		panic(vm.NewGoError(errUnavailable()))
	})
	vdisplayObj.Set("showPreviewWindow", func(call goja.FunctionCall) goja.Value {
		panic(vm.NewGoError(errUnavailable()))
	})
	vdisplayObj.Set("hidePreviewWindow", func(call goja.FunctionCall) goja.Value {
		panic(vm.NewGoError(errUnavailable()))
	})
	vdisplayObj.Set("setPreviewWindowSize", func(call goja.FunctionCall) goja.Value {
		panic(vm.NewGoError(errUnavailable()))
	})
	vdisplayObj.Set("setPreviewWindowPos", func(call goja.FunctionCall) goja.Value {
		panic(vm.NewGoError(errUnavailable()))
	})
	vdisplayObj.Set("destroy", func(call goja.FunctionCall) goja.Value {
		panic(vm.NewGoError(errUnavailable()))
	})

	engine.RegisterMethod("vdisplay.create", "AutoGo/vdisplay remote package unavailable", func(args ...interface{}) (interface{}, error) {
		return nil, errUnavailable()
	}, true)
	engine.RegisterMethod("vdisplay.getDisplayId", "AutoGo/vdisplay remote package unavailable", func(args ...interface{}) (interface{}, error) {
		return nil, errUnavailable()
	}, true)
	engine.RegisterMethod("vdisplay.launchApp", "AutoGo/vdisplay remote package unavailable", func(args ...interface{}) (interface{}, error) {
		return nil, errUnavailable()
	}, true)
	engine.RegisterMethod("vdisplay.setTitle", "AutoGo/vdisplay remote package unavailable", func(args ...interface{}) (interface{}, error) {
		return nil, errUnavailable()
	}, true)
	engine.RegisterMethod("vdisplay.setTouchCallback", "AutoGo/vdisplay remote package unavailable", func(args ...interface{}) (interface{}, error) {
		return nil, errUnavailable()
	}, true)
	engine.RegisterMethod("vdisplay.showPreviewWindow", "AutoGo/vdisplay remote package unavailable", func(args ...interface{}) (interface{}, error) {
		return nil, errUnavailable()
	}, true)
	engine.RegisterMethod("vdisplay.hidePreviewWindow", "AutoGo/vdisplay remote package unavailable", func(args ...interface{}) (interface{}, error) {
		return nil, errUnavailable()
	}, true)
	engine.RegisterMethod("vdisplay.setPreviewWindowSize", "AutoGo/vdisplay remote package unavailable", func(args ...interface{}) (interface{}, error) {
		return nil, errUnavailable()
	}, true)
	engine.RegisterMethod("vdisplay.setPreviewWindowPos", "AutoGo/vdisplay remote package unavailable", func(args ...interface{}) (interface{}, error) {
		return nil, errUnavailable()
	}, true)
	engine.RegisterMethod("vdisplay.destroy", "AutoGo/vdisplay remote package unavailable", func(args ...interface{}) (interface{}, error) {
		return nil, errUnavailable()
	}, true)
	return nil
}

func errUnavailable() error {
	return fmt.Errorf("AutoGo/%s package is unavailable in the remote github.com/Dasongzi1366/AutoGo module", "vdisplay")
}
