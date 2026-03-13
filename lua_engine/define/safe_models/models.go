package safe_models

import (
	"app/lua_engine/model"
	"app/lua_engine/model/app"
	"app/lua_engine/model/coroutine"
	"app/lua_engine/model/device"
	"app/lua_engine/model/dotocr"
	"app/lua_engine/model/files"
	"app/lua_engine/model/http"
	"app/lua_engine/model/images"
	"app/lua_engine/model/ime"
	"app/lua_engine/model/json"
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
	"app/lua_engine/model/yolo"
)

// SafeModules 包含安全的模块（排除 console、hud、imgui、vdisplay）
// 这四个模块在 Android16 下会出现不安全的内存访问报错
var SafeModules = []model.Module{
	&app.AppModule{},
	&coroutine.CoroutineModule{},
	&device.DeviceModule{},
	&dotocr.DotocrModule{},
	&files.FilesModule{},
	&http.HttpModule{},
	&ime.ImeModule{},
	&images.ImagesModule{},
	&json.JsonModule{},
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
