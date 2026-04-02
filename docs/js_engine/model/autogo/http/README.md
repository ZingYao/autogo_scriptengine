# http 模块

## 模块简介

http 模块提供了 HTTP 网络请求功能，支持 GET、POST 等请求方式。

## 方法列表

### http.get
发送GET请求

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| url | string | 请求的 URL |
| timeout | number | 超时时间（毫秒） |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| code | number | HTTP 状态码 |
| data | string | 响应数据 |

**使用示例：**
```javascript
// 发送 GET 请求
var response = http.get("https://api.example.com/data", 5000);
console.log("状态码: " + response.code);
console.log("数据: " + response.data);
```

---

### http.post
发送POST请求

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| url | string | 请求的 URL |
| data | string/array | 请求数据（字符串或字节数组） |
| headers | object | 请求头（可选） |
| timeout | number | 超时时间（毫秒，可选，默认 5000） |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| code | number | HTTP 状态码 |
| data | string | 响应数据 |

**使用示例：**
```javascript
// 发送 POST 请求
var response = http.post("https://api.example.com/data", "Hello World", {"Content-Type": "text/plain"}, 5000);
console.log("状态码: " + response.code);
console.log("数据: " + response.data);
```

---

### http.postMultipart
发送Multipart POST请求

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| url | string | 请求的 URL |
| fileName | string | 文件名 |
| fileData | string/array | 文件数据（字符串或字节数组） |
| timeout | number | 超时时间（毫秒） |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| code | number | HTTP 状态码 |
| data | string | 响应数据 |

**使用示例：**
```javascript
// 发送 Multipart POST 请求
var response = http.postMultipart("https://api.example.com/upload", "test.txt", "Hello World", 5000);
console.log("状态码: " + response.code);
console.log("数据: " + response.data);
```

---

## 综合使用示例

### 示例1：发送HTTP请求
```javascript
var response = http.get("https://api.example.com/data", 5000);
console.log("状态码: " + response.code);
console.log("数据: " + response.data);
```