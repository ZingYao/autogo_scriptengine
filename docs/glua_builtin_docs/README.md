# GLua Builtin Docs

本目录由 `go run ./tools/glua_builtin_docs` 生成，供 `go-lua-vm`/`gluals`、VS Code 和 JetBrains 导入。

## 文件

- `autogo-scriptengine-common.json`：9 个条目
- `autogo-scriptengine-android.json`：5513 个条目
- `autogo-scriptengine-ios.json`：5306 个条目
- `autogo-scriptengine-lrappsoft.json`：170 个条目

## 重新生成

```bash
go run ./tools/glua_builtin_docs
```

## VS Code

```json
{
  "glua.builtinDocs": [
    "docs/glua_builtin_docs/autogo-scriptengine-common.json",
    "docs/glua_builtin_docs/autogo-scriptengine-android.json",
    "docs/glua_builtin_docs/autogo-scriptengine-lrappsoft.json"
  ]
}
```

iOS 项目将 `android` 文件替换为 `autogo-scriptengine-ios.json`。修改 JSON 后需要重新加载 IDE 窗口或重启语言服务。
