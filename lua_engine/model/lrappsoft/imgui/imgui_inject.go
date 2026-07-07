//go:build ignore
// +build ignore

package imgui

import (
	"fmt"
	"sync"

	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"
	lua "github.com/yuin/gopher-lua"
)

// ImGuiModule imgui 模块（懒人精灵兼容）
type ImGuiModule struct {
	widgets       map[string]*Widget
	widgetMutex   sync.RWMutex
	widgetCounter int
}

// Widget 控件信息
type Widget struct {
	handle     string
	widgetType string
	properties map[string]interface{}
	callbacks  map[string]lua.LGFunction
}

// New 创建一个新的 ImGuiModule 实例
func New() *ImGuiModule {
	return &ImGuiModule{
		widgets:       make(map[string]*Widget),
		widgetCounter: 0,
	}
}

// convertHandle 将 Lua 句柄转换为字符串
func (m *ImGuiModule) convertHandle(handle lua.LValue) string {
	switch v := handle.(type) {
	case lua.LNumber:
		return fmt.Sprintf("%d", int(v))
	case lua.LString:
		return string(v)
	default:
		return ""
	}
}

// Name 返回模块名称
func (m *ImGuiModule) Name() string {
	return "imgui"
}

// IsAvailable 检查模块是否可用
func (m *ImGuiModule) IsAvailable() bool {
	return true
}

// Register 向引擎注册方法
func (m *ImGuiModule) Register(engine model.Engine) error {
	state := engine.GetState()
	if m.widgets == nil {
		m.widgets = make(map[string]*Widget)
	}

	// 创建 imgui 表
	imguiTable := state.NewTable()

	// 注册基础方法
	imguiTable.RawSetString("isSupport", state.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LBool(true))
		return 1
	}))

	imguiTable.RawSetString("getLastError", state.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LString(""))
		return 1
	}))

	// 注册控件创建方法
	imguiTable.RawSetString("createButton", state.NewFunction(func(L *lua.LState) int {
		return m.createButton(L)
	}))

	imguiTable.RawSetString("createCheckBox", state.NewFunction(func(L *lua.LState) int {
		return m.createCheckBox(L)
	}))

	imguiTable.RawSetString("createColorPicker", state.NewFunction(func(L *lua.LState) int {
		return m.createColorPicker(L)
	}))

	imguiTable.RawSetString("createSwitch", state.NewFunction(func(L *lua.LState) int {
		return m.createSwitch(L)
	}))

	imguiTable.RawSetString("createLabel", state.NewFunction(func(L *lua.LState) int {
		return m.createLabel(L)
	}))

	imguiTable.RawSetString("createInputText", state.NewFunction(func(L *lua.LState) int {
		return m.createInputText(L)
	}))

	imguiTable.RawSetString("createProgressBar", state.NewFunction(func(L *lua.LState) int {
		return m.createProgressBar(L)
	}))

	imguiTable.RawSetString("createComboBox", state.NewFunction(func(L *lua.LState) int {
		return m.createComboBox(L)
	}))

	imguiTable.RawSetString("createRadioGroup", state.NewFunction(func(L *lua.LState) int {
		return m.createRadioGroup(L)
	}))

	imguiTable.RawSetString("createTableView", state.NewFunction(func(L *lua.LState) int {
		return m.createTableView(L)
	}))

	imguiTable.RawSetString("createSlider", state.NewFunction(func(L *lua.LState) int {
		return m.createSlider(L)
	}))

	imguiTable.RawSetString("createBitmapShape", state.NewFunction(func(L *lua.LState) int {
		return m.createBitmapShape(L)
	}))

	// 注册事件回调方法
	imguiTable.RawSetString("setOnClick", state.NewFunction(func(L *lua.LState) int {
		return m.setOnClick(L)
	}))

	imguiTable.RawSetString("setOnCheck", state.NewFunction(func(L *lua.LState) int {
		return m.setOnCheck(L)
	}))

	imguiTable.RawSetString("setOnSelectEvent", state.NewFunction(func(L *lua.LState) int {
		return m.setOnSelectEvent(L)
	}))

	imguiTable.RawSetString("setOnSelectEventEx", state.NewFunction(func(L *lua.LState) int {
		return m.setOnSelectEventEx(L)
	}))

	// 注册控件操作方法
	imguiTable.RawSetString("setChecked", state.NewFunction(func(L *lua.LState) int {
		return m.setChecked(L)
	}))

	imguiTable.RawSetString("isChecked", state.NewFunction(func(L *lua.LState) int {
		return m.isChecked(L)
	}))

	imguiTable.RawSetString("setWidgetText", state.NewFunction(func(L *lua.LState) int {
		return m.setWidgetText(L)
	}))

	imguiTable.RawSetString("getWidgetText", state.NewFunction(func(L *lua.LState) int {
		return m.getWidgetText(L)
	}))

	imguiTable.RawSetString("getInputText", state.NewFunction(func(L *lua.LState) int {
		return m.getInputText(L)
	}))

	imguiTable.RawSetString("setInputText", state.NewFunction(func(L *lua.LState) int {
		return m.setInputText(L)
	}))

	imguiTable.RawSetString("setInputType", state.NewFunction(func(L *lua.LState) int {
		return m.setInputType(L)
	}))

	imguiTable.RawSetString("setProgressBarPos", state.NewFunction(func(L *lua.LState) int {
		return m.setProgressBarPos(L)
	}))

	imguiTable.RawSetString("getProgressBarPos", state.NewFunction(func(L *lua.LState) int {
		return m.getProgressBarPos(L)
	}))

	imguiTable.RawSetString("getItemText", state.NewFunction(func(L *lua.LState) int {
		return m.getItemText(L)
	}))

	imguiTable.RawSetString("removeItemAt", state.NewFunction(func(L *lua.LState) int {
		return m.removeItemAt(L)
	}))

	imguiTable.RawSetString("removeAllItems", state.NewFunction(func(L *lua.LState) int {
		return m.removeAllItems(L)
	}))

	imguiTable.RawSetString("getSelectedItemIndex", state.NewFunction(func(L *lua.LState) int {
		return m.getSelectedItemIndex(L)
	}))

	imguiTable.RawSetString("setItemSelected", state.NewFunction(func(L *lua.LState) int {
		return m.setItemSelected(L)
	}))

	imguiTable.RawSetString("getItemCount", state.NewFunction(func(L *lua.LState) int {
		return m.getItemCount(L)
	}))

	imguiTable.RawSetString("setTableHeaderItem", state.NewFunction(func(L *lua.LState) int {
		return m.setTableHeaderItem(L)
	}))

	imguiTable.RawSetString("insertTableRow", state.NewFunction(func(L *lua.LState) int {
		return m.insertTableRow(L)
	}))

	imguiTable.RawSetString("getTableItemText", state.NewFunction(func(L *lua.LState) int {
		return m.getTableItemText(L)
	}))

	imguiTable.RawSetString("setTableItemText", state.NewFunction(func(L *lua.LState) int {
		return m.setTableItemText(L)
	}))

	imguiTable.RawSetString("deleteTableRow", state.NewFunction(func(L *lua.LState) int {
		return m.deleteTableRow(L)
	}))

	imguiTable.RawSetString("clearTable", state.NewFunction(func(L *lua.LState) int {
		return m.clearTable(L)
	}))

	imguiTable.RawSetString("setSlider", state.NewFunction(func(L *lua.LState) int {
		return m.setSlider(L)
	}))

	imguiTable.RawSetString("getSlider", state.NewFunction(func(L *lua.LState) int {
		return m.getSlider(L)
	}))

	imguiTable.RawSetString("addRadioBox", state.NewFunction(func(L *lua.LState) int {
		return m.addRadioBox(L)
	}))

	imguiTable.RawSetString("addOptionItem", state.NewFunction(func(L *lua.LState) int {
		return m.addOptionItem(L)
	}))

	imguiTable.RawSetString("addTabBarItem", state.NewFunction(func(L *lua.LState) int {
		return m.addTabBarItem(L)
	}))

	imguiTable.RawSetString("isValidHandle", state.NewFunction(func(L *lua.LState) int {
		return m.isValidHandle(L)
	}))

	imguiTable.RawSetString("setStyleColor", state.NewFunction(func(L *lua.LState) int {
		return m.setStyleColor(L)
	}))

	imguiTable.RawSetString("close", state.NewFunction(func(L *lua.LState) int {
		return m.close(L)
	}))

	// 注册图形绘制方法
	imguiTable.RawSetString("createRectangle", state.NewFunction(func(L *lua.LState) int {
		return m.createRectangle(L)
	}))

	imguiTable.RawSetString("createCircle", state.NewFunction(func(L *lua.LState) int {
		return m.createCircle(L)
	}))

	imguiTable.RawSetString("createPolygon", state.NewFunction(func(L *lua.LState) int {
		return m.createPolygon(L)
	}))

	imguiTable.RawSetString("createLine", state.NewFunction(func(L *lua.LState) int {
		return m.createLine(L)
	}))

	imguiTable.RawSetString("createShapeText", state.NewFunction(func(L *lua.LState) int {
		return m.createShapeText(L)
	}))

	// 注册图形操作方法
	imguiTable.RawSetString("setShapePosition", state.NewFunction(func(L *lua.LState) int {
		return m.setShapePosition(L)
	}))

	imguiTable.RawSetString("setShapeVisibility", state.NewFunction(func(L *lua.LState) int {
		return m.setShapeVisibility(L)
	}))

	imguiTable.RawSetString("isShapeVisibility", state.NewFunction(func(L *lua.LState) int {
		return m.isShapeVisibility(L)
	}))

	imguiTable.RawSetString("setShapeTextString", state.NewFunction(func(L *lua.LState) int {
		return m.setShapeTextString(L)
	}))

	imguiTable.RawSetString("setShapeTextColor", state.NewFunction(func(L *lua.LState) int {
		return m.setShapeTextColor(L)
	}))

	imguiTable.RawSetString("setShapeTextBackground", state.NewFunction(func(L *lua.LState) int {
		return m.setShapeTextBackground(L)
	}))

	imguiTable.RawSetString("setShapeTextFontScale", state.NewFunction(func(L *lua.LState) int {
		return m.setShapeTextFontScale(L)
	}))

	imguiTable.RawSetString("removeShape", state.NewFunction(func(L *lua.LState) int {
		return m.removeShape(L)
	}))

	imguiTable.RawSetString("setBitmapShape", state.NewFunction(func(L *lua.LState) int {
		return m.setBitmapShape(L)
	}))

	imguiTable.RawSetString("setShapeThickness", state.NewFunction(func(L *lua.LState) int {
		return m.setShapeThickness(L)
	}))

	// 注册控件操作方法
	imguiTable.RawSetString("setWidgetSize", state.NewFunction(func(L *lua.LState) int {
		return m.setWidgetSize(L)
	}))

	imguiTable.RawSetString("setWidgetVisible", state.NewFunction(func(L *lua.LState) int {
		return m.setWidgetVisible(L)
	}))

	imguiTable.RawSetString("isWidgetVisible", state.NewFunction(func(L *lua.LState) int {
		return m.isWidgetVisible(L)
	}))

	imguiTable.RawSetString("setLayoutBorderVisible", state.NewFunction(func(L *lua.LState) int {
		return m.setLayoutBorderVisible(L)
	}))

	imguiTable.RawSetString("setWindowPos", state.NewFunction(func(L *lua.LState) int {
		return m.setWindowPos(L)
	}))

	imguiTable.RawSetString("setWindowSize", state.NewFunction(func(L *lua.LState) int {
		return m.setWindowSize(L)
	}))

	imguiTable.RawSetString("setWidgetStyle", state.NewFunction(func(L *lua.LState) int {
		return m.setWidgetStyle(L)
	}))

	imguiTable.RawSetString("setWidgetColor", state.NewFunction(func(L *lua.LState) int {
		return m.setWidgetColor(L)
	}))

	imguiTable.RawSetString("setWindowFlags", state.NewFunction(func(L *lua.LState) int {
		return m.setWindowFlags(L)
	}))

	imguiTable.RawSetString("getWindowPos", state.NewFunction(func(L *lua.LState) int {
		return m.getWindowPos(L)
	}))

	imguiTable.RawSetString("setOnClose", state.NewFunction(func(L *lua.LState) int {
		return m.setOnClose(L)
	}))

	imguiTable.RawSetString("destroyWindow", state.NewFunction(func(L *lua.LState) int {
		return m.destroyWindow(L)
	}))

	imguiTable.RawSetString("sameLine", state.NewFunction(func(L *lua.LState) int {
		return m.sameLine(L)
	}))

	// 注册窗口和布局方法
	imguiTable.RawSetString("createWindow", state.NewFunction(func(L *lua.LState) int {
		return m.createWindow(L)
	}))

	imguiTable.RawSetString("createVerticalLayout", state.NewFunction(func(L *lua.LState) int {
		return m.createVerticalLayout(L)
	}))

	imguiTable.RawSetString("createHorticalLayout", state.NewFunction(func(L *lua.LState) int {
		return m.createHorticalLayout(L)
	}))

	imguiTable.RawSetString("createTreeBoxLayout", state.NewFunction(func(L *lua.LState) int {
		return m.createTreeBoxLayout(L)
	}))

	imguiTable.RawSetString("createTabBar", state.NewFunction(func(L *lua.LState) int {
		return m.createTabBar(L)
	}))

	// 注册图片管理方法
	imguiTable.RawSetString("createImage", state.NewFunction(func(L *lua.LState) int {
		return m.createImage(L)
	}))

	imguiTable.RawSetString("setImage", state.NewFunction(func(L *lua.LState) int {
		return m.setImage(L)
	}))

	imguiTable.RawSetString("setImageFromBitmap", state.NewFunction(func(L *lua.LState) int {
		return m.setImageFromBitmap(L)
	}))

	// 注册主题系统方法
	imguiTable.RawSetString("setColorTheme", state.NewFunction(func(L *lua.LState) int {
		return m.setColorTheme(L)
	}))

	// 注册框架方法
	imguiTable.RawSetString("show", state.NewFunction(func(L *lua.LState) int {
		return m.show(L)
	}))

	imguiTable.RawSetString("showWindow", state.NewFunction(func(L *lua.LState) int {
		return m.showWindow(L)
	}))

	// 注册滑块回调方法
	imguiTable.RawSetString("setOnSliderEvent", state.NewFunction(func(L *lua.LState) int {
		return m.setOnSliderEvent(L)
	}))

	// 注册 imgui 表到全局
	state.SetGlobal("imgui", imguiTable)

	engine.RegisterMethod("imgui.isSupport", "检查 imgui 是否支持", func() bool {
		return true
	}, true)
	engine.RegisterMethod("imgui.getLastError", "获取最后错误", func() string {
		return ""
	}, true)
	engine.RegisterMethod("imgui.createButton", "创建按钮", m.createButtonValue, true)
	engine.RegisterMethod("imgui.createCheckBox", "创建复选框", m.createCheckBoxValue, true)
	engine.RegisterMethod("imgui.createColorPicker", "创建颜色选择器", m.createColorPickerValue, true)
	engine.RegisterMethod("imgui.createSwitch", "创建开关", m.createSwitchValue, true)
	engine.RegisterMethod("imgui.createLabel", "创建文本标签", m.createLabelValue, true)
	engine.RegisterMethod("imgui.createInputText", "创建输入框", m.createInputTextValue, true)
	engine.RegisterMethod("imgui.createProgressBar", "创建进度条", m.createProgressBarValue, true)
	engine.RegisterMethod("imgui.createComboBox", "创建组合框", m.createComboBoxValue, true)
	engine.RegisterMethod("imgui.createRadioGroup", "创建单选按钮组", m.createRadioGroupValue, true)
	engine.RegisterMethod("imgui.createTableView", "创建表格视图", m.createTableViewValue, true)
	engine.RegisterMethod("imgui.createSlider", "创建滑块", m.createSliderValue, true)
	engine.RegisterMethod("imgui.setChecked", "设置选中状态", m.setCheckedValue, true)
	engine.RegisterMethod("imgui.isChecked", "获取选中状态", m.isCheckedValue, true)
	engine.RegisterMethod("imgui.setWidgetText", "设置控件文本", m.setWidgetTextValue, true)
	engine.RegisterMethod("imgui.getWidgetText", "获取控件文本", m.getWidgetTextValue, true)
	engine.RegisterMethod("imgui.getInputText", "获取输入框文本", m.getInputTextValue, true)
	engine.RegisterMethod("imgui.setInputText", "设置输入框文本", m.setInputTextValue, true)
	engine.RegisterMethod("imgui.setProgressBarPos", "设置进度条进度", m.setProgressBarPosValue, true)
	engine.RegisterMethod("imgui.getProgressBarPos", "获取进度条进度", m.getProgressBarPosValue, true)
	engine.RegisterMethod("imgui.getItemText", "获取列表项文本", m.getItemTextValue, true)
	engine.RegisterMethod("imgui.removeItemAt", "删除列表项", m.removeItemAtValue, true)
	engine.RegisterMethod("imgui.removeAllItems", "清空列表项", m.removeAllItemsValue, true)
	engine.RegisterMethod("imgui.getSelectedItemIndex", "获取选中项索引", m.getSelectedItemIndexValue, true)
	engine.RegisterMethod("imgui.setItemSelected", "设置选中项索引", m.setItemSelectedValue, true)
	engine.RegisterMethod("imgui.getItemCount", "获取列表项数量", m.getItemCountValue, true)
	engine.RegisterMethod("imgui.addRadioBox", "添加单选项", m.addRadioBoxValue, true)
	engine.RegisterMethod("imgui.addOptionItem", "添加下拉项", m.addOptionItemValue, true)
	engine.RegisterMethod("imgui.setSlider", "设置滑块位置", m.setSliderValue, true)
	engine.RegisterMethod("imgui.getSlider", "获取滑块位置", m.getSliderValue, true)
	engine.RegisterMethod("imgui.setTableHeaderItem", "设置表头文本", m.setTableHeaderItemValue, true)
	engine.RegisterMethod("imgui.insertTableRow", "插入表格行", m.insertTableRowValue, true)
	engine.RegisterMethod("imgui.getTableItemText", "获取表格单元格文本", m.getTableItemTextValue, true)
	engine.RegisterMethod("imgui.setTableItemText", "设置表格单元格文本", m.setTableItemTextValue, true)
	engine.RegisterMethod("imgui.deleteTableRow", "删除表格行", m.deleteTableRowValue, true)
	engine.RegisterMethod("imgui.clearTable", "清空表格", m.clearTableValue, true)
	engine.RegisterMethod("imgui.createRectangle", "创建矩形", m.createRectangleValue, true)
	engine.RegisterMethod("imgui.createCircle", "创建圆形", m.createCircleValue, true)
	engine.RegisterMethod("imgui.createLine", "创建线段", m.createLineValue, true)
	engine.RegisterMethod("imgui.createShapeText", "创建文本图形", m.createShapeTextValue, true)
	engine.RegisterMethod("imgui.createBitmapShape", "创建位图图形", m.createBitmapShapeValue, true)
	engine.RegisterMethod("imgui.setShapePosition", "设置图形位置", m.setShapePositionValue, true)
	engine.RegisterMethod("imgui.setShapeVisibility", "设置图形可见性", m.setShapeVisibilityValue, true)
	engine.RegisterMethod("imgui.isShapeVisibility", "获取图形可见性", m.isShapeVisibilityValue, true)
	engine.RegisterMethod("imgui.setShapeTextString", "设置文本图形文字", m.setShapeTextStringValue, true)
	engine.RegisterMethod("imgui.removeShape", "删除图形", m.removeShapeValue, true)
	engine.RegisterMethod("imgui.setWidgetVisible", "设置控件可见性", m.setWidgetVisibleValue, true)
	engine.RegisterMethod("imgui.isWidgetVisible", "获取控件可见性", m.isWidgetVisibleValue, true)
	engine.RegisterMethod("imgui.createWindow", "创建窗口", m.createWindowValue, true)
	engine.RegisterMethod("imgui.createVerticalLayout", "创建垂直布局", m.createVerticalLayoutValue, true)
	engine.RegisterMethod("imgui.createHorticalLayout", "创建水平布局", m.createHorticalLayoutValue, true)
	engine.RegisterMethod("imgui.createTreeBoxLayout", "创建树布局", m.createTreeBoxLayoutValue, true)
	engine.RegisterMethod("imgui.createTabBar", "创建标签栏", m.createTabBarValue, true)
	engine.RegisterMethod("imgui.setInputType", "设置输入框类型", m.setInputTypeValue, true)
	engine.RegisterMethod("imgui.addTabBarItem", "添加标签栏项目", m.addTabBarItemValue, true)
	engine.RegisterMethod("imgui.setShapeTextColor", "设置文本图形文字颜色", m.setShapeTextColorValue, true)
	engine.RegisterMethod("imgui.setShapeTextBackground", "设置文本图形背景颜色", m.setShapeTextBackgroundValue, true)
	engine.RegisterMethod("imgui.setShapeTextFontScale", "设置文本图形字体缩放", m.setShapeTextFontScaleValue, true)
	engine.RegisterMethod("imgui.setBitmapShape", "设置位图图形", m.setBitmapShapeValue, true)
	engine.RegisterMethod("imgui.setShapeThickness", "设置图形线宽", m.setShapeThicknessValue, true)
	engine.RegisterMethod("imgui.setWidgetSize", "设置控件尺寸", m.setWidgetSizeValue, true)
	engine.RegisterMethod("imgui.setLayoutBorderVisible", "设置布局边框可见性", m.setLayoutBorderVisibleValue, true)
	engine.RegisterMethod("imgui.setWindowPos", "设置窗口位置", m.setWindowPosValue, true)
	engine.RegisterMethod("imgui.setWindowSize", "设置窗口尺寸", m.setWindowSizeValue, true)
	engine.RegisterMethod("imgui.setWidgetStyle", "设置控件样式", m.setWidgetStyleValue, true)
	engine.RegisterMethod("imgui.setWidgetColor", "设置控件颜色", m.setWidgetColorValue, true)
	engine.RegisterMethod("imgui.setWindowFlags", "设置窗口标志", m.setWindowFlagsValue, true)
	engine.RegisterMethod("imgui.getWindowPos", "获取窗口位置", m.getWindowPosValue, true)
	engine.RegisterMethod("imgui.sameLine", "同一行布局占位", func() {}, true)
	engine.RegisterMethod("imgui.createImage", "创建图片控件", m.createImageValue, true)
	engine.RegisterMethod("imgui.setImage", "设置图片路径", m.setImageValue, true)
	engine.RegisterMethod("imgui.setImageFromBitmap", "设置图片位图", m.setImageFromBitmapValue, true)
	engine.RegisterMethod("imgui.setColorTheme", "设置颜色主题", m.setColorThemeValue, true)
	engine.RegisterMethod("imgui.setStyleColor", "设置全局样式颜色", m.setStyleColorValue, true)
	engine.RegisterMethod("imgui.show", "显示 imgui", func() {}, true)
	engine.RegisterMethod("imgui.showWindow", "显示窗口", m.showWindowValue, true)
	engine.RegisterMethod("imgui.destroyWindow", "销毁窗口", m.destroyWindowValue, true)
	engine.RegisterMethod("imgui.isValidHandle", "检查控件句柄是否有效", m.isValidHandleValue, true)
	engine.RegisterMethod("imgui.close", "关闭 imgui 框架", func() {}, true)

	return nil
}

