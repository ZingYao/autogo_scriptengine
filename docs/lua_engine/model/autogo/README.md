# Lua 引擎 autogo 风格包文档

## 1. 风格包简介

autogo 风格包是 Lua 引擎的一种风格实现，基于 AutoGo 原生 API 开发。它提供了一套简洁、高效的 API 接口，方便开发者快速编写脚本。

## 2. 实现基础

autogo 风格包基于 AutoGo 原生 API 实现，主要特点包括：

- 直接映射 AutoGo 原生 API 功能
- 提供简洁的函数命名和参数结构
- 保持与 AutoGo 核心功能的一致性
- 支持所有 AutoGo 原生 API 的功能

## 3. 目录结构

autogo 风格包的目录结构如下：

```
autogo/
├── app/         # 应用相关操作
├── console/     # 控制台输出
├── device/      # 设备操作
├── files/       # 文件操作
└── http/        # 网络请求
```

## 4. 使用方法

### 4.1 导入模块

在 Lua 脚本中，可以使用 `require` 函数导入 autogo 风格包的模块：

```lua
-- 导入 app 模块
local app = require("autogo.app")

-- 导入 device 模块
local device = require("autogo.device")

-- 导入 console 模块
local console = require("autogo.console")

-- 导入 files 模块
local files = require("autogo.files")

-- 导入 http 模块
local http = require("autogo.http")
```

### 4.2 基本使用示例

```lua
-- 导入模块
local app = require("autogo.app")
local device = require("autogo.device")
local console = require("autogo.console")

-- 启动应用
app.start("com.example.app")

-- 等待应用启动
os.sleep(3000)

-- 点击屏幕
device.click(500, 500)

-- 输出日志
console.log("点击成功")

-- 等待 2 秒
os.sleep(2000)

-- 退出应用
app.stop("com.example.app")
```

## 5. 模块说明

### 5.1 app 模块

提供应用相关的操作，如启动、停止、切换应用等。

### 5.2 console 模块

提供控制台输出功能，如日志输出、错误提示等。

### 5.3 device 模块

提供设备操作功能，如点击、滑动、输入等。

### 5.4 files 模块

提供文件操作功能，如读写文件、创建目录等。

### 5.5 http 模块

提供网络请求功能，如 GET、POST 请求等。

## 6. 注意事项

1. autogo 风格包的模块路径是 `autogo.模块名`
2. 所有函数的参数和返回值与 AutoGo 原生 API 保持一致
3. 使用前请确保已导入所需的模块
4. 详细的模块 API 文档请参考各模块的 README.md 文件