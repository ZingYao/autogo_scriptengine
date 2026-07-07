package app

import "github.com/ZingYao/autogo_scriptengine/lua_engine/model"

// AppModule 是 go-lua-vm 迁移后的模块壳。
type AppModule struct{}

func New() *AppModule { return &AppModule{} }

func (m *AppModule) Name() string { return "app" }

func (m *AppModule) IsAvailable() bool { return true }

func (m *AppModule) Register(engine model.Engine) error { return nil }

func GetModule() model.Module { return &AppModule{} }