// generateHandle 生成唯一的句柄
func (m *ImGuiModule) generateHandle() string {
	m.widgetCounter++
	return fmt.Sprintf("widget_%d", m.widgetCounter)
}

func (m *ImGuiModule) createWidget(widgetType string, properties map[string]interface{}) string {
	handle := m.generateHandle()
	m.widgetMutex.Lock()
	m.widgets[handle] = &Widget{
		handle:     handle,
		widgetType: widgetType,
		properties: properties,
		callbacks:  make(map[string]lua.LGFunction),
	}
	m.widgetMutex.Unlock()
	return handle
}

func (m *ImGuiModule) createButtonValue(x, y, width, height float64, text string) string {
	return m.createWidget("button", map[string]interface{}{
		"x":      x,
		"y":      y,
		"width":  width,
		"height": height,
		"text":   text,
	})
}

func (m *ImGuiModule) createCheckBoxValue(parent, label string, args ...interface{}) string {
	checked := false
	if len(args) > 0 {
		if value, ok := args[0].(bool); ok {
			checked = value
		}
	}
	return m.createWidget("checkbox", map[string]interface{}{
		"parent":  parent,
		"label":   label,
		"checked": checked,
	})
}

func (m *ImGuiModule) createColorPickerValue(parent string, args ...interface{}) string {
	title := "Color"
	if len(args) > 0 {
		if value, ok := args[0].(string); ok {
			title = value
		}
	}
	color := float64(0xFF000000)
	if len(args) > 1 {
		if number, ok := toFloat64(args[1]); ok {
			color = number
		}
	}
	width := float64(0)
	if len(args) > 2 {
		if number, ok := toFloat64(args[2]); ok {
			width = number
		}
	}
	height := float64(0)
	if len(args) > 3 {
		if number, ok := toFloat64(args[3]); ok {
			height = number
		}
	}
	return m.createWidget("colorpicker", map[string]interface{}{
		"parent": parent,
		"title":  title,
		"color":  color,
		"width":  width,
		"height": height,
	})
}

func (m *ImGuiModule) createSwitchValue(parent, label string, args ...interface{}) string {
	checked := false
	if len(args) > 0 {
		if value, ok := args[0].(bool); ok {
			checked = value
		}
	}
	height := float64(0)
	if len(args) > 1 {
		if number, ok := toFloat64(args[1]); ok {
			height = number
		}
	}
	return m.createWidget("switch", map[string]interface{}{
		"parent":  parent,
		"label":   label,
		"checked": checked,
		"height":  height,
	})
}

func (m *ImGuiModule) createLabelValue(parent, text string, args ...interface{}) string {
	singleline := true
	if len(args) > 0 {
		if value, ok := args[0].(bool); ok {
			singleline = value
		}
	}
	return m.createWidget("label", map[string]interface{}{
		"parent":     parent,
		"text":       text,
		"singleline": singleline,
	})
}

