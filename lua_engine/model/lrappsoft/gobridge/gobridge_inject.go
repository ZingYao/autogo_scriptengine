//go:build ignore
// +build ignore

package gobridge

/*
#cgo LDFLAGS: -ldl
#include <dlfcn.h>
#include <stdlib.h>
#include <string.h>

// 定义函数调用包装器
typedef void (*void_func_t)();
typedef int (*int_func_t)(int);
typedef int (*int_int_func_t)(int, int);
typedef char* (*string_func_t)(char*);
typedef char* (*bytes_func_t)(void*, size_t);

void call_void(void* ptr) {
    ((void_func_t)ptr)();
}

int call_int_int(void* ptr, int arg1) {
    return ((int_func_t)ptr)(arg1);
}

int call_int_int_int(void* ptr, int arg1, int arg2) {
    return ((int_int_func_t)ptr)(arg1, arg2);
}

char* call_string(void* ptr, char* arg) {
    return ((string_func_t)ptr)(arg);
}

char* call_bytes(void* ptr, void* data, size_t len) {
    return ((bytes_func_t)ptr)(data, len);
}
*/
import "C"
import (
	"encoding/hex"
	"fmt"
	"sync"
	"unsafe"

	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"
	lua "github.com/yuin/gopher-lua"
)

// GoBridgeModule gobridge 模块
type GoBridgeModule struct {
	engine      model.Engine
	callbacks   map[string]lua.LValue
	callbackMux sync.RWMutex
	libraries   map[string]unsafe.Pointer
	libMux      sync.RWMutex
}

// NewGoBridgeModule 创建 gobridge 模块实例
func NewGoBridgeModule(engine model.Engine) *GoBridgeModule {
	return &GoBridgeModule{
		engine:    engine,
		callbacks: make(map[string]lua.LValue),
		libraries: make(map[string]unsafe.Pointer),
	}
}

// Name 返回模块名称
func (m *GoBridgeModule) Name() string {
	return "gobridge"
}

// IsAvailable 返回模块是否可用
func (m *GoBridgeModule) IsAvailable() bool {
	return true
}

// Register 向引擎注册模块的方法
func (m *GoBridgeModule) Register(engine model.Engine) error {
	m.engine = engine
	state := engine.GetState()

	// 初始化 maps（如果尚未初始化）
	if m.callbacks == nil {
		m.callbacks = make(map[string]lua.LValue)
	}
	if m.libraries == nil {
		m.libraries = make(map[string]unsafe.Pointer)
	}

	// 创建 gobridge 表
	gobridgeTable := state.NewTable()
	state.SetGlobal("gobridge", gobridgeTable)

	// 注册 gobridge.register - 注册Lua回调函数
	gobridgeTable.RawSetString("register", state.NewFunction(func(L *lua.LState) int {
		funcName := L.CheckString(1)
		callback := L.CheckFunction(2)

		m.callbackMux.Lock()
		m.callbacks[funcName] = callback
		m.callbackMux.Unlock()

		return 0
	}))

	// 注册 gobridge.call - 调用Go动态库中的函数
	gobridgeTable.RawSetString("call", state.NewFunction(func(L *lua.LState) int {
		libPath := L.CheckString(1)
		funcName := L.CheckString(2)

		numArgs := L.GetTop() - 2
		args := make([]interface{}, numArgs)
		for i := 0; i < numArgs; i++ {
			args[i] = LuaValueToGo(L.CheckAny(i + 3))
		}

		results, err := m.callSoFunction(libPath, funcName, args...)
		if err != nil {
			L.Push(lua.LNil)
			L.Push(lua.LString(err.Error()))
			return 2
		}

		for _, result := range results {
			L.Push(GoValueToLua(result))
		}

		return len(results)
	}))

	// 注册 gobridge.tobytes - 将字符串转换为字节数组
	gobridgeTable.RawSetString("tobytes", state.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LString(bytesToHex(L.CheckString(1))))
		return 1
	}))

	// 注册 gobridge.tostring - 将字节数组转换为字符串
	gobridgeTable.RawSetString("tostring", state.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LString(hexToString(L.CheckString(1))))
		return 1
	}))

	engine.RegisterMethod("gobridge.tobytes", "将字符串转换为十六进制字节串", bytesToHex, true)
	engine.RegisterMethod("gobridge.tostring", "将十六进制字节串转换为字符串", hexToString, true)
	engine.RegisterMethod("gobridge.call", "调用动态库函数", m.callSoFunctionValue, true)

	return nil
}

