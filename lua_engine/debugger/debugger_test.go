package debugger_test

import (
	"testing"
	"time"

	"github.com/ZingYao/autogo_scriptengine/lua_engine/debugger"
)

func TestDebuggerBreakpointContinue(t *testing.T) {
	dbg := debugger.New(debugger.Config{
		Enabled:        true,
		CollectGlobals: true,
	})

	file := "test.lua"
	dbg.SetBreakpoints(file, []int{2})

	errCh := make(chan error, 1)
	go func() {
		errCh <- dbg.Hit(file, 2)
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
			t.Fatalf("hit: %v", err)
		}
	case <-time.After(time.Second):
		t.Fatal("timeout waiting for hit")
	}
}
