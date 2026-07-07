//go:build ignore
// +build ignore

package extension

import (
	"crypto/md5"
	"database/sql"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/Dasongzi1366/AutoGo/files"
	"github.com/Dasongzi1366/AutoGo/images"
	"github.com/Dasongzi1366/AutoGo/plugin"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"

	lua "github.com/yuin/gopher-lua"
)

// ExtensionModule extension 模块（懒人精灵兼容）
type ExtensionModule struct {
	threads        map[int]*ThreadInfo
	threadMutex    sync.Mutex
	threadCounter  int
	cache          map[string]interface{}
	cacheMutex     sync.RWMutex
	mysqlConns     map[string]*sql.DB
	mysqlMutex     sync.Mutex
	ThrowException bool // 是否抛出异常
	ShowWarning    bool // 是否显示警告信息
	Debug          bool // 是否开启调试模式（打印脚本堆栈信息）
}

// ThreadInfo 线程信息
type ThreadInfo struct {
	ID     int
	Stop   bool
	Cancel chan struct{}
}

// New 创建一个新的 ExtensionModule 实例
func New() *ExtensionModule {
	return &ExtensionModule{
		threads:        make(map[int]*ThreadInfo),
		cache:          make(map[string]interface{}),
		mysqlConns:     make(map[string]*sql.DB),
		ThrowException: false,
		ShowWarning:    true,
		Debug:          false,
	}
}

// handleUnimplementedMethod 处理未实现方法的通用逻辑
func (m *ExtensionModule) handleUnimplementedMethod(methodName string, L *lua.LState) int {
	// 如果开启了调试模式，打印脚本堆栈信息
	if m.Debug {
		fmt.Println("\n=== 调试信息: 调用未实现方法 ===")
		fmt.Printf("方法名: extension.%s\n", methodName)
		fmt.Println("=== Lua 调用堆栈 ===")
		for i := 1; ; i++ {
			dbg, ok := L.GetStack(i)
			if !ok {
				break
			}
			fmt.Printf("  [%d] 函数: %s, 行号: %d\n", i, dbg.Name, dbg.CurrentLine)
			fmt.Printf("       源文件: %s\n", dbg.Source)
			fmt.Printf("       What: %s, 调用行: %d\n", dbg.What, dbg.LineDefined)
		}
		fmt.Println("=== 堆栈结束 ===\n")
	}

	// 打印警告信息
	if m.ShowWarning {
		fmt.Printf("[警告] extension.%s 方法暂未实现\n", methodName)
	}

	// 根据配置决定是否抛出异常
	if m.ThrowException {
		L.RaiseError(fmt.Sprintf("extension.%s 方法暂未实现", methodName))
		return 0
	}

	// 默认返回 nil
	L.Push(lua.LNil)
	return 1
}

// getBaiduAccessToken 获取百度 Access Token
func (m *ExtensionModule) getBaiduAccessToken(apiKey, secretKey string) (string, error) {
	tokenURL := fmt.Sprintf("https://aip.baidubce.com/oauth/2.0/token?grant_type=client_credentials&client_id=%s&client_secret=%s", apiKey, secretKey)

	resp, err := http.Get(tokenURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return "", err
	}

	accessToken, ok := result["access_token"].(string)
	if !ok {
		return "", fmt.Errorf("failed to get access token")
	}

	return accessToken, nil
}

