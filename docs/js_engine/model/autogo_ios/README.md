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

- [app](#app)
- [console](#console)
- [device](#device)
- [dotocr](#dotocr)
- [files](#files)
- [https](#https)
- [hud](#hud)
- [images](#images)
- [ime](#ime)
- [imgui](#imgui)
- [motion](#motion)
- [opencv](#opencv)
- [ppocr](#ppocr)
- [storages](#storages)
- [system](#system)
- [utils](#utils)
- [yolo](#yolo)

`uiacc` 与 `apkctl` 当前不注入：AutoGo iOS 参考目录暂未提供对应实现。

## 5. 模块导航

### app

应用管理模块，按 `app.xxx(...)` 调用。常用方法包括 `currentPackage`、`selfPackage`、`launch`、`forceStop`、`getList`、`getName`、`getVersion`、`getIcon`、`isInstalled`、`install`、`uninstall`、`clear`、`openUrl`。`getList(false)` 返回应用信息数组。

### console

控制台对象模块，模块入口为 `console.new(...)`，返回对象继续调用实例方法。对象生命周期由返回对象持有，复杂显示行为与 AutoGo iOS `console` 包保持一致。

### device

设备信息模块，按 `device.xxx(...)` 调用。常用方法包括 `getDisplayInfo`、`getBattery`、`getBatteryStatus`、`isScreenOn`、`isScreenUnlock`、`getBrightness`、`getTotalMem`、`getAvailMem`、`wakeUp`、`keepScreenOn`、`getIp`、`getWifiMac`、`reboot`。`getDisplayInfo()` 返回 `{width, height, scale, rotation}`。

### dotocr

点阵 OCR 模块，按 `dotocr.xxx(...)` 调用。iOS 版本不带 Android `displayId` 参数，图像入参可使用 `images.captureScreen(...)` 或 `images.readFromPath(...)` 的返回对象。

### files

文件模块，按 `files.xxx(...)` 调用。常用方法包括 `isFile`、`isDir`、`create`、`ensureDir`、`exists`、`read`、`readBytes`、`write`、`writeBytes`、`append`、`appendBytes`、`copy`、`move`、`rename`、`getName`、`getExtension`、`getMd5`、`remove`、`path`、`listDir`。

### https

网络请求模块，按 `https.xxx(...)` 调用。`get(url, timeout)`、`post(url, data, headers, timeout)`、`postMultipart(url, fileName, fileData, timeout)` 返回 `{code, data}` 或 `{code, body}` 结构，字节结果在 JavaScript 中按数组/字节桥接值处理。

### hud

悬浮显示对象模块，模块入口为 `hud.new(...)`，返回对象继续调用实例方法。用于 iOS 悬浮提示、状态展示和销毁控制。

### images

图像处理模块，按 `images.xxx(...)` 调用。支持截图、读写图像、取色找色、多点找色、裁剪、缩放、旋转、灰度、阈值化、Base64/bytes 编解码。返回图像对象可继续传给 `opencv`、`ppocr`、`yolo`、`dotocr`。

### ime

输入法模块，按 `ime.xxx(...)` 调用。包括 `getClipText`、`setClipText`、`inputText`。

### imgui

界面绘制模块，按 `imgui.xxx(...)` 调用。大部分函数通过反射桥按 AutoGo iOS `imgui` 导出方法注册；对象返回值继续以对象方法调用。

### motion

动作模块，按 `motion.xxx(...)` 调用，不注册全局触控入口。常用方法包括 `touchDown`、`touchMove`、`touchUp`、`click`、`longClick`、`swipe`、`swipe2`、`home`、`recents`、`volumeUp`、`volumeDown`、`keyAction`。

### opencv

OpenCV 模块，按 `opencv.xxx(...)` 调用。`Mat`、`Scalar`、`Point` 等对象返回后继续调用对象方法；使用完成后优先调用 `close()` 释放原生资源。

### ppocr

飞浆 OCR 模块，按 `ppocr.xxx(...)` 调用。iOS 版本不带 Android `displayId` 参数，识别结果统一为 JavaScript object/array。

### storages

本地存储模块，按 `storages.xxx(...)` 调用。用于创建命名存储、读写键值和清理数据。

### system

系统模块，按 `system.xxx(...)` 调用。提供 AutoGo iOS 当前导出的系统级能力，避免引入 Android 专属字段或 shell 语义。

### utils

工具模块，按 `utils.xxx(...)` 调用。用于延时、随机、编码、哈希等通用工具能力，入参与返回值按 Go 签名转换。

### yolo

目标检测模块，按 `yolo.xxx(...)` 调用。iOS 版本不带 Android `displayId` 参数，检测结果统一返回 `x/y/width/height/label/score/centerX/centerY` 等字段。

## 6. 示例

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

## 7. 注意事项

1. iOS 项目必须导入 `define/ios/autogo/...`，不要导入 Android define。
2. iOS 示例不要使用 `app.startActivity`、`app.getBrowserPackage`、`uiacc`、`apkctl` 或 `displayId` 参数。
3. `opencv`、`imgui` 等对象模块通过反射桥转换参数与返回值，复杂对象优先按返回对象继续调用方法。
