package lua_engine

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	autogoallmodels "github.com/ZingYao/autogo_scriptengine/lua_engine/define/android/autogo/all_models"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"
	autogocoroutine "github.com/ZingYao/autogo_scriptengine/lua_engine/model/autogo/coroutine"
	autogojson "github.com/ZingYao/autogo_scriptengine/lua_engine/model/autogo/json"
	lrconsole "github.com/ZingYao/autogo_scriptengine/lua_engine/model/lrappsoft/console"
	lrcrypt "github.com/ZingYao/autogo_scriptengine/lua_engine/model/lrappsoft/crypt"
	lrdynamicui "github.com/ZingYao/autogo_scriptengine/lua_engine/model/lrappsoft/dynamicui"
	lrffi "github.com/ZingYao/autogo_scriptengine/lua_engine/model/lrappsoft/ffi"
	lrgobridge "github.com/ZingYao/autogo_scriptengine/lua_engine/model/lrappsoft/gobridge"
	lrhttp "github.com/ZingYao/autogo_scriptengine/lua_engine/model/lrappsoft/http"
	lrimgui "github.com/ZingYao/autogo_scriptengine/lua_engine/model/lrappsoft/imgui"
	lrjson "github.com/ZingYao/autogo_scriptengine/lua_engine/model/lrappsoft/json"
	lrlfs "github.com/ZingYao/autogo_scriptengine/lua_engine/model/lrappsoft/lfs"
	lrmath "github.com/ZingYao/autogo_scriptengine/lua_engine/model/lrappsoft/math"
	lrnetwork "github.com/ZingYao/autogo_scriptengine/lua_engine/model/lrappsoft/network"
	lrstring "github.com/ZingYao/autogo_scriptengine/lua_engine/model/lrappsoft/string_module"
	lrtime "github.com/ZingYao/autogo_scriptengine/lua_engine/model/lrappsoft/time"
)

func TestLuaEngineExecuteStringWithGoLuaVM(t *testing.T) {
	engine := NewLuaEngine(nil)
	defer engine.Close()

	if err := engine.ExecuteString(`
		local value = 40 + 2
		if value ~= 42 then
			error("unexpected value")
		end
	`); err != nil {
		t.Fatalf("ExecuteString() error = %v", err)
	}
}

func TestLuaEngineExecuteStringWithGoLuaVMAsyncMode(t *testing.T) {
	engine := NewLuaEngine(nil)
	defer engine.Close()

	var executed atomic.Bool
	engine.RegisterMethod("async.mark", "mark async execution", func() {
		executed.Store(true)
	}, true)

	if err := engine.ExecuteStringWithMode(`async.mark()`, ExecuteModeAsync); err != nil {
		t.Fatalf("ExecuteStringWithMode(async) error = %v", err)
	}
	deadline := time.Now().Add(time.Second)
	for time.Now().Before(deadline) {
		if executed.Load() {
			return
		}
		time.Sleep(10 * time.Millisecond)
	}
	t.Fatal("async script did not execute before deadline")
}

func TestLuaEngineExecuteFileWithGoLuaVMAsyncMode(t *testing.T) {
	engine := NewLuaEngine(nil)
	defer engine.Close()

	var executed atomic.Bool
	engine.RegisterMethod("asyncFile.mark", "mark async file execution", func() {
		executed.Store(true)
	}, true)

	scriptPath := filepath.Join(t.TempDir(), "async.lua")
	if err := os.WriteFile(scriptPath, []byte(`asyncFile.mark()`), 0644); err != nil {
		t.Fatalf("WriteFile() error = %v", err)
	}
	if err := engine.ExecuteFileWithMode(scriptPath, ExecuteModeAsync); err != nil {
		t.Fatalf("ExecuteFileWithMode(async) error = %v", err)
	}
	deadline := time.Now().Add(time.Second)
	for time.Now().Before(deadline) {
		if executed.Load() {
			return
		}
		time.Sleep(10 * time.Millisecond)
	}
	t.Fatal("async file script did not execute before deadline")
}

func TestLuaEngineGoLuaVMCloseRejectsFurtherExecution(t *testing.T) {
	engine := NewLuaEngine(nil)
	engine.Start()
	engine.Close()

	if state := engine.GetEngineState(); state != StateStopped {
		t.Fatalf("state after Close = %s, want stopped", state.String())
	}
	if err := engine.ExecuteString(`return 1`); err == nil {
		t.Fatal("ExecuteString() after Close error = nil, want error")
	}
	if _, err := engine.CompileString(`return 1`); err == nil {
		t.Fatal("CompileString() after Close error = nil, want error")
	}
}

func TestLuaEngineGoLuaVMLifecycleStateTransitions(t *testing.T) {
	engine := NewLuaEngine(nil)
	defer engine.Close()

	if state := engine.GetEngineState(); state != StateStopped {
		t.Fatalf("initial state = %s, want stopped", state.String())
	}
	engine.Start()
	if state := engine.GetEngineState(); state != StateRunning {
		t.Fatalf("state after Start = %s, want running", state.String())
	}
	engine.Pause()
	if state := engine.GetEngineState(); state != StatePaused {
		t.Fatalf("state after Pause = %s, want paused", state.String())
	}
	engine.Resume()
	if state := engine.GetEngineState(); state != StateRunning {
		t.Fatalf("state after Resume = %s, want running", state.String())
	}
	engine.Stop()
	if state := engine.GetEngineState(); state != StateStopped {
		t.Fatalf("state after Stop = %s, want stopped", state.String())
	}
}

func TestLuaEngineGoLuaVMExitMinusOneSkipsExitAction(t *testing.T) {
	engine := NewLuaEngine(nil)
	defer engine.Close()

	if err := engine.ExecuteString(`os.exit(-1)`); err != nil {
		t.Fatalf("ExecuteString() error = %v", err)
	}
	if !engine.skipExitAction {
		t.Fatal("skipExitAction = false, want true")
	}
}

func TestLuaEngineGoLuaVMRegisterMethodBridge(t *testing.T) {
	engine := NewLuaEngine(nil)
	defer engine.Close()

	engine.RegisterMethod("mathx.add", "add two integers", func(left, right int) int {
		return left + right
	}, true)

	if err := engine.ExecuteString(`
		local value = mathx.add(20, 22)
		if value ~= 42 then
			error("unexpected bridge value")
		end
	`); err != nil {
		t.Fatalf("ExecuteString() error = %v", err)
	}
}

func TestLuaEngineGoLuaVMGlobalConvenienceFunctions(t *testing.T) {
	previousEngine := engine
	engine = NewLuaEngine(nil)
	defer func() {
		engine.Close()
		engine = previousEngine
	}()

	RegisterMethod("global.echo", "global echo", func(value string) string {
		return value
	}, true)
	if err := ExecuteString(`
		if global.echo("string") ~= "string" then
			error("global ExecuteString failed")
		end
	`); err != nil {
		t.Fatalf("ExecuteString() global error = %v", err)
	}
	bytecode, err := CompileString(`
		if global.echo("bytecode") ~= "bytecode" then
			error("global ExecuteBytecode failed")
		end
	`, "global_bytecode.lua")
	if err != nil {
		t.Fatalf("CompileString() global error = %v", err)
	}
	if err := ExecuteBytecode(bytecode); err != nil {
		t.Fatalf("ExecuteBytecode() global error = %v", err)
	}

	scriptPath := filepath.Join(t.TempDir(), "global_file.lua")
	if err := os.WriteFile(scriptPath, []byte(`
		if global.echo("file") ~= "file" then
			error("global ExecuteFile failed")
		end
	`), 0644); err != nil {
		t.Fatalf("WriteFile() global file error = %v", err)
	}
	if err := ExecuteFile(scriptPath); err != nil {
		t.Fatalf("ExecuteFile() global error = %v", err)
	}
	fileBytecode, err := CompileFile(scriptPath)
	if err != nil {
		t.Fatalf("CompileFile() global error = %v", err)
	}
	if err := ExecuteBytecode(fileBytecode); err != nil {
		t.Fatalf("ExecuteBytecode(file) global error = %v", err)
	}
}

func TestLuaEngineGoLuaVMNewEngineInstallsPreRegisteredMethods(t *testing.T) {
	previousEngine := engine
	engine = nil
	defer func() {
		engine = previousEngine
	}()

	RegisterMethod("preinit.echo", "preinit echo", func(value string) string {
		return value
	}, true)

	createdEngine := NewLuaEngine(nil)
	defer createdEngine.Close()
	if err := createdEngine.ExecuteString(`
		if preinit.echo("ready") ~= "ready" then
			error("pre-registered method was not installed")
		end
	`); err != nil {
		t.Fatalf("ExecuteString() pre-registered method error = %v", err)
	}
}

