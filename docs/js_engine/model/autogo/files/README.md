# files 模块

## 模块简介

files 模块提供了文件系统操作功能，包括文件读写、复制、移动、删除等。

## 方法列表

### files.isDir
返回路径path是否是文件夹

**使用示例：**
```javascript
// 调用 files.isDir 方法
files.isDir();
```

---

### files.isFile
返回路径path是否是文件

**使用示例：**
```javascript
// 调用 files.isFile 方法
files.isFile();
```

---

### files.isEmptyDir
返回文件夹path是否为空文件夹

**使用示例：**
```javascript
// 调用 files.isEmptyDir 方法
files.isEmptyDir();
```

---

### files.create
创建一个文件或文件夹

**使用示例：**
```javascript
// 调用 files.create 方法
files.create();
```

---

### files.exists
返回在路径path处的文件是否存在

**使用示例：**
```javascript
// 调用 files.exists 方法
files.exists();
```

---

### files.ensureDir
确保路径path所在的文件夹存在

**使用示例：**
```javascript
// 调用 files.ensureDir 方法
files.ensureDir();
```

---

### files.read
读取文本文件path的所有内容并返回

**使用示例：**
```javascript
// 调用 files.read 方法
files.read();
```

---

### files.readBytes
读取文件path的所有内容并返回

**使用示例：**
```javascript
// 调用 files.readBytes 方法
files.readBytes();
```

---

### files.write
把text写入到文件path中

**使用示例：**
```javascript
// 调用 files.write 方法
files.write();
```

---

### files.writeBytes
把bytes写入到文件path中

**使用示例：**
```javascript
// 调用 files.writeBytes 方法
files.writeBytes();
```

---

### files.append
把text追加到文件path的末尾

**使用示例：**
```javascript
// 调用 files.append 方法
files.append();
```

---

### files.appendBytes
把bytes追加到文件path的末尾

**使用示例：**
```javascript
// 调用 files.appendBytes 方法
files.appendBytes();
```

---

### files.copy
复制文件

**使用示例：**
```javascript
// 调用 files.copy 方法
files.copy();
```

---

### files.move
移动文件

**使用示例：**
```javascript
// 调用 files.move 方法
files.move();
```

---

### files.rename
重命名文件

**使用示例：**
```javascript
// 调用 files.rename 方法
files.rename();
```

---

### files.remove
删除文件或文件夹

**使用示例：**
```javascript
// 调用 files.remove 方法
files.remove();
```

---

### files.getName
返回文件的文件名

**使用示例：**
```javascript
// 调用 files.getName 方法
files.getName();
```

---

### files.getNameWithoutExtension
返回不含拓展名的文件的文件名

**使用示例：**
```javascript
// 调用 files.getNameWithoutExtension 方法
files.getNameWithoutExtension();
```

---

### files.getExtension
返回文件的拓展名

**使用示例：**
```javascript
// 调用 files.getExtension 方法
files.getExtension();
```

---

### files.getMd5
返回文件的MD5值

**使用示例：**
```javascript
// 调用 files.getMd5 方法
files.getMd5();
```

---

### files.path
返回相对路径对应的绝对路径

**使用示例：**
```javascript
// 调用 files.path 方法
files.path();
```

---

### files.listDir
列出文件夹path下的所有文件和文件夹

**使用示例：**
```javascript
// 调用 files.listDir 方法
files.listDir();
```

---

## 综合使用示例

### 示例1：读写文件
```javascript
files.write("/sdcard/test.txt", "Hello World");
var content = files.read("/sdcard/test.txt");
console.log(content);
```