package main

import (
	"log"
	"net"

	"github.com/ZingYao/autogo_scriptengine/lua_engine"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/debugger"
)

func main() {
	config := lua_engine.DefaultConfig()
	config.Debug = &debugger.Config{
		Enabled:          true,
		BreakOnError:     true,
		CollectGlobals:   true,
		CollectLocals:    true,
		MaxVariableDepth: 2,
	}

	engine := lua_engine.NewLuaEngine(&config)
	listener, err := net.Listen("tcp", "127.0.0.1:38697")
	if err != nil {
		log.Fatalf("listen dap: %v", err)
	}
	defer listener.Close()

	go func() {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept dap: %v", err)
			return
		}
		defer conn.Close()
		session := debugger.NewDAPSession(engine.GetDebugger(), conn, conn)
		if err := session.Serve(); err != nil {
			log.Printf("dap session: %v", err)
		}
	}()

	if err := engine.ExecuteFile("scripts/main.lua"); err != nil {
		log.Fatalf("execute lua: %v", err)
	}
}
