//go:build ignore
// +build ignore

package virtualscreen

import (
	"fmt"

	"github.com/Dasongzi1366/AutoGo/motion"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"

	lua "github.com/yuin/gopher-lua"
)

// VirtualScreenModule virtualscreen 模块（懒人精灵兼容）
type VirtualScreenModule struct {
	virtualDisplays map[int]interface{}
	nextDisplayId   int
}

// Name 返回模块名称
func (m *VirtualScreenModule) Name() string {
	return "virtualscreen"
}

// IsAvailable 检查模块是否可用
func (m *VirtualScreenModule) IsAvailable() bool {
	return true
}

// Register 向引擎注册方法
func (m *VirtualScreenModule) Register(engine model.Engine) error {
	state := engine.GetState()

	// 初始化虚拟显示存储
	m.virtualDisplays = make(map[int]interface{})
	m.nextDisplayId = 1

	// 创建 virtualDisplay 表
	virtualDisplayTable := state.NewTable()

	// 注册 virtualDisplay.createVirtualDisplay - 创建虚拟屏
	virtualDisplayTable.RawSetString("createVirtualDisplay", state.NewFunction(func(L *lua.LState) int {
		width := L.CheckInt(1)
		height := L.CheckInt(2)
		dpi := 320
		if L.GetTop() >= 3 {
			dpi = L.CheckInt(3)
		}

		displayId := m.createVirtualDisplay(L, width, height, dpi)
		L.Push(lua.LNumber(displayId))
		return 1
	}))

	// 注册 virtualDisplay.showVirtualDisplay - 显示虚拟屏
	virtualDisplayTable.RawSetString("showVirtualDisplay", state.NewFunction(func(L *lua.LState) int {
		displayId := L.CheckInt(1)
		config := L.CheckAny(2)

		result := m.showVirtualDisplay(L, displayId, config)
		L.Push(lua.LBool(result))
		return 1
	}))

	// 注册 virtualDisplay.closeVirtualDisplay - 关闭虚拟屏
	virtualDisplayTable.RawSetString("closeVirtualDisplay", state.NewFunction(func(L *lua.LState) int {
		displayId := L.CheckInt(1)
		m.closeVirtualDisplay(L, displayId)
		return 0
	}))

	// 注册 virtualDisplay.snapToBitmap - 虚拟屏截图
	virtualDisplayTable.RawSetString("snapToBitmap", state.NewFunction(func(L *lua.LState) int {
		displayId := L.CheckInt(1)
		x := L.CheckInt(2)
		y := L.CheckInt(3)
		w := L.CheckInt(4)
		h := L.CheckInt(5)

		bitmap := m.snapToBitmap(L, displayId, x, y, w, h)
		L.Push(bitmap)
		return 1
	}))

	// 注册 virtualDisplay.switchToVirtualDisplay - 切换虚拟屏幕
	virtualDisplayTable.RawSetString("switchToVirtualDisplay", state.NewFunction(func(L *lua.LState) int {
		displayId := L.CheckInt(1)
		result := m.switchToVirtualDisplay(L, displayId)
		L.Push(lua.LBool(result))
		return 1
	}))

	// 注册 virtualDisplay.runAppWithVirtualDisplay - 在虚拟屏里面运行指定应用
	virtualDisplayTable.RawSetString("runAppWithVirtualDisplay", state.NewFunction(func(L *lua.LState) int {
		displayId := L.CheckInt(1)
		pkg := L.CheckString(2)
		forceStopApp := true
		if L.GetTop() >= 3 {
			forceStopApp = L.CheckBool(3)
		}

		result := m.runAppWithVirtualDisplay(L, displayId, pkg, forceStopApp)
		L.Push(lua.LBool(result))
		return 1
	}))

	// 注册 virtualDisplay.tap - 虚拟屏里面点击
	virtualDisplayTable.RawSetString("tap", state.NewFunction(func(L *lua.LState) int {
		displayId := L.CheckInt(1)
		x := L.CheckInt(2)
		y := L.CheckInt(3)

		result := m.tap(L, displayId, x, y)
		L.Push(lua.LBool(result))
		return 1
	}))

	// 注册 virtualDisplay.touchDown - 虚拟屏里面模拟手指按下
	virtualDisplayTable.RawSetString("touchDown", state.NewFunction(func(L *lua.LState) int {
		displayId := L.CheckInt(1)
		id := L.CheckInt(2)
		x := L.CheckInt(3)
		y := L.CheckInt(4)

		result := m.touchDown(L, displayId, id, x, y)
		L.Push(lua.LBool(result))
		return 1
	}))

	// 注册 virtualDisplay.touchUp - 虚拟屏里面模拟手指弹起
	virtualDisplayTable.RawSetString("touchUp", state.NewFunction(func(L *lua.LState) int {
		displayId := L.CheckInt(1)
		id := L.CheckInt(2)

		result := m.touchUp(L, displayId, id)
		L.Push(lua.LBool(result))
		return 1
	}))

	// 注册 virtualDisplay.touchMove - 虚拟屏里面模拟手指滑动
	virtualDisplayTable.RawSetString("touchMove", state.NewFunction(func(L *lua.LState) int {
		displayId := L.CheckInt(1)
		id := L.CheckInt(2)
		x := L.CheckInt(3)
		y := L.CheckInt(4)

		result := m.touchMove(L, displayId, id, x, y)
		L.Push(lua.LBool(result))
		return 1
	}))

	// 注册 virtualDisplay.touchMoveEx - 虚拟屏里面模拟手指增强滑动
	virtualDisplayTable.RawSetString("touchMoveEx", state.NewFunction(func(L *lua.LState) int {
		displayId := L.CheckInt(1)
		id := L.CheckInt(2)
		x := L.CheckInt(3)
		y := L.CheckInt(4)
		time := L.CheckInt(5)

		result := m.touchMoveEx(L, displayId, id, x, y, time)
		L.Push(lua.LBool(result))
		return 1
	}))

	// 注册 virtualDisplay.swipe - 虚拟屏里面模拟滑动
	virtualDisplayTable.RawSetString("swipe", state.NewFunction(func(L *lua.LState) int {
		displayId := L.CheckInt(1)
		x1 := L.CheckInt(2)
		y1 := L.CheckInt(3)
		x2 := L.CheckInt(4)
		y2 := L.CheckInt(5)
		time := 500
		if L.GetTop() >= 6 {
			time = L.CheckInt(6)
		}

		result := m.swipe(L, displayId, x1, y1, x2, y2, time)
		L.Push(lua.LBool(result))
		return 1
	}))

	// 注册 virtualDisplay 到全局
	state.SetGlobal("virtualDisplay", virtualDisplayTable)

	return nil
}

