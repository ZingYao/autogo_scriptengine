-- AutoGo Lua 风格示例：模块对象入口、用户 require、复杂参数和返回值解析。
-- AutoGo 的 utils 是全局模块对象；用户工具模块建议使用业务化变量名，避免遮蔽全局 utils。
local helpers = require('utils')

local sum = helpers.add(5, 3)
local difference = helpers.subtract(10, 4)

local function safeCall(title, fn)
    local ok, result = pcall(fn)
    if ok then
        console.log('[OK] ' .. title)
        return result
    end
    console.log('[SKIP] ' .. title .. ': ' .. tostring(result))
    return nil
end

function main()
    console.log('Hello from Lua autogo style')
    console.log('helpers.add: ' .. sum)
    console.log('helpers.subtract: ' .. difference)

    safeCall('basic modules', function()
        console.log('screen: ' .. device.width .. 'x' .. device.height)
        console.log('current package: ' .. app.currentPackage())
        motion.click(100, 200)
    end)

    safeCall('https return table and map argument', function()
        local getResp = https.get('https://example.com', 5000)
        console.log('GET code: ' .. tostring(getResp.code))

        local postResp = https.post(
            'https://example.com/api',
            '{"hello":"autogo"}',
            {['Content-Type'] = 'application/json'},
            5000
        )
        console.log('POST code: ' .. tostring(postResp.code))
    end)

    safeCall('struct argument', function()
        app.startActivity({
            action = 'android.intent.action.VIEW',
            data = 'https://example.com',
            packageName = app.getBrowserPackage()
        })
    end)

    safeCall('slice and struct return value', function()
        local appList = app.getList(false)
        if #appList > 0 then
            console.log('first app: ' .. appList[1].packageName .. ' / ' .. appList[1].appName)
        end
    end)

    safeCall('callback argument', function()
        images.setCallback(function(x, y, color)
            console.log('image callback: ' .. x .. ',' .. y .. ',' .. color)
        end)
    end)

    safeCall('object lifecycle', function()
        local acc = uiacc.new()
        local node = acc.text('确定')
        if node ~= nil then
            node.click()
        end
    end)

    safeCall('opencv and imgui object constructors', function()
        local point = opencv.newPoint2f(10, 20)
        console.log('opencv point: ' .. tostring(point))

        local vec2 = imgui.newVec2(10, 20)
        console.log('imgui vec2: ' .. tostring(vec2))
    end)

    return 'Script executed successfully!'
end

main()
