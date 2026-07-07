package node

import "github.com/ZingYao/autogo_scriptengine/lua_engine/model"

// NodeModule 是 go-lua-vm 迁移后的模块壳。
type NodeModule struct {
	ThrowException bool
	ShowWarning    bool
	Debug          bool
}

func New() *NodeModule { return &NodeModule{} }

func (m *NodeModule) Name() string { return "node" }

func (m *NodeModule) IsAvailable() bool { return true }

func (m *NodeModule) Register(engine model.Engine) error { return nil }

func GetModule() model.Module { return &NodeModule{} }
