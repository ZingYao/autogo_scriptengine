-- Lua 引擎模块化测试脚本
-- 支持选择要测试的模块

-- 测试配置
local testConfig = {
    -- 设置为 true 测试所有模块，或指定要测试的模块列表
    testAll = true,
    -- 指定要测试的模块列表
    testModules = {
        "app",
        "device",
        "motion",
        "files",
        "images",
        "storages",
        "system",
        "http",
        "media",
        "opencv",
        "ppocr",
        "coroutine",
        "dotocr",
        "ime",
        "plugin",
        "rhino",
        "uiacc",
        "utils",
        "yolo",
        "imgui"
    }
}

-- 判断是否测试指定模块
local function shouldTestModule(moduleName)
    if testConfig.testAll then
        return true
    end
    for _, module in ipairs(testConfig.testModules) do
        if module == moduleName then
            return true
        end
    end
    return false
end

-- 测试 app 模块
local function testAppModule()
    if not shouldTestModule("app") then
        return
    end
    print("\n=== 测试 app 模块 ===")
    
    local currentPackage = app.currentPackage()
    print("当前应用包名: " .. tostring(currentPackage))
    
    local currentActivity = app.currentActivity()
    print("当前应用类名: " .. tostring(currentActivity))
    
    local browser = app.getBrowserPackage()
    print("默认浏览器: " .. tostring(browser))
    
    local success = app.launch("com.android.settings", 0)
    print("启动设置应用: " .. tostring(success))
    
    local installed = app.isInstalled("com.android.settings")
    print("设置应用是否已安装: " .. tostring(installed))
    
    if installed then
        local appName = app.getName("com.android.settings")
        print("设置应用名称: " .. tostring(appName))
        
        local version = app.getVersion("com.android.settings")
        print("设置应用版本: " .. tostring(version))
    end
    
    app.openUrl("https://www.example.com")
    print("已打开 URL")
end

-- 测试 device 模块
local function testDeviceModule()
    if not shouldTestModule("device") then
        return
    end
    print("\n=== 测试 device 模块 ===")
    
    print("屏幕宽度: " .. tostring(device.width))
    print("屏幕高度: " .. tostring(device.height))
    print("SDK版本: " .. tostring(device.sdkInt))
    print("CPU架构: " .. tostring(device.cpuAbi))
    print("品牌: " .. tostring(device.brand))
    print("型号: " .. tostring(device.model))
    
    local imei = device.getImei()
    print("IMEI: " .. tostring(imei))
    
    local androidId = device.getAndroidId()
    print("Android ID: " .. tostring(androidId))
    
    local wifiMac = device.getWifiMac()
    print("WIFI MAC: " .. tostring(wifiMac))
    
    local ip = device.getIp()
    print("IP地址: " .. tostring(ip))
    
    local brightness = device.getBrightness()
    print("当前亮度: " .. tostring(brightness))
    
    local brightnessMode = device.getBrightnessMode()
    print("亮度模式: " .. tostring(brightnessMode))
    
    local musicVolume = device.getMusicVolume()
    print("当前音乐音量: " .. tostring(musicVolume))
    
    local notificationVolume = device.getNotificationVolume()
    print("当前通知音量: " .. tostring(notificationVolume))
    
    local alarmVolume = device.getAlarmVolume()
    print("当前闹钟音量: " .. tostring(alarmVolume))
    
    local musicMaxVolume = device.getMusicMaxVolume()
    print("音乐音量最大值: " .. tostring(musicMaxVolume))
    
    local notificationMaxVolume = device.getNotificationMaxVolume()
    print("通知音量最大值: " .. tostring(notificationMaxVolume))
    
    local alarmMaxVolume = device.getAlarmMaxVolume()
    print("闹钟音量最大值: " .. tostring(alarmMaxVolume))
    
    device.setMusicVolume(50)
    print("音乐音量已设置为50")
    
    device.setNotificationVolume(50)
    print("通知音量已设置为50")
    
    device.setAlarmVolume(50)
    print("闹钟音量已设置为50")
    
    local battery = device.getBattery()
    print("电量: " .. tostring(battery) .. "%")
    
    local batteryStatus = device.getBatteryStatus()
    print("电池状态: " .. tostring(batteryStatus))
    
    local totalMem = device.getTotalMem()
    print("总内存: " .. tostring(totalMem / 1024) .. "MB")
    
    local availMem = device.getAvailMem()
    print("可用内存: " .. tostring(availMem / 1024) .. "MB")
    
    local screenOn = device.isScreenOn()
    print("屏幕是否亮着: " .. tostring(screenOn))
    
    local screenUnlock = device.isScreenUnlock()
    print("屏幕是否已解锁: " .. tostring(screenUnlock))
    
    device.wakeUp()
    print("设备已唤醒")
    
    device.keepScreenOn()
    print("屏幕常亮已启用")
    
    device.vibrate(500)
    print("设备震动500毫秒")
    
    device.cancelVibration()
    print("震动已取消")
