# PPOCR 模块

PPOCR 模块提供了基于 PaddleOCR 的文字识别功能，支持从屏幕、图像文件、Base64 字符串等多种来源进行文字识别。

## 方法列表

### ppocr.new(version)
创建一个新的 PPOCR 实例。

**入参：**
- `version`: OCR 模型版本（字符串）

**出参：** PPOCR 对象

**调用示例：**
```javascript
// 创建 PPOCR 实例
var ocr = ppocr.new("v3");
```

### ppocr.ocr(ocrObj, x1, y1, x2, y2, colorStr, displayId)
识别屏幕指定区域的文字。

**入参：**
- `ocrObj`: PPOCR 对象
- `x1`: 区域左上角 X 坐标（整数）
- `y1`: 区域左上角 Y 坐标（整数）
- `x2`: 区域右下角 X 坐标（整数）
- `y2`: 区域右下角 Y 坐标（整数）
- `colorStr`: 颜色过滤字符串（字符串，可选）
- `displayId`: 显示器 ID（整数，可选，默认 0）

**出参：** 识别结果数组，每个元素包含 text、x、y、w、h 字段

**调用示例：**
```javascript
// 创建 PPOCR 实例
var ocr = ppocr.new("v3");

// 识别屏幕区域
var results = ppocr.ocr(ocr, 0, 0, 1080, 1920, "", 0);
for (var i = 0; i < results.length; i++) {
    var item = results[i];
    console.println("文字:", item.text, "位置:", item.x, item.y);
}
```

### ppocr.ocrFromImage(ocrObj, img, colorStr)
从图像对象识别文字。

**入参：**
- `ocrObj`: PPOCR 对象
- `img`: 图像对象（NRGBA 格式）
- `colorStr`: 颜色过滤字符串（字符串，可选）

**出参：** 识别结果数组，每个元素包含 text、x、y、w、h 字段

**调用示例：**
```javascript
// 创建 PPOCR 实例
var ocr = ppocr.new("v3");

// 从图像识别
var img = images.captureScreen();
var results = ppocr.ocrFromImage(ocr, img, "");
for (var i = 0; i < results.length; i++) {
    var item = results[i];
    console.println("文字:", item.text);
}
```

### ppocr.ocrFromBase64(ocrObj, b64, colorStr)
从 Base64 编码的图像字符串识别文字。

**入参：**
- `ocrObj`: PPOCR 对象
- `b64`: Base64 编码的图像字符串
- `colorStr`: 颜色过滤字符串（字符串，可选）

**出参：** 识别结果数组，每个元素包含 text、x、y、w、h 字段

**调用示例：**
```javascript
// 创建 PPOCR 实例
var ocr = ppocr.new("v3");

// 从 Base64 图像识别
var base64Img = "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVR42mP8/5+hHgAHggJ/PchI7wAAAABJRU5ErkJggg==";
var results = ppocr.ocrFromBase64(ocr, base64Img, "");
for (var i = 0; i < results.length; i++) {
    console.println("文字:", results[i].text);
}
```

### ppocr.ocrFromPath(ocrObj, path, colorStr)
从图像文件路径识别文字。

**入参：**
- `ocrObj`: PPOCR 对象
- `path`: 图像文件路径（字符串）
- `colorStr`: 颜色过滤字符串（字符串，可选）

**出参：** 识别结果数组，每个元素包含 text、x、y、w、h 字段

**调用示例：**
```javascript
// 创建 PPOCR 实例
var ocr = ppocr.new("v3");

// 从文件识别
var results = ppocr.ocrFromPath(ocr, "/sdcard/image.png", "");
for (var i = 0; i < results.length; i++) {
    console.println("文字:", results[i].text);
}
```

### ppocr.close(ocrObj)
关闭 PPOCR 实例，释放资源。

**入参：**
- `ocrObj`: PPOCR 对象

**出参：** 无

**调用示例：**
```javascript
// 关闭 PPOCR 实例
ppocr.close(ocr);
```

## 完整示例

