# app 模块

## 模块简介

app 模块提供了应用管理功能，包括应用启动、安装、卸载、信息获取等。

## 方法列表

### app.currentPackage
获取当前页面应用包名

**参数：**

无参数

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| packageName | string | 当前应用的包名 |

**使用示例：**
```javascript
// 获取当前应用包名
var pkg = app.currentPackage();
console.log("当前应用包名: " + pkg);
```

***

### app.currentActivity
获取当前页面应用类名

**参数：**

无参数

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| activityName | string | 当前应用的类名 |

**使用示例：**
```javascript
// 获取当前应用类名
var activity = app.currentActivity();
console.log("当前应用类名: " + activity);
```

***

### app.launch
通过应用包名启动应用

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| packageName | string | 应用包名 |
| displayId | number | 显示ID（可选，默认为0） |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| success | boolean | 是否启动成功 |

**使用示例：**
```javascript
// 启动应用
var result = app.launch("com.example.app", 0);
if (result) {
    console.log("应用启动成功");
}
```

***

### app.getList
获取应用列表

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| includeSystemApps | boolean | 是否包含系统应用（可选，默认为true） |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| apps | array | 应用信息数组，每个元素包含包名、名称等信息 |

**使用示例：**
```javascript
// 获取应用列表
var apps = app.getList(true);
for (var i = 0; i < apps.length; i++) {
    console.log("应用: " + apps[i].name + " (" + apps[i].packageName + ")");
}
```

***

### app.getName
获取应用名称

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| packageName | string | 应用包名 |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| appName | string | 应用名称 |

**使用示例：**
```javascript
// 获取应用名称
var name = app.getName("com.example.app");
console.log("应用名称: " + name);
```

***

### app.getIcon
获取应用图标

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| packageName | string | 应用包名 |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| iconData | array | 图标数据（字节数组） |

**使用示例：**
```javascript
// 获取应用图标
var iconData = app.getIcon("com.example.app");
console.log("图标数据长度: " + iconData.length);
```

***

### app.getVersion
获取应用版本

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| packageName | string | 应用包名 |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| version | string | 应用版本号 |

**使用示例：**
```javascript
// 获取应用版本
var version = app.getVersion("com.example.app");
console.log("应用版本: " + version);
```

***

### app.openSetting
打开应用的详情页(设置页)

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| packageName | string | 应用包名 |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| success | boolean | 是否打开成功 |

**使用示例：**
```javascript
// 打开应用设置页
var result = app.openSetting("com.example.app");
if (result) {
    console.log("设置页打开成功");
}
```

***

### app.viewFile
用其他应用查看文件

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| path | string | 文件路径 |

**返回值：**

无返回值

**使用示例：**
```javascript
// 用其他应用查看文件
app.viewFile("/sdcard/test.txt");
```

***

### app.editFile
用其他应用编辑文件

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| path | string | 文件路径 |

**返回值：**

无返回值

**使用示例：**
```javascript
// 用其他应用编辑文件
app.editFile("/sdcard/test.txt");
```

***

### app.uninstall
卸载应用

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| packageName | string | 应用包名 |

**返回值：**

无返回值

**使用示例：**
```javascript
// 卸载应用
app.uninstall("com.example.app");
```

***

### app.install
安装应用

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| path | string | APK 文件路径 |

**返回值：**

无返回值

**使用示例：**
```javascript
// 安装应用
app.install("/sdcard/app.apk");
```

***

### app.isInstalled
判断是否已经安装某个应用

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| packageName | string | 应用包名 |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| installed | boolean | 是否已安装 |

**使用示例：**
```javascript
// 判断应用是否已安装
var installed = app.isInstalled("com.example.app");
if (installed) {
    console.log("应用已安装");
}
```

***

### app.clear
清除应用数据

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| packageName | string | 应用包名 |

**返回值：**

无返回值

**使用示例：**
```javascript
// 清除应用数据
app.clear("com.example.app");
```

***

### app.forceStop
强制停止应用

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| packageName | string | 应用包名 |

**返回值：**

无返回值

**使用示例：**
```javascript
// 强制停止应用
app.forceStop("com.example.app");
```

***

### app.disable
禁用应用

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| packageName | string | 应用包名 |

**返回值：**

无返回值

**使用示例：**
```javascript
// 禁用应用
app.disable("com.example.app");
```

***

### app.enableAccessibility
启用无障碍服务

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| packageName | string | 应用包名 |

**返回值：**

无返回值

**使用示例：**
```javascript
// 启用无障碍服务
app.enableAccessibility("com.example.app");
```

***

### app.disableAccessibility
禁用无障碍服务

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| packageName | string | 应用包名 |

**返回值：**

无返回值

**使用示例：**
```javascript
// 禁用无障碍服务
app.disableAccessibility("com.example.app");
```

***

### app.ignoreBattOpt
忽略电池优化

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| packageName | string | 应用包名 |

**返回值：**

无返回值

**使用示例：**
```javascript
// 忽略电池优化
app.ignoreBattOpt("com.example.app");
```

***

### app.enable
启用应用

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| packageName | string | 应用包名 |

**返回值：**

无返回值

**使用示例：**
```javascript
// 启用应用
app.enable("com.example.app");
```

***

### app.getBrowserPackage
获取系统默认浏览器包名

**参数：**

无参数

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| packageName | string | 默认浏览器包名 |

**使用示例：**
```javascript
// 获取默认浏览器包名
var browserPkg = app.getBrowserPackage();
console.log("默认浏览器: " + browserPkg);
```

***

### app.openUrl
用浏览器打开网站url

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| url | string | 要打开的网址 |

**返回值：**

无返回值

**使用示例：**
```javascript
// 用浏览器打开网址
app.openUrl("https://www.example.com");
```

***

### app.startActivity
通过 Intent 启动 Activity

**参数：**

- `options` (object): Intent 配置选项
  - `action` (string): Intent 动作
  - `type` (string): Intent 类型
  - `data` (string): Intent 数据
  - `packageName` (string): 目标包名

**使用示例：**

```javascript
// 调用 app.startActivity 方法
app.startActivity({
    action: "android.intent.action.VIEW",
    type: "text/plain",
    data: "Hello World"
});
```

***

### app.sendBroadcast
发送广播消息

**参数：**

- `options` (object): Intent 配置选项
  - `action` (string): Intent 动作
  - `type` (string): Intent 类型
  - `data` (string): Intent 数据
  - `packageName` (string): 目标包名

**使用示例：**

```javascript
// 调用 app.sendBroadcast 方法
app.sendBroadcast({
    action: "com.example.ACTION",
    type: "text/plain",
    data: "Broadcast Message"
});
```

***

### app.startService
启动服务

**参数：**

- `options` (object): Intent 配置选项
  - `action` (string): Intent 动作
  - `type` (string): Intent 类型
  - `data` (string): Intent 数据
  - `packageName` (string): 目标包名

**使用示例：**

```javascript
// 调用 app.startService 方法
app.startService({
    action: "com.example.SERVICE_ACTION",
    type: "text/plain",
    data: "Service Data"
});
```

***

## 综合使用示例

### 示例1：启动应用

```javascript
app.launch("com.example.app", 0);
var currentPkg = app.currentPackage();
console.log("当前应用: " + currentPkg);
```

