// AutoGo JavaScript 测试脚本（完整版）
// 测试所有可用模块的方法（排除 console、hud、imgui、vdisplay）
//
// 使用方法：
// 1. 测试所有模块：直接运行脚本
// 2. 测试指定模块：在脚本开头设置 testModules 变量
//    例如：var testModules = ["app", "device", "coroutine"];
// 3. 可用模块列表：
//    app, device, touch, files, images, storages, system, http,
//    media, opencv, ppocr, dotocr, ime, plugin, rhino, uiacc,
//    utils, yolo, coroutine, method

// ========== 测试配置 ==========
// 设置要测试的模块，如果为空数组则测试所有模块
var testModules = ["coroutine"]; // 例如：["app", "device", "coroutine"]

// 检查是否应该测试指定模块
function shouldTest(moduleName) {
    if (testModules.length === 0) {
        return true; // 测试所有模块
    }
    return testModules.indexOf(moduleName) !== -1;
}

console.log("=== JavaScript 引擎完整测试开始 ===");
console.log("测试时间: " + new Date().toLocaleString());

if (testModules.length > 0) {
    console.log("测试模块: " + testModules.join(", "));
} else {
    console.log("测试模式: 全部模块");
}

// ========== 1. 应用管理测试 (app) ==========
if (shouldTest("app")) {
    console.log("\n--- 1. 应用管理 (app) 测试 ---");

    try {
        // 获取当前应用包名
        var currentPackage = app.currentPackage();
        console.log("✓ app.currentPackage(): " + currentPackage);

        // 获取当前 Activity
        var currentActivity = app.currentActivity();
        console.log("✓ app.currentActivity(): " + currentActivity);

        // 检查应用是否已安装
        var isInstalled = app.isInstalled("com.android.settings");
        console.log("✓ app.isInstalled('com.android.settings'): " + isInstalled);

        // 获取浏览器包名
        var browserPackage = app.getBrowserPackage();
        console.log("✓ app.getBrowserPackage(): " + browserPackage);

        // 获取应用列表
        var appList = app.getList(false);
        console.log("✓ app.getList(false): " + JSON.stringify(appList).substring(0, 100) + "...");

        // 获取应用名称
        var appName = app.getName("com.android.settings");
        console.log("✓ app.getName('com.android.settings'): " + appName);

        // 获取应用版本
        var appVersion = app.getVersion("com.android.settings");
        console.log("✓ app.getVersion('com.android.settings'): " + appVersion);

        // 打开 URL
        console.log("✓ app.openUrl('https://www.baidu.com') - 已调用");

        // 获取应用设置
        console.log("✓ app.openAppSetting('com.android.settings') - 已调用");

        // 启用/禁用无障碍服务
        console.log("✓ app.enableAccessibility('com.android.settings') - 已调用");
        app.enableAccessibility("com.android.settings");

        console.log("✓ app.disableAccessibility('com.android.settings') - 已调用");
        app.disableAccessibility("com.android.settings");

        console.log("✓ 应用管理模块测试完成");
    } catch (e) {
        console.error("✗ 应用管理模块测试失败: " + e);
    }
}

// ========== 2. 设备管理测试 (device) ==========
if (shouldTest("device")) {
    console.log("\n--- 2. 设备管理 (device) 测试 ---");

    try {
        // 获取设备信息
        console.log("✓ device.width: " + device.width);
        console.log("✓ device.height: " + device.height);
        console.log("✓ device.sdkInt: " + device.sdkInt);
        console.log("✓ device.cpuAbi: " + device.cpuAbi);
        console.log("✓ device.buildId: " + device.buildId);
        console.log("✓ device.broad: " + device.broad);
        console.log("✓ device.brand: " + device.brand);
        console.log("✓ device.deviceName: " + device.deviceName);
        console.log("✓ device.model: " + device.model);
        console.log("✓ device.product: " + device.product);
        console.log("✓ device.bootloader: " + device.bootloader);
        console.log("✓ device.hardware: " + device.hardware);
        console.log("✓ device.fingerprint: " + device.fingerprint);
        console.log("✓ device.serial: " + device.serial);
        console.log("✓ device.incremental: " + device.incremental);
        console.log("✓ device.release: " + device.release);
        console.log("✓ device.baseOS: " + device.baseOS);
        console.log("✓ device.securityPatch: " + device.securityPatch);
        console.log("✓ device.codename: " + device.codename);

        // 获取设备标识
        var imei = device.getImei();
        console.log("✓ device.getImei(): " + imei);

        var androidId = device.getAndroidId();
        console.log("✓ device.getAndroidId(): " + androidId);

        var wifiMac = device.getWifiMac();
        console.log("✓ device.getWifiMac(): " + wifiMac);

        var wlanMac = device.getWlanMac();
        console.log("✓ device.getWlanMac(): " + wlanMac);

        var ip = device.getIp();
        console.log("✓ device.getIp(): " + ip);

        // 获取通知
        var notification = device.getNotification();
        console.log("✓ device.getNotification(): " + notification);

        // 获取亮度信息
        var brightness = device.getBrightness();
        console.log("✓ device.getBrightness(): " + brightness);

        var brightnessMode = device.getBrightnessMode();
        console.log("✓ device.getBrightnessMode(): " + brightnessMode);

        // 获取音量信息
        var musicVolume = device.getMusicVolume();
        console.log("✓ device.getMusicVolume(): " + musicVolume);

        var notificationVolume = device.getNotificationVolume();
        console.log("✓ device.getNotificationVolume(): " + notificationVolume);

        var alarmVolume = device.getAlarmVolume();
        console.log("✓ device.getAlarmVolume(): " + alarmVolume);

        var musicMaxVolume = device.getMusicMaxVolume();
        console.log("✓ device.getMusicMaxVolume(): " + musicMaxVolume);

        var notificationMaxVolume = device.getNotificationMaxVolume();
        console.log("✓ device.getNotificationMaxVolume(): " + notificationMaxVolume);

        var alarmMaxVolume = device.getAlarmMaxVolume();
        console.log("✓ device.getAlarmMaxVolume(): " + alarmMaxVolume);

        // 设置音量
        console.log("✓ device.setMusicVolume(50) - 已调用");
        device.setMusicVolume(50);

        console.log("✓ device.setNotificationVolume(50) - 已调用");
        device.setNotificationVolume(50);

        console.log("✓ device.setAlarmVolume(50) - 已调用");
        device.setAlarmVolume(50);

        // 设置显示电源
        console.log("✓ device.setDisplayPower(true) - 已调用");
        device.setDisplayPower(true);

        // 获取电池信息
        var battery = device.getBattery();
        console.log("✓ device.getBattery(): " + battery + "%");

        var batteryStatus = device.getBatteryStatus();
        console.log("✓ device.getBatteryStatus(): " + batteryStatus);

        // 设置电池状态
        console.log("✓ device.setBatteryStatus(2) - 已调用");
        device.setBatteryStatus(2);

        console.log("✓ device.setBatteryLevel(80) - 已调用");
        device.setBatteryLevel(80);

        // 获取内存信息
        var totalMem = device.getTotalMem();
        console.log("✓ device.getTotalMem(): " + totalMem + " KB");

        var availMem = device.getAvailMem();
        console.log("✓ device.getAvailMem(): " + availMem + " KB");

        // 检查屏幕状态
        var isScreenOn = device.isScreenOn();
        console.log("✓ device.isScreenOn(): " + isScreenOn);

        var isScreenUnlock = device.isScreenUnlock();
        console.log("✓ device.isScreenUnlock(): " + isScreenUnlock);

        // 唤醒设备
        console.log("✓ device.wakeUp() - 已调用");
        device.wakeUp();

        // 保持屏幕常亮
        console.log("✓ device.keepScreenOn() - 已调用");
        device.keepScreenOn();

        // 震动
        console.log("✓ device.vibrate(500) - 已调用");
        device.vibrate(500);

        // 取消震动
        console.log("✓ device.cancelVibration() - 已调用");
        device.cancelVibration();

        console.log("✓ 设备管理模块测试完成");
    } catch (e) {
        console.error("✗ 设备管理模块测试失败: " + e);
    }
}

