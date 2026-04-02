# storages 模块

## 模块简介

storages 模块提供了本地存储功能，支持键值对的持久化存储。

## 方法列表

### storages.get
从指定表中获取键值

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| name | string | 表名称 |
| key | string | 键名 |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| value | any | 键对应的值 |

**使用示例：**
```javascript
// 从指定表中获取键值
var value = storages.get("config", "username");
console.log("用户名: " + value);
```

---

### storages.put
写入键值对

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| name | string | 表名称 |
| key | string | 键名 |
| value | any | 要存储的值 |

**返回值：**

无返回值

**使用示例：**
```javascript
// 写入键值对
storages.put("config", "username", "admin");
```

---

### storages.remove
删除指定键

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| name | string | 表名称 |
| key | string | 键名 |

**返回值：**

无返回值

**使用示例：**
```javascript
// 删除指定键
storages.remove("config", "username");
```

---

### storages.contains
判断键是否存在

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| name | string | 表名称 |
| key | string | 键名 |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| exists | boolean | 键是否存在 |

**使用示例：**
```javascript
// 判断键是否存在
var exists = storages.contains("config", "username");
console.log("键是否存在: " + exists);
```

---

### storages.getAll
获取所有键值对

**使用示例：**
```javascript
// 调用 storages.getAll 方法
storages.getAll();
```

---

### storages.clear
清空指定表数据

**使用示例：**
```javascript
// 调用 storages.clear 方法
storages.clear();
```

---

## 综合使用示例

### 示例1：存储数据
```javascript
storages.put("config", "username", "admin");
var username = storages.get("config", "username");
console.log("用户名: " + username);
```