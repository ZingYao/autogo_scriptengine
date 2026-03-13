# Files 模块

Files 模块提供了文件和文件夹操作功能。

## 方法列表

### files.isDir(path)

返回路径 path 是否是文件夹。

**参数**:
- `path` (string): 文件路径

**返回值**: `boolean` - 如果是文件夹返回 `true`，否则返回 `false`

**调用示例**:
```javascript
if (files.isDir("/sdcard/Download")) {
    console.log("是文件夹");
}
```

### files.isFile(path)

返回路径 path 是否是文件。

**参数**:
- `path` (string): 文件路径

**返回值**: `boolean` - 如果是文件返回 `true`，否则返回 `false`

**调用示例**:
```javascript
if (files.isFile("/sdcard/test.txt")) {
    console.log("是文件");
}
```

### files.isEmptyDir(path)

返回文件夹 path 是否为空文件夹。如果该路径并非文件夹，则直接返回 `false`。

**参数**:
- `path` (string): 文件夹路径

**返回值**: `boolean` - 如果是空文件夹返回 `true`，否则返回 `false`

**调用示例**:
```javascript
if (files.isEmptyDir("/sdcard/empty")) {
    console.log("是空文件夹");
}
```

### files.create(path)

创建一个文件或文件夹并返回是否创建成功。如果文件已经存在，则直接返回 `false`。

**参数**:
- `path` (string): 文件或文件夹路径

**返回值**: `boolean` - 如果创建成功返回 `true`，否则返回 `false`

**调用示例**:
```javascript
if (files.create("/sdcard/test.txt")) {
    console.log("文件创建成功");
}
```

### files.exists(path)

返回在路径 path 处的文件是否存在。

**参数**:
- `path` (string): 文件路径

**返回值**: `boolean` - 如果文件存在返回 `true`，否则返回 `false`

**调用示例**:
```javascript
if (files.exists("/sdcard/test.txt")) {
    console.log("文件存在");
}
```

### files.ensureDir(path)

确保路径 path 所在的文件夹存在。如果该路径所在文件夹不存在，则创建该文件夹。

**参数**:
- `path` (string): 文件路径

**返回值**: `boolean` - 如果成功返回 `true`，否则返回 `false`

**调用示例**:
```javascript
files.ensureDir("/sdcard/test/subdir");
console.log("文件夹已确保存在");
```

### files.read(path)

读取文本文件 path 的所有内容并返回。

**参数**:
- `path` (string): 文件路径

**返回值**: `string` - 文件内容

**调用示例**:
```javascript
const content = files.read("/sdcard/test.txt");
console.log("文件内容: " + content);
```

### files.readBytes(path)

读取文件 path 的所有内容并返回字节数组。

**参数**:
- `path` (string): 文件路径

**返回值**: `Array` - 字节数组，如果文件不存在返回 `null`

**调用示例**:
```javascript
const bytes = files.readBytes("/sdcard/image.png");
console.log("文件大小: " + bytes.length + " bytes");
```

### files.write(path, text)

把 text 写入到文件 path 中。如果文件存在则覆盖，不存在则创建。

**参数**:
- `path` (string): 文件路径
- `text` (string): 要写入的文本内容

**返回值**: `undefined`

**调用示例**:
```javascript
files.write("/sdcard/test.txt", "Hello, World!");
console.log("文件已写入");
```

### files.writeBytes(path, bytes)

把 bytes 写入到文件 path 中。如果文件存在则覆盖，不存在则创建。

**参数**:
- `path` (string): 文件路径
- `bytes` (Array): 要写入的字节数组

**返回值**: `undefined`

**调用示例**:
```javascript
const bytes = [72, 101, 108, 108, 111]; // "Hello"
files.writeBytes("/sdcard/test.bin", bytes);
console.log("字节数据已写入");
```

### files.append(path, text)

把 text 追加到文件 path 的末尾。如果文件不存在则创建。

**参数**:
- `path` (string): 文件路径
- `text` (string): 要追加的文本内容

**返回值**: `undefined`

**调用示例**:
```javascript
files.append("/sdcard/test.txt", "\n追加的内容");
console.log("内容已追加");
```

### files.appendBytes(path, bytes)

把 bytes 追加到文件 path 的末尾。如果文件不存在则创建。

**参数**:
- `path` (string): 文件路径
- `bytes` (Array): 要追加的字节数组

**返回值**: `undefined`

**调用示例**:
```javascript
const bytes = [33, 34, 35]; // "!\"#"
files.appendBytes("/sdcard/test.bin", bytes);
console.log("字节数据已追加");
```

### files.copy(fromPath, toPath)

复制文件，返回是否复制成功。

**参数**:
- `fromPath` (string): 源文件路径
- `toPath` (string): 目标文件路径

**返回值**: `boolean` - 如果复制成功返回 `true`，否则返回 `false`

**调用示例**:
```javascript
if (files.copy("/sdcard/source.txt", "/sdcard/dest.txt")) {
    console.log("文件复制成功");
}
```

