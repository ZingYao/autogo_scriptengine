package touch

import "github.com/ZingYao/autogo_scriptengine/lua_engine/model"

// TouchModule 是 go-lua-vm 迁移后的模块壳。
type TouchModule struct {
	ThrowException bool
	ShowWarning    bool
	Debug          bool
}

func New() *TouchModule { return &TouchModule{} }

func (m *TouchModule) Name() string { return "touch" }

func (m *TouchModule) IsAvailable() bool { return true }

func (m *TouchModule) Register(engine model.Engine) error { return nil }

func GetModule() model.Module { return &TouchModule{} }
