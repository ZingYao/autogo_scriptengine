# Images 模块

Images 模块提供了图像处理和颜色识别功能。

## 方法列表

### images.pixel(x, y, displayId?)

获取指定坐标的像素颜色。

**参数**:
- `x` (number): X坐标
- `y` (number): Y坐标
- `displayId` (number, 可选): 显示器ID，默认为0

**返回值**: `string` - 像素颜色值

**调用示例**:
```javascript
const color = images.pixel(500, 1000, 0);
console.log("颜色: " + color);
```

### images.setCallback(callback)

设置图像回调函数。

**参数**:
- `callback` (function): 回调函数，接收 (image, displayId) 参数

**返回值**: `undefined`

**调用示例**:
```javascript
images.setCallback(function(img, displayId) {
    console.log("图像回调触发");
});
```

### images.captureScreen(x1, y1, x2, y2, displayId?)

截取屏幕。

**参数**:
- `x1` (number): 区域左上角X坐标
- `y1` (number): 区域左上角Y坐标
- `x2` (number): 区域右下角X坐标
- `y2` (number): 区域右下角Y坐标
- `displayId` (number, 可选): 显示器ID，默认为0

**返回值**: `object` - 截取的图像对象，如果失败返回 `null`

**调用示例**:
```javascript
const img = images.captureScreen(0, 0, 1080, 1920, 0);
if (img !== null) {
    console.log("屏幕截图成功");
}
```

### images.cmpColor(x, y, colorStr, sim, displayId?)

比较颜色。

**参数**:
- `x` (number): X坐标
- `y` (number): Y坐标
- `colorStr` (string): 颜色字符串
- `sim` (number): 相似度阈值 (0-1)
- `displayId` (number, 可选): 显示器ID，默认为0

**返回值**: `boolean` - 如果颜色匹配返回 `true`，否则返回 `false`

**调用示例**:
```javascript
if (images.cmpColor(500, 1000, "#FF0000", 0.9, 0)) {
    console.log("颜色匹配");
    click(500, 1000, 1, 0);
}
```

### images.findColor(x1, y1, x2, y2, colorStr, sim, dir, displayId?)

查找颜色。

**参数**:
- `x1` (number): 搜索区域左上角X坐标
- `y1` (number): 搜索区域左上角Y坐标
- `x2` (number): 搜索区域右下角X坐标
- `y2` (number): 搜索区域右下角Y坐标
- `colorStr` (string): 要查找的颜色字符串
- `sim` (number): 相似度阈值 (0-1)
- `dir` (number): 搜索方向 (0=从左到右, 1=从右到左, 2=从上到下, 3=从下到上)
- `displayId` (number, 可选): 显示器ID，默认为0

**返回值**: `number`, `number` - 找到的坐标 (x, y)，如果未找到返回 (-1, -1)

**调用示例**:
```javascript
const [x, y] = images.findColor(0, 0, 1080, 1920, "#FF0000", 0.9, 0, 0);
if (x !== -1 && y !== -1) {
    console.log("找到颜色在: " + x + ", " + y);
    click(x, y, 1, 0);
}
```

### images.getColorCountInRegion(x1, y1, x2, y2, colorStr, sim, displayId?)

获取区域内指定颜色的数量。

**参数**:
- `x1` (number): 区域左上角X坐标
- `y1` (number): 区域左上角Y坐标
- `x2` (number): 区域右下角X坐标
- `y2` (number): 区域右下角Y坐标
- `colorStr` (string): 颜色字符串
- `sim` (number): 相似度阈值 (0-1)
- `displayId` (number, 可选): 显示器ID，默认为0

**返回值**: `number` - 颜色数量

**调用示例**:
```javascript
const count = images.getColorCountInRegion(0, 0, 1080, 1920, "#FF0000", 0.9, 0);
console.log("颜色数量: " + count);
```

### images.detectsMultiColors(colors, sim, displayId?)

检测多点颜色。

**参数**:
- `colors` (string): 多点颜色字符串，格式为 "x1,y1,color1,x2,y2,color2,..."
- `sim` (number): 相似度阈值 (0-1)
- `displayId` (number, 可选): 显示器ID，默认为0

