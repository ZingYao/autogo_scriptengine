package media

import (
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"

	autogomedia "github.com/Dasongzi1366/AutoGo/media"
)

// MediaModule 是 go-lua-vm 迁移后的模块壳。
type MediaModule struct{}

func New() *MediaModule { return &MediaModule{} }

func (m *MediaModule) Name() string { return "media" }

func (m *MediaModule) IsAvailable() bool { return true }

func (m *MediaModule) Register(engine model.Engine) error {
	engine.RegisterMethod("media.scanFile", "扫描路径的媒体文件", autogomedia.ScanFile, true)
	engine.RegisterMethod("media.playMP3", "播放 MP3 文件", autogomedia.PlayMP3, true)
	engine.RegisterMethod("media.sendSMS", "发送短信", autogomedia.SendSMS, true)
	return nil
}

func GetModule() model.Module { return &MediaModule{} }