// callBaiduOCR 调用百度 OCR API
func (m *ExtensionModule) callBaiduOCR(apiURL, accessToken string, imageData []byte) (string, error) {
	// 将图片转换为 base64
	imageBase64 := base64.StdEncoding.EncodeToString(imageData)

	// 构造请求参数
	params := url.Values{}
	params.Set("image", imageBase64)

	// 添加 access token 到 URL
	fullURL := fmt.Sprintf("%s?access_token=%s", apiURL, accessToken)

	// 发送 POST 请求
	resp, err := http.PostForm(fullURL, params)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// Name 返回模块名称
func (m *ExtensionModule) Name() string {
	return "extension"
}

// IsAvailable 检查模块是否可用
func (m *ExtensionModule) IsAvailable() bool {
	return true
}

// Register 向引擎注册方法
func (m *ExtensionModule) Register(engine model.Engine) error {
	state := engine.GetState()
	m.threads = make(map[int]*ThreadInfo)
	m.threadCounter = 0

	// 创建 extension 表
	extensionTable := state.NewTable()
	state.SetGlobal("extension", extensionTable)

	// ========== 加密编码类方法 ==========

	// 注册 extension.md5 - 字符串的 MD5
	extensionTable.RawSetString("md5", state.NewFunction(func(L *lua.LState) int {
		str := L.CheckString(1)
		hash := md5.Sum([]byte(str))
		L.Push(lua.LString(hex.EncodeToString(hash[:])))
		return 1
	}))

	// 注册 extension.fileMD5 - 获取文件的 MD5 码
	extensionTable.RawSetString("fileMD5", state.NewFunction(func(L *lua.LState) int {
		filepath := L.CheckString(1)
		md5Str := extensionFileMD5(filepath)
		L.Push(lua.LString(md5Str))
		return 1
	}))

	// 注册 extension.base64Encode - base64 编码
	extensionTable.RawSetString("base64Encode", state.NewFunction(func(L *lua.LState) int {
		str := L.CheckString(1)
		encoded := base64.StdEncoding.EncodeToString([]byte(str))
		L.Push(lua.LString(encoded))
		return 1
	}))

	// 注册 extension.base64Decode - base64 解码
	extensionTable.RawSetString("base64Decode", state.NewFunction(func(L *lua.LState) int {
		str := L.CheckString(1)
		decoded, err := base64.StdEncoding.DecodeString(str)
		if err != nil {
			L.Push(lua.LNil)
			return 1
		}
		L.Push(lua.LString(string(decoded)))
		return 1
	}))

	// 注册 extension.encodeUrl - url 格式编码
	extensionTable.RawSetString("encodeUrl", state.NewFunction(func(L *lua.LState) int {
		str := L.CheckString(1)
		encoded := url.QueryEscape(str)
		L.Push(lua.LString(encoded))
		return 1
	}))

	// 注册 extension.decodeUrl - url 格式解码
	extensionTable.RawSetString("decodeUrl", state.NewFunction(func(L *lua.LState) int {
		str := L.CheckString(1)
		decoded, err := url.QueryUnescape(str)
		if err != nil {
			L.Push(lua.LNil)
			return 1
		}
		L.Push(lua.LString(decoded))
		return 1
	}))

	// 注册 extension.urlEncode - url 格式编码（别名）
	extensionTable.RawSetString("urlEncode", state.NewFunction(func(L *lua.LState) int {
		str := L.CheckString(1)
		encoded := url.QueryEscape(str)
		L.Push(lua.LString(encoded))
		return 1
	}))

	// 注册 extension.urlDecode - url 格式解码（别名）
	extensionTable.RawSetString("urlDecode", state.NewFunction(func(L *lua.LState) int {
		str := L.CheckString(1)
		decoded, err := url.QueryUnescape(str)
		if err != nil {
			L.Push(lua.LNil)
			return 1
		}
		L.Push(lua.LString(decoded))
		return 1
	}))

	// 注册 extension.htmlEncode - html 编码
	extensionTable.RawSetString("htmlEncode", state.NewFunction(func(L *lua.LState) int {
		str := L.CheckString(1)
		encoded := ""
		for _, c := range str {
			switch c {
			case '<':
				encoded += "&lt;"
			case '>':
				encoded += "&gt;"
			case '&':
				encoded += "&amp;"
			case '"':
				encoded += "&quot;"
			case '\'':
				encoded += "&#39;"
			default:
				encoded += string(c)
			}
		}
		L.Push(lua.LString(encoded))
		return 1
	}))

	// 注册 extension.htmlDecode - html 解码
	extensionTable.RawSetString("htmlDecode", state.NewFunction(func(L *lua.LState) int {
		str := L.CheckString(1)
		decoded := str
		decoded = strings.ReplaceAll(decoded, "&lt;", "<")
		decoded = strings.ReplaceAll(decoded, "&gt;", ">")
		decoded = strings.ReplaceAll(decoded, "&amp;", "&")
		decoded = strings.ReplaceAll(decoded, "&quot;", "\"")
		decoded = strings.ReplaceAll(decoded, "&#39;", "'")
		L.Push(lua.LString(decoded))
		return 1
	}))

	// 注册 extension.jsonEncode - JSON 编码
	extensionTable.RawSetString("jsonEncode", state.NewFunction(func(L *lua.LState) int {
		value := L.CheckAny(1)

		// 将 Lua 值转换为 JSON
		jsonStr, err := luaValueToJson(L, value)
		if err != nil {
			L.Push(lua.LNil)
			return 1
		}
		L.Push(lua.LString(jsonStr))
		return 1
	}))

	// 注册 extension.jsonDecode - JSON 解码
	extensionTable.RawSetString("jsonDecode", state.NewFunction(func(L *lua.LState) int {
		jsonStr := L.CheckString(1)

		// 将 JSON 字符串转换为 Lua 值
		value, err := jsonToLua(L, jsonStr)
		if err != nil {
			L.Push(lua.LNil)
			return 1
		}
		L.Push(value)
		return 1
	}))

	// 注册 extension.getFileBase64 - 获取文件 base64 编码
	extensionTable.RawSetString("getFileBase64", state.NewFunction(func(L *lua.LState) int {
		filepath := L.CheckString(1)
		data := files.ReadBytes(filepath)
		if data == nil {
			L.Push(lua.LNil)
			return 1
		}
		encoded := base64.StdEncoding.EncodeToString(data)
		L.Push(lua.LString(encoded))
		return 1
	}))

	// ========== 文件操作类方法 ==========

	// 注册 extension.fileExist - 文件或文件夹是否存在
	extensionTable.RawSetString("fileExist", state.NewFunction(func(L *lua.LState) int {
		filepath := L.CheckString(1)
		exists := files.Exists(filepath)
		L.Push(lua.LBool(exists))
		return 1
	}))

	// 注册 extension.mkdir - 创建文件夹
	extensionTable.RawSetString("mkdir", state.NewFunction(func(L *lua.LState) int {
		dir := L.CheckString(1)
		success := files.Create(dir)
		L.Push(lua.LBool(success))
		return 1
	}))

	// 注册 extension.delfile - 删除文件或文件夹
	extensionTable.RawSetString("delfile", state.NewFunction(func(L *lua.LState) int {
		path := L.CheckString(1)
		success := files.Remove(path)
		L.Push(lua.LBool(success))
		return 1
	}))

	// 注册 extension.readFile - 读取文件所有内容
	extensionTable.RawSetString("readFile", state.NewFunction(func(L *lua.LState) int {
		path := L.CheckString(1)
		content := files.Read(path)
		L.Push(lua.LString(content))
		return 1
	}))

	// 注册 extension.writeFile - 写字符串到文件
	extensionTable.RawSetString("writeFile", state.NewFunction(func(L *lua.LState) int {
		path := L.CheckString(1)
		str := L.CheckString(2)
		append := false
		if L.GetTop() >= 3 {
			append = L.CheckBool(3)
		}

		if append {
			files.Append(path, str)
		} else {
			files.Write(path, str)
		}
		L.Push(lua.LBool(true))
		return 1
	}))

	// 注册 extension.fileSize - 获取文件大小
	extensionTable.RawSetString("fileSize", state.NewFunction(func(L *lua.LState) int {
		path := L.CheckString(1)
		data := files.ReadBytes(path)
		if data == nil {
			L.Push(lua.LNumber(0))
			return 1
		}
		L.Push(lua.LNumber(len(data)))
		return 1
	}))

	// ========== 图片处理类方法 ==========

	// 注册 extension.getImage - 获取本地图片数据
	extensionTable.RawSetString("getImage", state.NewFunction(func(L *lua.LState) int {
		path := L.CheckString(1)
		img := images.ReadFromPath(path)
		if img == nil {
			L.Push(lua.LNumber(0))
			L.Push(lua.LNumber(0))
			L.Push(lua.LNil)
			return 3
		}
		L.Push(lua.LNumber(img.Bounds().Dx()))
		L.Push(lua.LNumber(img.Bounds().Dy()))
		L.Push(lua.LString("image"))
		return 3
	}))

	// 注册 extension.rotateImage - 旋转本地图片
	extensionTable.RawSetString("rotateImage", state.NewFunction(func(L *lua.LState) int {
		path := L.CheckString(1)
		rotate := L.CheckInt(2)
		img := images.ReadFromPath(path)
		if img == nil {
			L.Push(lua.LBool(false))
			return 1
		}
		rotated := images.Rotate(img, rotate)
		if rotated == nil {
			L.Push(lua.LBool(false))
			return 1
		}
		success := images.Save(rotated, path, 100)
		L.Push(lua.LBool(success))
		return 1
	}))

	// 注册 extension.scaleImage - 缩放图片
	extensionTable.RawSetString("scaleImage", state.NewFunction(func(L *lua.LState) int {
		src := L.CheckString(1)
		dst := L.CheckString(2)
		w := L.CheckInt(3)
		h := L.CheckInt(4)
		img := images.ReadFromPath(src)
		if img == nil {
			L.Push(lua.LBool(false))
			return 1
		}
		scaled := images.Resize(img, w, h)
		if scaled == nil {
			L.Push(lua.LBool(false))
			return 1
		}
		success := images.Save(scaled, dst, 100)
		L.Push(lua.LBool(success))
		return 1
	}))

	// 注册 extension.binaryImage - 二值化本地图片
	extensionTable.RawSetString("binaryImage", state.NewFunction(func(L *lua.LState) int {
		srcimage := L.CheckString(1)
		dstimage := L.CheckString(2)
		threshold := 150
		if L.GetTop() >= 3 {
			threshold = L.CheckInt(3)
		}
		_ = 255
		if L.GetTop() >= 4 {
			_ = L.CheckInt(4)
		}

		img := images.ReadFromPath(srcimage)
		if img == nil {
			L.Push(lua.LNumber(0))
			return 1
		}
		binary := images.ApplyBinarization(img, threshold)
		if binary == nil {
			L.Push(lua.LNumber(0))
			return 1
		}
		// 将 Gray 转换为 NRGBA
		nrgba := images.ToNrgba(binary)
		if nrgba == nil {
			L.Push(lua.LNumber(0))
			return 1
		}
		success := images.Save(nrgba, dstimage, 100)
		if success {
			L.Push(lua.LNumber(1))
		} else {
			L.Push(lua.LNumber(0))
		}
		return 1
	}))

	// 注册 extension.binaryRect - 区域截图并二值化
	extensionTable.RawSetString("binaryRect", state.NewFunction(func(L *lua.LState) int {
		save := L.CheckString(1)
		l := L.CheckInt(2)
		t := L.CheckInt(3)
		r := L.CheckInt(4)
		b := L.CheckInt(5)
		threshold := 150
		if L.GetTop() >= 6 {
			threshold = L.CheckInt(6)
		}
		_ = 255
		if L.GetTop() >= 7 {
			_ = L.CheckInt(7)
		}

		// 截图
		img := images.CaptureScreen(l, t, r, b)
		if img == nil {
			L.Push(lua.LNumber(0))
			return 1
		}
		// 二值化
		binary := images.ApplyBinarization(img, threshold)
		if binary == nil {
			L.Push(lua.LNumber(0))
			return 1
		}
		// 将 Gray 转换为 NRGBA
		nrgba := images.ToNrgba(binary)
		if nrgba == nil {
			L.Push(lua.LNumber(0))
			return 1
		}
		// 保存
		success := images.Save(nrgba, save, 100)
		if success {
			L.Push(lua.LNumber(1))
		} else {
			L.Push(lua.LNumber(0))
		}
		return 1
	}))

	// ========== OCR 相关方法 ==========

	// 注册 extension.qrDecode - 二维码解析
	extensionTable.RawSetString("qrDecode", state.NewFunction(func(L *lua.LState) int {
		filePath := L.CheckString(1)

		// 读取图片文件
		file, err := os.Open(filePath)
		if err != nil {
			// 文件打开失败，返回 nil
			L.Push(lua.LNil)
			return 1
		}
		defer file.Close()

		// 解码图片
		img, _, err := image.Decode(file)
		if err != nil {
			// 图片解码失败，返回 nil
			L.Push(lua.LNil)
			return 1
		}

		// 使用 gozxing 库解析二维码
		bmp, err := gozxing.NewBinaryBitmapFromImage(img)
		if err != nil {
			// 创建位图失败，返回 nil
			L.Push(lua.LNil)
			return 1
		}

		// 创建二维码读取器
		qrReader := qrcode.NewQRCodeReader()

		// 解析二维码
		result, err := qrReader.Decode(bmp, nil)
		if err != nil {
			// 解析失败，返回 nil
			L.Push(lua.LNil)
			return 1
		}

		// 返回二维码内容
		L.Push(lua.LString(result.GetText()))
		return 1
	}))

	// 注册 extension.bdOcr - 百度 ocr
	extensionTable.RawSetString("bdOcr", state.NewFunction(func(L *lua.LState) int {
		filePath := L.CheckString(1)
		apiKey := L.CheckString(2)
		secretKey := L.CheckString(3)
		ocrType := L.CheckInt(4)

		// 根据不同的 type 参数调用不同的 OCR 接口
		var apiURL string
		switch ocrType {
		case 0:
			apiURL = "https://aip.baidubce.com/rest/2.0/ocr/v1/general_basic" // 通用文字识别
		case 1:
			apiURL = "https://aip.baidubce.com/rest/2.0/ocr/v1/general" // 通用文字识别（含位置信息版）
		case 2:
			apiURL = "https://aip.baidubce.com/rest/2.0/ocr/v1/accurate_basic" // 通用文字识别（高精度版）
		case 3:
			apiURL = "https://aip.baidubce.com/rest/2.0/ocr/v1/accurate" // 通用文字识别（高精度含位置版）
		default:
			apiURL = "https://aip.baidubce.com/rest/2.0/ocr/v1/general_basic" // 默认使用通用文字识别
		}

		// 读取图片文件并转换为 base64
		imageData, err := os.ReadFile(filePath)
		if err != nil {
			L.Push(lua.LNil)
			return 1
		}

		// 获取 Access Token
		accessToken, err := m.getBaiduAccessToken(apiKey, secretKey)
		if err != nil {
			L.Push(lua.LNil)
			return 1
		}

		// 调用 OCR API
		result, err := m.callBaiduOCR(apiURL, accessToken, imageData)
		if err != nil {
			L.Push(lua.LNil)
			return 1
		}

		L.Push(lua.LString(result))
		return 1
	}))

	// ========== Plugin 相关方法 ==========

	// 创建 LuaEngine 表
	luaEngineTable := state.NewTable()

	// 注册 LuaEngine.loadApk - 加载一个 apk 插件
	luaEngineTable.RawSetString("loadApk", state.NewFunction(func(L *lua.LState) int {
		name := L.CheckString(1)
		loader := plugin.LoadApk(name)
		if loader == nil {
			L.Push(lua.LNil)
			return 1
		}
		// 返回一个包装的 Lua 表
		table := L.NewTable()
		L.Push(table)
		return 1
	}))

	// 注册 LuaEngine.getContext - 获取 android 上下文对象
	luaEngineTable.RawSetString("getContext", state.NewFunction(func(L *lua.LState) int {
		// 返回一个包装的 Lua 表
		table := L.NewTable()
		L.Push(table)
		return 1
	}))

	// 将 LuaEngine 表注册到全局
	state.SetGlobal("LuaEngine", luaEngineTable)

	// ========== import 方法（抛出异常） ==========

	// 注册 import - 加载 java 类
	state.SetGlobal("import", state.NewFunction(func(L *lua.LState) int {
		className := L.CheckString(1)
		// 抛出异常，提示用户此功能不可用
		L.RaiseError(fmt.Sprintf("import('%s') 不可用：AutoGo 不支持直接调用 Java 类。请使用 AutoGo 的 plugin 模块或 rhino 模块。", className))
		return 0
	}))

	// ========== MySQL 相关方法 ==========

	// 初始化 MySQL 连接池
	m.mysqlConns = make(map[string]*sql.DB)

	// 创建 mysql 表
	mysqlTable := state.NewTable()

	// 注册 mysql.connectSQL - 连接 mysql 数据库
	mysqlTable.RawSetString("connectSQL", state.NewFunction(func(L *lua.LState) int {
		host := L.CheckString(1)
		port := L.CheckInt(2)
		database := L.CheckString(3)
		user := "root"
		if L.GetTop() >= 4 {
			user = L.CheckString(4)
		}
		password := ""
		if L.GetTop() >= 5 {
			password = L.CheckString(5)
		}
		timeout := L.CheckInt(6)

		// 构建 DSN
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?timeout=%ds&parseTime=true", user, password, host, port, database, timeout/1000)

		// 连接数据库
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			L.Push(lua.LNil)
			L.Push(lua.LString(fmt.Sprintf("连接失败: %v", err)))
			return 2
		}

		// 测试连接
		err = db.Ping()
		if err != nil {
			L.Push(lua.LNil)
			L.Push(lua.LString(fmt.Sprintf("连接失败: %v", err)))
			return 2
		}

		// 设置连接池参数
		db.SetMaxOpenConns(10)
		db.SetMaxIdleConns(5)
		db.SetConnMaxLifetime(time.Hour)

		// 生成连接 ID
		connID := fmt.Sprintf("%s:%d:%s", host, port, database)

		// 保存连接
		m.mysqlMutex.Lock()
		m.mysqlConns[connID] = db
		m.mysqlMutex.Unlock()

		// 返回连接对象（Lua 表）
		connTable := L.NewTable()
		connTable.RawSetString("id", lua.LString(connID))
		connTable.RawSetString("host", lua.LString(host))
		connTable.RawSetString("port", lua.LNumber(port))
		connTable.RawSetString("database", lua.LString(database))

		L.Push(connTable)
		return 1
	}))

	// 注册 mysql.closeSQL - 关闭数据库连接
	mysqlTable.RawSetString("closeSQL", state.NewFunction(func(L *lua.LState) int {
		connTable := L.CheckTable(1)
		connID := L.GetField(connTable, "id").String()

		m.mysqlMutex.Lock()
		if db, ok := m.mysqlConns[connID]; ok {
			db.Close()
			delete(m.mysqlConns, connID)
		}
		m.mysqlMutex.Unlock()

		return 0
	}))

	// 注册 mysql.executeQuerySQL - 执行 sql 语句并返回结果
	mysqlTable.RawSetString("executeQuerySQL", state.NewFunction(func(L *lua.LState) int {
		connTable := L.CheckTable(1)
		sqlStr := L.CheckString(2)
		connID := L.GetField(connTable, "id").String()

		m.mysqlMutex.Lock()
		db, ok := m.mysqlConns[connID]
		m.mysqlMutex.Unlock()

		if !ok {
			L.Push(lua.LNil)
			L.Push(lua.LString("数据库连接不存在"))
			return 2
		}

		// 执行查询
		rows, err := db.Query(sqlStr)
		if err != nil {
			L.Push(lua.LNil)
			L.Push(lua.LString(fmt.Sprintf("查询失败: %v", err)))
			return 2
		}
		defer rows.Close()

		// 获取列名
		columns, err := rows.Columns()
		if err != nil {
			L.Push(lua.LNil)
			L.Push(lua.LString(fmt.Sprintf("获取列名失败: %v", err)))
			return 2
		}

		// 创建结果表
		resultTable := L.NewTable()
		rowIndex := 1

		// 读取每一行
		for rows.Next() {
			// 创建值数组
			values := make([]interface{}, len(columns))
			valuePtrs := make([]interface{}, len(columns))
			for i := range columns {
				valuePtrs[i] = &values[i]
			}

			// 扫描行数据
			if err := rows.Scan(valuePtrs...); err != nil {
				L.Push(lua.LNil)
				L.Push(lua.LString(fmt.Sprintf("扫描数据失败: %v", err)))
				return 2
			}

			// 创建行表
			rowTable := L.NewTable()
			for i, col := range columns {
				var value interface{}
				val := values[i]
				b, ok := val.([]byte)
				if ok {
					value = string(b)
				} else {
					value = val
				}

				// 根据类型设置值
				switch v := value.(type) {
				case string:
					rowTable.RawSetString(col, lua.LString(v))
				case int:
					rowTable.RawSetString(col, lua.LNumber(v))
				case int64:
					rowTable.RawSetString(col, lua.LNumber(v))
				case float64:
					rowTable.RawSetString(col, lua.LNumber(v))
				case bool:
					rowTable.RawSetString(col, lua.LBool(v))
				case nil:
					rowTable.RawSetString(col, lua.LNil)
				default:
					rowTable.RawSetString(col, lua.LNil)
				}
			}

			// 添加到结果表
			resultTable.RawSetInt(rowIndex, rowTable)
			rowIndex++
		}

		// 检查是否有错误
		if err := rows.Err(); err != nil {
			L.Push(lua.LNil)
			L.Push(lua.LString(fmt.Sprintf("读取数据失败: %v", err)))
			return 2
		}

		L.Push(resultTable)
		return 1
	}))

	// 注册 mysql.executeSQL - 执行 sql 语句
	mysqlTable.RawSetString("executeSQL", state.NewFunction(func(L *lua.LState) int {
		connTable := L.CheckTable(1)
		sqlStr := L.CheckString(2)
		connID := L.GetField(connTable, "id").String()

		m.mysqlMutex.Lock()
		db, ok := m.mysqlConns[connID]
		m.mysqlMutex.Unlock()

		if !ok {
			L.Push(lua.LNil)
			L.Push(lua.LString("数据库连接不存在"))
			return 2
		}

		// 执行 SQL
		result, err := db.Exec(sqlStr)
		if err != nil {
			L.Push(lua.LNil)
			L.Push(lua.LString(fmt.Sprintf("执行失败: %v", err)))
			return 2
		}

		// 获取影响的行数
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			L.Push(lua.LNil)
			L.Push(lua.LString(fmt.Sprintf("获取影响行数失败: %v", err)))
			return 2
		}

		// 返回影响的行数
		L.Push(lua.LNumber(rowsAffected))
		return 1
	}))

	// 将 mysql 表注册到全局
	state.SetGlobal("mysql", mysqlTable)

	// ========== 线程管理相关方法 ==========

	// 初始化缓存
	m.cache = make(map[string]interface{})

	// 注册 setCache - 设置缓存（线程安全）
	state.SetGlobal("setCache", state.NewFunction(func(L *lua.LState) int {
		key := L.CheckString(1)
		value := L.CheckAny(2)

		m.cacheMutex.Lock()
		m.cache[key] = value
		m.cacheMutex.Unlock()

		return 0
	}))

	// 注册 getCache - 获取缓存（线程安全）
	state.SetGlobal("getCache", state.NewFunction(func(L *lua.LState) int {
		key := L.CheckString(1)

		m.cacheMutex.RLock()
		value, ok := m.cache[key]
		m.cacheMutex.RUnlock()

		if ok {
			// 将 interface{} 转换为 lua.LValue
			switch v := value.(type) {
			case string:
				L.Push(lua.LString(v))
			case int:
				L.Push(lua.LNumber(v))
			case int64:
				L.Push(lua.LNumber(v))
			case float64:
				L.Push(lua.LNumber(v))
			case bool:
				L.Push(lua.LBool(v))
			default:
				L.Push(lua.LNil)
			}
		} else {
			L.Push(lua.LNil)
		}
		return 1
	}))

	// 注册 delCache - 删除缓存（线程安全）
	state.SetGlobal("delCache", state.NewFunction(func(L *lua.LState) int {
		key := L.CheckString(1)

		m.cacheMutex.Lock()
		delete(m.cache, key)
		m.cacheMutex.Unlock()

		return 0
	}))

	// 注册 beginThread - 启动一个线程
	state.SetGlobal("beginThread", state.NewFunction(func(L *lua.LState) int {
		callback := L.CheckFunction(1)

		m.threadMutex.Lock()
		m.threadCounter++
		threadID := m.threadCounter
		threadInfo := &ThreadInfo{
			ID:     threadID,
			Stop:   false,
			Cancel: make(chan struct{}),
		}
		m.threads[threadID] = threadInfo
		m.threadMutex.Unlock()

		// 启动 goroutine
		go func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Printf("线程 %d 发生 panic: %v\n", threadID, r)
				}
			}()

			// 创建一个新的 Lua 状态来执行线程函数
			newL := lua.NewState()
			defer newL.Close()

			// 将函数推送到新的 Lua 状态并调用
			newL.Push(callback)
			if err := newL.PCall(0, 0, nil); err != nil {
				fmt.Printf("线程 %d 执行失败: %v\n", threadID, err)
			}
		}()

		L.Push(lua.LNumber(threadID))
		return 1
	}))

	// 创建 Thread 表
	threadTable := state.NewTable()

	// 注册 Thread.newThread - 启动一个新的线程
	threadTable.RawSetString("newThread", state.NewFunction(func(luaState *lua.LState) int {
		callback := luaState.CheckFunction(1)

		m.threadMutex.Lock()
		m.threadCounter++
		threadID := m.threadCounter
		threadInfo := &ThreadInfo{
			ID:     threadID,
			Stop:   false,
			Cancel: make(chan struct{}),
		}
		m.threads[threadID] = threadInfo
		m.threadMutex.Unlock()

		// 启动 goroutine
		go func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Printf("线程 %d 发生 panic: %v\n", threadID, r)
				}
			}()

			// 创建一个新的 Lua 状态来执行线程函数
			newL := lua.NewState()
			defer newL.Close()

			// 将函数推送到新的 Lua 状态并调用
			newL.Push(callback)
			if err := newL.PCall(0, 0, nil); err != nil {
				fmt.Printf("线程 %d 执行失败: %v\n", threadID, err)
			}
		}()

		// 返回线程对象
		threadObj := luaState.NewTable()
		threadObj.RawSetString("id", lua.LNumber(threadID))
		threadObj.RawSetString("stopThread", luaState.NewFunction(func(luaState *lua.LState) int {
			m.threadMutex.Lock()
			if info, ok := m.threads[threadID]; ok {
				info.Stop = true
				close(info.Cancel)
				delete(m.threads, threadID)
			}
			m.threadMutex.Unlock()
			return 0
		}))
		luaState.Push(threadObj)
		return 1
	}))

	state.SetGlobal("Thread", threadTable)

	engine.RegisterMethod("extension.md5", "字符串 MD5", extensionMD5, true)
	engine.RegisterMethod("extension.fileMD5", "获取文件 MD5", extensionFileMD5, true)
	engine.RegisterMethod("extension.base64Encode", "Base64 编码", extensionBase64Encode, true)
	engine.RegisterMethod("extension.base64Decode", "Base64 解码", extensionBase64Decode, true)
	engine.RegisterMethod("extension.encodeUrl", "URL 编码", extensionURLEncode, true)
	engine.RegisterMethod("extension.decodeUrl", "URL 解码", extensionURLDecode, true)
	engine.RegisterMethod("extension.urlEncode", "URL 编码别名", extensionURLEncode, true)
	engine.RegisterMethod("extension.urlDecode", "URL 解码别名", extensionURLDecode, true)
	engine.RegisterMethod("extension.htmlEncode", "HTML 编码", extensionHTMLEncode, true)
	engine.RegisterMethod("extension.htmlDecode", "HTML 解码", extensionHTMLDecode, true)
	engine.RegisterMethod("extension.jsonEncode", "JSON 编码", extensionJSONEncode, true)
	engine.RegisterMethod("extension.jsonDecode", "JSON 解码", extensionJSONDecode, true)
	engine.RegisterMethod("extension.getFileBase64", "获取文件 Base64 编码", extensionGetFileBase64, true)
	engine.RegisterMethod("extension.fileExist", "文件或目录是否存在", extensionFileExist, true)
	engine.RegisterMethod("extension.mkdir", "创建目录", extensionMkdir, true)
	engine.RegisterMethod("extension.delfile", "删除文件或目录", extensionDeleteFile, true)
	engine.RegisterMethod("extension.readFile", "读取文件内容", extensionReadFile, true)
	engine.RegisterMethod("extension.writeFile", "写入文件内容", extensionWriteFile, true)
	engine.RegisterMethod("extension.fileSize", "获取文件大小", extensionFileSize, true)

	return nil
}

