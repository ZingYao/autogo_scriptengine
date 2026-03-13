# AutoGo JavaScript Engine

AutoGo JavaScript Engine 是一个高性能的 JavaScript 脚本引擎，为 AutoGo 框架提供了完整的 JavaScript 脚本支持。它允许用户使用 JavaScript 语言编写自动化脚本，并提供了丰富的 API 接口。

## 特性

- **完整的 API 注入**: 将 AutoGo 的所有功能模块注入到 JavaScript 引擎中
- **方法管理**: 支持动态注册、移除、列出方法
- **方法重写**: 允许用 JavaScript 函数重写已注册的方法
- **线程安全**: 所有操作都是线程安全的
- **丰富的功能**: 支持应用管理、设备控制、触摸操作、文件操作、图像处理、文字识别等

## 快速开始

### 初始化引擎

```go
import "github.com/Dasongzi1366/AutoGo/js_engine"

func main() {
    engine := js_engine.GetEngine()
    defer js_engine.Close()
    
    // 使用引擎...
}
```

### 引擎配置选项

引擎支持配置是否自动注入所有方法，适用于需要按需加载模块或自定义注入的场景：

```go
import js_engine "github.com/ZingYao/autogo_scriptengine/js_engine"

func main() {
    // 方式1: 使用默认配置（自动注入所有方法）
    engine := js_engine.GetEngine()
    defer js_engine.Close()

    // 方式2: 使用自定义配置创建引擎
    config := &js_engine.EngineConfig{
        AutoInjectMethods: false, // 禁用自动注入
    }
    engine := js_engine.NewEngine(config)
    defer engine.Close()

    // 按需注入模块
    engine.InjectModule("app")
    engine.InjectModule("device")
    engine.InjectModule("motion")

    // 或注入多个模块
    engine.InjectModules([]string{"files", "images", "ppocr"})

    // 获取可用模块列表
    modules := engine.GetAvailableModules()
    // ["app", "device", "motion", "files", "images", "storages", "system", "http", "media", "opencv", "ppocr", "console", "dotocr", "hud", "ime", "plugin", "rhino", "uiacc", "utils", "vdisplay", "yolo", "imgui"]

    // 手动注入所有方法
    engine.InjectAllMethods()
}
```

### 模块列表

| 模块 | 说明 | 详细文档 |
|------|------|----------|
| `app` | 应用管理 | [README](model/app/README.md) |
| `device` | 设备信息 | [README](model/device/README.md) |
| `motion` | 触摸操作 | [README](model/motion/README.md) |
| `files` | 文件操作 | [README](model/files/README.md) |
| `images` | 图像处理 | [README](model/images/README.md) |
| `storages` | 数据存储 | [README](model/storages/README.md) |
| `system` | 系统功能 | [README](model/system/README.md) |
| `http` | 网络请求 | [README](model/http/README.md) |
| `media` | 媒体控制 | [README](model/media/README.md) |
| `opencv` | 计算机视觉 | [README](model/opencv/README.md) |
| `ppocr` | OCR 文字识别 | [README](model/ppocr/README.md) |
| `console` | 控制台窗口 | [README](model/console/README.md) |
| `dotocr` | 点字 OCR | [README](model/dotocr/README.md) |
| `hud` | HUD 悬浮显示 | [README](model/hud/README.md) |
| `ime` | 输入法控制 | [README](model/ime/README.md) |
| `plugin` | 插件加载 | [README](model/plugin/README.md) |
| `rhino` | JavaScript 执行引擎 | [README](model/rhino/README.md) |
| `uiacc` | 无障碍 UI 操作 | [README](model/uiacc/README.md) |
| `utils` | 工具方法 | [README](model/utils/README.md) |
| `vdisplay` | 虚拟显示 | [README](model/vdisplay/README.md) |
| `yolo` | YOLO 目标检测 | [README](model/yolo/README.md) |
| `coroutine` | 协程管理 | [README](model/coroutine/README.md) |
| `imgui` | Dear ImGui GUI 库 | [README](model/imgui/README.md) |

### 执行 JavaScript 代码

```go
// 执行 JavaScript 字符串
err := js_engine.ExecuteString(`
    const packageName = app.currentPackage();
    console.log("当前应用包名: " + packageName);
    
    click(500, 1000, 1);
`)

// 执行 JavaScript 文件
err = js_engine.ExecuteFile("/path/to/script.js")
```

## API 模块

### 基础函数

```javascript
// 延迟执行
sleep(1000); // 延迟 1000 毫秒
```

### 应用管理 (app)

```javascript
// 获取当前应用包名
const packageName = app.currentPackage();

// 获取当前应用类名
const activity = app.currentActivity();

// 启动应用
app.launch("com.example.app", 0);

// 打开应用设置
app.openAppSetting("com.example.app");

// 查看文件
app.viewFile("/sdcard/test.txt");

// 编辑文件
app.editFile("/sdcard/test.txt");

// 卸载应用
app.uninstall("com.example.app");

// 安装应用
app.install("/sdcard/app.apk");

// 检查应用是否已安装
if (app.isInstalled("com.example.app")) {
    console.log("应用已安装");
}

// 清除应用数据
app.clear("com.example.app");

// 强制停止应用
app.forceStop("com.example.app");

// 禁用应用
app.disable("com.example.app");

// 忽略电池优化
app.ignoreBattOpt("com.example.app");

// 启用应用
app.enable("com.example.app");

// 获取默认浏览器包名
const browser = app.getBrowserPackage();

// 打开 URL
app.openUrl("https://example.com");
```

