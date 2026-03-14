# AutoGo JavaScript Engine

AutoGo JavaScript Engine 是一个高性能的 JavaScript 脚本引擎，为 AutoGo 框架提供了完整的 JavaScript 脚本支持。

## 特性

- **完整的 API 注入**: 将 AutoGo 的所有功能模块注入到 JavaScript 引擎中
- **模块化设计**: 支持按需加载模块
- **线程安全**: 所有操作都是线程安全的
- **文件系统支持**: 支持真实文件系统和虚拟文件系统 (embed.FS)
- **模块导入**: 支持 CommonJS 风格的 require 功能

## 快速开始

### 初始化引擎

```go
import "github.com/ZingYao/autogo_scriptengine/js_engine"

func main() {
    // 使用默认配置创建引擎（自动注入所有方法）
    engine := js_engine.NewEngine(&js_engine.DefaultConfig())
    defer engine.Close()
    
    // 使用引擎...
}
```

### 引擎配置选项

```go
import (
    "embed"
    "os"
    "github.com/ZingYao/autogo_scriptengine/js_engine"
)

//go:embed scripts/*
var scriptsFS embed.FS

func main() {
    // 使用自定义配置
    config := js_engine.DefaultConfig()
    config.AutoInjectMethods = false  // 禁用自动注入
    config.FileSystem = scriptsFS      // 设置虚拟文件系统
    
    engine := js_engine.NewEngine(&config)
    defer engine.Close()
    
    // 按需注入模块
    engine.InjectModule("app")
    engine.InjectModule("device")
}
```

### 执行 JavaScript 代码

```go
// 执行 JavaScript 字符串
err := engine.ExecuteString(`
    console.log("Hello, JavaScript!");
    const packageName = app.currentPackage();
    console.log("当前应用包名: " + packageName);
`)

// 执行 JavaScript 文件（从真实文件系统）
err = engine.ExecuteFile("/sdcard/script.js")

// 执行 JavaScript 文件（从虚拟文件系统）
err = engine.ExecuteFile("scripts/script.js")
```

## 模块列表

| 模块 | 说明 | 详细文档 |
|------|------|----------|
| `app` | 应用管理 | [README](model/app/README.md) |
| `device` | 设备信息 | [README](model/device/README.md) |
| `motion` | 触摸操作 | [README](model/motion/README.md) |
| `files` | 文件操作 | [README](model/files/README.md) |
| `images` | 图像处理 | [README](model/images/README.md) |
| `storages` | 数据存储 | [README](model/storages/README.md) |
| `system` | 系统功能 | [README](model/system/README.md) |
| `http` | 网络请求 | [README](model/http/README.md) |
| `media` | 媒体控制 | [README](model/media/README.md) |
| `opencv` | 计算机视觉 | [README](model/opencv/README.md) |
| `ppocr` | OCR 文字识别 | [README](model/ppocr/README.md) |
| `console` | 控制台窗口 | [README](model/console/README.md) |
| `coroutine` | 协程管理 | [README](model/coroutine/README.md) |
| `dotocr` | 点字 OCR | [README](model/dotocr/README.md) |
| `hud` | HUD 悬浮显示 | [README](model/hud/README.md) |
| `ime` | 输入法控制 | [README](model/ime/README.md) |
| `plugin` | 插件加载 | [README](model/plugin/README.md) |
| `rhino` | JavaScript 执行引擎 | [README](model/rhino/README.md) |
| `uiacc` | 无障碍 UI 操作 | [README](model/uiacc/README.md) |
| `utils` | 工具方法 | [README](model/utils/README.md) |
| `vdisplay` | 虚拟显示 | [README](model/vdisplay/README.md) |
| `yolo` | YOLO 目标检测 | [README](model/yolo/README.md) |
| `imgui` | Dear ImGui GUI 库 | [README](model/imgui/README.md) |

## DemoCode

### 1. 加载网络的脚本并执行

