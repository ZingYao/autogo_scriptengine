package images

import (
	"app/lua_engine/model"
	"image"

	"github.com/Dasongzi1366/AutoGo/images"
	lua "github.com/yuin/gopher-lua"
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
	state := engine.GetState()

	imagesObj := state.NewTable()
	state.SetGlobal("images", imagesObj)

	imagesObj.RawSetString("pixel", state.NewFunction(func(L *lua.LState) int {
		x := L.CheckInt(1)
		y := L.CheckInt(2)
		displayId := 0
		if L.GetTop() > 2 {
			displayId = L.CheckInt(3)
		}
		result := images.Pixel(x, y, displayId)
		L.Push(lua.LString(result))
		return 1
	}))

	imagesObj.RawSetString("setCallback", state.NewFunction(func(L *lua.LState) int {
		L.CheckFunction(1)
		images.SetCallback(func(img *image.NRGBA, displayId int) {
			ud := L.NewUserData()
			ud.Value = img
			L.Push(ud)
			L.Push(lua.LNumber(displayId))
			L.Call(2, 0)
		})
		return 0
	}))

	imagesObj.RawSetString("captureScreen", state.NewFunction(func(L *lua.LState) int {
		x1 := L.CheckInt(1)
		y1 := L.CheckInt(2)
		x2 := L.CheckInt(3)
		y2 := L.CheckInt(4)
		displayId := 0
		if L.GetTop() > 4 {
			displayId = L.CheckInt(5)
		}
		result := images.CaptureScreen(x1, y1, x2, y2, displayId)
		if result != nil {
			ud := L.NewUserData()
			ud.Value = result
			L.Push(ud)
		} else {
			L.Push(lua.LNil)
		}
		return 1
	}))

	imagesObj.RawSetString("cmpColor", state.NewFunction(func(L *lua.LState) int {
		x := L.CheckInt(1)
		y := L.CheckInt(2)
		colorStr := L.CheckString(3)
		sim := float32(L.CheckNumber(4))
		displayId := 0
		if L.GetTop() > 4 {
			displayId = L.CheckInt(5)
		}
		result := images.CmpColor(x, y, colorStr, sim, displayId)
		L.Push(lua.LBool(result))
		return 1
	}))

	imagesObj.RawSetString("findColor", state.NewFunction(func(L *lua.LState) int {
		x1 := L.CheckInt(1)
		y1 := L.CheckInt(2)
		x2 := L.CheckInt(3)
		y2 := L.CheckInt(4)
		colorStr := L.CheckString(5)
		sim := float32(L.CheckNumber(6))
		dir := L.CheckInt(7)
		displayId := 0
		if L.GetTop() > 7 {
			displayId = L.CheckInt(8)
		}
		x, y := images.FindColor(x1, y1, x2, y2, colorStr, sim, dir, displayId)
		result := L.NewTable()
		L.SetField(result, "x", lua.LNumber(x))
		L.SetField(result, "y", lua.LNumber(y))
		L.Push(result)
		return 1
	}))

	imagesObj.RawSetString("getColorCountInRegion", state.NewFunction(func(L *lua.LState) int {
		x1 := L.CheckInt(1)
		y1 := L.CheckInt(2)
		x2 := L.CheckInt(3)
		y2 := L.CheckInt(4)
		colorStr := L.CheckString(5)
		sim := float32(L.CheckNumber(6))
		displayId := 0
		if L.GetTop() > 6 {
			displayId = L.CheckInt(7)
		}
		result := images.GetColorCountInRegion(x1, y1, x2, y2, colorStr, sim, displayId)
		L.Push(lua.LNumber(result))
		return 1
	}))

	imagesObj.RawSetString("detectsMultiColors", state.NewFunction(func(L *lua.LState) int {
		colors := L.CheckString(1)
		sim := float32(L.CheckNumber(2))
		displayId := 0
		if L.GetTop() > 2 {
			displayId = L.CheckInt(3)
		}
		result := images.DetectsMultiColors(colors, sim, displayId)
		L.Push(lua.LBool(result))
		return 1
	}))

	imagesObj.RawSetString("findMultiColors", state.NewFunction(func(L *lua.LState) int {
		x1 := L.CheckInt(1)
		y1 := L.CheckInt(2)
		x2 := L.CheckInt(3)
		y2 := L.CheckInt(4)
		colors := L.CheckString(5)
		sim := float32(L.CheckNumber(6))
		dir := L.CheckInt(7)
		displayId := 0
		if L.GetTop() > 7 {
			displayId = L.CheckInt(8)
		}
		x, y := images.FindMultiColors(x1, y1, x2, y2, colors, sim, dir, displayId)
		result := L.NewTable()
		L.SetField(result, "x", lua.LNumber(x))
		L.SetField(result, "y", lua.LNumber(y))
		L.Push(result)
		return 1
	}))

	imagesObj.RawSetString("readFromPath", state.NewFunction(func(L *lua.LState) int {
		path := L.CheckString(1)
		result := images.ReadFromPath(path)
		if result != nil {
			ud := L.NewUserData()
			ud.Value = result
			L.Push(ud)
		} else {
			L.Push(lua.LNil)
		}
		return 1
	}))

	imagesObj.RawSetString("readFromUrl", state.NewFunction(func(L *lua.LState) int {
		url := L.CheckString(1)
		result := images.ReadFromUrl(url)
		if result != nil {
			ud := L.NewUserData()
			ud.Value = result
			L.Push(ud)
		} else {
			L.Push(lua.LNil)
		}
		return 1
	}))

	imagesObj.RawSetString("readFromBase64", state.NewFunction(func(L *lua.LState) int {
		base64Str := L.CheckString(1)
		result := images.ReadFromBase64(base64Str)
		if result != nil {
			ud := L.NewUserData()
			ud.Value = result
			L.Push(ud)
		} else {
			L.Push(lua.LNil)
		}
		return 1
	}))

	imagesObj.RawSetString("readFromBytes", state.NewFunction(func(L *lua.LState) int {
		data := L.CheckString(1)
		result := images.ReadFromBytes([]byte(data))
		if result != nil {
			ud := L.NewUserData()
			ud.Value = result
			L.Push(ud)
		} else {
			L.Push(lua.LNil)
		}
		return 1
	}))

	imagesObj.RawSetString("toNrgba", state.NewFunction(func(L *lua.LState) int {
		img := L.CheckUserData(1).Value.(image.Image)
		result := images.ToNrgba(img)
		if result != nil {
			ud := L.NewUserData()
			ud.Value = result
			L.Push(ud)
		} else {
			L.Push(lua.LNil)
		}
		return 1
	}))

	imagesObj.RawSetString("save", state.NewFunction(func(L *lua.LState) int {
		img := L.CheckUserData(1).Value.(*image.NRGBA)
		path := L.CheckString(2)
		quality := 90
		if L.GetTop() > 2 {
			quality = L.CheckInt(3)
		}
		result := images.Save(img, path, quality)
		L.Push(lua.LBool(result))
		return 1
	}))

	imagesObj.RawSetString("encodeToBase64", state.NewFunction(func(L *lua.LState) int {
		img := L.CheckUserData(1).Value.(*image.NRGBA)
		format := "png"
		if L.GetTop() > 1 {
			format = L.CheckString(2)
		}
		quality := 90
		if L.GetTop() > 2 {
			quality = L.CheckInt(3)
		}
		result := images.EncodeToBase64(img, format, quality)
		L.Push(lua.LString(result))
		return 1
	}))

	imagesObj.RawSetString("encodeToBytes", state.NewFunction(func(L *lua.LState) int {
		img := L.CheckUserData(1).Value.(*image.NRGBA)
		format := "png"
		if L.GetTop() > 1 {
			format = L.CheckString(2)
		}
		quality := 90
		if L.GetTop() > 2 {
			quality = L.CheckInt(3)
		}
		result := images.EncodeToBytes(img, format, quality)
		if result != nil {
			L.Push(lua.LString(string(result)))
		} else {
			L.Push(lua.LNil)
		}
		return 1
	}))

	imagesObj.RawSetString("clip", state.NewFunction(func(L *lua.LState) int {
		img := L.CheckUserData(1).Value.(*image.NRGBA)
		x1 := L.CheckInt(2)
		y1 := L.CheckInt(3)
		x2 := L.CheckInt(4)
		y2 := L.CheckInt(5)
		result := images.Clip(img, x1, y1, x2, y2)
		if result != nil {
			ud := L.NewUserData()
			ud.Value = result
			L.Push(ud)
		} else {
			L.Push(lua.LNil)
		}
		return 1
	}))

	imagesObj.RawSetString("resize", state.NewFunction(func(L *lua.LState) int {
		img := L.CheckUserData(1).Value.(*image.NRGBA)
		width := L.CheckInt(2)
		height := L.CheckInt(3)
		result := images.Resize(img, width, height)
		if result != nil {
			ud := L.NewUserData()
			ud.Value = result
			L.Push(ud)
		} else {
			L.Push(lua.LNil)
		}
		return 1
	}))

	imagesObj.RawSetString("rotate", state.NewFunction(func(L *lua.LState) int {
		img := L.CheckUserData(1).Value.(*image.NRGBA)
		degree := L.CheckInt(2)
		result := images.Rotate(img, degree)
		if result != nil {
			ud := L.NewUserData()
			ud.Value = result
			L.Push(ud)
		} else {
			L.Push(lua.LNil)
		}
		return 1
	}))

	imagesObj.RawSetString("grayscale", state.NewFunction(func(L *lua.LState) int {
		img := L.CheckUserData(1).Value.(*image.NRGBA)
		result := images.Grayscale(img)
		if result != nil {
			ud := L.NewUserData()
			ud.Value = result
			L.Push(ud)
		} else {
			L.Push(lua.LNil)
		}
		return 1
	}))

	imagesObj.RawSetString("applyThreshold", state.NewFunction(func(L *lua.LState) int {
		img := L.CheckUserData(1).Value.(*image.NRGBA)
		threshold := L.CheckInt(2)
		maxVal := L.CheckInt(3)
		typ := L.CheckString(4)
		result := images.ApplyThreshold(img, threshold, maxVal, typ)
		if result != nil {
			ud := L.NewUserData()
			ud.Value = result
			L.Push(ud)
		} else {
			L.Push(lua.LNil)
		}
		return 1
	}))

	imagesObj.RawSetString("applyAdaptiveThreshold", state.NewFunction(func(L *lua.LState) int {
		img := L.CheckUserData(1).Value.(*image.NRGBA)
		maxValue := float64(L.CheckNumber(2))
		adaptiveMethod := L.CheckString(3)
		thresholdType := L.CheckString(4)
		blockSize := L.CheckInt(5)
		C := float64(L.CheckNumber(6))
		result := images.ApplyAdaptiveThreshold(img, maxValue, adaptiveMethod, thresholdType, blockSize, C)
		if result != nil {
			ud := L.NewUserData()
			ud.Value = result
			L.Push(ud)
		} else {
			L.Push(lua.LNil)
		}
		return 1
	}))

	imagesObj.RawSetString("applyBinarization", state.NewFunction(func(L *lua.LState) int {
		img := L.CheckUserData(1).Value.(*image.NRGBA)
		threshold := L.CheckInt(2)
		result := images.ApplyBinarization(img, threshold)
		if result != nil {
			ud := L.NewUserData()
			ud.Value = result
			L.Push(ud)
		} else {
			L.Push(lua.LNil)
		}
		return 1
	}))

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