func (m *ImGuiModule) createInputTextValue(parent, label string, args ...interface{}) string {
	value := ""
	if len(args) > 0 {
		if text, ok := args[0].(string); ok {
			value = text
		}
	}
	inputType := 0
	if len(args) > 1 {
		if number, ok := toInt(args[1]); ok {
			inputType = number
		}
	}
	width := float64(0)
	if len(args) > 2 {
		if number, ok := toFloat64(args[2]); ok {
			width = number
		}
	}
	height := float64(0)
	if len(args) > 3 {
		if number, ok := toFloat64(args[3]); ok {
			height = number
		}
	}
	return m.createWidget("inputtext", map[string]interface{}{
		"parent":    parent,
		"label":     label,
		"value":     value,
		"inputType": inputType,
		"width":     width,
		"height":    height,
	})
}

func (m *ImGuiModule) createProgressBarValue(parent string, args ...interface{}) string {
	progress := float64(0)
	if len(args) > 0 {
		if number, ok := toFloat64(args[0]); ok {
			progress = number
		}
	}
	width := float64(0)
	if len(args) > 1 {
		if number, ok := toFloat64(args[1]); ok {
			width = number
		}
	}
	height := float64(0)
	if len(args) > 2 {
		if number, ok := toFloat64(args[2]); ok {
			height = number
		}
	}
	return m.createWidget("progressbar", map[string]interface{}{
		"parent":   parent,
		"progress": progress,
		"width":    width,
		"height":   height,
	})
}

func (m *ImGuiModule) createComboBoxValue(parent string, args ...interface{}) string {
	items := ""
	if len(args) > 0 {
		if value, ok := args[0].(string); ok {
			items = value
		}
	}
	width := float64(0)
	if len(args) > 1 {
		if number, ok := toFloat64(args[1]); ok {
			width = number
		}
	}
	return m.createWidget("combobox", map[string]interface{}{
		"parent":        parent,
		"items":         items,
		"width":         width,
		"selectedIndex": -1,
	})
}

func (m *ImGuiModule) createRadioGroupValue(parent, label string) string {
	return m.createWidget("radiogroup", map[string]interface{}{
		"parent":        parent,
		"label":         label,
		"items":         make([]string, 0),
		"selectedIndex": -1,
	})
}

func (m *ImGuiModule) createTableViewValue(parent, title string, columns int, showheader bool, args ...interface{}) string {
	width := float64(0)
	if len(args) > 0 {
		if number, ok := toFloat64(args[0]); ok {
			width = number
		}
	}
	height := float64(0)
	if len(args) > 1 {
		if number, ok := toFloat64(args[1]); ok {
			height = number
		}
	}
	return m.createWidget("tableview", map[string]interface{}{
		"parent":        parent,
		"title":         title,
		"columns":       columns,
		"showheader":    showheader,
		"width":         width,
		"height":        height,
		"rows":          make([][]string, 0),
		"headers":       make([]string, columns),
		"selectedIndex": -1,
	})
}

func (m *ImGuiModule) createSliderValue(parent, label string, min, max, initialPos float64, args ...interface{}) string {
	width := float64(0)
	if len(args) > 0 {
		if number, ok := toFloat64(args[0]); ok {
			width = number
		}
	}
	return m.createWidget("slider", map[string]interface{}{
		"parent": parent,
		"label":  label,
		"min":    min,
		"max":    max,
		"value":  initialPos,
		"width":  width,
	})
}

func (m *ImGuiModule) isValidHandleValue(handle string) bool {
	m.widgetMutex.RLock()
	defer m.widgetMutex.RUnlock()
	_, exists := m.widgets[handle]
	return exists
}

func (m *ImGuiModule) setCheckedValue(handle string, state bool) {
	m.widgetMutex.Lock()
	defer m.widgetMutex.Unlock()
	if widget, exists := m.widgets[handle]; exists {
		widget.properties["checked"] = state
	}
}

func (m *ImGuiModule) isCheckedValue(handle string) interface{} {
	m.widgetMutex.RLock()
	defer m.widgetMutex.RUnlock()
	widget, exists := m.widgets[handle]
	if !exists {
		return nil
	}
	checked, ok := widget.properties["checked"].(bool)
	if !ok {
		return nil
	}
	return checked
}

func (m *ImGuiModule) setWidgetTextValue(handle, text string) {
	m.widgetMutex.Lock()
	defer m.widgetMutex.Unlock()
	if widget, exists := m.widgets[handle]; exists {
		widget.properties["text"] = text
	}
}

func (m *ImGuiModule) getWidgetTextValue(handle string) string {
	m.widgetMutex.RLock()
	defer m.widgetMutex.RUnlock()
	widget, exists := m.widgets[handle]
	if !exists {
		return ""
	}
	text, ok := widget.properties["text"].(string)
	if !ok {
		return ""
	}
	return text
}

func (m *ImGuiModule) getInputTextValue(handle string) string {
	m.widgetMutex.RLock()
	defer m.widgetMutex.RUnlock()
	widget, exists := m.widgets[handle]
	if !exists {
		return ""
	}
	value, ok := widget.properties["value"].(string)
	if !ok {
		return ""
	}
	return value
}

func (m *ImGuiModule) setInputTextValue(handle, text string) {
	m.widgetMutex.Lock()
	defer m.widgetMutex.Unlock()
	if widget, exists := m.widgets[handle]; exists {
		widget.properties["value"] = text
	}
}

func (m *ImGuiModule) setProgressBarPosValue(handle string, progress float64) {
	m.widgetMutex.Lock()
	defer m.widgetMutex.Unlock()
	if widget, exists := m.widgets[handle]; exists {
		widget.properties["progress"] = progress
	}
}

func (m *ImGuiModule) getProgressBarPosValue(handle string) float64 {
	m.widgetMutex.RLock()
	defer m.widgetMutex.RUnlock()
	widget, exists := m.widgets[handle]
	if !exists {
		return 0
	}
	progress, ok := toFloat64(widget.properties["progress"])
	if !ok {
		return 0
	}
	return progress
}

func (m *ImGuiModule) getItemTextValue(handle string, itemIndex int) interface{} {
	m.widgetMutex.RLock()
	defer m.widgetMutex.RUnlock()
	widget, exists := m.widgets[handle]
	if !exists {
		return ""
	}
	items, ok := widget.properties["items"].(string)
	if !ok {
		return ""
	}
	itemList := splitItems(items)
	if itemIndex < 0 || itemIndex >= len(itemList) {
		return nil
	}
	return itemList[itemIndex]
}

func (m *ImGuiModule) removeItemAtValue(handle string, itemIndex int) {
	m.widgetMutex.Lock()
	defer m.widgetMutex.Unlock()
	widget, exists := m.widgets[handle]
	if !exists {
		return
	}
	items, ok := widget.properties["items"].(string)
	if !ok {
		return
	}
	itemList := splitItems(items)
	if itemIndex < 0 || itemIndex >= len(itemList) {
		return
	}
	widget.properties["items"] = joinItems(append(itemList[:itemIndex], itemList[itemIndex+1:]...))
}

func (m *ImGuiModule) removeAllItemsValue(handle string) {
	m.widgetMutex.Lock()
	defer m.widgetMutex.Unlock()
	if widget, exists := m.widgets[handle]; exists {
		widget.properties["items"] = ""
		widget.properties["selectedIndex"] = -1
	}
}

func (m *ImGuiModule) getSelectedItemIndexValue(handle string) int {
	m.widgetMutex.RLock()
	defer m.widgetMutex.RUnlock()
	widget, exists := m.widgets[handle]
	if !exists {
		return -2
	}
	index, ok := widget.properties["selectedIndex"].(int)
	if !ok {
		return -2
	}
	return index
}

func (m *ImGuiModule) setItemSelectedValue(handle string, index int) {
	m.widgetMutex.Lock()
	defer m.widgetMutex.Unlock()
	if widget, exists := m.widgets[handle]; exists {
		widget.properties["selectedIndex"] = index
	}
}

func (m *ImGuiModule) getItemCountValue(handle string) int {
	m.widgetMutex.RLock()
	defer m.widgetMutex.RUnlock()
	widget, exists := m.widgets[handle]
	if !exists {
		return -1
	}
	switch widget.widgetType {
	case "combobox":
		items, ok := widget.properties["items"].(string)
		if !ok {
			return -1
		}
		return len(splitItems(items))
	case "radiogroup":
		items, ok := widget.properties["items"].([]string)
		if !ok {
			return -1
		}
		return len(items)
	case "tableview":
		rows, ok := widget.properties["rows"].([][]string)
		if !ok {
			return -1
		}
		return len(rows)
	default:
		return -1
	}
}

func (m *ImGuiModule) addRadioBoxValue(handle, text string, _ bool) {
	m.widgetMutex.Lock()
	defer m.widgetMutex.Unlock()
	widget, exists := m.widgets[handle]
	if !exists {
		return
	}
	items, ok := widget.properties["items"].([]string)
	if !ok {
		return
	}
	widget.properties["items"] = append(items, text)
}

func (m *ImGuiModule) addOptionItemValue(handle, itemText string) {
	m.widgetMutex.Lock()
	defer m.widgetMutex.Unlock()
	widget, exists := m.widgets[handle]
	if !exists {
		return
	}
	items, ok := widget.properties["items"].(string)
	if !ok {
		return
	}
	if items == "" {
		widget.properties["items"] = itemText
		return
	}
	widget.properties["items"] = items + "|" + itemText
}

func (m *ImGuiModule) setSliderValue(handle string, pos float64) {
	m.widgetMutex.Lock()
	defer m.widgetMutex.Unlock()
	if widget, exists := m.widgets[handle]; exists {
		widget.properties["value"] = pos
	}
}

func (m *ImGuiModule) getSliderValue(handle string) float64 {
	m.widgetMutex.RLock()
	defer m.widgetMutex.RUnlock()
	widget, exists := m.widgets[handle]
	if !exists {
		return 0
	}
	value, ok := toFloat64(widget.properties["value"])
	if !ok {
		return 0
	}
	return value
}

func (m *ImGuiModule) setTableHeaderItemValue(handle string, col int, text string) {
	m.widgetMutex.Lock()
	defer m.widgetMutex.Unlock()
	widget, exists := m.widgets[handle]
	if !exists {
		return
	}
	headers, ok := widget.properties["headers"].([]string)
	if !ok || col < 0 || col >= len(headers) {
		return
	}
	headers[col] = text
	widget.properties["headers"] = headers
}

func (m *ImGuiModule) insertTableRowValue(handle string, after int) interface{} {
	m.widgetMutex.Lock()
	defer m.widgetMutex.Unlock()
	widget, exists := m.widgets[handle]
	if !exists {
		return nil
	}
	rows, ok := widget.properties["rows"].([][]string)
	if !ok {
		return nil
	}
	columns, ok := widget.properties["columns"].(int)
	if !ok {
		return nil
	}
	newRow := make([]string, columns)
	var newRows [][]string
	switch after {
	case -1:
		newRows = append([][]string{newRow}, rows...)
	case -2:
		newRows = append(rows, newRow)
	default:
		if after < 0 || after >= len(rows) {
			return nil
		}
		newRows = append(rows[:after+1], append([][]string{newRow}, rows[after+1:]...)...)
	}
	widget.properties["rows"] = newRows
	return len(newRows) - 1
}

func (m *ImGuiModule) getTableItemTextValue(handle string, row int, col int) string {
	m.widgetMutex.RLock()
	defer m.widgetMutex.RUnlock()
	widget, exists := m.widgets[handle]
	if !exists {
		return ""
	}
	rows, ok := widget.properties["rows"].([][]string)
	if !ok || row < 0 || row >= len(rows) {
		return ""
	}
	rowData := rows[row]
	if col < 0 || col >= len(rowData) {
		return ""
	}
	return rowData[col]
}

func (m *ImGuiModule) setTableItemTextValue(handle string, row int, col int, text string) {
	m.widgetMutex.Lock()
	defer m.widgetMutex.Unlock()
	widget, exists := m.widgets[handle]
	if !exists {
		return
	}
	rows, ok := widget.properties["rows"].([][]string)
	if !ok || row < 0 || row >= len(rows) {
		return
	}
	rowData := rows[row]
	if col < 0 || col >= len(rowData) {
		return
	}
	rowData[col] = text
	rows[row] = rowData
	widget.properties["rows"] = rows
}

func (m *ImGuiModule) deleteTableRowValue(handle string, row int) {
	m.widgetMutex.Lock()
	defer m.widgetMutex.Unlock()
	widget, exists := m.widgets[handle]
	if !exists {
		return
	}
	rows, ok := widget.properties["rows"].([][]string)
	if !ok || row < 0 || row >= len(rows) {
		return
	}
	widget.properties["rows"] = append(rows[:row], rows[row+1:]...)
}

func (m *ImGuiModule) clearTableValue(handle string) {
	m.widgetMutex.Lock()
	defer m.widgetMutex.Unlock()
	if widget, exists := m.widgets[handle]; exists {
		widget.properties["rows"] = make([][]string, 0)
	}
}

func (m *ImGuiModule) createRectangleValue(x, y, x1, y1, color float64, filled bool, args ...interface{}) string {
	rounding := float64(0)
	if len(args) > 0 {
		if number, ok := toFloat64(args[0]); ok {
			rounding = number
		}
	}
	return m.createWidget("rectangle", map[string]interface{}{
		"x":        x,
		"y":        y,
		"x1":       x1,
		"y1":       y1,
		"color":    color,
		"filled":   filled,
		"rounding": rounding,
		"visible":  true,
	})
}

