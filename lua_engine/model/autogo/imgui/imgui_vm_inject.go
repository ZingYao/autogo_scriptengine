//go:build !android || !cgo

package imgui

import (
	"fmt"

	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"
)

// ImGuiModule 是非 Android 环境下的 ImGui 模块占位。
type ImGuiModule struct{}

func New() *ImGuiModule { return &ImGuiModule{} }

func (m *ImGuiModule) Name() string { return "imgui" }

func (m *ImGuiModule) IsAvailable() bool { return true }

func (m *ImGuiModule) Register(engine model.Engine) error {
	engine.RegisterMethod("imgui.newVec2", "创建 Vec2", func(x, y float32) (interface{}, error) {
		return nil, fmt.Errorf("imgui.newVec2 requires Android AutoGo ImGui runtime")
	}, true)
	return nil
}

func GetModule() model.Module { return &ImGuiModule{} }
