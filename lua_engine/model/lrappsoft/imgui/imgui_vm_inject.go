package imgui

import (
	"fmt"
	"strings"

	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"
)

// ImGuiModule 是 go-lua-vm 迁移后的模块壳。
type ImGuiModule struct {
	nextID int
	values map[string]*imguiValue
}

type imguiValue struct {
	text     string
	checked  bool
	visible  bool
	progress float64
	slider   int
	items    []string
	selected int
	rows     [][]string
	x        int
	y        int
}

func New() *ImGuiModule { return &ImGuiModule{} }

func (m *ImGuiModule) Name() string { return "imgui" }

func (m *ImGuiModule) IsAvailable() bool { return true }

func (m *ImGuiModule) Register(engine model.Engine) error {
	m.values = make(map[string]*imguiValue)
	engine.RegisterMethod("imgui.isSupport", "返回 imgui 支持状态", func() bool { return true }, true)
	engine.RegisterMethod("imgui.getLastError", "返回最后错误", func() string { return "" }, true)
	engine.RegisterMethod("imgui.isValidHandle", "判断句柄是否有效", m.isValidHandle, true)
	engine.RegisterMethod("imgui.createButton", "创建按钮", func(_ int, _ int, _ int, _ int, text string) string {
		return m.newHandle(&imguiValue{text: text, visible: true, selected: -1})
	}, true)
	engine.RegisterMethod("imgui.getWidgetText", "获取控件文本", func(handle string) string {
		if value := m.value(handle); value != nil {
			return value.text
		}
		return ""
	}, true)
	engine.RegisterMethod("imgui.setWidgetText", "设置控件文本", func(handle, text string) {
		m.ensure(handle).text = text
	}, true)
	engine.RegisterMethod("imgui.createCheckBox", "创建复选框", func(_ string, text string, checked bool) string {
		return m.newHandle(&imguiValue{text: text, checked: checked, visible: true, selected: -1})
	}, true)
	engine.RegisterMethod("imgui.createSwitch", "创建开关", func(_ string, text string, checked bool, _ int) string {
		return m.newHandle(&imguiValue{text: text, checked: checked, visible: true, selected: -1})
	}, true)
	engine.RegisterMethod("imgui.isChecked", "获取选中状态", func(handle string) interface{} {
		if value := m.value(handle); value != nil {
			return value.checked
		}
		return nil
	}, true)
	engine.RegisterMethod("imgui.setChecked", "设置选中状态", func(handle string, checked bool) {
		m.ensure(handle).checked = checked
	}, true)
	engine.RegisterMethod("imgui.createColorPicker", "创建颜色选择器", func(_ string, text string, _ int, _ int, _ int) string {
		return m.newHandle(&imguiValue{text: text, visible: true, selected: -1})
	}, true)
	engine.RegisterMethod("imgui.createInputText", "创建输入框", func(_ string, _ string, text string) string {
		return m.newHandle(&imguiValue{text: text, visible: true, selected: -1})
	}, true)
	engine.RegisterMethod("imgui.getInputText", "获取输入文本", func(handle string) string { return m.ensure(handle).text }, true)
	engine.RegisterMethod("imgui.setInputText", "设置输入文本", func(handle, text string) { m.ensure(handle).text = text }, true)
	engine.RegisterMethod("imgui.setInputType", "设置输入类型", func(_ string, _ int) {}, true)
	engine.RegisterMethod("imgui.createProgressBar", "创建进度条", func(_ string, progress float64) string {
		return m.newHandle(&imguiValue{progress: progress, visible: true, selected: -1})
	}, true)
	engine.RegisterMethod("imgui.getProgressBarPos", "获取进度条位置", func(handle string) float64 { return m.ensure(handle).progress }, true)
	engine.RegisterMethod("imgui.setProgressBarPos", "设置进度条位置", func(handle string, progress float64) { m.ensure(handle).progress = progress }, true)
	engine.RegisterMethod("imgui.createLabel", "创建标签", func(_ string, text string) string {
		return m.newHandle(&imguiValue{text: text, visible: true, selected: -1})
	}, true)
	engine.RegisterMethod("imgui.createComboBox", "创建下拉框", func(_ string, items string, _ int) string {
		return m.newHandle(&imguiValue{items: splitImguiItems(items), selected: -1, visible: true})
	}, true)
	engine.RegisterMethod("imgui.addOptionItem", "添加选项", func(handle, item string) { m.ensure(handle).items = append(m.ensure(handle).items, item) }, true)
	engine.RegisterMethod("imgui.getItemCount", "获取条目数量", func(handle string) int {
		value := m.ensure(handle)
		if value.rows != nil {
			return len(value.rows)
		}
		return len(value.items)
	}, true)
	engine.RegisterMethod("imgui.getItemText", "获取条目文本", func(handle string, index int) string {
		items := m.ensure(handle).items
		if index < 0 || index >= len(items) {
			return ""
		}
		return items[index]
	}, true)
	engine.RegisterMethod("imgui.setItemSelected", "设置选中条目", func(handle string, index int) { m.ensure(handle).selected = index }, true)
	engine.RegisterMethod("imgui.getSelectedItemIndex", "获取选中条目", func(handle string) int { return m.ensure(handle).selected }, true)
	engine.RegisterMethod("imgui.removeItemAt", "删除条目", func(handle string, index int) {
		value := m.ensure(handle)
		if index >= 0 && index < len(value.items) {
			value.items = append(value.items[:index], value.items[index+1:]...)
		}
	}, true)
	engine.RegisterMethod("imgui.removeAllItems", "删除所有条目", func(handle string) {
		value := m.ensure(handle)
		value.items = nil
		value.selected = -1
	}, true)
	engine.RegisterMethod("imgui.createRadioGroup", "创建单选组", func(_ string, text string) string {
		return m.newHandle(&imguiValue{text: text, visible: true, selected: -1})
	}, true)
	engine.RegisterMethod("imgui.addRadioBox", "添加单选项", func(handle, text string, checked bool) {
		value := m.ensure(handle)
		value.items = append(value.items, text)
		if checked {
			value.selected = len(value.items) - 1
		}
	}, true)
	engine.RegisterMethod("imgui.createSlider", "创建滑块", func(_ string, _ string, _ int, _ int, value int) string {
		return m.newHandle(&imguiValue{slider: value, visible: true, selected: -1})
	}, true)
	engine.RegisterMethod("imgui.getSlider", "获取滑块值", func(handle string) int { return m.ensure(handle).slider }, true)
	engine.RegisterMethod("imgui.setSlider", "设置滑块值", func(handle string, value int) { m.ensure(handle).slider = value }, true)
	engine.RegisterMethod("imgui.createTableView", "创建表格", func(_ string, _ string, _ int, _ bool) string {
		return m.newHandle(&imguiValue{rows: [][]string{}, visible: true, selected: -1})
	}, true)
	engine.RegisterMethod("imgui.setTableHeaderItem", "设置表头", func(_ string, _ int, _ string) {}, true)
	engine.RegisterMethod("imgui.insertTableRow", "插入表格行", func(handle string, _ int) int {
		value := m.ensure(handle)
		value.rows = append(value.rows, []string{})
		return len(value.rows) - 1
	}, true)
	engine.RegisterMethod("imgui.setTableItemText", "设置表格单元格", func(handle string, row int, column int, text string) {
		value := m.ensure(handle)
		for len(value.rows) <= row {
			value.rows = append(value.rows, []string{})
		}
		for len(value.rows[row]) <= column {
			value.rows[row] = append(value.rows[row], "")
		}
		value.rows[row][column] = text
	}, true)
	engine.RegisterMethod("imgui.getTableItemText", "获取表格单元格", func(handle string, row int, column int) string {
		value := m.ensure(handle)
		if row < 0 || row >= len(value.rows) || column < 0 || column >= len(value.rows[row]) {
			return ""
		}
		return value.rows[row][column]
	}, true)
	engine.RegisterMethod("imgui.deleteTableRow", "删除表格行", func(handle string, row int) {
		value := m.ensure(handle)
		if row >= 0 && row < len(value.rows) {
			value.rows = append(value.rows[:row], value.rows[row+1:]...)
		}
	}, true)
	engine.RegisterMethod("imgui.clearTable", "清空表格", func(handle string) { m.ensure(handle).rows = [][]string{} }, true)
	m.registerShapeMethods(engine)
	m.registerWindowMethods(engine)
	return nil
}

