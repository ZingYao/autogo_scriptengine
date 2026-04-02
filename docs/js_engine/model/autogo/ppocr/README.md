# ppocr 模块

## 模块简介

ppocr 模块提供了基于 PaddleOCR 的文字识别功能，支持高精度的中文和英文识别。

## 方法列表

### ppocr.new
创建一个新的PPOCR实例

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| version | string | OCR 模型版本 |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| ocr | object | PPOCR 实例 |

**使用示例：**
```javascript
// 创建 PPOCR 实例
var ocr = ppocr.new("v3");
```

---

### ppocr.ocr
识别屏幕文字

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| x1 | number | 区域左上角 X 坐标 |
| y1 | number | 区域左上角 Y 坐标 |
| x2 | number | 区域右下角 X 坐标 |
| y2 | number | 区域右下角 Y 坐标 |
| colorStr | string | 颜色字符串（可选） |
| displayId | number | 显示设备 ID（可选，默认为0） |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| results | array | 识别结果数组，每个元素包含 text, x, y, w, h 属性 |

**使用示例：**
```javascript
// 识别屏幕文字
var results = ppocr.ocr(0, 0, 1920, 1080);
for (var i = 0; i < results.length; i++) {
    console.log("文字: " + results[i].text + ", 位置: (" + results[i].x + ", " + results[i].y + ")");
}
```

---

### ppocr.ocrFromImage
识别图片文字

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| img | object | 图片对象 |
| colorStr | string | 颜色字符串（可选） |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| results | array | 识别结果数组，每个元素包含 text, x, y, w, h 属性 |

**使用示例：**
```javascript
// 识别图片文字
var img = images.captureScreen(0, 0, 1920, 1080);
var results = ppocr.ocrFromImage(img);
```

---

### ppocr.ocrFromBase64
识别Base64图片文字

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| b64 | string | Base64 编码的图片数据 |
| colorStr | string | 颜色字符串（可选） |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| results | array | 识别结果数组，每个元素包含 text, x, y, w, h 属性 |

**使用示例：**
```javascript
// 识别 Base64 图片文字
var results = ppocr.ocrFromBase64(base64String);
```

---

### ppocr.ocrFromPath
识别文件图片文字

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| path | string | 图片文件路径 |
| colorStr | string | 颜色字符串（可选） |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| results | array | 识别结果数组，每个元素包含 text, x, y, w, h 属性 |

**使用示例：**
```javascript
// 识别文件图片文字
var results = ppocr.ocrFromPath("/sdcard/test.png");
```

---

### ppocr.close
关闭PPOCR实例

**参数：**

无参数

**返回值：**

无返回值

**使用示例：**
```javascript
// 关闭 PPOCR 实例
ppocr.close();
```

---

## 综合使用示例

### 示例1：创建实例并识别屏幕文字
```javascript
// 创建 PPOCR 实例
var ocr = ppocr.new("v3");

// 识别屏幕文字
var results = ocr.ocr(0, 0, 1920, 1080);
for (var i = 0; i < results.length; i++) {
    console.log("文字: " + results[i].text + ", 位置: (" + results[i].x + ", " + results[i].y + ")");
}

// 关闭实例
ocr.close();
```

### 示例2：识别图片文字
```javascript
// 创建 PPOCR 实例
var ocr = ppocr.new("v3");

// 截取屏幕
var img = images.captureScreen(0, 0, 1920, 1080);

// 识别图片文字
var results = ocr.ocrFromImage(img);
for (var i = 0; i < results.length; i++) {
    console.log("文字: " + results[i].text);
}

// 关闭实例
ocr.close();
```