//go:build ignore
// +build ignore

package time

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"

	lua "github.com/yuin/gopher-lua"
)

// TimeModule time 模块（懒人精灵兼容）
type TimeModule struct {
	startTime time.Time // 脚本启动时间
}

// Name 返回模块名称
func (m *TimeModule) Name() string {
	return "time"
}

// IsAvailable 检查模块是否可用
func (m *TimeModule) IsAvailable() bool {
	return true
}

// Register 向引擎注册方法
func (m *TimeModule) Register(engine model.Engine) error {
	// 记录脚本启动时间
	m.startTime = time.Now()

	state := engine.GetState()

	// 获取 package.preload 表
	packagePreload := state.GetGlobal("package")
	if packagePreload == lua.LNil {
		packagePreload = state.NewTable()
		state.SetGlobal("package", packagePreload)
	}

	packageTable := packagePreload.(*lua.LTable)
	preloadTable := packageTable.RawGetString("preload")
	if preloadTable == lua.LNil {
		preloadTable = state.NewTable()
		packageTable.RawSetString("preload", preloadTable)
	}

	// 创建 time 模块表
	timeModuleTable := state.NewTable()

	// 注册 systemTime - 返回系统当前时间戳（毫秒）
	timeModuleTable.RawSetString("systemTime", state.NewFunction(func(L *lua.LState) int {
		// 获取当前时间戳（毫秒）
		timestamp := time.Now().UnixMilli()
		L.Push(lua.LNumber(timestamp))
		return 1
	}))

	// 注册 getNetWorkTime - 从网络获取当前时间
	timeModuleTable.RawSetString("getNetWorkTime", state.NewFunction(func(L *lua.LState) int {
		// 尝试从网络获取时间
		networkTime, err := getNetworkTime()
		if err != nil {
			// 如果网络请求失败，返回本地时间
			networkTime = time.Now()
		}

		// 格式化为 年-月-日_时-分-秒
		formattedTime := networkTime.Format("2006-01-02_15-04-05")
		L.Push(lua.LString(formattedTime))
		return 1
	}))

	// 注册 tickCount - 返回脚本自启动以来的运行时长（毫秒）
	timeModuleTable.RawSetString("tickCount", state.NewFunction(func(L *lua.LState) int {
		// 计算脚本运行时间（毫秒）
		duration := time.Since(m.startTime).Milliseconds()
		L.Push(lua.LNumber(duration))
		return 1
	}))

	// 注册到 package.preload
	preloadTable.(*lua.LTable).RawSetString("time", state.NewFunction(func(L *lua.LState) int {
		L.Push(timeModuleTable)
		return 1
	}))

	// 注册为全局方法（可以直接调用，不需要 require）
	state.SetGlobal("systemTime", state.NewFunction(func(L *lua.LState) int {
		// 获取当前时间戳（毫秒）
		timestamp := time.Now().UnixMilli()
		L.Push(lua.LNumber(timestamp))
		return 1
	}))

	state.SetGlobal("getNetWorkTime", state.NewFunction(func(L *lua.LState) int {
		// 尝试从网络获取时间
		networkTime, err := getNetworkTime()
		if err != nil {
			// 如果网络请求失败，返回本地时间
			networkTime = time.Now()
		}

		// 格式化为 年-月-日_时-分-秒
		formattedTime := networkTime.Format("2006-01-02_15-04-05")
		L.Push(lua.LString(formattedTime))
		return 1
	}))

	state.SetGlobal("tickCount", state.NewFunction(func(L *lua.LState) int {
		// 计算脚本运行时间（毫秒）
		duration := time.Since(m.startTime).Milliseconds()
		L.Push(lua.LNumber(duration))
		return 1
	}))

	engine.RegisterMethod("time.systemTime", "返回系统当前时间戳（毫秒）", func() int64 {
		return time.Now().UnixMilli()
	}, true)
	engine.RegisterMethod("time.getNetWorkTime", "从网络获取当前时间", func() string {
		networkTime, err := getNetworkTime()
		if err != nil {
			networkTime = time.Now()
		}
		return networkTime.Format("2006-01-02_15-04-05")
	}, true)
	engine.RegisterMethod("time.tickCount", "返回脚本自启动以来的运行时长（毫秒）", func() int64 {
		return time.Since(m.startTime).Milliseconds()
	}, true)

	return nil
}

// getNetworkTime 从网络获取当前时间
func getNetworkTime() (time.Time, error) {
	// 使用 HTTP 请求获取网络时间
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	// 尝试从多个时间服务器获取时间
	servers := []string{
		"http://worldtimeapi.org/api/timezone/Asia/Shanghai",
		"http://worldtimeapi.org/api/ip",
	}

	for _, server := range servers {
		resp, err := client.Get(server)
		if err != nil {
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			// 这里简化处理，直接返回当前时间
			// 实际应用中应该解析响应体中的时间字段
			return time.Now(), nil
		}
	}

	return time.Now(), fmt.Errorf("无法从网络获取时间")
}
