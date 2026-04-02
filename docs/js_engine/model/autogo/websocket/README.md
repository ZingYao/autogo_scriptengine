# websocket 模块

## 模块简介

websocket 模块提供了 WebSocket 客户端功能，用于与 WebSocket 服务器进行双向通信。

## 调用方式

websocket 模块支持两种调用方式：

### 1. 传统方式（使用句柄）

使用 `connect` 方法连接，返回句柄 ID，后续操作需要使用句柄 ID。

### 2. 面向对象方式（推荐）

使用 `connect` 方法连接，返回 WebSocket 对象，可以直接调用对象的方法。

## 方法列表

### websocket.connect
连接 WebSocket 服务器（面向对象方式）

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| url | string | WebSocket 服务器地址 |
| onOpened | function | 连接成功回调，参数为 WebSocket 对象 |
| onClosed | function | 连接关闭回调，参数为 WebSocket 对象 |
| onError | function | 错误回调，参数为 WebSocket 对象和错误信息 |
| onRecv | function | 接收消息回调，参数为 WebSocket 对象和消息内容 |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| ws | object | WebSocket 对象，包含 send、close 方法和 handle 属性 |

**使用示例：**
```javascript
var ws = websocket.connect(
    "ws://echo.websocket.org",
    function(conn) { console.log("连接成功，句柄: " + conn.handle); },
    function(conn) { console.log("连接关闭，句柄: " + conn.handle); },
    function(conn, err) { console.log("错误: " + err); },
    function(conn, msg) { console.log("收到消息: " + msg); }
);

// 发送消息
ws.send("Hello WebSocket!");

// 关闭连接
ws.close();
```

### websocket.connectObject

此方法已废弃，请使用 `connect` 方法。`connect` 方法现在直接返回 WebSocket 对象，支持面向对象调用方式。

### websocket.send
发送消息（传统方式）

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| handle | number | WebSocket 连接句柄 |
| text | string | 要发送的消息内容 |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| success | boolean | 发送成功返回 true，失败返回 false |

**使用示例：**
```javascript
var success = websocket.send(handle, "Hello WebSocket!");
if (success) {
    console.log("发送成功");
} else {
    console.log("发送失败");
}
```

### websocket.close
关闭连接（传统方式）

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| handle | number | WebSocket 连接句柄 |

**返回值：**

无返回值

**使用示例：**
```javascript
websocket.close(handle);
```

## WebSocket 对象方法

使用 `connect` 返回的 WebSocket 对象包含以下方法：

### send
发送消息

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| text | string | 要发送的消息内容 |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| success | boolean | 发送成功返回 true，失败返回 false |

**使用示例：**
```javascript
ws.send("Hello WebSocket!");
```

### close
关闭连接

**参数：**

无参数

**返回值：**

无返回值

**使用示例：**
```javascript
ws.close();
```

### handle

获取连接句柄 ID

**使用示例：**

```javascript
console.log("句柄 ID: " + ws.handle);
```

## 完整示例

### 传统方式示例

```javascript
// WebSocket 测试脚本（传统方式）

var handle = websocket.connect(
    "ws://echo.websocket.org",
    function(h) { 
        console.log("连接成功，句柄: " + h); 
        // 连接成功后发送消息
        websocket.send(h, "Hello from traditional mode!");
    },
    function(h) { 
        console.log("连接关闭，句柄: " + h); 
    },
    function(h, err) { 
        console.log("错误: " + err); 
    },
    function(h, msg) { 
        console.log("收到消息: " + msg); 
    }
);

// 等待一段时间
sleep(5000);

// 关闭连接
websocket.close(handle);
```

### 面向对象方式示例

```javascript
// WebSocket 测试脚本（面向对象方式）
var ws = websocket.connect(
    "ws://echo.websocket.org",
    function(conn) { 
        console.log("连接成功，句柄: " + conn.handle); 
        // 连接成功后发送消息
        conn.send("Hello from object mode!");
    },
    function(conn) { 
        console.log("连接关闭，句柄: " + conn.handle); 
    },
    function(conn, err) { 
        console.log("错误: " + err); 
    },
    function(conn, msg) { 
        console.log("收到消息: " + msg); 
    }
);

// 等待一段时间
sleep(5000);

// 关闭连接
ws.close();
```

### 多连接示例（面向对象方式）

```javascript
// WebSocket 多连接测试脚本
var ws1 = websocket.connect(
    "ws://echo.websocket.org",
    function(conn) { 
        console.log("连接1成功"); 
        conn.send("Message from connection 1");
    },
    function(conn) { console.log("连接1关闭"); },
    function(conn, err) { console.log("连接1错误: " + err); },
    function(conn, msg) { console.log("连接1收到: " + msg); }
);

var ws2 = websocket.connect(
    "ws://echo.websocket.org",
    function(conn) { 
        console.log("连接2成功"); 
        conn.send("Message from connection 2");
    },
    function(conn) { console.log("连接2关闭"); },
    function(conn, err) { console.log("连接2错误: " + err); },
    function(conn, msg) { console.log("连接2收到: " + msg); }
);

// 等待一段时间
sleep(10000);

// 关闭所有连接
ws1.close();
ws2.close();
```

## 注意事项

1. **推荐使用面向对象方式**：`connect` 返回的对象更易用，不需要记住句柄 ID
2. **回调函数参数**：
   - 传统方式：回调函数的第一个参数是句柄 ID
   - 面向对象方式：回调函数的第一个参数是 WebSocket 对象
3. **错误处理**：建议在 `onError` 回调中处理连接错误
4. **资源释放**：使用完毕后记得调用 `close()` 方法关闭连接
5. **异步操作**：WebSocket 连接和消息接收都是异步的，需要在回调函数中处理
6. **多连接支持**：可以同时创建多个 WebSocket 连接，每个连接都是独立的

## 两种方式对比

| 特性   | 传统方式                           | 面向对象方式          |
| ---- | ------------------------------ | --------------- |
| 连接方法 | `connect`                      | `connect` |
| 返回值  | 句柄 ID                          | WebSocket 对象    |
| 发送消息 | `websocket.send(handle, text)` | `ws.send(text)` |
| 关闭连接 | `websocket.close(handle)`      | `ws.close()`    |
| 回调参数 | 句柄 ID                          | WebSocket 对象    |
| 易用性  | 需要记住句柄                         | 直接使用对象          |
| 推荐度  | 一般                             | **推荐**          |

## WebSocket 服务器地址示例

```javascript
// 公共测试服务器
"ws://echo.websocket.org"           // WebSocket Echo 服务器
"wss://echo.websocket.org"          // 安全 WebSocket Echo 服务器

// 本地测试服务器
"ws://localhost:8080"
"wss://localhost:8443"

// 自定义服务器
"ws://your-server.com/path"
"wss://your-server.com/path"
```

