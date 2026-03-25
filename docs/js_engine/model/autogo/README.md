# JavaScript 引擎 autogo 风格包文档

## 1. 风格包简介

autogo 风格包是 JavaScript 引擎的一种风格实现，基于 AutoGo 原生 API 开发。它提供了一套简洁、高效的 API 接口，方便开发者快速编写脚本。

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
├── coroutine/   # 协程支持
├── device/      # 设备操作
├── dotocr/      # 点字 OCR 识别
├── files/       # 文件操作
├── http/        # 网络请求
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
└── yolo/        # YOLO 目标检测
```

## 4. 使用方法

### 4.1 导入模块

在 JavaScript 脚本中，可以使用 `require` 函数导入 autogo 风格包的模块：

```javascript
// 导入 app 模块
const app = require("autogo.app");

// 导入 device 模块
const device = require("autogo.device");

// 导入 console 模块
const console = require("autogo.console");

// 导入其他模块
const files = require("autogo.files");
const http = require("autogo.http");
// 更多模块...
```

### 4.2 基本使用示例

```javascript
// 导入模块
const app = require("autogo.app");
const device = require("autogo.device");
const console = require("autogo.console");

// 启动应用
app.start("com.example.app");

// 等待应用启动
os.sleep(3000);

// 点击屏幕
device.click(500, 500);

// 输出日志
console.log("点击成功");

// 等待 2 秒
os.sleep(2000);

// 退出应用
app.stop("com.example.app");
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

### 5.6 其他模块

- **coroutine**：协程支持
- **dotocr**：点字 OCR 识别
- **hud**：HUD 悬浮显示
- **images**：图像处理
- **ime**：输入法控制
- **imgui**：ImGui GUI 库
- **media**：媒体控制
- **motion**：触摸操作
- **opencv**：计算机视觉
- **plugin**：插件加载
- **ppocr**：OCR 文字识别
- **rhino**：JavaScript 执行引擎
- **storages**：数据存储
- **system**：系统功能
- **uiacc**：无障碍 UI 操作
- **utils**：工具方法
- **vdisplay**：虚拟显示
- **yolo**：YOLO 目标检测

## 6. 注意事项

1. autogo 风格包的模块路径是 `autogo.模块名`
2. 所有函数的参数和返回值与 AutoGo 原生 API 保持一致
3. 使用前请确保已导入所需的模块
4. 详细的模块 API 文档请参考各模块的 README.md 文件