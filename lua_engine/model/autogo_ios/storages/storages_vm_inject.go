package storages

import "github.com/ZingYao/autogo_scriptengine/lua_engine/model"

// StoragesModule 是 go-lua-vm 迁移后的模块壳。
type StoragesModule struct{}

func New() *StoragesModule { return &StoragesModule{} }

func (m *StoragesModule) Name() string { return "storages" }

func (m *StoragesModule) IsAvailable() bool { return true }

func (m *StoragesModule) Register(engine model.Engine) error { return nil }

func GetModule() model.Module { return &StoragesModule{} }
