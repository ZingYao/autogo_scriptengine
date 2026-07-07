package utils

import "github.com/ZingYao/autogo_scriptengine/lua_engine/model"

// UtilsModule 是 go-lua-vm 迁移后的模块壳。
type UtilsModule struct{}

func New() *UtilsModule { return &UtilsModule{} }

func (m *UtilsModule) Name() string { return "utils" }

func (m *UtilsModule) IsAvailable() bool { return true }

func (m *UtilsModule) Register(engine model.Engine) error { return nil }

func GetModule() model.Module { return &UtilsModule{} }
