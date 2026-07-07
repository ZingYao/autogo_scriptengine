package dotocr

import "github.com/ZingYao/autogo_scriptengine/lua_engine/model"

// DotocrModule 是 go-lua-vm 迁移后的模块壳。
type DotocrModule struct{}

func New() *DotocrModule { return &DotocrModule{} }

func (m *DotocrModule) Name() string { return "dotocr" }

func (m *DotocrModule) IsAvailable() bool { return true }

func (m *DotocrModule) Register(engine model.Engine) error { return nil }

func GetModule() model.Module { return &DotocrModule{} }
