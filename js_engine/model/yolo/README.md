# Yolo 模块

Yolo 模块提供了基于 YOLO（You Only Look Once）的目标检测功能，支持从屏幕、图像文件、Base64 字符串等多种来源进行对象检测。

## 方法列表

### yolo.new(version, cpuThreadNum, paramPath, binPath, labels)
创建一个新的 YOLO 实例。

**入参：**
- `version`: YOLO 模型版本（字符串）
- `cpuThreadNum`: CPU 线程数（整数）
- `paramPath`: 参数文件路径（字符串）
- `binPath`: 模型文件路径（字符串）
- `labels`: 标签文件路径（字符串）

**出参：** YOLO 对象

**调用示例：**
```javascript
// 创建 YOLO 实例
var yoloObj = yolo.new("v5", 4, "/sdcard/model.param", "/sdcard/model.bin", "/sdcard/labels.txt");
```

### yolo.detect(yoloObj, x1, y1, x2, y2, displayId)
检测屏幕指定区域内的对象。

**入参：**
- `yoloObj`: YOLO 对象
- `x1`: 区域左上角 X 坐标（整数）
- `y1`: 区域左上角 Y 坐标（整数）
- `x2`: 区域右下角 X 坐标（整数）
- `y2`: 区域右下角 Y 坐标（整数）
- `displayId`: 显示器 ID（整数，可选，默认 0）

**出参：** 检测结果数组，每个元素包含 x、y、w、h、label、score 字段

**调用示例：**
```javascript
// 创建 YOLO 实例
var yoloObj = yolo.new("v5", 4, "/sdcard/model.param", "/sdcard/model.bin", "/sdcard/labels.txt");

// 检测屏幕区域
var results = yolo.detect(yoloObj, 0, 0, 1080, 1920, 0);
for (var i = 0; i < results.length; i++) {
    var item = results[i];
    console.println("标签:", item.label, "位置:", item.x, item.y, "大小:", item.w, item.h, "置信度:", item.score);
}
```

### yolo.detectFromImage(yoloObj, img)
从图像对象检测对象。

**入参：**
- `yoloObj`: YOLO 对象
- `img`: 图像对象（NRGBA 格式）

**出参：** 检测结果数组，每个元素包含 x、y、w、h、label、score 字段

**调用示例：**
```javascript
// 创建 YOLO 实例
var yoloObj = yolo.new("v5", 4, "/sdcard/model.param", "/sdcard/model.bin", "/sdcard/labels.txt");

// 从图像检测
var img = images.captureScreen();
var results = yolo.detectFromImage(yoloObj, img);
for (var i = 0; i < results.length; i++) {
    console.println("标签:", results[i].label);
}
```

### yolo.detectFromBase64(yoloObj, b64, colorStr)
从 Base64 编码的图像字符串检测对象。

**入参：**
- `yoloObj`: YOLO 对象
- `b64`: Base64 编码的图像字符串
- `colorStr`: 颜色过滤字符串（字符串，可选）

**出参：** 检测结果数组，每个元素包含 x、y、w、h、label、score 字段

**调用示例：**
```javascript
// 创建 YOLO 实例
var yoloObj = yolo.new("v5", 4, "/sdcard/model.param", "/sdcard/model.bin", "/sdcard/labels.txt");

// 从 Base64 图像检测
var base64Img = "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVR42mP8/5+hHgAHggJ/PchI7wAAAABJRU5ErkJggg==";
var results = yolo.detectFromBase64(yoloObj, base64Img, "");
for (var i = 0; i < results.length; i++) {
    console.println("标签:", results[i].label);
}
```

### yolo.detectFromPath(yoloObj, path, colorStr)
从图像文件路径检测对象。

**入参：**
- `yoloObj`: YOLO 对象
- `path`: 图像文件路径（字符串）
- `colorStr`: 颜色过滤字符串（字符串，可选）

**出参：** 检测结果数组，每个元素包含 x、y、w、h、label、score 字段

**调用示例：**
```javascript
// 创建 YOLO 实例
var yoloObj = yolo.new("v5", 4, "/sdcard/model.param", "/sdcard/model.bin", "/sdcard/labels.txt");

// 从文件检测
var results = yolo.detectFromPath(yoloObj, "/sdcard/image.png", "");
for (var i = 0; i < results.length; i++) {
    console.println("标签:", results[i].label);
}
```

