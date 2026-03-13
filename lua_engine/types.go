package lua_engine

import (
	"encoding/json"
	"fmt"
	"sync"

	lua "github.com/yuin/gopher-lua"
)

// EngineConfig 引擎配置选项
type EngineConfig struct {
	AutoInjectMethods bool     // 是否自动注入所有方法，默认为 true
	WhiteList         []string // 白名单：只加载这些模块，空列表 = 加载所有
	BlackList         []string // 黑名单：跳过这些模块，空列表 = 不跳过任何
	FailFast          bool     // 是否在模块加载失败时立即失败，false = 跳过失败模块继续
}

// LuaEngine Lua 引擎
type LuaEngine struct {
	state  *lua.LState
	mu     sync.RWMutex
	config EngineConfig
}

// DefaultConfig 返回默认配置
func DefaultConfig() EngineConfig {
	return EngineConfig{
		AutoInjectMethods: true,
		WhiteList:         []string{}, // 默认为空，加载所有模块
		BlackList:         []string{}, // 默认为空，不跳过任何模块
		FailFast:          false,      // 默认为 false，模块加载失败时跳过继续
	}
}

type MethodRegistry struct {
	methods map[string]MethodInfo
	mu      sync.RWMutex
}

type MethodInfo struct {
	Name        string
	Description string
	GoFunc      interface{}
	Overridable bool
	Overridden  bool
	LuaFunc     *lua.LFunction
}

var (
	registry     *MethodRegistry
	registryOnce sync.Once
)

func initRegistry() {
	registryOnce.Do(func() {
		registry = &MethodRegistry{
			methods: make(map[string]MethodInfo),
		}
	})
}

func GetRegistry() *MethodRegistry {
	initRegistry()
	return registry
}

func (r *MethodRegistry) ListMethods() []MethodInfo {
	r.mu.RLock()
	defer r.mu.RUnlock()

	methods := make([]MethodInfo, 0, len(r.methods))
	for _, method := range r.methods {
		methods = append(methods, method)
	}
	return methods
}

func (r *MethodRegistry) GetMethod(name string) (MethodInfo, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	method, exists := r.methods[name]
	return method, exists
}

func (r *MethodRegistry) OverrideMethod(name string, luaFunc *lua.LFunction) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	method, exists := r.methods[name]
	if !exists || !method.Overridable {
		return false
	}

	method.Overridden = true
	method.LuaFunc = luaFunc
	r.methods[name] = method
	return true
}

func (r *MethodRegistry) RestoreMethod(name string) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	method, exists := r.methods[name]
	if !exists {
		return false
	}

	method.Overridden = false
	method.LuaFunc = nil
	r.methods[name] = method
	return true
}

func (r *MethodRegistry) RemoveMethod(name string) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, exists := r.methods[name]
	if !exists {
		return false
	}

	delete(r.methods, name)
	return true
}

func (r *MethodRegistry) Contains(name string) bool {
	r.mu.RLock()
	defer r.mu.RUnlock()

	_, exists := r.methods[name]
	return exists
}

func (r *MethodRegistry) Clear() {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.methods = make(map[string]MethodInfo)
}

func (r *MethodRegistry) Count() int {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return len(r.methods)
}

func RegisterMethod(name, description string, goFunc interface{}, overridable bool) {
	initRegistry()
	registry.mu.Lock()
	defer registry.mu.Unlock()

	registry.methods[name] = MethodInfo{
		Name:        name,
		Description: description,
		GoFunc:      goFunc,
		Overridable: overridable,
		Overridden:  false,
	}
}

