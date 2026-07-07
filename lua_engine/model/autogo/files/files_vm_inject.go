package files

import (
	"fmt"

	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"

	autogofiles "github.com/Dasongzi1366/AutoGo/files"
)

// FilesModule 是 go-lua-vm 迁移后的模块壳。
type FilesModule struct{}

func New() *FilesModule { return &FilesModule{} }

func (m *FilesModule) Name() string { return "files" }

func (m *FilesModule) IsAvailable() bool { return true }

func (m *FilesModule) Register(engine model.Engine) error {
	engine.RegisterMethod("files.isDir", "返回路径是否是文件夹", autogofiles.IsDir, true)
	engine.RegisterMethod("files.isFile", "返回路径是否是文件", autogofiles.IsFile, true)
	engine.RegisterMethod("files.isEmptyDir", "返回文件夹是否为空", autogofiles.IsEmptyDir, true)
	engine.RegisterMethod("files.create", "创建文件或文件夹", autogofiles.Create, true)
	engine.RegisterMethod("files.exists", "返回路径是否存在", autogofiles.Exists, true)
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
	engine.RegisterMethod("files.remove", "删除文件或文件夹", autogofiles.Remove, true)
	engine.RegisterMethod("files.getName", "返回文件名", autogofiles.GetName, true)
	engine.RegisterMethod("files.getNameWithoutExtension", "返回不含扩展名的文件名", autogofiles.GetNameWithoutExtension, true)
	engine.RegisterMethod("files.getExtension", "返回文件扩展名", autogofiles.GetExtension, true)
	engine.RegisterMethod("files.getMd5", "远程 AutoGo/files 不包含 GetMd5", func(path string) (string, error) {
		return "", fmt.Errorf("files.getMd5 is unavailable in remote AutoGo/files")
	}, true)
	engine.RegisterMethod("files.path", "返回相对路径对应的绝对路径", autogofiles.Path, true)
	engine.RegisterMethod("files.listDir", "列出文件夹下的所有文件和文件夹", autogofiles.ListDir, true)
	return nil
}

func GetModule() model.Module { return &FilesModule{} }
