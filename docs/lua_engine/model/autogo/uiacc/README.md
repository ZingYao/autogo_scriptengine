# uiacc 模块

## 模块简介

uiacc 模块提供了无障碍服务（Accessibility）的功能，用于查找和操作界面元素。

## 方法列表

### uiacc.new
创建一个新的Accessibility对象

**参数：**
- `displayId` (int): 显示设备 ID（可选，默认 0）

**返回值：**
- Uiacc 对象: 用于查询和操作界面元素

**使用示例：**
```lua
-- 创建 Accessibility 对象
local acc = uiacc.new()
local acc2 = uiacc.new(1)  -- 指定显示设备 ID
```

---

### uiacc.text
根据文本内容查找元素

**参数：**
- `u` (Uiacc): Accessibility 对象
- `value` (string): 要查找的文本内容

**返回值：**
- Uiacc 对象: 匹配的元素

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.text("确定")
```

---

### uiacc.textContains
根据文本包含内容查找元素

**参数：**
- `u` (Uiacc): Accessibility 对象
- `value` (string): 要查找的文本内容（包含关系）

**返回值：**
- Uiacc 对象: 匹配的元素

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.textContains("登录")
```

---

### uiacc.textStartsWith
根据文本前缀查找元素

**参数：**
- `u` (Uiacc): Accessibility 对象
- `value` (string): 要查找的文本前缀

**返回值：**
- Uiacc 对象: 匹配的元素

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.textStartsWith("用户")
```

---

### uiacc.textEndsWith
根据文本后缀查找元素

**参数：**
- `u` (Uiacc): Accessibility 对象
- `value` (string): 要查找的文本后缀

**返回值：**
- Uiacc 对象: 匹配的元素

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.textEndsWith("按钮")
```

---

### uiacc.textMatches
根据文本正则表达式查找元素

**参数：**
- `u` (Uiacc): Accessibility 对象
- `value` (string): 正则表达式

**返回值：**
- Uiacc 对象: 匹配的元素

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.textMatches(".*登录.*")
```

---

### uiacc.desc
根据描述内容查找元素

**参数：**
- `u` (Uiacc): Accessibility 对象
- `value` (string): 要查找的描述内容

**返回值：**
- Uiacc 对象: 匹配的元素

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.desc("点击登录")
```

---

### uiacc.descContains
根据描述包含内容查找元素

**参数：**
- `u` (Uiacc): Accessibility 对象
- `value` (string): 要查找的描述内容（包含关系）

**返回值：**
- Uiacc 对象: 匹配的元素

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.descContains("登录")
```

---

### uiacc.descStartsWith
根据描述前缀查找元素

**参数：**
- `u` (Uiacc): Accessibility 对象
- `value` (string): 要查找的描述前缀

**返回值：**
- Uiacc 对象: 匹配的元素

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.descStartsWith("请")
```

---

### uiacc.descEndsWith
根据描述后缀查找元素

**参数：**
- `u` (Uiacc): Accessibility 对象
- `value` (string): 要查找的描述后缀

**返回值：**
- Uiacc 对象: 匹配的元素

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.descEndsWith("按钮")
```

---

### uiacc.descMatches
根据描述正则表达式查找元素

**参数：**
- `u` (Uiacc): Accessibility 对象
- `value` (string): 正则表达式

**返回值：**
- Uiacc 对象: 匹配的元素

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.descMatches(".*点击.*")
```

---

### uiacc.id
根据 ID 查找元素

**参数：**
- `u` (Uiacc): Accessibility 对象
- `value` (string): 要查找的元素 ID

**返回值：**
- Uiacc 对象: 匹配的元素

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.id("login_button")
```

---

### uiacc.idContains
根据 ID 包含内容查找元素

**参数：**
- `u` (Uiacc): Accessibility 对象
- `value` (string): 要查找的 ID 内容（包含关系）

**返回值：**
- Uiacc 对象: 匹配的元素

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.idContains("button")
```

---

### uiacc.idStartsWith
根据 ID 前缀查找元素

