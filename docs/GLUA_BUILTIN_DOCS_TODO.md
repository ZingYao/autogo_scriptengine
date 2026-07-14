# GLua 自定义方法提示目录 TODO

## 目标

为 AutoGo ScriptEngine 暴露给 Lua 的全局函数和 table 方法生成 `go-lua-vm`/`gluals` 可直接导入的 builtin docs JSON，使 VS Code、JetBrains 和 `gluals` 能提供方法补全、签名提示、悬浮文档与定义跳转。

上游导入格式以 `go-lua-vm` 的 `builtin-functions.json` 为准，自定义目录通过以下入口加载：

- `gluals --gluals-builtin-docs <path>`
- VS Code 设置 `glua.builtinDocs`
- JetBrains 设置 `Builtin docs JSON files`

## 范围

### 包含

- `lua_engine` 提供的公共全局函数。
- `console` 等引擎核心 table 方法。
- Android autogo 模块注册的方法。
- iOS autogo 模块注册的方法。
- lrappsoft 兼容模块注册的方法。
- `RegisterMethod` 静态注册的方法名、说明、参数和返回值。
- ImGui 等生成代码中可静态解析的批量注册方法。

### 不包含

- `go-lua-vm` 已内置的 `glua.*`、Lua 标准库和 Event 提示定义。
- 运行时由用户脚本通过 `registerMethod` 动态注册的方法。
- 删除、重命名或改变现有 Lua API。
- 修改 `go-lua-vm` 仓库；本任务只生成可供其导入的外部目录文件。

## 需求到产物映射

| 需求 | 代码或产物 |
| --- | --- |
| 公共方法提示 | `autogo-scriptengine-common.json` |
| Android autogo 方法提示 | `autogo-scriptengine-android.json` |
| iOS autogo 方法提示 | `autogo-scriptengine-ios.json` |
| lrappsoft 方法提示 | `autogo-scriptengine-lrappsoft.json` |
| 可重复生成 | builtin docs 生成工具或脚本 |
| 冲突可审计 | 生成统计与冲突报告 |
| 用户可导入 | 文档中的 `glua.builtinDocs` 和 `--gluals-builtin-docs` 示例 |

## 输出格式

每个 JSON 使用上游支持的多语言结构：

```json
{
  "functions": {
    "table.method": {
      "signature": {
        "en": "table.method(value)",
        "zh-CN": "table.method(value)"
      },
      "description": {
        "en": "Method description.",
        "zh-CN": "方法说明。"
      },
      "params": {
        "en": ["value: input value"],
        "zh-CN": ["value：输入值"]
      },
      "returns": {
        "en": "returns: result value.",
        "zh-CN": "返回：结果值。"
      },
      "example": {
        "en": "local result = table.method(value)",
        "zh-CN": "local result = table.method(value)"
      }
    }
  }
}
```

## 实施清单

### 1. 建立生成器

- [ ] 使用 Go AST 扫描 `lua_engine`，不通过启动脚本引擎来收集方法，避免 Android/CGO 环境依赖。
- [ ] 识别 `RegisterMethod("table.method", description, function, overridable)` 调用。
- [ ] 识别 `glua.Register`、`installVMCoreTable` 等核心全局/table 注册入口。
- [ ] 支持函数字面量、函数标识符、包选择器和方法表达式。
- [ ] 对字符串拼接或循环生成的方法名进行常量求值；无法解析时写入报告并使校验失败。
- [ ] 对生成结果按限定方法名稳定排序，保证重复运行得到一致 diff。

### 2. 转换 Lua 签名

- [ ] 保留源码中可获得的参数名；匿名参数使用 `arg1`、`arg2` 等稳定名称。
- [ ] 将 Go 可变参数转换为 Lua `...` 表达。
- [ ] 将 Go 类型映射为 Lua 提示类型：
  - `string`、`[]byte` → `string`
  - 整数、浮点数 → `number`
  - `bool` → `boolean`
  - slice、array、map、struct → `table`
  - 指针及宿主对象 → `userdata`
  - `interface{}`/`any` → `any`
  - `func` → `function`
