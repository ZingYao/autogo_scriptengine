# device 模块

## 模块简介

device 模块提供了设备信息获取和设置功能，包括屏幕、音量、电池等信息。

## 方法列表

### device.width
设备分辨率宽度

**参数：**

无参数

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| width | number | 屏幕宽度（像素） |

**使用示例：**
```javascript
// 获取屏幕宽度
var width = device.width;
console.log("屏幕宽度: " + width);
```

---

### device.height
设备分辨率高度

**参数：**

无参数

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| height | number | 屏幕高度（像素） |

**使用示例：**
```javascript
// 获取屏幕高度
var height = device.height;
console.log("屏幕高度: " + height);
```

---

### device.sdkInt
安卓系统API版本

**参数：**

无参数

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| sdkInt | number | 安卓系统API版本 |

**使用示例：**
```javascript
// 获取安卓系统API版本
var sdkInt = device.sdkInt();
console.log("API版本: " + sdkInt);
```

---

### device.cpuAbi
设备的CPU架构

**参数：**

无参数

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| cpuAbi | string | 设备的CPU架构 |

**使用示例：**
```javascript
// 获取设备的CPU架构
var cpuAbi = device.cpuAbi();
console.log("CPU架构: " + cpuAbi);
```

***

### device.buildId
设备构建ID

**参数：**

无参数

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| buildId | string | 设备构建ID |

**使用示例：**
```javascript
// 获取设备构建ID
var buildId = device.buildId();
console.log("构建ID: " + buildId);
```

***

### device.broad
设备主板

**参数：**

无参数

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| broad | string | 设备主板 |

**使用示例：**
```javascript
// 获取设备主板
var broad = device.broad();
console.log("主板: " + broad);
```

***

### device.brand
设备品牌

**参数：**

无参数

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| brand | string | 设备品牌 |

**使用示例：**
```javascript
// 获取设备品牌
var brand = device.brand();
console.log("品牌: " + brand);
```

***

### device.deviceName
设备名称

**参数：**

无参数

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| deviceName | string | 设备名称 |

**使用示例：**
```javascript
// 获取设备名称
var deviceName = device.deviceName();
console.log("设备名称: " + deviceName);
```

***

### device.model
设备型号

**参数：**

无参数

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| model | string | 设备型号 |

**使用示例：**
```javascript
// 获取设备型号
var model = device.model();
console.log("设备型号: " + model);
```

***

### device.product
设备产品名称

**参数：**

无参数

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| product | string | 设备产品名称 |

**使用示例：**
```javascript
// 获取设备产品名称
var product = device.product();
console.log("产品名称: " + product);
```

***

### device.bootloader
设备引导加载程序

**参数：**

无参数

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| bootloader | string | 设备引导加载程序 |

**使用示例：**
```javascript
// 获取设备引导加载程序
var bootloader = device.bootloader();
console.log("引导加载程序: " + bootloader);
```

***

### device.hardware
设备硬件信息

**参数：**

无参数

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| hardware | string | 设备硬件信息 |

**使用示例：**
```javascript
// 获取设备硬件信息
var hardware = device.hardware();
console.log("硬件信息: " + hardware);
```

***

### device.fingerprint
设备指纹

**参数：**

无参数

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| fingerprint | string | 设备指纹 |

**使用示例：**
```javascript
// 获取设备指纹
var fingerprint = device.fingerprint();
console.log("设备指纹: " + fingerprint);
```

***

### device.serial
设备序列号

**参数：**

无参数

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| serial | string | 设备序列号 |

**使用示例：**
```javascript
// 获取设备序列号
var serial = device.serial();
console.log("序列号: " + serial);
```

***

### device.incremental
设备增量版本

**参数：**

无参数

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| incremental | string | 设备增量版本 |

**使用示例：**
```javascript
// 获取设备增量版本
var incremental = device.incremental();
console.log("增量版本: " + incremental);
```

***

### device.release
设备发布版本

**参数：**

无参数

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| release | string | 设备发布版本 |

**使用示例：**
```javascript
// 获取设备发布版本
var release = device.release();
console.log("发布版本: " + release);
```

***

### device.baseOS
设备基础操作系统

**参数：**

无参数

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| baseOS | string | 设备基础操作系统 |

**使用示例：**
```javascript
// 获取设备基础操作系统
var baseOS = device.baseOS();
console.log("基础操作系统: " + baseOS);
```

***

### device.securityPatch
设备安全补丁级别

**参数：**

无参数

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| securityPatch | string | 设备安全补丁级别 |

**使用示例：**
```javascript
// 获取设备安全补丁级别
var securityPatch = device.securityPatch();
console.log("安全补丁: " + securityPatch);
```

