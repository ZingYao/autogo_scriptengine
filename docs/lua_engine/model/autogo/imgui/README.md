# imgui 模块

## 模块简介

imgui 模块提供了 ImGui 图形用户界面库的绑定，用于创建复杂的交互式 GUI 界面。

## 方法列表

### imgui.init
初始化ImGui

**使用示例：**
```lua
-- 调用 imgui.init 方法
imgui.init();
```

---

### imgui.close
关闭ImGui

**使用示例：**
```lua
-- 调用 imgui.close 方法
imgui.close();
```

---

### imgui.run
运行ImGui主循环

**使用示例：**
```lua
-- 调用 imgui.run 方法
imgui.run();
```

---

### imgui.vertexBufferLayout
获取顶点缓冲区布局

**使用示例：**
```lua
-- 调用 imgui.vertexBufferLayout 方法
imgui.vertexBufferLayout();
```

---

### imgui.indexBufferLayout
获取索引缓冲区布局

**使用示例：**
```lua
-- 调用 imgui.indexBufferLayout 方法
imgui.indexBufferLayout();
```

---

### imgui.newGlyphRange
创建字形范围

**使用示例：**
```lua
-- 调用 imgui.newGlyphRange 方法
imgui.newGlyphRange();
```

---

### imgui.newContext
创建ImGui上下文

**使用示例：**
```lua
-- 调用 imgui.newContext 方法
imgui.newContext();
```

---

### imgui.newIO
创建IO对象

**使用示例：**
```lua
-- 调用 imgui.newIO 方法
imgui.newIO();
```

---

### imgui.newStyle
创建样式对象

**使用示例：**
```lua
-- 调用 imgui.newStyle 方法
imgui.newStyle();
```

---

### imgui.newDrawList
创建绘制列表

**使用示例：**
```lua
-- 调用 imgui.newDrawList 方法
imgui.newDrawList();
```

---

### imgui.newFont
创建字体对象

**使用示例：**
```lua
-- 调用 imgui.newFont 方法
imgui.newFont();
```

---

### imgui.newFontAtlas
创建字体图集

**使用示例：**
```lua
-- 调用 imgui.newFontAtlas 方法
imgui.newFontAtlas();
```

---

### imgui.newFontConfig
创建字体配置

**使用示例：**
```lua
-- 调用 imgui.newFontConfig 方法
imgui.newFontConfig();
```

---

### imgui.newDrawData
创建绘制数据

**使用示例：**
```lua
-- 调用 imgui.newDrawData 方法
imgui.newDrawData();
```

---

### imgui.newStorage
创建存储对象

**使用示例：**
```lua
-- 调用 imgui.newStorage 方法
imgui.newStorage();
```

---

### imgui.newPlatformIO
创建平台IO对象

**使用示例：**
```lua
-- 调用 imgui.newPlatformIO 方法
imgui.newPlatformIO();
```

---

### imgui.createContext
创建ImGui上下文

**使用示例：**
```lua
-- 调用 imgui.createContext 方法
imgui.createContext();
```

---

### imgui.destroyContext
销毁ImGui上下文

**使用示例：**
```lua
-- 调用 imgui.destroyContext 方法
imgui.destroyContext();
```

---

### imgui.setCurrentContext
设置当前ImGui上下文

**使用示例：**
```lua
-- 调用 imgui.setCurrentContext 方法
imgui.setCurrentContext();
```

---

### imgui.newFrame
开始新帧

**使用示例：**
```lua
-- 调用 imgui.newFrame 方法
imgui.newFrame();
```

---

### imgui.render
渲染ImGui

**使用示例：**
```lua
-- 调用 imgui.render 方法
imgui.render();
```

---

### imgui.endFrame
结束帧

**使用示例：**
```lua
-- 调用 imgui.endFrame 方法
imgui.endFrame();
```

---

### imgui.begin
开始窗口

**使用示例：**
```lua
-- 调用 imgui.begin 方法
imgui.begin();
```

---

### imgui.end
结束窗口

**使用示例：**
```lua
-- 调用 imgui.end 方法
imgui.end();
```

---

### imgui.button
创建按钮

**使用示例：**
```lua
-- 调用 imgui.button 方法
imgui.button();
```

---

### imgui.text
显示文本

**使用示例：**
```lua
-- 调用 imgui.text 方法
imgui.text();
```

---

### imgui.inputText
输入文本框

**使用示例：**
```lua
-- 调用 imgui.inputText 方法
imgui.inputText();
```

---

### imgui.spacing
添加间距

**使用示例：**
```lua
-- 调用 imgui.spacing 方法
imgui.spacing();
```

---

### imgui.sameLine
同下一行

**使用示例：**
```lua
-- 调用 imgui.sameLine 方法
imgui.sameLine();
```

---

### imgui.setNextWindowPos
设置下一个窗口位置

**使用示例：**
```lua
-- 调用 imgui.setNextWindowPos 方法
imgui.setNextWindowPos();
```

---

### imgui.setNextWindowSize
设置下一个窗口大小

**使用示例：**
```lua
-- 调用 imgui.setNextWindowSize 方法
imgui.setNextWindowSize();
```

---

### imgui.colorConvertFloat4ToU32
颜色转换

**使用示例：**
```lua
-- 调用 imgui.colorConvertFloat4ToU32 方法
imgui.colorConvertFloat4ToU32();
```

---

### imgui.colorConvertU32ToFloat4
颜色转换

**使用示例：**
```lua
-- 调用 imgui.colorConvertU32ToFloat4 方法
imgui.colorConvertU32ToFloat4();
```

---

### imgui.setCursorPos
设置光标位置

**使用示例：**
```lua
-- 调用 imgui.setCursorPos 方法
imgui.setCursorPos();
```

---

### imgui.setWindowFontScale
设置窗口字体缩放

**使用示例：**
```lua
-- 调用 imgui.setWindowFontScale 方法
imgui.setWindowFontScale();
```

---

### imgui.newIO
创建IO对象

**使用示例：**
```lua
-- 调用 imgui.newIO 方法
imgui.newIO();
```

---

### imgui.newPlatformIO
创建平台IO对象

**使用示例：**
```lua
-- 调用 imgui.newPlatformIO 方法
imgui.newPlatformIO();
```

---

### imgui.checkbox
复选框

**使用示例：**
```lua
-- 调用 imgui.checkbox 方法
imgui.checkbox();
```

---

### imgui.sliderFloat
浮点滑块

**使用示例：**
```lua
-- 调用 imgui.sliderFloat 方法
imgui.sliderFloat();
```

---

### imgui.sliderInt
整数滑块

**使用示例：**
```lua
-- 调用 imgui.sliderInt 方法
imgui.sliderInt();
```

---

### imgui.separator
分隔符

**使用示例：**
```lua
-- 调用 imgui.separator 方法
imgui.separator();
```

---

### imgui.separatorText
分隔文本

**使用示例：**
```lua
-- 调用 imgui.separatorText 方法
imgui.separatorText();
```

---

### imgui.beginMenu
开始菜单

**使用示例：**
```lua
-- 调用 imgui.beginMenu 方法
imgui.beginMenu();
```

---

### imgui.endMenu
结束菜单

**使用示例：**
```lua
-- 调用 imgui.endMenu 方法
imgui.endMenu();
```

---

### imgui.beginMenuBar
开始菜单栏

**使用示例：**
```lua
-- 调用 imgui.beginMenuBar 方法
imgui.beginMenuBar();
```

---

### imgui.endMenuBar
结束菜单栏

**使用示例：**
```lua
-- 调用 imgui.endMenuBar 方法
imgui.endMenuBar();
```

---

### imgui.beginPopup
开始弹出窗口

**使用示例：**
```lua
-- 调用 imgui.beginPopup 方法
imgui.beginPopup();
```

---

### imgui.endPopup
结束弹出窗口

**使用示例：**
```lua
-- 调用 imgui.endPopup 方法
imgui.endPopup();
```

---

### imgui.colorEdit3
颜色编辑器(3通道)

**使用示例：**
```lua
-- 调用 imgui.colorEdit3 方法
imgui.colorEdit3();
```

---

### imgui.colorEdit4
颜色编辑器(4通道)

**使用示例：**
```lua
-- 调用 imgui.colorEdit4 方法
imgui.colorEdit4();
```

---

### imgui.progressBar
进度条

**使用示例：**
```lua
-- 调用 imgui.progressBar 方法
imgui.progressBar();
```

---

### imgui.bullet
项目符号

**使用示例：**
```lua
-- 调用 imgui.bullet 方法
imgui.bullet();
```

---

### imgui.bulletText
带项目符号的文本

**使用示例：**
```lua
-- 调用 imgui.bulletText 方法
imgui.bulletText();
```

---

### imgui.smallButton
小按钮

**使用示例：**
```lua
-- 调用 imgui.smallButton 方法
imgui.smallButton();
```

---

### imgui.arrowButton
箭头按钮

**使用示例：**
```lua
-- 调用 imgui.arrowButton 方法
imgui.arrowButton();
```

---

### imgui.beginTooltip
开始工具提示

**使用示例：**
```lua
-- 调用 imgui.beginTooltip 方法
imgui.beginTooltip();
```

---

### imgui.endTooltip
结束工具提示

**使用示例：**
```lua
-- 调用 imgui.endTooltip 方法
imgui.endTooltip();
```

---

### imgui.setTooltip
设置工具提示

**使用示例：**
```lua
-- 调用 imgui.setTooltip 方法
imgui.setTooltip();
```

---

### imgui.beginGroup
开始组

**使用示例：**
```lua
-- 调用 imgui.beginGroup 方法
imgui.beginGroup();
```

---

### imgui.endGroup
结束组

**使用示例：**
```lua
-- 调用 imgui.endGroup 方法
imgui.endGroup();
```

---

### imgui.popID
弹出ID

**使用示例：**
```lua
-- 调用 imgui.popID 方法
imgui.popID();
```

---

### imgui.popStyleVar
弹出样式变量

**使用示例：**
```lua
-- 调用 imgui.popStyleVar 方法
imgui.popStyleVar();
```

---

### imgui.popStyleColor
弹出样式颜色

**使用示例：**
```lua
-- 调用 imgui.popStyleColor 方法
imgui.popStyleColor();
```

---

### imgui.alert
显示对话框

**使用示例：**
```lua
-- 调用 imgui.alert 方法
imgui.alert();
```

---

### imgui.toast
显示Toast提示

**使用示例：**
```lua
-- 调用 imgui.toast 方法
imgui.toast();
```

---

### imgui.drawRect
绘制矩形

**使用示例：**
```lua
-- 调用 imgui.drawRect 方法
imgui.drawRect();
```

---

### imgui.beginChild
开始子窗口

**使用示例：**
```lua
-- 调用 imgui.beginChild 方法
imgui.beginChild();
```

---

### imgui.endChild
结束子窗口

**使用示例：**
```lua
-- 调用 imgui.endChild 方法
imgui.endChild();
```

---

### imgui.setNextWindowSizeConstraints
设置下一个窗口大小约束

**使用示例：**
```lua
-- 调用 imgui.setNextWindowSizeConstraints 方法
imgui.setNextWindowSizeConstraints();
```

---

### imgui.setNextWindowContentSize
设置下一个窗口内容大小

**使用示例：**
```lua
-- 调用 imgui.setNextWindowContentSize 方法
imgui.setNextWindowContentSize();
```

---

### imgui.setNextWindowCollapsed
设置下一个窗口折叠状态

**使用示例：**
```lua
-- 调用 imgui.setNextWindowCollapsed 方法
imgui.setNextWindowCollapsed();
```

---

### imgui.setNextWindowFocus
设置下一个窗口焦点

**使用示例：**
```lua
-- 调用 imgui.setNextWindowFocus 方法
imgui.setNextWindowFocus();
```

---

### imgui.setNextWindowBgAlpha
设置下一个窗口背景透明度

**使用示例：**
```lua
-- 调用 imgui.setNextWindowBgAlpha 方法
imgui.setNextWindowBgAlpha();
```

---

### imgui.setWindowPos
设置窗口位置

**使用示例：**
```lua
-- 调用 imgui.setWindowPos 方法
imgui.setWindowPos();
```

---

### imgui.setWindowSize
设置窗口大小

**使用示例：**
```lua
-- 调用 imgui.setWindowSize 方法
imgui.setWindowSize();
```

---

### imgui.setWindowCollapsed
设置窗口折叠状态

**使用示例：**
```lua
-- 调用 imgui.setWindowCollapsed 方法
imgui.setWindowCollapsed();
```

---

### imgui.setWindowFocus
设置窗口焦点

**使用示例：**
```lua
-- 调用 imgui.setWindowFocus 方法
imgui.setWindowFocus();
```

---

### imgui.pushFont
推入字体

**使用示例：**
```lua
-- 调用 imgui.pushFont 方法
imgui.pushFont();
```

---

### imgui.popFont
弹出字体

**使用示例：**
```lua
-- 调用 imgui.popFont 方法
imgui.popFont();
```

---

### imgui.pushStyleColor
推入样式颜色

**使用示例：**
```lua
-- 调用 imgui.pushStyleColor 方法
imgui.pushStyleColor();
```

---

### imgui.pushStyleVar
推入样式变量

**使用示例：**
```lua
-- 调用 imgui.pushStyleVar 方法
imgui.pushStyleVar();
```

---

### imgui.pushItemWidth
推入项宽度

**使用示例：**
```lua
-- 调用 imgui.pushItemWidth 方法
imgui.pushItemWidth();
```

---

### imgui.popItemWidth
弹出项宽度

**使用示例：**
```lua
-- 调用 imgui.popItemWidth 方法
imgui.popItemWidth();
```

---

### imgui.pushTextWrapPos
推入文本换行位置

**使用示例：**
```lua
-- 调用 imgui.pushTextWrapPos 方法
imgui.pushTextWrapPos();
```

---

### imgui.popTextWrapPos
弹出文本换行位置

**使用示例：**
```lua
-- 调用 imgui.popTextWrapPos 方法
imgui.popTextWrapPos();
```

---

### imgui.pushID
推入ID

**使用示例：**
```lua
-- 调用 imgui.pushID 方法
imgui.pushID();
```

---

### imgui.newLine
新行

**使用示例：**
```lua
-- 调用 imgui.newLine 方法
imgui.newLine();
```

---

### imgui.dummy
虚拟占位符

**使用示例：**
```lua
-- 调用 imgui.dummy 方法
imgui.dummy();
```

---

### imgui.indent
缩进

**使用示例：**
```lua
-- 调用 imgui.indent 方法
imgui.indent();
```

---

### imgui.unindent
取消缩进

**使用示例：**
```lua
-- 调用 imgui.unindent 方法
imgui.unindent();
```

---

### imgui.setCursorPos
设置光标位置

**使用示例：**
```lua
-- 调用 imgui.setCursorPos 方法
imgui.setCursorPos();
```

---

### imgui.setCursorPosX
设置光标X位置

**使用示例：**
```lua
-- 调用 imgui.setCursorPosX 方法
imgui.setCursorPosX();
```

---

### imgui.setCursorPosY
设置光标Y位置

**使用示例：**
```lua
-- 调用 imgui.setCursorPosY 方法
imgui.setCursorPosY();
```

---

### imgui.setCursorScreenPos
设置光标屏幕位置

**使用示例：**
```lua
-- 调用 imgui.setCursorScreenPos 方法
imgui.setCursorScreenPos();
```

---

### imgui.alignTextToFramePadding
对齐文本到帧内边距

**使用示例：**
```lua
-- 调用 imgui.alignTextToFramePadding 方法
imgui.alignTextToFramePadding();
```

---

### imgui.textColored
彩色文本

**使用示例：**
```lua
-- 调用 imgui.textColored 方法
imgui.textColored();
```

---

### imgui.textDisabled
禁用文本

**使用示例：**
```lua
-- 调用 imgui.textDisabled 方法
imgui.textDisabled();
```

---

### imgui.textWrapped
换行文本

**使用示例：**
```lua
-- 调用 imgui.textWrapped 方法
imgui.textWrapped();
```

---

### imgui.labelText
标签文本

**使用示例：**
```lua
-- 调用 imgui.labelText 方法
imgui.labelText();
```

---

### imgui.calcTextSize
计算文本大小

**使用示例：**
```lua
-- 调用 imgui.calcTextSize 方法
imgui.calcTextSize();
```

---

### imgui.calcItemWidth
计算项宽度

**使用示例：**
```lua
-- 调用 imgui.calcItemWidth 方法
imgui.calcItemWidth();
```

---

### imgui.setScrollHereX
设置滚动到此处X

**使用示例：**
```lua
-- 调用 imgui.setScrollHereX 方法
imgui.setScrollHereX();
```

---

### imgui.setScrollHereY
设置滚动到此处Y

**使用示例：**
```lua
-- 调用 imgui.setScrollHereY 方法
imgui.setScrollHereY();
```

---

### imgui.setScrollFromPosX
从位置X滚动

**使用示例：**
```lua
-- 调用 imgui.setScrollFromPosX 方法
imgui.setScrollFromPosX();
```

---

### imgui.setScrollFromPosY
从位置Y滚动

**使用示例：**
```lua
-- 调用 imgui.setScrollFromPosY 方法
imgui.setScrollFromPosY();
```

