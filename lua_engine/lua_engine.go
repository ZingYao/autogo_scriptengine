package lua_engine

import (
	"fmt"
	"sync"
	"time"

	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"

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

	// 初始化模块缓存
	e.moduleCache = make(map[string]lua.LValue)

	// 设置模块搜索路径
	e.setupPackagePath()

	e.registerCoreFunctions()
	if e.config.AutoInjectMethods {
		e.injectAllMethods()
	}

	// 注册自定义 require 函数，支持从 embed.FS 加载模块
	e.registerCustomRequire()
}

// setupPackagePath 设置 Lua 的模块搜索路径
func (e *LuaEngine) setupPackagePath() {
	// 获取当前的 package.path
	packageTable := e.state.GetGlobal("package").(*lua.LTable)
	currentPathValue := packageTable.RawGetString("path")
	currentPath := ""
	if currentPathValue != lua.LNil {
		currentPath = currentPathValue.String()
	}

	// 添加配置的搜索路径
	for _, searchPath := range e.config.SearchPaths {
		if currentPath != "" {
			currentPath += ";"
		}
		currentPath += searchPath + "/?.lua;" + searchPath + "/?/init.lua"
	}

	// 设置新的 package.path
	packageTable.RawSetString("path", lua.LString(currentPath))
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

// registerCustomRequire 注册自定义 require 函数，支持从 embed.FS 加载模块
func (e *LuaEngine) registerCustomRequire() {
	if e.config.FileSystem == nil {
		return
	}

	// 保存原始 require 函数
	originalRequire := e.state.GetGlobal("require")

	// 注册自定义 require 函数
	e.state.Register("require", func(L *lua.LState) int {
		moduleName := L.CheckString(1)

		// 检查缓存
		if cachedModule, exists := e.moduleCache[moduleName]; exists {
			L.Push(cachedModule)
			return 1
		}

		// 尝试从 embed.FS 加载模块
		if moduleValue, ok := e.loadModuleFromFS(L, moduleName); ok {
			// 缓存模块
			e.moduleCache[moduleName] = moduleValue

			L.Push(moduleValue)
			return 1
		}

		// 如果 embed.FS 中没有找到，使用原始 require
		if originalRequire != lua.LNil {
			L.Push(originalRequire)
			L.Push(lua.LString(moduleName))
			L.Call(1, 1)
			return 1
		}

		// 返回 nil 表示模块未找到
		L.Push(lua.LNil)
		return 1
	})
}

// loadModuleFromFS 从 embed.FS 加载模块
func (e *LuaEngine) loadModuleFromFS(L *lua.LState, moduleName string) (lua.LValue, bool) {
	if e.config.FileSystem == nil {
		return lua.LNil, false
	}

	// 尝试不同的路径模式
	possiblePaths := []string{
		moduleName + ".lua",
		moduleName + "/init.lua",
	}

	// 在所有搜索路径中查找
	for _, searchPath := range e.config.SearchPaths {
		for _, pathPattern := range possiblePaths {
			var fullPath string
			// 如果搜索路径是 "."，直接使用文件名（表示文件系统根目录）
			if searchPath == "." {
				fullPath = pathPattern
			} else {
				fullPath = searchPath + "/" + pathPattern
			}
			
			content, err := e.config.FileSystem.Open(fullPath)
			if err == nil {
				defer content.Close()
				data := make([]byte, 0)
				buf := make([]byte, 1024)
				for {
					n, err := content.Read(buf)
					if err != nil {
						break
					}
					data = append(data, buf[:n]...)
				}

				// 使用 LoadString 加载代码，然后调用以获取返回值
				fn, err := L.LoadString(string(data))
				if err != nil {
					continue
				}

				// 调用函数并获取返回值
				L.Push(fn)
				if err := L.PCall(0, 1, nil); err == nil {
					// 获取模块的返回值（模块应该返回一个 table）
					result := L.Get(-1)
					L.Pop(1)
					return result, true
				}
			}
		}
	}

	return lua.LNil, false
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

// ExecuteString 执行 Lua 代码字符串
// script: 要执行的 Lua 代码
// searchPaths: 可选参数，添加模块搜索路径（用于 require）
func (e *LuaEngine) ExecuteString(script string, searchPaths ...string) error {
	e.mu.Lock()
	defer e.mu.Unlock()

	if e.state == nil {
		return fmt.Errorf("Lua engine not initialized")
	}

	// 如果提供了搜索路径，添加到 package.path
	if len(searchPaths) > 0 {
		e.addSearchPaths(searchPaths...)
	}

	return e.state.DoString(script)
}

// addSearchPaths 添加模块搜索路径
func (e *LuaEngine) addSearchPaths(paths ...string) {
	// 只更新 config.SearchPaths，不更新 package.path
	// 因为我们使用自定义的 require 函数，它会从 config.SearchPaths 和文件系统中加载模块
	for _, searchPath := range paths {
		// 避免重复
		found := false
		for _, existingPath := range e.config.SearchPaths {
			if existingPath == searchPath {
				found = true
				break
			}
		}
		if !found {
			e.config.SearchPaths = append(e.config.SearchPaths, searchPath)
		}
	}
}

func (e *LuaEngine) ExecuteFile(path string) error {
	e.mu.Lock()
	defer e.mu.Unlock()

	if e.state == nil {
		return fmt.Errorf("Lua engine not initialized")
	}

	// 如果配置了文件系统，从文件系统读取并执行
	if e.config.FileSystem != nil {
		// 读取文件内容
		content, err := e.config.FileSystem.Open(path)
		if err != nil {
			return fmt.Errorf("failed to read file '%s': %v", path, err)
		}
		defer content.Close()

		data := make([]byte, 0)
		buf := make([]byte, 1024)
		for {
			n, err := content.Read(buf)
			if err != nil {
				break
			}
			data = append(data, buf[:n]...)
		}

		// 自动检测文件所在目录并添加到搜索路径
		e.addSearchPathsFromPath(path)

		// 执行文件内容
		return e.state.DoString(string(data))
	}

	// 自动检测文件所在目录并添加到搜索路径
	e.addSearchPathsFromPath(path)

	return e.state.DoFile(path)
}

// addSearchPathsFromPath 从文件路径中提取目录并添加到搜索路径
func (e *LuaEngine) addSearchPathsFromPath(path string) {
	// 提取目录（去掉文件名）
	lastSlash := -1
	for i := len(path) - 1; i >= 0; i-- {
		if path[i] == '/' || path[i] == '\\' {
			lastSlash = i
			break
		}
	}

	var dir string
	if lastSlash >= 0 {
		dir = path[:lastSlash]
	} else {
		// 如果路径中没有目录分隔符，说明是相对路径
		// 如果配置了文件系统，使用当前目录（.）
		// 否则使用空字符串
		if e.config.FileSystem != nil {
			dir = "."
		} else {
			dir = ""
		}
	}

	if dir != "" {
		e.addSearchPaths(dir)
	}
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

// ExecuteString 执行 Lua 代码字符串（全局函数）
// script: 要执行的 Lua 代码
// searchPaths: 可选参数，添加模块搜索路径（用于 require）
func ExecuteString(script string, searchPaths ...string) error {
	if engine != nil {
		return engine.ExecuteString(script, searchPaths...)
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
