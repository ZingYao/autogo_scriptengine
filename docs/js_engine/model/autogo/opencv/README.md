# opencv 模块

## 模块简介

opencv 模块提供了 OpenCV 计算机视觉库的功能，支持图像匹配、特征检测等高级图像处理功能。

## 方法列表

### opencv.findImage
在指定区域内查找匹配的图片模板

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| x1 | number | 区域左上角 X 坐标 |
| y1 | number | 区域左上角 Y 坐标 |
| x2 | number | 区域右下角 X 坐标 |
| y2 | number | 区域右下角 Y 坐标 |
| template | string/array | 模板数据（可以是字符串路径或字节数组） |
| isGray | boolean | 是否灰度化 |
| scalingFactor | number | 缩放因子 |
| sim | number | 相似度 (0-1) |
| displayId | number | 显示设备 ID（可选，默认为0） |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| x | number | 找到的 X 坐标（未找到返回 -1） |
| y | number | 找到的 Y 坐标（未找到返回 -1） |

**使用示例：**
```javascript
// 在指定区域内查找匹配的图片模板
var result = opencv.findImage(0, 0, 1000, 1000, "/sdcard/template.png", false, 1.0, 0.9, 0);
if (result.x !== -1) {
    console.log("找到图片在: " + result.x + ", " + result.y);
}
```

---

## 综合使用示例

### 示例1：基本使用
```javascript
// opencv 模块的基本使用示例
```