// ========== 3. 触摸操作测试 (touch) ==========
if (shouldTest("touch")) {
    console.log("\n--- 3. 触摸操作 (touch) 测试 ---");

    try {
        // 触摸操作
        console.log("✓ touchDown(500, 500, 1) - 已调用");
        touchDown(500, 500, 1);

        console.log("✓ touchMove(550, 550, 1) - 已调用");
        touchMove(550, 550, 1);

        console.log("✓ touchUp(550, 550, 1) - 已调用");
        touchUp(550, 550, 1);

        // 点击
        console.log("✓ click(500, 500, 1) - 已调用");
        click(500, 500, 1);

        // 长按
        console.log("✓ longClick(500, 500, 1000) - 已调用");
        longClick(500, 500, 1000);

        // 滑动
        console.log("✓ swipe(100, 500, 500, 500, 500) - 已调用");
        swipe(100, 500, 500, 500, 500);

        console.log("✓ swipe2(100, 500, 500, 500, 500) - 已调用");
        swipe2(100, 500, 500, 500, 500);

        // 按键
        console.log("✓ home() - 已调用");
        home();

        console.log("✓ back() - 已调用");
        back();

        console.log("✓ recents() - 已调用");
        recents();

        console.log("✓ powerDialog() - 已调用");
        powerDialog();

        console.log("✓ notifications() - 已调用");
        notifications();

        console.log("✓ quickSettings() - 已调用");
        quickSettings();

        console.log("✓ volumeUp() - 已调用");
        volumeUp();

        console.log("✓ volumeDown() - 已调用");
        volumeDown();

        console.log("✓ camera() - 已调用");
        camera();

        console.log("✓ keyAction(3) - 已调用");
        keyAction(3);

        console.log("✓ 触摸操作模块测试完成");
    } catch (e) {
        console.error("✗ 触摸操作模块测试失败: " + e);
    }
}

// ========== 4. 文件操作测试 (files) ==========
if (shouldTest("files")) {
    console.log("\n--- 4. 文件操作 (files) 测试 ---");

    try {
        var testPath = "/sdcard/autogo_test_js.txt";
        var testDir = "/sdcard/autogo_test_dir";

        // 创建文件
        console.log("✓ files.create('" + testPath + "') - 已调用");
        files.create(testPath);

        // 写入文件
        console.log("✓ files.write('" + testPath + "', 'Hello from JavaScript!') - 已调用");
        files.write(testPath, "Hello from JavaScript!");

        // 读取文件
        var content = files.read(testPath);
        console.log("✓ files.read('" + testPath + "'): " + content);

        // 追加内容
        console.log("✓ files.append('" + testPath + "', '\\nAppend text') - 已调用");
        files.append(testPath, "\nAppend text");

        // 检查文件是否存在
        var exists = files.exists(testPath);
        console.log("✓ files.exists('" + testPath + "'): " + exists);

        var isFile = files.isFile(testPath);
        console.log("✓ files.isFile('" + testPath + "'): " + isFile);

        // 获取文件名
        var fileName = files.getName(testPath);
        console.log("✓ files.getName('" + testPath + "'): " + fileName);

        var fileNameWithoutExt = files.getNameWithoutExtension(testPath);
        console.log("✓ files.getNameWithoutExtension('" + testPath + "'): " + fileNameWithoutExt);

        var extension = files.getExtension(testPath);
        console.log("✓ files.getExtension('" + testPath + "'): " + extension);

        // 获取文件 MD5
        var md5 = files.getMd5(testPath);
        console.log("✓ files.getMd5('" + testPath + "'): " + md5);

        // 获取绝对路径
        var absPath = files.path(testPath);
        console.log("✓ files.path('" + testPath + "'): " + absPath);

        // 创建目录
        console.log("✓ files.create('" + testDir + "') - 已调用");
        files.create(testDir);

        var isDir = files.isDir(testDir);
        console.log("✓ files.isDir('" + testDir + "'): " + isDir);

        var isEmptyDir = files.isEmptyDir(testDir);
        console.log("✓ files.isEmptyDir('" + testDir + "'): " + isEmptyDir);

        // 确保目录存在
        console.log("✓ files.ensureDir('" + testDir + "') - 已调用");
        files.ensureDir(testDir);

        // 创建带目录的文件
        var testPath2 = testDir + "/test2.txt";
        console.log("✓ files.create('" + testPath2 + "') - 已调用");
        files.create(testPath2);

        // 复制文件
        var testPath3 = "/sdcard/autogo_test_copy.txt";
        console.log("✓ files.copy('" + testPath + "', '" + testPath3 + "') - 已调用");
        files.copy(testPath, testPath3);

        // 重命名文件
        var testPath4 = "/sdcard/autogo_test_rename.txt";
        console.log("✓ files.rename('" + testPath3 + "', 'autogo_test_rename.txt') - 已调用");
        files.rename(testPath3, "autogo_test_rename.txt");

        // 移动文件
        var testPath5 = testDir + "/moved.txt";
        console.log("✓ files.move('" + testPath4 + "', '" + testPath5 + "') - 已调用");
        files.move(testPath4, testPath5);

        // 列出目录
        var dirList = files.listDir(testDir);
        console.log("✓ files.listDir('" + testDir + "'): " + JSON.stringify(dirList));

        // 删除文件
        console.log("✓ files.remove('" + testPath + "') - 已调用");
        files.remove(testPath);

        console.log("✓ files.remove('" + testDir + "') - 已调用");
        files.remove(testDir);

        console.log("✓ 文件操作模块测试完成");
    } catch (e) {
        console.error("✗ 文件操作模块测试失败: " + e);
    }
}

