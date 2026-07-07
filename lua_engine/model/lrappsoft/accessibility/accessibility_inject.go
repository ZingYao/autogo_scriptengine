//go:build ignore
// +build ignore

package accessibility

import (
	"fmt"

	"github.com/Dasongzi1366/AutoGo/uiacc"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"

	lua "github.com/yuin/gopher-lua"
)

type AccessibilityModule struct {
	ThrowException bool // 是否抛出异常
	ShowWarning    bool // 是否显示警告信息
	Debug          bool // 是否开启调试模式（打印脚本堆栈信息）
}

func New() *AccessibilityModule {
	return &AccessibilityModule{
		ThrowException: false,
		ShowWarning:    true,
		Debug:          false,
	}
}

// handleUnimplementedMethod 处理未实现方法的通用逻辑
func (m *AccessibilityModule) handleUnimplementedMethod(methodName string, L *lua.LState) int {
	// 如果开启了调试模式，打印脚本堆栈信息
	if m.Debug {
		fmt.Println("\n=== 调试信息: 调用未实现方法 ===")
		fmt.Printf("方法名: accessibility.%s\n", methodName)
		fmt.Println("=== Lua 调用堆栈 ===")
		for i := 1; ; i++ {
			dbg, ok := L.GetStack(i)
			if !ok {
				break
			}
			fmt.Printf("  [%d] 函数: %s, 行号: %d\n", i, dbg.Name, dbg.CurrentLine)
			fmt.Printf("       源文件: %s\n", dbg.Source)
			fmt.Printf("       What: %s, 调用行: %d\n", dbg.What, dbg.LineDefined)
		}
		fmt.Println("=== 堆栈结束 ===\n")
	}

	// 打印警告信息
	if m.ShowWarning {
		fmt.Printf("[警告] accessibility.%s 方法未实现，此功能在 AutoGo 中不支持\n", methodName)
	}

	// 根据配置决定是否抛出异常
	if m.ThrowException {
		L.RaiseError(fmt.Sprintf("accessibility.%s 方法未实现，此功能在 AutoGo 中不支持", methodName))
		return 0
	}

	// 默认返回 nil
	L.Push(lua.LNil)
	return 1
}

