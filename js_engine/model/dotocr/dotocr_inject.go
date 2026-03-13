package dotocr

import (
	"image"

	"github.com/ZingYao/autogo_scriptengine/js_engine/model"

	"github.com/Dasongzi1366/AutoGo/dotocr"
	"github.com/dop251/goja"
)

// DotocrModule dotocr 模块
type DotocrModule struct{}

// Name 返回模块名称
func (m *DotocrModule) Name() string {
	return "dotocr"
}

// IsAvailable 检查模块是否可用
func (m *DotocrModule) IsAvailable() bool {
	return true
}

// Register 向引擎注册方法
func (m *DotocrModule) Register(engine model.Engine) error {
	vm := engine.GetVM()

	dotocrObj := vm.NewObject()
	vm.Set("dotocr", dotocrObj)

	dotocrObj.Set("setDict", func(call goja.FunctionCall) goja.Value {
		name := call.Argument(0).String()
		dict := call.Argument(1).String()
		dotocr.SetDict(name, dict)
		return goja.Undefined()
	})

	dotocrObj.Set("ocr", func(call goja.FunctionCall) goja.Value {
		x1 := int(call.Argument(0).ToInteger())
		y1 := int(call.Argument(1).ToInteger())
		x2 := int(call.Argument(2).ToInteger())
		y2 := int(call.Argument(3).ToInteger())
		threshold := call.Argument(4).String()
		colGap := int(call.Argument(5).ToInteger())
		rowGap := int(call.Argument(6).ToInteger())
		sim := float32(call.Argument(7).ToFloat())
		mode := int(call.Argument(8).ToInteger())
		dictName := call.Argument(9).String()
		displayId := int(call.Argument(10).ToInteger())
		result := dotocr.Ocr(x1, y1, x2, y2, threshold, colGap, rowGap, sim, mode, dictName, displayId)
		return vm.ToValue(result)
	})

	dotocrObj.Set("ocrFromImage", func(call goja.FunctionCall) goja.Value {
		img := call.Argument(0).Export().(*image.NRGBA)
		threshold := call.Argument(1).String()
		colGap := int(call.Argument(2).ToInteger())
		rowGap := int(call.Argument(3).ToInteger())
		sim := float32(call.Argument(4).ToFloat())
		mode := int(call.Argument(5).ToInteger())
		dictName := call.Argument(6).String()
		result := dotocr.OcrFromImage(img, threshold, colGap, rowGap, sim, mode, dictName)
		return vm.ToValue(result)
	})

	dotocrObj.Set("ocrFromBase64", func(call goja.FunctionCall) goja.Value {
		b64 := call.Argument(0).String()
		threshold := call.Argument(1).String()
		colGap := int(call.Argument(2).ToInteger())
		rowGap := int(call.Argument(3).ToInteger())
		sim := float32(call.Argument(4).ToFloat())
		mode := int(call.Argument(5).ToInteger())
		dictName := call.Argument(6).String()
		result := dotocr.OcrFromBase64(b64, threshold, colGap, rowGap, sim, mode, dictName)
		return vm.ToValue(result)
	})

	dotocrObj.Set("ocrFromPath", func(call goja.FunctionCall) goja.Value {
		path := call.Argument(0).String()
		threshold := call.Argument(1).String()
		colGap := int(call.Argument(2).ToInteger())
		rowGap := int(call.Argument(3).ToInteger())
		sim := float32(call.Argument(4).ToFloat())
		mode := int(call.Argument(5).ToInteger())
		dictName := call.Argument(6).String()
		result := dotocr.OcrFromPath(path, threshold, colGap, rowGap, sim, mode, dictName)
		return vm.ToValue(result)
	})

	dotocrObj.Set("findStr", func(call goja.FunctionCall) goja.Value {
		x1 := int(call.Argument(0).ToInteger())
		y1 := int(call.Argument(1).ToInteger())
		x2 := int(call.Argument(2).ToInteger())
		y2 := int(call.Argument(3).ToInteger())
		text := call.Argument(4).String()
		threshold := call.Argument(5).String()
		colGap := int(call.Argument(6).ToInteger())
		rowGap := int(call.Argument(7).ToInteger())
		sim := float32(call.Argument(8).ToFloat())
		dictName := call.Argument(9).String()
		displayId := int(call.Argument(10).ToInteger())
		x, y := dotocr.FindStr(x1, y1, x2, y2, text, threshold, colGap, rowGap, sim, dictName, displayId)
		result := vm.NewObject()
		result.Set("x", x)
		result.Set("y", y)
		return result
	})

	dotocrObj.Set("findStrFromImage", func(call goja.FunctionCall) goja.Value {
		img := call.Argument(0).Export().(*image.NRGBA)
		text := call.Argument(1).String()
		threshold := call.Argument(2).String()
		colGap := int(call.Argument(3).ToInteger())
		rowGap := int(call.Argument(4).ToInteger())
		sim := float32(call.Argument(5).ToFloat())
		dictName := call.Argument(6).String()
		x, y := dotocr.FindStrFromImage(img, text, threshold, colGap, rowGap, sim, dictName)
		result := vm.NewObject()
		result.Set("x", x)
		result.Set("y", y)
		return result
	})

	dotocrObj.Set("findStrFromBase64", func(call goja.FunctionCall) goja.Value {
		b64 := call.Argument(0).String()
		text := call.Argument(1).String()
		threshold := call.Argument(2).String()
		colGap := int(call.Argument(3).ToInteger())
		rowGap := int(call.Argument(4).ToInteger())
		sim := float32(call.Argument(5).ToFloat())
		dictName := call.Argument(6).String()
		x, y := dotocr.FindStrFromBase64(b64, text, threshold, colGap, rowGap, sim, dictName)
		result := vm.NewObject()
		result.Set("x", x)
		result.Set("y", y)
		return result
	})

	dotocrObj.Set("findStrFromPath", func(call goja.FunctionCall) goja.Value {
		path := call.Argument(0).String()
		text := call.Argument(1).String()
		threshold := call.Argument(2).String()
		colGap := int(call.Argument(3).ToInteger())
		rowGap := int(call.Argument(4).ToInteger())
		sim := float32(call.Argument(5).ToFloat())
		dictName := call.Argument(6).String()
		x, y := dotocr.FindStrFromPath(path, text, threshold, colGap, rowGap, sim, dictName)
		result := vm.NewObject()
		result.Set("x", x)
		result.Set("y", y)
		return result
	})

	engine.RegisterMethod("dotocr.setDict", "设置字库", dotocr.SetDict, true)
	engine.RegisterMethod("dotocr.ocr", "从屏幕指定区域进行OCR识别", func(x1, y1, x2, y2 int, threshold string, colGap, rowGap int, sim float32, mode int, dictName string, displayId int) string {
		return dotocr.Ocr(x1, y1, x2, y2, threshold, colGap, rowGap, sim, mode, dictName, displayId)
	}, true)
	engine.RegisterMethod("dotocr.ocrFromImage", "从图像进行OCR识别", func(img *image.NRGBA, threshold string, colGap, rowGap int, sim float32, mode int, dictName string) string {
		return dotocr.OcrFromImage(img, threshold, colGap, rowGap, sim, mode, dictName)
	}, true)
	engine.RegisterMethod("dotocr.ocrFromBase64", "从Base64编码的图像字符串进行OCR识别", dotocr.OcrFromBase64, true)
	engine.RegisterMethod("dotocr.ocrFromPath", "从图像文件路径进行OCR识别", dotocr.OcrFromPath, true)
	engine.RegisterMethod("dotocr.findStr", "在屏幕指定区域中查找指定字符串的位置", dotocr.FindStr, true)
	engine.RegisterMethod("dotocr.findStrFromImage", "在图像中查找指定字符串的位置", dotocr.FindStrFromImage, true)
	engine.RegisterMethod("dotocr.findStrFromBase64", "在Base64编码的图像中查找指定字符串的位置", dotocr.FindStrFromBase64, true)
	engine.RegisterMethod("dotocr.findStrFromPath", "在图像文件中查找指定字符串的位置", dotocr.FindStrFromPath, true)

	return nil
}
