# ime 模块

## 模块简介

ime 模块提供了输入法相关的功能，包括剪贴板操作、文本输入、输入法切换等。

## 方法列表

### ime.getClipText
获取剪切板内容

**参数：**

无参数

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| text | string | 剪切板内容 |

**使用示例：**
```javascript
// 获取剪切板内容
var text = ime.getClipText();
console.log("剪切板内容: " + text);
```

---

### ime.setClipText
设置剪切板内容

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| text | string | 要设置的剪切板内容 |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| success | boolean | 是否设置成功 |

**使用示例：**
```javascript
// 设置剪切板内容
var success = ime.setClipText("Hello World");
console.log("设置结果: " + success);
```

---

### ime.keyAction
模拟按键

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| code | number | 按键码 |

**返回值：**

无返回值

**使用示例：**
```javascript
// 模拟按键
ime.keyAction(4); // 返回键
```

---

### ime.inputText
输入文本

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| text | string | 要输入的文本 |
| displayId | number | 显示设备 ID（可选，默认为 0） |

**返回值：**

无返回值

**使用示例：**
```javascript
// 输入文本
ime.inputText("Hello World");
```

---

### ime.getIMEList
获取输入法列表

**参数：**

无参数

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| imeList | array | 输入法列表，每个元素包含包名和名称 |

**使用示例：**
```javascript
// 获取输入法列表
var imeList = ime.getIMEList();
for (var i = 0; i < imeList.length; i++) {
    console.log(imeList[i].packageName + ": " + imeList[i].name);
}
```

---

### ime.setCurrentIME
设置当前输入法

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| packageName | string | 输入法包名 |

**返回值：**

无返回值

**使用示例：**
```javascript
// 设置当前输入法
ime.setCurrentIME("com.example.inputmethod");
```

---

## 综合使用示例

### 示例1：基本使用
```javascript
// ime 模块的基本使用示例
```