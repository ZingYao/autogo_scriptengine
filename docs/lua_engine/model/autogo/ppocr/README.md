# ppocr 模块

## 模块简介

ppocr 模块提供了基于 PaddleOCR 的文字识别功能，支持高精度的中文和英文识别。

## 方法列表

### ppocr.new
创建Ppocr对象

**使用示例：**
```lua
-- 调用 ppocr.new 方法
ppocr.new();
```

---

### ppocr.ocr
识别屏幕文字

**使用示例：**
```lua
-- 调用 ppocr.ocr 方法
ppocr.ocr();
```

---

### ppocr.ocrFromImage
识别图片文字

**使用示例：**
```lua
-- 调用 ppocr.ocrFromImage 方法
ppocr.ocrFromImage();
```

---

### ppocr.ocrFromBase64
识别Base64图片文字

**使用示例：**
```lua
-- 调用 ppocr.ocrFromBase64 方法
ppocr.ocrFromBase64();
```

---

### ppocr.ocrFromPath
识别文件图片文字

**使用示例：**
```lua
-- 调用 ppocr.ocrFromPath 方法
ppocr.ocrFromPath();
```

---

### ppocr.close
关闭Ppocr对象

**使用示例：**
```lua
-- 调用 ppocr.close 方法
ppocr.close();
```

---

## 综合使用示例

### 示例1：基本使用
```lua
-- ppocr 模块的基本使用示例
```