# console 模块

## 模块简介

console 模块提供了控制台窗口的功能，用于显示日志和调试信息。

## 方法列表

### console.new
创建控制台对象

**参数：**

无参数

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| console | object | 控制台对象 |

**使用示例：**
```javascript
// 创建控制台对象
var console = console.new();
console.show();
```

***

### console.println
打印多个参数到控制台

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| console | object | 控制台对象 |
| ...args | any | 要打印的参数列表（可变参数） |

**返回值：**

无返回值

**使用示例：**
```javascript
// 打印多个参数到控制台
var console = console.new();
console.println(console, "Hello", "World", 123);
```

***

### console.setTextSize
设置文本大小

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| size | number | 文本大小 |

**返回值：**

无返回值

**使用示例：**
```javascript
// 设置文本大小
console.setTextSize(16);
```

***

### console.setTextColor
设置文本颜色

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| color | string | 文本颜色（如 "#FF0000"） |

**返回值：**

无返回值

**使用示例：**
```javascript
// 设置文本颜色
console.setTextColor("#FF0000");
```

***

### console.setWindowSize
设置窗口大小

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| width | number | 窗口宽度 |
| height | number | 窗口高度 |

**返回值：**

无返回值

**使用示例：**
```javascript
// 设置窗口大小
console.setWindowSize(800, 600);
```

***

### console.setWindowPosition
设置窗口位置

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| x | number | 窗口 X 坐标 |
| y | number | 窗口 Y 坐标 |

**返回值：**

无返回值

**使用示例：**
```javascript
// 设置窗口位置
console.setWindowPosition(100, 100);
```

***

### console.setWindowColor
设置窗口颜色

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| color | string | 窗口颜色（如 "#00FF00"） |

**返回值：**

无返回值

**使用示例：**
```javascript
// 设置窗口颜色
console.setWindowColor("#00FF00");
```

***

### console.show
显示控制台

**参数：**

无参数

**返回值：**

无返回值

**使用示例：**
```javascript
// 显示控制台
console.show();
```

***

### console.hide
隐藏控制台

**参数：**

无参数

**返回值：**

无返回值

**使用示例：**
```javascript
// 隐藏控制台
console.hide();
```

***

### console.clear
清空控制台

**参数：**

无参数

**返回值：**

无返回值

**使用示例：**
```javascript
// 清空控制台
console.clear();
```

***

### console.isVisible
检查控制台是否可见

**参数：**

无参数

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| visible | boolean | 控制台可见返回 true，否则返回 false |

**使用示例：**
```javascript
// 检查控制台是否可见
var visible = console.isVisible();
console.log("控制台可见: " + visible);
```

***

### console.destroy
销毁控制台

**参数：**

无参数

**返回值：**

无返回值

**使用示例：**
```javascript
// 销毁控制台
console.destroy();
```

### 示例1：基本使用
```javascript
// console 模块的基本使用示例
```