### yolo.close(yoloObj)
关闭 YOLO 实例，释放资源。

**入参：**
- `yoloObj`: YOLO 对象

**出参：** 无

**调用示例：**
```javascript
// 关闭 YOLO 实例
yolo.close(yoloObj);
```

## 完整示例

```javascript
// 示例1：基本目标检测
function basicDetection() {
    var yoloObj = yolo.new("v5", 4, "/sdcard/model.param", "/sdcard/model.bin", "/sdcard/labels.txt");
    
    var results = yolo.detect(yoloObj, 0, 0, 1080, 1920, 0);
    console.println("检测到 " + results.length + " 个对象");
    
    for (var i = 0; i < results.length; i++) {
        var item = results[i];
        console.println("标签:", item.label, "位置:", item.x, item.y, "大小:", item.w, item.h, "置信度:", item.score);
    }
    
    yolo.close(yoloObj);
}

// 示例2：从图像文件检测
function detectFromFile() {
    var yoloObj = yolo.new("v5", 4, "/sdcard/model.param", "/sdcard/model.bin", "/sdcard/labels.txt");
    var imagePath = "/sdcard/image.png";
    
    var results = yolo.detectFromPath(yoloObj, imagePath, "");
    console.println("文件中检测到 " + results.length + " 个对象");
    
    for (var i = 0; i < results.length; i++) {
        console.println("标签:", results[i].label, "置信度:", results[i].score);
    }
    
    yolo.close(yoloObj);
}

// 示例3：从屏幕截图检测
function detectFromScreenshot() {
    var yoloObj = yolo.new("v5", 4, "/sdcard/model.param", "/sdcard/model.bin", "/sdcard/labels.txt");
    var img = images.captureScreen();
    
    var results = yolo.detectFromImage(yoloObj, img);
    console.println("截图中检测到 " + results.length + " 个对象");
    
    for (var i = 0; i < results.length; i++) {
        console.println("标签:", results[i].label);
    }
    
    yolo.close(yoloObj);
}

// 示例4：查找并点击目标
function findAndClickTarget(targetLabel) {
    var yoloObj = yolo.new("v5", 4, "/sdcard/model.param", "/sdcard/model.bin", "/sdcard/labels.txt");
    
    var results = yolo.detect(yoloObj, 0, 0, 1080, 1920, 0);
    
    for (var i = 0; i < results.length; i++) {
        var item = results[i];
        if (item.label === targetLabel) {
            console.println("找到目标:", targetLabel);
            var centerX = item.x + item.w / 2;
            var centerY = item.y + item.h / 2;
            motion.tap(centerX, centerY);
            yolo.close(yoloObj);
            return true;
        }
    }
    
    console.println("未找到目标:", targetLabel);
    yolo.close(yoloObj);
    return false;
}

// 示例5：按置信度过滤
function filterByConfidence(minScore) {
    var yoloObj = yolo.new("v5", 4, "/sdcard/model.param", "/sdcard/model.bin", "/sdcard/labels.txt");
    
    var results = yolo.detect(yoloObj, 0, 0, 1080, 1920, 0);
    var filteredResults = [];
    
    for (var i = 0; i < results.length; i++) {
        var item = results[i];
        if (item.score >= minScore) {
            filteredResults.push(item);
        }
    }
    
    console.println("过滤后剩余 " + filteredResults.length + " 个对象");
    for (var i = 0; i < filteredResults.length; i++) {
        console.println("标签:", filteredResults[i].label, "置信度:", filteredResults[i].score);
    }
    
    yolo.close(yoloObj);
}

// 示例6：统计检测到的对象
function countObjects() {
    var yoloObj = yolo.new("v5", 4, "/sdcard/model.param", "/sdcard/model.bin", "/sdcard/labels.txt");
    
    var results = yolo.detect(yoloObj, 0, 0, 1080, 1920, 0);
    var counts = {};
    
    for (var i = 0; i < results.length; i++) {
        var label = results[i].label;
        if (!counts[label]) {
            counts[label] = 0;
        }
        counts[label]++;
    }
    
    console.println("对象统计:");
    for (var label in counts) {
        console.println(label + ":", counts[label]);
    }
    
    yolo.close(yoloObj);
}

// 示例7：批量处理多个图像
function batchDetect(imagePaths) {
    var yoloObj = yolo.new("v5", 4, "/sdcard/model.param", "/sdcard/model.bin", "/sdcard/labels.txt");
    
    for (var i = 0; i < imagePaths.length; i++) {
        var path = imagePaths[i];
        console.println("处理图像:", path);
        
        var results = yolo.detectFromPath(yoloObj, path, "");
        console.println("检测到 " + results.length + " 个对象");
        
        for (var j = 0; j < results.length; j++) {
            console.println("标签:", results[j].label, "置信度:", results[j].score);
        }
    }
    
    yolo.close(yoloObj);
}

// 示例8：YOLO 检测器
function YoloDetector(version, cpuThreadNum, paramPath, binPath, labelsPath) {
    this.yoloObj = yolo.new(version, cpuThreadNum, paramPath, binPath, labelsPath);
    
    this.detect = function(x1, y1, x2, y2, displayId) {
        return yolo.detect(this.yoloObj, x1, y1, x2, y2, displayId);
    };
    
    this.detectFromImage = function(img) {
        return yolo.detectFromImage(this.yoloObj, img);
    };
    
    this.detectFromPath = function(path, colorStr) {
        return yolo.detectFromPath(this.yoloObj, path, colorStr);
    };
    
    this.detectFromBase64 = function(b64, colorStr) {
        return yolo.detectFromBase64(this.yoloObj, b64, colorStr);
    };
    
    this.findByLabel = function(results, targetLabel) {
        var found = [];
        for (var i = 0; i < results.length; i++) {
            if (results[i].label === targetLabel) {
                found.push(results[i]);
            }
        }
        return found;
    };
    
    this.filterByScore = function(results, minScore) {
        var filtered = [];
        for (var i = 0; i < results.length; i++) {
            if (results[i].score >= minScore) {
                filtered.push(results[i]);
            }
        }
        return filtered;
    };
    
    this.countByLabel = function(results) {
        var counts = {};
        for (var i = 0; i < results.length; i++) {
            var label = results[i].label;
            if (!counts[label]) {
                counts[label] = 0;
            }
            counts[label]++;
        }
        return counts;
    };
    
    this.close = function() {
        yolo.close(this.yoloObj);
        this.yoloObj = null;
    };
}

// 使用 YOLO 检测器
function useYoloDetector() {
    var detector = new YoloDetector("v5", 4, "/sdcard/model.param", "/sdcard/model.bin", "/sdcard/labels.txt");
    
    // 检测屏幕
    var results = detector.detect(0, 0, 1080, 1920, 0);
    console.println("检测到 " + results.length + " 个对象");
    
    // 查找特定标签
    var persons = detector.findByLabel(results, "person");
    console.println("检测到 " + persons.length + " 个人");
    
    // 按置信度过滤
    var highConfidence = detector.filterByScore(results, 0.8);
    console.println("高置信度对象: " + highConfidence.length);
    
    // 统计对象
    var counts = detector.countByLabel(results);
    console.println("对象统计:", JSON.stringify(counts));
    
    // 关闭检测器
    detector.close();
}

// 调用示例
basicDetection();
detectFromFile();
detectFromScreenshot();
findAndClickTarget("person");
filterByConfidence(0.8);
countObjects();
batchDetect(["/sdcard/img1.png", "/sdcard/img2.png", "/sdcard/img3.png"]);
useYoloDetector();
```

## 注意事项

1. 使用 YOLO 功能前，需要先创建实例
2. 创建实例时需要提供模型文件和标签文件路径
3. cpuThreadNum 参数影响检测速度，建议根据设备性能设置
4. 使用完毕后，建议调用 `close()` 方法释放资源
5. 检测结果包含对象的位置、大小、标签和置信度
6. colorStr 参数用于颜色过滤，可以指定只检测特定颜色的对象
7. displayId 用于多显示器环境，通常为 0 表示主显示器
8. 检测速度取决于图像大小和复杂度
9. 对于大图像，建议先裁剪到感兴趣的区域再检测
10. 置信度（score）范围是 0-1，值越大表示检测结果越可靠
