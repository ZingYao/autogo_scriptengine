# plugin 模块

## 模块简介

plugin 模块提供了加载外部 APK 插件的功能，允许动态加载和执行外部代码。

## 方法列表

### plugin.loadApk
加载外部APK

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| path | string | APK 文件路径 |

**返回值：**

无返回值

**使用示例：**
```javascript
// 加载外部 APK
plugin.loadApk("/sdcard/app.apk");
```

---

### plugin.newContext
创建Context参数

**参数：**

无参数

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| context | object | Context 对象 |

**使用示例：**
```javascript
// 创建 Context 参数
var context = plugin.newContext();
```

---

### plugin.newAssetManager
创建AssetManager参数

**参数：**

无参数

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| assetManager | object | AssetManager 对象 |

**使用示例：**
```javascript
// 创建 AssetManager 参数
var assetManager = plugin.newAssetManager();
```

---

## 综合使用示例

### 示例1：基本使用
```javascript
// plugin 模块的基本使用示例
```