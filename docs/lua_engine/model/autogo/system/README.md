# system 模块

## 模块简介

system 模块提供了系统级别的功能，包括进程管理、内存监控、系统设置等。

## 方法列表

### system.getPid
获取进程ID

**参数：**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| processName | string | 是 | 进程名称 |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| result | number | 进程ID |

**使用示例：**
```lua
-- 调用 system.getPid 方法
local pid = system.getPid("com.example.app")
```

---

### system.getMemoryUsage
获取内存使用

**参数：**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| pid | number | 是 | 进程ID |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| result | number | 内存使用量（KB） |

**使用示例：**
```lua
-- 调用 system.getMemoryUsage 方法
local memory = system.getMemoryUsage(12345)
```

---

### system.getCpuUsage
获取CPU使用率

**参数：**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| pid | number | 是 | 进程ID |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| result | number | CPU使用率（百分比） |

**使用示例：**
```lua
-- 调用 system.getCpuUsage 方法
local cpu = system.getCpuUsage(12345)
```

---

### system.restartSelf
重启自身

**参数：**

无

**返回值：**

无

**使用示例：**
```lua
-- 调用 system.restartSelf 方法
system.restartSelf()
```

---

### system.setBootStart
设置开机自启

**参数：**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| enable | boolean | 是 | 是否开启自启 |

**返回值：**

无

**使用示例：**
```lua
-- 调用 system.setBootStart 方法
system.setBootStart(true)
```

---

## 综合使用示例

### 示例1：基本使用
```lua
-- system 模块的基本使用示例
```