***

### device.codename
设备代号

**参数：**

无参数

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| codename | string | 设备代号 |

**使用示例：**
```javascript
// 获取设备代号
var codename = device.codename();
console.log("设备代号: " + codename);
```

***

### device.getDisplayInfo
获取指定屏幕的分辨率信息

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| displayId | number | 屏幕ID（0表示主屏幕） |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| width | number | 屏幕宽度（像素） |
| height | number | 屏幕高度（像素） |
| dpi | number | 屏幕DPI |
| rotation | number | 屏幕旋转角度 |

**使用示例：**
```javascript
// 获取主屏幕信息
var info = device.getDisplayInfo(0);
console.log("分辨率: " + info.width + "x" + info.height);
console.log("DPI: " + info.dpi);
```

---

### device.getImei
返回设备的IMEI

**参数：**

无参数

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| imei | string | 设备的IMEI |

**使用示例：**
```javascript
// 获取设备的IMEI
var imei = device.getImei();
console.log("IMEI: " + imei);
```

---

### device.getAndroidId
返回设备的Android ID

**参数：**

无参数

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| androidId | string | 设备的Android ID |

**使用示例：**
```javascript
// 获取设备的Android ID
var androidId = device.getAndroidId();
console.log("Android ID: " + androidId);
```

---

### device.getWifiMac
获取设备WIFI-MAC

**参数：**

无参数

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| wifiMac | string | 设备WIFI-MAC |

**使用示例：**
```javascript
// 获取设备WIFI-MAC
var wifiMac = device.getWifiMac();
console.log("WIFI-MAC: " + wifiMac);
```

---

### device.getWlanMac
获取设备以太网MAC

**参数：**

无参数

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| wlanMac | string | 设备以太网MAC |

**使用示例：**
```javascript
// 获取设备以太网MAC
var wlanMac = device.getWlanMac();
console.log("以太网MAC: " + wlanMac);
```

---

### device.getNotification
获取通知

**参数：**

无参数

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| notification | object | 通知信息 |

**使用示例：**
```javascript
// 获取通知
var notification = device.getNotification();
console.log("通知: " + notification);
```

---

### device.getIp
获取设备局域网IP地址

**参数：**

无参数

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| ip | string | 设备局域网IP地址 |

**使用示例：**
```javascript
// 获取设备局域网IP地址
var ip = device.getIp();
console.log("局域网IP: " + ip);
```

---

### device.getBrightness
返回当前的(手动)亮度

**参数：**

无参数

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| brightness | number | 当前亮度值 |

**使用示例：**
```javascript
// 获取当前亮度
var brightness = device.getBrightness();
console.log("当前亮度: " + brightness);
```

---

### device.getBrightnessMode
返回当前亮度模式

**参数：**

无参数

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| mode | number | 当前亮度模式 |

**使用示例：**
```javascript
// 获取当前亮度模式
var mode = device.getBrightnessMode();
console.log("亮度模式: " + mode);
```

---

### device.getMusicVolume
返回当前媒体音量

**参数：**

无参数

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| volume | number | 当前媒体音量 |

**使用示例：**
```javascript
// 获取当前媒体音量
var volume = device.getMusicVolume();
console.log("媒体音量: " + volume);
```

---

### device.getNotificationVolume
返回当前通知音量

**参数：**

无参数

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| volume | number | 当前通知音量 |

**使用示例：**
```javascript
// 获取当前通知音量
var volume = device.getNotificationVolume();
console.log("通知音量: " + volume);
```

---

### device.getAlarmVolume
返回当前闹钟音量

**参数：**

无参数

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| volume | number | 当前闹钟音量 |

**使用示例：**
```javascript
// 获取当前闹钟音量
var volume = device.getAlarmVolume();
console.log("闹钟音量: " + volume);
```

---

### device.getMusicMaxVolume
返回媒体音量的最大值

**参数：**

无参数

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| maxVolume | number | 媒体音量的最大值 |

**使用示例：**
```javascript
// 获取媒体音量的最大值
var maxVolume = device.getMusicMaxVolume();
console.log("媒体音量最大值: " + maxVolume);
```

---

### device.getNotificationMaxVolume
返回通知音量的最大值

**参数：**

无参数

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| maxVolume | number | 通知音量的最大值 |

**使用示例：**
```javascript
// 获取通知音量的最大值
var maxVolume = device.getNotificationMaxVolume();
console.log("通知音量最大值: " + maxVolume);
```

---

### device.getAlarmMaxVolume
返回闹钟音量的最大值

**参数：**

无参数

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| maxVolume | number | 闹钟音量的最大值 |

