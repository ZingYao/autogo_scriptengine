package accessibility

import "github.com/ZingYao/autogo_scriptengine/lua_engine/model"

// AccessibilityModule 是 go-lua-vm 迁移后的模块壳。
type AccessibilityModule struct {
	ThrowException bool
	ShowWarning    bool
	Debug          bool
}

func New() *AccessibilityModule { return &AccessibilityModule{} }

func (m *AccessibilityModule) Name() string { return "accessibility" }

func (m *AccessibilityModule) IsAvailable() bool { return true }

func (m *AccessibilityModule) Register(engine model.Engine) error { return nil }

func GetModule() model.Module { return &AccessibilityModule{} }
