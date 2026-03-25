-- 主脚本文件

-- 使用 require 引入工具函数模块（用户自定义的 Lua 文件）
local utils = require('utils')

-- console 已经通过 Go 注入为全局变量，无需 require
-- 测试工具函数
local sum = utils.add(5, 3)
local difference = utils.subtract(10, 4)

-- 输出结果
console.log('5 + 3 = ' .. sum)
console.log('10 - 4 = ' .. difference)

--