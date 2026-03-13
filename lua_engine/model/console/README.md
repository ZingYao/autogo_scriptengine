# Console 模块

Console 模块提供了控制台窗口的创建和管理功能，用于在屏幕上显示日志信息和调试输出。

## 方法列表

### console.new()
创建一个新的控制台对象。

**入参：** 无

**出参：** Console 对象

**调用示例：**
```javascript
// 创建控制台对象
var consoleObj = console.new();
```

### console.println(consoleObj, ...)
在控制台输出一行文本。

**入参：**
- `consoleObj`: Console 对象
- `...`: 要输出的文本内容，支持多个参数

**出参：** 无

**调用示例：**
```javascript
// 输出文本
console.println(consoleObj, "Hello", "World");
console.println(consoleObj, "这是一条日志信息");
```

### console.setTextSize(consoleObj, size)
设置控制台文本的字体大小。

**入参：**
- `consoleObj`: Console 对象
- `size`: 字体大小（整数）

**出参：** 无

**调用示例：**
```javascript
// 设置字体大小为 16
console.setTextSize(consoleObj, 16);
```

### console.setTextColor(consoleObj, color)
设置控制台文本的颜色。

**入参：**
- `consoleObj`: Console 对象
- `color`: 颜色字符串（如 "#FF0000" 表示红色）

**出参：** 无

**调用示例：**
```javascript
// 设置文本颜色为红色
console.setTextColor(consoleObj, "#FF0000");
```

### console.setWindowSize(consoleObj, width, height)
设置控制台窗口的大小。

**入参：**
- `consoleObj`: Console 对象
- `width`: 窗口宽度（整数）
- `height`: 窗口高度（整数）

**出参：** 无

**调用示例：**
```javascript
// 设置窗口大小为 800x600
console.setWindowSize(consoleObj, 800, 600);
```

### console.setWindowPosition(consoleObj, x, y)
设置控制台窗口的位置。

**入参：**
- `consoleObj`: Console 对象
- `x`: 窗口左上角 X 坐标（整数）
- `y`: 窗口左上角 Y 坐标（整数）

**出参：** 无

**调用示例：**
```javascript
// 设置窗口位置为 (100, 100)
console.setWindowPosition(consoleObj, 100, 100);
```

### console.setWindowColor(consoleObj, color)
设置控制台窗口的背景颜色。

**入参：**
- `consoleObj`: Console 对象
- `color`: 颜色字符串（如 "#000000" 表示黑色）

**出参：** 无

**调用示例：**
```javascript
// 设置窗口背景为黑色
console.setWindowColor(consoleObj, "#000000");
```

### console.show(consoleObj)
显示控制台窗口。

**入参：**
- `consoleObj`: Console 对象

**出参：** 无

**调用示例：**
```javascript
// 显示控制台窗口
console.show(consoleObj);
```

### console.hide(consoleObj)
隐藏控制台窗口。

**入参：**
- `consoleObj`: Console 对象

**出参：** 无

**调用示例：**
```javascript
// 隐藏控制台窗口
console.hide(consoleObj);
```

### console.clear(consoleObj)
清空控制台内容。

**入参：**
- `consoleObj`: Console 对象

**出参：** 无

**调用示例：**
```javascript
// 清空控制台内容
console.clear(consoleObj);
```

### console.isVisible(consoleObj)
检查控制台窗口是否可见。

**入参：**
- `consoleObj`: Console 对象

**出参：** 布尔值，true 表示可见，false 表示不可见

**调用示例：**
```javascript
// 检查控制台是否可见
var visible = console.isVisible(consoleObj);
console.println(consoleObj, "控制台可见:", visible);
```

### console.destroy(consoleObj)
销毁控制台对象，释放资源。

**入参：**
- `consoleObj`: Console 对象

**出参：** 无

**调用示例：**
```javascript
// 销毁控制台对象
console.destroy(consoleObj);
```

## 完整示例

```javascript
// 创建控制台对象
var consoleObj = console.new();

// 设置窗口属性
console.setWindowSize(consoleObj, 600, 400);
console.setWindowPosition(consoleObj, 100, 100);
console.setWindowColor(consoleObj, "#000000");
console.setTextSize(consoleObj, 14);
console.setTextColor(consoleObj, "#00FF00");

// 显示控制台
console.show(consoleObj);

// 输出日志信息
console.println(consoleObj, "程序启动...");
console.println(consoleObj, "当前时间:", new Date().toLocaleString());
console.println(consoleObj, "正在处理数据...");

// 等待一段时间
utils.sleep(3000);

// 清空控制台
console.clear(consoleObj);
console.println(consoleObj, "控制台已清空");

// 等待一段时间
utils.sleep(2000);

// 销毁控制台
console.destroy(consoleObj);
```

## 注意事项

1. 使用完控制台后，建议调用 `destroy()` 方法释放资源
2. 颜色字符串格式为十六进制，如 "#RRGGBB"
3. 控制台窗口的位置和大小应根据屏幕分辨率合理设置
4. 文本颜色和窗口背景颜色应有足够的对比度，确保可读性
