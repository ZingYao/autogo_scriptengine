//go:build ignore
// +build ignore

package dynamicui

import (
	"fmt"

	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"

	lua "github.com/yuin/gopher-lua"
)

// DynamicUIModule dynamicui 模块
type DynamicUIModule struct {
	ThrowException bool // 是否抛出异常
	ShowWarning    bool // 是否显示警告信息
	Debug          bool // 是否开启调试模式（打印脚本堆栈信息）
}

// Name 返回模块名称
func (m *DynamicUIModule) Name() string {
	return "dynamicui"
}

// IsAvailable 检查模块是否可用
func (m *DynamicUIModule) IsAvailable() bool {
	return true
}

// Register 向引擎注册方法
func (m *DynamicUIModule) Register(engine model.Engine) error {
	// 确保 config 不为 nil
	if m.ThrowException == false && m.ShowWarning == false && m.Debug == false {
		m.ThrowException = false
		m.ShowWarning = true
		m.Debug = false
	}

	state := engine.GetState()
	m.Inject(state)
	for _, methodName := range dynamicUIMethodNames {
		engine.RegisterMethod("ui."+methodName, "dynamicui 占位方法", m.placeholderMethod(methodName), true)
	}
	return nil
}

// New 创建一个新的 DynamicUIModule 实例
func New() *DynamicUIModule {
	return &DynamicUIModule{
		ThrowException: false,
		ShowWarning:    true,
		Debug:          false,
	}
}

// handleEmptyMethod 处理空方法的通用逻辑
func (m *DynamicUIModule) handleEmptyMethod(methodName string, L *lua.LState) int {
	// 如果开启了调试模式，打印脚本堆栈信息
	if m.Debug {
		fmt.Println("\n=== 调试信息: 调用未实现方法 ===")
		fmt.Printf("方法名: dynamicui.%s\n", methodName)
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
		fmt.Println("=== 堆栈结束 ===")
	}

	// 打印警告信息
	if m.ShowWarning {
		fmt.Printf("[警告] dynamicui.%s 方法未实现，此功能在 AutoGo 中不支持\n", methodName)
	}

	// 根据配置决定是否抛出异常
	if m.ThrowException {
		L.RaiseError(fmt.Sprintf("dynamicui.%s 方法未实现，此功能在 AutoGo 中不支持", methodName))
		return 0
	}

	// 默认返回 true（兼容懒人精灵的返回值）
	L.Push(lua.LBool(true))
	return 1
}

