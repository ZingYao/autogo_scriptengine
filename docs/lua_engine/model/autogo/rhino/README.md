# rhino 模块

## 模块简介

rhino 模块提供了 Rhino JavaScript 引擎的集成功能，用于执行 JavaScript 代码。

## 方法列表

### rhino.eval
执行指定的JavaScript脚本并返回结果

**参数：**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| contextId | string | 是 | JavaScript执行上下文ID，用于隔离不同的执行环境 |
| script | string | 是 | 要执行的JavaScript代码字符串 |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| result | string | JavaScript脚本的执行结果（字符串形式） |

**使用示例：**
```lua
-- 调用 rhino.eval 方法
local result = rhino.eval("context1", "1 + 2 + 3")
print("计算结果：" .. result)  -- 输出：计算结果：6

-- 执行更复杂的JavaScript代码
local jsCode = [[
    function add(a, b) {
        return a + b;
    }
    add(10, 20);
]]
local result2 = rhino.eval("context2", jsCode)
print("函数调用结果：" .. result2)  -- 输出：函数调用结果：30
```

---

## 综合使用示例

### 示例1：基本使用
```lua
-- rhino 模块的基本使用示例
```