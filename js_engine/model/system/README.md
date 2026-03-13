# System 模块

System 模块提供了系统级别的功能，包括进程管理、资源监控和系统设置。

## 方法列表

### system.getPid(processName)
根据进程名称获取进程 ID。

**入参：**
- `processName`: 进程名称（字符串）

**出参：** 进程 ID（整数），未找到返回 0

**调用示例：**
```javascript
// 获取进程 ID
var pid = system.getPid("com.android.settings");
console.println("进程 ID:", pid);
```

### system.getMemoryUsage(pid)
获取指定进程的内存使用情况。

**入参：**
- `pid`: 进程 ID（整数）

**出参：** 内存使用量（MB，整数）

**调用示例：**
```javascript
// 获取内存使用
var pid = system.getPid("com.android.settings");
if (pid !== 0) {
    var memory = system.getMemoryUsage(pid);
    console.println("内存使用:", memory, "MB");
}
```

### system.getCpuUsage(pid)
获取指定进程的 CPU 使用率。

**入参：**
- `pid`: 进程 ID（整数）

**出参：** CPU 使用率（浮点数，百分比）

**调用示例：**
```javascript
// 获取 CPU 使用率
var pid = system.getPid("com.android.settings");
if (pid !== 0) {
    var cpu = system.getCpuUsage(pid);
    console.println("CPU 使用率:", cpu.toFixed(2), "%");
}
```

### system.restartSelf()
重启当前应用程序。

**入参：** 无

**出参：** 无

**调用示例：**
```javascript
// 重启应用
system.restartSelf();
```

### system.setBootStart(enable)
设置是否开机自启动。

**入参：**
- `enable`: 是否启用开机自启动（布尔值）

**出参：** 无

**调用示例：**
```javascript
// 启用开机自启动
system.setBootStart(true);

// 禁用开机自启动
system.setBootStart(false);
```

## 完整示例

