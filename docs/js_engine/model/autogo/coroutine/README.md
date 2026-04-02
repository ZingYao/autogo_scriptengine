# coroutine 模块

## 模块简介

coroutine 模块提供了 JavaScript 协程（并发）功能的支持。该模块允许在脚本中创建和管理协程，实现异步任务的并发执行、协程池管理以及任务调度等功能。

## 常量

### 协程状态 (CoroutineState)

| 常量名 | 值 | 说明 |
|--------|-----|------|
| StatePending | 0 | 等待中 |
| StateRunning | 1 | 运行中 |
| StateCompleted | 2 | 已完成 |
| StateCancelled | 3 | 已取消 |
| StateError | 4 | 错误 |

## 方法列表

### 基础协程方法

#### coroutine.launch(fn, name, priority)
启动一个新的协程

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| fn | Function | 要执行的函数 |
| name | String | 协程名称（可选） |
| priority | Number | 协程优先级，默认为 0（可选） |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| coroutineId | String | 协程ID |

**使用示例：**
```javascript
// 启动一个简单的协程
var coroId = coroutine.launch(function() {
    console.log("协程开始执行");
    coroutine.sleep(1000);
    console.log("协程执行完成");
}, "myCoroutine", 1);
```

---

#### coroutine.delay(ms, fn)
延迟执行函数

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| ms | Number | 延迟毫秒数 |
| fn | Function | 要执行的函数 |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| coroutineId | String | 协程ID |

**使用示例：**
```javascript
// 3秒后执行
var delayId = coroutine.delay(3000, function() {
    console.log("3秒后执行");
});
```

---

#### coroutine.async(fn)
异步执行函数并返回结果（同步等待）

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| fn | Function | 要执行的函数 |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| result | Any | 函数执行结果 |

**使用示例：**
```javascript
// 异步执行并等待结果
var result = coroutine.async(function() {
    coroutine.sleep(500);
    return "完成";
});
console.log(result); // 输出: 完成
```

---

#### coroutine.await(value)
等待并返回传入的值（简化版本）

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| value | Any | 要等待的值 |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| value | Any | 传入的值 |

**使用示例：**
```javascript
var value = coroutine.await(somePromise);
```

---

#### coroutine.sleep(ms)
在协程中睡眠指定的毫秒数

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| ms | Number | 睡眠毫秒数 |

**返回值：**

无返回值

**使用示例：**
```javascript
coroutine.sleep(1000); // 睡眠1秒
```

---

### 协程管理方法

#### coroutine.cancel(coroutineId)
取消指定的协程

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| coroutineId | String | 协程ID |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| success | Boolean | 是否成功取消 |

**使用示例：**
```javascript
var coroId = coroutine.launch(function() {
    // 长时间运行的任务
});

// 取消协程
var success = coroutine.cancel(coroId);
```

---

#### coroutine.getActiveCoroutines()
获取活跃的协程数量

**参数：**

无参数

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| count | Number | 活跃协程数量 |

**使用示例：**
```javascript
var count = coroutine.getActiveCoroutines();
console.log("活跃协程数: " + count);
```

---

#### coroutine.getCoroutineList()
获取协程列表

**参数：**

无参数

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| list | Array | 协程信息数组，每个元素包含：id, name, state, priority, duration |

**使用示例：**
```javascript
var list = coroutine.getCoroutineList();
for (var i = 0; i < list.length; i++) {
    console.log("协程: " + list[i].name + ", 状态: " + list[i].state);
}
```

---

#### coroutine.getCoroutineInfo(coroutineId)
获取指定协程的详细信息

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| coroutineId | String | 协程ID |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| id | String | 协程ID |
| name | String | 协程名称 |
| state | Number | 协程状态 |
| priority | Number | 优先级 |
| startedAt | Number | 开始时间 |
| completedAt | Number | 完成时间（如果有） |
| duration | Number | 运行时长 |
| error | String | 错误信息（如果有） |
| result | Any | 执行结果（如果有） |

**使用示例：**
```javascript
var info = coroutine.getCoroutineInfo(coroId);
console.log("协程状态: " + info.state);
```

---

#### coroutine.cancelAll()
取消所有协程

**参数：**

无参数

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| count | Number | 取消的协程数量 |

**使用示例：**
```javascript
var count = coroutine.cancelAll();
console.log("取消了 " + count + " 个协程");
```

---

#### coroutine.getStats()
获取全局统计信息

**参数：**

无参数

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| totalTasks | Number | 总任务数 |
| completed | Number | 已完成数 |
| failed | Number | 失败数 |
| cancelled | Number | 取消数 |
| active | Number | 活跃协程数 |
| pools | Number | 协程池数量 |