---

### imgui.setScrollX
设置滚动X

**使用示例：**
```lua
-- 调用 imgui.setScrollX 方法
imgui.setScrollX();
```

---

### imgui.setScrollY
设置滚动Y

**使用示例：**
```lua
-- 调用 imgui.setScrollY 方法
imgui.setScrollY();
```

---

### imgui.checkbox
复选框

**使用示例：**
```lua
-- 调用 imgui.checkbox 方法
imgui.checkbox();
```

---

### imgui.checkboxFlags
复选框标志

**使用示例：**
```lua
-- 调用 imgui.checkboxFlags 方法
imgui.checkboxFlags();
```

---

### imgui.radioButton
单选按钮

**使用示例：**
```lua
-- 调用 imgui.radioButton 方法
imgui.radioButton();
```

---

### imgui.inputFloat
浮点输入

**使用示例：**
```lua
-- 调用 imgui.inputFloat 方法
imgui.inputFloat();
```

---

### imgui.inputInt
整数输入

**使用示例：**
```lua
-- 调用 imgui.inputInt 方法
imgui.inputInt();
```

---

### imgui.inputDouble
双精度浮点输入

**使用示例：**
```lua
-- 调用 imgui.inputDouble 方法
imgui.inputDouble();
```

---

### imgui.combo
下拉框

**使用示例：**
```lua
-- 调用 imgui.combo 方法
imgui.combo();
```

---

### imgui.dragFloat
浮点拖动

**使用示例：**
```lua
-- 调用 imgui.dragFloat 方法
imgui.dragFloat();
```

---

### imgui.dragInt
整数拖动

**使用示例：**
```lua
-- 调用 imgui.dragInt 方法
imgui.dragInt();
```

---

### imgui.sliderFloat
浮点滑块

**使用示例：**
```lua
-- 调用 imgui.sliderFloat 方法
imgui.sliderFloat();
```

---

### imgui.sliderInt
整数滑块

**使用示例：**
```lua
-- 调用 imgui.sliderInt 方法
imgui.sliderInt();
```

---

### imgui.sliderAngle
角度滑块

**使用示例：**
```lua
-- 调用 imgui.sliderAngle 方法
imgui.sliderAngle();
```

---

### imgui.colorEdit3
颜色编辑3

**使用示例：**
```lua
-- 调用 imgui.colorEdit3 方法
imgui.colorEdit3();
```

---

### imgui.colorEdit4
颜色编辑4

**使用示例：**
```lua
-- 调用 imgui.colorEdit4 方法
imgui.colorEdit4();
```

---

### imgui.colorPicker3
颜色选择器3

**使用示例：**
```lua
-- 调用 imgui.colorPicker3 方法
imgui.colorPicker3();
```

---

### imgui.colorPicker4
颜色选择器4

**使用示例：**
```lua
-- 调用 imgui.colorPicker4 方法
imgui.colorPicker4();
```

---

### imgui.colorButton
颜色按钮

**使用示例：**
```lua
-- 调用 imgui.colorButton 方法
imgui.colorButton();
```

---

### imgui.colorConvertHSVtoRGB
HSV转RGB

**使用示例：**
```lua
-- 调用 imgui.colorConvertHSVtoRGB 方法
imgui.colorConvertHSVtoRGB();
```

---

### imgui.colorConvertRGBtoHSV
RGB转HSV

**使用示例：**
```lua
-- 调用 imgui.colorConvertRGBtoHSV 方法
imgui.colorConvertRGBtoHSV();
```

---

### imgui.treeNode
树节点

**使用示例：**
```lua
-- 调用 imgui.treeNode 方法
imgui.treeNode();
```

---

### imgui.treePop
弹出树节点

**使用示例：**
```lua
-- 调用 imgui.treePop 方法
imgui.treePop();
```

---

### imgui.treePush
推入树节点

**使用示例：**
```lua
-- 调用 imgui.treePush 方法
imgui.treePush();
```

---

### imgui.collapsingHeader
折叠头

**使用示例：**
```lua
-- 调用 imgui.collapsingHeader 方法
imgui.collapsingHeader();
```

---

### imgui.selectable
可选项

**使用示例：**
```lua
-- 调用 imgui.selectable 方法
imgui.selectable();
```

---

### imgui.beginTable
开始表格

**使用示例：**
```lua
-- 调用 imgui.beginTable 方法
imgui.beginTable();
```

---

### imgui.endTable
结束表格

**使用示例：**
```lua
-- 调用 imgui.endTable 方法
imgui.endTable();
```

---

### imgui.tableNextRow
表格下一行

**使用示例：**
```lua
-- 调用 imgui.tableNextRow 方法
imgui.tableNextRow();
```

---

### imgui.tableNextColumn
表格下一列

**使用示例：**
```lua
-- 调用 imgui.tableNextColumn 方法
imgui.tableNextColumn();
```

---

### imgui.tableSetColumnIndex
设置表格列索引

**使用示例：**
```lua
-- 调用 imgui.tableSetColumnIndex 方法
imgui.tableSetColumnIndex();
```

---

### imgui.tableSetupColumn
设置表格列

**使用示例：**
```lua
-- 调用 imgui.tableSetupColumn 方法
imgui.tableSetupColumn();
```

---

### imgui.tableHeadersRow
表格标题行

**使用示例：**
```lua
-- 调用 imgui.tableHeadersRow 方法
imgui.tableHeadersRow();
```

---

### imgui.tableHeader
表格标题

**使用示例：**
```lua
-- 调用 imgui.tableHeader 方法
imgui.tableHeader();
```

---

### imgui.tableGetColumnCount
获取表格列数

**使用示例：**
```lua
-- 调用 imgui.tableGetColumnCount 方法
imgui.tableGetColumnCount();
```

---

### imgui.tableGetColumnIndex
获取表格列索引

**使用示例：**
```lua
-- 调用 imgui.tableGetColumnIndex 方法
imgui.tableGetColumnIndex();
```

---

### imgui.tableGetRowIndex
获取表格行索引

**使用示例：**
```lua
-- 调用 imgui.tableGetRowIndex 方法
imgui.tableGetRowIndex();
```

---

### imgui.beginMenu
开始菜单

**使用示例：**
```lua
-- 调用 imgui.beginMenu 方法
imgui.beginMenu();
```

---

### imgui.endMenu
结束菜单

**使用示例：**
```lua
-- 调用 imgui.endMenu 方法
imgui.endMenu();
```

---

### imgui.beginMenuBar
开始菜单栏

**使用示例：**
```lua
-- 调用 imgui.beginMenuBar 方法
imgui.beginMenuBar();
```

---

### imgui.endMenuBar
结束菜单栏

**使用示例：**
```lua
-- 调用 imgui.endMenuBar 方法
imgui.endMenuBar();
```

---

### imgui.menuItem
菜单项

**使用示例：**
```lua
-- 调用 imgui.menuItem 方法
imgui.menuItem();
```

---

### imgui.beginTabBar
开始标签栏

**使用示例：**
```lua
-- 调用 imgui.beginTabBar 方法
imgui.beginTabBar();
```

---

### imgui.endTabBar
结束标签栏

**使用示例：**
```lua
-- 调用 imgui.endTabBar 方法
imgui.endTabBar();
```

---

### imgui.beginTabItem
开始标签项

**使用示例：**
```lua
-- 调用 imgui.beginTabItem 方法
imgui.beginTabItem();
```

---

### imgui.endTabItem
结束标签项

**使用示例：**
```lua
-- 调用 imgui.endTabItem 方法
imgui.endTabItem();
```

---

### imgui.tabItemButton
标签项按钮

**使用示例：**
```lua
-- 调用 imgui.tabItemButton 方法
imgui.tabItemButton();
```

---

### imgui.beginDragDropSource
开始拖放源

**使用示例：**
```lua
-- 调用 imgui.beginDragDropSource 方法
imgui.beginDragDropSource();
```

---

### imgui.endDragDropSource
结束拖放源

**使用示例：**
```lua
-- 调用 imgui.endDragDropSource 方法
imgui.endDragDropSource();
```

---

### imgui.beginDragDropTarget
开始拖放目标

**使用示例：**
```lua
-- 调用 imgui.beginDragDropTarget 方法
imgui.beginDragDropTarget();
```

---

### imgui.endDragDropTarget
结束拖放目标

**使用示例：**
```lua
-- 调用 imgui.endDragDropTarget 方法
imgui.endDragDropTarget();
```

---

### imgui.setDragDropPayload
设置拖放数据

**使用示例：**
```lua
-- 调用 imgui.setDragDropPayload 方法
imgui.setDragDropPayload();
```

---

### imgui.beginDisabled
开始禁用

**使用示例：**
```lua
-- 调用 imgui.beginDisabled 方法
imgui.beginDisabled();
```

---

### imgui.endDisabled
结束禁用

**使用示例：**
```lua
-- 调用 imgui.endDisabled 方法
imgui.endDisabled();
```

---

### imgui.showDemoWindow
显示演示窗口

**使用示例：**
```lua
-- 调用 imgui.showDemoWindow 方法
imgui.showDemoWindow();
```

---

### imgui.showMetricsWindow
显示指标窗口

**使用示例：**
```lua
-- 调用 imgui.showMetricsWindow 方法
imgui.showMetricsWindow();
```

---

### imgui.showAboutWindow
显示关于窗口

**使用示例：**
```lua
-- 调用 imgui.showAboutWindow 方法
imgui.showAboutWindow();
```

---

### imgui.showStyleEditor
显示样式编辑器

**使用示例：**
```lua
-- 调用 imgui.showStyleEditor 方法
imgui.showStyleEditor();
```

---

### imgui.showStyleSelector
显示样式选择器

**使用示例：**
```lua
-- 调用 imgui.showStyleSelector 方法
imgui.showStyleSelector();
```

---

### imgui.showFontSelector
显示字体选择器

**使用示例：**
```lua
-- 调用 imgui.showFontSelector 方法
imgui.showFontSelector();
```

---

### imgui.showUserGuide
显示用户指南

**使用示例：**
```lua
-- 调用 imgui.showUserGuide 方法
imgui.showUserGuide();
```

---

### imgui.showDebugLogWindow
显示调试日志窗口

**使用示例：**
```lua
-- 调用 imgui.showDebugLogWindow 方法
imgui.showDebugLogWindow();
```

---

### imgui.showIDStackToolWindow
显示 ID 栈工具窗口

**使用示例：**
```lua
-- 调用 imgui.showIDStackToolWindow 方法
imgui.showIDStackToolWindow();
```

---

### imgui.beginCombo
开始组合框

**使用示例：**
```lua
-- 调用 imgui.beginCombo 方法
imgui.beginCombo();
```

---

### imgui.endCombo
结束组合框

**使用示例：**
```lua
-- 调用 imgui.endCombo 方法
imgui.endCombo();
```

---

### imgui.beginPopupModal
开始模态弹出窗口

**使用示例：**
```lua
-- 调用 imgui.beginPopupModal 方法
imgui.beginPopupModal();
```

---

### imgui.openPopup
打开弹出窗口

**使用示例：**
```lua
-- 调用 imgui.openPopup 方法
imgui.openPopup();
```

---

### imgui.closeCurrentPopup
关闭当前弹出窗口

**使用示例：**
```lua
-- 调用 imgui.closeCurrentPopup 方法
imgui.closeCurrentPopup();
```

---

### imgui.beginPopupContextItem
开始项上下文弹出窗口

**使用示例：**
```lua
-- 调用 imgui.beginPopupContextItem 方法
imgui.beginPopupContextItem();
```

---

### imgui.beginPopupContextWindow
开始窗口上下文弹出窗口

**使用示例：**
```lua
-- 调用 imgui.beginPopupContextWindow 方法
imgui.beginPopupContextWindow();
```

---

### imgui.beginPopupContextVoid
开始空白上下文弹出窗口

**使用示例：**
```lua
-- 调用 imgui.beginPopupContextVoid 方法
imgui.beginPopupContextVoid();
```

---

### imgui.isItemHovered
项是否悬停

**使用示例：**
```lua
-- 调用 imgui.isItemHovered 方法
imgui.isItemHovered();
```

---

### imgui.isItemActive
项是否激活

**使用示例：**
```lua
-- 调用 imgui.isItemActive 方法
imgui.isItemActive();
```

---

### imgui.isItemFocused
项是否聚焦

**使用示例：**
```lua
-- 调用 imgui.isItemFocused 方法
imgui.isItemFocused();
```

---

### imgui.isItemClicked
项是否被点击

**使用示例：**
```lua
-- 调用 imgui.isItemClicked 方法
imgui.isItemClicked();
```

---

### imgui.isItemVisible
项是否可见

**使用示例：**
```lua
-- 调用 imgui.isItemVisible 方法
imgui.isItemVisible();
```

---

### imgui.isItemEdited
项是否被编辑

**使用示例：**
```lua
-- 调用 imgui.isItemEdited 方法
imgui.isItemEdited();
```

---

### imgui.isItemActivated
项是否被激活

**使用示例：**
```lua
-- 调用 imgui.isItemActivated 方法
imgui.isItemActivated();
```

---

### imgui.isItemDeactivated
项是否被停用

**使用示例：**
```lua
-- 调用 imgui.isItemDeactivated 方法
imgui.isItemDeactivated();
```

---

### imgui.isItemDeactivatedAfterEdit
项是否在编辑后被停用

**使用示例：**
```lua
-- 调用 imgui.isItemDeactivatedAfterEdit 方法
imgui.isItemDeactivatedAfterEdit();
```

---

### imgui.isItemToggledOpen
项是否切换打开

**使用示例：**
```lua
-- 调用 imgui.isItemToggledOpen 方法
imgui.isItemToggledOpen();
```

---

### imgui.isMouseDragging
鼠标是否拖动

**使用示例：**
```lua
-- 调用 imgui.isMouseDragging 方法
imgui.isMouseDragging();
```

---

### imgui.isMouseHoveringRect
鼠标是否悬停在矩形上

**使用示例：**
```lua
-- 调用 imgui.isMouseHoveringRect 方法
imgui.isMouseHoveringRect();
```

---

### imgui.isMousePosValid
鼠标位置是否有效

**使用示例：**
```lua
-- 调用 imgui.isMousePosValid 方法
imgui.isMousePosValid();
```

---

### imgui.resetMouseDragDelta
重置鼠标拖动增量

**使用示例：**
```lua
-- 调用 imgui.resetMouseDragDelta 方法
imgui.resetMouseDragDelta();
```

---

### imgui.setKeyboardFocusHere
设置键盘焦点到此处

**使用示例：**
```lua
-- 调用 imgui.setKeyboardFocusHere 方法
imgui.setKeyboardFocusHere();
```

---

### imgui.setItemDefaultFocus
设置项默认焦点

**使用示例：**
```lua
-- 调用 imgui.setItemDefaultFocus 方法
imgui.setItemDefaultFocus();
```

---

### imgui.isAnyItemHovered
任何项是否悬停

**使用示例：**
```lua
-- 调用 imgui.isAnyItemHovered 方法
imgui.isAnyItemHovered();
```

---

### imgui.isAnyItemActive
任何项是否激活

**使用示例：**
```lua
-- 调用 imgui.isAnyItemActive 方法
imgui.isAnyItemActive();
```

---

### imgui.isAnyItemFocused
任何项是否聚焦

**使用示例：**
```lua
-- 调用 imgui.isAnyItemFocused 方法
imgui.isAnyItemFocused();
```

---

### imgui.isWindowHovered
窗口是否悬停

**使用示例：**
```lua
-- 调用 imgui.isWindowHovered 方法
imgui.isWindowHovered();
```

---

### imgui.isWindowFocused
窗口是否聚焦

**使用示例：**
```lua
-- 调用 imgui.isWindowFocused 方法
imgui.isWindowFocused();
```

---

### imgui.isWindowCollapsed
窗口是否折叠

**使用示例：**
```lua
-- 调用 imgui.isWindowCollapsed 方法
imgui.isWindowCollapsed();
```

---

### imgui.isRectVisible
矩形是否可见

**使用示例：**
```lua
-- 调用 imgui.isRectVisible 方法
imgui.isRectVisible();
```

---

### imgui.listBox
列表框

**使用示例：**
```lua
-- 调用 imgui.listBox 方法
imgui.listBox();
```

---

### imgui.beginMainMenuBar
开始主菜单栏

**使用示例：**
```lua
-- 调用 imgui.beginMainMenuBar 方法
imgui.beginMainMenuBar();
```

---

### imgui.endMainMenuBar
结束主菜单栏

**使用示例：**
```lua
-- 调用 imgui.endMainMenuBar 方法
imgui.endMainMenuBar();
```

---

### imgui.beginItemTooltip
开始项工具提示

**使用示例：**
```lua
-- 调用 imgui.beginItemTooltip 方法
imgui.beginItemTooltip();
```

---

### imgui.setNextItemOpen
设置下一项打开状态

**使用示例：**
```lua
-- 调用 imgui.setNextItemOpen 方法
imgui.setNextItemOpen();
```

---

### imgui.setColorEditOptions
设置颜色编辑选项

**使用示例：**
```lua
-- 调用 imgui.setColorEditOptions 方法
imgui.setColorEditOptions();
```

---

