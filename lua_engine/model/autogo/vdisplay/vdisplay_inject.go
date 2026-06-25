package vdisplay

import (
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"

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
		v := vdisplay.Create(L.CheckInt(1), L.CheckInt(2), L.CheckInt(3))
		if v == nil {
			L.Push(lua.LNil)
			return 1
		}
		L.Push(wrapVdisplay(L, v))
		return 1
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

func wrapVdisplay(L *lua.LState, v *vdisplay.Vdisplay) lua.LValue {
	obj := L.NewTable()
	obj.RawSetString("getDisplayId", L.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LNumber(v.GetDisplayId()))
		return 1
	}))
	obj.RawSetString("launchApp", L.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LBool(v.LaunchApp(L.CheckString(1))))
		return 1
	}))
	obj.RawSetString("setTitle", L.NewFunction(func(L *lua.LState) int {
		v.SetTitle(L.CheckString(1))
		return 0
	}))
	obj.RawSetString("setTouchCallback", L.NewFunction(func(L *lua.LState) int {
		callback := L.CheckFunction(1)
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
	obj.RawSetString("showPreviewWindow", L.NewFunction(func(L *lua.LState) int {
		v.ShowPreviewWindow(L.OptBool(1, false))
		return 0
	}))
	obj.RawSetString("hidePreviewWindow", L.NewFunction(func(L *lua.LState) int {
		v.HidePreviewWindow()
		return 0
	}))
	obj.RawSetString("setPreviewWindowSize", L.NewFunction(func(L *lua.LState) int {
		v.SetPreviewWindowSize(L.CheckInt(1), L.CheckInt(2))
		return 0
	}))
	obj.RawSetString("setPreviewWindowPos", L.NewFunction(func(L *lua.LState) int {
		v.SetPreviewWindowPos(L.CheckInt(1), L.CheckInt(2))
		return 0
	}))
	obj.RawSetString("destroy", L.NewFunction(func(L *lua.LState) int {
		v.Destroy()
		return 0
	}))
	return obj
}
