package all_models

import (
	"app/lua_engine/model"
	"app/lua_engine/model/app"
	"app/lua_engine/model/console"
	"app/lua_engine/model/coroutine"
	"app/lua_engine/model/device"
	"app/lua_engine/model/dotocr"
	"app/lua_engine/model/files"
	"app/lua_engine/model/http"
	"app/lua_engine/model/hud"
	"app/lua_engine/model/images"
	"app/lua_engine/model/ime"
	"app/lua_engine/model/imgui"
	"app/lua_engine/model/media"
	"app/lua_engine/model/motion"
	"app/lua_engine/model/opencv"
	"app/lua_engine/model/plugin"
	"app/lua_engine/model/ppocr"
	"app/lua_engine/model/rhino"
	"app/lua_engine/model/storages"
	"app/lua_engine/model/system"
	"app/lua_engine/model/uiacc"
	"app/lua_engine/model/utils"
	"app/lua_engine/model/vdisplay"
	"app/lua_engine/model/yolo"
)

// AllModules 包含所有可用模块的数组
var AllModules = []model.Module{
	&app.AppModule{},
	&device.DeviceModule{},
	&console.ConsoleModule{},
	&hud.HUDModule{},
	&vdisplay.VdisplayModule{},
	&coroutine.CoroutineModule{},
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
