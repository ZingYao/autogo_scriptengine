# utils 模块

## 模块简介

utils 模块提供了各种实用工具函数，包括日志记录、数据转换、加密解密、编码解码等。

## 方法列表

### utils.logI
记录一条INFO级别的日志

**参数：**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| label | string | 是 | 日志标签，用于标识日志来源 |
| message | string | 是 | 日志消息内容，支持多个参数自动拼接 |

**返回值：**

| 类型 | 说明 |
|------|------|
| undefined | 无返回值 |

**使用示例：**
```javascript
// 记录一条简单的日志
utils.logI("TAG", "这是一条日志");

// 记录多条消息
utils.logI("TAG", "操作开始", "用户ID:", 123, "时间:", "2024-01-01");
```

---

### utils.logE
记录一条ERROR级别的日志

**参数：**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| label | string | 是 | 日志标签，用于标识日志来源 |
| message | string | 是 | 错误消息内容，支持多个参数自动拼接 |

**返回值：**

| 类型 | 说明 |
|------|------|
| undefined | 无返回值 |

**使用示例：**
```javascript
// 记录一条错误日志
utils.logE("TAG", "操作失败");

// 记录详细的错误信息
utils.logE("TAG", "连接超时", "服务器:", "192.168.1.1", "端口:", 8080);
```

---

### utils.toast
显示Toast提示

**参数：**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| message | string | 是 | 要显示的提示消息 |
| x | number | 否 | 在界面上显示的 X 坐标（传递-1使用默认坐标） |
| y | number | 否 | 在界面上显示的 Y 坐标（传递-1使用默认坐标） |
| duration | number | 否 | 提示显示的持续时间，单位为毫秒（传递-1使用默认2000毫秒） |

**返回值：**

| 类型 | 说明 |
|------|------|
| undefined | 无返回值 |

**使用示例：**
```javascript
// 显示一条提示
utils.toast("操作成功");

// 显示错误提示
utils.toast("网络连接失败，请检查网络设置");

// 带位置和时长的调用
utils.toast("操作成功", 100, 200, 3000);
```

---

### utils.alert
显示Alert对话框

**参数：**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| title | string | 是 | 对话框标题 |
| content | string | 是 | 对话框内容 |
| btn1Text | string | 否 | 第一个按钮文本（可选） |
| btn2Text | string | 否 | 第二个按钮文本（可选） |

**返回值：**

| 类型 | 说明 |
|------|------|
| int | 用户点击的按钮索引（0表示第一个按钮，1表示第二个按钮） |

**使用示例：**
```javascript
// 显示简单的确认对话框
let result = utils.alert("确认", "确定要删除吗？", "确定", "取消");
if (result === 0) {
    // 用户点击了确定
}

// 显示提示对话框
utils.alert("提示", "操作已完成", "知道了");
```

---

### utils.shell
执行shell命令并返回输出

**参数：**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| cmd | string | 是 | 要执行的shell命令 |

**返回值：**

| 类型 | 说明 |
|------|------|
| string | 命令执行的输出结果 |

**使用示例：**
```javascript
// 获取当前目录
let result = utils.shell("pwd");
console.log(result);

// 列出文件
let files = utils.shell("ls -la");
console.log(files);

// 执行自定义命令
let output = utils.shell("echo 'Hello World'");
```

---

### utils.random
返回指定范围内的随机整数

**参数：**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| min | int | 是 | 随机数的最小值（包含） |
| max | int | 是 | 随机数的最大值（包含） |

**返回值：**

| 类型 | 说明 |
|------|------|
| int | 在[min, max]范围内的随机整数 |

**使用示例：**
```javascript
// 生成1-100之间的随机数
let num = utils.random(1, 100);
console.log(num);

// 生成0-9之间的随机数
let digit = utils.random(0, 9);
console.log(digit);
```

---

### utils.sleep
暂停当前线程指定的毫秒数

**参数：**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| ms | int | 是 | 暂停的毫秒数 |

**返回值：**

| 类型 | 说明 |
|------|------|
| undefined | 无返回值 |

**使用示例：**
```javascript
// 暂停1秒
utils.sleep(1000);

// 暂停500毫秒
utils.sleep(500);

// 在循环中使用
for (let i = 0; i < 5; i++) {
    console.log("执行第" + i + "次");
    utils.sleep(1000); // 每次暂停1秒
}
```