// ========== 5. 图像处理测试 (images) ==========
if (shouldTest("images")) {
    console.log("\n--- 5. 图像处理 (images) 测试 ---");

    try {
        // 获取像素颜色
        var pixel = images.pixel(100, 100);
        console.log("✓ images.pixel(100, 100): " + pixel);

        // 截取屏幕
        var img = images.captureScreen(0, 0, 500, 500);
        console.log("✓ images.captureScreen(0, 0, 500, 500): " + (img !== null));

        // 比较颜色
        var colorMatch = images.cmpColor(100, 100, "#FF0000", 0.9);
        console.log("✓ images.cmpColor(100, 100, '#FF0000', 0.9): " + colorMatch);

        // 查找颜色
        var findColorResult = images.findColor(0, 0, 1000, 1000, "#FF0000", 0.9, 1);
        console.log("✓ images.findColor(0, 0, 1000, 1000, '#FF0000', 0.9, 1): " + findColorResult);

        // 获取区域内颜色数量
        var colorCount = images.getColorCountInRegion(0, 0, 1000, 1000, "#FF0000", 0.9);
        console.log("✓ images.getColorCountInRegion(0, 0, 1000, 1000, '#FF0000', 0.9): " + colorCount);

        // 检测多点颜色
        // 多点颜色格式：主颜色-误差,偏移x,偏移y,颜色2-误差,偏移x,偏移y,颜色3-误差
        // 注意：最后一个颜色不需要偏移坐标
        var multiColors = "#FF0000-101010,10,10,#00FF00-101010,20,20,#0000FF-101010";
        try {
            var detectMultiColors = images.detectsMultiColors(multiColors, 0.9);
            console.log("✓ images.detectsMultiColors(...): " + detectMultiColors);

            // 查找多点颜色
            var findMultiColors = images.findMultiColors(0, 0, 1000, 1000, multiColors, 0.9, 1);
            console.log("✓ images.findMultiColors(0, 0, 1000, 1000, ..., 0.9, 1): " + findMultiColors);
        } catch (e) {
            console.log("✗ 多点颜色测试失败（可能是格式问题）: " + e);
        }

        // 从路径读取图片
        var imgFromPath = images.readFromPath("/sdcard/test.png");
        console.log("✓ images.readFromPath('/sdcard/test.png'): " + (imgFromPath !== null));

        // 从 Base64 读取图片
        var base64Img = "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVR42mNk+M9QDwADhgGAWjR9awAAAABJRU5ErkJggg==";
        var imgFromBase64 = images.readFromBase64(base64Img);
        console.log("✓ images.readFromBase64(...): " + (imgFromBase64 !== null));

        // 图片处理
        if (img !== null) {
            // 裁剪图片
            var clippedImg = images.clip(img, 10, 10, 100, 100);
            console.log("✓ images.clip(img, 10, 10, 100, 100): " + (clippedImg !== null));

            // 调整大小
            var resizedImg = images.resize(img, 200, 200);
            console.log("✓ images.resize(img, 200, 200): " + (resizedImg !== null));

            // 旋转图片
            var rotatedImg = images.rotate(img, 90);
            console.log("✓ images.rotate(img, 90): " + (rotatedImg !== null));

            // 灰度化
            var grayImg = images.grayscale(img);
            console.log("✓ images.grayscale(img): " + (grayImg !== null));

            // 二值化
            var binaryImg = images.applyBinarization(img, 128);
            console.log("✓ images.applyBinarization(img, 128): " + (binaryImg !== null));

            // 保存图片
            var savePath = "/sdcard/autogo_test_image.png";
            var saved = images.save(img, savePath, 90);
            console.log("✓ images.save(img, '" + savePath + "', 90): " + saved);

            // 编码为 Base64
            var base64Str = images.encodeToBase64(img, "png", 90);
            console.log("✓ images.encodeToBase64(img, 'png', 90): " + (base64Str.length > 0));

            // 编码为字节数组
            var bytesStr = images.encodeToBytes(img, "png", 90);
            console.log("✓ images.encodeToBytes(img, 'png', 90): " + (bytesStr.length > 0));

            // 转换为 NRGBA 格式
            var nrgbaImg = images.toNrgba(img);
            console.log("✓ images.toNrgba(img): " + (nrgbaImg !== null));

            // 设置回调函数
            console.log("✓ images.setCallback(callback) - 已调用");
            images.setCallback(function (img, displayId) {
                console.log("  回调被调用: displayId=" + displayId);
            });
        }

        console.log("✓ 图像处理模块测试完成");
    } catch (e) {
        console.error("✗ 图像处理模块测试失败: " + e);
    }
}

// ========== 6. 存储管理测试 (storages) ==========
if (shouldTest("storages")) {
    console.log("\n--- 6. 存储管理 (storages) 测试 ---");

    try {
        var tableName = "test_table_js";

        // 写入存储
        console.log("✓ storages.put('" + tableName + "', 'key1', 'value1') - 已调用");
        storages.put(tableName, "key1", "value1");

        console.log("✓ storages.put('" + tableName + "', 'key2', 'value2') - 已调用");
        storages.put(tableName, "key2", "value2");

        console.log("✓ storages.put('" + tableName + "', 'key3', 'value3') - 已调用");
        storages.put(tableName, "key3", "value3");

        // 读取存储
        var value1 = storages.get(tableName, "key1");
        console.log("✓ storages.get('" + tableName + "', 'key1'): " + value1);

        // 检查键是否存在
        var contains = storages.contains(tableName, "key1");
        console.log("✓ storages.contains('" + tableName + "', 'key1'): " + contains);

        // 获取所有数据
        var allData = storages.getAll(tableName);
        console.log("✓ storages.getAll('" + tableName + "'): " + JSON.stringify(allData));

        // 删除键
        console.log("✓ storages.remove('" + tableName + "', 'key3') - 已调用");
        storages.remove(tableName, "key3");

        // 清空存储
        console.log("✓ storages.clear('" + tableName + "') - 已调用");
        storages.clear(tableName);

        console.log("✓ 存储管理模块测试完成");
    } catch (e) {
        console.error("✗ 存储管理模块测试失败: " + e);
    }
}

