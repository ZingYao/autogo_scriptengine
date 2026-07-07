package system

import "github.com/ZingYao/autogo_scriptengine/lua_engine/model"

// SystemModule 是 go-lua-vm 迁移后的模块壳。
type SystemModule struct{}

func New() *SystemModule { return &SystemModule{} }

func (m *SystemModule) Name() string { return "system" }

func (m *SystemModule) IsAvailable() bool { return true }

func (m *SystemModule) Register(engine model.Engine) error { return nil }

func GetModule() model.Module { return &SystemModule{} }