**返回值**: `boolean` - 如果多点颜色匹配返回 `true`，否则返回 `false`

**调用示例**:
```javascript
if (images.detectsMultiColors("0,0,#FF0000,10,10,#00FF00", 0.9, 0)) {
    console.log("多点颜色匹配");
}
```

### images.findMultiColors(x1, y1, x2, y2, colors, sim, dir, displayId?)

查找多点颜色。

**参数**:
- `x1` (number): 搜索区域左上角X坐标
- `y1` (number): 搜索区域左上角Y坐标
- `x2` (number): 搜索区域右下角X坐标
- `y2` (number): 搜索区域右下角Y坐标
- `colors` (string): 多点颜色字符串，格式为 "x1,y1,color1,x2,y2,color2,..."
- `sim` (number): 相似度阈值 (0-1)
- `dir` (number): 搜索方向 (0=从左到右, 1=从右到左, 2=从上到下, 3=从下到上)
- `displayId` (number, 可选): 显示器ID，默认为0

**返回值**: `number`, `number` - 找到的坐标 (x, y)，如果未找到返回 (-1, -1)

**调用示例**:
```javascript
const [x, y] = images.findMultiColors(0, 0, 1080, 1920, "0,0,#FF0000,10,10,#00FF00", 0.9, 0, 0);
if (x !== -1 && y !== -1) {
    console.log("找到多点颜色在: " + x + ", " + y);
}
```

### images.readFromPath(path)

从路径读取图片。

**参数**:
- `path` (string): 图片文件路径

**返回值**: `object` - 图片对象，如果失败返回 `null`

**调用示例**:
```javascript
const img = images.readFromPath("/sdcard/image.png");
if (img !== null) {
    console.log("图片读取成功");
}
```

### images.readFromUrl(url)

从 URL 读取图片。

**参数**:
- `url` (string): 图片URL

**返回值**: `object` - 图片对象，如果失败返回 `null`

**调用示例**:
```javascript
const img = images.readFromUrl("https://example.com/image.png");
if (img !== null) {
    console.log("图片读取成功");
}
```

### images.readFromBase64(base64Str)

从 Base64 读取图片。

**参数**:
- `base64Str` (string): Base64编码的图片数据

**返回值**: `object` - 图片对象，如果失败返回 `null`

**调用示例**:
```javascript
const img = images.readFromBase64("iVBORw0KG...");
if (img !== null) {
    console.log("图片读取成功");
}
```

### images.readFromBytes(bytes)

从字节数组读取图片。

**参数**:
- `bytes` (Array): 图片字节数组

**返回值**: `object` - 图片对象，如果失败返回 `null`

**调用示例**:
```javascript
const img = images.readFromBytes([72, 101, 108, 108, 111]);
if (img !== null) {
    console.log("图片读取成功");
}
```

### images.save(img, path, quality?)

保存图片。

**参数**:
- `img` (object): 图片对象
- `path` (string): 保存路径
- `quality` (number, 可选): 图片质量 (1-100)，默认为90

**返回值**: `boolean` - 是否保存成功

**调用示例**:
```javascript
const success = images.save(img, "/sdcard/output.png", 90);
if (success) {
    console.log("图片保存成功");
}
```

### images.encodeToBase64(img, format?, quality?)

编码为 Base64。

**参数**:
- `img` (object): 图片对象
- `format` (string, 可选): 图片格式 (png, jpeg)，默认为png
- `quality` (number, 可选): 图片质量 (1-100)，默认为90

**返回值**: `string` - Base64编码的图片数据

**调用示例**:
```javascript
const base64 = images.encodeToBase64(img, "png", 90);
console.log("Base64长度: " + base64.length);
```

### images.encodeToBytes(img, format?, quality?)

编码为字节数组。

**参数**:
- `img` (object): 图片对象
- `format` (string, 可选): 图片格式 (png, jpeg)，默认为png
- `quality` (number, 可选): 图片质量 (1-100)，默认为90

**返回值**: `Array` - 图片字节数组

**调用示例**:
```javascript
const bytes = images.encodeToBytes(img, "png", 90);
console.log("字节数组长度: " + bytes.length);
```

