package js_engine

import (
	"os"
	"path/filepath"
	"sync"
	"testing"

	"github.com/dop251/goja"
)

type recordingDebugger struct {
	mu     sync.Mutex
	hits   int
	frames []goja.DebugFrame
}

func (d *recordingDebugger) BeforeInstruction(_ *goja.Runtime, state goja.DebugState) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	d.hits++
	if len(state.Frames) > 0 {
		d.frames = append(d.frames, state.Frames[0])
	}
	return nil
}

func (d *recordingDebugger) hitCount() int {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.hits
}

func TestJSEngineDebuggerReceivesExecutionState(t *testing.T) {
	debugger := &recordingDebugger{}
	engine := NewJSEngine(&EngineConfig{Debugger: debugger})

	if err := engine.ExecuteString("let value = 1; value += 2; globalThis.debugResult = value;"); err != nil {
		t.Fatalf("ExecuteString returned error: %v", err)
	}
	if debugger.hitCount() == 0 {
		t.Fatal("expected debugger to receive at least one execution callback")
	}
	if got := engine.GetVM().Get("debugResult").ToInteger(); got != 3 {
		t.Fatalf("debugResult = %d, want 3", got)
	}
}

func TestJSEngineImportModuleUsesRequireResolution(t *testing.T) {
	dir := t.TempDir()
	modulePath := filepath.Join(dir, "math.js")
	if err := os.WriteFile(modulePath, []byte("module.exports = { add: function(a, b) { return a + b; } };"), 0o644); err != nil {
		t.Fatalf("write module: %v", err)
	}

	engine := NewJSEngine(nil)
	script := `
		importModule("./math").then(function(math) {
			globalThis.importResult = math.add(4, 5);
		});
	`
	if err := engine.ExecuteString(script, dir); err != nil {
		t.Fatalf("ExecuteString returned error: %v", err)
	}
	if got := engine.GetVM().Get("importResult").ToInteger(); got != 9 {
		t.Fatalf("importResult = %d, want 9", got)
	}
}
