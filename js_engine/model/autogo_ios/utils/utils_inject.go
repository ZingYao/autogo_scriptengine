package utils

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/ZingYao/autogo_scriptengine/js_engine/model"
	"github.com/dop251/goja"
)

// UtilsModule 在远程 AutoGo/utils 为 Android CGO 包时提供 iOS 安全实现。
type UtilsModule struct{}

func (m *UtilsModule) Name() string { return "utils" }

func (m *UtilsModule) IsAvailable() bool { return true }

func (m *UtilsModule) Register(engine model.Engine) error {
	vm := engine.GetVM()
	utilsObj := vm.NewObject()
	vm.Set("utils", utilsObj)

	utilsObj.Set("logI", func(call goja.FunctionCall) goja.Value {
		fmt.Println(joinArgs(call.Arguments))
		return goja.Undefined()
	})
	utilsObj.Set("logE", func(call goja.FunctionCall) goja.Value {
		fmt.Println(joinArgs(call.Arguments))
		return goja.Undefined()
	})
	utilsObj.Set("toast", func(call goja.FunctionCall) goja.Value {
		fmt.Println(call.Argument(0).String())
		return goja.Undefined()
	})
	utilsObj.Set("alert", func(call goja.FunctionCall) goja.Value { return vm.ToValue(0) })
	utilsObj.Set("shell", func(call goja.FunctionCall) goja.Value {
		panic(vm.NewGoError(fmt.Errorf("utils.shell requires Android AutoGo runtime")))
	})
	utilsObj.Set("random", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(random(int(call.Argument(0).ToInteger()), int(call.Argument(1).ToInteger())))
	})
	utilsObj.Set("sleep", func(call goja.FunctionCall) goja.Value {
		time.Sleep(time.Duration(call.Argument(0).ToInteger()) * time.Millisecond)
		return goja.Undefined()
	})
	utilsObj.Set("i2s", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(strconv.Itoa(int(call.Argument(0).ToInteger())))
	})
	utilsObj.Set("s2i", func(call goja.FunctionCall) goja.Value {
		v, _ := strconv.Atoi(strings.TrimSpace(call.Argument(0).String()))
		return vm.ToValue(v)
	})
	utilsObj.Set("f2s", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(strconv.FormatFloat(call.Argument(0).ToFloat(), 'f', -1, 64))
	})
	utilsObj.Set("s2f", func(call goja.FunctionCall) goja.Value {
		v, _ := strconv.ParseFloat(strings.TrimSpace(call.Argument(0).String()), 64)
		return vm.ToValue(v)
	})
	utilsObj.Set("b2s", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(strconv.FormatBool(call.Argument(0).ToBoolean()))
	})
	utilsObj.Set("s2b", func(call goja.FunctionCall) goja.Value {
		v, _ := strconv.ParseBool(strings.TrimSpace(call.Argument(0).String()))
		return vm.ToValue(v)
	})

	engine.RegisterMethod("utils.logI", "记录一条 INFO 级别日志", func(args ...interface{}) { fmt.Println(args...) }, true)
	engine.RegisterMethod("utils.logE", "记录一条 ERROR 级别日志", func(args ...interface{}) { fmt.Println(args...) }, true)
	engine.RegisterMethod("utils.toast", "非 Android 环境输出 Toast 内容", func(message string, args ...int) { fmt.Println(message) }, true)
	engine.RegisterMethod("utils.alert", "非 Android 环境返回默认按钮", func(title, content string, buttons ...string) int { return 0 }, true)
	engine.RegisterMethod("utils.shell", "非 Android 环境禁用 shell 桥", func(cmd string) (string, error) { return "", fmt.Errorf("utils.shell requires Android AutoGo runtime") }, true)
	engine.RegisterMethod("utils.random", "返回指定范围内的随机整数", random, true)
	engine.RegisterMethod("utils.sleep", "暂停指定毫秒数", func(i int) { time.Sleep(time.Duration(i) * time.Millisecond) }, true)
	engine.RegisterMethod("utils.i2s", "整数转字符串", strconv.Itoa, true)
	engine.RegisterMethod("utils.s2i", "字符串转整数", func(s string) int { v, _ := strconv.Atoi(strings.TrimSpace(s)); return v }, true)
	engine.RegisterMethod("utils.f2s", "浮点数转字符串", func(f float64) string { return strconv.FormatFloat(f, 'f', -1, 64) }, true)
	engine.RegisterMethod("utils.s2f", "字符串转浮点数", func(s string) float64 { v, _ := strconv.ParseFloat(strings.TrimSpace(s), 64); return v }, true)
	engine.RegisterMethod("utils.b2s", "布尔值转字符串", strconv.FormatBool, true)
	engine.RegisterMethod("utils.s2b", "字符串转布尔值", func(s string) bool { v, _ := strconv.ParseBool(strings.TrimSpace(s)); return v }, true)
	return nil
}

func joinArgs(values []goja.Value) string {
	parts := make([]string, 0, len(values))
	for _, value := range values {
		parts = append(parts, value.String())
	}
	return strings.Join(parts, " ")
}

func random(min, max int) int {
	if min > max {
		min, max = max, min
	}
	return rand.Intn(max-min+1) + min
}
