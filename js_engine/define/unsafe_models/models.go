package unsafe_models

import (
	"app/js_engine/model"
	"app/js_engine/model/console"
	"app/js_engine/model/hud"
	"app/js_engine/model/imgui"
	"app/js_engine/model/vdisplay"
)

// UnsafeModules 包含不安全的模块（console、hud、vdisplay、 imgui）
// 这四个模块在 Android16 下会出现不安全的内存访问报错
var UnsafeModules = []model.Module{
	&console.ConsoleModule{},
	&imgui.ImGuiModule{},
	&hud.HUDModule{},
	&vdisplay.VdisplayModule{},
}
