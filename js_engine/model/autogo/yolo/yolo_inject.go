package yolo

import (
	"fmt"
	"image"
	"reflect"
	"strconv"

	"github.com/ZingYao/autogo_scriptengine/js_engine/model"

	"github.com/Dasongzi1366/AutoGo/yolo"
	"github.com/ZingYao/goja"
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
		if result == nil {
			panic("创建 YOLO 实例失败，请检查模型文件路径")
		}
		return vm.ToValue(result)
	})

	yoloObj.Set("detect", func(call goja.FunctionCall) goja.Value {
		y := call.Argument(0).Export().(*yolo.Yolo)
		x1 := int(call.Argument(1).ToInteger())
		y1 := int(call.Argument(2).ToInteger())
		x2 := int(call.Argument(3).ToInteger())
		y2 := int(call.Argument(4).ToInteger())
		return yoloResultsToArray(vm, detectYolo(y, x1, y1, x2, y2, displayID(call, 5)))
	})

	yoloObj.Set("detectFromImage", func(call goja.FunctionCall) goja.Value {
		y := call.Argument(0).Export().(*yolo.Yolo)
		img := call.Argument(1).Export().(*image.NRGBA)
		result, err := detectYoloFromImage(y, img, 0)
		if err != nil {
			panic(vm.NewGoError(err))
		}
		return yoloResultsToArray(vm, result)
	})

	yoloObj.Set("setImage", func(call goja.FunctionCall) goja.Value {
		y := call.Argument(0).Export().(*yolo.Yolo)
		img := call.Argument(1).Export().(*image.NRGBA)
		if err := setYoloImage(y, img); err != nil {
			panic(vm.NewGoError(err))
		}
		return goja.Undefined()
	})

	yoloObj.Set("detectFromBase64", func(call goja.FunctionCall) goja.Value {
		y := call.Argument(0).Export().(*yolo.Yolo)
		b64 := call.Argument(1).String()
		result, err := detectYoloByMethod(y, "DetectFromBase64", b64)
		if err != nil {
			panic(vm.NewGoError(err))
		}
		return yoloResultsToArray(vm, result)
	})

	yoloObj.Set("detectFromPath", func(call goja.FunctionCall) goja.Value {
		y := call.Argument(0).Export().(*yolo.Yolo)
		path := call.Argument(1).String()
		result, err := detectYoloByMethod(y, "DetectFromPath", path)
		if err != nil {
			panic(vm.NewGoError(err))
		}
		return yoloResultsToArray(vm, result)
	})

	yoloObj.Set("close", func(call goja.FunctionCall) goja.Value {
		y := call.Argument(0).Export().(*yolo.Yolo)
		y.Close()
		return goja.Undefined()
	})

	engine.RegisterMethod("yolo.new", "创建一个新的YOLO实例", func(version string, cpuThreadNum int, paramPath, binPath, labels string) *yolo.Yolo {
		return yolo.New(version, cpuThreadNum, paramPath, binPath, labels)
	}, true)
	engine.RegisterMethod("yolo.detect", "检测屏幕上的对象", func(y *yolo.Yolo, x1, y1, x2, y2, displayId int) []map[string]any {
		return detectYolo(y, x1, y1, x2, y2, displayId)
	}, true)
	engine.RegisterMethod("yolo.detectFromImage", "检测图片中的对象", func(y *yolo.Yolo, img *image.NRGBA) []map[string]any {
		result, err := detectYoloFromImage(y, img, 0)
		if err != nil {
			panic(err)
		}
		return result
	}, true)
	engine.RegisterMethod("yolo.setImage", "设置下次Detect方法的原始图像", func(y *yolo.Yolo, img *image.NRGBA) {
		if err := setYoloImage(y, img); err != nil {
			panic(err)
		}
	}, true)
	engine.RegisterMethod("yolo.detectFromBase64", "检测Base64图片中的对象", func(y *yolo.Yolo, b64 string) []map[string]any {
		result, err := detectYoloByMethod(y, "DetectFromBase64", b64)
		if err != nil {
			panic(err)
		}
		return result
	}, true)
	engine.RegisterMethod("yolo.detectFromPath", "检测文件图片中的对象", func(y *yolo.Yolo, path string) []map[string]any {
		result, err := detectYoloByMethod(y, "DetectFromPath", path)
		if err != nil {
			panic(err)
		}
		return result
	}, true)
	engine.RegisterMethod("yolo.close", "关闭YOLO实例", (*yolo.Yolo).Close, true)

	return nil
}

