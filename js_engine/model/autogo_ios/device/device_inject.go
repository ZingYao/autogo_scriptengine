package device

import (
	"github.com/ZingYao/autogo_scriptengine/js_engine/model"

	autogodevice "github.com/Dasongzi1366/AutoGo/device"
	"github.com/dop251/goja"
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
	vm := engine.GetVM()
	width, height, scale, rotation := autogodevice.GetDisplayInfo()
	deviceObj := vm.NewObject()
	vm.Set("device", deviceObj)

	deviceObj.Set("width", width)
	deviceObj.Set("height", height)
	deviceObj.Set("scale", scale)
	deviceObj.Set("rotation", rotation)
	deviceObj.Set("model", autogodevice.Model)
	deviceObj.Set("release", autogodevice.Release)
	deviceObj.Set("serial", autogodevice.Serial)

	deviceObj.Set("getDisplayInfo", func(call goja.FunctionCall) goja.Value {
		width, height, scale, rotation := autogodevice.GetDisplayInfo()
		result := vm.NewObject()
		result.Set("width", width)
		result.Set("height", height)
		result.Set("scale", scale)
		result.Set("rotation", rotation)
		return result
	})
	deviceObj.Set("getBattery", func(call goja.FunctionCall) goja.Value { return vm.ToValue(autogodevice.GetBattery()) })
	deviceObj.Set("getBatteryStatus", func(call goja.FunctionCall) goja.Value { return vm.ToValue(autogodevice.GetBatteryStatus()) })
	deviceObj.Set("isScreenOn", func(call goja.FunctionCall) goja.Value { return vm.ToValue(autogodevice.IsScreenOn()) })
	deviceObj.Set("isScreenUnlock", func(call goja.FunctionCall) goja.Value { return vm.ToValue(autogodevice.IsScreenUnlock()) })
	deviceObj.Set("getBrightness", func(call goja.FunctionCall) goja.Value { return vm.ToValue(autogodevice.GetBrightness()) })
	deviceObj.Set("getTotalMem", func(call goja.FunctionCall) goja.Value { return vm.ToValue(autogodevice.GetTotalMem()) })
	deviceObj.Set("getAvailMem", func(call goja.FunctionCall) goja.Value { return vm.ToValue(autogodevice.GetAvailMem()) })
	deviceObj.Set("wakeUp", func(call goja.FunctionCall) goja.Value { autogodevice.WakeUp(); return goja.Undefined() })
	deviceObj.Set("keepScreenOn", func(call goja.FunctionCall) goja.Value { autogodevice.KeepScreenOn(); return goja.Undefined() })
	deviceObj.Set("getIp", func(call goja.FunctionCall) goja.Value { return vm.ToValue(autogodevice.GetIp()) })
	deviceObj.Set("getWifiMac", func(call goja.FunctionCall) goja.Value { return vm.ToValue(autogodevice.GetWifiMac()) })
	deviceObj.Set("reboot", func(call goja.FunctionCall) goja.Value { autogodevice.Reboot(); return goja.Undefined() })

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