func TestLuaEngineGoLuaVMDefaultEngineInstallsPreRegisteredMethods(t *testing.T) {
	previousEngine := engine
	engine = nil
	once = sync.Once{}
	defer func() {
		if engine != nil {
			engine.Close()
		}
		engine = previousEngine
		once = sync.Once{}
	}()

	RegisterMethod("defaultPreinit.echo", "default preinit echo", func(value string) string {
		return value
	}, true)

	defaultEngine := GetLuaEngine()
	if err := defaultEngine.ExecuteString(`
		if defaultPreinit.echo("ready") ~= "ready" then
			error("default pre-registered method was not installed")
		end
	`); err != nil {
		t.Fatalf("ExecuteString() default pre-registered method error = %v", err)
	}
}

func TestLuaEngineGoLuaVMGlobalCloseResetsDefaultEngine(t *testing.T) {
	previousEngine := engine
	engine = NewLuaEngine(nil)
	once = sync.Once{}
	defer func() {
		if engine != nil {
			engine.Close()
		}
		engine = previousEngine
		once = sync.Once{}
	}()

	RegisterMethod("closedDefault.echo", "closed default echo", func(value string) string {
		return value
	}, true)
	Close()
	if engine != nil {
		t.Fatal("global Close() did not clear default engine")
	}

	reopenedEngine := GetLuaEngine()
	if reopenedEngine == nil {
		t.Fatal("GetLuaEngine() after Close returned nil")
	}
	if err := reopenedEngine.ExecuteString(`
		if closedDefault.echo("ready") ~= "ready" then
			error("closed default method was not reinstalled")
		end
	`); err != nil {
		t.Fatalf("ExecuteString() after global Close error = %v", err)
	}
}

func TestLuaEngineGoLuaVMRegisterMethodBridgeWithErrorReturn(t *testing.T) {
	engine := NewLuaEngine(nil)
	defer engine.Close()

	engine.RegisterMethod("mathx.ok", "return value with nil error", func(value string) (string, error) {
		return value + "-ok", nil
	}, true)
	engine.RegisterMethod("mathx.fail", "return go error", func() error {
		return errors.New("bridge failure")
	}, true)

	if err := engine.ExecuteString(`
		local value = mathx.ok("lua")
		if value ~= "lua-ok" then
			error("unexpected nil-error bridge value")
		end
	`); err != nil {
		t.Fatalf("ExecuteString() nil error bridge failed: %v", err)
	}

	err := engine.ExecuteString(`mathx.fail()`)
	if err == nil {
		t.Fatal("ExecuteString() error = nil, want bridge failure")
	}
	if !strings.Contains(err.Error(), "bridge failure") {
		t.Fatalf("ExecuteString() error = %v, want bridge failure", err)
	}
}

func TestLuaEngineGoLuaVMRegisterModuleBridge(t *testing.T) {
	engine := NewLuaEngine(nil)
	defer engine.Close()

	engine.RegisterModule(testVMModule{})

	if err := engine.ExecuteString(`
		local value = vmtest.concat("go", "lua")
		if value ~= "go-lua" then
			error("unexpected module bridge value")
		end
	`); err != nil {
		t.Fatalf("ExecuteString() error = %v", err)
	}
}

func TestLuaEngineGoLuaVMAutogoAllModulesExample(t *testing.T) {
	engine := NewLuaEngine(nil)
	defer engine.Close()

	engine.RegisterModule(autogoallmodels.AllModules...)
	if err := engine.ExecuteFile(filepath.Join("..", "examples", "lua_engine", "autogo", "scripts", "main.lua")); err != nil {
		t.Fatalf("ExecuteFile(autogo example) error = %v", err)
	}
}

func TestLuaEngineGoLuaVMRegistersLrTimeModule(t *testing.T) {
	engine := NewLuaEngine(nil)
	defer engine.Close()

	engine.RegisterModule(&lrtime.TimeModule{})

	if err := engine.ExecuteString(`
		if type(time) ~= "table" then
			error("time module was not registered")
		end
		local now = time.systemTime()
		if type(now) ~= "number" or now <= 0 then
			error("unexpected systemTime result")
		end
		local tick = time.tickCount()
		if type(tick) ~= "number" or tick < 0 then
			error("unexpected tickCount result")
		end
		local requiredTime = require("time")
		if requiredTime ~= time or type(requiredTime.systemTime()) ~= "number" then
			error("unexpected require('time') result")
		end
	`); err != nil {
		t.Fatalf("ExecuteString() lr time module error = %v", err)
	}
}

func TestLuaEngineGoLuaVMRegistersLrMathModule(t *testing.T) {
	engine := NewLuaEngine(nil)
	defer engine.Close()

	engine.RegisterModule(&lrmath.MathModule{})

	if err := engine.ExecuteString(`
		if type(math) ~= "table" then
			error("math module was not registered")
		end
		if math.tointeger("42") ~= 42 then
			error("unexpected tointeger result")
		end
		if math.type(1.5) ~= "float" then
			error("unexpected math.type float result")
		end
		if math.type(2) ~= "integer" then
			error("unexpected math.type integer result")
		end
		if not math.ult(1, 2) then
			error("unexpected math.ult result")
		end
	`); err != nil {
		t.Fatalf("ExecuteString() lr math module error = %v", err)
	}
}

func TestLuaEngineGoLuaVMRegistersLrStringModule(t *testing.T) {
	engine := NewLuaEngine(nil)
	defer engine.Close()

	engine.RegisterModule(&lrstring.StringModule{})

	if err := engine.ExecuteString(`
		local parts = string.splitStr("a,b,c", ",")
		if parts[1] ~= "a" or parts[3] ~= "c" then
			error("unexpected splitStr result")
		end
		if utf8.length("你好ab") ~= 4 then
			error("unexpected utf8.length result")
		end
		if utf8.left("你好世界", 2) ~= "你好" then
			error("unexpected utf8.left result")
		end
		if utf8.right("你好世界", 2) ~= "世界" then
			error("unexpected utf8.right result")
		end
		if utf8.mid("你好世界", 1, 2) ~= "好世" then
			error("unexpected utf8.mid result")
		end
		if utf8.strCut("你好世界", 1, 2) ~= "你界" then
			error("unexpected utf8.strCut result")
		end
		if utf8.inStr(1, "你好世界", "世界") ~= 3 then
			error("unexpected utf8.inStr result")
		end
		if utf8.inStrRev("你好世界你好", "你好", 6) ~= 5 then
			error("unexpected utf8.inStrRev result")
		end
		if utf8.strReverse("你好ab") ~= "ba好你" then
			error("unexpected utf8.strReverse result")
		end
	`); err != nil {
		t.Fatalf("ExecuteString() lr string module error = %v", err)
	}
}

func TestLuaEngineGoLuaVMRegistersLrCryptModule(t *testing.T) {
	engine := NewLuaEngine(nil)
	defer engine.Close()

	engine.RegisterModule(&lrcrypt.CryptModule{})

	if err := engine.ExecuteString(`
		if type(cryptLib) ~= "table" then
			error("cryptLib module was not registered")
		end
		if cryptLib.base64_encode("hello") ~= "aGVsbG8=" then
			error("unexpected base64 encode result")
		end
		if cryptLib.base64_decode("aGVsbG8=") ~= "hello" then
			error("unexpected base64 decode result")
		end
		if cryptLib.md5("hello") ~= "5d41402abc4b2a76b9719d911017c592" then
			error("unexpected md5 result")
		end
		if cryptLib.sha256("hello") ~= "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824" then
			error("unexpected sha256 result")
		end
		if cryptLib.hmac_sha256("hello", "key") ~= "9307b3b915efb5171ff14d8cb55fbcc798c6c0ef1456d66ded1a6aa723a58b7b" then
			error("unexpected hmac_sha256 result")
		end
		local encrypted = cryptLib.aes_crypt("hello", "1234567890abcdef", "encrypt", "ecb")
		if type(encrypted) ~= "string" or encrypted == "" then
			error("unexpected aes_crypt encrypt result")
		end
		local decrypted = cryptLib.aes_crypt(encrypted, "1234567890abcdef", "decrypt", "ecb")
		if decrypted ~= "hello" then
			error("unexpected aes_crypt decrypt result")
		end
		local publicKey, privateKey = cryptLib.rsa_generate_key()
		if type(publicKey) ~= "string" or not string.find(publicKey, "RSA PUBLIC KEY", 1, true) then
			error("unexpected rsa_generate_key public key")
		end
		if type(privateKey) ~= "string" or not string.find(privateKey, "RSA PRIVATE KEY", 1, true) then
			error("unexpected rsa_generate_key private key")
		end
		local aliasPublicKey, aliasPrivateKey = cryptLib.rsa_keygen(1024)
		if type(aliasPublicKey) ~= "string" or not string.find(aliasPublicKey, "RSA PUBLIC KEY", 1, true) then
			error("unexpected rsa_keygen public key")
		end
		if type(aliasPrivateKey) ~= "string" or not string.find(aliasPrivateKey, "RSA PRIVATE KEY", 1, true) then
			error("unexpected rsa_keygen private key")
		end
		local ok, err = pcall(function()
			return cryptLib.base64_decode("%%%")
		end)
		if ok or not err then
			error("unexpected base64 decode error handling")
		end
	`); err != nil {
		t.Fatalf("ExecuteString() lr crypt module error = %v", err)
	}
}

