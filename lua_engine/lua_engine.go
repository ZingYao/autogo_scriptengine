package lua_engine

import (
	"context"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/ZingYao/autogo_scriptengine/lua_engine/debugger"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"
	"github.com/ZingYao/go-lua-vm/bridge"
	glua "github.com/ZingYao/go-lua-vm/lua"
	gruntime "github.com/ZingYao/go-lua-vm/runtime"
	gluaos "github.com/ZingYao/go-lua-vm/stdlib/os"
)

var (
	engine          *LuaEngine
	once            sync.Once
	errorReturnType = reflect.TypeOf((*error)(nil)).Elem()
)

// GetLuaEngine 获取默认引擎实例（使用默认配置，自动注入所有方法）
func GetLuaEngine() *LuaEngine {
	once.Do(func() {
		engine = &LuaEngine{
			config: DefaultConfig(),
		}
		engine.moduleRegistry = model.NewModuleRegistry()
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
	e.moduleRegistry = model.NewModuleRegistry()
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

	options := glua.DefaultOptions()
	options.AllowHostFilesystem = true
	options.AllowEnvironment = true
	options.VirtualFilesystem = e.config.FileSystem
	options.PreferHostFilesystem = e.config.FileSystem == nil
	e.vmState = glua.NewStateWithOptions(options)
	if err := glua.OpenLibs(e.vmState); err != nil {
		fmt.Printf("[ERROR] open go-lua-vm libs failed: %v\n", err)
	}

	// 设置模块搜索路径
	e.setupPackagePath()

	// 初始化状态管理字段
	e.engineState = StateStopped
	e.ctx, e.cancel = context.WithCancel(context.Background())
	e.pauseChan = make(chan struct{})
	if e.config.Debug != nil && e.config.Debug.Enabled {
		e.debugger = debugger.New(*e.config.Debug)
		e.debugger.InstallVM(e.vmState)
	}

	e.registerCoreFunctionsForVM()
	e.reinstallVMMethodsFromRegistry()
}

// Start 启动引擎
func (e *LuaEngine) Start() {
	e.mu.Lock()
	defer e.mu.Unlock()

	if e.engineState == StateStopped {
		e.ctx, e.cancel = context.WithCancel(context.Background())
		e.pauseChan = make(chan struct{})
		e.engineState = StateRunning
	}
}

// Pause 暂停引擎执行
func (e *LuaEngine) Pause() {
	e.mu.Lock()
	defer e.mu.Unlock()

	if e.engineState == StateRunning {
		e.engineState = StatePaused
	}
}

// Resume 恢复引擎执行
func (e *LuaEngine) Resume() {
	e.mu.Lock()
	defer e.mu.Unlock()

	if e.engineState == StatePaused {
		e.engineState = StateRunning
		close(e.pauseChan)
		e.pauseChan = make(chan struct{})
	}
}

// Stop 停止引擎执行
func (e *LuaEngine) Stop() {
	e.mu.Lock()
	defer e.mu.Unlock()

	if e.engineState != StateStopped {
		if e.cancel != nil {
			e.cancel()
		}
		e.engineState = StateStopped
		close(e.pauseChan)
	}
}

// GetEngineState 获取引擎状态
func (e *LuaEngine) GetEngineState() EngineState {
	e.mu.RLock()
	defer e.mu.RUnlock()

	return e.engineState
}

// AddRequirePath 添加自定义 require 路径
func (e *LuaEngine) AddRequirePath(path string) {
	e.mu.Lock()
	defer e.mu.Unlock()

	for _, existingPath := range e.config.RequirePaths {
		if existingPath == path {
			return
		}
	}
	e.config.RequirePaths = append(e.config.RequirePaths, path)
	e.setupPackagePath()
}

// SetRequirePaths 设置自定义 require 路径
func (e *LuaEngine) SetRequirePaths(paths []string) {
	e.mu.Lock()
	defer e.mu.Unlock()

	e.config.RequirePaths = paths
	e.setupPackagePath()
}

// setupPackagePath 设置 Lua 的模块搜索路径
func (e *LuaEngine) setupPackagePath() {
	e.setupVMPackagePath()
}

// GetDebugger 获取当前 Lua 调试器实例，未启用调试时返回 nil。
func (e *LuaEngine) GetDebugger() *debugger.Debugger {
	return e.debugger
}

func (e *LuaEngine) setupVMPackagePath() {
	if e.vmState == nil {
		return
	}
	packageValue, err := glua.GetGlobal(e.vmState, "package")
	if err != nil || packageValue.Kind != gruntime.KindTable {
		return
	}
	packageTable, ok := packageValue.Ref.(*gruntime.Table)
	if !ok || packageTable == nil {
		return
	}
	currentPathValue := packageTable.RawGetString("path")
	if e.vmPackagePath == "" && currentPathValue.Kind == gruntime.KindString {
		e.vmPackagePath = currentPathValue.String
	}
	currentPath := e.vmPackagePath
	if currentPathValue.Kind == gruntime.KindString {
		currentPath = e.vmPackagePath
	}
	for _, searchPath := range e.config.SearchPaths {
		currentPath = appendLuaPackagePath(currentPath, searchPath)
	}
	for _, requirePath := range e.config.RequirePaths {
		currentPath = appendLuaPackagePath(currentPath, requirePath)
	}
	packageTable.RawSetString("path", gruntime.StringValue(currentPath))
}

func appendLuaPackagePath(currentPath string, root string) string {
	if root == "" {
		return currentPath
	}
	if currentPath != "" {
		currentPath += ";"
	}
	return currentPath + root + "/?.lua;" + root + "/?/init.lua"
}

// GetVMState 获取 go-lua-vm 状态机，供新 Lua 引擎桥接模块逐步迁移使用。
func (e *LuaEngine) GetVMState() *glua.State {
	return e.vmState
}

// InjectModule 注入指定模块的方法
// module: 模块名称，支持: app, device, motion, files, images, storages, system, http, media, opencv, ppocr, console, dotocr, hud, ime, plugin, rhino, uiacc, utils, vdisplay, yolo, imgui
func (e *LuaEngine) InjectModule(moduleName string) {
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
func (e *LuaEngine) InjectModules(modules []string) {
	for _, module := range modules {
		e.InjectModule(module)
	}
}

// GetAvailableModules 获取所有可用模块列表
func (e *LuaEngine) GetAvailableModules() []string {
	return e.moduleRegistry.ListModules()
}

// RegisterModule 注册一个或多个模块到当前引擎实例
// 用户可以在自己的代码中调用此方法来注册需要的模块
// 支持可变长参数，可以一次注册多个模块
func (e *LuaEngine) RegisterModule(modules ...model.Module) {
	for _, module := range modules {
		e.moduleRegistry.RegisterModule(module)
		module.Register(e)
	}
}

func (e *LuaEngine) RegisterMethod(name, description string, goFunc interface{}, overridable bool) {
	RegisterMethod(name, description, goFunc, overridable)
	if err := e.installVMMethod(name, goFunc); err != nil && e.config.FailFast {
		fmt.Printf("[ERROR] install go-lua-vm method %s failed: %v\n", name, err)
	}
}

// RegisterValue 注册 Lua 模块字段值，支持 device.width 这类非函数字段。
func (e *LuaEngine) RegisterValue(name string, value interface{}) error {
	if e.vmState == nil {
		return fmt.Errorf("Lua engine not initialized")
	}
	parts := strings.Split(name, ".")
	if len(parts) < 2 || parts[0] == "" || parts[len(parts)-1] == "" {
		return fmt.Errorf("invalid lua value name: %s", name)
	}
	moduleTable, err := e.ensureVMModuleTable(parts[:len(parts)-1])
	if err != nil {
		return err
	}
	luaValue, err := e.reflectToVMValue(reflect.ValueOf(value))
	if err != nil {
		return err
	}
	return moduleTable.RawSet(gruntime.StringValue(parts[len(parts)-1]), luaValue)
}

// ExecuteString 执行 Lua 代码字符串（实例方法）
// script: 要执行的 Lua 代码
// searchPaths: 可选参数，添加模块搜索路径（用于 require）
// 支持脚本退出后的动作：
//   - os.exit(0): 正常退出，执行配置的退出动作（重启/自定义/无动作）
//   - os.exit(-1): 强制退出，不执行任何退出动作
//   - os.exit(其他值): 正常退出，执行配置的退出动作
//
// 脚本异常退出时始终打印日志
func (e *LuaEngine) ExecuteString(script string, searchPaths ...string) error {
	return e.ExecuteStringWithMode(script, e.config.ExecuteMode, searchPaths...)
}

// ExecuteStringWithMode 带执行模式的执行 Lua 代码字符串
// script: 要执行的 Lua 代码
// mode: 执行模式
// searchPaths: 可选参数，添加模块搜索路径（用于 require）
func (e *LuaEngine) ExecuteStringWithMode(script string, mode ExecuteMode, searchPaths ...string) error {
	e.Start()
	e.currentScript = script
	e.currentSearchPaths = searchPaths
	e.skipExitAction = false

	// 异步执行
	if mode == ExecuteModeAsync {
		go func() {
			e.executeStringLoop(script, searchPaths...)
		}()
		return nil
	}

	// 同步执行
	return e.executeStringLoop(script, searchPaths...)
}

// executeStringLoop 执行脚本循环
func (e *LuaEngine) executeStringLoop(script string, searchPaths ...string) error {
	for {
		err := e.executeStringOnce(script, searchPaths...)

		// 如果脚本异常退出，打印错误日志
		if err != nil {
			fmt.Printf("脚本异常退出: %v\n", err)
			e.Stop()
			return err
		}

		// 如果跳过退出动作（os.exit(-1)），直接返回
		if e.skipExitAction {
			e.Stop()
			return nil
		}

		// 根据配置的退出动作执行相应操作
		switch e.config.OnExit {
		case ExitActionNone:
			// 无动作，直接退出
			e.Stop()
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
			e.Stop()
			return nil
		}
	}
}

// executeStringOnce 执行一次 Lua 代码字符串
func (e *LuaEngine) executeStringOnce(script string, searchPaths ...string) error {
	e.mu.Lock()
	defer e.mu.Unlock()

	if e.vmState == nil {
		return fmt.Errorf("Lua engine not initialized")
	}

	// 如果提供了搜索路径，添加到 package.path
	if len(searchPaths) > 0 {
		e.addSearchPaths(searchPaths...)
	}

	if e.debugger != nil && e.debugger.Enabled() {
		e.debugger.InstallVM(e.vmState)
		script = debugger.InstrumentSource(script, "<string>")
	}

	err := e.handleLuaVMError(glua.DoString(e.vmState, script))
	if err != nil && e.debugger != nil {
		e.debugger.NotifyError("<string>", err)
	}
	return err
}

// addSearchPaths 添加模块搜索路径
func (e *LuaEngine) addSearchPaths(paths ...string) {
	changed := false
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
			changed = true
		}
	}
	if changed {
		e.setupPackagePath()
	}
}

// ExecuteFile 执行 Lua 文件
// path: 要执行的 Lua 文件路径
// 支持脚本退出后的动作：
//   - os.exit(0): 正常退出，执行配置的退出动作（重启/自定义/无动作）
//   - os.exit(-1): 强制退出，不执行任何退出动作
//   - os.exit(其他值): 正常退出，执行配置的退出动作
//
// 脚本异常退出时始终打印日志
func (e *LuaEngine) ExecuteFile(path string) error {
	return e.ExecuteFileWithMode(path, e.config.ExecuteMode)
}

// ExecuteFileWithMode 带执行模式的执行 Lua 文件
// path: 要执行的 Lua 文件路径
// mode: 执行模式
func (e *LuaEngine) ExecuteFileWithMode(path string, mode ExecuteMode) error {
	// 读取文件内容
	var content string

	if e.config.FileSystem != nil {
		// 从文件系统读取
		file, err := e.config.FileSystem.Open(path)
		if err != nil {
			return fmt.Errorf("failed to read file '%s': %v", path, err)
		}
		defer file.Close()

		data := make([]byte, 0)
		buf := make([]byte, 1024)
		for {
			n, err := file.Read(buf)
			if err != nil {
				break
			}
			data = append(data, buf[:n]...)
		}
		content = string(data)
	} else {
		// 从本地文件系统读取
		data, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("failed to read file '%s': %v", path, err)
		}
		content = string(data)
	}

	// 提取文件所在目录作为搜索路径
	dir := ""
	lastSlash := -1
	for i := len(path) - 1; i >= 0; i-- {
		if path[i] == '/' || path[i] == '\\' {
			lastSlash = i
			break
		}
	}

	if lastSlash >= 0 {
		dir = path[:lastSlash]
	} else {
		// 如果路径中没有目录分隔符，说明是相对路径
		if e.config.FileSystem != nil {
			dir = "."
		} else {
			dir = ""
		}
	}

	// 构建搜索路径，包括文件所在目录和自定义 require 路径
	searchPaths := []string{}
	if dir != "" {
		searchPaths = append(searchPaths, dir)
	}
	searchPaths = append(searchPaths, e.config.RequirePaths...)

	// 异步执行
	if mode == ExecuteModeAsync {
		go func() {
			e.ExecuteStringWithMode(content, ExecuteModeSync, searchPaths...)
		}()
		return nil
	}

	// 同步执行
	return e.ExecuteStringWithMode(content, ExecuteModeSync, searchPaths...)
}

