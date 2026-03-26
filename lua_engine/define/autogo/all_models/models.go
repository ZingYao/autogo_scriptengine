package all_models

import (
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/autogo/app"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/autogo/console"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/autogo/coroutine"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/autogo/device"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/autogo/dotocr"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/autogo/files"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/autogo/http"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/autogo/hud"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/autogo/images"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/autogo/ime"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/autogo/imgui"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/autogo/media"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/autogo/motion"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/autogo/opencv"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/autogo/plugin"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/autogo/ppocr"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/autogo/rhino"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/autogo/storages"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/autogo/system"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/autogo/uiacc"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/autogo/utils"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/autogo/vdisplay"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/autogo/websocket"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model/autogo/yolo"
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
	&websocket.WebSocketModule{},
	&yolo.YoloModule{},
}
