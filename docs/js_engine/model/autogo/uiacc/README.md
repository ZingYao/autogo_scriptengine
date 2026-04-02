# uiacc 模块

## 模块简介

uiacc 模块提供了无障碍服务（Accessibility）的功能，用于查找和操作界面元素。

## 方法列表

### uiacc.new
创建一个新的Accessibility对象

**参数：**

| 参数名 | 类型 | 说明 |
|--------|------|------|
| displayId | number | 显示设备 ID（可选，默认为 0） |

**返回值：**

| 返回值 | 类型 | 说明 |
|--------|------|------|
| acc | object | Accessibility 对象 |

**使用示例：**
```javascript
// 创建 Accessibility 对象
var acc = uiacc.new();
var acc2 = uiacc.new(1);  // 指定显示设备 ID
```

---

## Accessibility 对象方法（查询方法）

### text
根据文本内容查找元素

**参数：**
- `value` (string): 要查找的文本内容

**返回值：**
- Accessibility 对象: 匹配的元素

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.text("确定");
```

---

### textContains
根据文本包含内容查找元素

**参数：**
- `value` (string): 要查找的文本内容（包含关系）

**返回值：**
- Accessibility 对象: 匹配的元素

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.textContains("登录");
```

---

### textStartsWith
根据文本前缀查找元素

**参数：**
- `value` (string): 要查找的文本前缀

**返回值：**
- Accessibility 对象: 匹配的元素

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.textStartsWith("用户");
```

---

### textEndsWith
根据文本后缀查找元素

**参数：**
- `value` (string): 要查找的文本后缀

**返回值：**
- Accessibility 对象: 匹配的元素

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.textEndsWith("按钮");
```

---

### textMatches
根据文本正则表达式查找元素

**参数：**
- `value` (string): 正则表达式

**返回值：**
- Accessibility 对象: 匹配的元素

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.textMatches(".*登录.*");
```

---

### desc
根据描述内容查找元素

**参数：**
- `value` (string): 要查找的描述内容

**返回值：**
- Accessibility 对象: 匹配的元素

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.desc("点击登录");
```

---

### descContains
根据描述包含内容查找元素

**参数：**
- `value` (string): 要查找的描述内容（包含关系）

**返回值：**
- Accessibility 对象: 匹配的元素

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.descContains("登录");
```

---

### descStartsWith
根据描述前缀查找元素

**参数：**
- `value` (string): 要查找的描述前缀

**返回值：**
- Accessibility 对象: 匹配的元素

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.descStartsWith("请");
```

---

### descEndsWith
根据描述后缀查找元素

**参数：**
- `value` (string): 要查找的描述后缀

**返回值：**
- Accessibility 对象: 匹配的元素

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.descEndsWith("按钮");
```

---

### descMatches
根据描述正则表达式查找元素

**参数：**
- `value` (string): 正则表达式

**返回值：**
- Accessibility 对象: 匹配的元素

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.descMatches(".*点击.*");
```

---

### id
根据 ID 查找元素

**参数：**
- `value` (string): 要查找的元素 ID

**返回值：**
- Accessibility 对象: 匹配的元素

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.id("login_button");
```

---

### idContains
根据 ID 包含内容查找元素

**参数：**
- `value` (string): 要查找的 ID 内容（包含关系）

**返回值：**
- Accessibility 对象: 匹配的元素

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.idContains("button");
```

---

### idStartsWith
根据 ID 前缀查找元素

**参数：**
- `value` (string): 要查找的 ID 前缀

**返回值：**
- Accessibility 对象: 匹配的元素

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.idStartsWith("btn_");
```

---

### idEndsWith
根据 ID 后缀查找元素

**参数：**
- `value` (string): 要查找的 ID 后缀

**返回值：**
- Accessibility 对象: 匹配的元素

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.idEndsWith("_submit");
```

---

### idMatches
根据 ID 正则表达式查找元素

**参数：**
- `value` (string): 正则表达式

**返回值：**
- Accessibility 对象: 匹配的元素

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.idMatches("btn_.*");
```

---

### className
根据类名查找元素

**参数：**
- `value` (string): 要查找的类名

**返回值：**
- Accessibility 对象: 匹配的元素

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.className("Button");
```

---

### classNameContains
根据类名包含内容查找元素

**参数：**
- `value` (string): 要查找的类名内容（包含关系）

