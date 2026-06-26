package images

import (
	"image"

	"github.com/ZingYao/autogo_scriptengine/js_engine/model"

	autogoimages "github.com/Dasongzi1366/AutoGo/images"
	"github.com/dop251/goja"
)

type ImagesModule struct{}

func (m *ImagesModule) Name() string      { return "images" }
func (m *ImagesModule) IsAvailable() bool { return true }

func bytesFromValue(value goja.Value) []byte {
	switch data := value.Export().(type) {
	case []byte:
		return data
	case string:
		return []byte(data)
	case []interface{}:
		result := make([]byte, 0, len(data))
		for _, item := range data {
			switch number := item.(type) {
			case int:
				result = append(result, byte(number))
			case int64:
				result = append(result, byte(number))
			case float64:
				result = append(result, byte(number))
			}
		}
		return result
	default:
		return nil
	}
}

func nrgbaFromValue(value goja.Value) *image.NRGBA {
	if goja.IsUndefined(value) || goja.IsNull(value) {
		return nil
	}
	return value.Export().(*image.NRGBA)
}

func imageFromValue(value goja.Value) image.Image {
	if goja.IsUndefined(value) || goja.IsNull(value) {
		return nil
	}
	return value.Export().(image.Image)
}

func pointToJS(vm *goja.Runtime, x, y int) goja.Value {
	return vm.ToValue(map[string]int{"x": x, "y": y})
}

func pointsToJS(vm *goja.Runtime, points []autogoimages.Point) goja.Value {
	rows := make([]map[string]int, 0, len(points))
	for _, point := range points {
		rows = append(rows, map[string]int{"x": point.X, "y": point.Y})
	}
	return vm.ToValue(rows)
}

func imageValue(vm *goja.Runtime, img any) goja.Value {
	if img == nil {
		return goja.Null()
	}
	return vm.ToValue(img)
}

