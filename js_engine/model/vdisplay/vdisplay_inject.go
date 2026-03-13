package vdisplay

import (
	"github.com/Dasongzi1366/AutoGo/vdisplay"
	"github.com/ZingYao/autogo_scriptengine/js_engine/model"
	"github.com/dop251/goja"
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
	vm := engine.GetVM()

	vdisplayObj := vm.NewObject()
	vm.Set("vdisplay", vdisplayObj)

	vdisplayObj.Set("create", func(call goja.FunctionCall) goja.Value {
		width := int(call.Argument(0).ToInteger())
		height := int(call.Argument(1).ToInteger())
		dpi := int(call.Argument(2).ToInteger())
		result := vdisplay.Create(width, height, dpi)
		return vm.ToValue(result)
	})

	vdisplayObj.Set("getDisplayId", func(call goja.FunctionCall) goja.Value {
		v := call.Argument(0).Export().(*vdisplay.Vdisplay)
		result := v.GetDisplayId()
		return vm.ToValue(result)
	})

	vdisplayObj.Set("launchApp", func(call goja.FunctionCall) goja.Value {
		v := call.Argument(0).Export().(*vdisplay.Vdisplay)
		packageName := call.Argument(1).String()
		result := v.LaunchApp(packageName)
		return vm.ToValue(result)
	})

	vdisplayObj.Set("setTitle", func(call goja.FunctionCall) goja.Value {
		v := call.Argument(0).Export().(*vdisplay.Vdisplay)
		title := call.Argument(1).String()
		v.SetTitle(title)
		return goja.Undefined()
	})

	vdisplayObj.Set("setTouchCallback", func(call goja.FunctionCall) goja.Value {
		v := call.Argument(0).Export().(*vdisplay.Vdisplay)
		fn, ok := goja.AssertFunction(call.Argument(1))
		if !ok {
			return goja.Undefined()
		}
		callback := func(x, y, action, displayId int) {
			fn(nil, vm.ToValue(x), vm.ToValue(y), vm.ToValue(action), vm.ToValue(displayId))
		}
		v.SetTouchCallback(callback)
		return goja.Undefined()
	})

	vdisplayObj.Set("showPreviewWindow", func(call goja.FunctionCall) goja.Value {
		v := call.Argument(0).Export().(*vdisplay.Vdisplay)
		rotated := call.Argument(1).ToBoolean()
		v.ShowPreviewWindow(rotated)
		return goja.Undefined()
	})

	vdisplayObj.Set("hidePreviewWindow", func(call goja.FunctionCall) goja.Value {
		v := call.Argument(0).Export().(*vdisplay.Vdisplay)
		v.HidePreviewWindow()
		return goja.Undefined()
	})

	vdisplayObj.Set("setPreviewWindowSize", func(call goja.FunctionCall) goja.Value {
		v := call.Argument(0).Export().(*vdisplay.Vdisplay)
		width := int(call.Argument(1).ToInteger())
		height := int(call.Argument(2).ToInteger())
		v.SetPreviewWindowSize(width, height)
		return goja.Undefined()
	})

	vdisplayObj.Set("setPreviewWindowPos", func(call goja.FunctionCall) goja.Value {
		v := call.Argument(0).Export().(*vdisplay.Vdisplay)
		x := int(call.Argument(1).ToInteger())
		y := int(call.Argument(2).ToInteger())
		v.SetPreviewWindowPos(x, y)
		return goja.Undefined()
	})

	vdisplayObj.Set("destroy", func(call goja.FunctionCall) goja.Value {
		v := call.Argument(0).Export().(*vdisplay.Vdisplay)
		v.Destroy()
		return goja.Undefined()
	})

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
