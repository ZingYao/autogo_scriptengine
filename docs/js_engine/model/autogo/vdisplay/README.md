# vdisplay 模块

## 模块简介

vdisplay 模块提供了虚拟显示设备的功能，可以创建和管理虚拟屏幕。

## 方法列表

### vdisplay.create
创建一个虚拟显示设备

**参数：**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| width | int | 是 | 虚拟屏幕的宽度（像素） |
| height | int | 是 | 虚拟屏幕的高度（像素） |
| dpi | int | 是 | 虚拟屏幕的DPI（每英寸点数） |

**返回值：**

| 类型 | 说明 |
|------|------|
| Vdisplay | 虚拟显示设备对象 |

**使用示例：**
```javascript
// 创建一个1080p的虚拟屏幕
let vd = vdisplay.create(1920, 1080, 320);

// 创建一个720p的虚拟屏幕
let vd2 = vdisplay.create(1280, 720, 240);

// 创建一个小尺寸的虚拟屏幕
let vd3 = vdisplay.create(800, 600, 160);
```

---

### vdisplay.getDisplayId
获取虚拟显示设备的DisplayId

**参数：**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| v | Vdisplay | 是 | 虚拟显示设备对象 |

**返回值：**

| 类型 | 说明 |
|------|------|
| int | 虚拟显示设备的DisplayId |

**使用示例：**
```javascript
// 创建虚拟显示设备
let vd = vdisplay.create(1920, 1080, 320);

// 获取DisplayId
let displayId = vdisplay.getDisplayId(vd);
console.log("DisplayId: " + displayId);
```

---

### vdisplay.launchApp
启动指定包名的应用到虚拟显示设备内

**参数：**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| v | Vdisplay | 是 | 虚拟显示设备对象 |
| packageName | string | 是 | 要启动的应用包名 |

**返回值：**

| 类型 | 说明 |
|------|------|
| undefined | 无返回值 |

**使用示例：**
```javascript
// 创建虚拟显示设备
let vd = vdisplay.create(1920, 1080, 320);

// 启动应用到虚拟显示设备
vdisplay.launchApp(vd, "com.android.settings");

// 启动浏览器
vdisplay.launchApp(vd, "com.android.browser");

// 启动自定义应用
vdisplay.launchApp(vd, "com.example.myapp");
```

---

### vdisplay.setTitle
设置预览窗口标题

**参数：**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| v | Vdisplay | 是 | 虚拟显示设备对象 |
| title | string | 是 | 预览窗口的标题 |

**返回值：**

| 类型 | 说明 |
|------|------|
| undefined | 无返回值 |

**使用示例：**
```javascript
// 创建虚拟显示设备
let vd = vdisplay.create(1920, 1080, 320);

// 设置预览窗口标题
vdisplay.setTitle(vd, "我的虚拟屏幕");

// 设置带时间戳的标题
let timestamp = utils.timestamp();
vdisplay.setTitle(vd, "虚拟屏幕 - " + timestamp);
```

---

### vdisplay.setTouchCallback
设置触控回调

**参数：**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| v | Vdisplay | 是 | 虚拟显示设备对象 |
| callback | function | 是 | 触控回调函数，接收参数：x, y, action, displayId |

**返回值：**

| 类型 | 说明 |
|------|------|
| undefined | 无返回值 |

**使用示例：**
```javascript
// 创建虚拟显示设备
let vd = vdisplay.create(1920, 1080, 320);

// 设置触控回调
vdisplay.setTouchCallback(vd, function(x, y, action, displayId) {
    console.log("触控事件:");
    console.log("  位置: (" + x + ", " + y + ")");
    console.log("  动作: " + action);
    console.log("  DisplayId: " + displayId);
});

// 带业务逻辑的回调
vdisplay.setTouchCallback(vd, function(x, y, action, displayId) {
    if (action === 0) {
        console.log("按下: (" + x + ", " + y + ")");
    } else if (action === 1) {
        console.log("抬起: (" + x + ", " + y + ")");
    }
});
```

---

### vdisplay.showPreviewWindow
显示预览窗口

**参数：**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| v | Vdisplay | 是 | 虚拟显示设备对象 |
| rotated | boolean | 否 | 是否旋转显示（可选，默认false） |

**返回值：**

| 类型 | 说明 |
|------|------|
| undefined | 无返回值 |

**使用示例：**
```javascript
// 创建虚拟显示设备
let vd = vdisplay.create(1920, 1080, 320);

// 显示预览窗口（不旋转）
vdisplay.showPreviewWindow(vd, false);

// 显示预览窗口（旋转90度）
vdisplay.showPreviewWindow(vd, true);

// 使用默认参数显示
vdisplay.showPreviewWindow(vd);
```

---

### vdisplay.hidePreviewWindow
隐藏预览窗口

**参数：**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| v | Vdisplay | 是 | 虚拟显示设备对象 |

**返回值：**

| 类型 | 说明 |
|------|------|
| undefined | 无返回值 |

**使用示例：**
```javascript
// 创建虚拟显示设备
let vd = vdisplay.create(1920, 1080, 320);

// 显示预览窗口
vdisplay.showPreviewWindow(vd);

// 隐藏预览窗口
vdisplay.hidePreviewWindow(vd);

// 再次显示
vdisplay.showPreviewWindow(vd);
```

