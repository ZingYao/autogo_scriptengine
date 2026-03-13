# Device 模块

Device 模块提供了设备信息获取和设备控制功能。

## 属性

### device.width

设备分辨率宽度，横屏和竖屏时的数值不同。

**类型**: `number` (只读)

**调用示例**:
```javascript
const width = device.width;
console.log("屏幕宽度: " + width);
```

### device.height

设备分辨率高度，横屏和竖屏时的数值不同。

**类型**: `number` (只读)

**调用示例**:
```javascript
const height = device.height;
console.log("屏幕高度: " + height);
```

### device.sdkInt

安卓系统 API 版本。例如安卓 4.4 的 sdkInt 为 19。

**类型**: `number` (只读)

**调用示例**:
```javascript
const sdkInt = device.sdkInt;
console.log("SDK版本: " + sdkInt);
```

### device.cpuAbi

设备的 CPU 架构，如 `"arm64-v8a"`, `"x86"`, `"x86_64"`。

**类型**: `string` (只读)

**调用示例**:
```javascript
const cpuAbi = device.cpuAbi;
console.log("CPU架构: " + cpuAbi);
```

### device.buildId

设备构建ID。

**类型**: `string` (只读)

**调用示例**:
```javascript
const buildId = device.buildId;
console.log("构建ID: " + buildId);
```

### device.broad

设备主板。

**类型**: `string` (只读)

**调用示例**:
```javascript
const broad = device.broad;
console.log("主板: " + broad);
```

### device.brand

设备品牌。

**类型**: `string` (只读)

**调用示例**:
```javascript
const brand = device.brand;
console.log("品牌: " + brand);
```

### device.deviceName

设备名称。

**类型**: `string` (只读)

**调用示例**:
```javascript
const deviceName = device.deviceName;
console.log("设备名称: " + deviceName);
```

### device.model

设备型号。

**类型**: `string` (只读)

**调用示例**:
```javascript
const model = device.model;
console.log("型号: " + model);
```

### device.product

设备产品名。

**类型**: `string` (只读)

**调用示例**:
```javascript
const product = device.product;
console.log("产品名: " + product);
```

### device.bootloader

设备引导程序版本。

**类型**: `string` (只读)

**调用示例**:
```javascript
const bootloader = device.bootloader;
console.log("引导程序: " + bootloader);
```

### device.hardware

设备硬件名称。

**类型**: `string` (只读)

**调用示例**:
```javascript
const hardware = device.hardware;
console.log("硬件: " + hardware);
```

### device.fingerprint

设备指纹。

**类型**: `string` (只读)

**调用示例**:
```javascript
const fingerprint = device.fingerprint;
console.log("指纹: " + fingerprint);
```

### device.serial

设备序列号。

**类型**: `string` (只读)

**调用示例**:
```javascript
const serial = device.serial;
console.log("序列号: " + serial);
```

### device.incremental

设备增量版本。

**类型**: `string` (只读)

**调用示例**:
```javascript
const incremental = device.incremental;
console.log("增量版本: " + incremental);
```

### device.release

设备发布版本。

**类型**: `string` (只读)

**调用示例**:
```javascript
const release = device.release;
console.log("发布版本: " + release);
```

### device.baseOS

设备基础操作系统。

**类型**: `string` (只读)

**调用示例**:
```javascript
const baseOS = device.baseOS;
console.log("基础操作系统: " + baseOS);
```

### device.securityPatch

设备安全补丁级别。

**类型**: `string` (只读)

**调用示例**:
```javascript
const securityPatch = device.securityPatch;
console.log("安全补丁: " + securityPatch);
```

### device.codename

设备代号。

**类型**: `string` (只读)

**调用示例**:
```javascript
const codename = device.codename;
console.log("代号: " + codename);
```

## 方法列表

### device.getImei()

返回设备的 IMEI。

**返回值**: `string` - IMEI 号码

**调用示例**:
```javascript
const imei = device.getImei();
console.log("IMEI: " + imei);
```

### device.getAndroidId()

返回设备的 Android ID。

**返回值**: `string` - Android ID

**调用示例**:
```javascript
const androidId = device.getAndroidId();
console.log("Android ID: " + androidId);
```

### device.getWifiMac()

获取设备 WIFI-MAC。

