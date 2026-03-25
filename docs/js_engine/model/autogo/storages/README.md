# storages 模块

## 模块简介

storages 模块提供了本地存储功能，支持键值对的持久化存储。

## 方法列表

### storages.get
从指定表中获取键值

**使用示例：**
```javascript
// 调用 storages.get 方法
storages.get();
```

---

### storages.put
写入键值对

**使用示例：**
```javascript
// 调用 storages.put 方法
storages.put();
```

---

### storages.remove
删除指定键

**使用示例：**
```javascript
// 调用 storages.remove 方法
storages.remove();
```

---

### storages.contains
判断键是否存在

**使用示例：**
```javascript
// 调用 storages.contains 方法
storages.contains();
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