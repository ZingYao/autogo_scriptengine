# opencv 模块

## 模块简介

opencv 模块提供了 OpenCV 计算机视觉库的功能，支持图像匹配、特征检测等高级图像处理功能。

## 方法列表

### opencv.findImage
在指定区域内查找匹配的图片模板

**参数：**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| x1 | int | 是 | 搜索区域左上角X坐标 |
| y1 | int | 是 | 搜索区域左上角Y坐标 |
| x2 | int | 是 | 搜索区域右下角X坐标 |
| y2 | int | 是 | 搜索区域右下角Y坐标 |
| template | string | 是 | 模板图片的字节数据（Base64编码或二进制字符串） |
| isGray | boolean | 是 | 是否使用灰度模式进行匹配 |
| scalingFactor | number | 是 | 缩放因子，用于调整搜索精度 |
| sim | number | 是 | 相似度阈值（0-1之间，值越大匹配越严格） |
| displayId | int | 是 | 显示器ID，用于指定在哪个屏幕上搜索 |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| x | number | 找到的图片中心点X坐标（未找到时为0） |
| y | number | 找到的图片中心点Y坐标（未找到时为0） |

**使用示例：**
```lua
-- 调用 opencv.findImage 方法
local x, y = opencv.findImage(0, 0, 1920, 1080, templateBytes, false, 1.0, 0.8, 0)
if x > 0 and y > 0 then
    print("找到图片，坐标：(" .. x .. ", " .. y .. ")")
end
```

---

## 综合使用示例

### 示例1：基本使用
```lua
-- opencv 模块的基本使用示例
```