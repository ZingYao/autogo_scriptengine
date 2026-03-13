# Coroutine 模块

Coroutine 模块提供了协程管理功能，包括协程创建、协程池、调度器等功能。

## 方法列表

### 基础协程方法

#### coroutine_launch(fn, name?, priority?)

启动一个新的协程并返回协程ID。

**参数**:
- `fn` (function): 要执行的函数
- `name` (string, 可选): 协程名称
- `priority` (number, 可选): 协程优先级，默认为0

**返回值**: `string` - 协程ID

**调用示例**:
```lua
local coroutineId = coroutine_launch(function()
    console.log("协程开始执行")
    coroutine_sleep(1000)
    console.log("协程执行完成")
    return "协程结果"
end, "myCoroutine", 10)
console.log("协程ID: " .. coroutineId)
```

#### coroutine_delay(ms, fn)

延迟指定毫秒数后执行函数，返回协程ID。

**参数**:
- `ms` (number): 延迟的毫秒数
- `fn` (function): 要执行的函数

**返回值**: `string` - 协程ID

**调用示例**:
```lua
local delayId = coroutine_delay(2000, function()
    console.log("2秒后执行")
end)
console.log("延迟协程ID: " .. delayId)
```

#### coroutine_async(fn)

异步执行函数并返回结果（同步等待完成）。

**参数**:
- `fn` (function): 要执行的函数

**返回值**: `any` - 函数执行结果

**调用示例**:
```lua
local asyncResult = coroutine_async(function()
    return "异步执行结果"
end)
console.log("异步结果: " .. asyncResult)
```

#### coroutine_await(value)

等待并返回传入的值（简化版本，直接返回参数值）。

**参数**:
- `value` (any): 要等待的值

**返回值**: `any` - 传入的值

**注意**: 当前实现为简化版本，直接返回传入的值，不进行异步等待。

**调用示例**:
```lua
local value = coroutine_await("普通值")
console.log("等待结果: " .. value)
```

#### coroutine_sleep(ms)

在协程中睡眠指定的毫秒数。

**参数**:
- `ms` (number): 睡眠的毫秒数

**返回值**: `nil`

**调用示例**:
```lua
console.log("开始睡眠")
coroutine_sleep(1000)
console.log("睡眠结束")
```

### 协程管理方法

#### coroutine_cancel(coroutineId)

取消指定的协程。

**参数**:
- `coroutineId` (string): 协程ID

**返回值**: `boolean` - 是否成功取消

**调用示例**:
```lua
local success = coroutine_cancel("coro_1")
if success then
    console.log("协程已取消")
end
```

#### coroutine_getActiveCoroutines()

获取活跃的协程数量。

**返回值**: `number` - 活跃协程数量

**调用示例**:
```lua
local activeCount = coroutine_getActiveCoroutines()
console.log("活跃协程数量: " .. activeCount)
```

#### coroutine_getCoroutineList()

获取协程列表。

**返回值**: `table` - 协程信息数组，每个协程包含 id、name、state、priority、duration

**调用示例**:
```lua
local coroutineList = coroutine_getCoroutineList()
for i, coro in ipairs(coroutineList) do
    console.log("协程ID: " .. coro.id)
    console.log("协程名称: " .. coro.name)
    console.log("协程状态: " .. coro.state)
    console.log("优先级: " .. coro.priority)
    console.log("运行时长: " .. coro.duration .. "ms")
end
```

#### coroutine_getCoroutineInfo(coroutineId)

获取指定协程的详细信息。

**参数**:
- `coroutineId` (string): 协程ID

**返回值**: `table` - 协程详细信息，包含 id、name、state、priority、startedAt、completedAt、duration、error、result

**调用示例**:
```lua
local info = coroutine_getCoroutineInfo("coro_1")
console.log("协程ID: " .. info.id)
console.log("协程名称: " .. info.name)
console.log("协程状态: " .. info.state)
console.log("开始时间: " .. info.startedAt)
console.log("完成时间: " .. info.completedAt)
console.log("运行时长: " .. info.duration .. "ms")
```

#### coroutine_cancelAll()

取消所有协程。

**返回值**: `number` - 取消的协程数量

**调用示例**:
```lua
local cancelledCount = coroutine_cancelAll()
console.log("已取消协程数量: " .. cancelledCount)
```

#### coroutine_getStats()

获取全局统计信息。

**返回值**: `table` - 统计信息，包含 totalTasks、completed、failed、cancelled、active、pools

**调用示例**:
```lua
local stats = coroutine_getStats()
console.log("总任务数: " .. stats.totalTasks)
console.log("已完成: " .. stats.completed)
console.log("失败: " .. stats.failed)
console.log("已取消: " .. stats.cancelled)
console.log("活跃协程: " .. stats.active)
console.log("协程池数量: " .. stats.pools)
```

### 协程池方法

#### coroutine_createPool(name, maxWorkers, maxTasks?)

创建协程池并返回池名称。

**参数**:
- `name` (string): 协程池名称
- `maxWorkers` (number): 最大工作协程数
- `maxTasks` (number, 可选): 最大任务数，默认为100

**返回值**: `string` - 协程池名称

**调用示例**:
```lua
local poolName = coroutine_createPool("myPool", 5, 100)
console.log("创建协程池: " .. poolName)
```

#### coroutine_submitToPool(poolName, fn, priority?)

提交任务到协程池。