### imgui.setTabItemClosed
设置标签项关闭

**使用示例：**
```lua
-- 调用 imgui.setTabItemClosed 方法
imgui.setTabItemClosed();
```

---

### imgui.newDrawList
创建绘制列表

**使用示例：**
```lua
-- 调用 imgui.newDrawList 方法
imgui.newDrawList();
```

---

### imgui.newFont
创建字体对象

**使用示例：**
```lua
-- 调用 imgui.newFont 方法
imgui.newFont();
```

---

### imgui.newFontAtlas
创建字体图集

**使用示例：**
```lua
-- 调用 imgui.newFontAtlas 方法
imgui.newFontAtlas();
```

---

### imgui.newFontConfig
创建字体配置

**使用示例：**
```lua
-- 调用 imgui.newFontConfig 方法
imgui.newFontConfig();
```

---

### imgui.newDrawData
创建绘制数据

**使用示例：**
```lua
-- 调用 imgui.newDrawData 方法
imgui.newDrawData();
```

---

### imgui.newStorage
创建存储对象

**使用示例：**
```lua
-- 调用 imgui.newStorage 方法
imgui.newStorage();
```

---

### imgui.newPlatformIO
创建平台IO对象

**使用示例：**
```lua
-- 调用 imgui.newPlatformIO 方法
imgui.newPlatformIO();
```

---

### imgui.newDrawCmd
创建绘制命令

**使用示例：**
```lua
-- 调用 imgui.newDrawCmd 方法
imgui.newDrawCmd();
```

---

### imgui.newDrawVert
创建绘制顶点

**使用示例：**
```lua
-- 调用 imgui.newDrawVert 方法
imgui.newDrawVert();
```

---

### imgui.newFontGlyph
创建字体字形

**使用示例：**
```lua
-- 调用 imgui.newFontGlyph 方法
imgui.newFontGlyph();
```

---

### imgui.newFontBaked
创建烘焙字体

**使用示例：**
```lua
-- 调用 imgui.newFontBaked 方法
imgui.newFontBaked();
```

---

### imgui.newFontLoader
创建字体加载器

**使用示例：**
```lua
-- 调用 imgui.newFontLoader 方法
imgui.newFontLoader();
```

---

### imgui.newDrawListSharedData
创建绘制列表共享数据

**使用示例：**
```lua
-- 调用 imgui.newDrawListSharedData 方法
imgui.newDrawListSharedData();
```

---

### imgui.newDrawListSplitter
创建绘制列表分割器

**使用示例：**
```lua
-- 调用 imgui.newDrawListSplitter 方法
imgui.newDrawListSplitter();
```

---

### imgui.newDrawChannel
创建绘制通道

**使用示例：**
```lua
-- 调用 imgui.newDrawChannel 方法
imgui.newDrawChannel();
```

---

### imgui.newFontAtlasBuilder
创建字体图集构建器

**使用示例：**
```lua
-- 调用 imgui.newFontAtlasBuilder 方法
imgui.newFontAtlasBuilder();
```

---

### imgui.newFontAtlasRect
创建字体图集矩形

**使用示例：**
```lua
-- 调用 imgui.newFontAtlasRect 方法
imgui.newFontAtlasRect();
```

---

### imgui.newFontAtlasRectEntry
创建字体图集矩形条目

**使用示例：**
```lua
-- 调用 imgui.newFontAtlasRectEntry 方法
imgui.newFontAtlasRectEntry();
```

---

### imgui.newFontAtlasPostProcessData
创建字体图集后处理数据

**使用示例：**
```lua
-- 调用 imgui.newFontAtlasPostProcessData 方法
imgui.newFontAtlasPostProcessData();
```

---

### imgui.newFontGlyphRangesBuilder
创建字体字形范围构建器

**使用示例：**
```lua
-- 调用 imgui.newFontGlyphRangesBuilder 方法
imgui.newFontGlyphRangesBuilder();
```

---

### imgui.newFontStackData
创建字体栈数据

**使用示例：**
```lua
-- 调用 imgui.newFontStackData 方法
imgui.newFontStackData();
```

---

### imgui.newBoxSelectState
创建框选状态

**使用示例：**
```lua
-- 调用 imgui.newBoxSelectState 方法
imgui.newBoxSelectState();
```

---

### imgui.newColorMod
创建颜色修改

**使用示例：**
```lua
-- 调用 imgui.newColorMod 方法
imgui.newColorMod();
```

---

### imgui.newComboPreviewData
创建组合预览数据

**使用示例：**
```lua
-- 调用 imgui.newComboPreviewData 方法
imgui.newComboPreviewData();
```

---

### imgui.newContextHook
创建上下文钩子

**使用示例：**
```lua
-- 调用 imgui.newContextHook 方法
imgui.newContextHook();
```

---

### imgui.newDataTypeInfo
创建数据类型信息

**使用示例：**
```lua
-- 调用 imgui.newDataTypeInfo 方法
imgui.newDataTypeInfo();
```

---

### imgui.newDataTypeStorage
创建数据类型存储

**使用示例：**
```lua
-- 调用 imgui.newDataTypeStorage 方法
imgui.newDataTypeStorage();
```

---

### imgui.newDeactivatedItemData
创建停用项数据

**使用示例：**
```lua
-- 调用 imgui.newDeactivatedItemData 方法
imgui.newDeactivatedItemData();
```

---

### imgui.newDebugAllocEntry
创建调试分配条目

**使用示例：**
```lua
-- 调用 imgui.newDebugAllocEntry 方法
imgui.newDebugAllocEntry();
```

---

### imgui.newDebugAllocInfo
创建调试分配信息

**使用示例：**
```lua
-- 调用 imgui.newDebugAllocInfo 方法
imgui.newDebugAllocInfo();
```

---

### imgui.newDockContext
创建停靠上下文

**使用示例：**
```lua
-- 调用 imgui.newDockContext 方法
imgui.newDockContext();
```

---

### imgui.newDockNode
创建停靠节点

**使用示例：**
```lua
-- 调用 imgui.newDockNode 方法
imgui.newDockNode();
```

---

### imgui.newErrorRecoveryState
创建错误恢复状态

**使用示例：**
```lua
-- 调用 imgui.newErrorRecoveryState 方法
imgui.newErrorRecoveryState();
```

---

### imgui.newFocusScopeData
创建焦点范围数据

**使用示例：**
```lua
-- 调用 imgui.newFocusScopeData 方法
imgui.newFocusScopeData();
```

---

### imgui.newGroupData
创建组数据

**使用示例：**
```lua
-- 调用 imgui.newGroupData 方法
imgui.newGroupData();
```

---

### imgui.newIDStackTool
创建ID栈工具

**使用示例：**
```lua
-- 调用 imgui.newIDStackTool 方法
imgui.newIDStackTool();
```

---

### imgui.newInputEvent
创建输入事件

**使用示例：**
```lua
-- 调用 imgui.newInputEvent 方法
imgui.newInputEvent();
```

---

### imgui.newInputEventAppFocused
创建应用聚焦输入事件

**使用示例：**
```lua
-- 调用 imgui.newInputEventAppFocused 方法
imgui.newInputEventAppFocused();
```

---

### imgui.newInputEventKey
创建键盘输入事件

**使用示例：**
```lua
-- 调用 imgui.newInputEventKey 方法
imgui.newInputEventKey();
```

---

### imgui.newInputEventMouseButton
创建鼠标按钮输入事件

**使用示例：**
```lua
-- 调用 imgui.newInputEventMouseButton 方法
imgui.newInputEventMouseButton();
```

---

### imgui.newInputEventMousePos
创建鼠标位置输入事件

**使用示例：**
```lua
-- 调用 imgui.newInputEventMousePos 方法
imgui.newInputEventMousePos();
```

---

### imgui.newInputEventMouseViewport
创建鼠标视口输入事件

**使用示例：**
```lua
-- 调用 imgui.newInputEventMouseViewport 方法
imgui.newInputEventMouseViewport();
```

---

### imgui.newInputEventMouseWheel
创建鼠标滚轮输入事件

**使用示例：**
```lua
-- 调用 imgui.newInputEventMouseWheel 方法
imgui.newInputEventMouseWheel();
```

---

### imgui.newInputEventText
创建文本输入事件

**使用示例：**
```lua
-- 调用 imgui.newInputEventText 方法
imgui.newInputEventText();
```

---

### imgui.newInputTextCallbackData
创建输入文本回调数据

**使用示例：**
```lua
-- 调用 imgui.newInputTextCallbackData 方法
imgui.newInputTextCallbackData();
```

---

### imgui.newInputTextDeactivatedState
创建输入文本停用状态

**使用示例：**
```lua
-- 调用 imgui.newInputTextDeactivatedState 方法
imgui.newInputTextDeactivatedState();
```

---

### imgui.newInputTextState
创建输入文本状态

**使用示例：**
```lua
-- 调用 imgui.newInputTextState 方法
imgui.newInputTextState();
```

---

### imgui.newKeyData
创建键数据

**使用示例：**
```lua
-- 调用 imgui.newKeyData 方法
imgui.newKeyData();
```

---

### imgui.newKeyOwnerData
创建键拥有者数据

**使用示例：**
```lua
-- 调用 imgui.newKeyOwnerData 方法
imgui.newKeyOwnerData();
```

---

### imgui.newKeyRoutingData
创建键路由数据

**使用示例：**
```lua
-- 调用 imgui.newKeyRoutingData 方法
imgui.newKeyRoutingData();
```

---

### imgui.newKeyRoutingTable
创建键路由表

**使用示例：**
```lua
-- 调用 imgui.newKeyRoutingTable 方法
imgui.newKeyRoutingTable();
```

---

### imgui.newLastItemData
创建最后项数据

**使用示例：**
```lua
-- 调用 imgui.newLastItemData 方法
imgui.newLastItemData();
```

---

### imgui.newListClipper
创建列表剪裁器

**使用示例：**
```lua
-- 调用 imgui.newListClipper 方法
imgui.newListClipper();
```

---

### imgui.newListClipperData
创建列表剪裁器数据

**使用示例：**
```lua
-- 调用 imgui.newListClipperData 方法
imgui.newListClipperData();
```

---

### imgui.newListClipperRange
创建列表剪裁器范围

**使用示例：**
```lua
-- 调用 imgui.newListClipperRange 方法
imgui.newListClipperRange();
```

---

### imgui.newLocEntry
创建本地化条目

**使用示例：**
```lua
-- 调用 imgui.newLocEntry 方法
imgui.newLocEntry();
```

---

### imgui.newMenuColumns
创建菜单列

**使用示例：**
```lua
-- 调用 imgui.newMenuColumns 方法
imgui.newMenuColumns();
```

---

### imgui.newMetricsConfig
创建指标配置

**使用示例：**
```lua
-- 调用 imgui.newMetricsConfig 方法
imgui.newMetricsConfig();
```

---

### imgui.newMultiSelectIO
创建多选IO

**使用示例：**
```lua
-- 调用 imgui.newMultiSelectIO 方法
imgui.newMultiSelectIO();
```

---

### imgui.newMultiSelectState
创建多选状态

**使用示例：**
```lua
-- 调用 imgui.newMultiSelectState 方法
imgui.newMultiSelectState();
```

---

### imgui.newMultiSelectTempData
创建多选临时数据

**使用示例：**
```lua
-- 调用 imgui.newMultiSelectTempData 方法
imgui.newMultiSelectTempData();
```

---

### imgui.newNavItemData
创建导航项数据

**使用示例：**
```lua
-- 调用 imgui.newNavItemData 方法
imgui.newNavItemData();
```

---

### imgui.newNextItemData
创建下一项数据

**使用示例：**
```lua
-- 调用 imgui.newNextItemData 方法
imgui.newNextItemData();
```

---

### imgui.newNextWindowData
创建下一窗口数据

**使用示例：**
```lua
-- 调用 imgui.newNextWindowData 方法
imgui.newNextWindowData();
```

---

### imgui.newOldColumnData
创建旧列数据

**使用示例：**
```lua
-- 调用 imgui.newOldColumnData 方法
imgui.newOldColumnData();
```

---

### imgui.newOldColumns
创建旧列

**使用示例：**
```lua
-- 调用 imgui.newOldColumns 方法
imgui.newOldColumns();
```

---

### imgui.newOnceUponAFrame
创建每帧一次

**使用示例：**
```lua
-- 调用 imgui.newOnceUponAFrame 方法
imgui.newOnceUponAFrame();
```

---

### imgui.newPayload
创建载荷

**使用示例：**
```lua
-- 调用 imgui.newPayload 方法
imgui.newPayload();
```

---

### imgui.newPlatformImeData
创建平台IME数据

**使用示例：**
```lua
-- 调用 imgui.newPlatformImeData 方法
imgui.newPlatformImeData();
```

---

### imgui.newPlatformMonitor
创建平台监视器

**使用示例：**
```lua
-- 调用 imgui.newPlatformMonitor 方法
imgui.newPlatformMonitor();
```

---

### imgui.newPopupData
创建弹出数据

**使用示例：**
```lua
-- 调用 imgui.newPopupData 方法
imgui.newPopupData();
```

---

### imgui.newPtrOrIndex
创建指针或索引

**使用示例：**
```lua
-- 调用 imgui.newPtrOrIndex 方法
imgui.newPtrOrIndex();
```

---

### imgui.newSelectionBasicStorage
创建选择基本存储

**使用示例：**
```lua
-- 调用 imgui.newSelectionBasicStorage 方法
imgui.newSelectionBasicStorage();
```

---

### imgui.newSelectionExternalStorage
创建选择外部存储

**使用示例：**
```lua
-- 调用 imgui.newSelectionExternalStorage 方法
imgui.newSelectionExternalStorage();
```

---

### imgui.newSelectionRequest
创建选择请求

**使用示例：**
```lua
-- 调用 imgui.newSelectionRequest 方法
imgui.newSelectionRequest();
```

---

### imgui.newSettingsHandler
创建设置处理器

**使用示例：**
```lua
-- 调用 imgui.newSettingsHandler 方法
imgui.newSettingsHandler();
```

---

### imgui.newShrinkWidthItem
创建收缩宽度项

**使用示例：**
```lua
-- 调用 imgui.newShrinkWidthItem 方法
imgui.newShrinkWidthItem();
```

---

### imgui.newSizeCallbackData
创建大小回调数据

**使用示例：**
```lua
-- 调用 imgui.newSizeCallbackData 方法
imgui.newSizeCallbackData();
```

---

### imgui.newStackLevelInfo
创建栈级别信息

**使用示例：**
```lua
-- 调用 imgui.newStackLevelInfo 方法
imgui.newStackLevelInfo();
```

---

### imgui.newStoragePair
创建存储对

**使用示例：**
```lua
-- 调用 imgui.newStoragePair 方法
imgui.newStoragePair();
```

---

### imgui.newStyleMod
创建样式修改

**使用示例：**
```lua
-- 调用 imgui.newStyleMod 方法
imgui.newStyleMod();
```

---

### imgui.newStyleVarInfo
创建样式变量信息

**使用示例：**
```lua
-- 调用 imgui.newStyleVarInfo 方法
imgui.newStyleVarInfo();
```

---

### imgui.newTabBar
创建标签栏

**使用示例：**
```lua
-- 调用 imgui.newTabBar 方法
imgui.newTabBar();
```

---

### imgui.newTabItem
创建标签项

**使用示例：**
```lua
-- 调用 imgui.newTabItem 方法
imgui.newTabItem();
```

---

### imgui.newTable
创建表格

**使用示例：**
```lua
-- 调用 imgui.newTable 方法
imgui.newTable();
```

---

### imgui.newTableCellData
创建表格单元格数据

**使用示例：**
```lua
-- 调用 imgui.newTableCellData 方法
imgui.newTableCellData();
```

---

### imgui.newTableColumn
创建表格列

**使用示例：**
```lua
-- 调用 imgui.newTableColumn 方法
imgui.newTableColumn();
```

---

### imgui.newViewport
创建视口

**使用示例：**
```lua
-- 调用 imgui.newViewport 方法
imgui.newViewport();
```

---

### imgui.newWindowClass
创建窗口类

**使用示例：**
```lua
-- 调用 imgui.newWindowClass 方法
imgui.newWindowClass();
```

---

### imgui.newTextBuffer
创建文本缓冲区

**使用示例：**
```lua
-- 调用 imgui.newTextBuffer 方法
imgui.newTextBuffer();
```

---

### imgui.newTextFilter
创建文本过滤器

**使用示例：**
```lua
-- 调用 imgui.newTextFilter 方法
imgui.newTextFilter();
```

---

### imgui.newTextRange
创建文本范围

**使用示例：**
```lua
-- 调用 imgui.newTextRange 方法
imgui.newTextRange();
```

---

### imgui.newTypingSelectState
创建打字选择状态

**使用示例：**
```lua
-- 调用 imgui.newTypingSelectState 方法
imgui.newTypingSelectState();
```

---

### imgui.newViewportP
创建视口指针

**使用示例：**
```lua
-- 调用 imgui.newViewportP 方法
imgui.newViewportP();
```

---

### imgui.newTextureData
创建纹理数据

**使用示例：**
```lua
-- 调用 imgui.newTextureData 方法
imgui.newTextureData();
```

---

### imgui.newTextureRef
创建纹理引用

