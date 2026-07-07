package hud

import (
	"fmt"

	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"
)

// HUDModule 在远程 AutoGo 缺少 hud 包时保留同名入口并返回明确错误。
type HUDModule struct{}

func New() *HUDModule { return &HUDModule{} }

func (m *HUDModule) Name() string { return "hud" }

func (m *HUDModule) IsAvailable() bool { return false }

func (m *HUDModule) Register(engine model.Engine) error {
	engine.RegisterMethod("hud.new", "AutoGo/hud remote package unavailable", unavailable, true)
	engine.RegisterMethod("hud.setPosition", "AutoGo/hud remote package unavailable", unavailable, true)
	engine.RegisterMethod("hud.setBackgroundColor", "AutoGo/hud remote package unavailable", unavailable, true)
	engine.RegisterMethod("hud.setTextSize", "AutoGo/hud remote package unavailable", unavailable, true)
	engine.RegisterMethod("hud.setText", "AutoGo/hud remote package unavailable", unavailable, true)
	engine.RegisterMethod("hud.show", "AutoGo/hud remote package unavailable", unavailable, true)
	engine.RegisterMethod("hud.hide", "AutoGo/hud remote package unavailable", unavailable, true)
	engine.RegisterMethod("hud.isVisible", "AutoGo/hud remote package unavailable", unavailable, true)
	engine.RegisterMethod("hud.destroy", "AutoGo/hud remote package unavailable", unavailable, true)
	return nil
}

func unavailable(args ...interface{}) (interface{}, error) {
	return nil, fmt.Errorf("AutoGo/%s package is unavailable in the remote github.com/Dasongzi1366/AutoGo module", "hud")
}

func GetModule() model.Module { return &HUDModule{} }
