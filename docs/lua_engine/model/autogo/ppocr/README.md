# ppocr 模块

## 模块简介

ppocr 模块提供了基于 PaddleOCR 的文字识别功能，支持高精度的中文和英文识别。

## 方法列表

### ppocr.new
创建一个新的PPOCR实例

**参数：**
- `version` (string): OCR 模型版本

**返回值：**
- Ppocr 对象: PPOCR 实例

**使用示例：**
```lua
-- 创建 PPOCR 实例
local ocr = ppocr.new("v3")
```

---

### ppocr.ocr
识别屏幕文字

**参数：**
- `p` (Ppocr): PPOCR 实例
- `x1` (int): 区域左上角 X 坐标
- `y1` (int): 区域左上角 Y 坐标
- `x2` (int): 区域右下角 X 坐标
- `y2` (int): 区域右下角 Y 坐标
- `colorStr` (string): 颜色字符串（可选）
- `displayId` (int): 显示设备 ID（可选，默认 0）

**返回值：**
- Array: 识别结果数组，每个元素包含 text, confidence 属性

**使用示例：**
```lua
-- 创建 PPOCR 实例
local ocr = ppocr.new("v3")

-- 识别屏幕文字
local results = ocr.ocr(0, 0, 1920, 1080, "", 0)
for i, result in ipairs(results) do
    print("文字: " .. result.text .. ", 置信度: " .. result.confidence)
end
```

---

### ppocr.ocrFromImage
识别图片文字

**参数：**
- `p` (Ppocr): PPOCR 实例
- `img` (image): 图片对象
- `colorStr` (string): 颜色字符串（可选）

**返回值：**
- Array: 识别结果数组，每个元素包含 text, confidence 属性

**使用示例：**
```lua
-- 创建 PPOCR 实例
local ocr = ppocr.new("v3")

-- 识别图片文字
local img = images.captureScreen(0, 0, 1920, 1080)
local results = ocr.ocrFromImage(img, "")
for i, result in ipairs(results) do
    print("文字: " .. result.text)
end
```

---

### ppocr.ocrFromBase64
识别Base64图片文字

**参数：**
- `p` (Ppocr): PPOCR 实例
- `b64` (string): Base64 编码的图片数据
- `colorStr` (string): 颜色字符串（可选）

**返回值：**
- Array: 识别结果数组，每个元素包含 text, confidence 属性

**使用示例：**
```lua
-- 创建 PPOCR 实例
local ocr = ppocr.new("v3")

-- 识别 Base64 图片文字
local results = ocr.ocrFromBase64(base64String, "")
for i, result in ipairs(results) do
    print("文字: " .. result.text)
end
```

---

### ppocr.ocrFromPath
识别文件图片文字

**参数：**
- `p` (Ppocr): PPOCR 实例
- `path` (string): 图片文件路径
- `colorStr` (string): 颜色字符串（可选）

**返回值：**
- Array: 识别结果数组，每个元素包含 text, confidence 属性

**使用示例：**
```lua
-- 创建 PPOCR 实例
local ocr = ppocr.new("v3")

-- 识别文件图片文字
local results = ocr.ocrFromPath("/sdcard/test.png", "")
for i, result in ipairs(results) do
    print("文字: " .. result.text)
end
```

---

### ppocr.close
关闭PPOCR实例

**参数：**
- `p` (Ppocr): PPOCR 实例

**使用示例：**
```lua
-- 创建 PPOCR 实例
local ocr = ppocr.new("v3")

-- 使用完毕后关闭实例
ocr.close()
```

---

## 综合使用示例

### 示例1：创建实例并识别屏幕文字
```lua
-- 创建 PPOCR 实例
local ocr = ppocr.new("v3")

-- 识别屏幕文字
local results = ocr.ocr(0, 0, 1920, 1080, "", 0)
for i, result in ipairs(results) do
    print("文字: " .. result.text .. ", 置信度: " .. result.confidence)
end

-- 关闭实例
ocr.close()
```

### 示例2：识别图片文字
```lua
-- 创建 PPOCR 实例
local ocr = ppocr.new("v3")

-- 截取屏幕
local img = images.captureScreen(0, 0, 1920, 1080)

-- 识别图片文字
local results = ocr.ocrFromImage(img, "")
for i, result in ipairs(results) do
    print("文字: " .. result.text)
end

-- 关闭实例
ocr.close()
```