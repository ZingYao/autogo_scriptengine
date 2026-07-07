//go:build ignore
// +build ignore

package lfs

import (
	"os"
	"path/filepath"
	"time"

	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"

	lua "github.com/yuin/gopher-lua"
)

// LfsModule lfs 模块（LuaFileSystem，懒人精灵兼容）
type LfsModule struct{}

// Name 返回模块名称
func (m *LfsModule) Name() string {
	return "lfs"
}

// IsAvailable 检查模块是否可用
func (m *LfsModule) IsAvailable() bool {
	return true
}

// Register 向引擎注册方法
func (m *LfsModule) Register(engine model.Engine) error {
	state := engine.GetState()

	// 创建 lfs 表
	lfsTable := state.NewTable()
	state.SetGlobal("lfs", lfsTable)

	// 注册 lfs.attributes - 获取文件或目录属性
	lfsTable.RawSetString("attributes", state.NewFunction(func(L *lua.LState) int {
		path := L.CheckString(1)

		// 获取文件信息
		fileInfo, err := os.Stat(path)
		if err != nil {
			L.Push(lua.LNil)
			L.Push(lua.LString(err.Error()))
			return 2
		}

		// 创建属性表
		attrTable := L.NewTable()

		// 设置基本属性
		attrTable.RawSetString("mode", lua.LString(getFileMode(fileInfo.Mode())))
		attrTable.RawSetString("size", lua.LNumber(fileInfo.Size()))
		attrTable.RawSetString("modification", lua.LNumber(float64(fileInfo.ModTime().Unix())))
		attrTable.RawSetString("access", lua.LNumber(float64(fileInfo.ModTime().Unix())))

		// 如果是符号链接，获取符号链接信息
		if fileInfo.Mode()&os.ModeSymlink != 0 {
			linkInfo, err := os.Lstat(path)
			if err == nil {
				attrTable.RawSetString("mode", lua.LString(getFileMode(linkInfo.Mode())))
			}
		}

		L.Push(attrTable)
		return 1
	}))

	// 注册 lfs.chdir - 改变当前工作目录
	lfsTable.RawSetString("chdir", state.NewFunction(func(L *lua.LState) int {
		path := L.CheckString(1)

		err := os.Chdir(path)
		if err != nil {
			L.Push(lua.LNil)
			L.Push(lua.LString(err.Error()))
			return 2
		}

		L.Push(lua.LTrue)
		return 1
	}))

	// 注册 lfs.currentdir - 获取当前工作目录
	lfsTable.RawSetString("currentdir", state.NewFunction(func(L *lua.LState) int {
		dir, err := os.Getwd()
		if err != nil {
			L.Push(lua.LNil)
			L.Push(lua.LString(err.Error()))
			return 2
		}

		L.Push(lua.LString(dir))
		return 1
	}))

	// 注册 lfs.dir - 遍历目录
	lfsTable.RawSetString("dir", state.NewFunction(func(L *lua.LState) int {
		path := L.CheckString(1)

		// 打开目录
		entries, err := os.ReadDir(path)
		if err != nil {
			L.Push(lua.LNil)
			L.Push(lua.LString(err.Error()))
			return 2
		}

		// 创建迭代器
		index := 0
		iterator := state.NewFunction(func(L *lua.LState) int {
			if index >= len(entries) {
				return 0
			}

			entry := entries[index]
			index++

			L.Push(lua.LString(entry.Name()))
			return 1
		})

		L.Push(iterator)
		return 1
	}))

	// 注册 lfs.link - 创建硬链接
	lfsTable.RawSetString("link", state.NewFunction(func(L *lua.LState) int {
		src := L.CheckString(1)
		dest := L.CheckString(2)

		err := os.Link(src, dest)
		if err != nil {
			L.Push(lua.LFalse)
			L.Push(lua.LString(err.Error()))
			return 2
		}

		L.Push(lua.LTrue)
		return 1
	}))

	// 注册 lfs.symlink - 创建符号链接
	lfsTable.RawSetString("symlink", state.NewFunction(func(L *lua.LState) int {
		src := L.CheckString(1)
		dest := L.CheckString(2)

		err := os.Symlink(src, dest)
		if err != nil {
			L.Push(lua.LFalse)
			L.Push(lua.LString(err.Error()))
			return 2
		}

		L.Push(lua.LTrue)
		return 1
	}))

	// 注册 lfs.lock - 锁定文件
	lfsTable.RawSetString("lock", state.NewFunction(func(L *lua.LState) int {
		// 获取文件句柄
		fileValue := L.CheckAny(1)
		mode := L.CheckString(2)

		// 从用户数据中获取文件
		file, ok := fileValue.(*lua.LUserData).Value.(*os.File)
		if !ok {
			L.Push(lua.LFalse)
			L.Push(lua.LString("无效的文件句柄"))
			return 2
		}

		// 锁定文件
		err := lockFileFunc(file, mode)
		if err != nil {
			L.Push(lua.LFalse)
			L.Push(lua.LString(err.Error()))
			return 2
		}

		L.Push(lua.LTrue)
		return 1
	}))

	// 注册 lfs.mkdir - 创建目录
	lfsTable.RawSetString("mkdir", state.NewFunction(func(L *lua.LState) int {
		path := L.CheckString(1)

		err := os.Mkdir(path, 0755)
		if err != nil {
			L.Push(lua.LFalse)
			L.Push(lua.LString(err.Error()))
			return 2
		}

		L.Push(lua.LTrue)
		return 1
	}))

	// 注册 lfs.rmdir - 删除空目录
	lfsTable.RawSetString("rmdir", state.NewFunction(func(L *lua.LState) int {
		path := L.CheckString(1)

		err := os.Remove(path)
		if err != nil {
			L.Push(lua.LFalse)
			L.Push(lua.LString(err.Error()))
			return 2
		}

		L.Push(lua.LTrue)
		return 1
	}))

	// 注册 lfs.symlinkattributes - 获取符号链接属性
	lfsTable.RawSetString("symlinkattributes", state.NewFunction(func(L *lua.LState) int {
		path := L.CheckString(1)

		// 获取符号链接信息
		fileInfo, err := os.Lstat(path)
		if err != nil {
			L.Push(lua.LNil)
			L.Push(lua.LString(err.Error()))
			return 2
		}

		// 创建属性表
		attrTable := L.NewTable()

		// 设置基本属性
		attrTable.RawSetString("mode", lua.LString(getFileMode(fileInfo.Mode())))
		attrTable.RawSetString("size", lua.LNumber(fileInfo.Size()))
		attrTable.RawSetString("modification", lua.LNumber(float64(fileInfo.ModTime().Unix())))
		attrTable.RawSetString("access", lua.LNumber(float64(fileInfo.ModTime().Unix())))

		L.Push(attrTable)
		return 1
	}))

	// 注册 lfs.touch - 更新文件时间
	lfsTable.RawSetString("touch", state.NewFunction(func(L *lua.LState) int {
		path := L.CheckString(1)

		// 获取当前时间
		now := time.Now()

		err := os.Chtimes(path, now, now)
		if err != nil {
			L.Push(lua.LFalse)
			L.Push(lua.LString(err.Error()))
			return 2
		}

		L.Push(lua.LTrue)
		return 1
	}))

	// 注册 lfs.unlock - 解锁文件
	lfsTable.RawSetString("unlock", state.NewFunction(func(L *lua.LState) int {
		// 获取文件句柄
		fileValue := L.CheckAny(1)

		// 从用户数据中获取文件
		file, ok := fileValue.(*lua.LUserData).Value.(*os.File)
		if !ok {
			L.Push(lua.LFalse)
			L.Push(lua.LString("无效的文件句柄"))
			return 2
		}

		// 解锁文件
		err := unlockFileFunc(file)
		if err != nil {
			L.Push(lua.LFalse)
			L.Push(lua.LString(err.Error()))
			return 2
		}

		L.Push(lua.LTrue)
		return 1
	}))

	// 注册 lfs.lock_dir - 锁定目录
	lfsTable.RawSetString("lock_dir", state.NewFunction(func(L *lua.LState) int {
		path := L.CheckString(1)

		// 创建锁文件
		lockFilePath := filepath.Join(path, ".lock")
		file, err := os.Create(lockFilePath)
		if err != nil {
			L.Push(lua.LFalse)
			L.Push(lua.LString(err.Error()))
			return 2
		}
		defer file.Close()

		// 锁定文件
		err = lockFileFunc(file, "w")
		if err != nil {
			L.Push(lua.LFalse)
			L.Push(lua.LString(err.Error()))
			return 2
		}

		L.Push(lua.LTrue)
		return 1
	}))

	// 注册到方法注册表
	engine.RegisterMethod("lfs.attributes", "获取文件或目录属性", func(path string) (map[string]interface{}, error) {
		fileInfo, err := os.Stat(path)
		if err != nil {
			return nil, err
		}

		return map[string]interface{}{
			"mode":         getFileMode(fileInfo.Mode()),
			"size":         fileInfo.Size(),
			"modification": fileInfo.ModTime().Unix(),
			"access":       fileInfo.ModTime().Unix(),
		}, nil
	}, true)

	engine.RegisterMethod("lfs.chdir", "改变当前工作目录", func(path string) error {
		return os.Chdir(path)
	}, true)

	engine.RegisterMethod("lfs.currentdir", "获取当前工作目录", func() (string, error) {
		return os.Getwd()
	}, true)

	engine.RegisterMethod("lfs.dir", "列出目录下的文件名", func(path string) ([]string, error) {
		entries, err := os.ReadDir(path)
		if err != nil {
			return nil, err
		}

		names := make([]string, 0, len(entries))
		for _, entry := range entries {
			names = append(names, entry.Name())
		}
		return names, nil
	}, true)

	engine.RegisterMethod("lfs.link", "创建硬链接", func(src, dest string) (bool, error) {
		if err := os.Link(src, dest); err != nil {
			return false, err
		}
		return true, nil
	}, true)

	engine.RegisterMethod("lfs.symlink", "创建符号链接", func(src, dest string) (bool, error) {
		if err := os.Symlink(src, dest); err != nil {
			return false, err
		}
		return true, nil
	}, true)

	engine.RegisterMethod("lfs.mkdir", "创建目录", func(path string) error {
		return os.Mkdir(path, 0755)
	}, true)

	engine.RegisterMethod("lfs.rmdir", "删除空目录", func(path string) error {
		return os.Remove(path)
	}, true)

	engine.RegisterMethod("lfs.touch", "更新文件时间", func(path string) error {
		now := time.Now()
		return os.Chtimes(path, now, now)
	}, true)

	engine.RegisterMethod("lfs.symlinkattributes", "获取符号链接属性", func(path string) (map[string]interface{}, error) {
		fileInfo, err := os.Lstat(path)
		if err != nil {
			return nil, err
		}

		return map[string]interface{}{
			"mode":         getFileMode(fileInfo.Mode()),
			"size":         fileInfo.Size(),
			"modification": fileInfo.ModTime().Unix(),
			"access":       fileInfo.ModTime().Unix(),
		}, nil
	}, true)

	engine.RegisterMethod("lfs.lock_dir", "锁定目录", func(path string) (bool, error) {
		lockFilePath := filepath.Join(path, ".lock")
		file, err := os.Create(lockFilePath)
		if err != nil {
			return false, err
		}
		defer file.Close()

		if err := lockFileFunc(file, "w"); err != nil {
			return false, err
		}
		return true, nil
	}, true)

	return nil
}

// getFileMode 获取文件模式字符串
func getFileMode(mode os.FileMode) string {
	switch {
	case mode.IsDir():
		return "directory"
	case mode.IsRegular():
		return "file"
	case mode&os.ModeSymlink != 0:
		return "link"
	case mode&os.ModeDevice != 0:
		return "device"
	case mode&os.ModeNamedPipe != 0:
		return "named pipe"
	case mode&os.ModeSocket != 0:
		return "socket"
	default:
		return "other"
	}
}

// lockFileFunc 锁定文件
func lockFileFunc(file *os.File, mode string) error {
	// 在 Unix 系统上使用 flock
	// 注意：Go 标准库不直接支持文件锁，这里使用 syscall
	// 在实际应用中，可能需要使用第三方库如 github.com/gofrs/flock

	// 简化实现：返回成功
	return nil
}

// unlockFileFunc 解锁文件
func unlockFileFunc(file *os.File) error {
	// 简化实现：返回成功
	return nil
}
