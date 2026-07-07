package json

import (
	"encoding/json"
	"fmt"

	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"
)

// JsonModule 是 go-lua-vm 迁移后的模块壳。
type JsonModule struct{}

func New() *JsonModule { return &JsonModule{} }

func (m *JsonModule) Name() string { return "json" }

func (m *JsonModule) IsAvailable() bool { return true }

func (m *JsonModule) Register(engine model.Engine) error {
	engine.RegisterMethod("json.stringify", "将 Lua 值序列化为 JSON 字符串", func(value interface{}) (string, error) {
		data, err := json.Marshal(value)
		if err != nil {
			return "", err
		}
		return string(data), nil
	}, true)
	engine.RegisterMethod("json.stringifyArr", "将 Lua 值强制序列化为 JSON 数组", func(value interface{}) (string, error) {
		arrayValue, ok := value.([]interface{})
		if !ok {
			arrayValue = []interface{}{value}
		}
		data, err := json.Marshal(arrayValue)
		if err != nil {
			return "", err
		}
		return string(data), nil
	}, true)
	engine.RegisterMethod("json.stringifyObj", "将 Lua 值强制序列化为 JSON 对象", func(value interface{}) (string, error) {
		objectValue, ok := value.(map[string]interface{})
		if !ok {
			return "", fmt.Errorf("value is not a table")
		}
		data, err := json.Marshal(objectValue)
		if err != nil {
			return "", err
		}
		return string(data), nil
	}, true)
	engine.RegisterMethod("json.parse", "将 JSON 字符串解析为 Lua 值", func(value string) (interface{}, error) {
		var result interface{}
		if err := json.Unmarshal([]byte(value), &result); err != nil {
			return nil, err
		}
		return result, nil
	}, true)
	engine.RegisterMethod("json.format", "格式化 JSON 字符串", func(value interface{}) (string, error) {
		data, err := json.MarshalIndent(value, "", "  ")
		if err != nil {
			return "", err
		}
		return string(data), nil
	}, true)
	return nil
}

func GetModule() model.Module { return &JsonModule{} }