// executeFileOnce 执行一次 Lua 文件
func (e *LuaEngine) executeFileOnce(path string) error {
	e.mu.Lock()
	defer e.mu.Unlock()

	if e.vmState == nil {
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
		return e.handleLuaVMError(glua.DoString(e.vmState, string(data)))
	}

	// 自动检测文件所在目录并添加到搜索路径
	e.addSearchPathsFromPath(path)

	return e.handleLuaVMError(glua.DoFile(e.vmState, path))
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

	if e.engineState != StateStopped {
		if e.cancel != nil {
			e.cancel()
		}
		e.engineState = StateStopped
		close(e.pauseChan)
		e.pauseChan = make(chan struct{})
	}
	if e.vmState != nil {
		_ = glua.Close(e.vmState)
		e.vmState = nil
	}
}

func (e *LuaEngine) handleLuaVMError(err error) error {
	if err == nil {
		return nil
	}
	var exitErr *gluaos.ExitError
	if errors.As(err, &exitErr) {
		if exitErr.Code == -1 {
			e.skipExitAction = true
		}
		return nil
	}
	return err
}

// Restart 重启 Lua 引擎
// 关闭当前状态并重新初始化，保留模块缓存
func (e *LuaEngine) Restart() error {
	e.mu.Lock()
	defer e.mu.Unlock()

	// 关闭当前状态
	if e.vmState != nil {
		_ = glua.Close(e.vmState)
	}
	e.vmPackagePath = ""

	options := glua.DefaultOptions()
	options.AllowHostFilesystem = true
	options.AllowEnvironment = true
	options.VirtualFilesystem = e.config.FileSystem
	options.PreferHostFilesystem = e.config.FileSystem == nil
	e.vmState = glua.NewStateWithOptions(options)
	if err := glua.OpenLibs(e.vmState); err != nil {
		return fmt.Errorf("open go-lua-vm libs: %w", err)
	}

	// 重新设置模块搜索路径
	e.setupPackagePath()

	// 重新注册核心函数
	if e.config.Debug != nil && e.config.Debug.Enabled {
		e.debugger = debugger.New(*e.config.Debug)
		e.debugger.InstallVM(e.vmState)
	}
	e.registerCoreFunctionsForVM()
	e.reinstallVMMethodsFromRegistry()

	return nil
}

