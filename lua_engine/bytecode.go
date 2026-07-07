package lua_engine

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
	"os"

	gluaBytecode "github.com/ZingYao/go-lua-vm/bytecode"
	"github.com/ZingYao/go-lua-vm/compiler/codegen"
	"github.com/ZingYao/go-lua-vm/compiler/parser"
	glua "github.com/ZingYao/go-lua-vm/lua"
)

// Bytecode 表示编译后的 Lua 字节码
// 当前可执行内容使用 go-lua-vm 的 Lua 5.3 binary chunk，Proto 仅保留旧 API 兼容。
type Bytecode struct {
	// Chunk 是 go-lua-vm 可直接加载执行的 Lua 5.3 binary chunk。
	Chunk []byte
	// Source 保存原始源码，便于后续调试和旧序列化兼容。
	Source string
	// Name 是字节码的名称（通常是文件名或脚本名）
	Name string
}

// CompileString 将 Lua 源码字符串编译为字节码
// source: Lua 源码字符串
// name: 可选参数，字节码的名称，用于调试和错误信息
// 返回编译后的字节码对象和可能的错误
func (e *LuaEngine) CompileString(source string, name ...string) (*Bytecode, error) {
	e.mu.RLock()
	defer e.mu.RUnlock()

	if e.vmState == nil {
		return nil, fmt.Errorf("Lua engine not initialized")
	}

	// 设置默认名称
	scriptName := "string"
	if len(name) > 0 && name[0] != "" {
		scriptName = name[0]
	}

	// 使用 go-lua-vm parser 解析源码，确保预编译阶段暴露语法错误。
	chunk, err := parser.NewWithSyntax(source, e.vmState.Options().SyntaxExtensions).ParseChunk()
	if err != nil {
		return nil, fmt.Errorf("failed to parse Lua source: %v", err)
	}

	// 编译为 Lua 5.3 Proto 并 dump 成 binary chunk，执行时不再回到 gopher-lua。
	proto, err := codegen.CompileChunk(chunk, scriptName)
	if err != nil {
		return nil, fmt.Errorf("failed to compile Lua source: %v", err)
	}

	return &Bytecode{
		Chunk:  gluaBytecode.DumpBinaryChunk(proto),
		Source: source,
		Name:   scriptName,
	}, nil
}

// CompileFile 将 Lua 源码文件编译为字节码
// path: Lua 源码文件路径
// 返回编译后的字节码对象和可能的错误
func (e *LuaEngine) CompileFile(path string) (*Bytecode, error) {
	e.mu.RLock()
	defer e.mu.RUnlock()

	if e.vmState == nil {
		return nil, fmt.Errorf("Lua engine not initialized")
	}

	// 读取文件内容
	var content []byte
	var err error

	if e.config.FileSystem != nil {
		// 从配置的文件系统读取
		file, err := e.config.FileSystem.Open(path)
		if err != nil {
			return nil, fmt.Errorf("failed to open file '%s': %v", path, err)
		}
		defer file.Close()

		content, err = io.ReadAll(file)
		if err != nil {
			return nil, fmt.Errorf("failed to read file '%s': %v", path, err)
		}
	} else {
		// 从操作系统文件系统读取
		content, err = os.ReadFile(path)
		if err != nil {
			return nil, fmt.Errorf("failed to read file '%s': %v", path, err)
		}
	}

	// 使用文件名作为字节码名称
	return e.CompileString(string(content), path)
}

// ExecuteBytecode 执行预编译的字节码
// bytecode: 编译后的字节码对象
// 返回执行结果和可能的错误
func (e *LuaEngine) ExecuteBytecode(bytecode *Bytecode) error {
	return e.ExecuteBytecodeWithMode(bytecode, e.config.ExecuteMode)
}

// ExecuteBytecodeWithMode 带执行模式的执行字节码
// bytecode: 编译后的字节码对象
// mode: 执行模式（同步或异步）
func (e *LuaEngine) ExecuteBytecodeWithMode(bytecode *Bytecode, mode ExecuteMode) error {
	if bytecode == nil || len(bytecode.Chunk) == 0 {
		return fmt.Errorf("invalid bytecode: nil")
	}

	e.Start()

	// 异步执行
	if mode == ExecuteModeAsync {
		go func() {
			e.executeBytecodeLoop(bytecode)
		}()
		return nil
	}

	// 同步执行
	return e.executeBytecodeLoop(bytecode)
}

// executeBytecodeLoop 执行字节码循环
func (e *LuaEngine) executeBytecodeLoop(bytecode *Bytecode) error {
	for {
		err := e.executeBytecodeOnce(bytecode)

		// 如果脚本异常退出，打印错误日志
		if err != nil {
			fmt.Printf("脚本异常退出: %v\n", err)
			e.Stop()
			return err
		}

		// 如果跳过退出动作（os.exit(-1)），直接返回
		if e.skipExitAction {
			e.Stop()
			return nil
		}

		// 根据配置的退出动作执行相应操作
		switch e.config.OnExit {
		case ExitActionNone:
			// 无动作，直接退出
			e.Stop()
			return nil
		case ExitActionRestart:
			// 重启脚本
			fmt.Println("脚本正常退出，正在重新启动...")
			// 继续循环，重新执行脚本
		case ExitActionCustom:
			// 执行自定义退出动作
			if e.config.CustomExitAction != nil {
				e.config.CustomExitAction()
			}
			e.Stop()
			return nil
		}
	}
}