**使用示例：**
```lua
-- 调用 imgui.newTextureRef 方法
imgui.newTextureRef();
```

---

### imgui.newDrawCmdHeader
创建绘制命令头

**使用示例：**
```lua
-- 调用 imgui.newDrawCmdHeader 方法
imgui.newDrawCmdHeader();
```

---

### imgui.newDrawDataBuilder
创建绘制数据构建器

**使用示例：**
```lua
-- 调用 imgui.newDrawDataBuilder 方法
imgui.newDrawDataBuilder();
```

---

### imgui.newBitVector
创建位向量

**使用示例：**
```lua
-- 调用 imgui.newBitVector 方法
imgui.newBitVector();
```

---

### imgui.newBitArrayForNamedKeys
创建命名键位数组

**使用示例：**
```lua
-- 调用 imgui.newBitArrayForNamedKeys 方法
imgui.newBitArrayForNamedKeys();
```

---

### imgui.colorHSVV
HSV颜色转RGBA

**使用示例：**
```lua
-- 调用 imgui.colorHSVV 方法
imgui.colorHSVV();
```

---

### imgui.alignTextToFramePadding
对齐文本到框架内边距

**使用示例：**
```lua
-- 调用 imgui.alignTextToFramePadding 方法
imgui.alignTextToFramePadding();
```

---

### imgui.arrowButton
箭头按钮

**使用示例：**
```lua
-- 调用 imgui.arrowButton 方法
imgui.arrowButton();
```

---

### imgui.progressBar
进度条

**使用示例：**
```lua
-- 调用 imgui.progressBar 方法
imgui.progressBar();
```

---

### imgui.progressBarWithSize
带大小的进度条

**使用示例：**
```lua
-- 调用 imgui.progressBarWithSize 方法
imgui.progressBarWithSize();
```

---

### imgui.progressBarWithOverlay
带覆盖的进度条

**使用示例：**
```lua
-- 调用 imgui.progressBarWithOverlay 方法
imgui.progressBarWithOverlay();
```

---

### imgui.setScrollHereX
设置滚动到此处X

**使用示例：**
```lua
-- 调用 imgui.setScrollHereX 方法
imgui.setScrollHereX();
```

---

### imgui.setScrollHereY
设置滚动到此处Y

**使用示例：**
```lua
-- 调用 imgui.setScrollHereY 方法
imgui.setScrollHereY();
```

---

### imgui.setNextWindowScroll
设置下一窗口滚动

**使用示例：**
```lua
-- 调用 imgui.setNextWindowScroll 方法
imgui.setNextWindowScroll();
```

---

### imgui.setTooltip
设置工具提示

**使用示例：**
```lua
-- 调用 imgui.setTooltip 方法
imgui.setTooltip();
```

---

### imgui.beginTooltip
开始工具提示

**使用示例：**
```lua
-- 调用 imgui.beginTooltip 方法
imgui.beginTooltip();
```

---

### imgui.endTooltip
结束工具提示

**使用示例：**
```lua
-- 调用 imgui.endTooltip 方法
imgui.endTooltip();
```

---

### imgui.openPopup
打开弹出窗口

**使用示例：**
```lua
-- 调用 imgui.openPopup 方法
imgui.openPopup();
```

---

### imgui.openPopupOnItemClick
点击项时打开弹出窗口

**使用示例：**
```lua
-- 调用 imgui.openPopupOnItemClick 方法
imgui.openPopupOnItemClick();
```

---

### imgui.closeCurrentPopup
关闭当前弹出窗口

**使用示例：**
```lua
-- 调用 imgui.closeCurrentPopup 方法
imgui.closeCurrentPopup();
```

---

### imgui.beginPopupModal
开始模态弹出窗口

**使用示例：**
```lua
-- 调用 imgui.beginPopupModal 方法
imgui.beginPopupModal();
```

---

### imgui.beginPopupContextWindow
开始窗口上下文弹出窗口

**使用示例：**
```lua
-- 调用 imgui.beginPopupContextWindow 方法
imgui.beginPopupContextWindow();
```

---

### imgui.beginPopupContextVoid
开始空白上下文弹出窗口

**使用示例：**
```lua
-- 调用 imgui.beginPopupContextVoid 方法
imgui.beginPopupContextVoid();
```

---

### imgui.beginPopup
开始弹出窗口

**使用示例：**
```lua
-- 调用 imgui.beginPopup 方法
imgui.beginPopup();
```

---

### imgui.endPopup
结束弹出窗口

**使用示例：**
```lua
-- 调用 imgui.endPopup 方法
imgui.endPopup();
```

---

### imgui.calcTextSize
计算文本大小

**使用示例：**
```lua
-- 调用 imgui.calcTextSize 方法
imgui.calcTextSize();
```

---

### imgui.pushClipRect
推入剪裁矩形

**使用示例：**
```lua
-- 调用 imgui.pushClipRect 方法
imgui.pushClipRect();
```

---

### imgui.popClipRect
弹出剪裁矩形

**使用示例：**
```lua
-- 调用 imgui.popClipRect 方法
imgui.popClipRect();
```

---

### imgui.setItemKeyOwner
设置项键拥有者

**使用示例：**
```lua
-- 调用 imgui.setItemKeyOwner 方法
imgui.setItemKeyOwner();
```

---

### imgui.setNextItemAllowOverlap
允许下一项重叠

**使用示例：**
```lua
-- 调用 imgui.setNextItemAllowOverlap 方法
imgui.setNextItemAllowOverlap();
```

---

### imgui.isItemToggledSelection
项是否切换选择

**使用示例：**
```lua
-- 调用 imgui.isItemToggledSelection 方法
imgui.isItemToggledSelection();
```

---

### imgui.beginMultiSelect
开始多选

**使用示例：**
```lua
-- 调用 imgui.beginMultiSelect 方法
imgui.beginMultiSelect();
```

---

### imgui.endMultiSelect
结束多选

**使用示例：**
```lua
-- 调用 imgui.endMultiSelect 方法
imgui.endMultiSelect();
```

---

### imgui.resetMouseDragDelta
重置鼠标拖拽增量

**使用示例：**
```lua
-- 调用 imgui.resetMouseDragDelta 方法
imgui.resetMouseDragDelta();
```

---

### imgui.setNavCursorVisible
设置导航光标可见

**使用示例：**
```lua
-- 调用 imgui.setNavCursorVisible 方法
imgui.setNavCursorVisible();
```

---

### imgui.setMouseCursor
设置鼠标光标

**使用示例：**
```lua
-- 调用 imgui.setMouseCursor 方法
imgui.setMouseCursor();
```

---

### imgui.setNextFrameWantCaptureMouse
设置下一帧想要捕获鼠标

**使用示例：**
```lua
-- 调用 imgui.setNextFrameWantCaptureMouse 方法
imgui.setNextFrameWantCaptureMouse();
```

---

### imgui.setNextFrameWantCaptureKeyboard
设置下一帧想要捕获键盘

**使用示例：**
```lua
-- 调用 imgui.setNextFrameWantCaptureKeyboard 方法
imgui.setNextFrameWantCaptureKeyboard();
```

---

### imgui.pushStyleVarFloat
推入样式变量（浮点）

**使用示例：**
```lua
-- 调用 imgui.pushStyleVarFloat 方法
imgui.pushStyleVarFloat();
```

---

### imgui.pushStyleVarVec2
推入样式变量（向量）

**使用示例：**
```lua
-- 调用 imgui.pushStyleVarVec2 方法
imgui.pushStyleVarVec2();
```

---

### imgui.popStyleVar
弹出样式变量

**使用示例：**
```lua
-- 调用 imgui.popStyleVar 方法
imgui.popStyleVar();
```

---

### imgui.pushFont
推入字体

**使用示例：**
```lua
-- 调用 imgui.pushFont 方法
imgui.pushFont();
```

---

### imgui.popFont
弹出字体

**使用示例：**
```lua
-- 调用 imgui.popFont 方法
imgui.popFont();
```

---

### imgui.pushItemFlag
推入项标志

**使用示例：**
```lua
-- 调用 imgui.pushItemFlag 方法
imgui.pushItemFlag();
```

---

### imgui.popItemFlag
弹出项标志

**使用示例：**
```lua
-- 调用 imgui.popItemFlag 方法
imgui.popItemFlag();
```

---

### imgui.pushIDStr
推入ID（字符串）

**使用示例：**
```lua
-- 调用 imgui.pushIDStr 方法
imgui.pushIDStr();
```

---

### imgui.pushIDInt
推入ID（整数）

**使用示例：**
```lua
-- 调用 imgui.pushIDInt 方法
imgui.pushIDInt();
```

---

### imgui.pushIDPtr
推入ID（指针）

**使用示例：**
```lua
-- 调用 imgui.pushIDPtr 方法
imgui.pushIDPtr();
```

---

### imgui.popID
弹出ID

**使用示例：**
```lua
-- 调用 imgui.popID 方法
imgui.popID();
```

---

### imgui.pushTextWrapPos
推入文本换行位置

**使用示例：**
```lua
-- 调用 imgui.pushTextWrapPos 方法
imgui.pushTextWrapPos();
```

---

### imgui.popTextWrapPos
弹出文本换行位置

**使用示例：**
```lua
-- 调用 imgui.popTextWrapPos 方法
imgui.popTextWrapPos();
```

---

### imgui.newLine
新行

**使用示例：**
```lua
-- 调用 imgui.newLine 方法
imgui.newLine();
```

---

### imgui.spacing
间距

**使用示例：**
```lua
-- 调用 imgui.spacing 方法
imgui.spacing();
```

---

### imgui.dummy
虚拟元素

**使用示例：**
```lua
-- 调用 imgui.dummy 方法
imgui.dummy();
```

---

### imgui.indent
缩进

**使用示例：**
```lua
-- 调用 imgui.indent 方法
imgui.indent();
```

---

### imgui.unindent
取消缩进

**使用示例：**
```lua
-- 调用 imgui.unindent 方法
imgui.unindent();
```

---

### imgui.beginGroup
开始组

**使用示例：**
```lua
-- 调用 imgui.beginGroup 方法
imgui.beginGroup();
```

---

### imgui.endGroup
结束组

**使用示例：**
```lua
-- 调用 imgui.endGroup 方法
imgui.endGroup();
```

---

### imgui.getCursorPos
获取光标位置

**使用示例：**
```lua
-- 调用 imgui.getCursorPos 方法
imgui.getCursorPos();
```

---

### imgui.getCursorPosX
获取光标X位置

**使用示例：**
```lua
-- 调用 imgui.getCursorPosX 方法
imgui.getCursorPosX();
```

---

### imgui.getCursorPosY
获取光标Y位置

**使用示例：**
```lua
-- 调用 imgui.getCursorPosY 方法
imgui.getCursorPosY();
```

---

### imgui.setCursorPos
设置光标位置

**使用示例：**
```lua
-- 调用 imgui.setCursorPos 方法
imgui.setCursorPos();
```

---

### imgui.setCursorPosX
设置光标X位置

**使用示例：**
```lua
-- 调用 imgui.setCursorPosX 方法
imgui.setCursorPosX();
```

---

### imgui.setCursorPosY
设置光标Y位置

**使用示例：**
```lua
-- 调用 imgui.setCursorPosY 方法
imgui.setCursorPosY();
```

---

### imgui.getCursorStartPos
获取光标起始位置

**使用示例：**
```lua
-- 调用 imgui.getCursorStartPos 方法
imgui.getCursorStartPos();
```

---

### imgui.getCursorScreenPos
获取屏幕光标位置

**使用示例：**
```lua
-- 调用 imgui.getCursorScreenPos 方法
imgui.getCursorScreenPos();
```

---

### imgui.setCursorScreenPos
设置屏幕光标位置

**使用示例：**
```lua
-- 调用 imgui.setCursorScreenPos 方法
imgui.setCursorScreenPos();
```

---

### imgui.getTextLineHeight
获取文本行高度

**使用示例：**
```lua
-- 调用 imgui.getTextLineHeight 方法
imgui.getTextLineHeight();
```

---

### imgui.getTextLineHeightWithSpacing
获取带间距的文本行高度

**使用示例：**
```lua
-- 调用 imgui.getTextLineHeightWithSpacing 方法
imgui.getTextLineHeightWithSpacing();
```

---

### imgui.getFrameHeight
获取框架高度

**使用示例：**
```lua
-- 调用 imgui.getFrameHeight 方法
imgui.getFrameHeight();
```

---

### imgui.getFrameHeightWithSpacing
获取带间距的框架高度

**使用示例：**
```lua
-- 调用 imgui.getFrameHeightWithSpacing 方法
imgui.getFrameHeightWithSpacing();
```

---

### imgui.getContentRegionMax
获取内容区域最大值

**使用示例：**
```lua
-- 调用 imgui.getContentRegionMax 方法
imgui.getContentRegionMax();
```

---

### imgui.getContentRegionAvail
获取可用内容区域

**使用示例：**
```lua
-- 调用 imgui.getContentRegionAvail 方法
imgui.getContentRegionAvail();
```

---

### imgui.popStyleColor
弹出样式颜色

**使用示例：**
```lua
-- 调用 imgui.popStyleColor 方法
imgui.popStyleColor();
```

---

### imgui.getStyleColor
获取样式颜色

**使用示例：**
```lua
-- 调用 imgui.getStyleColor 方法
imgui.getStyleColor();
```

---

### imgui.getStyleColorName
获取样式颜色名称

**使用示例：**
```lua
-- 调用 imgui.getStyleColorName 方法
imgui.getStyleColorName();
```

---

### imgui.beginMainMenuBar
开始主菜单栏

**使用示例：**
```lua
-- 调用 imgui.beginMainMenuBar 方法
imgui.beginMainMenuBar();
```

---

### imgui.endMainMenuBar
结束主菜单栏

**使用示例：**
```lua
-- 调用 imgui.endMainMenuBar 方法
imgui.endMainMenuBar();
```

---

### imgui.beginMenuBar
开始菜单栏

**使用示例：**
```lua
-- 调用 imgui.beginMenuBar 方法
imgui.beginMenuBar();
```

---

### imgui.endMenuBar
结束菜单栏

**使用示例：**
```lua
-- 调用 imgui.endMenuBar 方法
imgui.endMenuBar();
```

---

### imgui.beginMenu
开始菜单

**使用示例：**
```lua
-- 调用 imgui.beginMenu 方法
imgui.beginMenu();
```

---

### imgui.beginMenuDisabled
开始禁用菜单

**使用示例：**
```lua
-- 调用 imgui.beginMenuDisabled 方法
imgui.beginMenuDisabled();
```

---

### imgui.endMenu
结束菜单

**使用示例：**
```lua
-- 调用 imgui.endMenu 方法
imgui.endMenu();
```

---

### imgui.beginTabBar
开始标签栏

**使用示例：**
```lua
-- 调用 imgui.beginTabBar 方法
imgui.beginTabBar();
```

---

### imgui.beginTabBarWithFlags
开始标签栏（带标志）

**使用示例：**
```lua
-- 调用 imgui.beginTabBarWithFlags 方法
imgui.beginTabBarWithFlags();
```

---

### imgui.endTabBar
结束标签栏

**使用示例：**
```lua
-- 调用 imgui.endTabBar 方法
imgui.endTabBar();
```

---

### imgui.beginTabItem
开始标签项

**使用示例：**
```lua
-- 调用 imgui.beginTabItem 方法
imgui.beginTabItem();
```

---

### imgui.beginTabItemWithOpen
开始标签项（带打开状态）

**使用示例：**
```lua
-- 调用 imgui.beginTabItemWithOpen 方法
imgui.beginTabItemWithOpen();
```

---

### imgui.beginTabItemWithFlags
开始标签项（带标志）

**使用示例：**
```lua
-- 调用 imgui.beginTabItemWithFlags 方法
imgui.beginTabItemWithFlags();
```

---

### imgui.endTabItem
结束标签项

**使用示例：**
```lua
-- 调用 imgui.endTabItem 方法
imgui.endTabItem();
```

---

### imgui.setTabItemClosed
设置标签项关闭

**使用示例：**
```lua
-- 调用 imgui.setTabItemClosed 方法
imgui.setTabItemClosed();
```

---

### imgui.isItemHovered
项是否悬停

**使用示例：**
```lua
-- 调用 imgui.isItemHovered 方法
imgui.isItemHovered();
```

---

### imgui.isItemActive
项是否激活

**使用示例：**
```lua
-- 调用 imgui.isItemActive 方法
imgui.isItemActive();
```

---

### imgui.isItemFocused
项是否聚焦

**使用示例：**
```lua
-- 调用 imgui.isItemFocused 方法
imgui.isItemFocused();
```

---

### imgui.isItemClicked
项是否被点击

**使用示例：**
```lua
-- 调用 imgui.isItemClicked 方法
imgui.isItemClicked();
```

---

### imgui.isItemVisible
项是否可见

**使用示例：**
```lua
-- 调用 imgui.isItemVisible 方法
imgui.isItemVisible();
```

---

### imgui.isItemEdited
项是否被编辑

**使用示例：**
```lua
-- 调用 imgui.isItemEdited 方法
imgui.isItemEdited();
```

---

### imgui.isItemActivated
项是否被激活

