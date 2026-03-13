package yolo

import (
	"app/lua_engine/model"
	"image"

	"github.com/Dasongzi1366/AutoGo/yolo"
	lua "github.com/yuin/gopher-lua"
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
	state := engine.GetState()

	yoloObj := state.NewTable()
	state.SetGlobal("yolo", yoloObj)

	yoloObj.RawSetString("new", state.NewFunction(func(L *lua.LState) int {
		version := L.CheckString(1)
		cpuThreadNum := L.CheckInt(2)
		paramPath := L.CheckString(3)
		binPath := L.CheckString(4)
		labels := L.CheckString(5)
		y := yolo.New(version, cpuThreadNum, paramPath, binPath, labels)
		ud := L.NewUserData()
		ud.Value = y
		L.Push(ud)
		return 1
	}))

	yoloObj.RawSetString("detect", state.NewFunction(func(L *lua.LState) int {
		y := L.CheckUserData(1).Value.(*yolo.Yolo)
		x1 := L.CheckInt(2)
		y1 := L.CheckInt(3)
		x2 := L.CheckInt(4)
		y2 := L.CheckInt(5)
		displayId := L.CheckInt(6)
		result := y.Detect(x1, y1, x2, y2, displayId)
		resultTable := L.NewTable()
		for i, item := range result {
			itemTable := L.NewTable()
			itemTable.RawSetString("x", lua.LNumber(item.X))
			itemTable.RawSetString("y", lua.LNumber(item.Y))
			itemTable.RawSetString("width", lua.LNumber(item.Width))
			itemTable.RawSetString("height", lua.LNumber(item.Height))
			itemTable.RawSetString("label", lua.LString(item.Label))
			itemTable.RawSetString("confidence", lua.LNumber(item.Score))
			resultTable.RawSetInt(i+1, itemTable)
		}
		L.Push(resultTable)
		return 1
	}))

	yoloObj.RawSetString("detectFromImage", state.NewFunction(func(L *lua.LState) int {
		y := L.CheckUserData(1).Value.(*yolo.Yolo)
		img := L.CheckUserData(2).Value.(*image.NRGBA)
		result := y.DetectFromImage(img)
		resultTable := L.NewTable()
		for i, item := range result {
			itemTable := L.NewTable()
			itemTable.RawSetString("x", lua.LNumber(item.X))
			itemTable.RawSetString("y", lua.LNumber(item.Y))
			itemTable.RawSetString("width", lua.LNumber(item.Width))
			itemTable.RawSetString("height", lua.LNumber(item.Height))
			itemTable.RawSetString("label", lua.LString(item.Label))
			itemTable.RawSetString("confidence", lua.LNumber(item.Score))
			resultTable.RawSetInt(i+1, itemTable)
		}
		L.Push(resultTable)
		return 1
	}))

	yoloObj.RawSetString("detectFromBase64", state.NewFunction(func(L *lua.LState) int {
		y := L.CheckUserData(1).Value.(*yolo.Yolo)
		b64 := L.CheckString(2)
		colorStr := L.CheckString(3)
		result := y.DetectFromBase64(b64, colorStr)
		resultTable := L.NewTable()
		for i, item := range result {
			itemTable := L.NewTable()
			itemTable.RawSetString("x", lua.LNumber(item.X))
			itemTable.RawSetString("y", lua.LNumber(item.Y))
			itemTable.RawSetString("width", lua.LNumber(item.Width))
			itemTable.RawSetString("height", lua.LNumber(item.Height))
			itemTable.RawSetString("label", lua.LString(item.Label))
			itemTable.RawSetString("confidence", lua.LNumber(item.Score))
			resultTable.RawSetInt(i+1, itemTable)
		}
		L.Push(resultTable)
		return 1
	}))

	yoloObj.RawSetString("detectFromPath", state.NewFunction(func(L *lua.LState) int {
		y := L.CheckUserData(1).Value.(*yolo.Yolo)
		path := L.CheckString(2)
		colorStr := L.CheckString(3)
		result := y.DetectFromPath(path, colorStr)
		resultTable := L.NewTable()
		for i, item := range result {
			itemTable := L.NewTable()
			itemTable.RawSetString("x", lua.LNumber(item.X))
			itemTable.RawSetString("y", lua.LNumber(item.Y))
			itemTable.RawSetString("width", lua.LNumber(item.Width))
			itemTable.RawSetString("height", lua.LNumber(item.Height))
			itemTable.RawSetString("label", lua.LString(item.Label))
			itemTable.RawSetString("confidence", lua.LNumber(item.Score))
			resultTable.RawSetInt(i+1, itemTable)
		}
		L.Push(resultTable)
		return 1
	}))

	yoloObj.RawSetString("close", state.NewFunction(func(L *lua.LState) int {
		y := L.CheckUserData(1).Value.(*yolo.Yolo)
		y.Close()
		return 0
	}))

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
