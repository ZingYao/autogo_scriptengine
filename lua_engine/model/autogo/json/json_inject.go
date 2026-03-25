package json

import (
	"encoding/json"
	"fmt"

	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"

	lua "github.com/yuin/gopher-lua"
)

// JsonModule json 模块
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

	jsonObj := state.NewTable()
	state.SetGlobal("json", jsonObj)

	jsonObj.RawSetString("stringify", state.NewFunction(func(L *lua.LState) int {
		value := L.CheckAny(1)
		result, err := luaValueToJSON(value)
		if err != nil {
			L.Push(lua.LNil)
			L.Push(lua.LString(err.Error()))
			return 2
		}
		L.Push(lua.LString(result))
		return 1
	}))

	jsonObj.RawSetString("stringifyArr", state.NewFunction(func(L *lua.LState) int {
		value := L.CheckAny(1)
		result, err := luaValueToJSONArray(value)
		if err != nil {
			L.Push(lua.LNil)
			L.Push(lua.LString(err.Error()))
			return 2
		}
		L.Push(lua.LString(result))
		return 1
	}))

	jsonObj.RawSetString("stringifyObj", state.NewFunction(func(L *lua.LState) int {
		value := L.CheckAny(1)
		result, err := luaValueToJSONObject(value)
		if err != nil {
			L.Push(lua.LNil)
			L.Push(lua.LString(err.Error()))
			return 2
		}
		L.Push(lua.LString(result))
		return 1
	}))

	jsonObj.RawSetString("parse", state.NewFunction(func(L *lua.LState) int {
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

	jsonObj.RawSetString("format", state.NewFunction(func(L *lua.LState) int {
		value := L.CheckAny(1)
		result, err := luaValueToJSONFormatted(value)
		if err != nil {
			L.Push(lua.LNil)
			L.Push(lua.LString(err.Error()))
			return 2
		}
		L.Push(lua.LString(result))
		return 1
	}))

	engine.RegisterMethod("json.stringify", "将 Lua 值序列化为 JSON 字符串（自动判断数组或对象）", func(value lua.LValue) (string, error) {
		return luaValueToJSON(value)
	}, true)
	engine.RegisterMethod("json.stringifyArr", "将 Lua 值强制序列化为 JSON 数组", func(value lua.LValue) (string, error) {
		return luaValueToJSONArray(value)
	}, true)
	engine.RegisterMethod("json.stringifyObj", "将 Lua 值强制序列化为 JSON 对象", func(value lua.LValue) (string, error) {
		return luaValueToJSONObject(value)
	}, true)
	engine.RegisterMethod("json.parse", "将 JSON 字符串解析为 Lua 值", func(jsonStr string) (lua.LValue, error) {
		L := state
		var result interface{}
		err := json.Unmarshal([]byte(jsonStr), &result)
		if err != nil {
			return nil, err
		}
		return jsonToLuaValue(L, result)
	}, true)
	engine.RegisterMethod("json.format", "将 Lua 值格式化序列化为 JSON 字符串", func(value lua.LValue) (string, error) {
		return luaValueToJSONFormatted(value)
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

func luaValueToJSONArray(value lua.LValue) (string, error) {
	data, err := luaValueToGoArray(value)
	if err != nil {
		return "", err
	}
	result, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(result), nil
}

func luaValueToJSONObject(value lua.LValue) (string, error) {
	data, err := luaValueToGoObject(value)
	if err != nil {
		return "", err
	}
	result, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(result), nil
}

func luaValueToJSONFormatted(value lua.LValue) (string, error) {
	data, err := luaValueToGoValue(value)
	if err != nil {
		return "", err
	}
	result, err := json.MarshalIndent(data, "", "  ")
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

func luaValueToGoArray(value lua.LValue) (interface{}, error) {
	switch v := value.(type) {
	case *lua.LNilType:
		return []interface{}{}, nil
	case *lua.LTable:
		return luaTableToArray(v)
	default:
		// 非表类型，包装成数组
		goValue, err := luaValueToGoValue(value)
		if err != nil {
			return nil, err
		}
		return []interface{}{goValue}, nil
	}
}

func luaValueToGoObject(value lua.LValue) (interface{}, error) {
	switch v := value.(type) {
	case *lua.LNilType:
		return map[string]interface{}{}, nil
	case *lua.LTable:
		return luaTableToMap(v)
	default:
		return nil, fmt.Errorf("value is not a table")
	}
}

func isArray(table *lua.LTable) bool {
	if table == nil {
		return false
	}

	// 检查是否有字符串键，如果有字符串键则是对象
	hasStringKey := false
	table.ForEach(func(key, value lua.LValue) {
		if _, ok := key.(lua.LString); ok {
			hasStringKey = true
		}
	})
	if hasStringKey {
		return false
	}

	// 检查是否是连续的数字数组
	length := table.Len()
	if length == 0 {
		// 空 table，检查是否有数字键
		hasNumberKey := false
		table.ForEach(func(key, value lua.LValue) {
			if _, ok := key.(lua.LNumber); ok {
				hasNumberKey = true
			}
		})
		return hasNumberKey
	}

	// 检查从 1 到 length 的所有索引是否存在
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