end

-- 测试 motion 模块
local function testMotionModule()
    if not shouldTestModule("motion") then
        return
    end
    print("\n=== 测试 motion 模块 ===")
    
    print("点击操作测试")
    click(500, 1000, 1, 0)
    
    print("长按操作测试")
    longClick(500, 1000, 1000, 1, 0)
    
    print("滑动操作测试")
    swipe(500, 1000, 600, 1100, 500, 1, 0)
    
    print("swipe2 操作测试")
    swipe2(500, 1000, 600, 1100, 500, 1, 0)
    
    print("触摸操作测试")
    touchDown(500, 1000, 1, 0)
    touchMove(600, 1100, 1, 0)
    touchUp(600, 1100, 1, 0)
    
    print("系统按键测试")
    home(0)
    sleep(500)
    back(0)
    sleep(500)
    recents(0)
    
    print("音量按键测试")
    volumeUp(0)
    sleep(200)
    volumeDown(0)
    
    print("其他按键测试")
    powerDialog()
    notifications()
    quickSettings()
    camera()
    
    print("按键动作测试")
    keyAction(3, 0) -- KEYCODE_HOME
end

-- 测试 files 模块
local function testFilesModule()
    if not shouldTestModule("files") then
        return
    end
    print("\n=== 测试 files 模块 ===")
    
    local testPath = "/sdcard/test_lua.txt"
    
    files.write(testPath, "Hello from Lua!")
    print("文件已写入")
    
    local content = files.read(testPath)
    print("文件内容: " .. tostring(content))
    
    local exists = files.exists(testPath)
    print("文件是否存在: " .. tostring(exists))
    
    local isFile = files.isFile(testPath)
    print("是否是文件: " .. tostring(isFile))
    
    local isDir = files.isDir("/sdcard/Download")
    print("Download 是否是文件夹: " .. tostring(isDir))
    
    local name = files.getName(testPath)
    print("文件名: " .. tostring(name))
    
    local nameWithoutExt = files.getNameWithoutExtension(testPath)
    print("文件名(不含扩展名): " .. tostring(nameWithoutExt))
    
    local ext = files.getExtension(testPath)
    print("文件扩展名: " .. tostring(ext))
    
    files.append(testPath, "\n追加的内容")
    print("内容已追加")
    
    local newContent = files.read(testPath)
    print("新文件内容: " .. tostring(newContent))
    
    local copyPath = "/sdcard/test_lua_copy.txt"
    files.copy(testPath, copyPath)
    print("文件已复制")
    
    local movePath = "/sdcard/test_lua_moved.txt"
    files.move(copyPath, movePath)
    print("文件已移动")
    
    files.rename(movePath, "test_lua_renamed.txt")
    print("文件已重命名")
    
    local absPath = files.path("./test.txt")
    print("绝对路径: " .. tostring(absPath))
    
    local fileList = files.listDir("/sdcard")
    print("sdcard 文件夹内容数量: " .. tostring(#fileList))
    
    files.remove("/sdcard/test_lua_renamed.txt")
    print("文件已删除")
    
    files.remove(testPath)
    print("测试文件已删除")
end

-- 测试 images 模块
local function testImagesModule()
    if not shouldTestModule("images") then
        return
    end
    print("\n=== 测试 images 模块 ===")
    
    local color = images.pixel(500, 1000, 0)
    print("像素颜色: " .. tostring(color))
    
    local img = images.captureScreen(0, 0, 1080, 1920, 0)
    if img ~= nil then
        print("屏幕截图成功")
        
        local testImgPath = "/sdcard/test_lua_image.png"
        images.save(img, testImgPath, 90)
        print("图片已保存")
        
        local clipped = images.clip(img, 100, 100, 200, 200)
        print("图片已裁剪")
        
        local resized = images.resize(clipped, 100, 100)
        print("图片已调整大小")
        
        local rotated = images.rotate(resized, 90)
        print("图片已旋转")
        
        local gray = images.grayscale(rotated)
        print("图片已灰度化")
        
        local binary = images.applyBinarization(rotated, 128)
        print("图片已二值化")
        
        files.remove(testImgPath)
        print("测试图片已删除")
    end
    
    local cmpResult = images.cmpColor(500, 1000, "#FF0000", 0.9, 0)
    print("颜色比较结果: " .. tostring(cmpResult))
    
    local x, y = images.findColor(0, 0, 1080, 1920, "#FF0000", 0.9, 0, 0)
    print("找到颜色在: " .. tostring(x) .. ", " .. tostring(y))
    
    local count = images.getColorCountInRegion(0, 0, 1080, 1920, "#FF0000", 0.9, 0)
    print("颜色数量: " .. tostring(count))
end

-- 测试 storages 模块
local function testStoragesModule()
    if not shouldTestModule("storages") then
        return
    end
    print("\n=== 测试 storages 模块 ===")
    
    local tableName = "test_table_lua"
    
    storages.put(tableName, "key1", "value1")
    print("值已存储")
    
    storages.put(tableName, "key2", "value2")
    print("值已存储")
    
    storages.put(tableName, "key3", "value3")
    print("值已存储")
    
    local value1 = storages.get(tableName, "key1")
    print("存储的值: " .. tostring(value1))
    
    local contains = storages.contains(tableName, "key1")
    print("是否包含键: " .. tostring(contains))
    
    local allData = storages.getAll(tableName)
    print("所有数据: " .. tostring(allData))
    
    storages.remove(tableName, "key3")
    print("键已删除")
    
    storages.clear(tableName)
    print("存储已清空")
end

-- 测试 json 模块
local function testJsonModule()
    if not shouldTestModule("json") then
        return
    end
    print("\n=== 测试 json 模块 ===")
    
    -- 测试序列化
    local testData = {
        name = "test",
        value = 123,
        flag = true,
        items = {1, 2, 3},
        nested = {
            key = "value"
        }
    }
    
    local jsonStr = json.stringify(testData)
    print("序列化结果: " .. tostring(jsonStr))
    
    -- 测试解析
    local parsedData = json.parse(jsonStr)
    print("解析结果: " .. tostring(parsedData))
    
    -- 测试格式化序列化
    local formattedJson = json.format(testData)
    print("格式化结果:\n" .. tostring(formattedJson))
    
    -- 测试复杂对象
    local complexData = {
        array = {10, 20, 30, 40},
        object = {
            a = 1,
            b = 2,
            c = 3
        },
        null_value = nil,
        bool_value = false
    }
    
    local complexJson = json.stringify(complexData)
    print("复杂对象序列化: " .. tostring(complexJson))
    
    local complexParsed = json.parse(complexJson)
    print("复杂对象解析: " .. tostring(complexParsed))
end

-- 测试 system 模块
local function testSystemModule()
    if not shouldTestModule("system") then
        return
    end
    print("\n=== 测试 system 模块 ===")
    
    local pid = system.getPid("com.android.systemui")
    print("SystemUI 进程ID: " .. tostring(pid))
    
    if pid > 0 then
        local memUsage = system.getMemoryUsage(pid)
        print("内存使用: " .. tostring(memUsage))
        
        local cpuUsage = system.getCpuUsage(pid)
        print("CPU使用率: " .. tostring(cpuUsage))
    end
end

-- 测试 http 模块
local function testHttpModule()
    if not shouldTestModule("http") then
        return
    end
    print("\n=== 测试 http 模块 ===")
    
    local url = "https://www.example.com"
    local response = http.get(url, 5000)
    print("HTTP GET 响应码: " .. tostring(response.code))
    print("HTTP GET 响应数据: " .. tostring(response.data))
    
    local postResponse = http.post(url, "test data", {}, 5000)
    print("HTTP POST 响应码: " .. tostring(postResponse.code))
    print("HTTP POST 响应数据: " .. tostring(postResponse.data))
end

-- 测试 media 模块
local function testMediaModule()
    if not shouldTestModule("media") then
        return
    end
    print("\n=== 测试 media 模块 ===")
    
    local mediaPath = "/sdcard/autogo_test_media.jpg"
    local mp3Path = "/sdcard/test.mp3"
    
    media.scanFile(mediaPath)
    print("文件已扫描")
    
    local playResult = media.playMP3(mp3Path)
    if playResult ~= nil then
        print("播放结果: " .. tostring(playResult))
    else
        print("MP3 播放成功")
    end
    
    media.sendSMS("10086", "测试短信")
    print("短信已发送")
end

-- 测试 opencv 模块
local function testOpencvModule()
    if not shouldTestModule("opencv") then
        return
    end
    print("\n=== 测试 opencv 模块 ===")
    print("OpenCV 模块已加载")
end

-- 测试 ppocr 模块
local function testPpocrModule()
    if not shouldTestModule("ppocr") then
        return
    end
    print("\n=== 测试 ppocr 模块 ===")
    print("PP-OCR 模块已加载")
end

-- 测试 coroutine 模块
local function testCoroutineModule()
    if not shouldTestModule("coroutine") then
        return
    end
    print("\n=== 测试 coroutine 模块 ===")
    
    -- 测试启动协程
    print("测试启动协程...")
    local coroutineId = coroutine.launch(function()
        print("协程开始执行")
        coroutine.sleep(500)
        print("协程执行完成")
    end, "testCoroutine", 0)
    print("协程ID: " .. tostring(coroutineId))
    
    -- 测试延迟执行
    print("测试延迟执行...")
    local delayId = coroutine.delay(1000, function()
        print("延迟任务执行")
    end)
    print("延迟任务ID: " .. tostring(delayId))
    
    -- 测试异步执行
    print("测试异步执行...")
    local asyncResult = coroutine.async(function()
        return "异步结果"
    end)
    print("异步结果: " .. tostring(asyncResult))
    
    -- 测试 await（简化版本）
    print("测试 await...")
    local value = coroutine.await("普通值")
    print("等待结果: " .. tostring(value))
    
    -- 等待协程完成
    coroutine.sleep(1500)
    
    -- 测试获取活跃协程数量
    print("测试获取活跃协程数量...")
    local activeCount = coroutine.getActiveCoroutines()
    print("活跃协程数量: " .. tostring(activeCount))
    
    -- 测试获取协程列表
    print("测试获取协程列表...")
    local coroutineList = coroutine.getCoroutineList()
    for i, coro in ipairs(coroutineList) do
        print("协程 " .. i .. ": ID=" .. tostring(coro.id) .. ", 名称=" .. tostring(coro.name) .. ", 状态=" .. tostring(coro.state))
    end
    
    -- 测试获取统计信息
    print("测试获取统计信息...")
    local stats = coroutine.getStats()
    print("总任务数: " .. tostring(stats.totalTasks))
    print("已完成: " .. tostring(stats.completed))
    print("失败: " .. tostring(stats.failed))
    print("已取消: " .. tostring(stats.cancelled))
    print("活跃协程: " .. tostring(stats.active))
    print("协程池数量: " .. tostring(stats.pools))
    
    -- 测试创建协程池
    print("测试创建协程池...")
    local poolName = coroutine.createPool("testPool", 2, 5)
    print("协程池名称: " .. tostring(poolName))
    
    -- 测试提交任务到协程池
    print("测试提交任务到协程池...")
    local submitSuccess = coroutine.submitToPool("testPool", function()
        print("协程池任务执行")
        coroutine.sleep(300)
        print("协程池任务完成")
    end, 0)
    print("任务提交成功: " .. tostring(submitSuccess))
    
    -- 等待任务完成
    coroutine.sleep(500)
    
    -- 测试获取协程池统计信息
    print("测试获取协程池统计信息...")
    local poolStats = coroutine.getPoolStats("testPool")
    print("协程池名称: " .. tostring(poolStats.name))
    print("最大工作协程数: " .. tostring(poolStats.maxWorkers))
    print("最大任务数: " .. tostring(poolStats.maxTasks))
    print("活跃工作协程数: " .. tostring(poolStats.active))
    print("队列中任务数: " .. tostring(poolStats.queued))
    print("工作协程总数: " .. tostring(poolStats.workers))
    print("是否已关闭: " .. tostring(poolStats.closed))
    
    -- 测试关闭协程池
    print("测试关闭协程池...")
    local closeSuccess = coroutine.closePool("testPool")
    print("协程池关闭成功: " .. tostring(closeSuccess))
    
    -- 测试列出所有协程池
    print("测试列出所有协程池...")
    local poolList = coroutine.listPools()
    for i, pool in ipairs(poolList) do
        print("协程池 " .. i .. ": " .. tostring(pool))
    end
    
    -- 测试调度器方法
    print("测试调度器方法...")
    coroutine.setScheduleStrategy("fifo")
    local strategy = coroutine.getScheduleStrategy()
    print("当前调度策略: " .. tostring(strategy))
    
    coroutine.setPriority("testPriority", 10)
    local priority = coroutine.getPriority("testPriority")
    print("协程优先级: " .. tostring(priority))
    
    print("coroutine 模块测试完成")
end

-- 测试 dotocr 模块
local function testDotocrModule()
    if not shouldTestModule("dotocr") then
        return
    end
    print("\n=== 测试 dotocr 模块 ===")
    print("DotOCR 模块已加载")
end

-- 测试 ime 模块
local function testImeModule()
    if not shouldTestModule("ime") then
        return
    end
    print("\n=== 测试 ime 模块 ===")
    print("IME 模块已加载")
end

-- 测试 plugin 模块
local function testPluginModule()
    if not shouldTestModule("plugin") then
        return
    end
    print("\n=== 测试 plugin 模块 ===")
    print("Plugin 模块已加载")
end

-- 测试 rhino 模块
local function testRhinoModule()
    if not shouldTestModule("rhino") then
        return
    end
    print("\n=== 测试 rhino 模块 ===")
    print("Rhino 模块已加载")
end

-- 测试 uiacc 模块
local function testUiaccModule()
    if not shouldTestModule("uiacc") then
        return
    end
    print("\n=== 测试 uiacc 模块 ===")
    print("UIACC 模块已加载")
end

-- 测试 utils 模块
local function testUtilsModule()
    if not shouldTestModule("utils") then
        return
    end
    print("\n=== 测试 utils 模块 ===")
    print("Utils 模块已加载")
end

-- 测试 yolo 模块
local function testYoloModule()
    if not shouldTestModule("yolo") then
        return
    end
    print("\n=== 测试 yolo 模块 ===")
    print("YOLO 模块已加载")
end

-- 主函数
local function main()
    print("=== Lua 引擎模块化测试开始 ===")
    print("测试配置:")
    print("  测试所有模块: " .. tostring(testConfig.testAll))
    print("  要测试的模块数量: " .. tostring(#testConfig.testModules))
    
    testAppModule()
    testDeviceModule()
    testMotionModule()
    testFilesModule()
    testImagesModule()
    testStoragesModule()
    testJsonModule()
    testSystemModule()
    testHttpModule()
    testMediaModule()
    testOpencvModule()
    testPpocrModule()
    testCoroutineModule()
    testDotocrModule()
    testImeModule()
    testPluginModule()
    testRhinoModule()
    testUiaccModule()
    testUtilsModule()
    testYoloModule()
    
    print("\n=== Lua 引擎模块化测试完成 ===")
end

-- 运行主函数
main()