func yoloResultsToArray(vm *goja.Runtime, result []map[string]any) goja.Value {
	arr := vm.NewArray()
	for i, item := range result {
		obj := vm.NewObject()
		for key, value := range item {
			obj.Set(key, value)
		}
		arr.Set(strconv.Itoa(i), obj)
	}
	return arr
}

func detectYoloFromImage(y *yolo.Yolo, img *image.NRGBA, displayId int) ([]map[string]any, error) {
	if img == nil {
		return nil, fmt.Errorf("image is nil")
	}
	if result, err := detectYoloByMethod(y, "DetectFromImage", img); err == nil {
		return result, nil
	}
	if err := setYoloImage(y, img); err != nil {
		return nil, err
	}
	return detectYolo(y, 0, 0, img.Rect.Dx(), img.Rect.Dy(), displayId), nil
}

func detectYolo(y *yolo.Yolo, x1 int, y1 int, x2 int, y2 int, displayId int) []map[string]any {
	method := reflect.ValueOf(y).MethodByName("Detect")
	if !method.IsValid() {
		return nil
	}
	args := []reflect.Value{reflect.ValueOf(x1), reflect.ValueOf(y1), reflect.ValueOf(x2), reflect.ValueOf(y2)}
	if method.Type().NumIn() >= 5 {
		args = append(args, reflect.ValueOf(displayId))
	}
	return exportYoloResults(method.Call(args))
}

func detectYoloByMethod(y *yolo.Yolo, methodName string, arg any) ([]map[string]any, error) {
	method := reflect.ValueOf(y).MethodByName(methodName)
	if !method.IsValid() {
		return nil, fmt.Errorf("current AutoGo yolo.Yolo does not provide %s", methodName)
	}
	return exportYoloResults(method.Call([]reflect.Value{reflect.ValueOf(arg)})), nil
}

func exportYoloResults(values []reflect.Value) []map[string]any {
	if len(values) == 0 || values[0].Kind() != reflect.Slice {
		return nil
	}
	result := make([]map[string]any, 0, values[0].Len())
	for i := 0; i < values[0].Len(); i++ {
		item := values[0].Index(i)
		result = append(result, map[string]any{
			"x":     intField(item, "X"),
			"y":     intField(item, "Y"),
			"w":     intField(item, "Width"),
			"h":     intField(item, "Height"),
			"label": stringField(item, "Label"),
			"score": floatField(item, "Score"),
		})
	}
	return result
}

func intField(value reflect.Value, name string) int {
	field := value.FieldByName(name)
	if field.IsValid() && field.CanInt() {
		return int(field.Int())
	}
	return 0
}

func stringField(value reflect.Value, name string) string {
	field := value.FieldByName(name)
	if field.IsValid() && field.Kind() == reflect.String {
		return field.String()
	}
	return ""
}

func floatField(value reflect.Value, name string) float64 {
	field := value.FieldByName(name)
	if field.IsValid() && field.CanFloat() {
		return field.Float()
	}
	return 0
}

func displayID(call goja.FunctionCall, index int) int {
	if len(call.Arguments) > index {
		return int(call.Argument(index).ToInteger())
	}
	return 0
}

func setYoloImage(y *yolo.Yolo, img *image.NRGBA) error {
	method := reflect.ValueOf(y).MethodByName("SetImage")
	if !method.IsValid() {
		return fmt.Errorf("current AutoGo yolo.Yolo does not provide SetImage")
	}
	method.Call([]reflect.Value{reflect.ValueOf(img)})
	return nil
}
