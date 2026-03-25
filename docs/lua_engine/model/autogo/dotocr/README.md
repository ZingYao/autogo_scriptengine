# dotocr 模块

## 模块简介

dotocr 模块提供了 OCR（光学字符识别）功能，支持从屏幕、图像文件、Base64 编码的图像等多种来源进行文字识别。该模块适用于自动化脚本中需要识别屏幕文字的场景。

## 方法列表

### dotocr.setDict
设置字库

**使用示例：**
```lua
-- 调用 dotocr.setDict 方法
dotocr.setDict();
```

---

### dotocr.ocr
从屏幕指定区域进行OCR识别

**使用示例：**
```lua
-- 调用 dotocr.ocr 方法
dotocr.ocr();
```

---

### dotocr.ocrFromImage
从图像进行OCR识别

**使用示例：**
```lua
-- 调用 dotocr.ocrFromImage 方法
dotocr.ocrFromImage();
```

---

### dotocr.ocrFromBase64
从Base64编码的图像字符串进行OCR识别

**使用示例：**
```lua
-- 调用 dotocr.ocrFromBase64 方法
dotocr.ocrFromBase64();
```

---

### dotocr.ocrFromPath
从图像文件路径进行OCR识别

**使用示例：**
```lua
-- 调用 dotocr.ocrFromPath 方法
dotocr.ocrFromPath();
```

---

### dotocr.findStr
在屏幕指定区域中查找指定字符串的位置

**使用示例：**
```lua
-- 调用 dotocr.findStr 方法
dotocr.findStr();
```

---

### dotocr.findStrFromImage
在图像中查找指定字符串的位置

**使用示例：**
```lua
-- 调用 dotocr.findStrFromImage 方法
dotocr.findStrFromImage();
```

---

### dotocr.findStrFromBase64
在Base64编码的图像中查找指定字符串的位置

**使用示例：**
```lua
-- 调用 dotocr.findStrFromBase64 方法
dotocr.findStrFromBase64();
```

---

### dotocr.findStrFromPath
在图像文件中查找指定字符串的位置

**使用示例：**
```lua
-- 调用 dotocr.findStrFromPath 方法
dotocr.findStrFromPath();
```

---

## 综合使用示例

### 示例1：识别屏幕文字
```lua
var text = dotocr.ocr(0, 0, 500, 200, "FFFFFF-000000", 5, 5, 0.9, 0, "default", 0);
console.log("识别结果: " + text);
```