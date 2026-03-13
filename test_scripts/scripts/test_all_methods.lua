--AutoGo Lua引擎全量方法测试脚本
-- 测试所有注入的方法

print("========================================")
print("AutoGo Lua 引擎全量方法测试")
print("========================================")

-- ==================== 核心函数测试 ====================
print("\n[核心函数] sleep测试")
print("sleep(100) - 等待100ms")
sleep(100)
print("sleep测试通过")

-- ==================== app模块测试 ====================
print("\n[app模块] 测试开始")

-- app.currentPackage
local pkg = app_currentPackage()
print("当前包名: " .. tostring(pkg))

-- app.currentActivity
local activity = app_currentActivity()
print("当前Activity: " .. tostring(activity))

-- app.getBrowserPackage
local browser = app_getBrowserPackage()
print("默认浏览器包名: " .. tostring(browser))

-- app.isInstalled
local isInstalled = app_isInstalled("com.android.settings")
print("设置应用是否安装: " .. tostring(isInstalled))

-- app.launch (需要实际包名)
-- local launched = app_launch("com.android.settings", 0)
-- print("启动设置应用: " .. tostring(launched))

print("[app模块] 测试完成")

-- ==================== device模块测试 ====================
print("\n[device模块] 测试开始")

-- 设备信息
print("SDK版本: " .. tostring(device_sdkInt()))
print("CPU架构: " .. tostring(device_cpuAbi()))
print("构建ID: " .. tostring(device_buildId()))
print("主板型号: " .. tostring(device_broad()))
print("品牌: " .. tostring(device_brand()))
print("设备名: " .. tostring(device_device()))
print("型号: " .. tostring(device_model()))
print("产品名: " .. tostring(device_product()))
print("Bootloader版本: " .. tostring(device_bootloader()))
print("硬件名: " .. tostring(device_hardware()))
print("指纹: " .. tostring(device_fingerprint()))
print("序列号: " .. tostring(device_serial()))
print("增量版本: " .. tostring(device_incremental()))
print("Android版本: " .. tostring(device_release()))
print("基础OS: " .. tostring(device_baseOS()))
print("安全补丁: " .. tostring(device_securityPatch()))
print("代号: " .. tostring(device_codename()))

-- 设备标识
print("IMEI: " .. tostring(device_getImei()))
print("Android ID: " .. tostring(device_getAndroidId()))
print("WiFi MAC: " .. tostring(device_getWifiMac()))
print("以太网MAC: " .. tostring(device_getWlanMac()))
print("IP地址: " .. tostring(device_getIp()))

-- 音量和亮度
print("亮度: " .. tostring(device_getBrightness()))
print("亮度模式: " .. tostring(device_getBrightnessMode()))
print("媒体音量: " .. tostring(device_getMusicVolume()))
print("通知音量: " .. tostring(device_getNotificationVolume()))
print("闹钟音量: " .. tostring(device_getAlarmVolume()))
print("媒体最大音量: " .. tostring(device_getMusicMaxVolume()))
print("通知最大音量: " .. tostring(device_getNotificationMaxVolume()))
print("闹钟最大音量: " .. tostring(device_getAlarmMaxVolume()))

-- 电池和内存
print("电量: " .. tostring(device_getBattery()) .. "%")
print("电池状态: " .. tostring(device_getBatteryStatus()))
print("总内存: " .. tostring(device_getTotalMem()))
print("可用内存: " .. tostring(device_getAvailMem()))

-- 屏幕状态
print("屏幕亮起: " .. tostring(device_isScreenOn()))
print("屏幕解锁: " .. tostring(device_isScreenUnlock()))

print("[device模块] 测试完成")

-- ==================== motion模块测试 ====================
print("\n[motion模块] 测试开始")

-- 触摸操作测试 (需要实际坐标)
print("touchDown/touchMove/touchUp - 需要实际坐标，跳过")
-- touchDown(100, 100, 0, 0)
-- touchMove(200, 200, 0, 0)
-- touchUp(200, 200, 0, 0)

-- 点击测试
print("click - 需要实际坐标，跳过")
-- click(100, 100, 0, 0)

-- 长按测试
print("longClick - 需要实际坐标，跳过")
-- longClick(100, 100, 500, 0, 0)

-- 滑动测试
print("swipe - 需要实际坐标，跳过")
-- swipe(100, 100, 300, 300, 500, 0, 0)

-- 按键测试
print("home/back/recents - 需要实际设备，跳过")
-- home(0)
-- back(0)
-- recents(0)

-- 音量键测试
print("volumeUp/volumeDown - 需要实际设备，跳过")
-- volumeUp(0)
-- volumeDown(0)

print("[motion模块] 测试完成")

-- ==================== files模块测试 ====================
print("\n[files模块] 测试开始")

-- 创建测试文件
local testPath = "/data/local/tmp/test_autogo.txt"
local testContent = "Hello AutoGo Lua Engine!"

-- 写入文件
files_write(testPath, testContent)
print("写入文件: " .. testPath)

-- 检查文件是否存在
local exists = files_exists(testPath)
print("文件存在: " .. tostring(exists))

-- 读取文件
local content = files_read(testPath)
print("文件内容: " .. tostring(content))

-- 检查是否是文件
local isFile = files_isFile(testPath)
print("是文件: " .. tostring(isFile))

-- 获取文件名
local name = files_getName(testPath)
print("文件名: " .. tostring(name))

-- 获取不含扩展名的文件名
local nameWithoutExt = files_getNameWithoutExtension(testPath)
print("不含扩展名的文件名: " .. tostring(nameWithoutExt))