// GetModule 获取模块实例
func GetModule() model.Module {
	return New()
}

func extensionMD5(str string) string {
	hash := md5.Sum([]byte(str))
	return hex.EncodeToString(hash[:])
}

func extensionFileMD5(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		return ""
	}
	hash := md5.Sum(data)
	return hex.EncodeToString(hash[:])
}

func extensionBase64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func extensionBase64Decode(str string) interface{} {
	decoded, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return nil
	}
	return string(decoded)
}

func extensionURLEncode(str string) string {
	return url.QueryEscape(str)
}

func extensionURLDecode(str string) interface{} {
	decoded, err := url.QueryUnescape(str)
	if err != nil {
		return nil
	}
	return decoded
}

func extensionHTMLEncode(str string) string {
	encoded := ""
	for _, c := range str {
		switch c {
		case '<':
			encoded += "&lt;"
		case '>':
			encoded += "&gt;"
		case '&':
			encoded += "&amp;"
		case '"':
			encoded += "&quot;"
		case '\'':
			encoded += "&#39;"
		default:
			encoded += string(c)
		}
	}
	return encoded
}

func extensionHTMLDecode(str string) string {
	decoded := str
	decoded = strings.ReplaceAll(decoded, "&lt;", "<")
	decoded = strings.ReplaceAll(decoded, "&gt;", ">")
	decoded = strings.ReplaceAll(decoded, "&amp;", "&")
	decoded = strings.ReplaceAll(decoded, "&quot;", "\"")
	decoded = strings.ReplaceAll(decoded, "&#39;", "'")
	return decoded
}

