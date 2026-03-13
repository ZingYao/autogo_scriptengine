# OpenCV 模块

OpenCV 模块提供了基于 OpenCV 的图像识别功能，支持在屏幕或图像中查找匹配的图片模板。

## 方法列表

### opencv.findImage(x1, y1, x2, y2, template, isGray, scalingFactor, sim, displayId)
在指定区域内查找匹配的图片模板。

**入参：**
- `x1`: 区域左上角 X 坐标（整数）
- `y1`: 区域左上角 Y 坐标（整数）
- `x2`: 区域右下角 X 坐标（整数）
- `y2`: 区域右下角 Y 坐标（整数）
- `template`: 模板图片数据（字符串路径或字节数组）
- `isGray`: 是否使用灰度模式（布尔值）
- `scalingFactor`: 缩放因子（浮点数）
- `sim`: 相似度阈值（浮点数，0-1）
- `displayId`: 显示器 ID（整数，可选，默认 0）

**出参：** 包含 x 和 y 坐标的对象，未找到时为 (-1, -1)

**调用示例：**
```javascript
// 在屏幕区域查找图片
var result = opencv.findImage(0, 0, 1080, 1920, "/sdcard/template.png", false, 1.0, 0.8, 0);
if (result.x !== -1) {
    console.println("找到图片，位置: (" + result.x + ", " + result.y + ")");
    motion.tap(result.x, result.y);
} else {
    console.println("未找到图片");
}

// 使用字节数组作为模板
var templateData = files.readBytes("/sdcard/template.png");
var result2 = opencv.findImage(0, 0, 1080, 1920, templateData, false, 1.0, 0.8, 0);
```

## 完整示例

```javascript
// 示例1：基本图片查找
function findImageDemo() {
    var templatePath = "/sdcard/template.png";
    var result = opencv.findImage(0, 0, 1080, 1920, templatePath, false, 1.0, 0.8, 0);
    
    if (result.x !== -1 && result.y !== -1) {
        console.println("找到图片，位置: (" + result.x + ", " + result.y + ")");
        motion.tap(result.x, result.y);
    } else {
        console.println("未找到图片");
    }
}

// 示例2：使用字节数组查找
function findImageWithBytes() {
    var templatePath = "/sdcard/template.png";
    var templateData = files.readBytes(templatePath);
    
    var result = opencv.findImage(0, 0, 1080, 1920, templateData, false, 1.0, 0.8, 0);
    
    if (result.x !== -1 && result.y !== -1) {
        console.println("找到图片，位置: (" + result.x + ", " + result.y + ")");
    } else {
        console.println("未找到图片");
    }
}

// 示例3：使用灰度模式查找
function findImageGray() {
    var templatePath = "/sdcard/template.png";
    var result = opencv.findImage(0, 0, 1080, 1920, templatePath, true, 1.0, 0.8, 0);
    
    if (result.x !== -1 && result.y !== -1) {
        console.println("灰度模式找到图片，位置: (" + result.x + ", " + result.y + ")");
    } else {
        console.println("未找到图片");
    }
}

// 示例4：调整相似度阈值
function findImageWithSimilarity() {
    var templatePath = "/sdcard/template.png";
    
    // 高相似度要求
    var result1 = opencv.findImage(0, 0, 1080, 1920, templatePath, false, 1.0, 0.95, 0);
    console.println("高相似度结果:", result1.x, result1.y);
    
    // 低相似度要求
    var result2 = opencv.findImage(0, 0, 1080, 1920, templatePath, false, 1.0, 0.7, 0);
    console.println("低相似度结果:", result2.x, result2.y);
}

// 示例5：使用缩放因子
function findImageWithScaling() {
    var templatePath = "/sdcard/template.png";
    
    // 不缩放
    var result1 = opencv.findImage(0, 0, 1080, 1920, templatePath, false, 1.0, 0.8, 0);
    console.println("不缩放结果:", result1.x, result1.y);
    
    // 缩放 0.5 倍
    var result2 = opencv.findImage(0, 0, 1080, 1920, templatePath, false, 0.5, 0.8, 0);
    console.println("缩放0.5倍结果:", result2.x, result2.y);
}

// 示例6：在指定区域查找
function findImageInRegion() {
    var templatePath = "/sdcard/template.png";
    var x1 = 100;
    var y1 = 100;
    var x2 = 500;
    var y2 = 500;
    
    var result = opencv.findImage(x1, y1, x2, y2, templatePath, false, 1.0, 0.8, 0);
    
    if (result.x !== -1 && result.y !== -1) {
        console.println("在指定区域找到图片，位置: (" + result.x + ", " + result.y + ")");
    } else {
        console.println("在指定区域未找到图片");
    }
}

// 示例7：循环查找多个模板
function findMultipleImages() {
    var templates = [
        "/sdcard/template1.png",
        "/sdcard/template2.png",
        "/sdcard/template3.png"
    ];
    
    for (var i = 0; i < templates.length; i++) {
        var result = opencv.findImage(0, 0, 1080, 1920, templates[i], false, 1.0, 0.8, 0);
        
        if (result.x !== -1 && result.y !== -1) {
            console.println("找到模板 " + (i + 1) + "，位置: (" + result.x + ", " + result.y + ")");
            motion.tap(result.x, result.y);
            break;
        }
    }
}

// 示例8：带重试的查找
function findImageWithRetry(templatePath, maxRetries) {
    for (var i = 0; i < maxRetries; i++) {
        var result = opencv.findImage(0, 0, 1080, 1920, templatePath, false, 1.0, 0.8, 0);
        
        if (result.x !== -1 && result.y !== -1) {
            console.println("找到图片，位置: (" + result.x + ", " + result.y + ")");
            return result;
        }
        
        console.println("第 " + (i + 1) + " 次查找失败，重试...");
        utils.sleep(1000);
    }
    
    console.println("查找失败，已重试 " + maxRetries + " 次");
    return { x: -1, y: -1 };
}

// 示例9：自动点击找到的图片
function autoClickImage(templatePath) {
    var result = opencv.findImage(0, 0, 1080, 1920, templatePath, false, 1.0, 0.8, 0);
    
    if (result.x !== -1 && result.y !== -1) {
        console.println("找到图片，点击位置: (" + result.x + ", " + result.y + ")");
        motion.tap(result.x, result.y);
        return true;
    } else {
        console.println("未找到图片");
        return false;
    }
}

// 调用示例
findImageDemo();
findImageWithBytes();
findImageGray();
findImageWithSimilarity();
findImageWithScaling();
findImageInRegion();
findMultipleImages();
findImageWithRetry("/sdcard/template.png", 5);
autoClickImage("/sdcard/template.png");
```

## 注意事项

1. template 参数支持字符串路径和字节数组两种格式
2. isGray 参数设置为 true 可以提高查找速度，但可能降低准确率
3. scalingFactor 参数用于调整模板大小，1.0 表示原始大小
4. sim 参数范围是 0-1，值越大要求越严格，建议设置为 0.7-0.9
5. displayId 用于多显示器环境，通常为 0 表示主显示器
6. 查找图片时，如果未找到会返回 (-1, -1)
7. 对于大屏幕或复杂图像，查找可能需要较长时间
8. 建议使用较小的模板图片以提高查找速度
9. 灰度模式可以减少计算量，适合颜色不重要的场景
10. 在实际使用中，建议根据具体情况调整相似度阈值和缩放因子
