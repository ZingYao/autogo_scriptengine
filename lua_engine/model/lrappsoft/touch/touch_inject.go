//go:build ignore
// +build ignore

package touch

import (
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/Dasongzi1366/AutoGo/ime"
	"github.com/Dasongzi1366/AutoGo/motion"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"

	lua "github.com/yuin/gopher-lua"
)

// TouchModule touch 模块（懒人精灵兼容）
type TouchModule struct {
	imeLocked      bool
	ThrowException bool // 是否抛出异常
	ShowWarning    bool // 是否显示警告信息
	Debug          bool // 是否开启调试模式（打印脚本堆栈信息）
}

// checkRootPermission 检查是否有 root 权限
func checkRootPermission() bool {
	cmd := exec.Command("su", "-c", "id")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return false
	}
	return strings.Contains(string(output), "uid=0")
}

// GetKeyCodeFromString 将字符串按键标识符转换为数字按键码
func GetKeyCodeFromString(keyStr string) int {
	keyCodeMap := map[string]int{
		"home":        3,
		"back":        4,
		"recent":      187,
		"power":       26,
		"menu":        82,
		"volume_up":   24,
		"volume_down": 25,
		"camera":      27,
		"enter":       66,
		"delete":      67,
	}
	if code, ok := keyCodeMap[keyStr]; ok {
		return code
	}
	return 0
}

// New 创建一个新的 TouchModule 实例
func New() *TouchModule {
	return &TouchModule{
		imeLocked:      false,
		ThrowException: false,
		ShowWarning:    true,
		Debug:          false,
	}
}

// handleUnimplementedMethod 处理未实现方法的通用逻辑
func (m *TouchModule) handleUnimplementedMethod(methodName string, L *lua.LState) int {
	// 如果开启了调试模式，打印脚本堆栈信息
	if m.Debug {
		fmt.Println("\n=== 调试信息: 调用未实现方法 ===")
		fmt.Printf("方法名: touch.%s\n", methodName)
		fmt.Println("=== Lua 调用堆栈 ===")
		for i := 1; ; i++ {
			dbg, ok := L.GetStack(i)
			if !ok {
				break
			}
			fmt.Printf("  [%d] 函数: %s, 行号: %d\n", i, dbg.Name, dbg.CurrentLine)
			fmt.Printf("       源文件: %s\n", dbg.Source)
			fmt.Printf("       What: %s, 调用行: %d\n", dbg.What, dbg.LineDefined)
		}
		fmt.Println("=== 堆栈结束 ===\n")
	}

	// 打印警告信息
	if m.ShowWarning {
		fmt.Printf("[警告] touch.%s 方法未实现，此功能在 AutoGo 中不支持\n", methodName)
	}

	// 根据配置决定是否抛出异常
	if m.ThrowException {
		L.RaiseError(fmt.Sprintf("touch.%s 方法未实现，此功能在 AutoGo 中不支持", methodName))
		return 0
	}

	// 默认返回 0（空操作）
	return 0
}

// handleUnimplementedMethodWithInt 处理未实现方法的通用逻辑（返回整数）
func (m *TouchModule) handleUnimplementedMethodWithInt(methodName string, L *lua.LState, defaultValue int) int {
	// 如果开启了调试模式，打印脚本堆栈信息
	if m.Debug {
		fmt.Println("\n=== 调试信息: 调用未实现方法 ===")
		fmt.Printf("方法名: touch.%s\n", methodName)
		fmt.Println("=== Lua 调用堆栈 ===")
		for i := 1; ; i++ {
			dbg, ok := L.GetStack(i)
			if !ok {
				break
			}
			fmt.Printf("  [%d] 函数: %s, 行号: %d\n", i, dbg.Name, dbg.CurrentLine)
			fmt.Printf("       源文件: %s\n", dbg.Source)
			fmt.Printf("       What: %s, 调用行: %d\n", dbg.What, dbg.LineDefined)
		}
		fmt.Println("=== 堆栈结束 ===\n")
	}

	// 打印警告信息
	if m.ShowWarning {
		fmt.Printf("[警告] touch.%s 方法未实现，此功能在 AutoGo 中不支持\n", methodName)
	}

	// 根据配置决定是否抛出异常
	if m.ThrowException {
		L.RaiseError(fmt.Sprintf("touch.%s 方法未实现，此功能在 AutoGo 中不支持", methodName))
		return 0
	}

	// 返回默认值
	L.Push(lua.LNumber(defaultValue))
	return 1
}