func extensionJSONEncode(value interface{}) interface{} {
	data, err := json.Marshal(value)
	if err != nil {
		return nil
	}
	return string(data)
}

func extensionJSONDecode(jsonStr string) interface{} {
	var result interface{}
	if err := json.Unmarshal([]byte(jsonStr), &result); err != nil {
		return nil
	}
	return result
}

func extensionGetFileBase64(path string) interface{} {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil
	}
	return base64.StdEncoding.EncodeToString(data)
}

func extensionFileExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func extensionMkdir(path string) bool {
	return os.MkdirAll(path, 0755) == nil
}

func extensionDeleteFile(path string) bool {
	return os.RemoveAll(path) == nil
}

func extensionReadFile(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		return ""
	}
	return string(data)
}

func extensionWriteFile(path, str string, appendFlag ...bool) bool {
	flag := os.O_CREATE | os.O_WRONLY
	if len(appendFlag) > 0 && appendFlag[0] {
		flag |= os.O_APPEND
	} else {
		flag |= os.O_TRUNC
	}
	file, err := os.OpenFile(path, flag, 0644)
	if err != nil {
		return false
	}
	defer file.Close()
	_, err = file.WriteString(str)
	return err == nil
}

func extensionFileSize(path string) int {
	info, err := os.Stat(path)
	if err != nil {
		return 0
	}
	return int(info.Size())
}

