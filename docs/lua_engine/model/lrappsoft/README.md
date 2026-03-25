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

lrappsoft 风格包的模块通过 Go 注入到 Lua 引擎中，无需在 Lua 脚本中使用 require 导入。

### 4.1 直接调用模块

由于 lrappsoft 风格包的模块已经通过 Go 注入到 Lua 引擎中，您可以直接在 Lua 脚本中调用这些模块，无需使用 require 导入。

```lua
-- 直接调用模块
device.click(500, 500)
time.sleep(2000)
console.log("点击成功")
```

### 4.2 模块说明

lrappsoft 风格包的模块已经通过 Go 注入到 Lua 全局环境中，**无需使用 require**，您可以直接调用模块方法。

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
-- 直接调用模块（无需 require）
device.click(500, 500)
time.sleep(2000)
console.log("点击成功")
```

## 6. 模块说明

### 6.1 device 模块

提供设备操作功能，如点击、滑动、输入等，兼容懒人脚本的 device 模块。

### 6.2 time 模块

提供时间操作功能，如延迟、定时等，兼容懒人脚本的 time 模块。

### 6.3 console 模块

提供控制台输出功能，如日志输出、错误提示等，兼容懒人脚本的 console 模块。

## 7. 注意事项

1. **所有 lrappsoft 模块都已通过 Go 代码注入到 Lua 全局环境中，无需使用 require**
2. 所有函数的参数和返回值与懒人脚本保持一致
3. 使用前请确保已在 Go 代码中注册了所需的模块
4. 详细的模块 API 文档请参考各模块的 README.md 文件
