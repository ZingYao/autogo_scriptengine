package imgui

import (
	"app/js_engine/model"
	"runtime"
	"strconv"

	"github.com/Dasongzi1366/AutoGo/imgui"
	"github.com/Dasongzi1366/AutoGo/utils"
	"github.com/dop251/goja"
)

// ImGuiModule imgui 模块
type ImGuiModule struct{}

// Name 返回模块名称
func (m *ImGuiModule) Name() string {
	return "imgui"
}

// IsAvailable 检查模块是否可用（只支持 Android 平台）
func (m *ImGuiModule) IsAvailable() bool {
	return runtime.GOOS == "android"
}

// Register 向引擎注册方法
func (m *ImGuiModule) Register(engine model.Engine) error {
	vm := engine.GetVM()

	imguiObj := vm.NewObject()

	imguiObj.Set("init", func(call goja.FunctionCall) goja.Value {
		imgui.Init()
		return goja.Undefined()
	})

	imguiObj.Set("close", func(call goja.FunctionCall) goja.Value {
		imgui.Close()
		return goja.Undefined()
	})

	imguiObj.Set("alert", func(call goja.FunctionCall) goja.Value {
		title := call.Argument(0).String()
		content := call.Argument(1).String()
		btn1Text := ""
		btn2Text := ""
		if len(call.Arguments) >= 3 {
			btn1Text = call.Argument(2).String()
		}
		if len(call.Arguments) >= 4 {
			btn2Text = call.Argument(3).String()
		}
		result := utils.Alert(title, content, btn1Text, btn2Text)
		return vm.ToValue(result)
	})

	imguiObj.Set("toast", func(call goja.FunctionCall) goja.Value {
		message := call.Argument(0).String()
		utils.Toast(message)
		return goja.Undefined()
	})

	imguiObj.Set("drawRect", func(call goja.FunctionCall) goja.Value {
		x1 := int(call.Argument(0).ToInteger())
		y1 := int(call.Argument(1).ToInteger())
		x2 := int(call.Argument(2).ToInteger())
		y2 := int(call.Argument(3).ToInteger())
		colorStr := call.Argument(4).String()
		thickness := float32(1.0)
		if len(call.Arguments) >= 6 {
			thickness = float32(call.Argument(5).ToFloat())
		}
		color := parseColorString(colorStr)
		imgui.DrawRect(x1, y1, x2, y2, color, thickness)
		return goja.Undefined()
	})

	vm.Set("imgui", imguiObj)

	return nil
}

func parseColorString(colorStr string) uint32 {
	if len(colorStr) == 0 {
		return 0xFFFFFFFF
	}
	if colorStr[0] == '#' {
		colorStr = colorStr[1:]
	}
	color, err := strconv.ParseUint(colorStr, 16, 32)
	if err != nil {
		return 0xFFFFFFFF
	}
	return uint32(color)
}