**使用示例：**
```lua
-- 调用 imgui.isItemActivated 方法
imgui.isItemActivated();
```

---

### imgui.isItemDeactivated
项是否被停用

**使用示例：**
```lua
-- 调用 imgui.isItemDeactivated 方法
imgui.isItemDeactivated();
```

---

### imgui.isItemDeactivatedAfterEdit
项是否在编辑后被停用

**使用示例：**
```lua
-- 调用 imgui.isItemDeactivatedAfterEdit 方法
imgui.isItemDeactivatedAfterEdit();
```

---

### imgui.isItemToggledOpen
项是否切换打开状态

**使用示例：**
```lua
-- 调用 imgui.isItemToggledOpen 方法
imgui.isItemToggledOpen();
```

---

### imgui.isAnyItemHovered
是否有任何项悬停

**使用示例：**
```lua
-- 调用 imgui.isAnyItemHovered 方法
imgui.isAnyItemHovered();
```

---

### imgui.isAnyItemActive
是否有任何项激活

**使用示例：**
```lua
-- 调用 imgui.isAnyItemActive 方法
imgui.isAnyItemActive();
```

---

### imgui.isAnyItemFocused
是否有任何项聚焦

**使用示例：**
```lua
-- 调用 imgui.isAnyItemFocused 方法
imgui.isAnyItemFocused();
```

---

### imgui.getItemRectMin
获取项矩形最小值

**使用示例：**
```lua
-- 调用 imgui.getItemRectMin 方法
imgui.getItemRectMin();
```

---

### imgui.getItemRectMax
获取项矩形最大值

**使用示例：**
```lua
-- 调用 imgui.getItemRectMax 方法
imgui.getItemRectMax();
```

---

### imgui.getItemRectSize
获取项矩形大小

**使用示例：**
```lua
-- 调用 imgui.getItemRectSize 方法
imgui.getItemRectSize();
```

---

### imgui.setItemDefaultFocus
设置项默认焦点

**使用示例：**
```lua
-- 调用 imgui.setItemDefaultFocus 方法
imgui.setItemDefaultFocus();
```

---

### imgui.setKeyboardFocusHere
设置键盘焦点到此处

**使用示例：**
```lua
-- 调用 imgui.setKeyboardFocusHere 方法
imgui.setKeyboardFocusHere();
```

---

### imgui.isMouseHoveringRect
鼠标是否悬停在矩形上

**使用示例：**
```lua
-- 调用 imgui.isMouseHoveringRect 方法
imgui.isMouseHoveringRect();
```

---

### imgui.isMousePosValid
鼠标位置是否有效

**使用示例：**
```lua
-- 调用 imgui.isMousePosValid 方法
imgui.isMousePosValid();
```

---

### imgui.getMousePos
获取鼠标位置

**使用示例：**
```lua
-- 调用 imgui.getMousePos 方法
imgui.getMousePos();
```

---

### imgui.getMousePosOnOpeningCurrentPopup
获取打开当前弹出窗口时的鼠标位置

**使用示例：**
```lua
-- 调用 imgui.getMousePosOnOpeningCurrentPopup 方法
imgui.getMousePosOnOpeningCurrentPopup();
```

---

### imgui.getMouseDragDelta
获取鼠标拖拽增量

**使用示例：**
```lua
-- 调用 imgui.getMouseDragDelta 方法
imgui.getMouseDragDelta();
```

---

### imgui.resetMouseDragDelta
重置鼠标拖拽增量

**使用示例：**
```lua
-- 调用 imgui.resetMouseDragDelta 方法
imgui.resetMouseDragDelta();
```

---

### imgui.getMouseCursor
获取鼠标光标

**使用示例：**
```lua
-- 调用 imgui.getMouseCursor 方法
imgui.getMouseCursor();
```

---

### imgui.setMouseCursor
设置鼠标光标

**使用示例：**
```lua
-- 调用 imgui.setMouseCursor 方法
imgui.setMouseCursor();
```

---

### imgui.getClipboardText
获取剪贴板文本

**使用示例：**
```lua
-- 调用 imgui.getClipboardText 方法
imgui.getClipboardText();
```

---

### imgui.setClipboardText
设置剪贴板文本

**使用示例：**
```lua
-- 调用 imgui.setClipboardText 方法
imgui.setClipboardText();
```

---

### imgui.getTime
获取时间

**使用示例：**
```lua
-- 调用 imgui.getTime 方法
imgui.getTime();
```

---

### imgui.getFrameCount
获取帧数

**使用示例：**
```lua
-- 调用 imgui.getFrameCount 方法
imgui.getFrameCount();
```

---

### imgui.getStyleColorVec4
获取样式颜色Vec4

**使用示例：**
```lua
-- 调用 imgui.getStyleColorVec4 方法
imgui.getStyleColorVec4();
```

---

### imgui.getColorU32
获取颜色U32

**使用示例：**
```lua
-- 调用 imgui.getColorU32 方法
imgui.getColorU32();
```

---

### imgui.getColorU32FromRGBA
从RGBA获取颜色U32

**使用示例：**
```lua
-- 调用 imgui.getColorU32FromRGBA 方法
imgui.getColorU32FromRGBA();
```

---

### imgui.getColorU32FromU32
从U32获取颜色U32

**使用示例：**
```lua
-- 调用 imgui.getColorU32FromU32 方法
imgui.getColorU32FromU32();
```

---

### imgui.getFont
获取字体

**使用示例：**
```lua
-- 调用 imgui.getFont 方法
imgui.getFont();
```

---

### imgui.getFontSize
获取字体大小

**使用示例：**
```lua
-- 调用 imgui.getFontSize 方法
imgui.getFontSize();
```

---

### imgui.getFontTexUvWhitePixel
获取字体纹理UV白色像素

**使用示例：**
```lua
-- 调用 imgui.getFontTexUvWhitePixel 方法
imgui.getFontTexUvWhitePixel();
```

---

### imgui.getColorConvertRGBtoHSV
RGB转HSV

**使用示例：**
```lua
-- 调用 imgui.getColorConvertRGBtoHSV 方法
imgui.getColorConvertRGBtoHSV();
```

---

### imgui.getColorConvertHSVtoRGB
HSV转RGB

**使用示例：**
```lua
-- 调用 imgui.getColorConvertHSVtoRGB 方法
imgui.getColorConvertHSVtoRGB();
```

---

### imgui.getWindowDrawList
获取窗口绘制列表

**使用示例：**
```lua
-- 调用 imgui.getWindowDrawList 方法
imgui.getWindowDrawList();
```

---

### imgui.getBackgroundDrawList
获取背景绘制列表

**使用示例：**
```lua
-- 调用 imgui.getBackgroundDrawList 方法
imgui.getBackgroundDrawList();
```

---

### imgui.getForegroundDrawList
获取前景绘制列表

**使用示例：**
```lua
-- 调用 imgui.getForegroundDrawList 方法
imgui.getForegroundDrawList();
```

---

### imgui.getMainViewport
获取主视口

**使用示例：**
```lua
-- 调用 imgui.getMainViewport 方法
imgui.getMainViewport();
```

---

### imgui.findViewportByID
通过ID查找视口

**使用示例：**
```lua
-- 调用 imgui.findViewportByID 方法
imgui.findViewportByID();
```

---

### imgui.findViewportByPlatformHandle
通过平台句柄查找视口

**使用示例：**
```lua
-- 调用 imgui.findViewportByPlatformHandle 方法
imgui.findViewportByPlatformHandle();
```

---

### imgui.getPlatformIO
获取平台IO

**使用示例：**
```lua
-- 调用 imgui.getPlatformIO 方法
imgui.getPlatformIO();
```

---

### imgui.getIO
获取IO

**使用示例：**
```lua
-- 调用 imgui.getIO 方法
imgui.getIO();
```

---

### imgui.getStyle
获取样式

**使用示例：**
```lua
-- 调用 imgui.getStyle 方法
imgui.getStyle();
```

---

### imgui.getDragDropPayload
获取拖放载荷

**使用示例：**
```lua
-- 调用 imgui.getDragDropPayload 方法
imgui.getDragDropPayload();
```

---

### imgui.beginTooltip
开始工具提示

**使用示例：**
```lua
-- 调用 imgui.beginTooltip 方法
imgui.beginTooltip();
```

---

### imgui.endTooltip
结束工具提示

**使用示例：**
```lua
-- 调用 imgui.endTooltip 方法
imgui.endTooltip();
```

---

### imgui.setTooltip
设置工具提示

**使用示例：**
```lua
-- 调用 imgui.setTooltip 方法
imgui.setTooltip();
```

---

### imgui.beginItemTooltip
开始项工具提示

**使用示例：**
```lua
-- 调用 imgui.beginItemTooltip 方法
imgui.beginItemTooltip();
```

---

### imgui.beginPopup
开始弹出窗口

**使用示例：**
```lua
-- 调用 imgui.beginPopup 方法
imgui.beginPopup();
```

---

### imgui.beginPopupWithFlags
开始弹出窗口（带标志）

**使用示例：**
```lua
-- 调用 imgui.beginPopupWithFlags 方法
imgui.beginPopupWithFlags();
```

---

### imgui.endPopup
结束弹出窗口

**使用示例：**
```lua
-- 调用 imgui.endPopup 方法
imgui.endPopup();
```

---

### imgui.openPopup
打开弹出窗口

**使用示例：**
```lua
-- 调用 imgui.openPopup 方法
imgui.openPopup();
```

---

### imgui.openPopupWithFlags
打开弹出窗口（带标志）

**使用示例：**
```lua
-- 调用 imgui.openPopupWithFlags 方法
imgui.openPopupWithFlags();
```

---

### imgui.closeCurrentPopup
关闭当前弹出窗口

**使用示例：**
```lua
-- 调用 imgui.closeCurrentPopup 方法
imgui.closeCurrentPopup();
```

---

### imgui.beginPopupModal
开始模态弹出窗口

**使用示例：**
```lua
-- 调用 imgui.beginPopupModal 方法
imgui.beginPopupModal();
```

---

### imgui.beginPopupModalWithOpen
开始模态弹出窗口（带打开状态）

**使用示例：**
```lua
-- 调用 imgui.beginPopupModalWithOpen 方法
imgui.beginPopupModalWithOpen();
```

---

### imgui.beginPopupModalWithFlags
开始模态弹出窗口（带标志）

**使用示例：**
```lua
-- 调用 imgui.beginPopupModalWithFlags 方法
imgui.beginPopupModalWithFlags();
```

---

### imgui.beginPopupContextItem
开始项上下文弹出窗口

**使用示例：**
```lua
-- 调用 imgui.beginPopupContextItem 方法
imgui.beginPopupContextItem();
```

---

### imgui.beginPopupContextWindow
开始窗口上下文弹出窗口

**使用示例：**
```lua
-- 调用 imgui.beginPopupContextWindow 方法
imgui.beginPopupContextWindow();
```

---

### imgui.beginPopupContextVoid
开始空白上下文弹出窗口

**使用示例：**
```lua
-- 调用 imgui.beginPopupContextVoid 方法
imgui.beginPopupContextVoid();
```

---

### imgui.beginColumns
开始列

**使用示例：**
```lua
-- 调用 imgui.beginColumns 方法
imgui.beginColumns();
```

---

### imgui.getColumnIndex
获取列索引

**使用示例：**
```lua
-- 调用 imgui.getColumnIndex 方法
imgui.getColumnIndex();
```

---

### imgui.getColumnCount
获取列数

**使用示例：**
```lua
-- 调用 imgui.getColumnCount 方法
imgui.getColumnCount();
```

---

### imgui.getColumnOffset
获取列偏移

**使用示例：**
```lua
-- 调用 imgui.getColumnOffset 方法
imgui.getColumnOffset();
```

---

### imgui.setColumnOffset
设置列偏移

**使用示例：**
```lua
-- 调用 imgui.setColumnOffset 方法
imgui.setColumnOffset();
```

---

### imgui.getColumnWidth
获取列宽度

**使用示例：**
```lua
-- 调用 imgui.getColumnWidth 方法
imgui.getColumnWidth();
```

---

### imgui.setColumnWidth
设置列宽度

**使用示例：**
```lua
-- 调用 imgui.setColumnWidth 方法
imgui.setColumnWidth();
```

---

### imgui.pushIDStr
推入ID（字符串）

**使用示例：**
```lua
-- 调用 imgui.pushIDStr 方法
imgui.pushIDStr();
```

---

### imgui.pushIDInt
推入ID（整数）

**使用示例：**
```lua
-- 调用 imgui.pushIDInt 方法
imgui.pushIDInt();
```

---

### imgui.pushIDPtr
推入ID（指针）

**使用示例：**
```lua
-- 调用 imgui.pushIDPtr 方法
imgui.pushIDPtr();
```

---

### imgui.popID
弹出ID

**使用示例：**
```lua
-- 调用 imgui.popID 方法
imgui.popID();
```

---

### imgui.getID
获取ID

**使用示例：**
```lua
-- 调用 imgui.getID 方法
imgui.getID();
```

---

### imgui.getIDFromInt
从整数获取ID

**使用示例：**
```lua
-- 调用 imgui.getIDFromInt 方法
imgui.getIDFromInt();
```

---

### imgui.getIDFromPtr
从指针获取ID

**使用示例：**
```lua
-- 调用 imgui.getIDFromPtr 方法
imgui.getIDFromPtr();
```

---

### imgui.text
文本

**使用示例：**
```lua
-- 调用 imgui.text 方法
imgui.text();
```

---

### imgui.textColored
彩色文本

**使用示例：**
```lua
-- 调用 imgui.textColored 方法
imgui.textColored();
```

---

### imgui.textDisabled
禁用文本

**使用示例：**
```lua
-- 调用 imgui.textDisabled 方法
imgui.textDisabled();
```

---

### imgui.textWrapped
换行文本

**使用示例：**
```lua
-- 调用 imgui.textWrapped 方法
imgui.textWrapped();
```

---

### imgui.textUnformatted
无格式文本

**使用示例：**
```lua
-- 调用 imgui.textUnformatted 方法
imgui.textUnformatted();
```

---

### imgui.labelText
标签文本

**使用示例：**
```lua
-- 调用 imgui.labelText 方法
imgui.labelText();
```

---

### imgui.bullet
项目符号

**使用示例：**
```lua
-- 调用 imgui.bullet 方法
imgui.bullet();
```

---

### imgui.bulletText
项目符号文本

**使用示例：**
```lua
-- 调用 imgui.bulletText 方法
imgui.bulletText();
```

---

### imgui.separator
分隔符

**使用示例：**
```lua
-- 调用 imgui.separator 方法
imgui.separator();
```

---

### imgui.separatorText
分隔符文本

**使用示例：**
```lua
-- 调用 imgui.separatorText 方法
imgui.separatorText();
```

---

### imgui.sameLine
同一行

**使用示例：**
```lua
-- 调用 imgui.sameLine 方法
imgui.sameLine();
```

---

### imgui.newLine
新行

**使用示例：**
```lua
-- 调用 imgui.newLine 方法
imgui.newLine();
```

---

### imgui.spacing
间距

**使用示例：**
```lua
-- 调用 imgui.spacing 方法
imgui.spacing();
```

---

### imgui.dummy
虚拟元素

**使用示例：**
```lua
-- 调用 imgui.dummy 方法
imgui.dummy();
```

---

### imgui.beginGroup
开始组

**使用示例：**
```lua
-- 调用 imgui.beginGroup 方法
imgui.beginGroup();
```

---

### imgui.endGroup
结束组

**使用示例：**
```lua
-- 调用 imgui.endGroup 方法
imgui.endGroup();
```

---

### imgui.getCursorPos
获取光标位置

**使用示例：**
```lua
-- 调用 imgui.getCursorPos 方法
imgui.getCursorPos();
```

---

### imgui.getCursorPosX
获取光标X位置

**使用示例：**
```lua
-- 调用 imgui.getCursorPosX 方法
imgui.getCursorPosX();
```

---

### imgui.getCursorPosY
获取光标Y位置

**使用示例：**
```lua
-- 调用 imgui.getCursorPosY 方法
imgui.getCursorPosY();
```

---

### imgui.setCursorPos
设置光标位置

**使用示例：**
```lua
-- 调用 imgui.setCursorPos 方法
imgui.setCursorPos();
```

---

### imgui.setCursorPosX
设置光标X位置

**使用示例：**
```lua
-- 调用 imgui.setCursorPosX 方法
imgui.setCursorPosX();
```

---

### imgui.setCursorPosY
设置光标Y位置

**使用示例：**
```lua
-- 调用 imgui.setCursorPosY 方法
imgui.setCursorPosY();
```

---

### imgui.getCursorStartPos
获取光标起始位置

**使用示例：**
```lua
-- 调用 imgui.getCursorStartPos 方法
imgui.getCursorStartPos();
```

---

### imgui.getCursorScreenPos
获取屏幕光标位置

**使用示例：**
```lua
-- 调用 imgui.getCursorScreenPos 方法
imgui.getCursorScreenPos();
```

---

### imgui.setCursorScreenPos
设置屏幕光标位置

**使用示例：**
```lua
-- 调用 imgui.setCursorScreenPos 方法
imgui.setCursorScreenPos();
```

---

### imgui.getTextLineHeight
获取文本行高度