func (e *LuaEngine) reinstallVMMethodsFromRegistry() {
	initRegistry()
	for _, method := range registry.ListMethods() {
		if method.GoFunc == nil {
			continue
		}
		if err := e.installVMMethod(method.Name, method.GoFunc); err != nil && e.config.FailFast {
			fmt.Printf("[ERROR] reinstall go-lua-vm method %s failed: %v\n", method.Name, err)
		}
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
		engine = nil
		once = sync.Once{}
	}
}

func (e *LuaEngine) registerCoreFunctionsForVM() {
	if e.vmState == nil {
		return
	}
	_ = glua.Register(e.vmState, "registerMethod", func(args ...glua.Value) ([]glua.Value, error) {
		name := gluaStringArg(args, 0)
		description := gluaStringArg(args, 1)
		overridable := gluaBoolArg(args, 2)
		e.RegisterMethod(name, description, nil, overridable)
		return []glua.Value{gruntime.BooleanValue(true)}, nil
	})
	_ = glua.Register(e.vmState, "unregisterMethod", func(args ...glua.Value) ([]glua.Value, error) {
		name := gluaStringArg(args, 0)
		result := registry.RemoveMethod(name)
		_ = e.uninstallVMMethod(name)
		return []glua.Value{gruntime.BooleanValue(result)}, nil
	})
	_ = glua.Register(e.vmState, "listMethods", func(args ...glua.Value) ([]glua.Value, error) {
		table := gruntime.NewTable()
		for index, method := range registry.ListMethods() {
			item := gruntime.NewTable()
			_ = item.RawSet(gruntime.StringValue("name"), gruntime.StringValue(method.Name))
			_ = item.RawSet(gruntime.StringValue("description"), gruntime.StringValue(method.Description))
			_ = item.RawSet(gruntime.StringValue("overridable"), gruntime.BooleanValue(method.Overridable))
			_ = item.RawSet(gruntime.StringValue("overridden"), gruntime.BooleanValue(method.Overridden))
			_ = table.RawSet(gruntime.IntegerValue(int64(index+1)), gruntime.ReferenceValue(gruntime.KindTable, item))
		}
		return []glua.Value{gruntime.ReferenceValue(gruntime.KindTable, table)}, nil
	})
	_ = glua.Register(e.vmState, "sleep", func(args ...glua.Value) ([]glua.Value, error) {
		ms := int64(0)
		if len(args) > 0 {
			ms, _ = args[0].ToInteger()
		}
		time.Sleep(time.Duration(ms) * time.Millisecond)
		return nil, nil
	})
	_ = glua.Register(e.vmState, "overrideMethod", func(args ...glua.Value) ([]glua.Value, error) {
		return []glua.Value{gruntime.BooleanValue(false)}, nil
	})
	_ = glua.Register(e.vmState, "restoreMethod", func(args ...glua.Value) ([]glua.Value, error) {
		name := gluaStringArg(args, 0)
		return []glua.Value{gruntime.BooleanValue(registry.RestoreMethod(name))}, nil
	})
	_ = e.installVMCoreTable("console", map[string]gruntime.GoResultsFunction{
		"log": func(args ...gruntime.Value) ([]gruntime.Value, error) {
			e.printVMConsole("", args)
			return nil, nil
		},
		"error": func(args ...gruntime.Value) ([]gruntime.Value, error) {
			e.printVMConsole("[ERROR] ", args)
			return nil, nil
		},
	})
}

