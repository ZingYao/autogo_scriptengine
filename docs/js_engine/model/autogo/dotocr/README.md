# dotocr 模块

## 模块简介

dotocr 模块提供了 OCR（光学字符识别）功能，支持从屏幕、图像文件、Base64 编码的图像等多种来源进行文字识别。该模块适用于自动化脚本中需要识别屏幕文字的场景。

## 方法列表

### dotocr.setDict(name, dict)
设置字库

**参数：**
- `name` (String): 字库名称
- `dict` (String): 字库内容或路径

**返回值：** undefined

**使用示例：**
```javascript
dotocr.setDict("default", "/path/to/dict.txt");
```

---

### dotocr.ocr(x1, y1, x2, y2, threshold, colGap, rowGap, sim, mode, dictName, displayId)
从屏幕指定区域进行 OCR 识别

**参数：**
- `x1` (Number): 区域左上角 X 坐标
- `y1` (Number): 区域左上角 Y 坐标
- `x2` (Number): 区域右下角 X 坐标
- `y2` (Number): 区域右下角 Y 坐标
- `threshold` (String): 颜色阈值
- `colGap` (Number): 列间隔
- `rowGap` (Number): 行间隔
- `sim` (Number): 相似度阈值 (0-1)
- `mode` (Number): 识别模式
- `dictName` (String): 使用的字库名称
- `displayId` (Number): 显示设备 ID

**返回值：** (String) 识别到的文字

**使用示例：**
```javascript
// 识别屏幕左上角区域的文字
var text = dotocr.ocr(0, 0, 500, 200, "FFFFFF-000000", 5, 5, 0.9, 0, "default", 0);
console.log("识别结果: " + text);
```

---

### dotocr.ocrFromImage(img, threshold, colGap, rowGap, sim, mode, dictName)
从图像进行 OCR 识别

**参数：**
- `img` (Image): NRGBA 格式的图像对象
- `threshold` (String): 颜色阈值
- `colGap` (Number): 列间隔
- `rowGap` (Number): 行间隔
- `sim` (Number): 相似度阈值 (0-1)
- `mode` (Number): 识别模式
- `dictName` (String): 使用的字库名称

**返回值：** (String) 识别到的文字

**使用示例：**
```javascript
var img = images.captureScreen(0, 0, 500, 200);
var text = dotocr.ocrFromImage(img, "FFFFFF-000000", 5, 5, 0.9, 0, "default");
```

---

### dotocr.ocrFromBase64(b64, threshold, colGap, rowGap, sim, mode, dictName)
从 Base64 编码的图像字符串进行 OCR 识别

**参数：**
- `b64` (String): Base64 编码的图像字符串
- `threshold` (String): 颜色阈值
- `colGap` (Number): 列间隔
- `rowGap` (Number): 行间隔
- `sim` (Number): 相似度阈值 (0-1)
- `mode` (Number): 识别模式
- `dictName` (String): 使用的字库名称

**返回值：** (String) 识别到的文字

**使用示例：**
```javascript
var base64Str = images.encodeToBase64(img, "png", 100);
var text = dotocr.ocrFromBase64(base64Str, "FFFFFF-000000", 5, 5, 0.9, 0, "default");
```

---

### dotocr.ocrFromPath(path, threshold, colGap, rowGap, sim, mode, dictName)
从图像文件路径进行 OCR 识别

**参数：**
- `path` (String): 图像文件路径
- `threshold` (String): 颜色阈值
- `colGap` (Number): 列间隔
- `rowGap` (Number): 行间隔
- `sim` (Number): 相似度阈值 (0-1)
- `mode` (Number): 识别模式
- `dictName` (String): 使用的字库名称

**返回值：** (String) 识别到的文字

**使用示例：**
```javascript
var text = dotocr.ocrFromPath("/sdcard/screenshot.png", "FFFFFF-000000", 5, 5, 0.9, 0, "default");
```

---

### dotocr.findStr(x1, y1, x2, y2, text, threshold, colGap, rowGap, sim, dictName, displayId)
在屏幕指定区域中查找指定字符串的位置

