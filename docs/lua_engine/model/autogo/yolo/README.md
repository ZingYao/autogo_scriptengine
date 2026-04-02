# yolo 模块

## 模块简介

yolo 模块提供了 YOLO（You Only Look Once）目标检测功能，用于实时对象检测。

## 方法列表

### yolo.new
创建一个新的YOLO实例

**参数：**
- `version` (string): YOLO 模型版本
- `cpuThreadNum` (int): CPU 线程数
- `paramPath` (string): 参数文件路径
- `binPath` (string): 模型文件路径
- `labels` (string): 标签文件路径

**返回值：**
- Yolo 对象: YOLO 实例

**使用示例：**
```lua
-- 创建 YOLO 实例
local yolo = yolo.new("yolov5n", 4, "/sdcard/yolo/yolov5n.param", "/sdcard/yolo/yolov5n.bin", "/sdcard/yolo/coco.names")
```

---

### yolo.detect
检测屏幕上的对象

**参数：**
- `y` (Yolo): YOLO 实例
- `x1` (int): 区域左上角 X 坐标
- `y1` (int): 区域左上角 Y 坐标
- `x2` (int): 区域右下角 X 坐标
- `y2` (int): 区域右下角 Y 坐标
- `displayId` (int): 显示设备 ID（可选，默认 0）

**返回值：**
- Array: 检测结果数组，每个元素包含 x, y, width, height, label, confidence 属性

**使用示例：**
```lua
-- 创建 YOLO 实例
local yolo = yolo.new("yolov5n", 4, "/sdcard/yolo/yolov5n.param", "/sdcard/yolo/yolov5n.bin", "/sdcard/yolo/coco.names")

-- 检测屏幕上的对象
local results = yolo.detect(0, 0, 1920, 1080, 0)
for i, result in ipairs(results) do
    print("对象 " .. i .. ":")
    print("  标签: " .. result.label)
    print("  位置: (" .. result.x .. ", " .. result.y .. ")")
    print("  大小: " .. result.width .. "x" .. result.height)
    print("  置信度: " .. result.confidence)
end
```

---

### yolo.detectFromImage
检测图片中的对象

**参数：**
- `y` (Yolo): YOLO 实例
- `img` (image): 图片对象

**返回值：**
- Array: 检测结果数组，每个元素包含 x, y, width, height, label, confidence 属性

**使用示例：**
```lua
-- 创建 YOLO 实例
local yolo = yolo.new("yolov5n", 4, "/sdcard/yolo/yolov5n.param", "/sdcard/yolo/yolov5n.bin", "/sdcard/yolo/coco.names")

-- 截取屏幕
local img = images.captureScreen(0, 0, 1920, 1080)

-- 检测图片中的对象
local results = yolo.detectFromImage(img)
for i, result in ipairs(results) do
    print("对象 " .. i .. ": " .. result.label)
end
```

---

### yolo.detectFromBase64
检测Base64图片中的对象

**参数：**
- `y` (Yolo): YOLO 实例
- `b64` (string): Base64 编码的图片数据
- `colorStr` (string): 颜色字符串（可选）

**返回值：**
- Array: 检测结果数组，每个元素包含 x, y, width, height, label, confidence 属性

**使用示例：**
```lua
-- 创建 YOLO 实例
local yolo = yolo.new("yolov5n", 4, "/sdcard/yolo/yolov5n.param", "/sdcard/yolo/yolov5n.bin", "/sdcard/yolo/coco.names")

-- 检测 Base64 图片中的对象
local results = yolo.detectFromBase64(base64String, "")
for i, result in ipairs(results) do
    print("对象 " .. i .. ": " .. result.label)
end
```

---

### yolo.detectFromPath
检测文件图片中的对象

**参数：**
- `y` (Yolo): YOLO 实例
- `path` (string): 图片文件路径
- `colorStr` (string): 颜色字符串（可选）

**返回值：**
- Array: 检测结果数组，每个元素包含 x, y, width, height, label, confidence 属性

**使用示例：**
```lua
-- 创建 YOLO 实例
local yolo = yolo.new("yolov5n", 4, "/sdcard/yolo/yolov5n.param", "/sdcard/yolo/yolov5n.bin", "/sdcard/yolo/coco.names")

-- 检测文件图片中的对象
local results = yolo.detectFromPath("/sdcard/test.jpg", "")
for i, result in ipairs(results) do
    print("对象 " .. i .. ": " .. result.label)
end
```

---

### yolo.close
关闭YOLO实例

**参数：**
- `y` (Yolo): YOLO 实例

**使用示例：**
```lua
-- 创建 YOLO 实例
local yolo = yolo.new("yolov5n", 4, "/sdcard/yolo/yolov5n.param", "/sdcard/yolo/yolov5n.bin", "/sdcard/yolo/coco.names")

-- 使用完毕后关闭实例
yolo.close()
```

---

## 综合使用示例

### 示例1：创建实例并检测屏幕对象
```lua
-- 创建 YOLO 实例
local yolo = yolo.new("yolov5n", 4, "/sdcard/yolo/yolov5n.param", "/sdcard/yolo/yolov5n.bin", "/sdcard/yolo/coco.names")

-- 检测屏幕上的对象
local results = yolo.detect(0, 0, 1920, 1080, 0)
for i, result in ipairs(results) do
    print("对象 " .. i .. ": " .. result.label .. " (" .. result.confidence .. ")")
end

-- 关闭实例
yolo.close()
```

### 示例2：检测图片中的对象
```lua
-- 创建 YOLO 实例
local yolo = yolo.new("yolov5n", 4, "/sdcard/yolo/yolov5n.param", "/sdcard/yolo/yolov5n.bin", "/sdcard/yolo/coco.names")

-- 截取屏幕
local img = images.captureScreen(0, 0, 1920, 1080)

-- 检测图片中的对象
local results = yolo.detectFromImage(img)
for i, result in ipairs(results) do
    print("对象 " .. i .. ": " .. result.label)
end

-- 关闭实例
yolo.close()
```

### 示例3：检测文件图片中的对象
```lua
-- 创建 YOLO 实例
local yolo = yolo.new("yolov5n", 4, "/sdcard/yolo/yolov5n.param", "/sdcard/yolo/yolov5n.bin", "/sdcard/yolo/coco.names")

-- 检测文件图片中的对象
local results = yolo.detectFromPath("/sdcard/test.jpg", "")
for i, result in ipairs(results) do
    print("对象 " .. i .. ": " .. result.label .. " 位置: (" .. result.x .. ", " .. result.y .. ")")
end

-- 关闭实例
yolo.close()
```