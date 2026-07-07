//go:build ignore
// +build ignore

package files

import (
	"fmt"

	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"

	autogofiles "github.com/Dasongzi1366/AutoGo/files"
	lua "github.com/yuin/gopher-lua"
)

// FilesModule iOS files 模块。
type FilesModule struct{}

// Name 返回模块名称。
func (m *FilesModule) Name() string {
	return "files"
}

// IsAvailable 检查模块是否可用。
func (m *FilesModule) IsAvailable() bool {
	return true
}

// bytesFromTable 将 Lua 数组表转换为 Go 字节切片。
func bytesFromTable(table *lua.LTable) []byte {
	bytes := make([]byte, table.Len())
	for i := 1; i <= table.Len(); i++ {
		value := table.RawGetInt(i)
		if value != lua.LNil {
			bytes[i-1] = byte(lua.LVAsNumber(value))
		}
	}
	return bytes
}

// pushStringSlice 将字符串切片按 Lua 数组形式返回。
func pushStringSlice(L *lua.LState, values []string) {
	table := L.NewTable()
	for index, value := range values {
		table.RawSetInt(index+1, lua.LString(value))
	}
	L.Push(table)
}

// pushByteSlice 将字节切片按 Lua 数组形式返回。
func pushByteSlice(L *lua.LState, values []byte) {
	if values == nil {
		L.Push(lua.LNil)
		return
	}
	table := L.NewTable()
	for index, value := range values {
		table.RawSetInt(index+1, lua.LNumber(value))
	}
	L.Push(table)
}

