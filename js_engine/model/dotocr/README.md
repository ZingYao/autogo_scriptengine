# DotOCR 模块

DotOCR 模块提供了基于点阵的 OCR（光学字符识别）功能，支持从屏幕、图像文件、Base64 字符串等多种来源进行文字识别和查找。

## 方法列表

### dotocr.setDict(name, dict)
设置字库，用于后续的 OCR 识别。

**入参：**
- `name`: 字库名称（字符串）
- `dict`: 字库内容（字符串）

**出参：** 无

**调用示例：**
```javascript
// 设置字库
dotocr.setDict("my_dict", "字库内容...");
```

### dotocr.ocr(x1, y1, x2, y2, threshold, colGap, rowGap, sim, mode, dictName, displayId)
从屏幕指定区域进行 OCR 识别。

**入参：**
- `x1`: 区域左上角 X 坐标（整数）
- `y1`: 区域左上角 Y 坐标（整数）
- `x2`: 区域右下角 X 坐标（整数）
- `y2`: 区域右下角 Y 坐标（整数）
- `threshold`: 二值化阈值（字符串）
- `colGap`: 列间距（整数）
- `rowGap`: 行间距（整数）
- `sim`: 相似度阈值（浮点数，0-1）
- `mode`: 识别模式（整数）
- `dictName`: 字库名称（字符串）
- `displayId`: 显示器 ID（整数）

**出参：** 识别到的文本字符串

**调用示例：**
```javascript
// 识别屏幕区域 (0, 0, 500, 300) 内的文字
var text = dotocr.ocr(0, 0, 500, 300, "200", 5, 5, 0.8, 0, "my_dict", 0);
console.println(text);
```

### dotocr.ocrFromImage(img, threshold, colGap, rowGap, sim, mode, dictName)
从图像对象进行 OCR 识别。

**入参：**
- `img`: 图像对象（NRGBA 格式）
- `threshold`: 二值化阈值（字符串）
- `colGap`: 列间距（整数）
- `rowGap`: 行间距（整数）
- `sim`: 相似度阈值（浮点数，0-1）
- `mode`: 识别模式（整数）
- `dictName`: 字库名称（字符串）

**出参：** 识别到的文本字符串

**调用示例：**
```javascript
// 从图像对象识别文字
var img = images.captureScreen();
var text = dotocr.ocrFromImage(img, "200", 5, 5, 0.8, 0, "my_dict");
console.println(text);
```

### dotocr.ocrFromBase64(b64, threshold, colGap, rowGap, sim, mode, dictName)
从 Base64 编码的图像字符串进行 OCR 识别。

**入参：**
- `b64`: Base64 编码的图像字符串
- `threshold`: 二值化阈值（字符串）
- `colGap`: 列间距（整数）
- `rowGap`: 行间距（整数）
- `sim`: 相似度阈值（浮点数，0-1）
- `mode`: 识别模式（整数）
- `dictName`: 字库名称（字符串）

**出参：** 识别到的文本字符串

**调用示例：**
```javascript
// 从 Base64 图像识别文字
var base64Img = "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVR42mP8/5+hHgAHggJ/PchI7wAAAABJRU5ErkJggg==";
var text = dotocr.ocrFromBase64(base64Img, "200", 5, 5, 0.8, 0, "my_dict");
console.println(text);
```

### dotocr.ocrFromPath(path, threshold, colGap, rowGap, sim, mode, dictName)
从图像文件路径进行 OCR 识别。

**入参：**
- `path`: 图像文件路径（字符串）
- `threshold`: 二值化阈值（字符串）
- `colGap`: 列间距（整数）
- `rowGap`: 行间距（整数）
- `sim`: 相似度阈值（浮点数，0-1）
- `mode`: 识别模式（整数）
- `dictName`: 字库名称（字符串）

**出参：** 识别到的文本字符串

**调用示例：**
```javascript
// 从图像文件识别文字
var text = dotocr.ocrFromPath("/sdcard/image.png", "200", 5, 5, 0.8, 0, "my_dict");
console.println(text);
```

### dotocr.findStr(x1, y1, x2, y2, text, threshold, colGap, rowGap, sim, dictName, displayId)
在屏幕指定区域中查找指定字符串的位置。

**入参：**
- `x1`: 区域左上角 X 坐标（整数）
- `y1`: 区域左上角 Y 坐标（整数）
- `x2`: 区域右下角 X 坐标（整数）
- `y2`: 区域右下角 Y 坐标（整数）
- `text`: 要查找的文本（字符串）
- `threshold`: 二值化阈值（字符串）
- `colGap`: 列间距（整数）
- `rowGap`: 行间距（整数）
- `sim`: 相似度阈值（浮点数，0-1）
- `dictName`: 字库名称（字符串）
- `displayId`: 显示器 ID（整数）

**出参：** 包含 x 和 y 坐标的对象，未找到时为 (-1, -1)

**调用示例：**
```javascript
// 在屏幕区域查找文字
var pos = dotocr.findStr(0, 0, 500, 300, "开始", "200", 5, 5, 0.8, "my_dict", 0);
if (pos.x !== -1) {
    console.println("找到文字，位置: (" + pos.x + ", " + pos.y + ")");
    motion.tap(pos.x, pos.y);
} else {
    console.println("未找到文字");
}
```

