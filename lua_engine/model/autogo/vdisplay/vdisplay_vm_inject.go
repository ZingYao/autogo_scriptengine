package vdisplay

import (
	"fmt"

	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"
)

// VdisplayModule 在远程 AutoGo 缺少 vdisplay 包时保留同名入口并返回明确错误。
type VdisplayModule struct{}

func New() *VdisplayModule { return &VdisplayModule{} }

func (m *VdisplayModule) Name() string { return "vdisplay" }

func (m *VdisplayModule) IsAvailable() bool { return false }

func (m *VdisplayModule) Register(engine model.Engine) error {
	engine.RegisterMethod("vdisplay.create", "AutoGo/vdisplay remote package unavailable", unavailable, true)
	engine.RegisterMethod("vdisplay.getDisplayId", "AutoGo/vdisplay remote package unavailable", unavailable, true)
	engine.RegisterMethod("vdisplay.launchApp", "AutoGo/vdisplay remote package unavailable", unavailable, true)
	engine.RegisterMethod("vdisplay.setTitle", "AutoGo/vdisplay remote package unavailable", unavailable, true)
	engine.RegisterMethod("vdisplay.setTouchCallback", "AutoGo/vdisplay remote package unavailable", unavailable, true)
	engine.RegisterMethod("vdisplay.showPreviewWindow", "AutoGo/vdisplay remote package unavailable", unavailable, true)
	engine.RegisterMethod("vdisplay.hidePreviewWindow", "AutoGo/vdisplay remote package unavailable", unavailable, true)
	engine.RegisterMethod("vdisplay.setPreviewWindowSize", "AutoGo/vdisplay remote package unavailable", unavailable, true)
	engine.RegisterMethod("vdisplay.setPreviewWindowPos", "AutoGo/vdisplay remote package unavailable", unavailable, true)
	engine.RegisterMethod("vdisplay.destroy", "AutoGo/vdisplay remote package unavailable", unavailable, true)
	return nil
}

func unavailable(args ...interface{}) (interface{}, error) {
	return nil, fmt.Errorf("AutoGo/%s package is unavailable in the remote github.com/Dasongzi1366/AutoGo module", "vdisplay")
}

func GetModule() model.Module { return &VdisplayModule{} }
