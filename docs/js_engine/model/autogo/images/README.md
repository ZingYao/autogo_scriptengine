# images 模块

## 模块简介

images 模块提供了图像处理功能，包括屏幕截图、颜色识别、图像变换（裁剪、缩放、旋转等）以及图像格式转换等功能。

## 方法列表

### images.pixel
获取指定坐标的像素颜色

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| x | number | X 坐标 |
| y | number | Y 坐标 |
| displayId | number | 显示设备 ID（可选，默认为 0） |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| color | string | 颜色字符串（格式："#RRGGBB"） |

**使用示例：**
```javascript
// 获取指定坐标的像素颜色
var color = images.pixel(100, 100);
console.log("颜色: " + color);
```

---

### images.setCallback
设置回调函数

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| callback | function | 回调函数，接收图像和显示设备 ID 参数 |

**返回值：**

无返回值

**使用示例：**
```javascript
// 设置回调函数
images.setCallback(function(img, displayId) {
    console.log("收到图像，显示设备 ID: " + displayId);
});
```

---

### images.captureScreen
截取屏幕

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| x1 | number | 区域左上角 X 坐标 |
| y1 | number | 区域左上角 Y 坐标 |
| x2 | number | 区域右下角 X 坐标 |
| y2 | number | 区域右下角 Y 坐标 |
| displayId | number | 显示设备 ID（可选，默认为 0） |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| image | object | NRGBA 格式的图像对象 |

**使用示例：**
```javascript
// 截取屏幕
var img = images.captureScreen(0, 0, 500, 500);
```

---

### images.cmpColor
比较颜色

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| x | number | X 坐标 |
| y | number | Y 坐标 |
| colorStr | string | 颜色字符串（格式："#RRGGBB"） |
| sim | number | 相似度阈值 (0-1) |
| displayId | number | 显示设备 ID（可选，默认为 0） |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| match | boolean | 是否匹配 |

**使用示例：**
```javascript
// 比较颜色
var match = images.cmpColor(100, 100, "#FF0000", 0.9);
console.log("颜色匹配: " + match);
```

---

### images.findColor
查找颜色

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| x1 | number | 区域左上角 X 坐标 |
| y1 | number | 区域左上角 Y 坐标 |
| x2 | number | 区域右下角 X 坐标 |
| y2 | number | 区域右下角 Y 坐标 |
| colorStr | string | 要查找的颜色字符串（格式："#RRGGBB"） |
| sim | number | 相似度阈值 (0-1) |
| dir | number | 搜索方向 |
| displayId | number | 显示设备 ID（可选，默认为 0） |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| x | number | 找到的 X 坐标（未找到返回 -1） |
| y | number | 找到的 Y 坐标（未找到返回 -1） |

**使用示例：**
```javascript
// 查找颜色
var pos = images.findColor(0, 0, 1000, 500, "#FF0000", 0.9, 0, 0);
if (pos.x !== -1) {
    console.log("找到颜色在: " + pos.x + ", " + pos.y);
}
```

---

### images.getColorCountInRegion
获取区域内指定颜色的数量

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| x1 | number | 区域左上角 X 坐标 |
| y1 | number | 区域左上角 Y 坐标 |
| x2 | number | 区域右下角 X 坐标 |
| y2 | number | 区域右下角 Y 坐标 |
| colorStr | string | 要统计的颜色字符串（格式："#RRGGBB"） |
| sim | number | 相似度阈值 (0-1) |
| displayId | number | 显示设备 ID（可选，默认为 0） |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| count | number | 颜色数量 |

**使用示例：**
```javascript
// 获取区域内指定颜色的数量
var count = images.getColorCountInRegion(0, 0, 500, 500, "#FF0000", 0.9, 0);
console.log("颜色数量: " + count);
```

---

### images.detectsMultiColors
检测多点颜色

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| colors | string | 多点颜色字符串 |
| sim | number | 相似度阈值 (0-1) |
| displayId | number | 显示设备 ID（可选，默认为 0） |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| result | boolean | 是否检测到 |

**使用示例：**
```javascript
// 检测多点颜色
var result = images.detectsMultiColors("#FF0000,#00FF00", 0.9, 0);
console.log("检测结果: " + result);
```

---

### images.findMultiColors
查找多点颜色

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| x1 | number | 区域左上角 X 坐标 |
| y1 | number | 区域左上角 Y 坐标 |
| x2 | number | 区域右下角 X 坐标 |
| y2 | number | 区域右下角 Y 坐标 |
| colors | string | 多点颜色字符串 |
| sim | number | 相似度阈值 (0-1) |
| dir | number | 搜索方向 |
| displayId | number | 显示设备 ID（可选，默认为 0） |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| x | number | 找到的 X 坐标（未找到返回 -1） |
| y | number | 找到的 Y 坐标（未找到返回 -1） |

