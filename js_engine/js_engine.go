package js_engine

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"

	"github.com/ZingYao/autogo_scriptengine/js_engine/model"
	"github.com/ZingYao/autogo_scriptengine/js_engine/model/require"

	"github.com/dop251/goja"
)

var (
	engine *JSEngine
	once   sync.Once
)

// GetJSEngine 获取默认引擎实例（使用默认配置，自动注入所有方法）
func GetJSEngine() *JSEngine {
	once.Do(func() {
		engine = &JSEngine{
			config: DefaultConfig(),
		}
		engine.moduleRegistry = model.NewModuleRegistry()
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
	e.moduleRegistry = model.NewModuleRegistry()
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
	vm.Set("load", e.loadJS)

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

func (e *JSEngine) loadJS(call goja.FunctionCall) goja.Value {
	if len(call.Arguments) < 1 {
		return goja.Undefined()
	}

	path := call.Argument(0).String()

	// 读取文件内容
	var content []byte
	var err error

	if e.config.FileSystem != nil {
		// 使用配置的文件系统
		content, err = fs.ReadFile(e.config.FileSystem, path)
	} else {
		// 使用操作系统的文件系统
		content, err = os.ReadFile(path)
	}

	if err != nil {
		panic(fmt.Sprintf("failed to load file '%s': %v", path, err))
	}

	// 从路径中提取目录和文件名
	dir := filepath.Dir(path)
	filename := filepath.Base(path)

	// 保存当前的 __dirname 和 __filename
	oldDirname := e.vm.Get("__dirname")
	oldFilename := e.vm.Get("__filename")

	// 设置新的 __dirname 和 __filename
	e.vm.Set("__dirname", dir)
	e.vm.Set("__filename", filename)

	// 执行文件内容
	defer func() {
		// 恢复原来的 __dirname 和 __filename
		if oldDirname != goja.Undefined() {
			e.vm.Set("__dirname", oldDirname)
		} else {
			e.vm.Set("__dirname", goja.Undefined())
		}
		if oldFilename != goja.Undefined() {
			e.vm.Set("__filename", oldFilename)
		} else {
			e.vm.Set("__filename", goja.Undefined())
		}
	}()

	_, err = e.vm.RunString(string(content))
	if err != nil {
		panic(fmt.Sprintf("failed to execute file '%s': %v", path, err))
	}

	return goja.Undefined()
}

// InjectModule 注入指定模块的方法
// module: 模块名称，支持: app, device, motion, files, images, storages, system, http, media, opencv, ppocr, console, dotocr, hud, ime, plugin, rhino, uiacc, utils, vdisplay, yolo, imgui
func (e *JSEngine) InjectModule(moduleName string) {
	e.mu.Lock()
	defer e.mu.Unlock()

	module, ok := e.moduleRegistry.GetModule(moduleName)
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
	return e.moduleRegistry.ListModules()
}

// RegisterModule 注册一个或多个模块到当前引擎实例
// 用户可以在自己的代码中调用此方法来注册需要的模块
// 支持可变长参数，可以一次注册多个模块
func (e *JSEngine) RegisterModule(modules ...model.Module) {
	for _, module := range modules {
		e.moduleRegistry.RegisterModule(module)
	}
}

func (e *JSEngine) RegisterMethod(name, description string, goFunc interface{}, overridable bool) {
	RegisterMethod(name, description, goFunc, overridable)
}

// ExecuteString 执行 JavaScript 代码字符串
// script: 要执行的 JavaScript 代码
// dir: 可选参数，指定 __dirname（用于 require），如果为空则使用默认值 "scripts"
// 支持脚本退出后的动作：
//   - process.exit(0): 正常退出，执行配置的退出动作（重启/自定义/无动作）
//   - process.exit(-1): 强制退出，不执行任何退出动作
//   - process.exit(其他值): 正常退出，执行配置的退出动作
//
// 脚本异常退出时始终打印日志
func (e *JSEngine) ExecuteString(script string, dir ...string) error {
	e.currentScript = script
	if len(dir) > 0 && dir[0] != "" {
		e.currentDir = dir[0]
	} else {
		e.currentDir = "scripts"
	}
	e.skipExitAction = false

	for {
		err := e.executeStringOnce(script, dir...)

		// 如果脚本异常退出，打印错误日志
		if err != nil {
			fmt.Printf("脚本异常退出: %v\n", err)
			return err
		}

		// 如果跳过退出动作（process.exit(-1)），直接返回
		if e.skipExitAction {
			return nil
		}

		// 根据配置的退出动作执行相应操作
		switch e.config.OnExit {
		case ExitActionNone:
			// 无动作，直接退出
			return nil
		case ExitActionRestart:
			// 重启脚本
			fmt.Println("脚本正常退出，正在重新启动...")
			time.Sleep(1 * time.Second)
			// 继续循环，重新执行脚本
		case ExitActionCustom:
			// 执行自定义退出动作
			if e.config.CustomExitAction != nil {
				e.config.CustomExitAction()
			}
			return nil
		}
	}
}

// executeStringOnce 执行一次 JavaScript 代码字符串
func (e *JSEngine) executeStringOnce(script string, dir ...string) error {
	e.mu.Lock()
	defer e.mu.Unlock()

	if e.vm == nil {
		return fmt.Errorf("JavaScript engine not initialized")
	}

	// 如果配置了文件系统且 __dirname 未设置，设置 __dirname
	if e.config.FileSystem != nil {
		currentDir := e.vm.Get("__dirname")
		if currentDir == goja.Undefined() || currentDir.String() == "" {
			__dirname := "scripts"
			if len(dir) > 0 && dir[0] != "" {
				__dirname = dir[0]
			}
			e.vm.Set("__dirname", __dirname)
		}
	}

	// 注册特殊的 process.exit 函数，用于控制退出动作
	e.registerExitControl()

	// 使用 IIFE 包装脚本，避免全局作用域污染
	wrappedScript := "(function() {\n" + script + "\n})();"

	_, err := e.vm.RunString(wrappedScript)
	return err
}

func (e *JSEngine) ExecuteFile(path string) error {
	// 使用 load 函数来加载文件
	_, err := e.vm.RunString("load('" + path + "')")
	return err
}

// registerExitControl 注册特殊的 process.exit 函数，用于控制退出动作
// process.exit(0) - 正常退出，执行配置的退出动作（重启/自定义/无动作）
// process.exit(-1) - 强制退出，不执行任何退出动作
// process.exit(其他值) - 正常退出，执行配置的退出动作
func (e *JSEngine) registerExitControl() {
	// 获取 process 对象
	process := e.vm.Get("process")
	if process == goja.Undefined() {
		// 如果 process 对象不存在，创建一个
		processObj := e.vm.NewObject()
		e.vm.Set("process", processObj)
		process = processObj
	}

	processObj, ok := process.(*goja.Object)
	if !ok {
		// 如果 process 不是对象，无法注册退出控制
		return
	}

	// 保存原始的 exit 函数
	originalExit := processObj.Get("exit")

	// 注册新的 exit 函数
	processObj.Set("exit", e.vm.ToValue(func(code int) {
		// 如果退出码为 -1，跳过退出动作
		if code == -1 {
			e.skipExitAction = true
		}

		// 调用原始的 exit 函数
		if originalExit != goja.Undefined() {
			if exitFunc, ok := goja.AssertFunction(originalExit); ok {
				_, _ = exitFunc(goja.Null(), e.vm.ToValue(code))
			}
		}
	}))
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
