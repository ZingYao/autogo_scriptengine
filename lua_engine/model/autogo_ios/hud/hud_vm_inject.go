package hud

import "github.com/ZingYao/autogo_scriptengine/lua_engine/model"

// HUDModule 是 go-lua-vm 迁移后的模块壳。
type HUDModule struct{}

func New() *HUDModule { return &HUDModule{} }

func (m *HUDModule) Name() string { return "hud" }

func (m *HUDModule) IsAvailable() bool { return true }

func (m *HUDModule) Register(engine model.Engine) error { return nil }

func GetModule() model.Module { return &HUDModule{} }
