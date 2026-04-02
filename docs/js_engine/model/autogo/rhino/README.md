# rhino 模块

## 模块简介

rhino 模块提供了 Rhino JavaScript 引擎的集成功能，用于执行 JavaScript 代码。

## 方法列表

### rhino.eval
执行指定的JavaScript脚本并返回结果

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| script | string | 要执行的 JavaScript 脚本 |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| result | any | 执行结果 |

**使用示例：**
```javascript
// 执行 JavaScript 脚本
var result = rhino.eval("var x = 1 + 2;");
console.log("结果: " + result);
```

---

## 综合使用示例

### 示例1：基本使用
```javascript
// rhino 模块的基本使用示例
```