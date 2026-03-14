# AutoGo Lua Engine

AutoGo Lua Engine 是一个高性能的 Lua 脚本引擎，为 AutoGo 框架提供了完整的 Lua 脚本支持。

## 特性

- **完整的 API 注入**: 将 AutoGo 的所有功能模块注入到 Lua 引擎中
- **模块化设计**: 支持按需加载模块
- **线程安全**: 所有操作都是线程安全的
- **文件系统支持**: 支持真实文件系统和虚拟文件系统 (embed.FS)
- **模块导入**: 支持 Lua require 功能

## 快速开始

### 初始化引擎

```go
import "github.com/ZingYao/autogo_scriptengine/lua_engine"

func main() {
    // 使用默认配置创建引擎（自动注入所有方法）
    engine := lua_engine.NewLuaEngine(&lua_engine.DefaultConfig())
    defer engine.Close()
    
    // 使用引擎...
}
```

### 引擎配置选项

```go
import (
    "embed"
    "os"
    "github.com/ZingYao/autogo_scriptengine/lua_engine"
)

//go:embed scripts/*
var scriptsFS embed.FS

func main() {
    // 使用自定义配置
    config := lua_engine.DefaultConfig()
    config.AutoInjectMethods = false  // 禁用自动注入
    config.FileSystem = scriptsFS      // 设置虚拟文件系统
    config.SearchPaths = []string{"scripts"}  // 设置模块搜索路径
    
    engine := lua_engine.NewLuaEngine(&config)
    defer engine.Close()
    
    // 按需注入模块
    engine.InjectModule("app")
    engine.InjectModule("device")
}
```

### 执行 Lua 代码

```go
// 执行 Lua 字符串
err := engine.ExecuteString(`
    console.log("Hello, Lua!")
    local packageName = app_currentPackage()
    console.log("当前应用包名: " .. packageName)
`)

// 执行 Lua 文件（从真实文件系统）
err = engine.ExecuteFile("/sdcard/script.lua")

// 执行 Lua 文件（从虚拟文件系统）
err = engine.ExecuteFile("scripts/script.lua")
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
| `json` | JSON 处理 | [README](model/json/README.md) |

## DemoCode

### 1. 加载网络的脚本并执行

```go
package main

import (
    "fmt"
    "io"
    "net/http"
    "github.com/ZingYao/autogo_scriptengine/lua_engine"
    lua_safe_models "github.com/ZingYao/autogo_scriptengine/lua_engine/define/safe_models"
)

func main() {
    // 注册 Lua 引擎模块
    lua_engine.RegisterModule(lua_safe_models.SafeModules...)
    
    // 创建引擎
    engine := lua_engine.NewLuaEngine(&lua_engine.DefaultConfig())
    defer engine.Close()
    
    // 从网络下载脚本
    resp, err := http.Get("https://example.com/script.lua")
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
    "github.com/ZingYao/autogo_scriptengine/lua_engine"
    lua_safe_models "github.com/ZingYao/autogo_scriptengine/lua_engine/define/safe_models"
)

func main() {
    // 注册 Lua 引擎模块
    lua_engine.RegisterModule(lua_safe_models.SafeModules...)
    
    // 创建引擎，配置真实文件系统
    config := lua_engine.DefaultConfig()
    config.FileSystem = os.DirFS("/sdcard/script")  // 设置真实文件系统目录
    engine := lua_engine.NewLuaEngine(&config)
    defer engine.Close()
    
    // 执行脚本文件
    err := engine.ExecuteFile("main.lua")
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
    "github.com/ZingYao/autogo_scriptengine/lua_engine"
    lua_safe_models "github.com/ZingYao/autogo_scriptengine/lua_engine/define/safe_models"
)

//go:embed scripts/*
var scriptsFS embed.FS

func main() {
    // 注册 Lua 引擎模块
    lua_engine.RegisterModule(lua_safe_models.SafeModules...)
    
    // 创建引擎，配置虚拟文件系统
    config := lua_engine.DefaultConfig()
    config.FileSystem = scriptsFS      // 设置虚拟文件系统
    config.SearchPaths = []string{"scripts"}  // 设置模块搜索路径
    engine := lua_engine.NewLuaEngine(&config)
    defer engine.Close()
    
    // 执行脚本文件
    err := engine.ExecuteFile("scripts/main.lua")
    if err != nil {
        fmt.Printf("执行脚本失败: %v\n", err)
    } else {
        fmt.Println("脚本执行成功")
    }
}
```

### 4. Main 脚本加载 Model 脚本并运行

**main.lua** (主脚本)
```lua
-- main.lua - 主脚本，加载并调用 model 脚本

