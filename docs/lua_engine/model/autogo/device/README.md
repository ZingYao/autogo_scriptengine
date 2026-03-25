# device 模块

## 模块简介

device 模块提供了设备信息获取和设置功能，包括屏幕、音量、电池等信息。

## 方法列表

### device.width
设备分辨率宽度

**使用示例：**
```lua
-- 调用 device.width 方法
device.width();
```

---

### device.height
设备分辨率高度

**使用示例：**
```lua
-- 调用 device.height 方法
device.height();
```

---

### device.sdkInt
安卓系统API版本

**使用示例：**
```lua
-- 调用 device.sdkInt 方法
device.sdkInt();
```

---

### device.cpuAbi
设备的CPU架构

**使用示例：**
```lua
-- 调用 device.cpuAbi 方法
device.cpuAbi();
```

---

### device.getImei
返回设备的IMEI

**使用示例：**
```lua
-- 调用 device.getImei 方法
device.getImei();
```

---

### device.getAndroidId
返回设备的Android ID

**使用示例：**
```lua
-- 调用 device.getAndroidId 方法
device.getAndroidId();
```

---

### device.getWifiMac
获取设备WIFI-MAC

**使用示例：**
```lua
-- 调用 device.getWifiMac 方法
device.getWifiMac();
```

---

### device.getWlanMac
获取设备以太网MAC

**使用示例：**
```lua
-- 调用 device.getWlanMac 方法
device.getWlanMac();
```

---

### device.getIp
获取设备局域网IP地址

**使用示例：**
```lua
-- 调用 device.getIp 方法
device.getIp();
```

---

### device.getNotification
获取当前所有通知消息

**使用示例：**
```lua
-- 调用 device.getNotification 方法
device.getNotification();
```

---

### device.getBrightness
返回当前的(手动)亮度

**使用示例：**
```lua
-- 调用 device.getBrightness 方法
device.getBrightness();
```

---

### device.getBrightnessMode
返回当前亮度模式

**使用示例：**
```lua
-- 调用 device.getBrightnessMode 方法
device.getBrightnessMode();
```

---

### device.getMusicVolume
返回当前媒体音量

**使用示例：**
```lua
-- 调用 device.getMusicVolume 方法
device.getMusicVolume();
```

---

### device.getNotificationVolume
返回当前通知音量

**使用示例：**
```lua
-- 调用 device.getNotificationVolume 方法
device.getNotificationVolume();
```

---

### device.getAlarmVolume
返回当前闹钟音量

**使用示例：**
```lua
-- 调用 device.getAlarmVolume 方法
device.getAlarmVolume();
```

---

### device.getMusicMaxVolume
返回媒体音量的最大值

**使用示例：**
```lua
-- 调用 device.getMusicMaxVolume 方法
device.getMusicMaxVolume();
```

---

### device.getNotificationMaxVolume
返回通知音量的最大值

**使用示例：**
```lua
-- 调用 device.getNotificationMaxVolume 方法
device.getNotificationMaxVolume();
```

---

### device.getAlarmMaxVolume
返回闹钟音量的最大值

**使用示例：**
```lua
-- 调用 device.getAlarmMaxVolume 方法
device.getAlarmMaxVolume();
```

---

### device.setMusicVolume
设置当前媒体音量

**使用示例：**
```lua
-- 调用 device.setMusicVolume 方法
device.setMusicVolume();
```

---

### device.setNotificationVolume
设置当前通知音量

**使用示例：**
```lua
-- 调用 device.setNotificationVolume 方法
device.setNotificationVolume();
```

---

### device.setAlarmVolume
设置当前闹钟音量

**使用示例：**
```lua
-- 调用 device.setAlarmVolume 方法
device.setAlarmVolume();
```

---

### device.getBattery
返回当前电量百分比

**使用示例：**
```lua
-- 调用 device.getBattery 方法
device.getBattery();
```

---

### device.getBatteryStatus
获取电池状态

**使用示例：**
```lua
-- 调用 device.getBatteryStatus 方法
device.getBatteryStatus();
```

---

### device.setBatteryStatus
模拟电池状态

**使用示例：**
```lua
-- 调用 device.setBatteryStatus 方法
device.setBatteryStatus();
```

---

### device.setBatteryLevel
模拟电池电量百分百

**使用示例：**
```lua
-- 调用 device.setBatteryLevel 方法
device.setBatteryLevel();
```

---

### device.getTotalMem
返回设备内存总量

**使用示例：**
```lua
-- 调用 device.getTotalMem 方法
device.getTotalMem();
```

---

### device.getAvailMem
返回设备当前可用的内存

**使用示例：**
```lua
-- 调用 device.getAvailMem 方法
device.getAvailMem();
```

---

### device.isScreenOn
返回设备屏幕是否是亮着的

**使用示例：**
```lua
-- 调用 device.isScreenOn 方法
device.isScreenOn();
```

---

### device.isScreenUnlock
返回屏幕锁是否已经解开

**使用示例：**
```lua
-- 调用 device.isScreenUnlock 方法
device.isScreenUnlock();
```

---

### device.wakeUp
唤醒设备

**使用示例：**
```lua
-- 调用 device.wakeUp 方法
device.wakeUp();
```

---

### device.keepScreenOn
保持屏幕常亮

**使用示例：**
```lua
-- 调用 device.keepScreenOn 方法
device.keepScreenOn();
```

---

### device.vibrate
使设备震动一段时间

**使用示例：**
```lua
-- 调用 device.vibrate 方法
device.vibrate();
```

---

### device.cancelVibration
如果设备处于震动状态，则取消震动

**使用示例：**
```lua
-- 调用 device.cancelVibration 方法
device.cancelVibration();
```

---

### device.setDisplayPower
设置屏幕电源模式

**使用示例：**
```lua
-- 调用 device.setDisplayPower 方法
device.setDisplayPower();
```

---

### device.reboot
重启设备

**使用示例：**
```lua
-- 调用 device.reboot 方法
device.reboot();
```

---

## 综合使用示例

### 示例1：获取设备信息
```lua
console.log("屏幕宽度: " + device.width);
console.log("屏幕高度: " + device.height);
console.log("电量: " + device.getBattery() + "%");
```