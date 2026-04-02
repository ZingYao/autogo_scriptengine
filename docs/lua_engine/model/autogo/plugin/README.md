# plugin 模块

## 模块简介

plugin 模块提供了加载外部 APK 插件的功能，允许动态加载和执行外部代码。

## 方法列表

### plugin.loadApk
加载外部APK

**参数：**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| path | string | 是 | APK文件的完整路径 |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| plugin | userdata | 加载的插件对象（失败时为nil） |

**使用示例：**
```lua
-- 调用 plugin.loadApk 方法
local plugin = plugin.loadApk("/sdcard/Download/myplugin.apk")
if plugin then
    print("APK加载成功")
else
    print("APK加载失败")
end
```

---

## 综合使用示例

### 示例1：基本使用
```lua
-- plugin 模块的基本使用示例
```