- [ ] `error` 只进入错误语义说明，不作为正常 Lua 返回值展示。
- [ ] 多返回值按 Lua 实际返回顺序写入 `returns`。
- [ ] 无法精确判断的类型回退到 `any`，并记录回退位置。

### 3. 拆分目录

- [ ] `common`：引擎公共全局函数与核心 table。
- [ ] `android`：`lua_engine/model/autogo` 及 Android define 实际导出的模块。
- [ ] `ios`：`lua_engine/model/autogo_ios` 及 iOS define 实际导出的模块。
- [ ] `lrappsoft`：`lua_engine/model/lrappsoft` 实际导出的模块。
- [ ] 每个目录可独立加载；`common + 平台 + lrappsoft` 可以按需叠加。

### 4. 处理重复和冲突

- [ ] 同目录内同名且签名相同的条目合并。
- [ ] 同目录内同名但签名不同的条目按照当前实际注册顺序确定最终生效定义。
- [ ] 为被覆盖定义输出方法名、来源文件、来源行、候选签名和最终选择。
- [ ] Android 与 iOS 同名差异保留在各自目录，不跨平台覆盖。
- [ ] 不覆盖上游已有 `glua.*` 定义。
- [ ] 对 ImGui 大量同名注册进行专项审计，确保提示与运行时最后生效的方法一致。

### 5. 补全文档

- [ ] 在 Lua 引擎文档中说明四个目录的用途和组合方式。
- [ ] 添加 VS Code 配置示例：

```json
{
  "glua.builtinDocs": [
    "/absolute/path/autogo-scriptengine-common.json",
    "/absolute/path/autogo-scriptengine-android.json",
    "/absolute/path/autogo-scriptengine-lrappsoft.json"
  ]
}
```

- [ ] 添加 `gluals` 命令行示例：

```bash
gluals \
  --gluals-builtin-docs /absolute/path/autogo-scriptengine-common.json \
  --gluals-builtin-docs /absolute/path/autogo-scriptengine-android.json
```

- [ ] 添加 JetBrains `Builtin docs JSON files` 配置说明。
- [ ] 说明修改提示文件后需要重启语言服务或重新加载 IDE 窗口。

### 6. 验证

- [ ] 所有输出均可由标准 JSON 解析器加载。
- [ ] 使用 `gluals --gluals-builtin-docs` 验证每个文件可被接受。
- [ ] 比较静态注册方法集合与生成条目集合，缺失项必须为零或列入明确排除项。
- [ ] 检查每个条目至少包含 `signature`、`description`、`params` 和 `returns`。
- [ ] 抽查 `app`、`files`、`images`、`cryptLib`、`json`、`console` 和 `imgui`。
- [ ] 验证重复运行生成器不会产生无意义差异。
- [ ] 运行生成器自身测试以及当前环境可执行的 Go 测试。

## 风险与约束

- 当前源码约有数千个唯一注册名称，手工维护不可行，必须自动生成。
- Android、iOS 和 lrappsoft 存在同名但不同签名的方法，单一合并目录会产生错误提示。
- 反射桥和生成代码包含大量重复注册；提示目录只能表达一个最终签名，需要与运行时覆盖顺序保持一致。
- Go 反射无法获得参数名，因此参数名应优先来自 AST；无法取得时使用稳定占位名。
- 生成过程不得启动额外 `gopls`、后台服务或文件 watcher。
- 生成产物只描述当前 API，不应被解释为新增业务能力。

## 验收标准

- [ ] 四个 builtin docs JSON 均已生成并可独立导入。
- [ ] 公共、目标平台和 lrappsoft 目录可以组合加载。
- [ ] 无未解释的静态注册方法缺失。
- [ ] 所有冲突和类型回退均有可审计报告。
- [ ] IDE 能对代表性 table 提供成员补全、签名和中文悬浮说明。
- [ ] 没有修改现有 Lua API 的运行时行为。