func (e *LuaEngine) installVMMethod(name string, goFunc interface{}) error {
	if e.vmState == nil || goFunc == nil {
		return nil
	}
	parts := strings.Split(name, ".")
	if len(parts) < 2 || parts[0] == "" || parts[len(parts)-1] == "" {
		return nil
	}
	functionValue, err := e.bindVMMethodValue(goFunc)
	if err != nil {
		return err
	}
	moduleTable, err := e.ensureVMModuleTable(parts[:len(parts)-1])
	if err != nil {
		return err
	}
	return moduleTable.RawSet(gruntime.StringValue(parts[len(parts)-1]), functionValue)
}

func (e *LuaEngine) ensureVMModuleTable(parts []string) (*gruntime.Table, error) {
	var currentTable *gruntime.Table
	for index, part := range parts {
		if part == "" {
			return nil, nil
		}
		var currentValue gruntime.Value
		var err error
		if index == 0 {
			currentValue, err = glua.GetGlobal(e.vmState, part)
			if err != nil {
				return nil, err
			}
		} else {
			currentValue = currentTable.RawGetString(part)
		}
		nextTable, _ := currentValue.Ref.(*gruntime.Table)
		if currentValue.Kind != gruntime.KindTable || nextTable == nil {
			nextTable = gruntime.NewTable()
			nextValue := gruntime.ReferenceValue(gruntime.KindTable, nextTable)
			if index == 0 {
				if err := glua.SetGlobal(e.vmState, part, nextValue); err != nil {
					return nil, err
				}
			} else {
				currentTable.RawSetString(part, nextValue)
			}
		}
		currentTable = nextTable
		if err := e.setVMLoadedModule(strings.Join(parts[:index+1], "."), currentTable); err != nil {
			return nil, err
		}
	}
	return currentTable, nil
}

func (e *LuaEngine) setVMLoadedModule(name string, moduleTable *gruntime.Table) error {
	if name == "" || moduleTable == nil {
		return nil
	}
	packageValue, err := glua.GetGlobal(e.vmState, "package")
	if err != nil {
		return err
	}
	packageTable, _ := packageValue.Ref.(*gruntime.Table)
	if packageValue.Kind != gruntime.KindTable || packageTable == nil {
		packageTable = gruntime.NewTable()
		packageValue = gruntime.ReferenceValue(gruntime.KindTable, packageTable)
		if err := glua.SetGlobal(e.vmState, "package", packageValue); err != nil {
			return err
		}
	}
	loadedValue := packageTable.RawGetString("loaded")
	loadedTable, _ := loadedValue.Ref.(*gruntime.Table)
	if loadedValue.Kind != gruntime.KindTable || loadedTable == nil {
		loadedTable = gruntime.NewTable()
		packageTable.RawSetString("loaded", gruntime.ReferenceValue(gruntime.KindTable, loadedTable))
	}
	loadedTable.RawSetString(name, gruntime.ReferenceValue(gruntime.KindTable, moduleTable))
	return nil
}

func (e *LuaEngine) lookupVMModuleTable(parts []string) (*gruntime.Table, error) {
	var currentTable *gruntime.Table
	for index, part := range parts {
		if part == "" {
			return nil, nil
		}
		var currentValue gruntime.Value
		var err error
		if index == 0 {
			currentValue, err = glua.GetGlobal(e.vmState, part)
			if err != nil {
				return nil, err
			}
		} else {
			currentValue = currentTable.RawGetString(part)
		}
		nextTable, _ := currentValue.Ref.(*gruntime.Table)
		if currentValue.Kind != gruntime.KindTable || nextTable == nil {
			return nil, nil
		}
		currentTable = nextTable
	}
	return currentTable, nil
}

