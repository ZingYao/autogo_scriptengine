package opencv

import "github.com/ZingYao/autogo_scriptengine/lua_engine/model"

// OpencvModule 是 go-lua-vm 迁移后的模块壳。
type OpencvModule struct{}

func New() *OpencvModule { return &OpencvModule{} }

func (m *OpencvModule) Name() string { return "opencv" }

func (m *OpencvModule) IsAvailable() bool { return true }

func (m *OpencvModule) Register(engine model.Engine) error { return nil }

func GetModule() model.Module { return &OpencvModule{} }
