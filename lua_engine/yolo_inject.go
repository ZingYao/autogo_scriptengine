package lua_engine

import (
	"github.com/Dasongzi1366/AutoGo/yolo"
	lua "github.com/yuin/gopher-lua"
)

func injectYoloMethods(engine *LuaEngine) {

	engine.RegisterMethod("yolo.new", "创建一个新的YOLO实例", yolo.New, true)
	engine.RegisterMethod("yolo.detect", "在指定的屏幕区域执行目标检测", (*yolo.Yolo).Detect, true)
	engine.RegisterMethod("yolo.detectFromBase64", "从Base64编码的图像进行识别", (*yolo.Yolo).DetectFromBase64, true)
	engine.RegisterMethod("yolo.detectFromPath", "从文件路径进行识别", (*yolo.Yolo).DetectFromPath, true)
	engine.RegisterMethod("yolo.close", "关闭YOLO模型实例，释放相关资源", (*yolo.Yolo).Close, true)

	registerYoloLuaFunctions(engine)
}

func registerYoloLuaFunctions(engine *LuaEngine) {
	state := engine.GetState()

	state.Register("yolo_new", func(L *lua.LState) int {
		version := L.CheckString(1)
		cpuThreadNum := L.CheckInt(2)
		paramPath := L.CheckString(3)
		binPath := L.CheckString(4)
		labels := L.CheckString(5)
		result := yolo.New(version, cpuThreadNum, paramPath, binPath, labels)
		if result == nil {
			L.Push(lua.LNil)
			return 1
		}
		ud := L.NewUserData()
		ud.Value = result
		L.Push(ud)
		return 1
	})

	state.Register("yolo_detect", func(L *lua.LState) int {
		ud := L.CheckUserData(1)
		y := ud.Value.(*yolo.Yolo)
		x1 := L.CheckInt(2)
		y1 := L.CheckInt(3)
		x2 := L.CheckInt(4)
		y2 := L.CheckInt(5)
		displayId := L.CheckInt(6)
		results := y.Detect(x1, y1, x2, y2, displayId)
		tbl := L.NewTable()
		for i, r := range results {
			item := L.NewTable()
			L.SetField(item, "X", lua.LNumber(r.X))
			L.SetField(item, "Y", lua.LNumber(r.Y))
			L.SetField(item, "宽", lua.LNumber(r.Width))
			L.SetField(item, "高", lua.LNumber(r.Height))
			L.SetField(item, "标签", lua.LString(r.Label))
			L.SetField(item, "精度", lua.LNumber(r.Score))
			L.SetTable(tbl, lua.LNumber(i+1), item)
		}
		L.Push(tbl)
		return 1
	})

	state.Register("yolo_detectFromBase64", func(L *lua.LState) int {
		ud := L.CheckUserData(1)
		y := ud.Value.(*yolo.Yolo)
		b64 := L.CheckString(2)
		colorStr := L.CheckString(3)
		results := y.DetectFromBase64(b64, colorStr)
		tbl := L.NewTable()
		for i, r := range results {
			item := L.NewTable()
			L.SetField(item, "X", lua.LNumber(r.X))
			L.SetField(item, "Y", lua.LNumber(r.Y))
			L.SetField(item, "宽", lua.LNumber(r.Width))
			L.SetField(item, "高", lua.LNumber(r.Height))
			L.SetField(item, "标签", lua.LString(r.Label))
			L.SetField(item, "精度", lua.LNumber(r.Score))
			L.SetTable(tbl, lua.LNumber(i+1), item)
		}
		L.Push(tbl)
		return 1
	})

	state.Register("yolo_detectFromPath", func(L *lua.LState) int {
		ud := L.CheckUserData(1)
		y := ud.Value.(*yolo.Yolo)
		path := L.CheckString(2)
		colorStr := L.CheckString(3)
		results := y.DetectFromPath(path, colorStr)
		tbl := L.NewTable()
		for i, r := range results {
			item := L.NewTable()
			L.SetField(item, "X", lua.LNumber(r.X))
			L.SetField(item, "Y", lua.LNumber(r.Y))
			L.SetField(item, "宽", lua.LNumber(r.Width))
			L.SetField(item, "高", lua.LNumber(r.Height))
			L.SetField(item, "标签", lua.LString(r.Label))
			L.SetField(item, "精度", lua.LNumber(r.Score))
			L.SetTable(tbl, lua.LNumber(i+1), item)
		}
		L.Push(tbl)
		return 1
	})

	state.Register("yolo_close", func(L *lua.LState) int {
		ud := L.CheckUserData(1)
		y := ud.Value.(*yolo.Yolo)
		y.Close()
		return 0
	})
}
