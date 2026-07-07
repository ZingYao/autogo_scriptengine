package device

import "github.com/ZingYao/autogo_scriptengine/lua_engine/model"

// DeviceModule 是 go-lua-vm 迁移后的模块壳。
type DeviceModule struct{}

func New() *DeviceModule { return &DeviceModule{} }

func (m *DeviceModule) Name() string { return "device" }

func (m *DeviceModule) IsAvailable() bool { return true }

func (m *DeviceModule) Register(engine model.Engine) error { return nil }

func GetModule() model.Module { return &DeviceModule{} }
