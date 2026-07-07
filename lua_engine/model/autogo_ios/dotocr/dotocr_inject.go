//go:build ignore
// +build ignore

package dotocr

import (
	"image"

	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"

	autogodotocr "github.com/Dasongzi1366/AutoGo/dotocr"
	lua "github.com/yuin/gopher-lua"
)

type DotocrModule struct{}

func (m *DotocrModule) Name() string      { return "dotocr" }
func (m *DotocrModule) IsAvailable() bool { return true }

func pushPoint(L *lua.LState, x, y int) {
	table := L.NewTable()
	table.RawSetString("x", lua.LNumber(x))
	table.RawSetString("y", lua.LNumber(y))
	L.Push(table)
}

func (m *DotocrModule) Register(engine model.Engine) error {
	state := engine.GetState()
	dotocrObj := state.NewTable()
	state.SetGlobal("dotocr", dotocrObj)

	dotocrObj.RawSetString("setDict", state.NewFunction(func(L *lua.LState) int {
		autogodotocr.SetDict(L.OptString(1, ""), L.CheckString(2))
		return 0
	}))
	dotocrObj.RawSetString("ocr", state.NewFunction(func(L *lua.LState) int {
		result := autogodotocr.Ocr(L.CheckInt(1), L.CheckInt(2), L.CheckInt(3), L.CheckInt(4), L.CheckString(5), L.CheckInt(6), L.CheckInt(7), float32(L.CheckNumber(8)), L.CheckInt(9), L.OptString(10, ""))
		L.Push(lua.LString(result))
		return 1
	}))
	dotocrObj.RawSetString("ocrFromImage", state.NewFunction(func(L *lua.LState) int {
		img := L.CheckUserData(1).Value.(*image.NRGBA)
		result := autogodotocr.OcrFromImage(img, L.CheckString(2), L.CheckInt(3), L.CheckInt(4), float32(L.CheckNumber(5)), L.CheckInt(6), L.OptString(7, ""))
		L.Push(lua.LString(result))
		return 1
	}))
	dotocrObj.RawSetString("ocrFromBase64", state.NewFunction(func(L *lua.LState) int {
		result := autogodotocr.OcrFromBase64(L.CheckString(1), L.CheckString(2), L.CheckInt(3), L.CheckInt(4), float32(L.CheckNumber(5)), L.CheckInt(6), L.OptString(7, ""))
		L.Push(lua.LString(result))
		return 1
	}))
	dotocrObj.RawSetString("ocrFromPath", state.NewFunction(func(L *lua.LState) int {
		result := autogodotocr.OcrFromPath(L.CheckString(1), L.CheckString(2), L.CheckInt(3), L.CheckInt(4), float32(L.CheckNumber(5)), L.CheckInt(6), L.OptString(7, ""))
		L.Push(lua.LString(result))
		return 1
	}))
	dotocrObj.RawSetString("findStr", state.NewFunction(func(L *lua.LState) int {
		x, y := autogodotocr.FindStr(L.CheckInt(1), L.CheckInt(2), L.CheckInt(3), L.CheckInt(4), L.CheckString(5), L.CheckString(6), L.CheckInt(7), L.CheckInt(8), float32(L.CheckNumber(9)), L.OptString(10, ""))
		pushPoint(L, x, y)
		return 1
	}))
	dotocrObj.RawSetString("findStrFromImage", state.NewFunction(func(L *lua.LState) int {
		img := L.CheckUserData(1).Value.(*image.NRGBA)
		x, y := autogodotocr.FindStrFromImage(img, L.CheckString(2), L.CheckString(3), L.CheckInt(4), L.CheckInt(5), float32(L.CheckNumber(6)), L.OptString(7, ""))
		pushPoint(L, x, y)
		return 1
	}))
	dotocrObj.RawSetString("findStrFromBase64", state.NewFunction(func(L *lua.LState) int {
		x, y := autogodotocr.FindStrFromBase64(L.CheckString(1), L.CheckString(2), L.CheckString(3), L.CheckInt(4), L.CheckInt(5), float32(L.CheckNumber(6)), L.OptString(7, ""))
		pushPoint(L, x, y)
		return 1
	}))
	dotocrObj.RawSetString("findStrFromPath", state.NewFunction(func(L *lua.LState) int {
		x, y := autogodotocr.FindStrFromPath(L.CheckString(1), L.CheckString(2), L.CheckString(3), L.CheckInt(4), L.CheckInt(5), float32(L.CheckNumber(6)), L.OptString(7, ""))
		pushPoint(L, x, y)
		return 1
	}))

	engine.RegisterMethod("dotocr.setDict", "设置字库", autogodotocr.SetDict, true)
	engine.RegisterMethod("dotocr.ocr", "从屏幕区域进行点阵 OCR", autogodotocr.Ocr, true)
	engine.RegisterMethod("dotocr.ocrFromImage", "从图像对象进行点阵 OCR", autogodotocr.OcrFromImage, true)
	engine.RegisterMethod("dotocr.ocrFromBase64", "从 Base64 图像进行点阵 OCR", autogodotocr.OcrFromBase64, true)
	engine.RegisterMethod("dotocr.ocrFromPath", "从文件图像进行点阵 OCR", autogodotocr.OcrFromPath, true)
	engine.RegisterMethod("dotocr.findStr", "从屏幕区域查找字符串位置", autogodotocr.FindStr, true)
	engine.RegisterMethod("dotocr.findStrFromImage", "从图像对象查找字符串位置", autogodotocr.FindStrFromImage, true)
	engine.RegisterMethod("dotocr.findStrFromBase64", "从 Base64 图像查找字符串位置", autogodotocr.FindStrFromBase64, true)
	engine.RegisterMethod("dotocr.findStrFromPath", "从文件图像查找字符串位置", autogodotocr.FindStrFromPath, true)
	return nil
}
