package ui

import "github.com/ZingYao/autogo_scriptengine/lua_engine/model"

// UIModule 是 go-lua-vm 迁移后的模块壳。
type UIModule struct {
	ThrowException bool
	ShowWarning    bool
	Debug          bool
}

func New() *UIModule { return &UIModule{} }

func (m *UIModule) Name() string { return "ui" }

func (m *UIModule) IsAvailable() bool { return true }

func (m *UIModule) Register(engine model.Engine) error { return nil }

func GetModule() model.Module { return &UIModule{} }
