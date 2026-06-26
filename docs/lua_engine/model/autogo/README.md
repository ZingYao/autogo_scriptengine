# Lua 引擎 autogo 风格包文档

## 1. 风格包简介

autogo 风格包是 Lua 引擎的一种风格实现，基于 AutoGo 原生 API 开发。它提供了一套简洁、高效的 API 接口，方便开发者快速编写脚本。

## 2. 实现基础

autogo 风格包基于 AutoGo 原生 API 实现，主要特点包括：

- Go 包函数统一导出为脚本侧模块对象方法，例如 `app.OpenSetting` 对应 `app.openSetting`
- 方法名按 Go 导出名做常规小驼峰转换，不保留历史别名，例如不再使用 `app.openAppSetting`
- 不注册模块方法的全局入口，例如触控必须使用 `motion.click(...)`，不使用 `click(...)`
- Lua 表用于构造 Go struct、map、slice 入参；Go struct、map、slice 返回值会转换为 Lua table

## 3. 目录结构

autogo 风格包的目录结构如下：

```
autogo/
├── app/         # 应用相关操作
├── console/     # 控制台输出
├── device/      # 设备操作
├── files/       # 文件操作
├── https/       # 网络请求
├── images/      # 图像处理
├── imgui/       # ImGui GUI 库
├── motion/      # 触摸操作
├── opencv/      # 计算机视觉
├── plugin/      # 插件加载
├── ppocr/       # OCR 文字识别
├── uiacc/       # 无障碍 UI 操作
├── vdisplay/    # 虚拟显示
└── yolo/        # YOLO 目标检测
```

## 4. 注册方式

Android 与 iOS 的 Lua 注入包已按系统隔离，避免 Android 专用模块污染 iOS：

```go
// Android 全量 autogo 模块
import "github.com/ZingYao/autogo_scriptengine/lua_engine/define/android/autogo/all_models"

// iOS 全量 autogo 模块
import "github.com/ZingYao/autogo_scriptengine/lua_engine/define/ios/autogo/all_models"
```

如果只需要安全模块或非安全模块，可以把 `all_models` 替换为 `safe_models` 或 `unsafe_models`。

iOS 专用说明见 [Lua iOS autogo 概述](../autogo_ios/README.md)。iOS 当前不会注入 `uiacc`、`apkctl` 等 AutoGo iOS 参考目录不存在的模块，也不使用 Android 的 `displayId` 参数。

## 5. 使用方法

### 5.1 直接使用模块

**重要提示**：所有 autogo 风格包的模块都已通过 Go 代码注入到 Lua 全局环境中，**无需使用 require**，可以直接使用。

```lua
-- 直接使用全局模块，无需 require
-- app 模块：应用相关操作
app.launch("com.example.app", 0)
app.openSetting("com.example.app")

-- device 模块：设备信息
console.log("屏幕宽度: " .. device.width)
console.log("屏幕高度: " .. device.height)

-- console 模块：控制台输出
console.log("Hello, AutoGo!")

-- files 模块：文件操作
files.read("/sdcard/test.txt")
files.write("/sdcard/test.txt", "Hello")

-- https 模块：网络请求
local resp = https.get("https://example.com", 5000)
console.log("HTTP 状态码: " .. resp.code)

-- motion 模块：触摸操作
motion.click(100, 200)
motion.swipe(100, 500, 500, 500, 300)
```

### 5.2 struct、map、slice 与回调

```lua
-- Go struct 入参：使用 Lua table 构造字段
app.startActivity({
    action = "android.intent.action.VIEW",
    data = "https://example.com",
    packageName = app.getBrowserPackage()
})

-- Go map 入参：使用 Lua table 的字符串 key
local postResp = https.post(
    "https://example.com/api",
    '{"hello":"autogo"}',
    {["Content-Type"] = "application/json"},
    5000
)
console.log("POST 状态码: " .. postResp.code)

-- Go slice/struct 返回值：按 Lua table 解析
local apps = app.getList(false)
if #apps > 0 then
    console.log(apps[1].packageName .. " / " .. apps[1].appName)
end

-- Go callback 入参：直接传 Lua function
images.setCallback(function(x, y, color)
    console.log("image callback: " .. x .. "," .. y .. "," .. color)
end)
```

### 5.3 对象生命周期

```lua
-- uiacc、hud、vdisplay、opencv、imgui 等模块会返回对象
-- 返回对象的方法仍然挂在对象本身，不再额外注册全局别名
local acc = uiacc.new()
local node = acc.text("确定")
if node ~= nil then
    node.click()
end
```

## 6. 模块说明

### 6.1 app 模块

提供应用相关的操作，如启动、停止、切换应用等。

### 6.2 console 模块

提供控制台输出功能，如日志输出、错误提示等。

### 6.3 device 模块

提供设备信息和设备控制功能。

### 6.4 files 模块

提供文件操作功能，如读写文件、创建目录等。

### 6.5 https 模块

提供网络请求功能，如 GET、POST 请求等。

## 7. 注意事项

1. **所有模块都已通过 Go 代码注入到 Lua 全局环境中，无需使用 require**
2. 所有函数的参数和返回值与 AutoGo 原生 API 保持一致，脚本侧仅做小驼峰命名转换
3. 使用前请确保已在 Go 代码中注册了所需的模块
4. 复杂对象优先按模块返回对象继续调用，例如 `uiacc.new().text("OK").click()`
5. 详细的模块 API 文档请参考各模块的 README.md 文件
