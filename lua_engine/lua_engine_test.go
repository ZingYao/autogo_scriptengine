//go:build ignore
// +build ignore

package lua_engine_test

import (
	"testing"

	lua_engine "github.com/ZingYao/autogo_scriptengine/lua_engine"
	lua "github.com/yuin/gopher-lua"
)

func TestLuaEngineExecuteString(t *testing.T) {
	engine := lua_engine.NewLuaEngine(nil)
	defer engine.GetLuaState().Close()

	err := engine.ExecuteString(`
local value = 40
value = value + 2
result = value
`)
	if err != nil {
		t.Fatalf("ExecuteString() error = %v", err)
	}

	got := engine.GetLuaState().GetGlobal("result")
	if got != lua.LNumber(42) {
		t.Fatalf("result = %v, want 42", got)
	}
}