**参数：**
- `u` (Uiacc): Accessibility 对象
- `value` (string): 要查找的 ID 前缀

**返回值：**
- Uiacc 对象: 匹配的元素

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.idStartsWith("btn_")
```

---

### uiacc.idEndsWith
根据 ID 后缀查找元素

**参数：**
- `u` (Uiacc): Accessibility 对象
- `value` (string): 要查找的 ID 后缀

**返回值：**
- Uiacc 对象: 匹配的元素

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.idEndsWith("_submit")
```

---

### uiacc.idMatches
根据 ID 正则表达式查找元素

**参数：**
- `u` (Uiacc): Accessibility 对象
- `value` (string): 正则表达式

**返回值：**
- Uiacc 对象: 匹配的元素

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.idMatches("btn_.*")
```

---

### uiacc.className
根据类名查找元素

**参数：**
- `u` (Uiacc): Accessibility 对象
- `value` (string): 要查找的类名

**返回值：**
- Uiacc 对象: 匹配的元素

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.className("Button")
```

---

### uiacc.classNameContains
根据类名包含内容查找元素

**参数：**
- `u` (Uiacc): Accessibility 对象
- `value` (string): 要查找的类名内容（包含关系）

**返回值：**
- Uiacc 对象: 匹配的元素

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.classNameContains("Button")
```

---

### uiacc.classNameStartsWith
根据类名前缀查找元素

**参数：**
- `u` (Uiacc): Accessibility 对象
- `value` (string): 要查找的类名前缀

**返回值：**
- Uiacc 对象: 匹配的元素

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.classNameStartsWith("android.widget")
```

---

### uiacc.classNameEndsWith
根据类名后缀查找元素

**参数：**
- `u` (Uiacc): Accessibility 对象
- `value` (string): 要查找的类名后缀

**返回值：**
- Uiacc 对象: 匹配的元素

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.classNameEndsWith("Button")
```

---

### uiacc.classNameMatches
根据类名正则表达式查找元素

**参数：**
- `u` (Uiacc): Accessibility 对象
- `value` (string): 正则表达式

**返回值：**
- Uiacc 对象: 匹配的元素

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.classNameMatches(".*Button.*")
```

---

### uiacc.packageName
根据包名查找元素

**参数：**
- `u` (Uiacc): Accessibility 对象
- `value` (string): 要查找的包名

**返回值：**
- Uiacc 对象: 匹配的元素

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.packageName("com.example.app")
```

---

### uiacc.packageNameContains
根据包名包含内容查找元素

**参数：**
- `u` (Uiacc): Accessibility 对象
- `value` (string): 要查找的包名内容（包含关系）

**返回值：**
- Uiacc 对象: 匹配的元素

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.packageNameContains("example")
```

---

### uiacc.packageNameStartsWith
根据包名前缀查找元素

**参数：**
- `u` (Uiacc): Accessibility 对象
- `value` (string): 要查找的包名前缀

**返回值：**
- Uiacc 对象: 匹配的元素

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.packageNameStartsWith("com.example")
```

---

### uiacc.packageNameEndsWith
根据包名后缀查找元素

**参数：**
- `u` (Uiacc): Accessibility 对象
- `value` (string): 要查找的包名后缀

**返回值：**
- Uiacc 对象: 匹配的元素

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.packageNameEndsWith(".app")
```

---

### uiacc.packageNameMatches
根据包名正则表达式查找元素

**参数：**
- `u` (Uiacc): Accessibility 对象
- `value` (string): 正则表达式

**返回值：**
- Uiacc 对象: 匹配的元素

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.packageNameMatches("com\\..*\\.app")
```

---

### uiacc.bounds
根据边界区域查找元素

**参数：**
- `u` (Uiacc): Accessibility 对象
- `left` (int): 左边界
- `top` (int): 上边界
- `right` (int): 右边界
- `bottom` (int): 下边界

**返回值：**
- Uiacc 对象: 匹配的元素

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.bounds(100, 100, 200, 200)
```

---

### uiacc.boundsInside
查找在指定边界区域内的元素

