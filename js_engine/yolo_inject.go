package js_engine

import (
	"github.com/Dasongzi1366/AutoGo/yolo"
	"github.com/dop251/goja"
)

func injectYoloMethods(engine *JSEngine) {
	vm := engine.GetVM()

	yoloObj := vm.NewObject()
	vm.Set("yolo", yoloObj)

	yoloObj.Set("new", func(call goja.FunctionCall) goja.Value {
		version := call.Argument(0).String()
		cpuThreadNum := int(call.Argument(1).ToInteger())
		paramPath := call.Argument(2).String()
		binPath := call.Argument(3).String()
		labels := call.Argument(4).String()
		result := yolo.New(version, cpuThreadNum, paramPath, binPath, labels)
		return vm.ToValue(result)
	})

	yoloObj.Set("setImage", func(call goja.FunctionCall) goja.Value {
		y := call.Argument(0).Export().(*yolo.Yolo)
		img := call.Argument(1).Export()
		if imgNRGBA, ok := img.(*yoloImgWrapper); ok {
			if nrgba, ok := imgNRGBA.img.(interface{ SetImage(interface{}) }); ok {
				nrgba.SetImage(y)
			}
		}
		return goja.Undefined()
	})

	yoloObj.Set("detect", func(call goja.FunctionCall) goja.Value {
		y := call.Argument(0).Export().(*yolo.Yolo)
		x1 := int(call.Argument(1).ToInteger())
		y1 := int(call.Argument(2).ToInteger())
		x2 := int(call.Argument(3).ToInteger())
		y2 := int(call.Argument(4).ToInteger())
		displayId := int(call.Argument(5).ToInteger())
		results := y.Detect(x1, y1, x2, y2, displayId)
		return vm.ToValue(results)
	})

	yoloObj.Set("detectFromBase64", func(call goja.FunctionCall) goja.Value {
		y := call.Argument(0).Export().(*yolo.Yolo)
		b64 := call.Argument(1).String()
		colorStr := call.Argument(2).String()
		results := y.DetectFromBase64(b64, colorStr)
		return vm.ToValue(results)
	})

	yoloObj.Set("detectFromPath", func(call goja.FunctionCall) goja.Value {
		y := call.Argument(0).Export().(*yolo.Yolo)
		path := call.Argument(1).String()
		colorStr := call.Argument(2).String()
		results := y.DetectFromPath(path, colorStr)
		return vm.ToValue(results)
	})

	yoloObj.Set("close", func(call goja.FunctionCall) goja.Value {
		y := call.Argument(0).Export().(*yolo.Yolo)
		y.Close()
		return goja.Undefined()
	})

	// 注册方法到文档
	engine.RegisterMethod("yolo.new", "创建一个新的YOLO实例", yolo.New, true)
	engine.RegisterMethod("yolo.detect", "在指定的屏幕区域执行目标检测", (*yolo.Yolo).Detect, true)
	engine.RegisterMethod("yolo.detectFromBase64", "从Base64编码的图像进行识别", (*yolo.Yolo).DetectFromBase64, true)
	engine.RegisterMethod("yolo.detectFromPath", "从文件路径进行识别", (*yolo.Yolo).DetectFromPath, true)
	engine.RegisterMethod("yolo.close", "关闭YOLO模型实例，释放相关资源", (*yolo.Yolo).Close, true)
}

// yoloImgWrapper 用于包装图像对象
type yoloImgWrapper struct {
	img interface{}
}
