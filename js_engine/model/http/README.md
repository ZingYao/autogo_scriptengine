# HTTP 模块

HTTP 模块提供了 HTTP 请求功能，支持 GET、POST 和 Multipart POST 请求，用于与网络服务进行交互。

## 方法列表

### http.get(url, timeout)
发送 GET 请求到指定的 URL。

**入参：**
- `url`: 请求的 URL 地址（字符串）
- `timeout`: 超时时间（毫秒，整数）

**出参：** 包含 code 和 data 字段的对象
- `code`: HTTP 状态码（整数）
- `data`: 响应数据（字符串），如果失败则为 null

**调用示例：**
```javascript
// 发送 GET 请求
var response = http.get("https://api.example.com/data", 5000);
console.println("状态码:", response.code);
console.println("响应数据:", response.data);
```

### http.post(url, data, headers, timeout)
发送 POST 请求到指定的 URL。

**入参：**
- `url`: 请求的 URL 地址（字符串）
- `data`: 请求体数据（字符串或字节数组）
- `headers`: 请求头（对象，可选）
- `timeout`: 超时时间（毫秒，整数，默认 5000）

**出参：** 包含 code 和 data 字段的对象
- `code`: HTTP 状态码（整数）
- `data`: 响应数据（字符串），如果失败则为 null

**调用示例：**
```javascript
// 发送 POST 请求（字符串数据）
var response = http.post(
    "https://api.example.com/submit",
    JSON.stringify({ name: "张三", age: 25 }),
    { "Content-Type": "application/json" },
    5000
);
console.println("状态码:", response.code);
console.println("响应数据:", response.data);

// 发送 POST 请求（字节数组数据）
var byteData = new Uint8Array([72, 101, 108, 108, 111]);
var response2 = http.post(
    "https://api.example.com/upload",
    byteData,
    { "Content-Type": "application/octet-stream" },
    5000
);
```

### http.postMultipart(url, fileName, fileData, timeout)
发送 Multipart POST 请求，用于文件上传。

**入参：**
- `url`: 请求的 URL 地址（字符串）
- `fileName`: 文件名（字符串）
- `fileData`: 文件数据（字符串或字节数组）
- `timeout`: 超时时间（毫秒，整数）

**出参：** 包含 code 和 data 字段的对象
- `code`: HTTP 状态码（整数）
- `data`: 响应数据（字符串），如果失败则为 null

**调用示例：**
```javascript
// 上传文件
var fileData = files.readBytes("/sdcard/image.png");
var response = http.postMultipart(
    "https://api.example.com/upload",
    "image.png",
    fileData,
    10000
);
console.println("状态码:", response.code);
console.println("响应数据:", response.data);
```

## 完整示例

```javascript
// 示例1：简单的 GET 请求
function fetchData() {
    var response = http.get("https://api.example.com/users", 5000);
    if (response.code === 200) {
        console.println("获取数据成功:", response.data);
        var users = JSON.parse(response.data);
        console.println("用户数量:", users.length);
    } else {
        console.println("请求失败，状态码:", response.code);
    }
}

// 示例2：发送 JSON 数据的 POST 请求
function submitData() {
    var data = {
        username: "testuser",
        password: "123456",
        email: "test@example.com"
    };
    
    var response = http.post(
        "https://api.example.com/register",
        JSON.stringify(data),
        {
            "Content-Type": "application/json",
            "User-Agent": "MyApp/1.0"
        },
        5000
    );
    
    if (response.code === 200) {
        console.println("注册成功:", response.data);
    } else {
        console.println("注册失败:", response.code);
    }
}

// 示例3：上传文件
function uploadFile(filePath) {
    var fileData = files.readBytes(filePath);
    var fileName = files.getName(filePath);
    
    var response = http.postMultipart(
        "https://api.example.com/upload",
        fileName,
        fileData,
        10000
    );
    
    if (response.code === 200) {
        console.println("上传成功:", response.data);
    } else {
        console.println("上传失败:", response.code);
    }
}

// 示例4：带重试的请求
function requestWithRetry(url, maxRetries) {
    for (var i = 0; i < maxRetries; i++) {
        var response = http.get(url, 5000);
        if (response.code === 200) {
            return response.data;
        }
        console.println("请求失败，重试第 " + (i + 1) + " 次");
        utils.sleep(1000);
    }
    return null;
}

// 调用示例
fetchData();
submitData();
uploadFile("/sdcard/test.png");
var result = requestWithRetry("https://api.example.com/data", 3);
```

## 注意事项

1. timeout 参数单位为毫秒，建议根据网络情况设置合理的超时时间
2. POST 请求的 data 参数支持字符串和字节数组两种格式
3. headers 参数用于设置自定义请求头，如 Content-Type、User-Agent 等
4. Multipart POST 专门用于文件上传，会自动设置正确的 Content-Type
5. 响应数据为字符串格式，如果需要解析 JSON，需要使用 JSON.parse()
6. 建议在请求前检查网络连接状态
7. 对于大文件上传，建议适当增加超时时间
8. 注意处理网络异常和超时情况
