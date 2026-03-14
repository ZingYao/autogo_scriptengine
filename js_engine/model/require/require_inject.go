package require

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
	"sync"

	"github.com/dop251/goja"
)

// RequireModule 处理 CommonJS 的 require 功能
type RequireModule struct {
	vm         *goja.Runtime
	cache      map[string]goja.Value
	mu         sync.RWMutex
	fileSys    fs.FS
	initialDir string // 初始目录，用于主脚本
}

// NewRequireModule 创建新的 require 模块
func NewRequireModule(vm *goja.Runtime, fileSys fs.FS) *RequireModule {
	return &RequireModule{
		vm:         vm,
		cache:      make(map[string]goja.Value),
		fileSys:    fileSys,
		initialDir: "scripts", // 默认初始目录
	}
}

// SetInitialDir 设置初始目录
func (r *RequireModule) SetInitialDir(dir string) {
	r.initialDir = dir
}

// Register 注册 require 函数到 JavaScript 运行时
func (r *RequireModule) Register() error {
	// 注册 require 函数
	r.vm.Set("require", r.requireFunc)

	// 注册 module 对象
	moduleObj := r.vm.NewObject()
	moduleObj.Set("exports", r.vm.NewObject())
	r.vm.Set("module", moduleObj)

	// 注册 exports 对象（module.exports 的简写）
	r.vm.Set("exports", r.vm.NewObject())

	// 注册 __dirname 和 __filename（当前脚本路径）
	r.vm.Set("__dirname", "")
	r.vm.Set("__filename", "")

	return nil
}

// requireFunc 实现 require 函数
func (r *RequireModule) requireFunc(call goja.FunctionCall) goja.Value {
	if len(call.Arguments) == 0 {
		panic(r.vm.NewGoError(fmt.Errorf("require() requires a module name")))
	}

	modulePath := call.Argument(0).String()

	// 解析模块路径
	resolvedPath, err := r.resolveModule(modulePath)
	if err != nil {
		panic(r.vm.NewGoError(fmt.Errorf("cannot find module '%s': %v", modulePath, err)))
	}

	// 检查缓存
	r.mu.RLock()
	if cached, exists := r.cache[resolvedPath]; exists {
		r.mu.RUnlock()
		return cached
	}
	r.mu.RUnlock()

	// 加载模块
	moduleValue, err := r.loadModule(resolvedPath)
	if err != nil {
		panic(r.vm.NewGoError(fmt.Errorf("failed to load module '%s': %v", modulePath, err)))
	}

	// 缓存模块
	r.mu.Lock()
	r.cache[resolvedPath] = moduleValue
	r.mu.Unlock()

	return moduleValue
}

// resolveModule 解析模块路径
func (r *RequireModule) resolveModule(modulePath string) (string, error) {
	// 处理相对路径
	if strings.HasPrefix(modulePath, "./") || strings.HasPrefix(modulePath, "../") {
		// 获取当前调用者的 __dirname
		currentDir := r.vm.Get("__dirname")
		baseDir := ""
		if currentDir != goja.Undefined() {
			baseDir = currentDir.String()
		}

		// 如果 __dirname 为空，使用初始目录
		if baseDir == "" {
			baseDir = r.initialDir
		}

		// 解析相对路径
		absPath := filepath.Join(baseDir, modulePath)
		return r.resolveExtension(absPath)
	}

	// 处理绝对路径
	if filepath.IsAbs(modulePath) {
		return r.resolveExtension(modulePath)
	}

	// 处理 node_modules 路径（简化版本，只查找当前目录）
	// 如果是纯文件名，先尝试在当前目录或初始目录查找
	if baseDir := r.vm.Get("__dirname"); baseDir != goja.Undefined() {
		dir := baseDir.String()
		if dir == "" {
			dir = r.initialDir
		}
		absPath := filepath.Join(dir, modulePath)
		if result, err := r.resolveExtension(absPath); err == nil {
			return result, nil
		}
	}

	return r.resolveExtension(modulePath)
}

// resolveExtension 解析文件扩展名
func (r *RequireModule) resolveExtension(path string) (string, error) {
	// 尝试直接读取
	if _, err := fs.Stat(r.fileSys, path); err == nil {
		return path, nil
	}

	// 尝试添加 .js 扩展名
	jsPath := path + ".js"
	if _, err := fs.Stat(r.fileSys, jsPath); err == nil {
		return jsPath, nil
	}

	// 尝试添加 .json 扩展名
	jsonPath := path + ".json"
	if _, err := fs.Stat(r.fileSys, jsonPath); err == nil {
		return jsonPath, nil
	}

	// 尝试 index.js
	indexPath := filepath.Join(path, "index.js")
	if _, err := fs.Stat(r.fileSys, indexPath); err == nil {
		return indexPath, nil
	}

	return "", fmt.Errorf("module not found: %s", path)
}

// loadModule 加载模块
func (r *RequireModule) loadModule(path string) (goja.Value, error) {
	// 读取文件内容
	content, err := fs.ReadFile(r.fileSys, path)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}

	// 获取模块目录和文件名
	moduleDir := filepath.Dir(path)
	moduleFile := filepath.Base(path)

	// 设置 __dirname 和 __filename
	r.vm.Set("__dirname", moduleDir)
	r.vm.Set("__filename", moduleFile)

	// 创建新的 module 和 exports 对象
	moduleObj := r.vm.NewObject()
	exportsObj := r.vm.NewObject()
	moduleObj.Set("exports", exportsObj)

	// 保存旧的 module 和 exports
	oldModule := r.vm.Get("module")
	oldExports := r.vm.Get("exports")

	// 设置新的 module 和 exports
	r.vm.Set("module", moduleObj)
	r.vm.Set("exports", exportsObj)

	// 包装代码为立即执行函数
	wrappedCode := fmt.Sprintf("(function(module, exports, __dirname, __filename) {\n%s\n})(module, module.exports, '%s', '%s')",
		string(content), moduleDir, moduleFile)

	// 执行代码
	_, err = r.vm.RunString(wrappedCode)
	if err != nil {
		// 恢复旧的 module 和 exports
		r.vm.Set("module", oldModule)
		r.vm.Set("exports", oldExports)
		return nil, fmt.Errorf("failed to execute module: %v", err)
	}

	// 获取导出的值
	exportsValue := moduleObj.Get("exports")

	// 恢复旧的 module 和 exports
	r.vm.Set("module", oldModule)
	r.vm.Set("exports", oldExports)

	return exportsValue, nil
}

// ClearCache 清除模块缓存
func (r *RequireModule) ClearCache() {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.cache = make(map[string]goja.Value)
}

// GetCacheSize 获取缓存大小
func (r *RequireModule) GetCacheSize() int {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return len(r.cache)
}
