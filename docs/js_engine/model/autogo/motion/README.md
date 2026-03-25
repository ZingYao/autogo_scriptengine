# motion 模块

## 模块简介

motion 模块提供了模拟用户操作的功能，包括点击、滑动、按键等手势操作。

## 方法列表

### motion.touchDown
按下屏幕

**使用示例：**
```javascript
// 调用 motion.touchDown 方法
motion.touchDown();
```

---

### motion.touchMove
移动手指

**使用示例：**
```javascript
// 调用 motion.touchMove 方法
motion.touchMove();
```

---

### motion.touchUp
抬起手指

**使用示例：**
```javascript
// 调用 motion.touchUp 方法
motion.touchUp();
```

---

### motion.click
点击

**使用示例：**
```javascript
// 调用 motion.click 方法
motion.click();
```

---

### motion.longClick
长按

**使用示例：**
```javascript
// 调用 motion.longClick 方法
motion.longClick();
```

---

### motion.swipe
滑动

**使用示例：**
```javascript
// 调用 motion.swipe 方法
motion.swipe();
```

---

### motion.swipe2
滑动(两点)

**使用示例：**
```javascript
// 调用 motion.swipe2 方法
motion.swipe2();
```

---

### motion.home
按下Home键

**使用示例：**
```javascript
// 调用 motion.home 方法
motion.home();
```

---

### motion.back
按下返回键

**使用示例：**
```javascript
// 调用 motion.back 方法
motion.back();
```

---

### motion.recents
按下最近任务键

**使用示例：**
```javascript
// 调用 motion.recents 方法
motion.recents();
```

---

### motion.powerDialog
长按电源键

**使用示例：**
```javascript
// 调用 motion.powerDialog 方法
motion.powerDialog();
```

---

### motion.notifications
下拉通知栏

**使用示例：**
```javascript
// 调用 motion.notifications 方法
motion.notifications();
```

---

### motion.quickSettings
下拉快捷设置

**使用示例：**
```javascript
// 调用 motion.quickSettings 方法
motion.quickSettings();
```

---

### motion.volumeUp
按下音量加键

**使用示例：**
```javascript
// 调用 motion.volumeUp 方法
motion.volumeUp();
```

---

### motion.volumeDown
按下音量减键

**使用示例：**
```javascript
// 调用 motion.volumeDown 方法
motion.volumeDown();
```

---

### motion.camera
按下相机键

**使用示例：**
```javascript
// 调用 motion.camera 方法
motion.camera();
```

---

### motion.keyAction
按键动作

**使用示例：**
```javascript
// 调用 motion.keyAction 方法
motion.keyAction();
```

---

## 综合使用示例

### 示例1：点击屏幕
```javascript
click(500, 500);  // 在坐标(500, 500)处点击
swipe(100, 500, 900, 500, 500);  // 从(100,500)滑动到(900,500)，耗时500ms
```