func (m *ImGuiModule) createCircleValue(x, y, radius, color float64, filled bool, args ...interface{}) string {
	segments := 0
	if len(args) > 0 {
		if number, ok := toInt(args[0]); ok {
			segments = number
		}
	}
	return m.createWidget("circle", map[string]interface{}{
		"x":        x,
		"y":        y,
		"radius":   radius,
		"color":    color,
		"filled":   filled,
		"segments": segments,
		"visible":  true,
	})
}

func (m *ImGuiModule) createLineValue(x1, y1, x2, y2, color float64, args ...interface{}) string {
	thickness := float64(1)
	if len(args) > 0 {
		if number, ok := toFloat64(args[0]); ok {
			thickness = number
		}
	}
	return m.createWidget("line", map[string]interface{}{
		"x1":        x1,
		"y1":        y1,
		"x2":        x2,
		"y2":        y2,
		"color":     color,
		"thickness": thickness,
		"visible":   true,
	})
}

func (m *ImGuiModule) createShapeTextValue(x, y, w, h float64, text string, textColor, bgColor float64, hasBackground bool, args ...interface{}) string {
	fontScale := float64(1)
	if len(args) > 0 {
		if number, ok := toFloat64(args[0]); ok {
			fontScale = number
		}
	}
	return m.createWidget("shapetext", map[string]interface{}{
		"x":             x,
		"y":             y,
		"w":             w,
		"h":             h,
		"text":          text,
		"textColor":     textColor,
		"bgColor":       bgColor,
		"hasBackground": hasBackground,
		"fontScale":     fontScale,
		"visible":       true,
	})
}

func (m *ImGuiModule) createBitmapShapeValue(x, y, width, height float64, bitmap string) string {
	return m.createWidget("bitmap", map[string]interface{}{
		"x":       x,
		"y":       y,
		"width":   width,
		"height":  height,
		"bitmap":  bitmap,
		"visible": true,
	})
}

func (m *ImGuiModule) setShapePositionValue(handle string, x, y float64) {
	m.widgetMutex.Lock()
	defer m.widgetMutex.Unlock()
	if widget, exists := m.widgets[handle]; exists {
		widget.properties["x"] = x
		widget.properties["y"] = y
	}
}

func (m *ImGuiModule) setShapeVisibilityValue(handle string, visible bool) {
	m.setVisibilityValue(handle, visible)
}

func (m *ImGuiModule) isShapeVisibilityValue(handle string) bool {
	return m.isVisibleValue(handle)
}

func (m *ImGuiModule) setShapeTextStringValue(handle, text string) {
	m.setWidgetTextValue(handle, text)
}

func (m *ImGuiModule) setInputTypeValue(handle string, inputType int) {
	m.setWidgetProperty(handle, "inputType", inputType)
}

func (m *ImGuiModule) addTabBarItemValue(tabBarHandle, title string) {
	m.setWidgetProperty(tabBarHandle, "title", title)
}

func (m *ImGuiModule) removeShapeValue(handle string) {
	m.widgetMutex.Lock()
	defer m.widgetMutex.Unlock()
	delete(m.widgets, handle)
}

func (m *ImGuiModule) setWidgetVisibleValue(handle string, visible bool) {
	m.setVisibilityValue(handle, visible)
}

func (m *ImGuiModule) isWidgetVisibleValue(handle string) bool {
	return m.isVisibleValue(handle)
}

func (m *ImGuiModule) createWindowValue(title string, x, y, width, height float64) string {
	return m.createWidget("window", map[string]interface{}{
		"title":   title,
		"x":       x,
		"y":       y,
		"width":   width,
		"height":  height,
		"visible": false,
	})
}

func (m *ImGuiModule) createVerticalLayoutValue(parent string, args ...interface{}) string {
	return m.createLayoutValue("verticalLayout", parent, args...)
}

func (m *ImGuiModule) createHorticalLayoutValue(parent string, args ...interface{}) string {
	return m.createLayoutValue("horticalLayout", parent, args...)
}

func (m *ImGuiModule) createTreeBoxLayoutValue(parent string, args ...interface{}) string {
	return m.createLayoutValue("treeBoxLayout", parent, args...)
}

func (m *ImGuiModule) createTabBarValue(parent string, args ...interface{}) string {
	return m.createLayoutValue("tabBar", parent, args...)
}

func (m *ImGuiModule) createLayoutValue(widgetType, parent string, args ...interface{}) string {
	width := float64(0)
	if len(args) > 0 {
		if number, ok := toFloat64(args[0]); ok {
			width = number
		}
	}
	height := float64(0)
	if len(args) > 1 {
		if number, ok := toFloat64(args[1]); ok {
			height = number
		}
	}
	return m.createWidget(widgetType, map[string]interface{}{
		"parent":  parent,
		"width":   width,
		"height":  height,
		"visible": true,
	})
}

func (m *ImGuiModule) showWindowValue(handle string) {
	m.setVisibilityValue(handle, true)
}

func (m *ImGuiModule) destroyWindowValue(handle string) {
	m.removeShapeValue(handle)
}

func (m *ImGuiModule) setShapeTextColorValue(handle string, color float64) {
	m.setWidgetProperty(handle, "textColor", color)
}

func (m *ImGuiModule) setShapeTextBackgroundValue(handle string, color float64, hasBackground bool) {
	m.setWidgetProperties(handle, map[string]interface{}{
		"bgColor":       color,
		"hasBackground": hasBackground,
	})
}

func (m *ImGuiModule) setShapeTextFontScaleValue(handle string, scale float64) {
	m.setWidgetProperty(handle, "fontScale", scale)
}

func (m *ImGuiModule) setBitmapShapeValue(handle, bitmap string) {
	m.setWidgetProperty(handle, "bitmap", bitmap)
}

func (m *ImGuiModule) setShapeThicknessValue(handle string, thickness float64) {
	m.setWidgetProperty(handle, "thickness", thickness)
}

func (m *ImGuiModule) setWidgetSizeValue(handle string, width, height float64) {
	m.setWidgetProperties(handle, map[string]interface{}{
		"width":  width,
		"height": height,
	})
}

func (m *ImGuiModule) setLayoutBorderVisibleValue(handle string, visible bool) {
	m.setWidgetProperty(handle, "borderVisible", visible)
}

func (m *ImGuiModule) setWindowPosValue(handle string, x, y float64) {
	m.setWidgetProperties(handle, map[string]interface{}{
		"x": x,
		"y": y,
	})
}

func (m *ImGuiModule) setWindowSizeValue(handle string, width, height float64) {
	m.setWidgetSizeValue(handle, width, height)
}

func (m *ImGuiModule) setWidgetStyleValue(handle string, style int, value float64) {
	m.setWidgetProperties(handle, map[string]interface{}{
		"style":      style,
		"styleValue": value,
	})
}

func (m *ImGuiModule) setWidgetColorValue(handle string, color float64) {
	m.setWidgetProperty(handle, "color", color)
}

func (m *ImGuiModule) setWindowFlagsValue(handle string, flags int) {
	m.setWidgetProperty(handle, "flags", flags)
}

func (m *ImGuiModule) getWindowPosValue(handle string) map[string]float64 {
	m.widgetMutex.RLock()
	defer m.widgetMutex.RUnlock()
	widget, exists := m.widgets[handle]
	if !exists {
		return map[string]float64{"x": 0, "y": 0}
	}
	x, _ := toFloat64(widget.properties["x"])
	y, _ := toFloat64(widget.properties["y"])
	return map[string]float64{"x": x, "y": y}
}

func (m *ImGuiModule) createImageValue(parent, path string, args ...interface{}) string {
	width := float64(0)
	if len(args) > 0 {
		if number, ok := toFloat64(args[0]); ok {
			width = number
		}
	}
	height := float64(0)
	if len(args) > 1 {
		if number, ok := toFloat64(args[1]); ok {
			height = number
		}
	}
	return m.createWidget("image", map[string]interface{}{
		"parent":  parent,
		"path":    path,
		"width":   width,
		"height":  height,
		"visible": true,
	})
}

func (m *ImGuiModule) setImageValue(handle, path string) {
	m.setWidgetProperty(handle, "path", path)
}

func (m *ImGuiModule) setImageFromBitmapValue(handle, bitmap string) {
	m.setWidgetProperty(handle, "bitmap", bitmap)
}

func (m *ImGuiModule) setColorThemeValue(theme int) {
	m.widgetMutex.Lock()
	defer m.widgetMutex.Unlock()
	m.widgets["__theme"] = &Widget{
		handle:     "__theme",
		widgetType: "theme",
		properties: map[string]interface{}{"theme": theme},
		callbacks:  make(map[string]lua.LGFunction),
	}
}

func (m *ImGuiModule) setStyleColorValue(index int, color float64) {
	m.widgetMutex.Lock()
	defer m.widgetMutex.Unlock()
	m.widgets[fmt.Sprintf("__style_color_%d", index)] = &Widget{
		handle:     fmt.Sprintf("__style_color_%d", index),
		widgetType: "styleColor",
		properties: map[string]interface{}{"index": index, "color": color},
		callbacks:  make(map[string]lua.LGFunction),
	}
}

func (m *ImGuiModule) setWidgetProperty(handle, key string, value interface{}) {
	m.setWidgetProperties(handle, map[string]interface{}{key: value})
}

func (m *ImGuiModule) setWidgetProperties(handle string, properties map[string]interface{}) {
	m.widgetMutex.Lock()
	defer m.widgetMutex.Unlock()
	widget, exists := m.widgets[handle]
	if !exists {
		return
	}
	for key, value := range properties {
		widget.properties[key] = value
	}
}

func (m *ImGuiModule) setVisibilityValue(handle string, visible bool) {
	m.widgetMutex.Lock()
	defer m.widgetMutex.Unlock()
	if widget, exists := m.widgets[handle]; exists {
		widget.properties["visible"] = visible
	}
}

func (m *ImGuiModule) isVisibleValue(handle string) bool {
	m.widgetMutex.RLock()
	defer m.widgetMutex.RUnlock()
	widget, exists := m.widgets[handle]
	if !exists {
		return false
	}
	visible, ok := widget.properties["visible"].(bool)
	return ok && visible
}

func toFloat64(value interface{}) (float64, bool) {
	switch typedValue := value.(type) {
	case float64:
		return typedValue, true
	case float32:
		return float64(typedValue), true
	case int:
		return float64(typedValue), true
	case int64:
		return float64(typedValue), true
	default:
		return 0, false
	}
}

func toInt(value interface{}) (int, bool) {
	switch typedValue := value.(type) {
	case int:
		return typedValue, true
	case int64:
		return int(typedValue), true
	case float64:
		return int(typedValue), true
	default:
		return 0, false
	}
}

// isValidHandle 检查句柄是否有效
func (m *ImGuiModule) isValidHandle(L *lua.LState) int {
	handle := L.CheckString(1)
	m.widgetMutex.RLock()
	defer m.widgetMutex.RUnlock()

	_, exists := m.widgets[handle]
	L.Push(lua.LBool(exists))
	return 1
}

// createButton 创建按钮
func (m *ImGuiModule) createButton(L *lua.LState) int {
	x := L.CheckNumber(1)
	y := L.CheckNumber(2)
	width := L.CheckNumber(3)
	height := L.CheckNumber(4)
	text := L.CheckString(5)

	handle := m.generateHandle()

	m.widgetMutex.Lock()
	m.widgets[handle] = &Widget{
		handle:     handle,
		widgetType: "button",
		properties: map[string]interface{}{
			"x":      x,
			"y":      y,
			"width":  width,
			"height": height,
			"text":   text,
		},
		callbacks: make(map[string]lua.LGFunction),
	}
	m.widgetMutex.Unlock()

	L.Push(lua.LString(handle))
	return 1
}

// createCheckBox 创建复选框
func (m *ImGuiModule) createCheckBox(L *lua.LState) int {
	parent := L.CheckString(1)
	label := L.CheckString(2)
	checked := false
	if L.GetTop() >= 3 {
		checked = L.CheckBool(3)
	}

	handle := m.generateHandle()

	m.widgetMutex.Lock()
	m.widgets[handle] = &Widget{
		handle:     handle,
		widgetType: "checkbox",
		properties: map[string]interface{}{
			"parent":  parent,
			"label":   label,
			"checked": checked,
		},
		callbacks: make(map[string]lua.LGFunction),
	}
	m.widgetMutex.Unlock()

	L.Push(lua.LString(handle))
	return 1
}

// createColorPicker 创建颜色选择器
func (m *ImGuiModule) createColorPicker(L *lua.LState) int {
	parent := L.CheckString(1)
	title := "Color"
	if L.GetTop() >= 2 {
		title = L.CheckString(2)
	}
	color := uint32(0xFF000000)
	if L.GetTop() >= 3 {
		color = uint32(L.CheckNumber(3))
	}
	width := float32(0)
	if L.GetTop() >= 4 {
		width = float32(L.CheckNumber(4))
	}
	height := float32(0)
	if L.GetTop() >= 5 {
		height = float32(L.CheckNumber(5))
	}

	handle := m.generateHandle()

	m.widgetMutex.Lock()
	m.widgets[handle] = &Widget{
		handle:     handle,
		widgetType: "colorpicker",
		properties: map[string]interface{}{
			"parent": parent,
			"title":  title,
			"color":  color,
			"width":  width,
			"height": height,
		},
		callbacks: make(map[string]lua.LGFunction),
	}
	m.widgetMutex.Unlock()

	L.Push(lua.LString(handle))
	return 1
}

