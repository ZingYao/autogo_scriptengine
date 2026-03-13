package yolo

import (
	"image"
	"strconv"

	"github.com/ZingYao/autogo_scriptengine/js_engine/model"

	"github.com/Dasongzi1366/AutoGo/yolo"
	"github.com/dop251/goja"
)

// YoloModule yolo 模块
type YoloModule struct{}

// Name 返回模块名称
func (m *YoloModule) Name() string {
	return "yolo"
}

// IsAvailable 检查模块是否可用
func (m *YoloModule) IsAvailable() bool {
	return true
}

// Register 向引擎注册方法
func (m *YoloModule) Register(engine model.Engine) error {
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

	yoloObj.Set("detect", func(call goja.FunctionCall) goja.Value {
		y := call.Argument(0).Export().(*yolo.Yolo)
		x1 := int(call.Argument(1).ToInteger())
		y1 := int(call.Argument(2).ToInteger())
		x2 := int(call.Argument(3).ToInteger())
		y2 := int(call.Argument(4).ToInteger())
		displayId := 0
		if len(call.Arguments) > 5 {
			displayId = int(call.Argument(5).ToInteger())
		}
		result := y.Detect(x1, y1, x2, y2, displayId)
		arr := vm.NewArray()
		for i, item := range result {
			obj := vm.NewObject()
			obj.Set("x", item.X)
			obj.Set("y", item.Y)
			obj.Set("w", item.Width)
			obj.Set("h", item.Height)
			obj.Set("label", item.Label)
			obj.Set("score", item.Score)
			arr.Set(strconv.Itoa(i), obj)
		}
		return arr
	})

	yoloObj.Set("detectFromImage", func(call goja.FunctionCall) goja.Value {
		y := call.Argument(0).Export().(*yolo.Yolo)
		img := call.Argument(1).Export().(*image.NRGBA)
		result := y.DetectFromImage(img)
		arr := vm.NewArray()
		for i, item := range result {
			obj := vm.NewObject()
			obj.Set("x", item.X)
			obj.Set("y", item.Y)
			obj.Set("w", item.Width)
			obj.Set("h", item.Height)
			obj.Set("label", item.Label)
			obj.Set("score", item.Score)
			arr.Set(strconv.Itoa(i), obj)
		}
		return arr
	})

	yoloObj.Set("detectFromBase64", func(call goja.FunctionCall) goja.Value {
		y := call.Argument(0).Export().(*yolo.Yolo)
		b64 := call.Argument(1).String()
		colorStr := ""
		if len(call.Arguments) > 2 {
			colorStr = call.Argument(2).String()
		}
		result := y.DetectFromBase64(b64, colorStr)
		arr := vm.NewArray()
		for i, item := range result {
			obj := vm.NewObject()
			obj.Set("x", item.X)
			obj.Set("y", item.Y)
			obj.Set("w", item.Width)
			obj.Set("h", item.Height)
			obj.Set("label", item.Label)
			obj.Set("score", item.Score)
			arr.Set(strconv.Itoa(i), obj)
		}
		return arr
	})

	yoloObj.Set("detectFromPath", func(call goja.FunctionCall) goja.Value {
		y := call.Argument(0).Export().(*yolo.Yolo)
		path := call.Argument(1).String()
		colorStr := ""
		if len(call.Arguments) > 2 {
			colorStr = call.Argument(2).String()
		}
		result := y.DetectFromPath(path, colorStr)
		arr := vm.NewArray()
		for i, item := range result {
			obj := vm.NewObject()
			obj.Set("x", item.X)
			obj.Set("y", item.Y)
			obj.Set("w", item.Width)
			obj.Set("h", item.Height)
			obj.Set("label", item.Label)
			obj.Set("score", item.Score)
			arr.Set(strconv.Itoa(i), obj)
		}
		return arr
	})

	yoloObj.Set("close", func(call goja.FunctionCall) goja.Value {
		y := call.Argument(0).Export().(*yolo.Yolo)
		y.Close()
		return goja.Undefined()
	})

	engine.RegisterMethod("yolo.new", "创建一个新的YOLO实例", func(version string, cpuThreadNum int, paramPath, binPath, labels string) *yolo.Yolo {
		return yolo.New(version, cpuThreadNum, paramPath, binPath, labels)
	}, true)
	engine.RegisterMethod("yolo.detect", "检测屏幕上的对象", func(y *yolo.Yolo, x1, y1, x2, y2, displayId int) []yolo.Result {
		return y.Detect(x1, y1, x2, y2, displayId)
	}, true)
	engine.RegisterMethod("yolo.detectFromImage", "检测图片中的对象", func(y *yolo.Yolo, img *image.NRGBA) []yolo.Result {
		return y.DetectFromImage(img)
	}, true)
	engine.RegisterMethod("yolo.detectFromBase64", "检测Base64图片中的对象", func(y *yolo.Yolo, b64 string, colorStr string) []yolo.Result {
		return y.DetectFromBase64(b64, colorStr)
	}, true)
	engine.RegisterMethod("yolo.detectFromPath", "检测文件图片中的对象", func(y *yolo.Yolo, path string, colorStr string) []yolo.Result {
		return y.DetectFromPath(path, colorStr)
	}, true)
	engine.RegisterMethod("yolo.close", "关闭YOLO实例", (*yolo.Yolo).Close, true)

	return nil
}