// Register 向引擎注册 iOS files 方法。
func (m *FilesModule) Register(engine model.Engine) error {
	state := engine.GetState()
	filesObj := state.NewTable()
	state.SetGlobal("files", filesObj)

	filesObj.RawSetString("isFile", state.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LBool(autogofiles.IsFile(L.CheckString(1))))
		return 1
	}))
	filesObj.RawSetString("isDir", state.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LBool(autogofiles.IsDir(L.CheckString(1))))
		return 1
	}))
	filesObj.RawSetString("isEmptyDir", state.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LBool(autogofiles.IsEmptyDir(L.CheckString(1))))
		return 1
	}))
	filesObj.RawSetString("create", state.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LBool(autogofiles.Create(L.CheckString(1))))
		return 1
	}))
	filesObj.RawSetString("exists", state.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LBool(autogofiles.Exists(L.CheckString(1))))
		return 1
	}))
	filesObj.RawSetString("ensureDir", state.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LBool(autogofiles.EnsureDir(L.CheckString(1))))
		return 1
	}))
	filesObj.RawSetString("read", state.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LString(autogofiles.Read(L.CheckString(1))))
		return 1
	}))
	filesObj.RawSetString("readBytes", state.NewFunction(func(L *lua.LState) int {
		pushByteSlice(L, autogofiles.ReadBytes(L.CheckString(1)))
		return 1
	}))
	filesObj.RawSetString("write", state.NewFunction(func(L *lua.LState) int {
		autogofiles.Write(L.CheckString(1), L.CheckString(2))
		return 0
	}))
	filesObj.RawSetString("writeBytes", state.NewFunction(func(L *lua.LState) int {
		autogofiles.WriteBytes(L.CheckString(1), bytesFromTable(L.CheckTable(2)))
		return 0
	}))
	filesObj.RawSetString("append", state.NewFunction(func(L *lua.LState) int {
		autogofiles.Append(L.CheckString(1), L.CheckString(2))
		return 0
	}))
	filesObj.RawSetString("appendBytes", state.NewFunction(func(L *lua.LState) int {
		autogofiles.AppendBytes(L.CheckString(1), bytesFromTable(L.CheckTable(2)))
		return 0
	}))
	filesObj.RawSetString("copy", state.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LBool(autogofiles.Copy(L.CheckString(1), L.CheckString(2))))
		return 1
	}))
	filesObj.RawSetString("move", state.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LBool(autogofiles.Move(L.CheckString(1), L.CheckString(2))))
		return 1
	}))
	filesObj.RawSetString("rename", state.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LBool(autogofiles.Rename(L.CheckString(1), L.CheckString(2))))
		return 1
	}))
	filesObj.RawSetString("getName", state.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LString(autogofiles.GetName(L.CheckString(1))))
		return 1
	}))
	filesObj.RawSetString("getNameWithoutExtension", state.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LString(autogofiles.GetNameWithoutExtension(L.CheckString(1))))
		return 1
	}))
	filesObj.RawSetString("getExtension", state.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LString(autogofiles.GetExtension(L.CheckString(1))))
		return 1
	}))
	filesObj.RawSetString("getMd5", state.NewFunction(func(L *lua.LState) int {
		L.RaiseError("files.getMd5 is unavailable in remote AutoGo/files")
		return 1
	}))
	filesObj.RawSetString("remove", state.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LBool(autogofiles.Remove(L.CheckString(1))))
		return 1
	}))
	filesObj.RawSetString("path", state.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LString(autogofiles.Path(L.CheckString(1))))
		return 1
	}))
	filesObj.RawSetString("listDir", state.NewFunction(func(L *lua.LState) int {
		pushStringSlice(L, autogofiles.ListDir(L.CheckString(1)))
		return 1
	}))

	engine.RegisterMethod("files.isFile", "返回路径是否是文件", autogofiles.IsFile, true)
	engine.RegisterMethod("files.isDir", "返回路径是否是文件夹", autogofiles.IsDir, true)
	engine.RegisterMethod("files.isEmptyDir", "返回文件夹是否为空", autogofiles.IsEmptyDir, true)
	engine.RegisterMethod("files.create", "创建一个文件或文件夹", autogofiles.Create, true)
	engine.RegisterMethod("files.exists", "返回路径处的文件是否存在", autogofiles.Exists, true)
	engine.RegisterMethod("files.ensureDir", "确保路径所在文件夹存在", autogofiles.EnsureDir, true)
	engine.RegisterMethod("files.read", "读取文本文件内容", autogofiles.Read, true)
	engine.RegisterMethod("files.readBytes", "读取文件字节内容", autogofiles.ReadBytes, true)
	engine.RegisterMethod("files.write", "写入文本文件", autogofiles.Write, true)
	engine.RegisterMethod("files.writeBytes", "写入文件字节内容", autogofiles.WriteBytes, true)
	engine.RegisterMethod("files.append", "追加文本到文件末尾", autogofiles.Append, true)
	engine.RegisterMethod("files.appendBytes", "追加字节到文件末尾", autogofiles.AppendBytes, true)
	engine.RegisterMethod("files.copy", "复制文件", autogofiles.Copy, true)
	engine.RegisterMethod("files.move", "移动文件", autogofiles.Move, true)
	engine.RegisterMethod("files.rename", "重命名文件", autogofiles.Rename, true)
	engine.RegisterMethod("files.getName", "返回文件名", autogofiles.GetName, true)
	engine.RegisterMethod("files.getNameWithoutExtension", "返回不含扩展名的文件名", autogofiles.GetNameWithoutExtension, true)
	engine.RegisterMethod("files.getExtension", "返回文件扩展名", autogofiles.GetExtension, true)
	engine.RegisterMethod("files.getMd5", "远程 AutoGo/files 不包含 GetMd5", func(path string) (string, error) {
		return "", fmt.Errorf("files.getMd5 is unavailable in remote AutoGo/files")
	}, true)
	engine.RegisterMethod("files.remove", "删除文件或文件夹", autogofiles.Remove, true)
	engine.RegisterMethod("files.path", "返回相对路径对应的绝对路径", autogofiles.Path, true)
	engine.RegisterMethod("files.listDir", "列出文件夹下的文件和文件夹", autogofiles.ListDir, true)
	return nil
}
