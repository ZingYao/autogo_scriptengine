package storages

import (
	"github.com/ZingYao/autogo_scriptengine/js_engine/model"

	"github.com/Dasongzi1366/AutoGo/storages"
	"github.com/dop251/goja"
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
	vm := engine.GetVM()

	storagesObj := vm.NewObject()
	vm.Set("storages", storagesObj)

	storagesObj.Set("get", func(call goja.FunctionCall) goja.Value {
		table := call.Argument(0).String()
		key := call.Argument(1).String()
		result := storages.Get(table, key)
		return vm.ToValue(result)
	})

	storagesObj.Set("put", func(call goja.FunctionCall) goja.Value {
		table := call.Argument(0).String()
		key := call.Argument(1).String()
		value := call.Argument(2).String()
		storages.Put(table, key, value)
		return goja.Undefined()
	})

	storagesObj.Set("remove", func(call goja.FunctionCall) goja.Value {
		table := call.Argument(0).String()
		key := call.Argument(1).String()
		storages.Remove(table, key)
		return goja.Undefined()
	})

	storagesObj.Set("contains", func(call goja.FunctionCall) goja.Value {
		table := call.Argument(0).String()
		key := call.Argument(1).String()
		result := storages.Contains(table, key)
		return vm.ToValue(result)
	})

	storagesObj.Set("getAll", func(call goja.FunctionCall) goja.Value {
		table := call.Argument(0).String()
		result := storages.GetAll(table)
		obj := vm.NewObject()
		for key, value := range result {
			obj.Set(key, value)
		}
		return obj
	})

	storagesObj.Set("clear", func(call goja.FunctionCall) goja.Value {
		table := call.Argument(0).String()
		storages.Clear(table)
		return goja.Undefined()
	})

	engine.RegisterMethod("storages.get", "从指定表中获取键值", func(table, key string) string { return storages.Get(table, key) }, true)
	engine.RegisterMethod("storages.put", "写入键值对", func(table, key, value string) { storages.Put(table, key, value) }, true)
	engine.RegisterMethod("storages.remove", "删除指定键", func(table, key string) { storages.Remove(table, key) }, true)
	engine.RegisterMethod("storages.contains", "判断键是否存在", func(table, key string) bool { return storages.Contains(table, key) }, true)
	engine.RegisterMethod("storages.getAll", "获取所有键值对", func(table string) map[string]string { return storages.GetAll(table) }, true)
	engine.RegisterMethod("storages.clear", "清空指定表数据", func(table string) { storages.Clear(table) }, true)

	return nil
}
