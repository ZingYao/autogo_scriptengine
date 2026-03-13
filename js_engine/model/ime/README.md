# IME 模块

IME（Input Method Editor）模块提供了输入法相关的功能，包括剪切板操作、按键模拟、文本输入和输入法管理。

## 方法列表

### ime.getClipText()
获取剪切板内容。

**入参：** 无

**出参：** 剪切板中的文本内容（字符串）

**调用示例：**
```javascript
// 获取剪切板内容
var clipText = ime.getClipText();
console.println("剪切板内容:", clipText);
```

### ime.setClipText(text)
设置剪切板内容。

**入参：**
- `text`: 要设置的文本内容（字符串）

**出参：** 布尔值，true 表示设置成功，false 表示失败

**调用示例：**
```javascript
// 设置剪切板内容
var success = ime.setClipText("这是新的剪切板内容");
if (success) {
    console.println("剪切板设置成功");
} else {
    console.println("剪切板设置失败");
}
```

### ime.keyAction(code)
模拟按键操作。

**入参：**
- `code`: 按键码（整数）

**出参：** 无

**调用示例：**
```javascript
// 模拟按键
ime.keyAction(3); // 模拟 HOME 键
ime.keyAction(4); // 模拟 BACK 键
```

### ime.inputText(text, displayId)
输入文本到当前焦点。

**入参：**
- `text`: 要输入的文本（字符串）
- `displayId`: 显示器 ID（整数，可选，默认 0）

**出参：** 无

**调用示例：**
```javascript
// 输入文本
ime.inputText("Hello World");
ime.inputText("测试文本", 0);
```

### ime.getIMEList()
获取系统中的输入法列表。

**入参：** 无

**出参：** 输入法列表（数组）

**调用示例：**
```javascript
// 获取输入法列表
var imeList = ime.getIMEList();
console.println("输入法列表:");
for (var i = 0; i < imeList.length; i++) {
    console.println((i + 1) + ". " + imeList[i]);
}
```

### ime.setCurrentIME(packageName)
设置当前输入法。

**入参：**
- `packageName`: 输入法的包名（字符串）

**出参：** 无

**调用示例：**
```javascript
// 设置当前输入法
ime.setCurrentIME("com.android.inputmethod.pinyin");
```

## 完整示例

```javascript
// 示例1：剪切板操作
function clipboardDemo() {
    // 设置剪切板内容
    ime.setClipText("这是测试文本");
    
    // 获取剪切板内容
    var clipText = ime.getClipText();
    console.println("剪切板内容:", clipText);
}

// 示例2：文本输入
function inputTextDemo() {
    // 点击输入框
    motion.tap(500, 500);
    utils.sleep(500);
    
    // 输入文本
    ime.inputText("Hello World");
    utils.sleep(1000);
    
    // 输入更多文本
    ime.inputText("这是一段测试文本");
}

// 示例3：按键模拟
function keyActionDemo() {
    // 模拟 HOME 键
    ime.keyAction(3);
    utils.sleep(1000);
    
    // 模拟 BACK 键
    ime.keyAction(4);
    utils.sleep(1000);
    
    // 模拟 MENU 键
    ime.keyAction(82);
}

// 示例4：输入法管理
function imeManagementDemo() {
    // 获取输入法列表
    var imeList = ime.getIMEList();
    console.println("当前系统输入法:");
    for (var i = 0; i < imeList.length; i++) {
        console.println((i + 1) + ". " + imeList[i]);
    }
    
    // 切换到指定输入法
    if (imeList.length > 0) {
        ime.setCurrentIME(imeList[0]);
        console.println("已切换到:", imeList[0]);
    }
}

// 示例5：剪切板复制粘贴
function copyPasteDemo() {
    // 复制文本到剪切板
    var textToCopy = "这是要复制的文本";
    ime.setClipText(textToCopy);
    console.println("已复制:", textToCopy);
    
    // 粘贴文本
    motion.tap(500, 500); // 点击输入框
    utils.sleep(500);
    
    var clipText = ime.getClipText();
    ime.inputText(clipText);
    console.println("已粘贴:", clipText);
}

// 示例6：批量输入
function batchInputDemo() {
    var texts = [
        "第一行文本",
        "第二行文本",
        "第三行文本"
    ];
    
    for (var i = 0; i < texts.length; i++) {
        ime.inputText(texts[i]);
        ime.keyAction(66); // 模拟回车键
        utils.sleep(500);
    }
}

// 调用示例
clipboardDemo();
inputTextDemo();
keyActionDemo();
imeManagementDemo();
copyPasteDemo();
batchInputDemo();
```

## 注意事项

1. 按键码对应 Android 的 KeyEvent 常量，如 3=HOME, 4=BACK, 66=ENTER
2. 输入文本前，需要确保焦点在输入框中
3. 切换输入法可能需要系统权限
4. 剪切板操作可能受到系统限制
5. displayId 用于多显示器环境，通常为 0 表示主显示器
6. 某些按键操作可能需要 root 权限或特殊权限
7. 输入法切换可能需要用户手动确认
8. 建议在输入文本前添加适当的延迟，确保焦点已获取