---

### utils.i2s
将整数转换为字符串

**参数：**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| i | int | 是 | 要转换的整数 |

**返回值：**

| 类型 | 说明 |
|------|------|
| string | 转换后的字符串 |

**使用示例：**
```javascript
// 将整数转换为字符串
let str = utils.i2s(123);
console.log(str); // 输出: "123"

// 在字符串拼接中使用
let num = 456;
let message = "数字是: " + utils.i2s(num);
console.log(message); // 输出: "数字是: 456"
```

---

### utils.s2i
将字符串转换为整数

**参数：**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| s | string | 是 | 要转换的字符串 |

**返回值：**

| 类型 | 说明 |
|------|------|
| int | 转换后的整数 |

**使用示例：**
```javascript
// 将字符串转换为整数
let num = utils.s2i("789");
console.log(num); // 输出: 789

// 处理用户输入
let userInput = "100";
let value = utils.s2i(userInput);
console.log(value * 2); // 输出: 200
```

---

### utils.f2s
将浮点数转换为字符串

**参数：**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| f | float | 是 | 要转换的浮点数 |

**返回值：**

| 类型 | 说明 |
|------|------|
| string | 转换后的字符串 |

**使用示例：**
```javascript
// 将浮点数转换为字符串
let str = utils.f2s(3.14159);
console.log(str); // 输出: "3.14159"

// 在显示中使用
let price = 99.99;
let message = "价格: " + utils.f2s(price);
console.log(message); // 输出: "价格: 99.99"
```

---

### utils.s2f
将字符串转换为浮点数

**参数：**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| s | string | 是 | 要转换的字符串 |

**返回值：**

| 类型 | 说明 |
|------|------|
| float | 转换后的浮点数 |

**使用示例：**
```javascript
// 将字符串转换为浮点数
let num = utils.s2f("2.71828");
console.log(num); // 输出: 2.71828

// 处理用户输入
let userInput = "1.5";
let value = utils.s2f(userInput);
console.log(value * 2); // 输出: 3.0
```

---

### utils.b2s
将布尔值转换为字符串

**参数：**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| b | boolean | 是 | 要转换的布尔值 |

**返回值：**

| 类型 | 说明 |
|------|------|
| string | 转换后的字符串（"true" 或 "false"） |

**使用示例：**
```javascript
// 将布尔值转换为字符串
let str = utils.b2s(true);
console.log(str); // 输出: "true"

let str2 = utils.b2s(false);
console.log(str2); // 输出: "false"

// 在日志中使用
let isLoggedIn = true;
console.log("登录状态: " + utils.b2s(isLoggedIn));
```

---

### utils.s2b
将字符串转换为布尔值

**参数：**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| s | string | 是 | 要转换的字符串 |

**返回值：**

| 类型 | 说明 |
|------|------|
| boolean | 转换后的布尔值 |

**使用示例：**
```javascript
// 将字符串转换为布尔值
let flag = utils.s2b("true");
console.log(flag); // 输出: true

let flag2 = utils.s2b("false");
console.log(flag2); // 输出: false

// 处理配置文件
let config = "enabled";
let isEnabled = utils.s2b(config);
```

---

### utils.md5
计算字符串的MD5哈希值

**参数：**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| s | string | 是 | 要计算哈希值的字符串 |

**返回值：**

| 类型 | 说明 |
|------|------|
| string | 32位的MD5哈希值（十六进制字符串） |

**使用示例：**
```javascript
// 计算字符串的MD5值
let hash = utils.md5("Hello World");
console.log(hash); // 输出: "b10a8db164e0754105b7a99be72e3fe5"

// 计算密码的哈希值
let password = "mypassword123";
let passwordHash = utils.md5(password);
console.log("密码哈希: " + passwordHash);
```

---

### utils.sha1
计算字符串的SHA1哈希值

**参数：**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| s | string | 是 | 要计算哈希值的字符串 |

**返回值：**

| 类型 | 说明 |
|------|------|
| string | 40位的SHA1哈希值（十六进制字符串） |

**使用示例：**
```javascript
// 计算字符串的SHA1值
let hash = utils.sha1("Hello World");
console.log(hash); // 输出: "0a0a9f2a6772942557ab5355d76af442f8f65e01"

// 计算文件的SHA1值
let content = "文件内容";
let fileHash = utils.sha1(content);
console.log("文件哈希: " + fileHash);
```

