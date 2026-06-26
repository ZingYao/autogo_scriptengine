package dotocr

import (
	"image"

	"github.com/ZingYao/autogo_scriptengine/js_engine/model"

	autogodotocr "github.com/Dasongzi1366/AutoGo/dotocr"
	"github.com/dop251/goja"
)

type DotocrModule struct{}

func (m *DotocrModule) Name() string      { return "dotocr" }
func (m *DotocrModule) IsAvailable() bool { return true }

func pointToJS(vm *goja.Runtime, x, y int) goja.Value {
	return vm.ToValue(map[string]int{"x": x, "y": y})
}

func (m *DotocrModule) Register(engine model.Engine) error {
	vm := engine.GetVM()
	dotocrObj := vm.NewObject()
	vm.Set("dotocr", dotocrObj)

	dotocrObj.Set("setDict", func(call goja.FunctionCall) goja.Value {
		autogodotocr.SetDict(call.Argument(0).String(), call.Argument(1).String())
		return goja.Undefined()
	})
	dotocrObj.Set("ocr", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(autogodotocr.Ocr(
			int(call.Argument(0).ToInteger()),
			int(call.Argument(1).ToInteger()),
			int(call.Argument(2).ToInteger()),
			int(call.Argument(3).ToInteger()),
			call.Argument(4).String(),
			int(call.Argument(5).ToInteger()),
			int(call.Argument(6).ToInteger()),
			float32(call.Argument(7).ToFloat()),
			int(call.Argument(8).ToInteger()),
			call.Argument(9).String(),
		))
	})
	dotocrObj.Set("ocrFromImage", func(call goja.FunctionCall) goja.Value {
		img := call.Argument(0).Export().(*image.NRGBA)
		return vm.ToValue(autogodotocr.OcrFromImage(img, call.Argument(1).String(), int(call.Argument(2).ToInteger()), int(call.Argument(3).ToInteger()), float32(call.Argument(4).ToFloat()), int(call.Argument(5).ToInteger()), call.Argument(6).String()))
	})
	dotocrObj.Set("ocrFromBase64", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(autogodotocr.OcrFromBase64(call.Argument(0).String(), call.Argument(1).String(), int(call.Argument(2).ToInteger()), int(call.Argument(3).ToInteger()), float32(call.Argument(4).ToFloat()), int(call.Argument(5).ToInteger()), call.Argument(6).String()))
	})
	dotocrObj.Set("ocrFromPath", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(autogodotocr.OcrFromPath(call.Argument(0).String(), call.Argument(1).String(), int(call.Argument(2).ToInteger()), int(call.Argument(3).ToInteger()), float32(call.Argument(4).ToFloat()), int(call.Argument(5).ToInteger()), call.Argument(6).String()))
	})
	dotocrObj.Set("findStr", func(call goja.FunctionCall) goja.Value {
		x, y := autogodotocr.FindStr(int(call.Argument(0).ToInteger()), int(call.Argument(1).ToInteger()), int(call.Argument(2).ToInteger()), int(call.Argument(3).ToInteger()), call.Argument(4).String(), call.Argument(5).String(), int(call.Argument(6).ToInteger()), int(call.Argument(7).ToInteger()), float32(call.Argument(8).ToFloat()), call.Argument(9).String())
		return pointToJS(vm, x, y)
	})
	dotocrObj.Set("findStrFromImage", func(call goja.FunctionCall) goja.Value {
		img := call.Argument(0).Export().(*image.NRGBA)
		x, y := autogodotocr.FindStrFromImage(img, call.Argument(1).String(), call.Argument(2).String(), int(call.Argument(3).ToInteger()), int(call.Argument(4).ToInteger()), float32(call.Argument(5).ToFloat()), call.Argument(6).String())
		return pointToJS(vm, x, y)
	})
	dotocrObj.Set("findStrFromBase64", func(call goja.FunctionCall) goja.Value {
		x, y := autogodotocr.FindStrFromBase64(call.Argument(0).String(), call.Argument(1).String(), call.Argument(2).String(), int(call.Argument(3).ToInteger()), int(call.Argument(4).ToInteger()), float32(call.Argument(5).ToFloat()), call.Argument(6).String())
		return pointToJS(vm, x, y)
	})
	dotocrObj.Set("findStrFromPath", func(call goja.FunctionCall) goja.Value {
		x, y := autogodotocr.FindStrFromPath(call.Argument(0).String(), call.Argument(1).String(), call.Argument(2).String(), int(call.Argument(3).ToInteger()), int(call.Argument(4).ToInteger()), float32(call.Argument(5).ToFloat()), call.Argument(6).String())
		return pointToJS(vm, x, y)
	})

	engine.RegisterMethod("dotocr.setDict", "设置字库", autogodotocr.SetDict, true)
	engine.RegisterMethod("dotocr.ocr", "从屏幕区域进行点阵 OCR", autogodotocr.Ocr, true)
	engine.RegisterMethod("dotocr.ocrFromImage", "从图像对象进行点阵 OCR", autogodotocr.OcrFromImage, true)
	engine.RegisterMethod("dotocr.ocrFromBase64", "从 Base64 图像进行点阵 OCR", autogodotocr.OcrFromBase64, true)
	engine.RegisterMethod("dotocr.ocrFromPath", "从文件图像进行点阵 OCR", autogodotocr.OcrFromPath, true)
	engine.RegisterMethod("dotocr.findStr", "从屏幕区域查找字符串位置", autogodotocr.FindStr, true)
	engine.RegisterMethod("dotocr.findStrFromImage", "从图像对象查找字符串位置", autogodotocr.FindStrFromImage, true)
	engine.RegisterMethod("dotocr.findStrFromBase64", "从 Base64 图像查找字符串位置", autogodotocr.FindStrFromBase64, true)
	engine.RegisterMethod("dotocr.findStrFromPath", "从文件图像查找字符串位置", autogodotocr.FindStrFromPath, true)
	return nil
}
