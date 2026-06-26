package safe_models

import (
	"github.com/ZingYao/autogo_scriptengine/js_engine/model"
	"github.com/ZingYao/autogo_scriptengine/js_engine/model/autogo_ios/app"
	"github.com/ZingYao/autogo_scriptengine/js_engine/model/autogo_ios/console"
	"github.com/ZingYao/autogo_scriptengine/js_engine/model/autogo_ios/device"
	"github.com/ZingYao/autogo_scriptengine/js_engine/model/autogo_ios/dotocr"
	"github.com/ZingYao/autogo_scriptengine/js_engine/model/autogo_ios/files"
	"github.com/ZingYao/autogo_scriptengine/js_engine/model/autogo_ios/https"
	"github.com/ZingYao/autogo_scriptengine/js_engine/model/autogo_ios/hud"
	"github.com/ZingYao/autogo_scriptengine/js_engine/model/autogo_ios/images"
	"github.com/ZingYao/autogo_scriptengine/js_engine/model/autogo_ios/ime"
	"github.com/ZingYao/autogo_scriptengine/js_engine/model/autogo_ios/imgui"
	"github.com/ZingYao/autogo_scriptengine/js_engine/model/autogo_ios/motion"
	"github.com/ZingYao/autogo_scriptengine/js_engine/model/autogo_ios/opencv"
	"github.com/ZingYao/autogo_scriptengine/js_engine/model/autogo_ios/ppocr"
	"github.com/ZingYao/autogo_scriptengine/js_engine/model/autogo_ios/storages"
	"github.com/ZingYao/autogo_scriptengine/js_engine/model/autogo_ios/system"
	"github.com/ZingYao/autogo_scriptengine/js_engine/model/autogo_ios/utils"
	"github.com/ZingYao/autogo_scriptengine/js_engine/model/autogo_ios/yolo"
)

// SafeModules 包含 iOS AutoGo 安全模块。
var SafeModules = []model.Module{
	&app.AppModule{},
	&console.ConsoleModule{},
	&device.DeviceModule{},
	&dotocr.DotocrModule{},
	&files.FilesModule{},
	&hud.HUDModule{},
	&https.HttpsModule{},
	&images.ImagesModule{},
	&ime.ImeModule{},
	&imgui.ImGuiModule{},
	&motion.MotionModule{},
	&opencv.OpencvModule{},
	&ppocr.PpocrModule{},
	&storages.StoragesModule{},
	&system.SystemModule{},
	&utils.UtilsModule{},
	&yolo.YoloModule{},
}
