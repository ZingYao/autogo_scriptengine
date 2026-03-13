package imgui

import (
	"runtime"
	"strconv"

	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"

	"github.com/Dasongzi1366/AutoGo/imgui"
	"github.com/Dasongzi1366/AutoGo/utils"
	lua "github.com/yuin/gopher-lua"
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
	state := engine.GetState()

	imguiObj := state.NewTable()
	state.SetGlobal("imgui", imguiObj)

	imguiObj.RawSetString("init", state.NewFunction(func(L *lua.LState) int {
		imgui.Init()
		return 0
	}))

	imguiObj.RawSetString("close", state.NewFunction(func(L *lua.LState) int {
		imgui.Close()
		return 0
	}))

	imguiObj.RawSetString("alert", state.NewFunction(func(L *lua.LState) int {
		title := L.CheckString(1)
		content := L.CheckString(2)
		btn1Text := ""
		if L.GetTop() > 2 {
			btn1Text = L.CheckString(3)
		}
		btn2Text := ""
		if L.GetTop() > 3 {
			btn2Text = L.CheckString(4)
		}
		result := utils.Alert(title, content, btn1Text, btn2Text)
		L.Push(lua.LNumber(result))
		return 1
	}))

	imguiObj.RawSetString("toast", state.NewFunction(func(L *lua.LState) int {
		message := L.CheckString(1)
		utils.Toast(message)
		return 0
	}))

	imguiObj.RawSetString("drawRect", state.NewFunction(func(L *lua.LState) int {
		x1 := L.CheckInt(1)
		y1 := L.CheckInt(2)
		x2 := L.CheckInt(3)
		y2 := L.CheckInt(4)
		colorStr := L.CheckString(5)
		thickness := float32(L.CheckNumber(6))
		color := parseColorString(colorStr)
		imgui.DrawRect(x1, y1, x2, y2, color, thickness)
		return 0
	}))

	engine.RegisterMethod("imgui.init", "初始化ImGui", imgui.Init, true)
	engine.RegisterMethod("imgui.close", "关闭ImGui", imgui.Close, true)
	engine.RegisterMethod("imgui.alert", "显示对话框", func(title, content, btn1Text, btn2Text string) int {
		return utils.Alert(title, content, btn1Text, btn2Text)
	}, true)
	engine.RegisterMethod("imgui.toast", "显示Toast提示", func(message string) {
		utils.Toast(message)
	}, true)
	engine.RegisterMethod("imgui.drawRect", "绘制矩形", func(x1, y1, x2, y2 int, colorStr string, thickness float32) {
		color := parseColorString(colorStr)
		imgui.DrawRect(x1, y1, x2, y2, color, thickness)
	}, true)

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
