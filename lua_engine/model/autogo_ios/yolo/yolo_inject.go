package yolo

import (
	"image"

	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"

	autogoyolo "github.com/Dasongzi1366/AutoGo/yolo"
	lua "github.com/yuin/gopher-lua"
)

type YoloModule struct{}

func (m *YoloModule) Name() string      { return "yolo" }
func (m *YoloModule) IsAvailable() bool { return true }

func resultsToLua(L *lua.LState, results []autogoyolo.Result) lua.LValue {
	table := L.NewTable()
	for i, item := range results {
		row := L.NewTable()
		row.RawSetString("x", lua.LNumber(item.X))
		row.RawSetString("y", lua.LNumber(item.Y))
		row.RawSetString("width", lua.LNumber(item.Width))
		row.RawSetString("height", lua.LNumber(item.Height))
		row.RawSetString("label", lua.LString(item.Label))
		row.RawSetString("score", lua.LNumber(item.Score))
		row.RawSetString("centerX", lua.LNumber(item.CenterX))
		row.RawSetString("centerY", lua.LNumber(item.CenterY))
		table.RawSetInt(i+1, row)
	}
	return table
}

func wrapYolo(L *lua.LState, y *autogoyolo.Yolo) lua.LValue {
	obj := L.NewTable()
	obj.RawSetString("setImage", L.NewFunction(func(L *lua.LState) int {
		y.SetImage(L.CheckUserData(1).Value.(*image.NRGBA))
		return 0
	}))
	obj.RawSetString("detect", L.NewFunction(func(L *lua.LState) int {
		L.Push(resultsToLua(L, y.Detect(L.CheckInt(1), L.CheckInt(2), L.CheckInt(3), L.CheckInt(4))))
		return 1
	}))
	obj.RawSetString("detectFromImage", L.NewFunction(func(L *lua.LState) int {
		L.Push(resultsToLua(L, y.DetectFromImage(L.CheckUserData(1).Value.(*image.NRGBA))))
		return 1
	}))
	obj.RawSetString("detectFromBase64", L.NewFunction(func(L *lua.LState) int {
		L.Push(resultsToLua(L, y.DetectFromBase64(L.CheckString(1))))
		return 1
	}))
	obj.RawSetString("detectFromPath", L.NewFunction(func(L *lua.LState) int {
		L.Push(resultsToLua(L, y.DetectFromPath(L.CheckString(1))))
		return 1
	}))
	obj.RawSetString("close", L.NewFunction(func(L *lua.LState) int {
		y.Close()
		return 0
	}))
	return obj
}

func (m *YoloModule) Register(engine model.Engine) error {
	state := engine.GetState()
	yoloObj := state.NewTable()
	state.SetGlobal("yolo", yoloObj)
	yoloObj.RawSetString("new", state.NewFunction(func(L *lua.LState) int {
		y := autogoyolo.New(L.CheckString(1), L.CheckInt(2), L.CheckString(3), L.CheckString(4), L.CheckString(5))
		if y == nil {
			L.Push(lua.LNil)
			return 1
		}
		L.Push(wrapYolo(L, y))
		return 1
	}))

	engine.RegisterMethod("yolo.new", "创建 YOLO 实例", autogoyolo.New, true)
	engine.RegisterMethod("yolo.setImage", "设置下次 detect 的原始图像", func(y *autogoyolo.Yolo, img *image.NRGBA) { y.SetImage(img) }, true)
	engine.RegisterMethod("yolo.detect", "检测屏幕区域中的对象", func(y *autogoyolo.Yolo, x1, y1, x2, y2 int) []autogoyolo.Result {
		return y.Detect(x1, y1, x2, y2)
	}, true)
	engine.RegisterMethod("yolo.detectFromImage", "检测图像对象中的对象", func(y *autogoyolo.Yolo, img *image.NRGBA) []autogoyolo.Result {
		return y.DetectFromImage(img)
	}, true)
	engine.RegisterMethod("yolo.detectFromBase64", "检测 Base64 图像中的对象", func(y *autogoyolo.Yolo, b64 string) []autogoyolo.Result {
		return y.DetectFromBase64(b64)
	}, true)
	engine.RegisterMethod("yolo.detectFromPath", "检测文件图像中的对象", func(y *autogoyolo.Yolo, path string) []autogoyolo.Result {
		return y.DetectFromPath(path)
	}, true)
	engine.RegisterMethod("yolo.close", "关闭 YOLO 实例", func(y *autogoyolo.Yolo) { y.Close() }, true)
	return nil
}
