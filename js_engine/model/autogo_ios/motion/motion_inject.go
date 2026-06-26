package motion

import (
	"github.com/ZingYao/autogo_scriptengine/js_engine/model"

	autogomotion "github.com/Dasongzi1366/AutoGo/motion"
	"github.com/dop251/goja"
)

// MotionModule iOS motion 模块。
type MotionModule struct{}

// Name 返回模块名称。
func (m *MotionModule) Name() string {
	return "motion"
}

// IsAvailable 检查模块是否可用。
func (m *MotionModule) IsAvailable() bool {
	return true
}

// Register 向引擎注册 iOS motion 方法。
func (m *MotionModule) Register(engine model.Engine) error {
	vm := engine.GetVM()
	motionObj := vm.NewObject()
	vm.Set("motion", motionObj)

	fingerID := func(call goja.FunctionCall, index int) int {
		if len(call.Arguments) > index {
			return int(call.Argument(index).ToInteger())
		}
		return 0
	}

	motionObj.Set("touchDown", func(call goja.FunctionCall) goja.Value {
		autogomotion.TouchDown(int(call.Argument(0).ToInteger()), int(call.Argument(1).ToInteger()), fingerID(call, 2))
		return goja.Undefined()
	})
	motionObj.Set("touchMove", func(call goja.FunctionCall) goja.Value {
		autogomotion.TouchMove(int(call.Argument(0).ToInteger()), int(call.Argument(1).ToInteger()), fingerID(call, 2))
		return goja.Undefined()
	})
	motionObj.Set("touchUp", func(call goja.FunctionCall) goja.Value {
		autogomotion.TouchUp(int(call.Argument(0).ToInteger()), int(call.Argument(1).ToInteger()), fingerID(call, 2))
		return goja.Undefined()
	})
	motionObj.Set("click", func(call goja.FunctionCall) goja.Value {
		autogomotion.Click(int(call.Argument(0).ToInteger()), int(call.Argument(1).ToInteger()), fingerID(call, 2))
		return goja.Undefined()
	})
	motionObj.Set("longClick", func(call goja.FunctionCall) goja.Value {
		duration := 500
		if len(call.Arguments) > 2 {
			duration = int(call.Argument(2).ToInteger())
		}
		autogomotion.LongClick(int(call.Argument(0).ToInteger()), int(call.Argument(1).ToInteger()), duration, fingerID(call, 3))
		return goja.Undefined()
	})
	motionObj.Set("swipe", func(call goja.FunctionCall) goja.Value {
		autogomotion.Swipe(int(call.Argument(0).ToInteger()), int(call.Argument(1).ToInteger()), int(call.Argument(2).ToInteger()), int(call.Argument(3).ToInteger()), int(call.Argument(4).ToInteger()), fingerID(call, 5))
		return goja.Undefined()
	})
	motionObj.Set("swipe2", func(call goja.FunctionCall) goja.Value {
		autogomotion.Swipe2(int(call.Argument(0).ToInteger()), int(call.Argument(1).ToInteger()), int(call.Argument(2).ToInteger()), int(call.Argument(3).ToInteger()), int(call.Argument(4).ToInteger()), fingerID(call, 5))
		return goja.Undefined()
	})
	motionObj.Set("home", func(call goja.FunctionCall) goja.Value { autogomotion.Home(); return goja.Undefined() })
	motionObj.Set("recents", func(call goja.FunctionCall) goja.Value { autogomotion.Recents(); return goja.Undefined() })
	motionObj.Set("volumeUp", func(call goja.FunctionCall) goja.Value { autogomotion.VolumeUp(); return goja.Undefined() })
	motionObj.Set("volumeDown", func(call goja.FunctionCall) goja.Value { autogomotion.VolumeDown(); return goja.Undefined() })
	motionObj.Set("keyAction", func(call goja.FunctionCall) goja.Value {
		autogomotion.KeyAction(int(call.Argument(0).ToInteger()))
		return goja.Undefined()
	})

	engine.RegisterMethod("motion.touchDown", "按下屏幕", func(x, y, fingerID int) { autogomotion.TouchDown(x, y, fingerID) }, true)
	engine.RegisterMethod("motion.touchMove", "移动手指", func(x, y, fingerID int) { autogomotion.TouchMove(x, y, fingerID) }, true)
	engine.RegisterMethod("motion.touchUp", "抬起手指", func(x, y, fingerID int) { autogomotion.TouchUp(x, y, fingerID) }, true)
	engine.RegisterMethod("motion.click", "点击", func(x, y, fingerID int) { autogomotion.Click(x, y, fingerID) }, true)
	engine.RegisterMethod("motion.longClick", "长按", func(x, y, duration, fingerID int) { autogomotion.LongClick(x, y, duration, fingerID) }, true)
	engine.RegisterMethod("motion.swipe", "滑动", func(x1, y1, x2, y2, duration, fingerID int) {
		autogomotion.Swipe(x1, y1, x2, y2, duration, fingerID)
	}, true)
	engine.RegisterMethod("motion.swipe2", "贝塞尔曲线滑动", func(x1, y1, x2, y2, duration, fingerID int) {
		autogomotion.Swipe2(x1, y1, x2, y2, duration, fingerID)
	}, true)
	engine.RegisterMethod("motion.home", "按下 Home 键", autogomotion.Home, true)
	engine.RegisterMethod("motion.recents", "显示最近任务", autogomotion.Recents, true)
	engine.RegisterMethod("motion.volumeUp", "按下音量上键", autogomotion.VolumeUp, true)
	engine.RegisterMethod("motion.volumeDown", "按下音量下键", autogomotion.VolumeDown, true)
	engine.RegisterMethod("motion.keyAction", "模拟按键", func(code int) { autogomotion.KeyAction(code) }, true)

	return nil
}
