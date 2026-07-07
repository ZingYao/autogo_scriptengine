package json

import (
	"encoding/json"

	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"
)

// JsonModule 是 go-lua-vm 迁移后的模块壳。
type JsonModule struct{}

func New() *JsonModule { return &JsonModule{} }

func (m *JsonModule) Name() string { return "json" }

func (m *JsonModule) IsAvailable() bool { return true }

func (m *JsonModule) Register(engine model.Engine) error {
	engine.RegisterMethod("jsonLib.encode", "将 Lua 值序列化为 JSON 字符串", func(value interface{}) (string, error) {
		data, err := json.Marshal(value)
		if err != nil {
			return "", err
		}
		return string(data), nil
	}, true)
	engine.RegisterMethod("jsonLib.decode", "将 JSON 字符串解析为 Lua 值", func(value string) (interface{}, error) {
		var result interface{}
		if err := json.Unmarshal([]byte(value), &result); err != nil {
			return nil, err
		}
		return result, nil
	}, true)
	return nil
}

func GetModule() model.Module { return &JsonModule{} }
