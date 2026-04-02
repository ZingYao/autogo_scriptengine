# yolo 模块

## 模块简介

yolo 模块提供了 YOLO（You Only Look Once）目标检测功能，用于实时对象检测。

## 方法列表

### yolo.new
创建一个新的YOLO实例

**参数：**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| version | string | 是 | YOLO模型版本（如 "yolov5n", "yolov5s" 等） |
| cpuThreadNum | int | 是 | CPU线程数 |
| paramPath | string | 是 | 参数文件路径 |
| binPath | string | 是 | 模型文件路径 |
| labels | string | 是 | 标签文件路径 |

**返回值：**

| 类型 | 说明 |
|------|------|
| Yolo | YOLO实例对象，创建失败会抛出异常 |

**使用示例：**
```javascript
// 创建YOLO实例
let yolo = yolo.new("yolov5n", 4, "/sdcard/yolo/yolov5n.param", "/sdcard/yolo/yolov5n.bin", "/sdcard/yolo/coco.names");

// 使用不同的模型版本
let yolo2 = yolo.new("yolov5s", 8, "/sdcard/yolo/yolov5s.param", "/sdcard/yolo/yolov5s.bin", "/sdcard/yolo/coco.names");

// 使用自定义线程数
let yolo3 = yolo.new("yolov5n", 2, "/sdcard/yolo/yolov5n.param", "/sdcard/yolo/yolov5n.bin", "/sdcard/yolo/custom.labels");
```

---

### yolo.detect
检测屏幕上的对象

**参数：**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| y | Yolo | 是 | YOLO实例 |
| x1 | int | 是 | 检测区域左上角X坐标 |
| y1 | int | 是 | 检测区域左上角Y坐标 |
| x2 | int | 是 | 检测区域右下角X坐标 |
| y2 | int | 是 | 检测区域右下角Y坐标 |
| displayId | int | 否 | 显示设备ID（可选，默认0） |

**返回值：**

| 类型 | 说明 |
|------|------|
| Array | 检测结果数组，每个元素包含 x, y, w, h, label, score 属性 |

**检测结果对象属性：**

| 属性名 | 类型 | 说明 |
|--------|------|------|
| x | number | 检测框左上角X坐标 |
| y | number | 检测框左上角Y坐标 |
| w | number | 检测框宽度 |
| h | number | 检测框高度 |
| label | string | 检测对象的标签 |
| score | number | 检测置信度（0-1之间） |

**使用示例：**
```javascript
// 创建YOLO实例
let yolo = yolo.new("yolov5n", 4, "/sdcard/yolo/yolov5n.param", "/sdcard/yolo/yolov5n.bin", "/sdcard/yolo/coco.names");

// 检测整个屏幕
let results = yolo.detect(yolo, 0, 0, 1920, 1080, 0);
for (let i = 0; i < results.length; i++) {
    let result = results[i];
    console.log("对象 " + i + ":");
    console.log("  标签: " + result.label);
    console.log("  位置: (" + result.x + ", " + result.y + ")");
    console.log("  大小: " + result.w + "x" + result.h);
    console.log("  置信度: " + result.score);
}

// 检测屏幕左半部分
let leftResults = yolo.detect(yolo, 0, 0, 960, 1080, 0);

// 检测屏幕右半部分
let rightResults = yolo.detect(yolo, 960, 0, 1920, 1080, 0);

// 检测特定区域
let regionResults = yolo.detect(yolo, 100, 100, 500, 500, 0);
```

---

### yolo.detectFromImage
检测图片中的对象

**参数：**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| y | Yolo | 是 | YOLO实例 |
| img | image | 是 | 图片对象（image.NRGBA类型） |

**返回值：**

| 类型 | 说明 |
|------|------|
| Array | 检测结果数组，每个元素包含 x, y, w, h, label, score 属性 |

**使用示例：**
```javascript
// 创建YOLO实例
let yolo = yolo.new("yolov5n", 4, "/sdcard/yolo/yolov5n.param", "/sdcard/yolo/yolov5n.bin", "/sdcard/yolo/coco.names");

// 截取屏幕
let img = images.captureScreen(0, 0, 1920, 1080);

// 检测图片中的对象
let results = yolo.detectFromImage(yolo, img);
for (let i = 0; i < results.length; i++) {
    let result = results[i];
    console.log("对象 " + i + ": " + result.label + " (" + result.score + ")");
}

// 处理检测结果
if (results.length > 0) {
    console.log("检测到 " + results.length + " 个对象");
    for (let i = 0; i < results.length; i++) {
        let result = results[i];
        // 在检测到的对象位置绘制矩形
        images.drawRect(img, result.x, result.y, result.w, result.h, "#FF0000", 2);
    }
}
```