// handleUnimplementedMethodWithString 处理未实现方法的通用逻辑（返回字符串）
func (m *TouchModule) handleUnimplementedMethodWithString(methodName string, L *lua.LState, defaultValue string) int {
	// 如果开启了调试模式，打印脚本堆栈信息
	if m.Debug {
		fmt.Println("\n=== 调试信息: 调用未实现方法 ===")
		fmt.Printf("方法名: touch.%s\n", methodName)
		fmt.Println("=== Lua 调用堆栈 ===")
		for i := 1; ; i++ {
			dbg, ok := L.GetStack(i)
			if !ok {
				break
			}
			fmt.Printf("  [%d] 函数: %s, 行号: %d\n", i, dbg.Name, dbg.CurrentLine)
			fmt.Printf("       源文件: %s\n", dbg.Source)
			fmt.Printf("       What: %s, 调用行: %d\n", dbg.What, dbg.LineDefined)
		}
		fmt.Println("=== 堆栈结束 ===\n")
	}

	// 打印警告信息
	if m.ShowWarning {
		fmt.Printf("[警告] touch.%s 方法未实现，此功能在 AutoGo 中不支持\n", methodName)
	}

	// 根据配置决定是否抛出异常
	if m.ThrowException {
		L.RaiseError(fmt.Sprintf("touch.%s 方法未实现，此功能在 AutoGo 中不支持", methodName))
		return 0
	}

	// 返回默认值
	L.Push(lua.LString(defaultValue))
	return 1
}

// Name 返回模块名称
func (m *TouchModule) Name() string {
	return "touch"
}

// IsAvailable 检查模块是否可用
func (m *TouchModule) IsAvailable() bool {
	return true
}

