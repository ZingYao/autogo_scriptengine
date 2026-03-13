package lua_engine

import (
	"app/lua_engine/model"
	"fmt"
	"sync"
	"time"

	lua "github.com/yuin/gopher-lua"
)

var (
	engine         *LuaEngine
	once           sync.Once
	moduleRegistry *model.ModuleRegistry
)

func initModuleRegistry() {
	if moduleRegistry == nil {
		moduleRegistry = model.NewModuleRegistry()
	}
}

// RegisterModule 注册一个或多个模块到引擎
// 用户可以在自己的代码中调用此方法来注册需要的模块
// 支持可变长参数，可以一次注册多个模块
func RegisterModule(modules ...model.Module) {
	initModuleRegistry()
	for _, module := range modules {
		moduleRegistry.RegisterModule(module)
	}
}

// GetLuaEngine 获取默认引擎实例（使用默认配置，自动注入所有方法）
func GetLuaEngine() *LuaEngine {
	once.Do(func() {
		engine = &LuaEngine{
			config: DefaultConfig(),
		}
		initModuleRegistry()
		engine.init()
	})
	return engine
}

// GetEngine 获取默认引擎实例（GetLuaEngine 的别名）
func GetEngine() *LuaEngine {
	return GetLuaEngine()
}

// NewLuaEngine 创建新的引擎实例
// config: 引擎配置，传入 nil 使用默认配置
func NewLuaEngine(config *EngineConfig) *LuaEngine {
	cfg := DefaultConfig()
	if config != nil {
		cfg = *config
	}

	e := &LuaEngine{
		config: cfg,
	}
	initModuleRegistry()
	e.init()
	return e
}

// NewEngine 创建新的引擎实例（NewLuaEngine 的别名）
func NewEngine(config *EngineConfig) *LuaEngine {
	return NewLuaEngine(config)
}

func (e *LuaEngine) init() {
	e.mu.Lock()
	defer e.mu.Unlock()

	e.state = lua.NewState(lua.Options{
		SkipOpenLibs:        false,
		IncludeGoStackTrace: true,
	})

	e.registerCoreFunctions()
	if e.config.AutoInjectMethods {
		e.injectAllMethods()
	}
}

func (e *LuaEngine) GetState() *lua.LState {
	return e.state
}

func (e *LuaEngine) registerCoreFunctions() {
	state := e.state

	state.Register("registerMethod", e.registerMethodLua)
	state.Register("unregisterMethod", e.unregisterMethodLua)
	state.Register("listMethods", e.listMethodsLua)
	state.Register("overrideMethod", e.overrideMethodLua)
	state.Register("restoreMethod", e.restoreMethodLua)
	state.Register("sleep", e.sleepLua)
	state.Register("console.log", e.consoleLogLua)
	state.Register("console.error", e.consoleErrorLua)
}

func (e *LuaEngine) consoleLogLua(L *lua.LState) int {
	n := L.GetTop()
	for i := 1; i <= n; i++ {
		fmt.Print(L.ToString(i), " ")
	}
	fmt.Println()
	return 0
}

func (e *LuaEngine) consoleErrorLua(L *lua.LState) int {
	n := L.GetTop()
	fmt.Print("[ERROR] ")
	for i := 1; i <= n; i++ {
		fmt.Print(L.ToString(i), " ")
	}
	fmt.Println()
	return 0
}

func (e *LuaEngine) injectAllMethods() {
	whiteList := e.config.WhiteList
	blackList := e.config.BlackList
	failFast := e.config.FailFast

	modules := moduleRegistry.ListModules()

	for _, name := range modules {
		module, ok := moduleRegistry.GetModule(name)
		if !ok {
			continue
		}

		// 检查白名单
		if len(whiteList) > 0 {
			found := false
			for _, w := range whiteList {
				if w == name {
					found = true
					break
				}
			}
			if !found {
				continue
			}
		}

		// 检查黑名单
		blacklisted := false
		for _, b := range blackList {
			if b == name {
				blacklisted = true
				break
			}
		}
		if blacklisted {
			continue
		}

		// 检查模块是否可用
		if !module.IsAvailable() {
			if failFast {
				panic(fmt.Sprintf("module %s is not available", name))
			} else {
				fmt.Printf("[WARN] module %s is not available, skipping\n", name)
				continue
			}
		}

		// 注册模块
		err := module.Register(e)
		if err != nil {
			if failFast {
				panic(fmt.Sprintf("failed to register module %s: %v", name, err))
			} else {
				fmt.Printf("[WARN] failed to register module %s: %v, skipping\n", name, err)
				continue
			}
		}

		fmt.Printf("[INFO] module %s registered successfully\n", name)
	}
}

