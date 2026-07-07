//go:build ignore
// +build ignore

package ffi

import (
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"

	lua "github.com/yuin/gopher-lua"
)

// FfiModule ffi 模块（Foreign Function Interface）
type FfiModule struct{}

// Name 返回模块名称
func (m *FfiModule) Name() string {
	return "ffi"
}

// IsAvailable 检查模块是否可用
func (m *FfiModule) IsAvailable() bool {
	return true
}

// Register 向引擎注册方法
func (m *FfiModule) Register(engine model.Engine) error {
	state := engine.GetState()

	// 获取 package.preload 表
	packagePreload := state.GetGlobal("package")
	if packagePreload == lua.LNil {
		packagePreload = state.NewTable()
		state.SetGlobal("package", packagePreload)
	}

	packageTable := packagePreload.(*lua.LTable)
	preloadTable := packageTable.RawGetString("preload")
	if preloadTable == lua.LNil {
		preloadTable = state.NewTable()
		packageTable.RawSetString("preload", preloadTable)
	}

	// 创建 ffi 模块表
	ffiModuleTable := state.NewTable()

	// 注册 ffi.cdef - 定义C语言类型和函数（占位符）
	ffiModuleTable.RawSetString("cdef", state.NewFunction(func(L *lua.LState) int {
		// ffi.cdef 在 gopher-lua 中不可用
		// 这是 LuaJIT 的特性
		L.Push(lua.LNil)
		L.Push(lua.LString("ffi.cdef 在 AutoGo 中不可用，因为 gopher-lua 不支持 LuaJIT 的 FFI 特性"))
		return 2
	}))

	// 注册 ffi.load - 加载动态库（占位符）
	ffiModuleTable.RawSetString("load", state.NewFunction(func(L *lua.LState) int {
		// ffi.load 在 gopher-lua 中不可用
		// 这是 LuaJIT 的特性
		L.Push(lua.LNil)
		L.Push(lua.LString("ffi.load 在 AutoGo 中不可用，因为 gopher-lua 不支持 LuaJIT 的 FFI 特性"))
		return 2
	}))

	// 注册 ffi.sizeof - 获取类型大小（占位符）
	ffiModuleTable.RawSetString("sizeof", state.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LNumber(0))
		return 1
	}))

	// 注册 ffi.new - 创建 cdata 对象（占位符）
	ffiModuleTable.RawSetString("new", state.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LNil)
		return 1
	}))

	// 注册到 package.preload
	preloadTable.(*lua.LTable).RawSetString("ffi", state.NewFunction(func(L *lua.LState) int {
		L.Push(ffiModuleTable)
		return 1
	}))

	// 同时注册为全局变量 ffi（兼容直接调用）
	state.SetGlobal("ffi", ffiModuleTable)

	engine.RegisterMethod("ffi.cdef", "定义 C 类型和函数占位实现", func(string) (interface{}, string) {
		return nil, "ffi.cdef 在 AutoGo 中不可用，因为 gopher-lua 不支持 LuaJIT 的 FFI 特性"
	}, true)
	engine.RegisterMethod("ffi.load", "加载动态库占位实现", func(string) (interface{}, string) {
		return nil, "ffi.load 在 AutoGo 中不可用，因为 gopher-lua 不支持 LuaJIT 的 FFI 特性"
	}, true)
	engine.RegisterMethod("ffi.sizeof", "获取类型大小占位实现", func(string) int {
		return 0
	}, true)
	engine.RegisterMethod("ffi.new", "创建 cdata 对象占位实现", func(string) interface{} {
		return nil
	}, true)

	return nil
}