func (e *LuaEngine) uninstallVMMethod(name string) error {
	if e.vmState == nil {
		return nil
	}
	parts := strings.Split(name, ".")
	if len(parts) < 2 || parts[0] == "" || parts[len(parts)-1] == "" {
		return nil
	}
	moduleTable, err := e.lookupVMModuleTable(parts[:len(parts)-1])
	if err != nil || moduleTable == nil {
		return err
	}
	return moduleTable.RawSet(gruntime.StringValue(parts[len(parts)-1]), gruntime.NilValue())
}

func (e *LuaEngine) bindVMMethodValue(goFunc interface{}) (glua.Value, error) {
	functionType := reflect.TypeOf(goFunc)
	if functionType == nil || functionType.Kind() != reflect.Func {
		return gruntime.NilValue(), fmt.Errorf("goFunc must be function")
	}
	if !needsCustomVMBridge(functionType) {
		return bridge.BindReflectFunction(e.vmState, goFunc)
	}
	functionValue := reflect.ValueOf(goFunc)
	return gruntime.ReferenceValue(gruntime.KindGoClosure, gruntime.GoResultsFunction(func(args ...gruntime.Value) (results []gruntime.Value, err error) {
		defer func() {
			if recovered := recover(); recovered != nil {
				results = nil
				err = fmt.Errorf("panic in bridge function: %v", recovered)
			}
		}()
		callArgs, err := e.vmCallArgsToReflect(args, functionType)
		if err != nil {
			return nil, err
		}
		var goResults []reflect.Value
		if functionType.IsVariadic() {
			goResults = functionValue.CallSlice(callArgs)
		} else {
			goResults = functionValue.Call(callArgs)
		}
		luaResults := make([]gruntime.Value, 0, len(goResults))
		for index, result := range goResults {
			if result.Type().Implements(errorReturnType) {
				if !result.IsNil() {
					return nil, result.Interface().(error)
				}
				continue
			}
			convertedResult, err := e.reflectToVMValue(result)
			if err != nil {
				return nil, fmt.Errorf("return %d: %w", index+1, err)
			}
			luaResults = append(luaResults, convertedResult)
		}
		return luaResults, nil
	})), nil
}

func (e *LuaEngine) vmCallArgsToReflect(args []gruntime.Value, functionType reflect.Type) ([]reflect.Value, error) {
	if !functionType.IsVariadic() {
		if len(args) != functionType.NumIn() {
			return nil, fmt.Errorf("argument count mismatch: got %d, want %d", len(args), functionType.NumIn())
		}
		callArgs := make([]reflect.Value, 0, functionType.NumIn())
		for index := 0; index < functionType.NumIn(); index++ {
			convertedArg, err := e.vmValueToReflect(args[index], functionType.In(index))
			if err != nil {
				return nil, fmt.Errorf("argument %d: %w", index+1, err)
			}
			callArgs = append(callArgs, convertedArg)
		}
		return callArgs, nil
	}

	fixedArgCount := functionType.NumIn() - 1
	if len(args) < fixedArgCount {
		return nil, fmt.Errorf("argument count mismatch: got %d, want at least %d", len(args), fixedArgCount)
	}
	callArgs := make([]reflect.Value, 0, functionType.NumIn())
	for index := 0; index < fixedArgCount; index++ {
		convertedArg, err := e.vmValueToReflect(args[index], functionType.In(index))
		if err != nil {
			return nil, fmt.Errorf("argument %d: %w", index+1, err)
		}
		callArgs = append(callArgs, convertedArg)
	}
	variadicType := functionType.In(functionType.NumIn() - 1)
	variadicValues := reflect.MakeSlice(variadicType, 0, len(args)-fixedArgCount)
	for index := fixedArgCount; index < len(args); index++ {
		convertedArg, err := e.vmValueToReflect(args[index], variadicType.Elem())
		if err != nil {
			return nil, fmt.Errorf("argument %d: %w", index+1, err)
		}
		variadicValues = reflect.Append(variadicValues, convertedArg)
	}
	callArgs = append(callArgs, variadicValues)
	return callArgs, nil
}

func needsCustomVMBridge(functionType reflect.Type) bool {
	for index := 0; index < functionType.NumIn(); index++ {
		if needsCustomVMBridgeType(functionType.In(index)) {
			return true
		}
	}
	for index := 0; index < functionType.NumOut(); index++ {
		outputType := functionType.Out(index)
		if outputType.Implements(errorReturnType) {
			continue
		}
		if needsCustomVMBridgeType(outputType) {
			return true
		}
	}
	return false
}

func needsCustomVMBridgeType(valueType reflect.Type) bool {
	if valueType.Kind() == reflect.Interface && valueType.NumMethod() == 0 {
		return true
	}
	switch valueType.Kind() {
	case reflect.Map:
		return true
	case reflect.Func:
		return true
	case reflect.Struct:
		return true
	case reflect.Pointer:
		return valueType.Elem().Kind() == reflect.Struct
	case reflect.Slice:
		return valueType.Elem().Kind() != reflect.Uint8
	default:
		return false
	}
}