### files.move(fromPath, toPath)

移动文件，返回是否移动成功。

**参数**:
- `fromPath` (string): 源文件路径
- `toPath` (string): 目标文件路径

**返回值**: `boolean` - 如果移动成功返回 `true`，否则返回 `false`

**调用示例**:
```javascript
if (files.move("/sdcard/source.txt", "/sdcard/newlocation.txt")) {
    console.log("文件移动成功");
}
```

### files.rename(path, newName)

重命名文件，并返回是否重命名成功。

**参数**:
- `path` (string): 文件路径
- `newName` (string): 新文件名

**返回值**: `boolean` - 如果重命名成功返回 `true`，否则返回 `false`

**调用示例**:
```javascript
if (files.rename("/sdcard/old.txt", "new.txt")) {
    console.log("文件重命名成功");
}
```

### files.remove(path)

删除文件或文件夹，如果是文件夹会删除整个文件夹包含里面的所有文件，返回是否删除成功。

**参数**:
- `path` (string): 文件或文件夹路径

**返回值**: `boolean` - 如果删除成功返回 `true`，否则返回 `false`

**调用示例**:
```javascript
if (files.remove("/sdcard/test.txt")) {
    console.log("文件删除成功");
}
```

### files.getName(path)

返回文件的文件名。例如 `files.getName("/sdcard/1.txt")` 返回 `"1.txt"`。

**参数**:
- `path` (string): 文件路径

**返回值**: `string` - 文件名

**调用示例**:
```javascript
const name = files.getName("/sdcard/test.txt");
console.log("文件名: " + name); // 输出: test.txt
```

### files.getNameWithoutExtension(path)

返回不含拓展名的文件的文件名。例如 `files.getName("/sdcard/1.txt")` 返回 `"1"`。

**参数**:
- `path` (string): 文件路径

**返回值**: `string` - 不含扩展名的文件名

**调用示例**:
```javascript
const name = files.getNameWithoutExtension("/sdcard/test.txt");
console.log("文件名(不含扩展名): " + name); // 输出: test
```

### files.getExtension(path)

返回文件的拓展名。例如 `files.getExtension("/sdcard/1.txt")` 返回 `"txt"`。

**参数**:
- `path` (string): 文件路径

**返回值**: `string` - 文件扩展名

**调用示例**:
```javascript
const ext = files.getExtension("/sdcard/test.txt");
console.log("文件扩展名: " + ext); // 输出: txt
```

### files.path(relativePath)

返回相对路径对应的绝对路径。例如 `files.path("./1.png")`，如果运行这个语句的脚本位于文件夹 `"/sdcard/脚本/"` 中，则返回 `"/sdcard/脚本/1.png"`。

**参数**:
- `relativePath` (string): 相对路径

**返回值**: `string` - 绝对路径

**调用示例**:
```javascript
const absPath = files.path("./test.txt");
console.log("绝对路径: " + absPath);
```

### files.listDir(path)

列出文件夹 path 下的所有文件和文件夹。

**参数**:
- `path` (string): 文件夹路径

**返回值**: `Array` - 文件和文件夹名称数组

**调用示例**:
```javascript
const fileList = files.listDir("/sdcard");
for (let i = 0; i < fileList.length; i++) {
    console.log(fileList[i]);
}
```

## 完整示例

```javascript
// 文件和文件夹检查
console.log("=== 文件检查 ===");
if (files.exists("/sdcard/test.txt")) {
    console.log("文件存在");
    if (files.isFile("/sdcard/test.txt")) {
        console.log("是文件");
    }
}

if (files.isDir("/sdcard/Download")) {
    console.log("是文件夹");
    const fileList = files.listDir("/sdcard/Download");
    console.log("文件夹内容: " + fileList.join(", "));
}

// 文件读写
console.log("\n=== 文件读写 ===");
files.write("/sdcard/test.txt", "Hello, World!");
const content = files.read("/sdcard/test.txt");
console.log("文件内容: " + content);

// 追加内容
files.append("/sdcard/test.txt", "\n追加的内容");
const newContent = files.read("/sdcard/test.txt");
console.log("新内容: " + newContent);

// 文件操作
console.log("\n=== 文件操作 ===");
files.copy("/sdcard/test.txt", "/sdcard/copy.txt");
files.move("/sdcard/copy.txt", "/sdcard/moved.txt");
files.rename("/sdcard/moved.txt", "renamed.txt");

// 获取文件信息
console.log("\n=== 文件信息 ===");
const fullPath = "/sdcard/test.txt";
console.log("文件名: " + files.getName(fullPath));
console.log("文件名(不含扩展名): " + files.getNameWithoutExtension(fullPath));
console.log("文件扩展名: " + files.getExtension(fullPath));
console.log("绝对路径: " + files.path("./test.txt"));

// 删除文件
console.log("\n=== 删除文件 ===");
files.remove("/sdcard/renamed.txt");
console.log("文件已删除");
```