-- 获取扩展名
local ext = files_getExtension(testPath)
print("扩展名: " .. tostring(ext))

-- 追加内容
files_append(testPath, "\nAppended line")
print("追加内容成功")

-- 删除文件
local removed = files_remove(testPath)
print("删除文件: " .. tostring(removed))

print("[files模块] 测试完成")

-- ==================== storages模块测试 ====================
print("\n[storages模块] 测试开始")

local tableName = "test_table"
local key = "test_key"
local value = "test_value"

-- 写入键值对
storages_put(tableName, key, value)
print("写入键值对: " .. key .. " = " .. value)

-- 读取键值
local readValue = storages_get(tableName, key)
print("读取键值: " .. tostring(readValue))

-- 检查键是否存在
local contains = storages_contains(tableName, key)
print("键存在: " .. tostring(contains))

-- 获取所有键值对
local allData = storages_getAll(tableName)
print("所有键值对: " .. tostring(allData))

-- 删除键
storages_remove(tableName, key)
print("删除键: " .. key)

-- 清空表
storages_clear(tableName)
print("清空表: " .. tableName)

print("[storages模块] 测试完成")

-- ==================== system模块测试 ====================
print("\n[system模块] 测试开始")

-- 获取进程ID
local pid = system_getPid("init")
print("init进程PID: " .. tostring(pid))

-- 获取内存使用
if pid > 0 then
    local memUsage = system_getMemoryUsage(pid)
    print("内存使用: " .. tostring(memUsage))
end

print("[system模块] 测试完成")

-- ==================== images模块测试 ====================
print("\n[images模块] 测试开始")

-- 截屏测试 (需要实际设备)
print("captureScreen - 需要实际设备，跳过")
-- local img = images_captureScreen(0, 0, 100, 100, 0)

-- 像素颜色测试
print("pixel - 需要实际设备，跳过")
-- local color = images_pixel(100, 100, 0)

-- 颜色比较测试
print("cmpColor - 需要实际设备，跳过")
-- local cmpResult = images_cmpColor(100, 100, "#FF0000", 0.9, 0)

-- 查找颜色测试
print("findColor - 需要实际设备，跳过")
-- local x, y = images_findColor(0, 0, 500, 500, "#FF0000", 0.9, 0, 0)

print("[images模块] 测试完成")

-- ==================== http模块测试 ====================
print("\n[http模块] 测试开始")

-- GET请求测试
print("http_get - 测试HTTP GET请求")
local code, data = http_get("https://httpbin.org/get", 5000)
print("响应码: " .. tostring(code))
if data then
    print("响应数据长度: " .. string.len(data))
end

print("[http模块] 测试完成")

-- ==================== console模块测试 ====================
print("\n[console模块] 测试开始")

-- 创建控制台
local console = console_new()
print("创建控制台成功")

-- 设置窗口大小
console_setWindowSize(console, 400, 300)
print("设置窗口大小: 400x300")

-- 设置窗口位置
console_setWindowPosition(console, 100, 100)
print("设置窗口位置: 100, 100")

-- 设置窗口颜色
console_setWindowColor(console, "#000000")
print("设置窗口颜色: #000000")

-- 设置文本颜色
console_setTextColor(console, "#FFFFFF")
print("设置文本颜色: #FFFFFF")

-- 设置文本大小
console_setTextSize(console, 14)
print("设置文本大小: 14")

-- 打印内容
console_println(console, "Hello from Lua!")
print("打印内容成功")

-- 显示控制台
console_show(console)
print("显示控制台")

-- 检查可见性
local visible = console_isVisible(console)
print("控制台可见: " .. tostring(visible))

-- 隐藏控制台
console_hide(console)
print("隐藏控制台")

-- 销毁控制台
console_destroy(console)
print("销毁控制台")

print("[console模块] 测试完成")

-- ==================== imgui模块测试 ====================
print("\n[imgui模块] 测试开始")

-- 测试版本
print("ImGui版本: " .. tostring(imgui.version))

-- 测试常量
print("WindowFlags.None = " .. tostring(imgui.WindowFlags.None))
print("Col.Text = " .. tostring(imgui.Col.Text))
print("MouseButton.Left = " .. tostring(imgui.MouseButton.Left))
print("Key.Tab = " .. tostring(imgui.Key.Tab))

print("[imgui模块] 测试完成")

-- ==================== 方法注册测试 ====================
print("\n[方法注册] 测试开始")

-- 列出所有方法
local methods = listMethods()
print("已注册方法数量: " .. #methods)

-- 注册自定义方法
registerMethod("custom.test", "自定义测试方法", true)
print("注册自定义方法: custom.test")

-- 再次列出方法
methods = listMethods()
print("注册后方法数量: " .. #methods)

print("[方法注册] 测试完成")

-- ==================== 协程测试 ====================
print("\n[协程] 测试开始")

-- 创建协程
local coId = createCoroutine([[
    print("协程开始执行")
    sleep(100)
    print("协程执行中...")
    sleep(100)
    print("协程执行结束")
    return "done"
]])
print("创建协程ID: " .. tostring(coId))

-- 恢复协程
local results, status = resumeCoroutine(coId)
print("协程状态: " .. tostring(status))

-- 列出协程
local coroutines = listCoroutines()
print("协程数量: " .. #coroutines)

-- 移除协程
removeCoroutine(coId)
print("移除协程: " .. tostring(coId))

print("[协程] 测试完成")

print("\n========================================")
print("AutoGo Lua 引擎全量方法测试完成")
print("========================================")
