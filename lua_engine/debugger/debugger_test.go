package debugger_test

import (
	"strings"
	"testing"
	"time"

	"github.com/ZingYao/autogo_scriptengine/lua_engine/debugger"
	lua "github.com/yuin/gopher-lua"
)

func TestDebuggerBreakpointContinue(t *testing.T) {
	dbg := debugger.New(debugger.Config{
		Enabled:        true,
		CollectGlobals: true,
	})
	L := lua.NewState()
	defer L.Close()
	dbg.Install(L)

	file := "test.lua"
	dbg.SetBreakpoints(file, []int{2})
	source := debugger.InstrumentSource("local x = 1\nx = x + 1\nresult = x", file)

	errCh := make(chan error, 1)
	go func() {
		fn, err := L.Load(strings.NewReader(source), file)
		if err != nil {
			errCh <- err
			return
		}
		L.Push(fn)
		errCh <- L.PCall(0, lua.MultRet, nil)
	}()

	select {
	case event := <-dbg.Events():
		if event.Type != "stopped" || event.Reason != debugger.StopReasonBreakpoint {
			t.Fatalf("unexpected event: %#v", event)
		}
		if event.Position.Line != 2 {
			t.Fatalf("unexpected line: %d", event.Position.Line)
		}
	case <-time.After(time.Second):
		t.Fatal("timeout waiting for breakpoint")
	}

	dbg.Continue()

	select {
	case err := <-errCh:
		if err != nil {
			t.Fatalf("execute: %v", err)
		}
	case <-time.After(time.Second):
		t.Fatal("timeout waiting for script")
	}

	if got := L.GetGlobal("result").String(); got != "2" {
		t.Fatalf("result = %s, want 2", got)
	}
}