func TestLuaEngineGoLuaVMRegistersLrJSONModule(t *testing.T) {
	engine := NewLuaEngine(nil)
	defer engine.Close()

	engine.RegisterModule(&lrjson.JsonModule{})

	if err := engine.ExecuteString(`
		if type(jsonLib) ~= "table" then
			error("jsonLib module was not registered")
		end
		local encoded = jsonLib.encode({ name = "lua", count = 2 })
		if encoded ~= '{"count":2,"name":"lua"}' then
			error("unexpected json encode result: " .. tostring(encoded))
		end
		local decoded = jsonLib.decode('{"name":"go","items":["a","b"]}')
		if decoded.name ~= "go" or decoded.items[2] ~= "b" then
			error("unexpected json object decode result")
		end
		local decodedArray = jsonLib.decode('["x","y"]')
		if decodedArray[1] ~= "x" or decodedArray[2] ~= "y" then
			error("unexpected json array decode result")
		end
		local ok, err = pcall(function()
			return jsonLib.decode('{')
		end)
		if ok or not err then
			error("unexpected json decode error handling")
		end
	`); err != nil {
		t.Fatalf("ExecuteString() lr json module error = %v", err)
	}
}

func TestLuaEngineGoLuaVMRegistersLrLfsModule(t *testing.T) {
	engine := NewLuaEngine(nil)
	defer engine.Close()

	rootDir := t.TempDir()
	filePath := filepath.Join(rootDir, "file.txt")
	if err := os.WriteFile(filePath, []byte("hello"), 0o644); err != nil {
		t.Fatalf("WriteFile() error = %v", err)
	}
	engine.RegisterModule(&lrlfs.LfsModule{})
	engine.RegisterMethod("test.root", "return test root", func() string {
		return rootDir
	}, true)

	if err := engine.ExecuteString(`
		local root = test.root()
		local child = root .. "/child"
		local file = root .. "/file.txt"
		if type(lfs) ~= "table" then
			error("lfs module was not registered")
		end
		if require("lfs") ~= lfs then
			error("unexpected require('lfs') result")
		end
		local cwd = lfs.currentdir()
		if type(cwd) ~= "string" or cwd == "" then
			error("unexpected currentdir result")
		end
		local ok = lfs.mkdir(child)
		if ok ~= nil then
			error("mkdir should not return a success value through error-only bridge")
		end
		local dirAttr = lfs.attributes(child)
		if dirAttr.mode ~= "directory" then
			error("unexpected directory attributes")
		end
		local fileAttr = lfs.attributes(file)
		if fileAttr.mode ~= "file" or fileAttr.size ~= 5 then
			error("unexpected file attributes")
		end
		local hardLink = root .. "/hard-link.txt"
		if lfs.link(file, hardLink) ~= true then
			error("link should return true")
		end
		local symlink = root .. "/symbolic-link.txt"
		if lfs.symlink(file, symlink) ~= true then
			error("symlink should return true")
		end
		local symlinkAttr = lfs.symlinkattributes(symlink)
		if symlinkAttr.mode ~= "link" then
			error("unexpected symlink attributes")
		end
		local entries = lfs.dir(root)
		if type(entries) ~= "table" then
			error("dir should return a table through go-lua-vm bridge")
		end
		local foundFile = false
		local foundHardLink = false
		local foundSymlink = false
		for _, name in ipairs(entries) do
			if name == "file.txt" then
				foundFile = true
			elseif name == "hard-link.txt" then
				foundHardLink = true
			elseif name == "symbolic-link.txt" then
				foundSymlink = true
			end
		end
		if not foundFile or not foundHardLink or not foundSymlink then
			error("dir did not include expected entries")
		end
		if lfs.lock_dir(root) ~= true then
			error("lock_dir should return true")
		end
		if lfs.attributes(root .. "/.lock").mode ~= "file" then
			error("lock_dir should create a lock file")
		end
		if lfs.touch(file) ~= nil then
			error("touch should not return a success value through error-only bridge")
		end
		if lfs.rmdir(child) ~= nil then
			error("rmdir should not return a success value through error-only bridge")
		end
	`); err != nil {
		t.Fatalf("ExecuteString() lr lfs module error = %v", err)
	}
}

func TestLuaEngineGoLuaVMRegistersLrConsoleModule(t *testing.T) {
	engine := NewLuaEngine(nil)
	defer engine.Close()

	engine.RegisterModule(&lrconsole.ConsoleModule{})

	if err := engine.ExecuteString(`
		if type(console) ~= "table" then
			error("console module was not registered")
		end
		if console.show() ~= true then
			error("unexpected console.show result")
		end
		if console.showTitle() ~= true then
			error("unexpected console.showTitle default result")
		end
		if console.showTitle(false) ~= true then
			error("unexpected console.showTitle explicit result")
		end
		console.setPos(1, 2)
		console.setPos(1, 2, 3, 4)
		console.println()
		console.println(1, "hello")
		console.lockConsole()
		console.unlockConsole()
		if console.dismiss() ~= true then
			error("unexpected console.dismiss result")
		end
		console.clearLog()
		console.setTitle("title")
	`); err != nil {
		t.Fatalf("ExecuteString() lr console module error = %v", err)
	}
}