### 设备管理 (device)

```javascript
// 获取设备信息
console.log("分辨率: " + device.width(0) + "x" + device.height(0));
console.log("SDK 版本: " + device.sdkInt);
console.log("CPU 架构: " + device.cpuAbi);

// 获取设备标识
const imei = device.getImei();
const androidId = device.getAndroidId();
const wifiMac = device.getWifiMac();
const wlanMac = device.getWlanMac();
const ip = device.getIp();

// 音量控制
const musicVolume = device.getMusicVolume();
device.setMusicVolume(50);

// 电池信息
const battery = device.getBattery();
const batteryStatus = device.getBatteryStatus();

// 屏幕控制
if (device.isScreenOn()) {
    console.log("屏幕已点亮");
}

if (device.isScreenUnlock()) {
    console.log("屏幕已解锁");
}

device.wakeUp();
device.keepScreenOn();

// 震动
device.vibrate(1000);
device.cancelVibration();
```

### 触摸操作 (touch)

```javascript
// 基本触摸操作
touchDown(500, 1000, 1, 0);
touchMove(600, 1100, 1, 0);
touchUp(600, 1100, 1, 0);

// 点击
click(500, 1000, 1, 0);

// 长按
longClick(500, 1000, 2000, 1, 0);

// 滑动
swipe(500, 1000, 600, 1100, 500, 1, 0);

// 系统按键
home(0);
back(0);
recents(0);
powerDialog();
notifications();
quickSettings();
volumeUp(0);
volumeDown(0);
camera();

// 自定义按键
keyAction(3, 0); // KEYCODE_HOME
```

### 文件操作 (files)

```javascript
// 检查文件/文件夹
if (files.isDir("/sdcard/test.txt")) {
    console.log("是文件");
}

if (files.isDir("/sdcard/Download")) {
    console.log("是文件夹");
}

// 创建文件/文件夹
files.create("/sdcard/test.txt");

// 检查文件是否存在
if (files.exists("/sdcard/test.txt")) {
    console.log("文件存在");
}

// 读写文件
const content = files.read("/sdcard/test.txt");
files.write("/sdcard/test.txt", "Hello, World!");
files.append("/sdcard/test.txt", "\nAppend text");

// 复制/移动/重命名/删除
files.copy("/sdcard/source.txt", "/sdcard/dest.txt");
files.move("/sdcard/source.txt", "/sdcard/newlocation.txt");
files.rename("/sdcard/old.txt", "/sdcard/new.txt");
files.remove("/sdcard/test.txt");

// 获取文件信息
const name = files.getName("/sdcard/test.txt");
const nameWithoutExt = files.getNameWithoutExtension("/sdcard/test.txt");
const ext = files.getExtension("/sdcard/test.txt");

// 获取绝对路径
const absPath = files.path("./test.txt");

// 列出目录
const fileList = files.listDir("/sdcard");
for (let i = 0; i < fileList.length; i++) {
    console.log(fileList[i]);
}
```

### 图像处理 (images)

```javascript
// 获取像素颜色
const color = images.pixel(500, 1000, 0);
console.log("颜色: " + color);

// 截取屏幕
const img = images.captureScreen(0, 0, 1080, 1920, 0);

// 比较颜色
if (images.cmpColor(500, 1000, "#FF0000", 0.9, 0)) {
    console.log("颜色匹配");
}

// 查找颜色
const [x, y] = images.findColor(0, 0, 1080, 1920, "#FF0000", 0.9, 0, 0);
if (x !== -1 && y !== -1) {
    console.log("找到颜色在: " + x + ", " + y);
    click(x, y, 1, 0);
}

// 获取颜色数量
const count = images.getColorCountInRegion(0, 0, 1080, 1920, "#FF0000", 0.9, 0);
console.log("颜色数量: " + count);

// 多点颜色检测
if (images.detectsMultiColors("0,0,#FF0000,10,10,#00FF00", 0.9, 0)) {
    console.log("多点颜色匹配");
}

// 查找多点颜色
const [x, y] = images.findMultiColors(0, 0, 1080, 1920, "0,0,#FF0000,10,10,#00FF00", 0.9, 0, 0);

// 读取图片
const img = images.readFromPath("/sdcard/image.png");
const img = images.readFromUrl("https://example.com/image.png");
const img = images.readFromBase64("iVBORw0KG...");
const img = images.readFromBytes(data);

// 保存图片
images.save(img, "/sdcard/output.png", 90);

// 编码图片
const base64 = images.encodeToBase64(img, "png", 90);
const bytes = images.encodeToBytes(img, "png", 90);

// 图像处理
const clipped = images.clip(img, 100, 100, 200, 200);
const resized = images.resize(img, 500, 500);
const rotated = images.rotate(img, 90);
const gray = images.grayscale(img);
const threshold = images.applyThreshold(img, 128, 255, "BINARY");
const adaptive = images.applyAdaptiveThreshold(img, 255, "GAUSSIAN_C", "BINARY", 11, 2);
const binary = images.applyBinarization(img, 128);
```

### 存储管理 (storages)

```javascript
// 存储键值对
storages.put("myTable", "key1", "value1");
storages.put("myTable", "key2", "value2");

// 读取键值
const value = storages.get("myTable", "key1");
console.log("key1 = " + value);

// 检查键是否存在
if (storages.contains("myTable", "key2")) {
    console.log("key2 存在");
}

// 获取所有键值对
const allData = storages.getAll("myTable");
for (const key in allData) {
    console.log(key + " = " + allData[key]);
}

// 删除键
storages.remove("myTable", "key1");

// 清空表
storages.clear("myTable");
```

