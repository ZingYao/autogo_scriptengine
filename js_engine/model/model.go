package model

import "github.com/dop251/goja"

// Module 模块接口，所有注入模块都需要实现此接口
type Module interface {
	// Name 返回模块名称
	Name() string

	// Register 向引擎注册模块的方法
	// 返回错误：如果注册失败（如 Cgo 模块在非 Android 平台）
	Register(engine Engine) error

	// IsAvailable 返回模块是否可用（用于检查依赖）
	IsAvailable() bool
}

// Engine 引擎接口
type Engine interface {
	// GetVM 获取 JavaScript 虚拟机
	GetVM() *goja.Runtime

	// RegisterMethod 注册方法
	RegisterMethod(name, description string, fn interface{}, isStatic bool)
}
