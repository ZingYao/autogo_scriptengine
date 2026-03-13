package safe_models

import (
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/app"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/coroutine"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/device"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/dotocr"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/files"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/http"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/images"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/ime"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/json"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/media"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/motion"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/opencv"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/plugin"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/ppocr"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/rhino"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/storages"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/system"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/uiacc"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/utils"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/yolo"
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