### 系统管理 (system)

```javascript
// 获取进程 ID
const pid = system.getPid("myapp");

// 获取内存使用
const memory = system.getMemoryUsage(pid);

// 获取 CPU 使用率
const cpu = system.getCpuUsage(pid);

// 重启自身
system.restartSelf();

// 设置开机自启
system.setBootStart(true);
```

### 网络请求 (http)

```javascript
// GET 请求
const [code, data] = http.get("https://example.com", 5000);
console.log("状态码: " + code);
console.log("响应: " + data);

// POST Multipart 请求
const fileData = files.readBytes("/sdcard/image.png");
const [code, data] = http.postMultipart("https://example.com/upload", "image.png", fileData, 5000);
```

### 媒体管理 (media)

```javascript
// 扫描媒体文件
media.scanFile("/sdcard/image.png");
```

### 图像识别 (opencv)

```javascript
// 查找图片
const template = files.readBytes("/sdcard/template.png");
const [x, y] = opencv.findImage(0, 0, 1080, 1920, template, false, 1.0, 0.8, 0);
if (x !== -1 && y !== -1) {
    console.log("找到图片在: " + x + ", " + y);
    click(x, y, 1, 0);
}
```

### 文字识别 (ppocr)

```javascript
// 识别屏幕文字
const results = ppocr.ocr(0, 0, 1080, 1920, "", 0);
for (let i = 0; i < results.length; i++) {
    console.log("文本: " + results[i]["标签"]);
    console.log("位置: (" + results[i]["X"] + ", " + results[i]["Y"] + ")");
    console.log("大小: " + results[i]["宽"] + "x" + results[i]["高"]);
    console.log("精度: " + results[i]["精度"]);
    console.log("中心: (" + results[i]["CenterX"] + ", " + results[i]["CenterY"] + ")");
}

// 识别 Base64 图片
const results = ppocr.ocrFromBase64(base64Str, "");

// 识别文件图片
const results = ppocr.ocrFromPath("/sdcard/image.png", "");
```

### 协程管理 (coroutine)

