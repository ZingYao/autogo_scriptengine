package motion

import (
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"

	autogomotion "github.com/Dasongzi1366/AutoGo/motion"
)

// MotionModule 是 go-lua-vm 迁移后的模块壳。
type MotionModule struct{}

func New() *MotionModule { return &MotionModule{} }

func (m *MotionModule) Name() string { return "motion" }

func (m *MotionModule) IsAvailable() bool { return true }

func (m *MotionModule) Register(engine model.Engine) error {
	engine.RegisterMethod("motion.touchDown", "按下屏幕", func(x, y int, args ...int) {
		fingerID, displayID := fingerAndDisplay(args...)
		autogomotion.TouchDown(x, y, fingerID, displayID)
	}, true)
	engine.RegisterMethod("motion.touchMove", "移动手指", func(x, y int, args ...int) {
		fingerID, displayID := fingerAndDisplay(args...)
		autogomotion.TouchMove(x, y, fingerID, displayID)
	}, true)
	engine.RegisterMethod("motion.touchUp", "抬起手指", func(x, y int, args ...int) {
		fingerID, displayID := fingerAndDisplay(args...)
		autogomotion.TouchUp(x, y, fingerID, displayID)
	}, true)
	engine.RegisterMethod("motion.click", "点击坐标", func(x, y int, args ...int) {
		fingerID, displayID := fingerAndDisplay(args...)
		autogomotion.Click(x, y, fingerID, displayID)
	}, true)
	engine.RegisterMethod("motion.longClick", "长按坐标", func(x, y int, args ...int) {
		duration := 500
		if len(args) > 0 {
			duration = args[0]
		}
		fingerID := 0
		if len(args) > 1 {
			fingerID = args[1]
		}
		displayID := 0
		if len(args) > 2 {
			displayID = args[2]
		}
		autogomotion.LongClick(x, y, duration, fingerID, displayID)
	}, true)
	engine.RegisterMethod("motion.swipe", "滑动坐标", func(x1, y1, x2, y2 int, args ...int) {
		duration, fingerID, displayID := swipeArgs(args...)
		autogomotion.Swipe(x1, y1, x2, y2, duration, fingerID, displayID)
	}, true)
	engine.RegisterMethod("motion.swipe2", "贝塞尔曲线滑动坐标", func(x1, y1, x2, y2 int, args ...int) {
		duration, fingerID, displayID := swipeArgs(args...)
		autogomotion.Swipe2(x1, y1, x2, y2, duration, fingerID, displayID)
	}, true)
	engine.RegisterMethod("motion.home", "按下 Home 键", func(displayID ...int) {
		autogomotion.Home(optionalInt(0, displayID...))
	}, true)
	engine.RegisterMethod("motion.back", "按下返回键", func(displayID ...int) {
		autogomotion.Back(optionalInt(0, displayID...))
	}, true)
	engine.RegisterMethod("motion.recents", "显示最近任务", func(displayID ...int) {
		autogomotion.Recents(optionalInt(0, displayID...))
	}, true)
	engine.RegisterMethod("motion.powerDialog", "弹出电源键菜单", autogomotion.PowerDialog, true)
	engine.RegisterMethod("motion.notifications", "下拉通知栏", autogomotion.Notifications, true)
	engine.RegisterMethod("motion.quickSettings", "下拉快捷设置", autogomotion.QuickSettings, true)
	engine.RegisterMethod("motion.volumeUp", "按下音量加键", func(displayID ...int) {
		autogomotion.VolumeUp(optionalInt(0, displayID...))
	}, true)
	engine.RegisterMethod("motion.volumeDown", "按下音量减键", func(displayID ...int) {
		autogomotion.VolumeDown(optionalInt(0, displayID...))
	}, true)
	engine.RegisterMethod("motion.camera", "按下相机键", autogomotion.Camera, true)
	engine.RegisterMethod("motion.keyAction", "模拟按键", func(code int, displayID ...int) {
		autogomotion.KeyAction(code, optionalInt(0, displayID...))
	}, true)
	return nil
}

func GetModule() model.Module { return &MotionModule{} }

func fingerAndDisplay(args ...int) (int, int) {
	fingerID := 0
	if len(args) > 0 {
		fingerID = args[0]
	}
	displayID := 0
	if len(args) > 1 {
		displayID = args[1]
	}
	return fingerID, displayID
}

func swipeArgs(args ...int) (int, int, int) {
	duration := 500
	if len(args) > 0 {
		duration = args[0]
	}
	fingerID := 0
	if len(args) > 1 {
		fingerID = args[1]
	}
	displayID := 0
	if len(args) > 2 {
		displayID = args[2]
	}
	return duration, fingerID, displayID
}

func optionalInt(defaultValue int, values ...int) int {
	if len(values) == 0 {
		return defaultValue
	}
	return values[0]
}
