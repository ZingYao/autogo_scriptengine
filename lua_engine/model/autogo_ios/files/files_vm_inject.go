package files

import "github.com/ZingYao/autogo_scriptengine/lua_engine/model"

// FilesModule 是 go-lua-vm 迁移后的模块壳。
type FilesModule struct{}

func New() *FilesModule { return &FilesModule{} }

func (m *FilesModule) Name() string { return "files" }

func (m *FilesModule) IsAvailable() bool { return true }

func (m *FilesModule) Register(engine model.Engine) error { return nil }

func GetModule() model.Module { return &FilesModule{} }
