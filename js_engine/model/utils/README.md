# Utils 模块

Utils 模块提供了常用的工具函数，包括日志记录、提示框、命令执行、类型转换和随机数生成等功能。

## 方法列表

### utils.logI(label, message)
记录一条 INFO 级别的日志。

**入参：**
- `label`: 日志标签（字符串）
- `message`: 日志消息（字符串，支持多个参数）

**出参：** 无

**调用示例：**
```javascript
// 记录 INFO 日志
utils.logI("MyApp", "应用启动");
utils.logI("MyApp", "处理数据", "完成");
```

### utils.logE(label, message)
记录一条 ERROR 级别的日志。

**入参：**
- `label`: 日志标签（字符串）
- `message`: 日志消息（字符串，支持多个参数）

**出参：** 无

**调用示例：**
```javascript
// 记录 ERROR 日志
utils.logE("MyApp", "发生错误");
utils.logE("MyApp", "文件不存在:", "/sdcard/test.txt");
```

### utils.toast(message)
显示 Toast 提示框。

**入参：**
- `message`: 提示消息（字符串）

**出参：** 无

**调用示例：**
```javascript
// 显示 Toast 提示
utils.toast("操作成功");
utils.toast("正在加载...");
```

### utils.alert(title, content, btn1Text, btn2Text)
显示 Alert 对话框。

**入参：**
- `title`: 对话框标题（字符串）
- `content`: 对话框内容（字符串）
- `btn1Text`: 第一个按钮文本（字符串，可选）
- `btn2Text`: 第二个按钮文本（字符串，可选）

**出参：** 整数，表示用户点击的按钮索引

**调用示例：**
```javascript
// 显示简单的对话框
var result = utils.alert("提示", "这是一个对话框");
console.println("用户点击了按钮:", result);

// 显示带两个按钮的对话框
var result = utils.alert("确认", "确定要删除吗？", "确定", "取消");
if (result === 0) {
    console.println("用户点击了确定");
} else {
    console.println("用户点击了取消");
}
```

### utils.shell(cmd)
执行 shell 命令并返回输出。

**入参：**
- `cmd`: shell 命令（字符串）

**出参：** 命令输出（字符串）

**调用示例：**
```javascript
// 执行 shell 命令
var output = utils.shell("ls -l /sdcard");
console.println("命令输出:", output);
```

### utils.random(min, max)
返回指定范围内的随机整数。

**入参：**
- `min`: 最小值（整数）
- `max`: 最大值（整数）

**出参：** 随机整数（在 min 和 max 之间）

**调用示例：**
```javascript
// 生成随机数
var num = utils.random(1, 100);
console.println("随机数:", num);
```

### utils.sleep(ms)
暂停当前线程指定的毫秒数。

**入参：**
- `ms`: 暂停时间（毫秒，整数）

**出参：** 无

**调用示例：**
```javascript
// 暂停 1 秒
utils.sleep(1000);

// 暂停 3 秒
utils.sleep(3000);
```

### utils.i2s(i)
将整数转换为字符串。

**入参：**
- `i`: 整数

**出参：** 字符串

**调用示例：**
```javascript
// 整数转字符串
var str = utils.i2s(123);
console.println("字符串:", str);
```

### utils.s2i(s)
将字符串转换为整数。

**入参：**
- `s`: 字符串

**出参：** 整数

**调用示例：**
```javascript
// 字符串转整数
var num = utils.s2i("456");
console.println("整数:", num);
```

### utils.f2s(f)
将浮点数转换为字符串。

**入参：**
- `f`: 浮点数

**出参：** 字符串

**调用示例：**
```javascript
// 浮点数转字符串
var str = utils.f2s(3.14);
console.println("字符串:", str);
```

### utils.s2f(s)
将字符串转换为浮点数。

**入参：**
- `s`: 字符串

**出参：** 浮点数

**调用示例：**
```javascript
// 字符串转浮点数
var num = utils.s2f("2.718");
console.println("浮点数:", num);
```

### utils.b2s(b)
将布尔值转换为字符串。

**入参：**
- `b`: 布尔值

**出参：** 字符串（"true" 或 "false"）

**调用示例：**
```javascript
// 布尔值转字符串
var str = utils.b2s(true);
console.println("字符串:", str);
```

### utils.s2b(s)
将字符串转换为布尔值。

**入参：**
- `s`: 字符串（"true"、"false"、"1"、"0" 等）

**出参：** 布尔值

**调用示例：**
```javascript
// 字符串转布尔值
var bool = utils.s2b("true");
console.println("布尔值:", bool);
```

## 完整示例

