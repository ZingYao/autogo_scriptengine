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

	return nil
}