```go
package main

import (
    "fmt"
    "io"
    "net/http"
    "github.com/ZingYao/autogo_scriptengine/js_engine"
    "github.com/ZingYao/autogo_scriptengine/js_engine/define/safe_models"
)

func main() {
    // 注册 JavaScript 引擎模块
    js_engine.RegisterModule(safe_models.SafeModules...)
    
    // 创建引擎
    engine := js_engine.NewEngine(&js_engine.DefaultConfig())
    defer engine.Close()
    
    // 从网络下载脚本
    resp, err := http.Get("https://example.com/script.js")
    if err != nil {
        fmt.Printf("下载脚本失败: %v\n", err)
        return
    }
    defer resp.Body.Close()
    
    // 读取脚本内容
    script, err := io.ReadAll(resp.Body)
    if err != nil {
        fmt.Printf("读取脚本失败: %v\n", err)
        return
    }
    
    // 执行脚本
    err = engine.ExecuteString(string(script))
    if err != nil {
        fmt.Printf("执行脚本失败: %v\n", err)
    } else {
        fmt.Println("脚本执行成功")
    }
}
```

### 2. 加载真实文件系统的脚本并执行

```go
package main

import (
    "fmt"
    "os"
    "github.com/ZingYao/autogo_scriptengine/js_engine"
    "github.com/ZingYao/autogo_scriptengine/js_engine/define/safe_models"
)

func main() {
    // 注册 JavaScript 引擎模块
    js_engine.RegisterModule(safe_models.SafeModules...)
    
    // 创建引擎，配置真实文件系统
    config := js_engine.DefaultConfig()
    config.FileSystem = os.DirFS("/sdcard/script")  // 设置真实文件系统目录
    engine := js_engine.NewEngine(&config)
    defer engine.Close()
    
    // 执行脚本文件
    err := engine.ExecuteFile("main.js")
    if err != nil {
        fmt.Printf("执行脚本失败: %v\n", err)
    } else {
        fmt.Println("脚本执行成功")
    }
}
```

### 3. 加载虚拟文件系统的脚本并运行

```go
package main

import (
    "embed"
    "fmt"
    "github.com/ZingYao/autogo_scriptengine/js_engine"
    "github.com/ZingYao/autogo_scriptengine/js_engine/define/safe_models"
)

//go:embed scripts/*
var scriptsFS embed.FS

func main() {
    // 注册 JavaScript 引擎模块
    js_engine.RegisterModule(safe_models.SafeModules...)
    
    // 创建引擎，配置虚拟文件系统
    config := js_engine.DefaultConfig()
    config.FileSystem = scriptsFS      // 设置虚拟文件系统
    engine := js_engine.NewEngine(&config)
    defer engine.Close()
    
    // 执行脚本文件
    err := engine.ExecuteFile("scripts/main.js")
    if err != nil {
        fmt.Printf("执行脚本失败: %v\n", err)
    } else {
        fmt.Println("脚本执行成功")
    }
}
```

### 4. Main 脚本加载 Model 脚本并运行

**main.js** (主脚本)
```javascript
// main.js - 主脚本，加载并调用 model 脚本

console.log("=== 主脚本开始执行 ===");

// 导入 model 脚本
const model = require("./model");

// 调用 model 中的函数
model.init();
model.process("测试数据");

console.log("=== 主脚本执行完成 ===");
```

**model.js** (模型脚本)
```javascript
// model.js - 模型脚本

function init() {
    console.log("模型初始化完成");
}

function process(data) {
    console.log("处理数据: " + data);
    // 在这里添加具体的处理逻辑
}

// 导出函数
module.exports = {
    init,
    process
};
```

**Go 代码**
```go
package main

import (
    "embed"
    "fmt"
    "github.com/ZingYao/autogo_scriptengine/js_engine"
    "github.com/ZingYao/autogo_scriptengine/js_engine/define/safe_models"
)

//go:embed scripts/*
var scriptsFS embed.FS

func main() {
    // 注册 JavaScript 引擎模块
    js_engine.RegisterModule(safe_models.SafeModules...)
    
    // 创建引擎，配置虚拟文件系统
    config := js_engine.DefaultConfig()
    config.FileSystem = scriptsFS
    engine := js_engine.NewEngine(&config)
    defer engine.Close()
    
    // 执行主脚本，主脚本会自动加载 model 脚本
    err := engine.ExecuteFile("scripts/main.js")
    if err != nil {
        fmt.Printf("执行脚本失败: %v\n", err)
    } else {
        fmt.Println("脚本执行成功")
    }
}
```