---

### yolo.detectFromBase64
检测Base64图片中的对象

**参数：**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| y | Yolo | 是 | YOLO实例 |
| b64 | string | 是 | Base64编码的图片数据 |
| colorStr | string | 否 | 颜色字符串（可选，用于调试） |

**返回值：**

| 类型 | 说明 |
|------|------|
| Array | 检测结果数组，每个元素包含 x, y, w, h, label, score 属性 |

**使用示例：**
```javascript
// 创建YOLO实例
let yolo = yolo.new("yolov5n", 4, "/sdcard/yolo/yolov5n.param", "/sdcard/yolo/yolov5n.bin", "/sdcard/yolo/coco.names");

// 从文件读取图片并转换为Base64
let img = images.read("/sdcard/test.jpg");
let base64String = images.toBase64(img);

// 检测Base64图片中的对象
let results = yolo.detectFromBase64(yolo, base64String, "");
for (let i = 0; i < results.length; i++) {
    let result = results[i];
    console.log("对象 " + i + ": " + result.label);
}

// 使用网络图片的Base64数据
let networkBase64 = "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVR42mP8/5+hHgAHggJ/PchI7wAAAABJRU5ErkJggg==";
let networkResults = yolo.detectFromBase64(yolo, networkBase64, "#FF0000");
```

---

### yolo.detectFromPath
检测文件图片中的对象

**参数：**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| y | Yolo | 是 | YOLO实例 |
| path | string | 是 | 图片文件路径 |
| colorStr | string | 否 | 颜色字符串（可选，用于调试） |

**返回值：**

| 类型 | 说明 |
|------|------|
| Array | 检测结果数组，每个元素包含 x, y, w, h, label, score 属性 |

**使用示例：**
```javascript
// 创建YOLO实例
let yolo = yolo.new("yolov5n", 4, "/sdcard/yolo/yolov5n.param", "/sdcard/yolo/yolov5n.bin", "/sdcard/yolo/coco.names");

// 检测文件图片中的对象
let results = yolo.detectFromPath(yolo, "/sdcard/test.jpg", "");
for (let i = 0; i < results.length; i++) {
    let result = results[i];
    console.log("对象 " + i + ":");
    console.log("  标签: " + result.label);
    console.log("  位置: (" + result.x + ", " + result.y + ")");
    console.log("  置信度: " + result.score);
}

// 检测多张图片
let imagePaths = [
    "/sdcard/image1.jpg",
    "/sdcard/image2.jpg",
    "/sdcard/image3.jpg"
];

for (let j = 0; j < imagePaths.length; j++) {
    let path = imagePaths[j];
    let imgResults = yolo.detectFromPath(yolo, path, "#00FF00");
    console.log("图片 " + path + " 检测到 " + imgResults.length + " 个对象");
}
```

---

### yolo.close
关闭YOLO实例

**参数：**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| y | Yolo | 是 | YOLO实例 |

**返回值：**

| 类型 | 说明 |
|------|------|
| undefined | 无返回值 |

**使用示例：**
```javascript
// 创建YOLO实例
let yolo = yolo.new("yolov5n", 4, "/sdcard/yolo/yolov5n.param", "/sdcard/yolo/yolov5n.bin", "/sdcard/yolo/coco.names");

// 使用YOLO进行检测
let results = yolo.detect(yolo, 0, 0, 1920, 1080, 0);
console.log("检测到 " + results.length + " 个对象");

// 使用完毕后关闭实例
yolo.close(yolo);

// 注意：关闭后不能再使用该实例
// let newResults = yolo.detect(yolo, 0, 0, 1920, 1080, 0); // 这会出错
```

---

## 综合使用示例

### 示例1：创建实例并检测屏幕对象
```javascript
// 创建YOLO实例
let yolo = yolo.new("yolov5n", 4, "/sdcard/yolo/yolov5n.param", "/sdcard/yolo/yolov5n.bin", "/sdcard/yolo/coco.names");

// 检测屏幕上的对象
let results = yolo.detect(yolo, 0, 0, 1920, 1080, 0);
for (let i = 0; i < results.length; i++) {
    let result = results[i];
    console.log("对象 " + i + ": " + result.label + " (" + result.score + ")");
}

// 关闭实例
yolo.close(yolo);
```

