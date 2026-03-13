// AutoGo JavaScript 引擎全量方法测试脚本
// 测试所有注入的方法

console.log("========================================");
console.log("AutoGo JavaScript 引擎全量方法测试");
console.log("========================================");

// ==================== 核心函数测试 ====================
console.log("\n[核心函数] sleep测试");
console.log("sleep(100) - 等待100ms");
sleep(100);
console.log("sleep测试通过");

// ==================== app模块测试 ====================
console.log("\n[app模块] 测试开始");

// app.currentPackage
var pkg = app.currentPackage();
console.log("当前包名: " + pkg);

// app.currentActivity
var activity = app.currentActivity();
console.log("当前Activity: " + activity);

// app.getBrowserPackage
var browser = app.getBrowserPackage();
console.log("默认浏览器包名: " + browser);

// app.isInstalled
var isInstalled = app.isInstalled("com.android.settings");
console.log("设置应用是否安装: " + isInstalled);

// app.launch (需要实际包名)
// var launched = app.launch("com.android.settings", 0);
// console.log("启动设置应用: " + launched);

console.log("[app模块] 测试完成");

// ==================== device模块测试 ====================
console.log("\n[device模块] 测试开始");

// 设备信息
console.log("SDK版本: " + device.sdkInt());
console.log("CPU架构: " + device.cpuAbi());
console.log("构建ID: " + device.buildId());
console.log("主板型号: " + device.broad());
console.log("品牌: " + device.brand());
console.log("设备名: " + device.device());
console.log("型号: " + device.model());
console.log("产品名: " + device.product());
console.log("Bootloader版本: " + device.bootloader());
console.log("硬件名: " + device.hardware());
console.log("指纹: " + device.fingerprint());
console.log("序列号: " + device.serial());
console.log("增量版本: " + device.incremental());
console.log("Android版本: " + device.release());
console.log("基础OS: " + device.baseOS());
console.log("安全补丁: " + device.securityPatch());
console.log("代号: " + device.codename());

// 设备标识
console.log("IMEI: " + device.getImei());
console.log("Android ID: " + device.getAndroidId());
console.log("WiFi MAC: " + device.getWifiMac());
console.log("以太网MAC: " + device.getWlanMac());
console.log("IP地址: " + device.getIp());

// 音量和亮度
console.log("亮度: " + device.getBrightness());
console.log("亮度模式: " + device.getBrightnessMode());
console.log("媒体音量: " + device.getMusicVolume());
console.log("通知音量: " + device.getNotificationVolume());
console.log("闹钟音量: " + device.getAlarmVolume());
console.log("媒体最大音量: " + device.getMusicMaxVolume());
console.log("通知最大音量: " + device.getNotificationMaxVolume());
console.log("闹钟最大音量: " + device.getAlarmMaxVolume());

// 电池和内存
console.log("电量: " + device.getBattery() + "%");
console.log("电池状态: " + device.getBatteryStatus());
console.log("总内存: " + device.getTotalMem());
console.log("可用内存: " + device.getAvailMem());

// 屏幕状态
console.log("屏幕亮起: " + device.isScreenOn());
console.log("屏幕解锁: " + device.isScreenUnlock());

console.log("[device模块] 测试完成");

// ==================== motion模块测试 ====================
console.log("\n[motion模块] 测试开始");

// 触摸操作测试 (需要实际坐标)
console.log("touchDown/touchMove/touchUp - 需要实际坐标，跳过");
// touchDown(100, 100, 0, 0);
// touchMove(200, 200, 0, 0);
// touchUp(200, 200, 0, 0);

// 点击测试
console.log("click - 需要实际坐标，跳过");
// click(100, 100, 0, 0);

// 长按测试
console.log("longClick - 需要实际坐标，跳过");
// longClick(100, 100, 500, 0, 0);

// 滑动测试
console.log("swipe - 需要实际坐标，跳过");
// swipe(100, 100, 300, 300, 500, 0, 0);

// 按键测试
console.log("home/back/recents - 需要实际设备，跳过");
// home(0);
// back(0);
// recents(0);

// 音量键测试
console.log("volumeUp/volumeDown - 需要实际设备，跳过");
// volumeUp(0);
// volumeDown(0);

console.log("[motion模块] 测试完成");

// ==================== files模块测试 ====================
console.log("\n[files模块] 测试开始");

// 创建测试文件
var testPath = "/data/local/tmp/test_autogo.txt";
var testContent = "Hello AutoGo JavaScript Engine!";

// 写入文件
files.write(testPath, testContent);
console.log("写入文件: " + testPath);

// 检查文件是否存在
var exists = files.exists(testPath);
console.log("文件存在: " + exists);

// 读取文件
var content = files.read(testPath);
console.log("文件内容: " + content);

// 检查是否是文件
var isFile = files.isFile(testPath);
console.log("是文件: " + isFile);

// 获取文件名
var name = files.getName(testPath);
console.log("文件名: " + name);