// Inject 将 dynamicui 模块注入到 Lua 状态中
func (m *DynamicUIModule) Inject(state *lua.LState) {
	// 创建 ui 表
	uiTable := state.NewTable()

	// 1. ui.newLayout 创建一个新的布局
	uiTable.RawSetString("newLayout", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("newLayout", L)
	}))

	// 2. ui.show 显示一个布局
	uiTable.RawSetString("show", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("show", L)
	}))

	// 3. ui.dismiss 关闭一个布局
	uiTable.RawSetString("dismiss", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("dismiss", L)
	}))

	// 4. ui.newRow 布局换行排列
	uiTable.RawSetString("newRow", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("newRow", L)
	}))

	// 5. ui.addButton 创建一个按钮
	uiTable.RawSetString("addButton", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("addButton", L)
	}))

	// 6. ui.addTextView 创建文字框控件
	uiTable.RawSetString("addTextView", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("addTextView", L)
	}))

	// 7. ui.addEditText 创建输入框控件
	uiTable.RawSetString("addEditText", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("addEditText", L)
	}))

	// 8. ui.addCheckBox 创建多选框控件
	uiTable.RawSetString("addCheckBox", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("addCheckBox", L)
	}))

	// 9. ui.addRadioGroup 创建单选框控件
	uiTable.RawSetString("addRadioGroup", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("addRadioGroup", L)
	}))

	// 10. ui.addSpinner 创建下拉框控件
	uiTable.RawSetString("addSpinner", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("addSpinner", L)
	}))

	// 11. ui.addSeekBar 创建滑动条控件
	uiTable.RawSetString("addSeekBar", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("addSeekBar", L)
	}))

	// 12. ui.addProgressBar 创建进度条控件
	uiTable.RawSetString("addProgressBar", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("addProgressBar", L)
	}))

	// 13. ui.addImageView 创建图像控件
	uiTable.RawSetString("addImageView", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("addImageView", L)
	}))

	// 12. ui.addLine 创建线控件
	uiTable.RawSetString("addLine", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("addLine", L)
	}))

	// 13. ui.addWebView 创建一个浏览器控件
	uiTable.RawSetString("addWebView", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("addWebView", L)
	}))

	// 14. ui.addTimePicker 创建时间选择器
	uiTable.RawSetString("addTimePicker", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("addTimePicker", L)
	}))

	// 15. ui.addDatePicker 创建日期选择器
	uiTable.RawSetString("addDatePicker", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("addDatePicker", L)
	}))

	// 16. ui.addSwitch 创建开关控件
	uiTable.RawSetString("addSwitch", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("addSwitch", L)
	}))

	// 17. ui.addColorPicker 创建颜色选择器
	uiTable.RawSetString("addColorPicker", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("addColorPicker", L)
	}))

	// 18. ui.addRatingBar 创建评分条
	uiTable.RawSetString("addRatingBar", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("addRatingBar", L)
	}))

	// 19. ui.addSpace 添加空白
	uiTable.RawSetString("addSpace", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("addSpace", L)
	}))

	// 20. ui.callJs 调用webview打开的网页中的js函数
	uiTable.RawSetString("callJs", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("callJs", L)
	}))

	// 15. ui.addTabView 创建标签页控件
	uiTable.RawSetString("addTabView", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("addTabView", L)
	}))

	// 16. ui.addTab 创建子标签页控件
	uiTable.RawSetString("addTab", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("addTab", L)
	}))

	// 17. ui.setLine 重设线控件
	uiTable.RawSetString("setLine", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("setLine", L)
	}))

	// 18. ui.setButton 重设按钮控件
	uiTable.RawSetString("setButton", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("setButton", L)
	}))

	// 19. ui.setEditText 重设输入框控件
	uiTable.RawSetString("setEditText", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("setEditText", L)
	}))

	// 20. ui.setEditHintText 设置输入框默认提示字符串
	uiTable.RawSetString("setEditHintText", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("setEditHintText", L)
	}))

	// 21. ui.setTextView 重设文本框控件
	uiTable.RawSetString("setTextView", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("setTextView", L)
	}))

	// 22. ui.setCheckBox 重设多选框控件
	uiTable.RawSetString("setCheckBox", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("setCheckBox", L)
	}))

	// 23. ui.setRadioGroup 重设单选框控件
	uiTable.RawSetString("setRadioGroup", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("setRadioGroup", L)
	}))

	// 24. ui.setSpinner 重设下拉框控件
	uiTable.RawSetString("setSpinner", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("setSpinner", L)
	}))

	// 25. ui.setWebView 重设浏览器控件
	uiTable.RawSetString("setWebView", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("setWebView", L)
	}))

	// 26. ui.setImageView 重设图像控件
	uiTable.RawSetString("setImageView", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("setImageView", L)
	}))

	// 27. ui.setImageViewEx 重设图像控件
	uiTable.RawSetString("setImageViewEx", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("setImageViewEx", L)
	}))

	// 28. ui.setText 控件设置文字
	uiTable.RawSetString("setText", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("setText", L)
	}))

	// 29. ui.setTitleText 设置布局标题
	uiTable.RawSetString("setTitleText", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("setTitleText", L)
	}))

	// 30. ui.setTextSize 设置文字大小
	uiTable.RawSetString("setTextSize", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("setTextSize", L)
	}))

	// 31. ui.setEnable 设置控件可用状态
	uiTable.RawSetString("setEnable", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("setEnable", L)
	}))

	// 32. ui.setVisiblity 设置控件显示状态
	uiTable.RawSetString("setVisiblity", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("setVisiblity", L)
	}))

	// 33. ui.setRowVisibleByGid 批量设置行控件显示状态通过gid
	uiTable.RawSetString("setRowVisibleByGid", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("setRowVisibleByGid", L)
	}))

	// 34. ui.setBackground 设置控件背景颜色
	uiTable.RawSetString("setBackground", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("setBackground", L)
	}))

	// 35. ui.setTitleBackground 设置标题栏背景颜色
	uiTable.RawSetString("setTitleBackground", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("setTitleBackground", L)
	}))

	// 36. ui.setTextColor 设置文字颜色
	uiTable.RawSetString("setTextColor", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("setTextColor", L)
	}))

	// 37. ui.setInputType 设置输入类型
	uiTable.RawSetString("setInputType", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("setInputType", L)
	}))

	// 38. ui.getText 获取文字
	uiTable.RawSetString("getText", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("getText", L)
	}))

	// 39. ui.getEnable 获取可用状态
	uiTable.RawSetString("getEnable", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("getEnable", L)
	}))

	// 40. ui.getVisible 获取显示状态
	uiTable.RawSetString("getVisible", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("getVisible", L)
	}))

	// 41. ui.getTextColor 获取文字颜色
	uiTable.RawSetString("getTextColor", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("getTextColor", L)
	}))

	// 42. ui.setFullScreen 设置控件宽度全屏
	uiTable.RawSetString("setFullScreen", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("setFullScreen", L)
	}))

	// 43. ui.setPadding 设置控件内边距
	uiTable.RawSetString("setPadding", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("setPadding", L)
	}))

	// 44. ui.setGravity 设置控件对齐方式
	uiTable.RawSetString("setGravity", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("setGravity", L)
	}))

	// 45. ui.setOnClick 设置控件单击事件
	uiTable.RawSetString("setOnClick", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("setOnClick", L)
	}))

	// 46. ui.setOnBackEvent 设置窗口监听返回按键消息
	uiTable.RawSetString("setOnBackEvent", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("setOnBackEvent", L)
	}))

	// 47. ui.setOnClose 设置窗口关闭事件
	uiTable.RawSetString("setOnClose", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("setOnClose", L)
	}))

	// 48. ui.setOnChange 设置控件改变事件
	uiTable.RawSetString("setOnChange", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("setOnChange", L)
	}))

	// 49. ui.getValue 获取控件值
	uiTable.RawSetString("getValue", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("getValue", L)
	}))

	// 50. ui.getData 获取当前界面所有控件的值
	uiTable.RawSetString("getData", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("getData", L)
	}))

	// 51. ui.loadProfile 读取设置
	uiTable.RawSetString("loadProfile", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("loadProfile", L)
	}))

	// 52. ui.saveProfile 保存配置
	uiTable.RawSetString("saveProfile", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("saveProfile", L)
	}))

	// 53. ui.beginUiQueue 创建一个快速ui命令队列
	uiTable.RawSetString("beginUiQueue", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("beginUiQueue", L)
	}))

	// 54. ui.endUiQueue 执行这个ui命令队列
	uiTable.RawSetString("endUiQueue", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("endUiQueue", L)
	}))

	// 55. ui.addTableView 添加表格控件
	uiTable.RawSetString("addTableView", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("addTableView", L)
	}))

	// 56. ui.setTableViewAttrib 设置表格属性
	uiTable.RawSetString("setTableViewAttrib", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("setTableViewAttrib", L)
	}))

	// 57. ui.getTableViewAllData 获取表格全部数据
	uiTable.RawSetString("getTableViewAllData", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("getTableViewAllData", L)
	}))

	// 58. ui.getTableViewRowData 获取表格指定行数据
	uiTable.RawSetString("getTableViewRowData", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("getTableViewRowData", L)
	}))

	// 59. ui.getTableViewRowData 获取表格指定行数据（重复方法名）
	uiTable.RawSetString("getTableViewRowData", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("getTableViewRowData", L)
	}))

	// 60. ui.addTableViewRow 添加表格行
	uiTable.RawSetString("addTableViewRow", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("addTableViewRow", L)
	}))

	// 61. ui.removeTableViewRow 删除指定行
	uiTable.RawSetString("removeTableViewRow", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("removeTableViewRow", L)
	}))

	// 62. ui.getTableViewRowCnt 获取表格行数
	uiTable.RawSetString("getTableViewRowCnt", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("getTableViewRowCnt", L)
	}))

	// 63. ui.getTableViewSelectIndex 获取表格选中行索引
	uiTable.RawSetString("getTableViewSelectIndex", state.NewFunction(func(L *lua.LState) int {
		return m.handleEmptyMethod("getTableViewSelectIndex", L)
	}))

	// 将 ui 表注册到全局环境
	state.SetGlobal("ui", uiTable)
}