### 示例2：检测图片中的对象
```javascript
// 创建YOLO实例
let yolo = yolo.new("yolov5n", 4, "/sdcard/yolo/yolov5n.param", "/sdcard/yolo/yolov5n.bin", "/sdcard/yolo/coco.names");

// 截取屏幕
let img = images.captureScreen(0, 0, 1920, 1080);

// 检测图片中的对象
let results = yolo.detectFromImage(yolo, img);
for (let i = 0; i < results.length; i++) {
    let result = results[i];
    console.log("对象 " + i + ": " + result.label);
}

// 关闭实例
yolo.close(yolo);
```

### 示例3：检测文件图片中的对象
```javascript
// 创建YOLO实例
let yolo = yolo.new("yolov5n", 4, "/sdcard/yolo/yolov5n.param", "/sdcard/yolo/yolov5n.bin", "/sdcard/yolo/coco.names");

// 检测文件图片中的对象
let results = yolo.detectFromPath(yolo, "/sdcard/test.jpg", "");
for (let i = 0; i < results.length; i++) {
    let result = results[i];
    console.log("对象 " + i + ": " + result.label + " 位置: (" + result.x + ", " + result.y + ")");
}

// 关闭实例
yolo.close(yolo);
```

### 示例4：实时屏幕检测
```javascript
// 创建YOLO实例
let yolo = yolo.new("yolov5n", 4, "/sdcard/yolo/yolov5n.param", "/sdcard/yolo/yolov5n.bin", "/sdcard/yolo/coco.names");

// 持续检测屏幕
for (let i = 0; i < 10; i++) {
    // 检测屏幕上的对象
    let results = yolo.detect(yolo, 0, 0, 1920, 1080, 0);

    console.log("第 " + (i + 1) + " 次检测，发现 " + results.length + " 个对象");

    // 处理检测结果
    for (let j = 0; j < results.length; j++) {
        let result = results[j];
        if (result.score > 0.8) {
            console.log("高置信度对象: " + result.label + " (" + result.score + ")");
        }
    }

    // 等待1秒
    utils.sleep(1000);
}

// 关闭实例
yolo.close(yolo);
```

### 示例5：批量处理图片
```javascript
// 创建YOLO实例
let yolo = yolo.new("yolov5n", 4, "/sdcard/yolo/yolov5n.param", "/sdcard/yolo/yolov5n.bin", "/sdcard/yolo/coco.names");

// 定义要处理的图片列表
let imageList = [
    "/sdcard/images/img1.jpg",
    "/sdcard/images/img2.jpg",
    "/sdcard/images/img3.jpg",
    "/sdcard/images/img4.jpg",
    "/sdcard/images/img5.jpg"
];

// 统计检测结果
let totalObjects = 0;
let labelCounts = {};

// 批量处理图片
for (let i = 0; i < imageList.length; i++) {
    let path = imageList[i];
    console.log("处理图片: " + path);

    // 检测图片中的对象
    let results = yolo.detectFromPath(yolo, path, "");

    // 统计结果
    totalObjects += results.length;
    for (let j = 0; j < results.length; j++) {
        let label = results[j].label;
        if (labelCounts[label] === undefined) {
            labelCounts[label] = 0;
        }
        labelCounts[label]++;
    }

    console.log("  检测到 " + results.length + " 个对象");
}

// 输出统计结果
console.log("总计检测到 " + totalObjects + " 个对象");
console.log("标签统计:");
for (let label in labelCounts) {
    console.log("  " + label + ": " + labelCounts[label] + " 个");
}

// 关闭实例
yolo.close(yolo);
```

### 示例6：基于检测结果进行操作
```javascript
// 创建YOLO实例
let yolo = yolo.new("yolov5n", 4, "/sdcard/yolo/yolov5n.param", "/sdcard/yolo/yolov5n.bin", "/sdcard/yolo/coco.names");

// 检测屏幕上的对象
let results = yolo.detect(yolo, 0, 0, 1920, 1080, 0);

// 查找特定对象
let targetLabel = "person";
let found = false;

for (let i = 0; i < results.length; i++) {
    let result = results[i];

    // 检查是否为目标对象且置信度足够高
    if (result.label === targetLabel && result.score > 0.7) {
        console.log("找到目标对象:");
        console.log("  位置: (" + result.x + ", " + result.y + ")");
        console.log("  大小: " + result.w + "x" + result.h);
        console.log("  置信度: " + result.score);

        // 计算中心点
        let centerX = result.x + result.w / 2;
        let centerY = result.y + result.h / 2;
        console.log("  中心点: (" + centerX + ", " + centerY + ")");

        // 在中心点点击
        touch.tap(centerX, centerY);

        found = true;
        break;
    }
}

if (!found) {
    console.log("未找到目标对象: " + targetLabel);
}

// 关闭实例
yolo.close(yolo);
```
