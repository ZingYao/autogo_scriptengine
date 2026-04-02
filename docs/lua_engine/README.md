# Lua 引擎使用文档

## 1. 引擎简介

Lua 引擎是 AutoGo 脚本引擎的一部分，用于执行 Lua 脚本。它提供了丰富的 API 接口，支持与设备交互、文件操作、网络请求等功能。

## 2. 使用方法

### 2.1 基本使用

创建一个 Lua 脚本文件，例如 `test.lua`，并编写以下代码：

```lua
-- 直接使用全局模块，无需 require
-- 输出日志
console.log("Hello, AutoGo!")

-- 获取设备信息
console.log("屏幕宽度: " .. device.width)
console.log("屏幕高度: " .. device.height)
```

### 2.2 执行方式

**重要提示**：AutoGo 项目的运行必须要使用 AutoGo 的 VSCode Extension 或者 GoLand 的 Extension 来执行，不能直接使用 `go run` 或者 `go build` 命令。

Lua 脚本可以通过以下两种方式执行：

1. **Embed 方式**：将脚本嵌入到 Go 程序中编译执行
2. **手动放置方式**：将脚本文件手动放到移动设备上，然后根据路径执行

## 3. 脚本打包到 Embed 中

### 3.1 打包步骤

1. 创建一个包含所有脚本文件的目录，例如 `scripts`
2. 在 Go 代码中使用 `embed` 指令嵌入脚本文件：

```go
package main

import (
    "embed"
    "log"

    "github.com/ZingYao/autogo_scriptengine/lua_engine"
    "github.com/ZingYao/autogo_scriptengine/lua_engine/define/autogo/all_models"
)

//go:embed scripts/*
var scriptsFS embed.FS

func main() {
    // 初始化 Lua 引擎，配置文件系统以支持 require
    config := lua_engine.DefaultConfig()
    config.FileSystem = scriptsFS
    engine := lua_engine.NewLuaEngine(&config)
    defer engine.Close()

    // 注册所有 autogo 风格模块
    engine.RegisterModule(all_models.AllModules...)

    // 执行主脚本
    err := engine.ExecuteFile("scripts/main.lua")
    if err != nil {
        log.Fatalf("Failed to execute main.lua: %v", err)
    }
}
```

### 3.2 执行 Embed 脚本

使用 AutoGo 的 VSCode Extension 或 GoLand Extension 来运行项目。

## 4. 手动放置脚本到设备

### 4.1 放置步骤

1. 将脚本文件手动复制到移动设备的指定目录
2. 在 Go 代码中指定脚本路径并执行：

```go
package main

import (
    "log"

    "github.com/ZingYao/autogo_scriptengine/lua_engine"
    "github.com/ZingYao/autogo_scriptengine/lua_engine/define/autogo/all_models"
)

func main() {
    // 初始化 Lua 引擎
    config := lua_engine.DefaultConfig()
    engine := lua_engine.NewLuaEngine(&config)
    defer engine.Close()

    // 注册所有 autogo 风格模块
    engine.RegisterModule(all_models.AllModules...)

    // 执行设备上的脚本文件
    err := engine.ExecuteFile("/sdcard/scripts/main.lua")
    if err != nil {
        log.Fatalf("Failed to execute main.lua: %v", err)
    }
}
```

## 5. 从网络中加载代码并执行

### 5.1 网络加载示例

以下示例演示如何从网络中加载 Lua 代码并执行：

```go
package main

import (
    "fmt"
    "io"
    "net/http"
    "log"

    "github.com/ZingYao/autogo_scriptengine/lua_engine"
    "github.com/ZingYao/autogo_scriptengine/lua_engine/define/autogo/all_models"
)

func main() {
    // 初始化 Lua 引擎
    config := lua_engine.DefaultConfig()
    engine := lua_engine.NewLuaEngine(&config)
    defer engine.Close()

    // 注册所有 autogo 风格模块
    engine.RegisterModule(all_models.AllModules...)

    // 从网络加载脚本代码
    url := "https://example.com/scripts/main.lua"
    resp, err := http.Get(url)
    if err != nil {
        log.Fatalf("Failed to download script: %v", err)
    }
    defer resp.Body.Close()

    // 读取脚本内容
    scriptContent, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Fatalf("Failed to read script: %v", err)
    }

    // 执行脚本
    err = engine.ExecuteString(string(scriptContent), "main.lua")
    if err != nil {
        log.Fatalf("Failed to execute script: %v", err)
    }

    fmt.Println("Script executed successfully!")
}
```

### 5.2 网络加载说明

1. 使用 `http.Get` 或其他 HTTP 客户端从网络下载脚本
2. 读取下载的内容为字符串
3. 使用 `ExecuteString` 方法执行脚本内容
4. 可以指定脚本文件名用于错误提示

