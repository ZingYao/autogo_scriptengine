package media

import (
	"app/lua_engine/model"

	"github.com/Dasongzi1366/AutoGo/media"
	lua "github.com/yuin/gopher-lua"
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
	state := engine.GetState()

	mediaObj := state.NewTable()
	state.SetGlobal("media", mediaObj)

	mediaObj.RawSetString("scanFile", state.NewFunction(func(L *lua.LState) int {
		path := L.CheckString(1)
		media.ScanFile(path)
		return 0
	}))

	mediaObj.RawSetString("playMP3", state.NewFunction(func(L *lua.LState) int {
		path := L.CheckString(1)
		err := media.PlayMP3(path)
		if err != nil {
			L.Push(lua.LString(err.Error()))
		} else {
			L.Push(lua.LNil)
		}
		return 1
	}))

	mediaObj.RawSetString("sendSMS", state.NewFunction(func(L *lua.LState) int {
		number := L.CheckString(1)
		message := L.CheckString(2)
		media.SendSMS(number, message)
		return 0
	}))

	engine.RegisterMethod("media.scanFile", "扫描路径path的媒体文件", func(path string) { media.ScanFile(path) }, true)
	engine.RegisterMethod("media.playMP3", "播放MP3文件", func(path string) error { return media.PlayMP3(path) }, true)
	engine.RegisterMethod("media.sendSMS", "发送短信", func(number, message string) { media.SendSMS(number, message) }, true)

	return nil
}
