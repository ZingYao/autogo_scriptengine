package ppocr

import (
	"app/js_engine/model"
	"image"
	"strconv"

	"github.com/Dasongzi1366/AutoGo/ppocr"
	"github.com/dop251/goja"
)

// PpocrModule ppocr 模块
type PpocrModule struct{}

// Name 返回模块名称
func (m *PpocrModule) Name() string {
	return "ppocr"
}

// IsAvailable 检查模块是否可用
func (m *PpocrModule) IsAvailable() bool {
	return true
}

// Register 向引擎注册方法
func (m *PpocrModule) Register(engine model.Engine) error {
	vm := engine.GetVM()

	ppocrObj := vm.NewObject()
	vm.Set("ppocr", ppocrObj)

	ppocrObj.Set("new", func(call goja.FunctionCall) goja.Value {
		version := call.Argument(0).String()
		result := ppocr.New(version)
		return vm.ToValue(result)
	})

	ppocrObj.Set("ocr", func(call goja.FunctionCall) goja.Value {
		p := call.Argument(0).Export().(*ppocr.Ppocr)
		x1 := int(call.Argument(1).ToInteger())
		y1 := int(call.Argument(2).ToInteger())
		x2 := int(call.Argument(3).ToInteger())
		y2 := int(call.Argument(4).ToInteger())
		colorStr := call.Argument(5).String()
		displayId := 0
		if len(call.Arguments) > 6 {
			displayId = int(call.Argument(6).ToInteger())
		}
		result := p.Ocr(x1, y1, x2, y2, colorStr, displayId)
		arr := vm.NewArray()
		for i, item := range result {
			obj := vm.NewObject()
			obj.Set("text", item.Label)
			obj.Set("x", item.X)
			obj.Set("y", item.Y)
			obj.Set("w", item.Width)
			obj.Set("h", item.Height)
			arr.Set(strconv.Itoa(i), obj)
		}
		return arr
	})

	ppocrObj.Set("ocrFromImage", func(call goja.FunctionCall) goja.Value {
		p := call.Argument(0).Export().(*ppocr.Ppocr)
		img := call.Argument(1).Export().(*image.NRGBA)
		colorStr := call.Argument(2).String()
		result := p.OcrFromImage(img, colorStr)
		arr := vm.NewArray()
		for i, item := range result {
			obj := vm.NewObject()
			obj.Set("text", item.Label)
			obj.Set("x", item.X)
			obj.Set("y", item.Y)
			obj.Set("w", item.Width)
			obj.Set("h", item.Height)
			arr.Set(strconv.Itoa(i), obj)
		}
		return arr
	})

	ppocrObj.Set("ocrFromBase64", func(call goja.FunctionCall) goja.Value {
		p := call.Argument(0).Export().(*ppocr.Ppocr)
		b64 := call.Argument(1).String()
		colorStr := call.Argument(2).String()
		result := p.OcrFromBase64(b64, colorStr)
		arr := vm.NewArray()
		for i, item := range result {
			obj := vm.NewObject()
			obj.Set("text", item.Label)
			obj.Set("x", item.X)
			obj.Set("y", item.Y)
			obj.Set("w", item.Width)
			obj.Set("h", item.Height)
			arr.Set(strconv.Itoa(i), obj)
		}
		return arr
	})

	ppocrObj.Set("ocrFromPath", func(call goja.FunctionCall) goja.Value {
		p := call.Argument(0).Export().(*ppocr.Ppocr)
		path := call.Argument(1).String()
		colorStr := call.Argument(2).String()
		result := p.OcrFromPath(path, colorStr)
		arr := vm.NewArray()
		for i, item := range result {
			obj := vm.NewObject()
			obj.Set("text", item.Label)
			obj.Set("x", item.X)
			obj.Set("y", item.Y)
			obj.Set("w", item.Width)
			obj.Set("h", item.Height)
			arr.Set(strconv.Itoa(i), obj)
		}
		return arr
	})

	ppocrObj.Set("close", func(call goja.FunctionCall) goja.Value {
		p := call.Argument(0).Export().(*ppocr.Ppocr)
		p.Close()
		return goja.Undefined()
	})

	engine.RegisterMethod("ppocr.ocr", "识别屏幕文字", func(x1, y1, x2, y2 int, colorStr string, displayId int) []map[string]interface{} {
		return []map[string]interface{}{}
	}, true)
	engine.RegisterMethod("ppocr.ocrFromImage", "识别图片文字", func(img interface{}, colorStr string) []map[string]interface{} {
		return []map[string]interface{}{}
	}, true)
	engine.RegisterMethod("ppocr.ocrFromBase64", "识别Base64图片文字", func(b64, colorStr string) []map[string]interface{} {
		return []map[string]interface{}{}
	}, true)
	engine.RegisterMethod("ppocr.ocrFromPath", "识别文件图片文字", func(path, colorStr string) []map[string]interface{} {
		return []map[string]interface{}{}
	}, true)

	return nil
}