## 6. Lua 中如何 require 其他代码文件

### 6.1 基本 require

在 Lua 中，可以使用 `require` 函数来加载其他 Lua 文件（仅限用户自定义的 Lua 文件）。例如：

```lua
-- 加载同目录下的 utils.lua 文件
local utils = require("utils")

-- 调用 utils 中的函数
utils.doSomething()
```

### 6.2 加载子目录中的文件

如果文件在子目录中，可以使用点号分隔路径：

```lua
-- 加载 lib/utils.lua 文件
local utils = require("lib.utils")
```

### 6.3 使用注入的模块

AutoGo 脚本引擎提供了两种风格包，所有模块都已通过 Go 代码注入到 Lua 全局环境中，**无需使用 require**：

1. **autogo 风格**：基于 AutoGo 原生 API
2. **lrappsoft 风格**：基于懒人脚本 API，兼容大部分懒人脚本的 Lua 方法

使用注入模块的示例：

```lua
-- autogo 风格：直接使用全局模块
console.log("Hello")
device.width
motion.click(100, 200)

-- lrappsoft 风格：直接使用全局模块
log("Hello")
getScreenWidth()
touchDown(1, 100, 200)
```

### 6.4 模块导出

在 Lua 中，使用 `return` 来导出模块：

```lua
-- utils.lua
local utils = {}

function utils.add(a, b)
    return a + b
end

function utils.subtract(a, b)
    return a - b
end

return utils
```

## 7. 执行模式

Lua 引擎支持两种执行模式：

### 7.1 同步执行
同步执行是默认模式，会阻塞等待脚本执行完成：

```go
// 同步执行文件
err := engine.ExecuteFile("script.lua")
if err != nil {
    log.Fatalf("执行失败: %v", err)
}

// 同步执行字符串
err = engine.ExecuteString("print('Hello')", "scripts")
if err != nil {
    log.Fatalf("执行失败: %v", err)
}
```

### 7.2 异步执行
异步执行会在后台运行脚本，不会阻塞当前线程：

```go
// 异步执行文件
err := engine.ExecuteFileWithMode("script.lua", lua_engine.ExecuteModeAsync)
if err != nil {
    log.Fatalf("执行失败: %v", err)
}

// 异步执行字符串
err = engine.ExecuteStringWithMode("print('Hello')", lua_engine.ExecuteModeAsync, "scripts")
if err != nil {
    log.Fatalf("执行失败: %v", err)
}
```

## 8. Require 路径配置

### 8.1 自动路径管理
- 脚本执行时，当前脚本所在目录会自动添加到 Lua 的搜索路径中
- 这意味着你可以直接 require 同目录下的文件，无需指定完整路径

### 8.2 自定义路径
你可以通过 API 添加自定义的 require 搜索路径：

```go
// 添加单个路径
engine.AddRequirePath("/path/to/modules")

// 设置多个路径
engine.SetRequirePaths([]string{
    "/path/to/modules",
    "/another/path"
})
```

## 9. 注意事项

1. Lua 脚本文件的扩展名必须是 `.lua`
2. require 时不需要添加 `.lua` 扩展名
3. 脚本执行时，当前目录会被添加到 Lua 的搜索路径中
4. **所有模块都已通过 Go 代码注入到 Lua 全局环境中，无需使用 require**
5. autogo 风格和 lrappsoft 风格的模块都直接在全局空间可用
6. 使用 `embed` 时，需要配置 `FileSystem` 为嵌入的文件系统
7. 手动放置脚本时，需要确保脚本文件路径正确
8. **AutoGo 项目的运行必须要使用 AutoGo 的 VSCode Extension 或 GoLand Extension 来执行**

## 10. 示例脚本

### 10.1 基本操作示例

```lua
-- 直接使用全局模块，无需 require
-- 输出日志
console.log("Hello, AutoGo!")

-- 获取设备信息
console.log("屏幕宽度: " .. device.width)
console.log("屏幕高度: " .. device.height)
```

### 10.2 多文件脚本示例

**utils.lua**：

```lua
-- 工具函数
local utils = {}

function utils.add(a, b)
    return a + b
end

function utils.subtract(a, b)
    return a - b
end

return utils
```

**main.lua**：

```lua
-- 加载工具模块（用户自定义的 Lua 文件）
local utils = require("utils")

-- 直接使用全局模块（无需 require）
console.log("Hello, AutoGo!")

-- 测试工具函数
local sum = utils.add(5, 3)
local difference = utils.subtract(10, 4)

-- 输出结果
console.log("5 + 3 = " .. sum)
console.log("10 - 4 = " .. difference)
```

### 10.3 完整的 Go 示例（Embed 方式 - autogo 风格）

