# Uiacc 模块

Uiacc（UI Accessibility）模块提供了无障碍服务功能，用于查找和操作界面元素。

## Uiacc 选择器方法

### uiacc.new(displayId)
创建一个新的 Uiacc 对象。

**入参：**
- `displayId`: 显示器 ID（整数，可选，默认 0）

**出参：** Uiacc 对象

**调用示例：**
```javascript
// 创建 Uiacc 对象
var u = uiacc.new(0);
```

### uiacc.text(u, value)
根据文本内容选择元素。

**入参：**
- `u`: Uiacc 对象
- `value`: 文本内容（字符串）

**出参：** Uiacc 对象（支持链式调用）

**调用示例：**
```javascript
var u = uiacc.new().text("设置");
```

### uiacc.textContains(u, value)
根据文本包含内容选择元素。

**入参：**
- `u`: Uiacc 对象
- `value`: 文本包含内容（字符串）

**出参：** Uiacc 对象（支持链式调用）

**调用示例：**
```javascript
var u = uiacc.new().textContains("设置");
```

### uiacc.textStartsWith(u, value)
根据文本前缀选择元素。

**入参：**
- `u`: Uiacc 对象
- `value`: 文本前缀（字符串）

**出参：** Uiacc 对象（支持链式调用）

**调用示例：**
```javascript
var u = uiacc.new().textStartsWith("设置");
```

### uiacc.textEndsWith(u, value)
根据文本后缀选择元素。

**入参：**
- `u`: Uiacc 对象
- `value`: 文本后缀（字符串）

**出参：** Uiacc 对象（支持链式调用）

**调用示例：**
```javascript
var u = uiacc.new().textEndsWith("设置");
```

### uiacc.textMatches(u, value)
根据正则表达式选择元素。

**入参：**
- `u`: Uiacc 对象
- `value`: 正则表达式（字符串）

**出参：** Uiacc 对象（支持链式调用）

**调用示例：**
```javascript
var u = uiacc.new().textMatches(".*设置.*");
```

### uiacc.desc(u, value)
根据描述内容选择元素。

**入参：**
- `u`: Uiacc 对象
- `value`: 描述内容（字符串）

**出参：** Uiacc 对象（支持链式调用）

**调用示例：**
```javascript
var u = uiacc.new().desc("设置按钮");
```

### uiacc.id(u, value)
根据 ID 选择元素。

**入参：**
- `u`: Uiacc 对象
- `value`: 元素 ID（字符串）

**出参：** Uiacc 对象（支持链式调用）

**调用示例：**
```javascript
var u = uiacc.new().id("com.android.settings:id/button");
```

### uiacc.className(u, value)
根据类名选择元素。

**入参：**
- `u`: Uiacc 对象
- `value`: 类名（字符串）

**出参：** Uiacc 对象（支持链式调用）

**调用示例：**
```javascript
var u = uiacc.new().className("android.widget.Button");
```

### uiacc.bounds(u, left, top, right, bottom)
根据边界选择元素。

**入参：**
- `u`: Uiacc 对象
- `left`: 左边界（整数）
- `top`: 上边界（整数）
- `right`: 右边界（整数）
- `bottom`: 下边界（整数）

**出参：** Uiacc 对象（支持链式调用）

**调用示例：**
```javascript
var u = uiacc.new().bounds(0, 0, 500, 500);
```

### uiacc.clickable(u, value)
根据是否可点击选择元素。

**入参：**
- `u`: Uiacc 对象
- `value`: 是否可点击（布尔值）

**出参：** Uiacc 对象（支持链式调用）

**调用示例：**
```javascript
var u = uiacc.new().clickable(true);
```

## Uiacc 查找方法

### uiacc.findOnce(u)
查找第一个匹配的元素。

**入参：**
- `u`: Uiacc 对象

**出参：** UiObject 对象，未找到返回 null

**调用示例：**
```javascript
var u = uiacc.new().text("设置");
var obj = uiacc.findOnce(u);
if (obj) {
    console.println("找到元素");
}
```

### uiacc.find(u)
查找所有匹配的元素。

**入参：**
- `u`: Uiacc 对象

**出参：** UiObject 数组

**调用示例：**
```javascript
var u = uiacc.new().className("android.widget.Button");
var objs = uiacc.find(u);
console.println("找到", objs.length, "个按钮");
```

### uiacc.click(u, text)
点击指定文本的元素。

**入参：**
- `u`: Uiacc 对象
- `text`: 要点击的文本（字符串）

**出参：** 布尔值，true 表示成功，false 表示失败

**调用示例：**
```javascript
var u = uiacc.new();
var success = uiacc.click(u, "设置");
if (success) {
    console.println("点击成功");
}
```

### uiacc.waitFor(u, timeout)
等待元素出现。

**入参：**
- `u`: Uiacc 对象
- `timeout`: 超时时间（毫秒，整数）

**出参：** UiObject 对象，超时返回 null

**调用示例：**
```javascript
var u = uiacc.new().text("设置");
var obj = uiacc.waitFor(u, 5000);
if (obj) {
    console.println("元素已出现");
} else {
    console.println("等待超时");
}
```

