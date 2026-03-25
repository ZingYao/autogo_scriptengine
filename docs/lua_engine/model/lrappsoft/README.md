# Lua 引擎 lrappsoft 风格包文档

## 1. 风格包简介

lrappsoft 风格包是 Lua 引擎的一种风格实现，基于懒人脚本 API 开发。它提供了一套兼容懒人脚本的 API 接口，方便开发者快速迁移和使用懒人的 Lua 脚本。

## 2. 实现基础

lrappsoft 风格包基于懒人脚本 API 实现，主要特点包括：

- 兼容懒人脚本的 API 接口
- 实现了大部分懒人的 Lua 方法
- 保持与懒人脚本的使用习惯一致
- 支持懒人脚本的核心功能

## 3. 目录结构

lrappsoft 风格包的目录结构如下：

```
lrappsoft/
├── device/      # 设备操作
├── time/        # 时间操作
└── console/     # 控制台输出
```

## 4. 使用方法

### 4.1 导入模块

在 Lua 脚本中，可以使用 `require` 函数导入 lrappsoft 风格包的模块：

```lua
-- 导入 device 模块
local device = require("lrappsoft.device")

-- 导入 time 模块
local time = require("lrappsoft.time")

-- 导入 console 模块
local console = require("lrappsoft.console")
```

### 4.2 基本使用示例

```lua
-- 导入模块
local device = require("lrappsoft.device")
local time = require("lrappsoft.time")
local console = require("lrappsoft.console")

-- 点击屏幕
device.click(500, 500)

-- 输出日志
console.log("点击成功")

-- 等待 2 秒
time.sleep(2000)
```

## 5. 如何迁移懒人的 Lua 脚本

由于 lrappsoft 包实现了大部分懒人的 Lua 方法，因此可以直接调用这些方法，无需修改太多代码。以下是迁移步骤：

### 5.1 迁移步骤

1. **替换导入路径**：将原来的导入路径替换为 lrappsoft 风格包的路径
2. **保持方法调用不变**：由于 lrappsoft 风格包实现了与懒人脚本相同的方法名和参数，因此可以保持方法调用不变
3. **测试脚本**：运行迁移后的脚本，确保功能正常

### 5.2 迁移示例

**原懒人脚本**：

```lua
-- 导入模块
local device = require("device")
local time = require("time")
local console = require("console")

-- 点击屏幕
device.click(500, 500)

-- 输出日志
console.log("点击成功")

-- 等待 2 秒
time.sleep(2000)
```

**迁移后脚本**：

```lua
-- 导入模块
local device = require("lrappsoft.device")
local time = require("lrappsoft.time")
local console = require("lrappsoft.console")

-- 点击屏幕
device.click(500, 500)

-- 输出日志
console.log("点击成功")

-- 等待 2 秒
time.sleep(2000)
```

### 5.3 注意事项

1. 部分懒人脚本的高级功能可能需要额外调整
2. 对于不支持的方法，需要使用 AutoGo 原生 API 进行替代
3. 迁移过程中建议逐步测试，确保每个功能正常

## 6. 模块说明

### 6.1 device 模块

提供设备操作功能，如点击、滑动、输入等，兼容懒人脚本的 device 模块。

### 6.2 time 模块

提供时间操作功能，如延迟、定时等，兼容懒人脚本的 time 模块。

### 6.3 console 模块

提供控制台输出功能，如日志输出、错误提示等，兼容懒人脚本的 console 模块。

## 7. 注意事项

1. lrappsoft 风格包的模块路径是 `lrappsoft.模块名`
2. 所有函数的参数和返回值与懒人脚本保持一致
3. 使用前请确保已导入所需的模块
4. 详细的模块 API 文档请参考各模块的 README.md 文件