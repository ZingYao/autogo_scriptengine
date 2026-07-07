package console

import "github.com/ZingYao/autogo_scriptengine/lua_engine/model"

// ConsoleModule 是 go-lua-vm 迁移后的模块壳。
type ConsoleModule struct{}

func New() *ConsoleModule { return &ConsoleModule{} }

func (m *ConsoleModule) Name() string { return "console" }

func (m *ConsoleModule) IsAvailable() bool { return true }

func (m *ConsoleModule) Register(engine model.Engine) error { return nil }

func GetModule() model.Module { return &ConsoleModule{} }