```javascript
// ========== 基础协程方法 ==========

// 启动一个新的协程
const coroutineId = coroutine.launch(function() {
    console.log("协程开始执行");
    sleep(1000);
    console.log("协程执行完成");
}, "myCoroutine", 0);

// 延迟执行函数
const delayId = coroutine.delay(2000, function() {
    console.log("2秒后执行");
});

// 异步执行函数并返回结果（同步等待）
const asyncResult = coroutine.async(function() {
    return "异步结果";
});
console.log("异步结果: " + asyncResult);

// 在协程中睡眠
coroutine.sleep(1000);

// 等待普通值（简化版本，直接返回）
const value = coroutine.await("普通值");
console.log("等待结果: " + value);

// 注意：当前 await 实现为简化版本，直接返回传入的值
// 不进行真正的异步等待

// ========== 协程管理方法 ==========

// 取消指定的协程
const success = coroutine.cancel(coroutineId);

// 获取活跃的协程数量
const activeCount = coroutine.getActiveCoroutines();
console.log("活跃协程数量: " + activeCount);

// 获取协程列表
const coroutineList = coroutine.getCoroutineList();
for (let i = 0; i < coroutineList.length; i++) {
    console.log("协程ID: " + coroutineList[i].id);
    console.log("协程名称: " + coroutineList[i].name);
    console.log("协程状态: " + coroutineList[i].state);
    console.log("优先级: " + coroutineList[i].priority);
    console.log("运行时长: " + coroutineList[i].duration + "ms");
}

// 获取指定协程的详细信息
const info = coroutine.getCoroutineInfo(coroutineId);
console.log("协程ID: " + info.id);
console.log("协程名称: " + info.name);
console.log("协程状态: " + info.state);
console.log("开始时间: " + info.startedAt);
console.log("完成时间: " + info.completedAt);
console.log("运行时长: " + info.duration + "ms");

// 取消所有协程
const cancelledCount = coroutine.cancelAll();
console.log("已取消协程数量: " + cancelledCount);

// 获取全局统计信息
const stats = coroutine.getStats();
console.log("总任务数: " + stats.totalTasks);
console.log("已完成: " + stats.completed);
console.log("失败: " + stats.failed);
console.log("已取消: " + stats.cancelled);
console.log("活跃协程: " + stats.active);
console.log("协程池数量: " + stats.pools);

// ========== 协程池方法 ==========

// 创建协程池
const poolName = coroutine.createPool("myPool", 5, 100);
console.log("创建协程池: " + poolName);

// 提交任务到协程池
const submitSuccess = coroutine.submitToPool("myPool", function() {
    console.log("协程池任务执行");
    return "任务完成";
}, 0);

// 获取协程池统计信息
const poolStats = coroutine.getPoolStats("myPool");
console.log("协程池名称: " + poolStats.name);
console.log("最大工作协程数: " + poolStats.maxWorkers);
console.log("最大任务数: " + poolStats.maxTasks);
console.log("活跃工作协程数: " + poolStats.active);
console.log("队列中任务数: " + poolStats.queued);
console.log("工作协程总数: " + poolStats.workers);
console.log("是否已关闭: " + poolStats.closed);

// 关闭协程池
const closeSuccess = coroutine.closePool("myPool");
console.log("关闭协程池: " + closeSuccess);

// 列出所有协程池
const poolList = coroutine.listPools();
for (let i = 0; i < poolList.length; i++) {
    console.log("协程池: " + poolList[i]);
}

// ========== 调度器方法 ==========

// 设置调度策略 (fifo 或 priority)
coroutine.setScheduleStrategy("fifo");

// 获取当前调度策略
const strategy = coroutine.getScheduleStrategy();
console.log("当前调度策略: " + strategy);

// 设置协程优先级
coroutine.setPriority("myCoroutine", 10);

// 获取协程优先级
const priority = coroutine.getPriority("myCoroutine");
console.log("协程优先级: " + priority);

// ========== 协程使用示例 ==========

// 示例1: 并发执行多个任务
function concurrentTasks() {
    console.log("开始并发执行任务");
    
    const task1 = coroutine.launch(function() {
        console.log("任务1开始");
        coroutine.sleep(1000);
        console.log("任务1完成");
        return "任务1结果";
    }, "task1", 1);
    
    const task2 = coroutine.launch(function() {
        console.log("任务2开始");
        coroutine.sleep(1500);
        console.log("任务2完成");
        return "任务2结果";
    }, "task2", 2);
    
    const task3 = coroutine.launch(function() {
        console.log("任务3开始");
        coroutine.sleep(2000);
        console.log("任务3完成");
        return "任务3结果";
    }, "task3", 3);
    
    console.log("所有任务已启动");
}

// 示例2: 使用协程池处理批量任务
function batchProcessingWithPool() {
    console.log("创建协程池处理批量任务");
    
    const poolName = coroutine.createPool("batchPool", 3, 10);
    
    for (let i = 0; i < 10; i++) {
        const taskNum = i + 1;
        coroutine.submitToPool("batchPool", function() {
            console.log("处理任务 " + taskNum);
            coroutine.sleep(500);
            console.log("任务 " + taskNum + " 完成");
        }, taskNum);
    }
    
    console.log("所有任务已提交到协程池");
    
    // 等待任务完成
    coroutine.sleep(5000);
    
    // 查看协程池状态
    const poolStats = coroutine.getPoolStats("batchPool");
    console.log("协程池状态: " + JSON.stringify(poolStats));
    
    // 关闭协程池
    coroutine.closePool("batchPool");
}

// 示例3: 延迟执行和定时任务
function delayedAndScheduledTasks() {
    console.log("设置延迟任务");
    
    // 延迟1秒执行
    coroutine.delay(1000, function() {
        console.log("1秒后执行的任务");
    });
    
    // 延迟2秒执行
    coroutine.delay(2000, function() {
        console.log("2秒后执行的任务");
    });
    
    // 延迟3秒执行
    coroutine.delay(3000, function() {
        console.log("3秒后执行的任务");
    });
    
    console.log("所有延迟任务已设置");
}

// 示例4: 协程状态监控
function monitorCoroutines() {
    console.log("启动协程监控");
    
    // 启动多个协程
    for (let i = 0; i < 5; i++) {
        const taskNum = i + 1;
        coroutine.launch(function() {
            console.log("协程 " + taskNum + " 开始执行");
            coroutine.sleep(1000 * taskNum);
            console.log("协程 " + taskNum + " 执行完成");
        }, "monitorTask" + taskNum, taskNum);
    }
    
    // 定期检查协程状态
    setInterval(function() {
        const activeCount = coroutine.getActiveCoroutines();
        const stats = coroutine.getStats();
        
        console.log("=== 协程状态 ===");
        console.log("活跃协程数: " + activeCount);
        console.log("总任务数: " + stats.totalTasks);
        console.log("已完成: " + stats.completed);
        console.log("失败: " + stats.failed);
        console.log("已取消: " + stats.cancelled);
        
        if (activeCount === 0) {
            console.log("所有协程已完成");
        }
    }, 500);
}
```

## 方法管理

### 注册新方法

```javascript
registerMethod("myMethod", "我的自定义方法", null, true);

function myMethod(param) {
    console.log("自定义方法被调用: " + param);
    return "返回值";
}
```

### 移除方法

```javascript
const success = unregisterMethod("myMethod");
```

### 列出所有方法

```javascript
const methods = listMethods();
for (let i = 0; i < methods.length; i++) {
    console.log(methods[i].name + " - " + methods[i].description);
    console.log("可重写: " + methods[i].overridable);
    console.log("已重写: " + methods[i].overridden);
}
```

### 重写方法

```javascript
// 方法 1: 直接重写
const originalClick = click;

function click(x, y, fingerID) {
    console.log("点击: (" + x + ", " + y + ")");
    originalClick(x, y, fingerID);
}

// 方法 2: 使用 overrideMethod
overrideMethod("click", function(x, y, fingerID) {
    console.log("点击: (" + x + ", " + y + ")");
    // 调用原始实现
});
```

### 恢复方法

```javascript
const success = restoreMethod("click");
```

## 生成文档

### 生成 JavaScript 文档

```go
docGen := js_engine.NewDocumentationGenerator()
err := docGen.SaveJSDocumentation("js_api.js")
```

### 生成 Markdown 文档

```go
docGen := js_engine.NewDocumentationGenerator()
err := docGen.SaveMarkdownDocumentation("js_api.md")
```

## 完整示例

### 将脚本嵌入到程序并运行

以下是一个完整的 Demo，展示如何将 JavaScript 脚本嵌入到 Go 程序中，打包到产物中，运行时释放并执行：

#### 1. 使用 Go embed 嵌入脚本文件（Go 1.16+）