func (e *LuaEngine) vmValueToReflect(value gruntime.Value, targetType reflect.Type) (reflect.Value, error) {
	if value.IsNil() {
		return reflect.Zero(targetType), nil
	}
	switch targetType.Kind() {
	case reflect.Bool:
		return reflect.ValueOf(value.Truthy()).Convert(targetType), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		intValue, ok := value.ToInteger()
		if !ok {
			return reflect.Value{}, fmt.Errorf("expected integer, got %s", value.DebugString())
		}
		return reflect.ValueOf(intValue).Convert(targetType), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		intValue, ok := value.ToInteger()
		if !ok || intValue < 0 {
			return reflect.Value{}, fmt.Errorf("expected unsigned integer, got %s", value.DebugString())
		}
		return reflect.ValueOf(uint64(intValue)).Convert(targetType), nil
	case reflect.Float32, reflect.Float64:
		if value.Kind == gruntime.KindNumber {
			return reflect.ValueOf(value.Number).Convert(targetType), nil
		}
		intValue, ok := value.ToInteger()
		if !ok {
			return reflect.Value{}, fmt.Errorf("expected number, got %s", value.DebugString())
		}
		return reflect.ValueOf(float64(intValue)).Convert(targetType), nil
	case reflect.String:
		text, err := gruntime.ToString(value)
		if err != nil {
			return reflect.Value{}, err
		}
		return reflect.ValueOf(text.String).Convert(targetType), nil
	case reflect.Slice:
		return e.vmTableToReflectSlice(value, targetType)
	case reflect.Map:
		return e.vmTableToReflectMap(value, targetType)
	case reflect.Struct:
		return e.vmTableToReflectStruct(value, targetType)
	case reflect.Pointer:
		return e.vmObjectToReflect(value, targetType)
	case reflect.Func:
		return e.vmFunctionToReflect(value, targetType)
	case reflect.Interface:
		return e.vmValueToInterface(value, targetType)
	default:
		return reflect.Value{}, fmt.Errorf("unsupported argument type %s", targetType.String())
	}
}

func (e *LuaEngine) vmTableToReflectSlice(value gruntime.Value, targetType reflect.Type) (reflect.Value, error) {
	if targetType.Elem().Kind() == reflect.Uint8 {
		text, err := gruntime.ToString(value)
		if err != nil {
			return reflect.Value{}, err
		}
		return reflect.ValueOf([]byte(text.String)).Convert(targetType), nil
	}
	if value.Kind != gruntime.KindTable {
		return reflect.Value{}, fmt.Errorf("expected table, got %s", value.DebugString())
	}
	table, ok := value.Ref.(*gruntime.Table)
	if !ok || table == nil {
		return reflect.Zero(targetType), nil
	}
	result := reflect.MakeSlice(targetType, 0, 0)
	for index := int64(1); ; index++ {
		item := table.RawGetInteger(index)
		if item.IsNil() {
			break
		}
		convertedItem, err := e.vmValueToReflect(item, targetType.Elem())
		if err != nil {
			return reflect.Value{}, err
		}
		result = reflect.Append(result, convertedItem)
	}
	return result, nil
}

func (e *LuaEngine) vmTableToReflectStruct(value gruntime.Value, targetType reflect.Type) (reflect.Value, error) {
	if value.Kind != gruntime.KindTable {
		return reflect.Value{}, fmt.Errorf("expected table, got %s", value.DebugString())
	}
	table, ok := value.Ref.(*gruntime.Table)
	if !ok || table == nil {
		return reflect.Zero(targetType), nil
	}
	result := reflect.New(targetType).Elem()
	for index := 0; index < targetType.NumField(); index++ {
		field := targetType.Field(index)
		if field.PkgPath != "" || !result.Field(index).CanSet() {
			continue
		}
		fieldValue := vmTableField(table, field)
		if fieldValue.IsNil() {
			continue
		}
		convertedValue, err := e.vmValueToReflect(fieldValue, field.Type)
		if err != nil {
			return reflect.Value{}, fmt.Errorf("field %s: %w", field.Name, err)
		}
		result.Field(index).Set(convertedValue)
	}
	return result, nil
}

func (e *LuaEngine) vmObjectToReflect(value gruntime.Value, targetType reflect.Type) (reflect.Value, error) {
	if object, ok := vmObjectFromValue(value); ok {
		objectValue := reflect.ValueOf(object)
		if objectValue.Type().AssignableTo(targetType) {
			return objectValue, nil
		}
		if objectValue.Type().ConvertibleTo(targetType) {
			return objectValue.Convert(targetType), nil
		}
		return reflect.Value{}, fmt.Errorf("object %s cannot convert to %s", objectValue.Type().String(), targetType.String())
	}
	if value.Kind == gruntime.KindTable && targetType.Elem().Kind() == reflect.Struct {
		structValue, err := e.vmTableToReflectStruct(value, targetType.Elem())
		if err != nil {
			return reflect.Value{}, err
		}
		pointerValue := reflect.New(targetType.Elem())
		pointerValue.Elem().Set(structValue)
		return pointerValue, nil
	}
	return reflect.Value{}, fmt.Errorf("reflect object proxy expected")
}

func (e *LuaEngine) vmFunctionToReflect(value gruntime.Value, targetType reflect.Type) (reflect.Value, error) {
	if value.Kind != gruntime.KindGoClosure && value.Kind != gruntime.KindLuaClosure {
		return reflect.Value{}, fmt.Errorf("callable expected, got %s", value.DebugString())
	}
	callback := reflect.MakeFunc(targetType, func(args []reflect.Value) []reflect.Value {
		luaArgs := make([]gruntime.Value, 0, len(args))
		for _, arg := range args {
			luaArg, err := e.reflectToVMValue(arg)
			if err != nil {
				return reflectCallbackErrorResults(targetType, err)
			}
			luaArgs = append(luaArgs, luaArg)
		}
		luaResults, err := glua.Call(e.vmState, value, luaArgs...)
		if err != nil {
			return reflectCallbackErrorResults(targetType, err)
		}
		results := make([]reflect.Value, targetType.NumOut())
		luaResultIndex := 0
		for index := 0; index < targetType.NumOut(); index++ {
			outputType := targetType.Out(index)
			if outputType.Implements(errorReturnType) {
				results[index] = reflect.Zero(outputType)
				continue
			}
			if luaResultIndex >= len(luaResults) {
				results[index] = reflect.Zero(outputType)
				continue
			}
			convertedResult, err := e.vmValueToReflect(luaResults[luaResultIndex], outputType)
			if err != nil {
				return reflectCallbackErrorResults(targetType, err)
			}
			results[index] = convertedResult
			luaResultIndex++
		}
		return results
	})
	return callback, nil
}

