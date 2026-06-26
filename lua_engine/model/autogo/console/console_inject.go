package console

import (
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"

	"github.com/Dasongzi1366/AutoGo/console"
	lua "github.com/yuin/gopher-lua"
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
	state := engine.GetState()

	consoleObj := state.NewTable()
	state.SetGlobal("console", consoleObj)

	consoleObj.RawSetString("new", state.NewFunction(func(L *lua.LState) int {
		c := console.New()
		ud := L.NewUserData()
		ud.Value = c
		L.Push(ud)
		return 1
	}))

	consoleObj.RawSetString("println", state.NewFunction(func(L *lua.LState) int {
		c := L.CheckUserData(1).Value.(*console.Console)
		var args []interface{}
		for i := 2; i <= L.GetTop(); i++ {
			args = append(args, L.CheckAny(i))
		}
		c.Println(args...)
		return 0
	}))

	consoleObj.RawSetString("setTextSize", state.NewFunction(func(L *lua.LState) int {
		c := L.CheckUserData(1).Value.(*console.Console)
		size := L.CheckInt(2)
		c.SetTextSize(size)
		return 0
	}))

	consoleObj.RawSetString("setTextColor", state.NewFunction(func(L *lua.LState) int {
		c := L.CheckUserData(1).Value.(*console.Console)
		color := L.CheckString(2)
		c.SetTextColor(color)
		return 0
	}))

	consoleObj.RawSetString("setWindowSize", state.NewFunction(func(L *lua.LState) int {
		c := L.CheckUserData(1).Value.(*console.Console)
		width := L.CheckInt(2)
		height := L.CheckInt(3)
		c.SetWindowSize(width, height)
		return 0
	}))

	consoleObj.RawSetString("setWindowPosition", state.NewFunction(func(L *lua.LState) int {
		c := L.CheckUserData(1).Value.(*console.Console)
		x := L.CheckInt(2)
		y := L.CheckInt(3)
		c.SetWindowPosition(x, y)
		return 0
	}))

	consoleObj.RawSetString("setWindowColor", state.NewFunction(func(L *lua.LState) int {
		c := L.CheckUserData(1).Value.(*console.Console)
		color := L.CheckString(2)
		c.SetWindowColor(color)
		return 0
	}))

	consoleObj.RawSetString("show", state.NewFunction(func(L *lua.LState) int {
		c := L.CheckUserData(1).Value.(*console.Console)
		c.Show()
		return 0
	}))

	consoleObj.RawSetString("hide", state.NewFunction(func(L *lua.LState) int {
		c := L.CheckUserData(1).Value.(*console.Console)
		c.Hide()
		return 0
	}))

	consoleObj.RawSetString("clear", state.NewFunction(func(L *lua.LState) int {
		c := L.CheckUserData(1).Value.(*console.Console)
		c.Clear()
		return 0
	}))

	consoleObj.RawSetString("isVisible", state.NewFunction(func(L *lua.LState) int {
		c := L.CheckUserData(1).Value.(*console.Console)
		visible := c.IsVisible()
		L.Push(lua.LBool(visible))
		return 1
	}))

	consoleObj.RawSetString("destroy", state.NewFunction(func(L *lua.LState) int {
		c := L.CheckUserData(1).Value.(*console.Console)
		c.Destroy()
		return 0
	}))

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