**使用示例：**
```lua
-- 调用 imgui.getTextLineHeight 方法
imgui.getTextLineHeight();
```

---

### imgui.getTextLineHeightWithSpacing
获取带间距的文本行高度

**使用示例：**
```lua
-- 调用 imgui.getTextLineHeightWithSpacing 方法
imgui.getTextLineHeightWithSpacing();
```

---

### imgui.getFrameHeight
获取框架高度

**使用示例：**
```lua
-- 调用 imgui.getFrameHeight 方法
imgui.getFrameHeight();
```

---

### imgui.getFrameHeightWithSpacing
获取带间距的框架高度

**使用示例：**
```lua
-- 调用 imgui.getFrameHeightWithSpacing 方法
imgui.getFrameHeightWithSpacing();
```

---

### imgui.getContentRegionMax
获取内容区域最大值

**使用示例：**
```lua
-- 调用 imgui.getContentRegionMax 方法
imgui.getContentRegionMax();
```

---

### imgui.getContentRegionAvail
获取可用内容区域

**使用示例：**
```lua
-- 调用 imgui.getContentRegionAvail 方法
imgui.getContentRegionAvail();
```

---

### imgui.popStyleColor
弹出样式颜色

**使用示例：**
```lua
-- 调用 imgui.popStyleColor 方法
imgui.popStyleColor();
```

---

### imgui.getStyleColor
获取样式颜色

**使用示例：**
```lua
-- 调用 imgui.getStyleColor 方法
imgui.getStyleColor();
```

---

### imgui.getStyleColorName
获取样式颜色名称

**使用示例：**
```lua
-- 调用 imgui.getStyleColorName 方法
imgui.getStyleColorName();
```

---

### imgui.getStyleVarInfo
获取样式变量信息

**使用示例：**
```lua
-- 调用 imgui.getStyleVarInfo 方法
imgui.getStyleVarInfo();
```

---

### imgui.beginMainMenuBar
开始主菜单栏

**使用示例：**
```lua
-- 调用 imgui.beginMainMenuBar 方法
imgui.beginMainMenuBar();
```

---

### imgui.endMainMenuBar
结束主菜单栏

**使用示例：**
```lua
-- 调用 imgui.endMainMenuBar 方法
imgui.endMainMenuBar();
```

---

### imgui.beginMenuBar
开始菜单栏

**使用示例：**
```lua
-- 调用 imgui.beginMenuBar 方法
imgui.beginMenuBar();
```

---

### imgui.endMenuBar
结束菜单栏

**使用示例：**
```lua
-- 调用 imgui.endMenuBar 方法
imgui.endMenuBar();
```

---

### imgui.beginMenu
开始菜单

**使用示例：**
```lua
-- 调用 imgui.beginMenu 方法
imgui.beginMenu();
```

---

### imgui.beginMenuDisabled
开始禁用菜单

**使用示例：**
```lua
-- 调用 imgui.beginMenuDisabled 方法
imgui.beginMenuDisabled();
```

---

### imgui.endMenu
结束菜单

**使用示例：**
```lua
-- 调用 imgui.endMenu 方法
imgui.endMenu();
```

---

### imgui.beginTabBar
开始标签栏

**使用示例：**
```lua
-- 调用 imgui.beginTabBar 方法
imgui.beginTabBar();
```

---

### imgui.beginTabBarWithFlags
开始标签栏（带标志）

**使用示例：**
```lua
-- 调用 imgui.beginTabBarWithFlags 方法
imgui.beginTabBarWithFlags();
```

---

### imgui.endTabBar
结束标签栏

**使用示例：**
```lua
-- 调用 imgui.endTabBar 方法
imgui.endTabBar();
```

---

### imgui.beginTabItem
开始标签项

**使用示例：**
```lua
-- 调用 imgui.beginTabItem 方法
imgui.beginTabItem();
```

---

### imgui.beginTabItemWithOpen
开始标签项（带打开状态）

**使用示例：**
```lua
-- 调用 imgui.beginTabItemWithOpen 方法
imgui.beginTabItemWithOpen();
```

---

### imgui.beginTabItemWithFlags
开始标签项（带标志）

**使用示例：**
```lua
-- 调用 imgui.beginTabItemWithFlags 方法
imgui.beginTabItemWithFlags();
```

---

### imgui.endTabItem
结束标签项

**使用示例：**
```lua
-- 调用 imgui.endTabItem 方法
imgui.endTabItem();
```

---

### imgui.setTabItemClosed
设置标签项关闭

**使用示例：**
```lua
-- 调用 imgui.setTabItemClosed 方法
imgui.setTabItemClosed();
```

---

### imgui.isItemHovered
项是否悬停

**使用示例：**
```lua
-- 调用 imgui.isItemHovered 方法
imgui.isItemHovered();
```

---

### imgui.isItemActive
项是否激活

**使用示例：**
```lua
-- 调用 imgui.isItemActive 方法
imgui.isItemActive();
```

---

### imgui.isItemFocused
项是否聚焦

**使用示例：**
```lua
-- 调用 imgui.isItemFocused 方法
imgui.isItemFocused();
```

---

### imgui.isItemClicked
项是否被点击

**使用示例：**
```lua
-- 调用 imgui.isItemClicked 方法
imgui.isItemClicked();
```

---

### imgui.isItemVisible
项是否可见

**使用示例：**
```lua
-- 调用 imgui.isItemVisible 方法
imgui.isItemVisible();
```

---

### imgui.isItemEdited
项是否被编辑

**使用示例：**
```lua
-- 调用 imgui.isItemEdited 方法
imgui.isItemEdited();
```

---

### imgui.isItemActivated
项是否被激活

**使用示例：**
```lua
-- 调用 imgui.isItemActivated 方法
imgui.isItemActivated();
```

---

### imgui.isItemDeactivated
项是否被停用

**使用示例：**
```lua
-- 调用 imgui.isItemDeactivated 方法
imgui.isItemDeactivated();
```

---

### imgui.isItemDeactivatedAfterEdit
项是否在编辑后被停用

**使用示例：**
```lua
-- 调用 imgui.isItemDeactivatedAfterEdit 方法
imgui.isItemDeactivatedAfterEdit();
```

---

### imgui.isItemToggledOpen
项是否切换打开状态

**使用示例：**
```lua
-- 调用 imgui.isItemToggledOpen 方法
imgui.isItemToggledOpen();
```

---

### imgui.isAnyItemHovered
是否有任何项悬停

**使用示例：**
```lua
-- 调用 imgui.isAnyItemHovered 方法
imgui.isAnyItemHovered();
```

---

### imgui.isAnyItemActive
是否有任何项激活

**使用示例：**
```lua
-- 调用 imgui.isAnyItemActive 方法
imgui.isAnyItemActive();
```

---

### imgui.isAnyItemFocused
是否有任何项聚焦

**使用示例：**
```lua
-- 调用 imgui.isAnyItemFocused 方法
imgui.isAnyItemFocused();
```

---

### imgui.getItemRectMin
获取项矩形最小值

**使用示例：**
```lua
-- 调用 imgui.getItemRectMin 方法
imgui.getItemRectMin();
```

---

### imgui.getItemRectMax
获取项矩形最大值

**使用示例：**
```lua
-- 调用 imgui.getItemRectMax 方法
imgui.getItemRectMax();
```

---

### imgui.getItemRectSize
获取项矩形大小

**使用示例：**
```lua
-- 调用 imgui.getItemRectSize 方法
imgui.getItemRectSize();
```

---

### imgui.setItemDefaultFocus
设置项默认焦点

**使用示例：**
```lua
-- 调用 imgui.setItemDefaultFocus 方法
imgui.setItemDefaultFocus();
```

---

### imgui.setKeyboardFocusHere
设置键盘焦点到此处

**使用示例：**
```lua
-- 调用 imgui.setKeyboardFocusHere 方法
imgui.setKeyboardFocusHere();
```

---

### imgui.isMouseHoveringRect
鼠标是否悬停在矩形上

**使用示例：**
```lua
-- 调用 imgui.isMouseHoveringRect 方法
imgui.isMouseHoveringRect();
```

---

### imgui.isMousePosValid
鼠标位置是否有效

**使用示例：**
```lua
-- 调用 imgui.isMousePosValid 方法
imgui.isMousePosValid();
```

---

### imgui.getMousePos
获取鼠标位置

**使用示例：**
```lua
-- 调用 imgui.getMousePos 方法
imgui.getMousePos();
```

---

### imgui.getMousePosOnOpeningCurrentPopup
获取打开当前弹出窗口时的鼠标位置

**使用示例：**
```lua
-- 调用 imgui.getMousePosOnOpeningCurrentPopup 方法
imgui.getMousePosOnOpeningCurrentPopup();
```

---

### imgui.getMouseDragDelta
获取鼠标拖拽增量

**使用示例：**
```lua
-- 调用 imgui.getMouseDragDelta 方法
imgui.getMouseDragDelta();
```

---

### imgui.resetMouseDragDelta
重置鼠标拖拽增量

**使用示例：**
```lua
-- 调用 imgui.resetMouseDragDelta 方法
imgui.resetMouseDragDelta();
```

---

### imgui.getMouseCursor
获取鼠标光标

**使用示例：**
```lua
-- 调用 imgui.getMouseCursor 方法
imgui.getMouseCursor();
```

---

### imgui.setMouseCursor
设置鼠标光标

**使用示例：**
```lua
-- 调用 imgui.setMouseCursor 方法
imgui.setMouseCursor();
```

---

### imgui.getClipboardText
获取剪贴板文本

**使用示例：**
```lua
-- 调用 imgui.getClipboardText 方法
imgui.getClipboardText();
```

---

### imgui.setClipboardText
设置剪贴板文本

**使用示例：**
```lua
-- 调用 imgui.setClipboardText 方法
imgui.setClipboardText();
```

---

### imgui.getTime
获取时间

**使用示例：**
```lua
-- 调用 imgui.getTime 方法
imgui.getTime();
```

---

### imgui.getFrameCount
获取帧数

**使用示例：**
```lua
-- 调用 imgui.getFrameCount 方法
imgui.getFrameCount();
```

---

### imgui.getStyleColorVec4
获取样式颜色Vec4

**使用示例：**
```lua
-- 调用 imgui.getStyleColorVec4 方法
imgui.getStyleColorVec4();
```

---

### imgui.getColorU32
获取颜色U32

**使用示例：**
```lua
-- 调用 imgui.getColorU32 方法
imgui.getColorU32();
```

---

### imgui.getColorU32FromRGBA
从RGBA获取颜色U32

**使用示例：**
```lua
-- 调用 imgui.getColorU32FromRGBA 方法
imgui.getColorU32FromRGBA();
```

---

### imgui.getColorU32FromU32
从U32获取颜色U32

**使用示例：**
```lua
-- 调用 imgui.getColorU32FromU32 方法
imgui.getColorU32FromU32();
```

---

### imgui.getFont
获取字体

**使用示例：**
```lua
-- 调用 imgui.getFont 方法
imgui.getFont();
```

---

### imgui.getFontSize
获取字体大小

**使用示例：**
```lua
-- 调用 imgui.getFontSize 方法
imgui.getFontSize();
```

---

### imgui.getFontTexUvWhitePixel
获取字体纹理UV白色像素

**使用示例：**
```lua
-- 调用 imgui.getFontTexUvWhitePixel 方法
imgui.getFontTexUvWhitePixel();
```

---

### imgui.getColorConvertRGBtoHSV
RGB转HSV

**使用示例：**
```lua
-- 调用 imgui.getColorConvertRGBtoHSV 方法
imgui.getColorConvertRGBtoHSV();
```

---

### imgui.getColorConvertHSVtoRGB
HSV转RGB

**使用示例：**
```lua
-- 调用 imgui.getColorConvertHSVtoRGB 方法
imgui.getColorConvertHSVtoRGB();
```

---

### imgui.getWindowDrawList
获取窗口绘制列表

**使用示例：**
```lua
-- 调用 imgui.getWindowDrawList 方法
imgui.getWindowDrawList();
```

---

### imgui.getBackgroundDrawList
获取背景绘制列表

**使用示例：**
```lua
-- 调用 imgui.getBackgroundDrawList 方法
imgui.getBackgroundDrawList();
```

---

### imgui.getForegroundDrawList
获取前景绘制列表

**使用示例：**
```lua
-- 调用 imgui.getForegroundDrawList 方法
imgui.getForegroundDrawList();
```

---

### imgui.getMainViewport
获取主视口

**使用示例：**
```lua
-- 调用 imgui.getMainViewport 方法
imgui.getMainViewport();
```

---

### imgui.findViewportByID
通过ID查找视口

**使用示例：**
```lua
-- 调用 imgui.findViewportByID 方法
imgui.findViewportByID();
```

---

### imgui.findViewportByPlatformHandle
通过平台句柄查找视口

**使用示例：**
```lua
-- 调用 imgui.findViewportByPlatformHandle 方法
imgui.findViewportByPlatformHandle();
```

---

### imgui.getPlatformIO
获取平台IO

**使用示例：**
```lua
-- 调用 imgui.getPlatformIO 方法
imgui.getPlatformIO();
```

---

### imgui.getIO
获取IO

**使用示例：**
```lua
-- 调用 imgui.getIO 方法
imgui.getIO();
```

---

### imgui.getStyle
获取样式

**使用示例：**
```lua
-- 调用 imgui.getStyle 方法
imgui.getStyle();
```

---

### imgui.getDragDropPayload
获取拖放载荷

**使用示例：**
```lua
-- 调用 imgui.getDragDropPayload 方法
imgui.getDragDropPayload();
```

---

### imgui.beginTooltip
开始工具提示

**使用示例：**
```lua
-- 调用 imgui.beginTooltip 方法
imgui.beginTooltip();
```

---

### imgui.endTooltip
结束工具提示

**使用示例：**
```lua
-- 调用 imgui.endTooltip 方法
imgui.endTooltip();
```

---

### imgui.setTooltip
设置工具提示

**使用示例：**
```lua
-- 调用 imgui.setTooltip 方法
imgui.setTooltip();
```

---

### imgui.beginItemTooltip
开始项工具提示

**使用示例：**
```lua
-- 调用 imgui.beginItemTooltip 方法
imgui.beginItemTooltip();
```

---

### imgui.beginPopup
开始弹出窗口

**使用示例：**
```lua
-- 调用 imgui.beginPopup 方法
imgui.beginPopup();
```

---

### imgui.beginPopupWithFlags
开始弹出窗口（带标志）

**使用示例：**
```lua
-- 调用 imgui.beginPopupWithFlags 方法
imgui.beginPopupWithFlags();
```

---

### imgui.endPopup
结束弹出窗口

**使用示例：**
```lua
-- 调用 imgui.endPopup 方法
imgui.endPopup();
```

---

### imgui.openPopup
打开弹出窗口

**使用示例：**
```lua
-- 调用 imgui.openPopup 方法
imgui.openPopup();
```

---

### imgui.openPopupWithFlags
打开弹出窗口（带标志）

**使用示例：**
```lua
-- 调用 imgui.openPopupWithFlags 方法
imgui.openPopupWithFlags();
```

---

### imgui.closeCurrentPopup
关闭当前弹出窗口

**使用示例：**
```lua
-- 调用 imgui.closeCurrentPopup 方法
imgui.closeCurrentPopup();
```

---

### imgui.beginPopupModal
开始模态弹出窗口

**使用示例：**
```lua
-- 调用 imgui.beginPopupModal 方法
imgui.beginPopupModal();
```

---

### imgui.beginPopupModalWithOpen
开始模态弹出窗口（带打开状态）

**使用示例：**
```lua
-- 调用 imgui.beginPopupModalWithOpen 方法
imgui.beginPopupModalWithOpen();
```

---

### imgui.beginPopupModalWithFlags
开始模态弹出窗口（带标志）

**使用示例：**
```lua
-- 调用 imgui.beginPopupModalWithFlags 方法
imgui.beginPopupModalWithFlags();
```

---

### imgui.beginPopupContextItem
开始项上下文弹出窗口

**使用示例：**
```lua
-- 调用 imgui.beginPopupContextItem 方法
imgui.beginPopupContextItem();
```

---

### imgui.beginPopupContextWindow
开始窗口上下文弹出窗口

**使用示例：**
```lua
-- 调用 imgui.beginPopupContextWindow 方法
imgui.beginPopupContextWindow();
```

---

### imgui.beginPopupContextVoid
开始空白上下文弹出窗口

**使用示例：**
```lua
-- 调用 imgui.beginPopupContextVoid 方法
imgui.beginPopupContextVoid();
```

---

### imgui.beginColumns
开始列

**使用示例：**
```lua
-- 调用 imgui.beginColumns 方法
imgui.beginColumns();
```

---

### imgui.getColumnIndex
获取列索引

**使用示例：**
```lua
-- 调用 imgui.getColumnIndex 方法
imgui.getColumnIndex();
```

---

### imgui.getColumnCount
获取列数

**使用示例：**
```lua
-- 调用 imgui.getColumnCount 方法
imgui.getColumnCount();
```

---

### imgui.getColumnOffset
获取列偏移

**使用示例：**
```lua
-- 调用 imgui.getColumnOffset 方法
imgui.getColumnOffset();
```

