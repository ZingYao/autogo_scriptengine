//go:build ignore
// +build ignore

package images

import (
	"image"

	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"

	autogoimages "github.com/Dasongzi1366/AutoGo/images"
	lua "github.com/yuin/gopher-lua"
)

type ImagesModule struct{}

func (m *ImagesModule) Name() string      { return "images" }
func (m *ImagesModule) IsAvailable() bool { return true }

func pushImage(L *lua.LState, img any) {
	if img == nil {
		L.Push(lua.LNil)
		return
	}
	ud := L.NewUserData()
	ud.Value = img
	L.Push(ud)
}

func pushPoint(L *lua.LState, x, y int) {
	table := L.NewTable()
	table.RawSetString("x", lua.LNumber(x))
	table.RawSetString("y", lua.LNumber(y))
	L.Push(table)
}

func pushPoints(L *lua.LState, points []autogoimages.Point) {
	table := L.NewTable()
	for i, point := range points {
		row := L.NewTable()
		row.RawSetString("x", lua.LNumber(point.X))
		row.RawSetString("y", lua.LNumber(point.Y))
		table.RawSetInt(i+1, row)
	}
	L.Push(table)
}

func pushBytes(L *lua.LState, values []byte) {
	if values == nil {
		L.Push(lua.LNil)
		return
	}
	table := L.NewTable()
	for i, value := range values {
		table.RawSetInt(i+1, lua.LNumber(value))
	}
	L.Push(table)
}

func bytesFromLua(value lua.LValue) []byte {
	if value.Type() == lua.LTString {
		return []byte(value.String())
	}
	table, ok := value.(*lua.LTable)
	if !ok {
		return nil
	}
	data := make([]byte, table.Len())
	for i := 1; i <= table.Len(); i++ {
		item := table.RawGetInt(i)
		if item != lua.LNil {
			data[i-1] = byte(lua.LVAsNumber(item))
		}
	}
	return data
}

func nrgbaFromUserData(L *lua.LState, index int) *image.NRGBA {
	return L.CheckUserData(index).Value.(*image.NRGBA)
}

func imageFromUserData(L *lua.LState, index int) image.Image {
	return L.CheckUserData(index).Value.(image.Image)
}

