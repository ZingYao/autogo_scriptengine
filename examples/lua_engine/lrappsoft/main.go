package main

import (
	"embed"
	"fmt"
	"log"

	"github.com/ZingYao/autogo_scriptengine/lua_engine"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/define/lrappsoft_models"
)

//go:embed scripts/*
var scriptsFS embed.FS

func main() {
	// 初始化 Lua 引擎，配置文件系统以支持 require
	config := lua_engine.DefaultConfig()
	config.FileSystem = scriptsFS
	engine := lua_engine.NewLuaEngine(&config)
	defer engine.Close()

	// 注册所有 lrappsoft 风格模块
	engine.RegisterModule(lrappsoft_models.LrappsoftModules...)

	// 执行工具脚本（以便主脚本可以引用）
	err := engine.ExecuteFile("scripts/utils.lua")
	if err != nil {
		log.Fatalf("Failed to execute utils.lua: %v", err)
	}

	// 执行主脚本
	err = engine.ExecuteFile("scripts/main.lua")
	if err != nil {
		log.Fatalf("Failed to execute main.lua: %v", err)
	}

	// 输出执行结果
	fmt.Println("Lua lrappsoft style example completed!")
}
