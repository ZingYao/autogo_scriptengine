# AutoGo Lua Engine

AutoGo Lua Engine 是一个高性能的 Lua 脚本引擎，为 AutoGo 框架提供了完整的 Lua 脚本支持。它允许用户使用 Lua 语言编写自动化脚本，并提供了丰富的 API 接口。

## 特性

- **完整的 API 注入**: 将 AutoGo 的所有功能模块注入到 Lua 引擎中
- **方法管理**: 支持动态注册、移除、列出方法
- **方法重写**: 允许用 Lua 函数重写已注册的方法
- **线程安全**: 所有操作都是线程安全的
- **丰富的功能**: 支持应用管理、设备控制、触摸操作、文件操作、图像处理、文字识别等

## 快速开始

### 初始化引擎

```go
import "github.com/Dasongzi1366/AutoGo/lua_engine"

func main( then
    engine := lua_engine.GetEngine()
    defer lua_engine.Close()
    
    -- 使用引擎...
end
```

### 引擎配置选项

引擎支持配置是否自动注入所有方法，适用于需要按需加载模块或自定义注入的场景：

```go
import lua_engine "github.com/ZingYao/autogo_scriptengine/lua_engine"

func main( then
    -- 方式1: 使用默认配置（自动注入所有方法）
    engine := lua_engine.GetEngine()
    defer lua_engine.Close()

    -- 方式2: 使用自定义配置创建引擎
    config := &lua_engine.EngineConfig{
        AutoInjectMethods: false, -- 禁用自动注入
    end
    engine := lua_engine.NewEngine(config)
    defer engine.Close()

    -- 按需注入模块
    engine.InjectModule("app")
    engine.InjectModule("device")
    engine.InjectModule("motion")

    -- 或注入多个模块
    engine.InjectModules([]string{"files", "images", "ppocr"end)

    -- 获取可用模块列表
    modules := engine.GetAvailableModules()
    -- ["app", "device", "motion", "files", "images", "storages", "system", "http", "media", "opencv", "ppocr", "console", "dotocr", "hud", "ime", "plugin", "rhino", "uiacc", "utils", "vdisplay", "yolo", "imgui"]

    -- 手动注入所有方法
    engine.InjectAllMethods()
end
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
| `coroutine` | 协程管理 | [README](model/coroutine/README.md) |
| `dotocr` | 点字 OCR | [README](model/dotocr/README.md) |
| `hud` | HUD 悬浮显示 | [README](model/hud/README.md) |
| `ime` | 输入法控制 | [README](model/ime/README.md) |
| `plugin` | 插件加载 | [README](model/plugin/README.md) |
| `rhino` | JavaScript 执行引擎 | [README](model/rhino/README.md) |
| `uiacc` | 无障碍 UI 操作 | [README](model/uiacc/README.md) |
| `utils` | 工具方法 | [README](model/utils/README.md) |
| `vdisplay` | 虚拟显示 | [README](model/vdisplay/README.md) |
| `yolo` | YOLO 目标检测 | [README](model/yolo/README.md) |
| `imgui` | Dear ImGui GUI 库 | [README](model/imgui/README.md) |

### 执行 Lua 代码

```go
-- 执行 Lua 字符串
err := lua_engine.ExecuteString(`
    local packageName = app_currentPackage()
    console.log("当前应用包名: " .. packageName)
    
    click(500, 1000, 1)
`)

-- 执行 Lua 文件
err = lua_engine.ExecuteFile("/path/to/script.lua")
```

## API 模块

### 基础函数

```lua
-- 延迟执行
sleep(1000) -- 延迟 1000 毫秒
```

### 应用管理 (app)

```lua
-- 获取当前应用包名
local packageName = app_currentPackage()

-- 获取当前应用类名
local activity = app_currentActivity()

-- 启动应用
app_launch("com.example.app", 0)

-- 打开应用设置
app_openAppSetting("com.example.app")

-- 查看文件
app_viewFile("/sdcard/test.txt")

-- 编辑文件
app_editFile("/sdcard/test.txt")

-- 卸载应用
app_uninstall("com.example.app")

-- 安装应用
app_install("/sdcard/app_apk")

-- 检查应用是否已安装
if app_isInstalled("com.example.app") then
    console.log("应用已安装")
end

-- 清除应用数据
app_clear("com.example.app")

-- 强制停止应用
app_forceStop("com.example.app") 

-- 禁用应用
app_disable("com.example.app") 

-- 忽略电池优化
app_ignoreBattOpt("com.example.app") 

-- 启用应用
app_enable("com.example.app") 

-- 获取默认浏览器包名
local browser = app_getBrowserPackage() 

-- 打开 URL
app_openUrl("https://example.com") 
```

### 设备管理 (device)

```lua
-- 获取设备信息
console.log("分辨率: " + device_width(0) + "x" + device_height(0)) 
console.log("SDK 版本: " + device_sdkInt) 
console.log("CPU 架构: " + device_cpuAbi) 

-- 获取设备标识
local imei = device_getImei() 
local androidId = device_getAndroidId() 
local wifiMac = device_getWifiMac() 
local wlanMac = device_getWlanMac() 
local ip = device_getIp() 

-- 音量控制
local musicVolume = device_getMusicVolume() 
device_setMusicVolume(50) 

-- 电池信息
local battery = device_getBattery() 
local batteryStatus = device_getBatteryStatus() 