```javascript
// 示例1：基本文字识别
function basicOCR() {
    var ocr = ppocr.new("v3");
    
    var results = ppocr.ocr(ocr, 0, 0, 1080, 1920, "", 0);
    console.println("识别到 " + results.length + " 处文字");
    
    for (var i = 0; i < results.length; i++) {
        var item = results[i];
        console.println("文字:", item.text, "位置:", item.x, item.y, "大小:", item.w, item.h);
    }
    
    ppocr.close(ocr);
}

// 示例2：从图像文件识别
function ocrFromFile() {
    var ocr = ppocr.new("v3");
    var imagePath = "/sdcard/screenshot.png";
    
    var results = ppocr.ocrFromPath(ocr, imagePath, "");
    console.println("文件中识别到 " + results.length + " 处文字");
    
    for (var i = 0; i < results.length; i++) {
        console.println("文字:", results[i].text);
    }
    
    ppocr.close(ocr);
}

// 示例3：从屏幕截图识别
function ocrFromScreenshot() {
    var ocr = ppocr.new("v3");
    var img = images.captureScreen();
    
    var results = ppocr.ocrFromImage(ocr, img, "");
    console.println("截图中识别到 " + results.length + " 处文字");
    
    for (var i = 0; i < results.length; i++) {
        console.println("文字:", results[i].text);
    }
    
    ppocr.close(ocr);
}

// 示例4：查找并点击文字
function findAndClickText(targetText) {
    var ocr = ppocr.new("v3");
    
    var results = ppocr.ocr(ocr, 0, 0, 1080, 1920, "", 0);
    
    for (var i = 0; i < results.length; i++) {
        var item = results[i];
        if (item.text === targetText) {
            console.println("找到目标文字:", targetText);
            var centerX = item.x + item.w / 2;
            var centerY = item.y + item.h / 2;
            motion.tap(centerX, centerY);
            ppocr.close(ocr);
            return true;
        }
    }
    
    console.println("未找到目标文字:", targetText);
    ppocr.close(ocr);
    return false;
}

// 示例5：提取所有文字
function extractAllText() {
    var ocr = ppocr.new("v3");
    
    var results = ppocr.ocr(ocr, 0, 0, 1080, 1920, "", 0);
    var allText = "";
    
    for (var i = 0; i < results.length; i++) {
        allText += results[i].text + " ";
    }
    
    console.println("提取的文字:", allText);
    ppocr.close(ocr);
    return allText;
}

// 示例6：按位置过滤文字
function filterTextByRegion(x1, y1, x2, y2) {
    var ocr = ppocr.new("v3");
    
    var results = ppocr.ocr(ocr, x1, y1, x2, y2, "", 0);
    console.println("指定区域内识别到 " + results.length + " 处文字");
    
    for (var i = 0; i < results.length; i++) {
        console.println("文字:", results[i].text);
    }
    
    ppocr.close(ocr);
}

// 示例7：批量处理多个图像
function batchOCR(imagePaths) {
    var ocr = ppocr.new("v3");
    
    for (var i = 0; i < imagePaths.length; i++) {
        var path = imagePaths[i];
        console.println("处理图像:", path);
        
        var results = ppocr.ocrFromPath(ocr, path, "");
        console.println("识别到 " + results.length + " 处文字");
        
        for (var j = 0; j < results.length; j++) {
            console.println("文字:", results[j].text);
        }
    }
    
    ppocr.close(ocr);
}

// 调用示例
basicOCR();
ocrFromFile();
ocrFromScreenshot();
findAndClickText("开始");
extractAllText();
filterTextByRegion(100, 100, 500, 500);
batchOCR(["/sdcard/img1.png", "/sdcard/img2.png", "/sdcard/img3.png"]);
```

## 注意事项

1. 使用 PPOCR 功能前，需要先创建实例
2. 使用完毕后，建议调用 `close()` 方法释放资源
3. 识别结果包含文字内容、位置和大小信息
4. colorStr 参数用于颜色过滤，可以指定只识别特定颜色的文字
5. displayId 用于多显示器环境，通常为 0 表示主显示器
6. 识别速度取决于图像大小和复杂度
7. 对于大图像，建议先裁剪到感兴趣的区域再识别
8. 识别准确率受图像质量、文字清晰度等因素影响
9. 建议根据实际场景选择合适的 OCR 模型版本
