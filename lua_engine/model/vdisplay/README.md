# Vdisplay 模块

Vdisplay（Virtual Display）模块提供了虚拟显示设备的创建和管理功能，用于在虚拟屏幕中运行应用。

## 方法列表

### vdisplay.create(width, height, dpi)
创建一个虚拟显示设备。

**入参：**
- `width`: 显示宽度（整数）
- `height`: 显示高度（整数）
- `dpi`: 显示密度（整数）

**出参：** Vdisplay 对象

**调用示例：**
```javascript
// 创建虚拟显示设备
var v = vdisplay.create(1080, 1920, 320);
console.println("虚拟显示设备已创建");
```

### vdisplay.getDisplayId(v)
获取虚拟显示设备的 DisplayId。

**入参：**
- `v`: Vdisplay 对象

**出参：** Display ID（整数）

**调用示例：**
```javascript
// 获取 Display ID
var displayId = vdisplay.getDisplayId(v);
console.println("Display ID:", displayId);
```

### vdisplay.launchApp(v, packageName)
启动指定包名的应用到虚拟显示设备内。

**入参：**
- `v`: Vdisplay 对象
- `packageName`: 应用包名（字符串）

**出参：** 布尔值，true 表示成功，false 表示失败

**调用示例：**
```javascript
// 启动应用到虚拟显示
var success = vdisplay.launchApp(v, "com.android.settings");
if (success) {
    console.println("应用启动成功");
} else {
    console.println("应用启动失败");
}
```

### vdisplay.setTitle(v, title)
设置预览窗口标题。

**入参：**
- `v`: Vdisplay 对象
- `title`: 窗口标题（字符串）

**出参：** 无

**调用示例：**
```javascript
// 设置窗口标题
vdisplay.setTitle(v, "虚拟显示预览");
```

### vdisplay.setTouchCallback(v, callback)
设置触控回调函数。

**入参：**
- `v`: Vdisplay 对象
- `callback`: 回调函数，参数为 (x, y, action, displayId)

**出参：** 无

**调用示例：**
```javascript
// 设置触控回调
vdisplay.setTouchCallback(v, function(x, y, action, displayId) {
    console.println("触控事件:", x, y, action, displayId);
});
```

### vdisplay.showPreviewWindow(v, rotated)
显示预览窗口。

**入参：**
- `v`: Vdisplay 对象
- `rotated`: 是否旋转（布尔值）

**出参：** 无

**调用示例：**
```javascript
// 显示预览窗口
vdisplay.showPreviewWindow(v, false);
```

### vdisplay.hidePreviewWindow(v)
隐藏预览窗口。

**入参：**
- `v`: Vdisplay 对象

**出参：** 无

**调用示例：**
```javascript
// 隐藏预览窗口
vdisplay.hidePreviewWindow(v);
```

### vdisplay.setPreviewWindowSize(v, width, height)
设置预览窗口大小。

**入参：**
- `v`: Vdisplay 对象
- `width`: 窗口宽度（整数）
- `height`: 窗口高度（整数）

**出参：** 无

**调用示例：**
```javascript
// 设置预览窗口大小
vdisplay.setPreviewWindowSize(v, 540, 960);
```

### vdisplay.setPreviewWindowPos(v, x, y)
设置预览窗口位置。

**入参：**
- `v`: Vdisplay 对象
- `x`: 窗口左上角 X 坐标（整数）
- `y`: 窗口左上角 Y 坐标（整数）

**出参：** 无

**调用示例：**
```javascript
// 设置预览窗口位置
vdisplay.setPreviewWindowPos(v, 100, 100);
```

### vdisplay.destroy(v)
销毁指定的虚拟显示设备。

**入参：**
- `v`: Vdisplay 对象

**出参：** 无

**调用示例：**
```javascript
// 销毁虚拟显示设备
vdisplay.destroy(v);
```

## 完整示例