**参数：**
- `u` (Uiacc): Accessibility 对象
- `left` (int): 左边界
- `top` (int): 上边界
- `right` (int): 右边界
- `bottom` (int): 下边界

**返回值：**
- Uiacc 对象: 匹配的元素

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.boundsInside(100, 100, 200, 200)
```

---

### uiacc.boundsContains
查找包含指定边界区域的元素

**参数：**
- `u` (Uiacc): Accessibility 对象
- `left` (int): 左边界
- `top` (int): 上边界
- `right` (int): 右边界
- `bottom` (int): 下边界

**返回值：**
- Uiacc 对象: 匹配的元素

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.boundsContains(100, 100, 200, 200)
```

---

### uiacc.drawingOrder
根据绘制顺序查找元素

**参数：**
- `u` (Uiacc): Accessibility 对象
- `value` (int): 绘制顺序值

**返回值：**
- Uiacc 对象: 匹配的元素

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.drawingOrder(0)
```

---

### uiacc.clickable
根据可点击属性查找元素

**参数：**
- `u` (Uiacc): Accessibility 对象
- `value` (bool): 是否可点击

**返回值：**
- Uiacc 对象: 匹配的元素

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.clickable(true)
```

---

### uiacc.longClickable
根据可长按点击属性查找元素

**参数：**
- `u` (Uiacc): Accessibility 对象
- `value` (bool): 是否可长按点击

**返回值：**
- Uiacc 对象: 匹配的元素

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.longClickable(true)
```

---

### uiacc.checkable
根据可勾选属性查找元素

**参数：**
- `u` (Uiacc): Accessibility 对象
- `value` (bool): 是否可勾选

**返回值：**
- Uiacc 对象: 匹配的元素

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.checkable(true)
```

---

### uiacc.selected
根据选中状态查找元素

**参数：**
- `u` (Uiacc): Accessibility 对象
- `value` (bool): 是否选中

**返回值：**
- Uiacc 对象: 匹配的元素

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.selected(true)
```

---

### uiacc.enabled
根据启用状态查找元素

**参数：**
- `u` (Uiacc): Accessibility 对象
- `value` (bool): 是否启用

**返回值：**
- Uiacc 对象: 匹配的元素

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.enabled(true)
```

---

### uiacc.scrollable
根据可滚动属性查找元素

**参数：**
- `u` (Uiacc): Accessibility 对象
- `value` (bool): 是否可滚动

**返回值：**
- Uiacc 对象: 匹配的元素

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.scrollable(true)
```

---

### uiacc.editable
根据可编辑属性查找元素

**参数：**
- `u` (Uiacc): Accessibility 对象
- `value` (bool): 是否可编辑

**返回值：**
- Uiacc 对象: 匹配的元素

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.editable(true)
```

---

### uiacc.multiLine
根据多行属性查找元素

**参数：**
- `u` (Uiacc): Accessibility 对象
- `value` (bool): 是否多行

**返回值：**
- Uiacc 对象: 匹配的元素

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.multiLine(true)
```

---

### uiacc.checked
根据勾选状态查找元素

**参数：**
- `u` (Uiacc): Accessibility 对象
- `value` (bool): 是否勾选

**返回值：**
- Uiacc 对象: 匹配的元素

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.checked(true)
```

---

### uiacc.focusable
根据可聚焦属性查找元素

**参数：**
- `u` (Uiacc): Accessibility 对象
- `value` (bool): 是否可聚焦

**返回值：**
- Uiacc 对象: 匹配的元素

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.focusable(true)
```

---

### uiacc.dismissable
根据可关闭属性查找元素

**参数：**
- `u` (Uiacc): Accessibility 对象
- `value` (bool): 是否可关闭

**返回值：**
- Uiacc 对象: 匹配的元素

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.dismissable(true)
```

---

### uiacc.focused
根据聚焦状态查找元素

**参数：**
- `u` (Uiacc): Accessibility 对象
- `value` (bool): 是否聚焦