func reflectCallbackErrorResults(targetType reflect.Type, err error) []reflect.Value {
	results := make([]reflect.Value, targetType.NumOut())
	for index := 0; index < targetType.NumOut(); index++ {
		outputType := targetType.Out(index)
		if outputType.Implements(errorReturnType) {
			results[index] = reflect.ValueOf(err)
			continue
		}
		results[index] = reflect.Zero(outputType)
	}
	return results
}

func (e *LuaEngine) vmTableToReflectMap(value gruntime.Value, targetType reflect.Type) (reflect.Value, error) {
	if value.Kind != gruntime.KindTable {
		return reflect.Value{}, fmt.Errorf("expected table, got %s", value.DebugString())
	}
	table, ok := value.Ref.(*gruntime.Table)
	if !ok || table == nil {
		return reflect.Zero(targetType), nil
	}
	result := reflect.MakeMap(targetType)
	key := gruntime.NilValue()
	for {
		nextKey, nextValue, ok, err := table.RawNext(key)
		if err != nil {
			return reflect.Value{}, err
		}
		if !ok {
			break
		}
		convertedKey, err := e.vmValueToReflect(nextKey, targetType.Key())
		if err != nil {
			return reflect.Value{}, err
		}
		convertedValue, err := e.vmValueToReflect(nextValue, targetType.Elem())
		if err != nil {
			return reflect.Value{}, err
		}
		result.SetMapIndex(convertedKey, convertedValue)
		key = nextKey
	}
	return result, nil
}

func (e *LuaEngine) vmValueToInterface(value gruntime.Value, targetType reflect.Type) (reflect.Value, error) {
	var result interface{}
	if object, ok := vmObjectFromValue(value); ok {
		result = object
		if result == nil {
			return reflect.Zero(targetType), nil
		}
		return reflect.ValueOf(result), nil
	}
	switch value.Kind {
	case gruntime.KindNil:
		result = nil
	case gruntime.KindBoolean:
		result = value.Bool
	case gruntime.KindInteger:
		result = value.Integer
	case gruntime.KindNumber:
		result = value.Number
	case gruntime.KindString:
		result = value.String
	case gruntime.KindTable:
		table, ok := value.Ref.(*gruntime.Table)
		if !ok || table == nil {
			result = nil
			break
		}
		var err error
		result, err = e.vmTableToInterface(table)
		if err != nil {
			return reflect.Value{}, err
		}
	default:
		result = value
	}
	if result == nil {
		return reflect.Zero(targetType), nil
	}
	return reflect.ValueOf(result), nil
}

func (e *LuaEngine) vmTableToInterface(table *gruntime.Table) (interface{}, error) {
	if vmTableLooksLikeArray(table) {
		values := make([]interface{}, 0)
		for index := int64(1); ; index++ {
			item := table.RawGetInteger(index)
			if item.IsNil() {
				break
			}
			convertedItem, err := e.vmValueToInterface(item, reflect.TypeOf((*interface{})(nil)).Elem())
			if err != nil {
				return nil, err
			}
			values = append(values, convertedItem.Interface())
		}
		return values, nil
	}
	values := make(map[string]interface{})
	key := gruntime.NilValue()
	for {
		nextKey, nextValue, ok, err := table.RawNext(key)
		if err != nil {
			return nil, err
		}
		if !ok {
			break
		}
		if nextKey.Kind != gruntime.KindString {
			key = nextKey
			continue
		}
		convertedValue, err := e.vmValueToInterface(nextValue, reflect.TypeOf((*interface{})(nil)).Elem())
		if err != nil {
			return nil, err
		}
		values[nextKey.String] = convertedValue.Interface()
		key = nextKey
	}
	return values, nil
}

func vmTableLooksLikeArray(table *gruntime.Table) bool {
	count := int64(0)
	key := gruntime.NilValue()
	for {
		nextKey, _, ok, err := table.RawNext(key)
		if err != nil || !ok {
			break
		}
		index, isInteger := nextKey.ToInteger()
		if !isInteger || index <= 0 {
			return false
		}
		count++
		key = nextKey
	}
	if count == 0 {
		return false
	}
	for index := int64(1); index <= count; index++ {
		if table.RawGetInteger(index).IsNil() {
			return false
		}
	}
	return true
}

func (e *LuaEngine) reflectToVMValue(value reflect.Value) (gruntime.Value, error) {
	if !value.IsValid() {
		return gruntime.NilValue(), nil
	}
	if value.Kind() == reflect.Interface {
		if value.IsNil() {
			return gruntime.NilValue(), nil
		}
		return e.reflectToVMValue(value.Elem())
	}
	switch value.Kind() {
	case reflect.Bool:
		return gruntime.BooleanValue(value.Bool()), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return gruntime.IntegerValue(value.Int()), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		uintValue := value.Uint()
		if uintValue > uint64(^uint64(0)>>1) {
			return gruntime.NilValue(), fmt.Errorf("uint value overflows lua integer")
		}
		return gruntime.IntegerValue(int64(uintValue)), nil
	case reflect.Float32, reflect.Float64:
		return gruntime.NumberValue(value.Convert(reflect.TypeOf(float64(0))).Float()), nil
	case reflect.String:
		return gruntime.StringValue(value.String()), nil
	case reflect.Slice:
		if value.IsNil() {
			return gruntime.NilValue(), nil
		}
		if value.Type().Elem().Kind() == reflect.Uint8 {
			return gruntime.StringValue(string(value.Bytes())), nil
		}
		table := gruntime.NewTable()
		for index := 0; index < value.Len(); index++ {
			item, err := e.reflectToVMValue(value.Index(index))
			if err != nil {
				return gruntime.NilValue(), err
			}
			table.RawSetInteger(int64(index+1), item)
		}
		return gruntime.ReferenceValue(gruntime.KindTable, table), nil
	case reflect.Map:
		return e.reflectMapToVMTable(value)
	case reflect.Struct:
		return e.reflectStructToVMTable(value)
	case reflect.Pointer:
		if value.IsNil() {
			return gruntime.NilValue(), nil
		}
		return bridge.BindReflectStruct(e.vmState, value.Interface())
	case reflect.Func:
		return e.bindVMMethodValue(value.Interface())
	default:
		valueInterface := value.Interface()
		return bridge.ValueOf(e.vmState, valueInterface)
	}
}

