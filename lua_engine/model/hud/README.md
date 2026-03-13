# HUD 模块

HUD（Heads-Up Display）模块提供了在屏幕上显示浮动文本的功能，用于显示实时信息、调试数据等。

## 方法列表

### hud.new()
创建一个新的 HUD 对象。

**入参：** 无

**出参：** HUD 对象

**调用示例：**
```javascript
// 创建 HUD 对象
var hudObj = hud.new();
```

### hud.setPosition(hudObj, x1, y1, x2, y2)
设置 HUD 的显示位置和大小。

**入参：**
- `hudObj`: HUD 对象
- `x1`: 左上角 X 坐标（整数）
- `y1`: 左上角 Y 坐标（整数）
- `x2`: 右下角 X 坐标（整数）
- `y2`: 右下角 Y 坐标（整数）

**出参：** HUD 对象（支持链式调用）

**调用示例：**
```javascript
// 设置 HUD 位置和大小
hud.setPosition(hudObj, 100, 100, 500, 300);
```

### hud.setBackgroundColor(hudObj, color)
设置 HUD 的背景颜色。

**入参：**
- `hudObj`: HUD 对象
- `color`: 颜色字符串（如 "#80000000" 表示半透明黑色）

**出参：** HUD 对象（支持链式调用）

**调用示例：**
```javascript
// 设置背景颜色为半透明黑色
hud.setBackgroundColor(hudObj, "#80000000");
```

### hud.setTextSize(hudObj, size)
设置 HUD 中文本的字体大小。

**入参：**
- `hudObj`: HUD 对象
- `size`: 字体大小（整数）

**出参：** HUD 对象（支持链式调用）

**调用示例：**
```javascript
// 设置字体大小为 16
hud.setTextSize(hudObj, 16);
```

### hud.setText(hudObj, items)
设置 HUD 中要显示的文本内容。

**入参：**
- `hudObj`: HUD 对象
- `items`: 文本项数组，每个项包含 TextColor 和 Text 字段

**出参：** HUD 对象（支持链式调用）

**调用示例：**
```javascript
// 设置文本内容
hud.setText(hudObj, [
    { TextColor: "#FFFFFF", Text: "系统状态: 正常" },
    { TextColor: "#00FF00", Text: "CPU: 45%" },
    { TextColor: "#00FF00", Text: "内存: 60%" }
]);
```

### hud.show(hudObj)
显示 HUD。

**入参：**
- `hudObj`: HUD 对象

**出参：** 无

**调用示例：**
```javascript
// 显示 HUD
hud.show(hudObj);
```

### hud.hide(hudObj)
隐藏 HUD。

**入参：**
- `hudObj`: HUD 对象

**出参：** 无

**调用示例：**
```javascript
// 隐藏 HUD
hud.hide(hudObj);
```

### hud.isVisible(hudObj)
返回 HUD 是否可见。

**入参：**
- `hudObj`: HUD 对象

**出参：** 布尔值，true 表示可见，false 表示不可见

**调用示例：**
```javascript
// 检查 HUD 是否可见
var visible = hud.isVisible(hudObj);
console.println("HUD 可见:", visible);
```

### hud.destroy(hudObj)
销毁 HUD 对象，释放资源。

**入参：**
- `hudObj`: HUD 对象

**出参：** 无

**调用示例：**
```javascript
// 销毁 HUD 对象
hud.destroy(hudObj);
```

## 完整示例

```javascript
// 示例1：创建简单的状态显示 HUD
function createStatusHUD() {
    var hudObj = hud.new();
    hud.setPosition(hudObj, 50, 50, 400, 200);
    hud.setBackgroundColor(hudObj, "#80000000");
    hud.setTextSize(hudObj, 14);
    
    hud.setText(hudObj, [
        { TextColor: "#FFFFFF", Text: "系统状态" },
        { TextColor: "#00FF00", Text: "运行中..." }
    ]);
    
    hud.show(hudObj);
    return hudObj;
}

// 示例2：创建系统监控 HUD
function createMonitorHUD() {
    var hudObj = hud.new();
    hud.setPosition(hudObj, 10, 10, 300, 250);
    hud.setBackgroundColor(hudObj, "#90000000");
    hud.setTextSize(hudObj, 12);
    
    updateMonitorHUD(hudObj);
    hud.show(hudObj);
    return hudObj;
}

function updateMonitorHUD(hudObj) {
    var cpuUsage = device.getCpuUsage();
    var memUsage = device.getMemoryUsage();
    var battery = device.getBattery();
    
    hud.setText(hudObj, [
        { TextColor: "#FFFFFF", Text: "=== 系统监控 ===" },
        { TextColor: "#FFFF00", Text: "CPU: " + cpuUsage.toFixed(1) + "%" },
        { TextColor: "#FFFF00", Text: "内存: " + memUsage + " MB" },
        { TextColor: "#00FF00", Text: "电池: " + battery.level + "%" },
        { TextColor: "#00FFFF", Text: "时间: " + new Date().toLocaleTimeString() }
    ]);
}

// 示例3：使用链式调用创建 HUD
function createHUDWithChain() {
    var hudObj = hud.new()
        .setPosition(100, 100, 500, 300)
        .setBackgroundColor("#80000000")
        .setTextSize(16)
        .setText([
            { TextColor: "#FFFFFF", Text: "欢迎使用" },
            { TextColor: "#00FF00", Text: "HUD 显示系统" }
        ])
        .show();
    
    return hudObj;
}

// 示例4：动态更新 HUD 内容
function dynamicHUD() {
    var hudObj = hud.new();
    hud.setPosition(hudObj, 20, 20, 350, 200);
    hud.setBackgroundColor(hudObj, "#90000000");
    hud.setTextSize(hudObj, 14);
    hud.show(hudObj);
    
    var count = 0;
    var interval = setInterval(function() {
        count++;
        hud.setText(hudObj, [
            { TextColor: "#FFFFFF", Text: "计数器: " + count },
            { TextColor: "#00FF00", Text: "运行时间: " + count + " 秒" },
            { TextColor: "#FFFF00", Text: "状态: 正常" }
        ]);
        
        if (count >= 10) {
            clearInterval(interval);
            hud.destroy(hudObj);
        }
    }, 1000);
}

// 调用示例
var statusHUD = createStatusHUD();
var monitorHUD = createMonitorHUD();
var chainHUD = createHUDWithChain();
dynamicHUD();

// 定时更新监控 HUD
setInterval(function() {
    updateMonitorHUD(monitorHUD);
}, 1000);
```

## 注意事项

1. 使用完 HUD 后，建议调用 `destroy()` 方法释放资源
2. 颜色字符串格式为十六进制，支持透明度，如 "#AARRGGBB"
3. setText 方法可以接受多个文本项，每个项可以设置不同的颜色
4. HUD 的位置和大小应根据屏幕分辨率合理设置
5. 背景颜色建议使用半透明颜色，避免遮挡底层内容
6. 文本颜色和背景颜色应有足够的对比度，确保可读性
7. HUD 支持链式调用，可以连续调用多个方法
8. 动态更新 HUD 内容时，建议使用定时器定期刷新
