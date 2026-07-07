//go:build !android || !cgo

package utils

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"
)

// UtilsModule 在非 Android CGO 环境下提供纯 Go 可验证实现。
type UtilsModule struct{}

func New() *UtilsModule { return &UtilsModule{} }

func (m *UtilsModule) Name() string { return "utils" }

func (m *UtilsModule) IsAvailable() bool { return true }

func (m *UtilsModule) Register(engine model.Engine) error {
	engine.RegisterMethod("utils.logI", "记录一条 INFO 级别的日志", func(label string, message ...interface{}) { fmt.Println(append([]interface{}{label}, message...)...) }, true)
	engine.RegisterMethod("utils.logE", "记录一条 ERROR 级别的日志", func(label string, message ...interface{}) { fmt.Println(append([]interface{}{label}, message...)...) }, true)
	engine.RegisterMethod("utils.toast", "非 Android 环境输出 Toast 内容", func(message string, args ...int) { fmt.Println(message) }, true)
	engine.RegisterMethod("utils.alert", "非 Android 环境返回默认按钮", func(title, content string, buttons ...string) int { return 0 }, true)
	engine.RegisterMethod("utils.shell", "非 Android 环境禁用 shell 桥", func(cmd string) (string, error) { return "", fmt.Errorf("utils.shell requires Android AutoGo runtime") }, true)
	engine.RegisterMethod("utils.random", "返回指定范围内的随机整数", random, true)
	engine.RegisterMethod("utils.sleep", "暂停当前线程指定的毫秒数", func(i int) { time.Sleep(time.Duration(i) * time.Millisecond) }, true)
	engine.RegisterMethod("utils.i2s", "将整数转换为字符串", strconv.Itoa, true)
	engine.RegisterMethod("utils.s2i", "将字符串转换为整数", func(s string) int { v, _ := strconv.Atoi(strings.TrimSpace(s)); return v }, true)
	engine.RegisterMethod("utils.f2s", "将浮点数转换为字符串", func(f float64) string { return strconv.FormatFloat(f, 'f', -1, 64) }, true)
	engine.RegisterMethod("utils.s2f", "将字符串转换为浮点数", func(s string) float64 { v, _ := strconv.ParseFloat(strings.TrimSpace(s), 64); return v }, true)
	engine.RegisterMethod("utils.b2s", "将布尔值转换为字符串", strconv.FormatBool, true)
	engine.RegisterMethod("utils.s2b", "将字符串转换为布尔值", func(s string) bool { v, _ := strconv.ParseBool(strings.TrimSpace(s)); return v }, true)
	return nil
}

func random(min, max int) int {
	if min > max {
		min, max = max, min
	}
	return rand.Intn(max-min+1) + min
}

func GetModule() model.Module { return &UtilsModule{} }
