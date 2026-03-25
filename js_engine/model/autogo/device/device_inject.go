package device

import (
	"github.com/ZingYao/autogo_scriptengine/js_engine/model"

	"github.com/Dasongzi1366/AutoGo/device"
	"github.com/dop251/goja"
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
	vm := engine.GetVM()

	deviceObj := vm.NewObject()
	vm.Set("device", deviceObj)

	width, height, _, _ := device.GetDisplayInfo(0)
	deviceObj.Set("width", width)
	deviceObj.Set("height", height)

	deviceObj.Set("sdkInt", device.SdkInt)
	deviceObj.Set("cpuAbi", device.CpuAbi)
	deviceObj.Set("buildId", device.BuildId)
	deviceObj.Set("broad", device.Broad)
	deviceObj.Set("brand", device.Brand)
	deviceObj.Set("deviceName", device.Device)
	deviceObj.Set("model", device.Model)
	deviceObj.Set("product", device.Product)
	deviceObj.Set("bootloader", device.Bootloader)
	deviceObj.Set("hardware", device.Hardware)
	deviceObj.Set("fingerprint", device.Fingerprint)
	deviceObj.Set("serial", device.Serial)
	deviceObj.Set("incremental", device.Incremental)
	deviceObj.Set("release", device.Release)
	deviceObj.Set("baseOS", device.BaseOS)
	deviceObj.Set("securityPatch", device.SecurityPatch)
	deviceObj.Set("codename", device.Codename)

	deviceObj.Set("getDisplayInfo", func(call goja.FunctionCall) goja.Value {
		displayId := int(call.Argument(0).ToInteger())
		width, height, dpi, rotation := device.GetDisplayInfo(displayId)
		result := vm.NewObject()
		result.Set("width", width)
		result.Set("height", height)
		result.Set("dpi", dpi)
		result.Set("rotation", rotation)
		return result
	})

	deviceObj.Set("getImei", func(call goja.FunctionCall) goja.Value {
		result := device.GetImei()
		return vm.ToValue(result)
	})

	deviceObj.Set("getAndroidId", func(call goja.FunctionCall) goja.Value {
		result := device.GetAndroidId()
		return vm.ToValue(result)
	})

	deviceObj.Set("getWifiMac", func(call goja.FunctionCall) goja.Value {
		result := device.GetWifiMac()
		return vm.ToValue(result)
	})

	deviceObj.Set("getWlanMac", func(call goja.FunctionCall) goja.Value {
		result := device.GetWlanMac()
		return vm.ToValue(result)
	})

	deviceObj.Set("getNotification", func(call goja.FunctionCall) goja.Value {
		result := device.GetNotification()
		return vm.ToValue(result)
	})

	deviceObj.Set("getIp", func(call goja.FunctionCall) goja.Value {
		result := device.GetIp()
		return vm.ToValue(result)
	})

	deviceObj.Set("getBrightness", func(call goja.FunctionCall) goja.Value {
		result := device.GetBrightness()
		return vm.ToValue(result)
	})

	deviceObj.Set("getBrightnessMode", func(call goja.FunctionCall) goja.Value {
		result := device.GetBrightnessMode()
		return vm.ToValue(result)
	})

	deviceObj.Set("getMusicVolume", func(call goja.FunctionCall) goja.Value {
		result := device.GetMusicVolume()
		return vm.ToValue(result)
	})

	deviceObj.Set("getNotificationVolume", func(call goja.FunctionCall) goja.Value {
		result := device.GetNotificationVolume()
		return vm.ToValue(result)
	})

	deviceObj.Set("getAlarmVolume", func(call goja.FunctionCall) goja.Value {
		result := device.GetAlarmVolume()
		return vm.ToValue(result)
	})

	deviceObj.Set("getMusicMaxVolume", func(call goja.FunctionCall) goja.Value {
		result := device.GetMusicMaxVolume()
		return vm.ToValue(result)
	})

	deviceObj.Set("getNotificationMaxVolume", func(call goja.FunctionCall) goja.Value {
		result := device.GetNotificationMaxVolume()
		return vm.ToValue(result)
	})

	deviceObj.Set("getAlarmMaxVolume", func(call goja.FunctionCall) goja.Value {
		result := device.GetAlarmMaxVolume()
		return vm.ToValue(result)
	})

	deviceObj.Set("setMusicVolume", func(call goja.FunctionCall) goja.Value {
		volume := int(call.Argument(0).ToInteger())
		device.SetMusicVolume(volume)
		return goja.Undefined()
	})

	deviceObj.Set("setNotificationVolume", func(call goja.FunctionCall) goja.Value {
		volume := int(call.Argument(0).ToInteger())
		device.SetNotificationVolume(volume)
		return goja.Undefined()
	})

	deviceObj.Set("setAlarmVolume", func(call goja.FunctionCall) goja.Value {
		volume := int(call.Argument(0).ToInteger())
		device.SetAlarmVolume(volume)
		return goja.Undefined()
	})

	deviceObj.Set("getBattery", func(call goja.FunctionCall) goja.Value {
		result := device.GetBattery()
		return vm.ToValue(result)
	})

	deviceObj.Set("getBatteryStatus", func(call goja.FunctionCall) goja.Value {
		result := device.GetBatteryStatus()
		return vm.ToValue(result)
	})

	deviceObj.Set("setBatteryStatus", func(call goja.FunctionCall) goja.Value {
		value := int(call.Argument(0).ToInteger())
		device.SetBatteryStatus(value)
		return goja.Undefined()
	})

	deviceObj.Set("setBatteryLevel", func(call goja.FunctionCall) goja.Value {
		value := int(call.Argument(0).ToInteger())
		device.SetBatteryLevel(value)
		return goja.Undefined()
	})

	deviceObj.Set("getTotalMem", func(call goja.FunctionCall) goja.Value {
		result := device.GetTotalMem()
		return vm.ToValue(result)
	})

	deviceObj.Set("getAvailMem", func(call goja.FunctionCall) goja.Value {
		result := device.GetAvailMem()
		return vm.ToValue(result)
	})

	deviceObj.Set("isScreenOn", func(call goja.FunctionCall) goja.Value {
		result := device.IsScreenOn()
		return vm.ToValue(result)
	})

	deviceObj.Set("isScreenUnlock", func(call goja.FunctionCall) goja.Value {
		result := device.IsScreenUnlock()
		return vm.ToValue(result)
	})

	deviceObj.Set("setDisplayPower", func(call goja.FunctionCall) goja.Value {
		on := call.Argument(0).ToBoolean()
		device.SetDisplayPower(on)
		return goja.Undefined()
	})

	deviceObj.Set("wakeUp", func(call goja.FunctionCall) goja.Value {
		device.WakeUp()
		return goja.Undefined()
	})

	deviceObj.Set("reboot", func(call goja.FunctionCall) goja.Value {
		device.Reboot()
		return goja.Undefined()
	})

	deviceObj.Set("keepScreenOn", func(call goja.FunctionCall) goja.Value {
		device.KeepScreenOn()
		return goja.Undefined()
	})

	deviceObj.Set("vibrate", func(call goja.FunctionCall) goja.Value {
		ms := int(call.Argument(0).ToInteger())
		device.Vibrate(ms)
		return goja.Undefined()
	})

	deviceObj.Set("cancelVibration", func(call goja.FunctionCall) goja.Value {
		device.CancelVibration()
		return goja.Undefined()
	})

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
	engine.RegisterMethod("device.getDisplayInfo", "获取指定屏幕的分辨率信息", func(displayId int) (int, int, int, int) {
		return device.GetDisplayInfo(displayId)
	}, true)
	engine.RegisterMethod("device.getImei", "返回设备的IMEI", device.GetImei, true)
	engine.RegisterMethod("device.getAndroidId", "返回设备的Android ID", device.GetAndroidId, true)
	engine.RegisterMethod("device.getWifiMac", "获取设备WIFI-MAC", device.GetWifiMac, true)
	engine.RegisterMethod("device.getWlanMac", "获取设备以太网MAC", device.GetWlanMac, true)
	engine.RegisterMethod("device.getNotification", "获取通知", device.GetNotification, true)
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
	engine.RegisterMethod("device.setDisplayPower", "设置显示电源", func(on bool) { device.SetDisplayPower(on) }, true)
	engine.RegisterMethod("device.wakeUp", "唤醒设备", device.WakeUp, true)
	engine.RegisterMethod("device.reboot", "重启设备", device.Reboot, true)
	engine.RegisterMethod("device.keepScreenOn", "保持屏幕常亮", device.KeepScreenOn, true)
	engine.RegisterMethod("device.vibrate", "使设备震动一段时间", func(ms int) { device.Vibrate(ms) }, true)
	engine.RegisterMethod("device.cancelVibration", "如果设备处于震动状态，则取消震动", device.CancelVibration, true)

	return nil
}
