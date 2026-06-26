package utils

import (
	"github.com/ZingYao/autogo_scriptengine/js_engine/model"

	autogoutils "github.com/Dasongzi1366/AutoGo/utils"
	"github.com/dop251/goja"
)

// UtilsModule iOS utils 模块。
type UtilsModule struct{}

// Name 返回模块名称。
func (m *UtilsModule) Name() string {
	return "utils"
}

// IsAvailable 检查模块是否可用。
func (m *UtilsModule) IsAvailable() bool {
	return true
}

// Register 向引擎注册 iOS utils 方法。
func (m *UtilsModule) Register(engine model.Engine) error {
	vm := engine.GetVM()
	utilsObj := vm.NewObject()
	vm.Set("utils", utilsObj)

	utilsObj.Set("toast", func(call goja.FunctionCall) goja.Value {
		x, y, duration := -1, -1, -1
		if len(call.Arguments) >= 2 {
			x = int(call.Argument(1).ToInteger())
		}
		if len(call.Arguments) >= 3 {
			y = int(call.Argument(2).ToInteger())
		}
		if len(call.Arguments) >= 4 {
			duration = int(call.Argument(3).ToInteger())
		}
		autogoutils.Toast(call.Argument(0).String(), x, y, duration)
		return goja.Undefined()
	})
	utilsObj.Set("alert", func(call goja.FunctionCall) goja.Value {
		btn1Text, btn2Text := "", ""
		if len(call.Arguments) >= 3 {
			btn1Text = call.Argument(2).String()
		}
		if len(call.Arguments) >= 4 {
			btn2Text = call.Argument(3).String()
		}
		return vm.ToValue(autogoutils.Alert(call.Argument(0).String(), call.Argument(1).String(), btn1Text, btn2Text))
	})
	utilsObj.Set("inputAlert", func(call goja.FunctionCall) goja.Value {
		placeholder, defaultText, btn1Text, btn2Text := "", "", "", ""
		if len(call.Arguments) >= 3 {
			placeholder = call.Argument(2).String()
		}
		if len(call.Arguments) >= 4 {
			defaultText = call.Argument(3).String()
		}
		if len(call.Arguments) >= 5 {
			btn1Text = call.Argument(4).String()
		}
		if len(call.Arguments) >= 6 {
			btn2Text = call.Argument(5).String()
		}
		text, ok := autogoutils.InputAlert(call.Argument(0).String(), call.Argument(1).String(), placeholder, defaultText, btn1Text, btn2Text)
		result := vm.NewObject()
		result.Set("text", text)
		result.Set("ok", ok)
		return result
	})
	utilsObj.Set("random", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(autogoutils.Random(int(call.Argument(0).ToInteger()), int(call.Argument(1).ToInteger())))
	})
	utilsObj.Set("sleep", func(call goja.FunctionCall) goja.Value {
		autogoutils.Sleep(int(call.Argument(0).ToInteger()))
		return goja.Undefined()
	})
	utilsObj.Set("i2s", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(autogoutils.I2s(int(call.Argument(0).ToInteger())))
	})
	utilsObj.Set("s2i", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(autogoutils.S2i(call.Argument(0).String()))
	})
	utilsObj.Set("f2s", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(autogoutils.F2s(call.Argument(0).ToFloat()))
	})
	utilsObj.Set("s2f", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(autogoutils.S2f(call.Argument(0).String()))
	})
	utilsObj.Set("b2s", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(autogoutils.B2s(call.Argument(0).ToBoolean()))
	})
	utilsObj.Set("s2b", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(autogoutils.S2b(call.Argument(0).String()))
	})

	engine.RegisterMethod("utils.toast", "显示 Toast", autogoutils.Toast, true)
	engine.RegisterMethod("utils.alert", "显示提示弹窗", autogoutils.Alert, true)
	engine.RegisterMethod("utils.inputAlert", "显示输入弹窗", autogoutils.InputAlert, true)
	engine.RegisterMethod("utils.random", "生成随机整数", autogoutils.Random, true)
	engine.RegisterMethod("utils.sleep", "睡眠指定毫秒数", autogoutils.Sleep, true)
	engine.RegisterMethod("utils.i2s", "整数转字符串", autogoutils.I2s, true)
	engine.RegisterMethod("utils.s2i", "字符串转整数", autogoutils.S2i, true)
	engine.RegisterMethod("utils.f2s", "浮点数转字符串", autogoutils.F2s, true)
	engine.RegisterMethod("utils.s2f", "字符串转浮点数", autogoutils.S2f, true)
	engine.RegisterMethod("utils.b2s", "布尔值转字符串", autogoutils.B2s, true)
	engine.RegisterMethod("utils.s2b", "字符串转布尔值", autogoutils.S2b, true)
	return nil
}
