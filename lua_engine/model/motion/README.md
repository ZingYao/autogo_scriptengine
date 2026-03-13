# Motion 模块

Motion 模块提供了触摸操作和系统按键功能。

## 方法列表

### touchDown(x, y, fingerId?, displayId?)

按下屏幕。

**参数**:
- `x` (number): X坐标
- `y` (number): Y坐标
- `fingerId` (number, 可选): 手指ID，默认为0
- `displayId` (number, 可选): 显示器ID，默认为0

**返回值**: `undefined`

**调用示例**:
```lua
touchDown(500, 1000, 1, 0)
```

### touchMove(x, y, fingerId?, displayId?)

移动手指。

**参数**:
- `x` (number): X坐标
- `y` (number): Y坐标
- `fingerId` (number, 可选): 手指ID，默认为0
- `displayId` (number, 可选): 显示器ID，默认为0

**返回值**: `undefined`

**调用示例**:
```lua
touchMove(600, 1100, 1, 0)
```

### touchUp(x, y, fingerId?, displayId?)

抬起手指。

**参数**:
- `x` (number): X坐标
- `y` (number): Y坐标
- `fingerId` (number, 可选): 手指ID，默认为0
- `displayId` (number, 可选): 显示器ID，默认为0

**返回值**: `undefined`

**调用示例**:
```lua
touchUp(600, 1100, 1, 0)
```

### click(x, y, fingerId?, displayId?)

点击。

**参数**:
- `x` (number): X坐标
- `y` (number): Y坐标
- `fingerId` (number, 可选): 手指ID，默认为0
- `displayId` (number, 可选): 显示器ID，默认为0

**返回值**: `undefined`

**调用示例**:
```lua
click(500, 1000, 1, 0)
```

### longClick(x, y, duration, fingerId?, displayId?)

长按。

**参数**:
- `x` (number): X坐标
- `y` (number): Y坐标
- `duration` (number): 长按时间(毫秒)
- `fingerId` (number, 可选): 手指ID，默认为0
- `displayId` (number, 可选): 显示器ID，默认为0

**返回值**: `undefined`

**调用示例**:
```lua
longClick(500, 1000, 2000, 1, 0)
```

### swipe(x1, y1, x2, y2, duration, fingerId?, displayId?)

滑动。

**参数**:
- `x1` (number): 起始X坐标
- `y1` (number): 起始Y坐标
- `x2` (number): 结束X坐标
- `y2` (number): 结束Y坐标
- `duration` (number): 滑动时间(毫秒)
- `fingerId` (number, 可选): 手指ID，默认为0
- `displayId` (number, 可选): 显示器ID，默认为0

**返回值**: `undefined`

**调用示例**:
```lua
swipe(500, 1000, 600, 1100, 500, 1, 0)
```

### swipe2(x1, y1, x2, y2, duration, fingerId?, displayId?)

滑动(两点)。

**参数**:
- `x1` (number): 起始X坐标
- `y1` (number): 起始Y坐标
- `x2` (number): 结束X坐标
- `y2` (number): 结束Y坐标
- `duration` (number): 滑动时间(毫秒)
- `fingerId` (number, 可选): 手指ID，默认为0
- `displayId` (number, 可选): 显示器ID，默认为0

**返回值**: `undefined`

**调用示例**:
```lua
swipe2(500, 1000, 600, 1100, 500, 1, 0)
```

### home(displayId?)

按下 Home 键。

**参数**:
- `displayId` (number, 可选): 显示器ID，默认为0

**返回值**: `undefined`

**调用示例**:
```lua
home(0)
```

### back(displayId?)

按下返回键。

**参数**:
- `displayId` (number, 可选): 显示器ID，默认为0

**返回值**: `undefined`

**调用示例**:
```lua
back(0)
```

### recents(displayId?)

按下最近任务键。

**参数**:
- `displayId` (number, 可选): 显示器ID，默认为0

**返回值**: `undefined`

**调用示例**:
```lua
recents(0)
```

### powerDialog()

长按电源键。

**返回值**: `undefined`

**调用示例**:
```lua
powerDialog()
```

### notifications()

下拉通知栏。

**返回值**: `undefined`

**调用示例**:
```lua
notifications()
```

### quickSettings()

下拉快捷设置。

**返回值**: `undefined`

**调用示例**:
```lua
quickSettings()
```

### volumeUp(displayId?)

按下音量加键。

**参数**:
- `displayId` (number, 可选): 显示器ID，默认为0

**返回值**: `undefined`

**调用示例**:
```lua
volumeUp(0)
```

### volumeDown(displayId?)

按下音量减键。

**参数**:
- `displayId` (number, 可选): 显示器ID，默认为0

**返回值**: `undefined`

**调用示例**:
```lua
volumeDown(0)
```

### camera()

按下相机键。

**返回值**: `undefined`

**调用示例**:
```lua
camera()
```

### keyAction(code, displayId?)

按键动作。

**参数**:
- `code` (number): 按键代码
- `displayId` (number, 可选): 显示器ID，默认为0

**返回值**: `undefined`

**调用示例**:
```lua
keyAction(3, 0) -- KEYCODE_HOME
```

## 完整示例

```lua
-- 基本触摸操作
print("=== 基本触摸操作 ===")
touchDown(500, 1000, 1, 0)
touchMove(600, 1100, 1, 0)
touchUp(600, 1100, 1, 0)

-- 点击
print("\n=== 点击操作 ===")
click(500, 1000, 1, 0)

-- 长按
print("\n=== 长按操作 ===")
longClick(500, 1000, 2000, 1, 0)

-- 滑动
print("\n=== 滑动操作 ===")
swipe(500, 1000, 600, 1100, 500, 1, 0)

-- 系统按键
print("\n=== 系统按键 ===")
home(0)
sleep(1000)
back(0)
sleep(1000)
recents(0)

-- 音量控制
print("\n=== 音量控制 ===")
volumeUp(0)
sleep(500)
volumeDown(0)

-- 其他按键
print("\n=== 其他按键 ===")
powerDialog()
notifications()
quickSettings()
```
