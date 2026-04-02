# system 模块

## 模块简介

system 模块提供了系统级别的功能，包括进程管理、内存监控、系统设置等。

## 方法列表

### system.getPid
获取进程ID

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| processName | string | 进程名称 |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| pid | number | 进程ID |

**使用示例：**
```javascript
// 获取进程ID
var pid = system.getPid("com.android.systemui");
console.log("进程ID: " + pid);
```

---

### system.getMemoryUsage
获取内存使用

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| pid | number | 进程ID |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| memoryUsage | number | 内存使用量（MB） |

**使用示例：**
```javascript
// 获取内存使用
var memoryUsage = system.getMemoryUsage(1234);
console.log("内存使用: " + memoryUsage + " MB");
```

---

### system.getCpuUsage
获取CPU使用率

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| pid | number | 进程ID |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| cpuUsage | number | CPU使用率（百分比） |

**使用示例：**
```javascript
// 获取CPU使用率
var cpuUsage = system.getCpuUsage(1234);
console.log("CPU使用率: " + cpuUsage + "%");
```

---

### system.restartSelf
重启自身

**参数：**

无参数

**返回值：**

无返回值

**使用示例：**
```javascript
// 重启自身
system.restartSelf();
```

---

### system.setBootStart
设置开机自启

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| enable | boolean | 是否开启开机自启 |

**返回值：**

无返回值

**使用示例：**
```javascript
// 设置开机自启
system.setBootStart(true); // 开启开机自启
```

---

## 综合使用示例

### 示例1：基本使用
```javascript
// system 模块的基本使用示例
```