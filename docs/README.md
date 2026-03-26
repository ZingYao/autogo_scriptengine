# AutoGo ScriptEngine

[AutoGo](https://github.com/Dasongzi1366/AutoGo) 的脚本引擎扩展方案，为 AutoGo 提供 JavaScript 和 Lua 脚本语言支持，让开发者可以用熟悉的脚本语言编写自动化任务。

## 为什么选择 ScriptEngine

1. **降低准入门槛** - 使用脚本语言开发，无需深入理解 Go 语言和 Android 开发，降低学习成本
2. **代码保护** - 脚本代码易于混淆加密，有效保护业务逻辑
3. **热更新支持** - 脚本可动态加载，无需重新编译即可更新功能
4. **无痛迁移** - 可以无痛迁移其他平台的代码，复用现有的脚本代码库

## 功能特性

- **双引擎支持**：同时支持 JavaScript 和 Lua 脚本语言
- **丰富的 API**：提供应用管理、设备控制、图像识别、OCR 等多种功能
- **方法注册系统**：支持动态注册、重写和恢复方法
- **协程支持**：Lua 引擎支持协程操作
- **风格包支持**：提供 autogo 和 lrappsoft 两种风格包
- **懒人脚本兼容**：lrappsoft 风格包兼容大部分懒人脚本的 Lua 方法
- **文档生成**：可自动生成 API 文档

## 安装

```bash
go get github.com/ZingYao/autogo_scriptengine@v0.0.15.14.14.14.14.14.14.13.12
```

## 📚 详细文档

> **🔥 重要提示**：查看以下详细文档以获取完整的 API 参考和使用指南

**📖 项目文档地址**：https://zingyao.github.io/autogo_scriptengine/

### 🌐 HTML 在线文档

> **推荐**：查看美观的 HTML 在线文档，提供更好的阅读体验

- [📖 文档索引](./docs/index.html) - 所有文档的导航页面
- [🏠 项目主页](./docs/README.md) - 项目介绍和功能特性

**使用方法**：
```bash
# 生成/更新 HTML 文档
python3 scripts/convert_to_html.py
```

### 📖 Markdown 文档

> 如果您更喜欢阅读 Markdown 格式的文档，可以查看以下链接：

#### JavaScript 引擎文档

- [JavaScript 引擎完整文档](./js_engine/README.md) - 包含所有 API、配置选项和高级用法

#### 模块文档

- [app 模块](./js_engine/model/app/README.md) - 应用管理（启动、安装、卸载、强制停止等）
- [device 模块](./js_engine/model/device/README.md) - 设备信息（分辨率、SDK 版本、屏幕方向等）
- [motion 模块](./js_engine/model/motion/README.md) - 触摸操作（点击、滑动、手势等）
- [files 模块](./js_engine/model/files/README.md) - 文件操作（读写、复制、删除等）
- [images 模块](./js_engine/model/images/README.md) - 图像处理（截图、找色、找图等）
- [storages 模块](./js_engine/model/storages/README.md) - 数据存储（键值对存储）
- [system 模块](./js_engine/model/system/README.md) - 系统功能（剪贴板、通知等）
- [http 模块](./js_engine/model/http/README.md) - 网络请求（GET、POST 等）
- [websocket 模块](./js_engine/model/websocket/README.md) - WebSocket 客户端（连接、发送、接收消息等）
- [media 模块](./js_engine/model/media/README.md) - 媒体控制（音量、播放等）
- [opencv 模块](./js_engine/model/opencv/README.md) - 计算机视觉（图像处理、特征检测等）
- [ppocr 模块](./js_engine/model/ppocr/README.md) - OCR 文字识别
- [console 模块](./js_engine/model/console/README.md) - 控制台窗口（显示、隐藏、日志输出等）
- [dotocr 模块](./js_engine/model/dotocr/README.md) - 点字 OCR 识别（基于字库的 OCR）
- [hud 模块](./js_engine/model/hud/README.md) - HUD 悬浮显示（脚本状态显示等）
- [ime 模块](./js_engine/model/ime/README.md) - 输入法控制（剪切板、文本输入等）
- [plugin 模块](./js_engine/model/plugin/README.md) - 插件加载（加载外部 APK 调用 Java 方法）
- [rhino 模块](./js_engine/model/rhino/README.md) - JavaScript 执行引擎（Rhino）
- [uiacc 模块](./js_engine/model/uiacc/README.md) - 无障碍 UI 操作（控件查找、点击、输入等）
- [utils 模块](./js_engine/model/utils/README.md) - 工具方法（日志、Toast、类型转换等）
- [vdisplay 模块](./js_engine/model/vdisplay/README.md) - 虚拟显示（虚拟屏操作）
- [yolo 模块](./js_engine/model/yolo/README.md) - YOLO 目标检测（v5/v8 模型）
- [imgui 模块](./js_engine/model/imgui/README.md) - Dear ImGui GUI 库（窗口、按钮、输入框等控件）
- [coroutine 模块](./js_engine/model/coroutine/README.md) - 协程支持

#### 📖 Lua 引擎文档

- [Lua 引擎完整文档](./lua_engine/README.md) - 包含所有 API、配置选项和高级用法

#### 模块文档

- [app 模块](./lua_engine/model/app/README.md) - 应用管理（启动、安装、卸载、强制停止等）
- [device 模块](./lua_engine/model/device/README.md) - 设备信息（分辨率、SDK 版本、屏幕方向等）
- [motion 模块](./lua_engine/model/motion/README.md) - 触摸操作（点击、滑动、手势等）
- [files 模块](./lua_engine/model/files/README.md) - 文件操作（读写、复制、删除等）
- [images 模块](./lua_engine/model/images/README.md) - 图像处理（截图、找色、找图等）
- [storages 模块](./lua_engine/model/storages/README.md) - 数据存储（键值对存储）
- [system 模块](./lua_engine/model/system/README.md) - 系统功能（剪贴板、通知等）
- [http 模块](./lua_engine/model/http/README.md) - 网络请求（GET、POST 等）
- [websocket 模块](./lua_engine/model/websocket/README.md) - WebSocket 客户端（连接、发送、接收消息等）
- [media 模块](./lua_engine/model/media/README.md) - 媒体控制（音量、播放等）
- [opencv 模块](./lua_engine/model/opencv/README.md) - 计算机视觉（图像处理、特征检测等）
- [ppocr 模块](./lua_engine/model/ppocr/README.md) - OCR 文字识别
- [console 模块](./lua_engine/model/console/README.md) - 控制台窗口（显示、隐藏、日志输出等）
- [dotocr 模块](./lua_engine/model/dotocr/README.md) - 点字 OCR 识别（基于字库的 OCR）
- [hud 模块](./lua_engine/model/hud/README.md) - HUD 悬浮显示（脚本状态显示等）
- [ime 模块](./lua_engine/model/ime/README.md) - 输入法控制（剪切板、文本输入等）
- [plugin 模块](./lua_engine/model/plugin/README.md) - 插件加载（加载外部 APK 调用 Java 方法）
- [rhino 模块](./lua_engine/model/rhino/README.md) - JavaScript 执行引擎（Rhino）
- [uiacc 模块](./lua_engine/model/uiacc/README.md) - 无障碍 UI 操作（控件查找、点击、输入等）
- [utils 模块](./lua_engine/model/utils/README.md) - 工具方法（日志、Toast、类型转换等）
- [vdisplay 模块](./lua_engine/model/vdisplay/README.md) - 虚拟显示（虚拟屏操作）
- [yolo 模块](./lua_engine/model/yolo/README.md) - YOLO 目标检测（v5/v8 模型）
- [imgui 模块](./lua_engine/model/imgui/README.md) - Dear ImGui GUI 库（窗口、按钮、输入框等控件）
- [coroutine 模块](./lua_engine/model/coroutine/README.md) - 协程支持
- [json 模块](./lua_engine/model/json/README.md) - JSON 处理

## 环境要求

- Go 1.25.0 或更高版本
- AutoGo 框架（已在项目中集成）
- Android 设备（用于实际运行自动化脚本）

## 兼容性说明

### Android 版本兼容性

某些依赖包在特定的 Android 版本下可能会出现内存引用错误（Memory Reference Error）。如果遇到此类问题，可以尝试以下解决方案：

1. **修改引入的包**：根据您的 Android 版本，可以修改项目中引入的相关包
2. **禁用问题模块**：在引擎配置中禁用导致问题的特定模块
3. **使用替代方案**：某些功能可能有多种实现方式，可以尝试使用替代方案

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
4. **或者使用 WSL (Windows Subsystem for Linux) 环境进行开发

**建议的开发流程**：

1. 在 Windows 环境下开发时，只启用核心模块（如 app、device、motion）
2. 需要使用其他模块时，临时切换到 macOS/Linux 环境编译
3. 或者使用 WSL (Windows Subsystem for Linux) 环境进行开发

详细的白名单使用示例和模块依赖列表，请查看：
- [JavaScript 引擎文档](./js_engine/README.md) - 包含完整的白名单使用示例
- [Lua 引擎文档](./lua_engine/README.md) - 包含完整的白名单使用示例

### 常见问题处理

如果遇到兼容性问题，建议：

1. 检查您的 Android 设备版本和 SDK 版本
2. 查看项目的 GitHub Issues，了解已知的兼容性问题
3. 根据错误信息调整配置或代码
4. 如有必要，可以 Fork 项目并根据您的环境进行修改

## 依赖

- [AutoGo](https://github.com/Dasongzi1366/AutoGo) - Android 自动化框架（核心依赖）
- [goja](https://github.com/dop251/goja) - JavaScript 解释器
- [gopher-lua](https://github.com/yuin/gopher-lua) - Lua 解释器

## 风格包说明

### autogo 风格包

基于 AutoGo 原生 API 开发的风格包，提供简洁、高效的 API 接口，方便开发者快速编写脚本。

### lrappsoft 风格包

基于懒人脚本 API 开发的风格包，兼容大部分懒人脚本的 Lua 方法，方便开发者快速迁移懒人的 Lua 脚本。

**主要特点**：
- 兼容懒人脚本的 API 接口
- 实现了大部分懒人的 Lua 方法
- 保持与懒人脚本的使用习惯一致
- 支持懒人脚本的核心功能

**迁移指南**：
1. 替换导入路径：将原来的导入路径替换为 lrappsoft 风格包的路径
2. 保持方法调用不变：由于 lrappsoft 风格包实现了与懒人脚本相同的方法名和参数，因此可以保持方法调用不变
3. 测试脚本：运行迁移后的脚本，确保功能正常

## 依赖架构的重大改变

在本次开发了 lrappsoft 风格包后，依赖架构发生了重大改变：

1. **风格包分离**：将 API 实现分为 autogo 和 lrappsoft 两种风格包，提供不同的编程风格选择
2. **模块化设计**：每个风格包内部采用模块化设计，便于维护和扩展
3. **兼容性增强**：通过 lrappsoft 风格包，增强了与懒人脚本的兼容性
4. **依赖管理**：优化了依赖管理，减少了不必要的依赖

## 与 AutoGo 的关系

本项目是 AutoGo 的扩展方案，通过封装 AutoGo 提供的原生 API，为开发者提供更灵活的脚本编写方式：

- **AutoGo** - 提供 Android 自动化的核心能力（无障碍服务、图像识别、触摸模拟等）
- **ScriptEngine** - 为 AutoGo 添加脚本语言支持，让开发者可以用 JavaScript 或 Lua 编写自动化脚本
- **风格包** - 提供不同风格的 API 接口，满足不同开发者的需求

## 问题反馈

如果您在使用过程中遇到任何问题、发现 bug 或者有功能建议，欢迎通过以下方式反馈：

### 📝 提交 Issue

我们非常欢迎您在 GitHub 上提交 Issue 来反馈问题。提交 Issue 时，请尽可能提供以下信息，以便我们更快地定位和解决问题：

**必填信息**：
- 问题描述：清晰描述您遇到的问题或建议
- 复现步骤：详细的操作步骤，让我们能够重现问题
- 预期行为：您期望看到的结果
- 实际行为：实际发生的情况

**环境信息**：
- AutoGo ScriptEngine 版本
- Go 语言版本
- Android 设备型号和系统版本
- 使用的脚本语言（JavaScript 或 Lua）
- 操作系统（Windows/macOS/Linux）

**可选信息**：
- 错误日志或堆栈信息
- 最小可复现代码示例
- 相关截图或录屏

### 🐛 Bug 反馈

发现 bug 时，请按照以下模板提交 Issue：

```markdown
**Bug 描述**
简要描述 bug 的表现

**复现步骤**
1. 执行操作 A
2. 执行操作 B
3. 观察到错误 C

**预期行为**
描述您期望的正常行为

**实际行为**
描述实际发生的错误行为

**环境信息**
- AutoGo ScriptEngine 版本：
- Go 版本：
- Android 版本：
- 脚本语言：
```

### 💡 功能建议

如果您有功能改进建议，请按照以下模板提交 Issue：

```markdown
**功能描述**
描述您希望添加或改进的功能

**使用场景**
描述这个功能的使用场景和价值

**建议方案**
如果您有具体的实现想法，请描述您的建议方案
```

### 📌 Issue 标签

提交 Issue 时，请使用合适的标签帮助我们分类：
- `bug` - 报告 bug
- `enhancement` - 功能增强建议
- `documentation` - 文档相关
- `question` - 使用问题
- `compatibility` - 兼容性问题

**GitHub Issues 地址**：https://github.com/ZingYao/autogo_scriptengine/issues

在提交 Issue 前，建议先搜索现有的 Issues，确认是否已有类似问题或建议。

## 许可证

MIT License
