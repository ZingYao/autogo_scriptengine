package storages

import (
	"github.com/ZingYao/autogo_scriptengine/js_engine/model"

	autogostorages "github.com/Dasongzi1366/AutoGo/storages"
	"github.com/ZingYao/goja"
)

// StoragesModule iOS storages 模块。
type StoragesModule struct{}

// Name 返回模块名称。
func (m *StoragesModule) Name() string {
	return "storages"
}

// IsAvailable 检查模块是否可用。
func (m *StoragesModule) IsAvailable() bool {
	return true
}

// Register 向引擎注册 iOS storages 方法。
func (m *StoragesModule) Register(engine model.Engine) error {
	vm := engine.GetVM()
	storagesObj := vm.NewObject()
	vm.Set("storages", storagesObj)

	storagesObj.Set("get", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(autogostorages.Get(call.Argument(0).String(), call.Argument(1).String()))
	})
	storagesObj.Set("put", func(call goja.FunctionCall) goja.Value {
		autogostorages.Put(call.Argument(0).String(), call.Argument(1).String(), call.Argument(2).String())
		return goja.Undefined()
	})
	storagesObj.Set("remove", func(call goja.FunctionCall) goja.Value {
		autogostorages.Remove(call.Argument(0).String(), call.Argument(1).String())
		return goja.Undefined()
	})
	storagesObj.Set("contains", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(autogostorages.Contains(call.Argument(0).String(), call.Argument(1).String()))
	})
	storagesObj.Set("getAll", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(autogostorages.GetAll(call.Argument(0).String()))
	})
	storagesObj.Set("clear", func(call goja.FunctionCall) goja.Value {
		autogostorages.Clear(call.Argument(0).String())
		return goja.Undefined()
	})

	engine.RegisterMethod("storages.get", "从指定表中获取键值", autogostorages.Get, true)
	engine.RegisterMethod("storages.put", "写入键值对", autogostorages.Put, true)
	engine.RegisterMethod("storages.remove", "删除指定键", autogostorages.Remove, true)
	engine.RegisterMethod("storages.contains", "判断键是否存在", autogostorages.Contains, true)
	engine.RegisterMethod("storages.getAll", "获取所有键值对", autogostorages.GetAll, true)
	engine.RegisterMethod("storages.clear", "清空指定表数据", autogostorages.Clear, true)
	return nil
}
