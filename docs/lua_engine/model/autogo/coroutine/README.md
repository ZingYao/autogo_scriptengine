# coroutine 模块

## 模块简介

coroutine 模块提供了 JavaScript 协程（并发）功能的支持。该模块允许在脚本中创建和管理协程，实现异步任务的并发执行、协程池管理以及任务调度等功能。

## 方法列表

### coroutine.launch
启动一个新的协程

**使用示例：**
```lua
-- 调用 coroutine.launch 方法
coroutine.launch();
```

---

### coroutine.delay
延迟执行函数

**使用示例：**
```lua
-- 调用 coroutine.delay 方法
coroutine.delay();
```

---

### coroutine.async
异步执行函数并返回结果（同步等待）

**使用示例：**
```lua
-- 调用 coroutine.async 方法
coroutine.async();
```

---

### coroutine.await
等待并返回传入的值（简化版本）

**使用示例：**
```lua
-- 调用 coroutine.await 方法
coroutine.await();
```

---

### coroutine.cancel
取消指定的协程

**使用示例：**
```lua
-- 调用 coroutine.cancel 方法
coroutine.cancel();
```

---

### coroutine.sleep
在协程中睡眠指定的毫秒数

**使用示例：**
```lua
-- 调用 coroutine.sleep 方法
coroutine.sleep();
```

---

### coroutine.getActiveCoroutines
获取活跃的协程数量

**使用示例：**
```lua
-- 调用 coroutine.getActiveCoroutines 方法
coroutine.getActiveCoroutines();
```

---

### coroutine.getCoroutineList
获取协程列表

**使用示例：**
```lua
-- 调用 coroutine.getCoroutineList 方法
coroutine.getCoroutineList();
```

---

### coroutine.getCoroutineInfo
获取指定协程的详细信息

**使用示例：**
```lua
-- 调用 coroutine.getCoroutineInfo 方法
coroutine.getCoroutineInfo();
```

---

### coroutine.cancelAll
取消所有协程

**使用示例：**
```lua
-- 调用 coroutine.cancelAll 方法
coroutine.cancelAll();
```

---

### coroutine.getStats
获取全局统计信息

**使用示例：**
```lua
-- 调用 coroutine.getStats 方法
coroutine.getStats();
```

---

### coroutine.createPool
创建协程池

**使用示例：**
```lua
-- 调用 coroutine.createPool 方法
coroutine.createPool();
```

---

### coroutine.submitToPool
提交任务到协程池

**使用示例：**
```lua
-- 调用 coroutine.submitToPool 方法
coroutine.submitToPool();
```

---

### coroutine.getPoolStats
获取协程池统计信息

**使用示例：**
```lua
-- 调用 coroutine.getPoolStats 方法
coroutine.getPoolStats();
```

---

### coroutine.closePool
关闭协程池

**使用示例：**
```lua
-- 调用 coroutine.closePool 方法
coroutine.closePool();
```

---

### coroutine.listPools
列出所有协程池

**使用示例：**
```lua
-- 调用 coroutine.listPools 方法
coroutine.listPools();
```

---

### coroutine.setScheduleStrategy
设置调度策略

**使用示例：**
```lua
-- 调用 coroutine.setScheduleStrategy 方法
coroutine.setScheduleStrategy();
```

---

### coroutine.getScheduleStrategy
获取当前调度策略

**使用示例：**
```lua
-- 调用 coroutine.getScheduleStrategy 方法
coroutine.getScheduleStrategy();
```

---

### coroutine.setPriority
设置协程优先级

**使用示例：**
```lua
-- 调用 coroutine.setPriority 方法
coroutine.setPriority();
```

---

### coroutine.getPriority
获取协程优先级

**使用示例：**
```lua
-- 调用 coroutine.getPriority 方法
coroutine.getPriority();
```

---

## 综合使用示例

### 示例1：启动协程
```lua
var coroId = coroutine.launch(function() {
    console.log("协程开始执行");
    coroutine.sleep(1000);
    console.log("协程执行完成");
}, "myCoroutine", 1);
```