func (e *LuaEngine) reflectMapToVMTable(value reflect.Value) (gruntime.Value, error) {
	if value.IsNil() {
		return gruntime.NilValue(), nil
	}
	table := gruntime.NewTable()
	for _, key := range value.MapKeys() {
		luaKey, err := e.reflectToVMValue(key)
		if err != nil {
			return gruntime.NilValue(), err
		}
		luaValue, err := e.reflectToVMValue(value.MapIndex(key))
		if err != nil {
			return gruntime.NilValue(), err
		}
		if err := table.RawSet(luaKey, luaValue); err != nil {
			return gruntime.NilValue(), err
		}
	}
	return gruntime.ReferenceValue(gruntime.KindTable, table), nil
}

func (e *LuaEngine) reflectStructToVMTable(value reflect.Value) (gruntime.Value, error) {
	table := gruntime.NewTable()
	valueType := value.Type()
	for index := 0; index < value.NumField(); index++ {
		field := valueType.Field(index)
		if field.PkgPath != "" {
			continue
		}
		fieldValue, err := e.reflectToVMValue(value.Field(index))
		if err != nil {
			return gruntime.NilValue(), fmt.Errorf("field %s: %w", field.Name, err)
		}
		for _, name := range vmStructFieldNames(field) {
			table.RawSetString(name, fieldValue)
		}
	}
	return gruntime.ReferenceValue(gruntime.KindTable, table), nil
}

func vmObjectFromValue(value gruntime.Value) (interface{}, bool) {
	if value.Kind != gruntime.KindTable {
		return nil, false
	}
	table, ok := value.Ref.(*gruntime.Table)
	if !ok || table == nil {
		return nil, false
	}
	userdataValue := table.RawGetString("__userdata")
	if userdataValue.Kind != gruntime.KindUserdata {
		return nil, false
	}
	userdata, ok := userdataValue.Ref.(*gruntime.Userdata)
	if !ok || userdata == nil {
		return nil, false
	}
	proxy, ok := userdata.Data.(*bridge.ObjectProxy)
	if !ok || proxy == nil {
		return nil, false
	}
	return proxy.Object(), true
}

func (e *LuaEngine) installVMCoreTable(name string, methods map[string]gruntime.GoResultsFunction) error {
	if e.vmState == nil {
		return nil
	}
	moduleValue, err := glua.GetGlobal(e.vmState, name)
	if err != nil {
		return err
	}
	var moduleTable *gruntime.Table
	if moduleValue.Kind == gruntime.KindTable {
		if table, ok := moduleValue.Ref.(*gruntime.Table); ok {
			moduleTable = table
		}
	}
	if moduleTable == nil {
		moduleTable = gruntime.NewTable()
		if err := glua.SetGlobal(e.vmState, name, gruntime.ReferenceValue(gruntime.KindTable, moduleTable)); err != nil {
			return err
		}
	}
	for methodName, method := range methods {
		moduleTable.RawSetString(methodName, gruntime.ReferenceValue(gruntime.KindGoClosure, method))
	}
	return nil
}

func (e *LuaEngine) printVMConsole(prefix string, args []gruntime.Value) {
	if prefix != "" {
		fmt.Print(prefix)
	}
	for index, arg := range args {
		if index > 0 {
			fmt.Print(" ")
		}
		text, err := gruntime.ToString(arg)
		if err != nil {
			fmt.Print(arg.DebugString())
			continue
		}
		fmt.Print(text.String)
	}
	fmt.Println()
}

func vmTableField(table *gruntime.Table, field reflect.StructField) gruntime.Value {
	for _, name := range vmStructFieldNames(field) {
		value := table.RawGetString(name)
		if !value.IsNil() {
			return value
		}
	}
	return gruntime.NilValue()
}

func vmStructFieldNames(field reflect.StructField) []string {
	names := []string{field.Name, lowerFirstASCII(field.Name)}
	for _, tagName := range []string{"lua", "json"} {
		tagValue := field.Tag.Get(tagName)
		if tagValue == "" || tagValue == "-" {
			continue
		}
		if commaIndex := strings.IndexByte(tagValue, ','); commaIndex >= 0 {
			tagValue = tagValue[:commaIndex]
		}
		if tagValue != "" {
			names = append(names, tagValue)
		}
	}
	return names
}

func lowerFirstASCII(value string) string {
	if value == "" {
		return ""
	}
	first := value[0]
	if first >= 'A' && first <= 'Z' {
		return string(first+'a'-'A') + value[1:]
	}
	return value
}

func gluaStringArg(args []glua.Value, index int) string {
	if index >= len(args) {
		return ""
	}
	if args[index].Kind == gruntime.KindString {
		return args[index].String
	}
	return args[index].DebugString()
}

func gluaBoolArg(args []glua.Value, index int) bool {
	if index >= len(args) {
		return false
	}
	return args[index].Truthy()
}
