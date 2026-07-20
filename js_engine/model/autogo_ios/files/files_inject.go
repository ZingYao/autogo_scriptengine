package files

import (
	"fmt"

	"github.com/ZingYao/autogo_scriptengine/js_engine/model"

	autogofiles "github.com/Dasongzi1366/AutoGo/files"
	"github.com/ZingYao/goja"
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

// bytesFromValue 将脚本侧字符串、字节切片或数字数组转换为 Go 字节切片。
func bytesFromValue(value goja.Value) []byte {
	switch data := value.Export().(type) {
	case []byte:
		return data
	case string:
		return []byte(data)
	case []interface{}:
		result := make([]byte, 0, len(data))
		for _, item := range data {
			switch number := item.(type) {
			case int:
				result = append(result, byte(number))
			case int64:
				result = append(result, byte(number))
			case float64:
				result = append(result, byte(number))
			}
		}
		return result
	default:
		return nil
	}
}

// Register 向引擎注册 iOS files 方法。
func (m *FilesModule) Register(engine model.Engine) error {
	vm := engine.GetVM()
	filesObj := vm.NewObject()
	vm.Set("files", filesObj)

	filesObj.Set("isFile", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(autogofiles.IsFile(call.Argument(0).String()))
	})
	filesObj.Set("isDir", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(autogofiles.IsDir(call.Argument(0).String()))
	})
	filesObj.Set("isEmptyDir", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(autogofiles.IsEmptyDir(call.Argument(0).String()))
	})
	filesObj.Set("create", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(autogofiles.Create(call.Argument(0).String()))
	})
	filesObj.Set("exists", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(autogofiles.Exists(call.Argument(0).String()))
	})
	filesObj.Set("ensureDir", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(autogofiles.EnsureDir(call.Argument(0).String()))
	})
	filesObj.Set("read", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(autogofiles.Read(call.Argument(0).String()))
	})
	filesObj.Set("readBytes", func(call goja.FunctionCall) goja.Value {
		data := autogofiles.ReadBytes(call.Argument(0).String())
		if data == nil {
			return goja.Null()
		}
		return vm.ToValue(data)
	})
	filesObj.Set("write", func(call goja.FunctionCall) goja.Value {
		autogofiles.Write(call.Argument(0).String(), call.Argument(1).String())
		return goja.Undefined()
	})
	filesObj.Set("writeBytes", func(call goja.FunctionCall) goja.Value {
		autogofiles.WriteBytes(call.Argument(0).String(), bytesFromValue(call.Argument(1)))
		return goja.Undefined()
	})
	filesObj.Set("append", func(call goja.FunctionCall) goja.Value {
		autogofiles.Append(call.Argument(0).String(), call.Argument(1).String())
		return goja.Undefined()
	})
	filesObj.Set("appendBytes", func(call goja.FunctionCall) goja.Value {
		autogofiles.AppendBytes(call.Argument(0).String(), bytesFromValue(call.Argument(1)))
		return goja.Undefined()
	})
	filesObj.Set("copy", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(autogofiles.Copy(call.Argument(0).String(), call.Argument(1).String()))
	})
	filesObj.Set("move", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(autogofiles.Move(call.Argument(0).String(), call.Argument(1).String()))
	})
	filesObj.Set("rename", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(autogofiles.Rename(call.Argument(0).String(), call.Argument(1).String()))
	})
	filesObj.Set("getName", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(autogofiles.GetName(call.Argument(0).String()))
	})
	filesObj.Set("getNameWithoutExtension", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(autogofiles.GetNameWithoutExtension(call.Argument(0).String()))
	})
	filesObj.Set("getExtension", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(autogofiles.GetExtension(call.Argument(0).String()))
	})
	filesObj.Set("getMd5", func(call goja.FunctionCall) goja.Value {
		panic(vm.NewGoError(fmt.Errorf("files.getMd5 is unavailable in remote AutoGo/files")))
	})
	filesObj.Set("remove", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(autogofiles.Remove(call.Argument(0).String()))
	})
	filesObj.Set("path", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(autogofiles.Path(call.Argument(0).String()))
	})
	filesObj.Set("listDir", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(autogofiles.ListDir(call.Argument(0).String()))
	})

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
