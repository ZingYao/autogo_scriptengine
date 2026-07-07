package yolo

import (
	"image"

	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"

	autogoyolo "github.com/Dasongzi1366/AutoGo/yolo"
)

// YoloModule 是 go-lua-vm 迁移后的模块壳。
type YoloModule struct{}

func New() *YoloModule { return &YoloModule{} }

func (m *YoloModule) Name() string { return "yolo" }

func (m *YoloModule) IsAvailable() bool { return true }

func (m *YoloModule) Register(engine model.Engine) error {
	engine.RegisterMethod("yolo.new", "创建 YOLO 实例", autogoyolo.New, true)
	engine.RegisterMethod("yolo.detect", "检测屏幕上的对象", func(y *autogoyolo.Yolo, x1, y1, x2, y2 int, displayID ...int) []autogoyolo.Result {
		return y.Detect(x1, y1, x2, y2, optionalInt(0, displayID...))
	}, true)
	engine.RegisterMethod("yolo.detectFromImage", "检测图片中的对象", func(y *autogoyolo.Yolo, img *image.NRGBA) []autogoyolo.Result {
		return y.DetectFromImage(img)
	}, true)
	engine.RegisterMethod("yolo.setImage", "设置下次 Detect 使用的原始图像", (*autogoyolo.Yolo).SetImage, true)
	engine.RegisterMethod("yolo.detectFromBase64", "检测 Base64 图片中的对象", (*autogoyolo.Yolo).DetectFromBase64, true)
	engine.RegisterMethod("yolo.detectFromPath", "检测文件图片中的对象", (*autogoyolo.Yolo).DetectFromPath, true)
	engine.RegisterMethod("yolo.close", "关闭 YOLO 实例", (*autogoyolo.Yolo).Close, true)
	return nil
}

func GetModule() model.Module { return &YoloModule{} }

func optionalInt(defaultValue int, values ...int) int {
	if len(values) == 0 {
		return defaultValue
	}
	return values[0]
}
