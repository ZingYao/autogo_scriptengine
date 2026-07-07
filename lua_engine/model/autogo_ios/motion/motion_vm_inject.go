package motion

import "github.com/ZingYao/autogo_scriptengine/lua_engine/model"

// MotionModule 是 go-lua-vm 迁移后的模块壳。
type MotionModule struct{}

func New() *MotionModule { return &MotionModule{} }

func (m *MotionModule) Name() string { return "motion" }

func (m *MotionModule) IsAvailable() bool { return true }

func (m *MotionModule) Register(engine model.Engine) error { return nil }

func GetModule() model.Module { return &MotionModule{} }