**使用示例：**
```javascript
// 获取闹钟音量的最大值
var maxVolume = device.getAlarmMaxVolume();
console.log("闹钟音量最大值: " + maxVolume);
```

---

### device.setMusicVolume
设置当前媒体音量

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| volume | number | 音量值 |

**返回值：**

无返回值

**使用示例：**
```javascript
// 设置当前媒体音量
device.setMusicVolume(50);
```

---

### device.setNotificationVolume
设置当前通知音量

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| volume | number | 音量值 |

**返回值：**

无返回值

**使用示例：**
```javascript
// 设置当前通知音量
device.setNotificationVolume(50);
```

---

### device.setAlarmVolume
设置当前闹钟音量

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| volume | number | 音量值 |

**返回值：**

无返回值

**使用示例：**
```javascript
// 设置当前闹钟音量
device.setAlarmVolume(50);
```

---

### device.getBattery
返回当前电量百分比

**参数：**

无参数

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| battery | number | 电量百分比（0-100） |

**使用示例：**
```javascript
// 获取电量
var battery = device.getBattery;
console.log("电量: " + battery + "%");
```

---

### device.getBatteryStatus
获取电池状态

**参数：**

无参数

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| status | number | 电池状态 |

**使用示例：**
```javascript
// 获取电池状态
var status = device.getBatteryStatus();
console.log("电池状态: " + status);
```

---

### device.setBatteryStatus
模拟电池状态

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| status | number | 电池状态 |

**返回值：**

无返回值

**使用示例：**
```javascript
// 模拟电池状态
device.setBatteryStatus(2); // 2 表示充电中
```

---

### device.setBatteryLevel
模拟电池电量百分百

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| level | number | 电量百分比（0-100） |

**返回值：**

无返回值

**使用示例：**
```javascript
// 模拟电池电量百分百
device.setBatteryLevel(80); // 设置为 80%
```

---

### device.getTotalMem
返回设备内存总量

**参数：**

无参数

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| totalMem | number | 设备内存总量（MB） |

**使用示例：**
```javascript
// 获取设备内存总量
var totalMem = device.getTotalMem();
console.log("内存总量: " + totalMem + " MB");
```

---

### device.getAvailMem
返回设备当前可用的内存

**参数：**

无参数

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| availMem | number | 设备当前可用的内存（MB） |

**使用示例：**
```javascript
// 获取设备当前可用的内存
var availMem = device.getAvailMem();
console.log("可用内存: " + availMem + " MB");
```

---

### device.isScreenOn
返回设备屏幕是否是亮着的

**参数：**

无参数

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| isOn | boolean | 屏幕是否亮着 |

**使用示例：**
```javascript
// 检查屏幕是否亮着
var isOn = device.isScreenOn();
console.log("屏幕是否亮着: " + isOn);
```

---

### device.isScreenUnlock
返回屏幕锁是否已经解开

**参数：**

无参数

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| isUnlock | boolean | 屏幕锁是否已经解开 |

**使用示例：**
```javascript
// 检查屏幕锁是否已经解开
var isUnlock = device.isScreenUnlock();
console.log("屏幕锁是否解开: " + isUnlock);
```

---

### device.setDisplayPower
设置显示电源

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| on | boolean | 是否开启显示 |

**返回值：**

无返回值

**使用示例：**
```javascript
// 设置显示电源
device.setDisplayPower(true); // 开启显示
```

---

### device.wakeUp
唤醒设备

**参数：**

无参数

**返回值：**

无返回值

**使用示例：**
```javascript
// 唤醒设备
device.wakeUp();
```

---

### device.reboot
重启设备

**参数：**

无参数

**返回值：**

无返回值

**使用示例：**
```javascript
// 重启设备
device.reboot();
```

---

### device.keepScreenOn
保持屏幕常亮

**参数：**

无参数

**返回值：**

无返回值

**使用示例：**
```javascript
// 保持屏幕常亮
device.keepScreenOn();
```

---

### device.vibrate
使设备震动一段时间

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| duration | number | 震动持续时间（毫秒） |

**返回值：**

无返回值

**使用示例：**
```javascript
// 使设备震动
device.vibrate(1000); // 震动1秒
```

---

### device.cancelVibration
如果设备处于震动状态，则取消震动

**参数：**

无参数

**返回值：**

无返回值

**使用示例：**
```javascript
// 取消震动
device.cancelVibration();
```

---

## 综合使用示例

### 示例1：获取设备信息
```javascript
console.log("屏幕宽度: " + device.width);
console.log("屏幕高度: " + device.height);
console.log("电量: " + device.getBattery() + "%");
```