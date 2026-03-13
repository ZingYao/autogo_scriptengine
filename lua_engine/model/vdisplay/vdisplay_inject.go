package vdisplay

import (
	"app/lua_engine/model"

	"github.com/Dasongzi1366/AutoGo/vdisplay"
	lua "github.com/yuin/gopher-lua"
)

// VdisplayModule vdisplay 模块
type VdisplayModule struct{}

// Name 返回模块名称
func (m *VdisplayModule) Name() string {
	return "vdisplay"
}

// IsAvailable 检查模块是否可用
func (m *VdisplayModule) IsAvailable() bool {
	return true
}

// Register 向引擎注册方法
func (m *VdisplayModule) Register(engine model.Engine) error {
	state := engine.GetState()

	vdisplayObj := state.NewTable()
	state.SetGlobal("vdisplay", vdisplayObj)

	vdisplayObj.RawSetString("create", state.NewFunction(func(L *lua.LState) int {
		width := L.CheckInt(1)
		height := L.CheckInt(2)
		dpi := L.CheckInt(3)
		v := vdisplay.Create(width, height, dpi)
		ud := L.NewUserData()
		ud.Value = v
		L.Push(ud)
		return 1
	}))

	vdisplayObj.RawSetString("getDisplayId", state.NewFunction(func(L *lua.LState) int {
		v := L.CheckUserData(1).Value.(*vdisplay.Vdisplay)
		result := v.GetDisplayId()
		L.Push(lua.LNumber(result))
		return 1
	}))

	vdisplayObj.RawSetString("launchApp", state.NewFunction(func(L *lua.LState) int {
		v := L.CheckUserData(1).Value.(*vdisplay.Vdisplay)
		packageName := L.CheckString(2)
		v.LaunchApp(packageName)
		return 0
	}))

	vdisplayObj.RawSetString("setTitle", state.NewFunction(func(L *lua.LState) int {
		v := L.CheckUserData(1).Value.(*vdisplay.Vdisplay)
		title := L.CheckString(2)
		v.SetTitle(title)
		return 0
	}))

	vdisplayObj.RawSetString("setTouchCallback", state.NewFunction(func(L *lua.LState) int {
		v := L.CheckUserData(1).Value.(*vdisplay.Vdisplay)
		callback := L.CheckFunction(2)
		v.SetTouchCallback(func(x, y, action, displayId int) {
			L.Push(callback)
			L.Push(lua.LNumber(x))
			L.Push(lua.LNumber(y))
			L.Push(lua.LNumber(action))
			L.Push(lua.LNumber(displayId))
			L.Call(4, 0)
		})
		return 0
	}))

	vdisplayObj.RawSetString("showPreviewWindow", state.NewFunction(func(L *lua.LState) int {
		v := L.CheckUserData(1).Value.(*vdisplay.Vdisplay)
		rotated := false
		if L.GetTop() > 1 {
			rotated = L.CheckBool(2)
		}
		v.ShowPreviewWindow(rotated)
		return 0
	}))

	vdisplayObj.RawSetString("hidePreviewWindow", state.NewFunction(func(L *lua.LState) int {
		v := L.CheckUserData(1).Value.(*vdisplay.Vdisplay)
		v.HidePreviewWindow()
		return 0
	}))

	vdisplayObj.RawSetString("setPreviewWindowSize", state.NewFunction(func(L *lua.LState) int {
		v := L.CheckUserData(1).Value.(*vdisplay.Vdisplay)
		width := L.CheckInt(2)
		height := L.CheckInt(3)
		v.SetPreviewWindowSize(width, height)
		return 0
	}))

	vdisplayObj.RawSetString("setPreviewWindowPos", state.NewFunction(func(L *lua.LState) int {
		v := L.CheckUserData(1).Value.(*vdisplay.Vdisplay)
		x := L.CheckInt(2)
		y := L.CheckInt(3)
		v.SetPreviewWindowPos(x, y)
		return 0
	}))

	vdisplayObj.RawSetString("destroy", state.NewFunction(func(L *lua.LState) int {
		v := L.CheckUserData(1).Value.(*vdisplay.Vdisplay)
		v.Destroy()
		return 0
	}))

	engine.RegisterMethod("vdisplay.create", "创建一个虚拟显示设备", vdisplay.Create, true)
	engine.RegisterMethod("vdisplay.getDisplayId", "获取虚拟显示设备的DisplayId", (*vdisplay.Vdisplay).GetDisplayId, true)
	engine.RegisterMethod("vdisplay.launchApp", "启动指定包名的应用到虚拟显示设备内", (*vdisplay.Vdisplay).LaunchApp, true)
	engine.RegisterMethod("vdisplay.setTitle", "设置预览窗口标题", (*vdisplay.Vdisplay).SetTitle, true)
	engine.RegisterMethod("vdisplay.setTouchCallback", "设置触控回调", (*vdisplay.Vdisplay).SetTouchCallback, true)
	engine.RegisterMethod("vdisplay.showPreviewWindow", "显示预览窗口", (*vdisplay.Vdisplay).ShowPreviewWindow, true)
	engine.RegisterMethod("vdisplay.hidePreviewWindow", "隐藏预览窗口", (*vdisplay.Vdisplay).HidePreviewWindow, true)
	engine.RegisterMethod("vdisplay.setPreviewWindowSize", "设置预览窗口大小", (*vdisplay.Vdisplay).SetPreviewWindowSize, true)
	engine.RegisterMethod("vdisplay.setPreviewWindowPos", "设置预览窗口位置", (*vdisplay.Vdisplay).SetPreviewWindowPos, true)
	engine.RegisterMethod("vdisplay.destroy", "销毁指定的虚拟显示设备", (*vdisplay.Vdisplay).Destroy, true)

	return nil
}
