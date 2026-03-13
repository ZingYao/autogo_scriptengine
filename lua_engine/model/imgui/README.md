# ImGui 模块

ImGui 模块提供了在 Android 平台上显示即时模式图形界面的功能，包括对话框、提示框和图形绘制。

**注意：** 此模块仅在 Android 平台上可用。

## 方法列表

### imgui.init()
初始化 ImGui 系统。

**入参：** 无

**出参：** 无

**调用示例：**
```javascript
// 初始化 ImGui
imgui.init();
```

### imgui.close()
关闭 ImGui 系统。

**入参：** 无

**出参：** 无

**调用示例：**
```javascript
// 关闭 ImGui
imgui.close();
```

### imgui.alert(title, content, btn1Text, btn2Text)
显示一个 Alert 对话框。

**入参：**
- `title`: 对话框标题（字符串）
- `content`: 对话框内容（字符串）
- `btn1Text`: 第一个按钮文本（字符串，可选）
- `btn2Text`: 第二个按钮文本（字符串，可选）

**出参：** 整数，表示用户点击的按钮索引

**调用示例：**
```javascript
// 显示简单的对话框
var result = imgui.alert("提示", "这是一个对话框");
console.println("用户点击了按钮:", result);

// 显示带两个按钮的对话框
var result = imgui.alert("确认", "确定要删除吗？", "确定", "取消");
if (result === 0) {
    console.println("用户点击了确定");
} else {
    console.println("用户点击了取消");
}
```

### imgui.toast(message)
显示一个 Toast 提示框。

**入参：**
- `message`: 提示消息（字符串）

**出参：** 无

**调用示例：**
```javascript
// 显示 Toast 提示
imgui.toast("操作成功");
imgui.toast("正在加载...");
```

### imgui.drawRect(x1, y1, x2, y2, colorStr, thickness)
在屏幕上绘制一个矩形。

**入参：**
- `x1`: 矩形左上角 X 坐标（整数）
- `y1`: 矩形左上角 Y 坐标（整数）
- `x2`: 矩形右下角 X 坐标（整数）
- `y2`: 矩形右下角 Y 坐标（整数）
- `colorStr`: 颜色字符串（十六进制格式，如 "#FF0000"）
- `thickness`: 线条粗细（浮点数，可选，默认 1.0）

**出参：** 无

**调用示例：**
```javascript
// 绘制红色矩形
imgui.drawRect(100, 100, 300, 200, "#FF0000", 2.0);

// 绘制绿色矩形（默认粗细）
imgui.drawRect(400, 100, 600, 200, "#00FF00");
```

## 完整示例

```javascript
// 示例1：初始化和关闭 ImGui
function initImGuiDemo() {
    imgui.init();
    console.println("ImGui 已初始化");
    
    // 执行一些操作...
    
    imgui.close();
    console.println("ImGui 已关闭");
}

// 示例2：显示对话框
function alertDialogDemo() {
    // 简单对话框
    imgui.alert("提示", "欢迎使用 ImGui");
    
    // 确认对话框
    var result = imgui.alert("删除确认", "确定要删除此文件吗？", "删除", "取消");
    if (result === 0) {
        console.println("用户确认删除");
        // 执行删除操作
    } else {
        console.println("用户取消删除");
    }
    
    // 信息对话框
    imgui.alert("系统信息", "当前版本: 1.0.0\n作者: AutoGo Team", "确定");
}

// 示例3：显示 Toast 提示
function toastDemo() {
    imgui.toast("开始处理...");
    utils.sleep(2000);
    imgui.toast("处理完成");
    utils.sleep(1000);
    imgui.toast("操作成功");
}

// 示例4：绘制图形
function drawRectDemo() {
    // 绘制不同颜色的矩形
    imgui.drawRect(50, 50, 150, 150, "#FF0000", 2.0);   // 红色
    imgui.drawRect(200, 50, 300, 150, "#00FF00", 2.0);  // 绿色
    imgui.drawRect(350, 50, 450, 150, "#0000FF", 2.0);  // 蓝色
    
    // 绘制不同粗细的矩形
    imgui.drawRect(50, 200, 150, 300, "#FFFF00", 1.0); // 黄色，细线
    imgui.drawRect(200, 200, 300, 300, "#FF00FF", 3.0); // 紫色，粗线
    imgui.drawRect(350, 200, 450, 300, "#00FFFF", 5.0); // 青色，更粗
}

// 示例5：绘制边框
function drawBorderDemo() {
    var margin = 20;
    var x1 = margin;
    var y1 = margin;
    var x2 = device.getScreenWidth() - margin;
    var y2 = device.getScreenHeight() - margin;
    
    // 绘制屏幕边框
    imgui.drawRect(x1, y1, x2, y2, "#FFFFFF", 3.0);
}

// 示例6：绘制区域标记
function drawRegionDemo() {
    // 标记重要区域
    imgui.drawRect(100, 300, 500, 600, "#FF0000", 2.0);
    imgui.toast("已标记重要区域");
    
    // 标记安全区域
    imgui.drawRect(600, 300, 1000, 600, "#00FF00", 2.0);
    imgui.toast("已标记安全区域");
}

// 示例7：用户确认流程
function confirmActionDemo() {
    imgui.toast("开始执行操作...");
    
    var confirmed = imgui.alert("确认", "确定要执行此操作吗？", "确定", "取消");
    
    if (confirmed === 0) {
        imgui.toast("操作已确认，正在执行...");
        // 执行操作
        utils.sleep(2000);
        imgui.toast("操作完成");
    } else {
        imgui.toast("操作已取消");
    }
}

// 调用示例
initImGuiDemo();
alertDialogDemo();
toastDemo();
drawRectDemo();
drawBorderDemo();
drawRegionDemo();
confirmActionDemo();
```

## 注意事项

1. 此模块仅在 Android 平台上可用，其他平台调用会报错
2. 使用 ImGui 功能前，必须先调用 `init()` 方法
3. 使用完毕后，建议调用 `close()` 方法释放资源
4. 颜色字符串格式为十六进制，如 "#RRGGBB"
5. alert 对话框会阻塞当前线程，直到用户点击按钮
6. toast 提示框会在短暂时间后自动消失
7. drawRect 的坐标系统以屏幕左上角为原点 (0, 0)
8. thickness 参数控制线条粗细，值越大线条越粗
9. 在多显示器环境中，绘制操作默认在主显示器上
10. 建议在绘制图形时考虑屏幕分辨率，确保在不同设备上正常显示
