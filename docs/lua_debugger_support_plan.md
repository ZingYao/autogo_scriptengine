# Lua Debug 调试支持规划

## 1. 目标

为 `lua_engine` 增加语言级 Debug 能力，使 AutoGo Lua 脚本可以在 Android 与 iOS 运行环境中支持断点、暂停、恢复、单步、异常定位和变量查看。

当前仓库已有快速调试工具文档，但现有能力主要是脚本运行、日志查看、暂停/恢复/停止进程级控制。该规划面向 Lua VM 级调试，不替代现有调试器，而是为现有调试器提供底层 Debug Core。

## 2. 非目标

- 不在第一阶段实现 JavaScript 断点调试。
- 不在第一阶段接入 VSCode Debug Adapter Protocol。
- 不支持在 Go 注入方法内部做 Lua 单步，例如 `motion.click()` 内部仍属于 Go/AutoGo 原生调用。
- 不改变 Android/iOS autogo API 调用形式。
- 不修改 `AutoGo/` 参考目录。

## 3. 技术基础

Lua 引擎基于 `github.com/yuin/gopher-lua`。可以通过 Lua VM hook 在执行行、调用、返回等事件时获得调试回调，从而实现：

- 行号命中检查。
- 断点暂停。
- 单步控制。
- 脚本停止。
- 异常堆栈捕获。

当前 `LuaEngine` 已有 `Start/Pause/Resume/Stop` 和 `EngineState`，但状态控制尚未接入 VM hook，因此不能精确停在 Lua 源码行。

## 4. 总体架构

新增 `lua_engine/debugger` 子包，负责 Debug Core，不直接绑定 TUI、WebSocket 或设备管理。

```text
调试器 UI / CLI / WebSocket
        |
        v
DebuggerController
        |
        v
lua_engine.Debugger
        |
        v
gopher-lua hook
        |
        v
Lua 脚本执行
```

核心职责拆分：

- `Debugger`：维护断点、运行状态、单步模式、暂停等待。
- `DebugConfig`：控制是否启用、hook 粒度、变量采集深度。
- `DebugEvent`：向外部调试器推送当前暂停、恢复、异常、退出等事件。
- `DebugCommand`：接收继续、暂停、停止、单步、设置断点等命令。
- `FrameSnapshot`：描述当前文件、行号、函数名、调用栈和变量快照。

## 5. 建议 API

### 5.1 EngineConfig 扩展

```go
type EngineConfig struct {
    // Debug 为 Lua VM 级调试配置。nil 表示关闭调试能力。
    Debug *debugger.DebugConfig
}
```

### 5.2 DebugConfig

```go
type DebugConfig struct {
    Enabled          bool
    BreakOnStart     bool
    BreakOnError     bool
    CollectGlobals   bool
    CollectLocals    bool
    MaxVariableDepth int
    EventBufferSize  int
}
```

### 5.3 Debugger 控制接口

```go
type Controller interface {
    SetBreakpoint(file string, line int) error
    RemoveBreakpoint(file string, line int) error
    ClearBreakpoints(file string) error
    Continue() error
    Pause() error
    Stop() error
    StepInto() error
    StepOver() error
    StepOut() error
    Evaluate(expr string) (ValueSnapshot, error)
    Events() <-chan DebugEvent
}
```

第一阶段可以先实现 `SetBreakpoint`、`RemoveBreakpoint`、`Continue`、`Pause`、`Stop`、`StepInto`、`Events`。

## 6. 调试状态机

```text
Disabled
  |
  v
Running <---- Continue
  |              ^
  | breakpoint   |
  | pause        |
  v              |
Paused ----------+
  |
  | stop
  v
Stopped
```

补充状态：

- `Stepping`：单步运行到下一条可见 Lua 行后进入 `Paused`。
- `ErrorPaused`：脚本异常时暂停，等待用户查看堆栈后停止或继续。

## 7. 断点规则

断点使用 `file + line` 匹配。

文件名标准化规则：

- `ExecuteFile(path)` 使用清理后的绝对路径或相对工程路径。
- `ExecuteString(script, name)` 后续需要提供脚本名；没有脚本名时使用 `<string>`。
- embed/require 加载的脚本记录其虚拟路径。

断点命中条件：

- 当前 hook 事件为 line。
- 当前文件路径标准化后匹配。
- 当前行号匹配。
- 引擎未处于 stop 状态。

## 8. 变量查看

第一阶段变量查看只做基础能力：

- 当前全局变量快照。
- 基础类型：nil、bool、number、string。
- table 按最大深度展开。
- function、userdata、thread 只显示类型与摘要。

