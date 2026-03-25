-- 主脚本文件

-- 使用 load 函数加载工具函数模块
load('scripts/utils.lua')

-- 使用 require 引入 lrappsoft 模块
local console = require('console')

-- 测试工具函数
local sum = utils.add(5, 3)
local difference = utils.subtract(10, 4)

-- 输出结果
console.log('5 + 3 = ' .. sum)
console.log('10 - 4 = ' .. difference)

--