# media 模块

## 模块简介

media 模块提供了媒体相关的功能，包括音频播放、文件扫描、短信发送等。

## 方法列表

### media.scanFile
扫描路径path的媒体文件

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| path | string | 要扫描的路径 |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| files | array | 媒体文件列表 |

**使用示例：**
```javascript
// 调用 media.scanFile 方法
media.scanFile("/sdcard/Music");
```

---

### media.playMP3
播放MP3文件

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| path | string | MP3 文件路径 |

**返回值：**

无返回值

**使用示例：**
```javascript
// 调用 media.playMP3 方法
media.playMP3("/sdcard/Music/test.mp3");
```

---

### media.sendSMS
发送短信

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| number | string | 手机号码 |
| message | string | 短信内容 |

**返回值：**

无返回值

**使用示例：**
```javascript
// 调用 media.sendSMS 方法
media.sendSMS("13800138000", "Hello World");
```

---

## 综合使用示例

### 示例1：基本使用
```javascript
// media 模块的基本使用示例
```