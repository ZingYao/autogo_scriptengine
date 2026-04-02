package ui

import (
	"fmt"

	"github.com/Dasongzi1366/AutoGo/utils"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"
	lua "github.com/yuin/gopher-lua"
)

// UIModule ui 模块（懒人精灵兼容）
type UIModule struct {
	ThrowException bool // 是否抛出异常
	ShowWarning    bool // 是否显示警告信息
	Debug         bool // 是否开启调试模式（打印脚本堆栈信息）
}

// New 创建一个新的 UIModule 实例
func New() *UIModule {
	return &UIModule{
		ThrowException: false,
		ShowWarning:    true,
		Debug:         false,
	}
}

// handleEmptyMethod 处理空方法的通用逻辑
func (m *UIModule) handleEmptyMethod(methodName string, L *lua.LState) int {
	// 如果开启了调试模式，打印脚本堆栈信息
	if m.Debug {
		fmt.Println("\n=== 调试信息: 调用未实现方法 ===")
		fmt.Printf("方法名: ui.%s\n", methodName)
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
		fmt.Printf("[警告] ui.%s 方法未实现，此功能在 AutoGo 中不支持\n", methodName)
	}

	// 根据配置决定是否抛出异常
	if m.ThrowException {
		L.RaiseError(fmt.Sprintf("ui.%s 方法未实现，此功能在 AutoGo 中不支持", methodName))
		return 0
	}

	// 默认返回 0（无返回值）
	return 0
}

// handleEmptyMethodWithBool 处理空方法的通用逻辑（返回布尔值）
func (m *UIModule) handleEmptyMethodWithBool(methodName string, L *lua.LState, defaultValue bool) int {
	// 如果开启了调试模式，打印脚本堆栈信息
	if m.Debug {
		fmt.Println("\n=== 调试信息: 调用未实现方法 ===")
		fmt.Printf("方法名: ui.%s\n", methodName)
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
		fmt.Printf("[警告] ui.%s 方法未实现，此功能在 AutoGo 中不支持\n", methodName)
	}

	// 根据配置决定是否抛出异常
	if m.ThrowException {
		L.RaiseError(fmt.Sprintf("ui.%s 方法未实现，此功能在 AutoGo 中不支持", methodName))
		return 0
	}

	// 返回默认值
	L.Push(lua.LBool(defaultValue))
	return 1
}

// handleEmptyMethodWithNumber 处理空方法的通用逻辑（返回数字）
func (m *UIModule) handleEmptyMethodWithNumber(methodName string, L *lua.LState, defaultValue int) int {
	// 如果开启了调试模式，打印脚本堆栈信息
	if m.Debug {
		fmt.Println("\n=== 调试信息: 调用未实现方法 ===")
		fmt.Printf("方法名: ui.%s\n", methodName)
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
		fmt.Printf("[警告] ui.%s 方法未实现，此功能在 AutoGo 中不支持\n", methodName)
	}

	// 根据配置决定是否抛出异常
	if m.ThrowException {
		L.RaiseError(fmt.Sprintf("ui.%s 方法未实现，此功能在 AutoGo 中不支持", methodName))
		return 0
	}

	// 返回默认值
	L.Push(lua.LNumber(defaultValue))
	return 1
}

// handleEmptyMethodWithString 处理空方法的通用逻辑（返回字符串）
func (m *UIModule) handleEmptyMethodWithString(methodName string, L *lua.LState, defaultValue string) int {
	// 如果开启了调试模式，打印脚本堆栈信息
	if m.Debug {
		fmt.Println("\n=== 调试信息: 调用未实现方法 ===")
		fmt.Printf("方法名: ui.%s\n", methodName)
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
		fmt.Printf("[警告] ui.%s 方法未实现，此功能在 AutoGo 中不支持\n", methodName)
	}

	// 根据配置决定是否抛出异常
	if m.ThrowException {
		L.RaiseError(fmt.Sprintf("ui.%s 方法未实现，此功能在 AutoGo 中不支持", methodName))
		return 0
	}

	// 返回默认值
	L.Push(lua.LString(defaultValue))
	return 1
}

// Name 返回模块名称
func (m *UIModule) Name() string {
	return "ui"
}

// IsAvailable 检查模块是否可用
func (m *UIModule) IsAvailable() bool {
	return true
}

