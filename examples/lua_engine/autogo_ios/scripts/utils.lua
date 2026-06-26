-- iOS 示例工具模块：仅演示用户自定义 Lua require。
local tools = {}

function tools.join(left, right)
    return tostring(left) .. tostring(right)
end

function tools.dumpTable(value)
    if type(value) ~= 'table' then
        return tostring(value)
    end
    local parts = {}
    for key, item in pairs(value) do
        parts[#parts + 1] = tostring(key) .. '=' .. tostring(item)
    end
    return table.concat(parts, ', ')
end

return tools