// createSwitch 创建开关控件
func (m *ImGuiModule) createSwitch(L *lua.LState) int {
	parent := L.CheckString(1)
	label := L.CheckString(2)
	checked := false
	if L.GetTop() >= 3 {
		checked = L.CheckBool(3)
	}
	height := float32(0)
	if L.GetTop() >= 4 {
		height = float32(L.CheckNumber(4))
	}

	handle := m.generateHandle()

	m.widgetMutex.Lock()
	m.widgets[handle] = &Widget{
		handle:     handle,
		widgetType: "switch",
		properties: map[string]interface{}{
			"parent":  parent,
			"label":   label,
			"checked": checked,
			"height":  height,
		},
		callbacks: make(map[string]lua.LGFunction),
	}
	m.widgetMutex.Unlock()

	L.Push(lua.LString(handle))
	return 1
}

// createLabel 创建文本标签
func (m *ImGuiModule) createLabel(L *lua.LState) int {
	parent := L.CheckString(1)
	text := L.CheckString(2)
	singleline := true
	if L.GetTop() >= 3 {
		singleline = L.CheckBool(3)
	}

	handle := m.generateHandle()

	m.widgetMutex.Lock()
	m.widgets[handle] = &Widget{
		handle:     handle,
		widgetType: "label",
		properties: map[string]interface{}{
			"parent":     parent,
			"text":       text,
			"singleline": singleline,
		},
		callbacks: make(map[string]lua.LGFunction),
	}
	m.widgetMutex.Unlock()

	L.Push(lua.LString(handle))
	return 1
}

// createInputText 创建输入框
func (m *ImGuiModule) createInputText(L *lua.LState) int {
	parent := L.CheckString(1)
	label := L.CheckString(2)
	value := ""
	if L.GetTop() >= 3 {
		value = L.CheckString(3)
	}
	inputType := 0
	if L.GetTop() >= 4 {
		inputType = L.CheckInt(4)
	}
	width := float32(0)
	if L.GetTop() >= 5 {
		width = float32(L.CheckNumber(5))
	}
	height := float32(0)
	if L.GetTop() >= 6 {
		height = float32(L.CheckNumber(6))
	}

	handle := m.generateHandle()

	m.widgetMutex.Lock()
	m.widgets[handle] = &Widget{
		handle:     handle,
		widgetType: "inputtext",
		properties: map[string]interface{}{
			"parent":    parent,
			"label":     label,
			"value":     value,
			"inputType": inputType,
			"width":     width,
			"height":    height,
		},
		callbacks: make(map[string]lua.LGFunction),
	}
	m.widgetMutex.Unlock()

	L.Push(lua.LString(handle))
	return 1
}

// createProgressBar 创建进度条
func (m *ImGuiModule) createProgressBar(L *lua.LState) int {
	parent := L.CheckString(1)
	progress := float32(0)
	if L.GetTop() >= 2 {
		progress = float32(L.CheckNumber(2))
	}
	width := float32(0)
	if L.GetTop() >= 3 {
		width = float32(L.CheckNumber(3))
	}
	height := float32(0)
	if L.GetTop() >= 4 {
		height = float32(L.CheckNumber(4))
	}

	handle := m.generateHandle()

	m.widgetMutex.Lock()
	m.widgets[handle] = &Widget{
		handle:     handle,
		widgetType: "progressbar",
		properties: map[string]interface{}{
			"parent":   parent,
			"progress": progress,
			"width":    width,
			"height":   height,
		},
		callbacks: make(map[string]lua.LGFunction),
	}
	m.widgetMutex.Unlock()

	L.Push(lua.LString(handle))
	return 1
}

// createComboBox 创建组合框
func (m *ImGuiModule) createComboBox(L *lua.LState) int {
	parent := L.CheckString(1)
	items := ""
	if L.GetTop() >= 2 {
		items = L.CheckString(2)
	}
	width := float32(0)
	if L.GetTop() >= 3 {
		width = float32(L.CheckNumber(3))
	}

	handle := m.generateHandle()

	m.widgetMutex.Lock()
	m.widgets[handle] = &Widget{
		handle:     handle,
		widgetType: "combobox",
		properties: map[string]interface{}{
			"parent":        parent,
			"items":         items,
			"width":         width,
			"selectedIndex": -1,
		},
		callbacks: make(map[string]lua.LGFunction),
	}
	m.widgetMutex.Unlock()

	L.Push(lua.LString(handle))
	return 1
}

// createRadioGroup 创建单选按钮组
func (m *ImGuiModule) createRadioGroup(L *lua.LState) int {
	parent := L.CheckString(1)
	label := L.CheckString(2)

	handle := m.generateHandle()

	m.widgetMutex.Lock()
	m.widgets[handle] = &Widget{
		handle:     handle,
		widgetType: "radiogroup",
		properties: map[string]interface{}{
			"parent":        parent,
			"label":         label,
			"items":         make([]string, 0),
			"selectedIndex": -1,
		},
		callbacks: make(map[string]lua.LGFunction),
	}
	m.widgetMutex.Unlock()

	L.Push(lua.LString(handle))
	return 1
}

// createTableView 创建表格视图
func (m *ImGuiModule) createTableView(L *lua.LState) int {
	parent := L.CheckString(1)
	title := L.CheckString(2)
	columns := L.CheckInt(3)
	showheader := L.CheckBool(4)
	width := float32(0)
	if L.GetTop() >= 5 {
		width = float32(L.CheckNumber(5))
	}
	height := float32(0)
	if L.GetTop() >= 6 {
		height = float32(L.CheckNumber(6))
	}

	handle := m.generateHandle()

	m.widgetMutex.Lock()
	m.widgets[handle] = &Widget{
		handle:     handle,
		widgetType: "tableview",
		properties: map[string]interface{}{
			"parent":     parent,
			"title":      title,
			"columns":    columns,
			"showheader": showheader,
			"width":      width,
			"height":     height,
			"rows":       make([][]string, 0),
			"headers":    make([]string, columns),
		},
		callbacks: make(map[string]lua.LGFunction),
	}
	m.widgetMutex.Unlock()

	L.Push(lua.LString(handle))
	return 1
}

// createSlider 创建滑动条
func (m *ImGuiModule) createSlider(L *lua.LState) int {
	parent := L.CheckString(1)
	label := L.CheckString(2)
	min := L.CheckNumber(3)
	max := L.CheckNumber(4)
	initialPos := L.CheckNumber(5)
	width := float32(0)
	if L.GetTop() >= 6 {
		width = float32(L.CheckNumber(6))
	}

	handle := m.generateHandle()

	m.widgetMutex.Lock()
	m.widgets[handle] = &Widget{
		handle:     handle,
		widgetType: "slider",
		properties: map[string]interface{}{
			"parent":     parent,
			"label":      label,
			"min":        min,
			"max":        max,
			"initialPos": initialPos,
			"width":      width,
			"value":      float64(initialPos),
		},
		callbacks: make(map[string]lua.LGFunction),
	}
	m.widgetMutex.Unlock()

	L.Push(lua.LString(handle))
	return 1
}

// createBitmapShape 创建位图形状
func (m *ImGuiModule) createBitmapShape(L *lua.LState) int {
	x := L.CheckNumber(1)
	y := L.CheckNumber(2)
	width := L.CheckNumber(3)
	height := L.CheckNumber(4)
	bitmap := L.CheckString(5)

	handle := m.generateHandle()

	m.widgetMutex.Lock()
	m.widgets[handle] = &Widget{
		handle:     handle,
		widgetType: "bitmap",
		properties: map[string]interface{}{
			"x":      x,
			"y":      y,
			"width":  width,
			"height": height,
			"bitmap": bitmap,
		},
		callbacks: make(map[string]lua.LGFunction),
	}
	m.widgetMutex.Unlock()

	L.Push(lua.LString(handle))
	return 1
}

// setOnClick 设置按钮点击回调
func (m *ImGuiModule) setOnClick(L *lua.LState) int {
	handle := L.CheckString(1)
	callback := L.CheckFunction(2)

	m.widgetMutex.Lock()
	defer m.widgetMutex.Unlock()

	widget, exists := m.widgets[handle]
	if !exists {
		return 0
	}

	widget.callbacks["onclick"] = func(L *lua.LState) int {
		L.Push(callback)
		L.Push(lua.LString(handle))
		L.Call(1, 0)
		return 0
	}

	return 0
}

// setOnCheck 设置复选框状态变更回调
func (m *ImGuiModule) setOnCheck(L *lua.LState) int {
	handle := L.CheckString(1)
	callback := L.CheckFunction(2)

	m.widgetMutex.Lock()
	defer m.widgetMutex.Unlock()

	widget, exists := m.widgets[handle]
	if !exists {
		return 0
	}

	widget.callbacks["oncheck"] = func(L *lua.LState) int {
		L.Push(callback)
		L.Push(lua.LString(handle))
		m.widgetMutex.RLock()
		checked := widget.properties["checked"].(bool)
		m.widgetMutex.RUnlock()
		L.Push(lua.LBool(checked))
		L.Call(2, 0)
		return 0
	}

	return 0
}

// setOnSelectEvent 设置组合框/单选框选择事件回调
func (m *ImGuiModule) setOnSelectEvent(L *lua.LState) int {
	handle := L.CheckString(1)
	callback := L.CheckFunction(2)

	m.widgetMutex.Lock()
	defer m.widgetMutex.Unlock()

	widget, exists := m.widgets[handle]
	if !exists {
		return 0
	}

	widget.callbacks["onselect"] = func(L *lua.LState) int {
		L.Push(callback)
		L.Push(lua.LString(handle))
		m.widgetMutex.RLock()
		index := widget.properties["selectedIndex"].(int)
		m.widgetMutex.RUnlock()
		L.Push(lua.LNumber(index))
		L.Call(2, 0)
		return 0
	}

	return 0
}

// setOnSelectEventEx 设置表格选择事件回调
func (m *ImGuiModule) setOnSelectEventEx(L *lua.LState) int {
	handle := L.CheckString(1)
	callback := L.CheckFunction(2)

	m.widgetMutex.Lock()
	defer m.widgetMutex.Unlock()

	widget, exists := m.widgets[handle]
	if !exists {
		return 0
	}

	widget.callbacks["onselectex"] = func(L *lua.LState) int {
		L.Push(callback)
		L.Push(lua.LString(handle))
		L.Push(lua.LNumber(0))
		L.Push(lua.LNumber(0))
		L.Push(lua.LString(""))
		L.Call(4, 0)
		return 0
	}

	return 0
}

// setChecked 设置复选框/开关控件的选中状态
func (m *ImGuiModule) setChecked(L *lua.LState) int {
	handle := m.convertHandle(L.CheckAny(1))
	if handle == "" {
		return 0
	}
	state := L.CheckBool(2)

	m.widgetMutex.Lock()
	defer m.widgetMutex.Unlock()

	widget, exists := m.widgets[handle]
	if !exists {
		return 0
	}

	widget.properties["checked"] = state
	return 0
}

// isChecked 获取复选框/开关控件的当前状态
func (m *ImGuiModule) isChecked(L *lua.LState) int {
	handle := m.convertHandle(L.CheckAny(1))
	if handle == "" {
		L.Push(lua.LNil)
		return 1
	}

	m.widgetMutex.RLock()
	defer m.widgetMutex.RUnlock()

	widget, exists := m.widgets[handle]
	if !exists {
		L.Push(lua.LNil)
		return 1
	}

	checked, ok := widget.properties["checked"].(bool)
	if !ok {
		L.Push(lua.LNil)
		return 1
	}

	L.Push(lua.LBool(checked))
	return 1
}

// setWidgetText 设置控件文本
func (m *ImGuiModule) setWidgetText(L *lua.LState) int {
	handle := m.convertHandle(L.CheckAny(1))
	if handle == "" {
		return 0
	}
	text := L.CheckString(2)

	m.widgetMutex.Lock()
	defer m.widgetMutex.Unlock()

	widget, exists := m.widgets[handle]
	if !exists {
		return 0
	}

	widget.properties["text"] = text
	return 0
}

// getWidgetText 获取控件文本
func (m *ImGuiModule) getWidgetText(L *lua.LState) int {
	handle := m.convertHandle(L.CheckAny(1))
	if handle == "" {
		L.Push(lua.LString(""))
		return 1
	}

	m.widgetMutex.RLock()
	defer m.widgetMutex.RUnlock()

	widget, exists := m.widgets[handle]
	if !exists {
		L.Push(lua.LString(""))
		return 1
	}

	text, ok := widget.properties["text"].(string)
	if !ok {
		L.Push(lua.LString(""))
		return 1
	}

	L.Push(lua.LString(text))
	return 1
}