func TestLuaEngineGoLuaVMRegistersLrNetworkModule(t *testing.T) {
	engine := NewLuaEngine(nil)
	defer engine.Close()

	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		switch request.URL.Path {
		case "/get":
			_, _ = writer.Write([]byte("get:" + request.URL.Query().Get("q")))
		case "/post":
			body, _ := io.ReadAll(request.Body)
			_, _ = writer.Write([]byte("post:" + string(body)))
		case "/post-data":
			body, _ := io.ReadAll(request.Body)
			_, _ = writer.Write([]byte(request.Header.Get("Content-Type") + ":" + string(body)))
		case "/download":
			_, _ = writer.Write([]byte("download-body"))
		case "/upload":
			if err := request.ParseMultipartForm(1024 * 1024); err != nil {
				writer.WriteHeader(http.StatusBadRequest)
				_, _ = writer.Write([]byte(err.Error()))
				return
			}
			file, header, err := request.FormFile("file")
			if err != nil {
				writer.WriteHeader(http.StatusBadRequest)
				_, _ = writer.Write([]byte(err.Error()))
				return
			}
			defer file.Close()
			body, _ := io.ReadAll(file)
			_, _ = writer.Write([]byte(header.Filename + ":" + string(body)))
		case "/http-upload":
			if request.Header.Get("X-Test") != "lua" {
				writer.WriteHeader(http.StatusBadRequest)
				return
			}
			if err := request.ParseMultipartForm(1024 * 1024); err != nil {
				writer.WriteHeader(http.StatusBadRequest)
				_, _ = writer.Write([]byte(err.Error()))
				return
			}
			file, header, err := request.FormFile("payload")
			if err != nil {
				writer.WriteHeader(http.StatusBadRequest)
				_, _ = writer.Write([]byte(err.Error()))
				return
			}
			defer file.Close()
			body, _ := io.ReadAll(file)
			_, _ = writer.Write([]byte(header.Filename + ":" + string(body)))
		default:
			writer.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()

	tempDir := t.TempDir()
	downloadPath := filepath.Join(tempDir, "download.txt")
	httpDownloadPath := filepath.Join(tempDir, "http-download.txt")
	uploadPath := filepath.Join(tempDir, "upload.txt")
	if err := os.WriteFile(uploadPath, []byte("upload-body"), 0o644); err != nil {
		t.Fatalf("WriteFile(upload) error = %v", err)
	}
	engine.RegisterModule(&lrnetwork.NetworkModule{})
	engine.RegisterMethod("test.networkBaseURL", "return test network base url", func() string {
		return server.URL
	}, true)
	engine.RegisterMethod("test.downloadPath", "return test download path", func() string {
		return downloadPath
	}, true)
	engine.RegisterMethod("test.httpDownloadPath", "return test http download path", func() string {
		return httpDownloadPath
	}, true)
	engine.RegisterMethod("test.uploadPath", "return test upload path", func() string {
		return uploadPath
	}, true)

	if err := engine.ExecuteString(`
		if type(network) ~= "table" then
			error("network module was not registered")
		end
		local base = test.networkBaseURL()
		local getDefaultTimeoutResult = network.httpGet(base .. "/get?q=default")
		if getDefaultTimeoutResult.body ~= "get:default" or getDefaultTimeoutResult.statusCode ~= 200 then
			error("unexpected network.httpGet default timeout result")
		end
		local getResult = network.httpGet(base .. "/get?q=lua", 5)
		if getResult.body ~= "get:lua" or getResult.statusCode ~= 200 then
			error("unexpected network.httpGet result")
		end
		local postDefaultTimeoutResult = network.httpPost(base .. "/post", "name=default")
		if postDefaultTimeoutResult.body ~= "post:name=default" or postDefaultTimeoutResult.statusCode ~= 200 then
			error("unexpected network.httpPost default timeout result")
		end
		local postResult = network.httpPost(base .. "/post", "name=lua", 5)
		if postResult.body ~= "post:name=lua" or postResult.statusCode ~= 200 then
			error("unexpected network.httpPost result")
		end
		local postDataResult = network.httpPostData(base .. "/post-data", '{"name":"lua"}', "application/json")
		if postDataResult ~= 'application/json:{"name":"lua"}' then
			error("unexpected network.httpPostData result: " .. postDataResult)
		end
		if not network.downloadFile(base .. "/download", test.downloadPath()) then
			error("unexpected network.downloadFile result")
		end
		if network.uploadFile(base .. "/upload", test.uploadPath()) ~= "upload.txt:upload-body" then
			error("unexpected network.uploadFile result")
		end
		local uploadBody, uploadStatus = network.httpUpload(base .. "/http-upload", "payload", test.uploadPath(), { ["X-Test"] = "lua" })
		if uploadBody ~= "upload.txt:upload-body" or uploadStatus ~= 200 then
			error("unexpected network.httpUpload result")
		end
		if not network.httpDownload(base .. "/download", test.httpDownloadPath()) then
			error("unexpected network.httpDownload result")
		end
	`); err != nil {
		t.Fatalf("ExecuteString() lr network module error = %v", err)
	}
	downloaded, err := os.ReadFile(downloadPath)
	if err != nil {
		t.Fatalf("ReadFile(download) error = %v", err)
	}
	if string(downloaded) != "download-body" {
		t.Fatalf("downloaded body = %q, want download-body", string(downloaded))
	}
	httpDownloaded, err := os.ReadFile(httpDownloadPath)
	if err != nil {
		t.Fatalf("ReadFile(httpDownload) error = %v", err)
	}
	if string(httpDownloaded) != "download-body" {
		t.Fatalf("http downloaded body = %q, want download-body", string(httpDownloaded))
	}
}

func TestLuaEngineGoLuaVMRegistersLrHTTPModule(t *testing.T) {
	engine := NewLuaEngine(nil)
	defer engine.Close()

	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("X-Test-Header", request.Header.Get("X-Test-Header"))
		body, _ := io.ReadAll(request.Body)
		_, _ = writer.Write([]byte(request.Method + ":" + request.URL.Path + ":" + string(body)))
	}))
	defer server.Close()

	engine.RegisterModule(&lrhttp.HttpModule{})
	engine.RegisterMethod("test.httpBaseURL", "return test http base url", func() string {
		return server.URL
	}, true)

	if err := engine.ExecuteString(`
		if type(http) ~= "table" then
			error("http module was not registered")
		end
		if type(https) ~= "table" then
			error("https module was not registered")
		end
		if type(ssl) ~= "table" or type(ssl.https) ~= "table" then
			error("ssl.https module was not registered")
		end
		if require("http") ~= http then
			error("unexpected require('http') result")
		end
		if require("ssl") ~= ssl then
			error("unexpected require('ssl') result")
		end
		if require("ssl.https") ~= ssl.https then
			error("unexpected require('ssl.https') result")
		end
		local base = test.httpBaseURL()
		local headers = { ["X-Test-Header"] = "lua-vm" }
		local code, status, respHeaders, body = http.request(base .. "/http", "POST", headers, "body", 5, true)
		if code ~= 200 or status ~= "200 OK" or respHeaders["X-Test-Header"] ~= "lua-vm" or body ~= "POST:/http:body" then
			error("unexpected http.request result")
		end
		local httpsCode, _, _, httpsBody = https.request(base .. "/https", "GET", {}, "", 5, true)
		if httpsCode ~= 200 or httpsBody ~= "GET:/https:" then
			error("unexpected https.request result")
		end
		local sslCode, _, _, sslBody = ssl.https.request(base .. "/ssl", "GET", {}, "", 5, true)
		if sslCode ~= 200 or sslBody ~= "GET:/ssl:" then
			error("unexpected ssl.https.request result")
		end
	`); err != nil {
		t.Fatalf("ExecuteString() lr http module error = %v", err)
	}
}

func TestLuaEngineGoLuaVMRegistersLrFFIModule(t *testing.T) {
	engine := NewLuaEngine(nil)
	defer engine.Close()

	engine.RegisterModule(&lrffi.FfiModule{})

	if err := engine.ExecuteString(`
		if type(ffi) ~= "table" then
			error("ffi module was not registered")
		end
		if require("ffi") ~= ffi then
			error("unexpected require('ffi') result")
		end
		local cdefValue, cdefMessage = ffi.cdef("int value;")
		if cdefValue ~= nil or type(cdefMessage) ~= "string" or cdefMessage == "" then
			error("unexpected ffi.cdef result")
		end
		local loadValue, loadMessage = ffi.load("libtest.so")
		if loadValue ~= nil or type(loadMessage) ~= "string" or loadMessage == "" then
			error("unexpected ffi.load result")
		end
		if ffi.sizeof("int") ~= 0 then
			error("unexpected ffi.sizeof result")
		end
		if ffi.new("int") ~= nil then
			error("unexpected ffi.new result")
		end
	`); err != nil {
		t.Fatalf("ExecuteString() lr ffi module error = %v", err)
	}
}

func TestLuaEngineGoLuaVMRegistersLrDynamicUIModule(t *testing.T) {
	engine := NewLuaEngine(nil)
	defer engine.Close()

	engine.RegisterModule(&lrdynamicui.DynamicUIModule{ShowWarning: false})

	if err := engine.ExecuteString(`
		if type(ui) ~= "table" then
			error("ui module was not registered")
		end
		if ui.newLayout("main") ~= true then
			error("unexpected ui.newLayout result")
		end
		if ui.addButton("ok", 1, 2, 3) ~= true then
			error("unexpected ui.addButton result")
		end
		if ui.getTableViewRowData("table", 1) ~= true then
			error("unexpected ui.getTableViewRowData result")
		end
		if ui.getTableViewSelectIndex("table") ~= true then
			error("unexpected ui.getTableViewSelectIndex result")
		end
	`); err != nil {
		t.Fatalf("ExecuteString() lr dynamicui module error = %v", err)
	}
}

func TestLuaEngineGoLuaVMRegistersLrGoBridgeModule(t *testing.T) {
	engine := NewLuaEngine(nil)
	defer engine.Close()

	engine.RegisterModule(lrgobridge.NewGoBridgeModule(engine))

	if err := engine.ExecuteString(`
		if type(gobridge) ~= "table" then
			error("gobridge module was not registered")
		end
		local bytes = gobridge.tobytes("Lua")
		if bytes ~= "4c7561" then
			error("unexpected gobridge.tobytes result")
		end
		if gobridge.tostring(bytes) ~= "Lua" then
			error("unexpected gobridge.tostring result")
		end
		if gobridge.tostring("4c756") ~= "Lu" then
			error("unexpected gobridge.tostring odd hex result")
		end
		local ok, err = pcall(function()
			return gobridge.call("/path/to/missing/library.so", "missing")
		end)
		if ok or not string.find(tostring(err), "failed to load library", 1, true) then
			error("unexpected gobridge.call error result")
		end
	`); err != nil {
		t.Fatalf("ExecuteString() lr gobridge module error = %v", err)
	}
}

func TestLuaEngineGoLuaVMRegistersAutogoJSONModule(t *testing.T) {
	engine := NewLuaEngine(nil)
	defer engine.Close()

	engine.RegisterModule(&autogojson.JsonModule{})

	if err := engine.ExecuteString(`
		if type(json) ~= "table" then
			error("json module was not registered")
		end
		local encoded = json.stringify({name = "lua", count = 2})
		if encoded ~= '{"count":2,"name":"lua"}' and encoded ~= '{"name":"lua","count":2}' then
			error("unexpected json.stringify result: " .. encoded)
		end
		if json.stringifyArr("lua") ~= '["lua"]' then
			error("unexpected json.stringifyArr scalar result")
		end
		if json.stringifyObj({name = "lua"}) ~= '{"name":"lua"}' then
			error("unexpected json.stringifyObj result")
		end
		local parsed = json.parse('{"items":[1,2],"ok":true}')
		if parsed.ok ~= true or parsed.items[1] ~= 1 or parsed.items[2] ~= 2 then
			error("unexpected json.parse result")
		end
		local formatted = json.format({name = "lua"})
		if type(formatted) ~= "string" or not string.find(formatted, "\n") then
			error("unexpected json.format result")
		end
	`); err != nil {
		t.Fatalf("ExecuteString() autogo json module error = %v", err)
	}
}

