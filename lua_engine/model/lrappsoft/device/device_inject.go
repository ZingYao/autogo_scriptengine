//go:build ignore
// +build ignore

package device

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/Dasongzi1366/AutoGo/app"
	"github.com/Dasongzi1366/AutoGo/device"
	"github.com/Dasongzi1366/AutoGo/ime"
	"github.com/Dasongzi1366/AutoGo/media"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"

	lua "github.com/yuin/gopher-lua"
)

// DeviceModule device 模块（懒人精灵兼容）
type DeviceModule struct {
	screenLocked bool
}

// Name 返回模块名称
func (m *DeviceModule) Name() string {
	return "device"
}

// IsAvailable 检查模块是否可用
func (m *DeviceModule) IsAvailable() bool {
	return true
}

// Register 向引擎注册方法
func (m *DeviceModule) Register(engine model.Engine) error {
	state := engine.GetState()

	// 注册 getCpuArch - 获取 CPU 架构
	state.SetGlobal("getCpuArch", state.NewFunction(func(L *lua.LState) int {
		// 获取 CPU 架构
		// 0=x86，1=arm，2=arm64
		var arch int
		switch runtime.GOARCH {
		case "amd64", "386":
			arch = 0 // x86
		case "arm":
			arch = 1 // arm
		case "arm64":
			arch = 2 // arm64
		default:
			arch = 0 // 默认 x86
		}
		L.Push(lua.LNumber(arch))
		return 1
	}))

	// 注册 getSdPath - 获取 SD 卡路径
	state.SetGlobal("getSdPath", state.NewFunction(func(L *lua.LState) int {
		// 返回 /sdcard 路径
		// 注意：这只是一个兼容性返回值，实际使用可能存在问题
		// 危害：直接使用 /sdcard 可能导致权限问题或路径错误
		// 建议：使用 getWorkPath() 获取正确的工作目录
		L.Push(lua.LString("/sdcard"))
		return 1
	}))

	// 注册 setControlBarPosNew - 设置悬浮窗位置（空操作）
	state.SetGlobal("setControlBarPosNew", state.NewFunction(func(L *lua.LState) int {
		// AutoGo 不支持悬浮窗
		// 空操作
		return 0
	}))

	// 注册 showControlBar - 显示或隐藏悬浮按钮（空操作）
	state.SetGlobal("showControlBar", state.NewFunction(func(L *lua.LState) int {
		// AutoGo 不支持悬浮窗
		// 空操作
		return 0
	}))

	// 注册 restartScript - 重启当前脚本
	state.SetGlobal("restartScript", state.NewFunction(func(L *lua.LState) int {
		// 使用 os.exit(0) 触发重启
		// 需要在引擎配置中设置 OnExit = ExitActionRestart
		// 使用 os.exit(-1) 可以强制退出，不执行任何退出动作
		osTable := L.GetGlobal("os").(*lua.LTable)
		exitFunc := osTable.RawGetString("exit")
		if exitFunc.Type() == lua.LTFunction {
			L.Push(exitFunc)
			L.Push(lua.LNumber(0))
			L.Call(1, 0)
		}
		return 0
	}))

	// 注册 getWorkPath - 获取脚本工作目录
	state.SetGlobal("getWorkPath", state.NewFunction(func(L *lua.LState) int {
		// 获取当前工作目录
		dir, err := os.Getwd()
		if err != nil {
			L.Push(lua.LString(""))
		} else {
			L.Push(lua.LString(dir))
		}
		return 1
	}))

	// 注册 getApkVerInt - 获取 apk 壳子版本（空操作）
	state.SetGlobal("getApkVerInt", state.NewFunction(func(L *lua.LState) int {
		// AutoGo 不是 Android 应用
		// 返回 0
		L.Push(lua.LNumber(0))
		return 1
	}))

	// 注册 getScriptVersion - 获取脚本版本号（空操作）
	state.SetGlobal("getScriptVersion", state.NewFunction(func(L *lua.LState) int {
		// AutoGo 不支持脚本版本号
		// 返回 0
		L.Push(lua.LNumber(0))
		return 1
	}))

	// 注册 vibrate - 设备震动
	state.SetGlobal("vibrate", state.NewFunction(func(L *lua.LState) int {
		during := L.CheckInt(1)
		// 使用 AutoGo 的 Vibrate 方法
		device.Vibrate(during)
		return 0
	}))

	// 注册 getSubscriberId - 获取设备 IMSI（空操作）
	state.SetGlobal("getSubscriberId", state.NewFunction(func(L *lua.LState) int {
		// AutoGo 运行在非 Android 环境，无法获取 IMSI
		// 返回空字符串
		L.Push(lua.LString(""))
		return 1
	}))

	// 注册 getSimSerialNumber - 获取 SIM 卡序列号（空操作）
	state.SetGlobal("getSimSerialNumber", state.NewFunction(func(L *lua.LState) int {
		// AutoGo 运行在非 Android 环境，无法获取 SIM 序列号
		// 返回空字符串
		L.Push(lua.LString(""))
		return 1
	}))

	// 注册 getDisplayDpi - 获取设备 DPI
	state.SetGlobal("getDisplayDpi", state.NewFunction(func(L *lua.LState) int {
		// 使用 AutoGo 的 GetDisplayInfo 方法
		_, _, dpi, _ := device.GetDisplayInfo(0)
		L.Push(lua.LNumber(dpi))
		return 1
	}))

	// 注册 getBatteryLevel - 获取设备电量
	state.SetGlobal("getBatteryLevel", state.NewFunction(func(L *lua.LState) int {
		// 使用 AutoGo 的 GetBattery 方法
		battery := device.GetBattery()
		L.Push(lua.LNumber(battery))
		return 1
	}))

	// 注册 getUIConfig - 读取 UI 配置（空操作）
	state.SetGlobal("getUIConfig", state.NewFunction(func(L *lua.LState) int {
		// AutoGo 不支持 UI 配置
		// 返回 nil
		L.Push(lua.LNil)
		return 1
	}))

	// 注册 playAudio - 播放音频文件
	state.SetGlobal("playAudio", state.NewFunction(func(L *lua.LState) int {
		audioPath := L.CheckString(1)
		// 使用 AutoGo 的 PlayMP3 方法
		err := media.PlayMP3(audioPath)
		if err != nil {
			fmt.Printf("播放音频失败: %v\n", err)
		}
		return 0
	}))

	// 注册 stopAudio - 停止播放音频（空操作）
	state.SetGlobal("stopAudio", state.NewFunction(func(L *lua.LState) int {
		// AutoGo 不支持停止音频播放
		// 空操作
		return 0
	}))

	// 注册 installLrPkg - 安装 lrj 更新包（空操作）
	state.SetGlobal("installLrPkg", state.NewFunction(func(L *lua.LState) int {
		// AutoGo 不支持 lrj 更新包
		// 空操作
		return 0
	}))

	// 注册 getPackageName - 获取当前应用包名
	state.SetGlobal("getPackageName", state.NewFunction(func(L *lua.LState) int {
		// 使用 AutoGo 的 CurrentPackage 方法
		packageName := app.CurrentPackage()
		L.Push(lua.LString(packageName))
		return 1
	}))

	// 注册 getActivityName - 获取当前 Activity 名称
	state.SetGlobal("getActivityName", state.NewFunction(func(L *lua.LState) int {
		// 使用 AutoGo 的 CurrentActivity 方法
		activity := app.CurrentActivity()
		L.Push(lua.LString(activity))
		return 1
	}))

	// 注册 getSystemInfo - 获取系统信息
	state.SetGlobal("getSystemInfo", state.NewFunction(func(L *lua.LState) int {
		// 使用 AutoGo 的 GetDisplayInfo 方法获取设备信息
		width, height, dpi, rotation := device.GetDisplayInfo(0)

		// 构建系统信息表
		result := L.NewTable()
		result.RawSetString("width", lua.LNumber(width))
		result.RawSetString("height", lua.LNumber(height))
		result.RawSetString("dpi", lua.LNumber(dpi))
		result.RawSetString("rotation", lua.LNumber(rotation))
		result.RawSetString("cpuAbi", lua.LString(device.CpuAbi))
		result.RawSetString("buildId", lua.LString(device.BuildId))
		result.RawSetString("brand", lua.LString(device.Brand))
		result.RawSetString("device", lua.LString(device.Device))
		result.RawSetString("model", lua.LString(device.Model))
		result.RawSetString("product", lua.LString(device.Product))
		result.RawSetString("sdkInt", lua.LNumber(device.SdkInt))
		result.RawSetString("release", lua.LString(device.Release))
		L.Push(result)
		return 1
	}))

	// 注册 getInstalledApk - 获取已安装 APK
	state.SetGlobal("getInstalledApk", state.NewFunction(func(L *lua.LState) int {
		// 使用 AutoGo 的 GetList 方法
		appList := app.GetList(false)
		if len(appList) == 0 {
			L.Push(lua.LString(""))
			return 1
		}

		// 构建应用列表字符串
		var result string
		for i, appInfo := range appList {
			if i > 0 {
				result += "\n"
			}
			result += fmt.Sprintf("%s (%s)", appInfo.PackageName, appInfo.AppName)
		}
		L.Push(lua.LString(result))
		return 1
	}))

	// 注册 getInstalledApps - 获取已安装应用信息
	state.SetGlobal("getInstalledApps", state.NewFunction(func(L *lua.LState) int {
		// 使用 AutoGo 的 GetList 方法
		appList := app.GetList(false)
		if len(appList) == 0 {
			L.Push(lua.LNil)
			return 1
		}

		// 构建应用信息表
		result := L.NewTable()
		for i, appInfo := range appList {
			appTable := L.NewTable()
			appTable.RawSetString("packageName", lua.LString(appInfo.PackageName))
			appTable.RawSetString("appName", lua.LString(appInfo.AppName))
			appTable.RawSetString("versionName", lua.LString(appInfo.VersionName))
			appTable.RawSetString("versionCode", lua.LString(appInfo.VersionCode))
			appTable.RawSetString("isSystemApp", lua.LBool(appInfo.IsSystemApp))
			appTable.RawSetString("enabled", lua.LBool(appInfo.Enabled))
			result.RawSetInt(i+1, appTable)
		}
		L.Push(result)
		return 1
	}))

	// 注册 getLrApi - 获取懒人开放的 API 接口指针（空操作）
	state.SetGlobal("getLrApi", state.NewFunction(func(L *lua.LState) int {
		// AutoGo 不支持 FFI
		// 返回 nil
		L.Push(lua.LNil)
		return 1
	}))

	// 注册 getOaid - 获取设备 OAID（空操作）
	state.SetGlobal("getOaid", state.NewFunction(func(L *lua.LState) int {
		// AutoGo 运行在非 Android 环境，无法获取 OAID
		// 返回 nil
		L.Push(lua.LNil)
		return 1
	}))

	// 注册 getInsallAppInfos - 获取所有已安装 app 的详细信息（空操作）
	state.SetGlobal("getInsallAppInfos", state.NewFunction(func(L *lua.LState) int {
		// AutoGo 运行在非 Android 环境，无法获取已安装应用信息
		// 返回 nil
		L.Push(lua.LNil)
		return 1
	}))

	// 注册 getDisplayInfo - 获取设备分辨率等详细信息
	state.SetGlobal("getDisplayInfo", state.NewFunction(func(L *lua.LState) int {
		// 使用 AutoGo 的 GetDisplayInfo 方法
		width, height, dpi, rotation := device.GetDisplayInfo(0)

		// 构建 JSON 格式的设备信息
		info := map[string]interface{}{
			"width":    width,
			"height":   height,
			"dpi":      dpi,
			"rotation": rotation,
			"cpuAbi":   device.CpuAbi,
			"buildId":  device.BuildId,
			"brand":    device.Brand,
			"device":   device.Device,
			"model":    device.Model,
			"product":  device.Product,
			"sdkInt":   device.SdkInt,
			"release":  device.Release,
		}

		jsonData, err := json.Marshal(info)
		if err != nil {
			L.Push(lua.LNil)
		} else {
			L.Push(lua.LString(string(jsonData)))
		}
		return 1
	}))

	// 注册 getSensorsInfo - 获取所有传感器信息（空操作）
	state.SetGlobal("getSensorsInfo", state.NewFunction(func(L *lua.LState) int {
		// AutoGo 不支持传感器信息
		// 返回 nil
		L.Push(lua.LNil)
		return 1
	}))

	// 注册 installApk - 安装 APK
	state.SetGlobal("installApk", state.NewFunction(func(L *lua.LState) int {
		apkPath := L.CheckString(1)
		// 使用 AutoGo 的 Install 方法
		app.Install(apkPath)
		return 0
	}))

	// 注册 checkIsDebug - 判断当前是否处于调试状态（空操作）
	state.SetGlobal("checkIsDebug", state.NewFunction(func(L *lua.LState) int {
		// AutoGo 不支持调试状态检测
		// 返回 false
		L.Push(lua.LBool(false))
		return 1
	}))

	// 注册 getDeviceId - 获取设备 ID（空操作）
	state.SetGlobal("getDeviceId", state.NewFunction(func(L *lua.LState) int {
		// AutoGo 运行在非 Android 环境，无法获取设备 ID
		// 返回空字符串
		L.Push(lua.LString(""))
		return 1
	}))

	// 注册 getBrand - 获取设备品牌
	state.SetGlobal("getBrand", state.NewFunction(func(L *lua.LState) int {
		// 使用 AutoGo 的 Brand 变量
		L.Push(lua.LString(device.Brand))
		return 1
	}))

	// 注册 getBootLoader - 获取 Bootloader 版本号
	state.SetGlobal("getBootLoader", state.NewFunction(func(L *lua.LState) int {
		// 使用 AutoGo 的 Bootloader 变量
		L.Push(lua.LString(device.Bootloader))
		return 1
	}))

	// 注册 getBoard - 获取主板编号
	state.SetGlobal("getBoard", state.NewFunction(func(L *lua.LState) int {
		// 使用 AutoGo 的 Broad 变量
		L.Push(lua.LString(device.Broad))
		return 1
	}))

	// 注册 getManufacturer - 获取制造商代号（空操作）
	state.SetGlobal("getManufacturer", state.NewFunction(func(L *lua.LState) int {
		// AutoGo 运行在非 Android 环境，无法获取制造商代号
		// 返回空字符串
		L.Push(lua.LString(""))
		return 1
	}))

	// 注册 getProduct - 获取产品代号
	state.SetGlobal("getProduct", state.NewFunction(func(L *lua.LState) int {
		// 使用 AutoGo 的 Product 变量
		L.Push(lua.LString(device.Product))
		return 1
	}))

	// 注册 getDevice - 获取设备别名
	state.SetGlobal("getDevice", state.NewFunction(func(L *lua.LState) int {
		// 使用 AutoGo 的 Device 变量
		L.Push(lua.LString(device.Device))
		return 1
	}))

	// 注册 getModel - 获取设备型号
	state.SetGlobal("getModel", state.NewFunction(func(L *lua.LState) int {
		// 使用 AutoGo 的 Model 变量
		L.Push(lua.LString(device.Model))
		return 1
	}))

	// 注册 getHardware - 获取硬件序列号
	state.SetGlobal("getHardware", state.NewFunction(func(L *lua.LState) int {
		// 使用 AutoGo 的 Hardware 变量
		L.Push(lua.LString(device.Hardware))
		return 1
	}))

	// 注册 getId - 获取 id 修订号
	state.SetGlobal("getId", state.NewFunction(func(L *lua.LState) int {
		// 使用 AutoGo 的 Serial 变量
		L.Push(lua.LString(device.Serial))
		return 1
	}))

	// 注册 getFingerprint - 获取编译指纹
	state.SetGlobal("getFingerprint", state.NewFunction(func(L *lua.LState) int {
		// 使用 AutoGo 的 Fingerprint 变量
		L.Push(lua.LString(device.Fingerprint))
		return 1
	}))

	// 注册 getCpuAbi - 获取 CPU 支持类型
	state.SetGlobal("getCpuAbi", state.NewFunction(func(L *lua.LState) int {
		// 获取 CPU 支持类型
		L.Push(lua.LString(runtime.GOARCH))
		return 1
	}))

	// 注册 getCpuAbi2 - 获取 CPU 支持类型
	state.SetGlobal("getCpuAbi2", state.NewFunction(func(L *lua.LState) int {
		// 获取 CPU 支持类型
		L.Push(lua.LString(runtime.GOARCH))
		return 1
	}))

	// 注册 getSdkVersion - 获取 SDK 版本
	state.SetGlobal("getSdkVersion", state.NewFunction(func(L *lua.LState) int {
		// 使用 AutoGo 的 SdkInt 变量
		L.Push(lua.LString(fmt.Sprintf("%d", device.SdkInt)))
		return 1
	}))

	// 注册 getOsVersionName - 获取系统版本名
	state.SetGlobal("getOsVersionName", state.NewFunction(func(L *lua.LState) int {
		// 使用 AutoGo 的 Release 变量
		L.Push(lua.LString(device.Release))
		return 1
	}))

	// 注册 getWifiMac - 获取 WiFi MAC 地址
	state.SetGlobal("getWifiMac", state.NewFunction(func(L *lua.LState) int {
		// 使用 AutoGo 的 GetWifiMac 方法
		mac := device.GetWifiMac()
		L.Push(lua.LString(mac))
		return 1
	}))

	// 注册 runApp - 打开指定应用
	state.SetGlobal("runApp", state.NewFunction(func(L *lua.LState) int {
		packageName := L.CheckString(1)
		// 使用 AutoGo 的 Launch 方法
		app.Launch(packageName, 0)
		return 0
	}))

	// 注册 stopApp - 关闭指定应用
	state.SetGlobal("stopApp", state.NewFunction(func(L *lua.LState) int {
		packageName := L.CheckString(1)
		// 使用 AutoGo 的 ForceStop 方法
		app.ForceStop(packageName)
		return 0
	}))

	// 注册 scanImage - 刷新图片到媒体库
	state.SetGlobal("scanImage", state.NewFunction(func(L *lua.LState) int {
		imagePath := L.CheckString(1)
		// 使用 AutoGo 的 ScanFile 方法
		media.ScanFile(imagePath)
		return 0
	}))

	// 注册 getCurrentActivity - 获取最顶层 Activity 名称
	state.SetGlobal("getCurrentActivity", state.NewFunction(func(L *lua.LState) int {
		// 使用 AutoGo 的 CurrentActivity 方法
		activity := app.CurrentActivity()
		L.Push(lua.LString(activity))
		return 1
	}))

	// 注册 frontAppName - 获取前台应用包名
	state.SetGlobal("frontAppName", state.NewFunction(func(L *lua.LState) int {
		// 使用 AutoGo 的 CurrentPackage 方法
		packageName := app.CurrentPackage()
		L.Push(lua.LString(packageName))
		return 1
	}))

	// 注册 appIsFront - 判断指定包名是否为前台运行
	state.SetGlobal("appIsFront", state.NewFunction(func(L *lua.LState) int {
		packageName := L.CheckString(1)
		// 使用 AutoGo 的 CurrentPackage 方法
		currentPackage := app.CurrentPackage()
		isFront := currentPackage == packageName
		L.Push(lua.LBool(isFront))
		return 1
	}))

	// 注册 setMainThreadPause - 暂停主线程运行（空操作）
	state.SetGlobal("setMainThreadPause", state.NewFunction(func(L *lua.LState) int {
		// AutoGo 不支持多线程控制
		// 空操作
		return 0
	}))

	// 注册 setMainThreadResume - 恢复主线程运行（空操作）
	state.SetGlobal("setMainThreadResume", state.NewFunction(func(L *lua.LState) int {
		// AutoGo 不支持多线程控制
		// 空操作
		return 0
	}))

	// 注册 readPasteboard - 读取剪贴板内容
	state.SetGlobal("readPasteboard", state.NewFunction(func(L *lua.LState) int {
		// 使用 AutoGo 的 IME 包中的 GetClipText 方法
		text := ime.GetClipText()
		L.Push(lua.LString(text))
		return 1
	}))

	// 注册 writePasteboard - 写入剪贴板内容
	state.SetGlobal("writePasteboard", state.NewFunction(func(L *lua.LState) int {
		text := L.CheckString(1)
		// 使用 AutoGo 的 IME 包中的 SetClipText 方法
		success := ime.SetClipText(text)
		if !success {
			fmt.Printf("写入剪贴板失败\n")
		}
		return 0
	}))

	// 注册 appIsRunning - 判断 app 是否正在运行
	state.SetGlobal("appIsRunning", state.NewFunction(func(L *lua.LState) int {
		packageName := L.CheckString(1)
		// 使用 AutoGo 的 IsInstalled 方法判断应用是否安装
		// 注意：这里只能判断应用是否安装，无法判断是否正在运行
		isInstalled := app.IsInstalled(packageName)
		L.Push(lua.LBool(isInstalled))
		return 1
	}))

	// 注册 exec - 以最高权限执行命令（空操作）
	state.SetGlobal("exec", state.NewFunction(func(L *lua.LState) int {
		// AutoGo 不支持命令执行
		// 返回空字符串
		L.Push(lua.LString(""))
		return 1
	}))

	// 注册 runIntent - 构建并发送 Intent
	state.SetGlobal("runIntent", state.NewFunction(func(L *lua.LState) int {
		// 获取 Intent 表
		intentTable := L.CheckTable(1)

		// 构建 IntentOptions
		options := app.IntentOptions{}

		// 解析 action
		if actionValue := intentTable.RawGet(lua.LString("action")); actionValue != lua.LNil {
			if action, ok := actionValue.(lua.LString); ok {
				options.Action = string(action)
			}
		}

		// 解析 uri/data
		if uriValue := intentTable.RawGet(lua.LString("uri")); uriValue != lua.LNil {
			if uri, ok := uriValue.(lua.LString); ok {
				options.Data = string(uri)
			}
		}
		if dataValue := intentTable.RawGet(lua.LString("data")); dataValue != lua.LNil {
			if data, ok := dataValue.(lua.LString); ok {
				options.Data = string(data)
			}
		}

		// 解析 packageName
		if packageNameValue := intentTable.RawGet(lua.LString("packageName")); packageNameValue != lua.LNil {
			if packageName, ok := packageNameValue.(lua.LString); ok {
				options.PackageName = string(packageName)
			}
		}

		// 解析 type
		if typeValue := intentTable.RawGet(lua.LString("type")); typeValue != lua.LNil {
			if typeStr, ok := typeValue.(lua.LString); ok {
				options.Type = string(typeStr)
			}
		}

		// 解析 extras
		if extrasValue := intentTable.RawGet(lua.LString("extras")); extrasValue != lua.LNil {
			if extras, ok := extrasValue.(*lua.LTable); ok {
				options.Extras = make(map[string]string)
				extras.ForEach(func(key, value lua.LValue) {
					if keyStr, ok := key.(lua.LString); ok {
						if valueStr, ok := value.(lua.LString); ok {
							options.Extras[string(keyStr)] = string(valueStr)
						}
					}
				})
			}
		}

		// 使用 AutoGo 的 StartActivity 方法
		app.StartActivity(options)
		return 0
	}))

	// 注册 sendSms - 发送短信
	state.SetGlobal("sendSms", state.NewFunction(func(L *lua.LState) int {
		number := L.CheckString(1)
		content := L.CheckString(2)
		// 使用 AutoGo 的 SendSMS 方法
		media.SendSMS(number, content)
		return 0
	}))

	// 注册 phoneCall - 拨号或直接打出电话（空操作）
	state.SetGlobal("phoneCall", state.NewFunction(func(L *lua.LState) int {
		// AutoGo 不支持拨打电话
		// 空操作
		return 0
	}))

	// 注册 rnd - 生成随机整数
	state.SetGlobal("rnd", state.NewFunction(func(L *lua.LState) int {
		begin := L.CheckInt(1)
		end := L.CheckInt(2)

		// 生成指定范围的随机整数（包含端点）
		// 使用 Go 的 rand.Intn
		random := begin + int(time.Now().UnixNano()%int64(end-begin+1))
		if random < begin {
			random = begin
		}
		if random > end {
			random = end
		}

		L.Push(lua.LNumber(random))
		return 1
	}))

	// 注册 getDisplaySize - 获取屏幕分辨率
	state.SetGlobal("getDisplaySize", state.NewFunction(func(L *lua.LState) int {
		// 使用 AutoGo 的 GetDisplayInfo 方法
		width, height, _, _ := device.GetDisplayInfo(0)
		L.Push(lua.LNumber(width))
		L.Push(lua.LNumber(height))
		return 2
	}))

	// 注册 getScreenInfo - 获取屏幕信息（兼容性方法）
	state.SetGlobal("getScreenInfo", state.NewFunction(func(L *lua.LState) int {
		// 使用 AutoGo 的 GetDisplayInfo 方法
		width, height, _, _ := device.GetDisplayInfo(0)
		L.Push(lua.LNumber(width))
		L.Push(lua.LNumber(height))
		return 2
	}))

	// 注册 getScreenScale - 获取屏幕缩放比例（兼容性方法）
	state.SetGlobal("getScreenScale", state.NewFunction(func(L *lua.LState) int {
		// AutoGo 不支持屏幕缩放比例
		// 返回默认值 1.0
		L.Push(lua.LNumber(1.0))
		return 1
	}))

	// 注册 getScreenDensity - 获取屏幕密度（兼容性方法）
	state.SetGlobal("getScreenDensity", state.NewFunction(func(L *lua.LState) int {
		// 使用 AutoGo 的 GetDisplayInfo 方法获取 DPI
		_, _, dpi, _ := device.GetDisplayInfo(0)
		L.Push(lua.LNumber(dpi))
		return 1
	}))

	// 注册 setDisplayPowerOff - 设置息屏运行
	state.SetGlobal("setDisplayPowerOff", state.NewFunction(func(L *lua.LState) int {
		isPoweroff := L.CheckBool(1)
		// 使用 AutoGo 的 SetDisplayPower 方法
		device.SetDisplayPower(!isPoweroff)
		return 0
	}))

	// 注册 getDisplayRotate - 获取屏幕旋转方向
	state.SetGlobal("getDisplayRotate", state.NewFunction(func(L *lua.LState) int {
		// 使用 AutoGo 的 GetDisplayInfo 方法
		_, _, _, rotation := device.GetDisplayInfo(0)
		L.Push(lua.LNumber(rotation))
		return 1
	}))

	// 注册 setDpiToVir - 切换虚拟分辨率（空操作）
	state.SetGlobal("setDpiToVir", state.NewFunction(func(L *lua.LState) int {
		// AutoGo 不支持虚拟分辨率
		// 空操作
		return 0
	}))

	// 注册 setDpiToRealy - 恢复真实分辨率（空操作）
	state.SetGlobal("setDpiToRealy", state.NewFunction(func(L *lua.LState) int {
		// AutoGo 不支持虚拟分辨率
		// 空操作
		return 0
	}))

	// 注册 setLogOff - 关闭或开启日志输出（空操作）
	state.SetGlobal("setLogOff", state.NewFunction(func(L *lua.LState) int {
		// AutoGo 不支持日志开关
		// 空操作
		return 0
	}))

	// 注册 extractAssets - 将打包资源释放到目录（空操作）
	state.SetGlobal("extractAssets", state.NewFunction(func(L *lua.LState) int {
		// AutoGo 不支持资源释放
		// 空操作
		return 0
	}))

	// 注册 extractApkAssets - 将指定资源释放到目录（空操作）
	state.SetGlobal("extractApkAssets", state.NewFunction(func(L *lua.LState) int {
		// AutoGo 不支持资源释放
		// 空操作
		return 0
	}))

	// 注册 zip - 压缩文件或文件夹
	state.SetGlobal("zip", state.NewFunction(func(L *lua.LState) int {
		filePath := L.CheckString(1)
		zipPath := L.CheckString(2)

		// 使用 Go 的 archive/zip 包压缩文件
		err := createZip(filePath, zipPath)
		if err != nil {
			fmt.Printf("压缩失败: %v\n", err)
		}
		return 0
	}))

	// 注册 unZip - 解压文件
	state.SetGlobal("unZip", state.NewFunction(func(L *lua.LState) int {
		zipPath := L.CheckString(1)
		outDir := L.CheckString(2)
		// 可选参数：密码和字符集
		pass := L.OptString(3, "")
		_ = pass // 暂不支持密码
		charset := L.OptString(4, "UTF-8")
		_ = charset // 暂不支持字符集

		// 使用 Go 的 archive/zip 包解压文件
		err := extractZip(zipPath, outDir)
		if err != nil {
			fmt.Printf("解压失败: %v\n", err)
		}
		return 0
	}))

	// 注册 setTimer - 定时执行指定函数（空操作）
	state.SetGlobal("setTimer", state.NewFunction(func(L *lua.LState) int {
		// AutoGo 不支持定时器
		// 空操作
		return 0
	}))

	// 注册 setBTEnable - 打开或关闭蓝牙（空操作）
	state.SetGlobal("setBTEnable", state.NewFunction(func(L *lua.LState) int {
		// AutoGo 运行在非 Android 环境，无法控制蓝牙
		// 空操作
		return 0
	}))

	// 注册 setWifiEnable - 打开或关闭 WiFi（空操作）
	state.SetGlobal("setWifiEnable", state.NewFunction(func(L *lua.LState) int {
		// AutoGo 运行在非 Android 环境，无法控制 WiFi
		// 空操作
		return 0
	}))

	// 注册 setAirplaneMode - 打开或关闭飞行模式（空操作）
	state.SetGlobal("setAirplaneMode", state.NewFunction(func(L *lua.LState) int {
		// AutoGo 运行在非 Android 环境，无法控制飞行模式
		// 空操作
		return 0
	}))

	// 注册 sleep - 线程休眠
	state.SetGlobal("sleep", state.NewFunction(func(L *lua.LState) int {
		sleepTime := L.CheckInt(1)
		// 线程休眠（毫秒）
		time.Sleep(time.Duration(sleepTime) * time.Millisecond)
		return 0
	}))

	// 注册 printEx - 调试输出（原生标准输出）
	state.SetGlobal("printEx", state.NewFunction(func(L *lua.LState) int {
		str := L.CheckString(1)
		// 原生标准输出
		fmt.Print(str)
		return 0
	}))

	// 注册 lockScreen - 保持屏幕常亮（空操作）
	state.SetGlobal("lockScreen", state.NewFunction(func(L *lua.LState) int {
		// AutoGo 运行在非 Android 环境，无法控制屏幕常亮
		// 空操作
		m.screenLocked = true
		return 0
	}))

	// 注册 unLockScreen - 释放常亮状态（空操作）
	state.SetGlobal("unLockScreen", state.NewFunction(func(L *lua.LState) int {
		// AutoGo 运行在非 Android 环境，无法控制屏幕常亮
		// 空操作
		m.screenLocked = false
		return 0
	}))

	// 注册 setUserEventCallBack - 设置悬浮按钮菜单自定义事件回调（空操作）
	state.SetGlobal("setUserEventCallBack", state.NewFunction(func(L *lua.LState) int {
		// AutoGo 不支持悬浮窗
		// 空操作
		return 0
	}))

	// 注册 setPluginEventCallBack - 设置插件事件回调（空操作）
	state.SetGlobal("setPluginEventCallBack", state.NewFunction(func(L *lua.LState) int {
		// AutoGo 不支持插件系统
		// 返回空字符串
		L.Push(lua.LString(""))
		return 1
	}))

	// 注册 sendPluginEvent - 发送插件事件（空操作）
	state.SetGlobal("sendPluginEvent", state.NewFunction(func(L *lua.LState) int {
		// AutoGo 不支持插件系统
		// 空操作
		return 0
	}))

	// 注册 setNotifyEventCallBack - 设置系统通知事件回调（空操作）
	state.SetGlobal("setNotifyEventCallBack", state.NewFunction(func(L *lua.LState) int {
		// AutoGo 不支持系统通知
		// 空操作
		return 0
	}))

	// 注册 setStopCallBack - 设置脚本结束回调（空操作）
	state.SetGlobal("setStopCallBack", state.NewFunction(func(L *lua.LState) int {
		// AutoGo 不支持脚本结束回调
		// 空操作
		return 0
	}))

	// 注册 exitScript - 停止脚本运行（空操作）
	state.SetGlobal("exitScript", state.NewFunction(func(L *lua.LState) int {
		// AutoGo 不支持脚本停止
		// 空操作
		return 0
	}))

	// 注册 getRunEnvType - 获取运行环境类型（空操作）
	state.SetGlobal("getRunEnvType", state.NewFunction(func(L *lua.LState) int {
		// AutoGo 不是 Android 应用
		// 返回默认值 0
		L.Push(lua.LNumber(0))
		return 1
	}))

	// 注册 setRootEnvMode - 设置 root 模式（空操作）
	state.SetGlobal("setRootEnvMode", state.NewFunction(func(L *lua.LState) int {
		// AutoGo 不是 Android 应用
		// 空操作
		return 0
	}))

	// 注册 setHandleEnvMode - 设置 handle 模式（空操作）
	state.SetGlobal("setHandleEnvMode", state.NewFunction(func(L *lua.LState) int {
		// AutoGo 不是 Android 应用
		// 空操作
		return 0
	}))

	// 注册 setAccessibilityEnvMode - 设置无障碍模式（空操作）
	state.SetGlobal("setAccessibilityEnvMode", state.NewFunction(func(L *lua.LState) int {
		// AutoGo 不是 Android 应用
		// 空操作
		return 0
	}))

	// 注册 LuaEngine.registerExitCallback - Java 层设置脚本结束回调（空操作）
	luaEngineTable := state.NewTable()
	luaEngineTable.RawSetString("registerExitCallback", state.NewFunction(func(L *lua.LState) int {
		// AutoGo 不支持 Java 层回调
		// 空操作
		return 0
	}))

	// 注册 LuaEngine.setSnapCacheBitmap - 设置截图缓存图片（空操作）
	luaEngineTable.RawSetString("setSnapCacheBitmap", state.NewFunction(func(L *lua.LState) int {
		// AutoGo 不支持截图缓存
		// 空操作
		return 0
	}))

	// 注册 setSnapCacheTime - 设置截图缓存时长（空操作）
	state.SetGlobal("setSnapCacheTime", state.NewFunction(func(L *lua.LState) int {
		// AutoGo 不支持截图缓存
		// 空操作
		return 0
	}))

	// 注册 LuaEngine 表到全局
	state.SetGlobal("LuaEngine", luaEngineTable)

	return nil
}