**返回值：**
- Uiacc 对象: 匹配的元素

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.focused(true)
```

---

### uiacc.contextClickable
根据上下文可点击属性查找元素

**参数：**
- `u` (Uiacc): Accessibility 对象
- `value` (bool): 是否上下文可点击

**返回值：**
- Uiacc 对象: 匹配的元素

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.contextClickable(true)
```

---

### uiacc.index
根据索引查找元素

**参数：**
- `u` (Uiacc): Accessibility 对象
- `value` (int): 元素索引

**返回值：**
- Uiacc 对象: 匹配的元素

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.index(0)
```

---

### uiacc.visible
根据可见性查找元素

**参数：**
- `u` (Uiacc): Accessibility 对象
- `value` (bool): 是否可见

**返回值：**
- Uiacc 对象: 匹配的元素

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.visible(true)
```

---

### uiacc.password
根据密码属性查找元素

**参数：**
- `u` (Uiacc): Accessibility 对象
- `value` (bool): 是否为密码输入框

**返回值：**
- Uiacc 对象: 匹配的元素

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.password(true)
```

---

### uiacc.click
点击匹配的元素

**参数：**
- `u` (Uiacc): Accessibility 对象
- `text` (string): 点击后的文本内容（可选）

**返回值：**
- boolean: 点击成功返回 true，失败返回 false

**使用示例：**
```lua
local acc = uiacc.new()
local success = acc.click("确定")
```

---

### uiacc.waitFor
等待元素出现

**参数：**
- `u` (Uiacc): Accessibility 对象
- `timeout` (int): 超时时间（毫秒）

**返回值：**
- UiObject: 找到的元素对象，超时返回 nil

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.waitFor(5000)
if element ~= nil then
    print("找到元素")
end
```

---

### uiacc.findOnce
查找第一个匹配的元素

**参数：**
- `u` (Uiacc): Accessibility 对象

**返回值：**
- UiObject: 找到的元素对象，未找到返回 nil

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.findOnce()
if element ~= nil then
    element:objClick()
end
```

---

### uiacc.find
查找所有匹配的元素

**参数：**
- `u` (Uiacc): Accessibility 对象

**返回值：**
- Array: 元素对象数组

**使用示例：**
```lua
local acc = uiacc.new()
local elements = acc.find()
for i, element in ipairs(elements) do
    print("元素 " .. i .. ": " .. element:getText())
end
```

---

### uiacc.release
释放 Accessibility 对象资源

**参数：**
- `u` (Uiacc): Accessibility 对象

**使用示例：**
```lua
local acc = uiacc.new()
-- 使用完毕后释放
acc.release()
```

---

## UiObject 对象方法（操作方法）

### uiacc.objClick
点击元素

**参数：**
- `obj` (UiObject): UiObject 对象

**返回值：**
- boolean: 点击成功返回 true，失败返回 false

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.findOnce()
element:objClick()
```

---

### uiacc.clickCenter
点击元素中心

**参数：**
- `obj` (UiObject): UiObject 对象

**返回值：**
- boolean: 点击成功返回 true，失败返回 false

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.findOnce()
element:clickCenter()
```

---

### uiacc.clickLongClick
长按点击元素

**参数：**
- `obj` (UiObject): UiObject 对象

**返回值：**
- boolean: 点击成功返回 true，失败返回 false

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.findOnce()
element:clickLongClick()
```

---

### uiacc.copy
复制元素内容

**参数：**
- `obj` (UiObject): UiObject 对象

**返回值：**
- boolean: 复制成功返回 true，失败返回 false

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.findOnce()
element:copy()
```

---

### uiacc.cut
剪切元素内容

**参数：**
- `obj` (UiObject): UiObject 对象

**返回值：**
- boolean: 剪切成功返回 true，失败返回 false

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.findOnce()
element:cut()
```

---

### uiacc.paste
粘贴内容到元素

**参数：**
- `obj` (UiObject): UiObject 对象

**返回值：**
- boolean: 粘贴成功返回 true，失败返回 false

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.findOnce()
element:paste()
```

---

### uiacc.scrollForward
向前滚动

**参数：**
- `obj` (UiObject): UiObject 对象

**返回值：**
- boolean: 滚动成功返回 true，失败返回 false

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.findOnce()
element:scrollForward()
```