func TestLuaEngineGoLuaVMRegistersAutogoCoroutineModule(t *testing.T) {
	engine := NewLuaEngine(nil)
	defer engine.Close()

	engine.RegisterModule(&autogocoroutine.CoroutineModule{})

	if err := engine.ExecuteString(`
		if type(coroutine) ~= "table" then
			error("coroutine module was not registered")
		end
		coroutine.sleep(0)
		if coroutine.cancel("missing") ~= false then
			error("unexpected coroutine.cancel missing result")
		end
		local asyncValue = coroutine.async(function()
			return "done"
		end)
		if asyncValue ~= "done" then
			error("unexpected coroutine.async result")
		end
		local asyncError = coroutine.async(function()
			error("async failed")
		end)
		if type(asyncError) ~= "string" or asyncError == "" then
			error("unexpected coroutine.async error result")
		end
		local awaited = coroutine.await({ ok = true, count = 3 })
		if type(awaited) ~= "table" or awaited.ok ~= true or awaited.count ~= 3 then
			error("unexpected coroutine.await result")
		end
		if coroutine.getActiveCoroutines() ~= 0 then
			error("unexpected coroutine.getActiveCoroutines result")
		end
		local list = coroutine.getCoroutineList()
		if type(list) ~= "table" or #list ~= 0 then
			error("unexpected coroutine.getCoroutineList result")
		end
		if coroutine.getCoroutineInfo("missing") ~= nil then
			error("unexpected coroutine.getCoroutineInfo missing result")
		end
		local stats = coroutine.getStats()
		if type(stats) ~= "table" or stats.active ~= 0 then
			error("unexpected coroutine.getStats result")
		end
		local poolName = coroutine.createPool("unit_pool", 1, 2)
		if poolName ~= "unit_pool" then
			error("unexpected coroutine.createPool result")
		end
		local pools = coroutine.listPools()
		if type(pools) ~= "table" or #pools < 1 then
			error("unexpected coroutine.listPools result")
		end
		local poolStats = coroutine.getPoolStats("unit_pool")
		if type(poolStats) ~= "table" or poolStats.name ~= "unit_pool" or poolStats.maxWorkers ~= 1 or poolStats.maxTasks ~= 2 then
			error("unexpected coroutine.getPoolStats result")
		end
		coroutine.setScheduleStrategy("priority")
		if coroutine.getScheduleStrategy() ~= "priority" then
			error("unexpected coroutine schedule strategy result")
		end
		coroutine.setPriority("job", 7)
		if coroutine.getPriority("job") ~= 7 then
			error("unexpected coroutine priority result")
		end
		if coroutine.closePool("unit_pool") ~= true then
			error("unexpected coroutine.closePool result")
		end
		if coroutine.getPoolStats("unit_pool") ~= nil then
			error("unexpected coroutine.getPoolStats closed result")
		end
	`); err != nil {
		t.Fatalf("ExecuteString() autogo coroutine module error = %v", err)
	}
}

func TestLuaEngineGoLuaVMRegistersLrImGuiModule(t *testing.T) {
	engine := NewLuaEngine(nil)
	defer engine.Close()

	engine.RegisterModule(&lrimgui.ImGuiModule{})

	if err := engine.ExecuteString(`
		if type(imgui) ~= "table" then
			error("imgui module was not registered")
		end
		if imgui.isSupport() ~= true then
			error("unexpected imgui.isSupport result")
		end
		if imgui.getLastError() ~= "" then
			error("unexpected imgui.getLastError result")
		end
		local button = imgui.createButton(1, 2, 100, 30, "ok")
		if type(button) ~= "string" or button == "" or not imgui.isValidHandle(button) then
			error("unexpected imgui.createButton result")
		end
		if imgui.getWidgetText(button) ~= "ok" then
			error("unexpected imgui.getWidgetText button result")
		end
		imgui.setWidgetText(button, "confirm")
		if imgui.getWidgetText(button) ~= "confirm" then
			error("unexpected imgui.setWidgetText result")
		end
		local checkbox = imgui.createCheckBox(button, "enabled", true)
		if imgui.isChecked(checkbox) ~= true then
			error("unexpected imgui.createCheckBox checked result")
		end
		imgui.setChecked(checkbox, false)
		if imgui.isChecked(checkbox) ~= false then
			error("unexpected imgui.setChecked result")
		end
		local color = imgui.createColorPicker(button, "color", 255, 12, 12)
		if not imgui.isValidHandle(color) then
			error("unexpected imgui.createColorPicker result")
		end
		local switch = imgui.createSwitch(button, "switch", true, 20)
		if imgui.isChecked(switch) ~= true then
			error("unexpected imgui.createSwitch result")
		end
		local input = imgui.createInputText(button, "name", "lua")
		if imgui.getInputText(input) ~= "lua" then
			error("unexpected imgui.getInputText result")
		end
		imgui.setInputType(input, 2)
		imgui.setInputText(input, "vm")
		if imgui.getInputText(input) ~= "vm" then
			error("unexpected imgui.setInputText result")
		end
		local progress = imgui.createProgressBar(button, 0.25)
		if imgui.getProgressBarPos(progress) ~= 0.25 then
			error("unexpected imgui.createProgressBar result")
		end
		imgui.setProgressBarPos(progress, 0.75)
		if imgui.getProgressBarPos(progress) ~= 0.75 then
			error("unexpected imgui.setProgressBarPos result")
		end
		local label = imgui.createLabel(button, "label")
		if imgui.getWidgetText(label) ~= "label" then
			error("unexpected imgui.createLabel result")
		end
		local combo = imgui.createComboBox(button, "one|two", 120)
		if imgui.getItemCount(combo) ~= 2 or imgui.getItemText(combo, 1) ~= "two" then
			error("unexpected imgui.createComboBox result")
		end
		imgui.addOptionItem(combo, "three")
		if imgui.getItemCount(combo) ~= 3 or imgui.getItemText(combo, 2) ~= "three" then
			error("unexpected imgui.addOptionItem result")
		end
		imgui.setItemSelected(combo, 2)
		if imgui.getSelectedItemIndex(combo) ~= 2 then
			error("unexpected imgui.setItemSelected result")
		end
		imgui.removeItemAt(combo, 1)
		if imgui.getItemCount(combo) ~= 2 or imgui.getItemText(combo, 1) ~= "three" then
			error("unexpected imgui.removeItemAt result")
		end
		imgui.removeAllItems(combo)
		if imgui.getItemCount(combo) ~= 0 or imgui.getSelectedItemIndex(combo) ~= -1 then
			error("unexpected imgui.removeAllItems result")
		end
		local radio = imgui.createRadioGroup(button, "mode")
		imgui.addRadioBox(radio, "a", false)
		imgui.addRadioBox(radio, "b", true)
		if imgui.getItemCount(radio) ~= 2 then
			error("unexpected imgui.addRadioBox result")
		end
		local slider = imgui.createSlider(button, "volume", 0, 100, 12)
		if imgui.getSlider(slider) ~= 12 then
			error("unexpected imgui.createSlider result")
		end
		imgui.setSlider(slider, 42)
		if imgui.getSlider(slider) ~= 42 then
			error("unexpected imgui.setSlider result")
		end
		local tableHandle = imgui.createTableView(button, "rows", 2, true)
		imgui.setTableHeaderItem(tableHandle, 0, "name")
		local row = imgui.insertTableRow(tableHandle, -2)
		if row ~= 0 or imgui.getItemCount(tableHandle) ~= 1 then
			error("unexpected imgui.insertTableRow result")
		end
		imgui.setTableItemText(tableHandle, row, 0, "lua")
		imgui.setTableItemText(tableHandle, row, 1, "vm")
		if imgui.getTableItemText(tableHandle, row, 0) ~= "lua" or imgui.getTableItemText(tableHandle, row, 1) ~= "vm" then
			error("unexpected imgui table text result")
		end
		imgui.deleteTableRow(tableHandle, row)
		if imgui.getItemCount(tableHandle) ~= 0 then
			error("unexpected imgui.deleteTableRow result")
		end
		imgui.insertTableRow(tableHandle, -2)
		imgui.clearTable(tableHandle)
		if imgui.getItemCount(tableHandle) ~= 0 then
			error("unexpected imgui.clearTable result")
		end
		local rect = imgui.createRectangle(0, 0, 10, 10, 0xff, true)
		if not imgui.isValidHandle(rect) or imgui.isShapeVisibility(rect) ~= true then
			error("unexpected imgui.createRectangle result")
		end
		imgui.setShapePosition(rect, 5, 6)
		imgui.setShapeVisibility(rect, false)
		if imgui.isShapeVisibility(rect) ~= false then
			error("unexpected imgui.setShapeVisibility result")
		end
		local circle = imgui.createCircle(1, 2, 3, 0xff, false)
		local line = imgui.createLine(1, 2, 3, 4, 0xff)
		if not imgui.isValidHandle(circle) or not imgui.isValidHandle(line) then
			error("unexpected imgui shape handle result")
		end
		local shapeText = imgui.createShapeText(1, 2, 30, 10, "shape", 0xff, 0, false)
		imgui.setShapeTextString(shapeText, "shape-next")
		if imgui.getWidgetText(shapeText) ~= "shape-next" then
			error("unexpected imgui.setShapeTextString result")
		end
		local bitmapShape = imgui.createBitmapShape(1, 2, 3, 4, "bitmap")
		if not imgui.isValidHandle(bitmapShape) or imgui.isShapeVisibility(bitmapShape) ~= true then
			error("unexpected imgui.createBitmapShape result")
		end
		imgui.setShapeTextColor(shapeText, 0xaa)
		imgui.setShapeTextBackground(shapeText, 0xbb, true)
		imgui.setShapeTextFontScale(shapeText, 1.5)
		imgui.setShapeThickness(rect, 2)
		imgui.setBitmapShape(rect, "bitmap-data")
		imgui.removeShape(line)
		if imgui.isValidHandle(line) ~= false then
			error("unexpected imgui.removeShape result")
		end
		imgui.setWidgetSize(button, 200, 40)
		imgui.setWidgetStyle(button, 1, 2)
		imgui.setWidgetColor(button, 0xcc)
		imgui.setWidgetVisible(button, false)
		if imgui.isWidgetVisible(button) ~= false then
			error("unexpected imgui.setWidgetVisible false result")
		end
		imgui.setWidgetVisible(button, true)
		if imgui.isWidgetVisible(button) ~= true then
			error("unexpected imgui.setWidgetVisible true result")
		end
		local window = imgui.createWindow("main", 0, 0, 320, 240)
		if imgui.isWidgetVisible(window) ~= false then
			error("unexpected imgui.createWindow visibility result")
		end
		imgui.showWindow(window)
		if imgui.isWidgetVisible(window) ~= true then
			error("unexpected imgui.showWindow result")
		end
		imgui.setWindowPos(window, 10, 20)
		local pos = imgui.getWindowPos(window)
		if type(pos) ~= "table" or pos.x ~= 10 or pos.y ~= 20 then
			error("unexpected imgui.getWindowPos result")
		end
		imgui.setWindowSize(window, 640, 480)
		imgui.setWindowFlags(window, 3)
		local vLayout = imgui.createVerticalLayout(window, 100, 200)
		local hLayout = imgui.createHorticalLayout(window, 100, 200)
		local tree = imgui.createTreeBoxLayout(window, 100, 200)
		local tab = imgui.createTabBar(window, 100, 200)
		if not imgui.isValidHandle(vLayout) or not imgui.isValidHandle(hLayout) or not imgui.isValidHandle(tree) or not imgui.isValidHandle(tab) then
			error("unexpected imgui layout result")
		end
		imgui.addTabBarItem(tab, "first")
		imgui.setLayoutBorderVisible(vLayout, true)
		local imageHandle = imgui.createImage(window, "/tmp/image.png", 32, 32)
		if not imgui.isValidHandle(imageHandle) then
			error("unexpected imgui.createImage result")
		end
		imgui.setImage(imageHandle, "/tmp/next.png")
		imgui.setImageFromBitmap(imageHandle, "bitmap")
		imgui.setColorTheme(1)
		imgui.setStyleColor(2, 0xdd)
		imgui.sameLine()
		imgui.show()
		imgui.destroyWindow(window)
		if imgui.isValidHandle(window) ~= false then
			error("unexpected imgui.destroyWindow result")
		end
		if imgui.isValidHandle("missing") ~= false then
			error("unexpected imgui.isValidHandle missing result")
		end
		if imgui.isChecked("missing") ~= nil then
			error("unexpected imgui.isChecked missing result")
		end
		imgui.close()
	`); err != nil {
		t.Fatalf("ExecuteString() lr imgui module error = %v", err)
	}
}

