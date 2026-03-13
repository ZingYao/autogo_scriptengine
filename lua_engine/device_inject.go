package lua_engine

import (
	"github.com/Dasongzi1366/AutoGo/device"
	lua "github.com/yuin/gopher-lua"
)

func injectDeviceMethods(engine *LuaEngine) {

	engine.RegisterMethod("device.width", "设备分辨率宽度", func(displayId int) int {
		width, _, _, _ := device.GetDisplayInfo(displayId)
		return width
	}, true)
	engine.RegisterMethod("device.height", "设备分辨率高度", func(displayId int) int {
		_, height, _, _ := device.GetDisplayInfo(displayId)
		return height
	}, true)
	engine.RegisterMethod("device.sdkInt", "安卓系统API版本", func() int { return device.SdkInt }, true)
	engine.RegisterMethod("device.cpuAbi", "设备的CPU架构", func() string { return device.CpuAbi }, true)
	engine.RegisterMethod("device.buildId", "修订版本号", func() string { return device.BuildId }, true)
	engine.RegisterMethod("device.broad", "设备的主板型号", func() string { return device.Broad }, true)
	engine.RegisterMethod("device.brand", "与产品或硬件相关的厂商品牌", func() string { return device.Brand }, true)
	engine.RegisterMethod("device.device", "设备在工业设计中的名称", func() string { return device.Device }, true)
	engine.RegisterMethod("device.model", "设备型号", func() string { return device.Model }, true)
	engine.RegisterMethod("device.product", "整个产品的名称", func() string { return device.Product }, true)
	engine.RegisterMethod("device.bootloader", "设备Bootloader的版本", func() string { return device.Bootloader }, true)
	engine.RegisterMethod("device.hardware", "设备的硬件名称", func() string { return device.Hardware }, true)
	engine.RegisterMethod("device.fingerprint", "构建的唯一标识码", func() string { return device.Fingerprint }, true)
	engine.RegisterMethod("device.serial", "硬件序列号", func() string { return device.Serial }, true)
	engine.RegisterMethod("device.incremental", "设备构建的内部版本号", func() string { return device.Incremental }, true)
	engine.RegisterMethod("device.release", "Android系统版本号", func() string { return device.Release }, true)
	engine.RegisterMethod("device.baseOS", "设备的基础操作系统版本", func() string { return device.BaseOS }, true)
	engine.RegisterMethod("device.securityPatch", "安全补丁程序级别", func() string { return device.SecurityPatch }, true)
	engine.RegisterMethod("device.codename", "开发代号", func() string { return device.Codename }, true)
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

	registerDeviceLuaFunctions(engine)
}

