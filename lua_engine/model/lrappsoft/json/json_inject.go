//go:build ignore
// +build ignore

package json

import (
	"encoding/json"

	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"

	lua "github.com/yuin/gopher-lua"
)

// JsonModule json 模块（懒人精灵兼容）
type JsonModule struct{}

// Name 返回模块名称
func (m *JsonModule) Name() string {
	return "json"
}

// IsAvailable 检查模块是否可用
func (m *JsonModule) IsAvailable() bool {
	return true
}

// Register 向引擎注册方法
func (m *JsonModule) Register(engine model.Engine) error {
	state := engine.GetState()

	// 创建 jsonLib 模块表
	jsonLib := state.NewTable()

	// 注册 jsonLib.encode - 把lua表格编码成json字符串
	jsonLib.RawSetString("encode", state.NewFunction(func(L *lua.LState) int {
		tb := L.CheckAny(1)
		result, err := luaValueToJSON(tb)
		if err != nil {
			L.Push(lua.LNil)
			L.Push(lua.LString(err.Error()))
			return 2
		}
		L.Push(lua.LString(result))
		return 1
	}))

	// 注册 jsonLib.decode - 把json字符串转换成lua表格
	jsonLib.RawSetString("decode", state.NewFunction(func(L *lua.LState) int {
		jsonStr := L.CheckString(1)
		var result interface{}
		err := json.Unmarshal([]byte(jsonStr), &result)
		if err != nil {
			L.Push(lua.LNil)
			L.Push(lua.LString(err.Error()))
			return 2
		}
		luaValue, err := jsonToLuaValue(L, result)
		if err != nil {
			L.Push(lua.LNil)
			L.Push(lua.LString(err.Error()))
			return 2
		}
		L.Push(luaValue)
		return 1
	}))

	// 注册 jsonLib 到全局
	state.SetGlobal("jsonLib", jsonLib)

	// 注册到方法注册表
	engine.RegisterMethod("jsonLib.encode", "把lua表格编码成json字符串", func(value interface{}) (string, error) {
		result, err := json.Marshal(value)
		if err != nil {
			return "", err
		}
		return string(result), nil
	}, true)
	engine.RegisterMethod("jsonLib.decode", "把json字符串转换成lua表格", func(jsonStr string) (interface{}, error) {
		var result interface{}
		err := json.Unmarshal([]byte(jsonStr), &result)
		if err != nil {
			return nil, err
		}
		return result, nil
	}, true)

	return nil
}

func luaValueToJSON(value lua.LValue) (string, error) {
	data, err := luaValueToGoValue(value)
	if err != nil {
		return "", err
	}
	result, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(result), nil
}

func luaValueToGoValue(value lua.LValue) (interface{}, error) {
	switch v := value.(type) {
	case *lua.LNilType:
		return nil, nil
	case lua.LBool:
		return bool(v), nil
	case lua.LNumber:
		return float64(v), nil
	case lua.LString:
		return string(v), nil
	case *lua.LTable:
		if isArray(v) {
			return luaTableToArray(v)
		}
		return luaTableToMap(v)
	case *lua.LUserData:
		return v.Value, nil
	default:
		return nil, nil
	}
}

func isArray(table *lua.LTable) bool {
	if table == nil {
		return false
	}

	hasStringKey := false
	table.ForEach(func(key, value lua.LValue) {
		if _, ok := key.(lua.LString); ok {
			hasStringKey = true
		}
	})
	if hasStringKey {
		return false
	}

	length := table.Len()
	if length == 0 {
		hasNumberKey := false
		table.ForEach(func(key, value lua.LValue) {
			if _, ok := key.(lua.LNumber); ok {
				hasNumberKey = true
			}
		})
		return hasNumberKey
	}

	for i := 1; i <= length; i++ {
		value := table.RawGetInt(i)
		if value == lua.LNil {
			return false
		}
	}

	return true
}

func luaTableToArray(table *lua.LTable) ([]interface{}, error) {
	result := make([]interface{}, 0, table.Len())
	for i := 1; i <= table.Len(); i++ {
		value := table.RawGetInt(i)
		converted, err := luaValueToGoValue(value)
		if err != nil {
			return nil, err
		}
		result = append(result, converted)
	}
	return result, nil
}

func luaTableToMap(table *lua.LTable) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	table.ForEach(func(key, value lua.LValue) {
		keyStr, ok := key.(lua.LString)
		if !ok {
			return
		}
		converted, err := luaValueToGoValue(value)
		if err != nil {
			return
		}
		result[string(keyStr)] = converted
	})
	return result, nil
}

func jsonToLuaValue(L *lua.LState, data interface{}) (lua.LValue, error) {
	switch v := data.(type) {
	case nil:
		return lua.LNil, nil
	case bool:
		return lua.LBool(v), nil
	case float64:
		return lua.LNumber(v), nil
	case string:
		return lua.LString(v), nil
	case []interface{}:
		table := L.NewTable()
		for _, item := range v {
			luaValue, err := jsonToLuaValue(L, item)
			if err != nil {
				return nil, err
			}
			table.Append(luaValue)
		}
		return table, nil
	case map[string]interface{}:
		table := L.NewTable()
		for key, value := range v {
			luaValue, err := jsonToLuaValue(L, value)
			if err != nil {
				return nil, err
			}
			L.RawSet(table, lua.LString(key), luaValue)
		}
		return table, nil
	default:
		return lua.LNil, nil
	}
}

// GetModule 获取模块实例
func GetModule() model.Module {
	return &JsonModule{}
}
