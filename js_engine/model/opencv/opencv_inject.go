package opencv

import (
	"app/js_engine/model"

	"github.com/Dasongzi1366/AutoGo/opencv"
	"github.com/dop251/goja"
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
	vm := engine.GetVM()

	opencvObj := vm.NewObject()
	vm.Set("opencv", opencvObj)

	opencvObj.Set("findImage", func(call goja.FunctionCall) goja.Value {
		x1 := int(call.Argument(0).ToInteger())
		y1 := int(call.Argument(1).ToInteger())
		x2 := int(call.Argument(2).ToInteger())
		y2 := int(call.Argument(3).ToInteger())

		// 处理第5个参数，支持字符串路径或字节数组
		arg4 := call.Argument(4)
		var templateData []byte
		if str, ok := arg4.Export().(string); ok {
			templateData = []byte(str)
		} else if bytes, ok := arg4.Export().([]byte); ok {
			templateData = bytes
		} else if bytesPtr, ok := arg4.Export().(*[]byte); ok {
			templateData = *bytesPtr
		}

		isGray := call.Argument(5).ToBoolean()
		scalingFactor := float32(call.Argument(6).ToFloat())
		sim := float32(call.Argument(7).ToFloat())
		displayId := 0
		if len(call.Arguments) > 8 {
			displayId = int(call.Argument(8).ToInteger())
		}

		x, y := opencv.FindImage(x1, y1, x2, y2, &templateData, isGray, scalingFactor, sim, displayId)
		result := vm.NewObject()
		result.Set("x", x)
		result.Set("y", y)
		return result
	})

	engine.RegisterMethod("opencv.findImage", "在指定区域内查找匹配的图片模板", func(x1, y1, x2, y2 int, template *[]byte, isGray bool, scalingFactor, sim float32, displayId int) (int, int) {
		return opencv.FindImage(x1, y1, x2, y2, template, isGray, scalingFactor, sim, displayId)
	}, true)

	return nil
}
