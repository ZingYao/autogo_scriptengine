package ffi

import (
	"fmt"

	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"
)

// FfiModule 是 go-lua-vm 迁移后的模块壳。
type FfiModule struct{}

func New() *FfiModule { return &FfiModule{} }

func (m *FfiModule) Name() string { return "ffi" }

func (m *FfiModule) IsAvailable() bool { return true }

func (m *FfiModule) Register(engine model.Engine) error {
	engine.RegisterMethod("ffi.cdef", "声明 C 定义（兼容占位）", func(_ string) (interface{}, string) {
		return nil, "ffi.cdef is not available in go-lua-vm"
	}, true)
	engine.RegisterMethod("ffi.load", "加载动态库（兼容占位）", func(name string) (interface{}, string) {
		return nil, fmt.Sprintf("ffi.load is not available: %s", name)
	}, true)
	engine.RegisterMethod("ffi.sizeof", "返回 C 类型大小（兼容占位）", func(_ string) int {
		return 0
	}, true)
	engine.RegisterMethod("ffi.new", "创建 C 对象（兼容占位）", func(_ string) interface{} {
		return nil
	}, true)
	return nil
}

func GetModule() model.Module { return &FfiModule{} }
