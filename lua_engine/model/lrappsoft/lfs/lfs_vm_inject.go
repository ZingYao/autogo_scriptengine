package lfs

import (
	"os"
	"path/filepath"
	"time"

	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"
)

// LfsModule 是 go-lua-vm 迁移后的模块壳。
type LfsModule struct{}

func New() *LfsModule { return &LfsModule{} }

func (m *LfsModule) Name() string { return "lfs" }

func (m *LfsModule) IsAvailable() bool { return true }

func (m *LfsModule) Register(engine model.Engine) error {
	engine.RegisterMethod("lfs.currentdir", "返回当前工作目录", os.Getwd, true)
	engine.RegisterMethod("lfs.chdir", "切换当前工作目录", os.Chdir, true)
	engine.RegisterMethod("lfs.mkdir", "创建目录", func(path string) error {
		return os.Mkdir(path, 0o755)
	}, true)
	engine.RegisterMethod("lfs.rmdir", "删除目录", os.Remove, true)
	engine.RegisterMethod("lfs.touch", "更新文件时间", func(path string) error {
		now := time.Now()
		file, err := os.OpenFile(path, os.O_CREATE, 0o644)
		if err != nil {
			return err
		}
		_ = file.Close()
		return os.Chtimes(path, now, now)
	}, true)
	engine.RegisterMethod("lfs.attributes", "读取文件属性", fileAttributes, true)
	engine.RegisterMethod("lfs.symlinkattributes", "读取符号链接属性", linkAttributes, true)
	engine.RegisterMethod("lfs.dir", "读取目录条目", func(path string) ([]string, error) {
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
	engine.RegisterMethod("lfs.link", "创建硬链接", func(oldName, newName string) (bool, error) {
		return true, os.Link(oldName, newName)
	}, true)
	engine.RegisterMethod("lfs.symlink", "创建符号链接", func(oldName, newName string) (bool, error) {
		return true, os.Symlink(oldName, newName)
	}, true)
	engine.RegisterMethod("lfs.lock_dir", "创建目录锁文件", func(path string) (bool, error) {
		lockPath := filepath.Join(path, ".lock")
		file, err := os.OpenFile(lockPath, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0o644)
		if err != nil {
			return false, err
		}
		return true, file.Close()
	}, true)
	return nil
}

func GetModule() model.Module { return &LfsModule{} }

func fileAttributes(path string) (map[string]interface{}, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	return fileInfoAttributes(info), nil
}

func linkAttributes(path string) (map[string]interface{}, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return nil, err
	}
	return fileInfoAttributes(info), nil
}

func fileInfoAttributes(info os.FileInfo) map[string]interface{} {
	mode := "file"
	if info.Mode()&os.ModeSymlink != 0 {
		mode = "link"
	} else if info.IsDir() {
		mode = "directory"
	}
	return map[string]interface{}{
		"mode":         mode,
		"size":         info.Size(),
		"modification": info.ModTime().Unix(),
		"access":       info.ModTime().Unix(),
	}
}
