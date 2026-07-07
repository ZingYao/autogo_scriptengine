package rhino

import (
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"

	autogorhino "github.com/Dasongzi1366/AutoGo/rhino"
)

// RhinoModule 是 go-lua-vm 迁移后的模块壳。
type RhinoModule struct{}

func New() *RhinoModule { return &RhinoModule{} }

func (m *RhinoModule) Name() string { return "rhino" }

func (m *RhinoModule) IsAvailable() bool { return true }

func (m *RhinoModule) Register(engine model.Engine) error {
	engine.RegisterMethod("rhino.eval", "执行指定 JavaScript 脚本并返回结果", autogorhino.Eval, true)
	return nil
}

func GetModule() model.Module { return &RhinoModule{} }
