package device

import (
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"

	autogodevice "github.com/Dasongzi1366/AutoGo/device"
)

// DeviceModule 是 go-lua-vm 迁移后的模块壳。
type DeviceModule struct{}

func New() *DeviceModule { return &DeviceModule{} }

func (m *DeviceModule) Name() string { return "device" }

func (m *DeviceModule) IsAvailable() bool { return true }

func (m *DeviceModule) Register(engine model.Engine) error {
	width, height, _, _ := autogodevice.GetDisplayInfo(0)

	_ = engine.RegisterValue("device.width", width)
	_ = engine.RegisterValue("device.height", height)
	_ = engine.RegisterValue("device.sdkInt", autogodevice.SdkInt)
	_ = engine.RegisterValue("device.cpuAbi", autogodevice.CpuAbi)
	_ = engine.RegisterValue("device.buildId", autogodevice.BuildId)
	_ = engine.RegisterValue("device.broad", autogodevice.Broad)
	_ = engine.RegisterValue("device.brand", autogodevice.Brand)
	_ = engine.RegisterValue("device.deviceName", autogodevice.Device)
	_ = engine.RegisterValue("device.model", autogodevice.Model)
	_ = engine.RegisterValue("device.product", autogodevice.Product)
	_ = engine.RegisterValue("device.bootloader", autogodevice.Bootloader)
	_ = engine.RegisterValue("device.hardware", autogodevice.Hardware)
	_ = engine.RegisterValue("device.fingerprint", autogodevice.Fingerprint)
	_ = engine.RegisterValue("device.serial", autogodevice.Serial)
	_ = engine.RegisterValue("device.incremental", autogodevice.Incremental)
	_ = engine.RegisterValue("device.release", autogodevice.Release)
	_ = engine.RegisterValue("device.baseOS", autogodevice.BaseOS)
	_ = engine.RegisterValue("device.securityPatch", autogodevice.SecurityPatch)
	_ = engine.RegisterValue("device.codename", autogodevice.Codename)

	engine.RegisterMethod("device.getDisplayInfo", "获取指定屏幕的分辨率信息", func(displayID ...int) map[string]interface{} {
		targetDisplayID := 0
		if len(displayID) > 0 {
			targetDisplayID = displayID[0]
		}
		width, height, dpi, rotation := autogodevice.GetDisplayInfo(targetDisplayID)
		return map[string]interface{}{"width": width, "height": height, "dpi": dpi, "rotation": rotation}
	}, true)
	engine.RegisterMethod("device.width", "返回默认屏幕宽度", func() int {
		width, _, _, _ := autogodevice.GetDisplayInfo(0)
		return width
	}, true)
	engine.RegisterMethod("device.height", "返回默认屏幕高度", func() int {
		_, height, _, _ := autogodevice.GetDisplayInfo(0)
		return height
	}, true)
	engine.RegisterMethod("device.sdkInt", "安卓系统 API 版本", func() int {
		return autogodevice.SdkInt
	}, true)
	engine.RegisterMethod("device.cpuAbi", "设备 CPU 架构", func() string {
		return autogodevice.CpuAbi
	}, true)
	engine.RegisterMethod("device.getImei", "返回设备的 IMEI", autogodevice.GetImei, true)
	engine.RegisterMethod("device.getAndroidId", "返回设备的 Android ID", autogodevice.GetAndroidId, true)
	engine.RegisterMethod("device.getWifiMac", "获取设备 WIFI MAC", autogodevice.GetWifiMac, true)
	engine.RegisterMethod("device.getWlanMac", "获取设备以太网 MAC", autogodevice.GetWlanMac, true)
	engine.RegisterMethod("device.getIp", "获取设备局域网 IP 地址", autogodevice.GetIp, true)
	engine.RegisterMethod("device.getNotification", "获取当前所有通知消息", autogodevice.GetNotification, true)
	engine.RegisterMethod("device.getBrightness", "返回当前手动亮度", autogodevice.GetBrightness, true)
	engine.RegisterMethod("device.getBrightnessMode", "返回当前亮度模式", autogodevice.GetBrightnessMode, true)
	engine.RegisterMethod("device.getMusicVolume", "返回当前媒体音量", autogodevice.GetMusicVolume, true)
	engine.RegisterMethod("device.getNotificationVolume", "返回当前通知音量", autogodevice.GetNotificationVolume, true)
	engine.RegisterMethod("device.getAlarmVolume", "返回当前闹钟音量", autogodevice.GetAlarmVolume, true)
	engine.RegisterMethod("device.getMusicMaxVolume", "返回媒体音量最大值", autogodevice.GetMusicMaxVolume, true)
	engine.RegisterMethod("device.getNotificationMaxVolume", "返回通知音量最大值", autogodevice.GetNotificationMaxVolume, true)
	engine.RegisterMethod("device.getAlarmMaxVolume", "返回闹钟音量最大值", autogodevice.GetAlarmMaxVolume, true)
	engine.RegisterMethod("device.setMusicVolume", "设置当前媒体音量", autogodevice.SetMusicVolume, true)
	engine.RegisterMethod("device.setNotificationVolume", "设置当前通知音量", autogodevice.SetNotificationVolume, true)
	engine.RegisterMethod("device.setAlarmVolume", "设置当前闹钟音量", autogodevice.SetAlarmVolume, true)
	engine.RegisterMethod("device.getBattery", "返回当前电量百分比", autogodevice.GetBattery, true)
	engine.RegisterMethod("device.getBatteryStatus", "获取电池状态", autogodevice.GetBatteryStatus, true)
	engine.RegisterMethod("device.setBatteryStatus", "模拟电池状态", autogodevice.SetBatteryStatus, true)
	engine.RegisterMethod("device.setBatteryLevel", "模拟电池电量百分比", autogodevice.SetBatteryLevel, true)
	engine.RegisterMethod("device.getTotalMem", "返回设备内存总量", autogodevice.GetTotalMem, true)
	engine.RegisterMethod("device.getAvailMem", "返回设备当前可用内存", autogodevice.GetAvailMem, true)
	engine.RegisterMethod("device.isScreenOn", "返回设备屏幕是否亮着", autogodevice.IsScreenOn, true)
	engine.RegisterMethod("device.isScreenUnlock", "返回屏幕锁是否已经解开", autogodevice.IsScreenUnlock, true)
	engine.RegisterMethod("device.wakeUp", "唤醒设备", autogodevice.WakeUp, true)
	engine.RegisterMethod("device.keepScreenOn", "保持屏幕常亮", autogodevice.KeepScreenOn, true)
	engine.RegisterMethod("device.vibrate", "使设备震动一段时间", autogodevice.Vibrate, true)
	engine.RegisterMethod("device.cancelVibration", "取消设备震动", autogodevice.CancelVibration, true)
	engine.RegisterMethod("device.setDisplayPower", "设置屏幕电源模式", autogodevice.SetDisplayPower, true)
	engine.RegisterMethod("device.reboot", "重启设备", autogodevice.Reboot, true)
	return nil
}

func GetModule() model.Module { return &DeviceModule{} }
