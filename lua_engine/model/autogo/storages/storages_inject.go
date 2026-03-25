package storages

import (
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"

	"github.com/Dasongzi1366/AutoGo/storages"
	lua "github.com/yuin/gopher-lua"
)

// StoragesModule storages 模块
type StoragesModule struct{}

// Name 返回模块名称
func (m *StoragesModule) Name() string {
	return "storages"
}

// IsAvailable 检查模块是否可用
func (m *StoragesModule) IsAvailable() bool {
	return true
}

// Register 向引擎注册方法
func (m *StoragesModule) Register(engine model.Engine) error {
	state := engine.GetState()

	storagesObj := state.NewTable()
	state.SetGlobal("storages", storagesObj)

	storagesObj.RawSetString("get", state.NewFunction(func(L *lua.LState) int {
		table := L.CheckString(1)
		key := L.CheckString(2)
		result := storages.Get(table, key)
		L.Push(lua.LString(result))
		return 1
	}))

	storagesObj.RawSetString("put", state.NewFunction(func(L *lua.LState) int {
		table := L.CheckString(1)
		key := L.CheckString(2)
		value := L.CheckString(3)
		storages.Put(table, key, value)
		return 0
	}))

	storagesObj.RawSetString("remove", state.NewFunction(func(L *lua.LState) int {
		table := L.CheckString(1)
		key := L.CheckString(2)
		storages.Remove(table, key)
		return 0
	}))

	storagesObj.RawSetString("contains", state.NewFunction(func(L *lua.LState) int {
		table := L.CheckString(1)
		key := L.CheckString(2)
		result := storages.Contains(table, key)
		L.Push(lua.LBool(result))
		return 1
	}))

	storagesObj.RawSetString("getAll", state.NewFunction(func(L *lua.LState) int {
		table := L.CheckString(1)
		result := storages.GetAll(table)
		obj := L.NewTable()
		for key, value := range result {
			L.SetField(obj, key, lua.LString(value))
		}
		L.Push(obj)
		return 1
	}))

	storagesObj.RawSetString("clear", state.NewFunction(func(L *lua.LState) int {
		table := L.CheckString(1)
		storages.Clear(table)
		return 0
	}))

	engine.RegisterMethod("storages.get", "从指定表中获取键值", func(table, key string) string { return storages.Get(table, key) }, true)
	engine.RegisterMethod("storages.put", "写入键值对", func(table, key, value string) { storages.Put(table, key, value) }, true)
	engine.RegisterMethod("storages.remove", "删除指定键", func(table, key string) { storages.Remove(table, key) }, true)
	engine.RegisterMethod("storages.contains", "判断键是否存在", func(table, key string) bool { return storages.Contains(table, key) }, true)
	engine.RegisterMethod("storages.getAll", "获取所有键值对", func(table string) map[string]string { return storages.GetAll(table) }, true)
	engine.RegisterMethod("storages.clear", "清空指定表数据", func(table string) { storages.Clear(table) }, true)

	return nil
}