## API 参考

### 引擎方法

| 方法 | 说明 |
|------|------|
| `NewEngine(config *EngineConfig) *JSEngine` | 创建新的 JavaScript 引擎实例 |
| `GetVM() *goja.Runtime` | 获取 JavaScript 运行时对象 |
| `InjectModule(moduleName string)` | 注入指定模块 |
| `InjectModules(modules []string)` | 注入多个模块 |
| `GetAvailableModules() []string` | 获取所有可用模块列表 |
| `InjectAllMethods()` | 注入所有模块的方法 |
| `RegisterMethod(name, description string, goFunc interface{}, overridable bool)` | 注册自定义方法 |
| `ExecuteString(script string, dir ...string) error` | 执行 JavaScript 字符串（可指定目录） |
| `ExecuteFile(path string) error` | 执行 JavaScript 文件 |
| `Close()` | 关闭引擎 |
| `GetRegistry() *MethodRegistry` | 获取方法注册表 |

### 方法详细说明

#### NewEngine
创建新的 JavaScript 引擎实例。

```go
config := js_engine.DefaultConfig()
engine := js_engine.NewEngine(&config)
```

#### GetVM
获取 JavaScript 运行时对象，用于直接操作 JavaScript 虚拟机。

```go
vm := engine.GetVM()
```

#### InjectModule
注入指定模块的方法到引擎中。

```go
engine.InjectModule("app")
engine.InjectModule("device")
```

支持的模块：app, device, motion, files, images, storages, system, http, media, opencv, ppocr, console, coroutine, dotocr, hud, ime, plugin, rhino, uiacc, utils, vdisplay, yolo, imgui

#### InjectModules
注入多个模块的方法到引擎中。

```go
engine.InjectModules([]string{"app", "device", "motion"})
```

#### GetAvailableModules
获取所有可用模块列表。

```go
modules := engine.GetAvailableModules()
fmt.Println(modules)
// 输出: [app device motion files images storages system http media opencv ppocr console coroutine dotocr hud ime plugin rhino uiacc utils vdisplay yolo imgui]
```

#### InjectAllMethods
注入所有模块的方法到引擎中。

```go
engine.InjectAllMethods()
```

#### RegisterMethod
注册自定义方法到引擎中。

```go
engine.RegisterMethod("myFunction", "我的自定义函数", func(call goja.FunctionCall) goja.Value {
    // 实现自定义逻辑
    return nil
}, true)
```

参数：
- `name`: 方法名称
- `description`: 方法描述
- `goFunc`: Go 函数
- `overridable`: 是否允许被 JavaScript 函数重写

#### ExecuteString
执行 JavaScript 代码字符串。

```go
// 基本用法
err := engine.ExecuteString(`
    console.log("Hello, JavaScript!");
`)

// 指定目录（用于 require）
err = engine.ExecuteString(`
    const model = require("./model");
    model.init();
`, "scripts")
```

参数：
- `script`: 要执行的 JavaScript 代码
- `dir`: 可选参数，指定工作目录（用于 require）

#### ExecuteFile
执行 JavaScript 文件。

```go
// 从真实文件系统执行
err := engine.ExecuteFile("/sdcard/script.js")

// 从虚拟文件系统执行
err = engine.ExecuteFile("scripts/script.js")
```

参数：
- `path`: 文件路径

#### Close
关闭引擎，释放资源。

```go
defer engine.Close()
```

#### GetRegistry
获取方法注册表，用于管理已注册的方法。

```go
registry := engine.GetRegistry()
```

### 配置选项

| 选项 | 类型 | 说明 |
|------|------|------|
| `AutoInjectMethods` | bool | 是否自动注入所有方法，默认为 true |
| `WhiteList` | []string | 白名单：只加载这些模块 |
| `BlackList` | []string | 黑名单：跳过这些模块 |
| `FailFast` | bool | 是否在模块加载失败时立即失败，false = 跳过失败模块继续 |
| `FileSystem` | fs.FS | 虚拟文件系统 (embed.FS)，用于从嵌入文件中加载模块 |

## License

MIT License