### images.clip(img, x1, y1, x2, y2)

裁剪图片。

**参数**:
- `img` (object): 图片对象
- `x1` (number): 裁剪区域左上角X坐标
- `y1` (number): 裁剪区域左上角Y坐标
- `x2` (number): 裁剪区域右下角X坐标
- `y2` (number): 裁剪区域右下角Y坐标

**返回值**: `object` - 裁剪后的图片对象

**调用示例**:
```javascript
const clipped = images.clip(img, 100, 100, 200, 200);
console.log("图片裁剪成功");
```

### images.resize(img, width, height)

调整图片大小。

**参数**:
- `img` (object): 图片对象
- `width` (number): 目标宽度
- `height` (number): 目标高度

**返回值**: `object` - 调整大小后的图片对象

**调用示例**:
```javascript
const resized = images.resize(img, 500, 500);
console.log("图片调整大小成功");
```

### images.rotate(img, degree)

旋转图片。

**参数**:
- `img` (object): 图片对象
- `degree` (number): 旋转角度

**返回值**: `object` - 旋转后的图片对象

**调用示例**:
```javascript
const rotated = images.rotate(img, 90);
console.log("图片旋转成功");
```

### images.grayscale(img)

灰度化。

**参数**:
- `img` (object): 图片对象

**返回值**: `object` - 灰度化后的图片对象

**调用示例**:
```javascript
const gray = images.grayscale(img);
console.log("图片灰度化成功");
```

### images.applyThreshold(img, threshold, maxVal, typ)

应用阈值。

**参数**:
- `img` (object): 图片对象
- `threshold` (number): 阈值
- `maxVal` (number): 最大值
- `typ` (string): 阈值类型 (BINARY, BINARY_INV, TRUNC, TOZERO, TOZERO_INV, OTSU)

**返回值**: `object` - 应用阈值后的图片对象

**调用示例**:
```javascript
const threshold = images.applyThreshold(img, 128, 255, "BINARY");
console.log("图片阈值化成功");
```

### images.applyAdaptiveThreshold(img, maxValue, adaptiveMethod, thresholdType, blockSize, C)

应用自适应阈值。

**参数**:
- `img` (object): 图片对象
- `maxValue` (number): 最大值
- `adaptiveMethod` (string): 自适应方法 (MEAN_C, GAUSSIAN_C)
- `thresholdType` (string): 阈值类型 (BINARY, BINARY_INV)
- `blockSize` (number): 块大小
- `C` (number): 常数

**返回值**: `object` - 应用自适应阈值后的图片对象

**调用示例**:
```javascript
const adaptive = images.applyAdaptiveThreshold(img, 255, "GAUSSIAN_C", "BINARY", 11, 2);
console.log("图片自适应阈值化成功");
```

### images.applyBinarization(img, threshold)

二值化。

**参数**:
- `img` (object): 图片对象
- `threshold` (number): 阈值

**返回值**: `object` - 二值化后的图片对象

**调用示例**:
```javascript
const binary = images.applyBinarization(img, 128);
console.log("图片二值化成功");
```

## 完整示例

```javascript
// 获取像素颜色
const color = images.pixel(500, 1000, 0);
console.log("像素颜色: " + color);

// 比较颜色
if (images.cmpColor(500, 1000, "#FF0000", 0.9, 0)) {
    console.log("颜色匹配");
}

// 查找颜色
const [x, y] = images.findColor(0, 0, 1080, 1920, "#FF0000", 0.9, 0, 0);
if (x !== -1 && y !== -1) {
    console.log("找到颜色在: " + x + ", " + y);
    click(x, y, 1, 0);
}

// 读取图片
const img = images.readFromPath("/sdcard/template.png");
if (img !== null) {
    // 裁剪图片
    const clipped = images.clip(img, 100, 100, 200, 200);
    
    // 调整大小
    const resized = images.resize(clipped, 500, 500);
    
    // 灰度化
    const gray = images.grayscale(resized);
    
    // 保存
    images.save(gray, "/sdcard/output.png", 90);
    console.log("图片处理完成");
}
```
