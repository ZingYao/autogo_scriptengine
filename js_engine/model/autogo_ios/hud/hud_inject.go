package hud

import (
	"fmt"

	"github.com/ZingYao/autogo_scriptengine/js_engine/model"
	"github.com/ZingYao/goja"
)

// HUDModule 在远程 AutoGo 缺少 hud 包时保留 iOS 同名入口并返回明确错误。
type HUDModule struct{}

func (m *HUDModule) Name() string { return "hud" }

func (m *HUDModule) IsAvailable() bool { return false }

func (m *HUDModule) Register(engine model.Engine) error {
	vm := engine.GetVM()
	hudObj := vm.NewObject()
	vm.Set("hud", hudObj)

	hudObj.Set("new", func(call goja.FunctionCall) goja.Value {
		panic(vm.NewGoError(errUnavailable()))
	})
	hudObj.Set("setText", func(call goja.FunctionCall) goja.Value {
		panic(vm.NewGoError(errUnavailable()))
	})
	hudObj.Set("setPosition", func(call goja.FunctionCall) goja.Value {
		panic(vm.NewGoError(errUnavailable()))
	})
	hudObj.Set("setBackgroundColor", func(call goja.FunctionCall) goja.Value {
		panic(vm.NewGoError(errUnavailable()))
	})
	hudObj.Set("setTextSize", func(call goja.FunctionCall) goja.Value {
		panic(vm.NewGoError(errUnavailable()))
	})
	hudObj.Set("show", func(call goja.FunctionCall) goja.Value {
		panic(vm.NewGoError(errUnavailable()))
	})
	hudObj.Set("hide", func(call goja.FunctionCall) goja.Value {
		panic(vm.NewGoError(errUnavailable()))
	})
	hudObj.Set("isVisible", func(call goja.FunctionCall) goja.Value {
		panic(vm.NewGoError(errUnavailable()))
	})
	hudObj.Set("destroy", func(call goja.FunctionCall) goja.Value {
		panic(vm.NewGoError(errUnavailable()))
	})

	engine.RegisterMethod("hud.new", "AutoGo/hud remote package unavailable", func(args ...interface{}) (interface{}, error) {
		return nil, errUnavailable()
	}, true)
	engine.RegisterMethod("hud.setText", "AutoGo/hud remote package unavailable", func(args ...interface{}) (interface{}, error) {
		return nil, errUnavailable()
	}, true)
	engine.RegisterMethod("hud.setPosition", "AutoGo/hud remote package unavailable", func(args ...interface{}) (interface{}, error) {
		return nil, errUnavailable()
	}, true)
	engine.RegisterMethod("hud.setBackgroundColor", "AutoGo/hud remote package unavailable", func(args ...interface{}) (interface{}, error) {
		return nil, errUnavailable()
	}, true)
	engine.RegisterMethod("hud.setTextSize", "AutoGo/hud remote package unavailable", func(args ...interface{}) (interface{}, error) {
		return nil, errUnavailable()
	}, true)
	engine.RegisterMethod("hud.show", "AutoGo/hud remote package unavailable", func(args ...interface{}) (interface{}, error) {
		return nil, errUnavailable()
	}, true)
	engine.RegisterMethod("hud.hide", "AutoGo/hud remote package unavailable", func(args ...interface{}) (interface{}, error) {
		return nil, errUnavailable()
	}, true)
	engine.RegisterMethod("hud.isVisible", "AutoGo/hud remote package unavailable", func(args ...interface{}) (interface{}, error) {
		return nil, errUnavailable()
	}, true)
	engine.RegisterMethod("hud.destroy", "AutoGo/hud remote package unavailable", func(args ...interface{}) (interface{}, error) {
		return nil, errUnavailable()
	}, true)
	return nil
}

func errUnavailable() error {
	return fmt.Errorf("AutoGo/%s package is unavailable in the remote github.com/Dasongzi1366/AutoGo module", "hud")
}
