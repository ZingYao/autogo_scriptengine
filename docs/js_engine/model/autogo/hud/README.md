# hud 模块

## 模块简介

hud 模块提供了在屏幕上显示浮动窗口（HUD）的功能。可以创建自定义的浮动界面，用于显示调试信息、状态提示等内容。

## 方法列表

### hud.new
创建一个新的HUD对象

**参数：**

无参数

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| hud | object | HUD 对象 |

**使用示例：**
```javascript
// 创建 HUD 对象
var hud = hud.new();
hud.show();
```

---

### hud.setPosition
设置HUD的显示位置

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| hud | object | HUD 对象 |
| x1 | number | 左上角 X 坐标 |
| y1 | number | 左上角 Y 坐标 |
| x2 | number | 右下角 X 坐标 |
| y2 | number | 右下角 Y 坐标 |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| hud | object | HUD 对象（支持链式调用） |

**使用示例：**
```javascript
// 设置 HUD 位置
var hud = hud.new();
hud.setPosition(100, 100, 500, 300);
hud.show();
```

---

### hud.setBackgroundColor
设置HUD的背景颜色

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| hud | object | HUD 对象 |
| color | string | 背景颜色（格式："#RRGGBB" 或 "#AARRGGBB"） |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| hud | object | HUD 对象（支持链式调用） |

**使用示例：**
```javascript
// 设置 HUD 背景颜色
var hud = hud.new();
hud.setBackgroundColor("#FF0000");
hud.show();
```

---

### hud.setTextSize
设置HUD中文本的字体大小

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| hud | object | HUD 对象 |
| size | number | 字体大小 |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| hud | object | HUD 对象（支持链式调用） |

**使用示例：**
```javascript
// 设置 HUD 字体大小
var hud = hud.new();
hud.setTextSize(24);
hud.show();
```

---

### hud.setText
设置HUD中要显示的文本内容

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| hud | object | HUD 对象 |
| items | array | 文本项数组，每个项包含 Text 和 TextColor 属性 |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| hud | object | HUD 对象（支持链式调用） |

**使用示例：**
```javascript
// 设置 HUD 文本内容
var hud = hud.new();
hud.setText([
    {Text: "Hello", TextColor: "#FFFFFF"},
    {Text: "World", TextColor: "#FF0000"}
]);
hud.show();
```

---

### hud.show
显示HUD

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| hud | object | HUD 对象 |

**返回值：**

无返回值

**使用示例：**
```javascript
// 显示 HUD
var hud = hud.new();
hud.show();
```

---

### hud.hide
隐藏HUD

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| hud | object | HUD 对象 |

**返回值：**

无返回值

**使用示例：**
```javascript
// 隐藏 HUD
var hud = hud.new();
hud.show();
// ... 稍后隐藏
hud.hide();
```

---

### hud.isVisible
返回HUD是否可见

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| hud | object | HUD 对象 |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| visible | boolean | 是否可见 |

**使用示例：**
```javascript
// 检查 HUD 是否可见
var hud = hud.new();
hud.show();
var visible = hud.isVisible(hud);
console.log("HUD 是否可见: " + visible);
```

---

### hud.destroy
销毁HUD对象，释放资源

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| hud | object | HUD 对象 |

**返回值：**

无返回值

**使用示例：**
```javascript
// 销毁 HUD 对象
var hud = hud.new();
hud.show();
// ... 使用完毕后销毁
hud.destroy();
```

---

## 综合使用示例

### 示例1：基本使用
```javascript
// hud 模块的基本使用示例
```