**返回值：**
- Accessibility 对象: 匹配的元素

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.classNameContains("Button");
```

---

### classNameStartsWith
根据类名前缀查找元素

**参数：**
- `value` (string): 要查找的类名前缀

**返回值：**
- Accessibility 对象: 匹配的元素

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.classNameStartsWith("android.widget");
```

---

### classNameEndsWith
根据类名后缀查找元素

**参数：**
- `value` (string): 要查找的类名后缀

**返回值：**
- Accessibility 对象: 匹配的元素

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.classNameEndsWith("Button");
```

---

### classNameMatches
根据类名正则表达式查找元素

**参数：**
- `value` (string): 正则表达式

**返回值：**
- Accessibility 对象: 匹配的元素

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.classNameMatches(".*Button.*");
```

---

### packageName
根据包名查找元素

**参数：**
- `value` (string): 要查找的包名

**返回值：**
- Accessibility 对象: 匹配的元素

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.packageName("com.example.app");
```

---

### packageNameContains
根据包名包含内容查找元素

**参数：**
- `value` (string): 要查找的包名内容（包含关系）

**返回值：**
- Accessibility 对象: 匹配的元素

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.packageNameContains("example");
```

---

### packageNameStartsWith
根据包名前缀查找元素

**参数：**
- `value` (string): 要查找的包名前缀

**返回值：**
- Accessibility 对象: 匹配的元素

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.packageNameStartsWith("com.example");
```

---

### packageNameEndsWith
根据包名后缀查找元素

**参数：**
- `value` (string): 要查找的包名后缀

**返回值：**
- Accessibility 对象: 匹配的元素

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.packageNameEndsWith(".app");
```

---

### packageNameMatches
根据包名正则表达式查找元素

**参数：**
- `value` (string): 正则表达式

**返回值：**
- Accessibility 对象: 匹配的元素

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.packageNameMatches("com\\..*\\.app");
```

---

### bounds
根据边界区域查找元素

**参数：**
- `left` (int): 左边界
- `top` (int): 上边界
- `right` (int): 右边界
- `bottom` (int): 下边界

**返回值：**
- Accessibility 对象: 匹配的元素

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.bounds(100, 100, 200, 200);
```

---

### boundsInside
查找在指定边界区域内的元素

**参数：**
- `left` (int): 左边界
- `top` (int): 上边界
- `right` (int): 右边界
- `bottom` (int): 下边界

**返回值：**
- Accessibility 对象: 匹配的元素

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.boundsInside(100, 100, 200, 200);
```

---

### boundsContains
查找包含指定边界区域的元素

**参数：**
- `left` (int): 左边界
- `top` (int): 上边界
- `right` (int): 右边界
- `bottom` (int): 下边界

**返回值：**
- Accessibility 对象: 匹配的元素

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.boundsContains(100, 100, 200, 200);
```

---

### drawingOrder
根据绘制顺序查找元素

**参数：**
- `value` (int): 绘制顺序值

**返回值：**
- Accessibility 对象: 匹配的元素

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.drawingOrder(0);
```

---

### clickable
根据可点击属性查找元素

**参数：**
- `value` (bool): 是否可点击

**返回值：**
- Accessibility 对象: 匹配的元素

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.clickable(true);
```

---

### longClickable
根据可长按点击属性查找元素

**参数：**
- `value` (bool): 是否可长按点击

**返回值：**
- Accessibility 对象: 匹配的元素

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.longClickable(true);
```

---

### checkable
根据可勾选属性查找元素

**参数：**
- `value` (bool): 是否可勾选

**返回值：**
- Accessibility 对象: 匹配的元素

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.checkable(true);
```

---

### selected
根据选中状态查找元素

**参数：**
- `value` (bool): 是否选中

**返回值：**
- Accessibility 对象: 匹配的元素

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.selected(true);
```

---

### enabled
根据启用状态查找元素

**参数：**
- `value` (bool): 是否启用

**返回值：**
- Accessibility 对象: 匹配的元素

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.enabled(true);
```

---

### scrollable
根据可滚动属性查找元素

**参数：**
- `value` (bool): 是否可滚动

**返回值：**
- Accessibility 对象: 匹配的元素

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.scrollable(true);
```

---

### editable
根据可编辑属性查找元素

**参数：**
- `value` (bool): 是否可编辑

**返回值：**
- Accessibility 对象: 匹配的元素

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.editable(true);
```