func registerDeviceLuaFunctions(engine *LuaEngine) {
	state := engine.GetState()

	// 创建 device 表
	deviceTable := state.NewTable()
	state.SetGlobal("device", deviceTable)

	// 设备属性 - 函数形式
	state.SetField(deviceTable, "width", state.NewFunction(func(L *lua.LState) int {
		displayId := 0
		if L.GetTop() >= 1 {
			displayId = L.CheckInt(1)
		}
		width, _, _, _ := device.GetDisplayInfo(displayId)
		L.Push(lua.LNumber(width))
		return 1
	}))

	state.SetField(deviceTable, "height", state.NewFunction(func(L *lua.LState) int {
		displayId := 0
		if L.GetTop() >= 1 {
			displayId = L.CheckInt(1)
		}
		_, height, _, _ := device.GetDisplayInfo(displayId)
		L.Push(lua.LNumber(height))
		return 1
	}))

	// 设备属性 - 静态值
	state.SetField(deviceTable, "sdkInt", lua.LNumber(device.SdkInt))
	state.SetField(deviceTable, "cpuAbi", lua.LString(device.CpuAbi))
	state.SetField(deviceTable, "buildId", lua.LString(device.BuildId))
	state.SetField(deviceTable, "broad", lua.LString(device.Broad))
	state.SetField(deviceTable, "brand", lua.LString(device.Brand))
	state.SetField(deviceTable, "device_", lua.LString(device.Device)) // 避免与表名冲突
	state.SetField(deviceTable, "model", lua.LString(device.Model))
	state.SetField(deviceTable, "product", lua.LString(device.Product))
	state.SetField(deviceTable, "bootloader", lua.LString(device.Bootloader))
	state.SetField(deviceTable, "hardware", lua.LString(device.Hardware))
	state.SetField(deviceTable, "fingerprint", lua.LString(device.Fingerprint))
	state.SetField(deviceTable, "serial", lua.LString(device.Serial))
	state.SetField(deviceTable, "incremental", lua.LString(device.Incremental))
	state.SetField(deviceTable, "release", lua.LString(device.Release))
	state.SetField(deviceTable, "baseOS", lua.LString(device.BaseOS))
	state.SetField(deviceTable, "securityPatch", lua.LString(device.SecurityPatch))
	state.SetField(deviceTable, "codename", lua.LString(device.Codename))

	// 设备方法
	state.SetField(deviceTable, "getImei", state.NewFunction(func(L *lua.LState) int {
		result := device.GetImei()
		L.Push(lua.LString(result))
		return 1
	}))

	state.SetField(deviceTable, "getAndroidId", state.NewFunction(func(L *lua.LState) int {
		result := device.GetAndroidId()
		L.Push(lua.LString(result))
		return 1
	}))

	state.SetField(deviceTable, "getWifiMac", state.NewFunction(func(L *lua.LState) int {
		result := device.GetWifiMac()
		L.Push(lua.LString(result))
		return 1
	}))

	state.SetField(deviceTable, "getWlanMac", state.NewFunction(func(L *lua.LState) int {
		result := device.GetWlanMac()
		L.Push(lua.LString(result))
		return 1
	}))

	state.SetField(deviceTable, "getIp", state.NewFunction(func(L *lua.LState) int {
		result := device.GetIp()
		L.Push(lua.LString(result))
		return 1
	}))

	state.SetField(deviceTable, "getBrightness", state.NewFunction(func(L *lua.LState) int {
		result := device.GetBrightness()
		L.Push(lua.LString(result))
		return 1
	}))

	state.SetField(deviceTable, "getBrightnessMode", state.NewFunction(func(L *lua.LState) int {
		result := device.GetBrightnessMode()
		L.Push(lua.LString(result))
		return 1
	}))

	state.SetField(deviceTable, "getMusicVolume", state.NewFunction(func(L *lua.LState) int {
		result := device.GetMusicVolume()
		L.Push(lua.LNumber(result))
		return 1
	}))

	state.SetField(deviceTable, "getNotificationVolume", state.NewFunction(func(L *lua.LState) int {
		result := device.GetNotificationVolume()
		L.Push(lua.LNumber(result))
		return 1
	}))

	state.SetField(deviceTable, "getAlarmVolume", state.NewFunction(func(L *lua.LState) int {
		result := device.GetAlarmVolume()
		L.Push(lua.LNumber(result))
		return 1
	}))

	state.SetField(deviceTable, "getMusicMaxVolume", state.NewFunction(func(L *lua.LState) int {
		result := device.GetMusicMaxVolume()
		L.Push(lua.LNumber(result))
		return 1
	}))

	state.SetField(deviceTable, "getNotificationMaxVolume", state.NewFunction(func(L *lua.LState) int {
		result := device.GetNotificationMaxVolume()
		L.Push(lua.LNumber(result))
		return 1
	}))

	state.SetField(deviceTable, "getAlarmMaxVolume", state.NewFunction(func(L *lua.LState) int {
		result := device.GetAlarmMaxVolume()
		L.Push(lua.LNumber(result))
		return 1
	}))

	state.SetField(deviceTable, "setMusicVolume", state.NewFunction(func(L *lua.LState) int {
		volume := L.CheckInt(1)
		device.SetMusicVolume(volume)
		return 0
	}))

	state.SetField(deviceTable, "setNotificationVolume", state.NewFunction(func(L *lua.LState) int {
		volume := L.CheckInt(1)
		device.SetNotificationVolume(volume)
		return 0
	}))

	state.SetField(deviceTable, "setAlarmVolume", state.NewFunction(func(L *lua.LState) int {
		volume := L.CheckInt(1)
		device.SetAlarmVolume(volume)
		return 0
	}))

	state.SetField(deviceTable, "getBattery", state.NewFunction(func(L *lua.LState) int {
		result := device.GetBattery()
		L.Push(lua.LNumber(result))
		return 1
	}))

	state.SetField(deviceTable, "getBatteryStatus", state.NewFunction(func(L *lua.LState) int {
		result := device.GetBatteryStatus()
		L.Push(lua.LNumber(result))
		return 1
	}))

	state.SetField(deviceTable, "setBatteryStatus", state.NewFunction(func(L *lua.LState) int {
		value := L.CheckInt(1)
		device.SetBatteryStatus(value)
		return 0
	}))

	state.SetField(deviceTable, "setBatteryLevel", state.NewFunction(func(L *lua.LState) int {
		value := L.CheckInt(1)
		device.SetBatteryLevel(value)
		return 0
	}))

	state.SetField(deviceTable, "getTotalMem", state.NewFunction(func(L *lua.LState) int {
		result := device.GetTotalMem()
		L.Push(lua.LNumber(result))
		return 1
	}))

	state.SetField(deviceTable, "getAvailMem", state.NewFunction(func(L *lua.LState) int {
		result := device.GetAvailMem()
		L.Push(lua.LNumber(result))
		return 1
	}))

	state.SetField(deviceTable, "isScreenOn", state.NewFunction(func(L *lua.LState) int {
		result := device.IsScreenOn()
		L.Push(lua.LBool(result))
		return 1
	}))

	state.SetField(deviceTable, "isScreenUnlock", state.NewFunction(func(L *lua.LState) int {
		result := device.IsScreenUnlock()
		L.Push(lua.LBool(result))
		return 1
	}))

	state.SetField(deviceTable, "wakeUp", state.NewFunction(func(L *lua.LState) int {
		device.WakeUp()
		return 0
	}))

	state.SetField(deviceTable, "keepScreenOn", state.NewFunction(func(L *lua.LState) int {
		device.KeepScreenOn()
		return 0
	}))

	state.SetField(deviceTable, "vibrate", state.NewFunction(func(L *lua.LState) int {
		ms := L.CheckInt(1)
		device.Vibrate(ms)
		return 0
	}))

	state.SetField(deviceTable, "cancelVibration", state.NewFunction(func(L *lua.LState) int {
		device.CancelVibration()
		return 0
	}))
}
