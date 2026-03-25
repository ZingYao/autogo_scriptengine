//go:build ignore

// 实际运行时需要修改这里为 main
package main_test

import (
	"embed"
	"fmt"
	"log"

	"github.com/ZingYao/autogo_scriptengine/js_engine"
	"github.com/ZingYao/autogo_scriptengine/js_engine/define/autogo/all_models"
)

//go:embed scripts/*
var scriptsFS embed.FS

func main() {
	// 初始化 JavaScript 引擎，配置文件系统以支持 require
	config := js_engine.DefaultConfig()
	config.FileSystem = scriptsFS
	engine := js_engine.NewJSEngine(&config)
	defer engine.Close()

	// 注册所有 autogo 风格模块
	engine.RegisterModule(all_models.AllModules...)

	// 执行工具脚本（以便主脚本可以引用）
	err := engine.ExecuteFile("scripts/utils.js")
	if err != nil {
		log.Fatalf("Failed to execute utils.js: %v", err)
	}

	// 执行主脚本
	err = engine.ExecuteFile("scripts/main.js")
	if err != nil {
		log.Fatalf("Failed to execute main.js: %v", err)
	}

	// 输出执行结果
	fmt.Println("JavaScript autogo style example completed!")
}
