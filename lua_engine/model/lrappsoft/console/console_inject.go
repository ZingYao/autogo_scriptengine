//go:build ignore
// +build ignore

package console

import (
	"fmt"
	"strconv"

	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"

	lua "github.com/yuin/gopher-lua"
)

// ConsoleModule console 模块（懒人精灵兼容）
type ConsoleModule struct {
	consoleInstance interface{}
	titleVisible    bool
	locked          bool
}

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

	// 创建 console 表
	consoleTable := state.NewTable()

	// 注册 console.show - 显示控制台悬浮窗
	consoleTable.RawSetString("show", state.NewFunction(func(L *lua.LState) int {
		m.consoleShow(L)
		L.Push(lua.LBool(true))
		return 1
	}))

	// 注册 console.showTitle - 显示或者隐藏控制台标题栏
	consoleTable.RawSetString("showTitle", state.NewFunction(func(L *lua.LState) int {
		show := true
		if L.GetTop() >= 1 {
			show = L.CheckBool(1)
		}
		m.titleVisible = show
		fmt.Printf("[console] 标题栏显示状态: %v\n", show)
		L.Push(lua.LBool(true))
		return 1
	}))

	// 注册 console.lockConsole - 锁定控制台窗口
	consoleTable.RawSetString("lockConsole", state.NewFunction(func(L *lua.LState) int {
		m.locked = true
		fmt.Println("[console] 控制台窗口已锁定")
		return 0
	}))

	// 注册 console.unlockConsole - 解除锁定控制台窗口
	consoleTable.RawSetString("unlockConsole", state.NewFunction(func(L *lua.LState) int {
		m.locked = false
		fmt.Println("[console] 控制台窗口已解除锁定")
		return 0
	}))

	// 注册 console.dismiss - 关闭控制台窗口
	consoleTable.RawSetString("dismiss", state.NewFunction(func(L *lua.LState) int {
		m.consoleDismiss(L)
		L.Push(lua.LBool(true))
		return 1
	}))

	// 注册 console.setPos - 设置控制台窗口的位置和大小
	consoleTable.RawSetString("setPos", state.NewFunction(func(L *lua.LState) int {
		x := L.CheckInt(1)
		y := L.CheckInt(2)
		width := 0
		height := 0

		if L.GetTop() >= 3 {
			width = L.CheckInt(3)
		}
		if L.GetTop() >= 4 {
			height = L.CheckInt(4)
		}

		m.consoleSetPos(L, x, y, width, height)
		return 0
	}))

	// 注册 console.println - 打印日志到控制台窗口
	consoleTable.RawSetString("println", state.NewFunction(func(L *lua.LState) int {
		level := 0
		if L.GetTop() >= 1 {
			level = L.CheckInt(1)
		}

		log := ""
		if L.GetTop() >= 2 {
			log = L.CheckString(2)
		}

		m.consolePrintln(L, level, log)
		return 0
	}))

	// 注册 console.clearLog - 清除日志
	consoleTable.RawSetString("clearLog", state.NewFunction(func(L *lua.LState) int {
		m.consoleClearLog(L)
		return 0
	}))

	// 注册 console.setTitle - 设置控制台标题
	consoleTable.RawSetString("setTitle", state.NewFunction(func(L *lua.LState) int {
		title := ""
		if L.GetTop() >= 1 {
			title = L.CheckString(1)
		}
		m.consoleSetTitle(L, title)
		return 0
	}))

	// 注册 console 到全局
	state.SetGlobal("console", consoleTable)

	// 注册到方法注册表
	engine.RegisterMethod("console.show", "显示控制台悬浮窗", func() bool {
		return true
	}, true)
	engine.RegisterMethod("console.showTitle", "显示或者隐藏控制台标题栏", func(args ...interface{}) bool {
		show := true
		if len(args) > 0 {
			show = consoleBoolArg(args[0])
		}
		m.titleVisible = show
		return true
	}, true)
	engine.RegisterMethod("console.lockConsole", "锁定控制台窗口", func() {
		m.locked = true
	}, true)
	engine.RegisterMethod("console.unlockConsole", "解除锁定控制台窗口", func() {
		m.locked = false
	}, true)
	engine.RegisterMethod("console.dismiss", "关闭控制台窗口", func() bool {
		return true
	}, true)
	engine.RegisterMethod("console.setPos", "设置控制台窗口的位置和大小", func(args ...interface{}) {
		_ = consoleIntArg(args, 0)
		_ = consoleIntArg(args, 1)
		_ = consoleIntArg(args, 2)
		_ = consoleIntArg(args, 3)
	}, true)
	engine.RegisterMethod("console.println", "打印日志到控制台窗口", func(args ...interface{}) {
		_ = consoleIntArg(args, 0)
		_ = consoleStringArg(args, 1)
	}, true)
	engine.RegisterMethod("console.clearLog", "清除日志", func() {
	}, true)
	engine.RegisterMethod("console.setTitle", "设置控制台标题", func(title string) {
	}, true)

	return nil
}

// consoleShow 显示控制台悬浮窗
func (m *ConsoleModule) consoleShow(L *lua.LState) {
	fmt.Println("[console] 显示控制台悬浮窗")
}

// consoleDismiss 关闭控制台窗口
func (m *ConsoleModule) consoleDismiss(L *lua.LState) {
	fmt.Println("[console] 关闭控制台窗口")
}

// consoleSetPos 设置控制台窗口的位置和大小
func (m *ConsoleModule) consoleSetPos(L *lua.LState, x, y, width, height int) {
	fmt.Printf("[console] 设置位置: x=%d, y=%d, width=%d, height=%d\n", x, y, width, height)
}

// consolePrintln 打印日志到控制台窗口
func (m *ConsoleModule) consolePrintln(L *lua.LState, level int, log string) {
	fmt.Printf("[console][level=%d] %s\n", level, log)
}

// consoleClearLog 清除日志
func (m *ConsoleModule) consoleClearLog(L *lua.LState) {
	fmt.Println("[console] 清除日志")
}

// consoleSetTitle 设置控制台标题
func (m *ConsoleModule) consoleSetTitle(L *lua.LState, title string) {
	fmt.Printf("[console] 设置标题: %s\n", title)
}

func consoleBoolArg(value interface{}) bool {
	switch typedValue := value.(type) {
	case bool:
		return typedValue
	case string:
		return typedValue == "true" || typedValue == "1"
	default:
		return false
	}
}

func consoleIntArg(args []interface{}, index int) int {
	if index >= len(args) {
		return 0
	}
	switch value := args[index].(type) {
	case int:
		return value
	case int64:
		return int(value)
	case float64:
		return int(value)
	case string:
		parsed, _ := strconv.Atoi(value)
		return parsed
	default:
		return 0
	}
}

func consoleStringArg(args []interface{}, index int) string {
	if index >= len(args) {
		return ""
	}
	return fmt.Sprint(args[index])
}

// GetModule 获取模块实例
func GetModule() model.Module {
	return &ConsoleModule{}
}
