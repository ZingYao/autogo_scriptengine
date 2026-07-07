package https

import "github.com/ZingYao/autogo_scriptengine/lua_engine/model"

// HttpsModule 是 go-lua-vm 迁移后的模块壳。
type HttpsModule struct{}

func New() *HttpsModule { return &HttpsModule{} }

func (m *HttpsModule) Name() string { return "https" }

func (m *HttpsModule) IsAvailable() bool { return true }

func (m *HttpsModule) Register(engine model.Engine) error { return nil }

func GetModule() model.Module { return &HttpsModule{} }
