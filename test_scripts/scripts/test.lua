-- AutoGo Lua 测试脚本
-- 测试所有可用模块的功能

print("=== Lua 引擎测试开始 ===")

-- ========== 1. 应用管理测试 (app) ==========
print("\n--- 1. 应用管理 (app) 测试 ---")

-- 获取当前应用包名
local currentPackage = app_currentPackage()
print("当前应用包名: " .. currentPackage)

-- 获取当前 Activity
local currentActivity = app_currentActivity()
print("当前 Activity: " .. currentActivity)

-- 检查应用是否已安装
local isInstalled = app_isInstalled("com.android.settings")
print("设置应用是否已安装: " .. tostring(isInstalled))

-- 获取浏览器包名
local browserPackage = app_getBrowserPackage()
print("默认浏览器包名: " .. browserPackage)

-- ========== 2. 设备管理测试 (device) ==========
print("\n--- 2. 设备管理 (device) 测试 ---")

-- 获取设备信息
print("设备分辨率: " .. device.width .. "x" .. device.height)
print("SDK 版本: " .. device.sdkInt)
print("CPU 架构: " .. device.cpuAbi)

-- 获取设备标识
local imei = device_getImei()
print("IMEI: " .. imei)

local androidId = device_getAndroidId()
print("Android ID: " .. androidId)

local ip = device_getIp()
print("IP 地址: " .. ip)

-- 获取电池信息
local battery = device_getBattery()
print("电量: " .. battery .. "%")

-- 获取内存信息
local totalMem = device_getTotalMem()
local availMem = device_getAvailMem()
print("总内存: " .. totalMem .. ", 可用内存: " .. availMem)

-- 获取音量信息
local musicVolume = device_getMusicVolume()
local musicMaxVolume = device_getMusicMaxVolume()
print("媒体音量: " .. musicVolume .. "/" .. musicMaxVolume)

-- 检查屏幕状态
local isScreenOn = device_isScreenOn()
print("屏幕是否亮着: " .. tostring(isScreenOn))

-- ========== 3. 触摸操作测试 (motion) ==========
print("\n--- 3. 触摸操作 (motion) 测试 ---")

-- 测试点击
print("测试点击操作...")
click(500, 500, 1)
print("点击完成")

-- 测试滑动
print("测试滑动操作...")
swipe(100, 500, 500, 500, 500)
print("滑动完成")

-- 测试按键
print("测试按键操作...")
back()
print("返回键按下")

-- ========== 4. 文件操作测试 (files) ==========
print("\n--- 4. 文件操作 (files) 测试 ---")

-- 创建测试文件
local testPath = "/sdcard/autogo_test_lua.txt"
files_write(testPath, "Hello from Lua!")
print("文件写入成功: " .. testPath)

-- 读取文件
local content = files_read(testPath)
print("文件内容: " .. content)

-- 检查文件是否存在
local exists = files_exists(testPath)
print("文件是否存在: " .. tostring(exists))

-- 获取文件名
local fileName = files_getName(testPath)
print("文件名: " .. fileName)

-- 删除测试文件
files_remove(testPath)
print("测试文件已删除")

-- ========== 5. 图像处理测试 (images) ==========
print("\n--- 5. 图像处理 (images) 测试 ---")

-- 获取像素颜色
local pixel = images_pixel(100, 100)
print("像素颜色 (100,100): " .. pixel)

-- 截取屏幕
local img = images_captureScreen(0, 0, 500, 500)
print("屏幕截图成功")

-- 比较颜色
local colorMatch = images_cmpColor(100, 100, "#FF0000", 0.9)
print("颜色匹配结果: " .. tostring(colorMatch))

-- ========== 6. 存储管理测试 (storages) ==========
print("\n--- 6. 存储管理 (storages) 测试 ---")

-- 写入存储
storages_put("test_table_lua", "key1", "value1")
storages_put("test_table_lua", "key2", "value2")
print("存储写入成功")

-- 读取存储
local value1 = storages_get("test_table_lua", "key1")
print("读取 key1: " .. value1)

-- 检查键是否存在
local contains = storages_contains("test_table_lua", "key1")
print("key1 是否存在: " .. tostring(contains))

-- 获取所有数据
local allData = storages_getAll("test_table_lua")
print("所有数据条数: " .. #allData)

-- 清空存储
storages_clear("test_table_lua")
print("存储已清空")

-- ========== 7. 系统管理测试 (system) ==========
print("\n--- 7. 系统管理 (system) 测试 ---")

-- 获取进程 ID
local pid = system_getPid("com.android.systemui")
print("systemui 进程 ID: " .. pid)

-- 获取 CPU 使用率
if pid > 0 then
    local cpuUsage = system_getCpuUsage(pid)
    print("CPU 使用率: " .. cpuUsage)
    
    local memUsage = system_getMemoryUsage(pid)
    print("内存使用: " .. memUsage)
end

-- ========== 8. 网络请求测试 (http) ==========
print("\n--- 8. 网络请求 (http) 测试 ---")

-- 发送 GET 请求
local statusCode, response = http_get("https://httpbin.org/get", 5000)
print("HTTP 状态码: " .. statusCode)
print("响应长度: " .. string.len(response))

-- ========== 9. 方法管理测试 ==========
print("\n--- 9. 方法管理测试 ---")

-- 列出所有方法
local methods = listMethods()
print("已注册方法数量: " .. #methods)

-- 显示部分方法
for i = 1, math.min(5, #methods) do
    local method = methods[i]
    print("方法 " .. i .. ": " .. method.name .. " - " .. method.description)
end

-- ========== 10. HUD 测试 ==========
print("\n--- 10. HUD 测试 ---")

-- 显示 HUD
hud_show("Lua 测试进行中...", 2000)
print("HUD 显示完成")

-- ========== 11. 工具类测试 (utils) ==========
print("\n--- 11. 工具类 (utils) 测试 ---")

-- 测试随机数
local randomNum = utils_random(1, 100)
print("随机数 (1-100): " .. randomNum)

-- ========== 12. 睡眠测试 ==========
print("\n--- 12. 睡眠测试 ---")
print("睡眠 500ms...")
sleep(500)
print("睡眠完成")

-- ========== 13. 方法重写测试 ==========
print("\n--- 13. 方法重写测试 ---")

-- Lua 中方法重写的示例
print("方法重写功能测试")

-- ========== 14. 协程测试 (Lua 特有) ==========
print("\n--- 14. 协程测试 (Lua 特有) ---")

-- 创建协程
local coroutineScript = [[
    print("协程开始执行")
    for i = 1, 3 do
        print("协程计数: " .. i)
        coroutine.yield()
    end
    print("协程执行完成")
]]

local coId = createCoroutine(coroutineScript)
print("创建协程 ID: " .. coId)

-- 恢复协程执行
for i = 1, 3 do
    local results, status = resumeCoroutine(coId)
    print("协程状态: " .. status)
end

-- 列出所有协程
local coroutines = listCoroutines()
print("协程数量: " .. #coroutines)

-- 移除协程
removeCoroutine(coId)
print("协程已移除")

print("\n=== Lua 引擎测试完成 ===")