**使用示例：**
```javascript
var stats = coroutine.getStats();
console.log("总任务: " + stats.totalTasks);
console.log("已完成: " + stats.completed);
console.log("失败: " + stats.failed);
```

---

### 协程池方法

#### coroutine.createPool(name, maxWorkers, maxTasks)
创建协程池

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| name | String | 协程池名称 |
| maxWorkers | Number | 最大工作协程数 |
| maxTasks | Number | 最大任务数，默认100（可选） |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| poolName | String | 协程池名称 |

**使用示例：**
```javascript
coroutine.createPool("myPool", 5, 50);
```

---

#### coroutine.submitToPool(poolName, fn, priority)
提交任务到协程池

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| poolName | String | 协程池名称 |
| fn | Function | 要执行的任务函数 |
| priority | Number | 任务优先级，默认0（可选） |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| success | Boolean | 是否提交成功 |

**使用示例：**
```javascript
coroutine.createPool("workerPool", 3);

for (var i = 0; i < 10; i++) {
    coroutine.submitToPool("workerPool", function() {
        console.log("执行任务");
    }, i);
}
```

---

#### coroutine.getPoolStats(poolName)
获取协程池统计信息

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| poolName | String | 协程池名称 |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| name | String | 名称 |
| maxWorkers | Number | 最大工作协程数 |
| maxTasks | Number | 最大任务数 |
| active | Number | 活跃工作数 |
| queued | Number | 队列中的任务数 |
| workers | Number | 工作协程数 |
| closed | Boolean | 是否已关闭 |

**使用示例：**
```javascript
var stats = coroutine.getPoolStats("myPool");
console.log("活跃: " + stats.active + ", 队列: " + stats.queued);
```

---

#### coroutine.closePool(poolName)
关闭协程池

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| poolName | String | 协程池名称 |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| success | Boolean | 是否成功关闭 |

**使用示例：**
```javascript
coroutine.closePool("myPool");
```

---

#### coroutine.listPools()
列出所有协程池

**参数：**

无参数

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| pools | Array | 协程池名称数组 |

**使用示例：**
```javascript
var pools = coroutine.listPools();
console.log("协程池列表: " + pools.join(", "));
```

---

### 调度器方法

#### coroutine.setScheduleStrategy(strategy)
设置调度策略

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| strategy | String | 调度策略名称（如 "fifo"） |

**返回值：**

无返回值

**使用示例：**
```javascript
coroutine.setScheduleStrategy("fifo");
```

---

#### coroutine.getScheduleStrategy()
获取当前调度策略

**参数：**

无参数

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| strategy | String | 当前调度策略 |

**使用示例：**
```javascript
var strategy = coroutine.getScheduleStrategy();
console.log("当前调度策略: " + strategy);
```

---

#### coroutine.setPriority(name, priority)
设置协程优先级

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| name | String | 协程名称 |
| priority | Number | 优先级值 |

**返回值：**

无返回值

**使用示例：**
```javascript
coroutine.setPriority("myCoroutine", 10);
```

---

#### coroutine.getPriority(name)
获取协程优先级

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| name | String | 协程名称 |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| priority | Number | 优先级值 |

**使用示例：**
```javascript
var priority = coroutine.getPriority("myCoroutine");
console.log("优先级: " + priority);
```console.log("优先级: " + priority);
```

---

## 综合使用示例

### 示例1：并发下载任务
```javascript
// 创建协程池处理下载任务
coroutine.createPool("downloadPool", 5);

var urls = ["url1", "url2", "url3", "url4", "url5"];

for (var i = 0; i < urls.length; i++) {
    (function(url) {
        coroutine.submitToPool("downloadPool", function() {
            console.log("开始下载: " + url);
            coroutine.sleep(1000); // 模拟下载
            console.log("下载完成: " + url);
        });
    })(urls[i]);
}

// 等待一段时间后关闭池
coroutine.sleep(10000);
coroutine.closePool("downloadPool");
```

### 示例2：定时任务
```javascript
// 每5秒执行一次任务
function scheduleTask() {
    coroutine.delay(5000, function() {
        console.log("执行定时任务: " + new Date());
        scheduleTask(); // 递归调用实现循环
    });
}

scheduleTask();
```

### 示例3：协程状态监控
```javascript
// 启动多个协程
for (var i = 0; i < 5; i++) {
    coroutine.launch(function() {
        coroutine.sleep(Math.random() * 3000);
    }, "task_" + i);
}

// 监控协程状态
coroutine.delay(1000, function checkStatus() {
    var list = coroutine.getCoroutineList();
    console.log("=== 协程状态 ===");
    for (var i = 0; i < list.length; i++) {
        console.log(list[i].name + ": " + list[i].state);
    }
    
    if (coroutine.getActiveCoroutines() > 0) {
        coroutine.delay(1000, checkStatus);
    }
});
```
