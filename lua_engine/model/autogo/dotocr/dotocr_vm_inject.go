package dotocr

import (
	"fmt"

	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"
)

// DotocrModule 在远程 AutoGo 缺少 dotocr 包时保留同名入口并返回明确错误。
type DotocrModule struct{}

func New() *DotocrModule { return &DotocrModule{} }

func (m *DotocrModule) Name() string { return "dotocr" }

func (m *DotocrModule) IsAvailable() bool { return false }

func (m *DotocrModule) Register(engine model.Engine) error {
	engine.RegisterMethod("dotocr.setDict", "AutoGo/dotocr remote package unavailable", unavailable, true)
	engine.RegisterMethod("dotocr.ocr", "AutoGo/dotocr remote package unavailable", unavailable, true)
	engine.RegisterMethod("dotocr.ocrFromImage", "AutoGo/dotocr remote package unavailable", unavailable, true)
	engine.RegisterMethod("dotocr.ocrFromBase64", "AutoGo/dotocr remote package unavailable", unavailable, true)
	engine.RegisterMethod("dotocr.ocrFromPath", "AutoGo/dotocr remote package unavailable", unavailable, true)
	engine.RegisterMethod("dotocr.findStr", "AutoGo/dotocr remote package unavailable", unavailable, true)
	engine.RegisterMethod("dotocr.findStrFromImage", "AutoGo/dotocr remote package unavailable", unavailable, true)
	engine.RegisterMethod("dotocr.findStrFromBase64", "AutoGo/dotocr remote package unavailable", unavailable, true)
	engine.RegisterMethod("dotocr.findStrFromPath", "AutoGo/dotocr remote package unavailable", unavailable, true)
	return nil
}

func unavailable(args ...interface{}) (interface{}, error) {
	return nil, fmt.Errorf("AutoGo/%s package is unavailable in the remote github.com/Dasongzi1366/AutoGo module", "dotocr")
}

func GetModule() model.Module { return &DotocrModule{} }