```javascript
// 示例1：获取进程信息
function getProcessInfo() {
    var processName = "com.android.settings";
    var pid = system.getPid(processName);
    
    if (pid !== 0) {
        console.println("进程名称:", processName);
        console.println("进程 ID:", pid);
        
        var memory = system.getMemoryUsage(pid);
        console.println("内存使用:", memory, "MB");
        
        var cpu = system.getCpuUsage(pid);
        console.println("CPU 使用率:", cpu.toFixed(2), "%");
    } else {
        console.println("进程未找到:", processName);
    }
}

// 示例2：监控进程资源
function monitorProcess(processName, duration) {
    var pid = system.getPid(processName);
    
    if (pid === 0) {
        console.println("进程未找到:", processName);
        return;
    }
    
    console.println("开始监控进程:", processName);
    console.println("进程 ID:", pid);
    
    var interval = setInterval(function() {
        var memory = system.getMemoryUsage(pid);
        var cpu = system.getCpuUsage(pid);
        
        console.println("内存:", memory, "MB", "CPU:", cpu.toFixed(2), "%");
    }, 1000);
    
    // 持续监控指定时间
    setTimeout(function() {
        clearInterval(interval);
        console.println("监控结束");
    }, duration * 1000);
}

// 示例3：批量获取进程信息
function getMultipleProcessInfo() {
    var processes = [
        "com.android.systemui",
        "com.android.settings",
        "com.android.phone"
    ];
    
    for (var i = 0; i < processes.length; i++) {
        var processName = processes[i];
        var pid = system.getPid(processName);
        
        console.println("---", processName, "---");
        
        if (pid !== 0) {
            console.println("进程 ID:", pid);
            console.println("内存:", system.getMemoryUsage(pid), "MB");
            console.println("CPU:", system.getCpuUsage(pid).toFixed(2), "%");
        } else {
            console.println("进程未运行");
        }
    }
}

// 示例4：进程资源监控器
function ProcessMonitor() {
    this.processes = {};
    
    this.addProcess = function(processName) {
        var pid = system.getPid(processName);
        if (pid !== 0) {
            this.processes[processName] = pid;
            console.println("添加进程:", processName, "PID:", pid);
        } else {
            console.println("进程未找到:", processName);
        }
    };
    
    this.removeProcess = function(processName) {
        delete this.processes[processName];
        console.println("移除进程:", processName);
    };
    
    this.monitor = function() {
        console.println("=== 进程监控 ===");
        for (var name in this.processes) {
            var pid = this.processes[name];
            var memory = system.getMemoryUsage(pid);
            var cpu = system.getCpuUsage(pid);
            console.println(name, "- 内存:", memory, "MB", "CPU:", cpu.toFixed(2), "%");
        }
    };
    
    this.start = function(interval) {
        var self = this;
        this.intervalId = setInterval(function() {
            self.monitor();
        }, interval);
        console.println("监控已启动，间隔:", interval, "ms");
    };
    
    this.stop = function() {
        if (this.intervalId) {
            clearInterval(this.intervalId);
            console.println("监控已停止");
        }
    };
}

// 使用进程监控器
function useProcessMonitor() {
    var monitor = new ProcessMonitor();
    
    // 添加要监控的进程
    monitor.addProcess("com.android.systemui");
    monitor.addProcess("com.android.settings");
    
    // 启动监控
    monitor.start(2000);
    
    // 运行一段时间后停止
    setTimeout(function() {
        monitor.stop();
    }, 10000);
}

// 示例5：系统资源统计
function systemResourceStats() {
    var processes = [
        "com.android.systemui",
        "com.android.settings",
        "com.android.phone",
        "com.android.launcher"
    ];
    
    var totalMemory = 0;
    var totalCpu = 0;
    var count = 0;
    
    console.println("=== 系统资源统计 ===");
    
    for (var i = 0; i < processes.length; i++) {
        var processName = processes[i];
        var pid = system.getPid(processName);
        
        if (pid !== 0) {
            var memory = system.getMemoryUsage(pid);
            var cpu = system.getCpuUsage(pid);
            
            console.println(processName);
            console.println("  内存:", memory, "MB");
            console.println("  CPU:", cpu.toFixed(2), "%");
            
            totalMemory += memory;
            totalCpu += cpu;
            count++;
        }
    }
    
    if (count > 0) {
        console.println("--- 统计 ---");
        console.println("监控进程数:", count);
        console.println("总内存使用:", totalMemory, "MB");
        console.println("平均内存:", (totalMemory / count).toFixed(2), "MB");
        console.println("总 CPU 使用:", totalCpu.toFixed(2), "%");
        console.println("平均 CPU:", (totalCpu / count).toFixed(2), "%");
    }
}

// 示例6：开机自启动设置
function manageBootStart() {
    // 启用开机自启动
    console.println("启用开机自启动...");
    system.setBootStart(true);
    console.println("已启用");
    
    // 检查状态（这里只是示例，实际可能需要其他方法检查）
    console.println("当前状态: 已启用");
    
    // 禁用开机自启动
    console.println("禁用开机自启动...");
    system.setBootStart(false);
    console.println("已禁用");
}

// 示例7：应用重启
function restartApplication() {
    console.println("应用将在 3 秒后重启...");
    utils.sleep(3000);
    system.restartSelf();
}

// 调用示例
getProcessInfo();
monitorProcess("com.android.settings", 5);
getMultipleProcessInfo();
useProcessMonitor();
systemResourceStats();
manageBootStart();
// restartApplication(); // 取消注释以测试重启功能
```

## 注意事项

1. 获取进程信息需要相应的权限
2. 进程名称通常是应用的包名
3. 内存使用量单位为 MB
4. CPU 使用率是百分比形式，范围 0-100
5. 重启应用会终止当前进程并重新启动
6. 开机自启动需要系统权限，可能需要用户授权
7. 监控进程资源时，建议设置合理的采样间隔
8. 某些系统进程可能无法获取详细信息
9. 进程 ID 可能会变化，建议动态获取
10. 重启应用前建议保存重要数据
