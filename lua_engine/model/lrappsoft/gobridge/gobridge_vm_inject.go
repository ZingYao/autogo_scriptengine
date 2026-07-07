package gobridge

import (
	"encoding/hex"
	"fmt"

	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"
)

// GoBridgeModule 是 go-lua-vm 迁移后的模块壳。
type GoBridgeModule struct{ engine model.Engine }

func New() *GoBridgeModule                                  { return &GoBridgeModule{} }
func NewGoBridgeModule(engine model.Engine) *GoBridgeModule { return &GoBridgeModule{engine: engine} }

func (m *GoBridgeModule) Name() string      { return "gobridge" }
func (m *GoBridgeModule) IsAvailable() bool { return true }
func (m *GoBridgeModule) Register(engine model.Engine) error {
	engine.RegisterMethod("gobridge.tobytes", "字符串转十六进制", func(value string) string {
		return hex.EncodeToString([]byte(value))
	}, true)
	engine.RegisterMethod("gobridge.tostring", "十六进制转字符串", func(value string) string {
		if len(value)%2 == 1 {
			value = value[:len(value)-1]
		}
		data, err := hex.DecodeString(value)
		if err != nil {
			return ""
		}
		return string(data)
	}, true)
	engine.RegisterMethod("gobridge.call", "调用动态库函数", func(libraryPath, functionName string, args ...interface{}) (interface{}, error) {
		return nil, fmt.Errorf("failed to load library %s for function %s", libraryPath, functionName)
	}, true)
	return nil
}
func GetModule() model.Module { return &GoBridgeModule{} }
