package utils

import (
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"

	"github.com/Dasongzi1366/AutoGo/utils"
	lua "github.com/yuin/gopher-lua"
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
	state := engine.GetState()

	utilsObj := state.NewTable()
	state.SetGlobal("utils", utilsObj)

	utilsObj.RawSetString("logI", state.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		message := ""
		for i := 2; i <= L.GetTop(); i++ {
			message += L.CheckString(i) + " "
		}
		utils.LogI(label, message)
		return 0
	}))

	utilsObj.RawSetString("logE", state.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		message := ""
		for i := 2; i <= L.GetTop(); i++ {
			message += L.CheckString(i) + " "
		}
		utils.LogE(label, message)
		return 0
	}))

	utilsObj.RawSetString("toast", state.NewFunction(func(L *lua.LState) int {
		message := L.CheckString(1)
		x := 0
		y := 0
		duration := -1
		if L.GetTop() > 1 {
			x = L.CheckInt(2)
		}
		if L.GetTop() > 2 {
			y = L.CheckInt(3)
		}
		if L.GetTop() > 3 {
			duration = L.CheckInt(4)
		}
		utils.Toast(message, x, y, duration)
		return 0
	}))

	utilsObj.RawSetString("alert", state.NewFunction(func(L *lua.LState) int {
		title := L.CheckString(1)
		content := L.CheckString(2)
		btn1Text := ""
		if L.GetTop() > 2 {
			btn1Text = L.CheckString(3)
		}
		btn2Text := ""
		if L.GetTop() > 3 {
			btn2Text = L.CheckString(4)
		}
		result := utils.Alert(title, content, btn1Text, btn2Text)
		L.Push(lua.LNumber(result))
		return 1
	}))

	utilsObj.RawSetString("shell", state.NewFunction(func(L *lua.LState) int {
		cmd := L.CheckString(1)
		result := utils.Shell(cmd)
		L.Push(lua.LString(result))
		return 1
	}))

	utilsObj.RawSetString("random", state.NewFunction(func(L *lua.LState) int {
		min := L.CheckInt(1)
		max := L.CheckInt(2)
		result := utils.Random(min, max)
		L.Push(lua.LNumber(result))
		return 1
	}))

	utilsObj.RawSetString("sleep", state.NewFunction(func(L *lua.LState) int {
		i := L.CheckInt(1)
		utils.Sleep(i)
		return 0
	}))

	utilsObj.RawSetString("i2s", state.NewFunction(func(L *lua.LState) int {
		i := L.CheckInt(1)
		result := utils.I2s(i)
		L.Push(lua.LString(result))
		return 1
	}))

	utilsObj.RawSetString("s2i", state.NewFunction(func(L *lua.LState) int {
		s := L.CheckString(1)
		result := utils.S2i(s)
		L.Push(lua.LNumber(result))
		return 1
	}))

	utilsObj.RawSetString("f2s", state.NewFunction(func(L *lua.LState) int {
		f := L.CheckNumber(1)
		result := utils.F2s(float64(f))
		L.Push(lua.LString(result))
		return 1
	}))

	utilsObj.RawSetString("s2f", state.NewFunction(func(L *lua.LState) int {
		s := L.CheckString(1)
		result := utils.S2f(s)
		L.Push(lua.LNumber(result))
		return 1
	}))

	utilsObj.RawSetString("b2s", state.NewFunction(func(L *lua.LState) int {
		b := L.CheckBool(1)
		result := utils.B2s(b)
		L.Push(lua.LString(result))
		return 1
	}))

	utilsObj.RawSetString("s2b", state.NewFunction(func(L *lua.LState) int {
		s := L.CheckString(1)
		result := utils.S2b(s)
		L.Push(lua.LBool(result))
		return 1
	}))

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

	return nil
}
