// AutoGo JavaScript 测试脚本
// 测试所有可用模块的功能

console.log("=== JavaScript 引擎测试开始 ===");

// ========== 1. 应用管理测试 (app) ==========
console.log("\n--- 1. 应用管理 (app) 测试 ---");

// 获取当前应用包名
var currentPackage = app.currentPackage();
console.log("当前应用包名: " + currentPackage);

// 获取当前 Activity
var currentActivity = app.currentActivity();
console.log("当前 Activity: " + currentActivity);

// 检查应用是否已安装
var isInstalled = app.isInstalled("com.android.settings");
console.log("设置应用是否已安装: " + isInstalled);

// 获取浏览器包名
var browserPackage = app.getBrowserPackage();
console.log("默认浏览器包名: " + browserPackage);

// ========== 2. 设备管理测试 (device) ==========
console.log("\n--- 2. 设备管理 (device) 测试 ---");

// 获取设备信息
console.log("设备分辨率: " + device.width + "x" + device.height);
console.log("SDK 版本: " + device.sdkInt);
console.log("CPU 架构: " + device.cpuAbi);

// 获取设备标识
var imei = device.getImei();
console.log("IMEI: " + imei);

var androidId = device.getAndroidId();
console.log("Android ID: " + androidId);

var ip = device.getIp();
console.log("IP 地址: " + ip);

// 获取电池信息
var battery = device.getBattery();
console.log("电量: " + battery + "%");

// 获取内存信息
var totalMem = device.getTotalMem();
var availMem = device.getAvailMem();
console.log("总内存: " + totalMem + ", 可用内存: " + availMem);

// 获取音量信息
var musicVolume = device.getMusicVolume();
var musicMaxVolume = device.getMusicMaxVolume();
console.log("媒体音量: " + musicVolume + "/" + musicMaxVolume);

// 检查屏幕状态
var isScreenOn = device.isScreenOn();
console.log("屏幕是否亮着: " + isScreenOn);

// ========== 3. 触摸操作测试 (motion) ==========
console.log("\n--- 3. 触摸操作 (motion) 测试 ---");

// 测试点击
console.log("测试点击操作...");
click(500, 500, 1);
console.log("点击完成");

// 测试滑动
console.log("测试滑动操作...");
swipe(100, 500, 500, 500, 500);
console.log("滑动完成");

// 测试按键
console.log("测试按键操作...");
back();
console.log("返回键按下");

// ========== 4. 文件操作测试 (files) ==========
console.log("\n--- 4. 文件操作 (files) 测试 ---");

// 创建测试文件
var testPath = "/sdcard/autogo_test_js.txt";
files.write(testPath, "Hello from JavaScript!");
console.log("文件写入成功: " + testPath);

// 读取文件
var content = files.read(testPath);
console.log("文件内容: " + content);

// 检查文件是否存在
var exists = files.exists(testPath);
console.log("文件是否存在: " + exists);

// 获取文件名
var fileName = files.getName(testPath);
console.log("文件名: " + fileName);

// 删除测试文件
files.remove(testPath);
console.log("测试文件已删除");

// ========== 5. 图像处理测试 (images) ==========
console.log("\n--- 5. 图像处理 (images) 测试 ---");

// 获取像素颜色
var pixel = images.pixel(100, 100);
console.log("像素颜色 (100,100): " + pixel);

// 截取屏幕
var img = images.captureScreen(0, 0, 500, 500);
console.log("屏幕截图成功");

// 比较颜色
var colorMatch = images.cmpColor(100, 100, "#FF0000", 0.9);
console.log("颜色匹配结果: " + colorMatch);

// ========== 6. 存储管理测试 (storages) ==========
console.log("\n--- 6. 存储管理 (storages) 测试 ---");

// 写入存储
storages.put("test_table", "key1", "value1");
storages.put("test_table", "key2", "value2");
console.log("存储写入成功");

// 读取存储
var value1 = storages.get("test_table", "key1");
console.log("读取 key1: " + value1);

// 检查键是否存在
var contains = storages.contains("test_table", "key1");
console.log("key1 是否存在: " + contains);

// 获取所有数据
var allData = storages.getAll("test_table");
console.log("所有数据: " + JSON.stringify(allData));

// 清空存储
storages.clear("test_table");
console.log("存储已清空");

// ========== 7. 系统管理测试 (system) ==========
console.log("\n--- 7. 系统管理 (system) 测试 ---");

// 获取进程 ID
var pid = system.getPid("com.android.systemui");
console.log("systemui 进程 ID: " + pid);

// 获取 CPU 使用率
if (pid > 0) {
    var cpuUsage = system.getCpuUsage(pid);
    console.log("CPU 使用率: " + cpuUsage);
    
    var memUsage = system.getMemoryUsage(pid);
    console.log("内存使用: " + memUsage);
}

// ========== 8. 网络请求测试 (http) ==========
console.log("\n--- 8. 网络请求 (http) 测试 ---");

// 发送 GET 请求
var statusCode, response = http.get("https://httpbin.org/get", 5000);
console.log("HTTP 状态码: " + statusCode);
console.log("响应长度: " + response.length);

// ========== 9. 方法管理测试 ==========
console.log("\n--- 9. 方法管理测试 ---");

// 列出所有方法
var methods = listMethods();
console.log("已注册方法数量: " + methods.length);

// 显示部分方法
for (var i = 0; i < Math.min(5, methods.length); i++) {
    var method = methods[i];
    console.log("方法 " + (i+1) + ": " + method.name + " - " + method.description);
}

// ========== 10. 控制台测试 (console) ==========
console.log("\n--- 10. 控制台 (console) 测试 ---");

console.log("普通日志");
console.log("格式化输出: %s = %d", "测试", 123);

// ========== 11. HUD 测试 ==========
console.log("\n--- 11. HUD 测试 ---");

// 显示 HUD
hud.show("JavaScript 测试进行中...", 2000);
console.log("HUD 显示完成");

// ========== 12. 工具类测试 (utils) ==========
console.log("\n--- 12. 工具类 (utils) 测试 ---");

// 测试随机数
var randomNum = utils.random(1, 100);
console.log("随机数 (1-100): " + randomNum);

// ========== 13. 睡眠测试 ==========
console.log("\n--- 13. 睡眠测试 ---");
console.log("睡眠 500ms...");
sleep(500);
console.log("睡眠完成");

// ========== 14. 方法重写测试 ==========
console.log("\n--- 14. 方法重写测试 ---");

// 尝试重写方法
var overrideResult = overrideMethod("click", function(x, y, fingerID) {
    console.log("重写的点击方法被调用: (" + x + ", " + y + ")");
    return true;
});
console.log("方法重写结果: " + overrideResult);

// 测试重写的方法
click(300, 300, 1);

// 恢复方法
var restoreResult = restoreMethod("click");
console.log("方法恢复结果: " + restoreResult);

console.log("\n=== JavaScript 引擎测试完成 ===");