// InjectModule 注入指定模块的方法
// module: 模块名称，支持: app, device, motion, files, images, storages, system, http, media, opencv, ppocr, console, dotocr, hud, ime, plugin, rhino, uiacc, utils, vdisplay, yolo, imgui
func (e *LuaEngine) InjectModule(moduleName string) {
	e.mu.Lock()
	defer e.mu.Unlock()

	module, ok := moduleRegistry.GetModule(moduleName)
	if !ok {
		panic(fmt.Sprintf("unknown module: %s", moduleName))
	}

	if !module.IsAvailable() {
		panic(fmt.Sprintf("module %s is not available", moduleName))
	}

	err := module.Register(e)
	if err != nil {
		panic(fmt.Sprintf("failed to register module %s: %v", moduleName, err))
	}

	fmt.Printf("[INFO] module %s registered successfully\n", moduleName)
}

// InjectModules 注入多个模块的方法
func (e *LuaEngine) InjectModules(modules []string) {
	for _, module := range modules {
		e.InjectModule(module)
	}
}

// GetAvailableModules 获取所有可用模块列表
func (e *LuaEngine) GetAvailableModules() []string {
	return moduleRegistry.ListModules()
}

// InjectAllMethods 注入所有方法（公开方法，允许手动调用）
func (e *LuaEngine) InjectAllMethods() {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.injectAllMethods()
}

func (e *LuaEngine) RegisterMethod(name, description string, goFunc interface{}, overridable bool) {
	RegisterMethod(name, description, goFunc, overridable)
}

func (e *LuaEngine) ExecuteString(script string) error {
	e.mu.Lock()
	defer e.mu.Unlock()

	if e.state == nil {
		return fmt.Errorf("Lua engine not initialized")
	}

	return e.state.DoString(script)
}

func (e *LuaEngine) ExecuteFile(path string) error {
	e.mu.Lock()
	defer e.mu.Unlock()

	if e.state == nil {
		return fmt.Errorf("Lua engine not initialized")
	}

	return e.state.DoFile(path)
}

func (e *LuaEngine) Close() {
	e.mu.Lock()
	defer e.mu.Unlock()

	if e.state != nil {
		e.state.Close()
		e.state = nil
	}
}

func (e *LuaEngine) GetRegistry() *MethodRegistry {
	return GetRegistry()
}

func ExecuteString(script string) error {
	if engine != nil {
		return engine.ExecuteString(script)
	}
	return fmt.Errorf("Lua engine not initialized")
}

func ExecuteFile(path string) error {
	if engine != nil {
		return engine.ExecuteFile(path)
	}
	return fmt.Errorf("Lua engine not initialized")
}

func Close() {
	if engine != nil {
		engine.Close()
	}
}

func (e *LuaEngine) registerMethodLua(L *lua.LState) int {
	name := L.CheckString(1)
	description := L.CheckString(2)
	overridable := L.CheckBool(3)

	e.RegisterMethod(name, description, nil, overridable)
	L.Push(lua.LBool(true))
	return 1
}

func (e *LuaEngine) unregisterMethodLua(L *lua.LState) int {
	name := L.CheckString(1)
	result := registry.RemoveMethod(name)
	L.Push(lua.LBool(result))
	return 1
}

func (e *LuaEngine) listMethodsLua(L *lua.LState) int {
	methods := registry.ListMethods()
	tbl := L.NewTable()
	for i, method := range methods {
		item := L.NewTable()
		L.SetField(item, "name", lua.LString(method.Name))
		L.SetField(item, "description", lua.LString(method.Description))
		L.SetField(item, "overridable", lua.LBool(method.Overridable))
		L.SetField(item, "overridden", lua.LBool(method.Overridden))
		L.SetTable(tbl, lua.LNumber(i+1), item)
	}
	L.Push(tbl)
	return 1
}

func (e *LuaEngine) overrideMethodLua(L *lua.LState) int {
	name := L.CheckString(1)
	fn := L.CheckFunction(2)
	result := registry.OverrideMethod(name, fn)
	L.Push(lua.LBool(result))
	return 1
}

func (e *LuaEngine) restoreMethodLua(L *lua.LState) int {
	name := L.CheckString(1)
	result := registry.RestoreMethod(name)
	L.Push(lua.LBool(result))
	return 1
}

func (e *LuaEngine) sleepLua(L *lua.LState) int {
	ms := L.CheckInt(1)
	time.Sleep(time.Duration(ms) * time.Millisecond)
	return 0
}
