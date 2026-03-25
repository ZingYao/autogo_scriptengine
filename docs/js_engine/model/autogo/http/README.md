# http 模块

## 模块简介

http 模块提供了 HTTP 网络请求功能，支持 GET、POST 等请求方式。

## 方法列表

### http.get
发送GET请求

**使用示例：**
```javascript
// 调用 http.get 方法
http.get();
```

---

### http.post
发送POST请求

**使用示例：**
```javascript
// 调用 http.post 方法
http.post();
```

---

### http.postMultipart
发送Multipart POST请求

**使用示例：**
```javascript
// 调用 http.postMultipart 方法
http.postMultipart();
```

---

## 综合使用示例

### 示例1：发送HTTP请求
```javascript
var response = http.get("https://api.example.com/data", 5000);
console.log("状态码: " + response.code);
console.log("数据: " + response.data);
```