**参数**:
- `poolName` (string): 协程池名称
- `fn` (function): 要执行的函数
- `priority` (number, 可选): 任务优先级，默认为0

**返回值**: `boolean` - 是否成功提交

**调用示例**:
```lua
local success = coroutine_submitToPool("myPool", function()
    console.log("协程池任务执行")
    return "任务完成"
end, 0)
if success then
    console.log("任务已提交")
end
```

#### coroutine_getPoolStats(poolName)

获取协程池统计信息。

**参数**:
- `poolName` (string): 协程池名称

**返回值**: `table` - 协程池统计信息，包含 name、maxWorkers、maxTasks、active、queued、workers、closed

**调用示例**:
```lua
local poolStats = coroutine_getPoolStats("myPool")
console.log("协程池名称: " .. poolStats.name)
console.log("最大工作协程数: " .. poolStats.maxWorkers)
console.log("最大任务数: " .. poolStats.maxTasks)
console.log("活跃工作协程数: " .. poolStats.active)
console.log("队列中任务数: " .. poolStats.queued)
console.log("工作协程总数: " .. poolStats.workers)
console.log("是否已关闭: " .. tostring(poolStats.closed))
```

#### coroutine_closePool(poolName)

关闭协程池。

**参数**:
- `poolName` (string): 协程池名称

**返回值**: `boolean` - 是否成功关闭

**调用示例**:
```lua
local success = coroutine_closePool("myPool")
if success then
    console.log("协程池已关闭")
end
```

#### coroutine_listPools()

列出所有协程池的名称。

**返回值**: `table` - 协程池名称数组

**调用示例**:
```lua
local poolList = coroutine_listPools()
for i, poolName in ipairs(poolList) do
    console.log("协程池: " .. poolName)
end
```

### 调度器方法

#### coroutine_setScheduleStrategy(strategy)

设置调度策略。

**参数**:
- `strategy` (string): 调度策略（fifo 或 priority）

**返回值**: `nil`

**调用示例**:
```lua
coroutine_setScheduleStrategy("fifo")
```

#### coroutine_getScheduleStrategy()

获取当前调度策略。

**返回值**: `string` - 当前调度策略

**调用示例**:
```lua
local strategy = coroutine_getScheduleStrategy()
console.log("当前调度策略: " .. strategy)
```

#### coroutine_setPriority(name, priority)

设置协程优先级。

**参数**:
- `name` (string): 协程名称
- `priority` (number): 优先级数值

**返回值**: `nil`

**调用示例**:
```lua
coroutine_setPriority("myCoroutine", 10)
```

#### coroutine_getPriority(name)

获取协程优先级。

**参数**:
- `name` (string): 协程名称

**返回值**: `number` - 优先级数值

**调用示例**:
```lua
local priority = coroutine_getPriority("myCoroutine")
console.log("协程优先级: " .. priority)
```

## 完整示例

### 示例1: 并发执行多个任务

```lua
console.log("开始并发执行任务")

local task1 = coroutine_launch(function()
    console.log("任务1开始")
    coroutine_sleep(1000)
    console.log("任务1完成")
    return "任务1结果"
end, "task1", 1)

local task2 = coroutine_launch(function()
    console.log("任务2开始")
    coroutine_sleep(1500)
    console.log("任务2完成")
    return "任务2结果"
end, "task2", 2)

local task3 = coroutine_launch(function()
    console.log("任务3开始")
    coroutine_sleep(2000)
    console.log("任务3完成")
    return "任务3结果"
end, "task3", 3)

console.log("所有任务已启动")
```

### 示例2: 使用协程池处理批量任务

```lua
console.log("创建协程池处理批量任务")

local poolName = coroutine_createPool("batchPool", 3, 10)

for i = 0, 9 do
    local taskNum = i + 1
    coroutine_submitToPool("batchPool", function()
        console.log("处理任务 " .. taskNum)
        coroutine_sleep(500)
        console.log("任务 " .. taskNum .. " 完成")
    end, taskNum)
end

console.log("所有任务已提交到协程池")

coroutine_sleep(5000)

local poolStats = coroutine_getPoolStats("batchPool")
console.log("协程池状态: " .. poolStats.name)

coroutine_closePool("batchPool")
```

### 示例3: 协程状态监控

```lua
console.log("启动协程监控")

for i = 0, 4 do
    local taskNum = i + 1
    coroutine_launch(function()
        console.log("协程 " .. taskNum .. " 开始执行")
        coroutine_sleep(1000 * taskNum)
        console.log("协程 " .. taskNum .. " 执行完成")
    end, "monitorTask" .. taskNum, taskNum)
end

local function monitor()
    local activeCount = coroutine_getActiveCoroutines()
    local stats = coroutine_getStats()

    console.log("=== 协程状态 ===")
    console.log("活跃协程数: " .. activeCount)
    console.log("总任务数: " .. stats.totalTasks)
    console.log("已完成: " .. stats.completed)
    console.log("失败: " .. stats.failed)
    console.log("已取消: " .. stats.cancelled)

    if activeCount == 0 then
        console.log("所有协程已完成")
    end
end

-- 使用定时器监控（需要配合其他模块实现定时器功能）
monitor()
```

## 协程状态

协程有以下几种状态：

- `pending`: 等待中
- `running`: 运行中
- `completed`: 已完成
- `cancelled`: 已取消
- `error`: 错误
