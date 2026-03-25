# images 模块

## 模块简介

images 模块提供了图像处理功能，包括屏幕截图、颜色识别、图像变换（裁剪、缩放、旋转等）以及图像格式转换等功能。

## 方法列表

### images.pixel
获取指定坐标的像素颜色

**使用示例：**
```lua
-- 调用 images.pixel 方法
images.pixel();
```

---

### images.setCallback
设置回调函数

**使用示例：**
```lua
-- 调用 images.setCallback 方法
images.setCallback();
```

---

### images.captureScreen
截取屏幕

**使用示例：**
```lua
-- 调用 images.captureScreen 方法
images.captureScreen();
```

---

### images.cmpColor
比较颜色

**使用示例：**
```lua
-- 调用 images.cmpColor 方法
images.cmpColor();
```

---

### images.findColor
查找颜色

**使用示例：**
```lua
-- 调用 images.findColor 方法
images.findColor();
```

---

### images.getColorCountInRegion
获取区域内指定颜色的数量

**使用示例：**
```lua
-- 调用 images.getColorCountInRegion 方法
images.getColorCountInRegion();
```

---

### images.detectsMultiColors
检测多点颜色

**使用示例：**
```lua
-- 调用 images.detectsMultiColors 方法
images.detectsMultiColors();
```

---

### images.readFromPath
从路径读取图片

**使用示例：**
```lua
-- 调用 images.readFromPath 方法
images.readFromPath();
```

---

### images.readFromUrl
从URL读取图片

**使用示例：**
```lua
-- 调用 images.readFromUrl 方法
images.readFromUrl();
```

---

### images.readFromBase64
从Base64读取图片

**使用示例：**
```lua
-- 调用 images.readFromBase64 方法
images.readFromBase64();
```

---

### images.readFromBytes
从字节数组读取图片

**使用示例：**
```lua
-- 调用 images.readFromBytes 方法
images.readFromBytes();
```

---

### images.toNrgba
转换为NRGBA格式

**使用示例：**
```lua
-- 调用 images.toNrgba 方法
images.toNrgba();
```

---

### images.save
保存图片

**使用示例：**
```lua
-- 调用 images.save 方法
images.save();
```

---

### images.encodeToBase64
编码为Base64

**使用示例：**
```lua
-- 调用 images.encodeToBase64 方法
images.encodeToBase64();
```

---

### images.encodeToBytes
编码为字节数组

**使用示例：**
```lua
-- 调用 images.encodeToBytes 方法
images.encodeToBytes();
```

---

### images.clip
裁剪图片

**使用示例：**
```lua
-- 调用 images.clip 方法
images.clip();
```

---

### images.resize
调整图片大小

**使用示例：**
```lua
-- 调用 images.resize 方法
images.resize();
```

---

### images.rotate
旋转图片

**使用示例：**
```lua
-- 调用 images.rotate 方法
images.rotate();
```

---

### images.grayscale
灰度化

**使用示例：**
```lua
-- 调用 images.grayscale 方法
images.grayscale();
```

---

### images.applyThreshold
应用阈值

**使用示例：**
```lua
-- 调用 images.applyThreshold 方法
images.applyThreshold();
```

---

### images.applyAdaptiveThreshold
应用自适应阈值

**使用示例：**
```lua
-- 调用 images.applyAdaptiveThreshold 方法
images.applyAdaptiveThreshold();
```

---

### images.applyBinarization
二值化

**使用示例：**
```lua
-- 调用 images.applyBinarization 方法
images.applyBinarization();
```

---

## 综合使用示例

### 示例1：截图并保存
```lua
var img = images.captureScreen(0, 0, device.width, device.height);
images.save(img, "/sdcard/screenshot.png", 100);
```