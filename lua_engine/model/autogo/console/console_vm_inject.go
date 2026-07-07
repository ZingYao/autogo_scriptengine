package console

import (
	"fmt"

	autogoconsole "github.com/Dasongzi1366/AutoGo/console"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"
)

// ConsoleModule 为 Lua VM 暴露 AutoGo 控制台对象能力。
type ConsoleModule struct{}

func New() *ConsoleModule { return &ConsoleModule{} }

func (m *ConsoleModule) Name() string { return "console" }

func (m *ConsoleModule) IsAvailable() bool { return true }

// Register 注册控制台构造、对象方法和模块方法。
func (m *ConsoleModule) Register(engine model.Engine) error {
	engine.RegisterMethod("console.new", "创建控制台对象", newConsoleObject, true)
	engine.RegisterMethod("console.println", "输出一行内容", consolePrintln, true)
	engine.RegisterMethod("console.setTextSize", "设置控制台文本大小", consoleSetTextSize, true)
	engine.RegisterMethod("console.setTextColor", "设置控制台文本颜色", consoleSetTextColor, true)
	engine.RegisterMethod("console.setWindowSize", "设置控制台窗口大小", consoleSetWindowSize, true)
	engine.RegisterMethod("console.setWindowPosition", "设置控制台窗口位置", consoleSetWindowPosition, true)
	engine.RegisterMethod("console.setWindowColor", "设置控制台窗口颜色", consoleSetWindowColor, true)
	engine.RegisterMethod("console.show", "显示控制台", consoleShow, true)
	engine.RegisterMethod("console.hide", "隐藏控制台", consoleHide, true)
	engine.RegisterMethod("console.clear", "清空控制台", consoleClear, true)
	engine.RegisterMethod("console.isVisible", "返回控制台是否可见", consoleIsVisible, true)
	engine.RegisterMethod("console.destroy", "销毁控制台", consoleDestroy, true)
	return nil
}

func GetModule() model.Module { return &ConsoleModule{} }

func newConsoleObject() map[string]interface{} {
	return wrapConsole(autogoconsole.New())
}

func wrapConsole(c *autogoconsole.Console) map[string]interface{} {
	c = normalizeConsole(c)
	return map[string]interface{}{
		"__console":         c,
		"println":           func(args ...interface{}) { c.Println(args...) },
		"setTextSize":       func(size int) map[string]interface{} { return wrapConsole(c.SetTextSize(size)) },
		"setTextColor":      func(color string) map[string]interface{} { return wrapConsole(c.SetTextColor(color)) },
		"setWindowSize":     func(width, height int) map[string]interface{} { return wrapConsole(c.SetWindowSize(width, height)) },
		"setWindowPosition": func(x, y int) map[string]interface{} { return wrapConsole(c.SetWindowPosition(x, y)) },
		"setWindowColor":    func(color string) map[string]interface{} { return wrapConsole(c.SetWindowColor(color)) },
		"show":              func() { c.Show() },
		"hide":              func() { c.Hide() },
		"clear":             func() { c.Clear() },
		"isVisible":         func() bool { return c.IsVisible() },
		"destroy":           func() { c.Destroy() },
	}
}

func normalizeConsole(c *autogoconsole.Console) *autogoconsole.Console {
	if c != nil {
		return c
	}
	return &autogoconsole.Console{}
}

func consoleFromValue(value map[string]interface{}) (*autogoconsole.Console, error) {
	raw, ok := value["__console"]
	if !ok {
		return nil, fmt.Errorf("console object expected")
	}
	c, ok := raw.(*autogoconsole.Console)
	if !ok {
		return nil, fmt.Errorf("invalid console object")
	}
	return c, nil
}

func consolePrintln(value map[string]interface{}, args ...interface{}) error {
	c, err := consoleFromValue(value)
	if err != nil {
		return err
	}
	c.Println(args...)
	return nil
}

func consoleSetTextSize(value map[string]interface{}, size int) (map[string]interface{}, error) {
	c, err := consoleFromValue(value)
	if err != nil {
		return nil, err
	}
	return wrapConsole(c.SetTextSize(size)), nil
}

func consoleSetTextColor(value map[string]interface{}, color string) (map[string]interface{}, error) {
	c, err := consoleFromValue(value)
	if err != nil {
		return nil, err
	}
	return wrapConsole(c.SetTextColor(color)), nil
}

func consoleSetWindowSize(value map[string]interface{}, width, height int) (map[string]interface{}, error) {
	c, err := consoleFromValue(value)
	if err != nil {
		return nil, err
	}
	return wrapConsole(c.SetWindowSize(width, height)), nil
}

func consoleSetWindowPosition(value map[string]interface{}, x, y int) (map[string]interface{}, error) {
	c, err := consoleFromValue(value)
	if err != nil {
		return nil, err
	}
	return wrapConsole(c.SetWindowPosition(x, y)), nil
}

func consoleSetWindowColor(value map[string]interface{}, color string) (map[string]interface{}, error) {
	c, err := consoleFromValue(value)
	if err != nil {
		return nil, err
	}
	return wrapConsole(c.SetWindowColor(color)), nil
}

func consoleShow(value map[string]interface{}) error {
	c, err := consoleFromValue(value)
	if err != nil {
		return err
	}
	c.Show()
	return nil
}

func consoleHide(value map[string]interface{}) error {
	c, err := consoleFromValue(value)
	if err != nil {
		return err
	}
	c.Hide()
	return nil
}

func consoleClear(value map[string]interface{}) error {
	c, err := consoleFromValue(value)
	if err != nil {
		return err
	}
	c.Clear()
	return nil
}

func consoleIsVisible(value map[string]interface{}) (bool, error) {
	c, err := consoleFromValue(value)
	if err != nil {
		return false, err
	}
	return c.IsVisible(), nil
}

func consoleDestroy(value map[string]interface{}) error {
	c, err := consoleFromValue(value)
	if err != nil {
		return err
	}
	c.Destroy()
	return nil
}
