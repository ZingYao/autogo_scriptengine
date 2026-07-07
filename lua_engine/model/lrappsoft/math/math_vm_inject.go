package math

import (
	stdmath "math"
	"strconv"

	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"
)

// MathModule 是 go-lua-vm 迁移后的模块壳。
type MathModule struct{}

func New() *MathModule { return &MathModule{} }

func (m *MathModule) Name() string { return "math" }

func (m *MathModule) IsAvailable() bool { return true }

func (m *MathModule) Register(engine model.Engine) error {
	engine.RegisterMethod("math.tointeger", "将值转换为整数", func(value interface{}) interface{} {
		switch typedValue := value.(type) {
		case int:
			return int64(typedValue)
		case int64:
			return typedValue
		case float64:
			return int64(typedValue)
		case string:
			integerValue, err := strconv.ParseInt(typedValue, 10, 64)
			if err != nil {
				return nil
			}
			return integerValue
		default:
			return nil
		}
	}, true)
	engine.RegisterMethod("math.type", "获取数字类型", func(value interface{}) interface{} {
		switch typedValue := value.(type) {
		case int, int64:
			return "integer"
		case float64:
			if typedValue == stdmath.Trunc(typedValue) {
				return "integer"
			}
			return "float"
		default:
			return nil
		}
	}, true)
	engine.RegisterMethod("math.ult", "无符号比较两个整数", func(left, right int64) bool {
		return uint64(left) < uint64(right)
	}, true)
	return nil
}

func GetModule() model.Module { return &MathModule{} }
