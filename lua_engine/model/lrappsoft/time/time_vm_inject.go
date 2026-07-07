package time

import (
	"time"

	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"
)

// TimeModule 是 go-lua-vm 迁移后的模块壳。
type TimeModule struct {
	startTime time.Time
}

func New() *TimeModule { return &TimeModule{} }

func (m *TimeModule) Name() string { return "time" }

func (m *TimeModule) IsAvailable() bool { return true }

func (m *TimeModule) Register(engine model.Engine) error {
	m.startTime = time.Now()
	engine.RegisterMethod("time.systemTime", "返回系统当前时间戳（毫秒）", func() int64 {
		return time.Now().UnixMilli()
	}, true)
	engine.RegisterMethod("time.getNetWorkTime", "返回当前时间字符串", func() string {
		return time.Now().Format("2006-01-02_15-04-05")
	}, true)
	engine.RegisterMethod("time.tickCount", "返回模块注册以来的运行时长（毫秒）", func() int64 {
		return time.Since(m.startTime).Milliseconds()
	}, true)
	return nil
}

func GetModule() model.Module { return &TimeModule{} }