// register 注册 Lua 回调函数
// 参数: funcname (string) - 注册的函数名称
//
//	function (function) - Lua回调函数
func (m *GoBridgeModule) register(L *lua.LState) int {
	funcName := L.CheckString(1)
	callback := L.CheckFunction(2)

	m.callbackMux.Lock()
	m.callbacks[funcName] = callback
	m.callbackMux.Unlock()

	return 0
}

// CallLuaCallback 从 Go 代码调用已注册的 Lua 回调函数
// 参数: funcName (string) - 函数名称
//
//	args (...interface{}) - 传递给 Lua 函数的参数
//
// 返回值: ([]interface{}, error) - Lua 函数的返回值和错误
func (m *GoBridgeModule) CallLuaCallback(funcName string, args ...interface{}) ([]interface{}, error) {
	m.callbackMux.RLock()
	callback, exists := m.callbacks[funcName]
	m.callbackMux.RUnlock()

	if !exists {
		return nil, fmt.Errorf("callback %s not found", funcName)
	}

	state := m.engine.GetState()

	state.Push(callback)

	for _, arg := range args {
		state.Push(GoValueToLua(arg))
	}

	state.Call(len(args), lua.MultRet)

	results := make([]interface{}, 0)
	for {
		result := state.Get(-1)
		if result == lua.LNil {
			state.Pop(1)
			break
		}
		results = append([]interface{}{LuaValueToGo(result)}, results...)
		state.Pop(1)
	}

	return results, nil
}

// call 调用 Go 动态库中的函数
// 参数: libpath (string) - Go编译的动态库路径(如"libgo.so")
//
//	funcname (string) - Go导出的函数名
//	... (any) - 传递给Go函数的参数(可变参数)
//
// 返回值: (any) - 返回Go函数的执行结果(多种类型)
func (m *GoBridgeModule) call(L *lua.LState) int {
	libPath := L.CheckString(1)
	funcName := L.CheckString(2)

	numArgs := L.GetTop() - 2
	args := make([]interface{}, numArgs)
	for i := 0; i < numArgs; i++ {
		args[i] = LuaValueToGo(L.CheckAny(i + 3))
	}

	results, err := m.callSoFunction(libPath, funcName, args...)
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	for _, result := range results {
		L.Push(GoValueToLua(result))
	}

	return len(results)
}

// callSoFunction 调用 .so 文件中的函数
// 参数: libPath (string) - 动态库路径
//
//	funcName (string) - 函数名称
//	args (...interface{}) - 函数参数
//
// 返回值: ([]interface{}, error) - 函数返回值和错误
func (m *GoBridgeModule) callSoFunction(libPath, funcName string, args ...interface{}) ([]interface{}, error) {
	m.libMux.Lock()
	handle, exists := m.libraries[libPath]
	if !exists {
		cLibPath := C.CString(libPath)
		defer C.free(unsafe.Pointer(cLibPath))

		handle = C.dlopen(cLibPath, C.RTLD_LAZY)
		if handle == nil {
			m.libMux.Unlock()
			return nil, fmt.Errorf("failed to load library %s: %s", libPath, C.GoString(C.dlerror()))
		}
		m.libraries[libPath] = handle
	}
	m.libMux.Unlock()

	cFuncName := C.CString(funcName)
	defer C.free(unsafe.Pointer(cFuncName))

	funcPtr := C.dlsym(handle, cFuncName)
	if funcPtr == nil {
		return nil, fmt.Errorf("failed to find function %s in library %s: %s", funcName, libPath, C.GoString(C.dlerror()))
	}

	return m.callFunction(funcPtr, args...)
}

func (m *GoBridgeModule) callSoFunctionValue(libPath, funcName string, args ...interface{}) (interface{}, error) {
	results, err := m.callSoFunction(libPath, funcName, args...)
	if err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return nil, nil
	}
	if len(results) == 1 {
		return results[0], nil
	}
	return results, nil
}