### dotocr.findStrFromImage(img, text, threshold, colGap, rowGap, sim, dictName)
在图像中查找指定字符串的位置。

**入参：**
- `img`: 图像对象（NRGBA 格式）
- `text`: 要查找的文本（字符串）
- `threshold`: 二值化阈值（字符串）
- `colGap`: 列间距（整数）
- `rowGap`: 行间距（整数）
- `sim`: 相似度阈值（浮点数，0-1）
- `dictName`: 字库名称（字符串）

**出参：** 包含 x 和 y 坐标的对象，未找到时为 (-1, -1)

**调用示例：**
```javascript
// 在图像中查找文字
var img = images.captureScreen();
var pos = dotocr.findStrFromImage(img, "开始", "200", 5, 5, 0.8, "my_dict");
if (pos.x !== -1) {
    console.println("找到文字，位置: (" + pos.x + ", " + pos.y + ")");
}
```

### dotocr.findStrFromBase64(b64, text, threshold, colGap, rowGap, sim, dictName)
在 Base64 编码的图像中查找指定字符串的位置。

**入参：**
- `b64`: Base64 编码的图像字符串
- `text`: 要查找的文本（字符串）
- `threshold`: 二值化阈值（字符串）
- `colGap`: 列间距（整数）
- `rowGap`: 行间距（整数）
- `sim`: 相似度阈值（浮点数，0-1）
- `dictName`: 字库名称（字符串）

**出参：** 包含 x 和 y 坐标的对象，未找到时为 (-1, -1)

**调用示例：**
```javascript
// 在 Base64 图像中查找文字
var base64Img = "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVR42mP8/5+hHgAHggJ/PchI7wAAAABJRU5ErkJggg==";
var pos = dotocr.findStrFromBase64(base64Img, "开始", "200", 5, 5, 0.8, "my_dict");
if (pos.x !== -1) {
    console.println("找到文字，位置: (" + pos.x + ", " + pos.y + ")");
}
```

### dotocr.findStrFromPath(path, text, threshold, colGap, rowGap, sim, dictName)
在图像文件中查找指定字符串的位置。

**入参：**
- `path`: 图像文件路径（字符串）
- `text`: 要查找的文本（字符串）
- `threshold`: 二值化阈值（字符串）
- `colGap`: 列间距（整数）
- `rowGap`: 行间距（整数）
- `sim`: 相似度阈值（浮点数，0-1）
- `dictName`: 字库名称（字符串）

**出参：** 包含 x 和 y 坐标的对象，未找到时为 (-1, -1)

**调用示例：**
```javascript
// 在图像文件中查找文字
var pos = dotocr.findStrFromPath("/sdcard/image.png", "开始", "200", 5, 5, 0.8, "my_dict");
if (pos.x !== -1) {
    console.println("找到文字，位置: (" + pos.x + ", " + pos.y + ")");
}
```

## 完整示例

```javascript
// 设置字库
dotocr.setDict("my_dict", "字库内容...");

// 示例1：识别屏幕上的文字
var screenText = dotocr.ocr(0, 0, 1080, 1920, "200", 5, 5, 0.8, 0, "my_dict", 0);
console.println("识别到的文字: " + screenText);

// 示例2：查找并点击屏幕上的文字
var pos = dotocr.findStr(0, 0, 1080, 1920, "开始游戏", "200", 5, 5, 0.8, "my_dict", 0);
if (pos.x !== -1 && pos.y !== -1) {
    console.println("找到文字，点击位置: (" + pos.x + ", " + pos.y + ")");
    motion.tap(pos.x, pos.y);
} else {
    console.println("未找到指定文字");
}

// 示例3：从图像文件识别文字
var imagePath = "/sdcard/screenshot.png";
var fileText = dotocr.ocrFromPath(imagePath, "200", 5, 5, 0.8, 0, "my_dict");
console.println("文件中的文字: " + fileText);

// 示例4：从图像对象识别文字
var img = images.captureScreen();
var imgText = dotocr.ocrFromImage(img, "200", 5, 5, 0.8, 0, "my_dict");
console.println("图像中的文字: " + imgText);

// 示例5：在图像中查找文字
var img = images.captureScreen();
var findPos = dotocr.findStrFromImage(img, "确认", "200", 5, 5, 0.8, "my_dict");
if (findPos.x !== -1) {
    console.println("在图像中找到文字，位置: (" + findPos.x + ", " + findPos.y + ")");
}
```

## 注意事项

1. 使用 OCR 功能前，需要先设置合适的字库
2. 二值化阈值（threshold）需要根据图像亮度调整，通常在 100-255 之间
3. 相似度阈值（sim）范围是 0-1，值越大要求越严格
4. colGap 和 rowGap 用于控制字符之间的间距，根据实际情况调整
5. displayId 用于多显示器环境，通常为 0 表示主显示器
6. 识别模式（mode）的不同值可能影响识别效果，需要根据场景选择
7. 查找文字时，如果未找到会返回 (-1, -1)
