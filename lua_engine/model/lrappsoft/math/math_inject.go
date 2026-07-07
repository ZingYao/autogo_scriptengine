//go:build ignore
// +build ignore

package math

import (
	"math"
	"strconv"

	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"

	lua "github.com/yuin/gopher-lua"
)

// MathModule math 模块（懒人精灵兼容）
type MathModule struct{}

// Name 返回模块名称
func (m *MathModule) Name() string {
	return "math"
}

// IsAvailable 检查模块是否可用
func (m *MathModule) IsAvailable() bool {
	return true
}

// Register 向引擎注册方法
func (m *MathModule) Register(engine model.Engine) error {
	state := engine.GetState()

	// 获取或创建 math 表
	mathTable := state.GetGlobal("math")
	if mathTable == lua.LNil {
		mathTable = state.NewTable()
		state.SetGlobal("math", mathTable)
	}

	// 转换为 LTable
	mathTbl, ok := mathTable.(*lua.LTable)
	if !ok {
		mathTbl = state.NewTable()
		state.SetGlobal("math", mathTbl)
	}

	// 注册 math.tointeger - 将值转换为整数
	mathTbl.RawSetString("tointeger", state.NewFunction(func(L *lua.LState) int {
		value := L.CheckAny(1)

		switch v := value.(type) {
		case lua.LNumber:
			// 数字类型，转换为整数
			L.Push(lua.LNumber(int64(v)))
		case lua.LString:
			// 字符串类型，尝试转换为整数
			str := string(v)
			if intVal, err := strconv.ParseInt(str, 10, 64); err == nil {
				L.Push(lua.LNumber(intVal))
			} else {
				// 转换失败，返回 nil
				L.Push(lua.LNil)
			}
		default:
			// 其他类型，返回 nil
			L.Push(lua.LNil)
		}
		return 1
	}))

	// 注册 math.type - 获取数字的类型
	mathTbl.RawSetString("type", state.NewFunction(func(L *lua.LState) int {
		value := L.CheckAny(1)

		switch v := value.(type) {
		case lua.LNumber:
			// 检查是否为整数
			if float64(v) == math.Trunc(float64(v)) {
				L.Push(lua.LString("integer"))
			} else {
				L.Push(lua.LString("float"))
			}
		default:
			// 其他类型，返回 nil
			L.Push(lua.LNil)
		}
		return 1
	}))

	// 注册 math.ult - 无符号比较两个整数
	mathTbl.RawSetString("ult", state.NewFunction(func(L *lua.LState) int {
		m := L.CheckInt64(1)
		n := L.CheckInt64(2)

		// 转换为无符号整数进行比较
		um := uint64(m)
		un := uint64(n)

		if um < un {
			L.Push(lua.LTrue)
		} else {
			L.Push(lua.LFalse)
		}
		return 1
	}))

	// 注册到方法注册表
	engine.RegisterMethod("math.tointeger", "将值转换为整数", func(value interface{}) (interface{}, error) {
		switch v := value.(type) {
		case int:
			return int64(v), nil
		case int64:
			return v, nil
		case float64:
			return int64(v), nil
		case string:
			if intVal, err := strconv.ParseInt(v, 10, 64); err == nil {
				return intVal, nil
			}
			return nil, nil
		default:
			return nil, nil
		}
	}, true)

	engine.RegisterMethod("math.type", "获取数字的类型", func(value interface{}) (string, error) {
		switch v := value.(type) {
		case int, int64:
			return "integer", nil
		case float64:
			if v == math.Trunc(v) {
				return "integer", nil
			}
			return "float", nil
		default:
			return "", nil
		}
	}, true)

	engine.RegisterMethod("math.ult", "无符号比较两个整数", func(m, n int64) (bool, error) {
		um := uint64(m)
		un := uint64(n)
		return um < un, nil
	}, true)

	return nil
}
