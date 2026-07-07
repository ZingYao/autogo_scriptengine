package image

import "github.com/ZingYao/autogo_scriptengine/lua_engine/model"

// ImageModule 是 go-lua-vm 迁移后的模块壳。
type ImageModule struct{}

func New() *ImageModule { return &ImageModule{} }

func (m *ImageModule) Name() string { return "image" }

func (m *ImageModule) IsAvailable() bool { return true }

func (m *ImageModule) Register(engine model.Engine) error { return nil }

func GetModule() model.Module { return &ImageModule{} }