// createVirtualDisplay 创建虚拟屏
func (m *VirtualScreenModule) createVirtualDisplay(L *lua.LState, width, height, dpi int) int {
	displayId := m.nextDisplayId
	m.nextDisplayId++
	m.virtualDisplays[displayId] = map[string]interface{}{
		"width":  width,
		"height": height,
		"dpi":    dpi,
	}
	fmt.Printf("[virtualDisplay] 创建虚拟屏: displayId=%d, width=%d, height=%d, dpi=%d\n", displayId, width, height, dpi)
	return displayId
}

// showVirtualDisplay 显示虚拟屏
func (m *VirtualScreenModule) showVirtualDisplay(L *lua.LState, displayId int, config lua.LValue) bool {
	fmt.Printf("[virtualDisplay] 显示虚拟屏: displayId=%d\n", displayId)
	return true
}

// closeVirtualDisplay 关闭虚拟屏
func (m *VirtualScreenModule) closeVirtualDisplay(L *lua.LState, displayId int) {
	delete(m.virtualDisplays, displayId)
	fmt.Printf("[virtualDisplay] 关闭虚拟屏: displayId=%d\n", displayId)
}

// snapToBitmap 虚拟屏截图
func (m *VirtualScreenModule) snapToBitmap(L *lua.LState, displayId, x, y, w, h int) lua.LValue {
	fmt.Printf("[virtualDisplay] 虚拟屏截图: displayId=%d, x=%d, y=%d, w=%d, h=%d\n", displayId, x, y, w, h)
	return lua.LNil
}