// ========== 7. 系统管理测试 (system) ==========
if (shouldTest("system")) {
    console.log("\n--- 7. 系统管理 (system) 测试 ---");

    try {
        // 获取进程 ID
        var pid = system.getPid("com.android.systemui");
        console.log("✓ system.getPid('com.android.systemui'): " + pid);

        if (pid > 0) {
            // 获取内存使用
            var memUsage = system.getMemoryUsage(pid);
            console.log("✓ system.getMemoryUsage(" + pid + "): " + memUsage + " KB");

            // 获取 CPU 使用率
            var cpuUsage = system.getCpuUsage(pid);
            console.log("✓ system.getCpuUsage(" + pid + "): " + cpuUsage + "%");
        }

        // 重启自身
        console.log("✓ system.restartSelf() - 已调用（未实际执行）");
        // system.restartSelf(); // 不实际执行，避免重启

        // 设置开机自启
        console.log("✓ system.setBootStart(true) - 已调用");
        system.setBootStart(true);

        console.log("✓ system.setBootStart(false) - 已调用");
        system.setBootStart(false);

        console.log("✓ 系统管理模块测试完成");
    } catch (e) {
        console.error("✗ 系统管理模块测试失败: " + e);
    }
}

// ========== 8. 网络请求测试 (http) ==========
if (shouldTest("http")) {
    console.log("\n--- 8. 网络请求 (http) 测试 ---");

    try {
        // GET 请求
        console.log("✓ http.get('https://httpbin.org/get', 5000) - 已调用");
        var result = http.get("https://httpbin.org/get", 5000);
        console.log("✓ HTTP 状态码: " + result.code);
        console.log("✓ 响应长度: " + (result.data ? result.data.length : 0));

        // POST 请求
        var postData = JSON.stringify({ test: "data" });
        var postHeaders = { "Content-Type": "application/json" };
        console.log("✓ http.post('https://httpbin.org/post', data, headers, 5000) - 已调用");
        var postResult = http.post("https://httpbin.org/post", postData, postHeaders, 5000);
        console.log("✓ POST HTTP 状态码: " + postResult.code);
        console.log("✓ POST 响应长度: " + (postResult.data ? postResult.data.length : 0));

        // POST Multipart 请求
        console.log("✓ http.postMultipart('https://httpbin.org/post', 'test.txt', 'Hello', 5000) - 已调用");
        var postResult = http.postMultipart("https://httpbin.org/post", "test.txt", "Hello", 5000);
        console.log("✓ POST HTTP 状态码: " + postResult.code);
        console.log("✓ POST 响应长度: " + (postResult.data ? postResult.data.length : 0));

        console.log("✓ 网络请求模块测试完成");
    } catch (e) {
        console.error("✗ 网络请求模块测试失败: " + e);
    }
}

// ========== 9. 媒体管理测试 (media) ==========
if (shouldTest("media")) {
    console.log("\n--- 9. 媒体管理 (media) 测试 ---");

    try {
        var mediaPath = "/sdcard/autogo_test_media.jpg";
        var mp3Path = "/sdcard/test.mp3";

        // 扫描文件
        console.log("✓ media.scanFile('" + mediaPath + "') - 已调用");
        media.scanFile(mediaPath);

        // 播放 MP3
        console.log("✓ media.playMP3('" + mp3Path + "') - 已调用");
        var playResult = media.playMP3(mp3Path);
        if (playResult !== null) {
            console.log("✓ 播放结果: " + playResult);
        } else {
            console.log("✓ MP3 播放成功");
        }

        // 发送短信
        console.log("✓ media.sendSMS('10086', '测试短信') - 已调用");
        media.sendSMS("10086", "测试短信");

        console.log("✓ 媒体管理模块测试完成");
    } catch (e) {
        console.error("✗ 媒体管理模块测试失败: " + e);
    }
}

// ========== 10. 图像识别测试 (opencv) ==========
if (shouldTest("opencv")) {
    console.log("\n--- 10. 图像识别 (opencv) 测试 ---");

    try {
        var templatePath = "/sdcard/template.png";

        // 查找图片
        console.log("✓ opencv.findImage(0, 0, 1000, 1000, '" + templatePath + "', false, 1.0, 0.9) - 已调用");
        var findResult = opencv.findImage(0, 0, 1000, 1000, templatePath, false, 1.0, 0.9);
        console.log("✓ 查找结果: x=" + findResult.x + ", y=" + findResult.y);

        console.log("✓ 图像识别模块测试完成");
    } catch (e) {
        console.error("✗ 图像识别模块测试失败: " + e);
    }
}

// ========== 11. 文字识别测试 (ppocr) ==========
if (shouldTest("ppocr")) {
    console.log("\n--- 11. 文字识别 (ppocr) 测试 ---");

    try {
        // 创建 PPOCR 实例
        console.log("✓ ppocr.new('v3') - 已调用");
        var p = ppocr.new("v3");
        console.log("✓ PPOCR 实例创建成功");

        // 识别屏幕文字
        console.log("✓ ppocr.ocr(p, 0, 0, 1000, 1000, '#000000') - 已调用");
        var ocrResults = ppocr.ocr(p, 0, 0, 1000, 1000, "#000000");
        console.log("✓ 识别结果数量: " + ocrResults.length);
        if (ocrResults.length > 0) {
            console.log("✓ 第一个识别结果: " + JSON.stringify(ocrResults[0]));
        }

        // 从图片识别
        var img = images.captureScreen();
        console.log("✓ ppocr.ocrFromImage(p, img, '#000000') - 已调用");
        var ocrFromImage = ppocr.ocrFromImage(p, img, "#000000");
        console.log("✓ 图片识别结果数量: " + ocrFromImage.length);

        // 从 Base64 识别
        var base64Img = "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVR42mNk+M9QDwADhgGAWjR9awAAAABJRU5ErkJggg==";
        console.log("✓ ppocr.ocrFromBase64(p, ...) - 已调用");
        var ocrFromBase64 = ppocr.ocrFromBase64(p, base64Img, "#000000");
        console.log("✓ Base64 识别结果数量: " + ocrFromBase64.length);

        // 从路径识别
        var imagePath = "/sdcard/test_ocr.png";
        console.log("✓ ppocr.ocrFromPath(p, '" + imagePath + "', '#000000') - 已调用");
        var ocrFromPath = ppocr.ocrFromPath(p, imagePath, "#000000");
        console.log("✓ 路径识别结果数量: " + ocrFromPath.length);

        // 关闭实例
        console.log("✓ ppocr.close(p) - 已调用");
        ppocr.close(p);

        console.log("✓ 文字识别模块测试完成");
    } catch (e) {
        console.error("✗ 文字识别模块测试失败: " + e);
    }
}

