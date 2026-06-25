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
		v := vdisplay.Create(
			int(call.Argument(0).ToInteger()),
			int(call.Argument(1).ToInteger()),
			int(call.Argument(2).ToInteger()),
		)
		if v == nil {
			return goja.Null()
		}
		return wrapVdisplay(vm, v)
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

func wrapVdisplay(vm *goja.Runtime, v *vdisplay.Vdisplay) goja.Value {
	obj := vm.NewObject()
	obj.Set("getDisplayId", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(v.GetDisplayId())
	})
	obj.Set("launchApp", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(v.LaunchApp(call.Argument(0).String()))
	})
	obj.Set("setTitle", func(call goja.FunctionCall) goja.Value {
		v.SetTitle(call.Argument(0).String())
		return goja.Undefined()
	})
	obj.Set("setTouchCallback", func(call goja.FunctionCall) goja.Value {
		fn, ok := goja.AssertFunction(call.Argument(0))
		if !ok {
			return goja.Undefined()
		}
		v.SetTouchCallback(func(x, y, action, displayId int) {
			_, _ = fn(nil, vm.ToValue(x), vm.ToValue(y), vm.ToValue(action), vm.ToValue(displayId))
		})
		return goja.Undefined()
	})
	obj.Set("showPreviewWindow", func(call goja.FunctionCall) goja.Value {
		rotated := false
		if len(call.Arguments) > 0 {
			rotated = call.Argument(0).ToBoolean()
		}
		v.ShowPreviewWindow(rotated)
		return goja.Undefined()
	})
	obj.Set("hidePreviewWindow", func(call goja.FunctionCall) goja.Value {
		v.HidePreviewWindow()
		return goja.Undefined()
	})
	obj.Set("setPreviewWindowSize", func(call goja.FunctionCall) goja.Value {
		v.SetPreviewWindowSize(int(call.Argument(0).ToInteger()), int(call.Argument(1).ToInteger()))
		return goja.Undefined()
	})
	obj.Set("setPreviewWindowPos", func(call goja.FunctionCall) goja.Value {
		v.SetPreviewWindowPos(int(call.Argument(0).ToInteger()), int(call.Argument(1).ToInteger()))
		return goja.Undefined()
	})
	obj.Set("destroy", func(call goja.FunctionCall) goja.Value {
		v.Destroy()
		return goja.Undefined()
	})
	return vm.ToValue(obj)
}