第二阶段再补：

- 局部变量。
- upvalue。
- 调用栈 frame 选择。
- 表循环引用检测。
- 大 table 分页。

## 9. 异常处理

当 Lua 执行返回错误时：

- 如果 `BreakOnError=true`，先发送 `DebugEventError`。
- 事件中包含错误文本、当前脚本、当前行号、调用栈。
- 外部调试器可选择停止、重启或导出日志。

Go 注入方法返回错误时，仍按现有 Lua 桥接方式处理；如果错误上抛到 Lua 执行层，则进入异常处理。

## 10. 与 Android/iOS 的关系

Debug Core 放在 `lua_engine` 通用层，不区分 Android/iOS。

平台差异只存在于：

- 设备部署。
- 脚本入口。
- autogo 模块注入目录。
- 调试器 UI 如何连接设备。

因此 Debug Core 不允许依赖 Android-only 或 iOS-only 模块。

## 11. 与现有快速调试工具的关系

现有 `autogo_scriptengine_debugger` 可以作为上层 UI 使用 Debug Core：

- TUI 显示运行状态、断点列表、当前行、调用栈。
- 通过设备日志或 RPC/WebSocket 接收 `DebugEvent`。
- 发送 `DebugCommand` 控制脚本继续、暂停、单步、停止。

该仓库先提供库能力和 example，调试工具仓库再接入 UI。

## 12. 阶段计划

### 阶段 1：Debug Core 最小闭环

- [ ] 新增 `lua_engine/debugger` 子包。
- [ ] 新增 `DebugConfig`、`DebugEvent`、`Breakpoint`、`Controller`。
- [ ] `LuaEngine` 支持启用 debug hook。
- [ ] 支持文件行断点。
- [ ] 支持 `continue`、`pause`、`stop`、`stepInto`。
- [ ] 支持命中断点后阻塞等待命令。
- [ ] 支持异常事件回传。
- [ ] 新增 `examples/lua_engine/debugger`。
- [ ] 更新 `docs/debugger/README.md` 与 Lua 引擎文档。

验收：

- Lua 脚本执行到指定行可以停住。
- 调用 `continue` 后继续执行。
- 调用 `stepInto` 后停在下一条 Lua 行。
- 调用 `stop` 后脚本退出。
- 错误脚本可以回传异常事件与堆栈。

### 阶段 2：变量与调用栈

- [ ] 支持当前全局变量快照。
- [ ] 支持 table 限深展开。
- [ ] 支持调用栈 frame 列表。
- [ ] 支持局部变量和 upvalue 读取。
- [ ] 支持变量快照大小限制。

验收：

- 断点暂停时可以看到当前关键变量。
- 大 table 不会导致调试器卡死。
- 循环引用不会导致递归死循环。

### 阶段 3：远程调试协议

- [ ] 设计轻量 Debug RPC/WebSocket 协议。
- [ ] 支持设置/删除断点。
- [ ] 支持暂停、继续、单步、停止。
- [ ] 支持事件订阅。
- [ ] 支持表达式求值。
- [ ] 在快速调试工具中接入 TUI 控制。

验收：

- 调试工具可以远程控制设备上的 Lua 脚本。
- Android/iOS 项目使用同一套 Debug Core。
- 网络断开后脚本可按策略继续、暂停或停止。

### 阶段 4：IDE 协议适配

- [ ] 评估 Debug Adapter Protocol。
- [ ] 映射断点、线程、栈帧、变量、继续、单步能力。
- [ ] 提供 VSCode 配置示例。

验收：

- VSCode 可以连接 AutoGo Lua 调试会话。
- 可以设置断点、单步、查看变量。

## 13. 风险与处理

- hook 性能开销：默认关闭 Debug；启用后允许配置 hook 粒度。
- 变量展开成本高：限制深度、数量和字符串长度。
- Go 注入方法内部不可单步：文档明确边界，只停在 Lua 调用前后。
- 多协程调试复杂：第一阶段只支持主 Lua 线程；协程调试单独排期。
- 脚本路径不一致：执行、require、embed 加载统一做路径标准化。
- 移动端长时间暂停：上层调试器需要提供超时或断线策略。

## 14. 待确认问题

- 第一阶段是否只支持主线程，协程统一视为后续能力。
- 断点文件路径以工程相对路径为准，还是以设备端绝对路径为准。
- 脚本暂停时，AutoGo 宿主是否允许长时间阻塞。
- 快速调试工具与运行中脚本之间优先使用 WebSocket、ADB 日志轮询还是本地 RPC。
