//go:build !android || !cgo

package opencv

import (
	"fmt"

	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"
)

// OpencvModule 是非 Android 环境下的 OpenCV 模块占位。
type OpencvModule struct{}

func New() *OpencvModule { return &OpencvModule{} }

func (m *OpencvModule) Name() string { return "opencv" }

func (m *OpencvModule) IsAvailable() bool { return true }

func (m *OpencvModule) Register(engine model.Engine) error {
	engine.RegisterMethod("opencv.newPoint2f", "创建 Point2f", func(x, y float32) (interface{}, error) {
		return nil, fmt.Errorf("opencv.newPoint2f requires Android AutoGo OpenCV runtime")
	}, true)
	return nil
}

func GetModule() model.Module { return &OpencvModule{} }