---

### imgui.setColumnOffset
设置列偏移

**使用示例：**
```lua
-- 调用 imgui.setColumnOffset 方法
imgui.setColumnOffset();
```

---

### imgui.getColumnWidth
获取列宽度

**使用示例：**
```lua
-- 调用 imgui.getColumnWidth 方法
imgui.getColumnWidth();
```

---

### imgui.setColumnWidth
设置列宽度

**使用示例：**
```lua
-- 调用 imgui.setColumnWidth 方法
imgui.setColumnWidth();
```

---

### imgui.pushIDStr
推入ID（字符串）

**使用示例：**
```lua
-- 调用 imgui.pushIDStr 方法
imgui.pushIDStr();
```

---

### imgui.pushIDInt
推入ID（整数）

**使用示例：**
```lua
-- 调用 imgui.pushIDInt 方法
imgui.pushIDInt();
```

---

### imgui.pushIDPtr
推入ID（指针）

**使用示例：**
```lua
-- 调用 imgui.pushIDPtr 方法
imgui.pushIDPtr();
```

---

### imgui.popID
弹出ID

**使用示例：**
```lua
-- 调用 imgui.popID 方法
imgui.popID();
```

---

### imgui.getID
获取ID

**使用示例：**
```lua
-- 调用 imgui.getID 方法
imgui.getID();
```

---

### imgui.getIDFromInt
从整数获取ID

**使用示例：**
```lua
-- 调用 imgui.getIDFromInt 方法
imgui.getIDFromInt();
```

---

### imgui.getIDFromPtr
从指针获取ID

**使用示例：**
```lua
-- 调用 imgui.getIDFromPtr 方法
imgui.getIDFromPtr();
```

---

### imgui.text
文本

**使用示例：**
```lua
-- 调用 imgui.text 方法
imgui.text();
```

---

### imgui.textColored
彩色文本

**使用示例：**
```lua
-- 调用 imgui.textColored 方法
imgui.textColored();
```

---

### imgui.textDisabled
禁用文本

**使用示例：**
```lua
-- 调用 imgui.textDisabled 方法
imgui.textDisabled();
```

---

### imgui.textWrapped
换行文本

**使用示例：**
```lua
-- 调用 imgui.textWrapped 方法
imgui.textWrapped();
```

---

### imgui.textUnformatted
无格式文本

**使用示例：**
```lua
-- 调用 imgui.textUnformatted 方法
imgui.textUnformatted();
```

---

### imgui.labelText
标签文本

**使用示例：**
```lua
-- 调用 imgui.labelText 方法
imgui.labelText();
```

---

### imgui.bullet
项目符号

**使用示例：**
```lua
-- 调用 imgui.bullet 方法
imgui.bullet();
```

---

### imgui.bulletText
项目符号文本

**使用示例：**
```lua
-- 调用 imgui.bulletText 方法
imgui.bulletText();
```

---

### imgui.separator
分隔符

**使用示例：**
```lua
-- 调用 imgui.separator 方法
imgui.separator();
```

---

### imgui.separatorText
分隔符文本

**使用示例：**
```lua
-- 调用 imgui.separatorText 方法
imgui.separatorText();
```

---

### imgui.sameLine
同一行

**使用示例：**
```lua
-- 调用 imgui.sameLine 方法
imgui.sameLine();
```

---

### imgui.image
绘制图像

**使用示例：**
```lua
-- 调用 imgui.image 方法
imgui.image();
```

---

### imgui.imageWithUV
绘制图像（带UV）

**使用示例：**
```lua
-- 调用 imgui.imageWithUV 方法
imgui.imageWithUV();
```

---

### imgui.imageButton
图像按钮

**使用示例：**
```lua
-- 调用 imgui.imageButton 方法
imgui.imageButton();
```

---

### imgui.imageButtonWithUV
图像按钮（带UV）

**使用示例：**
```lua
-- 调用 imgui.imageButtonWithUV 方法
imgui.imageButtonWithUV();
```

---

### imgui.imageWithBg
绘制图像（带背景）

**使用示例：**
```lua
-- 调用 imgui.imageWithBg 方法
imgui.imageWithBg();
```

---

### imgui.dockSpace
停靠空间

**使用示例：**
```lua
-- 调用 imgui.dockSpace 方法
imgui.dockSpace();
```

---

### imgui.dockSpaceWithSize
停靠空间（带大小）

**使用示例：**
```lua
-- 调用 imgui.dockSpaceWithSize 方法
imgui.dockSpaceWithSize();
```

---

### imgui.dockSpaceWithFlags
停靠空间（带标志）

**使用示例：**
```lua
-- 调用 imgui.dockSpaceWithFlags 方法
imgui.dockSpaceWithFlags();
```

---

### imgui.dockSpaceOverViewport
视口停靠空间

**使用示例：**
```lua
-- 调用 imgui.dockSpaceOverViewport 方法
imgui.dockSpaceOverViewport();
```

---

### imgui.dockSpaceOverViewportWithID
视口停靠空间（带ID）

**使用示例：**
```lua
-- 调用 imgui.dockSpaceOverViewportWithID 方法
imgui.dockSpaceOverViewportWithID();
```

---

### imgui.setNextWindowDockID
设置下一个窗口停靠ID

**使用示例：**
```lua
-- 调用 imgui.setNextWindowDockID 方法
imgui.setNextWindowDockID();
```

---

### imgui.setNextWindowDockIDWithCond
设置下一个窗口停靠ID（带条件）

**使用示例：**
```lua
-- 调用 imgui.setNextWindowDockIDWithCond 方法
imgui.setNextWindowDockIDWithCond();
```

---

### imgui.setNextWindowClass
设置下一个窗口类

**使用示例：**
```lua
-- 调用 imgui.setNextWindowClass 方法
imgui.setNextWindowClass();
```

---

### imgui.listBox
列表框

**使用示例：**
```lua
-- 调用 imgui.listBox 方法
imgui.listBox();
```

---

### imgui.listBoxWithHeight
列表框（带高度）

**使用示例：**
```lua
-- 调用 imgui.listBoxWithHeight 方法
imgui.listBoxWithHeight();
```

---

### imgui.tableSetupScrollFreeze
设置表格滚动冻结

**使用示例：**
```lua
-- 调用 imgui.tableSetupScrollFreeze 方法
imgui.tableSetupScrollFreeze();
```

---

### imgui.tableGetSortSpecs
获取表格排序规范

**使用示例：**
```lua
-- 调用 imgui.tableGetSortSpecs 方法
imgui.tableGetSortSpecs();
```

---

### imgui.tableGetColumnName
获取表格列名

**使用示例：**
```lua
-- 调用 imgui.tableGetColumnName 方法
imgui.tableGetColumnName();
```

---

### imgui.tableGetColumnNameByIndex
获取表格列名（按索引）

**使用示例：**
```lua
-- 调用 imgui.tableGetColumnNameByIndex 方法
imgui.tableGetColumnNameByIndex();
```

---

### imgui.tableSetColumnEnabled
设置表格列启用

**使用示例：**
```lua
-- 调用 imgui.tableSetColumnEnabled 方法
imgui.tableSetColumnEnabled();
```

---

### imgui.tableSetBgColor
设置表格背景色

**使用示例：**
```lua
-- 调用 imgui.tableSetBgColor 方法
imgui.tableSetBgColor();
```

---

### imgui.tableSetBgColorByColumn
设置表格背景色（按列）

**使用示例：**
```lua
-- 调用 imgui.tableSetBgColorByColumn 方法
imgui.tableSetBgColorByColumn();
```

---

### imgui.invisibleButton
不可见按钮

**使用示例：**
```lua
-- 调用 imgui.invisibleButton 方法
imgui.invisibleButton();
```

---

### imgui.pushTextWrapPos
推入文本换行位置

**使用示例：**
```lua
-- 调用 imgui.pushTextWrapPos 方法
imgui.pushTextWrapPos();
```

---

### imgui.pushTextWrapPosWithPos
推入文本换行位置（带位置）

**使用示例：**
```lua
-- 调用 imgui.pushTextWrapPosWithPos 方法
imgui.pushTextWrapPosWithPos();
```

---

### imgui.popTextWrapPos
弹出文本换行位置

**使用示例：**
```lua
-- 调用 imgui.popTextWrapPos 方法
imgui.popTextWrapPos();
```

---

### imgui.pushItemWidth
推入项宽度

**使用示例：**
```lua
-- 调用 imgui.pushItemWidth 方法
imgui.pushItemWidth();
```

---

### imgui.popItemWidth
弹出项宽度

**使用示例：**
```lua
-- 调用 imgui.popItemWidth 方法
imgui.popItemWidth();
```

---

### imgui.pushClipRect
推入剪裁矩形

**使用示例：**
```lua
-- 调用 imgui.pushClipRect 方法
imgui.pushClipRect();
```

---

### imgui.popClipRect
弹出剪裁矩形

**使用示例：**
```lua
-- 调用 imgui.popClipRect 方法
imgui.popClipRect();
```

---

### imgui.pushStyleVarX
推入样式变量（X）

**使用示例：**
```lua
-- 调用 imgui.pushStyleVarX 方法
imgui.pushStyleVarX();
```

---

### imgui.pushStyleVarY
推入样式变量（Y）

**使用示例：**
```lua
-- 调用 imgui.pushStyleVarY 方法
imgui.pushStyleVarY();
```

---

### imgui.pushStyleVarFloat
推入样式变量（浮点）

**使用示例：**
```lua
-- 调用 imgui.pushStyleVarFloat 方法
imgui.pushStyleVarFloat();
```

---

### imgui.pushStyleVarVec2
推入样式变量（Vec2）

**使用示例：**
```lua
-- 调用 imgui.pushStyleVarVec2 方法
imgui.pushStyleVarVec2();
```

---

### imgui.popStyleVar
弹出样式变量

**使用示例：**
```lua
-- 调用 imgui.popStyleVar 方法
imgui.popStyleVar();
```

---

### imgui.popStyleVarCount
弹出样式变量（带数量）

**使用示例：**
```lua
-- 调用 imgui.popStyleVarCount 方法
imgui.popStyleVarCount();
```

---

### imgui.loadIniSettingsFromDisk
从磁盘加载INI设置

**使用示例：**
```lua
-- 调用 imgui.loadIniSettingsFromDisk 方法
imgui.loadIniSettingsFromDisk();
```

---

### imgui.loadIniSettingsFromMemory
从内存加载INI设置

**使用示例：**
```lua
-- 调用 imgui.loadIniSettingsFromMemory 方法
imgui.loadIniSettingsFromMemory();
```

---

### imgui.saveIniSettingsToDisk
保存INI设置到磁盘

**使用示例：**
```lua
-- 调用 imgui.saveIniSettingsToDisk 方法
imgui.saveIniSettingsToDisk();
```

---

### imgui.saveIniSettingsToMemory
保存INI设置到内存

**使用示例：**
```lua
-- 调用 imgui.saveIniSettingsToMemory 方法
imgui.saveIniSettingsToMemory();
```

---

### imgui.setClipboardText
设置剪贴板文本

**使用示例：**
```lua
-- 调用 imgui.setClipboardText 方法
imgui.setClipboardText();
```

---

### imgui.pushStyleColorU32
推入样式颜色（U32）

**使用示例：**
```lua
-- 调用 imgui.pushStyleColorU32 方法
imgui.pushStyleColorU32();
```

---

### imgui.pushStyleColorVec4
推入样式颜色（Vec4）

**使用示例：**
```lua
-- 调用 imgui.pushStyleColorVec4 方法
imgui.pushStyleColorVec4();
```

---

### imgui.popStyleColor
弹出样式颜色

**使用示例：**
```lua
-- 调用 imgui.popStyleColor 方法
imgui.popStyleColor();
```

---

### imgui.popStyleColorCount
弹出样式颜色（带数量）

**使用示例：**
```lua
-- 调用 imgui.popStyleColorCount 方法
imgui.popStyleColorCount();
```

---

### imgui.inputTextWithHint
带提示的文本输入

**使用示例：**
```lua
-- 调用 imgui.inputTextWithHint 方法
imgui.inputTextWithHint();
```

---

### imgui.inputTextMultiline
多行文本输入

**使用示例：**
```lua
-- 调用 imgui.inputTextMultiline 方法
imgui.inputTextMultiline();
```

---

### imgui.treeNodeEx
扩展树节点

**使用示例：**
```lua
-- 调用 imgui.treeNodeEx 方法
imgui.treeNodeEx();
```

---

### imgui.treeNodeStr
树节点（字符串）

**使用示例：**
```lua
-- 调用 imgui.treeNodeStr 方法
imgui.treeNodeStr();
```

---

### imgui.treeNodeStrStr
树节点（双字符串）

**使用示例：**
```lua
-- 调用 imgui.treeNodeStrStr 方法
imgui.treeNodeStrStr();
```

---

### imgui.treeNodePtr
树节点（指针）

**使用示例：**
```lua
-- 调用 imgui.treeNodePtr 方法
imgui.treeNodePtr();
```

---

### imgui.setKeyboardFocusHere
设置键盘焦点到此处

**使用示例：**
```lua
-- 调用 imgui.setKeyboardFocusHere 方法
imgui.setKeyboardFocusHere();
```

---

### imgui.setKeyboardFocusHereWithOffset
设置键盘焦点到此处（带偏移）

**使用示例：**
```lua
-- 调用 imgui.setKeyboardFocusHereWithOffset 方法
imgui.setKeyboardFocusHereWithOffset();
```

---

### imgui.setItemDefaultFocus
设置项默认焦点

**使用示例：**
```lua
-- 调用 imgui.setItemDefaultFocus 方法
imgui.setItemDefaultFocus();
```

---

### imgui.setNextItemAllowOverlap
设置下一项允许重叠

**使用示例：**
```lua
-- 调用 imgui.setNextItemAllowOverlap 方法
imgui.setNextItemAllowOverlap();
```

---

### imgui.setNextItemOpen
设置下一项打开状态

**使用示例：**
```lua
-- 调用 imgui.setNextItemOpen 方法
imgui.setNextItemOpen();
```

---

### imgui.setNextItemOpenWithCond
设置下一项打开状态（带条件）

**使用示例：**
```lua
-- 调用 imgui.setNextItemOpenWithCond 方法
imgui.setNextItemOpenWithCond();
```

---

### imgui.setScrollHereX
设置滚动到此处 X

**使用示例：**
```lua
-- 调用 imgui.setScrollHereX 方法
imgui.setScrollHereX();
```

---

### imgui.setScrollHereXWithRatio
设置滚动到此处 X（带比例）

**使用示例：**
```lua
-- 调用 imgui.setScrollHereXWithRatio 方法
imgui.setScrollHereXWithRatio();
```

---

### imgui.setScrollHereY
设置滚动到此处 Y

**使用示例：**
```lua
-- 调用 imgui.setScrollHereY 方法
imgui.setScrollHereY();
```

---

### imgui.setScrollHereYWithRatio
设置滚动到此处 Y（带比例）

**使用示例：**
```lua
-- 调用 imgui.setScrollHereYWithRatio 方法
imgui.setScrollHereYWithRatio();
```

---

### imgui.setScrollFromPosX
从位置 X 滚动

**使用示例：**
```lua
-- 调用 imgui.setScrollFromPosX 方法
imgui.setScrollFromPosX();
```

---

### imgui.setScrollFromPosXWithRatio
从位置 X 滚动（带比例）

**使用示例：**
```lua
-- 调用 imgui.setScrollFromPosXWithRatio 方法
imgui.setScrollFromPosXWithRatio();
```

---

### imgui.setScrollFromPosY
从位置 Y 滚动

**使用示例：**
```lua
-- 调用 imgui.setScrollFromPosY 方法
imgui.setScrollFromPosY();
```

---

### imgui.setScrollFromPosYWithRatio
从位置 Y 滚动（带比例）

**使用示例：**
```lua
-- 调用 imgui.setScrollFromPosYWithRatio 方法
imgui.setScrollFromPosYWithRatio();
```

---

### imgui.openPopupOnItemClick
点击时打开弹出窗口

**使用示例：**
```lua
-- 调用 imgui.openPopupOnItemClick 方法
imgui.openPopupOnItemClick();
```

---

### imgui.openPopupOnItemClickWithID
点击时打开弹出窗口（带ID）

**使用示例：**
```lua
-- 调用 imgui.openPopupOnItemClickWithID 方法
imgui.openPopupOnItemClickWithID();
```

---

### imgui.openPopupID
打开弹出窗口（ID）

**使用示例：**
```lua
-- 调用 imgui.openPopupID 方法
imgui.openPopupID();
```

---

### imgui.openPopupIDWithFlags
打开弹出窗口（ID，带标志）

**使用示例：**
```lua
-- 调用 imgui.openPopupIDWithFlags 方法
imgui.openPopupIDWithFlags();
```

---

### imgui.beginTooltip
开始工具提示

**使用示例：**
```lua
-- 调用 imgui.beginTooltip 方法
imgui.beginTooltip();
```

---

### imgui.endTooltip
结束工具提示

**使用示例：**
```lua
-- 调用 imgui.endTooltip 方法
imgui.endTooltip();
```

---

### imgui.beginItemTooltip
开始项工具提示