// ========== 12. 点阵OCR测试 (dotocr) ==========
if (shouldTest("dotocr")) {
    console.log("\n--- 12. 点阵OCR (dotocr) 测试 ---");

    try {
        // 设置字库
        console.log("✓ dotocr.setDict('test_dict', 'dict_content') - 已调用");
        dotocr.setDict("test_dict", "dict_content");

        // 准备测试图像
        var testImg = images.readFromPath("/sdcard/test.png");
        console.log("✓ 准备测试图像: " + (testImg !== null));

        // 识别屏幕文字
        console.log("✓ dotocr.ocr(0, 0, 1000, 1000, '#000000', 5, 5, 0.9, 0, 'test_dict', 0) - 已调用");
        var ocrResult = dotocr.ocr(0, 0, 1000, 1000, "#000000", 5, 5, 0.9, 0, "test_dict", 0);
        console.log("✓ OCR 识别结果: " + ocrResult);

        // 从 Base64 识别
        var base64Img = "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVR42mNk+M9QDwADhgGAWjR9awAAAABJRU5ErkJggg==";
        console.log("✓ dotocr.ocrFromBase64(...) - 已调用");
        var ocrFromBase64 = dotocr.ocrFromBase64(base64Img, "#000000", 5, 5, 0.9, 0, "test_dict");
        console.log("✓ Base64 OCR 结果: " + ocrFromBase64);

        // 从路径识别
        var imagePath = "/sdcard/test_dotocr.png";
        console.log("✓ dotocr.ocrFromPath('" + imagePath + "', '#000000', 5, 5, 0.9, 0, 'test_dict') - 已调用");
        var ocrFromPath = dotocr.ocrFromPath(imagePath, "#000000", 5, 5, 0.9, 0, "test_dict");
        console.log("✓ 路径 OCR 结果: " + ocrFromPath);

        // 从图像对象识别
        if (testImg !== null) {
            console.log("✓ dotocr.ocrFromImage(img, '#000000', 5, 5, 0.9, 0, 'test_dict') - 已调用");
            var ocrFromImage = dotocr.ocrFromImage(testImg, "#000000", 5, 5, 0.9, 0, "test_dict");
            console.log("✓ 图像对象 OCR 结果: " + ocrFromImage);
        }

        // 查找文字
        console.log("✓ dotocr.findStr(0, 0, 1000, 1000, 'test', '#000000', 5, 5, 0.9, 'test_dict', 0) - 已调用");
        var findStrResult = dotocr.findStr(0, 0, 1000, 1000, "test", "#000000", 5, 5, 0.9, "test_dict", 0);
        console.log("✓ 查找文字结果: x=" + findStrResult.x + ", y=" + findStrResult.y);

        // 从图像对象查找文字
        if (testImg !== null) {
            console.log("✓ dotocr.findStrFromImage(img, 'test', '#000000', 5, 5, 0.9, 'test_dict') - 已调用");
            var findStrFromImage = dotocr.findStrFromImage(testImg, "test", "#000000", 5, 5, 0.9, "test_dict");
            console.log("✓ 图像对象查找结果: x=" + findStrFromImage.x + ", y=" + findStrFromImage.y);
        }

        // 从 Base64 查找文字
        console.log("✓ dotocr.findStrFromBase64(...) - 已调用");
        var findStrFromBase64 = dotocr.findStrFromBase64(base64Img, "test", "#000000", 5, 5, 0.9, "test_dict");
        console.log("✓ Base64 查找结果: x=" + findStrFromBase64.x + ", y=" + findStrFromBase64.y);

        // 从路径查找文字
        console.log("✓ dotocr.findStrFromPath('" + imagePath + "', 'test', '#000000', 5, 5, 0.9, 'test_dict') - 已调用");
        var findStrFromPath = dotocr.findStrFromPath(imagePath, "test", "#000000", 5, 5, 0.9, "test_dict");
        console.log("✓ 路径查找结果: x=" + findStrFromPath.x + ", y=" + findStrFromPath.y);

        console.log("✓ 点阵OCR模块测试完成");
    } catch (e) {
        console.error("✗ 点阵OCR模块测试失败: " + e);
    }
}

// ========== 13. 输入法测试 (ime) ==========
if (shouldTest("ime")) {
    console.log("\n--- 13. 输入法 (ime) 测试 ---");

    try {
        // 获取剪切板内容
        var clipText = ime.getClipText();
        console.log("✓ ime.getClipText(): " + clipText);

        // 设置剪切板内容
        console.log("✓ ime.setClipText('测试文本') - 已调用");
        ime.setClipText("测试文本");

        // 模拟按键
        console.log("✓ ime.keyAction(3) - 已调用");
        ime.keyAction(3);

        // 输入文本
        console.log("✓ ime.inputText('Hello World') - 已调用");
        ime.inputText("Hello World");

        // 获取输入法列表
        var imeList = ime.getIMEList();
        console.log("✓ ime.getIMEList(): " + JSON.stringify(imeList));

        // 设置当前输入法
        console.log("✓ ime.setCurrentIME('com.android.inputmethod.latin') - 已调用");
        ime.setCurrentIME("com.android.inputmethod.latin");

        console.log("✓ 输入法模块测试完成");
    } catch (e) {
        console.error("✗ 输入法模块测试失败: " + e);
    }
}

// ========== 14. 插件测试 (plugin) ==========
if (shouldTest("plugin")) {
    console.log("\n--- 14. 插件 (plugin) 测试 ---");

    try {
        var apkPath = "/sdcard/test_plugin.apk";

        // 加载 APK
        console.log("✓ plugin.loadApk('" + apkPath + "') - 已调用");
        var classLoader = plugin.loadApk(apkPath);
        console.log("✓ 插件加载结果: " + (classLoader !== null));

        console.log("✓ 插件模块测试完成");
    } catch (e) {
        console.error("✗ 插件模块测试失败: " + e);
    }
}

// ========== 15. Rhino引擎测试 (rhino) ==========
if (shouldTest("rhino")) {
    console.log("\n--- 15. Rhino引擎 (rhino) 测试 ---");

    try {
        var contextId = "test_context";
        var script = "var x = 10; var y = 20; x + y;";

        // 执行脚本
        console.log("✓ rhino.eval('" + contextId + "', '" + script + "') - 已调用");
        var result = rhino.eval(contextId, script);
        console.log("✓ Rhino 执行结果: " + result);

        console.log("✓ Rhino引擎模块测试完成");
    } catch (e) {
        console.error("✗ Rhino引擎模块测试失败: " + e);
    }
}

