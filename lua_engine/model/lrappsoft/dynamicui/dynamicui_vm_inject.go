package dynamicui

import "github.com/ZingYao/autogo_scriptengine/lua_engine/model"

// DynamicUIModule 是 go-lua-vm 迁移后的模块壳。
type DynamicUIModule struct {
	ThrowException bool
	ShowWarning    bool
	Debug          bool
}

func New() *DynamicUIModule { return &DynamicUIModule{} }

func (m *DynamicUIModule) Name() string { return "dynamicui" }

func (m *DynamicUIModule) IsAvailable() bool { return true }

func (m *DynamicUIModule) Register(engine model.Engine) error {
	engine.RegisterMethod("ui.newLayout", "创建布局", func(_ string) bool { return true }, true)
	engine.RegisterMethod("ui.addButton", "添加按钮", func(_ string, _ int, _ int, _ int) bool { return true }, true)
	engine.RegisterMethod("ui.getTableViewRowData", "获取表格行数据", func(_ string, _ int) bool { return true }, true)
	engine.RegisterMethod("ui.getTableViewSelectIndex", "获取表格选中索引", func(_ string) bool { return true }, true)
	return nil
}

func GetModule() model.Module { return &DynamicUIModule{} }
