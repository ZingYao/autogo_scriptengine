# JavaScript 引擎 iOS autogo 风格包文档

## 1. 使用范围

iOS autogo 风格包位于 `js_engine/define/ios/autogo` 与 `js_engine/model/autogo_ios`。它只注入 AutoGo iOS 文档中存在的模块，避免 Android-only 模块、`displayId` 参数和 Android 专属返回字段污染 iOS 脚本。

## 2. 注册方式

```go
import "github.com/ZingYao/autogo_scriptengine/js_engine/define/ios/autogo/all_models"

engine.RegisterModule(all_models.AllModules...)
```

如需限制模块集合，可以改用 `safe_models`。当前 iOS `unsafe_models` 为空。

## 3. API 映射规则

- Go 包函数统一挂到模块对象下，例如 `app.OpenUrl` 映射为 `app.openUrl(...)`。
- 方法名按 Go 导出名做常规小驼峰转换，不保留历史别名。
- 不注册无模块名前缀的全局入口，例如触控必须使用 `motion.click(...)`。
- Go struct/map 入参使用 JavaScript object 构造，slice 入参使用 array。
- Go struct、map、slice 返回值转换为 JavaScript object/array；返回对象的方法保留在对象本身。

## 4. iOS 模块清单

当前 iOS JavaScript define 注入：

`app`、`console`、`device`、`dotocr`、`files`、`https`、`hud`、`images`、`ime`、`imgui`、`motion`、`opencv`、`ppocr`、`storages`、`system`、`utils`、`yolo`。

`uiacc` 与 `apkctl` 当前不注入：AutoGo iOS 参考目录暂未提供对应实现。

## 5. 示例

完整示例见 `examples/js_engine/autogo_ios`。

```javascript
console.log('screen: ' + device.width + 'x' + device.height);

const info = device.getDisplayInfo();
console.log('rotation: ' + info.rotation);

const resp = https.post(
    'https://example.com/api',
    JSON.stringify({ hello: 'ios-js' }),
    { 'Content-Type': 'application/json' },
    5000
);
console.log('status: ' + resp.code);

const apps = app.getList(false);
if (apps.length > 0) {
    console.log(apps[0].packageName + ' / ' + apps[0].appName);
}

const mat = opencv.newMat();
if (mat) {
    console.log('mat empty: ' + mat.empty());
    mat.close();
}
```

## 6. 注意事项

1. iOS 项目必须导入 `define/ios/autogo/...`，不要导入 Android define。
2. iOS 示例不要使用 `app.startActivity`、`app.getBrowserPackage`、`uiacc`、`apkctl` 或 `displayId` 参数。
3. `opencv`、`imgui` 等对象模块通过反射桥转换参数与返回值，复杂对象优先按返回对象继续调用方法。
