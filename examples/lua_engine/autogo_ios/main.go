//go:build ignore

// 实际运行时需要修改这里为 main。
package main_test

import (
	"embed"
	"fmt"
	"log"

	"github.com/ZingYao/autogo_scriptengine/lua_engine"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/define/ios/autogo/all_models"
)

//go:embed scripts/*
var scriptsFS embed.FS

func main() {
	// 初始化 Lua 引擎，配置文件系统以支持 require。
	config := lua_engine.DefaultConfig()
	config.FileSystem = scriptsFS
	engine := lua_engine.NewLuaEngine(&config)
	defer engine.Close()

	// 注册 iOS autogo 风格模块，避免注入 Android-only 模块。
	engine.RegisterModule(all_models.AllModules...)

	// 执行工具脚本，主脚本可通过 require("utils") 引用。
	if err := engine.ExecuteFile("scripts/utils.lua"); err != nil {
		log.Fatalf("Failed to execute utils.lua: %v", err)
	}

	// 执行 iOS 主脚本。
	if err := engine.ExecuteFile("scripts/main.lua"); err != nil {
		log.Fatalf("Failed to execute main.lua: %v", err)
	}

	fmt.Println("Lua iOS autogo style example completed!")
}