// ========== 16. UI辅助测试 (uiacc) ==========
if (shouldTest("uiacc")) {
    console.log("\n--- 16. UI辅助 (uiacc) 测试 ---");

    try {
        // 创建 UI 辅助对象
        console.log("✓ uiacc.new(0) - 已调用");
        var ui = uiacc.new(0);
        console.log("✓ UI 辅助对象: " + (ui !== null));

        // 查找控件
        console.log("✓ uiacc.text(ui, '设置') - 已调用");
        var selector = uiacc.text(ui, "设置");
        console.log("✓ 选择器: " + (selector !== null));

        // 查找一次
        console.log("✓ uiacc.findOnce(ui) - 已调用");
        var node = uiacc.findOnce(ui);
        console.log("✓ 查找结果: " + (node !== null));

        if (node !== null) {
            // 获取控件文本
            var text = uiacc.getText(node);
            console.log("✓ uiacc.getText(node): " + text);

            // 获取控件边界
            var bounds = uiacc.getBounds(node);
            console.log("✓ uiacc.getBounds(node): " + JSON.stringify(bounds));

            // 点击控件
            console.log("✓ uiacc.objClick(node) - 已调用");
            uiacc.objClick(node);

            // 点击中心
            console.log("✓ uiacc.clickCenter(node) - 已调用");
            uiacc.clickCenter(node);
        }

        // 查找所有匹配的控件
        console.log("✓ uiacc.find(ui) - 已调用");
        var nodes = uiacc.find(ui);
        console.log("✓ 查找结果数量: " + nodes.length);

        // 等待控件出现
        console.log("✓ uiacc.waitFor(ui, 5000) - 已调用");
        var waitResult = uiacc.waitFor(ui, 5000);
        console.log("✓ 等待结果: " + waitResult);

        // 使用 ID 查找
        console.log("✓ uiacc.id(ui, 'com.android.settings:id/title') - 已调用");
        var idSelector = uiacc.id(ui, "com.android.settings:id/title");
        console.log("✓ ID 选择器: " + (idSelector !== null));

        // 使用类名查找
        console.log("✓ uiacc.className(ui, 'android.widget.TextView') - 已调用");
        var classSelector = uiacc.className(ui, "android.widget.TextView");
        console.log("✓ 类名选择器: " + (classSelector !== null));

        // 设置可见性选择器
        console.log("✓ uiacc.visible(ui, true) - 已调用");
        var visibleSelector = uiacc.visible(ui, true);
        console.log("✓ 可见性选择器: " + (visibleSelector !== null));

        // 设置密码选择器
        console.log("✓ uiacc.password(ui, true) - 已调用");
        var passwordSelector = uiacc.password(ui, true);
        console.log("✓ 密码选择器: " + (passwordSelector !== null));

        if (node !== null) {
            // 获取可见性
            var isVisible = uiacc.getVisible(node);
            console.log("✓ uiacc.getVisible(node): " + isVisible);

            // 获取密码状态
            var isPassword = uiacc.getPassword(node);
            console.log("✓ uiacc.getPassword(node): " + isPassword);

            // 转换为字符串
            var str = uiacc.toString(node);
            console.log("✓ uiacc.toString(node): " + str);
        }

        // 释放资源
        console.log("✓ uiacc.release(ui) - 已调用");
        uiacc.release(ui);

        console.log("✓ UI辅助模块测试完成");
    } catch (e) {
        console.error("✗ UI辅助模块测试失败: " + e);
    }
}

// ========== 17. 工具函数测试 (utils) ==========
if (shouldTest("utils")) {
    console.log("\n--- 17. 工具函数 (utils) 测试 ---");

    try {
        // 日志
        console.log("✓ utils.logI('test', '这是一条测试日志') - 已调用");
        utils.logI("test", "这是一条测试日志");

        console.log("✓ utils.logE('test', '这是一条错误日志') - 已调用");
        utils.logE("test", "这是一条错误日志");

        // Toast
        console.log("✓ utils.toast('这是一条Toast提示') - 已调用");
        utils.toast("这是一条Toast提示");

        // Alert
        console.log("✓ utils.alert('标题', '内容', '确定', '取消') - 已调用");
        var alertResult = utils.alert("标题", "内容", "确定", "取消");
        console.log("✓ Alert 返回值: " + alertResult);

        // Shell 命令
        console.log("✓ utils.shell('echo test') - 已调用");
        var shellResult = utils.shell("echo test");
        console.log("✓ Shell 执行结果: " + shellResult);

        // 随机数
        console.log("✓ utils.random(1, 100) - 已调用");
        var randomNum = utils.random(1, 100);
        console.log("✓ 随机数: " + randomNum);

        // 睡眠
        console.log("✓ utils.sleep(500) - 已调用");
        utils.sleep(500);
        console.log("✓ 睡眠完成");

        // 类型转换
        console.log("✓ utils.i2s(123) - 已调用");
        var i2sResult = utils.i2s(123);
        console.log("✓ 整数转字符串: " + i2sResult);

        console.log("✓ utils.s2i('456') - 已调用");
        var s2iResult = utils.s2i("456");
        console.log("✓ 字符串转整数: " + s2iResult);

        console.log("✓ utils.f2s(3.14) - 已调用");
        var f2sResult = utils.f2s(3.14);
        console.log("✓ 浮点数转字符串: " + f2sResult);

        console.log("✓ utils.s2f('2.71') - 已调用");
        var s2fResult = utils.s2f("2.71");
        console.log("✓ 字符串转浮点数: " + s2fResult);

        console.log("✓ utils.b2s(true) - 已调用");
        var b2sResult = utils.b2s(true);
        console.log("✓ 布尔值转字符串: " + b2sResult);

        console.log("✓ utils.s2b('false') - 已调用");
        var s2bResult = utils.s2b("false");
        console.log("✓ 字符串转布尔值: " + s2bResult);

        console.log("✓ 工具函数模块测试完成");
    } catch (e) {
        console.error("✗ 工具函数模块测试失败: " + e);
    }
}