func (m *AccessibilityModule) Inject(state *lua.LState) {
	// ========== nodeLib 相关方法 ==========

	// 创建 nodeLib 表
	nodeLibTable := state.NewTable()

	// 注册 nodeLib.isAccServiceOk - 检测无障碍服务是否开启了
	nodeLibTable.RawSetString("isAccServiceOk", state.NewFunction(func(L *lua.LState) int {
		// AutoGo 的 uiacc 模块无需开启无障碍服务
		// 直接返回 true 表示可用
		L.Push(lua.LBool(true))
		return 1
	}))

	// 注册 nodeLib.openSnapService - 打开截图服务
	nodeLibTable.RawSetString("openSnapService", state.NewFunction(func(L *lua.LState) int {
		// AutoGo 使用 images.CaptureScreen 截图，无需开启截图服务
		// 直接返回 true 表示可用
		L.Push(lua.LBool(true))
		return 1
	}))

	// 注册 nodeLib.isSnapServiceOk - 检测截图服务是否开启
	nodeLibTable.RawSetString("isSnapServiceOk", state.NewFunction(func(L *lua.LState) int {
		// AutoGo 使用 images.CaptureScreen 截图，无需检测截图服务
		// 直接返回 true 表示可用
		L.Push(lua.LBool(true))
		return 1
	}))

	// 注册 nodeLib.findNextNode - 获取指定节点的下一个兄弟节点
	nodeLibTable.RawSetString("findNextNode", state.NewFunction(func(L *lua.LState) int {
		node := L.CheckTable(1)
		_ = L.CheckBool(2)

		// 从节点表格中提取信息
		class := lua.LVAsString(L.GetField(node, "class"))
		id := lua.LVAsString(L.GetField(node, "id"))
		text := lua.LVAsString(L.GetField(node, "text"))
		packageName := lua.LVAsString(L.GetField(node, "package"))

		// 创建 uiacc 选择器来查找这个节点
		acc := uiacc.New(0)
		if acc == nil {
			L.Push(lua.LNil)
			return 1
		}
		defer acc.Release()

		// 设置选择器条件
		if class != "" {
			acc.ClassName(class)
		}
		if id != "" {
			acc.Id(id)
		}
		if text != "" {
			acc.Text(text)
		}
		if packageName != "" {
			acc.PackageName(packageName)
		}

		// 查找节点
		uiObj := acc.FindOnce()
		if uiObj == nil {
			L.Push(lua.LNil)
			return 1
		}

		// 获取父节点
		parent := uiObj.GetParent()
		if parent == nil {
			L.Push(lua.LNil)
			return 1
		}

		// 获取当前节点的索引
		currentIndex := uiObj.GetIndex()

		// 获取所有子节点
		children := parent.GetChildren()
		if children == nil || len(children) == 0 {
			L.Push(lua.LNil)
			return 1
		}

		// 查找下一个兄弟节点
		if currentIndex >= 0 && currentIndex < len(children)-1 {
			nextNode := children[currentIndex+1]
			if nextNode != nil {
				result := L.NewTable()
				result.RawSetString("class", lua.LString(nextNode.GetClassName()))
				result.RawSetString("id", lua.LString(nextNode.GetId()))
				result.RawSetString("package", lua.LString(nextNode.GetPackageName()))
				result.RawSetString("text", lua.LString(nextNode.GetText()))
				result.RawSetString("desc", lua.LString(nextNode.GetDesc()))

				bounds := nextNode.GetBounds()
				boundsTable := L.NewTable()
				boundsTable.RawSetString("left", lua.LNumber(bounds.Left))
				boundsTable.RawSetString("top", lua.LNumber(bounds.Top))
				boundsTable.RawSetString("right", lua.LNumber(bounds.Right))
				boundsTable.RawSetString("bottom", lua.LNumber(bounds.Bottom))
				boundsTable.RawSetString("width", lua.LNumber(bounds.Width))
				boundsTable.RawSetString("height", lua.LNumber(bounds.Height))
				result.RawSetString("bounds", boundsTable)

				L.Push(result)
				return 1
			}
		}

		L.Push(lua.LNil)
		return 1
	}))

	// 注册 nodeLib.findPreNode - 获取指定节点的上一个兄弟节点
	nodeLibTable.RawSetString("findPreNode", state.NewFunction(func(L *lua.LState) int {
		node := L.CheckTable(1)
		_ = L.CheckBool(2)

		// 从节点表格中提取信息
		class := lua.LVAsString(L.GetField(node, "class"))
		id := lua.LVAsString(L.GetField(node, "id"))
		text := lua.LVAsString(L.GetField(node, "text"))
		packageName := lua.LVAsString(L.GetField(node, "package"))

		// 创建 uiacc 选择器来查找这个节点
		acc := uiacc.New(0)
		if acc == nil {
			L.Push(lua.LNil)
			return 1
		}
		defer acc.Release()

		// 设置选择器条件
		if class != "" {
			acc.ClassName(class)
		}
		if id != "" {
			acc.Id(id)
		}
		if text != "" {
			acc.Text(text)
		}
		if packageName != "" {
			acc.PackageName(packageName)
		}

		// 查找节点
		uiObj := acc.FindOnce()
		if uiObj == nil {
			L.Push(lua.LNil)
			return 1
		}

		// 获取父节点
		parent := uiObj.GetParent()
		if parent == nil {
			L.Push(lua.LNil)
			return 1
		}

		// 获取当前节点的索引
		currentIndex := uiObj.GetIndex()

		// 获取所有子节点
		children := parent.GetChildren()
		if children == nil || len(children) == 0 {
			L.Push(lua.LNil)
			return 1
		}

		// 查找上一个兄弟节点
		if currentIndex > 0 && currentIndex < len(children) {
			preNode := children[currentIndex-1]
			if preNode != nil {
				result := L.NewTable()
				result.RawSetString("class", lua.LString(preNode.GetClassName()))
				result.RawSetString("id", lua.LString(preNode.GetId()))
				result.RawSetString("package", lua.LString(preNode.GetPackageName()))
				result.RawSetString("text", lua.LString(preNode.GetText()))
				result.RawSetString("desc", lua.LString(preNode.GetDesc()))

				bounds := preNode.GetBounds()
				boundsTable := L.NewTable()
				boundsTable.RawSetString("left", lua.LNumber(bounds.Left))
				boundsTable.RawSetString("top", lua.LNumber(bounds.Top))
				boundsTable.RawSetString("right", lua.LNumber(bounds.Right))
				boundsTable.RawSetString("bottom", lua.LNumber(bounds.Bottom))
				boundsTable.RawSetString("width", lua.LNumber(bounds.Width))
				boundsTable.RawSetString("height", lua.LNumber(bounds.Height))
				result.RawSetString("bounds", boundsTable)

				L.Push(result)
				return 1
			}
		}

		L.Push(lua.LNil)
		return 1
	}))

	// 注册 nodeLib.findInNode - 在指定节点中查找符合要求的子节点
	nodeLibTable.RawSetString("findInNode", state.NewFunction(func(L *lua.LState) int {
		parentNode := L.CheckTable(1)
		selector := L.CheckTable(2)
		findAll := L.CheckBool(3)
		_ = L.CheckBool(4)

		// 从父节点表格中提取信息
		parentClass := lua.LVAsString(L.GetField(parentNode, "class"))
		parentId := lua.LVAsString(L.GetField(parentNode, "id"))
		parentText := lua.LVAsString(L.GetField(parentNode, "text"))
		parentPackage := lua.LVAsString(L.GetField(parentNode, "package"))

		// 创建 uiacc 选择器来查找父节点
		acc := uiacc.New(0)
		if acc == nil {
			L.Push(lua.LNil)
			return 1
		}
		defer acc.Release()

		// 设置父节点选择器条件
		if parentClass != "" {
			acc.ClassName(parentClass)
		}
		if parentId != "" {
			acc.Id(parentId)
		}
		if parentText != "" {
			acc.Text(parentText)
		}
		if parentPackage != "" {
			acc.PackageName(parentPackage)
		}

		// 查找父节点
		parentObj := acc.FindOnce()
		if parentObj == nil {
			L.Push(lua.LNil)
			return 1
		}

		// 获取所有子节点
		children := parentObj.GetChildren()
		if children == nil || len(children) == 0 {
			L.Push(lua.LNil)
			return 1
		}

		// 从选择器中提取筛选条件
		selectorText := lua.LVAsString(L.GetField(selector, "text"))
		selectorId := lua.LVAsString(L.GetField(selector, "id"))
		selectorClass := lua.LVAsString(L.GetField(selector, "class"))
		selectorPackage := lua.LVAsString(L.GetField(selector, "package"))

		// 筛选子节点
		matchedChildren := []*uiacc.UiObject{}
		for _, child := range children {
			matched := true

			// 检查 text
			if selectorText != "" && child.GetText() != selectorText {
				matched = false
			}

			// 检查 id
			if matched && selectorId != "" && child.GetId() != selectorId {
				matched = false
			}

			// 检查 class
			if matched && selectorClass != "" && child.GetClassName() != selectorClass {
				matched = false
			}

			// 检查 package
			if matched && selectorPackage != "" && child.GetPackageName() != selectorPackage {
				matched = false
			}

			if matched {
				matchedChildren = append(matchedChildren, child)
			}
		}

		// 如果没有匹配的子节点
		if len(matchedChildren) == 0 {
			L.Push(lua.LNil)
			return 1
		}

		// 根据 findAll 参数决定返回一个还是所有
		if !findAll {
			// 只返回第一个匹配的子节点
			child := matchedChildren[0]
			result := L.NewTable()
			result.RawSetString("class", lua.LString(child.GetClassName()))
			result.RawSetString("id", lua.LString(child.GetId()))
			result.RawSetString("package", lua.LString(child.GetPackageName()))
			result.RawSetString("text", lua.LString(child.GetText()))
			result.RawSetString("desc", lua.LString(child.GetDesc()))

			bounds := child.GetBounds()
			boundsTable := L.NewTable()
			boundsTable.RawSetString("left", lua.LNumber(bounds.Left))
			boundsTable.RawSetString("top", lua.LNumber(bounds.Top))
			boundsTable.RawSetString("right", lua.LNumber(bounds.Right))
			boundsTable.RawSetString("bottom", lua.LNumber(bounds.Bottom))
			boundsTable.RawSetString("width", lua.LNumber(bounds.Width))
			boundsTable.RawSetString("height", lua.LNumber(bounds.Height))
			result.RawSetString("bounds", boundsTable)

			L.Push(result)
			return 1
		} else {
			// 返回所有匹配的子节点
			result := L.NewTable()
			for i, child := range matchedChildren {
				childTable := L.NewTable()
				childTable.RawSetString("class", lua.LString(child.GetClassName()))
				childTable.RawSetString("id", lua.LString(child.GetId()))
				childTable.RawSetString("package", lua.LString(child.GetPackageName()))
				childTable.RawSetString("text", lua.LString(child.GetText()))
				childTable.RawSetString("desc", lua.LString(child.GetDesc()))

				bounds := child.GetBounds()
				boundsTable := L.NewTable()
				boundsTable.RawSetString("left", lua.LNumber(bounds.Left))
				boundsTable.RawSetString("top", lua.LNumber(bounds.Top))
				boundsTable.RawSetString("right", lua.LNumber(bounds.Right))
				boundsTable.RawSetString("bottom", lua.LNumber(bounds.Bottom))
				boundsTable.RawSetString("width", lua.LNumber(bounds.Width))
				boundsTable.RawSetString("height", lua.LNumber(bounds.Height))
				childTable.RawSetString("bounds", boundsTable)

				result.RawSetInt(i+1, childTable)
			}

			L.Push(result)
			return 1
		}
	}))

	// 注册 nodeLib.findOne - 查找出一个节点
	nodeLibTable.RawSetString("findOne", state.NewFunction(func(L *lua.LState) int {
		selector := L.CheckTable(1)
		fuzzyMatch := L.CheckBool(2)

		// 创建 uiacc 选择器
		acc := uiacc.New(0)
		if acc == nil {
			L.Push(lua.LNil)
			return 1
		}
		defer acc.Release()

		// 解析 selector 表格
		if text := L.GetField(selector, "text"); text != lua.LNil {
			textStr := lua.LVAsString(text)
			if fuzzyMatch {
				acc.TextContains(textStr)
			} else {
				acc.Text(textStr)
			}
		}

		if id := L.GetField(selector, "id"); id != lua.LNil {
			idStr := lua.LVAsString(id)
			if fuzzyMatch {
				acc.IdContains(idStr)
			} else {
				acc.Id(idStr)
			}
		}

		if className := L.GetField(selector, "class"); className != lua.LNil {
			classStr := lua.LVAsString(className)
			if fuzzyMatch {
				acc.ClassNameContains(classStr)
			} else {
				acc.ClassName(classStr)
			}
		}

		if packageName := L.GetField(selector, "package"); packageName != lua.LNil {
			packageStr := lua.LVAsString(packageName)
			if fuzzyMatch {
				acc.PackageNameContains(packageStr)
			} else {
				acc.PackageName(packageStr)
			}
		}

		// 查找节点
		uiObj := acc.FindOnce()
		if uiObj == nil {
			L.Push(lua.LNil)
			return 1
		}

		// 将 UiObject 转换为 Lua 表
		result := L.NewTable()
		result.RawSetString("class", lua.LString(uiObj.GetClassName()))
		result.RawSetString("id", lua.LString(uiObj.GetId()))
		result.RawSetString("package", lua.LString(uiObj.GetPackageName()))
		result.RawSetString("text", lua.LString(uiObj.GetText()))
		result.RawSetString("desc", lua.LString(uiObj.GetDesc()))

		// 获取节点位置信息
		bounds := uiObj.GetBounds()
		boundsTable := L.NewTable()
		boundsTable.RawSetString("left", lua.LNumber(bounds.Left))
		boundsTable.RawSetString("top", lua.LNumber(bounds.Top))
		boundsTable.RawSetString("right", lua.LNumber(bounds.Right))
		boundsTable.RawSetString("bottom", lua.LNumber(bounds.Bottom))
		boundsTable.RawSetString("width", lua.LNumber(bounds.Width))
		boundsTable.RawSetString("height", lua.LNumber(bounds.Height))
		result.RawSetString("bounds", boundsTable)

		L.Push(result)
		return 1
	}))

	// 注册 nodeLib.findNode - findOne 的别名（兼容性）
	nodeLibTable.RawSetString("findNode", state.NewFunction(func(L *lua.LState) int {
		selector := L.CheckTable(1)
		fuzzyMatch := L.OptBool(2, false)

		// 创建 uiacc 选择器
		acc := uiacc.New(0)
		if acc == nil {
			L.Push(lua.LNil)
			return 1
		}
		defer acc.Release()

		// 解析 selector 表格
		if text := L.GetField(selector, "text"); text != lua.LNil {
			textStr := lua.LVAsString(text)
			if fuzzyMatch {
				acc.TextContains(textStr)
			} else {
				acc.Text(textStr)
			}
		}

		if id := L.GetField(selector, "id"); id != lua.LNil {
			idStr := lua.LVAsString(id)
			if fuzzyMatch {
				acc.IdContains(idStr)
			} else {
				acc.Id(idStr)
			}
		}

		if className := L.GetField(selector, "class"); className != lua.LNil {
			classStr := lua.LVAsString(className)
			if fuzzyMatch {
				acc.ClassNameContains(classStr)
			} else {
				acc.ClassName(classStr)
			}
		}

		if packageName := L.GetField(selector, "package"); packageName != lua.LNil {
			packageStr := lua.LVAsString(packageName)
			if fuzzyMatch {
				acc.PackageNameContains(packageStr)
			} else {
				acc.PackageName(packageStr)
			}
		}

		// 查找节点
		uiObj := acc.FindOnce()
		if uiObj == nil {
			L.Push(lua.LNil)
			return 1
		}

		// 将 UiObject 转换为 Lua 表
		result := L.NewTable()
		result.RawSetString("class", lua.LString(uiObj.GetClassName()))
		result.RawSetString("id", lua.LString(uiObj.GetId()))
		result.RawSetString("package", lua.LString(uiObj.GetPackageName()))
		result.RawSetString("text", lua.LString(uiObj.GetText()))
		result.RawSetString("desc", lua.LString(uiObj.GetDesc()))

		// 获取节点位置信息
		bounds := uiObj.GetBounds()
		boundsTable := L.NewTable()
		boundsTable.RawSetString("left", lua.LNumber(bounds.Left))
		boundsTable.RawSetString("top", lua.LNumber(bounds.Top))
		boundsTable.RawSetString("right", lua.LNumber(bounds.Right))
		boundsTable.RawSetString("bottom", lua.LNumber(bounds.Bottom))
		boundsTable.RawSetString("width", lua.LNumber(bounds.Width))
		boundsTable.RawSetString("height", lua.LNumber(bounds.Height))
		result.RawSetString("bounds", boundsTable)

		L.Push(result)
		return 1
	}))

	// 注册 nodeLib.findAll - 查找所有满足要求的节点
	nodeLibTable.RawSetString("findAll", state.NewFunction(func(L *lua.LState) int {
		selector := L.CheckTable(1)
		fuzzyMatch := L.CheckBool(2)

		// 创建 uiacc 选择器
		acc := uiacc.New(0)
		if acc == nil {
			L.Push(lua.LNil)
			return 1
		}
		defer acc.Release()

		// 解析 selector 表格
		if text := L.GetField(selector, "text"); text != lua.LNil {
			textStr := lua.LVAsString(text)
			if fuzzyMatch {
				acc.TextContains(textStr)
			} else {
				acc.Text(textStr)
			}
		}

		if id := L.GetField(selector, "id"); id != lua.LNil {
			idStr := lua.LVAsString(id)
			if fuzzyMatch {
				acc.IdContains(idStr)
			} else {
				acc.Id(idStr)
			}
		}

		if className := L.GetField(selector, "class"); className != lua.LNil {
			classStr := lua.LVAsString(className)
			if fuzzyMatch {
				acc.ClassNameContains(classStr)
			} else {
				acc.ClassName(classStr)
			}
		}

		if packageName := L.GetField(selector, "package"); packageName != lua.LNil {
			packageStr := lua.LVAsString(packageName)
			if fuzzyMatch {
				acc.PackageNameContains(packageStr)
			} else {
				acc.PackageName(packageStr)
			}
		}

		// 查找所有节点
		uiObjs := acc.Find()
		if uiObjs == nil || len(uiObjs) == 0 {
			L.Push(lua.LNil)
			return 1
		}

		// 将 UiObject 数组转换为 Lua 表数组
		result := L.NewTable()
		for i, uiObj := range uiObjs {
			nodeTable := L.NewTable()
			nodeTable.RawSetString("class", lua.LString(uiObj.GetClassName()))
			nodeTable.RawSetString("id", lua.LString(uiObj.GetId()))
			nodeTable.RawSetString("package", lua.LString(uiObj.GetPackageName()))
			nodeTable.RawSetString("text", lua.LString(uiObj.GetText()))
			nodeTable.RawSetString("desc", lua.LString(uiObj.GetDesc()))

			// 获取节点位置信息
			bounds := uiObj.GetBounds()
			boundsTable := L.NewTable()
			boundsTable.RawSetString("left", lua.LNumber(bounds.Left))
			boundsTable.RawSetString("top", lua.LNumber(bounds.Top))
			boundsTable.RawSetString("right", lua.LNumber(bounds.Right))
			boundsTable.RawSetString("bottom", lua.LNumber(bounds.Bottom))
			boundsTable.RawSetString("width", lua.LNumber(bounds.Width))
			boundsTable.RawSetString("height", lua.LNumber(bounds.Height))
			nodeTable.RawSetString("bounds", boundsTable)

			result.RawSetInt(i+1, nodeTable)
		}

		L.Push(result)
		return 1
	}))

	// 注册 nodeLib.findNodes - findAll 的别名（兼容性）
	nodeLibTable.RawSetString("findNodes", state.NewFunction(func(L *lua.LState) int {
		selector := L.CheckTable(1)
		fuzzyMatch := L.OptBool(2, false)

		// 创建 uiacc 选择器
		acc := uiacc.New(0)
		if acc == nil {
			L.Push(lua.LNil)
			return 1
		}
		defer acc.Release()

		// 解析 selector 表格
		if text := L.GetField(selector, "text"); text != lua.LNil {
			textStr := lua.LVAsString(text)
			if fuzzyMatch {
				acc.TextContains(textStr)
			} else {
				acc.Text(textStr)
			}
		}

		if id := L.GetField(selector, "id"); id != lua.LNil {
			idStr := lua.LVAsString(id)
			if fuzzyMatch {
				acc.IdContains(idStr)
			} else {
				acc.Id(idStr)
			}
		}

		if className := L.GetField(selector, "class"); className != lua.LNil {
			classStr := lua.LVAsString(className)
			if fuzzyMatch {
				acc.ClassNameContains(classStr)
			} else {
				acc.ClassName(classStr)
			}
		}

		if packageName := L.GetField(selector, "package"); packageName != lua.LNil {
			packageStr := lua.LVAsString(packageName)
			if fuzzyMatch {
				acc.PackageNameContains(packageStr)
			} else {
				acc.PackageName(packageStr)
			}
		}

		// 查找所有节点
		uiObjs := acc.Find()
		if uiObjs == nil || len(uiObjs) == 0 {
			L.Push(lua.LNil)
			return 1
		}

		// 将 UiObject 数组转换为 Lua 表数组
		result := L.NewTable()
		for i, uiObj := range uiObjs {
			nodeTable := L.NewTable()
			nodeTable.RawSetString("class", lua.LString(uiObj.GetClassName()))
			nodeTable.RawSetString("id", lua.LString(uiObj.GetId()))
			nodeTable.RawSetString("package", lua.LString(uiObj.GetPackageName()))
			nodeTable.RawSetString("text", lua.LString(uiObj.GetText()))
			nodeTable.RawSetString("desc", lua.LString(uiObj.GetDesc()))

			// 获取节点位置信息
			bounds := uiObj.GetBounds()
			boundsTable := L.NewTable()
			boundsTable.RawSetString("left", lua.LNumber(bounds.Left))
			boundsTable.RawSetString("top", lua.LNumber(bounds.Top))
			boundsTable.RawSetString("right", lua.LNumber(bounds.Right))
			boundsTable.RawSetString("bottom", lua.LNumber(bounds.Bottom))
			boundsTable.RawSetString("width", lua.LNumber(bounds.Width))
			boundsTable.RawSetString("height", lua.LNumber(bounds.Height))
			nodeTable.RawSetString("bounds", boundsTable)

			result.RawSetInt(i+1, nodeTable)
		}

		L.Push(result)
		return 1
	}))

	// 注册 nodeLib.findChildNodes - 查找一个节点的所有子节点
	nodeLibTable.RawSetString("findChildNodes", state.NewFunction(func(L *lua.LState) int {
		node := L.CheckTable(1)

		// 从节点表格中提取信息
		class := lua.LVAsString(L.GetField(node, "class"))
		id := lua.LVAsString(L.GetField(node, "id"))
		text := lua.LVAsString(L.GetField(node, "text"))
		packageName := lua.LVAsString(L.GetField(node, "package"))

		// 创建 uiacc 选择器来查找这个节点
		acc := uiacc.New(0)
		if acc == nil {
			L.Push(lua.LNil)
			return 1
		}
		defer acc.Release()

		// 设置选择器条件
		if class != "" {
			acc.ClassName(class)
		}
		if id != "" {
			acc.Id(id)
		}
		if text != "" {
			acc.Text(text)
		}
		if packageName != "" {
			acc.PackageName(packageName)
		}

		// 查找父节点
		parentObj := acc.FindOnce()
		if parentObj == nil {
			L.Push(lua.LNil)
			return 1
		}

		// 获取所有子节点
		children := parentObj.GetChildren()
		if children == nil || len(children) == 0 {
			L.Push(lua.LNil)
			return 1
		}

		// 将子节点数组转换为 Lua 表数组
		result := L.NewTable()
		for i, child := range children {
			nodeTable := L.NewTable()
			nodeTable.RawSetString("class", lua.LString(child.GetClassName()))
			nodeTable.RawSetString("id", lua.LString(child.GetId()))
			nodeTable.RawSetString("package", lua.LString(child.GetPackageName()))
			nodeTable.RawSetString("text", lua.LString(child.GetText()))
			nodeTable.RawSetString("desc", lua.LString(child.GetDesc()))

			// 获取节点位置信息
			bounds := child.GetBounds()
			boundsTable := L.NewTable()
			boundsTable.RawSetString("left", lua.LNumber(bounds.Left))
			boundsTable.RawSetString("top", lua.LNumber(bounds.Top))
			boundsTable.RawSetString("right", lua.LNumber(bounds.Right))
			boundsTable.RawSetString("bottom", lua.LNumber(bounds.Bottom))
			boundsTable.RawSetString("width", lua.LNumber(bounds.Width))
			boundsTable.RawSetString("height", lua.LNumber(bounds.Height))
			nodeTable.RawSetString("bounds", boundsTable)

			result.RawSetInt(i+1, nodeTable)
		}

		L.Push(result)
		return 1
	}))

	// 注册 nodeLib.findParentNode - 获取指定节点的父节点
	nodeLibTable.RawSetString("findParentNode", state.NewFunction(func(L *lua.LState) int {
		node := L.CheckTable(1)

		// 从节点表格中提取信息
		class := lua.LVAsString(L.GetField(node, "class"))
		id := lua.LVAsString(L.GetField(node, "id"))
		text := lua.LVAsString(L.GetField(node, "text"))
		packageName := lua.LVAsString(L.GetField(node, "package"))

		// 创建 uiacc 选择器来查找这个节点
		acc := uiacc.New(0)
		if acc == nil {
			L.Push(lua.LNil)
			return 1
		}
		defer acc.Release()

		// 设置选择器条件
		if class != "" {
			acc.ClassName(class)
		}
		if id != "" {
			acc.Id(id)
		}
		if text != "" {
			acc.Text(text)
		}
		if packageName != "" {
			acc.PackageName(packageName)
		}

		// 查找节点
		uiObj := acc.FindOnce()
		if uiObj == nil {
			L.Push(lua.LNil)
			return 1
		}

		// 获取父节点
		parent := uiObj.GetParent()
		if parent == nil {
			L.Push(lua.LNil)
			return 1
		}

		// 将父节点转换为 Lua 表
		result := L.NewTable()
		result.RawSetString("class", lua.LString(parent.GetClassName()))
		result.RawSetString("id", lua.LString(parent.GetId()))
		result.RawSetString("package", lua.LString(parent.GetPackageName()))
		result.RawSetString("text", lua.LString(parent.GetText()))
		result.RawSetString("desc", lua.LString(parent.GetDesc()))

		bounds := parent.GetBounds()
		boundsTable := L.NewTable()
		boundsTable.RawSetString("left", lua.LNumber(bounds.Left))
		boundsTable.RawSetString("top", lua.LNumber(bounds.Top))
		boundsTable.RawSetString("right", lua.LNumber(bounds.Right))
		boundsTable.RawSetString("bottom", lua.LNumber(bounds.Bottom))
		boundsTable.RawSetString("width", lua.LNumber(bounds.Width))
		boundsTable.RawSetString("height", lua.LNumber(bounds.Height))
		result.RawSetString("bounds", boundsTable)

		L.Push(result)
		return 1
	}))

	// 注册 nodeLib.getNodeInfo - 获取节点信息
	nodeLibTable.RawSetString("getNodeInfo", state.NewFunction(func(L *lua.LState) int {
		node := L.CheckTable(1)

		// 直接返回节点表格（因为节点表格已经包含了所有信息）
		L.Push(node)
		return 1
	}))

	// 注册 nodeLib.clickNode - 点击节点
	nodeLibTable.RawSetString("clickNode", state.NewFunction(func(L *lua.LState) int {
		node := L.CheckTable(1)

		// 从节点表格中提取信息
		class := lua.LVAsString(L.GetField(node, "class"))
		id := lua.LVAsString(L.GetField(node, "id"))
		text := lua.LVAsString(L.GetField(node, "text"))
		packageName := lua.LVAsString(L.GetField(node, "package"))

		// 创建 uiacc 选择器来查找这个节点
		acc := uiacc.New(0)
		if acc == nil {
			L.Push(lua.LBool(false))
			return 1
		}
		defer acc.Release()

		// 设置选择器条件
		if class != "" {
			acc.ClassName(class)
		}
		if id != "" {
			acc.Id(id)
		}
		if text != "" {
			acc.Text(text)
		}
		if packageName != "" {
			acc.PackageName(packageName)
		}

		// 查找并点击节点
		uiObj := acc.FindOnce()
		if uiObj == nil {
			L.Push(lua.LBool(false))
			return 1
		}

		// 点击节点
		clicked := uiObj.Click()
		L.Push(lua.LBool(clicked))
		return 1
	}))

	// 注册 nodeLib.longClickNode - 长按节点
	nodeLibTable.RawSetString("longClickNode", state.NewFunction(func(L *lua.LState) int {
		node := L.CheckTable(1)

		// 从节点表格中提取信息
		class := lua.LVAsString(L.GetField(node, "class"))
		id := lua.LVAsString(L.GetField(node, "id"))
		text := lua.LVAsString(L.GetField(node, "text"))
		packageName := lua.LVAsString(L.GetField(node, "package"))

		// 创建 uiacc 选择器来查找这个节点
		acc := uiacc.New(0)
		if acc == nil {
			L.Push(lua.LBool(false))
			return 1
		}
		defer acc.Release()

		// 设置选择器条件
		if class != "" {
			acc.ClassName(class)
		}
		if id != "" {
			acc.Id(id)
		}
		if text != "" {
			acc.Text(text)
		}
		if packageName != "" {
			acc.PackageName(packageName)
		}

		// 查找并长按节点
		uiObj := acc.FindOnce()
		if uiObj == nil {
			L.Push(lua.LBool(false))
			return 1
		}

		// 长按节点
		longClicked := uiObj.ClickLongClick()
		L.Push(lua.LBool(longClicked))
		return 1
	}))

	// 注册 nodeLib.getNodeText - 获取节点文本
	nodeLibTable.RawSetString("getNodeText", state.NewFunction(func(L *lua.LState) int {
		node := L.CheckTable(1)

		// 直接从节点表格中获取文本
		text := lua.LVAsString(L.GetField(node, "text"))
		L.Push(lua.LString(text))
		return 1
	}))

	// 注册 nodeLib.setNodeText - 设置节点文本
	nodeLibTable.RawSetString("setNodeText", state.NewFunction(func(L *lua.LState) int {
		node := L.CheckTable(1)
		newText := L.CheckString(2)

		// 从节点表格中提取信息
		class := lua.LVAsString(L.GetField(node, "class"))
		id := lua.LVAsString(L.GetField(node, "id"))
		text := lua.LVAsString(L.GetField(node, "text"))
		packageName := lua.LVAsString(L.GetField(node, "package"))

		// 创建 uiacc 选择器来查找这个节点
		acc := uiacc.New(0)
		if acc == nil {
			L.Push(lua.LBool(false))
			return 1
		}
		defer acc.Release()

		// 设置选择器条件
		if class != "" {
			acc.ClassName(class)
		}
		if id != "" {
			acc.Id(id)
		}
		if text != "" {
			acc.Text(text)
		}
		if packageName != "" {
			acc.PackageName(packageName)
		}

		// 查找并设置节点文本
		uiObj := acc.FindOnce()
		if uiObj == nil {
			L.Push(lua.LBool(false))
			return 1
		}

		// 设置节点文本
		set := uiObj.SetText(newText)
		L.Push(lua.LBool(set))
		return 1
	}))

	// 注册 nodeLib.getNodeBounds - 获取节点边界
	nodeLibTable.RawSetString("getNodeBounds", state.NewFunction(func(L *lua.LState) int {
		node := L.CheckTable(1)

		// 直接从节点表格中获取边界
		bounds := L.GetField(node, "bounds")
		if bounds == lua.LNil {
			L.Push(lua.LNil)
			return 1
		}

		L.Push(bounds)
		return 1
	}))

	// 将 nodeLib 表注册到全局
	state.SetGlobal("nodeLib", nodeLibTable)
}

// Name 返回模块名称
func (m *AccessibilityModule) Name() string {
	return "accessibility"
}

// Register 向引擎注册模块的方法
func (m *AccessibilityModule) Register(engine model.Engine) error {
	m.Inject(engine.GetState())
	return nil
}

// IsAvailable 返回模块是否可用
func (m *AccessibilityModule) IsAvailable() bool {
	return true
}

// GetModule 获取模块实例
func GetModule() model.Module {
	return New()
}
