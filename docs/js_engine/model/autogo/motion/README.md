# motion 模块

## 模块简介

motion 模块提供了模拟用户操作的功能，包括点击、滑动、按键等手势操作。

## 方法列表

### motion.touchDown
按下屏幕

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| x | number | X 坐标 |
| y | number | Y 坐标 |
| fingerID | number | 手指ID（可选，默认为0） |
| displayId | number | 显示设备 ID（可选，默认为0） |

**返回值：**

无返回值

**使用示例：**
```javascript
// 调用 motion.touchDown 方法
motion.touchDown(500, 500);
```

---

### motion.touchMove
移动手指

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| x | number | X 坐标 |
| y | number | Y 坐标 |
| fingerID | number | 手指ID（可选，默认为0） |
| displayId | number | 显示设备 ID（可选，默认为0） |

**返回值：**

无返回值

**使用示例：**
```javascript
// 调用 motion.touchMove 方法
motion.touchMove(600, 500);
```

---

### motion.touchUp
抬起手指

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| x | number | X 坐标 |
| y | number | Y 坐标 |
| fingerID | number | 手指ID（可选，默认为0） |
| displayId | number | 显示设备 ID（可选，默认为0） |

**返回值：**

无返回值

**使用示例：**
```javascript
// 调用 motion.touchUp 方法
motion.touchUp(500, 500);
```

---

### motion.click
点击

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| x | number | X 坐标 |
| y | number | Y 坐标 |
| fingerID | number | 手指ID（可选，默认为0） |
| displayId | number | 显示设备 ID（可选，默认为0） |

**返回值：**

无返回值

**使用示例：**
```javascript
// 调用 motion.click 方法
motion.click(500, 500);
```

---

### motion.longClick
长按

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| x | number | X 坐标 |
| y | number | Y 坐标 |
| fingerID | number | 手指ID（可选，默认为0） |
| displayId | number | 显示设备 ID（可选，默认为0） |

**返回值：**

无返回值

**使用示例：**
```javascript
// 调用 motion.longClick 方法
motion.longClick(500, 500);
```

---

### motion.swipe
滑动

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| x1 | number | 起点 X 坐标 |
| y1 | number | 起点 Y 坐标 |
| x2 | number | 终点 X 坐标 |
| y2 | number | 终点 Y 坐标 |
| duration | number | 滑动持续时间（毫秒，可选，默认为 300） |

**返回值：**

无返回值

**使用示例：**
```javascript
// 调用 motion.swipe 方法
motion.swipe(100, 500, 900, 500);
```

---

### motion.swipe2
滑动(两点)

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| x1 | number | 第一点 X 坐标 |
| y1 | number | 第一点 Y 坐标 |
| x2 | number | 第二点 X 坐标 |
| y2 | number | 第二点 Y 坐标 |
| duration | number | 滑动持续时间（毫秒，可选，默认为 300） |

**返回值：**

无返回值

**使用示例：**
```javascript
// 调用 motion.swipe2 方法
motion.swipe2(100, 500, 900, 500);
```

---

### motion.home
按下Home键

**参数：**

无参数

**返回值：**

无返回值

**使用示例：**
```javascript
// 调用 motion.home 方法
motion.home();
```

---

### motion.back
按下返回键

**参数：**

无参数

**返回值：**

无返回值

**使用示例：**
```javascript
// 调用 motion.back 方法
motion.back();
```

---

### motion.recents
按下最近任务键

**参数：**

无参数

**返回值：**

无返回值

**使用示例：**
```javascript
// 调用 motion.recents 方法
motion.recents();
```

---

### motion.powerDialog
长按电源键

**参数：**

无参数

**返回值：**

无返回值

**使用示例：**
```javascript
// 调用 motion.powerDialog 方法
motion.powerDialog();
```

---

### motion.notifications
下拉通知栏

**参数：**

无参数

**返回值：**

无返回值

**使用示例：**
```javascript
// 调用 motion.notifications 方法
motion.notifications();
```

---

### motion.quickSettings
下拉快捷设置

**参数：**

无参数

**返回值：**

无返回值

**使用示例：**
```javascript
// 调用 motion.quickSettings 方法
motion.quickSettings();
```

---

### motion.volumeUp
按下音量加键

**参数：**

无参数

**返回值：**

无返回值

**使用示例：**
```javascript
// 调用 motion.volumeUp 方法
motion.volumeUp();
```

---

### motion.volumeDown
按下音量减键

**参数：**

无参数

**返回值：**

无返回值

**使用示例：**
```javascript
// 调用 motion.volumeDown 方法
motion.volumeDown();
```

---

### motion.camera
按下相机键

**参数：**

无参数

**返回值：**

无返回值

**使用示例：**
```javascript
// 调用 motion.camera 方法
motion.camera();
```

---

### motion.keyAction
按键动作

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| code | number | 按键码 |

**返回值：**

无返回值

**使用示例：**
```javascript
// 调用 motion.keyAction 方法
motion.keyAction(4); // 返回键
```

---

## 综合使用示例

### 示例1：点击屏幕
```javascript
click(500, 500);  // 在坐标(500, 500)处点击
swipe(100, 500, 900, 500, 500);  // 从(100,500)滑动到(900,500)，耗时500ms
```