**返回值**: `string` - WIFI MAC 地址

**调用示例**:
```javascript
const wifiMac = device.getWifiMac();
console.log("WIFI MAC: " + wifiMac);
```

### device.getWlanMac()

获取设备以太网 MAC。

**返回值**: `string` - 以太网 MAC 地址

**调用示例**:
```javascript
const wlanMac = device.getWlanMac();
console.log("以太网 MAC: " + wlanMac);
```

### device.getNotification()

获取通知信息。

**返回值**: `string` - 通知信息

**调用示例**:
```javascript
const notification = device.getNotification();
console.log("通知: " + notification);
```

### device.getIp()

获取设备局域网 IP 地址。

**返回值**: `string` - IP 地址

**调用示例**:
```javascript
const ip = device.getIp();
console.log("IP地址: " + ip);
```

### device.getBrightness()

返回当前的(手动)亮度。范围为 0~255。

**返回值**: `number` - 当前亮度值

**调用示例**:
```javascript
const brightness = device.getBrightness();
console.log("当前亮度: " + brightness);
```

### device.getBrightnessMode()

返回当前亮度模式，0 为手动亮度，1 为自动亮度。

**返回值**: `number` - 亮度模式

**调用示例**:
```javascript
const brightnessMode = device.getBrightnessMode();
if (brightnessMode === 0) {
    console.log("手动亮度模式");
} else {
    console.log("自动亮度模式");
}
```

### device.getMusicVolume()

返回当前媒体音量。

**返回值**: `number` - 媒体音量值

**调用示例**:
```javascript
const musicVolume = device.getMusicVolume();
console.log("媒体音量: " + musicVolume);
```

### device.getNotificationVolume()

返回当前通知音量。

**返回值**: `number` - 通知音量值

**调用示例**:
```javascript
const notificationVolume = device.getNotificationVolume();
console.log("通知音量: " + notificationVolume);
```

### device.getAlarmVolume()

返回当前闹钟音量。

**返回值**: `number` - 闹钟音量值

**调用示例**:
```javascript
const alarmVolume = device.getAlarmVolume();
console.log("闹钟音量: " + alarmVolume);
```

### device.getMusicMaxVolume()

返回媒体音量的最大值。

**返回值**: `number` - 媒体音量最大值

**调用示例**:
```javascript
const maxMusicVolume = device.getMusicMaxVolume();
console.log("媒体音量最大值: " + maxMusicVolume);
```

### device.getNotificationMaxVolume()

返回通知音量的最大值。

**返回值**: `number` - 通知音量最大值

**调用示例**:
```javascript
const maxNotificationVolume = device.getNotificationMaxVolume();
console.log("通知音量最大值: " + maxNotificationVolume);
```

### device.getAlarmMaxVolume()

返回闹钟音量的最大值。

**返回值**: `number` - 闹钟音量最大值

**调用示例**:
```javascript
const maxAlarmVolume = device.getAlarmMaxVolume();
console.log("闹钟音量最大值: " + maxAlarmVolume);
```

### device.setMusicVolume(volume)

设置当前媒体音量。

**参数**:
- `volume` (number): 音量值

**返回值**: `undefined`

**调用示例**:
```javascript
device.setMusicVolume(50);
console.log("媒体音量已设置为50");
```

### device.setNotificationVolume(volume)

设置当前通知音量。

**参数**:
- `volume` (number): 音量值

**返回值**: `undefined`

**调用示例**:
```javascript
device.setNotificationVolume(50);
console.log("通知音量已设置为50");
```

### device.setAlarmVolume(volume)

设置当前闹钟音量。

**参数**:
- `volume` (number): 音量值

**返回值**: `undefined`

**调用示例**:
```javascript
device.setAlarmVolume(50);
console.log("闹钟音量已设置为50");
```

### device.getBattery()

返回当前电量百分比。

**返回值**: `number` - 电量百分比 (0-100)

**调用示例**:
```javascript
const battery = device.getBattery();
console.log("电量: " + battery + "%");
```

### device.getBatteryStatus()

获取电池状态。1：没有充电；2：正充电；3：没插充电器；4：不充电；5：电池充满。

**返回值**: `number` - 电池状态

**调用示例**:
```javascript
const batteryStatus = device.getBatteryStatus();
console.log("电池状态: " + batteryStatus);
```

