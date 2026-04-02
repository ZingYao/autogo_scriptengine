package utils

import (
	"github.com/ZingYao/autogo_scriptengine/common"
	"github.com/ZingYao/autogo_scriptengine/js_engine/model"

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
		x := 0
		y := 0
		duration := -1
		if len(call.Arguments) > 1 {
			x = int(call.Argument(1).ToInteger())
		}
		if len(call.Arguments) > 2 {
			y = int(call.Argument(2).ToInteger())
		}
		if len(call.Arguments) > 3 {
			duration = int(call.Argument(3).ToInteger())
		}
		utils.Toast(message, x, y, duration)
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

	utilsObj.Set("md5", func(call goja.FunctionCall) goja.Value {
		s := call.Argument(0).String()
		result := common.Md5(s)
		return vm.ToValue(result)
	})

	utilsObj.Set("sha1", func(call goja.FunctionCall) goja.Value {
		s := call.Argument(0).String()
		result := common.Sha1(s)
		return vm.ToValue(result)
	})

	utilsObj.Set("sha256", func(call goja.FunctionCall) goja.Value {
		s := call.Argument(0).String()
		result := common.Sha256(s)
		return vm.ToValue(result)
	})

	utilsObj.Set("base64Encode", func(call goja.FunctionCall) goja.Value {
		s := call.Argument(0).String()
		result := common.Base64Encode(s)
		return vm.ToValue(result)
	})

	utilsObj.Set("base64Decode", func(call goja.FunctionCall) goja.Value {
		s := call.Argument(0).String()
		result, err := common.Base64Decode(s)
		if err != nil {
			return vm.ToValue("")
		}
		return vm.ToValue(result)
	})

	utilsObj.Set("urlEncode", func(call goja.FunctionCall) goja.Value {
		s := call.Argument(0).String()
		result := common.UrlEncode(s)
		return vm.ToValue(result)
	})

	utilsObj.Set("urlDecode", func(call goja.FunctionCall) goja.Value {
		s := call.Argument(0).String()
		result, err := common.UrlDecode(s)
		if err != nil {
			return vm.ToValue("")
		}
		return vm.ToValue(result)
	})

	utilsObj.Set("htmlEncode", func(call goja.FunctionCall) goja.Value {
		s := call.Argument(0).String()
		result := common.HtmlEncode(s)
		return vm.ToValue(result)
	})

	utilsObj.Set("htmlDecode", func(call goja.FunctionCall) goja.Value {
		s := call.Argument(0).String()
		result := common.HtmlDecode(s)
		return vm.ToValue(result)
	})

	utilsObj.Set("jsonEncode", func(call goja.FunctionCall) goja.Value {
		v := call.Argument(0).Export()
		result, err := common.JsonEncode(v)
		if err != nil {
			return vm.ToValue("")
		}
		return vm.ToValue(result)
	})

	utilsObj.Set("jsonDecode", func(call goja.FunctionCall) goja.Value {
		s := call.Argument(0).String()
		var result interface{}
		err := common.JsonDecode(s, &result)
		if err != nil {
			return goja.Undefined()
		}
		return vm.ToValue(result)
	})

	utilsObj.Set("timestamp", func(call goja.FunctionCall) goja.Value {
		result := common.Timestamp()
		return vm.ToValue(result)
	})

	engine.RegisterMethod("utils.logI", "记录一条INFO级别的日志", utils.LogI, true)
	engine.RegisterMethod("utils.logE", "记录一条ERROR级别的日志", utils.LogE, true)
	engine.RegisterMethod("utils.toast", "显示Toast提示", func(message string, x, y, duration int) { utils.Toast(message, x, y, duration) }, true)
	engine.RegisterMethod("utils.alert", "显示Alert对话框", func(title, content, btn1Text, btn2Text string) int {
		return utils.Alert(title, content, btn1Text, btn2Text)
	}, true)
	engine.RegisterMethod("utils.shell", "执行shell命令并返回输出", utils.Shell, true)
	engine.RegisterMethod("utils.random", "返回指定范围内的随机整数", utils.Random, true)
	engine.RegisterMethod("utils.sleep", "暂停当前线程指定的毫秒数", utils.Sleep, true)
	engine.RegisterMethod("utils.i2s", "将整数转换为字符串", utils.I2s, true)
	engine.RegisterMethod("utils.s2i", "将字符串转换为整数", utils.S2i, true)
	engine.RegisterMethod("utils.f2s", "将浮点数转换为字符串", utils.F2s, true)
	engine.RegisterMethod("utils.s2f", "将字符串转换为浮点数", utils.S2f, true)
	engine.RegisterMethod("utils.b2s", "将布尔值转换为字符串", utils.B2s, true)
	engine.RegisterMethod("utils.s2b", "将字符串转换为布尔值", utils.S2b, true)
	engine.RegisterMethod("utils.md5", "计算字符串的MD5哈希值", nil, true)
	engine.RegisterMethod("utils.sha1", "计算字符串的SHA1哈希值", nil, true)
	engine.RegisterMethod("utils.sha256", "计算字符串的SHA256哈希值", nil, true)
	engine.RegisterMethod("utils.base64Encode", "将字符串进行Base64编码", nil, true)
	engine.RegisterMethod("utils.base64Decode", "将Base64编码的字符串解码", nil, true)
	engine.RegisterMethod("utils.urlEncode", "对字符串进行URL编码", nil, true)
	engine.RegisterMethod("utils.urlDecode", "对URL编码的字符串进行解码", nil, true)
	engine.RegisterMethod("utils.htmlEncode", "对字符串进行HTML实体编码", nil, true)
	engine.RegisterMethod("utils.htmlDecode", "对HTML实体编码的字符串进行解码", nil, true)
	engine.RegisterMethod("utils.jsonEncode", "将对象转换为JSON字符串", nil, true)
	engine.RegisterMethod("utils.jsonDecode", "将JSON字符串解析为对象", nil, true)
	engine.RegisterMethod("utils.timestamp", "获取当前时间戳（秒）", nil, true)

	return nil
}