func TestLuaEngineGoLuaVMBytecodeExecuteAndSerialize(t *testing.T) {
	engine := NewLuaEngine(nil)
	defer engine.Close()

	if err := engine.ExecuteBytecode(nil); err == nil {
		t.Fatal("ExecuteBytecode(nil) error = nil, want error")
	}
	if err := engine.ExecuteBytecode(&Bytecode{}); err == nil {
		t.Fatal("ExecuteBytecode(empty) error = nil, want error")
	}
	if _, err := SerializeBytecode(nil); err == nil {
		t.Fatal("SerializeBytecode(nil) error = nil, want error")
	}
	if _, err := DeserializeBytecode(nil); err == nil {
		t.Fatal("DeserializeBytecode(nil) error = nil, want error")
	}
	if _, err := engine.CompileString(`if then`, "bad.lua"); err == nil {
		t.Fatal("CompileString(invalid syntax) error = nil, want error")
	}
	if _, err := engine.CompileFile(filepath.Join(t.TempDir(), "missing.lua")); err == nil {
		t.Fatal("CompileFile(missing) error = nil, want error")
	}

	var fileBytecodeExecuted atomic.Bool
	var asyncBytecodeExecuted atomic.Bool
	engine.RegisterMethod("bytecode.mark", "mark bytecode execution", func(value string) string {
		return value + "-executed"
	}, true)
	engine.RegisterMethod("bytecode.fileMark", "mark file bytecode execution", func() {
		fileBytecodeExecuted.Store(true)
	}, true)
	engine.RegisterMethod("bytecode.asyncMark", "mark async bytecode execution", func() {
		asyncBytecodeExecuted.Store(true)
	}, true)

	bytecode, err := engine.CompileString(`
		local value = bytecode.mark("chunk")
		if value ~= "chunk-executed" then
			error("unexpected bytecode value")
		end
	`, "bytecode_test.lua")
	if err != nil {
		t.Fatalf("CompileString() error = %v", err)
	}
	if len(bytecode.Chunk) == 0 {
		t.Fatal("CompileString() chunk is empty")
	}
	if err := engine.ExecuteBytecode(bytecode); err != nil {
		t.Fatalf("ExecuteBytecode() error = %v", err)
	}

	encoded, err := SerializeBytecode(bytecode)
	if err != nil {
		t.Fatalf("SerializeBytecode() error = %v", err)
	}
	decoded, err := DeserializeBytecode(encoded)
	if err != nil {
		t.Fatalf("DeserializeBytecode() error = %v", err)
	}
	if err := engine.ExecuteBytecode(decoded); err != nil {
		t.Fatalf("ExecuteBytecode(decoded) error = %v", err)
	}

	scriptPath := filepath.Join(t.TempDir(), "bytecode_file.lua")
	if err := os.WriteFile(scriptPath, []byte(`bytecode.fileMark()`), 0644); err != nil {
		t.Fatalf("WriteFile() error = %v", err)
	}
	fileBytecode, err := engine.CompileFile(scriptPath)
	if err != nil {
		t.Fatalf("CompileFile() error = %v", err)
	}
	if fileBytecode.GetName() != scriptPath {
		t.Fatalf("CompileFile() name = %q, want %q", fileBytecode.GetName(), scriptPath)
	}
	if fileBytecode.GetFunctionProto() != nil {
		t.Fatal("go-lua-vm bytecode should not expose gopher-lua FunctionProto")
	}
	bytecodePath := filepath.Join(t.TempDir(), "bytecode_file.gluac")
	if err := engine.SaveBytecodeToFile(fileBytecode, bytecodePath); err != nil {
		t.Fatalf("SaveBytecodeToFile() error = %v", err)
	}
	loadedBytecode, err := engine.LoadBytecodeFromFile(bytecodePath)
	if err != nil {
		t.Fatalf("LoadBytecodeFromFile() error = %v", err)
	}
	if err := engine.ExecuteBytecode(loadedBytecode); err != nil {
		t.Fatalf("ExecuteBytecode(loaded file) error = %v", err)
	}
	if !fileBytecodeExecuted.Load() {
		t.Fatal("loaded file bytecode did not execute")
	}

	asyncBytecode, err := engine.CompileString(`bytecode.asyncMark()`, "async_bytecode.lua")
	if err != nil {
		t.Fatalf("CompileString(async) error = %v", err)
	}
	if err := engine.ExecuteBytecodeWithMode(asyncBytecode, ExecuteModeAsync); err != nil {
		t.Fatalf("ExecuteBytecodeWithMode(async) error = %v", err)
	}
	deadline := time.Now().Add(time.Second)
	for time.Now().Before(deadline) {
		if asyncBytecodeExecuted.Load() {
			return
		}
		time.Sleep(10 * time.Millisecond)
	}
	t.Fatal("async bytecode did not execute before deadline")
}

