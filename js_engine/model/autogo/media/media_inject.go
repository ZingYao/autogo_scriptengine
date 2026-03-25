package media

import (
	"github.com/ZingYao/autogo_scriptengine/js_engine/model"

	"github.com/Dasongzi1366/AutoGo/media"
	"github.com/dop251/goja"
)

// MediaModule media 模块
type MediaModule struct{}

// Name 返回模块名称
func (m *MediaModule) Name() string {
	return "media"
}

// IsAvailable 检查模块是否可用
func (m *MediaModule) IsAvailable() bool {
	return true
}

// Register 向引擎注册方法
func (m *MediaModule) Register(engine model.Engine) error {
	vm := engine.GetVM()

	mediaObj := vm.NewObject()
	vm.Set("media", mediaObj)

	mediaObj.Set("scanFile", func(call goja.FunctionCall) goja.Value {
		path := call.Argument(0).String()
		media.ScanFile(path)
		return goja.Undefined()
	})

	mediaObj.Set("playMP3", func(call goja.FunctionCall) goja.Value {
		path := call.Argument(0).String()
		err := media.PlayMP3(path)
		if err != nil {
			return vm.ToValue(err.Error())
		}
		return goja.Null()
	})

	mediaObj.Set("sendSMS", func(call goja.FunctionCall) goja.Value {
		number := call.Argument(0).String()
		message := call.Argument(1).String()
		media.SendSMS(number, message)
		return goja.Undefined()
	})

	engine.RegisterMethod("media.scanFile", "扫描路径path的媒体文件", func(path string) { media.ScanFile(path) }, true)
	engine.RegisterMethod("media.playMP3", "播放MP3文件", func(path string) error { return media.PlayMP3(path) }, true)
	engine.RegisterMethod("media.sendSMS", "发送短信", func(number, message string) { media.SendSMS(number, message) }, true)

	return nil
}
