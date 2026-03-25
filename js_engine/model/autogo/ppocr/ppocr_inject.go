package ppocr

import (
	"image"
	"strconv"

	"github.com/ZingYao/autogo_scriptengine/js_engine/model"

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
		if result != nil {
			return wrapPpocr(vm, result)
		}
		return goja.Null()
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

func wrapPpocr(vm *goja.Runtime, p *ppocr.Ppocr) goja.Value {
	obj := vm.NewObject()

	obj.Set("ocr", func(call goja.FunctionCall) goja.Value {
		x1 := int(call.Argument(0).ToInteger())
		y1 := int(call.Argument(1).ToInteger())
		x2 := int(call.Argument(2).ToInteger())
		y2 := int(call.Argument(3).ToInteger())
		colorStr := call.Argument(4).String()
		displayId := 0
		if len(call.Arguments) > 5 {
			displayId = int(call.Argument(5).ToInteger())
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

	obj.Set("ocrFromImage", func(call goja.FunctionCall) goja.Value {
		img := call.Argument(0).Export().(*image.NRGBA)
		colorStr := call.Argument(1).String()
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

	obj.Set("ocrFromBase64", func(call goja.FunctionCall) goja.Value {
		b64 := call.Argument(0).String()
		colorStr := call.Argument(1).String()
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

	obj.Set("ocrFromPath", func(call goja.FunctionCall) goja.Value {
		path := call.Argument(0).String()
		colorStr := call.Argument(1).String()
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

	obj.Set("close", func(call goja.FunctionCall) goja.Value {
		p.Close()
		return goja.Undefined()
	})

	return vm.ToValue(obj)
}