---

### multiLine
根据多行属性查找元素

**参数：**
- `value` (bool): 是否多行

**返回值：**
- Accessibility 对象: 匹配的元素

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.multiLine(true);
```

---

### checked
根据勾选状态查找元素

**参数：**
- `value` (bool): 是否勾选

**返回值：**
- Accessibility 对象: 匹配的元素

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.checked(true);
```

---

### focusable
根据可聚焦属性查找元素

**参数：**
- `value` (bool): 是否可聚焦

**返回值：**
- Accessibility 对象: 匹配的元素

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.focusable(true);
```

---

### dismissable
根据可关闭属性查找元素

**参数：**
- `value` (bool): 是否可关闭

**返回值：**
- Accessibility 对象: 匹配的元素

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.dismissable(true);
```

---

### focused
根据聚焦状态查找元素

**参数：**
- `value` (bool): 是否聚焦

**返回值：**
- Accessibility 对象: 匹配的元素

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.focused(true);
```

---

### contextClickable
根据上下文可点击属性查找元素

**参数：**
- `value` (bool): 是否上下文可点击

**返回值：**
- Accessibility 对象: 匹配的元素

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.contextClickable(true);
```

---

### index
根据索引查找元素

**参数：**
- `value` (int): 元素索引

**返回值：**
- Accessibility 对象: 匹配的元素

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.index(0);
```

---

### visible
根据可见性查找元素

**参数：**
- `value` (bool): 是否可见

**返回值：**
- Accessibility 对象: 匹配的元素

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.visible(true);
```

---

### password
根据密码属性查找元素

**参数：**
- `value` (bool): 是否为密码输入框

**返回值：**
- Accessibility 对象: 匹配的元素

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.password(true);
```

---

### click
点击匹配的元素

**参数：**
- `text` (string): 点击后的文本内容（可选）

**返回值：**
- boolean: 点击成功返回 true，失败返回 false

**使用示例：**
```javascript
var acc = uiacc.new();
var success = acc.click("确定");
```

---

### waitFor
等待元素出现

**参数：**
- `timeout` (int): 超时时间（毫秒）

**返回值：**
- UiObject: 找到的元素对象，超时返回 null

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.waitFor(5000);
if (element != null) {
    console.log("找到元素");
}
```

---

### findOnce
查找第一个匹配的元素

**返回值：**
- UiObject: 找到的元素对象，未找到返回 null

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.findOnce();
if (element != null) {
    element.click();
}
```

---

### find
查找所有匹配的元素

**返回值：**
- Array: 元素对象数组

**使用示例：**
```javascript
var acc = uiacc.new();
var elements = acc.find();
for (var i = 0; i < elements.length; i++) {
    console.log("元素 " + i + ": " + elements[i].getText());
}
```

---

### release
释放 Accessibility 对象资源

**使用示例：**
```javascript
var acc = uiacc.new();
// 使用完毕后释放
acc.release();
```

---

## UiObject 对象方法（操作方法）

### click
点击元素

**返回值：**
- boolean: 点击成功返回 true，失败返回 false

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.findOnce();
element.click();
```

---

### clickCenter
点击元素中心

**返回值：**
- boolean: 点击成功返回 true，失败返回 false

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.findOnce();
element.clickCenter();
```

---

### clickLongClick
长按点击元素

**返回值：**
- boolean: 点击成功返回 true，失败返回 false

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.findOnce();
element.clickLongClick();
```

---

### copy
复制元素内容

**返回值：**
- boolean: 复制成功返回 true，失败返回 false

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.findOnce();
element.copy();
```

---

### cut
剪切元素内容

**返回值：**
- boolean: 剪切成功返回 true，失败返回 false

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.findOnce();
element.cut();
```

---

### paste
粘贴内容到元素

**返回值：**
- boolean: 粘贴成功返回 true，失败返回 false

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.findOnce();
element.paste();
```

---

### scrollForward
向前滚动

**返回值：**
- boolean: 滚动成功返回 true，失败返回 false

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.findOnce();
element.scrollForward();
```

---

### scrollBackward
向后滚动

**返回值：**
- boolean: 滚动成功返回 true，失败返回 false

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.findOnce();
element.scrollBackward();
```

---

### collapse
折叠元素

**返回值：**
- boolean: 折叠成功返回 true，失败返回 false

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.findOnce();
element.collapse();
```