**参数：**
- `x1` (Number): 区域左上角 X 坐标
- `y1` (Number): 区域左上角 Y 坐标
- `x2` (Number): 区域右下角 X 坐标
- `y2` (Number): 区域右下角 Y 坐标
- `text` (String): 要查找的文字
- `threshold` (String): 颜色阈值
- `colGap` (Number): 列间隔
- `rowGap` (Number): 行间隔
- `sim` (Number): 相似度阈值 (0-1)
- `dictName` (String): 使用的字库名称
- `displayId` (Number): 显示设备 ID

**返回值：** (Object) 包含 x 和 y 坐标的对象，如果未找到则返回 (-1, -1)

**使用示例：**
```javascript
var pos = dotocr.findStr(0, 0, 1000, 500, "确定", "FFFFFF-000000", 5, 5, 0.9, "default", 0);
if (pos.x !== -1) {
    click(pos.x, pos.y);
}
```

---

### dotocr.findStrFromImage(img, text, threshold, colGap, rowGap, sim, dictName)
在图像中查找指定字符串的位置

**参数：**
- `img` (Image): NRGBA 格式的图像对象
- `text` (String): 要查找的文字
- `threshold` (String): 颜色阈值
- `colGap` (Number): 列间隔
- `rowGap` (Number): 行间隔
- `sim` (Number): 相似度阈值 (0-1)
- `dictName` (String): 使用的字库名称

**返回值：** (Object) 包含 x 和 y 坐标的对象

**使用示例：**
```javascript
var img = images.captureScreen(0, 0, 1000, 500);
var pos = dotocr.findStrFromImage(img, "登录", "FFFFFF-000000", 5, 5, 0.9, "default");
```

---

### dotocr.findStrFromBase64(b64, text, threshold, colGap, rowGap, sim, dictName)
在 Base64 编码的图像中查找指定字符串的位置

**参数：**
- `b64` (String): Base64 编码的图像字符串
- `text` (String): 要查找的文字
- `threshold` (String): 颜色阈值
- `colGap` (Number): 列间隔
- `rowGap` (Number): 行间隔
- `sim` (Number): 相似度阈值 (0-1)
- `dictName` (String): 使用的字库名称

**返回值：** (Object) 包含 x 和 y 坐标的对象

**使用示例：**
```javascript
var pos = dotocr.findStrFromBase64(base64Str, "确定", "FFFFFF-000000", 5, 5, 0.9, "default");
```

---

### dotocr.findStrFromPath(path, text, threshold, colGap, rowGap, sim, dictName)
在图像文件中查找指定字符串的位置

**参数：**
- `path` (String): 图像文件路径
- `text` (String): 要查找的文字
- `threshold` (String): 颜色阈值
- `colGap` (Number): 列间隔
- `rowGap` (Number): 行间隔
- `sim` (Number): 相似度阈值 (0-1)
- `dictName` (String): 使用的字库名称

**返回值：** (Object) 包含 x 和 y 坐标的对象

**使用示例：**
```javascript
var pos = dotocr.findStrFromPath("/sdcard/screenshot.png", "登录", "FFFFFF-000000", 5, 5, 0.9, "default");
```

---

## 综合使用示例

### 示例1：自动点击屏幕上的文字
```javascript
// 查找并点击"确定"按钮
function clickText(text, x1, y1, x2, y2) {
    var pos = dotocr.findStr(x1, y1, x2, y2, text, "FFFFFF-000000", 5, 5, 0.9, "default", 0);
    if (pos.x !== -1) {
        click(pos.x, pos.y);
        return true;
    }
    return false;
}

// 点击屏幕上的"确定"
if (clickText("确定", 0, 0, device.width, device.height)) {
    console.log("点击成功");
} else {
    console.log("未找到文字");
}
```

### 示例2：识别并处理屏幕文字
```javascript
// 识别屏幕顶部状态栏的文字
var text = dotocr.ocr(0, 0, device.width, 100, "FFFFFF-000000", 3, 3, 0.95, 0, "default", 0);
console.log("状态栏文字: " + text);

// 判断是否包含特定文字
if (text.indexOf("WiFi") !== -1) {
    console.log("WiFi已连接");
}
```

### 示例3：批量识别图片文件
```javascript
var imageFiles = ["/sdcard/1.png", "/sdcard/2.png", "/sdcard/3.png"];

for (var i = 0; i < imageFiles.length; i++) {
    var text = dotocr.ocrFromPath(imageFiles[i], "000000-FFFFFF", 5, 5, 0.9, 0, "default");
    console.log("图片 " + imageFiles[i] + " 识别结果: " + text);
}
```
