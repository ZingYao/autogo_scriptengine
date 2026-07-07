package plugin

import (
	"errors"

	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"

	autogoplugin "github.com/Dasongzi1366/AutoGo/plugin"
)

var errNilPluginInstance = errors.New("plugin instance is nil")

// PluginModule 是 go-lua-vm 迁移后的模块壳。
type PluginModule struct{}

func New() *PluginModule { return &PluginModule{} }

func (m *PluginModule) Name() string { return "plugin" }

func (m *PluginModule) IsAvailable() bool { return true }

func (m *PluginModule) Register(engine model.Engine) error {
	engine.RegisterMethod("plugin.loadApk", "加载外部 APK", autogoplugin.LoadApk, true)
	engine.RegisterMethod("plugin.newContext", "创建 Context 参数", autogoplugin.NewContext, true)
	engine.RegisterMethod("plugin.newAssetManager", "创建 AssetManager 参数", autogoplugin.NewAssetManager, true)
	engine.RegisterMethod("plugin.newInstance", "创建类实例", func(cl *autogoplugin.ClassLoader, className string, args ...interface{}) *autogoplugin.Instance {
		if cl == nil {
			return nil
		}
		return cl.NewInstance(className, args...)
	}, true)
	engine.RegisterMethod("plugin.callString", "调用返回 String 的方法", func(inst *autogoplugin.Instance, methodName string, args ...interface{}) (string, error) {
		if inst == nil {
			return "", errNilPluginInstance
		}
		return inst.CallString(methodName, args...)
	}, true)
	engine.RegisterMethod("plugin.callInt", "调用返回 int 的方法", func(inst *autogoplugin.Instance, methodName string, args ...interface{}) (int, error) {
		if inst == nil {
			return 0, errNilPluginInstance
		}
		return inst.CallInt(methodName, args...)
	}, true)
	engine.RegisterMethod("plugin.callLong", "调用返回 long 的方法", func(inst *autogoplugin.Instance, methodName string, args ...interface{}) (int64, error) {
		if inst == nil {
			return 0, errNilPluginInstance
		}
		return inst.CallLong(methodName, args...)
	}, true)
	engine.RegisterMethod("plugin.callFloat", "调用返回 float 的方法", func(inst *autogoplugin.Instance, methodName string, args ...interface{}) (float32, error) {
		if inst == nil {
			return 0, errNilPluginInstance
		}
		return inst.CallFloat(methodName, args...)
	}, true)
	engine.RegisterMethod("plugin.callDouble", "调用返回 double 的方法", func(inst *autogoplugin.Instance, methodName string, args ...interface{}) (float64, error) {
		if inst == nil {
			return 0, errNilPluginInstance
		}
		return inst.CallDouble(methodName, args...)
	}, true)
	engine.RegisterMethod("plugin.callBool", "调用返回 boolean 的方法", func(inst *autogoplugin.Instance, methodName string, args ...interface{}) (bool, error) {
		if inst == nil {
			return false, errNilPluginInstance
		}
		return inst.CallBool(methodName, args...)
	}, true)
	engine.RegisterMethod("plugin.callVoid", "调用无返回值的方法", func(inst *autogoplugin.Instance, methodName string, args ...interface{}) error {
		if inst == nil {
			return errNilPluginInstance
		}
		return inst.CallVoid(methodName, args...)
	}, true)
	engine.RegisterMethod("plugin.releaseInstance", "释放插件实例", func(inst *autogoplugin.Instance) {
		if inst != nil {
			inst.Release()
		}
	}, true)
	engine.RegisterMethod("plugin.releaseClassLoader", "释放类加载器", func(cl *autogoplugin.ClassLoader) {
		if cl != nil {
			cl.Release()
		}
	}, true)
	engine.RegisterMethod("plugin.release", "释放插件实例或类加载器", func(value interface{}) {
		switch target := value.(type) {
		case *autogoplugin.Instance:
			if target != nil {
				target.Release()
			}
		case *autogoplugin.ClassLoader:
			if target != nil {
				target.Release()
			}
		}
	}, true)
	return nil
}

func GetModule() model.Module { return &PluginModule{} }