```go
package main

import (
    "embed"
    "fmt"
    "os"
    "path/filepath"
    
    "github.com/Dasongzi1366/AutoGo/files"
    "github.com/Dasongzi1366/AutoGo/js_engine"
)

//go:embed scripts/*
var scriptFS embed.FS

// EmbeddedScriptManager 嵌入式脚本管理器
type EmbeddedScriptManager struct {
    scriptDir  string
    extracted bool
}

// NewEmbeddedScriptManager 创建嵌入式脚本管理器
func NewEmbeddedScriptManager(scriptDir string) *EmbeddedScriptManager {
    return &EmbeddedScriptManager{
        scriptDir:  scriptDir,
        extracted: false,
    }
}

// ExtractScripts 提取嵌入的脚本到文件系统
func (esm *EmbeddedScriptManager) ExtractScripts() error {
    if esm.extracted {
        return nil
    }
    
    // 确保目标目录存在
    if !files.Exists(esm.scriptDir) {
        if err := files.Create(esm.scriptDir); err != nil {
            return fmt.Errorf("创建目录失败: %v", err)
        }
    }
    
    // 读取嵌入的脚本目录
    entries, err := scriptFS.ReadDir("scripts")
    if err != nil {
        return fmt.Errorf("读取嵌入目录失败: %v", err)
    }
    
    // 提取所有脚本文件
    for _, entry := range entries {
        if entry.IsDir() {
            continue
        }
        
        srcPath := filepath.Join("scripts", entry.Name())
        dstPath := filepath.Join(esm.scriptDir, entry.Name())
        
        // 读取嵌入的文件内容
        content, err := scriptFS.ReadFile(srcPath)
        if err != nil {
            return fmt.Errorf("读取文件失败 %s: %v", srcPath, err)
        }
        
        // 写入到文件系统
        if err := files.Write(dstPath, string(content)); err != nil {
            return fmt.Errorf("写入文件失败 %s: %v", dstPath, err)
        }
        
        fmt.Printf("已提取脚本: %s\n", dstPath)
    }
    
    esm.extracted = true
    return nil
}

// ListScripts 列出所有嵌入的脚本
func (esm *EmbeddedScriptManager) ListScripts() ([]string, error) {
    entries, err := scriptFS.ReadDir("scripts")
    if err != nil {
        return nil, fmt.Errorf("读取嵌入目录失败: %v", err)
    }
    
    var scripts []string
    for _, entry := range entries {
        if !entry.IsDir() {
            scripts = append(scripts, entry.Name())
        }
    }
    
    return scripts, nil
}

// GetScriptContent 获取脚本内容（不释放到文件系统）
func (esm *EmbeddedScriptManager) GetScriptContent(name string) (string, error) {
    srcPath := filepath.Join("scripts", name)
    content, err := scriptFS.ReadFile(srcPath)
    if err != nil {
        return "", fmt.Errorf("读取脚本失败 %s: %v", srcPath, err)
    }
    
    return string(content), nil
}

// RunScript 运行脚本（从嵌入的文件系统直接执行）
func (esm *EmbeddedScriptManager) RunScript(name string) error {
    content, err := esm.GetScriptContent(name)
    if err != nil {
        return err
    }
    
    if err := js_engine.ExecuteString(content); err != nil {
        return fmt.Errorf("执行脚本失败: %v", err)
    }
    
    return nil
}

// RunScriptFromFile 运行脚本（从释放的文件执行）
func (esm *EmbeddedScriptManager) RunScriptFromFile(name string) error {
    if !esm.extracted {
        if err := esm.ExtractScripts(); err != nil {
            return err
        }
    }
    
    scriptPath := filepath.Join(esm.scriptDir, name)
    if err := js_engine.ExecuteFile(scriptPath); err != nil {
        return fmt.Errorf("执行脚本失败: %v", err)
    }
    
    return nil
}

func main() {
    // 初始化 JavaScript 引擎
    engine := js_engine.GetEngine()
    defer js_engine.Close()
    
    // 创建嵌入式脚本管理器
    scriptDir := "/sdcard/AutoGo/scripts"
    scriptManager := NewEmbeddedScriptManager(scriptDir)
    
    // 列出所有嵌入的脚本
    fmt.Println("嵌入的脚本列表:")
    scripts, err := scriptManager.ListScripts()
    if err != nil {
        fmt.Printf("列出脚本失败: %v\n", err)
        return
    }
    
    for i, script := range scripts {
        fmt.Printf("  %d. %s\n", i+1, script)
    }
    
    // 方法1: 直接从嵌入的文件系统执行脚本（不释放到文件系统）
    fmt.Println("\n方法1: 直接从嵌入的文件系统执行脚本")
    if err := scriptManager.RunScript("demo.js"); err != nil {
        fmt.Printf("执行脚本失败: %v\n", err)
    }
    
    // 方法2: 提取脚本到文件系统后执行
    fmt.Println("\n方法2: 提取脚本到文件系统后执行")
    if err := scriptManager.ExtractScripts(); err != nil {
        fmt.Printf("提取脚本失败: %v\n", err)
        return
    }
    
    if err := scriptManager.RunScriptFromFile("demo.js"); err != nil {
        fmt.Printf("执行脚本失败: %v\n", err)
    }
}
```

#### 2. 脚本文件结构

```
project/
├── main.go
└── scripts/
    ├── demo.js
    ├── auto_login.js
    ├── find_color.js
    └── ocr_text.js
```

#### 3. 示例脚本文件

