package main

import (
	"embed"
	"fmt"
	"os"
	"time"

	"github.com/ZingYao/autogo_scriptengine/js_engine"
	jsAppModel "github.com/ZingYao/autogo_scriptengine/js_engine/model/app_models"
	jsDeviceModel "github.com/ZingYao/autogo_scriptengine/js_engine/model/device_models"
	luaAppModel "github.com/ZingYao/autogo_scriptengine/lua_engine/model/app_models"
	luaDeviceModel "github.com/ZingYao/autogo_scriptengine/lua_engine/model/device_models"

	"github.com/ZingYao/autogo_scriptengine/lua_engine"
)

//go:embed scripts/*
var scriptsFS embed.FS

// 默认屏幕ID
const defaultDisplayId = 0

func init() {
	// 注册 JavaScript 引擎模块
	// 方式1: 使用 define 包中的预定义模块数组
	// 注册核心模块（排除 console、hud、imgui、vdisplay）
	js_engine.RegisterModule(jsAppModel.AppModels, jsDeviceModel.DeviceModels)

	// 注册 Lua 引擎模块
	// 方式1: 使用 define 包中的预定义模块数组
	// 注册核心模块（排除 console、hud、imgui、vdisplay）
	lua_engine.RegisterModule(luaAppModel.AppModels, luaDeviceModel.DeviceModels)
}

func main() {
	fmt.Println("╔════════════════════════════════════════════════════════════╗")
	fmt.Println("║       AutoGo 脚本引擎全量测试                               ║")
	fmt.Println("║       JavaScript Engine / Lua Engine / Native Go           ║")
	fmt.Println("╚════════════════════════════════════════════════════════════╝")
	fmt.Println()

	// 测试 JavaScript 引擎
	// fmt.Println("┌──────────────────────────────────────────────────────────┐")
	// fmt.Println("│  1. JavaScript 引擎测试                                   │")
	// fmt.Println("└──────────────────────────────────────────────────────────┘")
	// testJSEngine()
	// fmt.Println()

	// 测试 Lua 引擎
	fmt.Println("┌──────────────────────────────────────────────────────────┐")
	fmt.Println("│  2. Lua 引擎测试                                        │")
	fmt.Println("└──────────────────────────────────────────────────────────┘")
	testLuaEngine()
	fmt.Println()

	fmt.Println("╔════════════════════════════════════════════════════════════╗")
	fmt.Println("║       所有测试完成                                         ║")
	fmt.Println("╚════════════════════════════════════════════════════════════╝")
}

// testJSEngine 测试 JavaScript 引擎
func testJSEngine() {
	startTime := time.Now()
	fmt.Println("[JS] 初始化 JavaScript 引擎...")

	// 获取 JavaScript 引擎实例
	engine := js_engine.GetEngine()
	defer js_engine.Close()

	// 从嵌入的文件系统中读取脚本
	fmt.Println("[JS] 从 FS 中提取测试脚本...")
	scriptContent, err := scriptsFS.ReadFile("scripts/test_full.js")
	if err != nil {
		fmt.Printf("[JS] 读取脚本失败: %v\n", err)
		return
	}

	// 将脚本写入临时文件（模拟从 FS 中提取到实际文件系统）
	tmpPath := "/sdcard/autogo_test_extracted.js"
	scriptStr := string(scriptContent)
	fmt.Printf("[JS] 脚本大小: %d 字节\n", len(scriptStr))

	// 显示脚本内容预览
	preview := scriptStr
	if len(preview) > 200 {
		preview = preview[:200] + "..."
	}
	fmt.Printf("[JS] 脚本预览:\n%s\n", preview)

	// 执行脚本
	fmt.Println("[JS] 开始执行 JavaScript 脚本...")
	fmt.Println("────────────────────────────────────────────────────────────")

	err = engine.ExecuteString(scriptStr)
	if err != nil {
		fmt.Printf("[JS] 执行错误: %v\n", err)
	}

	fmt.Println("────────────────────────────────────────────────────────────")

	// 清理临时文件
	os.Remove(tmpPath)

	elapsed := time.Since(startTime)
	fmt.Printf("[JS] 测试完成，耗时: %v\n", elapsed)
}

// testLuaEngine 测试 Lua 引擎
func testLuaEngine() {
	startTime := time.Now()
	fmt.Println("[Lua] 初始化 Lua 引擎...")

	// 获取 Lua 引擎实例
	engine := lua_engine.GetEngine()
	defer lua_engine.Close()

	// 从嵌入的文件系统中读取脚本
	fmt.Println("[Lua] 从 FS 中提取测试脚本...")
	scriptContent, err := scriptsFS.ReadFile("scripts/test_full.lua")
	if err != nil {
		fmt.Printf("[Lua] 读取脚本失败: %v\n", err)
		return
	}

	// 将脚本写入临时文件（模拟从 FS 中提取到实际文件系统）
	tmpPath := "/sdcard/autogo_test_extracted.lua"
	scriptStr := string(scriptContent)
	fmt.Printf("[Lua] 脚本大小: %d 字节\n", len(scriptStr))

	// 显示脚本内容预览
	preview := scriptStr
	if len(preview) > 200 {
		preview = preview[:200] + "..."
	}
	fmt.Printf("[Lua] 脚本预览:\n%s\n", preview)

	// 执行脚本
	fmt.Println("[Lua] 开始执行 Lua 脚本...")
	fmt.Println("────────────────────────────────────────────────────────────")

	err = engine.ExecuteString(scriptStr)
	if err != nil {
		fmt.Printf("[Lua] 执行错误: %v\n", err)
	}

	fmt.Println("────────────────────────────────────────────────────────────")

	// 清理临时文件
	os.Remove(tmpPath)

	elapsed := time.Since(startTime)
	fmt.Printf("[Lua] 测试完成，耗时: %v\n", elapsed)
}
