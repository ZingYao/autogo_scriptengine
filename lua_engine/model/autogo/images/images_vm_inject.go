package images

import (
	"image"

	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"

	autogoimages "github.com/Dasongzi1366/AutoGo/images"
)

// ImagesModule 是 go-lua-vm 迁移后的模块壳。
type ImagesModule struct{}

func New() *ImagesModule { return &ImagesModule{} }

func (m *ImagesModule) Name() string { return "images" }

func (m *ImagesModule) IsAvailable() bool { return true }

func (m *ImagesModule) Register(engine model.Engine) error {
	engine.RegisterMethod("images.setCallback", "设置截图回调函数", autogoimages.SetCallback, true)
	engine.RegisterMethod("images.captureScreen", "截取屏幕", func(x1, y1, x2, y2 int, displayID ...int) *image.NRGBA {
		return autogoimages.CaptureScreen(x1, y1, x2, y2, optionalInt(0, displayID...))
	}, true)
	engine.RegisterMethod("images.pixel", "获取指定坐标的像素颜色", func(x, y int, displayID ...int) string {
		return autogoimages.Pixel(x, y, optionalInt(0, displayID...))
	}, true)
	engine.RegisterMethod("images.cmpColor", "比较颜色", func(x, y int, colorStr string, sim float32, displayID ...int) bool {
		return autogoimages.CmpColor(x, y, colorStr, sim, optionalInt(0, displayID...))
	}, true)
	engine.RegisterMethod("images.findColor", "查找颜色", func(x1, y1, x2, y2 int, colorStr string, sim float32, dir int, displayID ...int) map[string]int {
		x, y := autogoimages.FindColor(x1, y1, x2, y2, colorStr, sim, dir, optionalInt(0, displayID...))
		return map[string]int{"x": x, "y": y}
	}, true)
	engine.RegisterMethod("images.getColorCountInRegion", "获取区域内指定颜色的数量", func(x1, y1, x2, y2 int, colorStr string, sim float32, displayID ...int) int {
		return autogoimages.GetColorCountInRegion(x1, y1, x2, y2, colorStr, sim, optionalInt(0, displayID...))
	}, true)
	engine.RegisterMethod("images.detectsMultiColors", "检测多点颜色", func(colors string, sim float32, displayID ...int) bool {
		return autogoimages.DetectsMultiColors(colors, sim, optionalInt(0, displayID...))
	}, true)
	engine.RegisterMethod("images.findMultiColors", "查找多点颜色", func(x1, y1, x2, y2 int, colors string, sim float32, dir int, displayID ...int) map[string]int {
		x, y := autogoimages.FindMultiColors(x1, y1, x2, y2, colors, sim, dir, optionalInt(0, displayID...))
		return map[string]int{"x": x, "y": y}
	}, true)
	engine.RegisterMethod("images.findMultiColorsAll", "查找所有多点颜色", func(x1, y1, x2, y2 int, colors string, sim float32, dir int, displayID ...int) []autogoimages.Point {
		return autogoimages.FindMultiColorsAll(x1, y1, x2, y2, colors, sim, dir, optionalInt(0, displayID...))
	}, true)
	engine.RegisterMethod("images.readFromPath", "从路径读取图片", autogoimages.ReadFromPath, true)
	engine.RegisterMethod("images.readFromUrl", "从 URL 读取图片", autogoimages.ReadFromUrl, true)
	engine.RegisterMethod("images.readFromBase64", "从 Base64 读取图片", autogoimages.ReadFromBase64, true)
	engine.RegisterMethod("images.readFromBytes", "从字节数组读取图片", autogoimages.ReadFromBytes, true)
	engine.RegisterMethod("images.save", "保存图片", autogoimages.Save, true)
	engine.RegisterMethod("images.encodeToBase64", "编码为 Base64", autogoimages.EncodeToBase64, true)
	engine.RegisterMethod("images.encodeToBytes", "编码为字节数组", autogoimages.EncodeToBytes, true)
	engine.RegisterMethod("images.toNrgba", "转换为 NRGBA 格式", autogoimages.ToNrgba, true)
	engine.RegisterMethod("images.clip", "裁剪图片", autogoimages.Clip, true)
	engine.RegisterMethod("images.resize", "调整图片大小", autogoimages.Resize, true)
	engine.RegisterMethod("images.rotate", "旋转图片", autogoimages.Rotate, true)
	engine.RegisterMethod("images.grayscale", "灰度化", autogoimages.Grayscale, true)
	engine.RegisterMethod("images.applyThreshold", "应用阈值", autogoimages.ApplyThreshold, true)
	engine.RegisterMethod("images.applyAdaptiveThreshold", "应用自适应阈值", autogoimages.ApplyAdaptiveThreshold, true)
	engine.RegisterMethod("images.applyBinarization", "二值化", autogoimages.ApplyBinarization, true)
	return nil
}

func GetModule() model.Module { return &ImagesModule{} }

func optionalInt(defaultValue int, values ...int) int {
	if len(values) == 0 {
		return defaultValue
	}
	return values[0]
}