// switchToVirtualDisplay 切换虚拟屏幕
func (m *VirtualScreenModule) switchToVirtualDisplay(L *lua.LState, displayId int) bool {
	fmt.Printf("[virtualDisplay] 切换虚拟屏幕: displayId=%d\n", displayId)
	return true
}

// runAppWithVirtualDisplay 在虚拟屏里面运行指定应用
func (m *VirtualScreenModule) runAppWithVirtualDisplay(L *lua.LState, displayId int, pkg string, forceStopApp bool) bool {
	fmt.Printf("[virtualDisplay] 在虚拟屏运行应用: displayId=%d, pkg=%s, forceStopApp=%v\n", displayId, pkg, forceStopApp)
	return true
}

// tap 虚拟屏里面点击 - 使用 AutoGo 的 motion.Click
func (m *VirtualScreenModule) tap(L *lua.LState, displayId, x, y int) bool {
	fmt.Printf("[virtualDisplay] 虚拟屏点击: displayId=%d, x=%d, y=%d\n", displayId, x, y)
	motion.Click(x, y, 0, displayId)
	return true
}

// touchDown 虚拟屏里面模拟手指按下 - 使用 AutoGo 的 motion.TouchDown
func (m *VirtualScreenModule) touchDown(L *lua.LState, displayId, id, x, y int) bool {
	fmt.Printf("[virtualDisplay] 手指按下: displayId=%d, id=%d, x=%d, y=%d\n", displayId, id, x, y)
	motion.TouchDown(x, y, id, displayId)
	return true
}

// touchUp 虚拟屏里面模拟手指弹起 - 使用 AutoGo 的 motion.TouchUp
func (m *VirtualScreenModule) touchUp(L *lua.LState, displayId, id int) bool {
	fmt.Printf("[virtualDisplay] 手指弹起: displayId=%d, id=%d\n", displayId, id)
	motion.TouchUp(0, 0, id, displayId)
	return true
}

// touchMove 虚拟屏里面模拟手指滑动 - 使用 AutoGo 的 motion.TouchMove
func (m *VirtualScreenModule) touchMove(L *lua.LState, displayId, id, x, y int) bool {
	fmt.Printf("[virtualDisplay] 手指滑动: displayId=%d, id=%d, x=%d, y=%d\n", displayId, id, x, y)
	motion.TouchMove(x, y, id, displayId)
	return true
}

// touchMoveEx 虚拟屏里面模拟手指增强滑动 - 使用 AutoGo 的 motion.Swipe
func (m *VirtualScreenModule) touchMoveEx(L *lua.LState, displayId, id, x, y, time int) bool {
	fmt.Printf("[virtualDisplay] 手指增强滑动: displayId=%d, id=%d, x=%d, y=%d, time=%d\n", displayId, id, x, y, time)
	motion.Swipe(0, 0, x, y, time, id, displayId)
	return true
}

// swipe 虚拟屏里面模拟滑动 - 使用 AutoGo 的 motion.Swipe
func (m *VirtualScreenModule) swipe(L *lua.LState, displayId, x1, y1, x2, y2, time int) bool {
	fmt.Printf("[virtualDisplay] 滑动: displayId=%d, from=(%d,%d) to=(%d,%d), time=%d\n", displayId, x1, y1, x2, y2, time)
	motion.Swipe(x1, y1, x2, y2, time, 0, displayId)
	return true
}

// GetModule 获取模块实例
func GetModule() model.Module {
	return &VirtualScreenModule{}
}