// getInputText 获取输入框当前文本内容
func (m *ImGuiModule) getInputText(L *lua.LState) int {
	handle := m.convertHandle(L.CheckAny(1))
	if handle == "" {
		L.Push(lua.LString(""))
		return 1
	}

	m.widgetMutex.RLock()
	defer m.widgetMutex.RUnlock()

	widget, exists := m.widgets[handle]
	if !exists {
		L.Push(lua.LString(""))
		return 1
	}

	value, ok := widget.properties["value"].(string)
	if !ok {
		L.Push(lua.LString(""))
		return 1
	}

	L.Push(lua.LString(value))
	return 1
}

// setInputText 设置输入框文本内容
func (m *ImGuiModule) setInputText(L *lua.LState) int {
	handle := m.convertHandle(L.CheckAny(1))
	if handle == "" {
		return 0
	}
	text := L.CheckString(2)

	m.widgetMutex.Lock()
	defer m.widgetMutex.Unlock()

	widget, exists := m.widgets[handle]
	if !exists {
		return 0
	}

	widget.properties["value"] = text
	return 0
}

// setInputType 设置输入框类型
func (m *ImGuiModule) setInputType(L *lua.LState) int {
	handle := L.CheckString(1)
	inputType := L.CheckInt(2)

	m.widgetMutex.Lock()
	defer m.widgetMutex.Unlock()

	widget, exists := m.widgets[handle]
	if !exists {
		return 0
	}

	widget.properties["inputType"] = inputType
	return 0
}

// setProgressBarPos 设置进度条当前进度值
func (m *ImGuiModule) setProgressBarPos(L *lua.LState) int {
	handle := m.convertHandle(L.CheckAny(1))
	if handle == "" {
		return 0
	}
	progress := L.CheckNumber(2)

	m.widgetMutex.Lock()
	defer m.widgetMutex.Unlock()

	widget, exists := m.widgets[handle]
	if !exists {
		return 0
	}

	widget.properties["progress"] = float32(progress)
	return 0
}

// getProgressBarPos 获取进度条当前进度值
func (m *ImGuiModule) getProgressBarPos(L *lua.LState) int {
	handle := m.convertHandle(L.CheckAny(1))
	if handle == "" {
		L.Push(lua.LNumber(0))
		return 1
	}

	m.widgetMutex.RLock()
	defer m.widgetMutex.RUnlock()

	widget, exists := m.widgets[handle]
	if !exists {
		L.Push(lua.LNumber(0))
		return 1
	}

	progress, ok := widget.properties["progress"].(float32)
	if !ok {
		L.Push(lua.LNumber(0))
		return 1
	}

	L.Push(lua.LNumber(float64(progress)))
	return 1
}

// getItemText 获取下拉项文本
func (m *ImGuiModule) getItemText(L *lua.LState) int {
	handle := L.CheckString(1)
	itemIndex := L.CheckInt(2)

	m.widgetMutex.RLock()
	defer m.widgetMutex.RUnlock()

	widget, exists := m.widgets[handle]
	if !exists {
		L.Push(lua.LString(""))
		return 1
	}

	items, ok := widget.properties["items"].(string)
	if !ok {
		L.Push(lua.LString(""))
		return 1
	}

	itemList := splitItems(items)
	if itemIndex < 0 || itemIndex >= len(itemList) {
		L.Push(lua.LNil)
		return 1
	}

	L.Push(lua.LString(itemList[itemIndex]))
	return 1
}

// removeItemAt 删除组合框指定项
func (m *ImGuiModule) removeItemAt(L *lua.LState) int {
	handle := L.CheckString(1)
	itemIndex := L.CheckInt(2)

	m.widgetMutex.Lock()
	defer m.widgetMutex.Unlock()

	widget, exists := m.widgets[handle]
	if !exists {
		return 0
	}

	items, ok := widget.properties["items"].(string)
	if !ok {
		return 0
	}

	itemList := splitItems(items)
	if itemIndex < 0 || itemIndex >= len(itemList) {
		return 0
	}

	newItemList := append(itemList[:itemIndex], itemList[itemIndex+1:]...)
	widget.properties["items"] = joinItems(newItemList)

	return 0
}

// removeAllItems 清空组合框所有项
func (m *ImGuiModule) removeAllItems(L *lua.LState) int {
	handle := L.CheckString(1)

	m.widgetMutex.Lock()
	defer m.widgetMutex.Unlock()

	widget, exists := m.widgets[handle]
	if !exists {
		return 0
	}

	widget.properties["items"] = ""
	widget.properties["selectedIndex"] = -1

	return 0
}

// getSelectedItemIndex 获取组合框|表被选中的项
func (m *ImGuiModule) getSelectedItemIndex(L *lua.LState) int {
	handle := m.convertHandle(L.CheckAny(1))
	if handle == "" {
		L.Push(lua.LNumber(-2))
		return 1
	}

	m.widgetMutex.RLock()
	defer m.widgetMutex.RUnlock()

	widget, exists := m.widgets[handle]
	if !exists {
		L.Push(lua.LNumber(-2))
		return 1
	}

	index, ok := widget.properties["selectedIndex"].(int)
	if !ok {
		L.Push(lua.LNumber(-2))
		return 1
	}

	L.Push(lua.LNumber(index))
	return 1
}

// setItemSelected 设置组合框|表选中项
func (m *ImGuiModule) setItemSelected(L *lua.LState) int {
	handle := m.convertHandle(L.CheckAny(1))
	if handle == "" {
		return 0
	}
	index := L.CheckInt(2)

	m.widgetMutex.Lock()
	defer m.widgetMutex.Unlock()

	widget, exists := m.widgets[handle]
	if !exists {
		return 0
	}

	widget.properties["selectedIndex"] = index
	return 0
}

// getItemCount 获取表格行数,或者组合框子项数，单选框控件子项数
func (m *ImGuiModule) getItemCount(L *lua.LState) int {
	handle := m.convertHandle(L.CheckAny(1))
	if handle == "" {
		L.Push(lua.LNumber(-1))
		return 1
	}

	m.widgetMutex.RLock()
	defer m.widgetMutex.RUnlock()

	widget, exists := m.widgets[handle]
	if !exists {
		L.Push(lua.LNumber(-1))
		return 1
	}

	switch widget.widgetType {
	case "combobox":
		items, ok := widget.properties["items"].(string)
		if !ok {
			L.Push(lua.LNumber(-1))
			return 1
		}
		itemList := splitItems(items)
		L.Push(lua.LNumber(len(itemList)))
	case "radiogroup":
		items, ok := widget.properties["items"].([]string)
		if !ok {
			L.Push(lua.LNumber(-1))
			return 1
		}
		L.Push(lua.LNumber(len(items)))
	case "tableview":
		rows, ok := widget.properties["rows"].([][]string)
		if !ok {
			L.Push(lua.LNumber(-1))
			return 1
		}
		L.Push(lua.LNumber(len(rows)))
	default:
		L.Push(lua.LNumber(-1))
	}

	return 1
}

// setTableHeaderItem 设置表头文本
func (m *ImGuiModule) setTableHeaderItem(L *lua.LState) int {
	handle := L.CheckString(1)
	col := L.CheckInt(2)
	text := L.CheckString(3)

	m.widgetMutex.Lock()
	defer m.widgetMutex.Unlock()

	widget, exists := m.widgets[handle]
	if !exists {
		return 0
	}

	headers, ok := widget.properties["headers"].([]string)
	if !ok {
		return 0
	}

	if col < 0 || col >= len(headers) {
		return 0
	}

	headers[col] = text
	widget.properties["headers"] = headers

	return 0
}

// insertTableRow 插入表格行
func (m *ImGuiModule) insertTableRow(L *lua.LState) int {
	handle := L.CheckString(1)
	after := L.CheckInt(2)

	m.widgetMutex.Lock()
	defer m.widgetMutex.Unlock()

	widget, exists := m.widgets[handle]
	if !exists {
		L.Push(lua.LNil)
		return 1
	}

	rows, ok := widget.properties["rows"].([][]string)
	if !ok {
		L.Push(lua.LNil)
		return 1
	}

	columns := widget.properties["columns"].(int)
	newRow := make([]string, columns)

	var newRows [][]string
	switch after {
	case -1:
		newRows = append([][]string{newRow}, rows...)
	case -2:
		newRows = append(rows, newRow)
	default:
		if after < 0 || after >= len(rows) {
			L.Push(lua.LNil)
			return 1
		}
		newRows = append(rows[:after+1], append([][]string{newRow}, rows[after+1:]...)...)
	}

	widget.properties["rows"] = newRows
	L.Push(lua.LNumber(len(newRows) - 1))
	return 1
}

// getTableItemText 获取表格单元格文本
func (m *ImGuiModule) getTableItemText(L *lua.LState) int {
	handle := m.convertHandle(L.CheckAny(1))
	if handle == "" {
		L.Push(lua.LString(""))
		return 1
	}
	row := L.CheckInt(2)
	col := L.CheckInt(3)

	m.widgetMutex.RLock()
	defer m.widgetMutex.RUnlock()

	widget, exists := m.widgets[handle]
	if !exists {
		L.Push(lua.LString(""))
		return 1
	}

	rows, ok := widget.properties["rows"].([][]string)
	if !ok {
		L.Push(lua.LString(""))
		return 1
	}

	if row < 0 || row >= len(rows) {
		L.Push(lua.LString(""))
		return 1
	}

	rowData := rows[row]
	if col < 0 || col >= len(rowData) {
		L.Push(lua.LString(""))
		return 1
	}

	L.Push(lua.LString(rowData[col]))
	return 1
}

// setTableItemText 设置表格单元格文本
func (m *ImGuiModule) setTableItemText(L *lua.LState) int {
	handle := L.CheckString(1)
	row := L.CheckInt(2)
	col := L.CheckInt(3)
	text := L.CheckString(4)

	m.widgetMutex.Lock()
	defer m.widgetMutex.Unlock()

	widget, exists := m.widgets[handle]
	if !exists {
		return 0
	}

	rows, ok := widget.properties["rows"].([][]string)
	if !ok {
		return 0
	}

	if row < 0 || row >= len(rows) {
		return 0
	}

	rowData := rows[row]
	if col < 0 || col >= len(rowData) {
		return 0
	}

	rowData[col] = text
	rows[row] = rowData
	widget.properties["rows"] = rows

	return 0
}

// deleteTableRow 删除表格行
func (m *ImGuiModule) deleteTableRow(L *lua.LState) int {
	handle := L.CheckString(1)
	row := L.CheckInt(2)

	m.widgetMutex.Lock()
	defer m.widgetMutex.Unlock()

	widget, exists := m.widgets[handle]
	if !exists {
		return 0
	}

	rows, ok := widget.properties["rows"].([][]string)
	if !ok {
		return 0
	}

	if row < 0 || row >= len(rows) {
		return 0
	}

	newRows := append(rows[:row], rows[row+1:]...)
	widget.properties["rows"] = newRows

	return 0
}

// clearTable 清空表格所有行
func (m *ImGuiModule) clearTable(L *lua.LState) int {
	handle := L.CheckString(1)

	m.widgetMutex.Lock()
	defer m.widgetMutex.Unlock()

	widget, exists := m.widgets[handle]
	if !exists {
		return 0
	}

	widget.properties["rows"] = make([][]string, 0)

	return 0
}

// setSlider 设置滑动条位置
func (m *ImGuiModule) setSlider(L *lua.LState) int {
	handle := m.convertHandle(L.CheckAny(1))
	if handle == "" {
		return 0
	}
	pos := L.CheckNumber(2)

	m.widgetMutex.Lock()
	defer m.widgetMutex.Unlock()

	widget, exists := m.widgets[handle]
	if !exists {
		return 0
	}

	widget.properties["value"] = float64(pos)
	return 0
}

// getSlider 获取滑动条位置
func (m *ImGuiModule) getSlider(L *lua.LState) int {
	handle := m.convertHandle(L.CheckAny(1))
	if handle == "" {
		L.Push(lua.LNumber(0))
		return 1
	}

	m.widgetMutex.RLock()
	defer m.widgetMutex.RUnlock()

	widget, exists := m.widgets[handle]
	if !exists {
		L.Push(lua.LNumber(0))
		return 1
	}

	var value float64
	valueRaw, exists := widget.properties["value"]
	if !exists {
		L.Push(lua.LNumber(0))
		return 1
	}

	switch v := valueRaw.(type) {
	case float64:
		value = v
	case float32:
		value = float64(v)
	case int:
		value = float64(v)
	case int64:
		value = float64(v)
	default:
		L.Push(lua.LNumber(0))
		return 1
	}

	L.Push(lua.LNumber(value))
	return 1
}

// addRadioBox 添加单选框
func (m *ImGuiModule) addRadioBox(L *lua.LState) int {
	handle := L.CheckString(1)
	text := L.CheckString(2)
	_ = L.CheckBool(3)

	m.widgetMutex.Lock()
	defer m.widgetMutex.Unlock()

	widget, exists := m.widgets[handle]
	if !exists {
		return 0
	}

	items, ok := widget.properties["items"].([]string)
	if !ok {
		return 0
	}

	widget.properties["items"] = append(items, text)

	return 0
}

// addOptionItem 添加选项
func (m *ImGuiModule) addOptionItem(L *lua.LState) int {
	handle := L.CheckString(1)
	itemText := L.CheckString(2)

	m.widgetMutex.Lock()
	defer m.widgetMutex.Unlock()

	widget, exists := m.widgets[handle]
	if !exists {
		return 0
	}

	items, ok := widget.properties["items"].(string)
	if !ok {
		return 0
	}

	if items == "" {
		widget.properties["items"] = itemText
	} else {
		widget.properties["items"] = items + "|" + itemText
	}

	return 0
}

