# utils 模块

## 模块简介

utils 模块提供了各种实用工具函数，包括日志记录、数据转换、加密解密、编码解码等。

## 方法列表

### utils.logI
记录一条INFO级别的日志

**使用示例：**
```lua
-- 调用 utils.logI 方法
utils.logI();
```

---

### utils.logE
记录一条ERROR级别的日志

**使用示例：**
```lua
-- 调用 utils.logE 方法
utils.logE();
```

---

### utils.toast
显示Toast提示

**使用示例：**
```lua
-- 调用 utils.toast 方法
utils.toast();
```

---

### utils.alert
显示Alert对话框

**使用示例：**
```lua
-- 调用 utils.alert 方法
utils.alert();
```

---

### utils.shell
执行shell命令并返回输出

**使用示例：**
```lua
-- 调用 utils.shell 方法
utils.shell();
```

---

### utils.random
返回指定范围内的随机整数

**使用示例：**
```lua
-- 调用 utils.random 方法
utils.random();
```

---

### utils.sleep
暂停当前线程指定的毫秒数

**使用示例：**
```lua
-- 调用 utils.sleep 方法
utils.sleep();
```

---

### utils.i2s
将整数转换为字符串

**使用示例：**
```lua
-- 调用 utils.i2s 方法
utils.i2s();
```

---

### utils.s2i
将字符串转换为整数

**使用示例：**
```lua
-- 调用 utils.s2i 方法
utils.s2i();
```

---

### utils.f2s
将浮点数转换为字符串

**使用示例：**
```lua
-- 调用 utils.f2s 方法
utils.f2s();
```

---

### utils.s2f
将字符串转换为浮点数

**使用示例：**
```lua
-- 调用 utils.s2f 方法
utils.s2f();
```

---

### utils.b2s
将布尔值转换为字符串

**使用示例：**
```lua
-- 调用 utils.b2s 方法
utils.b2s();
```

---

### utils.s2b
将字符串转换为布尔值

**使用示例：**
```lua
-- 调用 utils.s2b 方法
utils.s2b();
```

---

## 综合使用示例

### 示例1：日志和提示
```lua
utils.logI("TAG", "这是一条日志");
utils.toast("操作成功");
```