// GetModule 获取模块实例
func GetModule() model.Module {
	return &DeviceModule{}
}

// createZip 创建 ZIP 文件
func createZip(filePath, zipPath string) error {
	// 创建 ZIP 文件
	zipFile, err := os.Create(zipPath)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// 检查是文件还是目录
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return err
	}

	if fileInfo.IsDir() {
		// 压缩目录
		return filepath.Walk(filePath, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}

			// 创建 ZIP 文件头
			header, err := zip.FileInfoHeader(info)
			if err != nil {
				return err
			}

			// 计算相对路径
			relPath, err := filepath.Rel(filePath, path)
			if err != nil {
				return err
			}
			header.Name = relPath

			// 写入文件头
			writer, err := zipWriter.CreateHeader(header)
			if err != nil {
				return err
			}

			// 写入文件内容
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			_, err = io.Copy(writer, file)
			return err
		})
	} else {
		// 压缩单个文件
		header, err := zip.FileInfoHeader(fileInfo)
		if err != nil {
			return err
		}
		header.Name = filepath.Base(filePath)

		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}

		file, err := os.Open(filePath)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = io.Copy(writer, file)
		return err
	}
}

// extractZip 解压 ZIP 文件
func extractZip(zipPath, outDir string) error {
	// 打开 ZIP 文件
	zipReader, err := zip.OpenReader(zipPath)
	if err != nil {
		return err
	}
	defer zipReader.Close()

	// 创建输出目录
	if err := os.MkdirAll(outDir, 0755); err != nil {
		return err
	}

	// 解压每个文件
	for _, file := range zipReader.File {
		// 构建输出文件路径
		filePath := filepath.Join(outDir, file.Name)

		// 如果是目录，创建目录
		if file.FileInfo().IsDir() {
			if err := os.MkdirAll(filePath, file.Mode()); err != nil {
				return err
			}
			continue
		}

		// 创建文件目录
		if err := os.MkdirAll(filepath.Dir(filePath), 0755); err != nil {
			return err
		}

		// 打开 ZIP 中的文件
		fileReader, err := file.Open()
		if err != nil {
			return err
		}
		defer fileReader.Close()

		// 创建输出文件
		outFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}
		defer outFile.Close()

		// 复制文件内容
		if _, err := io.Copy(outFile, fileReader); err != nil {
			return err
		}
	}

	return nil
}
