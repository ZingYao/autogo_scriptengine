package safe_models

import (
	"app/js_engine/model"
)

// SafeModules 包含安全的模块（排除 console、hud、imgui、vdisplay）
// 这四个模块在 Android16 下会出现不安全的内存访问报错
var SafeModules = []model.Module{
	// &app.AppModule{},
	// &coroutine.CoroutineModule{},
	// &device.DeviceModule{},
	// &dotocr.DotocrModule{},
	// &files.FilesModule{},
	// &http.HttpModule{},
	// &images.ImagesModule{},
	// &ime.ImeModule{},
	// &media.MediaModule{},
	// &motion.MotionModule{},
	// &opencv.OpencvModule{},
	// &plugin.PluginModule{},
	// &ppocr.PpocrModule{},
	// &rhino.RhinoModule{},
	// &storages.StoragesModule{},
	// &system.SystemModule{},
	// &uiacc.UiaccModule{},
	// &utils.UtilsModule{},
	// &yolo.YoloModule{},
}