---

### expand
展开元素

**返回值：**
- boolean: 展开成功返回 true，失败返回 false

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.findOnce();
element.expand();
```

---

### show
显示元素

**返回值：**
- boolean: 显示成功返回 true，失败返回 false

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.findOnce();
element.show();
```

---

### select
选中元素

**返回值：**
- boolean: 选中成功返回 true，失败返回 false

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.findOnce();
element.select();
```

---

### clearSelect
清除选中状态

**返回值：**
- boolean: 清除成功返回 true，失败返回 false

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.findOnce();
element.clearSelect();
```

---

### setSelection
设置选中范围

**参数：**
- `start` (int): 起始位置
- `end` (int): 结束位置

**返回值：**
- boolean: 设置成功返回 true，失败返回 false

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.findOnce();
element.setSelection(0, 10);
```

---

### setVisibleToUser
设置元素对用户可见性

**参数：**
- `isVisible` (bool): 是否可见

**返回值：**
- boolean: 设置成功返回 true，失败返回 false

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.findOnce();
element.setVisibleToUser(true);
```

---

### setText
设置元素文本

**参数：**
- `str` (string): 要设置的文本

**返回值：**
- boolean: 设置成功返回 true，失败返回 false

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.findOnce();
element.setText("新的文本");
```

---

### getClickable
获取可点击属性

**返回值：**
- boolean: 是否可点击

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.findOnce();
console.log("可点击: " + element.getClickable());
```

---

### getLongClickable
获取可长按点击属性

**返回值：**
- boolean: 是否可长按点击

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.findOnce();
console.log("可长按点击: " + element.getLongClickable());
```

---

### getCheckable
获取可勾选属性

**返回值：**
- boolean: 是否可勾选

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.findOnce();
console.log("可勾选: " + element.getCheckable());
```

---

### getSelected
获取选中状态

**返回值：**
- boolean: 是否选中

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.findOnce();
console.log("选中状态: " + element.getSelected());
```

---

### getEnabled
获取启用状态

**返回值：**
- boolean: 是否启用

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.findOnce();
console.log("启用状态: " + element.getEnabled());
```

---

### getScrollable
获取可滚动属性

**返回值：**
- boolean: 是否可滚动

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.findOnce();
console.log("可滚动: " + element.getScrollable());
```

---

### getEditable
获取可编辑属性

**返回值：**
- boolean: 是否可编辑

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.findOnce();
console.log("可编辑: " + element.getEditable());
```

---

### getMultiLine
获取多行属性

**返回值：**
- boolean: 是否多行

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.findOnce();
console.log("多行: " + element.getMultiLine());
```

---

### getChecked
获取勾选状态

**返回值：**
- boolean: 是否勾选

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.findOnce();
console.log("勾选状态: " + element.getChecked());
```

---

### getFocused
获取聚焦状态

**返回值：**
- boolean: 是否聚焦

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.findOnce();
console.log("聚焦状态: " + element.getFocused());
```

---

### getFocusable
获取可聚焦属性

**返回值：**
- boolean: 是否可聚焦

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.findOnce();
console.log("可聚焦: " + element.getFocusable());
```

---

### getDismissable
获取可关闭属性

**返回值：**
- boolean: 是否可关闭

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.findOnce();
console.log("可关闭: " + element.getDismissable());
```

---

### getContextClickable
获取上下文可点击属性

**返回值：**
- boolean: 是否上下文可点击

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.findOnce();
console.log("上下文可点击: " + element.getContextClickable());
```

---

### getAccessibilityFocused
获取无障碍聚焦状态

**返回值：**
- boolean: 是否无障碍聚焦

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.findOnce();
console.log("无障碍聚焦: " + element.getAccessibilityFocused());
```

---

### getChildCount
获取子元素数量

**返回值：**
- int: 子元素数量

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.findOnce();
console.log("子元素数量: " + element.getChildCount());
```

---

### getDrawingOrder
获取绘制顺序

**返回值：**
- int: 绘制顺序

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.findOnce();
console.log("绘制顺序: " + element.getDrawingOrder());
```

---

### getIndex
获取元素索引

**返回值：**
- int: 元素索引

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.findOnce();
console.log("元素索引: " + element.getIndex());
```

---

### getBounds
获取元素边界

