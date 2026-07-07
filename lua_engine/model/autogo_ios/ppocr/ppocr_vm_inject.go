package ppocr

import "github.com/ZingYao/autogo_scriptengine/lua_engine/model"

// PpocrModule 是 go-lua-vm 迁移后的模块壳。
type PpocrModule struct{}

func New() *PpocrModule { return &PpocrModule{} }

func (m *PpocrModule) Name() string { return "ppocr" }

func (m *PpocrModule) IsAvailable() bool { return true }

func (m *PpocrModule) Register(engine model.Engine) error { return nil }

func GetModule() model.Module { return &PpocrModule{} }