### uiacc.release(u)
释放 Uiacc 对象资源。

**入参：**
- `u`: Uiacc 对象

**出参：** 无

**调用示例：**
```javascript
var u = uiacc.new();
// 使用 u...
uiacc.release(u);
```

## UiObject 方法

### uiacc.objClick(obj)
点击元素。

**入参：**
- `obj`: UiObject 对象

**出参：** 布尔值，true 表示成功，false 表示失败

**调用示例：**
```javascript
var u = uiacc.new().text("设置");
var obj = uiacc.findOnce(u);
if (obj) {
    uiacc.objClick(obj);
}
```

### uiacc.clickCenter(obj)
点击元素中心。

**入参：**
- `obj`: UiObject 对象

**出参：** 布尔值，true 表示成功，false 表示失败

**调用示例：**
```javascript
var u = uiacc.new().text("设置");
var obj = uiacc.findOnce(u);
if (obj) {
    uiacc.clickCenter(obj);
}
```

### uiacc.getText(obj)
获取元素文本。

**入参：**
- `obj`: UiObject 对象

**出参：** 元素文本（字符串）

**调用示例：**
```javascript
var u = uiacc.new().className("android.widget.TextView");
var obj = uiacc.findOnce(u);
if (obj) {
    var text = uiacc.getText(obj);
    console.println("元素文本:", text);
}
```

### uiacc.getBounds(obj)
获取元素边界。

**入参：**
- `obj`: UiObject 对象

**出参：** 包含 left、top、right、bottom 的对象

**调用示例：**
```javascript
var u = uiacc.new().text("设置");
var obj = uiacc.findOnce(u);
if (obj) {
    var bounds = uiacc.getBounds(obj);
    console.println("边界:", bounds.left, bounds.top, bounds.right, bounds.bottom);
}
```

## 完整示例

```javascript
// 示例1：查找并点击元素
function findAndClick() {
    var u = uiacc.new().text("设置");
    var obj = uiacc.findOnce(u);
    
    if (obj) {
        console.println("找到设置按钮");
        uiacc.objClick(obj);
    } else {
        console.println("未找到设置按钮");
    }
}

// 示例2：等待元素出现
function waitForElement() {
    var u = uiacc.new().text("加载完成");
    var obj = uiacc.waitFor(u, 10000);
    
    if (obj) {
        console.println("元素已出现");
    } else {
        console.println("等待超时");
    }
}

// 示例3：查找多个元素
function findMultipleElements() {
    var u = uiacc.new().className("android.widget.Button");
    var objs = uiacc.find(u);
    
    console.println("找到", objs.length, "个按钮");
    for (var i = 0; i < objs.length; i++) {
        var text = uiacc.getText(objs[i]);
        console.println("按钮", i + 1, ":", text);
    }
}

// 示例4：使用多个条件
function useMultipleConditions() {
    var u = uiacc.new()
        .text("设置")
        .className("android.widget.Button")
        .clickable(true);
    
    var obj = uiacc.findOnce(u);
    if (obj) {
        console.println("找到符合条件的元素");
        uiacc.objClick(obj);
    }
}

// 示例5：获取元素信息
function getElementInfo() {
    var u = uiacc.new().text("设置");
    var obj = uiacc.findOnce(u);
    
    if (obj) {
        var text = uiacc.getText(obj);
        var bounds = uiacc.getBounds(obj);
        var clickable = uiacc.getClickable(obj);
        
        console.println("文本:", text);
        console.println("边界:", bounds.left, bounds.top, bounds.right, bounds.bottom);
        console.println("可点击:", clickable);
    }
}

// 示例6：点击指定文本
function clickByText(targetText) {
    var u = uiacc.new();
    var success = uiacc.click(u, targetText);
    
    if (success) {
        console.println("点击成功:", targetText);
    } else {
        console.println("点击失败:", targetText);
    }
}

// 示例7：遍历所有元素
function traverseAllElements() {
    var u = uiacc.new();
    var objs = uiacc.find(u);
    
    console.println("共有", objs.length, "个元素");
    for (var i = 0; i < objs.length; i++) {
        var obj = objs[i];
        var text = uiacc.getText(obj);
        var className = uiacc.getClassName(obj);
        console.println(i + ":", className, "-", text);
    }
}

// 调用示例
findAndClick();
waitForElement();
findMultipleElements();
useMultipleConditions();
getElementInfo();
clickByText("设置");
traverseAllElements();
```

## 注意事项

1. 使用 Uiacc 功能需要启用无障碍服务
2. 选择器方法支持链式调用，可以组合多个条件
3. findOnce 只返回第一个匹配的元素
4. find 返回所有匹配的元素数组
5. waitFor 会阻塞当前线程直到元素出现或超时
6. 使用完毕后，建议调用 release 释放资源
7. 某些应用可能限制了无障碍服务的访问
8. 元素属性可能会动态变化，建议实时获取
9. 点击操作可能需要适当的延迟
10. 建议在操作前检查元素是否存在