func (m *DynamicUIModule) placeholderMethod(methodName string) func(...interface{}) (bool, error) {
	return func(...interface{}) (bool, error) {
		if m.ShowWarning {
			fmt.Printf("[警告] dynamicui.%s 方法未实现，此功能在 AutoGo 中不支持\n", methodName)
		}
		if m.ThrowException {
			return false, fmt.Errorf("dynamicui.%s 方法未实现，此功能在 AutoGo 中不支持", methodName)
		}
		return true, nil
	}
}

var dynamicUIMethodNames = []string{
	"newLayout",
	"show",
	"dismiss",
	"newRow",
	"addButton",
	"addTextView",
	"addEditText",
	"addCheckBox",
	"addRadioGroup",
	"addSpinner",
	"addSeekBar",
	"addProgressBar",
	"addImageView",
	"addLine",
	"addWebView",
	"addTimePicker",
	"addDatePicker",
	"addSwitch",
	"addColorPicker",
	"addRatingBar",
	"addSpace",
	"callJs",
	"addTabView",
	"addTab",
	"setLine",
	"setButton",
	"setEditText",
	"setEditHintText",
	"setTextView",
	"setCheckBox",
	"setRadioGroup",
	"setSpinner",
	"setWebView",
	"setImageView",
	"setImageViewEx",
	"setText",
	"setTitleText",
	"setTextSize",
	"setEnable",
	"setVisiblity",
	"setRowVisibleByGid",
	"setBackground",
	"setTitleBackground",
	"setTextColor",
	"setInputType",
	"getText",
	"getEnable",
	"getVisible",
	"getTextColor",
	"setFullScreen",
	"setPadding",
	"setGravity",
	"setOnClick",
	"setOnBackEvent",
	"setOnClose",
	"setOnChange",
	"getValue",
	"getData",
	"loadProfile",
	"saveProfile",
	"beginUiQueue",
	"endUiQueue",
	"addTableView",
	"setTableViewAttrib",
	"getTableViewAllData",
	"getTableViewRowData",
	"addTableViewRow",
	"removeTableViewRow",
	"getTableViewRowCnt",
	"getTableViewSelectIndex",
}

// GetModule 获取模块实例
func GetModule() model.Module {
	return New()
}
