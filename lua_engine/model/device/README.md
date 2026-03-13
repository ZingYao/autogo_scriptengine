# Device 模块

Device 模块提供了设备信息获取和设备控制功能。

## 属性

### device.width

设备分辨率宽度，横屏和竖屏时的数值不同。

**类型**: `number` (只读)

**调用示例**:
```lua
local width = device.width
print("屏幕宽度: " .. width)
```

### device.height

设备分辨率高度，横屏和竖屏时的数值不同。

**类型**: `number` (只读)

**调用示例**:
```lua
local height = device.height
print("屏幕高度: " .. height)
```

### device.sdkInt

安卓系统 API 版本。例如安卓 4.4 的 sdkInt 为 19。

**类型**: `number` (只读)

**调用示例**:
```lua
local sdkInt = device.sdkInt
print("SDK版本: " .. sdkInt)
```

### device.cpuAbi

设备的 CPU 架构，如 `"arm64-v8a"`, `"x86"`, `"x86_64"`。

**类型**: `string` (只读)

**调用示例**:
```lua
local cpuAbi = device.cpuAbi
print("CPU架构: " .. cpuAbi)
```

### device.buildId

设备构建ID。

**类型**: `string` (只读)

**调用示例**:
```lua
local buildId = device.buildId
print("构建ID: " .. buildId)
```

### device.broad

设备主板。

**类型**: `string` (只读)

**调用示例**:
```lua
local broad = device.broad
print("主板: " .. broad)
```

### device.brand

设备品牌。

**类型**: `string` (只读)

**调用示例**:
```lua
local brand = device.brand
print("品牌: " .. brand)
```

### device.deviceName

设备名称。

**类型**: `string` (只读)

**调用示例**:
```lua
local deviceName = device.deviceName
print("设备名称: " .. deviceName)
```

### device.model

设备型号。

**类型**: `string` (只读)

**调用示例**:
```lua
local model = device.model
print("型号: " .. model)
```

### device.product

设备产品名。

**类型**: `string` (只读)

**调用示例**:
```lua
local product = device.product
print("产品名: " .. product)
```

### device.bootloader

设备引导程序版本。

**类型**: `string` (只读)

**调用示例**:
```lua
local bootloader = device.bootloader
print("引导程序: " .. bootloader)
```

### device.hardware

设备硬件名称。

**类型**: `string` (只读)

**调用示例**:
```lua
local hardware = device.hardware
print("硬件: " .. hardware)
```

### device.fingerprint

设备指纹。

**类型**: `string` (只读)

**调用示例**:
```lua
local fingerprint = device.fingerprint
print("指纹: " .. fingerprint)
```

### device.serial

设备序列号。

**类型**: `string` (只读)

**调用示例**:
```lua
local serial = device.serial
print("序列号: " .. serial)
```

### device.incremental

设备增量版本。

**类型**: `string` (只读)

**调用示例**:
```lua
local incremental = device.incremental
print("增量版本: " .. incremental)
```

### device.release

设备发布版本。

**类型**: `string` (只读)

**调用示例**:
```lua
local release = device.release
print("发布版本: " .. release)
```

### device.baseOS

设备基础操作系统。

**类型**: `string` (只读)

**调用示例**:
```lua
local baseOS = device.baseOS
print("基础操作系统: " .. baseOS)
```

### device.securityPatch

设备安全补丁级别。

**类型**: `string` (只读)

**调用示例**:
```lua
local securityPatch = device.securityPatch
print("安全补丁: " .. securityPatch)
```

### device.codename

设备代号。

**类型**: `string` (只读)

**调用示例**:
```lua
local codename = device.codename
print("代号: " .. codename)
```

## 方法列表

### device.getImei()

返回设备的 IMEI。

**返回值**: `string` - IMEI 号码

**调用示例**:
```lua
local imei = device.getImei()
print("IMEI: " .. imei)
```

### device.getAndroidId()

返回设备的 Android ID。

**返回值**: `string` - Android ID

**调用示例**:
```lua
local androidId = device.getAndroidId()
print("Android ID: " .. androidId)
```

