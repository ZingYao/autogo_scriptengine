//go:build ignore
// +build ignore

package console

import (
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"

	autogoconsole "github.com/Dasongzi1366/AutoGo/console"
	lua "github.com/yuin/gopher-lua"
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

// luaValuesToAnySlice 将 Lua 可变参数转换为 Println 可接受的参数切片。
func luaValuesToAnySlice(L *lua.LState, start int) []any {
	args := make([]any, 0, L.GetTop()-start+1)
	for index := start; index <= L.GetTop(); index++ {
		args = append(args, L.CheckAny(index))
	}
	return args
}

// consoleChainValue 返回链式调用对象。
func consoleChainValue(L *lua.LState, current lua.LValue, original, next *autogoconsole.Console) lua.LValue {
	if next != nil && next != original {
		return wrapConsole(L, next)
	}
	return current
}

// wrapConsole 将 Go Console 对象包装为 Lua 实例对象。
func wrapConsole(L *lua.LState, c *autogoconsole.Console) lua.LValue {
	obj := L.NewTable()
	obj.RawSetString("setWindowSize", L.NewFunction(func(L *lua.LState) int {
		L.Push(consoleChainValue(L, obj, c, c.SetWindowSize(L.CheckInt(1), L.CheckInt(2))))
		return 1
	}))
	obj.RawSetString("setWindowPosition", L.NewFunction(func(L *lua.LState) int {
		L.Push(consoleChainValue(L, obj, c, c.SetWindowPosition(L.CheckInt(1), L.CheckInt(2))))
		return 1
	}))
	obj.RawSetString("setWindowColor", L.NewFunction(func(L *lua.LState) int {
		L.Push(consoleChainValue(L, obj, c, c.SetWindowColor(L.CheckString(1))))
		return 1
	}))
	obj.RawSetString("setTextColor", L.NewFunction(func(L *lua.LState) int {
		L.Push(consoleChainValue(L, obj, c, c.SetTextColor(L.CheckString(1))))
		return 1
	}))
	obj.RawSetString("setTextSize", L.NewFunction(func(L *lua.LState) int {
		L.Push(consoleChainValue(L, obj, c, c.SetTextSize(L.CheckInt(1))))
		return 1
	}))
	obj.RawSetString("println", L.NewFunction(func(L *lua.LState) int {
		c.Println(luaValuesToAnySlice(L, 1)...)
		return 0
	}))
	obj.RawSetString("clear", L.NewFunction(func(L *lua.LState) int {
		c.Clear()
		return 0
	}))
	obj.RawSetString("show", L.NewFunction(func(L *lua.LState) int {
		c.Show()
		return 0
	}))
	obj.RawSetString("hide", L.NewFunction(func(L *lua.LState) int {
		c.Hide()
		return 0
	}))
	obj.RawSetString("isVisible", L.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LBool(c.IsVisible()))
		return 1
	}))
	obj.RawSetString("destroy", L.NewFunction(func(L *lua.LState) int {
		c.Destroy()
		return 0
	}))
	return obj
}

// Register 向引擎注册 iOS console 方法。
func (m *ConsoleModule) Register(engine model.Engine) error {
	state := engine.GetState()
	consoleObj := state.NewTable()
	state.SetGlobal("console", consoleObj)
	consoleObj.RawSetString("new", state.NewFunction(func(L *lua.LState) int {
		c := autogoconsole.New()
		if c == nil {
			L.Push(lua.LNil)
			return 1
		}
		L.Push(wrapConsole(L, c))
		return 1
	}))

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
