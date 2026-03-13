package unsafe_models

import (
	"github.com/ZingYao/autogo_scriptengine/js_engine/model"
	"github.com/ZingYao/autogo_scriptengine/js_engine/model/console"
	"github.com/ZingYao/autogo_scriptengine/js_engine/model/hud"
	"github.com/ZingYao/autogo_scriptengine/js_engine/model/imgui"
	"github.com/ZingYao/autogo_scriptengine/js_engine/model/vdisplay"
)

// UnsafeModules 包含不安全的模块（console、hud、vdisplay、 imgui）
// 这四个模块在 Android16 下会出现不安全的内存访问报错
var UnsafeModules = []model.Module{
	&console.ConsoleModule{},
	&imgui.ImGuiModule{},
	&hud.HUDModule{},
	&vdisplay.VdisplayModule{},
}
