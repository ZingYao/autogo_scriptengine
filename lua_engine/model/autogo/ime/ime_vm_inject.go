package ime

import (
	"fmt"

	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"
)

// ImeModule 是 go-lua-vm 迁移后的模块壳。
type ImeModule struct{}

func New() *ImeModule { return &ImeModule{} }

func (m *ImeModule) Name() string { return "ime" }

func (m *ImeModule) IsAvailable() bool { return true }

func (m *ImeModule) Register(engine model.Engine) error {
	// 当前 AutoGo 参考目录缺少 AutoGo/core，不能直接 import AutoGo/ime，否则主工程无法编译。
	engine.RegisterMethod("ime.getClipText", "获取剪切板内容", func() (string, error) { return "", errImeUnavailable() }, true)
	engine.RegisterMethod("ime.setClipText", "设置剪切板内容", func(string) (bool, error) { return false, errImeUnavailable() }, true)
	engine.RegisterMethod("ime.keyAction", "模拟按键", func(int) error { return errImeUnavailable() }, true)
	engine.RegisterMethod("ime.inputText", "输入文本", func(string, ...int) error { return errImeUnavailable() }, true)
	engine.RegisterMethod("ime.getIMEList", "获取输入法列表", func() ([]string, error) { return nil, errImeUnavailable() }, true)
	engine.RegisterMethod("ime.setCurrentIME", "设置当前输入法", func(string) error { return errImeUnavailable() }, true)
	return nil
}

func GetModule() model.Module { return &ImeModule{} }

func errImeUnavailable() error {
	return fmt.Errorf("AutoGo ime module is unavailable: missing github.com/Dasongzi1366/AutoGo/core dependency in AutoGo reference directory")
}