console.log("=== 主脚本开始执行 ===")

-- 导入 model 脚本
local model = require("model")

-- 调用 model 中的函数
model.init()
model.process("测试数据")

console.log("=== 主脚本执行完成 ===")
```

**model.lua** (模型脚本)
```lua
-- model.lua - 模型脚本

local M = {}

-- 初始化函数
function M.init()
    console.log("模型初始化完成")
end

-- 处理函数
function M.process(data)
    console.log("处理数据: " .. data)
    -- 在这里添加具体的处理逻辑
end

return M
```

**Go 代码**
```go
package main

import (
    "embed"
    "fmt"
    "github.com/ZingYao/autogo_scriptengine/lua_engine"
    lua_safe_models "github.com/ZingYao/autogo_scriptengine/lua_engine/define/safe_models"
)

//go:embed scripts/*
var scriptsFS embed.FS

func main() {
    // 注册 Lua 引擎模块
    lua_engine.RegisterModule(lua_safe_models.SafeModules...)
    
    // 创建引擎，配置虚拟文件系统
    config := lua_engine.DefaultConfig()
    config.FileSystem = scriptsFS
    config.SearchPaths = []string{"scripts"}
    engine := lua_engine.NewLuaEngine(&config)
    defer engine.Close()
    
    // 执行主脚本，主脚本会自动加载 model 脚本
    err := engine.ExecuteFile("scripts/main.lua")
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
| `NewLuaEngine(config *EngineConfig) *LuaEngine` | 创建新的 Lua 引擎实例 |
| `GetState() *lua.LState` | 获取 Lua 状态对象 |
| `InjectModule(moduleName string)` | 注入指定模块 |
| `InjectModules(modules []string)` | 注入多个模块 |
| `GetAvailableModules() []string` | 获取所有可用模块列表 |
| `InjectAllMethods()` | 注入所有模块的方法 |
| `RegisterMethod(name, description string, goFunc interface{}, overridable bool)` | 注册自定义方法 |
| `ExecuteString(script string, searchPaths ...string) error` | 执行 Lua 字符串（可指定搜索路径） |
| `ExecuteFile(path string) error` | 执行 Lua 文件 |
| `Close()` | 关闭引擎 |
| `GetRegistry() *MethodRegistry` | 获取方法注册表 |

### 方法详细说明

#### NewLuaEngine
创建新的 Lua 引擎实例。

```go
config := lua_engine.DefaultConfig()
engine := lua_engine.NewLuaEngine(&config)
```

#### GetState
获取 Lua 状态对象，用于直接操作 Lua 虚拟机。

```go
state := engine.GetState()
```

#### InjectModule
注入指定模块的方法到引擎中。

```go
engine.InjectModule("app")
engine.InjectModule("device")
```

支持的模块：app, device, motion, files, images, storages, system, http, media, opencv, ppocr, console, coroutine, dotocr, hud, ime, plugin, rhino, uiacc, utils, vdisplay, yolo, imgui, json

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
// 输出: [app device motion files images storages system http media opencv ppocr console coroutine dotocr hud ime plugin rhino uiacc utils vdisplay yolo imgui json]
```

#### InjectAllMethods
注入所有模块的方法到引擎中。

```go
engine.InjectAllMethods()
```

#### RegisterMethod
注册自定义方法到引擎中。

```go
engine.RegisterMethod("myFunction", "我的自定义函数", func(L *lua.LState) int {
    // 实现自定义逻辑
    return 1
}, true)
```

参数：
- `name`: 方法名称
- `description`: 方法描述
- `goFunc`: Go 函数
- `overridable`: 是否允许被 Lua 函数重写

#### ExecuteString
执行 Lua 代码字符串。

```go
// 基本用法
err := engine.ExecuteString(`
    console.log("Hello, Lua!")
`)

// 指定搜索路径（用于 require）
err = engine.ExecuteString(`
    local model = require("model")
    model.init()
`, "scripts")
```

参数：
- `script`: 要执行的 Lua 代码
- `searchPaths`: 可选参数，添加模块搜索路径（用于 require）

#### ExecuteFile
执行 Lua 文件。

```go
// 从真实文件系统执行
err := engine.ExecuteFile("/sdcard/script.lua")

// 从虚拟文件系统执行
err = engine.ExecuteFile("scripts/script.lua")
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
| `SearchPaths` | []string | 模块搜索路径，用于 require 查找模块 |
| `FileSystem` | fs.FS | 虚拟文件系统 (embed.FS)，用于从嵌入文件中加载模块 |

## 兼容性说明

### Android 版本兼容性

某些依赖包在特定的 Android 版本下可能会出现内存引用错误（Memory Reference Error）。如果遇到此类问题，可以尝试以下解决方案：

1. **修改引入的包**：根据您的 Android 版本，可以修改项目中引入的相关包
2. **禁用问题模块**：在引擎配置中使用黑名单禁用导致问题的特定模块
3. **使用替代方案**：某些功能可能有多种实现方式，可以尝试使用替代方案

示例：禁用可能导致问题的模块
```go
config := lua_engine.DefaultConfig()
config.BlackList = []string{"opencv", "ppocr"}  // 禁用可能导致内存问题的模块
engine := lua_engine.NewLuaEngine(&config)
```

### Windows 开发环境

在 Windows 环境下开发时，如果引入了超过 1 个以上的带 C 依赖的库，可能会导致编译命令过长，触发以下错误：

```
The command line is too long.
```

**重要说明**：这是受限于当前 AutoGo Extension 的实现，无法从根本上解决。

**解决方案**：

1. **避免过多使用带 C 的库**：尽量减少使用包含 C 代码的依赖包
2. **减少依赖库的引用**：遇到问题时，仅保留刚需依赖库，使用白名单手动指定需要加载的模块
3. **切换开发环境**：使用 macOS 或 Linux 系统进行编译

**示例：使用白名单手动指定依赖**

```go
import "github.com/ZingYao/autogo_scriptengine/lua_engine"

// 只加载必需的模块，避免引入过多的 C 依赖
config := lua_engine.DefaultConfig()
config.WhiteList = []string{
    "app",      // 应用管理
    "device",   // 设备信息
    "motion",   // 触摸操作
    "files",    // 文件操作
    "console",  // 控制台
}
engine := lua_engine.NewLuaEngine(&config)
defer engine.Close()

// 执行脚本
err := engine.ExecuteString(`
    console.log("Hello, Lua!")
    local packageName = app.currentPackage()
    console.log("当前应用包名: " .. packageName)
`)
```

**建议的开发流程**：

1. 在 Windows 环境下开发时，只启用核心模块（如 app、device、motion、files、console）
2. 需要使用其他模块（如 opencv、ppocr 等）时，临时切换到 macOS/Linux 环境编译
3. 或者使用 WSL (Windows Subsystem for Linux) 环境进行开发

**可用模块列表**：

| 模块 | 是否包含 C 依赖 | 说明 |
|------|----------------|------|
| `app` | 否 | 应用管理 |
| `device` | 否 | 设备信息 |
| `motion` | 否 | 触摸操作 |
| `files` | 否 | 文件操作 |
| `images` | 否 | 图像处理 |
| `storages` | 否 | 数据存储 |
| `system` | 否 | 系统功能 |
| `http` | 否 | 网络请求 |
| `media` | 否 | 媒体控制 |
| `console` | 否 | 控制台窗口 |
| `coroutine` | 否 | 协程支持 |
| `dotocr` | 否 | 点字 OCR |
| `hud` | 否 | HUD 悬浮显示 |
| `ime` | 否 | 输入法控制 |
| `plugin` | 否 | 插件加载 |
| `rhino` | 否 | JavaScript 执行引擎 |
| `uiacc` | 否 | 无障碍 UI 操作 |
| `utils` | 否 | 工具方法 |
| `vdisplay` | 否 | 虚拟显示 |
| `yolo` | 是 | YOLO 目标检测 |
| `imgui` | 否 | Dear ImGui GUI 库 |
| `opencv` | 是 | 计算机视觉 |
| `ppocr` | 是 | OCR 文字识别 |
| `json` | 否 | JSON 处理 |

### 常见问题处理

如果遇到兼容性问题，建议：

1. 检查您的 Android 设备版本和 SDK 版本
2. 查看项目的 GitHub Issues，了解已知的兼容性问题
3. 根据错误信息调整配置或代码
4. 如有必要，可以 Fork 项目并根据您的环境进行修改

## License

MIT License