// luaValueToJson 将 Lua 值转换为 JSON 字符串
func luaValueToJson(L *lua.LState, value lua.LValue) (string, error) {
	if value == lua.LNil {
		return "null", nil
	}

	switch v := value.(type) {
	case lua.LBool:
		if bool(v) {
			return "true", nil
		}
		return "false", nil
	case lua.LNumber:
		return fmt.Sprintf("%v", v), nil
	case lua.LString:
		escaped := strings.ReplaceAll(string(v), "\\", "\\\\")
		escaped = strings.ReplaceAll(escaped, "\"", "\\\"")
		escaped = strings.ReplaceAll(escaped, "\n", "\\n")
		escaped = strings.ReplaceAll(escaped, "\r", "\\r")
		escaped = strings.ReplaceAll(escaped, "\t", "\\t")
		return fmt.Sprintf("\"%s\"", escaped), nil
	case *lua.LTable:
		// 检查是否是数组
		isArray := true
		if v.Len() > 0 {
			for i := 1; i <= v.Len(); i++ {
				if v.RawGetInt(i) == lua.LNil {
					isArray = false
					break
				}
			}
		} else {
			// 空表默认是对象
			isArray = false
		}

		if isArray {
			// 处理数组
			var parts []string
			for i := 1; i <= v.Len(); i++ {
				item := v.RawGetInt(i)
				jsonItem, err := luaValueToJson(L, item)
				if err != nil {
					return "", err
				}
				parts = append(parts, jsonItem)
			}
			return fmt.Sprintf("[%s]", strings.Join(parts, ",")), nil
		} else {
			// 处理对象
			var parts []string
			v.ForEach(func(key, val lua.LValue) {
				if keyStr, ok := key.(lua.LString); ok {
					jsonKey, err := luaValueToJson(L, keyStr)
					if err != nil {
						return
					}
					jsonVal, err := luaValueToJson(L, val)
					if err != nil {
						return
					}
					parts = append(parts, fmt.Sprintf("%s:%s", jsonKey, jsonVal))
				}
			})
			return fmt.Sprintf("{%s}", strings.Join(parts, ",")), nil
		}
	default:
		return "", fmt.Errorf("unsupported type: %T", value)
	}
}

// jsonToLua 将 JSON 字符串转换为 Lua 值
func jsonToLua(L *lua.LState, jsonStr string) (lua.LValue, error) {
	var result interface{}
	if err := json.Unmarshal([]byte(jsonStr), &result); err != nil {
		return nil, err
	}
	return goValueToLua(L, result)
}

// goValueToLua 将 Go 值转换为 Lua 值
func goValueToLua(L *lua.LState, value interface{}) (lua.LValue, error) {
	switch v := value.(type) {
	case nil:
		return lua.LNil, nil
	case bool:
		return lua.LBool(v), nil
	case float64:
		return lua.LNumber(v), nil
	case string:
		return lua.LString(v), nil
	case []interface{}:
		table := L.NewTable()
		for i, item := range v {
			luaItem, err := goValueToLua(L, item)
			if err != nil {
				return nil, err
			}
			table.RawSetInt(i+1, luaItem)
		}
		return table, nil
	case map[string]interface{}:
		table := L.NewTable()
		for key, item := range v {
			luaItem, err := goValueToLua(L, item)
			if err != nil {
				return nil, err
			}
			table.RawSetString(key, luaItem)
		}
		return table, nil
	default:
		return nil, fmt.Errorf("unsupported type: %T", value)
	}
}