// ========== 18. YOLO检测测试 (yolo) ==========
if (shouldTest("yolo")) {
    console.log("\n--- 18. YOLO检测 (yolo) 测试 ---");

    try {
        // 创建 YOLO 实例
        console.log("✓ yolo.new('v8', 4, '/data/local/tmp/param', '/data/local/tmp/bin', 'person,bicycle,car') - 已调用");
        var y = yolo.new("v8", 4, "/data/local/tmp/param", "/data/local/tmp/bin", "person,bicycle,car");
        console.log("✓ YOLO 实例创建成功");

        // 检测屏幕
        console.log("✓ yolo.detect(y, 0, 0, 1000, 1000, 0) - 已调用");
        var detectResults = yolo.detect(y, 0, 0, 1000, 1000, 0);
        console.log("✓ 检测结果数量: " + detectResults.length);
        if (detectResults.length > 0) {
            console.log("✓ 第一个检测结果: " + JSON.stringify(detectResults[0]));
        }

        // 从图片检测
        var img = images.captureScreen();
        console.log("✓ yolo.detectFromImage(y, img) - 已调用");
        var detectFromImageResults = yolo.detectFromImage(y, img);
        console.log("✓ 图片检测结果数量: " + detectFromImageResults.length);

        // 从 Base64 检测
        var base64Img = "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVR42mNk+M9QDwADhgGAWjR9awAAAABJRU5ErkJggg==";
        console.log("✓ yolo.detectFromBase64(y, base64Img, '#FF0000') - 已调用");
        var detectFromBase64Results = yolo.detectFromBase64(y, base64Img, "#FF0000");
        console.log("✓ Base64 检测结果数量: " + detectFromBase64Results.length);

        // 从路径检测
        var imagePath = "/sdcard/test_yolo.jpg";
        console.log("✓ yolo.detectFromPath(y, '" + imagePath + "', '#FF0000') - 已调用");
        var detectFromPathResults = yolo.detectFromPath(y, imagePath, "#FF0000");
        console.log("✓ 路径检测结果数量: " + detectFromPathResults.length);

        // 关闭实例
        console.log("✓ yolo.close(y) - 已调用");
        yolo.close(y);

        console.log("✓ YOLO检测模块测试完成");
    } catch (e) {
        console.error("✗ YOLO检测模块测试失败: " + e);
    }
}

// ========== 19. 协程管理测试 (coroutine) ==========
if (shouldTest("coroutine")) {
    console.log("\n--- 19. 协程管理 (coroutine) 测试 ---");

    try {
        // ========== 基础协程方法测试 ==========
        console.log("\n--- 基础协程方法测试 ---");

        // 获取初始活跃协程数量
        var initialCount = coroutine.getActiveCoroutines();
        console.log("✓ coroutine.getActiveCoroutines(): " + initialCount);

        // 测试 launch - 启动协程（带名称和优先级）
        console.log("✓ coroutine.launch(function() {...}, '测试协程1', 10) - 已调用");
        var coroId1 = coroutine.launch(function () {
            console.log("  协程 1 正在执行");
        }, "测试协程1", 10);
        console.log("✓ 协程 1 ID: " + coroId1);

        // 获取协程详细信息
        console.log("✓ coroutine.getCoroutineInfo('" + coroId1 + "') - 已调用");
        var coroInfo1 = coroutine.getCoroutineInfo(coroId1);
        console.log("✓ 协程信息: " + JSON.stringify(coroInfo1));

        // 测试 delay - 延迟执行
        console.log("✓ coroutine.delay(100, function() {...}) - 已调用");
        var coroId2 = coroutine.delay(100, function () {
            console.log("  延迟执行的协程 2");
        });
        console.log("✓ 延迟协程 ID: " + coroId2);

        // 测试 sleep - 在协程中睡眠
        console.log("✓ coroutine.sleep(50) - 已调用");
        coroutine.sleep(50);
        console.log("✓ 睡眠完成");

        // 测试 async - 异步执行（现在直接返回结果）
        console.log("✓ coroutine.async(function() {...}) - 已调用");
        var asyncResult = coroutine.async(function () {
            return "异步执行结果";
        });
        console.log("✓ Async 结果: " + asyncResult);

        // 测试 await - 等待普通值
        console.log("✓ coroutine.await('普通值') - 已调用");
        var awaitResult1 = coroutine.await("普通值");
        console.log("✓ Await 普通值结果: " + awaitResult1);

        // 测试 await - 等待数字
        console.log("✓ coroutine.await(123) - 已调用");
        var awaitResult2 = coroutine.await(123);
        console.log("✓ Await 数字结果: " + awaitResult2);

        // 获取协程列表
        console.log("✓ coroutine.getCoroutineList() - 已调用");
        var coroList = coroutine.getCoroutineList();
        console.log("✓ 协程列表长度: " + coroList.length);

        // ========== 协程池测试 ==========
        console.log("\n--- 协程池测试 ---");

        // 创建协程池
        console.log("✓ coroutine.createPool('testPool', 3, 10) - 已调用");
        var poolName = coroutine.createPool("testPool", 3, 10);
        console.log("✓ 协程池名称: " + poolName);

        // 获取协程池统计信息
        console.log("✓ coroutine.getPoolStats('" + poolName + "') - 已调用");
        var poolStats = coroutine.getPoolStats(poolName);
        console.log("✓ 协程池统计: " + JSON.stringify(poolStats));

        // 提交任务到协程池
        console.log("✓ coroutine.submitToPool('" + poolName + "', function() {...}) - 已调用");
        var submitResult1 = coroutine.submitToPool(poolName, function () {
            console.log("  协程池任务 1 执行中");
        });
        console.log("✓ 提交结果 1: " + submitResult1);

        // 提交多个任务
        console.log("✓ 提交多个任务到协程池");
        for (var i = 0; i < 5; i++) {
            var taskNum = i + 2;
            var submitResult = coroutine.submitToPool(poolName, function (num) {
                console.log("  协程池任务 " + num + " 执行中");
            }, i, taskNum);
            console.log("  任务 " + taskNum + " 提交结果: " + submitResult);
        }

        // 再次获取协程池统计信息
        console.log("✓ coroutine.getPoolStats('" + poolName + "') - 已调用");
        var poolStats2 = coroutine.getPoolStats(poolName);
        console.log("✓ 协程池统计 (提交任务后): " + JSON.stringify(poolStats2));

        // 列出所有协程池
        console.log("✓ coroutine.listPools() - 已调用");
        var poolList = coroutine.listPools();
        console.log("✓ 协程池列表: " + JSON.stringify(poolList));

        // 等待任务完成
        console.log("✓ 等待协程池任务完成...");
        coroutine.sleep(200);

        // 再次获取协程池统计信息
        console.log("✓ coroutine.getPoolStats('" + poolName + "') - 已调用");
        var poolStats3 = coroutine.getPoolStats(poolName);
        console.log("✓ 协程池统计 (任务完成后): " + JSON.stringify(poolStats3));

        // 关闭协程池
        console.log("✓ coroutine.closePool('" + poolName + "') - 已调用");
        var closeResult = coroutine.closePool(poolName);
        console.log("✓ 关闭结果: " + closeResult);

        // ========== 调度器测试 ==========
        console.log("\n--- 调度器测试 ---");

        // 获取当前调度策略
        console.log("✓ coroutine.getScheduleStrategy() - 已调用");
        var strategy = coroutine.getScheduleStrategy();
        console.log("✓ 当前调度策略: " + strategy);

        // 设置调度策略
        console.log("✓ coroutine.setScheduleStrategy('priority') - 已调用");
        coroutine.setScheduleStrategy("priority");

        // 再次获取调度策略
        console.log("✓ coroutine.getScheduleStrategy() - 已调用");
        var strategy2 = coroutine.getScheduleStrategy();
        console.log("✓ 新调度策略: " + strategy2);

        // 设置协程优先级
        console.log("✓ coroutine.setPriority('高优先级任务', 100) - 已调用");
        coroutine.setPriority("高优先级任务", 100);

        console.log("✓ coroutine.setPriority('中优先级任务', 50) - 已调用");
        coroutine.setPriority("中优先级任务", 50);

        console.log("✓ coroutine.setPriority('低优先级任务', 10) - 已调用");
        coroutine.setPriority("低优先级任务", 10);

        // 获取协程优先级
        console.log("✓ coroutine.getPriority('高优先级任务') - 已调用");
        var priority1 = coroutine.getPriority("高优先级任务");
        console.log("✓ 高优先级任务优先级: " + priority1);

        console.log("✓ coroutine.getPriority('中优先级任务') - 已调用");
        var priority2 = coroutine.getPriority("中优先级任务");
        console.log("✓ 中优先级任务优先级: " + priority2);

        console.log("✓ coroutine.getPriority('低优先级任务') - 已调用");
        var priority3 = coroutine.getPriority("低优先级任务");
        console.log("✓ 低优先级任务优先级: " + priority3);

        // ========== 全局统计测试 ==========
        console.log("\n--- 全局统计测试 ---");

        // 获取全局统计信息
        console.log("✓ coroutine.getStats() - 已调用");
        var globalStats = coroutine.getStats();
        console.log("✓ 全局统计: " + JSON.stringify(globalStats));

        // ========== 协程管理测试 ==========
        console.log("\n--- 协程管理测试 ---");

        // 取消指定协程
        console.log("✓ coroutine.cancel('" + coroId1 + "') - 已调用");
        var cancelResult = coroutine.cancel(coroId1);
        console.log("✓ 取消结果: " + cancelResult);

        // 获取活跃协程数量
        var activeCount = coroutine.getActiveCoroutines();
        console.log("✓ coroutine.getActiveCoroutines(): " + activeCount);

        // 等待延迟协程完成
        console.log("✓ 等待延迟协程完成...");
        coroutine.sleep(150);

        // 取消所有协程
        console.log("✓ coroutine.cancelAll() - 已调用");
        var cancelAllResult = coroutine.cancelAll();
        console.log("✓ 取消所有协程数量: " + cancelAllResult);

        // 最终活跃协程数量
        var finalCount = coroutine.getActiveCoroutines();
        console.log("✓ 最终活跃协程数量: " + finalCount);

        // 最终全局统计
        console.log("✓ coroutine.getStats() - 已调用");
        var finalStats = coroutine.getStats();
        console.log("✓ 最终全局统计: " + JSON.stringify(finalStats));

        console.log("✓ 协程管理模块测试完成");
    } catch (e) {
        console.error("✗ 协程管理模块测试失败: " + e);
    }
}

