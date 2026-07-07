//go:build ignore

// 实际运行时需要修改这里为 main
package main_test

import (
	"embed"
	"fmt"
	"log"

	"github.com/ZingYao/autogo_scriptengine/lua_engine"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/define/android/autogo/all_models"
)

//go:embed scripts/*
var scriptsFS embed.FS

func main() {
	// 初始化 Lua 引擎，配置文件系统以支持 require
	config := lua_engine.DefaultConfig()
	config.FileSystem = scriptsFS
	engine := lua_engine.NewLuaEngine(&config)
	defer engine.Close()

	// 注册所有 autogo 风格模块
	engine.RegisterModule(all_models.AllModules...)

	// 执行主脚本；utils 名称保留给 AutoGo 模块对象，示例辅助函数放在脚本局部作用域。
	if err := engine.ExecuteFile("scripts/main.lua"); err != nil {
		log.Fatalf("Failed to execute main.lua: %v", err)
	}

	// 输出执行结果
	fmt.Println("Lua autogo style example completed!")
}
