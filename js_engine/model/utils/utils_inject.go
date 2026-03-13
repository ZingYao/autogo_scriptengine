package utils

import (
	"app/js_engine/model"

	"github.com/Dasongzi1366/AutoGo/utils"
	"github.com/dop251/goja"
)

// UtilsModule utils 模块
type UtilsModule struct{}

// Name 返回模块名称
func (m *UtilsModule) Name() string {
	return "utils"
}

// IsAvailable 检查模块是否可用
func (m *UtilsModule) IsAvailable() bool {
	return true
}

// Register 向引擎注册方法
func (m *UtilsModule) Register(engine model.Engine) error {
	vm := engine.GetVM()

	utilsObj := vm.NewObject()
	vm.Set("utils", utilsObj)

	utilsObj.Set("logI", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		message := ""
		for i := 1; i < len(call.Arguments); i++ {
			message += call.Argument(i).String() + " "
		}
		utils.LogI(label, message)
		return goja.Undefined()
	})

	utilsObj.Set("logE", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		message := ""
		for i := 1; i < len(call.Arguments); i++ {
			message += call.Argument(i).String() + " "
		}
		utils.LogE(label, message)
		return goja.Undefined()
	})

	utilsObj.Set("toast", func(call goja.FunctionCall) goja.Value {
		message := call.Argument(0).String()
		utils.Toast(message)
		return goja.Undefined()
	})

	utilsObj.Set("alert", func(call goja.FunctionCall) goja.Value {
		title := call.Argument(0).String()
		content := call.Argument(1).String()
		btn1Text := ""
		if len(call.Arguments) > 2 {
			btn1Text = call.Argument(2).String()
		}
		btn2Text := ""
		if len(call.Arguments) > 3 {
			btn2Text = call.Argument(3).String()
		}
		result := utils.Alert(title, content, btn1Text, btn2Text)
		return vm.ToValue(result)
	})

	utilsObj.Set("shell", func(call goja.FunctionCall) goja.Value {
		cmd := call.Argument(0).String()
		result := utils.Shell(cmd)
		return vm.ToValue(result)
	})

	utilsObj.Set("random", func(call goja.FunctionCall) goja.Value {
		min := int(call.Argument(0).ToInteger())
		max := int(call.Argument(1).ToInteger())
		result := utils.Random(min, max)
		return vm.ToValue(result)
	})

	utilsObj.Set("sleep", func(call goja.FunctionCall) goja.Value {
		i := int(call.Argument(0).ToInteger())
		utils.Sleep(i)
		return goja.Undefined()
	})

	utilsObj.Set("i2s", func(call goja.FunctionCall) goja.Value {
		i := int(call.Argument(0).ToInteger())
		result := utils.I2s(i)
		return vm.ToValue(result)
	})

	utilsObj.Set("s2i", func(call goja.FunctionCall) goja.Value {
		s := call.Argument(0).String()
		result := utils.S2i(s)
		return vm.ToValue(result)
	})

	utilsObj.Set("f2s", func(call goja.FunctionCall) goja.Value {
		f := call.Argument(0).ToFloat()
		result := utils.F2s(f)
		return vm.ToValue(result)
	})

	utilsObj.Set("s2f", func(call goja.FunctionCall) goja.Value {
		s := call.Argument(0).String()
		result := utils.S2f(s)
		return vm.ToValue(result)
	})

	utilsObj.Set("b2s", func(call goja.FunctionCall) goja.Value {
		b := call.Argument(0).ToBoolean()
		result := utils.B2s(b)
		return vm.ToValue(result)
	})

	utilsObj.Set("s2b", func(call goja.FunctionCall) goja.Value {
		s := call.Argument(0).String()
		result := utils.S2b(s)
		return vm.ToValue(result)
	})

	engine.RegisterMethod("utils.logI", "记录一条INFO级别的日志", utils.LogI, true)
	engine.RegisterMethod("utils.logE", "记录一条ERROR级别的日志", utils.LogE, true)
	engine.RegisterMethod("utils.toast", "显示Toast提示", utils.Toast, true)
	engine.RegisterMethod("utils.alert", "显示Alert对话框", func(title, content, btn1Text, btn2Text string) int { return utils.Alert(title, content, btn1Text, btn2Text) }, true)
	engine.RegisterMethod("utils.shell", "执行shell命令并返回输出", utils.Shell, true)
	engine.RegisterMethod("utils.random", "返回指定范围内的随机整数", utils.Random, true)
	engine.RegisterMethod("utils.sleep", "暂停当前线程指定的毫秒数", utils.Sleep, true)
	engine.RegisterMethod("utils.i2s", "将整数转换为字符串", utils.I2s, true)
	engine.RegisterMethod("utils.s2i", "将字符串转换为整数", utils.S2i, true)
	engine.RegisterMethod("utils.f2s", "将浮点数转换为字符串", utils.F2s, true)
	engine.RegisterMethod("utils.s2f", "将字符串转换为浮点数", utils.S2f, true)
	engine.RegisterMethod("utils.b2s", "将布尔值转换为字符串", utils.B2s, true)
	engine.RegisterMethod("utils.s2b", "将字符串转换为布尔值", utils.S2b, true)

	return nil
}