// Register 向引擎注册方法
func (m *UIModule) Register(engine model.Engine) error {
	state := engine.GetState()

	// 注册 closeWindow - 关闭窗口
	state.SetGlobal("closeWindow", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("closeWindow", L)
	}))

	// 注册 getUIWebViewUrl - 获取当前浏览器的地址
	state.SetGlobal("getUIWebViewUrl", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethodWithString("getUIWebViewUrl", L, "")
	}))

	// 注册 setUIWebViewUrl - 设置浏览器地址
	state.SetGlobal("setUIWebViewUrl", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("setUIWebViewUrl", L)
	}))

	// 注册 getUISelected - 获取单选框或者下拉框当前选中项
	state.SetGlobal("getUISelected", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethodWithNumber("getUISelected", L, 0)
	}))

	// 注册 getUISelectText - 获取单选框或者下拉框当前选中项的文本内容
	state.SetGlobal("getUISelectText", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethodWithString("getUISelectText", L, "")
	}))

	// 注册 getUIText - 获取控件显示的文本内容
	state.SetGlobal("getUIText", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethodWithString("getUIText", L, "")
	}))

	// 注册 setUIText - 设置控件文本
	state.SetGlobal("setUIText", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("setUIText", L)
	}))

	// 注册 setUISelect - 设置单选框或者下拉框选项被选中
	state.SetGlobal("setUISelect", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("setUISelect", L)
	}))

	// 注册 setUICheck - 设置多选框被选中或者反选
	state.SetGlobal("setUICheck", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("setUICheck", L)
	}))

	// 注册 getUICheck - 获取多选框状态
	state.SetGlobal("getUICheck", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethodWithBool("getUICheck", L, false)
	}))

	// 注册 getUIEnable - 获取控件是否可用
	state.SetGlobal("getUIEnable", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethodWithBool("getUIEnable", L, false)
	}))

	// 注册 setUIEnable - 设置控件是否可用
	state.SetGlobal("setUIEnable", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("setUIEnable", L)
	}))

	// 注册 getUIVisible - 获取当前控件可见的值
	state.SetGlobal("getUIVisible", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethodWithNumber("getUIVisible", L, 0)
	}))

	// 注册 setUIVisible - 设置控件隐藏或可见
	state.SetGlobal("setUIVisible", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("setUIVisible", L)
	}))

	// 注册 setUITextColor - 修改控件字体颜色
	state.SetGlobal("setUITextColor", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("setUITextColor", L)
	}))

	// 注册 setUIBackground - 设置窗口或者控件的背景
	state.SetGlobal("setUIBackground", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("setUIBackground", L)
	}))

	// 注册 setUIConfig - 给指定窗口ui加载一个配置
	state.SetGlobal("setUIConfig", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("setUIConfig", L)
	}))

	// 注册 toast - 弹窗显示信息
	state.SetGlobal("toast", state.NewFunction(func(L *lua.LState) int {
		// 获取消息文本
		message := L.CheckString(1)
		if message == "" {
			return 0
		}

		// 调用 AutoGo 的 utils.Toast
		utils.Toast(message, 0, 0, -1)

		return 0
	}))

	// 注册 hideToast - 关闭消息显示
	state.SetGlobal("hideToast", state.NewFunction(func(L *lua.LState) int {
		// AutoGo 的 utils.Toast 不需要手动隐藏，它会自动消失
		// 这里不做任何操作，保持兼容性
		return 0
	}))

	// 注册 showUI - 显示一个自定义的界面
	state.SetGlobal("showUI", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethodWithString("showUI", L, "")
	}))

	// 注册 showUIEx - 显示一个自定义的界面（扩展版本）
	state.SetGlobal("showUIEx", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethodWithString("showUIEx", L, "")
	}))

	// 注册 createHUD - 创建一个HUD用于显示
	state.SetGlobal("createHUD", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethodWithNumber("createHUD", L, 0)
	}))

	// 注册 showHUD - 显示HUD或者刷新
	state.SetGlobal("showHUD", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("showHUD", L)
	}))

	// 注册 hideHUD - 隐藏并销毁HUD
	state.SetGlobal("hideHUD", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethodWithBool("hideHUD", L, false)
	}))

	return nil
}

// GetModule 获取模块实例
func GetModule() model.Module {
	return New()
}
