# json 模块

## 模块简介

json 模块提供了相关的功能支持。

## 方法列表

### json.stringify
将 Lua 值序列化为 JSON 字符串（自动判断数组或对象）

**参数：**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| value | any | 是 | 要序列化的 Lua 值 |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| result | string | JSON 字符串 |
| error | error | 错误信息（失败时） |

**使用示例：**
```lua
-- 调用 json.stringify 方法
local result, err = json.stringify({name = "test", value = 123})
```

---

### json.stringifyArr
将 Lua 值强制序列化为 JSON 数组

**参数：**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| value | any | 是 | 要序列化的 Lua 值 |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| result | string | JSON 数组字符串 |
| error | error | 错误信息（失败时） |

**使用示例：**
```lua
-- 调用 json.stringifyArr 方法
local result, err = json.stringifyArr({1, 2, 3})
```

---

### json.stringifyObj
将 Lua 值强制序列化为 JSON 对象

**参数：**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| value | any | 是 | 要序列化的 Lua 值 |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| result | string | JSON 对象字符串 |
| error | error | 错误信息（失败时） |

**使用示例：**
```lua
-- 调用 json.stringifyObj 方法
local result, err = json.stringifyObj({name = "test", value = 123})
```

---

### json.parse
将 JSON 字符串解析为 Lua 值

**参数：**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| jsonStr | string | 是 | JSON 字符串 |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| result | any | 解析后的 Lua 值 |
| error | error | 错误信息（失败时） |

**使用示例：**
```lua
-- 调用 json.parse 方法
local result, err = json.parse('{"name":"test","value":123}')
```

---

### json.format
将 Lua 值格式化序列化为 JSON 字符串

**参数：**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| value | any | 是 | 要格式化的 Lua 值 |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| result | string | 格式化的 JSON 字符串 |
| error | error | 错误信息（失败时） |

**使用示例：**
```lua
-- 调用 json.format 方法
local result, err = json.format({name = "test", value = 123})
```

---

## 综合使用示例

### 示例1：基本使用
```lua
-- json 模块的基本使用示例
```