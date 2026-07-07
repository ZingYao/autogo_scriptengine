package system

import (
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"

	autogosystem "github.com/Dasongzi1366/AutoGo/system"
)

// SystemModule 是 go-lua-vm 迁移后的模块壳。
type SystemModule struct{}

func New() *SystemModule { return &SystemModule{} }

func (m *SystemModule) Name() string { return "system" }

func (m *SystemModule) IsAvailable() bool { return true }

func (m *SystemModule) Register(engine model.Engine) error {
	engine.RegisterMethod("system.getPid", "获取进程 ID", autogosystem.GetPid, true)
	engine.RegisterMethod("system.getMemoryUsage", "获取内存使用", autogosystem.GetMemoryUsage, true)
	engine.RegisterMethod("system.getCpuUsage", "获取 CPU 使用率", autogosystem.GetCpuUsage, true)
	engine.RegisterMethod("system.restartSelf", "重启自身", autogosystem.RestartSelf, true)
	engine.RegisterMethod("system.setBootStart", "设置开机自启", autogosystem.SetBootStart, true)
	return nil
}

func GetModule() model.Module { return &SystemModule{} }
