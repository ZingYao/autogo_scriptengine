# utils 模块

## 模块简介

utils 模块提供了各种实用工具函数，包括日志记录、数据转换、加密解密、编码解码等。

## 方法列表

### utils.logI
记录一条INFO级别的日志

**使用示例：**
```javascript
// 调用 utils.logI 方法
utils.logI();
```

---

### utils.logE
记录一条ERROR级别的日志

**使用示例：**
```javascript
// 调用 utils.logE 方法
utils.logE();
```

---

### utils.toast
显示Toast提示

**使用示例：**
```javascript
// 调用 utils.toast 方法
utils.toast();
```

---

### utils.alert
显示Alert对话框

**使用示例：**
```javascript
// 调用 utils.alert 方法
utils.alert();
```

---

### utils.shell
执行shell命令并返回输出

**使用示例：**
```javascript
// 调用 utils.shell 方法
utils.shell();
```

---

### utils.random
返回指定范围内的随机整数

**使用示例：**
```javascript
// 调用 utils.random 方法
utils.random();
```

---

### utils.sleep
暂停当前线程指定的毫秒数

**使用示例：**
```javascript
// 调用 utils.sleep 方法
utils.sleep();
```

---

### utils.i2s
将整数转换为字符串

**使用示例：**
```javascript
// 调用 utils.i2s 方法
utils.i2s();
```

---

### utils.s2i
将字符串转换为整数

**使用示例：**
```javascript
// 调用 utils.s2i 方法
utils.s2i();
```

---

### utils.f2s
将浮点数转换为字符串

**使用示例：**
```javascript
// 调用 utils.f2s 方法
utils.f2s();
```

---

### utils.s2f
将字符串转换为浮点数

**使用示例：**
```javascript
// 调用 utils.s2f 方法
utils.s2f();
```

---

### utils.b2s
将布尔值转换为字符串

**使用示例：**
```javascript
// 调用 utils.b2s 方法
utils.b2s();
```

---

### utils.s2b
将字符串转换为布尔值

**使用示例：**
```javascript
// 调用 utils.s2b 方法
utils.s2b();
```

---

### utils.md5
计算字符串的MD5哈希值

**使用示例：**
```javascript
// 调用 utils.md5 方法
utils.md5();
```

---

### utils.sha1
计算字符串的SHA1哈希值

**使用示例：**
```javascript
// 调用 utils.sha1 方法
utils.sha1();
```

---

### utils.sha256
计算字符串的SHA256哈希值

**使用示例：**
```javascript
// 调用 utils.sha256 方法
utils.sha256();
```

---

### utils.base64Encode
将字符串进行Base64编码

**使用示例：**
```javascript
// 调用 utils.base64Encode 方法
utils.base64Encode();
```

---

### utils.base64Decode
将Base64编码的字符串解码

**使用示例：**
```javascript
// 调用 utils.base64Decode 方法
utils.base64Decode();
```

---

### utils.urlEncode
对字符串进行URL编码

**使用示例：**
```javascript
// 调用 utils.urlEncode 方法
utils.urlEncode();
```

---

### utils.urlDecode
对URL编码的字符串进行解码

**使用示例：**
```javascript
// 调用 utils.urlDecode 方法
utils.urlDecode();
```

---

### utils.htmlEncode
对字符串进行HTML实体编码

**使用示例：**
```javascript
// 调用 utils.htmlEncode 方法
utils.htmlEncode();
```

---

### utils.htmlDecode
对HTML实体编码的字符串进行解码

**使用示例：**
```javascript
// 调用 utils.htmlDecode 方法
utils.htmlDecode();
```

---

### utils.jsonEncode
将对象转换为JSON字符串

**使用示例：**
```javascript
// 调用 utils.jsonEncode 方法
utils.jsonEncode();
```

---

### utils.jsonDecode
将JSON字符串解析为对象

**使用示例：**
```javascript
// 调用 utils.jsonDecode 方法
utils.jsonDecode();
```

---

### utils.timestamp
获取当前时间戳（秒）

**使用示例：**
```javascript
// 调用 utils.timestamp 方法
utils.timestamp();
```

---

## 综合使用示例

### 示例1：日志和提示
```javascript
utils.logI("TAG", "这是一条日志");
utils.toast("操作成功");
```