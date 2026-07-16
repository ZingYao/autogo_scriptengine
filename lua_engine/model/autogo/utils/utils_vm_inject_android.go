//go:build android && cgo

package utils

import (
	autogoutils "github.com/Dasongzi1366/AutoGo/utils"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"
)

// UtilsModule 为 Lua VM 暴露 AutoGo/utils 方法。
type UtilsModule struct{}

func New() *UtilsModule { return &UtilsModule{} }

func (m *UtilsModule) Name() string { return "utils" }

func (m *UtilsModule) IsAvailable() bool { return true }

// Register 注册 AutoGo/utils 导出函数。
func (m *UtilsModule) Register(engine model.Engine) error {
	engine.RegisterMethod("utils.logI", "记录一条 INFO 级别的日志", autogoutils.LogI, true)
	engine.RegisterMethod("utils.logE", "记录一条 ERROR 级别的日志", autogoutils.LogE, true)
	engine.RegisterMethod("utils.toast", "显示 Toast 提示", func(message string, args ...int) {
		x, y, duration := -1, -1, -1
		if len(args) > 0 {
			x = args[0]
		}
		if len(args) > 1 {
			y = args[1]
		}
		if len(args) > 2 {
			duration = args[2]
		}
		autogoutils.Toast(message, x, y, duration)
	}, true)
	engine.RegisterMethod("utils.alert", "显示 Alert 对话框", func(title, content string, buttons ...string) int {
		btn1Text, btn2Text := "", ""
		if len(buttons) > 0 {
			btn1Text = buttons[0]
		}
		if len(buttons) > 1 {
			btn2Text = buttons[1]
		}
		return autogoutils.Alert(title, content, btn1Text, btn2Text)
	}, true)
	engine.RegisterMethod("utils.shell", "执行 shell 命令并返回输出", autogoutils.Shell, true)
	engine.RegisterMethod("utils.random", "返回指定范围内的随机整数", autogoutils.Random, true)
	engine.RegisterMethod("utils.sleep", "暂停当前线程指定的毫秒数", autogoutils.Sleep, true)
	engine.RegisterMethod("utils.i2s", "将整数转换为字符串", autogoutils.I2s, true)
	engine.RegisterMethod("utils.s2i", "将字符串转换为整数", autogoutils.S2i, true)
	engine.RegisterMethod("utils.f2s", "将浮点数转换为字符串", autogoutils.F2s, true)
	engine.RegisterMethod("utils.s2f", "将字符串转换为浮点数", autogoutils.S2f, true)
	engine.RegisterMethod("utils.b2s", "将布尔值转换为字符串", autogoutils.B2s, true)
	engine.RegisterMethod("utils.s2b", "将字符串转换为布尔值", autogoutils.S2b, true)
	return nil
}

func GetModule() model.Module { return &UtilsModule{} }
