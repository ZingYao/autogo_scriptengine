package console

import (
	"fmt"

	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"
)

// ConsoleModule 是 go-lua-vm 迁移后的模块壳。
type ConsoleModule struct {
	titleVisible bool
	locked       bool
}

func New() *ConsoleModule { return &ConsoleModule{} }

func (m *ConsoleModule) Name() string { return "console" }

func (m *ConsoleModule) IsAvailable() bool { return true }

func (m *ConsoleModule) Register(engine model.Engine) error {
	engine.RegisterMethod("console.show", "显示控制台悬浮窗", func() bool {
		return true
	}, true)
	engine.RegisterMethod("console.showTitle", "显示或者隐藏控制台标题栏", func(args ...interface{}) bool {
		m.titleVisible = true
		if len(args) > 0 {
			m.titleVisible = consoleBoolArg(args[0])
		}
		return true
	}, true)
	engine.RegisterMethod("console.lockConsole", "锁定控制台窗口", func() {
		m.locked = true
	}, true)
	engine.RegisterMethod("console.unlockConsole", "解除锁定控制台窗口", func() {
		m.locked = false
	}, true)
	engine.RegisterMethod("console.dismiss", "关闭控制台窗口", func() bool {
		return true
	}, true)
	engine.RegisterMethod("console.setPos", "设置控制台窗口的位置和大小", func(args ...interface{}) {
	}, true)
	engine.RegisterMethod("console.println", "打印日志到控制台窗口", func(args ...interface{}) {
		if len(args) > 0 {
			fmt.Println(args...)
		}
	}, true)
	engine.RegisterMethod("console.clearLog", "清除日志", func() {
	}, true)
	engine.RegisterMethod("console.setTitle", "设置控制台标题", func(title string) {
	}, true)
	return nil
}

func GetModule() model.Module { return &ConsoleModule{} }

func consoleBoolArg(value interface{}) bool {
	if boolValue, ok := value.(bool); ok {
		return boolValue
	}
	return false
}
