# AutoGo Debugger

AutoGo ScriptEngine Debugger v1.0.0 已发布，项目地址：

<https://github.com/ZingYao/autogo_scriptengine_debugger/releases/tag/v1.0.0>

当前推荐使用 VSCode 或 JetBrains 插件中的 AutoGo Debugger。插件已经内置调试能力，普通脚本开发不需要手动下载命令行工具。

## 当前支持范围

| 能力 | Lua / GLua | JavaScript |
| --- | --- | --- |
| 调试启动 | 支持 | 支持 |
| 断点调试 | 支持 | 支持 |
| 单步调试 | 支持 | 支持 |
| 调用栈 | 支持 | 支持 |
| 变量查看 | 支持 | 支持 |
| 表达式求值 | 支持 | 支持 |
| 变量修改 | 支持 | 支持 |
| `require` 依赖 | 支持 | 支持 |
| `importModule` 依赖 | 不适用 | 支持 |

Debugger 当前通过 IDE 插件提供 Lua/GLua 与 JavaScript 代码的 DAP 调试。JavaScript 侧基于 AutoGo 维护的 `github.com/ZingYao/goja` fork，支持断点、单步、栈帧、变量面板、表达式求值、变量修改、`require` 和 `importModule`。

## IDE 集成

### VSCode

VSCode 插件提供 AutoGo Activity Bar 入口和调试按钮。打开 `.lua`、`.glua` 或 `.js` 文件后，使用 AutoGo 面板中的调试入口即可。

![VSCode Lua/GLua Debugger 真机截图](assets/vscode-lua-debug.png)

典型流程：

1. 打开 AutoGo 项目根目录。
2. 确认 AutoGo 插件已识别项目和设备。
3. 打开需要调试的 `.lua`、`.glua` 或 `.js` 文件。
4. 在行号左侧设置断点。
5. 点击 AutoGo 面板中的 Debug 当前文件。
6. 调试启动后使用 VSCode 原生调试面板查看变量、调用栈和断点。

### JetBrains

JetBrains 插件支持 IDEA、GoLand 等 IntelliJ Platform IDE。打开 `.lua`、`.glua` 或 `.js` 文件后，可通过顶部 AutoGo 菜单、工具栏或运行配置启动调试。

![JetBrains Lua/GLua Debugger 真机截图](assets/jetbrains-lua-debug.png)

典型流程：

1. 打开 AutoGo 项目根目录。
2. 在 AutoGo 设置中确认 AG、ADB、Go、GLuac 路径已自动发现或手动配置。
3. 打开需要调试的 `.lua`、`.glua` 或 `.js` 文件。
4. 在编辑器左侧设置断点。
5. 使用 AutoGo Debug 当前文件。
6. 调试启动后使用 JetBrains 原生 Debug 工具窗口查看栈帧、变量和日志。

## 调试流程

IDE 插件在启动 Lua/GLua/JavaScript 调试时会完成以下工作：

1. 解析当前入口文件的静态依赖闭包。Lua/GLua 解析 `require`，JavaScript 解析 `require`、`importModule` 和静态 `import`。
2. 将入口文件和依赖文件增量同步到移动端远程引擎。
3. 启动或复用移动端 AutoGo ScriptEngine 远程引擎。
4. 创建 DAP 会话并映射本地源码路径与设备端 release 路径。
5. 使用 IDE 原生 Debug UI 执行断点、单步、继续、暂停、变量查看等操作。

动态依赖无法在调试前完全推断。遇到动态 `require`、动态 `importModule` 或动态 `import()` 时，需要把额外脚本文件加入项目同步配置，或改为静态依赖。

## Lua DAP 能力

当前 Lua/GLua 调试支持：

- 文件行断点。
- `continue`、`pause`、`next`、`stepIn`、`stepOut`。
- `threads`、`stackTrace`、`scopes`、`variables`。
- 局部变量、Upvalues 和 Globals 快照。
- 异常事件回传。
- 源码路径映射。
- 多文件 `require` 依赖调试。

边界说明：

- Go 注入方法内部不能按 Lua 行单步，只能停在 Lua 调用前后。
- 使用 `gluac -s` 去除调试信息后的字节码不能做源码级调试。

## JavaScript DAP 能力

当前 JavaScript 调试支持：

- 文件行断点。
- `continue`、`pause`、`next`、`stepIn`、`stepOut`。
- `threads`、`stackTrace`、`scopes`、`variables`。
- 局部变量、闭包变量和 Globals 快照。
- 表达式求值和变量修改。
- 源码路径映射。
- 多文件 `require`、`importModule` 和静态 `import` 依赖调试。

边界说明：

- JavaScript 调试依赖 AutoGo 维护的 goja fork，宿主工程会锁定 `github.com/ZingYao/goja` 的调试版本。
- Go 注入方法内部不能按 JavaScript 行单步，只能停在 JavaScript 调用前后。
- 动态依赖需要显式加入同步配置，否则移动端调试时可能找不到文件。

示例代码见：

```text
examples/lua_engine/debugger
```

## 工具版本

IDE 插件已内置 debugger 使用路径。只有在需要核对插件内置 debugger 版本或排查工具自身问题时，才需要查看独立 Release。

Release 下载地址：

<https://github.com/ZingYao/autogo_scriptengine_debugger/releases/tag/v1.0.0>

v1.0.0 提供的平台包：

| 平台 | 文件 |
| --- | --- |
| Windows x64 | `AutoGoScriptEngineDebugger-Windows.zip` |
| macOS Apple Silicon | `AutoGoScriptEngineDebugger-macOS-ARM.tar.gz` |
| macOS Intel | `AutoGoScriptEngineDebugger-macOS-AMD.tar.gz` |

macOS 解压后如遇到执行权限问题：

```bash
chmod +x AutoGoScriptEngineDebugger*
```

## 常见问题

### JavaScript 能不能使用 Debugger？

可以。当前 debugger 支持 JavaScript DAP 调试，入口文件使用 `.js` 后缀即可。调试器会启动独立的 JavaScript DAP 端口，并通过 goja fork 暴露断点、单步、调用栈、变量、表达式求值和变量修改能力。

### 为什么断点没有命中？

优先检查：

1. 当前文件是否为 `.lua`、`.glua` 或 `.js`。
2. 调试启动前入口文件和依赖文件是否已成功同步。
3. 断点行是否是可执行语句。
4. Lua/GLua 是否使用了 strip 后的 `.gluac` 字节码。
5. 动态 `require` 的文件是否已加入额外同步配置。

### 截图中为什么没有真实路径？

文档截图已做脱敏处理。实际 IDE 中会显示你的本地项目路径、设备序列号和运行目录；对外反馈问题时请同样打码这些信息。
