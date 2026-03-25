// 主脚本文件

// 使用 require 引入工具函数模块
const utils = require('./utils');

// 测试工具函数
const sum = utils.add(5, 3);
const difference = utils.subtract(10, 4);

// 输出结果（console 已经通过 InjectAllMethods 注入）
console.log('5 + 3 = ' + sum);
console.log('10 - 4 = ' + difference);

// 导出主函数
function main() {
    console.log('Hello from JavaScript autogo style!');
    console.log('Sum: ' + sum);
    console.log('Difference: ' + difference);
    return 'Script executed successfully!';
}

// 执行主函数
main();