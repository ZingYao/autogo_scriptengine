package images

import (
	"image"
	"strconv"

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
		fn, ok := goja.AssertFunction(call.Argument(0))
		if !ok {
			images.SetCallback(nil)
			return goja.Undefined()
		}
		images.SetCallback(func(img *image.NRGBA, displayId int) {
			_, _ = fn(nil, vm.ToValue(img), vm.ToValue(displayId))
		})
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

	imagesObj.Set("findMultiColorsAll", func(call goja.FunctionCall) goja.Value {
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
		points := images.FindMultiColorsAll(x1, y1, x2, y2, colors, sim, dir, displayId)
		arr := vm.NewArray()
		for i, point := range points {
			obj := vm.NewObject()
			obj.Set("x", point.X)
			obj.Set("y", point.Y)
			arr.Set(strconv.Itoa(i), obj)
		}
		return arr
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
		jsBytes := call.Argument(0).Export()
		var data []byte

		switch v := jsBytes.(type) {
		case []byte:
			data = v
		case []interface{}:
			data = make([]byte, len(v))
			for i, val := range v {
				if num, ok := val.(float64); ok {
					data[i] = byte(num)
				}
			}
		}

		result := images.ReadFromBytes(data)
		if result != nil {
			return vm.ToValue(result)
		}
		return goja.Null()
	})

	imagesObj.Set("toNrgba", func(call goja.FunctionCall) goja.Value {
		arg0 := call.Argument(0).Export()
		if arg0 == nil {
			return goja.Null()
		}
		img := arg0.(image.Image)
		result := images.ToNrgba(img)
		if result != nil {
			return vm.ToValue(result)
		}
		return goja.Null()
	})

	imagesObj.Set("save", func(call goja.FunctionCall) goja.Value {
		arg0 := call.Argument(0).Export()
		if arg0 == nil {
			return goja.Null()
		}
		img := arg0.(*image.NRGBA)
		path := call.Argument(1).String()
		quality := int(call.Argument(2).ToInteger())
		result := images.Save(img, path, quality)
		return vm.ToValue(result)
	})

	imagesObj.Set("encodeToBase64", func(call goja.FunctionCall) goja.Value {
		arg0 := call.Argument(0).Export()
		if arg0 == nil {
			return goja.Null()
		}
		img := arg0.(*image.NRGBA)
		format := call.Argument(1).String()
		quality := int(call.Argument(2).ToInteger())
		result := images.EncodeToBase64(img, format, quality)
		return vm.ToValue(result)
	})

	imagesObj.Set("encodeToBytes", func(call goja.FunctionCall) goja.Value {
		arg0 := call.Argument(0).Export()
		if arg0 == nil {
			return goja.Null()
		}
		img := arg0.(*image.NRGBA)
		format := call.Argument(1).String()
		quality := int(call.Argument(2).ToInteger())
		result := images.EncodeToBytes(img, format, quality)
		if result != nil {
			return vm.ToValue(result)
		}
		return goja.Null()
	})

	imagesObj.Set("clip", func(call goja.FunctionCall) goja.Value {
		arg0 := call.Argument(0).Export()
		if arg0 == nil {
			return goja.Null()
		}
		img := arg0.(*image.NRGBA)
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
		arg0 := call.Argument(0).Export()
		if arg0 == nil {
			return goja.Null()
		}
		img := arg0.(*image.NRGBA)
		width := int(call.Argument(1).ToInteger())
		height := int(call.Argument(2).ToInteger())
		result := images.Resize(img, width, height)
		if result != nil {
			return vm.ToValue(result)
		}
		return goja.Null()
	})

	imagesObj.Set("rotate", func(call goja.FunctionCall) goja.Value {
		arg0 := call.Argument(0).Export()
		if arg0 == nil {
			return goja.Null()
		}
		img := arg0.(*image.NRGBA)
		degree := int(call.Argument(1).ToInteger())
		result := images.Rotate(img, degree)
		if result != nil {
			return vm.ToValue(result)
		}
		return goja.Null()
	})

	imagesObj.Set("grayscale", func(call goja.FunctionCall) goja.Value {
		arg0 := call.Argument(0).Export()
		if arg0 == nil {
			return goja.Null()
		}
		img := arg0.(*image.NRGBA)
		result := images.Grayscale(img)
		if result != nil {
			return vm.ToValue(result)
		}
		return goja.Null()
	})

	imagesObj.Set("applyThreshold", func(call goja.FunctionCall) goja.Value {
		arg0 := call.Argument(0).Export()
		if arg0 == nil {
			return goja.Null()
		}
		img := arg0.(*image.NRGBA)
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
		arg0 := call.Argument(0).Export()
		if arg0 == nil {
			return goja.Null()
		}
		img := arg0.(*image.NRGBA)
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
		arg0 := call.Argument(0).Export()
		if arg0 == nil {
			return goja.Null()
		}
		img := arg0.(*image.NRGBA)
		threshold := int(call.Argument(1).ToInteger())
		result := images.ApplyBinarization(img, threshold)
		if result != nil {
			return vm.ToValue(result)
		}
		return goja.Null()
	})

	engine.RegisterMethod("images.setCallback", "设置回调函数", images.SetCallback, true)
	engine.RegisterMethod("images.captureScreen", "截取屏幕", images.CaptureScreen, true)
	engine.RegisterMethod("images.pixel", "获取指定坐标的像素颜色", images.Pixel, true)
	engine.RegisterMethod("images.cmpColor", "比较颜色", images.CmpColor, true)
	engine.RegisterMethod("images.findColor", "查找颜色", images.FindColor, true)
	engine.RegisterMethod("images.getColorCountInRegion", "获取区域内指定颜色的数量", images.GetColorCountInRegion, true)
	engine.RegisterMethod("images.detectsMultiColors", "检测多点颜色", images.DetectsMultiColors, true)
	engine.RegisterMethod("images.findMultiColors", "查找多点颜色", images.FindMultiColors, true)
	engine.RegisterMethod("images.findMultiColorsAll", "查找所有多点颜色", images.FindMultiColorsAll, true)
	engine.RegisterMethod("images.readFromPath", "从路径读取图片", images.ReadFromPath, true)
	engine.RegisterMethod("images.readFromUrl", "从URL读取图片", images.ReadFromUrl, true)
	engine.RegisterMethod("images.readFromBase64", "从Base64读取图片", images.ReadFromBase64, true)
	engine.RegisterMethod("images.readFromBytes", "从字节数组读取图片", images.ReadFromBytes, true)
	engine.RegisterMethod("images.save", "保存图片", images.Save, true)
	engine.RegisterMethod("images.encodeToBase64", "编码为Base64", images.EncodeToBase64, true)
	engine.RegisterMethod("images.encodeToBytes", "编码为字节数组", images.EncodeToBytes, true)
	engine.RegisterMethod("images.toNrgba", "转换为NRGBA格式", images.ToNrgba, true)
	engine.RegisterMethod("images.clip", "裁剪图片", images.Clip, true)
	engine.RegisterMethod("images.resize", "调整图片大小", images.Resize, true)
	engine.RegisterMethod("images.rotate", "旋转图片", images.Rotate, true)
	engine.RegisterMethod("images.grayscale", "灰度化", images.Grayscale, true)
	engine.RegisterMethod("images.applyThreshold", "应用阈值", images.ApplyThreshold, true)
	engine.RegisterMethod("images.applyAdaptiveThreshold", "应用自适应阈值", images.ApplyAdaptiveThreshold, true)
	engine.RegisterMethod("images.applyBinarization", "二值化", images.ApplyBinarization, true)

	return nil
}
