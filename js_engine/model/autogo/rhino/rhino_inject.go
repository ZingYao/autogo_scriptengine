package rhino

import (
	"github.com/ZingYao/autogo_scriptengine/js_engine/model"

	"github.com/Dasongzi1366/AutoGo/rhino"
	"github.com/dop251/goja"
)

// RhinoModule rhino 模块
type RhinoModule struct{}

// Name 返回模块名称
func (m *RhinoModule) Name() string {
	return "rhino"
}

// IsAvailable 检查模块是否可用
func (m *RhinoModule) IsAvailable() bool {
	return true
}

// Register 向引擎注册方法
func (m *RhinoModule) Register(engine model.Engine) error {
	vm := engine.GetVM()

	rhinoObj := vm.NewObject()
	vm.Set("rhino", rhinoObj)

	rhinoObj.Set("eval", func(call goja.FunctionCall) goja.Value {
		contextId := call.Argument(0).String()
		script := call.Argument(1).String()
		result := rhino.Eval(contextId, script)
		return vm.ToValue(result)
	})

	engine.RegisterMethod("rhino.eval", "执行指定的JavaScript脚本并返回结果", rhino.Eval, true)

	return nil
}
