-- Android 示例工具模块：仅演示用户自定义 Lua require。
local tools = {}

-- 计算两个数的和
function tools.add(a, b)
    return a + b
end

-- 计算两个数的差
function tools.subtract(a, b)
    return a - b
end

return tools
