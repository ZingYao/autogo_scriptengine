// iOS 示例工具模块：仅演示用户自定义 JavaScript require。
function join(left, right) {
    return String(left) + String(right);
}

function dumpObject(value) {
    if (value === null || typeof value !== 'object') {
        return String(value);
    }
    return Object.keys(value).map(function(key) {
        return key + '=' + value[key];
    }).join(', ');
}

module.exports = {
    join: join,
    dumpObject: dumpObject
};