### device.setBatteryStatus(value)

模拟电池状态。1：没有充电；2：正充电；5：电池充满。

**参数**:
- `value` (number): 电池状态值

**返回值**: `undefined`

**调用示例**:
```javascript
device.setBatteryStatus(2);
console.log("电池状态已设置为充电中");
```

### device.setBatteryLevel(value)

模拟电池电量百分百 0-100。

**参数**:
- `value` (number): 电量百分比 (0-100)

**返回值**: `undefined`

**调用示例**:
```javascript
device.setBatteryLevel(80);
console.log("电池电量已设置为80%");
```

### device.getTotalMem()

返回设备内存总量，单位(KB)。1MB = 1024KB。

**返回值**: `number` - 内存总量(KB)

**调用示例**:
```javascript
const totalMem = device.getTotalMem();
console.log("总内存: " + (totalMem / 1024) + "MB");
```

### device.getAvailMem()

返回设备当前可用的内存，单位字节(KB)。

**返回值**: `number` - 可用内存(KB)

**调用示例**:
```javascript
const availMem = device.getAvailMem();
console.log("可用内存: " + (availMem / 1024) + "MB");
```

### device.isScreenOn()

返回设备屏幕是否是亮着的。如果屏幕亮着，返回 `true`；否则返回 `false`。

**返回值**: `boolean` - 屏幕是否亮着

**调用示例**:
```javascript
if (device.isScreenOn()) {
    console.log("屏幕已点亮");
} else {
    console.log("屏幕已关闭");
}
```

### device.isScreenUnlock()

返回屏幕锁是否已经解开。已经解开返回 `true`；否则返回 `false`。

**返回值**: `boolean` - 屏幕是否已解锁

**调用示例**:
```javascript
if (device.isScreenUnlock()) {
    console.log("屏幕已解锁");
} else {
    console.log("屏幕已锁定");
}
```

### device.wakeUp()

唤醒设备，包括唤醒设备 CPU、屏幕等，可以用来点亮屏幕。

**返回值**: `undefined`

**调用示例**:
```javascript
device.wakeUp();
console.log("设备已唤醒");
```

### device.keepScreenOn()

保持屏幕常亮。

**返回值**: `undefined`

**调用示例**:
```javascript
device.keepScreenOn();
console.log("屏幕常亮已启用");
```

### device.vibrate(ms)

使设备震动一段时间，单位毫秒，需要 root 权限。

**参数**:
- `ms` (number): 震动时间(毫秒)

**返回值**: `undefined`

**调用示例**:
```javascript
device.vibrate(1000);
console.log("设备震动1秒");
```

### device.cancelVibration()

如果设备处于震动状态，则取消震动，需要 root 权限。

**返回值**: `undefined`

**调用示例**:
```javascript
device.cancelVibration();
console.log("震动已取消");
```

## 完整示例

```javascript
// 获取设备基本信息
console.log("=== 设备信息 ===");
console.log("分辨率: " + device.width + "x" + device.height);
console.log("SDK版本: " + device.sdkInt);
console.log("CPU架构: " + device.cpuAbi);
console.log("品牌: " + device.brand);
console.log("型号: " + device.model);

// 获取设备标识
console.log("\n=== 设备标识 ===");
console.log("IMEI: " + device.getImei());
console.log("Android ID: " + device.getAndroidId());
console.log("WIFI MAC: " + device.getWifiMac());
console.log("IP地址: " + device.getIp());

// 音量控制
console.log("\n=== 音量控制 ===");
console.log("当前媒体音量: " + device.getMusicVolume());
console.log("媒体音量最大值: " + device.getMusicMaxVolume());
device.setMusicVolume(50);
console.log("媒体音量已设置为50");

// 电池信息
console.log("\n=== 电池信息 ===");
console.log("电量: " + device.getBattery() + "%");
console.log("电池状态: " + device.getBatteryStatus());

// 屏幕控制
console.log("\n=== 屏幕控制 ===");
if (device.isScreenOn()) {
    console.log("屏幕已点亮");
}
if (device.isScreenUnlock()) {
    console.log("屏幕已解锁");
}

// 震动
console.log("\n=== 震动 ===");
device.vibrate(1000);
```