```javascript
// 示例1：日志记录
function loggingDemo() {
    utils.logI("MyApp", "应用启动");
    utils.logI("MyApp", "正在初始化...");
    utils.logI("MyApp", "初始化完成");
    
    utils.logE("MyApp", "发生错误");
    utils.logE("MyApp", "文件不存在:", "/sdcard/test.txt");
}

// 示例2：用户提示
function userPromptDemo() {
    // 显示 Toast
    utils.toast("开始处理...");
    utils.sleep(1000);
    utils.toast("处理完成");
    
    // 显示对话框
    var result = utils.alert("提示", "操作已完成", "确定");
    console.println("用户点击了按钮:", result);
    
    // 确认对话框
    var result2 = utils.alert("确认", "确定要退出吗？", "退出", "取消");
    if (result2 === 0) {
        console.println("用户选择退出");
    } else {
        console.println("用户取消退出");
    }
}

// 示例3：命令执行
function commandExecutionDemo() {
    // 列出文件
    var output = utils.shell("ls -l /sdcard");
    console.println("文件列表:");
    console.println(output);
    
    // 获取系统信息
    var output2 = utils.shell("getprop ro.build.version.release");
    console.println("Android 版本:", output2.trim());
}

// 示例4：随机数生成
function randomDemo() {
    for (var i = 0; i < 10; i++) {
        var num = utils.random(1, 100);
        console.println("随机数:", num);
    }
}

// 示例5：延时操作
function delayDemo() {
    console.println("开始执行");
    utils.sleep(1000);
    console.println("1 秒后");
    utils.sleep(2000);
    console.println("3 秒后");
    utils.sleep(3000);
    console.println("6 秒后");
}

// 示例6：类型转换
function typeConversionDemo() {
    // 整数转换
    var i = 123;
    var s = utils.i2s(i);
    var i2 = utils.s2i(s);
    console.println("整数:", i, "-> 字符串:", s, "-> 整数:", i2);
    
    // 浮点数转换
    var f = 3.14;
    var s2 = utils.f2s(f);
    var f2 = utils.s2f(s2);
    console.println("浮点数:", f, "-> 字符串:", s2, "-> 浮点数:", f2);
    
    // 布尔值转换
    var b = true;
    var s3 = utils.b2s(b);
    var b2 = utils.s2b(s3);
    console.println("布尔值:", b, "-> 字符串:", s3, "-> 布尔值:", b2);
}

// 示例7：进度显示
function progressDemo() {
    var total = 10;
    for (var i = 1; i <= total; i++) {
        var progress = Math.floor(i / total * 100);
        utils.toast("进度: " + progress + "%");
        utils.logI("Progress", "处理进度:", progress + "%");
        utils.sleep(500);
    }
    utils.toast("处理完成");
}

// 示例8：错误处理
function errorHandlingDemo() {
    try {
        utils.logI("MyApp", "开始执行");
        
        // 模拟错误
        var result = 1 / 0;
        
        utils.logI("MyApp", "执行完成");
    } catch (e) {
        utils.logE("MyApp", "捕获到异常:", e.message);
        utils.alert("错误", "发生错误: " + e.message, "确定");
    }
}

// 示例9：用户输入验证
function inputValidationDemo() {
    var input = "123";
    
    // 验证是否为数字
    var num = utils.s2i(input);
    if (!isNaN(num)) {
        utils.logI("Validation", "输入是有效的数字:", num);
    } else {
        utils.logE("Validation", "输入不是有效的数字");
    }
    
    // 验证布尔值
    var boolStr = "true";
    var bool = utils.s2b(boolStr);
    utils.logI("Validation", "布尔值:", bool);
}

// 示例10：定时任务
function scheduledTask() {
    var count = 0;
    var interval = setInterval(function() {
        count++;
        utils.logI("Task", "执行第 " + count + " 次");
        utils.toast("任务执行中: " + count);
        
        if (count >= 5) {
            clearInterval(interval);
            utils.toast("任务完成");
        }
    }, 2000);
}

// 调用示例
loggingDemo();
userPromptDemo();
commandExecutionDemo();
randomDemo();
delayDemo();
typeConversionDemo();
progressDemo();
errorHandlingDemo();
inputValidationDemo();
scheduledTask();
```

## 注意事项

1. 日志记录会输出到系统日志，可以通过 logcat 查看
2. Toast 提示会在短暂时间后自动消失
3. Alert 对话框会阻塞当前线程，直到用户点击按钮
4. 执行 shell 命令需要相应的权限
5. random 函数返回的随机数包含 min 和 max
6. sleep 函数会阻塞当前线程，建议合理设置延时
7. 类型转换函数会尝试转换，失败可能返回默认值
8. s2b 函数支持 "true"、"false"、"1"、"0" 等多种格式
9. 建议在关键操作前后添加日志，便于调试
10. 使用 shell 命令时要注意安全性，避免执行危险命令
