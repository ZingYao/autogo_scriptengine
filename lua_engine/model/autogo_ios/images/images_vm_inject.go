package images

import "github.com/ZingYao/autogo_scriptengine/lua_engine/model"

// ImagesModule 是 go-lua-vm 迁移后的模块壳。
type ImagesModule struct{}

func New() *ImagesModule { return &ImagesModule{} }

func (m *ImagesModule) Name() string { return "images" }

func (m *ImagesModule) IsAvailable() bool { return true }

func (m *ImagesModule) Register(engine model.Engine) error { return nil }

func GetModule() model.Module { return &ImagesModule{} }
