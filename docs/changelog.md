# 更新日志

## Unreleased

### Debugger 文档

- 接入 AutoGo goja fork 的 JavaScript DAP 调试能力，补充 VSCode 与 JetBrains 的 JS 调试说明。
- 更新 AutoGo Debugger v1.0.0 使用说明，明确 VSCode 和 JetBrains 插件已内置 debugger 工具。
- 明确当前支持 Lua/GLua/JavaScript DAP 断点调试。
- 补充 VSCode 与 JetBrains 的脱敏截图指引，避免暴露本地路径、设备序列号等个人信息。
- 移除旧版终端工具使用路径描述，避免和当前 IDE Lua Debugger 定位混淆。

### iOS autogo 风格包

- 新增 Lua/JavaScript iOS autogo define 与 `autogo_ios` 模块隔离目录。
- iOS Lua/JavaScript API 按 AutoGo Go 导出名映射为模块对象小驼峰方法，不保留历史别名和全局入口。
- 补充 iOS examples，覆盖模块对象入口、struct/map/slice 入参、返回值解析、对象生命周期和 `opencv/imgui` 构造调用。
- 补充 iOS autogo 风格包文档，明确 `uiacc/apkctl` 暂不注入以及不得使用 Android-only `displayId` 参数。

## v0.0.15 (2026-03-25)

### 文档优化与 lrappsoft 风格兼容

- **重大文档更新**：全面优化和修正所有文档中的模块使用说明
- **修复模块导入问题**：明确说明所有模块通过 Go 代码注入到全局环境，无需使用 require
- **lrappsoft 风格兼容**：完整实现 lrappsoft 风格的 Lua 代码支持
- **更新示例代码**：修正所有文档和示例代码中的错误用法
- **澄清 require 用法**：明确区分用户自定义文件（需要 require）和注入模块（无需 require）
- **更新引擎文档**：
  - Lua 引擎文档：修正所有示例代码和说明
  - JavaScript 引擎文档：修正所有示例代码和说明
  - autogo 风格包文档：更新使用方法
  - lrappsoft 风格包文档：更新使用方法
- **修复示例代码**：更新 examples 目录中的示例脚本，确保与文档一致

## v0.0.9 (2026-03-16)

### 文档更新与脚本引擎支持

- 支持 JavaScript 和 Lua 脚本引擎
- 支持 require 功能，实现脚本模块化
- 完善文档结构和 API 参考
- 修复各种兼容性问题
- 添加模块白名单功能，解决 Windows 编译命令过长问题
- 优化代码结构和性能
- 添加无痛迁移特性，支持其他平台代码的无缝迁移

## v0.0.5 (2026-03-13)

### 完成全量测试并修复完善

- 对所有模块进行全量测试
- 修复测试中发现的问题
- 完善代码实现和文档
- 优化代码结构和性能

## v0.0.4 (2026-03-13)

### 功能增强与修复

- 修复初始化问题
- 优化方法注册机制
- 完善错误处理
- 提升代码稳定性

## v0.0.3 (2026-03-13)

### 修复编译问题

- 修复编译错误
- 优化依赖管理
- 提升构建稳定性

## v0.0.2 (2026-03-13)

### 完善 API

- 完善 API 实现
- 修复参数错误
- 优化代码结构

## v0.0.1 (2026-03-13)

### 项目初始化

- 初始化 AutoGo ScriptEngine 项目
- 实现基本架构
- 添加核心模块
