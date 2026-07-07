//go:build ignore
// +build ignore

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
		if p == nil {
			L.Push(lua.LNil)
			return 1
		}
		L.Push(wrapPpocr(L, p))
		return 1
	}))

	engine.RegisterMethod("ppocr.new", "创建Ppocr对象", ppocr.New, true)
	engine.RegisterMethod("ppocr.ocr", "识别屏幕文字", func(p *ppocr.Ppocr, x1, y1, x2, y2 int, colorStr string, displayId int) []ppocr.Result {
		return p.Ocr(x1, y1, x2, y2, colorStr, displayId)
	}, true)
	engine.RegisterMethod("ppocr.ocrFromImage", "识别图片文字", func(p *ppocr.Ppocr, img *image.NRGBA, colorStr string) []ppocr.Result {
		return p.OcrFromImage(img, colorStr)
	}, true)
	engine.RegisterMethod("ppocr.ocrFromBase64", "识别Base64图片文字", func(p *ppocr.Ppocr, b64, colorStr string) []ppocr.Result {
		return p.OcrFromBase64(b64, colorStr)
	}, true)
	engine.RegisterMethod("ppocr.ocrFromPath", "识别文件图片文字", func(p *ppocr.Ppocr, path, colorStr string) []ppocr.Result {
		return p.OcrFromPath(path, colorStr)
	}, true)
	engine.RegisterMethod("ppocr.close", "关闭Ppocr对象", func(p *ppocr.Ppocr) {
		p.Close()
	}, true)

	return nil
}

func wrapPpocr(L *lua.LState, p *ppocr.Ppocr) lua.LValue {
	obj := L.NewTable()
	obj.RawSetString("ocr", L.NewFunction(func(L *lua.LState) int {
		x1 := L.CheckInt(1)
		y1 := L.CheckInt(2)
		x2 := L.CheckInt(3)
		y2 := L.CheckInt(4)
		colorStr := L.CheckString(5)
		displayId := L.OptInt(6, 0)
		L.Push(ppocrResultsToLua(L, p.Ocr(x1, y1, x2, y2, colorStr, displayId)))
		return 1
	}))
	obj.RawSetString("ocrFromImage", L.NewFunction(func(L *lua.LState) int {
		img := L.CheckUserData(1).Value.(*image.NRGBA)
		colorStr := L.CheckString(2)
		L.Push(ppocrResultsToLua(L, p.OcrFromImage(img, colorStr)))
		return 1
	}))
	obj.RawSetString("ocrFromBase64", L.NewFunction(func(L *lua.LState) int {
		b64 := L.CheckString(1)
		colorStr := L.CheckString(2)
		L.Push(ppocrResultsToLua(L, p.OcrFromBase64(b64, colorStr)))
		return 1
	}))
	obj.RawSetString("ocrFromPath", L.NewFunction(func(L *lua.LState) int {
		path := L.CheckString(1)
		colorStr := L.CheckString(2)
		L.Push(ppocrResultsToLua(L, p.OcrFromPath(path, colorStr)))
		return 1
	}))
	obj.RawSetString("close", L.NewFunction(func(L *lua.LState) int {
		p.Close()
		return 0
	}))
	return obj
}

func ppocrResultsToLua(L *lua.LState, result []ppocr.Result) lua.LValue {
	resultTable := L.NewTable()
	for i, item := range result {
		itemTable := L.NewTable()
		itemTable.RawSetString("text", lua.LString(item.Label))
		itemTable.RawSetString("confidence", lua.LNumber(item.Score))
		itemTable.RawSetString("x", lua.LNumber(item.X))
		itemTable.RawSetString("y", lua.LNumber(item.Y))
		itemTable.RawSetString("w", lua.LNumber(item.Width))
		itemTable.RawSetString("h", lua.LNumber(item.Height))
		resultTable.RawSetInt(i+1, itemTable)
	}
	return resultTable
}
