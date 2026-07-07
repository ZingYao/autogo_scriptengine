//go:build ignore
// +build ignore

package storages

import (
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"

	autogostorages "github.com/Dasongzi1366/AutoGo/storages"
	lua "github.com/yuin/gopher-lua"
)

type StoragesModule struct{}

func (m *StoragesModule) Name() string      { return "storages" }
func (m *StoragesModule) IsAvailable() bool { return true }

func (m *StoragesModule) Register(engine model.Engine) error {
	state := engine.GetState()
	storagesObj := state.NewTable()
	state.SetGlobal("storages", storagesObj)
	storagesObj.RawSetString("get", state.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LString(autogostorages.Get(L.CheckString(1), L.CheckString(2))))
		return 1
	}))
	storagesObj.RawSetString("put", state.NewFunction(func(L *lua.LState) int {
		autogostorages.Put(L.CheckString(1), L.CheckString(2), L.CheckString(3))
		return 0
	}))
	storagesObj.RawSetString("remove", state.NewFunction(func(L *lua.LState) int {
		autogostorages.Remove(L.CheckString(1), L.CheckString(2))
		return 0
	}))
	storagesObj.RawSetString("contains", state.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LBool(autogostorages.Contains(L.CheckString(1), L.CheckString(2))))
		return 1
	}))
	storagesObj.RawSetString("getAll", state.NewFunction(func(L *lua.LState) int {
		result := L.NewTable()
		for key, value := range autogostorages.GetAll(L.CheckString(1)) {
			result.RawSetString(key, lua.LString(value))
		}
		L.Push(result)
		return 1
	}))
	storagesObj.RawSetString("clear", state.NewFunction(func(L *lua.LState) int {
		autogostorages.Clear(L.CheckString(1))
		return 0
	}))
	engine.RegisterMethod("storages.get", "从指定表中获取键值", autogostorages.Get, true)
	engine.RegisterMethod("storages.put", "写入键值对", func(table, key, value string) { autogostorages.Put(table, key, value) }, true)
	engine.RegisterMethod("storages.remove", "删除指定键", func(table, key string) { autogostorages.Remove(table, key) }, true)
	engine.RegisterMethod("storages.contains", "判断键是否存在", autogostorages.Contains, true)
	engine.RegisterMethod("storages.getAll", "获取所有键值对", autogostorages.GetAll, true)
	engine.RegisterMethod("storages.clear", "清空指定表数据", func(table string) { autogostorages.Clear(table) }, true)
	return nil
}
