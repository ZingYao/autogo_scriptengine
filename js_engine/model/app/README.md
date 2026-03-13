# App 模块

App 模块提供了应用管理相关的功能，包括获取应用信息、启动应用、安装卸载应用等操作。

## 方法列表

### app.currentPackage()

获取当前页面应用包名。

**返回值**: `string`

**调用示例**:
```javascript
const packageName = app.currentPackage();
console.log("当前应用包名: " + packageName);
```

### app.currentActivity()

获取当前页面应用类名。

**返回值**: `string`

**调用示例**:
```javascript
const activity = app.currentActivity();
console.log("当前应用类名: " + activity);
```

### app.launch(packageName, displayId)

通过应用包名启动应用。

**参数**:
- `packageName` (string): 应用包名
- `displayId` (number): 显示器ID，默认为0

**返回值**: `boolean` - 如果该包名对应的应用不存在，则返回 `false`；否则返回 `true`

**调用示例**:
```javascript
const success = app.launch("com.example.app", 0);
if (success) {
    console.log("应用启动成功");
} else {
    console.log("应用启动失败");
}
```

### app.getList(includeSystemApps)

获取应用列表。

**参数**:
- `includeSystemApps` (boolean): 是否包含系统应用，默认为true

**返回值**: `Array` - 应用信息数组

**调用示例**:
```javascript
const apps = app.getList(true);
for (let i = 0; i < apps.length; i++) {
    console.log("应用名称: " + apps[i].name);
    console.log("应用包名: " + apps[i].packageName);
}
```

### app.getName(packageName)

获取应用名称。

**参数**:
- `packageName` (string): 应用包名

**返回值**: `string` - 应用名称

**调用示例**:
```javascript
const appName = app.getName("com.example.app");
console.log("应用名称: " + appName);
```

### app.getIcon(packageName)

获取应用图标。

**参数**:
- `packageName` (string): 应用包名

**返回值**: `Array` - 图标字节数组

**调用示例**:
```javascript
const iconBytes = app.getIcon("com.example.app");
console.log("图标数据长度: " + iconBytes.length);
```

### app.getVersion(packageName)

获取应用版本。

**参数**:
- `packageName` (string): 应用包名

**返回值**: `string` - 应用版本号

**调用示例**:
```javascript
const version = app.getVersion("com.example.app");
console.log("应用版本: " + version);
```

### app.openAppSetting(packageName)

打开应用的详情页(设置页)。

**参数**:
- `packageName` (string): 应用包名

**返回值**: `boolean` - 如果找不到该应用，返回 `false`；否则返回 `true`

**调用示例**:
```javascript
const success = app.openAppSetting("com.example.app");
if (success) {
    console.log("已打开应用设置页");
}
```

### app.viewFile(path)

用其他应用查看文件。

**参数**:
- `path` (string): 文件路径

**返回值**: `undefined`

**调用示例**:
```javascript
app.viewFile("/sdcard/test.txt");
```

### app.editFile(path)

用其他应用编辑文件。

**参数**:
- `path` (string): 文件路径

**返回值**: `undefined`

**调用示例**:
```javascript
app.editFile("/sdcard/test.txt");
```

### app.uninstall(packageName)

卸载应用。

**参数**:
- `packageName` (string): 应用包名

**返回值**: `undefined`

**调用示例**:
```javascript
app.uninstall("com.example.app");
```

### app.install(path)

安装应用。

**参数**:
- `path` (string): APK文件路径

**返回值**: `undefined`

**调用示例**:
```javascript
app.install("/sdcard/app.apk");
```

### app.isInstalled(packageName)

判断是否已经安装某个应用。

**参数**:
- `packageName` (string): 应用包名

**返回值**: `boolean` - 如果已安装返回 `true`，否则返回 `false`

**调用示例**:
```javascript
if (app.isInstalled("com.example.app")) {
    console.log("应用已安装");
} else {
    console.log("应用未安装");
}
```

### app.clear(packageName)

