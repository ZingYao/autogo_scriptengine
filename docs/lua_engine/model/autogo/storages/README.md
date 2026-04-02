# storages 模块

## 模块简介

storages 模块提供了本地存储功能，支持键值对的持久化存储。

## 方法列表

### storages.get
从指定表中获取键值

**参数：**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| table | string | 是 | 存储表名称 |
| key | string | 是 | 键名 |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| result | string | 键对应的值 |

**使用示例：**
```lua
-- 调用 storages.get 方法
local value = storages.get("config", "username")
```

---

### storages.put
写入键值对

**参数：**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| table | string | 是 | 存储表名称 |
| key | string | 是 | 键名 |
| value | string | 是 | 键值 |

**返回值：**

无

**使用示例：**
```lua
-- 调用 storages.put 方法
storages.put("config", "username", "admin")
```

---

### storages.remove
删除指定键

**参数：**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| table | string | 是 | 存储表名称 |
| key | string | 是 | 键名 |

**返回值：**

无

**使用示例：**
```lua
-- 调用 storages.remove 方法
storages.remove("config", "username")
```

---

### storages.contains
判断键是否存在

**参数：**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| table | string | 是 | 存储表名称 |
| key | string | 是 | 键名 |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| result | boolean | 键是否存在 |

**使用示例：**
```lua
-- 调用 storages.contains 方法
local exists = storages.contains("config", "username")
```

---

### storages.getAll
获取所有键值对

**参数：**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| table | string | 是 | 存储表名称 |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| result | table | 包含所有键值对的Lua表 |

**使用示例：**
```lua
-- 调用 storages.getAll 方法
local allData = storages.getAll("config")
```

---

### storages.clear
清空指定表数据

**参数：**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| table | string | 是 | 存储表名称 |

**返回值：**

无

**使用示例：**
```lua
-- 调用 storages.clear 方法
storages.clear("config")
```

---

## 综合使用示例

### 示例1：存储数据
```lua
storages.put("config", "username", "admin");
var username = storages.get("config", "username");
console.log("用户名: " + username);
```