func TestLuaEngineGoLuaVMCoreFunctions(t *testing.T) {
	engine := NewLuaEngine(nil)
	defer engine.Close()

	engine.RegisterMethod("core.echo", "echo value", func(value string) string {
		return value
	}, true)

	if err := engine.ExecuteString(`
		console.log("core", "log")
		console.error("core", "error")
		local methods = listMethods()
		if type(methods) ~= "table" or #methods == 0 then
			error("listMethods did not return method table")
		end
		if restoreMethod("core.echo") ~= true then
			error("restoreMethod should return true for registered method")
		end
		if overrideMethod("core.echo") ~= false then
			error("overrideMethod should be disabled on go-lua-vm bridge")
		end
		if registerMethod("core.dynamicDoc", "dynamic doc method", true) ~= true then
			error("registerMethod should return true")
		end
		local foundDynamicDoc = false
		for _, method in ipairs(listMethods()) do
			if method.name == "core.dynamicDoc" and method.description == "dynamic doc method" and method.overridable == true then
				foundDynamicDoc = true
			end
		end
		if not foundDynamicDoc then
			error("registerMethod was not reflected in listMethods")
		end
		if unregisterMethod("core.dynamicDoc") ~= true then
			error("unregisterMethod should return true for dynamic doc method")
		end
	`); err != nil {
		t.Fatalf("ExecuteString() core functions error = %v", err)
	}

	info, ok := GetRegistry().GetMethod("core.echo")
	if !ok {
		t.Fatal("core.echo not registered")
	}
	if info.Overridden {
		t.Fatal("core.echo overridden = true, want false")
	}
	if _, ok := GetRegistry().GetMethod("core.dynamicDoc"); ok {
		t.Fatal("core.dynamicDoc still registered after unregisterMethod")
	}
}

func TestLuaEngineGoLuaVMRegisterMethodBridgeWithMap(t *testing.T) {
	engine := NewLuaEngine(nil)
	defer engine.Close()

	engine.RegisterMethod("maps.get", "return string map", func() map[string]string {
		return map[string]string{
			"name": "luavm",
			"kind": "table",
		}
	}, true)
	engine.RegisterMethod("maps.read", "read string map", func(values map[string]string) string {
		return values["name"] + ":" + values["kind"]
	}, true)
	engine.RegisterMethod("maps.dynamic", "read dynamic map", func(values map[string]interface{}) string {
		items := values["items"].([]interface{})
		meta := values["meta"].(map[string]interface{})
		return items[0].(string) + ":" + items[1].(string) + ":" + meta["name"].(string)
	}, true)
	engine.RegisterMethod("maps.dynamicReturn", "return dynamic map", func() map[string]interface{} {
		return map[string]interface{}{
			"items": []interface{}{"x", "y"},
			"meta": map[string]interface{}{
				"name": "go",
			},
		}
	}, true)
	engine.RegisterMethod("maps.intKeys", "return int-key map", func() map[int]string {
		return map[int]string{
			1: "one",
			2: "two",
		}
	}, true)
	engine.RegisterMethod("maps.readIntKeys", "read int-key map", func(values map[int]string) string {
		return values[2] + ":" + values[4]
	}, true)

	if err := engine.ExecuteString(`
		local values = maps.get()
		if values.name ~= "luavm" or values.kind ~= "table" then
			error("unexpected map return")
		end
		local text = maps.read({ name = "lua", kind = "input" })
		if text ~= "lua:input" then
			error("unexpected map argument")
		end
		local dynamic = maps.dynamic({ items = { "a", "b" }, meta = { name = "lua" } })
		if dynamic ~= "a:b:lua" then
			error("unexpected dynamic map argument")
		end
		local dynamicReturn = maps.dynamicReturn()
		if dynamicReturn.items[2] ~= "y" or dynamicReturn.meta.name ~= "go" then
			error("unexpected dynamic map return")
		end
		local intKeys = maps.intKeys()
		if intKeys[1] ~= "one" or intKeys[2] ~= "two" then
			error("unexpected int-key map return")
		end
		local intText = maps.readIntKeys({ [2] = "two", [4] = "four" })
		if intText ~= "two:four" then
			error("unexpected int-key map argument")
		end
	`); err != nil {
		t.Fatalf("ExecuteString() map bridge error = %v", err)
	}
}

func TestLuaEngineGoLuaVMRegisterMethodBridgeWithSliceArguments(t *testing.T) {
	engine := NewLuaEngine(nil)
	defer engine.Close()

	engine.RegisterMethod("slices.join", "join string slice", func(values []string) string {
		return strings.Join(values, ",")
	}, true)
	engine.RegisterMethod("slices.sum", "sum float32 slice", func(values []float32) float32 {
		var total float32
		for _, value := range values {
			total += value
		}
		return total
	}, true)
	engine.RegisterMethod("slices.bytes", "read bytes", func(values []byte) int {
		return len(values)
	}, true)
	engine.RegisterMethod("slices.bytesReturn", "return bytes", func() []byte {
		return []byte("lua-bytes")
	}, true)

	if err := engine.ExecuteString(`
		if slices.join({ "a", "b", "c" }) ~= "a,b,c" then
			error("unexpected string slice argument")
		end
		if slices.sum({ 1.5, 2.5 }) ~= 4.0 then
			error("unexpected float slice argument")
		end
		if slices.bytes("abc") ~= 3 then
			error("unexpected byte slice argument")
		end
		if slices.bytesReturn() ~= "lua-bytes" then
			error("unexpected byte slice return")
		end
	`); err != nil {
		t.Fatalf("ExecuteString() slice bridge error = %v", err)
	}
}

func TestLuaEngineGoLuaVMRegisterMethodBridgeWithStructTables(t *testing.T) {
	engine := NewLuaEngine(nil)
	defer engine.Close()

	engine.RegisterMethod("structs.read", "read struct slice", func(items []testVMTextItem) string {
		return items[0].Text + ":" + items[0].TextColor
	}, true)
	engine.RegisterMethod("structs.list", "return struct slice", func() []testVMTextItem {
		return []testVMTextItem{
			{Text: "hello", TextColor: "#fff"},
		}
	}, true)
	engine.RegisterMethod("structs.readPointer", "read struct pointer", func(item *testVMTextItem) string {
		return item.Text + ":" + item.TextColor
	}, true)
	engine.RegisterMethod("structs.isNilPointer", "read nil struct pointer", func(item *testVMTextItem) bool {
		return item == nil
	}, true)
	engine.RegisterMethod("structs.nilPointer", "return nil struct pointer", func() *testVMTextItem {
		return nil
	}, true)
	engine.RegisterMethod("structs.readTagged", "read tagged struct", func(item testVMTaggedItem) string {
		return item.Title + ":" + item.Detail
	}, true)
	engine.RegisterMethod("structs.tagged", "return tagged struct", func() testVMTaggedItem {
		return testVMTaggedItem{
			Title:  "tagged",
			Detail: "return",
		}
	}, true)

	if err := engine.ExecuteString(`
		local read = structs.read({ { text = "lua", textColor = "#000" } })
		if read ~= "lua:#000" then
			error("unexpected struct argument")
		end
		local list = structs.list()
		if list[1].Text ~= "hello" or list[1].textColor ~= "#fff" then
			error("unexpected struct return")
		end
		local pointerRead = structs.readPointer({ text = "pointer", textColor = "#333" })
		if pointerRead ~= "pointer:#333" then
			error("unexpected struct pointer argument")
		end
		if not structs.isNilPointer(nil) then
			error("unexpected nil struct pointer argument")
		end
		if structs.nilPointer() ~= nil then
			error("unexpected nil struct pointer return")
		end
		local taggedRead = structs.readTagged({ title_text = "tagged", detail_text = "input" })
		if taggedRead ~= "tagged:input" then
			error("unexpected tagged struct argument")
		end
		local tagged = structs.tagged()
		if tagged.title_text ~= "tagged" or tagged.detail_text ~= "return" then
			error("unexpected tagged struct return")
		end
	`); err != nil {
		t.Fatalf("ExecuteString() struct bridge error = %v", err)
	}
}

func TestLuaEngineGoLuaVMCustomBridgePreservesObjectProxy(t *testing.T) {
	engine := NewLuaEngine(nil)
	defer engine.Close()

	engine.RegisterMethod("objects.new", "create object", func(name string) *testVMWidget {
		return &testVMWidget{Name: name}
	}, true)
	engine.RegisterMethod("objects.setItems", "set struct items", func(widget *testVMWidget, items []testVMTextItem) *testVMWidget {
		widget.Items = items
		return widget
	}, true)
	engine.RegisterMethod("objects.count", "count items", func(widget *testVMWidget) int {
		return len(widget.Items)
	}, true)
	engine.RegisterMethod("objects.nilWidget", "return nil object", func() *testVMWidget {
		return nil
	}, true)
	engine.RegisterMethod("objects.isNil", "check nil object", func(widget *testVMWidget) bool {
		return widget == nil
	}, true)
	engine.RegisterMethod("objects.interfaceName", "read object through interface", func(value interface{}) string {
		widget, ok := value.(*testVMWidget)
		if !ok || widget == nil {
			return ""
		}
		return widget.Name
	}, true)
	engine.RegisterMethod("objects.interfaceTable", "read table through interface", func(value interface{}) string {
		table, ok := value.(map[string]interface{})
		if !ok {
			return ""
		}
		enabled, _ := table["enabled"].(bool)
		title, _ := table["title"].(string)
		if !enabled {
			return ""
		}
		return title
	}, true)

	if err := engine.ExecuteString(`
		local widget = objects.new("panel")
		local same = objects.setItems(widget, { { text = "one", textColor = "#111" }, { text = "two", textColor = "#222" } })
		if objects.count(widget) ~= 2 or objects.count(same) ~= 2 then
			error("object proxy was not preserved")
		end
		if objects.interfaceName(widget) ~= "panel" then
			error("object proxy was not passed through interface")
		end
		if objects.interfaceTable({ title = "plain", enabled = true }) ~= "plain" then
			error("plain table was not passed through interface")
		end
		if objects.nilWidget() ~= nil then
			error("nil pointer return was not converted to nil")
		end
		if objects.isNil(nil) ~= true then
			error("nil pointer argument was not converted to nil")
		end
	`); err != nil {
		t.Fatalf("ExecuteString() object proxy bridge error = %v", err)
	}
}