**返回值：**
- Object: 包含 left, top, right, bottom 属性的对象

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.findOnce();
var bounds = element.getBounds();
console.log("边界: " + bounds.left + ", " + bounds.top + ", " + bounds.right + ", " + bounds.bottom);
```

---

### getBoundsInParent
获取元素在父容器中的边界

**返回值：**
- Object: 包含 left, top, right, bottom 属性的对象

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.findOnce();
var bounds = element.getBoundsInParent();
console.log("父容器边界: " + bounds.left + ", " + bounds.top + ", " + bounds.right + ", " + bounds.bottom);
```

---

### getId
获取元素 ID

**返回值：**
- string: 元素 ID

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.findOnce();
console.log("元素 ID: " + element.getId());
```

---

### getText
获取元素文本

**返回值：**
- string: 元素文本

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.findOnce();
console.log("元素文本: " + element.getText());
```

---

### getDesc
获取元素描述

**返回值：**
- string: 元素描述

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.findOnce();
console.log("元素描述: " + element.getDesc());
```

---

### getPackageName
获取元素包名

**返回值：**
- string: 元素包名

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.findOnce();
console.log("元素包名: " + element.getPackageName());
```

---

### getClassName
获取元素类名

**返回值：**
- string: 元素类名

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.findOnce();
console.log("元素类名: " + element.getClassName());
```

---

### getParent
获取父元素

**返回值：**
- UiObject: 父元素对象，无父元素返回 null

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.findOnce();
var parent = element.getParent();
if (parent != null) {
    console.log("找到父元素");
}
```

---

### getChild
获取指定索引的子元素

**参数：**
- `index` (int): 子元素索引

**返回值：**
- UiObject: 子元素对象，无子元素返回 null

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.findOnce();
var child = element.getChild(0);
if (child != null) {
    console.log("找到第一个子元素");
}
```

---

### getChildren
获取所有子元素

**返回值：**
- Array: 子元素对象数组

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.findOnce();
var children = element.getChildren();
for (var i = 0; i < children.length; i++) {
    console.log("子元素 " + i + ": " + children[i].getText());
}
```

---

### getVisible
获取可见性

**返回值：**
- boolean: 是否可见

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.findOnce();
console.log("可见性: " + element.getVisible());
```

---

### getPassword
获取密码属性

**返回值：**
- boolean: 是否为密码输入框

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.findOnce();
console.log("密码属性: " + element.getPassword());
```

---

### toString
获取元素字符串表示

**返回值：**
- string: 元素的字符串表示

**使用示例：**
```javascript
var acc = uiacc.new();
var element = acc.findOnce();
console.log("元素信息: " + element.toString());
```

---

## 综合使用示例

### 示例1：查找并点击按钮
```javascript
// 创建 Accessibility 对象
var acc = uiacc.new();

// 根据文本查找按钮
var button = acc.text("登录");
if (button != null) {
    button.click();
    console.log("点击登录按钮成功");
}

// 释放资源
acc.release();
```

### 示例2：查找输入框并输入文本
```javascript
// 创建 Accessibility 对象
var acc = uiacc.new();

// 查找密码输入框
var passwordInput = acc.password(true);
if (passwordInput != null) {
    passwordInput.setText("123456");
    console.log("设置密码成功");
}

// 释放资源
acc.release();
```

### 示例3：遍历所有子元素
```javascript
// 创建 Accessibility 对象
var acc = uiacc.new();

// 查找所有元素
var elements = acc.find();
for (var i = 0; i < elements.length; i++) {
    console.log("元素 " + i + ":");
    console.log("  文本: " + elements[i].getText());
    console.log("  类名: " + elements[i].getClassName());
    console.log("  包名: " + elements[i].getPackageName());
}

// 释放资源
acc.release();
```

### 示例4：等待元素出现
```javascript
// 创建 Accessibility 对象
var acc = uiacc.new();

// 等待元素出现（最多等待 5 秒）
var element = acc.waitFor(5000);
if (element != null) {
    console.log("元素已出现: " + element.getText());
    element.click();
} else {
    console.log("等待超时，元素未出现");
}

// 释放资源
acc.release();
```

### 示例5：组合查询条件
```javascript
// 创建 Accessibility 对象
var acc = uiacc.new();

// 组合多个查询条件
var element = acc.text("确定").clickable(true).enabled(true);
if (element != null) {
    console.log("找到可点击的确定按钮");
    element.click();
}

// 释放资源
acc.release();
```