# device 模块

## 模块简介

device 模块提供了设备信息获取和设置功能，包括屏幕、音量、电池等信息。

## 方法列表

## 综合使用示例

### 示例1：获取设备信息
```lua
console.log("屏幕宽度: " + device.width);
console.log("屏幕高度: " + device.height);
console.log("电量: " + device.getBattery() + "%");
```