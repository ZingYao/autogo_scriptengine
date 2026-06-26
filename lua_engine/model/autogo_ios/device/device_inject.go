package device

import (
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"

	autogodevice "github.com/Dasongzi1366/AutoGo/device"
	lua "github.com/yuin/gopher-lua"
)

// DeviceModule iOS device 模块。
type DeviceModule struct{}

// Name 返回模块名称。
func (m *DeviceModule) Name() string {
	return "device"
}

// IsAvailable 检查模块是否可用。
func (m *DeviceModule) IsAvailable() bool {
	return true
}

// Register 向引擎注册 iOS device 方法。
func (m *DeviceModule) Register(engine model.Engine) error {
	state := engine.GetState()
	width, height, scale, rotation := autogodevice.GetDisplayInfo()
	deviceObj := state.NewTable()
	state.SetGlobal("device", deviceObj)

	deviceObj.RawSetString("width", lua.LNumber(width))
	deviceObj.RawSetString("height", lua.LNumber(height))
	deviceObj.RawSetString("scale", lua.LNumber(scale))
	deviceObj.RawSetString("rotation", lua.LNumber(rotation))
	deviceObj.RawSetString("model", lua.LString(autogodevice.Model))
	deviceObj.RawSetString("release", lua.LString(autogodevice.Release))
	deviceObj.RawSetString("serial", lua.LString(autogodevice.Serial))

	deviceObj.RawSetString("getDisplayInfo", state.NewFunction(func(L *lua.LState) int {
		width, height, scale, rotation := autogodevice.GetDisplayInfo()
		result := L.NewTable()
		result.RawSetString("width", lua.LNumber(width))
		result.RawSetString("height", lua.LNumber(height))
		result.RawSetString("scale", lua.LNumber(scale))
		result.RawSetString("rotation", lua.LNumber(rotation))
		L.Push(result)
		return 1
	}))

	deviceObj.RawSetString("getBattery", state.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LNumber(autogodevice.GetBattery()))
		return 1
	}))
	deviceObj.RawSetString("getBatteryStatus", state.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LNumber(autogodevice.GetBatteryStatus()))
		return 1
	}))
	deviceObj.RawSetString("isScreenOn", state.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LBool(autogodevice.IsScreenOn()))
		return 1
	}))
	deviceObj.RawSetString("isScreenUnlock", state.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LBool(autogodevice.IsScreenUnlock()))
		return 1
	}))
	deviceObj.RawSetString("getBrightness", state.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LNumber(autogodevice.GetBrightness()))
		return 1
	}))
	deviceObj.RawSetString("getTotalMem", state.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LNumber(autogodevice.GetTotalMem()))
		return 1
	}))
	deviceObj.RawSetString("getAvailMem", state.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LNumber(autogodevice.GetAvailMem()))
		return 1
	}))
	deviceObj.RawSetString("wakeUp", state.NewFunction(func(L *lua.LState) int {
		autogodevice.WakeUp()
		return 0
	}))
	deviceObj.RawSetString("keepScreenOn", state.NewFunction(func(L *lua.LState) int {
		autogodevice.KeepScreenOn()
		return 0
	}))
	deviceObj.RawSetString("getIp", state.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LString(autogodevice.GetIp()))
		return 1
	}))
	deviceObj.RawSetString("getWifiMac", state.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LString(autogodevice.GetWifiMac()))
		return 1
	}))
	deviceObj.RawSetString("reboot", state.NewFunction(func(L *lua.LState) int {
		autogodevice.Reboot()
		return 0
	}))

	engine.RegisterMethod("device.width", "设备分辨率宽度", func() int { width, _, _, _ := autogodevice.GetDisplayInfo(); return width }, true)
	engine.RegisterMethod("device.height", "设备分辨率高度", func() int { _, height, _, _ := autogodevice.GetDisplayInfo(); return height }, true)
	engine.RegisterMethod("device.scale", "屏幕缩放比例", func() float64 { _, _, scale, _ := autogodevice.GetDisplayInfo(); return scale }, true)
	engine.RegisterMethod("device.rotation", "屏幕旋转角度", func() int { _, _, _, rotation := autogodevice.GetDisplayInfo(); return rotation }, true)
	engine.RegisterMethod("device.model", "设备机型标识", func() string { return autogodevice.Model }, true)
	engine.RegisterMethod("device.release", "系统版本", func() string { return autogodevice.Release }, true)
	engine.RegisterMethod("device.serial", "设备序列号", func() string { return autogodevice.Serial }, true)
	engine.RegisterMethod("device.getDisplayInfo", "获取设备分辨率信息", autogodevice.GetDisplayInfo, true)
	engine.RegisterMethod("device.getBattery", "返回当前电量百分比", autogodevice.GetBattery, true)
	engine.RegisterMethod("device.getBatteryStatus", "返回电池状态", autogodevice.GetBatteryStatus, true)
	engine.RegisterMethod("device.isScreenOn", "返回设备屏幕是否亮着", autogodevice.IsScreenOn, true)
	engine.RegisterMethod("device.isScreenUnlock", "返回屏幕锁是否已经解开", autogodevice.IsScreenUnlock, true)
	engine.RegisterMethod("device.getBrightness", "返回屏幕亮度", autogodevice.GetBrightness, true)
	engine.RegisterMethod("device.getTotalMem", "返回设备内存总量", autogodevice.GetTotalMem, true)
	engine.RegisterMethod("device.getAvailMem", "返回设备当前可用内存", autogodevice.GetAvailMem, true)
	engine.RegisterMethod("device.wakeUp", "唤醒设备", autogodevice.WakeUp, true)
	engine.RegisterMethod("device.keepScreenOn", "保持屏幕常亮", autogodevice.KeepScreenOn, true)
	engine.RegisterMethod("device.getIp", "获取设备局域网 IP 地址", autogodevice.GetIp, true)
	engine.RegisterMethod("device.getWifiMac", "获取设备 WIFI MAC 地址", autogodevice.GetWifiMac, true)
	engine.RegisterMethod("device.reboot", "重启设备", autogodevice.Reboot, true)

	return nil
}
