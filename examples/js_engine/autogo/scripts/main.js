// AutoGo JavaScript 风格示例：模块对象入口、复杂参数和返回值解析
const utils = require('./utils');

const sum = utils.add(5, 3);
const difference = utils.subtract(10, 4);

function safeCall(title, fn) {
    try {
        const result = fn();
        console.log('[OK] ' + title);
        return result;
    } catch (err) {
        console.log('[SKIP] ' + title + ': ' + err);
        return null;
    }
}

function main() {
    console.log('Hello from JavaScript autogo style');
    console.log('utils.add: ' + sum);
    console.log('utils.subtract: ' + difference);

    safeCall('basic modules', function() {
        console.log('screen: ' + device.width + 'x' + device.height);
        console.log('current package: ' + app.currentPackage());
        motion.click(100, 200);
    });

    safeCall('https return object and map argument', function() {
        const getResp = https.get('https://example.com', 5000);
        console.log('GET code: ' + getResp.code);

        const postResp = https.post(
            'https://example.com/api',
            JSON.stringify({ hello: 'autogo' }),
            { 'Content-Type': 'application/json' },
            5000
        );
        console.log('POST code: ' + postResp.code);
    });

    safeCall('struct argument', function() {
        app.startActivity({
            action: 'android.intent.action.VIEW',
            data: 'https://example.com',
            packageName: app.getBrowserPackage()
        });
    });

    safeCall('slice and struct return value', function() {
        const appList = app.getList(false);
        if (appList.length > 0) {
            console.log('first app: ' + appList[0].packageName + ' / ' + appList[0].appName);
        }
    });

    safeCall('callback argument', function() {
        images.setCallback(function(x, y, color) {
            console.log('image callback: ' + x + ',' + y + ',' + color);
        });
    });

    safeCall('object lifecycle', function() {
        const acc = uiacc.new();
        const node = acc.text('确定');
        if (node) {
            node.click();
        }
    });

    safeCall('opencv and imgui object constructors', function() {
        const point = opencv.newPoint2f(10, 20);
        console.log('opencv point: ' + point);

        const vec2 = imgui.newVec2(10, 20);
        console.log('imgui vec2: ' + vec2);
    });

    return 'Script executed successfully!';
}

main();