// callFunction 调用 C 函数指针
// 参数: funcPtr (unsafe.Pointer) - 函数指针
//
//	args (...interface{}) - 函数参数
//
// 返回值: ([]interface{}, error) - 函数返回值和错误
func (m *GoBridgeModule) callFunction(funcPtr unsafe.Pointer, args ...interface{}) ([]interface{}, error) {
	if len(args) == 0 {
		C.call_void(funcPtr)
		return []interface{}{}, nil
	}

	if arg1, ok := bridgeIntArg(args[0]); ok {
		if len(args) == 1 {
			result := C.call_int_int(funcPtr, C.int(arg1))
			return []interface{}{int(result)}, nil
		}
		if len(args) == 2 {
			if arg2, ok := bridgeIntArg(args[1]); ok {
				result := C.call_int_int_int(funcPtr, C.int(arg1), C.int(arg2))
				return []interface{}{int(result)}, nil
			}
		}
	}

	switch v := args[0].(type) {
	case string:
		if len(args) == 1 {
			cStr := C.CString(v)
			defer C.free(unsafe.Pointer(cStr))
			result := C.call_string(funcPtr, cStr)
			return []interface{}{C.GoString(result)}, nil
		}
	case []byte:
		if len(args) == 1 {
			if len(v) > 0 {
				cBytes := C.CBytes(v)
				defer C.free(cBytes)
				result := C.call_bytes(funcPtr, cBytes, C.size_t(len(v)))
				return []interface{}{C.GoString(result)}, nil
			}
		}
	}

	return nil, fmt.Errorf("unsupported function signature or argument types")
}

func bridgeIntArg(value interface{}) (int, bool) {
	switch v := value.(type) {
	case int:
		return v, true
	case int8:
		return int(v), true
	case int16:
		return int(v), true
	case int32:
		return int(v), true
	case int64:
		return int(v), true
	case uint:
		return int(v), true
	case uint8:
		return int(v), true
	case uint16:
		return int(v), true
	case uint32:
		return int(v), true
	case uint64:
		return int(v), true
	case float64:
		return int(v), true
	default:
		return 0, false
	}
}

// tobytes 将字符串转换为字节数组
// 参数: str (string) - 要转换的字符串
// 返回值: (string) - 字节数组的十六进制表示
func (m *GoBridgeModule) tobytes(L *lua.LState) int {
	L.Push(lua.LString(bytesToHex(L.CheckString(1))))
	return 1
}

// tostring 将字节数组转换为字符串
// 参数: hexStr (string) - 字节数组的十六进制表示
// 返回值: (string) - 转换后的字符串
func (m *GoBridgeModule) tostring(L *lua.LState) int {
	L.Push(lua.LString(hexToString(L.CheckString(1))))
	return 1
}

func bytesToHex(str string) string {
	return hex.EncodeToString([]byte(str))
}

func hexToString(hexStr string) string {
	bytes, err := hex.DecodeString(evenHexPrefix(hexStr))
	if err != nil {
		return ""
	}
	return string(bytes)
}

func evenHexPrefix(hexStr string) string {
	if len(hexStr)%2 == 0 {
		return hexStr
	}
	return hexStr[:len(hexStr)-1]
}

// GoValueToLua 将 Go 值转换为 Lua 值
func GoValueToLua(value interface{}) lua.LValue {
	switch v := value.(type) {
	case nil:
		return lua.LNil
	case bool:
		if v {
			return lua.LTrue
		}
		return lua.LFalse
	case int:
		return lua.LNumber(v)
	case int64:
		return lua.LNumber(v)
	case float64:
		return lua.LNumber(v)
	case string:
		return lua.LString(v)
	case []byte:
		return lua.LString(string(v))
	default:
		return lua.LNil
	}
}

// LuaValueToGo 将 Lua 值转换为 Go 值
func LuaValueToGo(value lua.LValue) interface{} {
	switch value.Type() {
	case lua.LTNil:
		return nil
	case lua.LTBool:
		return lua.LVAsBool(value)
	case lua.LTNumber:
		return float64(lua.LVAsNumber(value))
	case lua.LTString:
		return lua.LVAsString(value)
	default:
		return nil
	}
}