**scripts/demo.js**
```javascript
// Demo 脚本
console.log("=== Demo 脚本开始执行 ===");

// 获取设备信息
const width = device.width(0);
const height = device.height(0);
console.log("屏幕分辨率: " + width + "x" + height);

// 获取当前应用
const packageName = app.currentPackage();
console.log("当前应用: " + packageName);

// 点击屏幕中心
click(width/2, height/2, 1, 0);

console.log("=== Demo 脚本执行完成 ===");
return true;
```

**scripts/auto_login.js**
```javascript
// 自动登录脚本
function autoLogin(username, password) {
    console.log("开始自动登录...");
    
    // 检查当前应用
    const currentApp = app.currentPackage();
    if (currentApp !== "com.example.app") {
        console.log("启动应用...");
        app.launch("com.example.app", 0);
        sleep(3000);
    }
    
    // 查找用户名输入框
    const [x, y] = images.findColor(0, 0, device.width(0), device.height(0), "#FF0000", 0.9, 0, 0);
    if (x !== -1) {
        click(x, y, 1, 0);
        sleep(500);
    }
    
    // 查找密码输入框
    const [x, y] = images.findColor(0, 0, device.width(0), device.height(0), "#00FF00", 0.9, 0, 0);
    if (x !== -1) {
        click(x, y, 1, 0);
        sleep(500);
    }
    
    // 查找登录按钮
    const [x, y] = images.findColor(0, 0, device.width(0), device.height(0), "#0000FF", 0.9, 0, 0);
    if (x !== -1) {
        click(x, y, 1, 0);
        sleep(2000);
    }
    
    console.log("登录完成");
    return true;
}

return autoLogin("user123", "pass456");
```

**scripts/find_color.js**
```javascript
// 查找颜色并点击
function findAndClick(color, sim) {
    const [x, y] = images.findColor(0, 0, device.width(0), device.height(0), color, sim, 0, 0);
    if (x !== -1 && y !== -1) {
        click(x, y, 1, 0);
        console.log("找到颜色 " + color + " 在: (" + x + ", " + y + ")");
        return true;
    }
    console.log("未找到颜色: " + color);
    return false;
}

return findAndClick("#FF0000", 0.9);
```

**scripts/ocr_text.js**
```javascript
// OCR 文字识别并点击
function findTextAndClick(text) {
    const results = ppocr.ocr(0, 0, device.width(0), device.height(0), "", 0);
    for (let i = 0; i < results.length; i++) {
        if (results[i]["标签"].includes(text)) {
            click(results[i]["CenterX"], results[i]["CenterY"], 1, 0);
            console.log("找到文字 '" + text + "' 在: (" + results[i]["CenterX"] + ", " + results[i]["CenterY"] + ")");
            return true;
        }
    }
    console.log("未找到文字: " + text);
    return false;
}

return findTextAndClick("确定");
```

#### 4. 使用 go-bindata 嵌入脚本文件（兼容旧版本）

```bash
# 安装 go-bindata
go get -u github.com/go-bindata/go-bindata/...

# 生成嵌入的文件
go-bindata -o scripts.go -pkg main scripts/
```

```go
package main

import (
    "fmt"
    "path/filepath"
    
    "github.com/Dasongzi1366/AutoGo/files"
    "github.com/Dasongzi1366/AutoGo/js_engine"
)

// BindataScriptManager 使用 go-bindata 的脚本管理器
type BindataScriptManager struct {
    scriptDir  string
    extracted bool
}

// NewBindataScriptManager 创建脚本管理器
func NewBindataScriptManager(scriptDir string) *BindataScriptManager {
    return &BindataScriptManager{
        scriptDir:  scriptDir,
        extracted: false,
    }
}

// ExtractScripts 提取嵌入的脚本到文件系统
func (bsm *BindataScriptManager) ExtractScripts() error {
    if bsm.extracted {
        return nil
    }
    
    // 确保目标目录存在
    if !files.Exists(bsm.scriptDir) {
        if err := files.Create(bsm.scriptDir); err != nil {
            return fmt.Errorf("创建目录失败: %v", err)
        }
    }
    
    // 遍历所有嵌入的文件
    for _, name := range AssetNames() {
        if filepath.Ext(name) !== ".js" {
            continue
        }
        
        // 读取嵌入的文件内容
        content, err := Asset(name)
        if err != nil {
            return fmt.Errorf("读取文件失败 %s: %v", name, err)
        }
        
        // 写入到文件系统
        dstPath := filepath.Join(bsm.scriptDir, filepath.Base(name));
        if err := files.Write(dstPath, string(content)); err != nil {
            return fmt.Errorf("写入文件失败 %s: %v", dstPath, err)
        }
        
        fmt.Printf("已提取脚本: %s\n", dstPath);
    }
    
    bsm.extracted = true
    return nil
}

// ListScripts 列出所有嵌入的脚本
func (bsm *BindataScriptManager) ListScripts() ([]string, error) {
    var scripts []string
    for _, name := range AssetNames() {
        if filepath.Ext(name) === ".js" {
            scripts = append(scripts, filepath.Base(name));
        }
    }
    return scripts, nil
}

// GetScriptContent 获取脚本内容
func (bsm *BindataScriptManager) GetScriptContent(name string) (string, error) {
    content, err := Asset(name)
    if err != nil {
        return "", fmt.Errorf("读取脚本失败 %s: %v", name, err)
    }
    return string(content), nil
}

// RunScript 运行脚本
func (bsm *BindataScriptManager) RunScript(name string) error {
    content, err := bsm.GetScriptContent(name)
    if err != nil {
        return err
    }
    
    if err := js_engine.ExecuteString(content); err != nil {
        return fmt.Errorf("执行脚本失败: %v", err)
    }
    
    return nil
}

func main() {
    // 初始化 JavaScript 引擎
    engine := js_engine.GetEngine()
    defer js_engine.Close()
    
    // 创建脚本管理器
    scriptDir := "/sdcard/AutoGo/scripts"
    scriptManager := NewBindataScriptManager(scriptDir)
    
    // 列出所有嵌入的脚本
    fmt.Println("嵌入的脚本列表:")
    scripts, err := scriptManager.ListScripts()
    if err != nil {
        fmt.Printf("列出脚本失败: %v\n", err)
        return
    }
    
    for (let i = 0; i < scripts.length; i++) {
        fmt.Printf("  %d. %s\n", i+1, scripts[i]);
    }
    
    // 提取并运行脚本
    if err := scriptManager.ExtractScripts(); err != nil {
        fmt.Printf("提取脚本失败: %v\n", err)
        return
    }
    
    if err := scriptManager.RunScript("demo.js"); err != nil {
        fmt.Printf("执行脚本失败: %v\n", err)
    }
}
```

