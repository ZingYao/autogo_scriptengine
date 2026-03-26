# websocket 模块

## 模块简介

websocket 模块提供了 WebSocket 客户端功能，支持连接到 WebSocket 服务器、发送和接收消息、处理连接状态等操作。

## 方法列表

### websocket.connect

连接到 WebSocket 服务器

**参数：**
- `url` (string): WebSocket 服务器地址
- `onOpened` (function): 连接成功回调函数，参数为连接句柄
- `onClosed` (function): 连接关闭回调函数，参数为连接句柄
- `onError` (function): 连接错误回调函数，参数为连接句柄和错误信息
- `onRecv` (function): 接收消息回调函数，参数为连接句柄和消息内容

**返回值：**
- `number`: 连接句柄，用于后续的发送和关闭操作

**使用示例：**
```lua
-- 连接到 WebSocket 服务器
local handle = websocket.connect(
    "ws://echo.websocket.org",
    function(h)
        console.log("连接成功，句柄: " .. h)
    end,
    function(h)
        console.log("连接关闭，句柄: " .. h)
    end,
    function(h, err)
        console.log("连接错误: " .. err)
    end,
    function(h, msg)
        console.log("收到消息: " .. msg)
    end
)
```

---

### websocket.send

向 WebSocket 服务器发送消息

**参数：**
- `handle` (number): 连接句柄（由 connect 方法返回）
- `text` (string): 要发送的消息内容

**返回值：**
- `boolean`: 发送成功返回 true，失败返回 false

**使用示例：**
```lua
-- 发送消息
local success = websocket.send(handle, "Hello WebSocket")
if success then
    console.log("发送成功")
else
    console.log("发送失败")
end
```

---

### websocket.close

关闭 WebSocket 连接

**参数：**
- `handle` (number): 连接句柄（由 connect 方法返回）

**返回值：**
- 无

**使用示例：**
```lua
-- 关闭连接
websocket.close(handle)
console.log("连接已关闭")
```

---

## 综合使用示例

### 示例1：完整的 WebSocket 客户端

```lua
-- 连接到 WebSocket 服务器
local handle = websocket.connect(
    "ws://echo.websocket.org",
    function(h)
        console.log("连接成功，句柄: " .. h)
    end,
    function(h)
        console.log("连接关闭，句柄: " .. h)
    end,
    function(h, err)
        console.log("连接错误: " .. err)
    end,
    function(h, msg)
        console.log("收到消息: " .. msg)
    end
)

-- 发送消息
websocket.send(handle, "Hello, WebSocket Server!")

-- 等待一段时间后关闭连接
os.sleep(5000)
websocket.close(handle)
```

### 示例2：WebSocket 聊天室客户端

```lua
local wsHandle = nil

-- 连接到聊天室
wsHandle = websocket.connect(
    "ws://chat.example.com:8080",
    function(h)
        console.log("已连接到聊天室")
    end,
    function(h)
        console.log("已断开连接")
    end,
    function(h, err)
        console.log("连接错误: " .. err)
    end,
    function(h, msg)
        console.log("收到消息: " .. msg)
    end
)

-- 发送聊天消息
function sendMessage(text)
    if wsHandle ~= nil then
        local success = websocket.send(wsHandle, text)
        if success then
            console.log("消息已发送: " .. text)
        else
            console.log("发送失败")
        end
    end
end

-- 模拟发送消息
sendMessage("大家好！")
```

## 注意事项

1. WebSocket 连接是异步的，消息接收通过回调函数处理
2. 每个连接都有一个唯一的句柄，用于标识和管理连接
3. 连接句柄从 1 开始递增
4. 关闭连接后，句柄会被释放，不能再用于发送消息
5. 错误回调会在连接失败或连接中断时触发
6. 消息接收回调会在收到服务器消息时触发
7. 建议在 onClosed 回调中清理相关资源
8. WebSocket 连接支持 wss:// 和 ws:// 协议
