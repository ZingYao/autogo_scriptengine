# Rhino 模块

Rhino 模块提供了执行 JavaScript 脚本的功能，支持在指定的上下文中执行代码并获取结果。

## 方法列表

### rhino.eval(contextId, script)
执行指定的 JavaScript 脚本并返回结果。

**入参：**
- `contextId`: 上下文 ID（字符串）
- `script`: 要执行的 JavaScript 脚本（字符串）

**出参：** 脚本执行结果

**调用示例：**
```javascript
// 执行简单的 JavaScript 代码
var result = rhino.eval("context1", "1 + 1");
console.println("计算结果:", result);

// 执行复杂的脚本
var script = "function add(a, b) { return a + b; } add(10, 20);";
var result2 = rhino.eval("context2", script);
console.println("函数调用结果:", result2);
```

## 完整示例

```javascript
// 示例1：基本计算
function basicCalculation() {
    var result = rhino.eval("calc", "2 * 3 + 4");
    console.println("计算结果:", result);
}

// 示例2：执行函数
function executeFunction() {
    var script = "function greet(name) { return 'Hello, ' + name + '!'; } greet('World');";
    var result = rhino.eval("greet", script);
    console.println("函数结果:", result);
}

// 示例3：使用变量
function useVariables() {
    var script = "var x = 10; var y = 20; x + y;";
    var result = rhino.eval("vars", script);
    console.println("变量计算结果:", result);
}

// 示例4：执行复杂逻辑
function complexLogic() {
    var script = "var sum = 0; for (var i = 1; i <= 10; i++) { sum += i; } sum;";
    var result = rhino.eval("loop", script);
    console.println("循环求和结果:", result);
}

// 示例5：使用对象
function useObjects() {
    var script = "var obj = { name: '张三', age: 25 }; JSON.stringify(obj);";
    var result = rhino.eval("obj", script);
    console.println("对象结果:", result);
}

// 示例6：使用数组
function useArrays() {
    var script = "var arr = [1, 2, 3, 4, 5]; arr.reduce(function(a, b) { return a + b; }, 0);";
    var result = rhino.eval("array", script);
    console.println("数组求和结果:", result);
}

// 示例7：条件判断
function conditionalLogic() {
    var script = "var score = 85; score >= 60 ? '及格' : '不及格';";
    var result = rhino.eval("condition", script);
    console.println("判断结果:", result);
}

// 示例8：字符串操作
function stringOperations() {
    var script = "'Hello World'.toUpperCase()";
    var result = rhino.eval("string", script);
    console.println("字符串操作结果:", result);
}

// 示例9：数学运算
function mathOperations() {
    var script = "Math.max(10, 20, 30, 40, 50)";
    var result = rhino.eval("math", script);
    console.println("数学运算结果:", result);
}

// 示例10：多个上下文
function multipleContexts() {
    var script1 = "var x = 100; x * 2;";
    var result1 = rhino.eval("context1", script1);
    console.println("上下文1结果:", result1);
    
    var script2 = "var y = 200; y / 2;";
    var result2 = rhino.eval("context2", script2);
    console.println("上下文2结果:", result2);
}

// 示例11：错误处理
function errorHandling() {
    try {
        var script = "throw new Error('测试错误');";
        var result = rhino.eval("error", script);
        console.println("结果:", result);
    } catch (e) {
        console.println("捕获到错误:", e.message);
    }
}

// 示例12：动态脚本执行
function dynamicScriptExecution() {
    var operations = [
        "10 + 20",
        "30 - 15",
        "5 * 6",
        "100 / 4"
    ];
    
    for (var i = 0; i < operations.length; i++) {
        var result = rhino.eval("dynamic", operations[i]);
        console.println("操作 " + (i + 1) + ":", operations[i], "=", result);
    }
}

// 示例13：使用闭包
function useClosure() {
    var script = "function counter() { var count = 0; return function() { count++; return count; }; } var c = counter(); c(); c(); c();";
    var result = rhino.eval("closure", script);
    console.println("闭包结果:", result);
}

// 示例14：递归函数
function recursiveFunction() {
    var script = "function factorial(n) { return n <= 1 ? 1 : n * factorial(n - 1); } factorial(5);";
    var result = rhino.eval("recursive", script);
    console.println("递归结果:", result);
}

// 调用示例
basicCalculation();
executeFunction();
useVariables();
complexLogic();
useObjects();
useArrays();
conditionalLogic();
stringOperations();
mathOperations();
multipleContexts();
errorHandling();
dynamicScriptExecution();
useClosure();
recursiveFunction();
```

## 注意事项

1. contextId 用于标识不同的执行上下文，相同 contextId 的脚本共享变量
2. 执行的脚本必须是有效的 JavaScript 代码
3. 脚本执行可能会抛出异常，需要做好错误处理
4. 复杂的脚本可能需要较长的执行时间
5. 建议在执行脚本前进行语法检查
6. 不同上下文之间的变量是隔离的
7. 脚本执行结果会自动转换为 JavaScript 值
8. 避免在脚本中执行无限循环或递归过深
9. 脚本中可以访问 Rhino 引擎提供的所有功能
10. 建议根据实际需求合理使用不同的上下文 ID
