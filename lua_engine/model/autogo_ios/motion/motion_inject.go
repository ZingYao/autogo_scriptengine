package motion

import (
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"

	autogomotion "github.com/Dasongzi1366/AutoGo/motion"
	lua "github.com/yuin/gopher-lua"
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
	state := engine.GetState()
	motionObj := state.NewTable()
	state.SetGlobal("motion", motionObj)

	fingerID := func(L *lua.LState, index int) int {
		if L.GetTop() >= index {
			return L.CheckInt(index)
		}
		return 0
	}

	motionObj.RawSetString("touchDown", state.NewFunction(func(L *lua.LState) int {
		autogomotion.TouchDown(L.CheckInt(1), L.CheckInt(2), fingerID(L, 3))
		return 0
	}))
	motionObj.RawSetString("touchMove", state.NewFunction(func(L *lua.LState) int {
		autogomotion.TouchMove(L.CheckInt(1), L.CheckInt(2), fingerID(L, 3))
		return 0
	}))
	motionObj.RawSetString("touchUp", state.NewFunction(func(L *lua.LState) int {
		autogomotion.TouchUp(L.CheckInt(1), L.CheckInt(2), fingerID(L, 3))
		return 0
	}))
	motionObj.RawSetString("click", state.NewFunction(func(L *lua.LState) int {
		autogomotion.Click(L.CheckInt(1), L.CheckInt(2), fingerID(L, 3))
		return 0
	}))
	motionObj.RawSetString("longClick", state.NewFunction(func(L *lua.LState) int {
		duration := 500
		if L.GetTop() >= 3 {
			duration = L.CheckInt(3)
		}
		autogomotion.LongClick(L.CheckInt(1), L.CheckInt(2), duration, fingerID(L, 4))
		return 0
	}))
	motionObj.RawSetString("swipe", state.NewFunction(func(L *lua.LState) int {
		autogomotion.Swipe(L.CheckInt(1), L.CheckInt(2), L.CheckInt(3), L.CheckInt(4), L.CheckInt(5), fingerID(L, 6))
		return 0
	}))
	motionObj.RawSetString("swipe2", state.NewFunction(func(L *lua.LState) int {
		autogomotion.Swipe2(L.CheckInt(1), L.CheckInt(2), L.CheckInt(3), L.CheckInt(4), L.CheckInt(5), fingerID(L, 6))
		return 0
	}))
	motionObj.RawSetString("home", state.NewFunction(func(L *lua.LState) int { autogomotion.Home(); return 0 }))
	motionObj.RawSetString("recents", state.NewFunction(func(L *lua.LState) int { autogomotion.Recents(); return 0 }))
	motionObj.RawSetString("volumeUp", state.NewFunction(func(L *lua.LState) int { autogomotion.VolumeUp(); return 0 }))
	motionObj.RawSetString("volumeDown", state.NewFunction(func(L *lua.LState) int { autogomotion.VolumeDown(); return 0 }))
	motionObj.RawSetString("keyAction", state.NewFunction(func(L *lua.LState) int {
		autogomotion.KeyAction(L.CheckInt(1))
		return 0
	}))

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
