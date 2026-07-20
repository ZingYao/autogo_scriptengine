package console

import (
	"github.com/Dasongzi1366/AutoGo/console"
	"github.com/ZingYao/autogo_scriptengine/js_engine/model"
	"github.com/ZingYao/goja"
)

// ConsoleModule console 模块
type ConsoleModule struct{}

// Name 返回模块名称
func (m *ConsoleModule) Name() string {
	return "console"
}

// IsAvailable 检查模块是否可用
func (m *ConsoleModule) IsAvailable() bool {
	return true
}

// Register 向引擎注册方法
func (m *ConsoleModule) Register(engine model.Engine) error {
	vm := engine.GetVM()

	consoleObj := vm.NewObject()
	vm.Set("console", consoleObj)

	consoleObj.Set("new", func(call goja.FunctionCall) goja.Value {
		c := console.New()
		return vm.ToValue(c)
	})

	consoleObj.Set("println", func(call goja.FunctionCall) goja.Value {
		c := call.Argument(0).Export().(*console.Console)
		var args []any
		for i := 1; i < len(call.Arguments); i++ {
			args = append(args, call.Argument(i).Export())
		}
		c.Println(args...)
		return goja.Undefined()
	})

	consoleObj.Set("setTextSize", func(call goja.FunctionCall) goja.Value {
		c := call.Argument(0).Export().(*console.Console)
		size := int(call.Argument(1).ToInteger())
		c.SetTextSize(size)
		return goja.Undefined()
	})

	consoleObj.Set("setTextColor", func(call goja.FunctionCall) goja.Value {
		c := call.Argument(0).Export().(*console.Console)
		color := call.Argument(1).String()
		c.SetTextColor(color)
		return goja.Undefined()
	})

	consoleObj.Set("setWindowSize", func(call goja.FunctionCall) goja.Value {
		c := call.Argument(0).Export().(*console.Console)
		width := int(call.Argument(1).ToInteger())
		height := int(call.Argument(2).ToInteger())
		c.SetWindowSize(width, height)
		return goja.Undefined()
	})

	consoleObj.Set("setWindowPosition", func(call goja.FunctionCall) goja.Value {
		c := call.Argument(0).Export().(*console.Console)
		x := int(call.Argument(1).ToInteger())
		y := int(call.Argument(2).ToInteger())
		c.SetWindowPosition(x, y)
		return goja.Undefined()
	})

	consoleObj.Set("setWindowColor", func(call goja.FunctionCall) goja.Value {
		c := call.Argument(0).Export().(*console.Console)
		color := call.Argument(1).String()
		c.SetWindowColor(color)
		return goja.Undefined()
	})

	consoleObj.Set("show", func(call goja.FunctionCall) goja.Value {
		c := call.Argument(0).Export().(*console.Console)
		c.Show()
		return goja.Undefined()
	})

	consoleObj.Set("hide", func(call goja.FunctionCall) goja.Value {
		c := call.Argument(0).Export().(*console.Console)
		c.Hide()
		return goja.Undefined()
	})

	consoleObj.Set("clear", func(call goja.FunctionCall) goja.Value {
		c := call.Argument(0).Export().(*console.Console)
		c.Clear()
		return goja.Undefined()
	})

	consoleObj.Set("isVisible", func(call goja.FunctionCall) goja.Value {
		c := call.Argument(0).Export().(*console.Console)
		visible := c.IsVisible()
		return vm.ToValue(visible)
	})

	consoleObj.Set("destroy", func(call goja.FunctionCall) goja.Value {
		c := call.Argument(0).Export().(*console.Console)
		c.Destroy()
		return goja.Undefined()
	})

	engine.RegisterMethod("console.new", "创建控制台对象", console.New, true)
	engine.RegisterMethod("console.println", "输出一行内容", func(c *console.Console, args ...any) { c.Println(args...) }, true)
	engine.RegisterMethod("console.setTextSize", "设置控制台文本大小", func(c *console.Console, size int) *console.Console {
		return c.SetTextSize(size)
	}, true)
	engine.RegisterMethod("console.setTextColor", "设置控制台文本颜色", func(c *console.Console, color string) *console.Console {
		return c.SetTextColor(color)
	}, true)
	engine.RegisterMethod("console.setWindowSize", "设置控制台窗口大小", func(c *console.Console, width, height int) *console.Console {
		return c.SetWindowSize(width, height)
	}, true)
	engine.RegisterMethod("console.setWindowPosition", "设置控制台窗口位置", func(c *console.Console, x, y int) *console.Console {
		return c.SetWindowPosition(x, y)
	}, true)
	engine.RegisterMethod("console.setWindowColor", "设置控制台窗口颜色", func(c *console.Console, color string) *console.Console {
		return c.SetWindowColor(color)
	}, true)
	engine.RegisterMethod("console.show", "显示控制台", func(c *console.Console) { c.Show() }, true)
	engine.RegisterMethod("console.hide", "隐藏控制台", func(c *console.Console) { c.Hide() }, true)
	engine.RegisterMethod("console.clear", "清空控制台", func(c *console.Console) { c.Clear() }, true)
	engine.RegisterMethod("console.isVisible", "返回控制台是否可见", func(c *console.Console) bool { return c.IsVisible() }, true)
	engine.RegisterMethod("console.destroy", "销毁控制台", func(c *console.Console) { c.Destroy() }, true)

	return nil
}
