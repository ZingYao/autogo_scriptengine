# Plugin 模块

Plugin 模块提供了加载外部 APK 插件的功能，用于扩展应用程序的功能。

## 方法列表

### plugin.loadApk(apkPath)
加载外部 APK 文件作为插件。

**入参：**
- `apkPath`: APK 文件路径（字符串）

**出参：** ClassLoader 对象，用于加载插件中的类

**调用示例：**
```javascript
// 加载 APK 插件
var classLoader = plugin.loadApk("/sdcard/plugin.apk");
console.println("插件加载成功");
```

## 完整示例

```javascript
// 示例1：加载单个插件
function loadSinglePlugin() {
    var apkPath = "/sdcard/my_plugin.apk";
    
    if (files.exists(apkPath)) {
        console.println("开始加载插件:", apkPath);
        var classLoader = plugin.loadApk(apkPath);
        
        if (classLoader) {
            console.println("插件加载成功");
            // 使用 classLoader 加载插件中的类
        } else {
            console.println("插件加载失败");
        }
    } else {
        console.println("插件文件不存在:", apkPath);
    }
}

// 示例2：加载多个插件
function loadMultiplePlugins() {
    var plugins = [
        "/sdcard/plugin1.apk",
        "/sdcard/plugin2.apk",
        "/sdcard/plugin3.apk"
    ];
    
    for (var i = 0; i < plugins.length; i++) {
        var apkPath = plugins[i];
        
        if (files.exists(apkPath)) {
            console.println("加载插件 " + (i + 1) + ":", apkPath);
            var classLoader = plugin.loadApk(apkPath);
            
            if (classLoader) {
                console.println("插件 " + (i + 1) + " 加载成功");
            } else {
                console.println("插件 " + (i + 1) + " 加载失败");
            }
        } else {
            console.println("插件文件不存在:", apkPath);
        }
    }
}

// 示例3：从目录加载所有插件
function loadPluginsFromDirectory(directory) {
    var pluginFiles = files.list(directory);
    var loadedCount = 0;
    
    console.println("在目录中查找插件:", directory);
    
    for (var i = 0; i < pluginFiles.length; i++) {
        var fileName = pluginFiles[i];
        
        if (fileName.endsWith(".apk")) {
            var apkPath = directory + "/" + fileName;
            console.println("加载插件:", fileName);
            
            var classLoader = plugin.loadApk(apkPath);
            
            if (classLoader) {
                console.println("插件加载成功:", fileName);
                loadedCount++;
            } else {
                console.println("插件加载失败:", fileName);
            }
        }
    }
    
    console.println("共加载 " + loadedCount + " 个插件");
}

// 示例4：带错误处理的插件加载
function loadPluginWithErrorHandling(apkPath) {
    try {
        if (!files.exists(apkPath)) {
            console.println("错误：插件文件不存在");
            return null;
        }
        
        console.println("加载插件:", apkPath);
        var classLoader = plugin.loadApk(apkPath);
        
        if (classLoader) {
            console.println("插件加载成功");
            return classLoader;
        } else {
            console.println("错误：插件加载失败");
            return null;
        }
    } catch (e) {
        console.println("加载插件时发生异常:", e.message);
        return null;
    }
}

// 示例5：插件管理器
function PluginManager() {
    this.plugins = {};
    
    this.load = function(apkPath, name) {
        if (this.plugins[name]) {
            console.println("插件已存在:", name);
            return false;
        }
        
        var classLoader = plugin.loadApk(apkPath);
        
        if (classLoader) {
            this.plugins[name] = {
                path: apkPath,
                loader: classLoader,
                loadedAt: new Date()
            };
            console.println("插件加载成功:", name);
            return true;
        } else {
            console.println("插件加载失败:", name);
            return false;
        }
    };
    
    this.unload = function(name) {
        if (this.plugins[name]) {
            delete this.plugins[name];
            console.println("插件已卸载:", name);
            return true;
        } else {
            console.println("插件不存在:", name);
            return false;
        }
    };
    
    this.list = function() {
        var names = [];
        for (var name in this.plugins) {
            names.push(name);
        }
        return names;
    };
    
    this.get = function(name) {
        return this.plugins[name];
    };
}

// 使用插件管理器
function usePluginManager() {
    var manager = new PluginManager();
    
    // 加载插件
    manager.load("/sdcard/plugin1.apk", "plugin1");
    manager.load("/sdcard/plugin2.apk", "plugin2");
    
    // 列出所有插件
    var pluginList = manager.list();
    console.println("已加载的插件:", pluginList.join(", "));
    
    // 获取插件信息
    var pluginInfo = manager.get("plugin1");
    if (pluginInfo) {
        console.println("插件路径:", pluginInfo.path);
        console.println("加载时间:", pluginInfo.loadedAt);
    }
    
    // 卸载插件
    manager.unload("plugin1");
}

// 调用示例
loadSinglePlugin();
loadMultiplePlugins();
loadPluginsFromDirectory("/sdcard/plugins");
loadPluginWithErrorHandling("/sdcard/test_plugin.apk");
usePluginManager();
```

## 注意事项

1. 加载 APK 插件需要相应的权限
2. 插件文件必须是有效的 APK 格式
3. 加载插件后，返回的 ClassLoader 对象可以用于加载插件中的类
4. 建议在加载插件前检查文件是否存在
5. 插件加载可能会失败，需要做好错误处理
6. 加载的插件会占用系统资源，使用完毕后应适当释放
7. 不同插件之间可能有依赖关系，需要按正确顺序加载
8. 插件的安全性需要自行验证，建议只加载可信来源的插件
