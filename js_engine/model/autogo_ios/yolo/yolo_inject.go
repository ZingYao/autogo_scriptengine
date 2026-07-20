package yolo

import (
	"image"

	"github.com/ZingYao/autogo_scriptengine/js_engine/model"

	autogoyolo "github.com/Dasongzi1366/AutoGo/yolo"
	"github.com/ZingYao/goja"
)

type YoloModule struct{}

func (m *YoloModule) Name() string      { return "yolo" }
func (m *YoloModule) IsAvailable() bool { return true }

func resultsToJS(vm *goja.Runtime, results []autogoyolo.Result) goja.Value {
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

func wrapYolo(vm *goja.Runtime, y *autogoyolo.Yolo) goja.Value {
	obj := vm.NewObject()
	obj.Set("setImage", func(call goja.FunctionCall) goja.Value {
		y.SetImage(call.Argument(0).Export().(*image.NRGBA))
		return goja.Undefined()
	})
	obj.Set("detect", func(call goja.FunctionCall) goja.Value {
		return resultsToJS(vm, y.Detect(
			int(call.Argument(0).ToInteger()),
			int(call.Argument(1).ToInteger()),
			int(call.Argument(2).ToInteger()),
			int(call.Argument(3).ToInteger()),
		))
	})
	obj.Set("detectFromImage", func(call goja.FunctionCall) goja.Value {
		return resultsToJS(vm, y.DetectFromImage(call.Argument(0).Export().(*image.NRGBA)))
	})
	obj.Set("detectFromBase64", func(call goja.FunctionCall) goja.Value {
		return resultsToJS(vm, y.DetectFromBase64(call.Argument(0).String()))
	})
	obj.Set("detectFromPath", func(call goja.FunctionCall) goja.Value {
		return resultsToJS(vm, y.DetectFromPath(call.Argument(0).String()))
	})
	obj.Set("close", func(call goja.FunctionCall) goja.Value {
		y.Close()
		return goja.Undefined()
	})
	return obj
}

func (m *YoloModule) Register(engine model.Engine) error {
	vm := engine.GetVM()
	yoloObj := vm.NewObject()
	vm.Set("yolo", yoloObj)
	yoloObj.Set("new", func(call goja.FunctionCall) goja.Value {
		y := autogoyolo.New(call.Argument(0).String(), int(call.Argument(1).ToInteger()), call.Argument(2).String(), call.Argument(3).String(), call.Argument(4).String())
		if y == nil {
			return goja.Null()
		}
		return wrapYolo(vm, y)
	})

	engine.RegisterMethod("yolo.new", "创建 YOLO 实例", autogoyolo.New, true)
	engine.RegisterMethod("yolo.setImage", "设置下次 detect 的原始图像", func(y *autogoyolo.Yolo, img *image.NRGBA) { y.SetImage(img) }, true)
	engine.RegisterMethod("yolo.detect", "检测屏幕区域中的对象", func(y *autogoyolo.Yolo, x1, y1, x2, y2 int) []autogoyolo.Result {
		return y.Detect(x1, y1, x2, y2)
	}, true)
	engine.RegisterMethod("yolo.detectFromImage", "检测图像对象中的对象", func(y *autogoyolo.Yolo, img *image.NRGBA) []autogoyolo.Result {
		return y.DetectFromImage(img)
	}, true)
	engine.RegisterMethod("yolo.detectFromBase64", "检测 Base64 图像中的对象", func(y *autogoyolo.Yolo, b64 string) []autogoyolo.Result {
		return y.DetectFromBase64(b64)
	}, true)
	engine.RegisterMethod("yolo.detectFromPath", "检测文件图像中的对象", func(y *autogoyolo.Yolo, path string) []autogoyolo.Result {
		return y.DetectFromPath(path)
	}, true)
	engine.RegisterMethod("yolo.close", "关闭 YOLO 实例", func(y *autogoyolo.Yolo) { y.Close() }, true)
	return nil
}