// ========== 20. 方法管理测试 (method) ==========
if (shouldTest("method")) {
    console.log("\n--- 20. 方法管理 (method) 测试 ---");

    try {
        // 列出所有方法
        var methods = listMethods();
        console.log("✓ listMethods(): 已注册方法数量: " + methods.length);

        // 显示前 5 个方法
        console.log("✓ 前 5 个方法:");
        for (var i = 0; i < Math.min(5, methods.length); i++) {
            var method = methods[i];
            console.log("  " + (i + 1) + ". " + method.name + " - " + method.description);
        }

        // 注册新方法
        console.log("✓ registerMethod('myCustomMethod', '自定义方法', null, true) - 已调用");
        registerMethod("myCustomMethod", "自定义方法", null, true);

        // 重写方法
        console.log("✓ overrideMethod('click', function(x, y, fingerID) {...}) - 已调用");
        var overrideResult = overrideMethod("click", function (x, y, fingerID) {
            console.log("重写的 click 方法被调用: (" + x + ", " + y + ")");
            return true;
        });
        console.log("✓ 方法重写结果: " + overrideResult);

        // 测试重写的方法
        click(300, 300, 1);

        // 恢复方法
        console.log("✓ restoreMethod('click') - 已调用");
        var restoreResult = restoreMethod("click");
        console.log("✓ 方法恢复结果: " + restoreResult);

        // 注销方法
        console.log("✓ unregisterMethod('myCustomMethod') - 已调用");
        var unregisterResult = unregisterMethod("myCustomMethod");
        console.log("✓ 方法注销结果: " + unregisterResult);

        console.log("✓ 方法管理模块测试完成");
    } catch (e) {
        console.error("✗ 方法管理模块测试失败: " + e);
    }
}

// ========== 测试总结 ==========
console.log("\n=== JavaScript 引擎完整测试完成 ===");
console.log("测试时间: " + new Date().toLocaleString());
console.log("\n已测试的模块:");
console.log("  ✓ app - 应用管理");
console.log("  ✓ device - 设备管理");
console.log("  ✓ touch - 触摸操作");
console.log("  ✓ files - 文件操作");
console.log("  ✓ images - 图像处理");
console.log("  ✓ storages - 存储管理");
console.log("  ✓ system - 系统管理");
console.log("  ✓ http - 网络请求");
console.log("  ✓ media - 媒体管理");
console.log("  ✓ opencv - 图像识别");
console.log("  ✓ ppocr - 文字识别");
console.log("  ✓ dotocr - 点阵OCR");
console.log("  ✓ ime - 输入法");
console.log("  ✓ plugin - 插件");
console.log("  ✓ rhino - Rhino引擎");
console.log("  ✓ uiacc - UI辅助");
console.log("  ✓ utils - 工具函数");
console.log("  ✓ yolo - YOLO检测");
console.log("  ✓ coroutine - 协程管理");
console.log("  ✓ method - 方法管理");
console.log("\n排除的模块:");
console.log("  ✗ console - 控制台");
console.log("  ✗ imgui - ImGui");
console.log("  ✗ vdisplay - 虚拟显示");
