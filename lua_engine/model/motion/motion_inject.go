package motion

import (
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"

	"github.com/Dasongzi1366/AutoGo/motion"
	lua "github.com/yuin/gopher-lua"
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
	state := engine.GetState()

	motionObj := state.NewTable()
	state.SetGlobal("motion", motionObj)

	state.SetGlobal("touchDown", state.NewFunction(func(L *lua.LState) int {
		x := L.CheckInt(1)
		y := L.CheckInt(2)
		fingerID := 0
		if L.GetTop() > 2 {
			fingerID = L.CheckInt(3)
		}
		displayId := 0
		if L.GetTop() > 3 {
			displayId = L.CheckInt(4)
		}
		motion.TouchDown(x, y, fingerID, displayId)
		return 0
	}))

	state.SetGlobal("touchMove", state.NewFunction(func(L *lua.LState) int {
		x := L.CheckInt(1)
		y := L.CheckInt(2)
		fingerID := 0
		if L.GetTop() > 2 {
			fingerID = L.CheckInt(3)
		}
		displayId := 0
		if L.GetTop() > 3 {
			displayId = L.CheckInt(4)
		}
		motion.TouchMove(x, y, fingerID, displayId)
		return 0
	}))

	state.SetGlobal("touchUp", state.NewFunction(func(L *lua.LState) int {
		x := L.CheckInt(1)
		y := L.CheckInt(2)
		fingerID := 0
		if L.GetTop() > 2 {
			fingerID = L.CheckInt(3)
		}
		displayId := 0
		if L.GetTop() > 3 {
			displayId = L.CheckInt(4)
		}
		motion.TouchUp(x, y, fingerID, displayId)
		return 0
	}))

	state.SetGlobal("click", state.NewFunction(func(L *lua.LState) int {
		x := L.CheckInt(1)
		y := L.CheckInt(2)
		fingerID := 0
		if L.GetTop() > 2 {
			fingerID = L.CheckInt(3)
		}
		displayId := 0
		if L.GetTop() > 3 {
			displayId = L.CheckInt(4)
		}
		motion.Click(x, y, fingerID, displayId)
		return 0
	}))

	state.SetGlobal("longClick", state.NewFunction(func(L *lua.LState) int {
		x := L.CheckInt(1)
		y := L.CheckInt(2)
		duration := 500
		if L.GetTop() > 2 {
			duration = L.CheckInt(3)
		}
		fingerID := 0
		if L.GetTop() > 3 {
			fingerID = L.CheckInt(4)
		}
		displayId := 0
		if L.GetTop() > 4 {
			displayId = L.CheckInt(5)
		}
		motion.LongClick(x, y, duration, fingerID, displayId)
		return 0
	}))

	state.SetGlobal("swipe", state.NewFunction(func(L *lua.LState) int {
		x1 := L.CheckInt(1)
		y1 := L.CheckInt(2)
		x2 := L.CheckInt(3)
		y2 := L.CheckInt(4)
		duration := L.CheckInt(5)
		fingerID := 0
		if L.GetTop() > 5 {
			fingerID = L.CheckInt(6)
		}
		displayId := 0
		if L.GetTop() > 6 {
			displayId = L.CheckInt(7)
		}
		motion.Swipe(x1, y1, x2, y2, duration, fingerID, displayId)
		return 0
	}))

	state.SetGlobal("swipe2", state.NewFunction(func(L *lua.LState) int {
		x1 := L.CheckInt(1)
		y1 := L.CheckInt(2)
		x2 := L.CheckInt(3)
		y2 := L.CheckInt(4)
		duration := L.CheckInt(5)
		fingerID := 0
		if L.GetTop() > 5 {
			fingerID = L.CheckInt(6)
		}
		displayId := 0
		if L.GetTop() > 6 {
			displayId = L.CheckInt(7)
		}
		motion.Swipe2(x1, y1, x2, y2, duration, fingerID, displayId)
		return 0
	}))

	state.SetGlobal("home", state.NewFunction(func(L *lua.LState) int {
		displayId := 0
		if L.GetTop() > 0 {
			displayId = L.CheckInt(1)
		}
		motion.Home(displayId)
		return 0
	}))

	state.SetGlobal("back", state.NewFunction(func(L *lua.LState) int {
		displayId := 0
		if L.GetTop() > 0 {
			displayId = L.CheckInt(1)
		}
		motion.Back(displayId)
		return 0
	}))

	state.SetGlobal("recents", state.NewFunction(func(L *lua.LState) int {
		displayId := 0
		if L.GetTop() > 0 {
			displayId = L.CheckInt(1)
		}
		motion.Recents(displayId)
		return 0
	}))

	state.SetGlobal("powerDialog", state.NewFunction(func(L *lua.LState) int {
		motion.PowerDialog()
		return 0
	}))

	state.SetGlobal("notifications", state.NewFunction(func(L *lua.LState) int {
		motion.Notifications()
		return 0
	}))

	state.SetGlobal("quickSettings", state.NewFunction(func(L *lua.LState) int {
		motion.QuickSettings()
		return 0
	}))

	state.SetGlobal("volumeUp", state.NewFunction(func(L *lua.LState) int {
		displayId := 0
		if L.GetTop() > 0 {
			displayId = L.CheckInt(1)
		}
		motion.VolumeUp(displayId)
		return 0
	}))

	state.SetGlobal("volumeDown", state.NewFunction(func(L *lua.LState) int {
		displayId := 0
		if L.GetTop() > 0 {
			displayId = L.CheckInt(1)
		}
		motion.VolumeDown(displayId)
		return 0
	}))

	state.SetGlobal("camera", state.NewFunction(func(L *lua.LState) int {
		motion.Camera()
		return 0
	}))

	state.SetGlobal("keyAction", state.NewFunction(func(L *lua.LState) int {
		code := L.CheckInt(1)
		displayId := 0
		if L.GetTop() > 1 {
			displayId = L.CheckInt(2)
		}
		motion.KeyAction(code, displayId)
		return 0
	}))

	motionObj.RawSetString("touchDown", state.NewFunction(func(L *lua.LState) int {
		x := L.CheckInt(1)
		y := L.CheckInt(2)
		fingerID := 0
		if L.GetTop() > 2 {
			fingerID = L.CheckInt(3)
		}
		displayId := 0
		if L.GetTop() > 3 {
			displayId = L.CheckInt(4)
		}
		motion.TouchDown(x, y, fingerID, displayId)
		return 0
	}))

	motionObj.RawSetString("touchMove", state.NewFunction(func(L *lua.LState) int {
		x := L.CheckInt(1)
		y := L.CheckInt(2)
		fingerID := 0
		if L.GetTop() > 2 {
			fingerID = L.CheckInt(3)
		}
		displayId := 0
		if L.GetTop() > 3 {
			displayId = L.CheckInt(4)
		}
		motion.TouchMove(x, y, fingerID, displayId)
		return 0
	}))

	motionObj.RawSetString("touchUp", state.NewFunction(func(L *lua.LState) int {
		x := L.CheckInt(1)
		y := L.CheckInt(2)
		fingerID := 0
		if L.GetTop() > 2 {
			fingerID = L.CheckInt(3)
		}
		displayId := 0
		if L.GetTop() > 3 {
			displayId = L.CheckInt(4)
		}
		motion.TouchUp(x, y, fingerID, displayId)
		return 0
	}))

	motionObj.RawSetString("click", state.NewFunction(func(L *lua.LState) int {
		x := L.CheckInt(1)
		y := L.CheckInt(2)
		fingerID := 0
		if L.GetTop() > 2 {
			fingerID = L.CheckInt(3)
		}
		displayId := 0
		if L.GetTop() > 3 {
			displayId = L.CheckInt(4)
		}
		motion.Click(x, y, fingerID, displayId)
		return 0
	}))

	motionObj.RawSetString("longClick", state.NewFunction(func(L *lua.LState) int {
		x := L.CheckInt(1)
		y := L.CheckInt(2)
		duration := 500
		if L.GetTop() > 2 {
			duration = L.CheckInt(3)
		}
		fingerID := 0
		if L.GetTop() > 3 {
			fingerID = L.CheckInt(4)
		}
		displayId := 0
		if L.GetTop() > 4 {
			displayId = L.CheckInt(5)
		}
		motion.LongClick(x, y, duration, fingerID, displayId)
		return 0
	}))

	motionObj.RawSetString("swipe", state.NewFunction(func(L *lua.LState) int {
		x1 := L.CheckInt(1)
		y1 := L.CheckInt(2)
		x2 := L.CheckInt(3)
		y2 := L.CheckInt(4)
		duration := L.CheckInt(5)
		fingerID := 0
		if L.GetTop() > 5 {
			fingerID = L.CheckInt(6)
		}
		displayId := 0
		if L.GetTop() > 6 {
			displayId = L.CheckInt(7)
		}
		motion.Swipe(x1, y1, x2, y2, duration, fingerID, displayId)
		return 0
	}))

	motionObj.RawSetString("swipe2", state.NewFunction(func(L *lua.LState) int {
		x1 := L.CheckInt(1)
		y1 := L.CheckInt(2)
		x2 := L.CheckInt(3)
		y2 := L.CheckInt(4)
		duration := L.CheckInt(5)
		fingerID := 0
		if L.GetTop() > 5 {
			fingerID = L.CheckInt(6)
		}
		displayId := 0
		if L.GetTop() > 6 {
			displayId = L.CheckInt(7)
		}
		motion.Swipe2(x1, y1, x2, y2, duration, fingerID, displayId)
		return 0
	}))

	motionObj.RawSetString("home", state.NewFunction(func(L *lua.LState) int {
		displayId := 0
		if L.GetTop() > 0 {
			displayId = L.CheckInt(1)
		}
		motion.Home(displayId)
		return 0
	}))

	motionObj.RawSetString("back", state.NewFunction(func(L *lua.LState) int {
		displayId := 0
		if L.GetTop() > 0 {
			displayId = L.CheckInt(1)
		}
		motion.Back(displayId)
		return 0
	}))

	motionObj.RawSetString("recents", state.NewFunction(func(L *lua.LState) int {
		displayId := 0
		if L.GetTop() > 0 {
			displayId = L.CheckInt(1)
		}
		motion.Recents(displayId)
		return 0
	}))

	motionObj.RawSetString("powerDialog", state.NewFunction(func(L *lua.LState) int {
		motion.PowerDialog()
		return 0
	}))

	motionObj.RawSetString("notifications", state.NewFunction(func(L *lua.LState) int {
		motion.Notifications()
		return 0
	}))

	motionObj.RawSetString("quickSettings", state.NewFunction(func(L *lua.LState) int {
		motion.QuickSettings()
		return 0
	}))

	motionObj.RawSetString("volumeUp", state.NewFunction(func(L *lua.LState) int {
		displayId := 0
		if L.GetTop() > 0 {
			displayId = L.CheckInt(1)
		}
		motion.VolumeUp(displayId)
		return 0
	}))

	motionObj.RawSetString("volumeDown", state.NewFunction(func(L *lua.LState) int {
		displayId := 0
		if L.GetTop() > 0 {
			displayId = L.CheckInt(1)
		}
		motion.VolumeDown(displayId)
		return 0
	}))

	motionObj.RawSetString("camera", state.NewFunction(func(L *lua.LState) int {
		motion.Camera()
		return 0
	}))

	motionObj.RawSetString("keyAction", state.NewFunction(func(L *lua.LState) int {
		code := L.CheckInt(1)
		displayId := 0
		if L.GetTop() > 1 {
			displayId = L.CheckInt(2)
		}
		motion.KeyAction(code, displayId)
		return 0
	}))

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