---

### vdisplay.setPreviewWindowSize
设置预览窗口大小

**参数：**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| v | Vdisplay | 是 | 虚拟显示设备对象 |
| width | int | 是 | 预览窗口的宽度 |
| height | int | 是 | 预览窗口的高度 |

**返回值：**

| 类型 | 说明 |
|------|------|
| undefined | 无返回值 |

**使用示例：**
```javascript
// 创建虚拟显示设备
let vd = vdisplay.create(1920, 1080, 320);

// 显示预览窗口
vdisplay.showPreviewWindow(vd);

// 设置预览窗口大小为800x600
vdisplay.setPreviewWindowSize(vd, 800, 600);

// 设置预览窗口大小为640x480
vdisplay.setPreviewWindowSize(vd, 640, 480);

// 设置预览窗口大小为全屏
vdisplay.setPreviewWindowSize(vd, 1920, 1080);
```

---

### vdisplay.setPreviewWindowPos
设置预览窗口位置

**参数：**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| v | Vdisplay | 是 | 虚拟显示设备对象 |
| x | int | 是 | 预览窗口的X坐标 |
| y | int | 是 | 预览窗口的Y坐标 |

**返回值：**

| 类型 | 说明 |
|------|------|
| undefined | 无返回值 |

**使用示例：**
```javascript
// 创建虚拟显示设备
let vd = vdisplay.create(1920, 1080, 320);

// 显示预览窗口
vdisplay.showPreviewWindow(vd);

// 设置预览窗口位置到左上角
vdisplay.setPreviewWindowPos(vd, 0, 0);

// 设置预览窗口位置到屏幕中央
vdisplay.setPreviewWindowPos(vd, 560, 240);

// 设置预览窗口位置到右下角
vdisplay.setPreviewWindowPos(vd, 1120, 480);
```

---

### vdisplay.destroy
销毁指定的虚拟显示设备

**参数：**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| v | Vdisplay | 是 | 虚拟显示设备对象 |

**返回值：**

| 类型 | 说明 |
|------|------|
| undefined | 无返回值 |

**使用示例：**
```javascript
// 创建虚拟显示设备
let vd = vdisplay.create(1920, 1080, 320);

// 使用虚拟显示设备
vdisplay.showPreviewWindow(vd);
vdisplay.launchApp(vd, "com.android.settings");

// 使用完毕后销毁
vdisplay.destroy(vd);

// 注意：销毁后不能再使用该对象
// vdisplay.showPreviewWindow(vd); // 这会出错
```

---

## 综合使用示例

### 示例1：基本使用
```javascript
// 创建虚拟显示设备
let vd = vdisplay.create(1920, 1080, 320);

// 设置预览窗口标题
vdisplay.setTitle(vd, "我的虚拟屏幕");

// 显示预览窗口
vdisplay.showPreviewWindow(vd);

// 启动应用
vdisplay.launchApp(vd, "com.android.settings");

// 使用完毕后销毁
vdisplay.destroy(vd);
```

### 示例2：完整的虚拟显示设备管理
```javascript
// 创建虚拟显示设备
let vd = vdisplay.create(1920, 1080, 320);

// 设置预览窗口标题
vdisplay.setTitle(vd, "测试虚拟屏幕");

// 设置预览窗口大小和位置
vdisplay.setPreviewWindowSize(vd, 800, 600);
vdisplay.setPreviewWindowPos(vd, 100, 100);

// 显示预览窗口（旋转）
vdisplay.showPreviewWindow(vd, true);

// 设置触控回调
vdisplay.setTouchCallback(vd, function(x, y, action, displayId) {
    console.log("触控: (" + x + ", " + y + ") action=" + action);
});

// 启动应用
vdisplay.launchApp(vd, "com.android.settings");

// 等待一段时间
utils.sleep(5000);

// 隐藏预览窗口
vdisplay.hidePreviewWindow(vd);

// 等待一段时间
utils.sleep(2000);

// 再次显示预览窗口
vdisplay.showPreviewWindow(vd);

// 使用完毕后销毁
vdisplay.destroy(vd);
```

### 示例3：多虚拟显示设备管理
```javascript
// 创建第一个虚拟显示设备
let vd1 = vdisplay.create(1920, 1080, 320);
vdisplay.setTitle(vd1, "虚拟屏幕1");
vdisplay.showPreviewWindow(vd1);
vdisplay.setPreviewWindowPos(vd1, 0, 0);

// 创建第二个虚拟显示设备
let vd2 = vdisplay.create(1280, 720, 240);
vdisplay.setTitle(vd2, "虚拟屏幕2");
vdisplay.showPreviewWindow(vd2);
vdisplay.setPreviewWindowPos(vd2, 800, 0);

// 在第一个屏幕启动应用
vdisplay.launchApp(vd1, "com.android.settings");

// 在第二个屏幕启动应用
vdisplay.launchApp(vd2, "com.android.browser");

// 等待一段时间
utils.sleep(10000);

// 销毁所有虚拟显示设备
vdisplay.destroy(vd1);
vdisplay.destroy(vd2);
```
