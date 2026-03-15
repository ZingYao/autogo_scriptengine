# 关于 AutoGo ScriptEngine

AutoGo ScriptEngine 是一个为 [AutoGo](https://github.com/Dasongzi1366/AutoGo) 提供脚本语言支持的扩展方案，让开发者可以使用 JavaScript 或 Lua 编写自动化任务。

## 项目简介

AutoGo ScriptEngine 通过封装 AutoGo 的底层 API，为开发者提供了简单易用的脚本接口。开发者无需深入了解 Go 语言和 Android 开发，只需使用熟悉的 JavaScript 或 Lua 脚本语言即可快速开发自动化脚本。

## 核心特性

- **双引擎支持**：同时支持 JavaScript 和 Lua 脚本语言
- **丰富的 API**：提供应用管理、设备控制、图像识别、OCR 等多种功能
- **模块化设计**：支持动态注册、重写和恢复方法
- **协程支持**：Lua 引擎支持协程操作
- **文档完善**：提供完整的 API 文档和使用示例
- **跨平台**：支持 Windows、macOS、Linux 开发环境

## 技术栈

- **Go 1.25.0**：核心开发语言
- **Goja**：JavaScript 执行引擎
- **Gopher-Lua**：Lua 执行引擎
- **AutoGo**：底层 Android 自动化框架

## 主要功能模块

- **app**：应用管理（启动、安装、卸载、强制停止等）
- **device**：设备信息（分辨率、SDK 版本、屏幕方向等）
- **motion**：触摸操作（点击、滑动、手势等）
- **files**：文件操作（读写、复制、删除等）
- **images**：图像处理（截图、找色、找图等）
- **storages**：数据存储（键值对存储）
- **system**：系统功能（剪贴板、通知等）
- **http**：网络请求（GET、POST 等）
- **media**：媒体控制（音量、播放等）
- **opencv**：计算机视觉（图像处理、特征检测等）
- **ppocr**：OCR 文字识别
- **console**：控制台窗口（显示、隐藏、日志输出等）
- **dotocr**：点字 OCR 识别（基于字库的 OCR）
- **hud**：HUD 悬浮显示（脚本状态显示等）
- **ime**：输入法控制（剪切板、文本输入等）
- **plugin**：插件加载（加载外部 APK 调用 Java 方法）
- **rhino**：JavaScript 执行引擎（Rhino）
- **uiacc**：无障碍 UI 操作（控件查找、点击、输入等）
- **utils**：工具方法（日志、Toast、类型转换等）
- **vdisplay**：虚拟显示（虚拟屏操作）
- **yolo**：YOLO 目标检测（v5/v8 模型）
- **imgui**：Dear ImGui GUI 库（窗口、按钮、输入框等控件）
- **coroutine**：协程支持

## 快速开始

### 安装

```bash
go get github.com/ZingYao/autogo_scriptengine@v0.0.9
```

### JavaScript 示例

```go
package main

import (
    "github.com/ZingYao/autogo_scriptengine/js_engine"
    jsAppModel "github.com/ZingYao/autogo_scriptengine/js_engine/model/app"
    jsDeviceModel "github.com/ZingYao/autogo_scriptengine/js_engine/model/device"
)

func init() {
    // 注册需要的模块
    js_engine.RegisterModule(
        jsAppModel.AppModule{},
        jsDeviceModel.DeviceModule{},
    )
}

func main() {
    engine := js_engine.GetEngine()
    defer js_engine.Close()

    err := engine.ExecuteString(`
        console.log("Hello, JavaScript!");
        const packageName = app.currentPackage();
        console.log("当前应用包名: " + packageName);
    `)
    if err != nil {
        panic(err)
    }
}
```

### Lua 示例

```go
package main

import (
    "github.com/ZingYao/autogo_scriptengine/lua_engine"
    luaAppModel "github.com/ZingYao/autogo_scriptengine/lua_engine/model/app"
    luaDeviceModel "github.com/ZingYao/autogo_scriptengine/lua_engine/model/device"
)

func init() {
    // 注册需要的模块
    lua_engine.RegisterModule(
        luaAppModel.AppModule{},
        luaDeviceModel.DeviceModule{},
    )
}

func main() {
    engine := lua_engine.GetEngine()
    defer lua_engine.Close()

    err := engine.ExecuteString(`
        console.log("Hello, Lua!")
        local packageName = app.currentPackage()
        console.log("当前应用包名: " .. packageName)
    `)
    if err != nil {
        panic(err)
    }
}
```

## 文档

- [📖 项目 README](./README.md) - 项目介绍和快速开始指南
- [📜 JavaScript 引擎文档](./js_engine/README.md) - JavaScript 引擎完整文档
- [📜 Lua 引擎文档](./lua_engine/README.md) - Lua 引擎完整文档
- [🌐 HTML 在线文档](./docs/index.html) - 美观的 HTML 在线文档
- [📋 更新日志](./UPDATE.md) - 版本更新记录

## 兼容性说明

### Windows 开发环境

在 Windows 环境下开发时，如果引入了超过 1 个以上的带 C 依赖的库，可能会导致编译命令过长。

**解决方案**：
1. 避免过多使用带 C 的库
2. 只注册刚需模块，使用 `RegisterModule` 函数手动注册
3. 切换到 macOS 或 Linux 系统进行编译

### Android 版本兼容性

某些包在某些 Android 版本下会有内存引用错误问题。

**解决方案**：
1. 遇到问题时，减少注册的模块
2. 只使用核心模块（app、device、motion、files、console）

## 许可证

本项目采用 MIT 许可证，详见 [LICENSE](./LICENSE) 文件。

## 贡献

欢迎提交 Issue 和 Pull Request！

1. Fork 本仓库
2. 创建特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 提交 Pull Request

## 联系方式

- **GitHub Issues**: [提交问题](https://github.com/ZingYao/autogo_scriptengine/issues)
- **项目主页**: [GitHub 仓库](https://github.com/ZingYao/autogo_scriptengine)

## 致谢

感谢 [AutoGo](https://github.com/Dasongzi1366/AutoGo) 项目提供的底层支持！
