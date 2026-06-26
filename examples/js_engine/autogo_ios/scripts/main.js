// AutoGo iOS JavaScript 风格示例：模块对象入口、复杂参数和返回值解析。
const utils = require('./utils');

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
    console.log('Hello from JavaScript iOS autogo style');

    safeCall('app and device modules', function() {
        console.log('self package: ' + app.selfPackage());
        console.log('current package: ' + app.currentPackage());
        console.log('screen: ' + device.width + 'x' + device.height);

        const display = device.getDisplayInfo();
        console.log('display info: ' + utils.dumpObject(display));
    });

    safeCall('files module and byte array conversion', function() {
        const path = files.path('autogo_ios_js_example.txt');
        files.write(path, 'hello ios js');
        files.append(path, '\nappend line');
        console.log('file exists: ' + files.exists(path));
        console.log('file name: ' + files.getName(path));
        console.log('file content: ' + files.read(path));

        files.writeBytes(path + '.bin', [65, 66, 67]);
        const bytes = files.readBytes(path + '.bin');
        console.log('bytes length: ' + bytes.length);
    });

    safeCall('https return object and map argument', function() {
        const getResp = https.get('https://example.com', 5000);
        console.log('GET code: ' + getResp.code);

        const postResp = https.post(
            'https://example.com/api',
            JSON.stringify({ hello: 'ios-js' }),
            { 'Content-Type': 'application/json' },
            5000
        );
        console.log('POST code: ' + postResp.code);
    });

    safeCall('app list array and struct return value', function() {
        const appList = app.getList(false);
        if (appList.length > 0) {
            const first = appList[0];
            console.log('first app: ' + first.packageName + ' / ' + first.appName);
        }
    });

    safeCall('iOS safe app operation', function() {
        app.openUrl('https://autogo.cc/ios/');
    });

    safeCall('opencv and imgui object constructors', function() {
        const point = opencv.newPoint2f(10, 20);
        console.log('opencv point: ' + point);

        const mat = opencv.newMat();
        if (mat) {
            console.log('opencv mat empty: ' + mat.empty());
            mat.close();
        }

        const vec2 = imgui.newVec2(10, 20);
        console.log('imgui vec2: ' + vec2);
    });

    safeCall('ocr object lifecycle placeholders', function() {
        // 按需传入模型路径；这里仅展示 iOS 模块对象入口和返回对象调用形态。
        // const detector = yolo.new('/path/to/model');
        // const results = detector.detectFromPath('/path/to/image.png');
        // detector.close();
        console.log(utils.join('object methods stay on returned objects: ', 'yolo/opencv/imgui'));
    });

    return 'iOS JavaScript script executed successfully!';
}

main();