---

### uiacc.scrollBackward
向后滚动

**参数：**
- `obj` (UiObject): UiObject 对象

**返回值：**
- boolean: 滚动成功返回 true，失败返回 false

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.findOnce()
element:scrollBackward()
```

---

### uiacc.collapse
折叠元素

**参数：**
- `obj` (UiObject): UiObject 对象

**返回值：**
- boolean: 折叠成功返回 true，失败返回 false

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.findOnce()
element:collapse()
```

---

### uiacc.expand
展开元素

**参数：**
- `obj` (UiObject): UiObject 对象

**返回值：**
- boolean: 展开成功返回 true，失败返回 false

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.findOnce()
element:expand()
```

---

### uiacc.show
显示元素

**参数：**
- `obj` (UiObject): UiObject 对象

**返回值：**
- boolean: 显示成功返回 true，失败返回 false

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.findOnce()
element:show()
```

---

### uiacc.select
选中元素

**参数：**
- `obj` (UiObject): UiObject 对象

**返回值：**
- boolean: 选中成功返回 true，失败返回 false

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.findOnce()
element:select()
```

---

### uiacc.clearSelect
清除选中状态

**参数：**
- `obj` (UiObject): UiObject 对象

**返回值：**
- boolean: 清除成功返回 true，失败返回 false

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.findOnce()
element:clearSelect()
```

---

### uiacc.setSelection
设置选中范围

**参数：**
- `obj` (UiObject): UiObject 对象
- `start` (int): 起始位置
- `end` (int): 结束位置

**返回值：**
- boolean: 设置成功返回 true，失败返回 false

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.findOnce()
element:setSelection(0, 10)
```

---

### uiacc.setVisibleToUser
设置元素对用户可见性

**参数：**
- `obj` (UiObject): UiObject 对象
- `isVisible` (bool): 是否可见

**返回值：**
- boolean: 设置成功返回 true，失败返回 false

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.findOnce()
element:setVisibleToUser(true)
```

---

### uiacc.setText
设置元素文本

**参数：**
- `obj` (UiObject): UiObject 对象
- `str` (string): 要设置的文本

**返回值：**
- boolean: 设置成功返回 true，失败返回 false

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.findOnce()
element:setText("新的文本")
```

---

### uiacc.getClickable
获取可点击属性

**参数：**
- `obj` (UiObject): UiObject 对象

**返回值：**
- boolean: 是否可点击

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.findOnce()
print("可点击: " .. tostring(element:getClickable()))
```

---

### uiacc.getLongClickable
获取可长按点击属性

**参数：**
- `obj` (UiObject): UiObject 对象

**返回值：**
- boolean: 是否可长按点击

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.findOnce()
print("可长按点击: " .. tostring(element:getLongClickable()))
```

---

### uiacc.getCheckable
获取可勾选属性

**参数：**
- `obj` (UiObject): UiObject 对象

**返回值：**
- boolean: 是否可勾选

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.findOnce()
print("可勾选: " .. tostring(element:getCheckable()))
```

---

### uiacc.getSelected
获取选中状态

**参数：**
- `obj` (UiObject): UiObject 对象

**返回值：**
- boolean: 是否选中

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.findOnce()
print("选中状态: " .. tostring(element:getSelected()))
```

---

### uiacc.getEnabled
获取启用状态

**参数：**
- `obj` (UiObject): UiObject 对象

**返回值：**
- boolean: 是否启用

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.findOnce()
print("启用状态: " .. tostring(element:getEnabled()))
```

---

### uiacc.getScrollable
获取可滚动属性

**参数：**
- `obj` (UiObject): UiObject 对象

**返回值：**
- boolean: 是否可滚动

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.findOnce()
print("可滚动: " .. tostring(element:getScrollable()))
```

---

### uiacc.getEditable
获取可编辑属性

**参数：**
- `obj` (UiObject): UiObject 对象

**返回值：**
- boolean: 是否可编辑

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.findOnce()
print("可编辑: " .. tostring(element:getEditable()))
```

