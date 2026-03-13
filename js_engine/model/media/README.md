# Media 模块

Media 模块提供了媒体文件相关的功能，包括文件扫描、音频播放和短信发送。

## 方法列表

### media.scanFile(path)
扫描指定路径的媒体文件，使其在系统中可见。

**入参：**
- `path`: 文件或目录路径（字符串）

**出参：** 无

**调用示例：**
```javascript
// 扫描单个文件
media.scanFile("/sdcard/Download/image.png");

// 扫描目录
media.scanFile("/sdcard/Music");
```

### media.playMP3(path)
播放指定的 MP3 音频文件。

**入参：**
- `path`: MP3 文件路径（字符串）

**出参：** 如果播放失败，返回错误信息字符串；成功则返回 null

**调用示例：**
```javascript
// 播放 MP3 文件
var result = media.playMP3("/sdcard/Music/song.mp3");
if (result === null) {
    console.println("开始播放音乐");
} else {
    console.println("播放失败:", result);
}
```

### media.sendSMS(number, message)
发送短信到指定号码。

**入参：**
- `number`: 手机号码（字符串）
- `message`: 短信内容（字符串）

**出参：** 无

**调用示例：**
```javascript
// 发送短信
media.sendSMS("13800138000", "这是一条测试短信");
```

## 完整示例

```javascript
// 示例1：扫描媒体文件
function scanMediaFiles() {
    console.println("开始扫描媒体文件...");
    
    // 扫描图片文件
    media.scanFile("/sdcard/Pictures");
    console.println("图片文件扫描完成");
    
    // 扫描音乐文件
    media.scanFile("/sdcard/Music");
    console.println("音乐文件扫描完成");
    
    // 扫描视频文件
    media.scanFile("/sdcard/Movies");
    console.println("视频文件扫描完成");
}

// 示例2：播放音乐
function playMusicDemo() {
    var musicPath = "/sdcard/Music/song.mp3";
    
    // 检查文件是否存在
    if (files.exists(musicPath)) {
        console.println("开始播放:", musicPath);
        var result = media.playMP3(musicPath);
        
        if (result === null) {
            console.println("音乐播放成功");
        } else {
            console.println("音乐播放失败:", result);
        }
    } else {
        console.println("音乐文件不存在:", musicPath);
    }
}

// 示例3：播放列表
function playMusicList() {
    var musicList = [
        "/sdcard/Music/song1.mp3",
        "/sdcard/Music/song2.mp3",
        "/sdcard/Music/song3.mp3"
    ];
    
    for (var i = 0; i < musicList.length; i++) {
        var musicPath = musicList[i];
        if (files.exists(musicPath)) {
            console.println("播放第 " + (i + 1) + " 首歌曲");
            var result = media.playMP3(musicPath);
            
            if (result === null) {
                console.println("播放成功");
                utils.sleep(3000); // 等待3秒
            } else {
                console.println("播放失败:", result);
            }
        } else {
            console.println("文件不存在:", musicPath);
        }
    }
}

// 示例4：发送短信
function sendSMSDemo() {
    var phoneNumber = "13800138000";
    var message = "这是一条来自 AutoGo 的测试短信";
    
    console.println("发送短信到:", phoneNumber);
    media.sendSMS(phoneNumber, message);
    console.println("短信已发送");
}

// 示例5：批量发送短信
function sendBulkSMS() {
    var contacts = [
        { name: "张三", phone: "13800138000" },
        { name: "李四", phone: "13900139000" },
        { name: "王五", phone: "13700137000" }
    ];
    
    var message = "这是群发短信内容";
    
    for (var i = 0; i < contacts.length; i++) {
        var contact = contacts[i];
        console.println("发送短信给:", contact.name);
        media.sendSMS(contact.phone, message);
        console.println("短信已发送给:", contact.name);
        utils.sleep(1000); // 避免发送过快
    }
}

// 示例6：媒体文件管理
function manageMediaFiles() {
    var sourcePath = "/sdcard/Download/new_song.mp3";
    var targetPath = "/sdcard/Music/new_song.mp3";
    
    // 移动文件到音乐目录
    if (files.exists(sourcePath)) {
        files.move(sourcePath, targetPath);
        console.println("文件已移动到:", targetPath);
        
        // 扫描新文件
        media.scanFile(targetPath);
        console.println("文件扫描完成");
        
        // 播放新音乐
        media.playMP3(targetPath);
    } else {
        console.println("源文件不存在:", sourcePath);
    }
}

// 示例7：音乐播放器
function simpleMusicPlayer() {
    var musicDir = "/sdcard/Music";
    var musicFiles = files.list(musicDir);
    
    if (musicFiles.length === 0) {
        console.println("音乐目录为空");
        return;
    }
    
    console.println("找到 " + musicFiles.length + " 首音乐");
    
    for (var i = 0; i < musicFiles.length; i++) {
        var musicFile = musicFiles[i];
        var musicPath = musicDir + "/" + musicFile;
        
        if (musicFile.endsWith(".mp3")) {
            console.println("播放:", musicFile);
            var result = media.playMP3(musicPath);
            
            if (result === null) {
                console.println("正在播放:", musicFile);
                utils.sleep(5000); // 播放5秒后切换
            } else {
                console.println("播放失败:", result);
            }
        }
    }
}

// 调用示例
scanMediaFiles();
playMusicDemo();
playMusicList();
sendSMSDemo();
sendBulkSMS();
manageMediaFiles();
simpleMusicPlayer();
```

## 注意事项

1. 扫描媒体文件后，文件才会在系统的媒体库中显示
2. 播放 MP3 文件需要确保文件路径正确且文件存在
3. 发送短信需要相应的权限，可能需要用户授权
4. 批量发送短信时，建议添加适当的延迟，避免被系统限制
5. 播放音乐是同步操作，会阻塞当前线程直到播放完成
6. 建议在播放音乐前检查文件是否存在
7. 短信发送可能产生费用，请谨慎使用
8. 某些设备可能对短信发送有频率限制
