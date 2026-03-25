# app 模块

## 模块简介

app 模块提供了应用管理功能，包括应用启动、安装、卸载、信息获取等。

## 方法列表

### app.currentPackage
获取当前页面应用包名

**使用示例：**
```javascript
// 调用 app.currentPackage 方法
app.currentPackage();
```

---

### app.currentActivity
获取当前页面应用类名

**使用示例：**
```javascript
// 调用 app.currentActivity 方法
app.currentActivity();
```

---

### app.launch
通过应用包名启动应用

**使用示例：**
```javascript
// 调用 app.launch 方法
app.launch();
```

---

### app.getList
获取应用列表

**使用示例：**
```javascript
// 调用 app.getList 方法
app.getList();
```

---

### app.getName
获取应用名称

**使用示例：**
```javascript
// 调用 app.getName 方法
app.getName();
```

---

### app.getIcon
获取应用图标

**使用示例：**
```javascript
// 调用 app.getIcon 方法
app.getIcon();
```

---

### app.getVersion
获取应用版本

**使用示例：**
```javascript
// 调用 app.getVersion 方法
app.getVersion();
```

---

### app.openSetting
打开应用的详情页(设置页)

**使用示例：**
```javascript
// 调用 app.openSetting 方法
app.openSetting();
```

---

### app.viewFile
用其他应用查看文件

**使用示例：**
```javascript
// 调用 app.viewFile 方法
app.viewFile();
```

---

### app.editFile
用其他应用编辑文件

**使用示例：**
```javascript
// 调用 app.editFile 方法
app.editFile();
```

---

### app.uninstall
卸载应用

**使用示例：**
```javascript
// 调用 app.uninstall 方法
app.uninstall();
```

---

### app.install
安装应用

**使用示例：**
```javascript
// 调用 app.install 方法
app.install();
```

---

### app.isInstalled
判断是否已经安装某个应用

**使用示例：**
```javascript
// 调用 app.isInstalled 方法
app.isInstalled();
```

---

### app.clear
清除应用数据

**使用示例：**
```javascript
// 调用 app.clear 方法
app.clear();
```

---

### app.forceStop
强制停止应用

**使用示例：**
```javascript
// 调用 app.forceStop 方法
app.forceStop();
```

---

### app.disable
禁用应用

**使用示例：**
```javascript
// 调用 app.disable 方法
app.disable();
```

---

### app.enableAccessibility
启用无障碍服务

**使用示例：**
```javascript
// 调用 app.enableAccessibility 方法
app.enableAccessibility();
```

---

### app.disableAccessibility
禁用无障碍服务

**使用示例：**
```javascript
// 调用 app.disableAccessibility 方法
app.disableAccessibility();
```

---

### app.ignoreBattOpt
忽略电池优化

**使用示例：**
```javascript
// 调用 app.ignoreBattOpt 方法
app.ignoreBattOpt();
```

---

### app.enable
启用应用

**使用示例：**
```javascript
// 调用 app.enable 方法
app.enable();
```

---

### app.getBrowserPackage
获取系统默认浏览器包名

**使用示例：**
```javascript
// 调用 app.getBrowserPackage 方法
app.getBrowserPackage();
```

---

### app.openUrl
用浏览器打开网站url

**使用示例：**
```javascript
// 调用 app.openUrl 方法
app.openUrl();
```

---

## 综合使用示例

### 示例1：启动应用
```javascript
app.launch("com.example.app", 0);
var currentPkg = app.currentPackage();
console.log("当前应用: " + currentPkg);
```