func (r *MethodRegistry) ExportMethodsJSON() (string, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	methods := make([]map[string]interface{}, 0, len(r.methods))
	for _, method := range r.methods {
		methodMap := map[string]interface{}{
			"name":        method.Name,
			"description": method.Description,
			"overridable": method.Overridable,
			"overridden":  method.Overridden,
		}
		methods = append(methods, methodMap)
	}

	jsonData, err := json.MarshalIndent(methods, "", "  ")
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

func (r *MethodRegistry) ExportMethodsLuaTable() string {
	r.mu.RLock()
	defer r.mu.RUnlock()

	luaTable := "AutoGo = {\n"
	luaTable += "  methods = {\n"

	for name, method := range r.methods {
		luaTable += fmt.Sprintf("    [\"%s\"] = {\n", name)
		luaTable += fmt.Sprintf("      description = \"%s\",\n", method.Description)
		luaTable += fmt.Sprintf("      overridable = %t,\n", method.Overridable)
		luaTable += fmt.Sprintf("      overridden = %t,\n", method.Overridden)
		luaTable += "    },\n"
	}

	luaTable += "  },\n"
	luaTable += "}\n"

	return luaTable
}

func (r *MethodRegistry) GenerateLuaDocumentation() string {
	r.mu.RLock()
	defer r.mu.RUnlock()

	doc := "-- AutoGo Lua API 文档\n"
	doc += "-- 自动生成时间: " + getTimestamp() + "\n\n"

	doc += "--------------------------------------------------\n"
	doc += "-- 应用管理 (app)\n"
	doc += "--------------------------------------------------\n"
	doc += "app_currentPackage() -> string\n"
	doc += "  获取当前页面应用包名\n\n"
	doc += "app_currentActivity() -> string\n"
	doc += "  获取当前页面应用类名\n\n"
	doc += "app_launch(packageName: string) -> boolean\n"
	doc += "  通过应用包名启动应用\n\n"
	doc += "app_openAppSetting(packageName: string) -> boolean\n"
	doc += "  打开应用的详情页(设置页)\n\n"
	doc += "app_viewFile(path: string)\n"
	doc += "  用其他应用查看文件\n\n"
	doc += "app_editFile(path: string)\n"
	doc += "  用其他应用编辑文件\n\n"
	doc += "app_uninstall(packageName: string)\n"
	doc += "  卸载应用\n\n"
	doc += "app_install(path: string)\n"
	doc += "  安装应用\n\n"
	doc += "app_isInstalled(packageName: string) -> boolean\n"
	doc += "  判断是否已经安装某个应用\n\n"
	doc += "app_clear(packageName: string)\n"
	doc += "  清除应用数据\n\n"
	doc += "app_forceStop(packageName: string)\n"
	doc += "  强制停止应用\n\n"
	doc += "app_disable(packageName: string)\n"
	doc += "  禁用应用\n\n"
	doc += "app_ignoreBattOpt(packageName: string)\n"
	doc += "  忽略电池优化\n\n"
	doc += "app_enable(packageName: string)\n"
	doc += "  启用应用\n\n"
	doc += "app_getBrowserPackage() -> string\n"
	doc += "  获取系统默认浏览器包名\n\n"
	doc += "app_openUrl(url: string)\n"
	doc += "  用浏览器打开网站url\n\n"

	doc += "--------------------------------------------------\n"
	doc += "-- 设备管理 (device)\n"
	doc += "--------------------------------------------------\n"
	doc += "device.width -> number (只读)\n"
	doc += "  设备分辨率宽度\n\n"
	doc += "device.height -> number (只读)\n"
	doc += "  设备分辨率高度\n\n"
	doc += "device.sdkInt -> number (只读)\n"
	doc += "  安卓系统API版本\n\n"
	doc += "device.cpuAbi -> string (只读)\n"
	doc += "  设备的CPU架构\n\n"
	doc += "device_getImei() -> string\n"
	doc += "  返回设备的IMEI\n\n"
	doc += "device_getAndroidId() -> string\n"
	doc += "  返回设备的Android ID\n\n"
	doc += "device_getWifiMac() -> string\n"
	doc += "  获取设备WIFI-MAC\n\n"
	doc += "device_getWlanMac() -> string\n"
	doc += "  获取设备以太网MAC\n\n"
	doc += "device_getIp() -> string\n"
	doc += "  获取设备局域网IP地址\n\n"
	doc += "device_getBrightness() -> string\n"
	doc += "  返回当前的(手动)亮度\n\n"
	doc += "device_getBrightnessMode() -> string\n"
	doc += "  返回当前亮度模式\n\n"
	doc += "device_getMusicVolume() -> number\n"
	doc += "  返回当前媒体音量\n\n"
	doc += "device_getNotificationVolume() -> number\n"
	doc += "  返回当前通知音量\n\n"
	doc += "device_getAlarmVolume() -> number\n"
	doc += "  返回当前闹钟音量\n\n"
	doc += "device_getMusicMaxVolume() -> number\n"
	doc += "  返回媒体音量的最大值\n\n"
	doc += "device_getNotificationMaxVolume() -> number\n"
	doc += "  返回通知音量的最大值\n\n"
	doc += "device_getAlarmMaxVolume() -> number\n"
	doc += "  返回闹钟音量的最大值\n\n"
	doc += "device_setMusicVolume(volume: number)\n"
	doc += "  设置当前媒体音量\n\n"
	doc += "device_setNotificationVolume(volume: number)\n"
	doc += "  设置当前通知音量\n\n"
	doc += "device_setAlarmVolume(volume: number)\n"
	doc += "  设置当前闹钟音量\n\n"
	doc += "device_getBattery() -> number\n"
	doc += "  返回当前电量百分比\n\n"
	doc += "device_getBatteryStatus() -> number\n"
	doc += "  获取电池状态\n\n"
	doc += "device_setBatteryStatus(value: number)\n"
	doc += "  模拟电池状态\n\n"
	doc += "device_setBatteryLevel(value: number)\n"
	doc += "  模拟电池电量百分百\n\n"
	doc += "device_getTotalMem() -> number\n"
	doc += "  返回设备内存总量\n\n"
	doc += "device_getAvailMem() -> number\n"
	doc += "  返回设备当前可用的内存\n\n"
	doc += "device_isScreenOn() -> boolean\n"
	doc += "  返回设备屏幕是否是亮着的\n\n"
	doc += "device_isScreenUnlock() -> boolean\n"
	doc += "  返回屏幕锁是否已经解开\n\n"
	doc += "device_setScreenMode(mode: number)\n"
	doc += "  设置屏幕显示模式\n\n"
	doc += "device_wakeUp()\n"
	doc += "  唤醒设备\n\n"
	doc += "device_keepScreenOn()\n"
	doc += "  保持屏幕常亮\n\n"
	doc += "device_vibrate(ms: number)\n"
	doc += "  使设备震动一段时间\n\n"
	doc += "device_cancelVibration()\n"
	doc += "  如果设备处于震动状态，则取消震动\n\n"

	doc += "--------------------------------------------------\n"
	doc += "-- 触摸操作 (touch)\n"
	doc += "--------------------------------------------------\n"
	doc += "touchDown(x: number, y: number, fingerID: number)\n"
	doc += "  按下屏幕\n\n"
	doc += "touchMove(x: number, y: number, fingerID: number)\n"
	doc += "  移动手指\n\n"
	doc += "touchUp(x: number, y: number, fingerID: number)\n"
	doc += "  抬起手指\n\n"
	doc += "click(x: number, y: number, fingerID: number)\n"
	doc += "  点击\n\n"
	doc += "longClick(x: number, y: number, duration: number)\n"
	doc += "  长按\n\n"
	doc += "swipe(x1: number, y1: number, x2: number, y2: number, duration: number)\n"
	doc += "  滑动\n\n"
	doc += "swipe2(x1: number, y1: number, x2: number, y2: number, duration: number)\n"
	doc += "  滑动(两点)\n\n"
	doc += "home()\n"
	doc += "  按下Home键\n\n"
	doc += "back()\n"
	doc += "  按下返回键\n\n"
	doc += "recents()\n"
	doc += "  按下最近任务键\n\n"
	doc += "powerDialog()\n"
	doc += "  长按电源键\n\n"
	doc += "notifications()\n"
	doc += "  下拉通知栏\n\n"
	doc += "quickSettings()\n"
	doc += "  下拉快捷设置\n\n"
	doc += "volumeUp()\n"
	doc += "  按下音量加键\n\n"
	doc += "volumeDown()\n"
	doc += "  按下音量减键\n\n"
	doc += "camera()\n"
	doc += "  按下相机键\n\n"
	doc += "keyAction(code: number)\n"
	doc += "  按键动作\n\n"

	doc += "--------------------------------------------------\n"
	doc += "-- 文件操作 (files)\n"
	doc += "--------------------------------------------------\n"
	doc += "files_isDir(path: string) -> boolean\n"
	doc += "  返回路径path是否是文件夹\n\n"
	doc += "files_isFile(path: string) -> boolean\n"
	doc += "  返回路径path是否是文件\n\n"
	doc += "files_isEmptyDir(path: string) -> boolean\n"
	doc += "  返回文件夹path是否为空文件夹\n\n"
	doc += "files_create(path: string) -> boolean\n"
	doc += "  创建一个文件或文件夹\n\n"
	doc += "files_createWithDirs(path: string) -> boolean\n"
	doc += "  创建一个文件或文件夹并确保所在文件夹存在\n\n"
	doc += "files_exists(path: string) -> boolean\n"
	doc += "  返回在路径path处的文件是否存在\n\n"
	doc += "files_ensureDir(path: string) -> boolean\n"
	doc += "  确保路径path所在的文件夹存在\n\n"
	doc += "files_read(path: string) -> string\n"
	doc += "  读取文本文件path的所有内容并返回\n\n"
	doc += "files_readBytes(path: string) -> string\n"
	doc += "  读取文件path的所有内容并返回\n\n"
	doc += "files_write(path: string, text: string)\n"
	doc += "  把text写入到文件path中\n\n"
	doc += "files_writeBytes(path: string, bytes: string)\n"
	doc += "  把bytes写入到文件path中\n\n"
	doc += "files_append(path: string, text: string)\n"
	doc += "  把text追加到文件path的末尾\n\n"
	doc += "files_appendBytes(path: string, bytes: string)\n"
	doc += "  把bytes追加到文件path的末尾\n\n"
	doc += "files_copy(fromPath: string, toPath: string) -> boolean\n"
	doc += "  复制文件\n\n"
	doc += "files_move(fromPath: string, toPath: string) -> boolean\n"
	doc += "  移动文件\n\n"
	doc += "files_rename(path: string, newName: string) -> boolean\n"
	doc += "  重命名文件\n\n"
	doc += "files_remove(path: string) -> boolean\n"
	doc += "  删除文件或文件夹\n\n"
	doc += "files_getName(path: string) -> string\n"
	doc += "  返回文件的文件名\n\n"
	doc += "files_getNameWithoutExtension(path: string) -> string\n"
	doc += "  返回不含拓展名的文件的文件名\n\n"
	doc += "files_getExtension(path: string) -> string\n"
	doc += "  返回文件的拓展名\n\n"
	doc += "files_path(relativePath: string) -> string\n"
	doc += "  返回相对路径对应的绝对路径\n\n"
	doc += "files_listDir(path: string) -> table\n"
	doc += "  列出文件夹path下的所有文件和文件夹\n\n"

	doc += "--------------------------------------------------\n"
	doc += "-- 图像处理 (images)\n"
	doc += "--------------------------------------------------\n"
	doc += "images_pixel(x: number, y: number) -> string\n"
	doc += "  获取指定坐标的像素颜色\n\n"
	doc += "images_captureScreen(x1: number, y1: number, x2: number, y2: number) -> userdata\n"
	doc += "  截取屏幕\n\n"
	doc += "images_cmpColor(x: number, y: number, colorStr: string, sim: number) -> boolean\n"
	doc += "  比较颜色\n\n"
	doc += "images_findColor(x1: number, y1: number, x2: number, y2: number, colorStr: string, sim: number, dir: number) -> number, number\n"
	doc += "  查找颜色\n\n"
	doc += "images_getColorCountInRegion(x1: number, y1: number, x2: number, y2: number, colorStr: string, sim: number) -> number\n"
	doc += "  获取区域内指定颜色的数量\n\n"
	doc += "images_detectsMultiColors(colors: string, sim: number) -> boolean\n"
	doc += "  检测多点颜色\n\n"
	doc += "images_findMultiColors(x1: number, y1: number, x2: number, y2: number, colors: string, sim: number, dir: number) -> number, number\n"
	doc += "  查找多点颜色\n\n"
	doc += "images_readFromPath(path: string) -> userdata\n"
	doc += "  从路径读取图片\n\n"
	doc += "images_readFromUrl(url: string) -> userdata\n"
	doc += "  从URL读取图片\n\n"
	doc += "images_readFromBase64(base64Str: string) -> userdata\n"
	doc += "  从Base64读取图片\n\n"
	doc += "images_readFromBytes(bytes: string) -> userdata\n"
	doc += "  从字节数组读取图片\n\n"
	doc += "images_save(img: userdata, path: string, quality: number) -> boolean\n"
	doc += "  保存图片\n\n"
	doc += "images_encodeToBase64(img: userdata, format: string, quality: number) -> string\n"
	doc += "  编码为Base64\n\n"
	doc += "images_encodeToBytes(img: userdata, format: string, quality: number) -> string\n"
	doc += "  编码为字节数组\n\n"
	doc += "images_clip(img: userdata, x1: number, y1: number, x2: number, y2: number) -> userdata\n"
	doc += "  裁剪图片\n\n"
	doc += "images_resize(img: userdata, width: number, height: number) -> userdata\n"
	doc += "  调整图片大小\n\n"
	doc += "images_rotate(img: userdata, degree: number) -> userdata\n"
	doc += "  旋转图片\n\n"
	doc += "images_grayscale(img: userdata) -> userdata\n"
	doc += "  灰度化\n\n"
	doc += "images_applyThreshold(img: userdata, threshold: number, maxVal: number, typ: string) -> userdata\n"
	doc += "  应用阈值\n\n"
	doc += "images_applyAdaptiveThreshold(img: userdata, maxValue: number, adaptiveMethod: string, thresholdType: string, blockSize: number, C: number) -> userdata\n"
	doc += "  应用自适应阈值\n\n"
	doc += "images_applyBinarization(img: userdata, threshold: number) -> userdata\n"
	doc += "  二值化\n\n"

	doc += "--------------------------------------------------\n"
	doc += "-- 存储管理 (storages)\n"
	doc += "--------------------------------------------------\n"
	doc += "storages_get(table: string, key: string) -> string\n"
	doc += "  从指定表中获取键值\n\n"
	doc += "storages_put(table: string, key: string, value: string)\n"
	doc += "  写入键值对\n\n"
	doc += "storages_remove(table: string, key: string)\n"
	doc += "  删除指定键\n\n"
	doc += "storages_contains(table: string, key: string) -> boolean\n"
	doc += "  判断键是否存在\n\n"
	doc += "storages_getAll(table: string) -> table\n"
	doc += "  获取所有键值对\n\n"
	doc += "storages_clear(table: string)\n"
	doc += "  清空指定表数据\n\n"

	doc += "--------------------------------------------------\n"
	doc += "-- 系统管理 (system)\n"
	doc += "--------------------------------------------------\n"
	doc += "system_getPid(processName: string) -> number\n"
	doc += "  获取进程ID\n\n"
	doc += "system_getMemoryUsage(pid: number) -> number\n"
	doc += "  获取内存使用\n\n"
	doc += "system_getCpuUsage(pid: number) -> number\n"
	doc += "  获取CPU使用率\n\n"
	doc += "system_restartSelf()\n"
	doc += "  重启自身\n\n"
	doc += "system_setBootStart(enable: boolean)\n"
	doc += "  设置开机自启\n\n"

	doc += "--------------------------------------------------\n"
	doc += "-- 网络请求 (http)\n"
	doc += "--------------------------------------------------\n"
	doc += "http_get(url: string, timeout: number) -> number, string\n"
	doc += "  发送GET请求\n\n"
	doc += "http_postMultipart(url: string, fileName: string, fileData: string) -> number, string\n"
	doc += "  发送Multipart POST请求\n\n"

	doc += "--------------------------------------------------\n"
	doc += "-- 媒体管理 (media)\n"
	doc += "--------------------------------------------------\n"
	doc += "media_scanFile(path: string)\n"
	doc += "  扫描路径path的媒体文件\n\n"

	doc += "--------------------------------------------------\n"
	doc += "-- 图像识别 (opencv)\n"
	doc += "--------------------------------------------------\n"
	doc += "opencv_findImage(x1: number, y1: number, x2: number, y2: number, template: string, isGray: boolean, scalingFactor: number, sim: number) -> number, number\n"
	doc += "  在指定区域内查找匹配的图片模板\n\n"

	doc += "--------------------------------------------------\n"
	doc += "-- 文字识别 (ppocr)\n"
	doc += "--------------------------------------------------\n"
	doc += "ppocr_ocr(x1: number, y1: number, x2: number, y2: number, colorStr: string) -> table\n"
	doc += "  识别屏幕文字\n\n"
	doc += "ppocr_ocrFromBase64(b64: string, colorStr: string) -> table\n"
	doc += "  识别Base64图片文字\n\n"
	doc += "ppocr_ocrFromPath(path: string, colorStr: string) -> table\n"
	doc += "  识别文件图片文字\n\n"

	doc += "--------------------------------------------------\n"
	doc += "-- 方法管理 (method)\n"
	doc += "--------------------------------------------------\n"
	doc += "registerMethod(name: string, description: string, goFunc: function, overridable: boolean)\n"
	doc += "  注册新方法到Lua引擎\n\n"
	doc += "unregisterMethod(name: string) -> boolean\n"
	doc += "  从Lua引擎中移除方法\n\n"
	doc += "listMethods() -> table\n"
	doc += "  列出所有已注册的方法\n\n"
	doc += "overrideMethod(name: string, luaFunc: function) -> boolean\n"
	doc += "  用Lua函数重写已注册的方法\n\n"
	doc += "restoreMethod(name: string) -> boolean\n"
	doc += "  恢复被重写的方法\n\n"

	doc += "--------------------------------------------------\n"
	doc += "-- 协程管理 (coroutine)\n"
	doc += "--------------------------------------------------\n"
	doc += "createCoroutine(script: string) -> number\n"
	doc += "  创建一个新的协程，返回协程ID\n\n"
	doc += "resumeCoroutine(id: number) -> table, string\n"
	doc += "  恢复协程执行，返回结果和状态\n\n"
	doc += "yieldCoroutine(id: number) -> boolean\n"
	doc += "  挂起当前协程\n\n"
	doc += "listCoroutines() -> table\n"
	doc += "  列出所有协程\n\n"
	doc += "removeCoroutine(id: number) -> boolean\n"
	doc += "  移除指定的协程\n\n"

	return doc
}

func getTimestamp() string {
	return "2026-03-12"
}