// executeBytecodeOnce 执行一次字节码
func (e *LuaEngine) executeBytecodeOnce(bytecode *Bytecode) error {
	e.mu.Lock()
	defer e.mu.Unlock()

	if e.vmState == nil {
		return fmt.Errorf("Lua engine not initialized")
	}

	// 从 go-lua-vm binary chunk 加载 closure，并在 protected call 边界内执行。
	chunkName := bytecode.Name
	if chunkName == "" {
		chunkName = "=(bytecode)"
	}
	err := glua.ProtectedCall(e.vmState, func(callState *glua.State) error {
		if err := glua.LoadString(callState, string(bytecode.Chunk), chunkName); err != nil {
			return err
		}
		closureValue, err := callState.Pop()
		if err != nil {
			return err
		}
		_, err = glua.Call(callState, closureValue)
		return err
	})
	return e.handleLuaVMError(err)
}

// SerializeBytecode 将字节码序列化为字节数组
// 这允许将编译后的字节码保存到文件或传输
// bytecode: 要序列化的字节码对象
// 返回序列化后的字节数组和可能的错误
func SerializeBytecode(bytecode *Bytecode) ([]byte, error) {
	if bytecode == nil || len(bytecode.Chunk) == 0 {
		return nil, fmt.Errorf("invalid bytecode: nil")
	}

	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)

	// 序列化字节码
	err := encoder.Encode(struct {
		Name   string
		Chunk  []byte
		Source string
	}{
		Name:   bytecode.Name,
		Chunk:  bytecode.Chunk,
		Source: bytecode.Source,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to serialize bytecode: %v", err)
	}

	return buf.Bytes(), nil
}

// DeserializeBytecode 从字节数组反序列化字节码
// 注意：由于 gopher-lua 的 FunctionProto 不支持直接序列化
// 这个函数主要用于演示，实际使用时需要重新编译源码
// data: 序列化的字节数组
// 返回字节码对象和可能的错误
func DeserializeBytecode(data []byte) (*Bytecode, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("empty data")
	}

	var result struct {
		Name   string
		Chunk  []byte
		Source string
	}

	buf := bytes.NewBuffer(data)
	decoder := gob.NewDecoder(buf)

	err := decoder.Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("failed to deserialize bytecode: %v", err)
	}

	return &Bytecode{
		Name:   result.Name,
		Chunk:  result.Chunk,
		Source: result.Source,
	}, nil
}

// SaveBytecodeToFile 将字节码保存到文件
// 保存的内容包含 go-lua-vm Lua 5.3 binary chunk，后续可直接加载执行。
// bytecode: 要保存的字节码对象
// path: 目标文件路径
func (e *LuaEngine) SaveBytecodeToFile(bytecode *Bytecode, path string) error {
	if bytecode == nil {
		return fmt.Errorf("invalid bytecode: nil")
	}

	// 序列化字节码
	data, err := SerializeBytecode(bytecode)
	if err != nil {
		return fmt.Errorf("failed to serialize bytecode: %v", err)
	}

	// 写入文件
	err = os.WriteFile(path, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write bytecode file: %v", err)
	}

	return nil
}

// LoadBytecodeFromFile 从文件加载字节码
// 加载结果包含 go-lua-vm Lua 5.3 binary chunk，可直接传给 ExecuteBytecode。
// path: 字节码文件路径
func (e *LuaEngine) LoadBytecodeFromFile(path string) (*Bytecode, error) {
	// 读取文件
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read bytecode file: %v", err)
	}

	// 反序列化字节码
	bytecode, err := DeserializeBytecode(data)
	if err != nil {
		return nil, fmt.Errorf("failed to deserialize bytecode: %v", err)
	}

	return bytecode, nil
}

// GetName 获取字节码的名称
func (b *Bytecode) GetName() string {
	if b == nil {
		return ""
	}
	return b.Name
}

// GetFunctionProto 兼容旧 API。
// go-lua-vm 不暴露 gopher-lua FunctionProto，迁移后固定返回 nil。
func (b *Bytecode) GetFunctionProto() interface{} {
	return nil
}

// CompileString 全局函数：将 Lua 源码字符串编译为字节码
func CompileString(source string, name ...string) (*Bytecode, error) {
	if engine != nil {
		return engine.CompileString(source, name...)
	}
	return nil, fmt.Errorf("Lua engine not initialized")
}

// CompileFile 全局函数：将 Lua 源码文件编译为字节码
func CompileFile(path string) (*Bytecode, error) {
	if engine != nil {
		return engine.CompileFile(path)
	}
	return nil, fmt.Errorf("Lua engine not initialized")
}

// ExecuteBytecode 全局函数：执行预编译的字节码
func ExecuteBytecode(bytecode *Bytecode) error {
	if engine != nil {
		return engine.ExecuteBytecode(bytecode)
	}
	return fmt.Errorf("Lua engine not initialized")
}
