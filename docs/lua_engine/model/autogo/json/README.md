# json 模块

## 模块简介

json 模块提供了相关的功能支持。

## 方法列表

### json.stringify
将 Lua 值序列化为 JSON 字符串（自动判断数组或对象）

**使用示例：**
```lua
-- 调用 json.stringify 方法
json.stringify();
```

---

### json.stringifyArr
将 Lua 值强制序列化为 JSON 数组

**使用示例：**
```lua
-- 调用 json.stringifyArr 方法
json.stringifyArr();
```

---

### json.stringifyObj
将 Lua 值强制序列化为 JSON 对象

**使用示例：**
```lua
-- 调用 json.stringifyObj 方法
json.stringifyObj();
```

---

### json.parse
将 JSON 字符串解析为 Lua 值

**使用示例：**
```lua
-- 调用 json.parse 方法
json.parse();
```

---

### json.format
将 Lua 值格式化序列化为 JSON 字符串

**使用示例：**
```lua
-- 调用 json.format 方法
json.format();
```

---

## 综合使用示例

### 示例1：基本使用
```lua
-- json 模块的基本使用示例
```