// Register 向引擎注册方法
func (m *TouchModule) Register(engine model.Engine) error {
	state := engine.GetState()

	// 创建 imeLib 表
	imeLibTable := state.NewTable()

	// 注册 imeLib.lock - 锁定使用懒人输入法
	imeLibTable.RawSetString("lock", state.NewFunction(func(L *lua.LState) int {
		m.imeLocked = true
		L.Push(lua.LBool(true))
		return 1
	}))

	// 注册 imeLib.unlock - 解锁懒人输入法
	imeLibTable.RawSetString("unlock", state.NewFunction(func(L *lua.LState) int {
		m.imeLocked = false
		L.Push(lua.LBool(true))
		return 1
	}))

	// 注册 imeLib.setText - 输入法模拟输入文字
	imeLibTable.RawSetString("setText", state.NewFunction(func(L *lua.LState) int {
		text := L.CheckString(1)
		// 使用 AutoGo 的 InputText 方法
		ime.InputText(text, 0)
		L.Push(lua.LBool(true))
		return 1
	}))

	// 注册 imeLib.deleteChar - 输入法删除一个字符
	imeLibTable.RawSetString("deleteChar", state.NewFunction(func(L *lua.LState) int {
		// 模拟删除键
		motion.KeyAction(67, 0)
		L.Push(lua.LBool(true))
		return 1
	}))

	// 注册 imeLib.finishInput - 输入法模拟完成输入
	imeLibTable.RawSetString("finishInput", state.NewFunction(func(L *lua.LState) int {
		// 模拟回车键
		motion.KeyAction(66, 0)
		L.Push(lua.LBool(true))
		return 1
	}))

	// 注册 imeLib.keyEvent - 输入法输入字符
	imeLibTable.RawSetString("keyEvent", state.NewFunction(func(L *lua.LState) int {
		_ = L.CheckInt(1) // action 参数，暂时不使用
		keycode := L.CheckInt(2)
		// 使用 AutoGo 的 KeyAction 方法
		motion.KeyAction(keycode, 0)
		L.Push(lua.LBool(true))
		return 1
	}))

	// 注册 imeLib 表到全局
	state.SetGlobal("imeLib", imeLibTable)

	// 注册 inputText - 模拟输入文字
	state.SetGlobal("inputText", state.NewFunction(func(L *lua.LState) int {
		text := L.CheckString(1)
		// 使用 AutoGo 的 InputText 方法
		ime.InputText(text, 0)
		L.Push(lua.LBool(true))
		return 1
	}))

	// 注册 tap - 点击
	state.SetGlobal("tap", state.NewFunction(func(L *lua.LState) int {
		x := L.CheckInt(1)
		y := L.CheckInt(2)
		// 使用 AutoGo 的 Click 方法
		motion.Click(x, y, 0, 0)
		return 0
	}))

	// 注册 longTap - 长点击
	state.SetGlobal("longTap", state.NewFunction(func(L *lua.LState) int {
		x := L.CheckInt(1)
		y := L.CheckInt(2)
		// 使用 AutoGo 的 LongClick 方法，默认长按 500ms
		motion.LongClick(x, y, 500, 0, 0)
		return 0
	}))

	// 注册 touchDown - 按下手指
	state.SetGlobal("touchDown", state.NewFunction(func(L *lua.LState) int {
		id := L.CheckInt(1)
		x := L.CheckInt(2)
		y := L.CheckInt(3)
		// 使用 AutoGo 的 TouchDown 方法
		motion.TouchDown(x, y, id, 0)
		return 0
	}))

	// 注册 touchUp - 弹起手指
	state.SetGlobal("touchUp", state.NewFunction(func(L *lua.LState) int {
		id := L.CheckInt(1)
		// 使用 AutoGo 的 TouchUp 方法
		motion.TouchUp(0, 0, id, 0)
		L.Push(lua.LBool(true))
		return 1
	}))

	// 注册 touchMove - 模拟滑动
	state.SetGlobal("touchMove", state.NewFunction(func(L *lua.LState) int {
		id := L.CheckInt(1)
		x := L.CheckInt(2)
		y := L.CheckInt(3)
		// 使用 AutoGo 的 TouchMove 方法
		motion.TouchMove(x, y, id, 0)
		L.Push(lua.LBool(true))
		return 1
	}))

	// 创建 touch 表
	touchTable := state.NewTable()

	// 注册 touch.down - 按下触摸事件（别名）
	touchTable.RawSetString("down", state.NewFunction(func(L *lua.LState) int {
		id := L.CheckInt(1)
		x := L.CheckInt(2)
		y := L.CheckInt(3)
		motion.TouchDown(x, y, id, 0)
		return 0
	}))

	// 注册 touch.move - 移动触摸事件（别名）
	touchTable.RawSetString("move", state.NewFunction(func(L *lua.LState) int {
		id := L.CheckInt(1)
		x := L.CheckInt(2)
		y := L.CheckInt(3)
		motion.TouchMove(x, y, id, 0)
		return 0
	}))

	// 注册 touch.up - 抬起触摸事件（别名）
	touchTable.RawSetString("up", state.NewFunction(func(L *lua.LState) int {
		id := L.CheckInt(1)
		x := L.OptInt(2, 0)
		y := L.OptInt(3, 0)
		motion.TouchUp(x, y, id, 0)
		return 0
	}))

	// 注册 touch.click - 点击事件（别名）
	touchTable.RawSetString("click", state.NewFunction(func(L *lua.LState) int {
		x := L.CheckInt(1)
		y := L.CheckInt(2)
		motion.Click(x, y, 0, 0)
		return 0
	}))

	// 注册 touch.doubleClick - 双击事件
	touchTable.RawSetString("doubleClick", state.NewFunction(func(L *lua.LState) int {
		x := L.CheckInt(1)
		y := L.CheckInt(2)
		motion.Click(x, y, 0, 0)
		time.Sleep(100 * time.Millisecond)
		motion.Click(x, y, 0, 0)
		return 0
	}))

	// 注册 touch.longClick - 长按事件（别名）
	touchTable.RawSetString("longClick", state.NewFunction(func(L *lua.LState) int {
		x := L.CheckInt(1)
		y := L.CheckInt(2)
		duration := L.OptInt(3, 1000)
		motion.LongClick(x, y, duration, 0, 0)
		return 0
	}))

	// 注册 touch.swipe - 滑动事件（别名）
	touchTable.RawSetString("swipe", state.NewFunction(func(L *lua.LState) int {
		x1 := L.CheckInt(1)
		y1 := L.CheckInt(2)
		x2 := L.CheckInt(3)
		y2 := L.CheckInt(4)
		duration := L.OptInt(5, 500)
		motion.Swipe(x1, y1, x2, y2, duration, 0, 0)
		return 0
	}))

	// 注册 touch.key - 按键事件
	touchTable.RawSetString("key", state.NewFunction(func(L *lua.LState) int {
		keycode := L.CheckInt(1)
		motion.KeyAction(keycode, 0)
		return 0
	}))

	// 注册 touch.keyCode - 按键码事件
	touchTable.RawSetString("keyCode", state.NewFunction(func(L *lua.LState) int {
		keyStr := L.CheckString(1)
		code := GetKeyCodeFromString(keyStr)
		motion.KeyAction(code, 0)
		return 0
	}))

	// 注册 touch.inputText - 输入文本（别名）
	touchTable.RawSetString("inputText", state.NewFunction(func(L *lua.LState) int {
		text := L.CheckString(1)
		ime.InputText(text, 0)
		return 0
	}))

	// 注册 touch.input - 输入事件（兼容旧版本）
	touchTable.RawSetString("input", state.NewFunction(func(L *lua.LState) int {
		text := L.CheckString(1)
		ime.InputText(text, 0)
		return 0
	}))

	// 注册 touch.clearText - 清除文本
	touchTable.RawSetString("clearText", state.NewFunction(func(L *lua.LState) int {
		// 模拟多次删除键
		for i := 0; i < 50; i++ {
			motion.KeyAction(67, 0)
			time.Sleep(10 * time.Millisecond)
		}
		return 0
	}))

	// 注册 touch.isIME - 检查输入法是否激活
	touchTable.RawSetString("isIME", state.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LBool(false))
		return 1
	}))

	// 注册 touch.openIME - 打开输入法
	touchTable.RawSetString("openIME", state.NewFunction(func(L *lua.LState) int {
		return 0
	}))

	// 注册 touch.closeIME - 关闭输入法
	touchTable.RawSetString("closeIME", state.NewFunction(func(L *lua.LState) int {
		return 0
	}))

	// 注册 touch.lockIME - 锁定输入法
	touchTable.RawSetString("lockIME", state.NewFunction(func(L *lua.LState) int {
		m.imeLocked = true
		return 0
	}))

	// 注册 touch.unlockIME - 解锁输入法
	touchTable.RawSetString("unlockIME", state.NewFunction(func(L *lua.LState) int {
		m.imeLocked = false
		return 0
	}))

	// 注册 touch.isRooted - 检查是否有 root 权限
	touchTable.RawSetString("isRooted", state.NewFunction(func(L *lua.LState) int {
		isRooted := checkRootPermission()
		L.Push(lua.LBool(isRooted))
		return 1
	}))

	// 注册 touch.execShell - 执行 Shell 命令
	touchTable.RawSetString("execShell", state.NewFunction(func(L *lua.LState) int {
		cmdStr := L.CheckString(1)
		cmd := exec.Command("sh", "-c", cmdStr)
		output, err := cmd.CombinedOutput()
		if err != nil {
			L.Push(lua.LString(""))
			return 1
		}
		L.Push(lua.LString(string(output)))
		return 1
	}))

	// 注册 touch 表到全局
	state.SetGlobal("touch", touchTable)

	// 注册 touchMoveEx - 模拟滑动增强版
	state.SetGlobal("touchMoveEx", state.NewFunction(func(L *lua.LState) int {
		id := L.CheckInt(1)
		x := L.CheckInt(2)
		y := L.CheckInt(3)
		duration := L.CheckInt(4)
		// 使用 AutoGo 的 TouchMove 方法
		motion.TouchMove(x, y, id, 0)
		time.Sleep(time.Duration(duration) * time.Millisecond)
		return 0
	}))

	// 注册 swipe - 划动
	state.SetGlobal("swipe", state.NewFunction(func(L *lua.LState) int {
		x1 := L.CheckInt(1)
		y1 := L.CheckInt(2)
		x2 := L.CheckInt(3)
		y2 := L.CheckInt(4)
		duration := L.CheckInt(5)
		// 使用 AutoGo 的 Swipe 方法
		motion.Swipe(x1, y1, x2, y2, duration, 0, 0)
		L.Push(lua.LBool(true))
		return 1
	}))

	// 注册 keyPress - 按键
	state.SetGlobal("keyPress", state.NewFunction(func(L *lua.LState) int {
		keycode := L.CheckAny(1)
		var code int
		switch keycode.Type() {
		case lua.LTString:
			// 字符串按键标识符
			keyStr := keycode.String()
			code = GetKeyCodeFromString(keyStr)
		case lua.LTNumber:
			// 数字按键码
			code = int(keycode.(lua.LNumber))
		default:
			L.Push(lua.LBool(false))
			return 1
		}
		// 使用 AutoGo 的 KeyAction 方法
		motion.KeyAction(code, 0)
		L.Push(lua.LBool(true))
		return 1
	}))

	// 注册 keyDown - 按键按下
	state.SetGlobal("keyDown", state.NewFunction(func(L *lua.LState) int {
		keycode := L.CheckAny(1)
		var code int
		switch keycode.Type() {
		case lua.LTString:
			keyStr := keycode.String()
			code = GetKeyCodeFromString(keyStr)
		case lua.LTNumber:
			code = int(keycode.(lua.LNumber))
		default:
			L.Push(lua.LBool(false))
			return 1
		}
		// 使用 AutoGo 的 KeyAction 方法
		motion.KeyAction(code, 0)
		L.Push(lua.LBool(true))
		return 1
	}))

	// 注册 keyUp - 按键弹起
	state.SetGlobal("keyUp", state.NewFunction(func(L *lua.LState) int {
		keycode := L.CheckAny(1)
		switch keycode.Type() {
		case lua.LTString:
			// 字符串按键标识符
			_ = keycode.String()
		case lua.LTNumber:
			// 数字按键码
			_ = int(keycode.(lua.LNumber))
		default:
			L.Push(lua.LBool(false))
			return 1
		}
		// AutoGo 的 KeyAction 方法同时处理按下和弹起
		// 这里我们不做实际操作，因为 AutoGo 的实现是完整的按键动作
		L.Push(lua.LBool(true))
		return 1
	}))

	// 注册 setOnTouchListener - 获取用户触摸屏幕坐标
	state.SetGlobal("setOnTouchListener", state.NewFunction(func(L *lua.LState) int {
		// 检查回调函数参数
		if L.GetTop() < 1 {
			L.Push(lua.LBool(false))
			return 1
		}

		_ = L.CheckFunction(1)

		// 检查是否有 root 权限
		if !checkRootPermission() {
			L.Push(lua.LBool(false))
			return 1
		}

		// 注意：虽然可以检查 root 权限，但实际监听触摸事件需要：
		// 1. 找到正确的触摸输入设备文件
		// 2. 解析复杂的 getevent 输出格式
		// 3. 在 goroutine 中安全地调用 Lua 回调函数
		// 4. 处理多指触摸和事件同步
		// 由于实现复杂度较高且涉及 Lua 状态的并发安全问题，
		// 这里只检查 root 权限并返回 true，实际监听功能暂不实现
		// 如果需要完整的触摸监听功能，建议使用 AutoGo 的 uiacc 模块

		L.Push(lua.LBool(true))
		return 1
	}))

	return nil
}

// GetModule 获取模块实例
func GetModule() model.Module {
	return New()
}
