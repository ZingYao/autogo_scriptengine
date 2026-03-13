package images

import (
	"image"

	"github.com/ZingYao/autogo_scriptengine/js_engine/model"

	"github.com/Dasongzi1366/AutoGo/images"
	"github.com/dop251/goja"
)

// ImagesModule images 模块
type ImagesModule struct{}

// Name 返回模块名称
func (m *ImagesModule) Name() string {
	return "images"
}

// IsAvailable 检查模块是否可用
func (m *ImagesModule) IsAvailable() bool {
	return true
}

// Register 向引擎注册方法
func (m *ImagesModule) Register(engine model.Engine) error {
	vm := engine.GetVM()

	imagesObj := vm.NewObject()
	vm.Set("images", imagesObj)

	imagesObj.Set("pixel", func(call goja.FunctionCall) goja.Value {
		x := int(call.Argument(0).ToInteger())
		y := int(call.Argument(1).ToInteger())
		displayId := 0
		if len(call.Arguments) >= 3 {
			displayId = int(call.Argument(2).ToInteger())
		}
		result := images.Pixel(x, y, displayId)
		return vm.ToValue(result)
	})

	imagesObj.Set("setCallback", func(call goja.FunctionCall) goja.Value {
		callback := call.Argument(0).Export()
		if fn, ok := callback.(goja.Callable); ok {
			images.SetCallback(func(img *image.NRGBA, displayId int) {
				fn(vm.ToValue(img), vm.ToValue(displayId))
			})
		}
		return goja.Undefined()
	})

	imagesObj.Set("captureScreen", func(call goja.FunctionCall) goja.Value {
		x1 := int(call.Argument(0).ToInteger())
		y1 := int(call.Argument(1).ToInteger())
		x2 := int(call.Argument(2).ToInteger())
		y2 := int(call.Argument(3).ToInteger())
		displayId := 0
		if len(call.Arguments) >= 5 {
			displayId = int(call.Argument(4).ToInteger())
		}
		result := images.CaptureScreen(x1, y1, x2, y2, displayId)
		if result != nil {
			return vm.ToValue(result)
		}
		return goja.Null()
	})

	imagesObj.Set("cmpColor", func(call goja.FunctionCall) goja.Value {
		x := int(call.Argument(0).ToInteger())
		y := int(call.Argument(1).ToInteger())
		colorStr := call.Argument(2).String()
		sim := float32(call.Argument(3).ToFloat())
		displayId := 0
		if len(call.Arguments) >= 5 {
			displayId = int(call.Argument(4).ToInteger())
		}
		result := images.CmpColor(x, y, colorStr, sim, displayId)
		return vm.ToValue(result)
	})

	imagesObj.Set("findColor", func(call goja.FunctionCall) goja.Value {
		x1 := int(call.Argument(0).ToInteger())
		y1 := int(call.Argument(1).ToInteger())
		x2 := int(call.Argument(2).ToInteger())
		y2 := int(call.Argument(3).ToInteger())
		colorStr := call.Argument(4).String()
		sim := float32(call.Argument(5).ToFloat())
		dir := int(call.Argument(6).ToInteger())
		displayId := 0
		if len(call.Arguments) >= 8 {
			displayId = int(call.Argument(7).ToInteger())
		}
		x, y := images.FindColor(x1, y1, x2, y2, colorStr, sim, dir, displayId)
		result := vm.NewObject()
		result.Set("x", x)
		result.Set("y", y)
		return vm.ToValue(result)
	})

	imagesObj.Set("getColorCountInRegion", func(call goja.FunctionCall) goja.Value {
		x1 := int(call.Argument(0).ToInteger())
		y1 := int(call.Argument(1).ToInteger())
		x2 := int(call.Argument(2).ToInteger())
		y2 := int(call.Argument(3).ToInteger())
		colorStr := call.Argument(4).String()
		sim := float32(call.Argument(5).ToFloat())
		displayId := 0
		if len(call.Arguments) >= 7 {
			displayId = int(call.Argument(6).ToInteger())
		}
		result := images.GetColorCountInRegion(x1, y1, x2, y2, colorStr, sim, displayId)
		return vm.ToValue(result)
	})

	imagesObj.Set("detectsMultiColors", func(call goja.FunctionCall) goja.Value {
		colors := call.Argument(0).String()
		sim := float32(call.Argument(1).ToFloat())
		displayId := 0
		if len(call.Arguments) >= 3 {
			displayId = int(call.Argument(2).ToInteger())
		}
		result := images.DetectsMultiColors(colors, sim, displayId)
		return vm.ToValue(result)
	})

	imagesObj.Set("findMultiColors", func(call goja.FunctionCall) goja.Value {
		x1 := int(call.Argument(0).ToInteger())
		y1 := int(call.Argument(1).ToInteger())
		x2 := int(call.Argument(2).ToInteger())
		y2 := int(call.Argument(3).ToInteger())
		colors := call.Argument(4).String()
		sim := float32(call.Argument(5).ToFloat())
		dir := int(call.Argument(6).ToInteger())
		displayId := 0
		if len(call.Arguments) >= 8 {
			displayId = int(call.Argument(7).ToInteger())
		}
		x, y := images.FindMultiColors(x1, y1, x2, y2, colors, sim, dir, displayId)
		result := vm.NewObject()
		result.Set("x", x)
		result.Set("y", y)
		return vm.ToValue(result)
	})

	imagesObj.Set("readFromPath", func(call goja.FunctionCall) goja.Value {
		path := call.Argument(0).String()
		result := images.ReadFromPath(path)
		if result != nil {
			return vm.ToValue(result)
		}
		return goja.Null()
	})

	imagesObj.Set("readFromUrl", func(call goja.FunctionCall) goja.Value {
		url := call.Argument(0).String()
		result := images.ReadFromUrl(url)
		if result != nil {
			return vm.ToValue(result)
		}
		return goja.Null()
	})

	imagesObj.Set("readFromBase64", func(call goja.FunctionCall) goja.Value {
		base64Str := call.Argument(0).String()
		result := images.ReadFromBase64(base64Str)
		if result != nil {
			return vm.ToValue(result)
		}
		return goja.Null()
	})

	imagesObj.Set("readFromBytes", func(call goja.FunctionCall) goja.Value {
		data := call.Argument(0).Export().([]byte)
		result := images.ReadFromBytes(data)
		if result != nil {
			return vm.ToValue(result)
		}
		return goja.Null()
	})

	imagesObj.Set("toNrgba", func(call goja.FunctionCall) goja.Value {
		img := call.Argument(0).Export().(image.Image)
		result := images.ToNrgba(img)
		if result != nil {
			return vm.ToValue(result)
		}
		return goja.Null()
	})

	imagesObj.Set("save", func(call goja.FunctionCall) goja.Value {
		img := call.Argument(0).Export().(*image.NRGBA)
		path := call.Argument(1).String()
		quality := int(call.Argument(2).ToInteger())
		result := images.Save(img, path, quality)
		return vm.ToValue(result)
	})

	imagesObj.Set("encodeToBase64", func(call goja.FunctionCall) goja.Value {
		img := call.Argument(0).Export().(*image.NRGBA)
		format := call.Argument(1).String()
		quality := int(call.Argument(2).ToInteger())
		result := images.EncodeToBase64(img, format, quality)
		return vm.ToValue(result)
	})

	imagesObj.Set("encodeToBytes", func(call goja.FunctionCall) goja.Value {
		img := call.Argument(0).Export().(*image.NRGBA)
		format := call.Argument(1).String()
		quality := int(call.Argument(2).ToInteger())
		result := images.EncodeToBytes(img, format, quality)
		if result != nil {
			return vm.ToValue(result)
		}
		return goja.Null()
	})

	imagesObj.Set("clip", func(call goja.FunctionCall) goja.Value {
		img := call.Argument(0).Export().(*image.NRGBA)
		x1 := int(call.Argument(1).ToInteger())
		y1 := int(call.Argument(2).ToInteger())
		x2 := int(call.Argument(3).ToInteger())
		y2 := int(call.Argument(4).ToInteger())
		result := images.Clip(img, x1, y1, x2, y2)
		if result != nil {
			return vm.ToValue(result)
		}
		return goja.Null()
	})

	imagesObj.Set("resize", func(call goja.FunctionCall) goja.Value {
		img := call.Argument(0).Export().(*image.NRGBA)
		width := int(call.Argument(1).ToInteger())
		height := int(call.Argument(2).ToInteger())
		result := images.Resize(img, width, height)
		if result != nil {
			return vm.ToValue(result)
		}
		return goja.Null()
	})

	imagesObj.Set("rotate", func(call goja.FunctionCall) goja.Value {
		img := call.Argument(0).Export().(*image.NRGBA)
		degree := int(call.Argument(1).ToInteger())
		result := images.Rotate(img, degree)
		if result != nil {
			return vm.ToValue(result)
		}
		return goja.Null()
	})

	imagesObj.Set("grayscale", func(call goja.FunctionCall) goja.Value {
		img := call.Argument(0).Export().(*image.NRGBA)
		result := images.Grayscale(img)
		if result != nil {
			return vm.ToValue(result)
		}
		return goja.Null()
	})

	imagesObj.Set("applyThreshold", func(call goja.FunctionCall) goja.Value {
		img := call.Argument(0).Export().(*image.NRGBA)
		threshold := int(call.Argument(1).ToInteger())
		maxVal := int(call.Argument(2).ToInteger())
		typ := call.Argument(3).String()
		result := images.ApplyThreshold(img, threshold, maxVal, typ)
		if result != nil {
			return vm.ToValue(result)
		}
		return goja.Null()
	})

	imagesObj.Set("applyAdaptiveThreshold", func(call goja.FunctionCall) goja.Value {
		img := call.Argument(0).Export().(*image.NRGBA)
		maxValue := call.Argument(1).ToFloat()
		adaptiveMethod := call.Argument(2).String()
		thresholdType := call.Argument(3).String()
		blockSize := int(call.Argument(4).ToInteger())
		C := call.Argument(5).ToFloat()
		result := images.ApplyAdaptiveThreshold(img, maxValue, adaptiveMethod, thresholdType, blockSize, C)
		if result != nil {
			return vm.ToValue(result)
		}
		return goja.Null()
	})

	imagesObj.Set("applyBinarization", func(call goja.FunctionCall) goja.Value {
		img := call.Argument(0).Export().(*image.NRGBA)
		threshold := int(call.Argument(1).ToInteger())
		result := images.ApplyBinarization(img, threshold)
		if result != nil {
			return vm.ToValue(result)
		}
		return goja.Null()
	})

	engine.RegisterMethod("images.pixel", "获取指定坐标的像素颜色", func(x, y int) string { return images.Pixel(x, y, 0) }, true)
	engine.RegisterMethod("images.setCallback", "设置回调函数", func(callback func(*image.NRGBA, int)) { images.SetCallback(callback) }, true)
	engine.RegisterMethod("images.captureScreen", "截取屏幕", func(x1, y1, x2, y2 int) *image.NRGBA {
		return images.CaptureScreen(x1, y1, x2, y2, 0)
	}, true)
	engine.RegisterMethod("images.cmpColor", "比较颜色", func(x, y int, colorStr string, sim float32) bool {
		return images.CmpColor(x, y, colorStr, sim, 0)
	}, true)
	engine.RegisterMethod("images.findColor", "查找颜色", func(x1, y1, x2, y2 int, colorStr string, sim float32, dir int) (int, int) {
		return images.FindColor(x1, y1, x2, y2, colorStr, sim, dir, 0)
	}, true)
	engine.RegisterMethod("images.getColorCountInRegion", "获取区域内指定颜色的数量", func(x1, y1, x2, y2 int, colorStr string, sim float32) int {
		return images.GetColorCountInRegion(x1, y1, x2, y2, colorStr, sim, 0)
	}, true)
	engine.RegisterMethod("images.detectsMultiColors", "检测多点颜色", func(colors string, sim float32) bool {
		return images.DetectsMultiColors(colors, sim, 0)
	}, true)
	engine.RegisterMethod("images.findMultiColors", "查找多点颜色", func(x1, y1, x2, y2 int, colors string, sim float32, dir int) (int, int) {
		return images.FindMultiColors(x1, y1, x2, y2, colors, sim, dir, 0)
	}, true)
	engine.RegisterMethod("images.readFromPath", "从路径读取图片", func(path string) *image.NRGBA { return images.ReadFromPath(path) }, true)
	engine.RegisterMethod("images.readFromUrl", "从URL读取图片", func(url string) *image.NRGBA { return images.ReadFromUrl(url) }, true)
	engine.RegisterMethod("images.readFromBase64", "从Base64读取图片", func(base64Str string) *image.NRGBA { return images.ReadFromBase64(base64Str) }, true)
	engine.RegisterMethod("images.readFromBytes", "从字节数组读取图片", func(data []byte) *image.NRGBA { return images.ReadFromBytes(data) }, true)
	engine.RegisterMethod("images.toNrgba", "转换为NRGBA格式", func(img image.Image) *image.NRGBA { return images.ToNrgba(img) }, true)
	engine.RegisterMethod("images.save", "保存图片", func(img *image.NRGBA, path string, quality int) bool { return images.Save(img, path, quality) }, true)
	engine.RegisterMethod("images.encodeToBase64", "编码为Base64", func(img *image.NRGBA, format string, quality int) string {
		return images.EncodeToBase64(img, format, quality)
	}, true)
	engine.RegisterMethod("images.encodeToBytes", "编码为字节数组", func(img *image.NRGBA, format string, quality int) []byte {
		return images.EncodeToBytes(img, format, quality)
	}, true)
	engine.RegisterMethod("images.clip", "裁剪图片", func(img *image.NRGBA, x1, y1, x2, y2 int) *image.NRGBA { return images.Clip(img, x1, y1, x2, y2) }, true)
	engine.RegisterMethod("images.resize", "调整图片大小", func(img *image.NRGBA, width, height int) *image.NRGBA { return images.Resize(img, width, height) }, true)
	engine.RegisterMethod("images.rotate", "旋转图片", func(img *image.NRGBA, degree int) *image.NRGBA { return images.Rotate(img, degree) }, true)
	engine.RegisterMethod("images.grayscale", "灰度化", func(img *image.NRGBA) *image.Gray { return images.Grayscale(img) }, true)
	engine.RegisterMethod("images.applyThreshold", "应用阈值", func(img *image.NRGBA, threshold, maxVal int, typ string) *image.Gray {
		return images.ApplyThreshold(img, threshold, maxVal, typ)
	}, true)
	engine.RegisterMethod("images.applyAdaptiveThreshold", "应用自适应阈值", func(img *image.NRGBA, maxValue float64, adaptiveMethod, thresholdType string, blockSize int, C float64) *image.Gray {
		return images.ApplyAdaptiveThreshold(img, maxValue, adaptiveMethod, thresholdType, blockSize, C)
	}, true)
	engine.RegisterMethod("images.applyBinarization", "二值化", func(img *image.NRGBA, threshold int) *image.Gray { return images.ApplyBinarization(img, threshold) }, true)

	return nil
}