---

### uiacc.getMultiLine
获取多行属性

**参数：**
- `obj` (UiObject): UiObject 对象

**返回值：**
- boolean: 是否多行

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.findOnce()
print("多行: " .. tostring(element:getMultiLine()))
```

---

### uiacc.getChecked
获取勾选状态

**参数：**
- `obj` (UiObject): UiObject 对象

**返回值：**
- boolean: 是否勾选

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.findOnce()
print("勾选状态: " .. tostring(element:getChecked()))
```

---

### uiacc.getFocused
获取聚焦状态

**参数：**
- `obj` (UiObject): UiObject 对象

**返回值：**
- boolean: 是否聚焦

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.findOnce()
print("聚焦状态: " .. tostring(element:getFocused()))
```

---

### uiacc.getFocusable
获取可聚焦属性

**参数：**
- `obj` (UiObject): UiObject 对象

**返回值：**
- boolean: 是否可聚焦

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.findOnce()
print("可聚焦: " .. tostring(element:getFocusable()))
```

---

### uiacc.getDismissable
获取可关闭属性

**参数：**
- `obj` (UiObject): UiObject 对象

**返回值：**
- boolean: 是否可关闭

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.findOnce()
print("可关闭: " .. tostring(element:getDismissable()))
```

---

### uiacc.getContextClickable
获取上下文可点击属性

**参数：**
- `obj` (UiObject): UiObject 对象

**返回值：**
- boolean: 是否上下文可点击

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.findOnce()
print("上下文可点击: " .. tostring(element:getContextClickable()))
```

---

### uiacc.getAccessibilityFocused
获取无障碍聚焦状态

**参数：**
- `obj` (UiObject): UiObject 对象

**返回值：**
- boolean: 是否无障碍聚焦

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.findOnce()
print("无障碍聚焦: " .. tostring(element:getAccessibilityFocused()))
```

---

### uiacc.getChildCount
获取子元素数量

**参数：**
- `obj` (UiObject): UiObject 对象

**返回值：**
- int: 子元素数量

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.findOnce()
print("子元素数量: " .. element:getChildCount())
```

---

### uiacc.getDrawingOrder
获取绘制顺序

**参数：**
- `obj` (UiObject): UiObject 对象

**返回值：**
- int: 绘制顺序

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.findOnce()
print("绘制顺序: " .. element:getDrawingOrder())
```

---

### uiacc.getIndex
获取元素索引

**参数：**
- `obj` (UiObject): UiObject 对象

**返回值：**
- int: 元素索引

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.findOnce()
print("元素索引: " .. element:getIndex())
```

---

### uiacc.getBounds
获取元素边界

**参数：**
- `obj` (UiObject): UiObject 对象

**返回值：**
- Table: 包含 left, top, right, bottom 属性的表

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.findOnce()
local bounds = element:getBounds()
print("边界: " .. bounds.left .. ", " .. bounds.top .. ", " .. bounds.right .. ", " .. bounds.bottom)
```

---

### uiacc.getBoundsInParent
获取元素在父容器中的边界

**参数：**
- `obj` (UiObject): UiObject 对象

**返回值：**
- Table: 包含 left, top, right, bottom 属性的表

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.findOnce()
local bounds = element:getBoundsInParent()
print("父容器边界: " .. bounds.left .. ", " .. bounds.top .. ", " .. bounds.right .. ", " .. bounds.bottom)
```

---

### uiacc.getId
获取元素 ID

**参数：**
- `obj` (UiObject): UiObject 对象

**返回值：**
- string: 元素 ID

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.findOnce()
print("元素 ID: " .. element:getId())
```

---

### uiacc.getText
获取元素文本

**参数：**
- `obj` (UiObject): UiObject 对象

**返回值：**
- string: 元素文本

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.findOnce()
print("元素文本: " .. element:getText())
```

---

### uiacc.getDesc
获取元素描述

**参数：**
- `obj` (UiObject): UiObject 对象

