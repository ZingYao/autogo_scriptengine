-- 主脚本文件

-- 使用 require 引入工具函数模块
local utils = require('utils')

-- 测试工具函数
local sum = utils.add(5,3)
local difference = utils.subtract(10, 4)

-- 输出结果（console 已经通过 InjectAllMethods 注入为全局变量）
console.log('5 + 3 = ' .. sum)
console.log('10 - 4 = ' .. difference)

-- 主函数
function main()
    console.log('Hello from Lua autogo style!')
    console.log('Sum: ' .. sum)
    console.log('Difference: ' .. difference)
    return 'Script executed successfully!'
end

-- 执行主函数
main()