# Storages 模块

Storages 模块提供了简单的键值对存储功能，用于持久化数据。

## 方法列表

### storages.get(table, key)
从指定表中获取键值。

**入参：**
- `table`: 表名（字符串）
- `key`: 键名（字符串）

**出参：** 键对应的值（字符串）

**调用示例：**
```javascript
// 获取值
var value = storages.get("user", "name");
console.println("用户名:", value);
```

### storages.put(table, key, value)
写入键值对。

**入参：**
- `table`: 表名（字符串）
- `key`: 键名（字符串）
- `value`: 值（字符串）

**出参：** 无

**调用示例：**
```javascript
// 写入键值对
storages.put("user", "name", "张三");
storages.put("user", "age", "25");
```

### storages.remove(table, key)
删除指定键。

**入参：**
- `table`: 表名（字符串）
- `key`: 键名（字符串）

**出参：** 无

**调用示例：**
```javascript
// 删除键
storages.remove("user", "name");
```

### storages.contains(table, key)
判断键是否存在。

**入参：**
- `table`: 表名（字符串）
- `key`: 键名（字符串）

**出参：** 布尔值，true 表示存在，false 表示不存在

**调用示例：**
```javascript
// 检查键是否存在
var exists = storages.contains("user", "name");
console.println("键存在:", exists);
```

### storages.getAll(table)
获取所有键值对。

**入参：**
- `table`: 表名（字符串）

**出参：** 包含所有键值对的对象

**调用示例：**
```javascript
// 获取所有键值对
var allData = storages.getAll("user");
console.println("所有数据:", JSON.stringify(allData));
```

### storages.clear(table)
清空指定表数据。

**入参：**
- `table`: 表名（字符串）

**出参：** 无

**调用示例：**
```javascript
// 清空表
storages.clear("user");
```

## 完整示例