清除应用数据。

**参数**:
- `packageName` (string): 应用包名

**返回值**: `undefined`

**调用示例**:
```javascript
app.clear("com.example.app");
```

### app.forceStop(packageName)

强制停止应用。

**参数**:
- `packageName` (string): 应用包名

**返回值**: `undefined`

**调用示例**:
```javascript
app.forceStop("com.example.app");
```

### app.disable(packageName)

禁用应用。

**参数**:
- `packageName` (string): 应用包名

**返回值**: `undefined`

**调用示例**:
```javascript
app.disable("com.example.app");
```

### app.enable(packageName)

启用应用。

**参数**:
- `packageName` (string): 应用包名

**返回值**: `undefined`

**调用示例**:
```javascript
app.enable("com.example.app");
```

### app.ignoreBattOpt(packageName)

忽略电池优化。

**参数**:
- `packageName` (string): 应用包名

**返回值**: `undefined`

**调用示例**:
```javascript
app.ignoreBattOpt("com.example.app");
```

### app.enableAccessibility(packageName)

启用无障碍服务。

**参数**:
- `packageName` (string): 应用包名

**返回值**: `undefined`

**调用示例**:
```javascript
app.enableAccessibility("com.example.app");
```

### app.disableAccessibility(packageName)

禁用无障碍服务。

**参数**:
- `packageName` (string): 应用包名

**返回值**: `undefined`

**调用示例**:
```javascript
app.disableAccessibility("com.example.app");
```

### app.getBrowserPackage()

获取系统默认浏览器包名。

**返回值**: `string` - 浏览器包名

**调用示例**:
```javascript
const browser = app.getBrowserPackage();
console.log("默认浏览器: " + browser);
```

### app.openUrl(url)

用浏览器打开网站url。

**参数**:
- `url` (string): 要打开的URL

**返回值**: `undefined`

**调用示例**:
```javascript
app.openUrl("https://example.com");
```

### app.startActivity(options)

启动Activity。

**参数**:
- `options` (object): Intent选项对象
  - `action` (string): Intent动作
  - `type` (string): MIME类型
  - `data` (string): Intent数据
  - `packageName` (string): 目标包名

**返回值**: `undefined`

**调用示例**:
```javascript
app.startActivity({
    action: "android.intent.action.VIEW",
    type: "text/plain",
    data: "content://test"
});
```

### app.sendBroadcast(options)

发送广播。

**参数**:
- `options` (object): Intent选项对象
  - `action` (string): Intent动作
  - `type` (string): MIME类型
  - `data` (string): Intent数据
  - `packageName` (string): 目标包名

**返回值**: `undefined`

**调用示例**:
```javascript
app.sendBroadcast({
    action: "com.example.ACTION_TEST",
    packageName: "com.example.app"
});
```

### app.startService(options)

启动Service。

**参数**:
- `options` (object): Intent选项对象
  - `action` (string): Intent动作
  - `type` (string): MIME类型
  - `data` (string): Intent数据
  - `packageName` (string): 目标包名

**返回值**: `undefined`

**调用示例**:
```javascript
app.startService({
    action: "com.example.SERVICE_ACTION",
    packageName: "com.example.app"
});
```

## 完整示例

```javascript
// 获取当前应用信息
const currentPackage = app.currentPackage();
const currentActivity = app.currentActivity();
console.log("当前应用: " + currentPackage);
console.log("当前Activity: " + currentActivity);

// 启动应用
if (app.launch("com.example.app", 0)) {
    console.log("应用启动成功");
}

// 检查应用是否安装
if (app.isInstalled("com.example.app")) {
    console.log("应用已安装");
    const appName = app.getName("com.example.app");
    const version = app.getVersion("com.example.app");
    console.log("应用名称: " + appName);
    console.log("应用版本: " + version);
}

// 打开URL
app.openUrl("https://example.com");

// 获取应用列表
const apps = app.getList(false);
console.log("非系统应用数量: " + apps.length);
```
