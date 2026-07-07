//go:build ignore
// +build ignore

package images

import (
	"fmt"
	"image"
	"runtime"

	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"

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
		callback := L.ToFunction(1)

		if callback == nil {
			images.SetCallback(nil)
			return 0
		}

		images.SetCallback(func(img *image.NRGBA, displayId int) {
			if callback == nil {
				fmt.Println("截图回调已清除, 不执行回调函数, 显示ID:", displayId)
				return
			}
			if img == nil {
				fmt.Println("截图回调: 图片为 nil, 不执行回调函数, 显示ID:", displayId)
				return
			}
			L.Push(callback)
			ud := L.NewUserData()
			ud.Value = img
			L.Push(ud)
			L.Push(lua.LNumber(displayId))

			// 使用 PCall 捕获错误
			err := L.PCall(2, 0, nil)
			if err != nil {
				fmt.Println("回调函数执行错误:", err.Error())

				// 打印 Golang 调用堆栈信息
				buf := make([]byte, 4096)
				n := runtime.Stack(buf, false)
				fmt.Printf("Golang 调用堆栈:\n%s\n", buf[:n])

				// 打印 Lua 脚本堆栈信息
				fmt.Println("Lua 脚本堆栈:")
				fmt.Println(L.Where(1))
			}
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

	imagesObj.RawSetString("findMultiColorsAll", state.NewFunction(func(L *lua.LState) int {
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
		points := images.FindMultiColorsAll(x1, y1, x2, y2, colors, sim, dir, displayId)
		result := L.NewTable()
		for i, point := range points {
			pointTable := L.NewTable()
			pointTable.RawSetString("x", lua.LNumber(point.X))
			pointTable.RawSetString("y", lua.LNumber(point.Y))
			result.RawSetInt(i+1, pointTable)
		}
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
