package files

import (
	"strconv"

	"github.com/ZingYao/autogo_scriptengine/js_engine/model"

	"github.com/Dasongzi1366/AutoGo/files"
	"github.com/dop251/goja"
)

// FilesModule files 模块
type FilesModule struct{}

// Name 返回模块名称
func (m *FilesModule) Name() string {
	return "files"
}

// IsAvailable 检查模块是否可用
func (m *FilesModule) IsAvailable() bool {
	return true
}

// Register 向引擎注册方法
func (m *FilesModule) Register(engine model.Engine) error {
	vm := engine.GetVM()

	filesObj := vm.NewObject()
	vm.Set("files", filesObj)

	filesObj.Set("isDir", func(call goja.FunctionCall) goja.Value {
		path := call.Argument(0).String()
		result := files.IsDir(path)
		return vm.ToValue(result)
	})

	filesObj.Set("isFile", func(call goja.FunctionCall) goja.Value {
		path := call.Argument(0).String()
		result := files.IsFile(path)
		return vm.ToValue(result)
	})

	filesObj.Set("isEmptyDir", func(call goja.FunctionCall) goja.Value {
		path := call.Argument(0).String()
		result := files.IsEmptyDir(path)
		return vm.ToValue(result)
	})

	filesObj.Set("create", func(call goja.FunctionCall) goja.Value {
		path := call.Argument(0).String()
		result := files.Create(path)
		return vm.ToValue(result)
	})

	filesObj.Set("exists", func(call goja.FunctionCall) goja.Value {
		path := call.Argument(0).String()
		result := files.Exists(path)
		return vm.ToValue(result)
	})

	filesObj.Set("ensureDir", func(call goja.FunctionCall) goja.Value {
		path := call.Argument(0).String()
		result := files.EnsureDir(path)
		return vm.ToValue(result)
	})

	filesObj.Set("read", func(call goja.FunctionCall) goja.Value {
		path := call.Argument(0).String()
		result := files.Read(path)
		return vm.ToValue(result)
	})

	filesObj.Set("readBytes", func(call goja.FunctionCall) goja.Value {
		path := call.Argument(0).String()
		result := files.ReadBytes(path)
		if result != nil {
			return vm.ToValue(result)
		}
		return goja.Null()
	})

	filesObj.Set("write", func(call goja.FunctionCall) goja.Value {
		path := call.Argument(0).String()
		text := call.Argument(1).String()
		files.Write(path, text)
		return goja.Undefined()
	})

	filesObj.Set("writeBytes", func(call goja.FunctionCall) goja.Value {
		path := call.Argument(0).String()
		jsBytes := call.Argument(1).Export()
		var bytes []byte

		switch v := jsBytes.(type) {
		case []byte:
			bytes = v
		case []interface{}:
			bytes = make([]byte, len(v))
			for i, val := range v {
				if num, ok := val.(float64); ok {
					bytes[i] = byte(num)
				}
			}
		}

		files.WriteBytes(path, bytes)
		return goja.Undefined()
	})

	filesObj.Set("append", func(call goja.FunctionCall) goja.Value {
		path := call.Argument(0).String()
		text := call.Argument(1).String()
		files.Append(path, text)
		return goja.Undefined()
	})

	filesObj.Set("appendBytes", func(call goja.FunctionCall) goja.Value {
		path := call.Argument(0).String()
		jsBytes := call.Argument(1).Export()
		var bytes []byte

		switch v := jsBytes.(type) {
		case []byte:
			bytes = v
		case []interface{}:
			bytes = make([]byte, len(v))
			for i, val := range v {
				if num, ok := val.(float64); ok {
					bytes[i] = byte(num)
				}
			}
		}

		files.AppendBytes(path, bytes)
		return goja.Undefined()
	})

	filesObj.Set("copy", func(call goja.FunctionCall) goja.Value {
		fromPath := call.Argument(0).String()
		toPath := call.Argument(1).String()
		result := files.Copy(fromPath, toPath)
		return vm.ToValue(result)
	})

	filesObj.Set("move", func(call goja.FunctionCall) goja.Value {
		fromPath := call.Argument(0).String()
		toPath := call.Argument(1).String()
		result := files.Move(fromPath, toPath)
		return vm.ToValue(result)
	})

	filesObj.Set("rename", func(call goja.FunctionCall) goja.Value {
		path := call.Argument(0).String()
		newName := call.Argument(1).String()
		result := files.Rename(path, newName)
		return vm.ToValue(result)
	})

	filesObj.Set("remove", func(call goja.FunctionCall) goja.Value {
		path := call.Argument(0).String()
		result := files.Remove(path)
		return vm.ToValue(result)
	})

	filesObj.Set("getName", func(call goja.FunctionCall) goja.Value {
		path := call.Argument(0).String()
		result := files.GetName(path)
		return vm.ToValue(result)
	})

	filesObj.Set("getNameWithoutExtension", func(call goja.FunctionCall) goja.Value {
		path := call.Argument(0).String()
		result := files.GetNameWithoutExtension(path)
		return vm.ToValue(result)
	})

	filesObj.Set("getExtension", func(call goja.FunctionCall) goja.Value {
		path := call.Argument(0).String()
		result := files.GetExtension(path)
		return vm.ToValue(result)
	})

	filesObj.Set("getMd5", func(call goja.FunctionCall) goja.Value {
		path := call.Argument(0).String()
		result := files.GetMd5(path)
		return vm.ToValue(result)
	})

	filesObj.Set("path", func(call goja.FunctionCall) goja.Value {
		relativePath := call.Argument(0).String()
		result := files.Path(relativePath)
		return vm.ToValue(result)
	})

	filesObj.Set("listDir", func(call goja.FunctionCall) goja.Value {
		path := call.Argument(0).String()
		result := files.ListDir(path)
		arr := vm.NewArray()
		for i, item := range result {
			arr.Set(strconv.Itoa(i), item)
		}
		return arr
	})

	engine.RegisterMethod("files.isDir", "返回路径path是否是文件夹", files.IsDir, true)
	engine.RegisterMethod("files.isFile", "返回路径path是否是文件", files.IsFile, true)
	engine.RegisterMethod("files.isEmptyDir", "返回文件夹path是否为空文件夹", files.IsEmptyDir, true)
	engine.RegisterMethod("files.create", "创建一个文件或文件夹", files.Create, true)
	engine.RegisterMethod("files.exists", "返回在路径path处的文件是否存在", files.Exists, true)
	engine.RegisterMethod("files.ensureDir", "确保路径path所在的文件夹存在", files.EnsureDir, true)
	engine.RegisterMethod("files.read", "读取文本文件path的所有内容并返回", files.Read, true)
	engine.RegisterMethod("files.readBytes", "读取文件path的所有内容并返回", files.ReadBytes, true)
	engine.RegisterMethod("files.write", "把text写入到文件path中", func(path, text string) { files.Write(path, text) }, true)
	engine.RegisterMethod("files.writeBytes", "把bytes写入到文件path中", func(path string, bytes []byte) { files.WriteBytes(path, bytes) }, true)
	engine.RegisterMethod("files.append", "把text追加到文件path的末尾", func(path, text string) { files.Append(path, text) }, true)
	engine.RegisterMethod("files.appendBytes", "把bytes追加到文件path的末尾", func(path string, bytes []byte) { files.AppendBytes(path, bytes) }, true)
	engine.RegisterMethod("files.copy", "复制文件", files.Copy, true)
	engine.RegisterMethod("files.move", "移动文件", files.Move, true)
	engine.RegisterMethod("files.rename", "重命名文件", files.Rename, true)
	engine.RegisterMethod("files.remove", "删除文件或文件夹", files.Remove, true)
	engine.RegisterMethod("files.getName", "返回文件的文件名", files.GetName, true)
	engine.RegisterMethod("files.getNameWithoutExtension", "返回不含拓展名的文件的文件名", files.GetNameWithoutExtension, true)
	engine.RegisterMethod("files.getExtension", "返回文件的拓展名", files.GetExtension, true)
	engine.RegisterMethod("files.getMd5", "返回文件的MD5值", files.GetMd5, true)
	engine.RegisterMethod("files.path", "返回相对路径对应的绝对路径", files.Path, true)
	engine.RegisterMethod("files.listDir", "列出文件夹path下的所有文件和文件夹", files.ListDir, true)

	return nil
}