**使用示例：**
```javascript
// 查找多点颜色
var pos = images.findMultiColors(0, 0, 1000, 500, "#FF0000,#00FF00", 0.9, 0, 0);
if (pos.x !== -1) {
    console.log("找到颜色在: " + pos.x + ", " + pos.y);
}
```

---

### images.readFromPath
从路径读取图片

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| path | string | 图片文件路径 |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| image | object | NRGBA 格式的图像对象 |

**使用示例：**
```javascript
// 从路径读取图片
var img = images.readFromPath("/sdcard/test.png");
```

---

### images.readFromUrl
从URL读取图片

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| url | string | 图片 URL |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| image | object | NRGBA 格式的图像对象 |

**使用示例：**
```javascript
// 从 URL 读取图片
var img = images.readFromUrl("https://example.com/image.png");
```

---

### images.readFromBase64
从Base64读取图片

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| base64Str | string | Base64 编码的图片字符串 |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| image | object | NRGBA 格式的图像对象 |

**使用示例：**
```javascript
// 从 Base64 读取图片
var img = images.readFromBase64("iVBORw0KGgoAAAANS...");
```

---

### images.readFromBytes
从字节数组读取图片

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| bytes | array | 字节数组 |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| image | object | NRGBA 格式的图像对象 |

**使用示例：**
```javascript
// 从字节数组读取图片
var bytes = [137, 80, 78, 71, 13, 10, 26, 10]; // PNG 文件头
var img = images.readFromBytes(bytes);
```

---

### images.toNrgba
转换为NRGBA格式

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| img | object | 图像对象 |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| image | object | NRGBA 格式的图像对象 |

**使用示例：**
```javascript
// 转换为 NRGBA 格式
var nrgbaImg = images.toNrgba(img);
```

---

### images.save
保存图片

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| img | object | 图像对象 |
| path | string | 保存路径 |
| quality | number | 图片质量（0-100） |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| success | boolean | 是否保存成功 |

**使用示例：**
```javascript
// 保存图片
var success = images.save(img, "/sdcard/test.png", 100);
```

---

### images.encodeToBase64
编码为Base64

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| img | object | 图像对象 |
| format | string | 图片格式（如 "png", "jpeg"） |
| quality | number | 图片质量（0-100） |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| base64Str | string | Base64 编码的图片字符串 |

**使用示例：**
```javascript
// 编码为 Base64
var base64Str = images.encodeToBase64(img, "png", 100);
```

---

### images.encodeToBytes
编码为字节数组

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| img | object | 图像对象 |
| format | string | 图片格式（如 "png", "jpeg"） |
| quality | number | 图片质量（0-100） |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| bytes | array | 字节数组 |

**使用示例：**
```javascript
// 编码为字节数组
var bytes = images.encodeToBytes(img, "png", 100);
```

---

### images.clip
裁剪图片

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| img | object | 图像对象 |
| x1 | number | 区域左上角 X 坐标 |
| y1 | number | 区域左上角 Y 坐标 |
| x2 | number | 区域右下角 X 坐标 |
| y2 | number | 区域右下角 Y 坐标 |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| image | object | 裁剪后的图像对象 |

**使用示例：**
```javascript
// 裁剪图片
var croppedImg = images.clip(img, 100, 100, 500, 500);
```

---

### images.resize
调整图片大小

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| img | object | 图像对象 |
| width | number | 目标宽度 |
| height | number | 目标高度 |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| image | object | 调整大小后的图像对象 |

**使用示例：**
```javascript
// 调整图片大小
var resizedImg = images.resize(img, 800, 600);
```

---

### images.rotate
旋转图片

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| img | object | 图像对象 |
| degree | number | 旋转角度（度数） |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| image | object | 旋转后的图像对象 |

**使用示例：**
```javascript
// 旋转图片
var rotatedImg = images.rotate(img, 90);
```

---

### images.grayscale
灰度化

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| img | object | 图像对象 |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| image | object | 灰度化后的图像对象 |

**使用示例：**
```javascript
// 灰度化图片
var grayImg = images.grayscale(img);
```

---

### images.applyThreshold
应用阈值

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| img | object | 图像对象 |
| threshold | number | 阈值 |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| image | object | 应用阈值后的图像对象 |

**使用示例：**
```javascript
// 应用阈值
var thresholdImg = images.applyThreshold(img, 128);
```

---

### images.applyAdaptiveThreshold
应用自适应阈值

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| img | object | 图像对象 |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| image | object | 应用自适应阈值后的图像对象 |

**使用示例：**
```javascript
// 应用自适应阈值
var adaptiveThresholdImg = images.applyAdaptiveThreshold(img);
```

---

### images.applyBinarization
二值化

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| img | object | 图像对象 |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| image | object | 二值化后的图像对象 |

**使用示例：**
```javascript
// 二值化图片
var binaryImg = images.applyBinarization(img);
```

---

## 综合使用示例

### 示例1：截图并保存
```javascript
var img = images.captureScreen(0, 0, device.width, device.height);
images.save(img, "/sdcard/screenshot.png", 100);
```