func (m *ImagesModule) Register(engine model.Engine) error {
	state := engine.GetState()
	imagesObj := state.NewTable()
	state.SetGlobal("images", imagesObj)

	imagesObj.RawSetString("setCallback", state.NewFunction(func(L *lua.LState) int {
		callback := L.ToFunction(1)
		if callback == nil {
			autogoimages.SetCallback(nil)
			return 0
		}
		autogoimages.SetCallback(func(img *image.NRGBA) {
			L.Push(callback)
			pushImage(L, img)
			_ = L.PCall(1, 0, nil)
		})
		return 0
	}))
	imagesObj.RawSetString("captureScreen", state.NewFunction(func(L *lua.LState) int {
		pushImage(L, autogoimages.CaptureScreen(L.CheckInt(1), L.CheckInt(2), L.CheckInt(3), L.CheckInt(4)))
		return 1
	}))
	imagesObj.RawSetString("pixel", state.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LString(autogoimages.Pixel(L.CheckInt(1), L.CheckInt(2))))
		return 1
	}))
	imagesObj.RawSetString("cmpColor", state.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LBool(autogoimages.CmpColor(L.CheckInt(1), L.CheckInt(2), L.CheckString(3), float32(L.CheckNumber(4)))))
		return 1
	}))
	imagesObj.RawSetString("findColor", state.NewFunction(func(L *lua.LState) int {
		x, y := autogoimages.FindColor(L.CheckInt(1), L.CheckInt(2), L.CheckInt(3), L.CheckInt(4), L.CheckString(5), float32(L.CheckNumber(6)), L.CheckInt(7))
		pushPoint(L, x, y)
		return 1
	}))
	imagesObj.RawSetString("getColorCountInRegion", state.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LNumber(autogoimages.GetColorCountInRegion(L.CheckInt(1), L.CheckInt(2), L.CheckInt(3), L.CheckInt(4), L.CheckString(5), float32(L.CheckNumber(6)))))
		return 1
	}))
	imagesObj.RawSetString("detectsMultiColors", state.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LBool(autogoimages.DetectsMultiColors(L.CheckString(1), float32(L.CheckNumber(2)))))
		return 1
	}))
	imagesObj.RawSetString("findMultiColors", state.NewFunction(func(L *lua.LState) int {
		x, y := autogoimages.FindMultiColors(L.CheckInt(1), L.CheckInt(2), L.CheckInt(3), L.CheckInt(4), L.CheckString(5), float32(L.CheckNumber(6)), L.CheckInt(7))
		pushPoint(L, x, y)
		return 1
	}))
	imagesObj.RawSetString("findMultiColorsAll", state.NewFunction(func(L *lua.LState) int {
		pushPoints(L, autogoimages.FindMultiColorsAll(L.CheckInt(1), L.CheckInt(2), L.CheckInt(3), L.CheckInt(4), L.CheckString(5), float32(L.CheckNumber(6)), L.CheckInt(7)))
		return 1
	}))
	imagesObj.RawSetString("readFromPath", state.NewFunction(func(L *lua.LState) int { pushImage(L, autogoimages.ReadFromPath(L.CheckString(1))); return 1 }))
	imagesObj.RawSetString("readFromUrl", state.NewFunction(func(L *lua.LState) int { pushImage(L, autogoimages.ReadFromUrl(L.CheckString(1))); return 1 }))
	imagesObj.RawSetString("readFromBase64", state.NewFunction(func(L *lua.LState) int { pushImage(L, autogoimages.ReadFromBase64(L.CheckString(1))); return 1 }))
	imagesObj.RawSetString("readFromBytes", state.NewFunction(func(L *lua.LState) int {
		pushImage(L, autogoimages.ReadFromBytes(bytesFromLua(L.CheckAny(1))))
		return 1
	}))
	imagesObj.RawSetString("save", state.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LBool(autogoimages.Save(nrgbaFromUserData(L, 1), L.CheckString(2), L.CheckInt(3))))
		return 1
	}))
	imagesObj.RawSetString("encodeToBase64", state.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LString(autogoimages.EncodeToBase64(nrgbaFromUserData(L, 1), L.CheckString(2), L.CheckInt(3))))
		return 1
	}))
	imagesObj.RawSetString("encodeToBytes", state.NewFunction(func(L *lua.LState) int {
		pushBytes(L, autogoimages.EncodeToBytes(nrgbaFromUserData(L, 1), L.CheckString(2), L.CheckInt(3)))
		return 1
	}))
	imagesObj.RawSetString("toNrgba", state.NewFunction(func(L *lua.LState) int { pushImage(L, autogoimages.ToNrgba(imageFromUserData(L, 1))); return 1 }))
	imagesObj.RawSetString("clip", state.NewFunction(func(L *lua.LState) int {
		pushImage(L, autogoimages.Clip(nrgbaFromUserData(L, 1), L.CheckInt(2), L.CheckInt(3), L.CheckInt(4), L.CheckInt(5)))
		return 1
	}))
	imagesObj.RawSetString("resize", state.NewFunction(func(L *lua.LState) int {
		pushImage(L, autogoimages.Resize(nrgbaFromUserData(L, 1), L.CheckInt(2), L.CheckInt(3)))
		return 1
	}))
	imagesObj.RawSetString("rotate", state.NewFunction(func(L *lua.LState) int {
		pushImage(L, autogoimages.Rotate(nrgbaFromUserData(L, 1), L.CheckInt(2)))
		return 1
	}))
	imagesObj.RawSetString("grayscale", state.NewFunction(func(L *lua.LState) int { pushImage(L, autogoimages.Grayscale(nrgbaFromUserData(L, 1))); return 1 }))
	imagesObj.RawSetString("applyThreshold", state.NewFunction(func(L *lua.LState) int {
		pushImage(L, autogoimages.ApplyThreshold(nrgbaFromUserData(L, 1), L.CheckInt(2), L.CheckInt(3), L.CheckString(4)))
		return 1
	}))
	imagesObj.RawSetString("applyAdaptiveThreshold", state.NewFunction(func(L *lua.LState) int {
		pushImage(L, autogoimages.ApplyAdaptiveThreshold(nrgbaFromUserData(L, 1), float64(L.CheckNumber(2)), L.CheckString(3), L.CheckString(4), L.CheckInt(5), float64(L.CheckNumber(6))))
		return 1
	}))
	imagesObj.RawSetString("applyBinarization", state.NewFunction(func(L *lua.LState) int {
		pushImage(L, autogoimages.ApplyBinarization(nrgbaFromUserData(L, 1), L.CheckInt(2)))
		return 1
	}))

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