```javascript
// 示例1：基本存储操作
function basicStorage() {
    // 写入数据
    storages.put("user", "name", "张三");
    storages.put("user", "age", "25");
    storages.put("user", "email", "zhangsan@example.com");
    
    // 读取数据
    var name = storages.get("user", "name");
    var age = storages.get("user", "age");
    var email = storages.get("user", "email");
    
    console.println("姓名:", name);
    console.println("年龄:", age);
    console.println("邮箱:", email);
}

// 示例2：检查键是否存在
function checkKeyExists() {
    var key = "name";
    var exists = storages.contains("user", key);
    
    if (exists) {
        var value = storages.get("user", key);
        console.println("键存在，值为:", value);
    } else {
        console.println("键不存在");
    }
}

// 示例3：获取所有数据
function getAllData() {
    var allData = storages.getAll("user");
    
    console.println("所有用户数据:");
    for (var key in allData) {
        console.println(key + ":", allData[key]);
    }
}

// 示例4：删除数据
function deleteData() {
    // 删除单个键
    storages.remove("user", "email");
    console.println("已删除邮箱");
    
    // 检查是否删除成功
    var exists = storages.contains("user", "email");
    console.println("邮箱键存在:", exists);
}

// 示例5：清空表
function clearTable() {
    console.println("清空前:");
    var before = storages.getAll("user");
    console.println("数据量:", Object.keys(before).length);
    
    // 清空表
    storages.clear("user");
    
    console.println("清空后:");
    var after = storages.getAll("user");
    console.println("数据量:", Object.keys(after).length);
}

// 示例6：配置管理
function ConfigManager() {
    this.table = "config";
    
    this.set = function(key, value) {
        storages.put(this.table, key, value);
    };
    
    this.get = function(key, defaultValue) {
        if (storages.contains(this.table, key)) {
            return storages.get(this.table, key);
        }
        return defaultValue;
    };
    
    this.remove = function(key) {
        storages.remove(this.table, key);
    };
    
    this.getAll = function() {
        return storages.getAll(this.table);
    };
    
    this.clear = function() {
        storages.clear(this.table);
    };
}

// 使用配置管理器
function useConfigManager() {
    var config = new ConfigManager();
    
    // 设置配置
    config.set("theme", "dark");
    config.set("language", "zh-CN");
    config.set("fontSize", "16");
    
    // 获取配置
    var theme = config.get("theme", "light");
    var language = config.get("language", "en-US");
    var fontSize = config.get("fontSize", "14");
    
    console.println("主题:", theme);
    console.println("语言:", language);
    console.println("字体大小:", fontSize);
    
    // 获取所有配置
    var allConfig = config.getAll();
    console.println("所有配置:", JSON.stringify(allConfig));
}

// 示例7：用户会话管理
function SessionManager() {
    this.table = "session";
    
    this.saveSession = function(userId, token) {
        storages.put(this.table, "userId", userId);
        storages.put(this.table, "token", token);
        storages.put(this.table, "loginTime", new Date().toISOString());
    };
    
    this.getSession = function() {
        if (!storages.contains(this.table, "userId")) {
            return null;
        }
        
        return {
            userId: storages.get(this.table, "userId"),
            token: storages.get(this.table, "token"),
            loginTime: storages.get(this.table, "loginTime")
        };
    };
    
    this.clearSession = function() {
        storages.clear(this.table);
    };
    
    this.isLoggedIn = function() {
        return storages.contains(this.table, "userId");
    };
}

// 使用会话管理器
function useSessionManager() {
    var session = new SessionManager();
    
    // 保存会话
    session.saveSession("12345", "abc123xyz");
    console.println("会话已保存");
    
    // 检查登录状态
    var loggedIn = session.isLoggedIn();
    console.println("已登录:", loggedIn);
    
    // 获取会话信息
    var sessionInfo = session.getSession();
    console.println("用户ID:", sessionInfo.userId);
    console.println("令牌:", sessionInfo.token);
    console.println("登录时间:", sessionInfo.loginTime);
    
    // 清除会话
    session.clearSession();
    console.println("会话已清除");
    
    // 再次检查登录状态
    loggedIn = session.isLoggedIn();
    console.println("已登录:", loggedIn);
}

// 示例8：计数器
function Counter() {
    this.table = "counter";
    this.key = "count";
    
    this.increment = function() {
        var current = this.get();
        var next = current + 1;
        storages.put(this.table, this.key, next.toString());
        return next;
    };
    
    this.decrement = function() {
        var current = this.get();
        var next = current - 1;
        storages.put(this.table, this.key, next.toString());
        return next;
    };
    
    this.get = function() {
        if (storages.contains(this.table, this.key)) {
            return parseInt(storages.get(this.table, this.key));
        }
        return 0;
    };
    
    this.reset = function() {
        storages.put(this.table, this.key, "0");
    };
}

// 使用计数器
function useCounter() {
    var counter = new Counter();
    
    console.println("初始计数:", counter.get());
    
    console.println("增加后:", counter.increment());
    console.println("增加后:", counter.increment());
    console.println("增加后:", counter.increment());
    
    console.println("减少后:", counter.decrement());
    
    console.println("重置前:", counter.get());
    counter.reset();
    console.println("重置后:", counter.get());
}

// 调用示例
basicStorage();
checkKeyExists();
getAllData();
deleteData();
clearTable();
useConfigManager();
useSessionManager();
useCounter();
```

## 注意事项

1. 所有存储的值都是字符串类型，其他类型需要转换
2. 表名用于区分不同的数据集合
3. 删除不存在的键不会报错
4. 获取不存在的键会返回空字符串
5. 清空表会删除表中的所有数据
6. 建议使用有意义的表名和键名
7. 存储的数据会持久化保存
8. 对于复杂对象，建议使用 JSON.stringify() 和 JSON.parse() 进行转换
9. 不同表之间的数据是隔离的
10. 建议定期清理不需要的数据