func GetModule() model.Module { return &ImGuiModule{} }

func (m *ImGuiModule) registerShapeMethods(engine model.Engine) {
	engine.RegisterMethod("imgui.createRectangle", "创建矩形", func(_ int, _ int, _ int, _ int, _ int, visible bool) string {
		return m.newHandle(&imguiValue{visible: visible, selected: -1})
	}, true)
	engine.RegisterMethod("imgui.createCircle", "创建圆形", func(_ int, _ int, _ int, _ int, visible bool) string {
		return m.newHandle(&imguiValue{visible: visible, selected: -1})
	}, true)
	engine.RegisterMethod("imgui.createLine", "创建线条", func(_ int, _ int, _ int, _ int, _ int) string {
		return m.newHandle(&imguiValue{visible: true, selected: -1})
	}, true)
	engine.RegisterMethod("imgui.createShapeText", "创建形状文本", func(_ int, _ int, _ int, _ int, text string, _ int, _ int, visible bool) string {
		return m.newHandle(&imguiValue{text: text, visible: visible, selected: -1})
	}, true)
	engine.RegisterMethod("imgui.createBitmapShape", "创建位图形状", func(_ int, _ int, _ int, _ int, text string) string {
		return m.newHandle(&imguiValue{text: text, visible: true, selected: -1})
	}, true)
	engine.RegisterMethod("imgui.setShapePosition", "设置形状位置", func(handle string, x int, y int) {
		value := m.ensure(handle)
		value.x = x
		value.y = y
	}, true)
	engine.RegisterMethod("imgui.setShapeVisibility", "设置形状可见性", func(handle string, visible bool) { m.ensure(handle).visible = visible }, true)
	engine.RegisterMethod("imgui.isShapeVisibility", "获取形状可见性", func(handle string) bool { return m.ensure(handle).visible }, true)
	engine.RegisterMethod("imgui.setShapeTextString", "设置形状文本", func(handle string, text string) { m.ensure(handle).text = text }, true)
	engine.RegisterMethod("imgui.setShapeTextColor", "设置形状文本颜色", func(_ string, _ int) {}, true)
	engine.RegisterMethod("imgui.setShapeTextBackground", "设置形状文本背景", func(_ string, _ int, _ bool) {}, true)
	engine.RegisterMethod("imgui.setShapeTextFontScale", "设置形状文本缩放", func(_ string, _ float64) {}, true)
	engine.RegisterMethod("imgui.setShapeThickness", "设置形状线宽", func(_ string, _ int) {}, true)
	engine.RegisterMethod("imgui.setBitmapShape", "设置位图形状", func(_ string, _ string) {}, true)
	engine.RegisterMethod("imgui.removeShape", "移除形状", func(handle string) { delete(m.values, handle) }, true)
}

