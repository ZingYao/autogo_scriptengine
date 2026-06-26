package ppocr

import (
	"image"

	"github.com/ZingYao/autogo_scriptengine/js_engine/model"

	autogoppocr "github.com/Dasongzi1366/AutoGo/ppocr"
	"github.com/dop251/goja"
)

type PpocrModule struct{}

func (m *PpocrModule) Name() string      { return "ppocr" }
func (m *PpocrModule) IsAvailable() bool { return true }

func resultsToJS(vm *goja.Runtime, results []autogoppocr.Result) goja.Value {
	rows := make([]map[string]any, 0, len(results))
	for _, item := range results {
		rows = append(rows, map[string]any{
			"x":       item.X,
			"y":       item.Y,
			"width":   item.Width,
			"height":  item.Height,
			"label":   item.Label,
			"score":   item.Score,
			"centerX": item.CenterX,
			"centerY": item.CenterY,
		})
	}
	return vm.ToValue(rows)
}

func wrapPpocr(vm *goja.Runtime, p *autogoppocr.Ppocr) goja.Value {
	obj := vm.NewObject()
	obj.Set("ocr", func(call goja.FunctionCall) goja.Value {
		return resultsToJS(vm, p.Ocr(
			int(call.Argument(0).ToInteger()),
			int(call.Argument(1).ToInteger()),
			int(call.Argument(2).ToInteger()),
			int(call.Argument(3).ToInteger()),
			call.Argument(4).String(),
		))
	})
	obj.Set("ocrFromImage", func(call goja.FunctionCall) goja.Value {
		img := call.Argument(0).Export().(*image.NRGBA)
		return resultsToJS(vm, p.OcrFromImage(img, call.Argument(1).String()))
	})
	obj.Set("ocrFromBase64", func(call goja.FunctionCall) goja.Value {
		return resultsToJS(vm, p.OcrFromBase64(call.Argument(0).String(), call.Argument(1).String()))
	})
	obj.Set("ocrFromPath", func(call goja.FunctionCall) goja.Value {
		return resultsToJS(vm, p.OcrFromPath(call.Argument(0).String(), call.Argument(1).String()))
	})
	obj.Set("close", func(call goja.FunctionCall) goja.Value {
		p.Close()
		return goja.Undefined()
	})
	return obj
}

func (m *PpocrModule) Register(engine model.Engine) error {
	vm := engine.GetVM()
	ppocrObj := vm.NewObject()
	vm.Set("ppocr", ppocrObj)
	ppocrObj.Set("new", func(call goja.FunctionCall) goja.Value {
		p := autogoppocr.New(call.Argument(0).String())
		if p == nil {
			return goja.Null()
		}
		return wrapPpocr(vm, p)
	})

	engine.RegisterMethod("ppocr.new", "创建 PPOCR 实例", autogoppocr.New, true)
	engine.RegisterMethod("ppocr.ocr", "从屏幕区域识别文字", func(p *autogoppocr.Ppocr, x1, y1, x2, y2 int, colorStr string) []autogoppocr.Result {
		return p.Ocr(x1, y1, x2, y2, colorStr)
	}, true)
	engine.RegisterMethod("ppocr.ocrFromImage", "从图像对象识别文字", func(p *autogoppocr.Ppocr, img *image.NRGBA, colorStr string) []autogoppocr.Result {
		return p.OcrFromImage(img, colorStr)
	}, true)
	engine.RegisterMethod("ppocr.ocrFromBase64", "从 Base64 图像识别文字", func(p *autogoppocr.Ppocr, b64, colorStr string) []autogoppocr.Result {
		return p.OcrFromBase64(b64, colorStr)
	}, true)
	engine.RegisterMethod("ppocr.ocrFromPath", "从文件图像识别文字", func(p *autogoppocr.Ppocr, path, colorStr string) []autogoppocr.Result {
		return p.OcrFromPath(path, colorStr)
	}, true)
	engine.RegisterMethod("ppocr.close", "关闭 PPOCR 实例", func(p *autogoppocr.Ppocr) { p.Close() }, true)
	return nil
}