**使用示例：**
```lua
-- 调用 imgui.beginItemTooltip 方法
imgui.beginItemTooltip();
```

---

### imgui.setTooltip
设置工具提示

**使用示例：**
```lua
-- 调用 imgui.setTooltip 方法
imgui.setTooltip();
```

---

### imgui.calcTextSize
计算文本大小

**使用示例：**
```lua
-- 调用 imgui.calcTextSize 方法
imgui.calcTextSize();
```

---

### imgui.calcTextSizeWithWrap
计算文本大小（带换行）

**使用示例：**
```lua
-- 调用 imgui.calcTextSizeWithWrap 方法
imgui.calcTextSizeWithWrap();
```

---

### imgui.calcItemWidth
计算项宽度

**使用示例：**
```lua
-- 调用 imgui.calcItemWidth 方法
imgui.calcItemWidth();
```

---

### imgui.progressBar
进度条

**使用示例：**
```lua
-- 调用 imgui.progressBar 方法
imgui.progressBar();
```

---

### imgui.progressBarWithSize
进度条（带大小）

**使用示例：**
```lua
-- 调用 imgui.progressBarWithSize 方法
imgui.progressBarWithSize();
```

---

### imgui.progressBarWithOverlay
进度条（带覆盖）

**使用示例：**
```lua
-- 调用 imgui.progressBarWithOverlay 方法
imgui.progressBarWithOverlay();
```

---

### imgui.setScrollX
设置滚动 X

**使用示例：**
```lua
-- 调用 imgui.setScrollX 方法
imgui.setScrollX();
```

---

### imgui.setScrollY
设置滚动 Y

**使用示例：**
```lua
-- 调用 imgui.setScrollY 方法
imgui.setScrollY();
```

---

### imgui.setWindowPos
设置窗口位置

**使用示例：**
```lua
-- 调用 imgui.setWindowPos 方法
imgui.setWindowPos();
```

---

### imgui.setWindowPosWithCond
设置窗口位置（带条件）

**使用示例：**
```lua
-- 调用 imgui.setWindowPosWithCond 方法
imgui.setWindowPosWithCond();
```

---

### imgui.setWindowPosByName
设置窗口位置（按名称）

**使用示例：**
```lua
-- 调用 imgui.setWindowPosByName 方法
imgui.setWindowPosByName();
```

---

### imgui.setWindowPosByNameWithCond
设置窗口位置（按名称，带条件）

**使用示例：**
```lua
-- 调用 imgui.setWindowPosByNameWithCond 方法
imgui.setWindowPosByNameWithCond();
```

---

### imgui.setWindowSize
设置窗口大小

**使用示例：**
```lua
-- 调用 imgui.setWindowSize 方法
imgui.setWindowSize();
```

---

### imgui.setWindowSizeWithCond
设置窗口大小（带条件）

**使用示例：**
```lua
-- 调用 imgui.setWindowSizeWithCond 方法
imgui.setWindowSizeWithCond();
```

---

### imgui.setWindowSizeByName
设置窗口大小（按名称）

**使用示例：**
```lua
-- 调用 imgui.setWindowSizeByName 方法
imgui.setWindowSizeByName();
```

---

### imgui.setWindowSizeByNameWithCond
设置窗口大小（按名称，带条件）

**使用示例：**
```lua
-- 调用 imgui.setWindowSizeByNameWithCond 方法
imgui.setWindowSizeByNameWithCond();
```

---

### imgui.setWindowCollapsed
设置窗口折叠状态

**使用示例：**
```lua
-- 调用 imgui.setWindowCollapsed 方法
imgui.setWindowCollapsed();
```

---

### imgui.setWindowCollapsedWithCond
设置窗口折叠状态（带条件）

**使用示例：**
```lua
-- 调用 imgui.setWindowCollapsedWithCond 方法
imgui.setWindowCollapsedWithCond();
```

---

### imgui.setWindowCollapsedByName
设置窗口折叠状态（按名称）

**使用示例：**
```lua
-- 调用 imgui.setWindowCollapsedByName 方法
imgui.setWindowCollapsedByName();
```

---

### imgui.setWindowCollapsedByNameWithCond
设置窗口折叠状态（按名称，带条件）

**使用示例：**
```lua
-- 调用 imgui.setWindowCollapsedByNameWithCond 方法
imgui.setWindowCollapsedByNameWithCond();
```

---

### imgui.setWindowFocus
设置窗口焦点

**使用示例：**
```lua
-- 调用 imgui.setWindowFocus 方法
imgui.setWindowFocus();
```

---

### imgui.setWindowFocusByName
设置窗口焦点（按名称）

**使用示例：**
```lua
-- 调用 imgui.setWindowFocusByName 方法
imgui.setWindowFocusByName();
```

---

### imgui.setWindowFontScale
设置窗口字体缩放

**使用示例：**
```lua
-- 调用 imgui.setWindowFontScale 方法
imgui.setWindowFontScale();
```

---

### imgui.beginDisabled
开始禁用

**使用示例：**
```lua
-- 调用 imgui.beginDisabled 方法
imgui.beginDisabled();
```

---

### imgui.columns
列

**使用示例：**
```lua
-- 调用 imgui.columns 方法
imgui.columns();
```

---

### imgui.debugFlashStyleColor
调试闪烁样式颜色

**使用示例：**
```lua
-- 调用 imgui.debugFlashStyleColor 方法
imgui.debugFlashStyleColor();
```

---

### imgui.debugLog
调试日志

**使用示例：**
```lua
-- 调用 imgui.debugLog 方法
imgui.debugLog();
```

---

### imgui.debugStartItemPicker
调试开始项目选择器

**使用示例：**
```lua
-- 调用 imgui.debugStartItemPicker 方法
imgui.debugStartItemPicker();
```

---

### imgui.debugTextEncoding
调试文本编码

**使用示例：**
```lua
-- 调用 imgui.debugTextEncoding 方法
imgui.debugTextEncoding();
```

---

### imgui.destroyContextV
销毁上下文（带参数）

**使用示例：**
```lua
-- 调用 imgui.destroyContextV 方法
imgui.destroyContextV();
```

---

### imgui.destroyPlatformWindows
销毁平台窗口

**使用示例：**
```lua
-- 调用 imgui.destroyPlatformWindows 方法
imgui.destroyPlatformWindows();
```

---

### imgui.endListBox
结束列表框

**使用示例：**
```lua
-- 调用 imgui.endListBox 方法
imgui.endListBox();
```

---

### imgui.image
图像

**使用示例：**
```lua
-- 调用 imgui.image 方法
imgui.image();
```

---

### imgui.imageWithBg
图像（带背景）

**使用示例：**
```lua
-- 调用 imgui.imageWithBg 方法
imgui.imageWithBg();
```

---

### imgui.loadIniSettingsFromMemoryV
从内存加载INI设置（带参数）

**使用示例：**
```lua
-- 调用 imgui.loadIniSettingsFromMemoryV 方法
imgui.loadIniSettingsFromMemoryV();
```

---

### imgui.logButtons
日志按钮

**使用示例：**
```lua
-- 调用 imgui.logButtons 方法
imgui.logButtons();
```

---

### imgui.logFinish
完成日志

**使用示例：**
```lua
-- 调用 imgui.logFinish 方法
imgui.logFinish();
```

---

### imgui.logText
日志文本

**使用示例：**
```lua
-- 调用 imgui.logText 方法
imgui.logText();
```

---

### imgui.logToClipboard
日志到剪贴板

**使用示例：**
```lua
-- 调用 imgui.logToClipboard 方法
imgui.logToClipboard();
```

---

### imgui.logToClipboardV
日志到剪贴板（带参数）

**使用示例：**
```lua
-- 调用 imgui.logToClipboardV 方法
imgui.logToClipboardV();
```

---

### imgui.logToFile
日志到文件

**使用示例：**
```lua
-- 调用 imgui.logToFile 方法
imgui.logToFile();
```

---

### imgui.logToFileV
日志到文件（带参数）

**使用示例：**
```lua
-- 调用 imgui.logToFileV 方法
imgui.logToFileV();
```

---

### imgui.logToTTY
日志到TTY

**使用示例：**
```lua
-- 调用 imgui.logToTTY 方法
imgui.logToTTY();
```

---

### imgui.logToTTYV
日志到TTY（带参数）

**使用示例：**
```lua
-- 调用 imgui.logToTTYV 方法
imgui.logToTTYV();
```

---

### imgui.memFree
释放内存

**使用示例：**
```lua
-- 调用 imgui.memFree 方法
imgui.memFree();
```

---

### imgui.nextColumn
下一列

**使用示例：**
```lua
-- 调用 imgui.nextColumn 方法
imgui.nextColumn();
```

---

### imgui.openPopupStrV
打开弹出窗口（带参数）

**使用示例：**
```lua
-- 调用 imgui.openPopupStrV 方法
imgui.openPopupStrV();
```

---

### imgui.plotHistogramFloatPtr
绘制直方图（浮点指针）

**使用示例：**
```lua
-- 调用 imgui.plotHistogramFloatPtr 方法
imgui.plotHistogramFloatPtr();
```

---

### imgui.plotHistogramFloatPtrV
绘制直方图（浮点指针，带参数）

**使用示例：**
```lua
-- 调用 imgui.plotHistogramFloatPtrV 方法
imgui.plotHistogramFloatPtrV();
```

---

### imgui.plotLinesFloatPtr
绘制线条（浮点指针）

**使用示例：**
```lua
-- 调用 imgui.plotLinesFloatPtr 方法
imgui.plotLinesFloatPtr();
```

---

### imgui.plotLinesFloatPtrV
绘制线条（浮点指针，带参数）

**使用示例：**
```lua
-- 调用 imgui.plotLinesFloatPtrV 方法
imgui.plotLinesFloatPtrV();
```

---

### imgui.pushIDStrStr
推送ID（字符串范围）

**使用示例：**
```lua
-- 调用 imgui.pushIDStrStr 方法
imgui.pushIDStrStr();
```

---

### imgui.renderPlatformWindowsDefault
渲染平台窗口（默认）

**使用示例：**
```lua
-- 调用 imgui.renderPlatformWindowsDefault 方法
imgui.renderPlatformWindowsDefault();
```

---

### imgui.renderPlatformWindowsDefaultV
渲染平台窗口（默认，带参数）

**使用示例：**
```lua
-- 调用 imgui.renderPlatformWindowsDefaultV 方法
imgui.renderPlatformWindowsDefaultV();
```

---

### imgui.resetMouseDragDeltaV
重置鼠标拖动增量（带参数）

**使用示例：**
```lua
-- 调用 imgui.resetMouseDragDeltaV 方法
imgui.resetMouseDragDeltaV();
```

---

### imgui.sameLineV
同一行（带参数）

**使用示例：**
```lua
-- 调用 imgui.sameLineV 方法
imgui.sameLineV();
```

---

### imgui.setItemTooltip
设置项目工具提示

**使用示例：**
```lua
-- 调用 imgui.setItemTooltip 方法
imgui.setItemTooltip();
```

---

### imgui.setNextItemSelectionUserData
设置下一项选择用户数据

**使用示例：**
```lua
-- 调用 imgui.setNextItemSelectionUserData 方法
imgui.setNextItemSelectionUserData();
```

---

### imgui.setNextItemShortcut
设置下一项快捷键

**使用示例：**
```lua
-- 调用 imgui.setNextItemShortcut 方法
imgui.setNextItemShortcut();
```

---

### imgui.setNextItemShortcutV
设置下一项快捷键（带参数）

**使用示例：**
```lua
-- 调用 imgui.setNextItemShortcutV 方法
imgui.setNextItemShortcutV();
```

---

### imgui.setNextItemStorageID
设置下一项存储ID

**使用示例：**
```lua
-- 调用 imgui.setNextItemStorageID 方法
imgui.setNextItemStorageID();
```

---

### imgui.setNextItemWidth
设置下一项宽度

**使用示例：**
```lua
-- 调用 imgui.setNextItemWidth 方法
imgui.setNextItemWidth();
```

---

### imgui.setNextWindowDockID
设置下一窗口停靠ID

**使用示例：**
```lua
-- 调用 imgui.setNextWindowDockID 方法
imgui.setNextWindowDockID();
```

---

### imgui.setNextWindowPosV
设置下一窗口位置（带参数）

**使用示例：**
```lua
-- 调用 imgui.setNextWindowPosV 方法
imgui.setNextWindowPosV();
```

---

### imgui.setNextWindowSizeV
设置下一窗口大小（带参数）

**使用示例：**
```lua
-- 调用 imgui.setNextWindowSizeV 方法
imgui.setNextWindowSizeV();
```

---

### imgui.setNextWindowViewport
设置下一窗口视口

**使用示例：**
```lua
-- 调用 imgui.setNextWindowViewport 方法
imgui.setNextWindowViewport();
```

---

### imgui.showAboutWindowV
显示关于窗口（带参数）

**使用示例：**
```lua
-- 调用 imgui.showAboutWindowV 方法
imgui.showAboutWindowV();
```

---

### imgui.showDebugLogWindowV
显示调试日志窗口（带参数）

**使用示例：**
```lua
-- 调用 imgui.showDebugLogWindowV 方法
imgui.showDebugLogWindowV();
```

---

### imgui.showDemoWindowV
显示演示窗口（带参数）

**使用示例：**
```lua
-- 调用 imgui.showDemoWindowV 方法
imgui.showDemoWindowV();
```

---

### imgui.showIDStackToolWindowV
显示ID堆栈工具窗口（带参数）

**使用示例：**
```lua
-- 调用 imgui.showIDStackToolWindowV 方法
imgui.showIDStackToolWindowV();
```

---

### imgui.showMetricsWindowV
显示指标窗口（带参数）

**使用示例：**
```lua
-- 调用 imgui.showMetricsWindowV 方法
imgui.showMetricsWindowV();
```

---

### imgui.showStyleEditorV
显示样式编辑器（带参数）

**使用示例：**
```lua
-- 调用 imgui.showStyleEditorV 方法
imgui.showStyleEditorV();
```

---

### imgui.styleColorsClassic
经典样式颜色

**使用示例：**
```lua
-- 调用 imgui.styleColorsClassic 方法
imgui.styleColorsClassic();
```

---

### imgui.styleColorsClassicV
经典样式颜色（带参数）

**使用示例：**
```lua
-- 调用 imgui.styleColorsClassicV 方法
imgui.styleColorsClassicV();
```

---

### imgui.styleColorsDark
深色样式颜色

**使用示例：**
```lua
-- 调用 imgui.styleColorsDark 方法
imgui.styleColorsDark();
```

---

### imgui.styleColorsDarkV
深色样式颜色（带参数）

**使用示例：**
```lua
-- 调用 imgui.styleColorsDarkV 方法
imgui.styleColorsDarkV();
```

---

### imgui.styleColorsLight
浅色样式颜色

**使用示例：**
```lua
-- 调用 imgui.styleColorsLight 方法
imgui.styleColorsLight();
```

---

### imgui.styleColorsLightV
浅色样式颜色（带参数）

**使用示例：**
```lua
-- 调用 imgui.styleColorsLightV 方法
imgui.styleColorsLightV();
```

---

### imgui.tableAngledHeadersRow
表格角度标题行

**使用示例：**
```lua
-- 调用 imgui.tableAngledHeadersRow 方法
imgui.tableAngledHeadersRow();
```

---

### imgui.tableNextRow
表格下一行

**使用示例：**
```lua
-- 调用 imgui.tableNextRow 方法
imgui.tableNextRow();
```

---

### imgui.tableSetBgColor
设置表格背景颜色

**使用示例：**
```lua
-- 调用 imgui.tableSetBgColor 方法
imgui.tableSetBgColor();
```

---

### imgui.tableSetupColumn
设置表格列

**使用示例：**
```lua
-- 调用 imgui.tableSetupColumn 方法
imgui.tableSetupColumn();
```

---

### imgui.textUnformattedV
无格式文本（带参数）

**使用示例：**
```lua
-- 调用 imgui.textUnformattedV 方法
imgui.textUnformattedV();
```

---

### imgui.treePushPtr
树推送（指针）

**使用示例：**
```lua
-- 调用 imgui.treePushPtr 方法
imgui.treePushPtr();
```

---

### imgui.updatePlatformWindows
更新平台窗口

**使用示例：**
```lua
-- 调用 imgui.updatePlatformWindows 方法
imgui.updatePlatformWindows();
```

---

### imgui.valueBool
布尔值

**使用示例：**
```lua
-- 调用 imgui.valueBool 方法
imgui.valueBool();
```

---

### imgui.valueFloat
浮点值

**使用示例：**
```lua
-- 调用 imgui.valueFloat 方法
imgui.valueFloat();
```

---

### imgui.valueFloatV
浮点值（带参数）

**使用示例：**
```lua
-- 调用 imgui.valueFloatV 方法
imgui.valueFloatV();
```

---

### imgui.valueInt
整数值

**使用示例：**
```lua
-- 调用 imgui.valueInt 方法
imgui.valueInt();
```

---

### imgui.valueUint
无符号整数值

**使用示例：**
```lua
-- 调用 imgui.valueUint 方法
imgui.valueUint();
```

---

## 综合使用示例

### 示例1：基本使用
```lua
-- imgui 模块的基本使用示例
```