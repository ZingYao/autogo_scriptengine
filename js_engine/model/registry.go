package model

import (
	"sync"
)

// ModuleRegistry 模块注册表
type ModuleRegistry struct {
	modules map[string]Module
	mu      sync.RWMutex
}

// NewModuleRegistry 创建新的模块注册表
func NewModuleRegistry() *ModuleRegistry {
	return &ModuleRegistry{
		modules: make(map[string]Module),
	}
}

// RegisterModule 注册模块
func (r *ModuleRegistry) RegisterModule(m Module) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.modules[m.Name()] = m
}

// GetModule 获取模块
func (r *ModuleRegistry) GetModule(name string) (Module, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	m, ok := r.modules[name]
	return m, ok
}

// ListModules 列出所有模块
func (r *ModuleRegistry) ListModules() []string {
	r.mu.RLock()
	defer r.mu.RUnlock()
	names := make([]string, 0, len(r.modules))
	for name := range r.modules {
		names = append(names, name)
	}
	return names
}

// Count 返回模块数量
func (r *ModuleRegistry) Count() int {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return len(r.modules)
}

// Clear 清空所有模块
func (r *ModuleRegistry) Clear() {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.modules = make(map[string]Module)
}

// RegisterModules 批量注册模块
func (r *ModuleRegistry) RegisterModules(modules []Module) {
	for _, m := range modules {
		r.RegisterModule(m)
	}
}

// GetModules 获取所有模块对象
func (r *ModuleRegistry) GetModules() map[string]Module {
	r.mu.RLock()
	defer r.mu.RUnlock()
	result := make(map[string]Module, len(r.modules))
	for name, module := range r.modules {
		result[name] = module
	}
	return result
}
