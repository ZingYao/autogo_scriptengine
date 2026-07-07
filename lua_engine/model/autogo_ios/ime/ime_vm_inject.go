package ime

import "github.com/ZingYao/autogo_scriptengine/lua_engine/model"

// ImeModule 是 go-lua-vm 迁移后的模块壳。
type ImeModule struct{}

func New() *ImeModule { return &ImeModule{} }

func (m *ImeModule) Name() string { return "ime" }

func (m *ImeModule) IsAvailable() bool { return true }

func (m *ImeModule) Register(engine model.Engine) error { return nil }

func GetModule() model.Module { return &ImeModule{} }