## 与 Lua 引擎的对比

### 相同点

1. **API 接口一致**: JavaScript 引擎提供了与 Lua 引擎完全相同的 API 接口
2. **功能模块相同**: 支持相同的功能模块（app、device、touch、files、images、storages、system、http、media、opencv、ppocr）
3. **方法管理机制**: 都支持方法注册、移除、列出、重写和恢复
4. **线程安全**: 所有操作都是线程安全的
5. **文档生成**: 都支持生成 JavaScript/Lua 和 Markdown 格式的文档

### 不同点

1. **脚本语言**: JavaScript vs Lua
2. **语法差异**: JavaScript 使用驼峰命名和对象属性访问，Lua 使用下划线命名和函数调用
3. **数据结构**: JavaScript 使用数组和对象，Lua 使用表（table）
4. **函数调用**: JavaScript 使用对象方法调用（如 `app.currentPackage()`），Lua 使用全局函数（如 `app_currentPackage()`）
5. **协程支持**: Lua 引擎支持协程管理，JavaScript 引擎暂不支持
6. **引擎库**: Lua 使用 gopher-lua，JavaScript 使用 goja

### API 调用对比

| 功能 | Lua | JavaScript |
|------|-----|------------|
| 获取当前应用包名 | `app_currentPackage()` | `app.currentPackage()` |
| 点击屏幕 | `click(x, y, fingerID, displayId)` | `click(x, y, fingerID, displayId)` |
| 读取文件 | `files.read(path)` | `files.read(path)` |
| 查找颜色 | `images_findColor(...)` | `images.findColor(...)` |
| 存储键值 | `storages_put(table, key, value)` | `storages.put(table, key, value)` |

### 性能对比

- **执行速度**: JavaScript (goja) 通常比 Lua (gopher-lua) 更快
- **内存占用**: JavaScript 引擎的内存占用略高于 Lua 引擎
- **启动时间**: 两者启动时间相近

### 适用场景

- **Lua 引擎**: 适合需要协程支持、对性能要求不极高的场景
- **JavaScript 引擎**: 适合需要更高性能、更广泛开发者基础的场景

## 自定义 Go 方法注入

除了使用内置的模块，你还可以注入自己实现的 Go 方法到 JavaScript 引擎中。

### 方式1: 通过 RegisterMethod 注册方法

```go
package main

import (
    "fmt"
    js_engine "github.com/ZingYao/autogo_scriptengine/js_engine"
)

func main() {
    // 创建引擎（不自动注入方法）
    config := &js_engine.EngineConfig{
        AutoInjectMethods: false,
    }
    engine := js_engine.NewEngine(config)
    defer engine.Close()

    // 注册自定义 Go 函数
    engine.RegisterMethod("myGreet", "打招呼", func(name string) string {
        return "Hello, " + name + "!"
    }, true)

    // 注册带多参数的函数
    engine.RegisterMethod("myAdd", "加法运算", func(a, b int) int {
        return a + b
    }, true)

    // 注册返回多个值的函数（通过返回 map 或对象）
    engine.RegisterMethod("myGetInfo", "获取信息", func() map[string]interface{} {
        return map[string]interface{}{
            "name":  "AutoGo",
            "version": "1.0.0",
            "status": "running",
        }
    }, true)

    // 在 JavaScript 中调用
    err := engine.ExecuteString(`
        // 调用自定义方法
        var greeting = myGreet("World");
        console.log(greeting); // Hello, World!

        var sum = myAdd(10, 20);
        console.log("Sum: " + sum); // Sum: 30

        var info = myGetInfo();
        console.log("Name: " + info.name);
        console.log("Version: " + info.version);
    `)
    if err != nil {
        fmt.Printf("执行错误: %v\n", err)
    }
}
```

### 方式2: 直接通过 VM 设置方法

```go
package main

import (
    "fmt"
    js_engine "github.com/ZingYao/autogo_scriptengine/js_engine"
    "github.com/dop251/goja"
)

func main() {
    engine := js_engine.NewEngine(nil) // 使用默认配置
    defer engine.Close()

    vm := engine.GetVM()

    // 设置全局函数
    vm.Set("myCustomFunc", func(call goja.FunctionCall) goja.Value {
        arg := call.Argument(0).String()
        result := "Processed: " + arg
        return vm.ToValue(result)
    })

    // 设置对象方法
    myModule := vm.NewObject()
    myModule.Set("method1", func(x, y int) int {
        return x * y
    })
    myModule.Set("method2", func(s string) string {
        return "Echo: " + s
    })
    vm.Set("myModule", myModule)

    // 在 JavaScript 中调用
    err := engine.ExecuteString(`
        var result1 = myCustomFunc("test");
        console.log(result1); // Processed: test

        var result2 = myModule.method1(5, 6);
        console.log(result2); // 30

        var result3 = myModule.method2("hello");
        console.log(result3); // Echo: hello
    `)
    if err != nil {
        fmt.Printf("执行错误: %v\n", err)
    }
}
```

