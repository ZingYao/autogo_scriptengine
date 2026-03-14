package js_engine

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strconv"
	"sync"
	"time"

	"github.com/ZingYao/autogo_scriptengine/js_engine/model"
	"github.com/ZingYao/autogo_scriptengine/js_engine/model/require"

	"github.com/dop251/goja"
)

var (
	engine         *JSEngine
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

// GetJSEngine 获取默认引擎实例（使用默认配置，自动注入所有方法）
func GetJSEngine() *JSEngine {
	once.Do(func() {
		engine = &JSEngine{
			config: DefaultConfig(),
		}
		initModuleRegistry()
		engine.init()
	})
	return engine
}

// GetEngine 获取默认引擎实例（GetJSEngine 的别名）
func GetEngine() *JSEngine {
	return GetJSEngine()
}

// NewJSEngine 创建新的引擎实例
// config: 引擎配置，传入 nil 使用默认配置
func NewJSEngine(config *EngineConfig) *JSEngine {
	cfg := DefaultConfig()
	if config != nil {
		cfg = *config
	}

	e := &JSEngine{
		config: cfg,
	}
	initModuleRegistry()
	e.init()
	return e
}

// NewEngine 创建新的引擎实例（NewJSEngine 的别名）
func NewEngine(config *EngineConfig) *JSEngine {
	return NewJSEngine(config)
}

func (e *JSEngine) init() {
	e.mu.Lock()
	defer e.mu.Unlock()

	e.vm = goja.New()

	e.registerCoreFunctions()
	if e.config.AutoInjectMethods {
		e.injectAllMethods()
	}
}

func (e *JSEngine) GetVM() *goja.Runtime {
	return e.vm
}

func (e *JSEngine) registerCoreFunctions() {
	vm := e.vm

	vm.Set("registerMethod", e.registerMethodJS)
	vm.Set("unregisterMethod", e.unregisterMethodJS)
	vm.Set("listMethods", e.listMethodsJS)
	vm.Set("overrideMethod", e.overrideMethodJS)
	vm.Set("restoreMethod", e.restoreMethodJS)
	vm.Set("sleep", e.sleepJS)

	// 注册 require 功能
	if e.config.FileSystem != nil {
		requireModule := require.NewRequireModule(vm, e.config.FileSystem)
		requireModule.Register()
	}

	consoleObj := vm.NewObject()
	consoleObj.Set("log", e.consoleLogJS)
	consoleObj.Set("error", e.consoleErrorJS)
	vm.Set("console", consoleObj)
}

func (e *JSEngine) consoleLogJS(call goja.FunctionCall) goja.Value {
	args := call.Arguments
	for _, arg := range args {
		fmt.Print(arg.Export(), " ")
	}
	fmt.Println()
	return goja.Undefined()
}

func (e *JSEngine) consoleErrorJS(call goja.FunctionCall) goja.Value {
	args := call.Arguments
	fmt.Print("[ERROR] ")
	for _, arg := range args {
		fmt.Print(arg.Export(), " ")
	}
	fmt.Println()
	return goja.Undefined()
}

func (e *JSEngine) injectAllMethods() {
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
func (e *JSEngine) InjectModule(moduleName string) {
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
func (e *JSEngine) InjectModules(modules []string) {
	for _, module := range modules {
		e.InjectModule(module)
	}
}

// GetAvailableModules 获取所有可用模块列表
func (e *JSEngine) GetAvailableModules() []string {
	return moduleRegistry.ListModules()
}

// InjectAllMethods 注入所有方法（公开方法，允许手动调用）
func (e *JSEngine) InjectAllMethods() {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.injectAllMethods()
}

func (e *JSEngine) RegisterMethod(name, description string, goFunc interface{}, overridable bool) {
	RegisterMethod(name, description, goFunc, overridable)
}

// ExecuteString 执行 JavaScript 代码字符串
// script: 要执行的 JavaScript 代码
// dir: 可选参数，指定 __dirname（用于 require），如果为空则使用默认值 "scripts"
func (e *JSEngine) ExecuteString(script string, dir ...string) error {
	e.mu.Lock()
	defer e.mu.Unlock()

	if e.vm == nil {
		return fmt.Errorf("JavaScript engine not initialized")
	}

	// 如果配置了文件系统且 __dirname 未设置，设置 __dirname
	// 这样在使用 ExecuteString 时也能正常使用 require
	if e.config.FileSystem != nil {
		currentDir := e.vm.Get("__dirname")
		if currentDir == goja.Undefined() || currentDir.String() == "" {
			// 如果提供了 dir 参数，使用它；否则使用默认值
			__dirname := "scripts"
			if len(dir) > 0 && dir[0] != "" {
				__dirname = dir[0]
			}
			e.vm.Set("__dirname", __dirname)
		}
	}

	_, err := e.vm.RunString(script)
	return err
}

func (e *JSEngine) ExecuteFile(path string) error {
	e.mu.Lock()
	defer e.mu.Unlock()

	if e.vm == nil {
		return fmt.Errorf("JavaScript engine not initialized")
	}

	// 如果配置了文件系统，从文件系统读取并自动设置 __dirname
	if e.config.FileSystem != nil {
		// 读取文件内容
		content, err := fs.ReadFile(e.config.FileSystem, path)
		if err != nil {
			return fmt.Errorf("failed to read file '%s': %v", path, err)
		}

		// 从路径中提取目录和文件名
		dir := filepath.Dir(path)
		filename := filepath.Base(path)

		// 设置 __dirname 和 __filename
		e.vm.Set("__dirname", dir)
		e.vm.Set("__filename", filename)

		// 执行文件内容
		_, err = e.vm.RunString(string(content))
		return err
	}

	// 如果没有配置文件系统，使用原来的方式（通过 load 函数）
	_, err := e.vm.RunString("load('" + path + "')")
	return err
}

func (e *JSEngine) Close() {
	e.mu.Lock()
	defer e.mu.Unlock()

	e.vm = nil
}

func (e *JSEngine) GetRegistry() *MethodRegistry {
	return GetRegistry()
}

// ExecuteString 执行 JavaScript 代码字符串（全局函数）
// script: 要执行的 JavaScript 代码
// dir: 可选参数，指定 __dirname（用于 require），如果为空则使用默认值 "scripts"
func ExecuteString(script string, dir ...string) error {
	if engine != nil {
		return engine.ExecuteString(script, dir...)
	}
	return fmt.Errorf("JavaScript engine not initialized")
}

func ExecuteFile(path string) error {
	if engine != nil {
		return engine.ExecuteFile(path)
	}
	return fmt.Errorf("JavaScript engine not initialized")
}

func Close() {
	if engine != nil {
		engine.Close()
	}
}

func (e *JSEngine) registerMethodJS(call goja.FunctionCall) goja.Value {
	name := call.Argument(0).String()
	description := call.Argument(1).String()
	overridable := call.Argument(2).ToBoolean()

	e.RegisterMethod(name, description, nil, overridable)
	return e.vm.ToValue(true)
}

func (e *JSEngine) unregisterMethodJS(call goja.FunctionCall) goja.Value {
	name := call.Argument(0).String()
	result := registry.RemoveMethod(name)
	return e.vm.ToValue(result)
}

func (e *JSEngine) listMethodsJS(call goja.FunctionCall) goja.Value {
	methods := registry.ListMethods()
	arr := e.vm.NewArray()
	for i, method := range methods {
		item := e.vm.NewObject()
		item.Set("name", method.Name)
		item.Set("description", method.Description)
		item.Set("overridable", method.Overridable)
		item.Set("overridden", method.Overridden)
		arr.Set(strconv.Itoa(i), item)
	}
	return arr
}

func (e *JSEngine) overrideMethodJS(call goja.FunctionCall) goja.Value {
	name := call.Argument(0).String()
	fn, ok := goja.AssertFunction(call.Argument(1))
	if !ok {
		panic("overrideMethod: second argument must be a function")
	}
	result := registry.OverrideMethod(name, fn)
	return e.vm.ToValue(result)
}

func (e *JSEngine) restoreMethodJS(call goja.FunctionCall) goja.Value {
	name := call.Argument(0).String()
	result := registry.RestoreMethod(name)
	return e.vm.ToValue(result)
}

func (e *JSEngine) sleepJS(call goja.FunctionCall) goja.Value {
	ms := int(call.Argument(0).ToInteger())
	time.Sleep(time.Duration(ms) * time.Millisecond)
	return goja.Undefined()
}
