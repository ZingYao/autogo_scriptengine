package opencv

import (
	"app/lua_engine/model"

	"github.com/Dasongzi1366/AutoGo/opencv"
	lua "github.com/yuin/gopher-lua"
)

// OpencvModule opencv 模块
type OpencvModule struct{}

// Name 返回模块名称
func (m *OpencvModule) Name() string {
	return "opencv"
}

// IsAvailable 检查模块是否可用
func (m *OpencvModule) IsAvailable() bool {
	return true
}

// Register 向引擎注册方法
func (m *OpencvModule) Register(engine model.Engine) error {
	state := engine.GetState()

	opencvObj := state.NewTable()
	state.SetGlobal("opencv", opencvObj)

	opencvObj.RawSetString("findImage", state.NewFunction(func(L *lua.LState) int {
		x1 := L.CheckInt(1)
		y1 := L.CheckInt(2)
		x2 := L.CheckInt(3)
		y2 := L.CheckInt(4)
		templateBytes := L.CheckString(5)
		template := []byte(templateBytes)
		isGray := L.CheckBool(6)
		scalingFactor := float32(L.CheckNumber(7))
		sim := float32(L.CheckNumber(8))
		displayId := L.CheckInt(9)
		x, y := opencv.FindImage(x1, y1, x2, y2, &template, isGray, scalingFactor, sim, displayId)
		result := L.NewTable()
		result.RawSetString("x", lua.LNumber(x))
		result.RawSetString("y", lua.LNumber(y))
		L.Push(result)
		return 1
	}))

	engine.RegisterMethod("opencv.findImage", "在指定区域内查找匹配的图片模板", func(x1, y1, x2, y2 int, template *[]byte, isGray bool, scalingFactor, sim float32, displayId int) (int, int) {
		return opencv.FindImage(x1, y1, x2, y2, template, isGray, scalingFactor, sim, displayId)
	}, true)

	return nil
}
