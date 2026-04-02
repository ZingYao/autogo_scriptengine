# vdisplay 模块

## 模块简介

vdisplay 模块提供了虚拟显示设备的功能，可以创建和管理虚拟屏幕。

## 方法列表

### vdisplay.create
创建一个虚拟显示设备

**参数：**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| width | number | 是 | 显示宽度（像素） |
| height | number | 是 | 显示高度（像素） |
| dpi | number | 是 | 显示DPI |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| vdisplay | userdata | 虚拟显示设备对象 |

**使用示例：**
```lua
-- 调用 vdisplay.create 方法
local vd = vdisplay.create(1080, 1920, 320)
```

---

### vdisplay.getDisplayId
获取虚拟显示设备的DisplayId

**参数：**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| vdisplay | userdata | 是 | 虚拟显示设备对象 |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| displayId | number | 虚拟显示设备的DisplayId |

**使用示例：**
```lua
-- 调用 vdisplay.getDisplayId 方法
local id = vdisplay.getDisplayId(vd)
```

---

### vdisplay.launchApp
启动指定包名的应用到虚拟显示设备内

**参数：**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| vdisplay | userdata | 是 | 虚拟显示设备对象 |
| packageName | string | 是 | 应用包名 |

**返回值：**

无

**使用示例：**
```lua
-- 调用 vdisplay.launchApp 方法
vdisplay.launchApp(vd, "com.example.app")
```

---

### vdisplay.setTitle
设置预览窗口标题

**参数：**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| vdisplay | userdata | 是 | 虚拟显示设备对象 |
| title | string | 是 | 窗口标题 |

**返回值：**

无

**使用示例：**
```lua
-- 调用 vdisplay.setTitle 方法
vdisplay.setTitle(vd, "虚拟屏幕")
```

---

### vdisplay.setTouchCallback
设置触控回调

**参数：**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| vdisplay | userdata | 是 | 虚拟显示设备对象 |
| callback | function | 是 | 触控回调函数，回调参数：(x, y, action, displayId) |

**返回值：**

无

**使用示例：**
```lua
-- 调用 vdisplay.setTouchCallback 方法
vdisplay.setTouchCallback(vd, function(x, y, action, displayId)
    print("触控事件:", x, y, action, displayId)
end)
```

---

### vdisplay.showPreviewWindow
显示预览窗口

**参数：**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| vdisplay | userdata | 是 | 虚拟显示设备对象 |
| rotated | boolean | 否 | 是否旋转显示（可选，默认false） |

**返回值：**

无

**使用示例：**
```lua
-- 调用 vdisplay.showPreviewWindow 方法
vdisplay.showPreviewWindow(vd, false)
```

---

### vdisplay.hidePreviewWindow
隐藏预览窗口

**参数：**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| vdisplay | userdata | 是 | 虚拟显示设备对象 |

**返回值：**

无

**使用示例：**
```lua
-- 调用 vdisplay.hidePreviewWindow 方法
vdisplay.hidePreviewWindow(vd)
```

---

### vdisplay.setPreviewWindowSize
设置预览窗口大小

**参数：**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| vdisplay | userdata | 是 | 虚拟显示设备对象 |
| width | number | 是 | 窗口宽度 |
| height | number | 是 | 窗口高度 |

**返回值：**

无

**使用示例：**
```lua
-- 调用 vdisplay.setPreviewWindowSize 方法
vdisplay.setPreviewWindowSize(vd, 540, 960)
```

---

### vdisplay.setPreviewWindowPos
设置预览窗口位置

**参数：**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| vdisplay | userdata | 是 | 虚拟显示设备对象 |
| x | number | 是 | 窗口X坐标 |
| y | number | 是 | 窗口Y坐标 |

**返回值：**

无

**使用示例：**
```lua
-- 调用 vdisplay.setPreviewWindowPos 方法
vdisplay.setPreviewWindowPos(vd, 100, 100)
```

---

### vdisplay.destroy
销毁指定的虚拟显示设备

**参数：**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| vdisplay | userdata | 是 | 虚拟显示设备对象 |

**返回值：**

无

**使用示例：**
```lua
-- 调用 vdisplay.destroy 方法
vdisplay.destroy(vd)
```

---

## 综合使用示例

### 示例1：基本使用
```lua
-- vdisplay 模块的基本使用示例
```