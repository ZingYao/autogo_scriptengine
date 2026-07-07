//go:build ignore
// +build ignore

package image

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"image"
	"image/draw"
	"io"
	"net/http"
	"strings"

	"github.com/Dasongzi1366/AutoGo/images"
	"github.com/Dasongzi1366/AutoGo/ppocr"
	"github.com/Dasongzi1366/AutoGo/yolo"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"

	lua "github.com/yuin/gopher-lua"
)

// ImageModule image 模块（懒人精灵兼容）
type ImageModule struct {
	ppocrEngine *ppocr.Ppocr
	yoloEngine  *yolo.Yolo
}

// Name 返回模块名称
func (m *ImageModule) Name() string {
	return "image"
}

// IsAvailable 检查模块是否可用
func (m *ImageModule) IsAvailable() bool {
	return true
}

// Register 向引擎注册方法
func (m *ImageModule) Register(engine model.Engine) error {
	state := engine.GetState()

	// 初始化 PPOCR 引擎
	m.ppocrEngine = ppocr.New("v2")

	// 创建 LuaEngine 表
	luaEngineTable := state.NewTable()

	// 注册 LuaEngine.snapShot - java层截图方法
	luaEngineTable.RawSetString("snapShot", state.NewFunction(func(L *lua.LState) int {
		l := L.CheckInt(1)
		t := L.CheckInt(2)
		r := L.CheckInt(3)
		b := L.CheckInt(4)

		// 使用 AutoGo 的 CaptureScreen 方法
		img := images.CaptureScreen(l, t, r, b, 0)
		if img == nil {
			L.Push(lua.LNil)
			return 1
		}

		// 将 Go 的 image.NRGBA 转换为 Lua 的 userdata
		ud := L.NewUserData()
		ud.Value = img
		L.Push(ud)
		return 1
	}))

	// 注册 LuaEngine.releaseBmp - 释放 bitmap 资源
	luaEngineTable.RawSetString("releaseBmp", state.NewFunction(func(L *lua.LState) int {
		// 获取 bitmap userdata
		ud := L.CheckUserData(1)
		if ud == nil {
			return 0
		}

		// 释放 image.NRGBA 资源
		// Go 的垃圾回收会自动处理，这里不需要手动释放
		return 0
	}))

	// 注册 LuaEngine 表到全局
	state.SetGlobal("LuaEngine", luaEngineTable)

	// 注册 YoloV5 表
	yoloV5Table := state.NewTable()

	// 注册 loadModel 方法
	yoloV5Table.RawSetString("loadModel", state.NewFunction(func(L *lua.LState) int {
		paramPath := L.CheckString(1)
		binPath := L.CheckString(2)
		labels := L.CheckString(3)

		// 使用 AutoGo 的 yolo.New 方法，支持 v5 和 v8
		// 默认使用 v5，CPU 线程数设为 4
		m.yoloEngine = yolo.New("v5", 4, paramPath, binPath, labels)
		if m.yoloEngine == nil {
			L.Push(lua.LBool(false))
			return 1
		}

		L.Push(lua.LBool(true))
		return 1
	}))

	// 注册 detect 方法
	yoloV5Table.RawSetString("detect", state.NewFunction(func(L *lua.LState) int {
		// 获取 bitmap userdata
		ud := L.CheckUserData(1)
		if ud == nil {
			L.Push(lua.LNil)
			return 1
		}

		// 获取 image.NRGBA 对象
		img, ok := ud.Value.(*image.NRGBA)
		if !ok || img == nil {
			L.Push(lua.LNil)
			return 1
		}

		// 检查 yolo 引擎是否已初始化
		if m.yoloEngine == nil {
			L.Push(lua.LNil)
			return 1
		}

		// 使用 AutoGo 的 yolo.DetectFromImage 方法
		results := m.yoloEngine.DetectFromImage(img)
		if results == nil {
			L.Push(lua.LNil)
			return 1
		}

		// 将结果转换为 JSON 字符串
		jsonData, err := json.Marshal(results)
		if err != nil {
			L.Push(lua.LNil)
			return 1
		}

		L.Push(lua.LString(string(jsonData)))
		return 1
	}))

	// 注册 close 方法
	yoloV5Table.RawSetString("close", state.NewFunction(func(L *lua.LState) int {
		// 关闭 YOLO 模型实例，释放相关资源
		if m.yoloEngine != nil {
			m.yoloEngine.Close()
			m.yoloEngine = nil
		}
		return 0
	}))

	state.SetGlobal("YoloV5", yoloV5Table)

	// 注册 PaddleOcr 表
	paddleOcrTable := state.NewTable()

	// 注册 loadModel 方法
	paddleOcrTable.RawSetString("loadModel", state.NewFunction(func(L *lua.LState) int {
		// AutoGo 使用 ppocr.New 来初始化引擎
		// 参数 isUseOnnxModel 不支持，默认使用 v2 模型
		L.Push(lua.LBool(true))
		return 1
	}))

	// 注册 detect 方法
	paddleOcrTable.RawSetString("detect", state.NewFunction(func(L *lua.LState) int {
		// 获取 bitmap userdata
		ud := L.CheckUserData(1)
		if ud == nil {
			L.Push(lua.LNil)
			return 1
		}

		// 获取 image.NRGBA 对象
		img, ok := ud.Value.(*image.NRGBA)
		if !ok || img == nil {
			L.Push(lua.LNil)
			return 1
		}

		// 使用 AutoGo 的 ppocr.OcrFromImage 方法
		if m.ppocrEngine == nil {
			m.ppocrEngine = ppocr.New("v2")
		}

		results := m.ppocrEngine.OcrFromImage(img, "")
		if results == nil {
			L.Push(lua.LNil)
			return 1
		}

		// 将结果转换为 JSON 字符串
		jsonData, err := json.Marshal(results)
		if err != nil {
			L.Push(lua.LNil)
			return 1
		}

		L.Push(lua.LString(string(jsonData)))
		return 1
	}))

	// 注册 loadOnnxModel 方法（空操作）
	paddleOcrTable.RawSetString("loadOnnxModel", state.NewFunction(func(L *lua.LState) int {
		// AutoGo 不支持自定义 ONNX 模型
		// 返回空字符串
		L.Push(lua.LString(""))
		return 1
	}))

	// 注册 loadNnccModel 方法（空操作）
	paddleOcrTable.RawSetString("loadNnccModel", state.NewFunction(func(L *lua.LState) int {
		// AutoGo 不支持自定义 NCNN 模型
		// 返回空字符串
		L.Push(lua.LString(""))
		return 1
	}))

	// 注册 detectWithPadding 方法
	paddleOcrTable.RawSetString("detectWithPadding", state.NewFunction(func(L *lua.LState) int {
		// 获取 bitmap userdata
		ud := L.CheckUserData(1)
		if ud == nil {
			L.Push(lua.LNil)
			return 1
		}

		// 获取 image.NRGBA 对象
		img, ok := ud.Value.(*image.NRGBA)
		if !ok || img == nil {
			L.Push(lua.LNil)
			return 1
		}

		// 获取 padding 参数
		padding := L.OptInt(2, 0)
		_ = padding // 暂不支持 padding

		// 使用 AutoGo 的 ppocr.OcrFromImage 方法
		if m.ppocrEngine == nil {
			m.ppocrEngine = ppocr.New("v2")
		}

		results := m.ppocrEngine.OcrFromImage(img, "")
		if results == nil {
			L.Push(lua.LNil)
			return 1
		}

		// 将结果转换为 JSON 字符串
		jsonData, err := json.Marshal(results)
		if err != nil {
			L.Push(lua.LNil)
			return 1
		}

		L.Push(lua.LString(string(jsonData)))
		return 1
	}))

	// 注册 PaddleOcr 表到全局
	state.SetGlobal("PaddleOcr", paddleOcrTable)

	// 注册 images 包的辅助函数

	// 注册 CaptureScreen - 截取屏幕的指定区域
	state.SetGlobal("CaptureScreen", state.NewFunction(func(L *lua.LState) int {
		x1 := L.CheckInt(1)
		y1 := L.CheckInt(2)
		x2 := L.CheckInt(3)
		y2 := L.CheckInt(4)

		// 使用 AutoGo 的 CaptureScreen 方法
		img := images.CaptureScreen(x1, y1, x2, y2, 0)
		if img == nil {
			L.Push(lua.LNil)
			return 1
		}

		// 将 Go 的 image.NRGBA 转换为 Lua 的 userdata
		ud := L.NewUserData()
		ud.Value = img
		L.Push(ud)
		return 1
	}))

	// 注册 Pixel - 获取指定坐标点的颜色值
	state.SetGlobal("Pixel", state.NewFunction(func(L *lua.LState) int {
		x := L.CheckInt(1)
		y := L.CheckInt(2)

		// 使用 AutoGo 的 Pixel 方法
		color := images.Pixel(x, y, 0)
		L.Push(lua.LString(color))
		return 1
	}))

	// 注册 CmpColor - 比较指定坐标点 (x, y) 的颜色
	state.SetGlobal("CmpColor", state.NewFunction(func(L *lua.LState) int {
		x := L.CheckInt(1)
		y := L.CheckInt(2)
		colorStr := L.CheckString(3)
		sim := L.CheckNumber(4)

		// 使用 AutoGo 的 CmpColor 方法
		result := images.CmpColor(x, y, colorStr, float32(sim), 0)
		L.Push(lua.LBool(result))
		return 1
	}))

	// 注册 FindColor - 在指定区域内查找目标颜色
	state.SetGlobal("FindColor", state.NewFunction(func(L *lua.LState) int {
		x1 := L.CheckInt(1)
		y1 := L.CheckInt(2)
		x2 := L.CheckInt(3)
		y2 := L.CheckInt(4)
		colorStr := L.CheckString(5)
		sim := L.CheckNumber(6)
		dir := L.CheckInt(7)

		// 使用 AutoGo 的 FindColor 方法
		x, y := images.FindColor(x1, y1, x2, y2, colorStr, float32(sim), dir, 0)
		L.Push(lua.LNumber(x))
		L.Push(lua.LNumber(y))
		return 2
	}))

	// 注册 GetColorCountInRegion - 计算指定区域内符合颜色条件的像素数量
	state.SetGlobal("GetColorCountInRegion", state.NewFunction(func(L *lua.LState) int {
		x1 := L.CheckInt(1)
		y1 := L.CheckInt(2)
		x2 := L.CheckInt(3)
		y2 := L.CheckInt(4)
		colorStr := L.CheckString(5)
		sim := L.CheckNumber(6)

		// 使用 AutoGo 的 GetColorCountInRegion 方法
		count := images.GetColorCountInRegion(x1, y1, x2, y2, colorStr, float32(sim), 0)
		L.Push(lua.LNumber(count))
		return 1
	}))

	// 注册 DetectsMultiColors - 根据指定的颜色串信息在屏幕进行多点颜色比对
	state.SetGlobal("DetectsMultiColors", state.NewFunction(func(L *lua.LState) int {
		colors := L.CheckString(1)
		sim := L.CheckNumber(2)

		// 使用 AutoGo 的 DetectsMultiColors 方法
		result := images.DetectsMultiColors(colors, float32(sim), 0)
		L.Push(lua.LBool(result))
		return 1
	}))

	// 注册 FindMultiColors - 在指定区域内查找匹配的多点颜色序列
	state.SetGlobal("FindMultiColors", state.NewFunction(func(L *lua.LState) int {
		x1 := L.CheckInt(1)
		y1 := L.CheckInt(2)
		x2 := L.CheckInt(3)
		y2 := L.CheckInt(4)
		colors := L.CheckString(5)
		sim := L.CheckNumber(6)
		dir := L.CheckInt(7)

		// 使用 AutoGo 的 FindMultiColors 方法
		x, y := images.FindMultiColors(x1, y1, x2, y2, colors, float32(sim), dir, 0)
		L.Push(lua.LNumber(x))
		L.Push(lua.LNumber(y))
		return 2
	}))

	// 注册 ReadFromPath - 读取在路径path的图片文件
	state.SetGlobal("ReadFromPath", state.NewFunction(func(L *lua.LState) int {
		path := L.CheckString(1)

		// 使用 AutoGo 的 ReadFromPath 方法
		img := images.ReadFromPath(path)
		if img == nil {
			L.Push(lua.LNil)
			return 1
		}

		// 将 Go 的 image.NRGBA 转换为 Lua 的 userdata
		ud := L.NewUserData()
		ud.Value = img
		L.Push(ud)
		return 1
	}))

	// 注册 ReadFromUrl - 加载地址URL的网络图片
	state.SetGlobal("ReadFromUrl", state.NewFunction(func(L *lua.LState) int {
		url := L.CheckString(1)

		// 使用 Go 标准库下载图片
		// AutoGo 的 ReadFromUrl 函数返回 nil，需要自己实现
		resp, err := http.Get(url)
		if err != nil {
			L.Push(lua.LNil)
			return 1
		}
		defer resp.Body.Close()

		// 读取图片数据
		data, err := io.ReadAll(resp.Body)
		if err != nil {
			L.Push(lua.LNil)
			return 1
		}

		// 使用 Go 标准库解码图片
		img, _, err := image.Decode(bytes.NewReader(data))
		if err != nil {
			L.Push(lua.LNil)
			return 1
		}

		// 转换为 image.NRGBA
		bounds := img.Bounds()
		nrgba := image.NewNRGBA(bounds)
		draw.Draw(nrgba, bounds, img, bounds.Min, draw.Src)

		// 将 Go 的 image.NRGBA 转换为 Lua 的 userdata
		ud := L.NewUserData()
		ud.Value = nrgba
		L.Push(ud)
		return 1
	}))

	// 注册 ReadFromBase64 - 解码Base64数据并返回解码后的图片
	state.SetGlobal("ReadFromBase64", state.NewFunction(func(L *lua.LState) int {
		base64Str := L.CheckString(1)

		// 使用 Go 标准库解码 Base64
		// AutoGo 的 ReadFromBase64 和 ReadFromBytes 函数都返回 nil，需要自己实现
		decoder := base64.NewDecoder(base64.StdEncoding, strings.NewReader(base64Str))
		data, err := io.ReadAll(decoder)
		if err != nil {
			L.Push(lua.LNil)
			return 1
		}

		// 使用 Go 标准库解码图片
		img, _, err := image.Decode(bytes.NewReader(data))
		if err != nil {
			L.Push(lua.LNil)
			return 1
		}

		// 转换为 image.NRGBA
		bounds := img.Bounds()
		nrgba := image.NewNRGBA(bounds)
		draw.Draw(nrgba, bounds, img, bounds.Min, draw.Src)

		// 将 Go 的 image.NRGBA 转换为 Lua 的 userdata
		ud := L.NewUserData()
		ud.Value = nrgba
		L.Push(ud)
		return 1
	}))

	// 注册 EncodeToBase64 - 把Image对象编码为Base64数据并返回
	state.SetGlobal("EncodeToBase64", state.NewFunction(func(L *lua.LState) int {
		// 获取 bitmap userdata
		ud := L.CheckUserData(1)
		if ud == nil {
			L.Push(lua.LString(""))
			return 1
		}

		// 获取 image.NRGBA 对象
		img, ok := ud.Value.(*image.NRGBA)
		if !ok || img == nil {
			L.Push(lua.LString(""))
			return 1
		}

		// 获取格式和质量参数
		fromat := L.OptString(2, "png")
		quality := L.OptInt(3, 90)

		// 使用 AutoGo 的 EncodeToBase64 方法
		base64Str := images.EncodeToBase64(img, fromat, quality)
		L.Push(lua.LString(base64Str))
		return 1
	}))

	// 注册 Save - 把图片image保存到path中
	state.SetGlobal("Save", state.NewFunction(func(L *lua.LState) int {
		// 获取 bitmap userdata
		ud := L.CheckUserData(1)
		if ud == nil {
			L.Push(lua.LBool(false))
			return 1
		}

		// 获取 image.NRGBA 对象
		img, ok := ud.Value.(*image.NRGBA)
		if !ok || img == nil {
			L.Push(lua.LBool(false))
			return 1
		}

		// 获取路径和质量参数
		path := L.CheckString(2)
		quality := L.OptInt(3, 90)

		// 使用 AutoGo 的 Save 方法
		result := images.Save(img, path, quality)
		L.Push(lua.LBool(result))
		return 1
	}))

	// 注册 Clip - 从图片img的位置(x1, y1)处剪切至(x2, y2)区域
	state.SetGlobal("Clip", state.NewFunction(func(L *lua.LState) int {
		// 获取 bitmap userdata
		ud := L.CheckUserData(1)
		if ud == nil {
			L.Push(lua.LNil)
			return 1
		}

		// 获取 image.NRGBA 对象
		img, ok := ud.Value.(*image.NRGBA)
		if !ok || img == nil {
			L.Push(lua.LNil)
			return 1
		}

		// 获取剪切区域参数
		x1 := L.CheckInt(2)
		y1 := L.CheckInt(3)
		x2 := L.CheckInt(4)
		y2 := L.CheckInt(5)

		// 使用 AutoGo 的 Clip 方法
		clippedImg := images.Clip(img, x1, y1, x2, y2)
		if clippedImg == nil {
			L.Push(lua.LNil)
			return 1
		}

		// 将 Go 的 image.NRGBA 转换为 Lua 的 userdata
		ud2 := L.NewUserData()
		ud2.Value = clippedImg
		L.Push(ud2)
		return 1
	}))

	// 注册 Resize - 调整图片大小
	state.SetGlobal("Resize", state.NewFunction(func(L *lua.LState) int {
		// 获取 bitmap userdata
		ud := L.CheckUserData(1)
		if ud == nil {
			L.Push(lua.LNil)
			return 1
		}

		// 获取 image.NRGBA 对象
		img, ok := ud.Value.(*image.NRGBA)
		if !ok || img == nil {
			L.Push(lua.LNil)
			return 1
		}

		// 获取宽高参数
		width := L.CheckInt(2)
		height := L.CheckInt(3)

		// 使用 AutoGo 的 Resize 方法
		resizedImg := images.Resize(img, width, height)
		if resizedImg == nil {
			L.Push(lua.LNil)
			return 1
		}

		// 将 Go 的 image.NRGBA 转换为 Lua 的 userdata
		ud2 := L.NewUserData()
		ud2.Value = resizedImg
		L.Push(ud2)
		return 1
	}))

	// 注册 Rotate - 将图片顺时针旋转degree度
	state.SetGlobal("Rotate", state.NewFunction(func(L *lua.LState) int {
		// 获取 bitmap userdata
		ud := L.CheckUserData(1)
		if ud == nil {
			L.Push(lua.LNil)
			return 1
		}

		// 获取 image.NRGBA 对象
		img, ok := ud.Value.(*image.NRGBA)
		if !ok || img == nil {
			L.Push(lua.LNil)
			return 1
		}

		// 获取旋转角度参数
		degree := L.CheckInt(2)

		// 使用 AutoGo 的 Rotate 方法
		rotatedImg := images.Rotate(img, degree)
		if rotatedImg == nil {
			L.Push(lua.LNil)
			return 1
		}

		// 将 Go 的 image.NRGBA 转换为 Lua 的 userdata
		ud2 := L.NewUserData()
		ud2.Value = rotatedImg
		L.Push(ud2)
		return 1
	}))

	// 注册 Grayscale - 将图片灰度化
	state.SetGlobal("Grayscale", state.NewFunction(func(L *lua.LState) int {
		// 获取 bitmap userdata
		ud := L.CheckUserData(1)
		if ud == nil {
			L.Push(lua.LNil)
			return 1
		}

		// 获取 image.NRGBA 对象
		img, ok := ud.Value.(*image.NRGBA)
		if !ok || img == nil {
			L.Push(lua.LNil)
			return 1
		}

		// 使用 AutoGo 的 Grayscale 方法
		grayImg := images.Grayscale(img)
		if grayImg == nil {
			L.Push(lua.LNil)
			return 1
		}

		// 将 Go 的 image.Gray 转换为 Lua 的 userdata
		ud2 := L.NewUserData()
		ud2.Value = grayImg
		L.Push(ud2)
		return 1
	}))

	// 注册 ApplyThreshold - 将图片阈值化
	state.SetGlobal("ApplyThreshold", state.NewFunction(func(L *lua.LState) int {
		// 获取 bitmap userdata
		ud := L.CheckUserData(1)
		if ud == nil {
			L.Push(lua.LNil)
			return 1
		}

		// 获取 image.NRGBA 对象
		img, ok := ud.Value.(*image.NRGBA)
		if !ok || img == nil {
			L.Push(lua.LNil)
			return 1
		}

		// 获取阈值参数
		threshold := L.CheckInt(2)
		maxVal := L.CheckInt(3)
		typ := L.CheckString(4)

		// 使用 AutoGo 的 ApplyThreshold 方法
		thresholdImg := images.ApplyThreshold(img, threshold, maxVal, typ)
		if thresholdImg == nil {
			L.Push(lua.LNil)
			return 1
		}

		// 将 Go 的 image.Gray 转换为 Lua 的 userdata
		ud2 := L.NewUserData()
		ud2.Value = thresholdImg
		L.Push(ud2)
		return 1
	}))

	// 注册 ApplyAdaptiveThreshold - 将图像进行自适应阈值化处理
	state.SetGlobal("ApplyAdaptiveThreshold", state.NewFunction(func(L *lua.LState) int {
		// 获取 bitmap userdata
		ud := L.CheckUserData(1)
		if ud == nil {
			L.Push(lua.LNil)
			return 1
		}

		// 获取 image.NRGBA 对象
		img, ok := ud.Value.(*image.NRGBA)
		if !ok || img == nil {
			L.Push(lua.LNil)
			return 1
		}

		// 获取自适应阈值参数
		maxValue := float64(L.CheckNumber(2))
		adaptiveMethod := L.CheckString(3)
		thresholdType := L.CheckString(4)
		blockSize := L.CheckInt(5)
		C := float64(L.CheckNumber(6))

		// 使用 AutoGo 的 ApplyAdaptiveThreshold 方法
		thresholdImg := images.ApplyAdaptiveThreshold(img, maxValue, adaptiveMethod, thresholdType, blockSize, C)
		if thresholdImg == nil {
			L.Push(lua.LNil)
			return 1
		}

		// 将 Go 的 image.Gray 转换为 Lua 的 userdata
		ud2 := L.NewUserData()
		ud2.Value = thresholdImg
		L.Push(ud2)
		return 1
	}))

	// 注册 ApplyBinarization - 将图像进行二值化处理
	state.SetGlobal("ApplyBinarization", state.NewFunction(func(L *lua.LState) int {
		// 获取 bitmap userdata
		ud := L.CheckUserData(1)
		if ud == nil {
			L.Push(lua.LNil)
			return 1
		}

		// 获取 image.NRGBA 对象
		img, ok := ud.Value.(*image.NRGBA)
		if !ok || img == nil {
			L.Push(lua.LNil)
			return 1
		}

		// 获取阈值参数
		threshold := L.CheckInt(2)

		// 使用 AutoGo 的 ApplyBinarization 方法
		binaryImg := images.ApplyBinarization(img, threshold)
		if binaryImg == nil {
			L.Push(lua.LNil)
			return 1
		}

		// 将 Go 的 image.Gray 转换为 Lua 的 userdata
		ud2 := L.NewUserData()
		ud2.Value = binaryImg
		L.Push(ud2)
		return 1
	}))

	return nil
}

// Close 释放模块资源
func (m *ImageModule) Close() {
	if m.ppocrEngine != nil {
		m.ppocrEngine.Close()
		m.ppocrEngine = nil
	}
}