---

### utils.sha256
计算字符串的SHA256哈希值

**参数：**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| s | string | 是 | 要计算哈希值的字符串 |

**返回值：**

| 类型 | 说明 |
|------|------|
| string | 64位的SHA256哈希值（十六进制字符串） |

**使用示例：**
```javascript
// 计算字符串的SHA256值
let hash = utils.sha256("Hello World");
console.log(hash); // 输出: "a591a6d40bf420404a011733cfb7b190d62c65bf0bcda32b57b277d9ad9f146e"

// 计算数据的SHA256值
let data = "重要数据";
let dataHash = utils.sha256(data);
console.log("数据哈希: " + dataHash);
```

---

### utils.base64Encode
将字符串进行Base64编码

**参数：**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| s | string | 是 | 要编码的字符串 |

**返回值：**

| 类型 | 说明 |
|------|------|
| string | Base64编码后的字符串 |

**使用示例：**
```javascript
// 将字符串进行Base64编码
let encoded = utils.base64Encode("Hello World");
console.log(encoded); // 输出: "SGVsbG8gV29ybGQ="

// 编码中文
let chinese = "你好世界";
let encodedChinese = utils.base64Encode(chinese);
console.log(encodedChinese);
```

---

### utils.base64Decode
将Base64编码的字符串解码

**参数：**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| s | string | 是 | Base64编码的字符串 |

**返回值：**

| 类型 | 说明 |
|------|------|
| string | 解码后的原始字符串，解码失败返回空字符串 |

**使用示例：**
```javascript
// 解码Base64字符串
let decoded = utils.base64Decode("SGVsbG8gV29ybGQ=");
console.log(decoded); // 输出: "Hello World"

// 解码中文
let encodedChinese = "5L2g5aW95LiW55WM";
let decodedChinese = utils.base64Decode(encodedChinese);
console.log(decodedChinese); // 输出: "你好世界"
```

---

### utils.urlEncode
对字符串进行URL编码

**参数：**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| s | string | 是 | 要编码的字符串 |

**返回值：**

| 类型 | 说明 |
|------|------|
| string | URL编码后的字符串 |

**使用示例：**
```javascript
// 对字符串进行URL编码
let encoded = utils.urlEncode("Hello World");
console.log(encoded); // 输出: "Hello+World"

// 编码中文
let chinese = "你好世界";
let encodedChinese = utils.urlEncode(chinese);
console.log(encodedChinese); // 输出: "%E4%BD%A0%E5%A5%BD%E4%B8%96%E7%95%8C"

// 编码特殊字符
let special = "a&b=c";
let encodedSpecial = utils.urlEncode(special);
console.log(encodedSpecial); // 输出: "a%26b%3Dc"
```

---

### utils.urlDecode
对URL编码的字符串进行解码

**参数：**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| s | string | 是 | URL编码的字符串 |

**返回值：**

| 类型 | 说明 |
|------|------|
| string | 解码后的原始字符串，解码失败返回空字符串 |

**使用示例：**
```javascript
// 解码URL字符串
let decoded = utils.urlDecode("Hello+World");
console.log(decoded); // 输出: "Hello World"

// 解码中文
let encodedChinese = "%E4%BD%A0%E5%A5%BD%E4%B8%96%E7%95%8C";
let decodedChinese = utils.urlDecode(encodedChinese);
console.log(decodedChinese); // 输出: "你好世界"

// 解码特殊字符
let encodedSpecial = "a%26b%3Dc";
let decodedSpecial = utils.urlDecode(encodedSpecial);
console.log(decodedSpecial); // 输出: "a&b=c"
```

---

### utils.htmlEncode
对字符串进行HTML实体编码

**参数：**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| s | string | 是 | 要编码的字符串 |

**返回值：**

| 类型 | 说明 |
|------|------|
| string | HTML实体编码后的字符串 |

**使用示例：**
```javascript
// 对字符串进行HTML编码
let encoded = utils.htmlEncode("<div>Hello</div>");
console.log(encoded); // 输出: "&lt;div&gt;Hello&lt;/div&gt;"

// 编码特殊字符
let special = "<script>alert('XSS')</script>";
let encodedSpecial = utils.htmlEncode(special);
console.log(encodedSpecial); // 输出: "&lt;script&gt;alert('XSS')&lt;/script&gt;"
```

