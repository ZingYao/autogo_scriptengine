package debugsource

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ZingYao/autogo_scriptengine/lua_engine"
	"github.com/ZingYao/go-lua-vm/bytecode"
	gruntime "github.com/ZingYao/go-lua-vm/runtime"
)

type sourceCapturingObserver struct {
	source string
}

// BeforeInstruction 保存文件入口第一条可见指令的源码路径。
func (o *sourceCapturingObserver) BeforeInstruction(_ *gruntime.State, _ *gruntime.VM,
	proto *bytecode.Proto, _ int) error {
	// 首次观察结果足以证明文件 chunk 没有退化成 =(string)。
	if o.source == "" && proto != nil {
		o.source = proto.Source
	}
	return nil
}

// TestExecuteFilePreservesDebugSource 验证文件执行向 DAP 暴露真实源码路径。
func TestExecuteFilePreservesDebugSource(t *testing.T) {
	// DAP 断点按 Proto.Source 匹配；退化成 =(string) 会导致所有文件断点失效。
	observer := &sourceCapturingObserver{}
	config := lua_engine.DefaultConfig()
	config.DebugObserver = observer
	engine := lua_engine.NewLuaEngine(&config)
	defer engine.Close()

	scriptPath := filepath.Join(t.TempDir(), "debug-source.lua")
	if err := os.WriteFile(scriptPath, []byte("local answer = 42\n"), 0644); err != nil {
		t.Fatalf("WriteFile() error = %v", err)
	}
	if err := engine.ExecuteFile(scriptPath); err != nil {
		t.Fatalf("ExecuteFile() error = %v", err)
	}
	if observer.source != "@"+scriptPath {
		// 精确匹配锁定 IDE 与设备端使用同一文件身份。
		t.Fatalf("Proto.Source = %q, want %q", observer.source, "@"+scriptPath)
	}
}
