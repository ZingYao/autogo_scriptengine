package storages

import (
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"

	autogostorages "github.com/Dasongzi1366/AutoGo/storages"
)

// StoragesModule 是 go-lua-vm 迁移后的模块壳。
type StoragesModule struct{}

func New() *StoragesModule { return &StoragesModule{} }

func (m *StoragesModule) Name() string { return "storages" }

func (m *StoragesModule) IsAvailable() bool { return true }

func (m *StoragesModule) Register(engine model.Engine) error {
	engine.RegisterMethod("storages.get", "从指定表中获取键值", autogostorages.Get, true)
	engine.RegisterMethod("storages.put", "写入键值对", autogostorages.Put, true)
	engine.RegisterMethod("storages.remove", "删除指定键", autogostorages.Remove, true)
	engine.RegisterMethod("storages.contains", "判断键是否存在", autogostorages.Contains, true)
	engine.RegisterMethod("storages.getAll", "获取所有键值对", autogostorages.GetAll, true)
	engine.RegisterMethod("storages.clear", "清空指定表数据", autogostorages.Clear, true)
	return nil
}

func GetModule() model.Module { return &StoragesModule{} }