---

### utils.htmlDecode
对HTML实体编码的字符串进行解码

**参数：**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| s | string | 是 | HTML实体编码的字符串 |

**返回值：**

| 类型 | 说明 |
|------|------|
| string | 解码后的原始字符串 |

**使用示例：**
```javascript
// 解码HTML字符串
let decoded = utils.htmlDecode("&lt;div&gt;Hello&lt;/div&gt;");
console.log(decoded); // 输出: "<div>Hello</div>"

// 解码特殊字符
let encodedSpecial = "&lt;script&gt;alert('XSS')&lt;/script&gt;";
let decodedSpecial = utils.htmlDecode(encodedSpecial);
console.log(decodedSpecial); // 输出: "<script>alert('XSS')</script>"
```

---

### utils.jsonEncode
将对象转换为JSON字符串

**参数：**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| v | object | 是 | 要转换的对象 |

**返回值：**

| 类型 | 说明 |
|------|------|
| string | JSON格式的字符串，转换失败返回空字符串 |

**使用示例：**
```javascript
// 将对象转换为JSON字符串
let obj = { name: "张三", age: 25, city: "北京" };
let jsonStr = utils.jsonEncode(obj);
console.log(jsonStr); // 输出: '{"name":"张三","age":25,"city":"北京"}'

// 转换数组
let arr = [1, 2, 3, 4, 5];
let jsonArr = utils.jsonEncode(arr);
console.log(jsonArr); // 输出: '[1,2,3,4,5]'
```

---

### utils.jsonDecode
将JSON字符串解析为对象

**参数：**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| s | string | 是 | JSON格式的字符串 |

**返回值：**

| 类型 | 说明 |
|------|------|
| object | 解析后的对象，解析失败返回undefined |

**使用示例：**
```javascript
// 解析JSON字符串
let jsonStr = '{"name":"张三","age":25,"city":"北京"}';
let obj = utils.jsonDecode(jsonStr);
console.log(obj.name); // 输出: "张三"
console.log(obj.age); // 输出: 25

// 解析数组
let jsonArr = '[1,2,3,4,5]';
let arr = utils.jsonDecode(jsonArr);
console.log(arr[0]); // 输出: 1
```

---

### utils.timestamp
获取当前时间戳（秒）

**参数：**

无

**返回值：**

| 类型 | 说明 |
|------|------|
| int | 当前时间戳（从1970年1月1日开始的秒数） |

**使用示例：**
```javascript
// 获取当前时间戳
let ts = utils.timestamp();
console.log(ts); // 输出: 当前时间戳，如 1704067200

// 记录操作时间
let startTime = utils.timestamp();
// 执行一些操作
utils.sleep(2000);
let endTime = utils.timestamp();
console.log("耗时: " + (endTime - startTime) + " 秒");
```

---

## 综合使用示例

### 示例1：日志和提示
```javascript
// 记录INFO日志
utils.logI("TAG", "这是一条日志");

// 显示Toast提示
utils.toast("操作成功");

// 记录ERROR日志
utils.logE("TAG", "操作失败");
```

### 示例2：数据转换
```javascript
// 整数转字符串
let num = 123;
let str = utils.i2s(num);
console.log(str); // "123"

// 字符串转整数
let strNum = "456";
let intNum = utils.s2i(strNum);
console.log(intNum); // 456

// 浮点数转字符串
let floatNum = 3.14;
let floatStr = utils.f2s(floatNum);
console.log(floatStr); // "3.14"
```

### 示例3：加密和编码
```javascript
// 计算MD5
let md5Hash = utils.md5("Hello World");
console.log(md5Hash);

// Base64编码
let encoded = utils.base64Encode("Hello World");
console.log(encoded);

// URL编码
let urlEncoded = utils.urlEncode("你好世界");
console.log(urlEncoded);
```

### 示例4：JSON处理
```javascript
// 对象转JSON
let obj = { name: "张三", age: 25 };
let jsonStr = utils.jsonEncode(obj);
console.log(jsonStr);

// JSON转对象
let parsedObj = utils.jsonDecode(jsonStr);
console.log(parsedObj.name);
```
