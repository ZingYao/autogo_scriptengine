package js_engine

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"sync"

	"github.com/dop251/goja"
)

// EngineConfig 引擎配置选项
type EngineConfig struct {
	AutoInjectMethods bool     // 是否自动注入所有方法，默认为 true
	WhiteList        []string // 白名单：只加载这些模块，空列表 = 加载所有
	BlackList        []string // 黑名单：跳过这些模块，空列表 = 不跳过任何
	FailFast        bool     // 是否在模块加载失败时立即失败，false = 跳过失败模块继续
	FileSystem      fs.FS    // 文件系统，用于 require 功能
}

// JSEngine JavaScript 引擎
type JSEngine struct {
	vm     *goja.Runtime
	mu     sync.RWMutex
	config EngineConfig
}

// DefaultConfig 返回默认配置
func DefaultConfig() EngineConfig {
	return EngineConfig{
		AutoInjectMethods: true,
		WhiteList:        []string{}, // 默认为空，加载所有模块
		BlackList:        []string{}, // 默认为空，不跳过任何模块
		FailFast:        false,    // 默认为 false，模块加载失败时跳过继续
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
	JSFunc      goja.Callable
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

func (r *MethodRegistry) OverrideMethod(name string, jsFunc goja.Callable) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	method, exists := r.methods[name]
	if !exists || !method.Overridable {
		return false
	}

	method.Overridden = true
	method.JSFunc = jsFunc
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
	method.JSFunc = nil
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

func (r *MethodRegistry) ExportMethodsJSObject() string {
	r.mu.RLock()
	defer r.mu.RUnlock()

	jsObject := "const AutoGo = {\n"
	jsObject += "  methods: {\n"

	for name, method := range r.methods {
		jsObject += fmt.Sprintf("    \"%s\": {\n", name)
		jsObject += fmt.Sprintf("      description: \"%s\",\n", method.Description)
		jsObject += fmt.Sprintf("      overridable: %t,\n", method.Overridable)
		jsObject += fmt.Sprintf("      overridden: %t,\n", method.Overridden)
		jsObject += "    },\n"
	}

	jsObject += "  },\n"
	jsObject += "};\n"

	return jsObject
}

func (r *MethodRegistry) GenerateJSDocumentation() string {
	r.mu.RLock()
	defer r.mu.RUnlock()

	doc := "// AutoGo JavaScript API 文档\n"
	doc += "// 自动生成时间: " + getTimestamp() + "\n\n"

	doc += "--------------------------------------------------\n"
	doc += "// 应用管理 (app)\n"
	doc += "--------------------------------------------------\n"
	doc += "app.currentPackage() -> string\n"
	doc += "  获取当前页面应用包名\n\n"
	doc += "app.currentActivity() -> string\n"
	doc += "  获取当前页面应用类名\n\n"
	doc += "app.launch(packageName: string) -> boolean\n"
	doc += "  通过应用包名启动应用\n\n"
	doc += "app.openAppSetting(packageName: string) -> boolean\n"
	doc += "  打开应用的详情页(设置页)\n\n"
	doc += "app.viewFile(path: string)\n"
	doc += "  用其他应用查看文件\n\n"
	doc += "app.editFile(path: string)\n"
	doc += "  用其他应用编辑文件\n\n"
	doc += "app.uninstall(packageName: string)\n"
	doc += "  卸载应用\n\n"
	doc += "app.install(path: string)\n"
	doc += "  安装应用\n\n"
	doc += "app.isInstalled(packageName: string) -> boolean\n"
	doc += "  判断是否已经安装某个应用\n\n"
	doc += "app.clear(packageName: string)\n"
	doc += "  清除应用数据\n\n"
	doc += "app.forceStop(packageName: string)\n"
	doc += "  强制停止应用\n\n"
	doc += "app.disable(packageName: string)\n"
	doc += "  禁用应用\n\n"
	doc += "app.ignoreBattOpt(packageName: string)\n"
	doc += "  忽略电池优化\n\n"
	doc += "app.enable(packageName: string)\n"
	doc += "  启用应用\n\n"
	doc += "app.getBrowserPackage() -> string\n"
	doc += "  获取系统默认浏览器包名\n\n"
	doc += "app.openUrl(url: string)\n"
	doc += "  用浏览器打开网站url\n\n"

	doc += "--------------------------------------------------\n"
	doc += "// 设备管理 (device)\n"
	doc += "--------------------------------------------------\n"
	doc += "device.width -> number (只读)\n"
	doc += "  设备分辨率宽度\n\n"
	doc += "device.height -> number (只读)\n"
	doc += "  设备分辨率高度\n\n"
	doc += "device.sdkInt -> number (只读)\n"
	doc += "  安卓系统API版本\n\n"
	doc += "device.cpuAbi -> string (只读)\n"
	doc += "  设备的CPU架构\n\n"
	doc += "device.getImei() -> string\n"
	doc += "  返回设备的IMEI\n\n"
	doc += "device.getAndroidId() -> string\n"
	doc += "  返回设备的Android ID\n\n"
	doc += "device.getWifiMac() -> string\n"
	doc += "  获取设备WIFI-MAC\n\n"
	doc += "device.getWlanMac() -> string\n"
	doc += "  获取设备以太网MAC\n\n"
	doc += "device.getIp() -> string\n"
	doc += "  获取设备局域网IP地址\n\n"
	doc += "device.getBrightness() -> string\n"
	doc += "  返回当前的(手动)亮度\n\n"
	doc += "device.getBrightnessMode() -> string\n"
	doc += "  返回当前亮度模式\n\n"
	doc += "device.getMusicVolume() -> number\n"
	doc += "  返回当前媒体音量\n\n"
	doc += "device.getNotificationVolume() -> number\n"
	doc += "  返回当前通知音量\n\n"
	doc += "device.getAlarmVolume() -> number\n"
	doc += "  返回当前闹钟音量\n\n"
	doc += "device.getMusicMaxVolume() -> number\n"
	doc += "  返回媒体音量的最大值\n\n"
	doc += "device.getNotificationMaxVolume() -> number\n"
	doc += "  返回通知音量的最大值\n\n"
	doc += "device.getAlarmMaxVolume() -> number\n"
	doc += "  返回闹钟音量的最大值\n\n"
	doc += "device.setMusicVolume(volume: number)\n"
	doc += "  设置当前媒体音量\n\n"
	doc += "device.setNotificationVolume(volume: number)\n"
	doc += "  设置当前通知音量\n\n"
	doc += "device.setAlarmVolume(volume: number)\n"
	doc += "  设置当前闹钟音量\n\n"
	doc += "device.getBattery() -> number\n"
	doc += "  返回当前电量百分比\n\n"
	doc += "device.getBatteryStatus() -> number\n"
	doc += "  获取电池状态\n\n"
	doc += "device.setBatteryStatus(value: number)\n"
	doc += "  模拟电池状态\n\n"
	doc += "device.setBatteryLevel(value: number)\n"
	doc += "  模拟电池电量百分百\n\n"
	doc += "device.getTotalMem() -> number\n"
	doc += "  返回设备内存总量\n\n"
	doc += "device.getAvailMem() -> number\n"
	doc += "  返回设备当前可用的内存\n\n"
	doc += "device.isScreenOn() -> boolean\n"
	doc += "  返回设备屏幕是否是亮着的\n\n"
	doc += "device.isScreenUnlock() -> boolean\n"
	doc += "  返回屏幕锁是否已经解开\n\n"
	doc += "device.setScreenMode(mode: number)\n"
	doc += "  设置屏幕显示模式\n\n"
	doc += "device.wakeUp()\n"
	doc += "  唤醒设备\n\n"
	doc += "device.keepScreenOn()\n"
	doc += "  保持屏幕常亮\n\n"
	doc += "device.vibrate(ms: number)\n"
	doc += "  使设备震动一段时间\n\n"
	doc += "device.cancelVibration()\n"
	doc += "  如果设备处于震动状态，则取消震动\n\n"

	doc += "--------------------------------------------------\n"
	doc += "// 触摸操作 (touch)\n"
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
	doc += "// 文件操作 (files)\n"
	doc += "--------------------------------------------------\n"
	doc += "files.isDir(path: string) -> boolean\n"
	doc += "  返回路径path是否是文件夹\n\n"
	doc += "files.isFile(path: string) -> boolean\n"
	doc += "  返回路径path是否是文件\n\n"
	doc += "files.isEmptyDir(path: string) -> boolean\n"
	doc += "  返回文件夹path是否为空文件夹\n\n"
	doc += "files.create(path: string) -> boolean\n"
	doc += "  创建一个文件或文件夹\n\n"
	doc += "files.createWithDirs(path: string) -> boolean\n"
	doc += "  创建一个文件或文件夹并确保所在文件夹存在\n\n"
	doc += "files.exists(path: string) -> boolean\n"
	doc += "  返回在路径path处的文件是否存在\n\n"
	doc += "files.ensureDir(path: string) -> boolean\n"
	doc += "  确保路径path所在的文件夹存在\n\n"
	doc += "files.read(path: string) -> string\n"
	doc += "  读取文本文件path的所有内容并返回\n\n"
	doc += "files.readBytes(path: string) -> string\n"
	doc += "  读取文件path的所有内容并返回\n\n"
	doc += "files.write(path: string, text: string)\n"
	doc += "  把text写入到文件path中\n\n"
	doc += "files.writeBytes(path: string, bytes: string)\n"
	doc += "  把bytes写入到文件path中\n\n"
	doc += "files.append(path: string, text: string)\n"
	doc += "  把text追加到文件path的末尾\n\n"
	doc += "files.appendBytes(path: string, bytes: string)\n"
	doc += "  把bytes追加到文件path的末尾\n\n"
	doc += "files.copy(fromPath: string, toPath: string) -> boolean\n"
	doc += "  复制文件\n\n"
	doc += "files.move(fromPath: string, toPath: string) -> boolean\n"
	doc += "  移动文件\n\n"
	doc += "files.rename(path: string, newName: string) -> boolean\n"
	doc += "  重命名文件\n\n"
	doc += "files.remove(path: string) -> boolean\n"
	doc += "  删除文件或文件夹\n\n"
	doc += "files.getName(path: string) -> string\n"
	doc += "  返回文件的文件名\n\n"
	doc += "files.getNameWithoutExtension(path: string) -> string\n"
	doc += "  返回不含拓展名的文件的文件名\n\n"
	doc += "files.getExtension(path: string) -> string\n"
	doc += "  返回文件的拓展名\n\n"
	doc += "files.path(relativePath: string) -> string\n"
	doc += "  返回相对路径对应的绝对路径\n\n"
	doc += "files.listDir(path: string) -> Array\n"
	doc += "  列出文件夹path下的所有文件和文件夹\n\n"

	doc += "--------------------------------------------------\n"
	doc += "// 图像处理 (images)\n"
	doc += "--------------------------------------------------\n"
	doc += "images.pixel(x: number, y: number) -> string\n"
	doc += "  获取指定坐标的像素颜色\n\n"
	doc += "images.captureScreen(x1: number, y1: number, x2: number, y2: number) -> object\n"
	doc += "  截取屏幕\n\n"
	doc += "images.cmpColor(x: number, y: number, colorStr: string, sim: number) -> boolean\n"
	doc += "  比较颜色\n\n"
	doc += "images.findColor(x1: number, y1: number, x2: number, y2: number, colorStr: string, sim: number, dir: number) -> number, number\n"
	doc += "  查找颜色\n\n"
	doc += "images.getColorCountInRegion(x1: number, y1: number, x2: number, y2: number, colorStr: string, sim: number) -> number\n"
	doc += "  获取区域内指定颜色的数量\n\n"
	doc += "images.detectsMultiColors(colors: string, sim: number) -> boolean\n"
	doc += "  检测多点颜色\n\n"
	doc += "images.findMultiColors(x1: number, y1: number, x2: number, y2: number, colors: string, sim: number, dir: number) -> number, number\n"
	doc += "  查找多点颜色\n\n"
	doc += "images.readFromPath(path: string) -> object\n"
	doc += "  从路径读取图片\n\n"
	doc += "images.readFromUrl(url: string) -> object\n"
	doc += "  从URL读取图片\n\n"
	doc += "images.readFromBase64(base64Str: string) -> object\n"
	doc += "  从Base64读取图片\n\n"
	doc += "images.readFromBytes(bytes: string) -> object\n"
	doc += "  从字节数组读取图片\n\n"
	doc += "images.save(img: object, path: string, quality: number) -> boolean\n"
	doc += "  保存图片\n\n"
	doc += "images.encodeToBase64(img: object, format: string, quality: number) -> string\n"
	doc += "  编码为Base64\n\n"
	doc += "images.encodeToBytes(img: object, format: string, quality: number) -> string\n"
	doc += "  编码为字节数组\n\n"
	doc += "images.clip(img: object, x1: number, y1: number, x2: number, y2: number) -> object\n"
	doc += "  裁剪图片\n\n"
	doc += "images.resize(img: object, width: number, height: number) -> object\n"
	doc += "  调整图片大小\n\n"
	doc += "images.rotate(img: object, degree: number) -> object\n"
	doc += "  旋转图片\n\n"
	doc += "images.grayscale(img: object) -> object\n"
	doc += "  灰度化\n\n"
	doc += "images.applyThreshold(img: object, threshold: number, maxVal: number, typ: string) -> object\n"
	doc += "  应用阈值\n\n"
	doc += "images.applyAdaptiveThreshold(img: object, maxValue: number, adaptiveMethod: string, thresholdType: string, blockSize: number, C: number) -> object\n"
	doc += "  应用自适应阈值\n\n"
	doc += "images.applyBinarization(img: object, threshold: number) -> object\n"
	doc += "  二值化\n\n"

	doc += "--------------------------------------------------\n"
	doc += "// 存储管理 (storages)\n"
	doc += "--------------------------------------------------\n"
	doc += "storages.get(table: string, key: string) -> string\n"
	doc += "  从指定表中获取键值\n\n"
	doc += "storages.put(table: string, key: string, value: string)\n"
	doc += "  写入键值对\n\n"
	doc += "storages.remove(table: string, key: string)\n"
	doc += "  删除指定键\n\n"
	doc += "storages.contains(table: string, key: string) -> boolean\n"
	doc += "  判断键是否存在\n\n"
	doc += "storages.getAll(table: string) -> object\n"
	doc += "  获取所有键值对\n\n"
	doc += "storages.clear(table: string)\n"
	doc += "  清空指定表数据\n\n"

	doc += "--------------------------------------------------\n"
	doc += "// 系统管理 (system)\n"
	doc += "--------------------------------------------------\n"
	doc += "system.getPid(processName: string) -> number\n"
	doc += "  获取进程ID\n\n"
	doc += "system.getMemoryUsage(pid: number) -> number\n"
	doc += "  获取内存使用\n\n"
	doc += "system.getCpuUsage(pid: number) -> number\n"
	doc += "  获取CPU使用率\n\n"
	doc += "system.restartSelf()\n"
	doc += "  重启自身\n\n"
	doc += "system.setBootStart(enable: boolean)\n"
	doc += "  设置开机自启\n\n"

	doc += "--------------------------------------------------\n"
	doc += "// 网络请求 (http)\n"
	doc += "--------------------------------------------------\n"
	doc += "http.get(url: string, timeout: number) -> number, string\n"
	doc += "  发送GET请求\n\n"
	doc += "http.postMultipart(url: string, fileName: string, fileData: string) -> number, string\n"
	doc += "  发送Multipart POST请求\n\n"

	doc += "--------------------------------------------------\n"
	doc += "// 媒体管理 (media)\n"
	doc += "--------------------------------------------------\n"
	doc += "media.scanFile(path: string)\n"
	doc += "  扫描路径path的媒体文件\n\n"

	doc += "--------------------------------------------------\n"
	doc += "// 图像识别 (opencv)\n"
	doc += "--------------------------------------------------\n"
	doc += "opencv.findImage(x1: number, y1: number, x2: number, y2: number, template: string, isGray: boolean, scalingFactor: number, sim: number) -> number, number\n"
	doc += "  在指定区域内查找匹配的图片模板\n\n"

	doc += "--------------------------------------------------\n"
	doc += "// 文字识别 (ppocr)\n"
	doc += "--------------------------------------------------\n"
	doc += "ppocr.ocr(x1: number, y1: number, x2: number, y2: number, colorStr: string) -> Array\n"
	doc += "  识别屏幕文字\n\n"
	doc += "ppocr.ocrFromBase64(b64: string, colorStr: string) -> Array\n"
	doc += "  识别Base64图片文字\n\n"
	doc += "ppocr.ocrFromPath(path: string, colorStr: string) -> Array\n"
	doc += "  识别文件图片文字\n\n"

	doc += "--------------------------------------------------\n"
	doc += "// 方法管理 (method)\n"
	doc += "--------------------------------------------------\n"
	doc += "registerMethod(name: string, description: string, goFunc: function, overridable: boolean)\n"
	doc += "  注册新方法到JavaScript引擎\n\n"
	doc += "unregisterMethod(name: string) -> boolean\n"
	doc += "  从JavaScript引擎中移除方法\n\n"
	doc += "listMethods() -> Array\n"
	doc += "  列出所有已注册的方法\n\n"
	doc += "overrideMethod(name: string, jsFunc: function) -> boolean\n"
	doc += "  用JavaScript函数重写已注册的方法\n\n"
	doc += "restoreMethod(name: string) -> boolean\n"
	doc += "  恢复被重写的方法\n\n"

	return doc
}

func getTimestamp() string {
	return "2026-03-13"
}
