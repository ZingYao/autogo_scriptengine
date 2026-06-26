-- AutoGo iOS Lua 风格示例：模块对象入口、复杂参数和返回值解析。
local utils = require('utils')

local function safeCall(title, fn)
    local ok, result = pcall(fn)
    if ok then
        console.log('[OK] ' .. title)
        return result
    end
    console.log('[SKIP] ' .. title .. ': ' .. tostring(result))
    return nil
end

local function logTable(prefix, value)
    console.log(prefix .. utils.dumpTable(value))
end

function main()
    console.log('Hello from Lua iOS autogo style')

    safeCall('app and device modules', function()
        console.log('self package: ' .. app.selfPackage())
        console.log('current package: ' .. app.currentPackage())
        console.log('screen: ' .. device.width .. 'x' .. device.height)

        local display = device.getDisplayInfo()
        logTable('display info: ', display)
    end)

    safeCall('files module and byte slice conversion', function()
        local path = files.path('autogo_ios_lua_example.txt')
        files.write(path, 'hello ios lua')
        files.append(path, '\nappend line')
        console.log('file exists: ' .. tostring(files.exists(path)))
        console.log('file name: ' .. files.getName(path))
        console.log('file content: ' .. files.read(path))

        files.writeBytes(path .. '.bin', {65, 66, 67})
        local bytes = files.readBytes(path .. '.bin')
        console.log('bytes length: ' .. tostring(#bytes))
    end)

    safeCall('https return table and map argument', function()
        local getResp = https.get('https://example.com', 5000)
        console.log('GET code: ' .. tostring(getResp.code))

        local postResp = https.post(
            'https://example.com/api',
            '{"hello":"ios-lua"}',
            {['Content-Type'] = 'application/json'},
            5000
        )
        console.log('POST code: ' .. tostring(postResp.code))
    end)

    safeCall('app list slice and struct return value', function()
        local appList = app.getList(false)
        if #appList > 0 then
            local first = appList[1]
            console.log('first app: ' .. first.packageName .. ' / ' .. first.appName)
        end
    end)

    safeCall('iOS safe app operation', function()
        app.openUrl('https://autogo.cc/ios/')
    end)

    safeCall('opencv and imgui object constructors', function()
        local point = opencv.newPoint2f(10, 20)
        console.log('opencv point: ' .. tostring(point))

        local mat = opencv.newMat()
        if mat ~= nil then
            console.log('opencv mat empty: ' .. tostring(mat.empty()))
            mat.close()
        end

        local vec2 = imgui.newVec2(10, 20)
        console.log('imgui vec2: ' .. tostring(vec2))
    end)

    safeCall('ocr object lifecycle placeholders', function()
        -- 按需传入模型路径；这里仅展示 iOS 模块对象入口和返回对象调用形态。
        -- local detector = yolo.new('/path/to/model')
        -- local results = detector.detectFromPath('/path/to/image.png')
        -- detector.close()
        console.log(utils.join('object methods stay on returned objects: ', 'yolo/opencv/imgui'))
    end)

    return 'iOS Lua script executed successfully!'
end

main()