### device.getWifiMac()

获取设备 WIFI-MAC。

**返回值**: `string` - WIFI MAC 地址

**调用示例**:
```lua
local wifiMac = device.getWifiMac()
print("WIFI MAC: " .. wifiMac)
```

### device.getWlanMac()

获取设备以太网 MAC。

**返回值**: `string` - 以太网 MAC 地址

**调用示例**:
```lua
local wlanMac = device.getWlanMac()
print("以太网 MAC: " .. wlanMac)
```

### device.getIp()

获取设备局域网 IP 地址。

**返回值**: `string` - IP 地址

**调用示例**:
```lua
local ip = device.getIp()
print("IP地址: " .. ip)
```

### device.getBrightness()

返回当前的(手动)亮度。范围为 0~255。

**返回值**: `number` - 当前亮度值

**调用示例**:
```lua
local brightness = device.getBrightness()
print("当前亮度: " .. brightness)
```

### device.getBrightnessMode()

返回当前亮度模式，0 为手动亮度，1 为自动亮度。

**返回值**: `number` - 亮度模式

**调用示例**:
```lua
local brightnessMode = device.getBrightnessMode()
if brightnessMode == 0 then
    print("手动亮度模式")
else
    print("自动亮度模式")
end
```

### device.getMusicVolume()

返回当前媒体音量。

**返回值**: `number` - 媒体音量值

**调用示例**:
```lua
local musicVolume = device.getMusicVolume()
print("媒体音量: " .. musicVolume)
```

### device.getNotificationVolume()

返回当前通知音量。

**返回值**: `number` - 通知音量值

**调用示例**:
```lua
local notificationVolume = device.getNotificationVolume()
print("通知音量: " .. notificationVolume)
```

### device.getAlarmVolume()

返回当前闹钟音量。

**返回值**: `number` - 闹钟音量值

**调用示例**:
```lua
local alarmVolume = device.getAlarmVolume()
print("闹钟音量: " .. alarmVolume)
```

### device.getMusicMaxVolume()

返回媒体音量的最大值。

**返回值**: `number` - 媒体音量最大值

**调用示例**:
```lua
local maxMusicVolume = device.getMusicMaxVolume()
print("媒体音量最大值: " .. maxMusicVolume)
```

### device.getNotificationMaxVolume()

返回通知音量的最大值。

**返回值**: `number` - 通知音量最大值

**调用示例**:
```lua
local maxNotificationVolume = device.getNotificationMaxVolume()
print("通知音量最大值: " .. maxNotificationVolume)
```

### device.getAlarmMaxVolume()

返回闹钟音量的最大值。

**返回值**: `number` - 闹钟音量最大值

**调用示例**:
```lua
local maxAlarmVolume = device.getAlarmMaxVolume()
print("闹钟音量最大值: " .. maxAlarmVolume)
```

### device.setMusicVolume(volume)

设置当前媒体音量。

**参数**:
- `volume` (number): 音量值

**返回值**: `undefined`

**调用示例**:
```lua
device.setMusicVolume(50)
print("媒体音量已设置为50")
```

### device.setNotificationVolume(volume)

设置当前通知音量。

**参数**:
- `volume` (number): 音量值

**返回值**: `undefined`

**调用示例**:
```lua
device.setNotificationVolume(50)
print("通知音量已设置为50")
```

### device.setAlarmVolume(volume)

设置当前闹钟音量。

**参数**:
- `volume` (number): 音量值

**返回值**: `undefined`

**调用示例**:
```lua
device.setAlarmVolume(50)
print("闹钟音量已设置为50")
```

### device.getBattery()

返回当前电量百分比。

**返回值**: `number` - 电量百分比 (0-100)

**调用示例**:
```lua
local battery = device.getBattery()
print("电量: " .. battery .. "%")
```

### device.getBatteryStatus()

获取电池状态。1：没有充电；2：正充电；3：没插充电器；4：不充电；5：电池充满。

**返回值**: `number` - 电池状态

**调用示例**:
```lua
local batteryStatus = device.getBatteryStatus()
print("电池状态: " .. batteryStatus)
```