// 获取不含扩展名的文件名
var nameWithoutExt = files.getNameWithoutExtension(testPath);
console.log("不含扩展名的文件名: " + nameWithoutExt);

// 获取扩展名
var ext = files.getExtension(testPath);
console.log("扩展名: " + ext);

// 追加内容
files.append(testPath, "\nAppended line");
console.log("追加内容成功");

// 删除文件
var removed = files.remove(testPath);
console.log("删除文件: " + removed);

console.log("[files模块] 测试完成");

// ==================== storages模块测试 ====================
console.log("\n[storages模块] 测试开始");

var tableName = "test_table";
var key = "test_key";
var value = "test_value";

// 写入键值对
storages.put(tableName, key, value);
console.log("写入键值对: " + key + " = " + value);

// 读取键值
var readValue = storages.get(tableName, key);
console.log("读取键值: " + readValue);

// 检查键是否存在
var contains = storages.contains(tableName, key);
console.log("键存在: " + contains);

// 获取所有键值对
var allData = storages.getAll(tableName);
console.log("所有键值对: " + JSON.stringify(allData));

// 删除键
storages.remove(tableName, key);
console.log("删除键: " + key);

// 清空表
storages.clear(tableName);
console.log("清空表: " + tableName);

console.log("[storages模块] 测试完成");

// ==================== system模块测试 ====================
console.log("\n[system模块] 测试开始");

// 获取进程ID
var pid = system.getPid("init");
console.log("init进程PID: " + pid);

// 获取内存使用
if (pid > 0) {
    var memUsage = system.getMemoryUsage(pid);
    console.log("内存使用: " + memUsage);
}

console.log("[system模块] 测试完成");

// ==================== images模块测试 ====================
console.log("\n[images模块] 测试开始");

// 截屏测试 (需要实际设备)
console.log("captureScreen - 需要实际设备，跳过");
// var img = images.captureScreen(0, 0, 100, 100, 0);

// 像素颜色测试
console.log("pixel - 需要实际设备，跳过");
// var color = images.pixel(100, 100, 0);

// 颜色比较测试
console.log("cmpColor - 需要实际设备，跳过");
// var cmpResult = images.cmpColor(100, 100, "#FF0000", 0.9, 0);

// 查找颜色测试
console.log("findColor - 需要实际设备，跳过");
// var result = images.findColor(0, 0, 500, 500, "#FF0000", 0.9, 0, 0);

console.log("[images模块] 测试完成");

// ==================== http模块测试 ====================
console.log("\n[http模块] 测试开始");

// GET请求测试
console.log("http.get - 测试HTTP GET请求");
var result = http.get("https://httpbin.org/get", 5000);
console.log("响应码: " + result.code);
if (result.data) {
    console.log("响应数据长度: " + result.data.length);
}

console.log("[http模块] 测试完成");

// ==================== console模块测试 ====================
console.log("\n[console模块] 测试开始");

// 创建控制台
var consoleObj = console.new();
console.log("创建控制台成功");

// 设置窗口大小
consoleObj.setWindowSize(400, 300);
console.log("设置窗口大小: 400x300");

// 设置窗口位置
consoleObj.setWindowPosition(100, 100);
console.log("设置窗口位置: 100, 100");

// 设置窗口颜色
consoleObj.setWindowColor("#000000");
console.log("设置窗口颜色: #000000");

// 设置文本颜色
consoleObj.setTextColor("#FFFFFF");
console.log("设置文本颜色: #FFFFFF");

// 设置文本大小
consoleObj.setTextSize(14);
console.log("设置文本大小: 14");

// 打印内容
consoleObj.println("Hello from JavaScript!");
console.log("打印内容成功");

// 显示控制台
consoleObj.show();
console.log("显示控制台");

// 检查可见性
var visible = consoleObj.isVisible();
console.log("控制台可见: " + visible);

// 隐藏控制台
consoleObj.hide();
console.log("隐藏控制台");

// 销毁控制台
consoleObj.destroy();
console.log("销毁控制台");

console.log("[console模块] 测试完成");

// ==================== imgui模块测试 ====================
console.log("\n[imgui模块] 测试开始");

// 测试版本
console.log("ImGui版本: " + imgui.version);

// 测试常量
console.log("WindowFlags.None = " + imgui.WindowFlags.None);
console.log("Col.Text = " + imgui.Col.Text);
console.log("MouseButton.Left = " + imgui.MouseButton.Left);
console.log("Key.Tab = " + imgui.Key.Tab);

console.log("[imgui模块] 测试完成");

// ==================== 方法注册测试 ====================
console.log("\n[方法注册] 测试开始");

// 列出所有方法
var methods = listMethods();
console.log("已注册方法数量: " + methods.length);

// 注册自定义方法
registerMethod("custom.test", "自定义测试方法", true);
console.log("注册自定义方法: custom.test");

// 再次列出方法
methods = listMethods();
console.log("注册后方法数量: " + methods.length);

console.log("[方法注册] 测试完成");

console.log("\n========================================");
console.log("AutoGo JavaScript 引擎全量方法测试完成");
console.log("========================================");