// addTabBarItem 向标签栏添加标签页
func (m *ImGuiModule) addTabBarItem(L *lua.LState) int {
	tabBarHandle := L.CheckString(1)
	title := L.CheckString(2)

	m.widgetMutex.Lock()
	defer m.widgetMutex.Unlock()

	widget, exists := m.widgets[tabBarHandle]
	if !exists {
		return 0
	}

	widget.properties["title"] = title

	return 0
}

// setStyleColor 设置样式颜色
func (m *ImGuiModule) setStyleColor(L *lua.LState) int {
	_ = L.CheckInt(1)
	_ = L.CheckNumber(2)
	return 0
}

// close 关闭imgui框架
func (m *ImGuiModule) close(L *lua.LState) int {
	return 0
}

// createRectangle 创建矩形
func (m *ImGuiModule) createRectangle(L *lua.LState) int {
	x := float32(L.CheckNumber(1))
	y := float32(L.CheckNumber(2))
	x1 := float32(L.CheckNumber(3))
	y1 := float32(L.CheckNumber(4))
	color := uint32(L.CheckNumber(5))
	filled := L.CheckBool(6)
	rounding := float32(0)
	if L.GetTop() >= 7 {
		rounding = float32(L.CheckNumber(7))
	}

	handle := m.generateHandle()

	m.widgetMutex.Lock()
	m.widgets[handle] = &Widget{
		handle:     handle,
		widgetType: "rectangle",
		properties: map[string]interface{}{
			"x":        x,
			"y":        y,
			"x1":       x1,
			"y1":       y1,
			"color":    color,
			"filled":   filled,
			"rounding": rounding,
			"visible":  true,
		},
		callbacks: make(map[string]lua.LGFunction),
	}
	m.widgetMutex.Unlock()

	L.Push(lua.LString(handle))
	return 1
}

// createCircle 创建圆形
func (m *ImGuiModule) createCircle(L *lua.LState) int {
	x := float32(L.CheckNumber(1))
	y := float32(L.CheckNumber(2))
	radius := float32(L.CheckNumber(3))
	color := uint32(L.CheckNumber(4))
	filled := L.CheckBool(5)
	segments := int32(0)
	if L.GetTop() >= 6 {
		segments = int32(L.CheckInt(6))
	}

	handle := m.generateHandle()

	m.widgetMutex.Lock()
	m.widgets[handle] = &Widget{
		handle:     handle,
		widgetType: "circle",
		properties: map[string]interface{}{
			"x":        x,
			"y":        y,
			"radius":   radius,
			"color":    color,
			"filled":   filled,
			"segments": segments,
			"visible":  true,
		},
		callbacks: make(map[string]lua.LGFunction),
	}
	m.widgetMutex.Unlock()

	L.Push(lua.LString(handle))
	return 1
}

// createPolygon 创建多边形
func (m *ImGuiModule) createPolygon(L *lua.LState) int {
	points := L.CheckTable(1)
	color := uint32(L.CheckNumber(2))
	closed := L.CheckBool(3)
	filled := L.CheckBool(4)
	thickness := float32(1)
	if L.GetTop() >= 5 {
		thickness = float32(L.CheckNumber(5))
	}

	pointList := make([][2]float32, 0)
	points.ForEach(func(key, value lua.LValue) {
		if value.Type() == lua.LTTable {
			table := value.(*lua.LTable)
			x := float32(table.RawGetInt(1).(lua.LNumber))
			y := float32(table.RawGetInt(2).(lua.LNumber))
			pointList = append(pointList, [2]float32{x, y})
		}
	})

	handle := m.generateHandle()

	m.widgetMutex.Lock()
	m.widgets[handle] = &Widget{
		handle:     handle,
		widgetType: "polygon",
		properties: map[string]interface{}{
			"points":    pointList,
			"color":     color,
			"closed":    closed,
			"filled":    filled,
			"thickness": thickness,
			"visible":   true,
		},
		callbacks: make(map[string]lua.LGFunction),
	}
	m.widgetMutex.Unlock()

	L.Push(lua.LString(handle))
	return 1
}

// createLine 创建线段
func (m *ImGuiModule) createLine(L *lua.LState) int {
	x1 := float32(L.CheckNumber(1))
	y1 := float32(L.CheckNumber(2))
	x2 := float32(L.CheckNumber(3))
	y2 := float32(L.CheckNumber(4))
	color := uint32(L.CheckNumber(5))
	thickness := float32(1)
	if L.GetTop() >= 6 {
		thickness = float32(L.CheckNumber(6))
	}

	handle := m.generateHandle()

	m.widgetMutex.Lock()
	m.widgets[handle] = &Widget{
		handle:     handle,
		widgetType: "line",
		properties: map[string]interface{}{
			"x1":        x1,
			"y1":        y1,
			"x2":        x2,
			"y2":        y2,
			"color":     color,
			"thickness": thickness,
			"visible":   true,
		},
		callbacks: make(map[string]lua.LGFunction),
	}
	m.widgetMutex.Unlock()

	L.Push(lua.LString(handle))
	return 1
}

// createShapeText 创建文本图形
func (m *ImGuiModule) createShapeText(L *lua.LState) int {
	x := float32(L.CheckNumber(1))
	y := float32(L.CheckNumber(2))
	w := float32(L.CheckNumber(3))
	h := float32(L.CheckNumber(4))
	text := L.CheckString(5)
	textColor := uint32(L.CheckNumber(6))
	bgColor := uint32(L.CheckNumber(7))
	hasBackground := L.CheckBool(8)
	fontScale := float32(1)
	if L.GetTop() >= 9 {
		fontScale = float32(L.CheckNumber(9))
	}

	handle := m.generateHandle()

	m.widgetMutex.Lock()
	m.widgets[handle] = &Widget{
		handle:     handle,
		widgetType: "shapetext",
		properties: map[string]interface{}{
			"x":             x,
			"y":             y,
			"w":             w,
			"h":             h,
			"text":          text,
			"textColor":     textColor,
			"bgColor":       bgColor,
			"hasBackground": hasBackground,
			"fontScale":     fontScale,
			"visible":       true,
		},
		callbacks: make(map[string]lua.LGFunction),
	}
	m.widgetMutex.Unlock()

	L.Push(lua.LString(handle))
	return 1
}

// setShapePosition 设置图形位置
func (m *ImGuiModule) setShapePosition(L *lua.LState) int {
	handle := L.CheckString(1)
	x := float32(L.CheckNumber(2))
	y := float32(L.CheckNumber(3))

	m.widgetMutex.Lock()
	defer m.widgetMutex.Unlock()

	widget, exists := m.widgets[handle]
	if !exists {
		return 0
	}

	widget.properties["x"] = x
	widget.properties["y"] = y
	return 0
}

// setShapeVisibility 设置图形可见性
func (m *ImGuiModule) setShapeVisibility(L *lua.LState) int {
	handle := L.CheckString(1)
	visible := L.CheckBool(2)

	m.widgetMutex.Lock()
	defer m.widgetMutex.Unlock()

	widget, exists := m.widgets[handle]
	if !exists {
		return 0
	}

	widget.properties["visible"] = visible
	return 0
}

// isShapeVisibility 判断图形是否可见
func (m *ImGuiModule) isShapeVisibility(L *lua.LState) int {
	handle := L.CheckString(1)

	m.widgetMutex.RLock()
	defer m.widgetMutex.RUnlock()

	widget, exists := m.widgets[handle]
	if !exists {
		L.Push(lua.LBool(false))
		return 1
	}

	visible, ok := widget.properties["visible"].(bool)
	if !ok {
		L.Push(lua.LBool(false))
		return 1
	}

	L.Push(lua.LBool(visible))
	return 1
}

// setShapeTextString 修改文本内容
func (m *ImGuiModule) setShapeTextString(L *lua.LState) int {
	handle := L.CheckString(1)
	newText := L.CheckString(2)

	m.widgetMutex.Lock()
	defer m.widgetMutex.Unlock()

	widget, exists := m.widgets[handle]
	if !exists {
		return 0
	}

	widget.properties["text"] = newText
	return 0
}

// setShapeTextColor 设置文本颜色
func (m *ImGuiModule) setShapeTextColor(L *lua.LState) int {
	handle := L.CheckString(1)
	newColor := uint32(L.CheckNumber(2))

	m.widgetMutex.Lock()
	defer m.widgetMutex.Unlock()

	widget, exists := m.widgets[handle]
	if !exists {
		return 0
	}

	widget.properties["textColor"] = newColor
	return 0
}

// setShapeTextBackground 设置文本背景
func (m *ImGuiModule) setShapeTextBackground(L *lua.LState) int {
	handle := L.CheckString(1)
	bgColor := uint32(L.CheckNumber(2))
	hasBg := L.CheckBool(3)

	m.widgetMutex.Lock()
	defer m.widgetMutex.Unlock()

	widget, exists := m.widgets[handle]
	if !exists {
		return 0
	}

	widget.properties["bgColor"] = bgColor
	widget.properties["hasBackground"] = hasBg
	return 0
}

// setShapeTextFontScale 设置文本图形字体缩放比例
func (m *ImGuiModule) setShapeTextFontScale(L *lua.LState) int {
	handle := L.CheckString(1)
	scale := float32(L.CheckNumber(2))

	m.widgetMutex.Lock()
	defer m.widgetMutex.Unlock()

	widget, exists := m.widgets[handle]
	if !exists {
		return 0
	}

	widget.properties["fontScale"] = scale
	return 0
}

// removeShape 删除图形对象
func (m *ImGuiModule) removeShape(L *lua.LState) int {
	handle := L.CheckString(1)

	m.widgetMutex.Lock()
	defer m.widgetMutex.Unlock()

	delete(m.widgets, handle)
	return 0
}

// setBitmapShape 设置位图形状
func (m *ImGuiModule) setBitmapShape(L *lua.LState) int {
	handle := L.CheckString(1)
	bitmap := L.CheckString(2)

	m.widgetMutex.Lock()
	defer m.widgetMutex.Unlock()

	widget, exists := m.widgets[handle]
	if !exists {
		return 0
	}

	widget.properties["bitmap"] = bitmap
	return 0
}

// setShapeThickness 设置形状边框厚度
func (m *ImGuiModule) setShapeThickness(L *lua.LState) int {
	handle := L.CheckString(1)
	thickness := float32(L.CheckNumber(2))

	m.widgetMutex.Lock()
	defer m.widgetMutex.Unlock()

	widget, exists := m.widgets[handle]
	if !exists {
		return 0
	}

	widget.properties["thickness"] = thickness
	return 0
}

// setWidgetSize 设置控件尺寸
func (m *ImGuiModule) setWidgetSize(L *lua.LState) int {
	handle := L.CheckString(1)
	width := float32(L.CheckNumber(2))
	height := float32(L.CheckNumber(3))

	m.widgetMutex.Lock()
	defer m.widgetMutex.Unlock()

	widget, exists := m.widgets[handle]
	if !exists {
		return 0
	}

	widget.properties["width"] = width
	widget.properties["height"] = height
	return 0
}

// setWidgetVisible 设置控件可见性
func (m *ImGuiModule) setWidgetVisible(L *lua.LState) int {
	handle := L.CheckString(1)
	visible := L.CheckBool(2)

	m.widgetMutex.Lock()
	defer m.widgetMutex.Unlock()

	widget, exists := m.widgets[handle]
	if !exists {
		return 0
	}

	widget.properties["visible"] = visible
	return 0
}

// isWidgetVisible 获取控件当前可见状态
func (m *ImGuiModule) isWidgetVisible(L *lua.LState) int {
	handle := L.CheckString(1)

	m.widgetMutex.RLock()
	defer m.widgetMutex.RUnlock()

	widget, exists := m.widgets[handle]
	if !exists {
		L.Push(lua.LBool(false))
		return 1
	}

	visible, ok := widget.properties["visible"].(bool)
	if !ok {
		L.Push(lua.LBool(true))
		return 1
	}

	L.Push(lua.LBool(visible))
	return 1
}

// setLayoutBorderVisible 设置容器布局控件是否显示边框
func (m *ImGuiModule) setLayoutBorderVisible(L *lua.LState) int {
	handle := L.CheckString(1)
	visible := L.CheckBool(2)

	m.widgetMutex.Lock()
	defer m.widgetMutex.Unlock()

	widget, exists := m.widgets[handle]
	if !exists {
		return 0
	}

	widget.properties["borderVisible"] = visible
	return 0
}

// setWindowPos 设置窗口位置
func (m *ImGuiModule) setWindowPos(L *lua.LState) int {
	handle := L.CheckString(1)
	x := float32(L.CheckNumber(2))
	y := float32(L.CheckNumber(3))

	m.widgetMutex.Lock()
	defer m.widgetMutex.Unlock()

	widget, exists := m.widgets[handle]
	if !exists {
		return 0
	}

	widget.properties["x"] = x
	widget.properties["y"] = y
	return 0
}