```javascript
// 示例1：创建虚拟显示设备
function createVirtualDisplay() {
    var width = 1080;
    var height = 1920;
    var dpi = 320;
    
    console.println("创建虚拟显示设备...");
    var v = vdisplay.create(width, height, dpi);
    
    var displayId = vdisplay.getDisplayId(v);
    console.println("虚拟显示设备已创建，Display ID:", displayId);
    
    return v;
}

// 示例2：启动应用到虚拟显示
function launchAppInVirtualDisplay(v, packageName) {
    console.println("启动应用:", packageName);
    var success = vdisplay.launchApp(v, packageName);
    
    if (success) {
        console.println("应用启动成功");
    } else {
        console.println("应用启动失败");
    }
    
    return success;
}

// 示例3：显示预览窗口
function showPreview(v) {
    console.println("显示预览窗口...");
    vdisplay.setTitle(v, "虚拟显示预览");
    vdisplay.showPreviewWindow(v, false);
    
    // 设置窗口大小和位置
    vdisplay.setPreviewWindowSize(v, 540, 960);
    vdisplay.setPreviewWindowPos(v, 50, 50);
    
    console.println("预览窗口已显示");
}

// 示例4：隐藏预览窗口
function hidePreview(v) {
    console.println("隐藏预览窗口...");
    vdisplay.hidePreviewWindow(v);
    console.println("预览窗口已隐藏");
}

// 示例5：设置触控回调
function setupTouchCallback(v) {
    console.println("设置触控回调...");
    
    vdisplay.setTouchCallback(v, function(x, y, action, displayId) {
        var actionText = "";
        if (action === 0) {
            actionText = "按下";
        } else if (action === 1) {
            actionText = "抬起";
        } else if (action === 2) {
            actionText = "移动";
        }
        
        console.println("触控事件:", "位置(" + x + ", " + y + ")", "动作:", actionText, "Display ID:", displayId);
    });
    
    console.println("触控回调已设置");
}

// 示例6：虚拟显示管理器
function VirtualDisplayManager() {
    this.display = null;
    
    this.create = function(width, height, dpi) {
        if (this.display) {
            console.println("虚拟显示设备已存在");
            return false;
        }
        
        this.display = vdisplay.create(width, height, dpi);
        console.println("虚拟显示设备已创建");
        return true;
    };
    
    this.destroy = function() {
        if (!this.display) {
            console.println("虚拟显示设备不存在");
            return false;
        }
        
        vdisplay.destroy(this.display);
        this.display = null;
        console.println("虚拟显示设备已销毁");
        return true;
    };
    
    this.launchApp = function(packageName) {
        if (!this.display) {
            console.println("虚拟显示设备不存在");
            return false;
        }
        
        return vdisplay.launchApp(this.display, packageName);
    };
    
    this.showPreview = function(rotated) {
        if (!this.display) {
            console.println("虚拟显示设备不存在");
            return false;
        }
        
        vdisplay.showPreviewWindow(this.display, rotated);
        return true;
    };
    
    this.hidePreview = function() {
        if (!this.display) {
            console.println("虚拟显示设备不存在");
            return false;
        }
        
        vdisplay.hidePreviewWindow(this.display);
        return true;
    };
    
    this.setPreviewSize = function(width, height) {
        if (!this.display) {
            console.println("虚拟显示设备不存在");
            return false;
        }
        
        vdisplay.setPreviewWindowSize(this.display, width, height);
        return true;
    };
    
    this.setPreviewPos = function(x, y) {
        if (!this.display) {
            console.println("虚拟显示设备不存在");
            return false;
        }
        
        vdisplay.setPreviewWindowPos(this.display, x, y);
        return true;
    };
    
    this.setTitle = function(title) {
        if (!this.display) {
            console.println("虚拟显示设备不存在");
            return false;
        }
        
        vdisplay.setTitle(this.display, title);
        return true;
    };
    
    this.setTouchCallback = function(callback) {
        if (!this.display) {
            console.println("虚拟显示设备不存在");
            return false;
        }
        
        vdisplay.setTouchCallback(this.display, callback);
        return true;
    };
    
    this.getDisplayId = function() {
        if (!this.display) {
            return -1;
        }
        
        return vdisplay.getDisplayId(this.display);
    };
}

// 使用虚拟显示管理器
function useVirtualDisplayManager() {
    var manager = new VirtualDisplayManager();
    
    // 创建虚拟显示设备
    manager.create(1080, 1920, 320);
    
    // 设置预览窗口
    manager.setTitle("我的虚拟显示");
    manager.setPreviewSize(540, 960);
    manager.setPreviewPos(100, 100);
    manager.showPreview(false);
    
    // 设置触控回调
    manager.setTouchCallback(function(x, y, action, displayId) {
        console.println("触控:", x, y, action, displayId);
    });
    
    // 启动应用
    manager.launchApp("com.android.settings");
    
    // 等待一段时间
    utils.sleep(10000);
    
    // 隐藏预览
    manager.hidePreview();
    
    // 销毁虚拟显示设备
    manager.destroy();
}

// 示例7：多窗口预览
function multiWindowPreview() {
    var v1 = vdisplay.create(1080, 1920, 320);
    var v2 = vdisplay.create(720, 1280, 320);
    
    // 设置第一个预览窗口
    vdisplay.setTitle(v1, "虚拟显示 1");
    vdisplay.setPreviewWindowSize(v1, 540, 960);
    vdisplay.setPreviewWindowPos(v1, 50, 50);
    vdisplay.showPreviewWindow(v1, false);
    
    // 设置第二个预览窗口
    vdisplay.setTitle(v2, "虚拟显示 2");
    vdisplay.setPreviewWindowSize(v2, 360, 640);
    vdisplay.setPreviewWindowPos(v2, 600, 50);
    vdisplay.showPreviewWindow(v2, false);
    
    console.println("两个预览窗口已显示");
    
    // 等待一段时间
    utils.sleep(10000);
    
    // 销毁虚拟显示设备
    vdisplay.destroy(v1);
    vdisplay.destroy(v2);
}

// 调用示例
var v = createVirtualDisplay();
launchAppInVirtualDisplay(v, "com.android.settings");
showPreview(v);
setupTouchCallback(v);
useVirtualDisplayManager();
multiWindowPreview();
```

## 注意事项

1. 创建虚拟显示设备需要相应的权限
2. 虚拟显示设备会占用系统资源，使用完毕后应销毁
3. 预览窗口的大小和位置应根据屏幕分辨率合理设置
4. 触控回调函数会在主线程中调用，避免执行耗时操作
5. 启动应用到虚拟显示前，确保应用已安装
6. 预览窗口可以旋转，rotated 参数控制旋转状态
7. 多个虚拟显示设备可以同时存在
8. 销毁虚拟显示设备会自动隐藏预览窗口
9. Display ID 用于区分不同的显示设备
10. 建议在创建虚拟显示设备前检查系统资源是否充足
