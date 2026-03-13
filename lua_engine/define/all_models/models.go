package all_models

import (
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/app"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/console"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/coroutine"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/device"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/dotocr"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/files"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/http"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/hud"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/images"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/ime"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/imgui"
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
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/vdisplay"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/yolo"
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
