package extension

import "github.com/ZingYao/autogo_scriptengine/lua_engine/model"

// ExtensionModule 是 go-lua-vm 迁移后的模块壳。
type ExtensionModule struct {
	ThrowException bool
	ShowWarning    bool
	Debug          bool
}

func New() *ExtensionModule { return &ExtensionModule{} }

func (m *ExtensionModule) Name() string { return "extension" }

func (m *ExtensionModule) IsAvailable() bool { return true }

func (m *ExtensionModule) Register(engine model.Engine) error { return nil }

func GetModule() model.Module { return &ExtensionModule{} }
