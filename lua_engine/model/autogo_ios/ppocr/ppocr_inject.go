package ppocr

import (
	"image"

	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"

	autogoppocr "github.com/Dasongzi1366/AutoGo/ppocr"
	lua "github.com/yuin/gopher-lua"
)

type PpocrModule struct{}

func (m *PpocrModule) Name() string      { return "ppocr" }
func (m *PpocrModule) IsAvailable() bool { return true }

func resultsToLua(L *lua.LState, results []autogoppocr.Result) lua.LValue {
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

func wrapPpocr(L *lua.LState, p *autogoppocr.Ppocr) lua.LValue {
	obj := L.NewTable()
	obj.RawSetString("ocr", L.NewFunction(func(L *lua.LState) int {
		L.Push(resultsToLua(L, p.Ocr(L.CheckInt(1), L.CheckInt(2), L.CheckInt(3), L.CheckInt(4), L.CheckString(5))))
		return 1
	}))
	obj.RawSetString("ocrFromImage", L.NewFunction(func(L *lua.LState) int {
		img := L.CheckUserData(1).Value.(*image.NRGBA)
		L.Push(resultsToLua(L, p.OcrFromImage(img, L.CheckString(2))))
		return 1
	}))
	obj.RawSetString("ocrFromBase64", L.NewFunction(func(L *lua.LState) int {
		L.Push(resultsToLua(L, p.OcrFromBase64(L.CheckString(1), L.CheckString(2))))
		return 1
	}))
	obj.RawSetString("ocrFromPath", L.NewFunction(func(L *lua.LState) int {
		L.Push(resultsToLua(L, p.OcrFromPath(L.CheckString(1), L.CheckString(2))))
		return 1
	}))
	obj.RawSetString("close", L.NewFunction(func(L *lua.LState) int {
		p.Close()
		return 0
	}))
	return obj
}

func (m *PpocrModule) Register(engine model.Engine) error {
	state := engine.GetState()
	ppocrObj := state.NewTable()
	state.SetGlobal("ppocr", ppocrObj)
	ppocrObj.RawSetString("new", state.NewFunction(func(L *lua.LState) int {
		p := autogoppocr.New(L.CheckString(1))
		if p == nil {
			L.Push(lua.LNil)
			return 1
		}
		L.Push(wrapPpocr(L, p))
		return 1
	}))

	engine.RegisterMethod("ppocr.new", "创建 PPOCR 实例", autogoppocr.New, true)
	engine.RegisterMethod("ppocr.ocr", "从屏幕区域识别文字", func(p *autogoppocr.Ppocr, x1, y1, x2, y2 int, colorStr string) []autogoppocr.Result {
		return p.Ocr(x1, y1, x2, y2, colorStr)
	}, true)
	engine.RegisterMethod("ppocr.ocrFromImage", "从图像对象识别文字", func(p *autogoppocr.Ppocr, img *image.NRGBA, colorStr string) []autogoppocr.Result {
		return p.OcrFromImage(img, colorStr)
	}, true)
	engine.RegisterMethod("ppocr.ocrFromBase64", "从 Base64 图像识别文字", func(p *autogoppocr.Ppocr, b64, colorStr string) []autogoppocr.Result {
		return p.OcrFromBase64(b64, colorStr)
	}, true)
	engine.RegisterMethod("ppocr.ocrFromPath", "从文件图像识别文字", func(p *autogoppocr.Ppocr, path, colorStr string) []autogoppocr.Result {
		return p.OcrFromPath(path, colorStr)
	}, true)
	engine.RegisterMethod("ppocr.close", "关闭 PPOCR 实例", func(p *autogoppocr.Ppocr) { p.Close() }, true)
	return nil
}
