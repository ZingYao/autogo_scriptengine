# 示例工程

示例工程用于验证脚本入口、模块注入、用户 `require`、复杂参数、返回值解析和对象生命周期。实际运行时需要按示例文件顶部说明将 `package main_test` 改为 `package main`，并通过 AutoGo Extension 构建运行。

## Android autogo

- [Lua Android autogo 示例](https://github.com/ZingYao/autogo_scriptengine/tree/master/examples/lua_engine/autogo)
- [JavaScript Android autogo 示例](https://github.com/ZingYao/autogo_scriptengine/tree/master/examples/js_engine/autogo)

这两个示例都会注册 Android autogo 全量模块，并展示以下调用形态：

- 模块对象入口：`app.currentPackage()`、`device.width`、`motion.click(...)`
- 用户脚本模块：Lua 使用 `require("utils")`，JavaScript 使用 `require("./utils")`
- 复杂参数：`https.post(url, body, headers, timeout)`、`app.startActivity({...})`
- 返回值解析：`https.get(...)`、`app.getList(false)`
- 对象生命周期：`uiacc.new()`、`opencv.newMat()`、`imgui.newVec2(...)`

## iOS autogo

- [Lua iOS autogo 示例](https://github.com/ZingYao/autogo_scriptengine/tree/master/examples/lua_engine/autogo_ios)
- [JavaScript iOS autogo 示例](https://github.com/ZingYao/autogo_scriptengine/tree/master/examples/js_engine/autogo_ios)

iOS 示例只注册 iOS autogo 模块，避免注入 Android-only 模块或参数。示例重点覆盖 `app.openUrl`、`device.getDisplayInfo`、`files.*`、`https.*`、`opencv/imgui` 对象构造，以及用户 `require`。

## lrappsoft

- [Lua lrappsoft 示例](https://github.com/ZingYao/autogo_scriptengine/tree/master/examples/lua_engine/lrappsoft)

该示例用于验证懒人脚本风格 API 的兼容入口，适合迁移 lrappsoft 风格脚本时对照。

## 字节码

- [Lua 字节码示例](https://github.com/ZingYao/autogo_scriptengine/tree/master/examples/lua_engine/bytecode)

该示例用于验证 Lua 源码编译、字节码加载和执行链路。
