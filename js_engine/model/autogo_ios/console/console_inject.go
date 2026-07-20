package console

import (
	"github.com/ZingYao/autogo_scriptengine/js_engine/model"

	autogoconsole "github.com/Dasongzi1366/AutoGo/console"
	"github.com/ZingYao/goja"
)

// ConsoleModule iOS console 模块。
type ConsoleModule struct{}

// Name 返回模块名称。
func (m *ConsoleModule) Name() string {
	return "console"
}

// IsAvailable 检查模块是否可用。
func (m *ConsoleModule) IsAvailable() bool {
	return true
}

// jsValuesToAnySlice 将 JavaScript 可变参数转换为 Println 可接受的参数切片。
func jsValuesToAnySlice(args []goja.Value) []any {
	values := make([]any, 0, len(args))
	for _, arg := range args {
		values = append(values, arg.Export())
	}
	return values
}

// consoleChainValue 返回链式调用对象。
func consoleChainValue(vm *goja.Runtime, current *goja.Object, original, next *autogoconsole.Console) goja.Value {
	if next != nil && next != original {
		return wrapConsole(vm, next)
	}
	return current
}

// wrapConsole 将 Go Console 对象包装为 JavaScript 实例对象。
func wrapConsole(vm *goja.Runtime, c *autogoconsole.Console) goja.Value {
	obj := vm.NewObject()
	obj.Set("setWindowSize", func(call goja.FunctionCall) goja.Value {
		next := c.SetWindowSize(int(call.Argument(0).ToInteger()), int(call.Argument(1).ToInteger()))
		return consoleChainValue(vm, obj, c, next)
	})
	obj.Set("setWindowPosition", func(call goja.FunctionCall) goja.Value {
		next := c.SetWindowPosition(int(call.Argument(0).ToInteger()), int(call.Argument(1).ToInteger()))
		return consoleChainValue(vm, obj, c, next)
	})
	obj.Set("setWindowColor", func(call goja.FunctionCall) goja.Value {
		next := c.SetWindowColor(call.Argument(0).String())
		return consoleChainValue(vm, obj, c, next)
	})
	obj.Set("setTextColor", func(call goja.FunctionCall) goja.Value {
		next := c.SetTextColor(call.Argument(0).String())
		return consoleChainValue(vm, obj, c, next)
	})
	obj.Set("setTextSize", func(call goja.FunctionCall) goja.Value {
		next := c.SetTextSize(int(call.Argument(0).ToInteger()))
		return consoleChainValue(vm, obj, c, next)
	})
	obj.Set("println", func(call goja.FunctionCall) goja.Value {
		c.Println(jsValuesToAnySlice(call.Arguments)...)
		return goja.Undefined()
	})
	obj.Set("clear", func(call goja.FunctionCall) goja.Value {
		c.Clear()
		return goja.Undefined()
	})
	obj.Set("show", func(call goja.FunctionCall) goja.Value {
		c.Show()
		return goja.Undefined()
	})
	obj.Set("hide", func(call goja.FunctionCall) goja.Value {
		c.Hide()
		return goja.Undefined()
	})
	obj.Set("isVisible", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(c.IsVisible())
	})
	obj.Set("destroy", func(call goja.FunctionCall) goja.Value {
		c.Destroy()
		return goja.Undefined()
	})
	return obj
}

// Register 向引擎注册 iOS console 方法。
func (m *ConsoleModule) Register(engine model.Engine) error {
	vm := engine.GetVM()
	consoleObj := vm.NewObject()
	vm.Set("console", consoleObj)
	consoleObj.Set("new", func(call goja.FunctionCall) goja.Value {
		c := autogoconsole.New()
		if c == nil {
			return goja.Null()
		}
		return wrapConsole(vm, c)
	})

	engine.RegisterMethod("console.new", "创建控制台对象", autogoconsole.New, true)
	engine.RegisterMethod("console.setWindowSize", "设置控制台窗口大小", func(c *autogoconsole.Console, width, height int) *autogoconsole.Console {
		return c.SetWindowSize(width, height)
	}, true)
	engine.RegisterMethod("console.setWindowPosition", "设置控制台窗口位置", func(c *autogoconsole.Console, x, y int) *autogoconsole.Console {
		return c.SetWindowPosition(x, y)
	}, true)
	engine.RegisterMethod("console.setWindowColor", "设置控制台窗口颜色", func(c *autogoconsole.Console, color string) *autogoconsole.Console {
		return c.SetWindowColor(color)
	}, true)
	engine.RegisterMethod("console.setTextColor", "设置控制台文本颜色", func(c *autogoconsole.Console, color string) *autogoconsole.Console {
		return c.SetTextColor(color)
	}, true)
	engine.RegisterMethod("console.setTextSize", "设置控制台文本大小", func(c *autogoconsole.Console, size int) *autogoconsole.Console {
		return c.SetTextSize(size)
	}, true)
	engine.RegisterMethod("console.println", "输出一行内容", func(c *autogoconsole.Console, args ...any) { c.Println(args...) }, true)
	engine.RegisterMethod("console.clear", "清空控制台", func(c *autogoconsole.Console) { c.Clear() }, true)
	engine.RegisterMethod("console.show", "显示控制台", func(c *autogoconsole.Console) { c.Show() }, true)
	engine.RegisterMethod("console.hide", "隐藏控制台", func(c *autogoconsole.Console) { c.Hide() }, true)
	engine.RegisterMethod("console.isVisible", "返回控制台是否可见", func(c *autogoconsole.Console) bool { return c.IsVisible() }, true)
	engine.RegisterMethod("console.destroy", "销毁控制台", func(c *autogoconsole.Console) { c.Destroy() }, true)
	return nil
}
