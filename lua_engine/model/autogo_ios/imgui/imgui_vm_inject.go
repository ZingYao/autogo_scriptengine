package imgui

import "github.com/ZingYao/autogo_scriptengine/lua_engine/model"

// ImGuiModule 是 go-lua-vm 迁移后的模块壳。
type ImGuiModule struct{}

func New() *ImGuiModule { return &ImGuiModule{} }

func (m *ImGuiModule) Name() string { return "imgui" }

func (m *ImGuiModule) IsAvailable() bool { return true }

func (m *ImGuiModule) Register(engine model.Engine) error { return nil }

func GetModule() model.Module { return &ImGuiModule{} }
