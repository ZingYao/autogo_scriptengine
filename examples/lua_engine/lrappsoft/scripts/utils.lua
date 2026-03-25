-- 工具函数模块

-- 计算两个数的和
function add(a, b)
    return a + b
end

-- 计算两个数的差
function subtract(a, b)
    return a - b
end

-- 返回模块
tools = {
    add = add,
    subtract = subtract
}
return tools