func TestLuaEngineGoLuaVMRestartReinstallsRegisteredMethods(t *testing.T) {
	engine := NewLuaEngine(nil)
	defer engine.Close()

	moduleDir := t.TempDir()
	modulePath := filepath.Join(moduleDir, "restart_module.lua")
	if err := os.WriteFile(modulePath, []byte(`return { value = "kept" }`), 0644); err != nil {
		t.Fatalf("WriteFile() error = %v", err)
	}
	engine.AddRequirePath(moduleDir)

	engine.RegisterMethod("restart.echo", "echo after restart", func(value string) string {
		return value
	}, true)

	if err := engine.ExecuteString(`
		if restart.echo("before") ~= "before" then
			error("method missing before restart")
		end
		local module = require("restart_module")
		if module.value ~= "kept" then
			error("require path missing before restart")
		end
	`); err != nil {
		t.Fatalf("ExecuteString() before restart error = %v", err)
	}
	if err := engine.Restart(); err != nil {
		t.Fatalf("Restart() error = %v", err)
	}
	if err := engine.ExecuteString(`
		if restart.echo("after") ~= "after" then
			error("method missing after restart")
		end
		local module = require("restart_module")
		if module.value ~= "kept" then
			error("require path missing after restart")
		end
	`); err != nil {
		t.Fatalf("ExecuteString() after restart error = %v", err)
	}
}

func TestLuaEngineGoLuaVMUnregisterRemovesModuleMethod(t *testing.T) {
	engine := NewLuaEngine(nil)
	defer engine.Close()

	engine.RegisterMethod("removable.echo", "echo before unregister", func(value string) string {
		return value
	}, true)

	if err := engine.ExecuteString(`
		if removable.echo("ok") ~= "ok" then
			error("method missing before unregister")
		end
		if unregisterMethod("removable.echo") ~= true then
			error("unregisterMethod should return true")
		end
		if removable.echo ~= nil then
			error("method still exists after unregister")
		end
	`); err != nil {
		t.Fatalf("ExecuteString() unregister error = %v", err)
	}
	if _, ok := GetRegistry().GetMethod("removable.echo"); ok {
		t.Fatal("removable.echo still exists in registry")
	}
}

func TestLuaEngineGoLuaVMRequireUsesSearchPaths(t *testing.T) {
	engine := NewLuaEngine(nil)
	defer engine.Close()

	moduleDir := t.TempDir()
	modulePath := filepath.Join(moduleDir, "sample_module.lua")
	if err := os.WriteFile(modulePath, []byte(`return { value = "loaded" }`), 0644); err != nil {
		t.Fatalf("WriteFile() error = %v", err)
	}

	if err := engine.ExecuteString(`
		local sample = require("sample_module")
		if sample.value ~= "loaded" then
			error("module was not loaded from search path")
		end
	`, moduleDir); err != nil {
		t.Fatalf("ExecuteString() require search path error = %v", err)
	}
}

func TestLuaEngineGoLuaVMExecuteFileUsesScriptDirectoryForRequire(t *testing.T) {
	engine := NewLuaEngine(nil)
	defer engine.Close()

	scriptDir := t.TempDir()
	modulePath := filepath.Join(scriptDir, "file_module.lua")
	if err := os.WriteFile(modulePath, []byte(`return { value = "file-loaded" }`), 0644); err != nil {
		t.Fatalf("WriteFile(module) error = %v", err)
	}
	scriptPath := filepath.Join(scriptDir, "main.lua")
	if err := os.WriteFile(scriptPath, []byte(`
		local module = require("file_module")
		if module.value ~= "file-loaded" then
			error("module was not loaded from script directory")
		end
	`), 0644); err != nil {
		t.Fatalf("WriteFile(script) error = %v", err)
	}

	if err := engine.ExecuteFile(scriptPath); err != nil {
		t.Fatalf("ExecuteFile() error = %v", err)
	}
}

func TestLuaEngineGoLuaVMRequirePathDoesNotDuplicate(t *testing.T) {
	engine := NewLuaEngine(nil)
	defer engine.Close()

	engine.AddRequirePath("dup_path")
	engine.AddRequirePath("dup_path")
	engine.setupPackagePath()

	if err := engine.ExecuteString(`
		local target = "dup_path/?.lua"
		local first = package.path:find(target, 1, true)
		if not first then
			error("require path was not added")
		end
		local second = package.path:find(target, first + #target, true)
		if second then
			error("require path was duplicated")
		end
	`); err != nil {
		t.Fatalf("ExecuteString() duplicate require path error = %v", err)
	}
}

func TestLuaEngineGoLuaVMRegisterMethodBridgeWithCallbackArgument(t *testing.T) {
	engine := NewLuaEngine(nil)
	defer engine.Close()

	called := false
	engine.RegisterMethod("callbacks.call", "call lua callback", func(callback func(string) string) string {
		return callback("go")
	}, true)
	engine.RegisterMethod("callbacks.notify", "call lua notify", func(callback func(int)) {
		callback(7)
		called = true
	}, true)
	engine.RegisterMethod("callbacks.multi", "call lua multi-result callback", func(callback func(string) (string, int, error)) string {
		text, count, err := callback("go")
		if err != nil {
			return "error:" + err.Error()
		}
		return text + ":" + strconv.Itoa(count)
	}, true)
	engine.RegisterMethod("callbacks.captureError", "capture lua callback error", func(callback func() (string, error)) string {
		text, err := callback()
		if err != nil {
			return "error"
		}
		return text
	}, true)

	if err := engine.ExecuteString(`
		local value = callbacks.call(function(input)
			return input .. "-lua"
		end)
		if value ~= "go-lua" then
			error("unexpected callback result")
		end
		local got = 0
		callbacks.notify(function(value)
			got = value
		end)
		if got ~= 7 then
			error("unexpected callback notify value")
		end
		local multi = callbacks.multi(function(input)
			return input .. "-multi", 3
		end)
		if multi ~= "go-multi:3" then
			error("unexpected callback multi result")
		end
		local captured = callbacks.captureError(function()
			error("callback failed")
		end)
		if captured ~= "error" then
			error("unexpected callback error result")
		end
	`); err != nil {
		t.Fatalf("ExecuteString() callback bridge error = %v", err)
	}
	if !called {
		t.Fatal("callback notify wrapper was not called")
	}
}

func TestLuaEngineGoLuaVMRegisterMethodBridgeWithVariadicArguments(t *testing.T) {
	engine := NewLuaEngine(nil)
	defer engine.Close()

	engine.RegisterMethod("variadic.join", "join variadic values", func(prefix string, values ...interface{}) string {
		parts := []string{prefix}
		for _, value := range values {
			parts = append(parts, value.(string))
		}
		return strings.Join(parts, ":")
	}, true)
	engine.RegisterMethod("variadic.count", "count variadic values", func(values ...interface{}) int {
		return len(values)
	}, true)
	engine.RegisterMethod("variadic.describe", "describe variadic table values", func(values ...interface{}) string {
		arrayValue := values[0].([]interface{})
		mapValue := values[1].(map[string]interface{})
		return arrayValue[0].(string) + ":" + arrayValue[1].(string) + ":" + mapValue["name"].(string)
	}, true)

	if err := engine.ExecuteString(`
		if variadic.join("root", "a", "b") ~= "root:a:b" then
			error("unexpected variadic join")
		end
		if variadic.count() ~= 0 or variadic.count("x", "y", "z") ~= 3 then
			error("unexpected variadic count")
		end
		if variadic.describe({ "a", "b" }, { name = "lua" }) ~= "a:b:lua" then
			error("unexpected variadic table conversion")
		end
	`); err != nil {
		t.Fatalf("ExecuteString() variadic bridge error = %v", err)
	}
}

type testVMTextItem struct {
	TextColor string
	Text      string
}

type testVMTaggedItem struct {
	Title  string `lua:"title_text"`
	Detail string `json:"detail_text"`
}

type testVMWidget struct {
	Name  string
	Items []testVMTextItem
}

type testVMModule struct{}

func (testVMModule) Name() string {
	return "vmtest"
}

func (testVMModule) Register(engine model.Engine) error {
	engine.RegisterMethod("vmtest.concat", "concat strings", func(left, right string) string {
		return left + "-" + right
	}, true)
	return nil
}

func (testVMModule) IsAvailable() bool {
	return true
}
