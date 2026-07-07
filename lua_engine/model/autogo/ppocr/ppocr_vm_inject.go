package ppocr

import (
	"image"

	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"

	autogoppocr "github.com/Dasongzi1366/AutoGo/ppocr"
)

// PpocrModule 是 go-lua-vm 迁移后的模块壳。
type PpocrModule struct{}

func New() *PpocrModule { return &PpocrModule{} }

func (m *PpocrModule) Name() string { return "ppocr" }

func (m *PpocrModule) IsAvailable() bool { return true }

func (m *PpocrModule) Register(engine model.Engine) error {
	engine.RegisterMethod("ppocr.new", "创建 Ppocr 对象", autogoppocr.New, true)
	engine.RegisterMethod("ppocr.ocr", "识别屏幕文字", func(p *autogoppocr.Ppocr, x1, y1, x2, y2 int, colorStr string, displayID ...int) []autogoppocr.Result {
		return p.Ocr(x1, y1, x2, y2, colorStr, optionalInt(0, displayID...))
	}, true)
	engine.RegisterMethod("ppocr.ocrFromImage", "识别图片文字", func(p *autogoppocr.Ppocr, img *image.NRGBA, colorStr string) []autogoppocr.Result {
		return p.OcrFromImage(img, colorStr)
	}, true)
	engine.RegisterMethod("ppocr.ocrFromBase64", "识别 Base64 图片文字", func(p *autogoppocr.Ppocr, b64, colorStr string) []autogoppocr.Result {
		return p.OcrFromBase64(b64, colorStr)
	}, true)
	engine.RegisterMethod("ppocr.ocrFromPath", "识别文件图片文字", func(p *autogoppocr.Ppocr, path, colorStr string) []autogoppocr.Result {
		return p.OcrFromPath(path, colorStr)
	}, true)
	engine.RegisterMethod("ppocr.close", "关闭 Ppocr 对象", (*autogoppocr.Ppocr).Close, true)
	return nil
}

func GetModule() model.Module { return &PpocrModule{} }

func optionalInt(defaultValue int, values ...int) int {
	if len(values) == 0 {
		return defaultValue
	}
	return values[0]
}