// setWindowSize 设置窗口大小
func (m *ImGuiModule) setWindowSize(L *lua.LState) int {
	handle := L.CheckString(1)
	width := float32(L.CheckNumber(2))
	height := float32(L.CheckNumber(3))

	m.widgetMutex.Lock()
	defer m.widgetMutex.Unlock()

	widget, exists := m.widgets[handle]
	if !exists {
		return 0
	}

	widget.properties["width"] = width
	widget.properties["height"] = height
	return 0
}

// setWidgetStyle 设置imgui控件一些属性
func (m *ImGuiModule) setWidgetStyle(L *lua.LState) int {
	handle := L.CheckString(1)
	style := L.CheckInt(2)
	v1 := L.CheckNumber(3)
	v2 := float32(0)
	if L.GetTop() >= 4 {
		v2 = float32(L.CheckNumber(4))
	}

	m.widgetMutex.Lock()
	defer m.widgetMutex.Unlock()

	widget, exists := m.widgets[handle]
	if !exists {
		return 0
	}

	widget.properties["style"] = style
	widget.properties["styleV1"] = v1
	widget.properties["styleV2"] = v2
	return 0
}

// setWidgetColor 设置imgui控件的相关颜色
func (m *ImGuiModule) setWidgetColor(L *lua.LState) int {
	handle := L.CheckString(1)
	colorType := L.CheckInt(2)
	color := uint32(L.CheckNumber(3))

	m.widgetMutex.Lock()
	defer m.widgetMutex.Unlock()

	widget, exists := m.widgets[handle]
	if !exists {
		return 0
	}

	widget.properties["colorType"] = colorType
	widget.properties["color"] = color
	return 0
}

// setWindowFlags 设置窗口标志
func (m *ImGuiModule) setWindowFlags(L *lua.LState) int {
	handle := L.CheckString(1)
	flags := L.CheckInt(2)

	m.widgetMutex.Lock()
	defer m.widgetMutex.Unlock()

	widget, exists := m.widgets[handle]
	if !exists {
		return 0
	}

	widget.properties["flags"] = flags
	return 0
}

// getWindowPos 获取窗口位置和尺寸
func (m *ImGuiModule) getWindowPos(L *lua.LState) int {
	handle := L.CheckString(1)

	m.widgetMutex.RLock()
	defer m.widgetMutex.RUnlock()

	widget, exists := m.widgets[handle]
	if !exists {
		return 0
	}

	x, _ := widget.properties["x"].(float32)
	y, _ := widget.properties["y"].(float32)
	width, _ := widget.properties["width"].(float32)
	height, _ := widget.properties["height"].(float32)

	L.Push(lua.LNumber(x))
	L.Push(lua.LNumber(y))
	L.Push(lua.LNumber(width))
	L.Push(lua.LNumber(height))
	return 4
}

// setOnClose 设置窗口关闭回调函数
func (m *ImGuiModule) setOnClose(L *lua.LState) int {
	handle := L.CheckString(1)
	callback := L.CheckFunction(2)

	m.widgetMutex.Lock()
	defer m.widgetMutex.Unlock()

	widget, exists := m.widgets[handle]
	if !exists {
		return 0
	}

	widget.callbacks["onclose"] = func(L *lua.LState) int {
		L.Push(callback)
		L.Push(lua.LString(handle))
		L.Call(1, 0)
		return 0
	}

	return 0
}

// destroyWindow 销毁一个imgui窗口
func (m *ImGuiModule) destroyWindow(L *lua.LState) int {
	handle := L.CheckString(1)

	m.widgetMutex.Lock()
	defer m.widgetMutex.Unlock()

	delete(m.widgets, handle)
	return 0
}

// sameLine 设置控件同行布局
func (m *ImGuiModule) sameLine(L *lua.LState) int {
	handle := L.CheckString(1)
	spacing := float32(-1)
	if L.GetTop() >= 2 {
		spacing = float32(L.CheckNumber(2))
	}

	m.widgetMutex.Lock()
	defer m.widgetMutex.Unlock()

	widget, exists := m.widgets[handle]
	if !exists {
		return 0
	}

	widget.properties["sameLine"] = true
	widget.properties["spacing"] = spacing
	return 0
}

// createWindow 创建一个imgui窗口
func (m *ImGuiModule) createWindow(L *lua.LState) int {
	title := L.CheckString(1)
	x := float32(L.CheckNumber(2))
	y := float32(L.CheckNumber(3))
	width := float32(L.CheckNumber(4))
	height := float32(L.CheckNumber(5))
	showclose := L.CheckBool(6)

	handle := m.generateHandle()

	m.widgetMutex.Lock()
	m.widgets[handle] = &Widget{
		handle:     handle,
		widgetType: "window",
		properties: map[string]interface{}{
			"title":     title,
			"x":         x,
			"y":         y,
			"width":     width,
			"height":    height,
			"showclose": showclose,
			"visible":   true,
		},
		callbacks: make(map[string]lua.LGFunction),
	}
	m.widgetMutex.Unlock()

	L.Push(lua.LString(handle))
	return 1
}

// createVerticalLayout 创建垂直布局容器
func (m *ImGuiModule) createVerticalLayout(L *lua.LState) int {
	parent := L.CheckString(1)
	width := float32(0)
	if L.GetTop() >= 2 {
		width = float32(L.CheckNumber(2))
	}
	height := float32(0)
	if L.GetTop() >= 3 {
		height = float32(L.CheckNumber(3))
	}

	handle := m.generateHandle()

	m.widgetMutex.Lock()
	m.widgets[handle] = &Widget{
		handle:     handle,
		widgetType: "verticallayout",
		properties: map[string]interface{}{
			"parent":  parent,
			"width":   width,
			"height":  height,
			"visible": true,
		},
		callbacks: make(map[string]lua.LGFunction),
	}
	m.widgetMutex.Unlock()

	L.Push(lua.LString(handle))
	return 1
}

// createHorticalLayout 创建水平布局容器
func (m *ImGuiModule) createHorticalLayout(L *lua.LState) int {
	parent := L.CheckString(1)
	width := float32(0)
	if L.GetTop() >= 2 {
		width = float32(L.CheckNumber(2))
	}
	height := float32(0)
	if L.GetTop() >= 3 {
		height = float32(L.CheckNumber(3))
	}

	handle := m.generateHandle()

	m.widgetMutex.Lock()
	m.widgets[handle] = &Widget{
		handle:     handle,
		widgetType: "horticallayout",
		properties: map[string]interface{}{
			"parent":  parent,
			"width":   width,
			"height":  height,
			"visible": true,
		},
		callbacks: make(map[string]lua.LGFunction),
	}
	m.widgetMutex.Unlock()

	L.Push(lua.LString(handle))
	return 1
}

// createTreeBoxLayout 创建树形布局容器
func (m *ImGuiModule) createTreeBoxLayout(L *lua.LState) int {
	parent := L.CheckString(1)
	title := L.CheckString(2)
	width := float32(0)
	if L.GetTop() >= 3 {
		width = float32(L.CheckNumber(3))
	}

	handle := m.generateHandle()

	m.widgetMutex.Lock()
	m.widgets[handle] = &Widget{
		handle:     handle,
		widgetType: "treebox",
		properties: map[string]interface{}{
			"parent":  parent,
			"title":   title,
			"width":   width,
			"visible": true,
		},
		callbacks: make(map[string]lua.LGFunction),
	}
	m.widgetMutex.Unlock()

	L.Push(lua.LString(handle))
	return 1
}

// createTabBar 创建标签栏控件
func (m *ImGuiModule) createTabBar(L *lua.LState) int {
	parent := L.CheckString(1)
	title := L.CheckString(2)

	handle := m.generateHandle()

	m.widgetMutex.Lock()
	m.widgets[handle] = &Widget{
		handle:     handle,
		widgetType: "tabbar",
		properties: map[string]interface{}{
			"parent":  parent,
			"title":   title,
			"tabs":    make([]string, 0),
			"visible": true,
		},
		callbacks: make(map[string]lua.LGFunction),
	}
	m.widgetMutex.Unlock()

	L.Push(lua.LString(handle))
	return 1
}

// createImage 创建图片显示控件
func (m *ImGuiModule) createImage(L *lua.LState) int {
	parent := L.CheckString(1)
	path := L.CheckString(2)
	width := float32(0)
	if L.GetTop() >= 3 {
		width = float32(L.CheckNumber(3))
	}
	height := float32(0)
	if L.GetTop() >= 4 {
		height = float32(L.CheckNumber(4))
	}

	handle := m.generateHandle()

	m.widgetMutex.Lock()
	m.widgets[handle] = &Widget{
		handle:     handle,
		widgetType: "image",
		properties: map[string]interface{}{
			"parent":  parent,
			"path":    path,
			"width":   width,
			"height":  height,
			"visible": true,
		},
		callbacks: make(map[string]lua.LGFunction),
	}
	m.widgetMutex.Unlock()

	L.Push(lua.LString(handle))
	return 1
}

// setImage 设置或更换显示图片
func (m *ImGuiModule) setImage(L *lua.LState) int {
	handle := L.CheckString(1)
	path := L.CheckString(2)

	m.widgetMutex.Lock()
	defer m.widgetMutex.Unlock()

	widget, exists := m.widgets[handle]
	if !exists {
		return 0
	}

	widget.properties["path"] = path
	return 0
}

// setImageFromBitmap 设置图像对象
func (m *ImGuiModule) setImageFromBitmap(L *lua.LState) int {
	handle := L.CheckString(1)
	bitmap := L.CheckString(2)

	m.widgetMutex.Lock()
	defer m.widgetMutex.Unlock()

	widget, exists := m.widgets[handle]
	if !exists {
		return 0
	}

	widget.properties["bitmap"] = bitmap
	return 0
}

// setColorTheme 设置imgui颜色主题
func (m *ImGuiModule) setColorTheme(L *lua.LState) int {
	style := L.CheckInt(1)

	m.widgetMutex.Lock()
	defer m.widgetMutex.Unlock()

	for _, widget := range m.widgets {
		widget.properties["theme"] = style
	}

	return 0
}

// show 显示UI框架
func (m *ImGuiModule) show(L *lua.LState) int {
	touchable := false
	if L.GetTop() >= 1 {
		touchable = L.CheckBool(1)
	}
	font := ""
	if L.GetTop() >= 2 {
		font = L.CheckString(2)
	}
	fontsize := float32(0)
	if L.GetTop() >= 3 {
		fontsize = float32(L.CheckNumber(3))
	}

	m.widgetMutex.Lock()
	defer m.widgetMutex.Unlock()

	for _, widget := range m.widgets {
		widget.properties["touchable"] = touchable
		widget.properties["font"] = font
		widget.properties["fontsize"] = fontsize
		widget.properties["visible"] = true
	}

	return 0
}

// showWindow 创建并显示独立窗口
func (m *ImGuiModule) showWindow(L *lua.LState) int {
	config := L.CheckTable(1)

	title := ""
	if config.RawGetString("title") != lua.LNil {
		title = config.RawGetString("title").String()
	}

	x := float32(0)
	if config.RawGetString("x") != lua.LNil {
		x = float32(config.RawGetString("x").(lua.LNumber))
	}

	y := float32(0)
	if config.RawGetString("y") != lua.LNil {
		y = float32(config.RawGetString("y").(lua.LNumber))
	}

	width := float32(0)
	if config.RawGetString("width") != lua.LNil {
		width = float32(config.RawGetString("width").(lua.LNumber))
	}

	height := float32(0)
	if config.RawGetString("height") != lua.LNil {
		height = float32(config.RawGetString("height").(lua.LNumber))
	}

	handle := m.generateHandle()

	m.widgetMutex.Lock()
	m.widgets[handle] = &Widget{
		handle:     handle,
		widgetType: "window",
		properties: map[string]interface{}{
			"title":   title,
			"x":       x,
			"y":       y,
			"width":   width,
			"height":  height,
			"visible": true,
		},
		callbacks: make(map[string]lua.LGFunction),
	}
	m.widgetMutex.Unlock()

	L.Push(lua.LString(handle))
	return 1
}

// setOnSliderEvent 设置滑动条值变化事件回调
func (m *ImGuiModule) setOnSliderEvent(L *lua.LState) int {
	handle := L.CheckString(1)
	callback := L.CheckFunction(2)

	m.widgetMutex.Lock()
	defer m.widgetMutex.Unlock()

	widget, exists := m.widgets[handle]
	if !exists {
		return 0
	}

	widget.callbacks["onslider"] = func(L *lua.LState) int {
		L.Push(callback)
		L.Push(lua.LString(handle))
		m.widgetMutex.RLock()
		value := widget.properties["value"]
		m.widgetMutex.RUnlock()
		L.Push(lua.LNumber(value.(float64)))
		L.Call(2, 0)
		return 0
	}

	return 0
}

// splitItems 分割选项字符串
func splitItems(items string) []string {
	if items == "" {
		return []string{}
	}
	result := []string{}
	current := ""
	for _, c := range items {
		if c == '|' {
			result = append(result, current)
			current = ""
		} else {
			current += string(c)
		}
	}
	if current != "" {
		result = append(result, current)
	}
	return result
}

// joinItems 连接选项字符串
func joinItems(items []string) string {
	if len(items) == 0 {
		return ""
	}
	result := items[0]
	for i := 1; i < len(items); i++ {
		result += "|" + items[i]
	}
	return result
}

// GetModule 获取模块实例
func GetModule() model.Module {
	return New()
}
