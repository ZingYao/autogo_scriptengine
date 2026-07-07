package cv

import "github.com/ZingYao/autogo_scriptengine/lua_engine/model"

// CvModule 是 go-lua-vm 迁移后的模块壳。
type CvModule struct{}

func New() *CvModule { return &CvModule{} }

func (m *CvModule) Name() string { return "cv" }

func (m *CvModule) IsAvailable() bool { return true }

func (m *CvModule) Register(engine model.Engine) error { return nil }

func GetModule() model.Module { return &CvModule{} }