-- 屏幕控制
if (device_isScreenOn() then
    console.log("屏幕已点亮") 
end

if (device_isScreenUnlock() then
    console.log("屏幕已解锁") 
end

device_wakeUp() 
device_keepScreenOn() 

-- 震动
device_vibrate(1000) 
device_cancelVibration() 
```

### 触摸操作 (touch)

```lua
-- 基本触摸操作
touchDown(500, 1000, 1, 0) 
touchMove(600, 1100, 1, 0) 
touchUp(600, 1100, 1, 0) 

-- 点击
click(500, 1000, 1, 0) 

-- 长按
longClick(500, 1000, 2000, 1, 0) 

-- 滑动
swipe(500, 1000, 600, 1100, 500, 1, 0) 

-- 系统按键
home(0) 
back(0) 
recents(0) 
powerDialog() 
notifications() 
quickSettings() 
volumeUp(0) 
volumeDown(0) 
camera() 

-- 自定义按键
keyAction(3, 0)  -- KEYCODE_HOME
```

### 文件操作 (files)

```lua
-- 检查文件/文件夹
if (files_isDir("/sdcard/test.txt") then
    console.log("是文件") 
end

if (files_isDir("/sdcard/Download") then
    console.log("是文件夹") 
end

-- 创建文件/文件夹
files_create("/sdcard/test.txt") 

-- 检查文件是否存在
if (files_exists("/sdcard/test.txt") then
    console.log("文件存在") 
end

-- 读写文件
local content = files_read("/sdcard/test.txt") 
files_write("/sdcard/test.txt", "Hello, World!") 
files_append("/sdcard/test.txt", "\nAppend text") 

-- 复制/移动/重命名/删除
files_copy("/sdcard/source.txt", "/sdcard/dest.txt") 
files_move("/sdcard/source.txt", "/sdcard/newlocation.txt") 
files_rename("/sdcard/old.txt", "/sdcard/new.txt") 
files_remove("/sdcard/test.txt") 

-- 获取文件信息
local name = files_getName("/sdcard/test.txt") 
local nameWithoutExt = files_getNameWithoutExtension("/sdcard/test.txt") 
local ext = files_getExtension("/sdcard/test.txt") 

-- 获取绝对路径
local absPath = files_path("./test.txt") 

-- 列出目录
local fileList = files_listDir("/sdcard") 
for (local i = 0  i < fileList.length  i++ then
    console.log(fileList[i]) 
end
```

### 协程管理 (coroutine)

```lua
-- ========== 基础协程方法 ==========

-- 启动一个新的协程
local coroutineId = coroutine_launch(function()
    console.log("协程开始执行")
    sleep(1000)
    console.log("协程执行完成")
end, "myCoroutine", 0)

-- 延迟执行函数
local delayId = coroutine_delay(2000, function()
    console.log("2秒后执行")
end)

-- 异步执行函数并返回结果（同步等待）
local asyncResult = coroutine_async(function()
    return "异步结果"
end)
console.log("异步结果: " .. asyncResult)

-- 在协程中睡眠
coroutine_sleep(1000)

-- 等待普通值（简化版本，直接返回）
local value = coroutine_await("普通值")
console.log("等待结果: " .. value)

-- 注意：当前 await 实现为简化版本，直接返回传入的值
-- 不进行真正的异步等待

-- ========== 协程管理方法 ==========

-- 取消指定的协程
local success = coroutine_cancel(coroutineId)

-- 获取活跃的协程数量
local activeCount = coroutine_getActiveCoroutines()
console.log("活跃协程数量: " .. activeCount)

-- 获取协程列表
local coroutineList = coroutine_getCoroutineList()
for i, coro in ipairs(coroutineList) do
    console.log("协程ID: " .. coro.id)
    console.log("协程名称: " .. coro.name)
    console.log("协程状态: " .. coro.state)
    console.log("优先级: " .. coro.priority)
    console.log("运行时长: " .. coro.duration .. "ms")
end

-- 获取指定协程的详细信息
local info = coroutine_getCoroutineInfo(coroutineId)
console.log("协程ID: " .. info.id)
console.log("协程名称: " .. info.name)
console.log("协程状态: " .. info.state)
console.log("开始时间: " .. info.startedAt)
console.log("完成时间: " .. info.completedAt)
console.log("运行时长: " .. info.duration .. "ms")

-- 取消所有协程
local cancelledCount = coroutine_cancelAll()
console.log("已取消协程数量: " .. cancelledCount)

-- 获取全局统计信息
local stats = coroutine_getStats()
console.log("总任务数: " .. stats.totalTasks)
console.log("已完成: " .. stats.completed)
console.log("失败: " .. stats.failed)
console.log("已取消: " .. stats.cancelled)
console.log("活跃协程: " .. stats.active)
console.log("协程池数量: " .. stats.pools)

-- ========== 协程池方法 ==========

-- 创建协程池
local poolName = coroutine_createPool("myPool", 5, 100)
console.log("创建协程池: " .. poolName)

-- 提交任务到协程池
local submitSuccess = coroutine_submitToPool("myPool", function()
    console.log("协程池任务执行")
    return "任务完成"
end, 0)

-- 获取协程池统计信息
local poolStats = coroutine_getPoolStats("myPool")
console.log("协程池名称: " .. poolStats.name)
console.log("最大工作协程数: " .. poolStats.maxWorkers)
console.log("最大任务数: " .. poolStats.maxTasks)
console.log("活跃工作协程数: " .. poolStats.active)
console.log("队列中任务数: " .. poolStats.queued)
console.log("工作协程总数: " .. poolStats.workers)
console.log("是否已关闭: " .. tostring(poolStats.closed))

-- 关闭协程池
local closeSuccess = coroutine_closePool("myPool")
console.log("关闭协程池: " .. tostring(closeSuccess))

-- ========== 调度器方法 ==========

-- 设置调度策略
coroutine_setScheduleStrategy("fifo")

-- 获取当前调度策略
local strategy = coroutine_getScheduleStrategy()
console.log("当前调度策略: " .. strategy)

-- 设置协程优先级
coroutine_setPriority("myCoroutine", 10)

-- 获取协程优先级
local priority = coroutine_getPriority("myCoroutine")
console.log("协程优先级: " .. priority)
```

### 图像处理 (images)

```lua
-- 获取像素颜色
local color = images_pixel(500, 1000, 0) 
console.log("颜色: " + color) 

-- 截取屏幕
local img = images_captureScreen(0, 0, 1080, 1920, 0) 

-- 比较颜色
if (images_cmpColor(500, 1000, "#FF0000", 0.9, 0) then
    console.log("颜色匹配") 
end

-- 查找颜色
local [x, y] = images_findColor(0, 0, 1080, 1920, "#FF0000", 0.9, 0, 0) 
if (x !== -1 && y !== -1 then
    console.log("找到颜色在: " + x + ", " + y) 
    click(x, y, 1, 0) 
end

-- 获取颜色数量
local count = images_getColorCountInRegion(0, 0, 1080, 1920, "#FF0000", 0.9, 0) 
console.log("颜色数量: " + count) 

-- 多点颜色检测
if (images_detectsMultiColors("0,0,#FF0000,10,10,#00FF00", 0.9, 0) then
    console.log("多点颜色匹配") 
end

-- 查找多点颜色
local [x, y] = images_findMultiColors(0, 0, 1080, 1920, "0,0,#FF0000,10,10,#00FF00", 0.9, 0, 0) 

-- 读取图片
local img = images_readFromPath("/sdcard/image.png") 
local img = images_readFromUrl("https://example.com/image.png") 
local img = images_readFromBase64("iVBORw0KG...") 
local img = images_readFromBytes(data) 

-- 保存图片
images_save(img, "/sdcard/output.png", 90) 

-- 编码图片
local base64 = images_encodeToBase64(img, "png", 90) 
local bytes = images_encodeToBytes(img, "png", 90) 

-- 图像处理
local clipped = images_clip(img, 100, 100, 200, 200) 
local resized = images_resize(img, 500, 500) 
local rotated = images_rotate(img, 90) 
local gray = images_grayscale(img) 
local threshold = images_applyThreshold(img, 128, 255, "BINARY") 
local adaptive = images_applyAdaptiveThreshold(img, 255, "GAUSSIAN_C", "BINARY", 11, 2) 
local binary = images_applyBinarization(img, 128) 
```

### 存储管理 (storages)

```lua
-- 存储键值对
storages_put("myTable", "key1", "value1") 
storages_put("myTable", "key2", "value2") 

-- 读取键值
local value = storages_get("myTable", "key1") 
console.log("key1 = " + value) 

-- 检查键是否存在
if (storages_contains("myTable", "key2") then
    console.log("key2 存在") 
end

-- 获取所有键值对
local allData = storages_getAll("myTable") 
for (local key in allData then
    console.log(key + " = " + allData[key]) 
end

-- 删除键
storages_remove("myTable", "key1") 

-- 清空表
storages_clear("myTable") 
```

### 系统管理 (system)

```lua
-- 获取进程 ID
local pid = system_getPid("myapp") 

-- 获取内存使用
local memory = system_getMemoryUsage(pid) 

-- 获取 CPU 使用率
local cpu = system_getCpuUsage(pid) 

-- 重启自身
system_restartSelf() 

-- 设置开机自启
system_setBootStart(true) 
```

### 网络请求 (http)

```lua
-- GET 请求
local [code, data] = http_get("https://example.com", 5000) 
console.log("状态码: " + code) 
console.log("响应: " + data) 

-- POST Multipart 请求
local fileData = files_readBytes("/sdcard/image.png") 
local [code, data] = http_postMultipart("https://example.com/upload", "image.png", fileData, 5000) 
```

### 媒体管理 (media)

```lua
-- 扫描媒体文件
media_scanFile("/sdcard/image.png") 
```

### 图像识别 (opencv)

```lua
-- 查找图片
local template = files_readBytes("/sdcard/template.png") 
local [x, y] = opencv_findImage(0, 0, 1080, 1920, template, false, 1.0, 0.8, 0) 
if (x !== -1 && y !== -1 then
    console.log("找到图片在: " + x + ", " + y) 
    click(x, y, 1, 0) 
end
```

### 文字识别 (ppocr)

```lua
-- 识别屏幕文字
local results = ppocr_ocr(0, 0, 1080, 1920, "", 0) 
for (local i = 0  i < results.length  i++ then
    console.log("文本: " + results[i]["标签"]) 
    console.log("位置: (" + results[i]["X"] + ", " + results[i]["Y"] + ")") 
    console.log("大小: " + results[i]["宽"] + "x" + results[i]["高"]) 
    console.log("精度: " + results[i]["精度"]) 
    console.log("中心: (" + results[i]["CenterX"] + ", " + results[i]["CenterY"] + ")") 
end

-- 识别 Base64 图片
local results = ppocr_ocrFromBase64(base64Str, "") 

-- 识别文件图片
local results = ppocr_ocrFromPath("/sdcard/image.png", "") 
```

### 协程管理 (coroutine)

```lua
-- ========== 基础协程方法 ==========

-- 启动一个新的协程
local coroutineId = coroutine_launch(function ( then
    console.log("协程开始执行") 
    sleep(1000) 
    console.log("协程执行完成") 
end, "myCoroutine", 0) 

-- 延迟执行函数
local delayId = coroutine_delay(2000, function ( then
    console.log("2秒后执行") 
end) 

-- 异步执行函数并返回结果（同步等待）
local asyncResult = coroutine_async(function ( then
    return "异步结果" 
end) 
console.log("异步结果: " + asyncResult) 

-- 在协程中睡眠
coroutine_sleep(1000) 

-- 等待普通值（简化版本，直接返回）
local value = coroutine_await("普通值") 
console.log("等待结果: " + value) 

-- 注意：当前 await 实现为简化版本，直接返回传入的值
-- 不进行真正的异步等待

-- ========== 协程管理方法 ==========

-- 取消指定的协程
local success = coroutine_cancel(coroutineId) 

-- 获取活跃的协程数量
local activeCount = coroutine_getActiveCoroutines() 
console.log("活跃协程数量: " + activeCount) 

-- 获取协程列表
local coroutineList = coroutine_getCoroutineList() 
for (local i = 0  i < coroutineList.length  i++ then
    console.log("协程ID: " + coroutineList[i].id) 
    console.log("协程名称: " + coroutineList[i].name) 
    console.log("协程状态: " + coroutineList[i].state) 
    console.log("优先级: " + coroutineList[i].priority) 
    console.log("运行时长: " + coroutineList[i].duration + "ms") 
end

-- 获取指定协程的详细信息
local info = coroutine_getCoroutineInfo(coroutineId) 
console.log("协程ID: " + info.id) 
console.log("协程名称: " + info.name) 
console.log("协程状态: " + info.state) 
console.log("开始时间: " + info.startedAt) 
console.log("完成时间: " + info.completedAt) 
console.log("运行时长: " + info.duration + "ms") 

-- 取消所有协程
local cancelledCount = coroutine_cancelAll() 
console.log("已取消协程数量: " + cancelledCount) 

-- 获取全局统计信息
local stats = coroutine_getStats() 
console.log("总任务数: " + stats.totalTasks) 
console.log("已完成: " + stats.completed) 
console.log("失败: " + stats.failed) 
console.log("已取消: " + stats.cancelled) 
console.log("活跃协程: " + stats.active) 
console.log("协程池数量: " + stats.pools) 

-- ========== 协程池方法 ==========

-- 创建协程池
local poolName = coroutine_createPool("myPool", 5, 100) 
console.log("创建协程池: " + poolName) 

-- 提交任务到协程池
local submitSuccess = coroutine_submitToPool("myPool", function ( then
    console.log("协程池任务执行") 
    return "任务完成" 
end, 0) 

-- 获取协程池统计信息
local poolStats = coroutine_getPoolStats("myPool") 
console.log("协程池名称: " + poolStats.name) 
console.log("最大工作协程数: " + poolStats.maxWorkers) 
console.log("最大任务数: " + poolStats.maxTasks) 
console.log("活跃工作协程数: " + poolStats.active) 
console.log("队列中任务数: " + poolStats.queued) 
console.log("工作协程总数: " + poolStats.workers) 
console.log("是否已关闭: " + poolStats.closed) 

-- 关闭协程池
local closeSuccess = coroutine_closePool("myPool") 
console.log("关闭协程池: " + closeSuccess) 

-- 列出所有协程池
local poolList = coroutine_listPools() 
for (local i = 0  i < poolList.length  i++ then
    console.log("协程池: " + poolList[i]) 
end

-- ========== 调度器方法 ==========

-- 设置调度策略 (fifo 或 priority)
coroutine_setScheduleStrategy("fifo") 

-- 获取当前调度策略
local strategy = coroutine_getScheduleStrategy() 
console.log("当前调度策略: " + strategy) 

-- 设置协程优先级
coroutine_setPriority("myCoroutine", 10) 

-- 获取协程优先级
local priority = coroutine_getPriority("myCoroutine") 
console.log("协程优先级: " + priority) 

-- ========== 协程使用示例 ==========

-- 示例1: 并发执行多个任务
function concurrentTasks( then
    console.log("开始并发执行任务") 
    
    local task1 = coroutine_launch(function ( then
        console.log("任务1开始") 
        coroutine_sleep(1000) 
        console.log("任务1完成") 
        return "任务1结果" 
    end, "task1", 1) 
    
    local task2 = coroutine_launch(function ( then
        console.log("任务2开始") 
        coroutine_sleep(1500) 
        console.log("任务2完成") 
        return "任务2结果" 
    end, "task2", 2) 
    
    local task3 = coroutine_launch(function ( then
        console.log("任务3开始") 
        coroutine_sleep(2000) 
        console.log("任务3完成") 
        return "任务3结果" 
    end, "task3", 3) 
    
    console.log("所有任务已启动") 
end

-- 示例2: 使用协程池处理批量任务
function batchProcessingWithPool( then
    console.log("创建协程池处理批量任务") 
    
    local poolName = coroutine_createPool("batchPool", 3, 10) 
    
    for (local i = 0  i < 10  i++ then
        local taskNum = i + 1 
        coroutine_submitToPool("batchPool", function ( then
            console.log("处理任务 " + taskNum) 
            coroutine_sleep(500) 
            console.log("任务 " + taskNum + " 完成") 
        end, taskNum) 
    end
    
    console.log("所有任务已提交到协程池") 
    
    -- 等待任务完成
    coroutine_sleep(5000) 
    
    -- 查看协程池状态
    local poolStats = coroutine_getPoolStats("batchPool") 
    console.log("协程池状态: " + JSON.stringify(poolStats)) 
    
    -- 关闭协程池
    coroutine_closePool("batchPool") 
end

-- 示例3: 延迟执行和定时任务
function delayedAndScheduledTasks( then
    console.log("设置延迟任务") 
    
    -- 延迟1秒执行
    coroutine_delay(1000, function ( then
        console.log("1秒后执行的任务") 
    end) 
    
    -- 延迟2秒执行
    coroutine_delay(2000, function ( then
        console.log("2秒后执行的任务") 
    end) 
    
    -- 延迟3秒执行
    coroutine_delay(3000, function ( then
        console.log("3秒后执行的任务") 
    end) 
    
    console.log("所有延迟任务已设置") 
end

-- 示例4: 协程状态监控
function monitorCoroutines( then
    console.log("启动协程监控") 
    
    -- 启动多个协程
    for (local i = 0  i < 5  i++ then
        local taskNum = i + 1 
        coroutine_launch(function ( then
            console.log("协程 " + taskNum + " 开始执行") 
            coroutine_sleep(1000 * taskNum) 
            console.log("协程 " + taskNum + " 执行完成") 
        end, "monitorTask" + taskNum, taskNum) 
    end
    
    -- 定期检查协程状态
    setInterval(function ( then
        local activeCount = coroutine_getActiveCoroutines() 
        local stats = coroutine_getStats() 
        
        console.log("=== 协程状态 ===") 
        console.log("活跃协程数: " + activeCount) 
        console.log("总任务数: " + stats.totalTasks) 
        console.log("已完成: " + stats.completed) 
        console.log("失败: " + stats.failed) 
        console.log("已取消: " + stats.cancelled) 
        
        if (activeCount === 0 then
            console.log("所有协程已完成") 
        end
    end, 500) 
end
```

## 方法管理

### 注册新方法

```lua
registerMethod("myMethod", "我的自定义方法", null, true) 

function myMethod(param then
    console.log("自定义方法被调用: " + param) 
    return "返回值" 
end
```

### 移除方法

```lua
local success = unregisterMethod("myMethod") 
```

### 列出所有方法

```lua
local methods = listMethods() 
for (local i = 0  i < methods.length  i++ then
    console.log(methods[i].name + " - " + methods[i].description) 
    console.log("可重写: " + methods[i].overridable) 
    console.log("已重写: " + methods[i].overridden) 
end
```

### 重写方法

```lua
-- 方法 1: 直接重写
local originalClick = click 

function click(x, y, fingerID then
    console.log("点击: (" + x + ", " + y + ")") 
    originalClick(x, y, fingerID) 
end

-- 方法 2: 使用 overrideMethod
overrideMethod("click", function (x, y, fingerID then
    console.log("点击: (" + x + ", " + y + ")") 
    -- 调用原始实现
end) 
```

### 恢复方法

```lua
local success = restoreMethod("click") 
```

## 生成文档

### 生成 JavaScript 文档

```go
docGen := lua_engine.NewDocumentationGenerator()
err := docGen.SaveJSDocumentation("js_api.lua")
```

### 生成 Markdown 文档

```go
docGen := lua_engine.NewDocumentationGenerator()
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
    "github.com/Dasongzi1366/AutoGo/lua_engine"
)

//go:embed scripts/*
var scriptFS embed.FS

-- EmbeddedScriptManager 嵌入式脚本管理器
type EmbeddedScriptManager struct {
    scriptDir  string
    extracted bool
end

-- NewEmbeddedScriptManager 创建嵌入式脚本管理器
func NewEmbeddedScriptManager(scriptDir string) *EmbeddedScriptManager {
    return &EmbeddedScriptManager{
        scriptDir:  scriptDir,
        extracted: false,
    end
end

-- ExtractScripts 提取嵌入的脚本到文件系统
func (esm *EmbeddedScriptManager) ExtractScripts() error {
    if esm.extracted {
        return nil
    end
    
    -- 确保目标目录存在
    if !files_Exists(esm.scriptDir then
        if err := files_Create(esm.scriptDir)  err != nil {
            return fmt.Errorf("创建目录失败: %v", err)
        end
    end
    
    -- 读取嵌入的脚本目录
    entries, err := scriptFS.ReadDir("scripts")
    if err != nil {
        return fmt.Errorf("读取嵌入目录失败: %v", err)
    end
    
    -- 提取所有脚本文件
    for _, entry := range entries {
        if entry.IsDir( then
            continue
        end
        
        srcPath := filepath.Join("scripts", entry.Name())
        dstPath := filepath.Join(esm.scriptDir, entry.Name())
        
        -- 读取嵌入的文件内容
        content, err := scriptFS.ReadFile(srcPath)
        if err != nil {
            return fmt.Errorf("读取文件失败 %s: %v", srcPath, err)
        end
        
        -- 写入到文件系统
        if err := files_Write(dstPath, string(content))  err != nil {
            return fmt.Errorf("写入文件失败 %s: %v", dstPath, err)
        end
        
        fmt.Printf("已提取脚本: %s\n", dstPath)
    end
    
    esm.extracted = true
    return nil
end

-- ListScripts 列出所有嵌入的脚本
func (esm *EmbeddedScriptManager) ListScripts() ([]string, error then
    entries, err := scriptFS.ReadDir("scripts")
    if err != nil {
        return nil, fmt.Errorf("读取嵌入目录失败: %v", err)
    end
    
    var scripts []string
    for _, entry := range entries {
        if !entry.IsDir( then
            scripts = append(scripts, entry.Name())
        end
    end
    
    return scripts, nil
end

-- GetScriptContent 获取脚本内容（不释放到文件系统）
func (esm *EmbeddedScriptManager) GetScriptContent(name string) (string, error then
    srcPath := filepath.Join("scripts", name)
    content, err := scriptFS.ReadFile(srcPath)
    if err != nil {
        return "", fmt.Errorf("读取脚本失败 %s: %v", srcPath, err)
    end
    
    return string(content), nil
end

-- RunScript 运行脚本（从嵌入的文件系统直接执行）
func (esm *EmbeddedScriptManager) RunScript(name string) error {
    content, err := esm.GetScriptContent(name)
    if err != nil {
        return err
    end
    
    if err := lua_engine.ExecuteString(content)  err != nil {
        return fmt.Errorf("执行脚本失败: %v", err)
    end
    
    return nil
end

-- RunScriptFromFile 运行脚本（从释放的文件执行）
func (esm *EmbeddedScriptManager) RunScriptFromFile(name string) error {
    if !esm.extracted {
        if err := esm.ExtractScripts()  err != nil {
            return err
        end
    end
    
    scriptPath := filepath.Join(esm.scriptDir, name)
    if err := lua_engine.ExecuteFile(scriptPath)  err != nil {
        return fmt.Errorf("执行脚本失败: %v", err)
    end
    
    return nil
end

func main( then
    -- 初始化 JavaScript 引擎
    engine := lua_engine.GetEngine()
    defer lua_engine.Close()
    
    -- 创建嵌入式脚本管理器
    scriptDir := "/sdcard/AutoGo/scripts"
    scriptManager := NewEmbeddedScriptManager(scriptDir)
    
    -- 列出所有嵌入的脚本
    fmt.Println("嵌入的脚本列表:")
    scripts, err := scriptManager.ListScripts()
    if err != nil {
        fmt.Printf("列出脚本失败: %v\n", err)
        return
    end
    
    for i, script := range scripts {
        fmt.Printf("  %d. %s\n", i+1, script)
    end
    
    -- 方法1: 直接从嵌入的文件系统执行脚本（不释放到文件系统）
    fmt.Println("\n方法1: 直接从嵌入的文件系统执行脚本")
    if err := scriptManager.RunScript("demo.lua")  err != nil {
        fmt.Printf("执行脚本失败: %v\n", err)
    end
    
    -- 方法2: 提取脚本到文件系统后执行
    fmt.Println("\n方法2: 提取脚本到文件系统后执行")
    if err := scriptManager.ExtractScripts()  err != nil {
        fmt.Printf("提取脚本失败: %v\n", err)
        return
    end
    
    if err := scriptManager.RunScriptFromFile("demo.lua")  err != nil {
        fmt.Printf("执行脚本失败: %v\n", err)
    end
end
```

#### 2. 脚本文件结构

```
project/
├── main.go
└── scripts/
    ├── demo.lua
    ├── auto_login.lua
    ├── find_color.lua
    └── ocr_text.lua
```

#### 3. 示例脚本文件

**scripts/demo.lua**
```lua
-- Demo 脚本
console.log("=== Demo 脚本开始执行 ===") 

-- 获取设备信息
local width = device_width(0) 
local height = device_height(0) 
console.log("屏幕分辨率: " + width + "x" + height) 

-- 获取当前应用
local packageName = app_currentPackage() 
console.log("当前应用: " + packageName) 

-- 点击屏幕中心
click(width/2, height/2, 1, 0) 

console.log("=== Demo 脚本执行完成 ===") 
return true 
```

**scripts/auto_login.lua**
```lua
-- 自动登录脚本
function autoLogin(username, password then
    console.log("开始自动登录...") 
    
    -- 检查当前应用
    local currentApp = app_currentPackage() 
    if (currentApp !== "com.example.app" then
        console.log("启动应用...") 
        app_launch("com.example.app", 0) 
        sleep(3000) 
    end
    
    -- 查找用户名输入框
    local [x, y] = images_findColor(0, 0, device_width(0), device_height(0), "#FF0000", 0.9, 0, 0) 
    if (x !== -1 then
        click(x, y, 1, 0) 
        sleep(500) 
    end
    
    -- 查找密码输入框
    local [x, y] = images_findColor(0, 0, device_width(0), device_height(0), "#00FF00", 0.9, 0, 0) 
    if (x !== -1 then
        click(x, y, 1, 0) 
        sleep(500) 
    end
    
    -- 查找登录按钮
    local [x, y] = images_findColor(0, 0, device_width(0), device_height(0), "#0000FF", 0.9, 0, 0) 
    if (x !== -1 then
        click(x, y, 1, 0) 
        sleep(2000) 
    end
    
    console.log("登录完成") 
    return true 
end

return autoLogin("user123", "pass456") 
```

**scripts/find_color.lua**
```lua
-- 查找颜色并点击
function findAndClick(color, sim then
    local [x, y] = images_findColor(0, 0, device_width(0), device_height(0), color, sim, 0, 0) 
    if (x !== -1 && y !== -1 then
        click(x, y, 1, 0) 
        console.log("找到颜色 " + color + " 在: (" + x + ", " + y + ")") 
        return true 
    end
    console.log("未找到颜色: " + color) 
    return false 
end

return findAndClick("#FF0000", 0.9) 
```

**scripts/ocr_text.lua**
```lua
-- OCR 文字识别并点击
function findTextAndClick(text then
    local results = ppocr_ocr(0, 0, device_width(0), device_height(0), "", 0) 
    for (local i = 0  i < results.length  i++ then
        if (results[i]["标签"].includes(text) then
            click(results[i]["CenterX"], results[i]["CenterY"], 1, 0) 
            console.log("找到文字 '" + text + "' 在: (" + results[i]["CenterX"] + ", " + results[i]["CenterY"] + ")") 
            return true 
        end
    end
    console.log("未找到文字: " + text) 
    return false 
end

return findTextAndClick("确定") 
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
    "github.com/Dasongzi1366/AutoGo/lua_engine"
)

-- BindataScriptManager 使用 go-bindata 的脚本管理器
type BindataScriptManager struct {
    scriptDir  string
    extracted bool
end

-- NewBindataScriptManager 创建脚本管理器
func NewBindataScriptManager(scriptDir string) *BindataScriptManager {
    return &BindataScriptManager{
        scriptDir:  scriptDir,
        extracted: false,
    end
end

-- ExtractScripts 提取嵌入的脚本到文件系统
func (bsm *BindataScriptManager) ExtractScripts() error {
    if bsm.extracted {
        return nil
    end
    
    -- 确保目标目录存在
    if !files_Exists(bsm.scriptDir then
        if err := files_Create(bsm.scriptDir)  err != nil {
            return fmt.Errorf("创建目录失败: %v", err)
        end
    end
    
    -- 遍历所有嵌入的文件
    for _, name := range AssetNames( then
        if filepath.Ext(name) !== ".lua" {
            continue
        end
        
        -- 读取嵌入的文件内容
        content, err := Asset(name)
        if err != nil {
            return fmt.Errorf("读取文件失败 %s: %v", name, err)
        end
        
        -- 写入到文件系统
        dstPath := filepath.Join(bsm.scriptDir, filepath.Base(name)) 
        if err := files_Write(dstPath, string(content))  err != nil {
            return fmt.Errorf("写入文件失败 %s: %v", dstPath, err)
        end
        
        fmt.Printf("已提取脚本: %s\n", dstPath) 
    end
    
    bsm.extracted = true
    return nil
end

-- ListScripts 列出所有嵌入的脚本
func (bsm *BindataScriptManager) ListScripts() ([]string, error then
    var scripts []string
    for _, name := range AssetNames( then
        if filepath.Ext(name) === ".lua" {
            scripts = append(scripts, filepath.Base(name)) 
        end
    end
    return scripts, nil
end

-- GetScriptContent 获取脚本内容
func (bsm *BindataScriptManager) GetScriptContent(name string) (string, error then
    content, err := Asset(name)
    if err != nil {
        return "", fmt.Errorf("读取脚本失败 %s: %v", name, err)
    end
    return string(content), nil
end

-- RunScript 运行脚本
func (bsm *BindataScriptManager) RunScript(name string) error {
    content, err := bsm.GetScriptContent(name)
    if err != nil {
        return err
    end
    
    if err := lua_engine.ExecuteString(content)  err != nil {
        return fmt.Errorf("执行脚本失败: %v", err)
    end
    
    return nil
end

func main( then
    -- 初始化 JavaScript 引擎
    engine := lua_engine.GetEngine()
    defer lua_engine.Close()
    
    -- 创建脚本管理器
    scriptDir := "/sdcard/AutoGo/scripts"
    scriptManager := NewBindataScriptManager(scriptDir)
    
    -- 列出所有嵌入的脚本
    fmt.Println("嵌入的脚本列表:")
    scripts, err := scriptManager.ListScripts()
    if err != nil {
        fmt.Printf("列出脚本失败: %v\n", err)
        return
    end
    
    for (local i = 0  i < scripts.length  i++ then
        fmt.Printf("  %d. %s\n", i+1, scripts[i]) 
    end
    
    -- 提取并运行脚本
    if err := scriptManager.ExtractScripts()  err != nil {
        fmt.Printf("提取脚本失败: %v\n", err)
        return
    end
    
    if err := scriptManager.RunScript("demo.lua")  err != nil {
        fmt.Printf("执行脚本失败: %v\n", err)
    end
end
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
4. **函数调用**: JavaScript 使用对象方法调用（如 `app_currentPackage()`），Lua 使用全局函数（如 `app_currentPackage()`）
5. **协程支持**: Lua 引擎支持协程管理，JavaScript 引擎暂不支持
6. **引擎库**: Lua 使用 gopher-lua，JavaScript 使用 goja

### API 调用对比

| 功能 | Lua | JavaScript |
|------|-----|------------|
| 获取当前应用包名 | `app_currentPackage()` | `app_currentPackage()` |
| 点击屏幕 | `click(x, y, fingerID, displayId)` | `click(x, y, fingerID, displayId)` |
| 读取文件 | `files_read(path)` | `files_read(path)` |
| 查找颜色 | `images_findColor(...)` | `images_findColor(...)` |
| 存储键值 | `storages_put(table, key, value)` | `storages_put(table, key, value)` |

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
    lua_engine "github.com/ZingYao/autogo_scriptengine/lua_engine"
)

func main( then
    -- 创建引擎（不自动注入方法）
    config := &lua_engine.EngineConfig{
        AutoInjectMethods: false,
    end
    engine := lua_engine.NewEngine(config)
    defer engine.Close()

    -- 注册自定义 Go 函数
    engine.RegisterMethod("myGreet", "打招呼", func(name string) string {
        return "Hello, " + name + "!"
    end, true)

    -- 注册带多参数的函数
    engine.RegisterMethod("myAdd", "加法运算", func(a, b int) int {
        return a + b
    end, true)

    -- 注册返回多个值的函数（通过返回 map 或对象）
    engine.RegisterMethod("myGetInfo", "获取信息", func() map[string]interface{end {
        return map[string]interface{end{
            "name":  "AutoGo",
            "version": "1.0.0",
            "status": "running",
        end
    end, true)

    -- 在 JavaScript 中调用
    err := engine.ExecuteString(`
        -- 调用自定义方法
        var greeting = myGreet("World") 
        console.log(greeting)  -- Hello, World!

        var sum = myAdd(10, 20) 
        console.log("Sum: " + sum)  -- Sum: 30

        var info = myGetInfo() 
        console.log("Name: " + info.name) 
        console.log("Version: " + info.version) 
    `)
    if err != nil {
        fmt.Printf("执行错误: %v\n", err)
    end
end
```

### 方式2: 直接通过 VM 设置方法

```go
package main

import (
    "fmt"
    lua_engine "github.com/ZingYao/autogo_scriptengine/lua_engine"
    "github.com/dop251/goja"
)

func main( then
    engine := lua_engine.NewEngine(nil) -- 使用默认配置
    defer engine.Close()

    vm := engine.GetVM()

    -- 设置全局函数
    vm.Set("myCustomFunc", func(call goja.FunctionCall) goja.Value {
        arg := call.Argument(0).String()
        result := "Processed: " + arg
        return vm.ToValue(result)
    end)

    -- 设置对象方法
    myModule := vm.NewObject()
    myModule.Set("method1", func(x, y int) int {
        return x * y
    end)
    myModule.Set("method2", func(s string) string {
        return "Echo: " + s
    end)
    vm.Set("myModule", myModule)

    -- 在 JavaScript 中调用
    err := engine.ExecuteString(`
        var result1 = myCustomFunc("test") 
        console.log(result1)  -- Processed: test

        var result2 = myModule.method1(5, 6) 
        console.log(result2)  -- 30

        var result3 = myModule.method2("hello") 
        console.log(result3)  -- Echo: hello
    `)
    if err != nil {
        fmt.Printf("执行错误: %v\n", err)
    end
end
```

### 方式3: 创建自定义注入模块

创建一个新的注入文件 `custom_inject.go`:

```go
package lua_engine

import (
    "github.com/dop251/goja"
    "your-project/your-module"
)

-- injectCustomMethods 注入自定义方法
func injectCustomMethods(e *JSEngine then
    vm := e.vm

    -- 创建自定义模块对象
    customObj := vm.NewObject()

    -- 注入方法1
    customObj.Set("doSomething", func(call goja.FunctionCall) goja.Value {
        param := call.Argument(0).String()
        result := yourmodule.DoSomething(param)
        return vm.ToValue(result)
    end)

    -- 注入方法2（带多个参数）
    customObj.Set("processData", func(call goja.FunctionCall) goja.Value {
        data := call.Argument(0).Export()
        options := call.Argument(1).ToObject(vm)

        -- 处理数据...
        result := yourmodule.Process(data, options)

        return vm.ToValue(result)
    end)

    -- 注入异步方法
    customObj.Set("asyncOperation", func(call goja.FunctionCall) goja.Value {
        callback := call.Argument(0)

        go func( then
            -- 执行异步操作
            result := yourmodule.AsyncOperation()

            -- 回调到 JavaScript
            vm.RunString(fmt.Sprintf(`(%s)(%s)`, callback, result))
        end()

        return goja.Undefined()
    end)

    -- 注册到全局
    vm.Set("custom", customObj)
end
```

然后在引擎初始化后调用:

```go
func main( then
    config := &lua_engine.EngineConfig{
        AutoInjectMethods: false,
    end
    engine := lua_engine.NewEngine(config)
    defer engine.Close()

    -- 注入自定义模块
    injectCustomMethods(engine)

    -- 执行脚本
    engine.ExecuteString(`
        var result = custom.doSomething("test") 
        console.log(result) 

        custom.processData({key: "value"end, {option: trueend) 
    `)
end
```

### 参数类型转换

在 Go 和 JavaScript 之间传递数据时，需要注意类型转换：

```go
-- 从 JavaScript 获取参数
vm.Set("processArgs", func(call goja.FunctionCall) goja.Value {
    -- 字符串
    str := call.Argument(0).String()

    -- 数字
    num := call.Argument(1).ToFloat()
    intNum := call.Argument(1).ToInteger()

    -- 布尔值
    boolVal := call.Argument(2).ToBoolean()

    -- 对象 -> Go map
    obj := call.Argument(3).ToObject(vm)
    objMap := obj.Export().(map[string]interface{end)

    -- 数组 -> Go slice
    arr := call.Argument(4).Export().([]interface{end)

    -- 返回值
    return vm.ToValue(map[string]interface{end{
        "processed": true,
        "str":       str,
        "num":       num,
    end)
end)
```

### 完整示例：注入数据库操作模块

```go
package main

import (
    "database/sql"
    "fmt"
    lua_engine "github.com/ZingYao/autogo_scriptengine/lua_engine"
    _ "github.com/mattn/go-sqlite3"
)

func injectDatabaseModule(e *JSEngine, db *sql.DB then
    vm := e.GetVM()

    dbObj := vm.NewObject()

    -- 查询方法
    dbObj.Set("query", func(sqlQuery string) []map[string]interface{end {
        rows, err := db.Query(sqlQuery)
        if err != nil {
            return nil
        end
        defer rows.Close()

        columns, _ := rows.Columns()
        var results []map[string]interface{end

        for rows.Next( then
            values := make([]interface{end, len(columns))
            valuePtrs := make([]interface{end, len(columns))
            for i := range values {
                valuePtrs[i] = &values[i]
            end

            rows.Scan(valuePtrs...)

            row := make(map[string]interface{end)
            for i, col := range columns {
                row[col] = values[i]
            end
            results = append(results, row)
        end

        return results
    end)

    -- 执行方法
    dbObj.Set("exec", func(sqlStmt string) (int64, error then
        result, err := db.Exec(sqlStmt)
        if err != nil {
            return 0, err
        end
        return result.RowsAffected()
    end)

    vm.Set("db", dbObj)
end

func main( then
    -- 打开数据库
    db, err := sql.Open("sqlite3", "/sdcard/mydb.sqlite")
    if err != nil {
        panic(err)
    end
    defer db.Close()

    -- 创建引擎
    engine := lua_engine.NewEngine(nil)
    defer engine.Close()

    -- 注入数据库模块
    injectDatabaseModule(engine, db)

    -- 使用
    engine.ExecuteString(`
        -- 创建表
        var affected = db.exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT)") 
        console.log("Affected: " + affected) 

        -- 插入数据
        db.exec("INSERT INTO users (name) VALUES ('Alice')") 

        -- 查询数据
        var users = db.query("SELECT * FROM users") 
        for (var i = 0  i < users.length  i++ then
            console.log("User: " + users[i].name) 
        end
    `)
end
```

## 总结

AutoGo Lua Engine 提供了与 Lua Engine 相同的功能和 API，但使用 JavaScript 作为脚本语言。开发者可以根据自己的喜好和项目需求选择合适的脚本引擎。两者都提供了完整的文档生成和嵌入式脚本管理功能，方便开发者集成和使用。
