package ppocr

import (
	"image"

	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"

	"github.com/Dasongzi1366/AutoGo/ppocr"
	lua "github.com/yuin/gopher-lua"
)

// PpocrModule ppocr 模块
type PpocrModule struct{}

// Name 返回模块名称
func (m *PpocrModule) Name() string {
	return "ppocr"
}

// IsAvailable 检查模块是否可用
func (m *PpocrModule) IsAvailable() bool {
	return true
}

// Register 向引擎注册方法
func (m *PpocrModule) Register(engine model.Engine) error {
	state := engine.GetState()

	ppocrObj := state.NewTable()
	state.SetGlobal("ppocr", ppocrObj)

	ppocrObj.RawSetString("new", state.NewFunction(func(L *lua.LState) int {
		version := L.CheckString(1)
		p := ppocr.New(version)
		ud := L.NewUserData()
		ud.Value = p
		L.Push(ud)
		return 1
	}))

	ppocrObj.RawSetString("ocr", state.NewFunction(func(L *lua.LState) int {
		p := L.CheckUserData(1).Value.(*ppocr.Ppocr)
		x1 := L.CheckInt(2)
		y1 := L.CheckInt(3)
		x2 := L.CheckInt(4)
		y2 := L.CheckInt(5)
		colorStr := L.CheckString(6)
		displayId := L.CheckInt(7)
		result := p.Ocr(x1, y1, x2, y2, colorStr, displayId)
		resultTable := L.NewTable()
		for i, item := range result {
			itemTable := L.NewTable()
			itemTable.RawSetString("text", lua.LString(item.Label))
			itemTable.RawSetString("confidence", lua.LNumber(item.Score))
			resultTable.RawSetInt(i+1, itemTable)
		}
		L.Push(resultTable)
		return 1
	}))

	ppocrObj.RawSetString("ocrFromImage", state.NewFunction(func(L *lua.LState) int {
		p := L.CheckUserData(1).Value.(*ppocr.Ppocr)
		img := L.CheckUserData(2).Value.(*image.NRGBA)
		colorStr := L.CheckString(3)
		result := p.OcrFromImage(img, colorStr)
		resultTable := L.NewTable()
		for i, item := range result {
			itemTable := L.NewTable()
			itemTable.RawSetString("text", lua.LString(item.Label))
			itemTable.RawSetString("confidence", lua.LNumber(item.Score))
			resultTable.RawSetInt(i+1, itemTable)
		}
		L.Push(resultTable)
		return 1
	}))

	ppocrObj.RawSetString("ocrFromBase64", state.NewFunction(func(L *lua.LState) int {
		p := L.CheckUserData(1).Value.(*ppocr.Ppocr)
		b64 := L.CheckString(2)
		colorStr := L.CheckString(3)
		result := p.OcrFromBase64(b64, colorStr)
		resultTable := L.NewTable()
		for i, item := range result {
			itemTable := L.NewTable()
			itemTable.RawSetString("text", lua.LString(item.Label))
			itemTable.RawSetString("confidence", lua.LNumber(item.Score))
			resultTable.RawSetInt(i+1, itemTable)
		}
		L.Push(resultTable)
		return 1
	}))

	ppocrObj.RawSetString("ocrFromPath", state.NewFunction(func(L *lua.LState) int {
		p := L.CheckUserData(1).Value.(*ppocr.Ppocr)
		path := L.CheckString(2)
		colorStr := L.CheckString(3)
		result := p.OcrFromPath(path, colorStr)
		resultTable := L.NewTable()
		for i, item := range result {
			itemTable := L.NewTable()
			itemTable.RawSetString("text", lua.LString(item.Label))
			itemTable.RawSetString("confidence", lua.LNumber(item.Score))
			resultTable.RawSetInt(i+1, itemTable)
		}
		L.Push(resultTable)
		return 1
	}))

	ppocrObj.RawSetString("close", state.NewFunction(func(L *lua.LState) int {
		p := L.CheckUserData(1).Value.(*ppocr.Ppocr)
		p.Close()
		return 0
	}))

	engine.RegisterMethod("ppocr.new", "创建Ppocr对象", ppocr.New, true)
	engine.RegisterMethod("ppocr.ocr", "识别屏幕文字", func(x1, y1, x2, y2 int, colorStr string, displayId int) []map[string]interface{} {
		return []map[string]interface{}{}
	}, true)
	engine.RegisterMethod("ppocr.ocrFromImage", "识别图片文字", func(img interface{}, colorStr string) []map[string]interface{} {
		return []map[string]interface{}{}
	}, true)
	engine.RegisterMethod("ppocr.ocrFromBase64", "识别Base64图片文字", func(b64, colorStr string) []map[string]interface{} {
		return []map[string]interface{}{}
	}, true)
	engine.RegisterMethod("ppocr.ocrFromPath", "识别文件图片文字", func(path, colorStr string) []map[string]interface{} {
		return []map[string]interface{}{}
	}, true)
	engine.RegisterMethod("ppocr.close", "关闭Ppocr对象", func(p *ppocr.Ppocr) {
		p.Close()
	}, true)

	return nil
}