func (m *ImagesModule) Register(engine model.Engine) error {
	vm := engine.GetVM()
	imagesObj := vm.NewObject()
	vm.Set("images", imagesObj)

	imagesObj.Set("setCallback", func(call goja.FunctionCall) goja.Value {
		fn, ok := goja.AssertFunction(call.Argument(0))
		if !ok {
			autogoimages.SetCallback(nil)
			return goja.Undefined()
		}
		autogoimages.SetCallback(func(img *image.NRGBA) {
			_, _ = fn(goja.Undefined(), imageValue(vm, img))
		})
		return goja.Undefined()
	})
	imagesObj.Set("captureScreen", func(call goja.FunctionCall) goja.Value {
		return imageValue(vm, autogoimages.CaptureScreen(int(call.Argument(0).ToInteger()), int(call.Argument(1).ToInteger()), int(call.Argument(2).ToInteger()), int(call.Argument(3).ToInteger())))
	})
	imagesObj.Set("pixel", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(autogoimages.Pixel(int(call.Argument(0).ToInteger()), int(call.Argument(1).ToInteger())))
	})
	imagesObj.Set("cmpColor", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(autogoimages.CmpColor(int(call.Argument(0).ToInteger()), int(call.Argument(1).ToInteger()), call.Argument(2).String(), float32(call.Argument(3).ToFloat())))
	})
	imagesObj.Set("findColor", func(call goja.FunctionCall) goja.Value {
		x, y := autogoimages.FindColor(int(call.Argument(0).ToInteger()), int(call.Argument(1).ToInteger()), int(call.Argument(2).ToInteger()), int(call.Argument(3).ToInteger()), call.Argument(4).String(), float32(call.Argument(5).ToFloat()), int(call.Argument(6).ToInteger()))
		return pointToJS(vm, x, y)
	})
	imagesObj.Set("getColorCountInRegion", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(autogoimages.GetColorCountInRegion(int(call.Argument(0).ToInteger()), int(call.Argument(1).ToInteger()), int(call.Argument(2).ToInteger()), int(call.Argument(3).ToInteger()), call.Argument(4).String(), float32(call.Argument(5).ToFloat())))
	})
	imagesObj.Set("detectsMultiColors", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(autogoimages.DetectsMultiColors(call.Argument(0).String(), float32(call.Argument(1).ToFloat())))
	})
	imagesObj.Set("findMultiColors", func(call goja.FunctionCall) goja.Value {
		x, y := autogoimages.FindMultiColors(int(call.Argument(0).ToInteger()), int(call.Argument(1).ToInteger()), int(call.Argument(2).ToInteger()), int(call.Argument(3).ToInteger()), call.Argument(4).String(), float32(call.Argument(5).ToFloat()), int(call.Argument(6).ToInteger()))
		return pointToJS(vm, x, y)
	})
	imagesObj.Set("findMultiColorsAll", func(call goja.FunctionCall) goja.Value {
		return pointsToJS(vm, autogoimages.FindMultiColorsAll(int(call.Argument(0).ToInteger()), int(call.Argument(1).ToInteger()), int(call.Argument(2).ToInteger()), int(call.Argument(3).ToInteger()), call.Argument(4).String(), float32(call.Argument(5).ToFloat()), int(call.Argument(6).ToInteger())))
	})
	imagesObj.Set("readFromPath", func(call goja.FunctionCall) goja.Value {
		return imageValue(vm, autogoimages.ReadFromPath(call.Argument(0).String()))
	})
	imagesObj.Set("readFromUrl", func(call goja.FunctionCall) goja.Value {
		return imageValue(vm, autogoimages.ReadFromUrl(call.Argument(0).String()))
	})
	imagesObj.Set("readFromBase64", func(call goja.FunctionCall) goja.Value {
		return imageValue(vm, autogoimages.ReadFromBase64(call.Argument(0).String()))
	})
	imagesObj.Set("readFromBytes", func(call goja.FunctionCall) goja.Value {
		return imageValue(vm, autogoimages.ReadFromBytes(bytesFromValue(call.Argument(0))))
	})
	imagesObj.Set("save", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(autogoimages.Save(nrgbaFromValue(call.Argument(0)), call.Argument(1).String(), int(call.Argument(2).ToInteger())))
	})
	imagesObj.Set("encodeToBase64", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(autogoimages.EncodeToBase64(nrgbaFromValue(call.Argument(0)), call.Argument(1).String(), int(call.Argument(2).ToInteger())))
	})
	imagesObj.Set("encodeToBytes", func(call goja.FunctionCall) goja.Value {
		data := autogoimages.EncodeToBytes(nrgbaFromValue(call.Argument(0)), call.Argument(1).String(), int(call.Argument(2).ToInteger()))
		if data == nil {
			return goja.Null()
		}
		return vm.ToValue(data)
	})
	imagesObj.Set("toNrgba", func(call goja.FunctionCall) goja.Value {
		return imageValue(vm, autogoimages.ToNrgba(imageFromValue(call.Argument(0))))
	})
	imagesObj.Set("clip", func(call goja.FunctionCall) goja.Value {
		return imageValue(vm, autogoimages.Clip(nrgbaFromValue(call.Argument(0)), int(call.Argument(1).ToInteger()), int(call.Argument(2).ToInteger()), int(call.Argument(3).ToInteger()), int(call.Argument(4).ToInteger())))
	})
	imagesObj.Set("resize", func(call goja.FunctionCall) goja.Value {
		return imageValue(vm, autogoimages.Resize(nrgbaFromValue(call.Argument(0)), int(call.Argument(1).ToInteger()), int(call.Argument(2).ToInteger())))
	})
	imagesObj.Set("rotate", func(call goja.FunctionCall) goja.Value {
		return imageValue(vm, autogoimages.Rotate(nrgbaFromValue(call.Argument(0)), int(call.Argument(1).ToInteger())))
	})
	imagesObj.Set("grayscale", func(call goja.FunctionCall) goja.Value {
		return imageValue(vm, autogoimages.Grayscale(nrgbaFromValue(call.Argument(0))))
	})
	imagesObj.Set("applyThreshold", func(call goja.FunctionCall) goja.Value {
		return imageValue(vm, autogoimages.ApplyThreshold(nrgbaFromValue(call.Argument(0)), int(call.Argument(1).ToInteger()), int(call.Argument(2).ToInteger()), call.Argument(3).String()))
	})
	imagesObj.Set("applyAdaptiveThreshold", func(call goja.FunctionCall) goja.Value {
		return imageValue(vm, autogoimages.ApplyAdaptiveThreshold(nrgbaFromValue(call.Argument(0)), call.Argument(1).ToFloat(), call.Argument(2).String(), call.Argument(3).String(), int(call.Argument(4).ToInteger()), call.Argument(5).ToFloat()))
	})
	imagesObj.Set("applyBinarization", func(call goja.FunctionCall) goja.Value {
		return imageValue(vm, autogoimages.ApplyBinarization(nrgbaFromValue(call.Argument(0)), int(call.Argument(1).ToInteger())))
	})

	engine.RegisterMethod("images.setCallback", "设置图像回调", autogoimages.SetCallback, true)
	engine.RegisterMethod("images.captureScreen", "截取屏幕", autogoimages.CaptureScreen, true)
	engine.RegisterMethod("images.pixel", "获取像素颜色", autogoimages.Pixel, true)
	engine.RegisterMethod("images.cmpColor", "比较颜色", autogoimages.CmpColor, true)
	engine.RegisterMethod("images.findColor", "查找颜色", autogoimages.FindColor, true)
	engine.RegisterMethod("images.getColorCountInRegion", "获取区域颜色数量", autogoimages.GetColorCountInRegion, true)
	engine.RegisterMethod("images.detectsMultiColors", "检测多点颜色", autogoimages.DetectsMultiColors, true)
	engine.RegisterMethod("images.findMultiColors", "查找多点颜色", autogoimages.FindMultiColors, true)
	engine.RegisterMethod("images.findMultiColorsAll", "查找所有多点颜色", autogoimages.FindMultiColorsAll, true)
	engine.RegisterMethod("images.readFromPath", "从路径读取图片", autogoimages.ReadFromPath, true)
	engine.RegisterMethod("images.readFromUrl", "从 URL 读取图片", autogoimages.ReadFromUrl, true)
	engine.RegisterMethod("images.readFromBase64", "从 Base64 读取图片", autogoimages.ReadFromBase64, true)
	engine.RegisterMethod("images.readFromBytes", "从字节数组读取图片", autogoimages.ReadFromBytes, true)
	engine.RegisterMethod("images.save", "保存图片", autogoimages.Save, true)
	engine.RegisterMethod("images.encodeToBase64", "编码为 Base64", autogoimages.EncodeToBase64, true)
	engine.RegisterMethod("images.encodeToBytes", "编码为字节数组", autogoimages.EncodeToBytes, true)
	engine.RegisterMethod("images.toNrgba", "转换为 NRGBA", autogoimages.ToNrgba, true)
	engine.RegisterMethod("images.clip", "裁剪图片", autogoimages.Clip, true)
	engine.RegisterMethod("images.resize", "调整图片大小", autogoimages.Resize, true)
	engine.RegisterMethod("images.rotate", "旋转图片", autogoimages.Rotate, true)
	engine.RegisterMethod("images.grayscale", "灰度化图片", autogoimages.Grayscale, true)
	engine.RegisterMethod("images.applyThreshold", "应用阈值", autogoimages.ApplyThreshold, true)
	engine.RegisterMethod("images.applyAdaptiveThreshold", "应用自适应阈值", autogoimages.ApplyAdaptiveThreshold, true)
	engine.RegisterMethod("images.applyBinarization", "二值化图片", autogoimages.ApplyBinarization, true)
	return nil
}