```go
package main

import (
    "embed"
    "fmt"
    "log"

    "github.com/ZingYao/autogo_scriptengine/lua_engine"
    "github.com/ZingYao/autogo_scriptengine/lua_engine/define/autogo/all_models"
)

//go:embed scripts/*
var scriptsFS embed.FS

func main() {
    // 初始化 Lua 引擎，配置文件系统以支持 require
    config := lua_engine.DefaultConfig()
    config.FileSystem = scriptsFS
    engine := lua_engine.NewLuaEngine(&config)
    defer engine.Close()

    // 注册所有 autogo 风格模块
    engine.RegisterModule(all_models.AllModules...)

    // 执行工具脚本（以便主脚本可以引用）
    err := engine.ExecuteFile("scripts/utils.lua")
    if err != nil {
        log.Fatalf("Failed to execute utils.lua: %v", err)
    }

    // 执行主脚本
    err = engine.ExecuteFile("scripts/main.lua")
    if err != nil {
        log.Fatalf("Failed to execute main.lua: %v", err)
    }

    // 输出执行结果
    fmt.Println("Lua autogo style example completed!")
}
```

### 10.4 完整的 Go 示例（手动放置方式 - autogo 风格）

```go
package main

import (
    "fmt"
    "log"

    "github.com/ZingYao/autogo_scriptengine/lua_engine"
    "github.com/ZingYao/autogo_scriptengine/lua_engine/define/autogo/all_models"
)

func main() {
    // 初始化 Lua 引擎
    config := lua_engine.DefaultConfig()
    engine := lua_engine.NewLuaEngine(&config)
    defer engine.Close()

    // 注册所有 autogo 风格模块
    engine.RegisterModule(all_models.AllModules...)

    // 执行设备上的脚本文件
    err := engine.ExecuteFile("/sdcard/scripts/main.lua")
    if err != nil {
        log.Fatalf("Failed to execute main.lua: %v", err)
    }

    // 输出执行结果
    fmt.Println("Lua autogo style example completed!")
}
```

### 10.5 完整的 Go 示例（Embed 方式 - lrappsoft 风格）

```go
package main

import (
    "embed"
    "fmt"
    "log"

    "github.com/ZingYao/autogo_scriptengine/lua_engine"
    "github.com/ZingYao/autogo_scriptengine/lua_engine/define/lrappsoft_models"
)

//go:embed scripts/*
var scriptsFS embed.FS

func main() {
    // 初始化 Lua 引擎，配置文件系统以支持 require
    config := lua_engine.DefaultConfig()
    config.FileSystem = scriptsFS
    engine := lua_engine.NewLuaEngine(&config)
    defer engine.Close()

    // 注册所有 lrappsoft 风格模块
    engine.RegisterModule(lrappsoft_models.LrappsoftModules...)

    // 执行工具脚本（以便主脚本可以引用）
    err := engine.ExecuteFile("scripts/utils.lua")
    if err != nil {
        log.Fatalf("Failed to execute utils.lua: %v", err)
    }

    // 执行主脚本
    err = engine.ExecuteFile("scripts/main.lua")
    if err != nil {
        log.Fatalf("Failed to execute main.lua: %v", err)
    }

    // 输出执行结果
    fmt.Println("Lua lrappsoft style example completed!")
}
```

### 10.6 完整的 Go 示例（手动放置方式 - lrappsoft 风格）

```go
package main

import (
    "fmt"
    "log"

    "github.com/ZingYao/autogo_scriptengine/lua_engine"
    "github.com/ZingYao/autogo_scriptengine/lua_engine/define/lrappsoft_models"
)

func main() {
    // 初始化 Lua 引擎
    config := lua_engine.DefaultConfig()
    engine := lua_engine.NewLuaEngine(&config)
    defer engine.Close()

    // 注册所有 lrappsoft 风格模块
    engine.RegisterModule(lrappsoft_models.LrappsoftModules...)

    // 执行设备上的脚本文件
    err := engine.ExecuteFile("/sdcard/scripts/main.lua")
    if err != nil {
        log.Fatalf("Failed to execute main.lua: %v", err)
    }

    // 输出执行结果
    fmt.Println("Lua lrappsoft style example completed!")
}
```

## 11. 示例代码 GitHub 地址

Lua 引擎的示例代码可以在以下 GitHub 地址找到：

- **Lua autogo 风格示例**：[examples/lua_engine/autogo](https://github.com/ZingYao/autogo_scriptengine/tree/main/examples/lua_engine/autogo)
- **Lua lrappsoft 风格示例**：[examples/lua_engine/lrappsoft](https://github.com/ZingYao/autogo_scriptengine/tree/main/examples/lua_engine/lrappsoft)

这些示例包含了完整的 Lua 脚本和 Go 代码，展示了如何使用 embed 和 require 功能。