### device.setBatteryStatus(value)

模拟电池状态。1：没有充电；2：正充电；5：电池充满。

**参数**:
- `value` (number): 电池状态值

**返回值**: `undefined`

**调用示例**:
```lua
device.setBatteryStatus(2)
print("电池状态已设置为充电中")
```

### device.setBatteryLevel(value)

模拟电池电量百分百 0-100。

**参数**:
- `value` (number): 电量百分比 (0-100)

**返回值**: `undefined`

**调用示例**:
```lua
device.setBatteryLevel(80)
print("电池电量已设置为80%")
```

### device.getTotalMem()

返回设备内存总量，单位(KB)。1MB = 1024KB。

**返回值**: `number` - 内存总量(KB)

**调用示例**:
```lua
local totalMem = device.getTotalMem()
print("总内存: " .. (totalMem / 1024) .. "MB")
```

### device.getAvailMem()

返回设备当前可用的内存，单位字节(KB)。

**返回值**: `number` - 可用内存(KB)

**调用示例**:
```lua
local availMem = device.getAvailMem()
print("可用内存: " .. (availMem / 1024) .. "MB")
```

### device.isScreenOn()

返回设备屏幕是否是亮着的。如果屏幕亮着，返回 `true`；否则返回 `false`。

**返回值**: `boolean` - 屏幕是否亮着

**调用示例**:
```lua
if device.isScreenOn() then
    print("屏幕已点亮")
else
    print("屏幕已关闭")
end
```

### device.isScreenUnlock()

返回屏幕锁是否已经解开。已经解开返回 `true`；否则返回 `false`。

**返回值**: `boolean` - 屏幕是否已解锁

**调用示例**:
```lua
if device.isScreenUnlock() then
    print("屏幕已解锁")
else
    print("屏幕已锁定")
end
```

### device.wakeUp()

唤醒设备，包括唤醒设备 CPU、屏幕等，可以用来点亮屏幕。

**返回值**: `undefined`

**调用示例**:
```lua
device.wakeUp()
print("设备已唤醒")
```

### device.keepScreenOn()

保持屏幕常亮。

**返回值**: `undefined`

**调用示例**:
```lua
device.keepScreenOn()
print("屏幕常亮已启用")
```

### device.vibrate(ms)

使设备震动一段时间，单位毫秒，需要 root 权限。

**参数**:
- `ms` (number): 震动时间(毫秒)

**返回值**: `undefined`

**调用示例**:
```lua
device.vibrate(1000)
print("设备震动1秒")
```

### device.cancelVibration()

如果设备处于震动状态，则取消震动，需要 root 权限。

**返回值**: `undefined`

**调用示例**:
```lua
device.cancelVibration()
print("震动已取消")
```

## 完整示例

```lua
-- 获取设备基本信息
print("=== 设备信息 ===")
print("分辨率: " .. device.width .. "x" .. device.height)
print("SDK版本: " .. device.sdkInt)
print("CPU架构: " .. device.cpuAbi)
print("品牌: " .. device.brand)
print("型号: " .. device.model)

-- 获取设备标识
print("\n=== 设备标识 ===")
print("IMEI: " .. device.getImei())
print("Android ID: " .. device.getAndroidId())
print("WIFI MAC: " .. device.getWifiMac())
print("IP地址: " .. device.getIp())

-- 音量控制
print("\n=== 音量控制 ===")
print("当前媒体音量: " .. device.getMusicVolume())
print("媒体音量最大值: " .. device.getMusicMaxVolume())
device.setMusicVolume(50)
print("媒体音量已设置为50")

-- 电池信息
print("\n=== 电池信息 ===")
print("电量: " .. device.getBattery() .. "%")
print("电池状态: " .. device.getBatteryStatus())

-- 屏幕控制
print("\n=== 屏幕控制 ===")
if device.isScreenOn() then
    print("屏幕已点亮")
end
if device.isScreenUnlock() then
    print("屏幕已解锁")
end

-- 震动
print("\n=== 震动 ===")
device.vibrate(1000)
```
