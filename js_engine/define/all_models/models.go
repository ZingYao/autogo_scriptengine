package all_models

import (
	"app/js_engine/model"
	"app/js_engine/model/app"
	"app/js_engine/model/console"
	"app/js_engine/model/device"
	"app/js_engine/model/dotocr"
	"app/js_engine/model/files"
	"app/js_engine/model/http"
	"app/js_engine/model/hud"
	"app/js_engine/model/images"
	"app/js_engine/model/ime"
	"app/js_engine/model/imgui"
	"app/js_engine/model/media"
	"app/js_engine/model/motion"
	"app/js_engine/model/opencv"
	"app/js_engine/model/plugin"
	"app/js_engine/model/ppocr"
	"app/js_engine/model/rhino"
	"app/js_engine/model/storages"
	"app/js_engine/model/system"
	"app/js_engine/model/uiacc"
	"app/js_engine/model/utils"
	"app/js_engine/model/vdisplay"
	"app/js_engine/model/yolo"
)

// AllModules 包含所有可用模块的数组
var AllModules = []model.Module{
	&app.AppModule{},
	&device.DeviceModule{},
	&console.ConsoleModule{},
	&hud.HUDModule{},
	&vdisplay.VdisplayModule{},
	&dotocr.DotocrModule{},
	&files.FilesModule{},
	&http.HttpModule{},
	&images.ImagesModule{},
	&ime.ImeModule{},
	&imgui.ImGuiModule{},
	&media.MediaModule{},
	&motion.MotionModule{},
	&opencv.OpencvModule{},
	&plugin.PluginModule{},
	&ppocr.PpocrModule{},
	&rhino.RhinoModule{},
	&storages.StoragesModule{},
	&system.SystemModule{},
	&uiacc.UiaccModule{},
	&utils.UtilsModule{},
	&yolo.YoloModule{},
}