### 方式3: 创建自定义注入模块

创建一个新的注入文件 `custom_inject.go`:

```go
package js_engine

import (
    "github.com/dop251/goja"
    "your-project/your-module"
)

// injectCustomMethods 注入自定义方法
func injectCustomMethods(e *JSEngine) {
    vm := e.vm

    // 创建自定义模块对象
    customObj := vm.NewObject()

    // 注入方法1
    customObj.Set("doSomething", func(call goja.FunctionCall) goja.Value {
        param := call.Argument(0).String()
        result := yourmodule.DoSomething(param)
        return vm.ToValue(result)
    })

    // 注入方法2（带多个参数）
    customObj.Set("processData", func(call goja.FunctionCall) goja.Value {
        data := call.Argument(0).Export()
        options := call.Argument(1).ToObject(vm)

        // 处理数据...
        result := yourmodule.Process(data, options)

        return vm.ToValue(result)
    })

    // 注入异步方法
    customObj.Set("asyncOperation", func(call goja.FunctionCall) goja.Value {
        callback := call.Argument(0)

        go func() {
            // 执行异步操作
            result := yourmodule.AsyncOperation()

            // 回调到 JavaScript
            vm.RunString(fmt.Sprintf(`(%s)(%s)`, callback, result))
        }()

        return goja.Undefined()
    })

    // 注册到全局
    vm.Set("custom", customObj)
}
```

然后在引擎初始化后调用:

```go
func main() {
    config := &js_engine.EngineConfig{
        AutoInjectMethods: false,
    }
    engine := js_engine.NewEngine(config)
    defer engine.Close()

    // 注入自定义模块
    injectCustomMethods(engine)

    // 执行脚本
    engine.ExecuteString(`
        var result = custom.doSomething("test");
        console.log(result);

        custom.processData({key: "value"}, {option: true});
    `)
}
```

### 参数类型转换

在 Go 和 JavaScript 之间传递数据时，需要注意类型转换：

```go
// 从 JavaScript 获取参数
vm.Set("processArgs", func(call goja.FunctionCall) goja.Value {
    // 字符串
    str := call.Argument(0).String()

    // 数字
    num := call.Argument(1).ToFloat()
    intNum := call.Argument(1).ToInteger()

    // 布尔值
    boolVal := call.Argument(2).ToBoolean()

    // 对象 -> Go map
    obj := call.Argument(3).ToObject(vm)
    objMap := obj.Export().(map[string]interface{})

    // 数组 -> Go slice
    arr := call.Argument(4).Export().([]interface{})

    // 返回值
    return vm.ToValue(map[string]interface{}{
        "processed": true,
        "str":       str,
        "num":       num,
    })
})
```

### 完整示例：注入数据库操作模块

```go
package main

import (
    "database/sql"
    "fmt"
    js_engine "github.com/ZingYao/autogo_scriptengine/js_engine"
    _ "github.com/mattn/go-sqlite3"
)

func injectDatabaseModule(e *JSEngine, db *sql.DB) {
    vm := e.GetVM()

    dbObj := vm.NewObject()

    // 查询方法
    dbObj.Set("query", func(sqlQuery string) []map[string]interface{} {
        rows, err := db.Query(sqlQuery)
        if err != nil {
            return nil
        }
        defer rows.Close()

        columns, _ := rows.Columns()
        var results []map[string]interface{}

        for rows.Next() {
            values := make([]interface{}, len(columns))
            valuePtrs := make([]interface{}, len(columns))
            for i := range values {
                valuePtrs[i] = &values[i]
            }

            rows.Scan(valuePtrs...)

            row := make(map[string]interface{})
            for i, col := range columns {
                row[col] = values[i]
            }
            results = append(results, row)
        }

        return results
    })

    // 执行方法
    dbObj.Set("exec", func(sqlStmt string) (int64, error) {
        result, err := db.Exec(sqlStmt)
        if err != nil {
            return 0, err
        }
        return result.RowsAffected()
    })

    vm.Set("db", dbObj)
}

func main() {
    // 打开数据库
    db, err := sql.Open("sqlite3", "/sdcard/mydb.sqlite")
    if err != nil {
        panic(err)
    }
    defer db.Close()

    // 创建引擎
    engine := js_engine.NewEngine(nil)
    defer engine.Close()

    // 注入数据库模块
    injectDatabaseModule(engine, db)

    // 使用
    engine.ExecuteString(`
        // 创建表
        var affected = db.exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT)");
        console.log("Affected: " + affected);

        // 插入数据
        db.exec("INSERT INTO users (name) VALUES ('Alice')");

        // 查询数据
        var users = db.query("SELECT * FROM users");
        for (var i = 0; i < users.length; i++) {
            console.log("User: " + users[i].name);
        }
    `)
}
```

## 总结

AutoGo JavaScript Engine 提供了与 Lua Engine 相同的功能和 API，但使用 JavaScript 作为脚本语言。开发者可以根据自己的喜好和项目需求选择合适的脚本引擎。两者都提供了完整的文档生成和嵌入式脚本管理功能，方便开发者集成和使用。
