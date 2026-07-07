//go:build ignore
// +build ignore

package plugin

import (
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"

	"github.com/Dasongzi1366/AutoGo/plugin"
	lua "github.com/yuin/gopher-lua"
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
	state := engine.GetState()

	pluginObj := state.NewTable()
	state.SetGlobal("plugin", pluginObj)

	pluginObj.RawSetString("loadApk", state.NewFunction(func(L *lua.LState) int {
		path := L.CheckString(1)
		result := plugin.LoadApk(path)
		if result != nil {
			L.Push(wrapClassLoader(L, result))
		} else {
			L.Push(lua.LNil)
		}
		return 1
	}))

	pluginObj.RawSetString("newContext", state.NewFunction(func(L *lua.LState) int {
		ud := L.NewUserData()
		ud.Value = plugin.NewContext()
		L.Push(ud)
		return 1
	}))

	pluginObj.RawSetString("newAssetManager", state.NewFunction(func(L *lua.LState) int {
		ud := L.NewUserData()
		ud.Value = plugin.NewAssetManager()
		L.Push(ud)
		return 1
	}))

	engine.RegisterMethod("plugin.loadApk", "加载外部APK", plugin.LoadApk, true)
	engine.RegisterMethod("plugin.newContext", "创建Context参数", plugin.NewContext, true)
	engine.RegisterMethod("plugin.newAssetManager", "创建AssetManager参数", plugin.NewAssetManager, true)
	engine.RegisterMethod("plugin.newInstance", "创建类实例", func(cl *plugin.ClassLoader, className string, args ...interface{}) *plugin.Instance {
		return cl.NewInstance(className, args...)
	}, true)
	engine.RegisterMethod("plugin.callString", "调用返回String的方法", func(inst *plugin.Instance, methodName string, args ...interface{}) (string, error) {
		return inst.CallString(methodName, args...)
	}, true)
	engine.RegisterMethod("plugin.callInt", "调用返回int的方法", func(inst *plugin.Instance, methodName string, args ...interface{}) (int, error) {
		return inst.CallInt(methodName, args...)
	}, true)
	engine.RegisterMethod("plugin.callLong", "调用返回long的方法", func(inst *plugin.Instance, methodName string, args ...interface{}) (int64, error) {
		return inst.CallLong(methodName, args...)
	}, true)
	engine.RegisterMethod("plugin.callFloat", "调用返回float的方法", func(inst *plugin.Instance, methodName string, args ...interface{}) (float32, error) {
		return inst.CallFloat(methodName, args...)
	}, true)
	engine.RegisterMethod("plugin.callDouble", "调用返回double的方法", func(inst *plugin.Instance, methodName string, args ...interface{}) (float64, error) {
		return inst.CallDouble(methodName, args...)
	}, true)
	engine.RegisterMethod("plugin.callBool", "调用返回boolean的方法", func(inst *plugin.Instance, methodName string, args ...interface{}) (bool, error) {
		return inst.CallBool(methodName, args...)
	}, true)
	engine.RegisterMethod("plugin.callVoid", "调用无返回值的方法", func(inst *plugin.Instance, methodName string, args ...interface{}) error {
		return inst.CallVoid(methodName, args...)
	}, true)
	engine.RegisterMethod("plugin.release", "释放实例或类加载器", nil, true)

	return nil
}

func wrapClassLoader(L *lua.LState, cl *plugin.ClassLoader) lua.LValue {
	obj := L.NewTable()
	obj.RawSetString("newInstance", L.NewFunction(func(L *lua.LState) int {
		className := L.CheckString(1)
		args := luaArgsToInterfaces(L, 2)
		inst := cl.NewInstance(className, args...)
		if inst == nil {
			L.Push(lua.LNil)
			return 1
		}
		L.Push(wrapInstance(L, inst))
		return 1
	}))
	obj.RawSetString("release", L.NewFunction(func(L *lua.LState) int {
		cl.Release()
		return 0
	}))
	return obj
}

func wrapInstance(L *lua.LState, inst *plugin.Instance) lua.LValue {
	obj := L.NewTable()
	obj.RawSetString("callString", L.NewFunction(func(L *lua.LState) int {
		result, err := inst.CallString(L.CheckString(1), luaArgsToInterfaces(L, 2)...)
		return pushLuaResult(L, lua.LString(result), err)
	}))
	obj.RawSetString("callInt", L.NewFunction(func(L *lua.LState) int {
		result, err := inst.CallInt(L.CheckString(1), luaArgsToInterfaces(L, 2)...)
		return pushLuaResult(L, lua.LNumber(result), err)
	}))
	obj.RawSetString("callLong", L.NewFunction(func(L *lua.LState) int {
		result, err := inst.CallLong(L.CheckString(1), luaArgsToInterfaces(L, 2)...)
		return pushLuaResult(L, lua.LNumber(result), err)
	}))
	obj.RawSetString("callFloat", L.NewFunction(func(L *lua.LState) int {
		result, err := inst.CallFloat(L.CheckString(1), luaArgsToInterfaces(L, 2)...)
		return pushLuaResult(L, lua.LNumber(result), err)
	}))
	obj.RawSetString("callDouble", L.NewFunction(func(L *lua.LState) int {
		result, err := inst.CallDouble(L.CheckString(1), luaArgsToInterfaces(L, 2)...)
		return pushLuaResult(L, lua.LNumber(result), err)
	}))
	obj.RawSetString("callBool", L.NewFunction(func(L *lua.LState) int {
		result, err := inst.CallBool(L.CheckString(1), luaArgsToInterfaces(L, 2)...)
		return pushLuaResult(L, lua.LBool(result), err)
	}))
	obj.RawSetString("callVoid", L.NewFunction(func(L *lua.LState) int {
		err := inst.CallVoid(L.CheckString(1), luaArgsToInterfaces(L, 2)...)
		return pushLuaResult(L, lua.LBool(err == nil), err)
	}))
	obj.RawSetString("release", L.NewFunction(func(L *lua.LState) int {
		inst.Release()
		return 0
	}))
	return obj
}

func luaArgsToInterfaces(L *lua.LState, start int) []interface{} {
	args := make([]interface{}, 0, max(L.GetTop()-start+1, 0))
	for i := start; i <= L.GetTop(); i++ {
		args = append(args, luaValueToInterface(L.Get(i)))
	}
	return args
}

func luaValueToInterface(value lua.LValue) interface{} {
	switch v := value.(type) {
	case lua.LBool:
		return bool(v)
	case lua.LNumber:
		return float64(v)
	case lua.LString:
		return string(v)
	case *lua.LUserData:
		return v.Value
	default:
		return value
	}
}

func pushLuaResult(L *lua.LState, value lua.LValue, err error) int {
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}
	L.Push(value)
	return 1
}
