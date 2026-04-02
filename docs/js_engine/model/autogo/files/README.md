# files 模块

## 模块简介

files 模块提供了文件系统操作功能，包括文件读写、复制、移动、删除等。

## 方法列表

### files.isDir
返回路径path是否是文件夹

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| path | string | 文件或文件夹路径 |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| isDir | boolean | 是否为文件夹 |

**使用示例：**
```javascript
// 判断路径是否为文件夹
var isDir = files.isDir("/sdcard/Download");
console.log("是否为文件夹: " + isDir);
```

---

### files.isFile
返回路径path是否是文件

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| path | string | 文件或文件夹路径 |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| isFile | boolean | 是否为文件 |

**使用示例：**
```javascript
// 判断路径是否为文件
var isFile = files.isFile("/sdcard/test.txt");
console.log("是否为文件: " + isFile);
```

---

### files.isEmptyDir
返回文件夹path是否为空文件夹

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| path | string | 文件夹路径 |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| isEmpty | boolean | 是否为空文件夹 |

**使用示例：**
```javascript
// 判断文件夹是否为空
var isEmpty = files.isEmptyDir("/sdcard/Download");
console.log("文件夹是否为空: " + isEmpty);
```

---

### files.create
创建一个文件或文件夹

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| path | string | 文件或文件夹路径 |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| success | boolean | 是否创建成功 |

**使用示例：**
```javascript
// 创建文件
var success = files.create("/sdcard/test.txt");
console.log("创建结果: " + success);
```

---

### files.exists
返回在路径path处的文件是否存在

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| path | string | 文件或文件夹路径 |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| exists | boolean | 文件是否存在 |

**使用示例：**
```javascript
// 检查文件是否存在
var exists = files.exists("/sdcard/test.txt");
console.log("文件是否存在: " + exists);
```

---

### files.ensureDir
确保路径path所在的文件夹存在

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| path | string | 文件或文件夹路径 |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| success | boolean | 是否创建成功 |

**使用示例：**
```javascript
// 确保文件夹存在
var success = files.ensureDir("/sdcard/test/subfolder");
console.log("创建结果: " + success);
```

---

### files.read
读取文本文件path的所有内容并返回

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| path | string | 文件路径 |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| content | string | 文件内容 |

**使用示例：**
```javascript
// 读取文件内容
var content = files.read("/sdcard/test.txt");
console.log("文件内容: " + content);
```

---

### files.readBytes
读取文件path的所有内容并返回

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| path | string | 文件路径 |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| data | array | 文件内容（字节数组） |

**使用示例：**
```javascript
// 读取文件内容（字节数组）
var data = files.readBytes("/sdcard/test.bin");
console.log("数据长度: " + data.length);
```

---

### files.write
把text写入到文件path中

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| path | string | 文件路径 |
| text | string | 要写入的文本内容 |

**返回值：**

无返回值

**使用示例：**
```javascript
// 写入文本到文件
files.write("/sdcard/test.txt", "Hello World");
```

---

### files.writeBytes
把bytes写入到文件path中

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| path | string | 文件路径 |
| bytes | array | 要写入的字节数组 |

**返回值：**

无返回值

**使用示例：**
```javascript
// 写入字节数组到文件
var data = [72, 101, 108, 108, 111]; // "Hello"
files.writeBytes("/sdcard/test.bin", data);
```

---

### files.append
把text追加到文件path的末尾

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| path | string | 文件路径 |
| text | string | 要追加的文本内容 |

**返回值：**

无返回值

**使用示例：**
```javascript
// 追加文本到文件
files.append("/sdcard/test.txt", "\nHello Again");
```

---

### files.appendBytes
把bytes追加到文件path的末尾

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| path | string | 文件路径 |
| bytes | array | 要追加的字节数组 |

**返回值：**

无返回值

**使用示例：**
```javascript
// 追加字节数组到文件
var data = [87, 111, 114, 108, 100]; // "World"
files.appendBytes("/sdcard/test.bin", data);
```

---

### files.copy
复制文件

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| fromPath | string | 源文件路径 |
| toPath | string | 目标文件路径 |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| success | boolean | 是否复制成功 |

**使用示例：**
```javascript
// 复制文件
var success = files.copy("/sdcard/test.txt", "/sdcard/test_copy.txt");
console.log("复制结果: " + success);
```

---

### files.move
移动文件

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| fromPath | string | 源文件路径 |
| toPath | string | 目标文件路径 |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| success | boolean | 是否移动成功 |

**使用示例：**
```javascript
// 移动文件
var success = files.move("/sdcard/test.txt", "/sdcard/Download/test.txt");
console.log("移动结果: " + success);
```

---

### files.rename
重命名文件

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| path | string | 文件路径 |
| newName | string | 新的文件名 |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| success | boolean | 是否重命名成功 |

**使用示例：**
```javascript
// 重命名文件
var success = files.rename("/sdcard/test.txt", "new_test.txt");
console.log("重命名结果: " + success);
```

---

### files.remove
删除文件或文件夹

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| path | string | 文件或文件夹路径 |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| success | boolean | 是否删除成功 |

**使用示例：**
```javascript
// 删除文件
var success = files.remove("/sdcard/test.txt");
console.log("删除结果: " + success);
```

---

### files.getName
返回文件的文件名

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| path | string | 文件或文件夹路径 |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| name | string | 文件名 |

**使用示例：**
```javascript
// 获取文件名
var name = files.getName("/sdcard/Download/test.txt");
console.log("文件名: " + name);
```

---

### files.getNameWithoutExtension
返回不含拓展名的文件的文件名

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| path | string | 文件或文件夹路径 |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| name | string | 不含扩展名的文件名 |

**使用示例：**
```javascript
// 获取不含扩展名的文件名
var name = files.getNameWithoutExtension("/sdcard/Download/test.txt");
console.log("文件名（不含扩展名）: " + name);
```

---

### files.getExtension
返回文件的拓展名

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| path | string | 文件或文件夹路径 |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| extension | string | 文件扩展名 |

**使用示例：**
```javascript
// 获取文件扩展名
var ext = files.getExtension("/sdcard/Download/test.txt");
console.log("文件扩展名: " + ext);
```

---

### files.getMd5
返回文件的MD5值

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| path | string | 文件路径 |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| md5 | string | 文件的 MD5 值 |

**使用示例：**
```javascript
// 获取文件的 MD5 值
var md5 = files.getMd5("/sdcard/test.txt");
console.log("文件 MD5: " + md5);
```

---

### files.path
返回相对路径对应的绝对路径

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| relativePath | string | 相对路径 |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| absolutePath | string | 绝对路径 |

**使用示例：**
```javascript
// 获取绝对路径
var absPath = files.path("test.txt");
console.log("绝对路径: " + absPath);
```

---

### files.listDir
列出文件夹path下的所有文件和文件夹

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| path | string | 文件夹路径 |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| items | array | 文件和文件夹名称数组 |

**使用示例：**
```javascript
// 列出文件夹内容
var items = files.listDir("/sdcard/Download");
for (var i = 0; i < items.length; i++) {
    console.log(items[i]);
}
```

---

## 综合使用示例

### 示例1：读写文件
```javascript
files.write("/sdcard/test.txt", "Hello World");
var content = files.read("/sdcard/test.txt");
console.log(content);
```