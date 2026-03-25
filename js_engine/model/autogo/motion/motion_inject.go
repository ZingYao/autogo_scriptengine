package motion

import (
	"github.com/ZingYao/autogo_scriptengine/js_engine/model"

	"github.com/Dasongzi1366/AutoGo/motion"
	"github.com/dop251/goja"
)

// MotionModule motion 模块
type MotionModule struct{}

// Name 返回模块名称
func (m *MotionModule) Name() string {
	return "motion"
}

// IsAvailable 检查模块是否可用
func (m *MotionModule) IsAvailable() bool {
	return true
}

// Register 向引擎注册方法
func (m *MotionModule) Register(engine model.Engine) error {
	vm := engine.GetVM()

	vm.Set("touchDown", func(call goja.FunctionCall) goja.Value {
		x := int(call.Argument(0).ToInteger())
		y := int(call.Argument(1).ToInteger())
		fingerID := 0
		if len(call.Arguments) > 2 {
			fingerID = int(call.Argument(2).ToInteger())
		}
		displayId := 0
		if len(call.Arguments) > 3 {
			displayId = int(call.Argument(3).ToInteger())
		}
		motion.TouchDown(x, y, fingerID, displayId)
		return goja.Undefined()
	})

	vm.Set("touchMove", func(call goja.FunctionCall) goja.Value {
		x := int(call.Argument(0).ToInteger())
		y := int(call.Argument(1).ToInteger())
		fingerID := 0
		if len(call.Arguments) > 2 {
			fingerID = int(call.Argument(2).ToInteger())
		}
		displayId := 0
		if len(call.Arguments) > 3 {
			displayId = int(call.Argument(3).ToInteger())
		}
		motion.TouchMove(x, y, fingerID, displayId)
		return goja.Undefined()
	})

	vm.Set("touchUp", func(call goja.FunctionCall) goja.Value {
		x := int(call.Argument(0).ToInteger())
		y := int(call.Argument(1).ToInteger())
		fingerID := 0
		if len(call.Arguments) > 2 {
			fingerID = int(call.Argument(2).ToInteger())
		}
		displayId := 0
		if len(call.Arguments) > 3 {
			displayId = int(call.Argument(3).ToInteger())
		}
		motion.TouchUp(x, y, fingerID, displayId)
		return goja.Undefined()
	})

	vm.Set("click", func(call goja.FunctionCall) goja.Value {
		x := int(call.Argument(0).ToInteger())
		y := int(call.Argument(1).ToInteger())
		fingerID := 0
		if len(call.Arguments) > 2 {
			fingerID = int(call.Argument(2).ToInteger())
		}
		displayId := 0
		if len(call.Arguments) > 3 {
			displayId = int(call.Argument(3).ToInteger())
		}
		motion.Click(x, y, fingerID, displayId)
		return goja.Undefined()
	})

	vm.Set("longClick", func(call goja.FunctionCall) goja.Value {
		x := int(call.Argument(0).ToInteger())
		y := int(call.Argument(1).ToInteger())
		duration := 500
		if len(call.Arguments) > 2 {
			duration = int(call.Argument(2).ToInteger())
		}
		fingerID := 0
		if len(call.Arguments) > 3 {
			fingerID = int(call.Argument(3).ToInteger())
		}
		displayId := 0
		if len(call.Arguments) > 4 {
			displayId = int(call.Argument(4).ToInteger())
		}
		motion.LongClick(x, y, duration, fingerID, displayId)
		return goja.Undefined()
	})

	vm.Set("swipe", func(call goja.FunctionCall) goja.Value {
		x1 := int(call.Argument(0).ToInteger())
		y1 := int(call.Argument(1).ToInteger())
		x2 := int(call.Argument(2).ToInteger())
		y2 := int(call.Argument(3).ToInteger())
		duration := int(call.Argument(4).ToInteger())
		fingerID := 0
		if len(call.Arguments) > 5 {
			fingerID = int(call.Argument(5).ToInteger())
		}
		displayId := 0
		if len(call.Arguments) > 6 {
			displayId = int(call.Argument(6).ToInteger())
		}
		motion.Swipe(x1, y1, x2, y2, duration, fingerID, displayId)
		return goja.Undefined()
	})

	vm.Set("swipe2", func(call goja.FunctionCall) goja.Value {
		x1 := int(call.Argument(0).ToInteger())
		y1 := int(call.Argument(1).ToInteger())
		x2 := int(call.Argument(2).ToInteger())
		y2 := int(call.Argument(3).ToInteger())
		duration := int(call.Argument(4).ToInteger())
		fingerID := 0
		if len(call.Arguments) > 5 {
			fingerID = int(call.Argument(5).ToInteger())
		}
		displayId := 0
		if len(call.Arguments) > 6 {
			displayId = int(call.Argument(6).ToInteger())
		}
		motion.Swipe2(x1, y1, x2, y2, duration, fingerID, displayId)
		return goja.Undefined()
	})

	vm.Set("home", func(call goja.FunctionCall) goja.Value {
		displayId := 0
		if len(call.Arguments) > 0 {
			displayId = int(call.Argument(0).ToInteger())
		}
		motion.Home(displayId)
		return goja.Undefined()
	})

	vm.Set("back", func(call goja.FunctionCall) goja.Value {
		displayId := 0
		if len(call.Arguments) > 0 {
			displayId = int(call.Argument(0).ToInteger())
		}
		motion.Back(displayId)
		return goja.Undefined()
	})

	vm.Set("recents", func(call goja.FunctionCall) goja.Value {
		displayId := 0
		if len(call.Arguments) > 0 {
			displayId = int(call.Argument(0).ToInteger())
		}
		motion.Recents(displayId)
		return goja.Undefined()
	})

	vm.Set("powerDialog", func(call goja.FunctionCall) goja.Value {
		motion.PowerDialog()
		return goja.Undefined()
	})

	vm.Set("notifications", func(call goja.FunctionCall) goja.Value {
		motion.Notifications()
		return goja.Undefined()
	})

	vm.Set("quickSettings", func(call goja.FunctionCall) goja.Value {
		motion.QuickSettings()
		return goja.Undefined()
	})

	vm.Set("volumeUp", func(call goja.FunctionCall) goja.Value {
		displayId := 0
		if len(call.Arguments) > 0 {
			displayId = int(call.Argument(0).ToInteger())
		}
		motion.VolumeUp(displayId)
		return goja.Undefined()
	})

	vm.Set("volumeDown", func(call goja.FunctionCall) goja.Value {
		displayId := 0
		if len(call.Arguments) > 0 {
			displayId = int(call.Argument(0).ToInteger())
		}
		motion.VolumeDown(displayId)
		return goja.Undefined()
	})

	vm.Set("camera", func(call goja.FunctionCall) goja.Value {
		motion.Camera()
		return goja.Undefined()
	})

	vm.Set("keyAction", func(call goja.FunctionCall) goja.Value {
		code := int(call.Argument(0).ToInteger())
		displayId := 0
		if len(call.Arguments) > 1 {
			displayId = int(call.Argument(1).ToInteger())
		}
		motion.KeyAction(code, displayId)
		return goja.Undefined()
	})

	motionObj := vm.NewObject()
	motionObj.Set("touchDown", func(call goja.FunctionCall) goja.Value {
		x := int(call.Argument(0).ToInteger())
		y := int(call.Argument(1).ToInteger())
		fingerID := 0
		if len(call.Arguments) > 2 {
			fingerID = int(call.Argument(2).ToInteger())
		}
		displayId := 0
		if len(call.Arguments) > 3 {
			displayId = int(call.Argument(3).ToInteger())
		}
		motion.TouchDown(x, y, fingerID, displayId)
		return goja.Undefined()
	})

	motionObj.Set("touchMove", func(call goja.FunctionCall) goja.Value {
		x := int(call.Argument(0).ToInteger())
		y := int(call.Argument(1).ToInteger())
		fingerID := 0
		if len(call.Arguments) > 2 {
			fingerID = int(call.Argument(2).ToInteger())
		}
		displayId := 0
		if len(call.Arguments) > 3 {
			displayId = int(call.Argument(3).ToInteger())
		}
		motion.TouchMove(x, y, fingerID, displayId)
		return goja.Undefined()
	})

	motionObj.Set("touchUp", func(call goja.FunctionCall) goja.Value {
		x := int(call.Argument(0).ToInteger())
		y := int(call.Argument(1).ToInteger())
		fingerID := 0
		if len(call.Arguments) > 2 {
			fingerID = int(call.Argument(2).ToInteger())
		}
		displayId := 0
		if len(call.Arguments) > 3 {
			displayId = int(call.Argument(3).ToInteger())
		}
		motion.TouchUp(x, y, fingerID, displayId)
		return goja.Undefined()
	})

	motionObj.Set("click", func(call goja.FunctionCall) goja.Value {
		x := int(call.Argument(0).ToInteger())
		y := int(call.Argument(1).ToInteger())
		fingerID := 0
		if len(call.Arguments) > 2 {
			fingerID = int(call.Argument(2).ToInteger())
		}
		displayId := 0
		if len(call.Arguments) > 3 {
			displayId = int(call.Argument(3).ToInteger())
		}
		motion.Click(x, y, fingerID, displayId)
		return goja.Undefined()
	})

	motionObj.Set("longClick", func(call goja.FunctionCall) goja.Value {
		x := int(call.Argument(0).ToInteger())
		y := int(call.Argument(1).ToInteger())
		duration := 500
		if len(call.Arguments) > 2 {
			duration = int(call.Argument(2).ToInteger())
		}
		fingerID := 0
		if len(call.Arguments) > 3 {
			fingerID = int(call.Argument(3).ToInteger())
		}
		displayId := 0
		if len(call.Arguments) > 4 {
			displayId = int(call.Argument(4).ToInteger())
		}
		motion.LongClick(x, y, duration, fingerID, displayId)
		return goja.Undefined()
	})

	motionObj.Set("swipe", func(call goja.FunctionCall) goja.Value {
		x1 := int(call.Argument(0).ToInteger())
		y1 := int(call.Argument(1).ToInteger())
		x2 := int(call.Argument(2).ToInteger())
		y2 := int(call.Argument(3).ToInteger())
		duration := int(call.Argument(4).ToInteger())
		fingerID := 0
		if len(call.Arguments) > 5 {
			fingerID = int(call.Argument(5).ToInteger())
		}
		displayId := 0
		if len(call.Arguments) > 6 {
			displayId = int(call.Argument(6).ToInteger())
		}
		motion.Swipe(x1, y1, x2, y2, duration, fingerID, displayId)
		return goja.Undefined()
	})

	motionObj.Set("swipe2", func(call goja.FunctionCall) goja.Value {
		x1 := int(call.Argument(0).ToInteger())
		y1 := int(call.Argument(1).ToInteger())
		x2 := int(call.Argument(2).ToInteger())
		y2 := int(call.Argument(3).ToInteger())
		duration := int(call.Argument(4).ToInteger())
		fingerID := 0
		if len(call.Arguments) > 5 {
			fingerID = int(call.Argument(5).ToInteger())
		}
		displayId := 0
		if len(call.Arguments) > 6 {
			displayId = int(call.Argument(6).ToInteger())
		}
		motion.Swipe2(x1, y1, x2, y2, duration, fingerID, displayId)
		return goja.Undefined()
	})

	motionObj.Set("home", func(call goja.FunctionCall) goja.Value {
		displayId := 0
		if len(call.Arguments) > 0 {
			displayId = int(call.Argument(0).ToInteger())
		}
		motion.Home(displayId)
		return goja.Undefined()
	})

	motionObj.Set("back", func(call goja.FunctionCall) goja.Value {
		displayId := 0
		if len(call.Arguments) > 0 {
			displayId = int(call.Argument(0).ToInteger())
		}
		motion.Back(displayId)
		return goja.Undefined()
	})

	motionObj.Set("recents", func(call goja.FunctionCall) goja.Value {
		displayId := 0
		if len(call.Arguments) > 0 {
			displayId = int(call.Argument(0).ToInteger())
		}
		motion.Recents(displayId)
		return goja.Undefined()
	})

	motionObj.Set("powerDialog", func(call goja.FunctionCall) goja.Value {
		motion.PowerDialog()
		return goja.Undefined()
	})

	motionObj.Set("notifications", func(call goja.FunctionCall) goja.Value {
		motion.Notifications()
		return goja.Undefined()
	})

	motionObj.Set("quickSettings", func(call goja.FunctionCall) goja.Value {
		motion.QuickSettings()
		return goja.Undefined()
	})

	motionObj.Set("volumeUp", func(call goja.FunctionCall) goja.Value {
		displayId := 0
		if len(call.Arguments) > 0 {
			displayId = int(call.Argument(0).ToInteger())
		}
		motion.VolumeUp(displayId)
		return goja.Undefined()
	})

	motionObj.Set("volumeDown", func(call goja.FunctionCall) goja.Value {
		displayId := 0
		if len(call.Arguments) > 0 {
			displayId = int(call.Argument(0).ToInteger())
		}
		motion.VolumeDown(displayId)
		return goja.Undefined()
	})

	motionObj.Set("camera", func(call goja.FunctionCall) goja.Value {
		motion.Camera()
		return goja.Undefined()
	})

	motionObj.Set("keyAction", func(call goja.FunctionCall) goja.Value {
		code := int(call.Argument(0).ToInteger())
		displayId := 0
		if len(call.Arguments) > 1 {
			displayId = int(call.Argument(1).ToInteger())
		}
		motion.KeyAction(code, displayId)
		return goja.Undefined()
	})

	vm.Set("motion", motionObj)

	engine.RegisterMethod("touchDown", "按下屏幕", func(x, y, fingerID, displayId int) { motion.TouchDown(x, y, fingerID, displayId) }, true)
	engine.RegisterMethod("touchMove", "移动手指", func(x, y, fingerID, displayId int) { motion.TouchMove(x, y, fingerID, displayId) }, true)
	engine.RegisterMethod("touchUp", "抬起手指", func(x, y, fingerID, displayId int) { motion.TouchUp(x, y, fingerID, displayId) }, true)
	engine.RegisterMethod("click", "点击", func(x, y, fingerID, displayId int) { motion.Click(x, y, fingerID, displayId) }, true)
	engine.RegisterMethod("longClick", "长按", func(x, y, duration, fingerID, displayId int) { motion.LongClick(x, y, duration, fingerID, displayId) }, true)
	engine.RegisterMethod("swipe", "滑动", func(x1, y1, x2, y2, duration, fingerID, displayId int) {
		motion.Swipe(x1, y1, x2, y2, duration, fingerID, displayId)
	}, true)
	engine.RegisterMethod("swipe2", "滑动(两点)", func(x1, y1, x2, y2, duration, fingerID, displayId int) {
		motion.Swipe2(x1, y1, x2, y2, duration, fingerID, displayId)
	}, true)
	engine.RegisterMethod("home", "按下Home键", func(displayId int) { motion.Home(displayId) }, true)
	engine.RegisterMethod("back", "按下返回键", func(displayId int) { motion.Back(displayId) }, true)
	engine.RegisterMethod("recents", "按下最近任务键", func(displayId int) { motion.Recents(displayId) }, true)
	engine.RegisterMethod("powerDialog", "长按电源键", motion.PowerDialog, true)
	engine.RegisterMethod("notifications", "下拉通知栏", motion.Notifications, true)
	engine.RegisterMethod("quickSettings", "下拉快捷设置", motion.QuickSettings, true)
	engine.RegisterMethod("volumeUp", "按下音量加键", func(displayId int) { motion.VolumeUp(displayId) }, true)
	engine.RegisterMethod("volumeDown", "按下音量减键", func(displayId int) { motion.VolumeDown(displayId) }, true)
	engine.RegisterMethod("camera", "按下相机键", motion.Camera, true)
	engine.RegisterMethod("keyAction", "按键动作", func(code, displayId int) { motion.KeyAction(code, displayId) }, true)

	return nil
}
