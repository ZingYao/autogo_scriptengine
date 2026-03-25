# http 模块

## 模块简介

http 模块提供了 HTTP 网络请求功能，支持 GET、POST 等请求方式。

## 方法列表

### http.request
发起网络请求

**使用示例：**
```lua
-- 调用 http.request 方法
http.request();
```

---

## 综合使用示例

### 示例1：发送HTTP请求
```lua
var response = http.get("https://api.example.com/data", 5000);
console.log("状态码: " + response.code);
console.log("数据: " + response.data);
```