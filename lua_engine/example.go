package lua_engine

import (
	"fmt"
	"os"
)

func main() {
	// 初始化 Lua 引擎
	engine := GetEngine()
	defer Close()

	// 生成 Lua API 文档
	docGen := NewDocumentationGenerator()

	// 生成 Lua 文档
	fmt.Println("生成 Lua API 文档...")
	err := docGen.SaveLuaDocumentation("lua_api.lua")
	if err != nil {
		fmt.Printf("生成 Lua 文档失败: %v\n", err)
		return
	}
	fmt.Println("✓ Lua 文档已生成: lua_api.lua")

	// 生成 Markdown 文档
	fmt.Println("生成 Markdown API 文档...")
	err = docGen.SaveMarkdownDocumentation("lua_api.md")
	if err != nil {
		fmt.Printf("生成 Markdown 文档失败: %v\n", err)
		return
	}
	fmt.Println("✓ Markdown 文档已生成: lua_api.md")

	// 导出方法列表为 JSON
	fmt.Println("导出方法列表...")
	registry := engine.GetRegistry()
	jsonData, err := registry.ExportMethodsJSON()
	if err != nil {
		fmt.Printf("导出方法列表失败: %v\n", err)
		return
	}

	err = os.WriteFile("methods.json", []byte(jsonData), 0644)
	if err != nil {
		fmt.Printf("写入方法列表失败: %v\n", err)
		return
	}
	fmt.Println("✓ 方法列表已导出: methods.json")

	// 导出方法列表为 Lua 表
	fmt.Println("导出方法列表为 Lua 表...")
	luaTable := registry.ExportMethodsLuaTable()
	err = os.WriteFile("methods.lua", []byte(luaTable), 0644)
	if err != nil {
		fmt.Printf("写入 Lua 表失败: %v\n", err)
		return
	}
	fmt.Println("✓ Lua 表已导出: methods.lua")

	fmt.Println("\n所有文档生成完成!")

	// ========== 基础示例 ==========
	fmt.Println("\n========== 基础示例 ==========")

	// 示例1: 基础 Lua 脚本执行
	fmt.Println("\n--- 示例1: 基础 Lua 脚本执行 ---")
	err = ExecuteString(`
		console.log("=== 基础 Lua 脚本执行 ===")
		
		-- 获取当前应用包名
		local packageName = app_currentPackage()
		console.log("当前应用包名: " .. packageName)
		
		-- 获取设备信息
		local width = device.width(0)
		local height = device.height(0)
		console.log("设备分辨率: " .. width .. "x" .. height)
		console.log("SDK 版本: " .. device.sdkInt)
		
		-- 点击屏幕
		click(width/2, height/2, 1, 0)
		console.log("点击屏幕中心")
		
		-- 延迟
		sleep(1000)
		console.log("延迟 1 秒后")
	`)
	if err != nil {
		fmt.Printf("执行基础示例失败: %v\n", err)
	}

	// 示例2: 文件操作
	fmt.Println("\n--- 示例2: 文件操作 ---")
	err = ExecuteString(`
		console.log("=== 文件操作示例 ===")
		
		-- 检查文件是否存在
		if files_exists("/sdcard/test.txt") then
			console.log("文件存在")
			local content = files_read("/sdcard/test.txt")
			console.log("文件内容: " .. content)
		else
			console.log("文件不存在，创建新文件")
			files_write("/sdcard/test.txt", "Hello from Lua!")
			console.log("文件已创建")
		end
		
		-- 列出目录
		local fileList = files_listDir("/sdcard")
		console.log("目录文件数量: " .. #fileList)
	`)
	if err != nil {
		fmt.Printf("执行文件操作示例失败: %v\n", err)
	}

	// 示例3: 图像处理
	fmt.Println("\n--- 示例3: 图像处理 ---")
	err = ExecuteString(`
		console.log("=== 图像处理示例 ===")
		
		-- 获取像素颜色
		local color = images_pixel(500, 1000, 0)
		console.log("像素颜色: " .. color)
		
		-- 比较颜色
		if images_cmpColor(500, 1000, "#FF0000", 0.9, 0) then
			console.log("颜色匹配")
		else
			console.log("颜色不匹配")
		end
		
		-- 查找颜色
		local x, y = images_findColor(0, 0, 1080, 1920, "#FF0000", 0.9, 0, 0)
		if x ~= -1 and y ~= -1 then
			console.log("找到颜色在: " .. x .. ", " .. y)
			click(x, y, 1, 0)
		else
			console.log("未找到颜色")
		end
	`)
	if err != nil {
		fmt.Printf("执行图像处理示例失败: %v\n", err)
	}

	// 示例4: 存储操作
	fmt.Println("\n--- 示例4: 存储操作 ---")
	err = ExecuteString(`
		console.log("=== 存储操作示例 ===")
		
		-- 存储数据
		storages_put("myTable", "key1", "value1")
		storages_put("myTable", "key2", "value2")
		console.log("数据已存储")
		
		-- 读取数据
		local value1 = storages_get("myTable", "key1")
		console.log("key1 = " .. value1)
		
		-- 检查键是否存在
		if storages_contains("myTable", "key2") then
			console.log("key2 存在")
		end
		
		-- 获取所有数据
		local allData = storages_getAll("myTable")
		for k, v in pairs(allData) do
			console.log(k .. " = " .. v)
		end
	`)
	if err != nil {
		fmt.Printf("执行存储操作示例失败: %v\n", err)
	}

	// 示例5: 网络请求
	fmt.Println("\n--- 示例5: 网络请求 ---")
	err = ExecuteString(`
		console.log("=== 网络请求示例 ===")
		
		-- 发送 GET 请求
		local code, data = http_get("https://example.com", 5000)
		console.log("状态码: " .. code)
		console.log("响应长度: " .. string.len(data))
	`)
	if err != nil {
		fmt.Printf("执行网络请求示例失败: %v\n", err)
	}

	// 示例6: 文字识别
	fmt.Println("\n--- 示例6: 文字识别 ---")
	err = ExecuteString(`
		console.log("=== 文字识别示例 ===")
		
		-- 识别屏幕文字
		local results = ppocr_ocr(0, 0, 1080, 1920, "", 0)
		console.log("识别到 " .. #results .. " 个文本")
		
		for i, result in ipairs(results) do
			console.log("文本: " .. result["标签"])
			console.log("位置: (" .. result["X"] .. ", " .. result["Y"] .. ")")
			console.log("大小: " .. result["宽"] .. "x" .. result["高"])
			console.log("精度: " .. result["精度"])
			console.log("中心: (" .. result["CenterX"] .. ", " .. result["CenterY"] .. ")")
		end
	`)
	if err != nil {
		fmt.Printf("执行文字识别示例失败: %v\n", err)
	}

	// 示例7: 方法管理
	fmt.Println("\n--- 示例7: 方法管理 ---")
	err = ExecuteString(`
		console.log("=== 方法管理示例 ===")
		
		-- 列出所有方法
		local methods = listMethods()
		console.log("已注册方法数量: " .. #methods)
		
		-- 注册自定义方法
		registerMethod("myCustomMethod", "我的自定义方法", true)
		
		function myCustomMethod(param)
			console.log("自定义方法被调用: " .. param)
			return "返回值: " .. param
		end
		
		-- 使用自定义方法
		local result = myCustomMethod("测试参数")
		console.log(result)
		
		-- 重写方法
		local originalClick = click
		
		function click(x, y, fingerID, displayId)
			console.log("点击: (" .. x .. ", " .. y .. ")")
			originalClick(x, y, fingerID, displayId)
		end
		
		-- 测试重写的方法
		click(100, 200, 1, 0)
		
		-- 恢复方法
		restoreMethod("click")
		console.log("方法已恢复")
	`)
	if err != nil {
		fmt.Printf("执行方法管理示例失败: %v\n", err)
	}

	fmt.Println("\n========== 基础示例完成 ==========")
}
