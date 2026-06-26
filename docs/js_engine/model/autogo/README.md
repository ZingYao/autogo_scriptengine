# JavaScript 引擎 Android autogo 风格包文档

## 1. 风格包简介

Android autogo 风格包是 JavaScript 引擎的一种 Android 风格实现，基于 AutoGo Android 原生 API 开发。它提供了一套简洁、高效的 API 接口，方便开发者快速编写 Android 自动化脚本。

## 2. 实现基础

autogo 风格包基于 AutoGo 原生 API 实现，主要特点包括：

- Go 包函数统一导出为脚本侧模块对象方法，例如 `app.OpenSetting` 对应 `app.openSetting`
- 方法名按 Go 导出名做常规小驼峰转换，不保留历史别名，例如不再使用 `app.openAppSetting`
- 不注册模块方法的全局入口，例如触控必须使用 `motion.click(...)`，不使用 `click(...)`
- JavaScript object 用于构造 Go struct、map 入参，array 用于构造 Go slice 入参；Go struct、map、slice 返回值会转换为 JavaScript object/array

## 3. 目录结构

Android autogo 风格包的目录结构如下：

```
autogo/
├── app/         # 应用相关操作
├── console/     # 控制台输出
├── coroutine/   # 协程支持
├── device/      # 设备操作
├── dotocr/      # 点字 OCR 识别
├── files/       # 文件操作
├── https/       # 网络请求
├── hud/         # HUD 悬浮显示
├── images/      # 图像处理
├── ime/         # 输入法控制
├── imgui/       # ImGui GUI 库
├── media/       # 媒体控制
├── motion/      # 触摸操作
├── opencv/      # 计算机视觉
├── plugin/      # 插件加载
├── ppocr/       # OCR 文字识别
├── rhino/       # JavaScript 执行引擎
├── storages/    # 数据存储
├── system/      # 系统功能
├── uiacc/       # 无障碍 UI 操作
├── utils/       # 工具方法
├── vdisplay/    # 虚拟显示
├── websocket/   # WebSocket 通信
└── yolo/        # YOLO 目标检测
```

## 4. 注册方式

JavaScript 的 autogo define 已按系统隔离。Android 项目使用 Android define，iOS 项目使用 iOS define：

```go
// Android 全量 autogo 模块
import "github.com/ZingYao/autogo_scriptengine/js_engine/define/android/autogo/all_models"

// iOS 全量 autogo 模块
import "github.com/ZingYao/autogo_scriptengine/js_engine/define/ios/autogo/all_models"
```

如果只需要安全模块或非安全模块，可以把 `all_models` 替换为 `safe_models` 或 `unsafe_models`。

iOS 专用说明见 [JavaScript iOS autogo 概述](../autogo_ios/README.md)。iOS 当前不会注入 `uiacc`、`apkctl` 等 AutoGo iOS 参考目录不存在的模块，也不使用 Android 的 `displayId` 参数。

## 5. 使用方法

### 5.1 直接使用模块

**重要提示**：所有 autogo 风格包的模块都已通过 Go 代码注入到 JavaScript 全局环境中，**无需使用 require**，可以直接使用。

```javascript
// 直接使用全局模块，无需 require
// app 模块：应用相关操作
app.launch("com.example.app", 0);
app.openSetting("com.example.app");

// device 模块：设备信息
console.log("屏幕宽度: " + device.width);
console.log("屏幕高度: " + device.height);

// console 模块：控制台输出
console.log("Hello, AutoGo!");

// files 模块：文件操作
files.read("/sdcard/test.txt");
files.write("/sdcard/test.txt", "Hello");

// https 模块：网络请求
const resp = https.get("https://example.com", 5000);
console.log("HTTP 状态码: " + resp.code);

// motion 模块：触摸操作
motion.click(100, 200);
motion.swipe(100, 500, 500, 500, 300);
```

### 5.2 struct、map、slice 与回调

```javascript
// Go struct 入参：使用 JavaScript object 构造字段
app.startActivity({
    action: "android.intent.action.VIEW",
    data: "https://example.com",
    packageName: app.getBrowserPackage()
});

// Go map 入参：使用普通 object 的字符串 key
const postResp = https.post(
    "https://example.com/api",
    JSON.stringify({ hello: "autogo" }),
    { "Content-Type": "application/json" },
    5000
);
console.log("POST 状态码: " + postResp.code);

// Go slice/struct 返回值：按 JavaScript array/object 解析
const apps = app.getList(false);
if (apps.length > 0) {
    console.log(apps[0].packageName + " / " + apps[0].appName);
}

// Go callback 入参：直接传 JavaScript function
images.setCallback(function(x, y, color) {
    console.log("image callback: " + x + "," + y + "," + color);
});
```

### 5.3 对象生命周期

```javascript
// uiacc、hud、vdisplay、opencv、imgui 等模块会返回对象
// 返回对象的方法仍然挂在对象本身，不再额外注册全局别名
const acc = uiacc.new();
const node = acc.text("确定");
if (node) {
    node.click();
}
```

## 6. 注意事项

1. **所有 autogo 模块都已通过 Go 代码注入到 JavaScript 全局环境中，无需使用 require**
2. 所有函数的参数和返回值与 AutoGo 原生 API 保持一致，脚本侧仅做小驼峰命名转换
3. 使用前请确保已在 Go 代码中注册了所需的模块
4. 复杂对象优先按模块返回对象继续调用，例如 `uiacc.new().text("OK").click()`
5. 详细的模块 API 文档请参考各模块的 README.md 文件
