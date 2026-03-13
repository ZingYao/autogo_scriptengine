package device

import (
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"

	"github.com/Dasongzi1366/AutoGo/device"
	lua "github.com/yuin/gopher-lua"
)

// DeviceModule device 模块
type DeviceModule struct{}

// Name 返回模块名称
func (m *DeviceModule) Name() string {
	return "device"
}

// IsAvailable 检查模块是否可用
func (m *DeviceModule) IsAvailable() bool {
	return true
}

// Register 向引擎注册方法
func (m *DeviceModule) Register(engine model.Engine) error {
	state := engine.GetState()

	width, height, _, _ := device.GetDisplayInfo(0)
	deviceObj := state.NewTable()
	state.SetGlobal("device", deviceObj)

	deviceObj.RawSetString("width", lua.LNumber(width))
	deviceObj.RawSetString("height", lua.LNumber(height))
	deviceObj.RawSetString("sdkInt", lua.LNumber(device.SdkInt))
	deviceObj.RawSetString("cpuAbi", lua.LString(device.CpuAbi))
	deviceObj.RawSetString("buildId", lua.LString(device.BuildId))
	deviceObj.RawSetString("broad", lua.LString(device.Broad))
	deviceObj.RawSetString("brand", lua.LString(device.Brand))
	deviceObj.RawSetString("deviceName", lua.LString(device.Device))
	deviceObj.RawSetString("model", lua.LString(device.Model))
	deviceObj.RawSetString("product", lua.LString(device.Product))
	deviceObj.RawSetString("bootloader", lua.LString(device.Bootloader))
	deviceObj.RawSetString("hardware", lua.LString(device.Hardware))
	deviceObj.RawSetString("fingerprint", lua.LString(device.Fingerprint))
	deviceObj.RawSetString("serial", lua.LString(device.Serial))
	deviceObj.RawSetString("incremental", lua.LString(device.Incremental))
	deviceObj.RawSetString("release", lua.LString(device.Release))
	deviceObj.RawSetString("baseOS", lua.LString(device.BaseOS))
	deviceObj.RawSetString("securityPatch", lua.LString(device.SecurityPatch))
	deviceObj.RawSetString("codename", lua.LString(device.Codename))

	deviceObj.RawSetString("getImei", state.NewFunction(func(L *lua.LState) int {
		result := device.GetImei()
		L.Push(lua.LString(result))
		return 1
	}))

	deviceObj.RawSetString("getAndroidId", state.NewFunction(func(L *lua.LState) int {
		result := device.GetAndroidId()
		L.Push(lua.LString(result))
		return 1
	}))

	deviceObj.RawSetString("getWifiMac", state.NewFunction(func(L *lua.LState) int {
		result := device.GetWifiMac()
		L.Push(lua.LString(result))
		return 1
	}))

	deviceObj.RawSetString("getWlanMac", state.NewFunction(func(L *lua.LState) int {
		result := device.GetWlanMac()
		L.Push(lua.LString(result))
		return 1
	}))

	deviceObj.RawSetString("getIp", state.NewFunction(func(L *lua.LState) int {
		result := device.GetIp()
		L.Push(lua.LString(result))
		return 1
	}))

	deviceObj.RawSetString("getBrightness", state.NewFunction(func(L *lua.LState) int {
		result := device.GetBrightness()
		L.Push(lua.LString(result))
		return 1
	}))

	deviceObj.RawSetString("getBrightnessMode", state.NewFunction(func(L *lua.LState) int {
		result := device.GetBrightnessMode()
		L.Push(lua.LString(result))
		return 1
	}))

	deviceObj.RawSetString("getMusicVolume", state.NewFunction(func(L *lua.LState) int {
		result := device.GetMusicVolume()
		L.Push(lua.LNumber(result))
		return 1
	}))

	deviceObj.RawSetString("getNotificationVolume", state.NewFunction(func(L *lua.LState) int {
		result := device.GetNotificationVolume()
		L.Push(lua.LNumber(result))
		return 1
	}))

	deviceObj.RawSetString("getAlarmVolume", state.NewFunction(func(L *lua.LState) int {
		result := device.GetAlarmVolume()
		L.Push(lua.LNumber(result))
		return 1
	}))

	deviceObj.RawSetString("getMusicMaxVolume", state.NewFunction(func(L *lua.LState) int {
		result := device.GetMusicMaxVolume()
		L.Push(lua.LNumber(result))
		return 1
	}))

	deviceObj.RawSetString("getNotificationMaxVolume", state.NewFunction(func(L *lua.LState) int {
		result := device.GetNotificationMaxVolume()
		L.Push(lua.LNumber(result))
		return 1
	}))

	deviceObj.RawSetString("getAlarmMaxVolume", state.NewFunction(func(L *lua.LState) int {
		result := device.GetAlarmMaxVolume()
		L.Push(lua.LNumber(result))
		return 1
	}))

	deviceObj.RawSetString("setMusicVolume", state.NewFunction(func(L *lua.LState) int {
		volume := L.CheckInt(1)
		device.SetMusicVolume(volume)
		return 0
	}))

	deviceObj.RawSetString("setNotificationVolume", state.NewFunction(func(L *lua.LState) int {
		volume := L.CheckInt(1)
		device.SetNotificationVolume(volume)
		return 0
	}))

	deviceObj.RawSetString("setAlarmVolume", state.NewFunction(func(L *lua.LState) int {
		volume := L.CheckInt(1)
		device.SetAlarmVolume(volume)
		return 0
	}))

	deviceObj.RawSetString("getBattery", state.NewFunction(func(L *lua.LState) int {
		result := device.GetBattery()
		L.Push(lua.LNumber(result))
		return 1
	}))

	deviceObj.RawSetString("getBatteryStatus", state.NewFunction(func(L *lua.LState) int {
		result := device.GetBatteryStatus()
		L.Push(lua.LNumber(result))
		return 1
	}))

	deviceObj.RawSetString("setBatteryStatus", state.NewFunction(func(L *lua.LState) int {
		value := L.CheckInt(1)
		device.SetBatteryStatus(value)
		return 0
	}))

	deviceObj.RawSetString("setBatteryLevel", state.NewFunction(func(L *lua.LState) int {
		value := L.CheckInt(1)
		device.SetBatteryLevel(value)
		return 0
	}))

	deviceObj.RawSetString("getTotalMem", state.NewFunction(func(L *lua.LState) int {
		result := device.GetTotalMem()
		L.Push(lua.LNumber(result))
		return 1
	}))

	deviceObj.RawSetString("getAvailMem", state.NewFunction(func(L *lua.LState) int {
		result := device.GetAvailMem()
		L.Push(lua.LNumber(result))
		return 1
	}))

	deviceObj.RawSetString("isScreenOn", state.NewFunction(func(L *lua.LState) int {
		result := device.IsScreenOn()
		L.Push(lua.LBool(result))
		return 1
	}))

	deviceObj.RawSetString("isScreenUnlock", state.NewFunction(func(L *lua.LState) int {
		result := device.IsScreenUnlock()
		L.Push(lua.LBool(result))
		return 1
	}))

	deviceObj.RawSetString("wakeUp", state.NewFunction(func(L *lua.LState) int {
		device.WakeUp()
		return 0
	}))

	deviceObj.RawSetString("keepScreenOn", state.NewFunction(func(L *lua.LState) int {
		device.KeepScreenOn()
		return 0
	}))

	deviceObj.RawSetString("vibrate", state.NewFunction(func(L *lua.LState) int {
		ms := L.CheckInt(1)
		device.Vibrate(ms)
		return 0
	}))

	deviceObj.RawSetString("cancelVibration", state.NewFunction(func(L *lua.LState) int {
		device.CancelVibration()
		return 0
	}))

	engine.RegisterMethod("device.width", "设备分辨率宽度", func() int {
		width, _, _, _ := device.GetDisplayInfo(0)
		return width
	}, true)
	engine.RegisterMethod("device.height", "设备分辨率高度", func() int {
		_, height, _, _ := device.GetDisplayInfo(0)
		return height
	}, true)
	engine.RegisterMethod("device.sdkInt", "安卓系统API版本", func() int { return device.SdkInt }, true)
	engine.RegisterMethod("device.cpuAbi", "设备的CPU架构", func() string { return device.CpuAbi }, true)
	engine.RegisterMethod("device.getImei", "返回设备的IMEI", device.GetImei, true)
	engine.RegisterMethod("device.getAndroidId", "返回设备的Android ID", device.GetAndroidId, true)
	engine.RegisterMethod("device.getWifiMac", "获取设备WIFI-MAC", device.GetWifiMac, true)
	engine.RegisterMethod("device.getWlanMac", "获取设备以太网MAC", device.GetWlanMac, true)
	engine.RegisterMethod("device.getIp", "获取设备局域网IP地址", device.GetIp, true)
	engine.RegisterMethod("device.getBrightness", "返回当前的(手动)亮度", device.GetBrightness, true)
	engine.RegisterMethod("device.getBrightnessMode", "返回当前亮度模式", device.GetBrightnessMode, true)
	engine.RegisterMethod("device.getMusicVolume", "返回当前媒体音量", device.GetMusicVolume, true)
	engine.RegisterMethod("device.getNotificationVolume", "返回当前通知音量", device.GetNotificationVolume, true)
	engine.RegisterMethod("device.getAlarmVolume", "返回当前闹钟音量", device.GetAlarmVolume, true)
	engine.RegisterMethod("device.getMusicMaxVolume", "返回媒体音量的最大值", device.GetMusicMaxVolume, true)
	engine.RegisterMethod("device.getNotificationMaxVolume", "返回通知音量的最大值", device.GetNotificationMaxVolume, true)
	engine.RegisterMethod("device.getAlarmMaxVolume", "返回闹钟音量的最大值", device.GetAlarmMaxVolume, true)
	engine.RegisterMethod("device.setMusicVolume", "设置当前媒体音量", func(volume int) { device.SetMusicVolume(volume) }, true)
	engine.RegisterMethod("device.setNotificationVolume", "设置当前通知音量", func(volume int) { device.SetNotificationVolume(volume) }, true)
	engine.RegisterMethod("device.setAlarmVolume", "设置当前闹钟音量", func(volume int) { device.SetAlarmVolume(volume) }, true)
	engine.RegisterMethod("device.getBattery", "返回当前电量百分比", device.GetBattery, true)
	engine.RegisterMethod("device.getBatteryStatus", "获取电池状态", device.GetBatteryStatus, true)
	engine.RegisterMethod("device.setBatteryStatus", "模拟电池状态", func(value int) { device.SetBatteryStatus(value) }, true)
	engine.RegisterMethod("device.setBatteryLevel", "模拟电池电量百分百", func(value int) { device.SetBatteryLevel(value) }, true)
	engine.RegisterMethod("device.getTotalMem", "返回设备内存总量", device.GetTotalMem, true)
	engine.RegisterMethod("device.getAvailMem", "返回设备当前可用的内存", device.GetAvailMem, true)
	engine.RegisterMethod("device.isScreenOn", "返回设备屏幕是否是亮着的", device.IsScreenOn, true)
	engine.RegisterMethod("device.isScreenUnlock", "返回屏幕锁是否已经解开", device.IsScreenUnlock, true)
	engine.RegisterMethod("device.wakeUp", "唤醒设备", device.WakeUp, true)
	engine.RegisterMethod("device.keepScreenOn", "保持屏幕常亮", device.KeepScreenOn, true)
	engine.RegisterMethod("device.vibrate", "使设备震动一段时间", func(ms int) { device.Vibrate(ms) }, true)
	engine.RegisterMethod("device.cancelVibration", "如果设备处于震动状态，则取消震动", device.CancelVibration, true)

	return nil
}
