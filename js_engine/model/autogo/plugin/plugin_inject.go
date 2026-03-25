package plugin

import (
	"github.com/ZingYao/autogo_scriptengine/js_engine/model"

	"github.com/Dasongzi1366/AutoGo/plugin"
	"github.com/dop251/goja"
)

// PluginModule plugin 模块
type PluginModule struct{}

// Name 返回模块名称
func (m *PluginModule) Name() string {
	return "plugin"
}

// IsAvailable 检查模块是否可用
func (m *PluginModule) IsAvailable() bool {
	return true
}

// Register 向引擎注册方法
func (m *PluginModule) Register(engine model.Engine) error {
	vm := engine.GetVM()

	pluginObj := vm.NewObject()
	vm.Set("plugin", pluginObj)

	pluginObj.Set("loadApk", func(call goja.FunctionCall) goja.Value {
		apkPath := call.Argument(0).String()
		cl := plugin.LoadApk(apkPath)
		if cl != nil {
			return wrapClassLoader(vm, cl)
		}
		return goja.Null()
	})

	pluginObj.Set("newContext", func(call goja.FunctionCall) goja.Value {
		ctx := plugin.NewContext()
		return vm.ToValue(ctx)
	})

	pluginObj.Set("newAssetManager", func(call goja.FunctionCall) goja.Value {
		am := plugin.NewAssetManager()
		return vm.ToValue(am)
	})

	engine.RegisterMethod("plugin.loadApk", "加载外部APK", plugin.LoadApk, true)
	engine.RegisterMethod("plugin.newContext", "创建Context参数", plugin.NewContext, true)
	engine.RegisterMethod("plugin.newAssetManager", "创建AssetManager参数", plugin.NewAssetManager, true)

	return nil
}

func wrapClassLoader(vm *goja.Runtime, cl *plugin.ClassLoader) goja.Value {
	obj := vm.NewObject()

	obj.Set("newInstance", func(call goja.FunctionCall) goja.Value {
		className := call.Argument(0).String()
		args := make([]interface{}, 0)
		for i := 1; i < len(call.Arguments); i++ {
			args = append(args, call.Argument(i).Export())
		}
		inst := cl.NewInstance(className, args...)
		if inst != nil {
			return wrapInstance(vm, inst)
		}
		return goja.Null()
	})

	obj.Set("release", func(call goja.FunctionCall) goja.Value {
		cl.Release()
		return goja.Undefined()
	})

	return vm.ToValue(obj)
}

func wrapInstance(vm *goja.Runtime, inst *plugin.Instance) goja.Value {
	obj := vm.NewObject()

	obj.Set("callString", func(call goja.FunctionCall) goja.Value {
		methodName := call.Argument(0).String()
		args := make([]interface{}, 0)
		for i := 1; i < len(call.Arguments); i++ {
			args = append(args, call.Argument(i).Export())
		}
		result, err := inst.CallString(methodName, args...)
		if err != nil {
			panic(err)
		}
		return vm.ToValue(result)
	})

	obj.Set("callInt", func(call goja.FunctionCall) goja.Value {
		methodName := call.Argument(0).String()
		args := make([]interface{}, 0)
		for i := 1; i < len(call.Arguments); i++ {
			args = append(args, call.Argument(i).Export())
		}
		result, err := inst.CallInt(methodName, args...)
		if err != nil {
			panic(err)
		}
		return vm.ToValue(result)
	})

	obj.Set("callLong", func(call goja.FunctionCall) goja.Value {
		methodName := call.Argument(0).String()
		args := make([]interface{}, 0)
		for i := 1; i < len(call.Arguments); i++ {
			args = append(args, call.Argument(i).Export())
		}
		result, err := inst.CallLong(methodName, args...)
		if err != nil {
			panic(err)
		}
		return vm.ToValue(result)
	})

	obj.Set("callFloat", func(call goja.FunctionCall) goja.Value {
		methodName := call.Argument(0).String()
		args := make([]interface{}, 0)
		for i := 1; i < len(call.Arguments); i++ {
			args = append(args, call.Argument(i).Export())
		}
		result, err := inst.CallFloat(methodName, args...)
		if err != nil {
			panic(err)
		}
		return vm.ToValue(result)
	})

	obj.Set("callDouble", func(call goja.FunctionCall) goja.Value {
		methodName := call.Argument(0).String()
		args := make([]interface{}, 0)
		for i := 1; i < len(call.Arguments); i++ {
			args = append(args, call.Argument(i).Export())
		}
		result, err := inst.CallDouble(methodName, args...)
		if err != nil {
			panic(err)
		}
		return vm.ToValue(result)
	})

	obj.Set("callBool", func(call goja.FunctionCall) goja.Value {
		methodName := call.Argument(0).String()
		args := make([]interface{}, 0)
		for i := 1; i < len(call.Arguments); i++ {
			args = append(args, call.Argument(i).Export())
		}
		result, err := inst.CallBool(methodName, args...)
		if err != nil {
			panic(err)
		}
		return vm.ToValue(result)
	})

	obj.Set("callVoid", func(call goja.FunctionCall) goja.Value {
		methodName := call.Argument(0).String()
		args := make([]interface{}, 0)
		for i := 1; i < len(call.Arguments); i++ {
			args = append(args, call.Argument(i).Export())
		}
		err := inst.CallVoid(methodName, args...)
		if err != nil {
			panic(err)
		}
		return goja.Undefined()
	})

	obj.Set("release", func(call goja.FunctionCall) goja.Value {
		inst.Release()
		return goja.Undefined()
	})

	return vm.ToValue(obj)
}