func (m *ImGuiModule) registerWindowMethods(engine model.Engine) {
	engine.RegisterMethod("imgui.setWidgetSize", "设置控件尺寸", func(_ string, _ int, _ int) {}, true)
	engine.RegisterMethod("imgui.setWidgetStyle", "设置控件样式", func(_ string, _ int, _ int) {}, true)
	engine.RegisterMethod("imgui.setWidgetColor", "设置控件颜色", func(_ string, _ int) {}, true)
	engine.RegisterMethod("imgui.setWidgetVisible", "设置控件可见性", func(handle string, visible bool) { m.ensure(handle).visible = visible }, true)
	engine.RegisterMethod("imgui.isWidgetVisible", "获取控件可见性", func(handle string) bool { return m.ensure(handle).visible }, true)
	engine.RegisterMethod("imgui.createWindow", "创建窗口", func(name string, x int, y int, _ int, _ int) string {
		return m.newHandle(&imguiValue{text: name, x: x, y: y, visible: false, selected: -1})
	}, true)
	engine.RegisterMethod("imgui.showWindow", "显示窗口", func(handle string) { m.ensure(handle).visible = true }, true)
	engine.RegisterMethod("imgui.setWindowPos", "设置窗口位置", func(handle string, x int, y int) {
		value := m.ensure(handle)
		value.x = x
		value.y = y
	}, true)
	engine.RegisterMethod("imgui.getWindowPos", "获取窗口位置", func(handle string) map[string]int {
		value := m.ensure(handle)
		return map[string]int{"x": value.x, "y": value.y}
	}, true)
	engine.RegisterMethod("imgui.setWindowSize", "设置窗口尺寸", func(_ string, _ int, _ int) {}, true)
	engine.RegisterMethod("imgui.setWindowFlags", "设置窗口标记", func(_ string, _ int) {}, true)
	engine.RegisterMethod("imgui.createVerticalLayout", "创建垂直布局", func(_ string, _ int, _ int) string { return m.newHandle(&imguiValue{visible: true, selected: -1}) }, true)
	engine.RegisterMethod("imgui.createHorticalLayout", "创建水平布局", func(_ string, _ int, _ int) string { return m.newHandle(&imguiValue{visible: true, selected: -1}) }, true)
	engine.RegisterMethod("imgui.createTreeBoxLayout", "创建树布局", func(_ string, _ int, _ int) string { return m.newHandle(&imguiValue{visible: true, selected: -1}) }, true)
	engine.RegisterMethod("imgui.createTabBar", "创建标签栏", func(_ string, _ int, _ int) string { return m.newHandle(&imguiValue{visible: true, selected: -1}) }, true)
	engine.RegisterMethod("imgui.addTabBarItem", "添加标签项", func(_ string, _ string) {}, true)
	engine.RegisterMethod("imgui.setLayoutBorderVisible", "设置布局边框", func(_ string, _ bool) {}, true)
	engine.RegisterMethod("imgui.createImage", "创建图片", func(_ string, path string, _ int, _ int) string {
		return m.newHandle(&imguiValue{text: path, visible: true, selected: -1})
	}, true)
	engine.RegisterMethod("imgui.setImage", "设置图片路径", func(handle, path string) { m.ensure(handle).text = path }, true)
	engine.RegisterMethod("imgui.setImageFromBitmap", "设置图片位图", func(handle, bitmap string) { m.ensure(handle).text = bitmap }, true)
	engine.RegisterMethod("imgui.setColorTheme", "设置颜色主题", func(_ int) {}, true)
	engine.RegisterMethod("imgui.setStyleColor", "设置样式颜色", func(_ int, _ int) {}, true)
	engine.RegisterMethod("imgui.sameLine", "同行布局", func() {}, true)
	engine.RegisterMethod("imgui.show", "显示 UI", func() {}, true)
	engine.RegisterMethod("imgui.destroyWindow", "销毁窗口", func(handle string) { delete(m.values, handle) }, true)
	engine.RegisterMethod("imgui.close", "关闭 imgui", func() {}, true)
}

func (m *ImGuiModule) newHandle(value *imguiValue) string {
	m.nextID++
	handle := fmt.Sprintf("imgui-%d", m.nextID)
	m.values[handle] = value
	return handle
}

func (m *ImGuiModule) value(handle string) *imguiValue {
	if m.values == nil {
		m.values = make(map[string]*imguiValue)
	}
	return m.values[handle]
}

func (m *ImGuiModule) ensure(handle string) *imguiValue {
	if value := m.value(handle); value != nil {
		return value
	}
	value := &imguiValue{selected: -1}
	m.values[handle] = value
	return value
}

func (m *ImGuiModule) isValidHandle(handle string) bool {
	return m.value(handle) != nil
}

func splitImguiItems(items string) []string {
	if items == "" {
		return nil
	}
	return strings.Split(items, "|")
}
