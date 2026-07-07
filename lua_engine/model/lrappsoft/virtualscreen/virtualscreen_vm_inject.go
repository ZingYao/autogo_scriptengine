package virtualscreen

import "github.com/ZingYao/autogo_scriptengine/lua_engine/model"

// VirtualScreenModule 是 go-lua-vm 迁移后的模块壳。
type VirtualScreenModule struct{}

func New() *VirtualScreenModule { return &VirtualScreenModule{} }

func (m *VirtualScreenModule) Name() string { return "virtualscreen" }

func (m *VirtualScreenModule) IsAvailable() bool { return true }

func (m *VirtualScreenModule) Register(engine model.Engine) error { return nil }

func GetModule() model.Module { return &VirtualScreenModule{} }