**返回值：**
- string: 元素描述

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.findOnce()
print("元素描述: " .. element:getDesc())
```

---

### uiacc.getPackageName
获取元素包名

**参数：**
- `obj` (UiObject): UiObject 对象

**返回值：**
- string: 元素包名

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.findOnce()
print("元素包名: " .. element:getPackageName())
```

---

### uiacc.getClassName
获取元素类名

**参数：**
- `obj` (UiObject): UiObject 对象

**返回值：**
- string: 元素类名

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.findOnce()
print("元素类名: " .. element:getClassName())
```

---

### uiacc.getParent
获取父元素

**参数：**
- `obj` (UiObject): UiObject 对象

**返回值：**
- UiObject: 父元素对象，无父元素返回 nil

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.findOnce()
local parent = element:getParent()
if parent ~= nil then
    print("找到父元素")
end
```

---

### uiacc.getChild
获取指定索引的子元素

**参数：**
- `obj` (UiObject): UiObject 对象
- `index` (int): 子元素索引

**返回值：**
- UiObject: 子元素对象，无子元素返回 nil

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.findOnce()
local child = element:getChild(0)
if child ~= nil then
    print("找到第一个子元素")
end
```

---

### uiacc.getChildren
获取所有子元素

**参数：**
- `obj` (UiObject): UiObject 对象

**返回值：**
- Array: 子元素对象数组

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.findOnce()
local children = element:getChildren()
for i, child in ipairs(children) do
    print("子元素 " .. i .. ": " .. child:getText())
end
```

---

### uiacc.getVisible
获取可见性

**参数：**
- `obj` (UiObject): UiObject 对象

**返回值：**
- boolean: 是否可见

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.findOnce()
print("可见性: " .. tostring(element:getVisible()))
```

---

### uiacc.getPassword
获取密码属性

**参数：**
- `obj` (UiObject): UiObject 对象

**返回值：**
- boolean: 是否为密码输入框

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.findOnce()
print("密码属性: " .. tostring(element:getPassword()))
```

---

### uiacc.toString
获取元素字符串表示

**参数：**
- `obj` (UiObject): UiObject 对象

**返回值：**
- string: 元素的字符串表示

**使用示例：**
```lua
local acc = uiacc.new()
local element = acc.findOnce()
print("元素信息: " .. element:toString())
```

---

## 综合使用示例

### 示例1：查找并点击按钮
```lua
-- 创建 Accessibility 对象
local acc = uiacc.new()

-- 根据文本查找按钮
local button = acc.text("登录")
if button ~= nil then
    button:objClick()
    print("点击登录按钮成功")
end

-- 释放资源
acc.release()
```

### 示例2：查找输入框并输入文本
```lua
-- 创建 Accessibility 对象
local acc = uiacc.new()

-- 查找密码输入框
local passwordInput = acc.password(true)
if passwordInput ~= nil then
    passwordInput:setText("123456")
    print("设置密码成功")
end

-- 释放资源
acc.release()
```

### 示例3：遍历所有子元素
```lua
-- 创建 Accessibility 对象
local acc = uiacc.new()

-- 查找所有元素
local elements = acc.find()
for i, element in ipairs(elements) do
    print("元素 " .. i .. ":")
    print("  文本: " .. element:getText())
    print("  类名: " .. element:getClassName())
    print("  包名: " .. element:getPackageName())
end

-- 释放资源
acc.release()
```

### 示例4：等待元素出现
```lua
-- 创建 Accessibility 对象
local acc = uiacc.new()

-- 等待元素出现（最多等待 5 秒）
local element = acc.waitFor(5000)
if element ~= nil then
    print("元素已出现: " .. element:getText())
    element:objClick()
else
    print("等待超时，元素未出现")
end

-- 释放资源
acc.release()
```

### 示例5：组合查询条件
```lua
-- 创建 Accessibility 对象
local acc = uiacc.new()

-- 组合多个查询条件
local element = acc.text("确定").clickable(true).enabled(true)
if element ~= nil then
    print("找到可点击的确定按钮")
    element:objClick()
end

-- 释放资源
acc.release()
```