//go:build ignore
// +build ignore

package node

import (
	"fmt"
	"sync"

	"github.com/Dasongzi1366/AutoGo/uiacc"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"

	lua "github.com/yuin/gopher-lua"
)

type NodeModule struct {
	uiObjectCache map[int64]*uiacc.UiObject
	cacheMutex    sync.RWMutex
	nextId        int64
}

func New() *NodeModule {
	return &NodeModule{
		uiObjectCache: make(map[int64]*uiacc.UiObject),
		nextId:        1,
	}
}

// 缓存 UIObject 并返回唯一 ID
func (m *NodeModule) cacheUiObject(uiObj *uiacc.UiObject) int64 {
	if uiObj == nil {
		return 0
	}

	m.cacheMutex.Lock()
	defer m.cacheMutex.Unlock()

	id := m.nextId
	m.nextId++
	m.uiObjectCache[id] = uiObj
	return id
}

// 从缓存中获取 UIObject
func (m *NodeModule) getUiObject(id int64) *uiacc.UiObject {
	m.cacheMutex.RLock()
	defer m.cacheMutex.RUnlock()

	return m.uiObjectCache[id]
}

// 从缓存中删除 UIObject
func (m *NodeModule) removeUiObject(id int64) {
	m.cacheMutex.Lock()
	defer m.cacheMutex.Unlock()

	delete(m.uiObjectCache, id)
}

// calculateDepth 计算节点的深度
func (m *NodeModule) calculateDepth(uiObj *uiacc.UiObject) int {
	if uiObj == nil {
		return -1
	}

	depth := 0
	current := uiObj

	for current != nil {
		parent := current.GetParent()
		if parent == nil {
			break
		}
		depth++
		current = parent
	}

	return depth
}

// Selector 选择器结构体
type Selector struct {
	acc       *uiacc.Uiacc
	luaState  *lua.LState
	module    *NodeModule
	condition *SelectorCondition
}

// SelectorCondition 选择器条件
type SelectorCondition struct {
	id                    string
	idContains            string
	idStartsWith          string
	idEndsWith            string
	idMatches             string
	text                  string
	textContains          string
	textStartsWith        string
	textEndsWith          string
	textMatches           string
	desc                  string
	descContains          string
	descStartsWith        string
	descEndsWith          string
	descMatches           string
	className             string
	classNameContains     string
	classNameStartsWith   string
	classNameEndsWith     string
	classNameMatches      string
	packageName           string
	packageNameContains   string
	packageNameStartsWith string
	packageNameEndsWith   string
	packageNameMatches    string
	bounds                *Bounds
	boundsInside          *Bounds
	drawingOrder          int
	depth                 int
	index                 int
	visibleToUser         *bool
	selected              *bool
	clickable             *bool
	longClickable         *bool
	enabled               *bool
	password              *bool
	scrollable            *bool
	checked               *bool
	checkable             *bool
	focusable             *bool
	focused               *bool
}

// Bounds 边界
type Bounds struct {
	Left   int
	Top    int
	Right  int
	Bottom int
}

func (m *NodeModule) Inject(state *lua.LState) {
	// ========== 选择器方法 ==========

	// 注册 id - 根据节点的id属性匹配
	state.SetGlobal("id", state.NewFunction(func(L *lua.LState) int {
		str := L.CheckString(1)
		sel := &Selector{
			acc:      uiacc.New(0),
			luaState: L,
			module:   m,
			condition: &SelectorCondition{
				id: str,
			},
		}
		sel.applyCondition()
		L.Push(sel.toLuaTable())
		return 1
	}))

	// 注册 idContains - 根据id包含的部分字符串匹配
	state.SetGlobal("idContains", state.NewFunction(func(L *lua.LState) int {
		str := L.CheckString(1)
		sel := &Selector{
			acc:      uiacc.New(0),
			luaState: L,
			module:   m,
			condition: &SelectorCondition{
				idContains: str,
			},
		}
		sel.applyCondition()
		L.Push(sel.toLuaTable())
		return 1
	}))

	// 注册 idStartsWith - 根据id的前缀去匹配
	state.SetGlobal("idStartsWith", state.NewFunction(func(L *lua.LState) int {
		str := L.CheckString(1)
		sel := &Selector{
			acc:      uiacc.New(0),
			luaState: L,
			module:   m,
			condition: &SelectorCondition{
				idStartsWith: str,
			},
		}
		sel.applyCondition()
		L.Push(sel.toLuaTable())
		return 1
	}))

	// 注册 idEndsWith - 根据id的后缀去匹配
	state.SetGlobal("idEndsWith", state.NewFunction(func(L *lua.LState) int {
		str := L.CheckString(1)
		sel := &Selector{
			acc:      uiacc.New(0),
			luaState: L,
			module:   m,
			condition: &SelectorCondition{
				idEndsWith: str,
			},
		}
		sel.applyCondition()
		L.Push(sel.toLuaTable())
		return 1
	}))

	// 注册 idMatches - 正则匹配id
	state.SetGlobal("idMatches", state.NewFunction(func(L *lua.LState) int {
		str := L.CheckString(1)
		sel := &Selector{
			acc:      uiacc.New(0),
			luaState: L,
			module:   m,
			condition: &SelectorCondition{
				idMatches: str,
			},
		}
		sel.applyCondition()
		L.Push(sel.toLuaTable())
		return 1
	}))

	// 注册 text - 根据节点的text属性匹配
	state.SetGlobal("text", state.NewFunction(func(L *lua.LState) int {
		str := L.CheckString(1)
		sel := &Selector{
			acc:      uiacc.New(0),
			luaState: L,
			module:   m,
			condition: &SelectorCondition{
				text: str,
			},
		}
		sel.applyCondition()
		L.Push(sel.toLuaTable())
		return 1
	}))

	// 注册 textContains - 根据text包含的部分字符串匹配
	state.SetGlobal("textContains", state.NewFunction(func(L *lua.LState) int {
		str := L.CheckString(1)
		sel := &Selector{
			acc:      uiacc.New(0),
			luaState: L,
			module:   m,
			condition: &SelectorCondition{
				textContains: str,
			},
		}
		sel.applyCondition()
		L.Push(sel.toLuaTable())
		return 1
	}))

	// 注册 textStartsWith - 根据text的前缀去匹配
	state.SetGlobal("textStartsWith", state.NewFunction(func(L *lua.LState) int {
		str := L.CheckString(1)
		sel := &Selector{
			acc:      uiacc.New(0),
			luaState: L,
			module:   m,
			condition: &SelectorCondition{
				textStartsWith: str,
			},
		}
		sel.applyCondition()
		L.Push(sel.toLuaTable())
		return 1
	}))

	// 注册 textEndsWith - 根据text的后缀去匹配
	state.SetGlobal("textEndsWith", state.NewFunction(func(L *lua.LState) int {
		str := L.CheckString(1)
		sel := &Selector{
			acc:      uiacc.New(0),
			luaState: L,
			module:   m,
			condition: &SelectorCondition{
				textEndsWith: str,
			},
		}
		sel.applyCondition()
		L.Push(sel.toLuaTable())
		return 1
	}))

	// 注册 textMatches - 正则匹配text
	state.SetGlobal("textMatches", state.NewFunction(func(L *lua.LState) int {
		str := L.CheckString(1)
		sel := &Selector{
			acc:      uiacc.New(0),
			luaState: L,
			module:   m,
			condition: &SelectorCondition{
				textMatches: str,
			},
		}
		sel.applyCondition()
		L.Push(sel.toLuaTable())
		return 1
	}))

	// 注册 desc - 根据节点的desc属性匹配
	state.SetGlobal("desc", state.NewFunction(func(L *lua.LState) int {
		str := L.CheckString(1)
		sel := &Selector{
			acc:      uiacc.New(0),
			luaState: L,
			module:   m,
			condition: &SelectorCondition{
				desc: str,
			},
		}
		sel.applyCondition()
		L.Push(sel.toLuaTable())
		return 1
	}))

	// 注册 descContains - 根据desc包含的部分字符串匹配
	state.SetGlobal("descContains", state.NewFunction(func(L *lua.LState) int {
		str := L.CheckString(1)
		sel := &Selector{
			acc:      uiacc.New(0),
			luaState: L,
			module:   m,
			condition: &SelectorCondition{
				descContains: str,
			},
		}
		sel.applyCondition()
		L.Push(sel.toLuaTable())
		return 1
	}))

	// 注册 descStartsWith - 根据text的前缀去匹配
	state.SetGlobal("descStartsWith", state.NewFunction(func(L *lua.LState) int {
		str := L.CheckString(1)
		sel := &Selector{
			acc:      uiacc.New(0),
			luaState: L,
			module:   m,
			condition: &SelectorCondition{
				descStartsWith: str,
			},
		}
		sel.applyCondition()
		L.Push(sel.toLuaTable())
		return 1
	}))

	// 注册 descEndsWith - 根据text的后缀去匹配
	state.SetGlobal("descEndsWith", state.NewFunction(func(L *lua.LState) int {
		str := L.CheckString(1)
		sel := &Selector{
			acc:      uiacc.New(0),
			luaState: L,
			module:   m,
			condition: &SelectorCondition{
				descEndsWith: str,
			},
		}
		sel.applyCondition()
		L.Push(sel.toLuaTable())
		return 1
	}))

	// 注册 descMatches - 正则匹配text
	state.SetGlobal("descMatches", state.NewFunction(func(L *lua.LState) int {
		str := L.CheckString(1)
		sel := &Selector{
			acc:      uiacc.New(0),
			luaState: L,
			module:   m,
			condition: &SelectorCondition{
				descMatches: str,
			},
		}
		sel.applyCondition()
		L.Push(sel.toLuaTable())
		return 1
	}))

	// 注册 className - 根据节点的className属性匹配
	state.SetGlobal("className", state.NewFunction(func(L *lua.LState) int {
		str := L.CheckString(1)
		sel := &Selector{
			acc:      uiacc.New(0),
			luaState: L,
			module:   m,
			condition: &SelectorCondition{
				className: str,
			},
		}
		sel.applyCondition()
		L.Push(sel.toLuaTable())
		return 1
	}))

	// 注册 classNameContains - 根据className属性所包含的字符串模糊匹配
	state.SetGlobal("classNameContains", state.NewFunction(func(L *lua.LState) int {
		str := L.CheckString(1)
		sel := &Selector{
			acc:      uiacc.New(0),
			luaState: L,
			module:   m,
			condition: &SelectorCondition{
				classNameContains: str,
			},
		}
		sel.applyCondition()
		L.Push(sel.toLuaTable())
		return 1
	}))

	// 注册 classNameStartsWith - 根据className属性前缀匹配
	state.SetGlobal("classNameStartsWith", state.NewFunction(func(L *lua.LState) int {
		str := L.CheckString(1)
		sel := &Selector{
			acc:      uiacc.New(0),
			luaState: L,
			module:   m,
			condition: &SelectorCondition{
				classNameStartsWith: str,
			},
		}
		sel.applyCondition()
		L.Push(sel.toLuaTable())
		return 1
	}))

	// 注册 classNameEndsWith - 根据className属性后缀匹配
	state.SetGlobal("classNameEndsWith", state.NewFunction(func(L *lua.LState) int {
		str := L.CheckString(1)
		sel := &Selector{
			acc:      uiacc.New(0),
			luaState: L,
			module:   m,
			condition: &SelectorCondition{
				classNameEndsWith: str,
			},
		}
		sel.applyCondition()
		L.Push(sel.toLuaTable())
		return 1
	}))

	// 注册 classNameMatches - 根据className属性正则匹配
	state.SetGlobal("classNameMatches", state.NewFunction(func(L *lua.LState) int {
		str := L.CheckString(1)
		sel := &Selector{
			acc:      uiacc.New(0),
			luaState: L,
			module:   m,
			condition: &SelectorCondition{
				classNameMatches: str,
			},
		}
		sel.applyCondition()
		L.Push(sel.toLuaTable())
		return 1
	}))

	// 注册 packageName - 根据packageName属性全字段匹配
	state.SetGlobal("packageName", state.NewFunction(func(L *lua.LState) int {
		str := L.CheckString(1)
		sel := &Selector{
			acc:      uiacc.New(0),
			luaState: L,
			module:   m,
			condition: &SelectorCondition{
				packageName: str,
			},
		}
		sel.applyCondition()
		L.Push(sel.toLuaTable())
		return 1
	}))

	// 注册 packageNameContains - 匹配包含指定字符串的的节点
	state.SetGlobal("packageNameContains", state.NewFunction(func(L *lua.LState) int {
		str := L.CheckString(1)
		sel := &Selector{
			acc:      uiacc.New(0),
			luaState: L,
			module:   m,
			condition: &SelectorCondition{
				packageNameContains: str,
			},
		}
		sel.applyCondition()
		L.Push(sel.toLuaTable())
		return 1
	}))

	// 注册 packageNameStartsWith - 匹配包名前缀为指定字符串的的节点
	state.SetGlobal("packageNameStartsWith", state.NewFunction(func(L *lua.LState) int {
		str := L.CheckString(1)
		sel := &Selector{
			acc:      uiacc.New(0),
			luaState: L,
			module:   m,
			condition: &SelectorCondition{
				packageNameStartsWith: str,
			},
		}
		sel.applyCondition()
		L.Push(sel.toLuaTable())
		return 1
	}))

	// 注册 packageNameEndsWith - 匹配包名后缀为指定字符串的的节点
	state.SetGlobal("packageNameEndsWith", state.NewFunction(func(L *lua.LState) int {
		str := L.CheckString(1)
		sel := &Selector{
			acc:      uiacc.New(0),
			luaState: L,
			module:   m,
			condition: &SelectorCondition{
				packageNameEndsWith: str,
			},
		}
		sel.applyCondition()
		L.Push(sel.toLuaTable())
		return 1
	}))

	// 注册 packageNameMatches - 包名正则匹配
	state.SetGlobal("packageNameMatches", state.NewFunction(func(L *lua.LState) int {
		str := L.CheckString(1)
		sel := &Selector{
			acc:      uiacc.New(0),
			luaState: L,
			module:   m,
			condition: &SelectorCondition{
				packageNameMatches: str,
			},
		}
		sel.applyCondition()
		L.Push(sel.toLuaTable())
		return 1
	}))

	// 注册 bounds - 根据节点的范围匹配
	state.SetGlobal("bounds", state.NewFunction(func(L *lua.LState) int {
		l := L.CheckInt(1)
		t := L.CheckInt(2)
		r := L.CheckInt(3)
		b := L.CheckInt(4)
		sel := &Selector{
			acc:      uiacc.New(0),
			luaState: L,
			module:   m,
			condition: &SelectorCondition{
				bounds: &Bounds{
					Left:   l,
					Top:    t,
					Right:  r,
					Bottom: b,
				},
			},
		}
		sel.applyCondition()
		L.Push(sel.toLuaTable())
		return 1
	}))

	// 注册 boundsInside - 匹配该范围内的节点
	state.SetGlobal("boundsInside", state.NewFunction(func(L *lua.LState) int {
		l := L.CheckInt(1)
		t := L.CheckInt(2)
		r := L.CheckInt(3)
		b := L.CheckInt(4)
		sel := &Selector{
			acc:      uiacc.New(0),
			luaState: L,
			module:   m,
			condition: &SelectorCondition{
				boundsInside: &Bounds{
					Left:   l,
					Top:    t,
					Right:  r,
					Bottom: b,
				},
			},
		}
		sel.applyCondition()
		L.Push(sel.toLuaTable())
		return 1
	}))

	// 注册 drawingOrder - 根据绘制顺序匹配
	state.SetGlobal("drawingOrder", state.NewFunction(func(L *lua.LState) int {
		level := L.CheckInt(1)
		sel := &Selector{
			acc:      uiacc.New(0),
			luaState: L,
			module:   m,
			condition: &SelectorCondition{
				drawingOrder: level,
			},
		}
		sel.applyCondition()
		L.Push(sel.toLuaTable())
		return 1
	}))

	// 注册 depth - 根据深度匹配（通过遍历 UI 树实现）
	state.SetGlobal("depth", state.NewFunction(func(L *lua.LState) int {
		level := L.CheckInt(1)
		sel := &Selector{
			acc:      uiacc.New(0),
			luaState: L,
			module:   m,
			condition: &SelectorCondition{
				depth: level,
			},
		}
		sel.applyCondition()
		L.Push(sel.toLuaTable())
		return 1
	}))

	// 注册 index - 根据索引匹配
	state.SetGlobal("index", state.NewFunction(func(L *lua.LState) int {
		idx := L.CheckInt(1)
		sel := &Selector{
			acc:      uiacc.New(0),
			luaState: L,
			module:   m,
			condition: &SelectorCondition{
				index: idx,
			},
		}
		sel.applyCondition()
		L.Push(sel.toLuaTable())
		return 1
	}))

	// 注册 visibleToUser - 根据是否可见匹配
	state.SetGlobal("visibleToUser", state.NewFunction(func(L *lua.LState) int {
		b := L.CheckBool(1)
		sel := &Selector{
			acc:      uiacc.New(0),
			luaState: L,
			module:   m,
			condition: &SelectorCondition{
				visibleToUser: &b,
			},
		}
		sel.applyCondition()
		L.Push(sel.toLuaTable())
		return 1
	}))

	// 注册 selected - 根据是否选中匹配
	state.SetGlobal("selected", state.NewFunction(func(L *lua.LState) int {
		b := L.CheckBool(1)
		sel := &Selector{
			acc:      uiacc.New(0),
			luaState: L,
			module:   m,
			condition: &SelectorCondition{
				selected: &b,
			},
		}
		sel.applyCondition()
		L.Push(sel.toLuaTable())
		return 1
	}))

	// 注册 clickable - 根据是否可点击匹配
	state.SetGlobal("clickable", state.NewFunction(func(L *lua.LState) int {
		b := L.CheckBool(1)
		sel := &Selector{
			acc:      uiacc.New(0),
			luaState: L,
			module:   m,
			condition: &SelectorCondition{
				clickable: &b,
			},
		}
		sel.applyCondition()
		L.Push(sel.toLuaTable())
		return 1
	}))

	// 注册 longClickable - 根据是否可长按点击匹配
	state.SetGlobal("longClickable", state.NewFunction(func(L *lua.LState) int {
		b := L.CheckBool(1)
		sel := &Selector{
			acc:      uiacc.New(0),
			luaState: L,
			module:   m,
			condition: &SelectorCondition{
				longClickable: &b,
			},
		}
		sel.applyCondition()
		L.Push(sel.toLuaTable())
		return 1
	}))

	// 注册 enabled - 根据是否可用匹配
	state.SetGlobal("enabled", state.NewFunction(func(L *lua.LState) int {
		b := L.CheckBool(1)
		sel := &Selector{
			acc:      uiacc.New(0),
			luaState: L,
			module:   m,
			condition: &SelectorCondition{
				enabled: &b,
			},
		}
		sel.applyCondition()
		L.Push(sel.toLuaTable())
		return 1
	}))

	// 注册 password - 根据是否是密码框匹配
	state.SetGlobal("password", state.NewFunction(func(L *lua.LState) int {
		b := L.CheckBool(1)
		sel := &Selector{
			acc:      uiacc.New(0),
			luaState: L,
			module:   m,
			condition: &SelectorCondition{
				password: &b,
			},
		}
		sel.applyCondition()
		L.Push(sel.toLuaTable())
		return 1
	}))

	// 注册 scrollable - 根据是否可以滚动来匹配
	state.SetGlobal("scrollable", state.NewFunction(func(L *lua.LState) int {
		b := L.CheckBool(1)
		sel := &Selector{
			acc:      uiacc.New(0),
			luaState: L,
			module:   m,
			condition: &SelectorCondition{
				scrollable: &b,
			},
		}
		sel.applyCondition()
		L.Push(sel.toLuaTable())
		return 1
	}))

	// 注册 checked - 是否被勾选来匹配
	state.SetGlobal("checked", state.NewFunction(func(L *lua.LState) int {
		b := L.CheckBool(1)
		sel := &Selector{
			acc:      uiacc.New(0),
			luaState: L,
			module:   m,
			condition: &SelectorCondition{
				checked: &b,
			},
		}
		sel.applyCondition()
		L.Push(sel.toLuaTable())
		return 1
	}))

	// 注册 checkable - 是否可以被勾选来匹配
	state.SetGlobal("checkable", state.NewFunction(func(L *lua.LState) int {
		b := L.CheckBool(1)
		sel := &Selector{
			acc:      uiacc.New(0),
			luaState: L,
			module:   m,
			condition: &SelectorCondition{
				checkable: &b,
			},
		}
		sel.applyCondition()
		L.Push(sel.toLuaTable())
		return 1
	}))

	// 注册 focusable - 根据是否允许抢占焦点来匹配
	state.SetGlobal("focusable", state.NewFunction(func(L *lua.LState) int {
		b := L.CheckBool(1)
		sel := &Selector{
			acc:      uiacc.New(0),
			luaState: L,
			module:   m,
			condition: &SelectorCondition{
				focusable: &b,
			},
		}
		sel.applyCondition()
		L.Push(sel.toLuaTable())
		return 1
	}))

	// 注册 focused - 根据是否抢占了焦点来匹配
	state.SetGlobal("focused", state.NewFunction(func(L *lua.LState) int {
		b := L.CheckBool(1)
		sel := &Selector{
			acc:      uiacc.New(0),
			luaState: L,
			module:   m,
			condition: &SelectorCondition{
				focused: &b,
			},
		}
		sel.applyCondition()
		L.Push(sel.toLuaTable())
		return 1
	}))
}

// applyCondition 应用选择器条件
func (s *Selector) applyCondition() {
	c := s.condition

	// 应用 ID 条件
	if c.id != "" {
		s.acc.Id(c.id)
	}
	if c.idContains != "" {
		s.acc.IdContains(c.idContains)
	}
	if c.idStartsWith != "" {
		s.acc.IdStartsWith(c.idStartsWith)
	}
	if c.idEndsWith != "" {
		s.acc.IdEndsWith(c.idEndsWith)
	}
	if c.idMatches != "" {
		s.acc.IdMatches(c.idMatches)
	}

	// 应用 Text 条件
	if c.text != "" {
		s.acc.Text(c.text)
	}
	if c.textContains != "" {
		s.acc.TextContains(c.textContains)
	}
	if c.textStartsWith != "" {
		s.acc.TextStartsWith(c.textStartsWith)
	}
	if c.textEndsWith != "" {
		s.acc.TextEndsWith(c.textEndsWith)
	}
	if c.textMatches != "" {
		s.acc.TextMatches(c.textMatches)
	}

	// 应用 Desc 条件
	if c.desc != "" {
		s.acc.Desc(c.desc)
	}
	if c.descContains != "" {
		s.acc.DescContains(c.descContains)
	}
	if c.descStartsWith != "" {
		s.acc.DescStartsWith(c.descStartsWith)
	}
	if c.descEndsWith != "" {
		s.acc.DescEndsWith(c.descEndsWith)
	}
	if c.descMatches != "" {
		s.acc.DescMatches(c.descMatches)
	}

	// 应用 ClassName 条件
	if c.className != "" {
		s.acc.ClassName(c.className)
	}
	if c.classNameContains != "" {
		s.acc.ClassNameContains(c.classNameContains)
	}
	if c.classNameStartsWith != "" {
		s.acc.ClassNameStartsWith(c.classNameStartsWith)
	}
	if c.classNameEndsWith != "" {
		s.acc.ClassNameEndsWith(c.classNameEndsWith)
	}
	if c.classNameMatches != "" {
		s.acc.ClassNameMatches(c.classNameMatches)
	}

	// 应用 PackageName 条件
	if c.packageName != "" {
		s.acc.PackageName(c.packageName)
	}
	if c.packageNameContains != "" {
		s.acc.PackageNameContains(c.packageNameContains)
	}
	if c.packageNameStartsWith != "" {
		s.acc.PackageNameStartsWith(c.packageNameStartsWith)
	}
	if c.packageNameEndsWith != "" {
		s.acc.PackageNameEndsWith(c.packageNameEndsWith)
	}
	if c.packageNameMatches != "" {
		s.acc.PackageNameMatches(c.packageNameMatches)
	}

	// 应用 Bounds 条件
	if c.bounds != nil {
		s.acc.Bounds(c.bounds.Left, c.bounds.Top, c.bounds.Right, c.bounds.Bottom)
	}
	if c.boundsInside != nil {
		s.acc.BoundsInside(c.boundsInside.Left, c.boundsInside.Top, c.boundsInside.Right, c.boundsInside.Bottom)
	}

	// 应用其他条件
	if c.drawingOrder != 0 {
		s.acc.DrawingOrder(c.drawingOrder)
	}
	if c.index != 0 {
		s.acc.Index(c.index)
	}
	if c.visibleToUser != nil {
		s.acc.Visible(*c.visibleToUser)
	}
	if c.selected != nil {
		s.acc.Selected(*c.selected)
	}
	if c.clickable != nil {
		s.acc.Clickable(*c.clickable)
	}
	if c.longClickable != nil {
		s.acc.LongClickable(*c.longClickable)
	}
	if c.enabled != nil {
		s.acc.Enabled(*c.enabled)
	}
	if c.password != nil {
		s.acc.Password(*c.password)
	}
	if c.scrollable != nil {
		s.acc.Scrollable(*c.scrollable)
	}
	if c.checked != nil {
		s.acc.Checked(*c.checked)
	}
	if c.checkable != nil {
		s.acc.Checkable(*c.checkable)
	}
	if c.focusable != nil {
		s.acc.Focusable(*c.focusable)
	}
	if c.focused != nil {
		s.acc.Focused(*c.focused)
	}

	// depth 在查找时通过 calculateDepth 实现，这里不需要调用 AutoGo API
}

// toLuaTable 将选择器转换为 Lua 表
func (s *Selector) toLuaTable() *lua.LTable {
	table := s.luaState.NewTable()

	// 注册 findOne 方法
	table.RawSetString("findOne", s.luaState.NewFunction(func(L *lua.LState) int {
		// 支持冒号语法：sel:findOne(timeout)
		// 检查第一个参数是否是选择器表本身
		if L.GetTop() >= 1 && L.CheckAny(1).Type() == lua.LTTable {
			_ = L.CheckTable(1) // 跳过 self 参数
		}

		// timeout 参数预留
		_ = 10000
		if L.GetTop() >= 2 {
			_ = L.CheckInt(2) // timeout 参数预留
		}

		uiObj := s.acc.FindOnce()
		if uiObj == nil {
			L.Push(lua.LNil)
			return 1
		}

		// 如果设置了 depth 条件，检查节点深度
		if s.condition.depth != 0 {
			nodeDepth := s.module.calculateDepth(uiObj)
			if nodeDepth != s.condition.depth {
				L.Push(lua.LNil)
				return 1
			}
		}

		// 缓存 UIObject 并返回 ID
		id := s.module.cacheUiObject(uiObj)
		result := s.luaState.NewTable()
		result.RawSetString("__node_id__", lua.LNumber(id))
		s.module.WrapNodeTable(s.luaState, result)
		L.Push(result)
		return 1
	}))

	// 注册 findAll 方法
	table.RawSetString("findAll", s.luaState.NewFunction(func(L *lua.LState) int {
		// 支持冒号语法：sel:findAll(timeout)
		// 检查第一个参数是否是选择器表本身
		if L.GetTop() >= 1 && L.CheckAny(1).Type() == lua.LTTable {
			_ = L.CheckTable(1) // 跳过 self 参数
		}

		// timeout 参数预留
		_ = 10000
		if L.GetTop() >= 2 {
			_ = L.CheckInt(2) // timeout 参数预留
		}

		uiObjs := s.acc.Find()
		if uiObjs == nil || len(uiObjs) == 0 {
			L.Push(lua.LNil)
			return 1
		}

		// 缓存所有 UIObject 并返回 ID 数组
		result := s.luaState.NewTable()
		count := 0
		for _, uiObj := range uiObjs {
			// 如果设置了 depth 条件，检查节点深度
			if s.condition.depth != 0 {
				nodeDepth := s.module.calculateDepth(uiObj)
				if nodeDepth != s.condition.depth {
					continue
				}
			}

			id := s.module.cacheUiObject(uiObj)
			nodeTable := s.luaState.NewTable()
			nodeTable.RawSetString("__node_id__", lua.LNumber(id))
			s.module.WrapNodeTable(s.luaState, nodeTable)
			count++
			result.RawSetInt(count, nodeTable)
		}

		if count == 0 {
			L.Push(lua.LNil)
		} else {
			L.Push(result)
		}
		return 1
	}))

	// 注册 findOnce 方法
	table.RawSetString("findOnce", s.luaState.NewFunction(func(L *lua.LState) int {
		// 支持冒号语法：sel:findOnce(index)
		// 检查第一个参数是否是选择器表本身
		if L.GetTop() >= 1 && L.CheckAny(1).Type() == lua.LTTable {
			_ = L.CheckTable(1) // 跳过 self 参数
		}

		index := 0
		if L.GetTop() >= 2 {
			index = L.CheckInt(2)
		}

		uiObjs := s.acc.Find()
		if uiObjs == nil || len(uiObjs) == 0 {
			L.Push(lua.LNil)
			return 1
		}

		// 如果设置了 depth 条件，过滤节点
		filteredObjs := uiObjs
		if s.condition.depth != 0 {
			filteredObjs = make([]*uiacc.UiObject, 0)
			for _, uiObj := range uiObjs {
				nodeDepth := s.module.calculateDepth(uiObj)
				if nodeDepth == s.condition.depth {
					filteredObjs = append(filteredObjs, uiObj)
				}
			}
		}

		if len(filteredObjs) == 0 || index >= len(filteredObjs) {
			L.Push(lua.LNil)
			return 1
		}

		// 缓存 UIObject 并返回 ID
		id := s.module.cacheUiObject(filteredObjs[index])
		result := s.luaState.NewTable()
		result.RawSetString("__node_id__", lua.LNumber(id))
		s.module.WrapNodeTable(s.luaState, result)
		L.Push(result)
		return 1
	}))

	// 注册 click 方法
	table.RawSetString("click", s.luaState.NewFunction(func(L *lua.LState) int {
		// 支持冒号语法：sel:click()
		// 检查第一个参数是否是选择器表本身
		if L.GetTop() >= 1 && L.CheckAny(1).Type() == lua.LTTable {
			_ = L.CheckTable(1) // 跳过 self 参数
		}

		// 只获取第一个匹配的节点
		uiObj := s.acc.FindOnce()
		if uiObj == nil {
			L.Push(lua.LBool(false))
			return 1
		}

		// 如果设置了 depth 条件，检查节点深度
		if s.condition.depth != 0 {
			nodeDepth := s.module.calculateDepth(uiObj)
			if nodeDepth != s.condition.depth {
				L.Push(lua.LBool(false))
				return 1
			}
		}

		// 使用 uiacc.UiObject 的 Click 方法点击节点
		success := uiObj.Click()
		L.Push(lua.LBool(success))
		return 1
	}))

	// 注册 longClick 方法
	table.RawSetString("longClick", s.luaState.NewFunction(func(L *lua.LState) int {
		// 支持冒号语法：sel:longClick()
		// 检查第一个参数是否是选择器表本身
		if L.GetTop() >= 1 && L.CheckAny(1).Type() == lua.LTTable {
			_ = L.CheckTable(1) // 跳过 self 参数
		}

		// 只获取第一个匹配的节点
		uiObj := s.acc.FindOnce()
		if uiObj == nil {
			L.Push(lua.LBool(false))
			return 1
		}

		// 如果设置了 depth 条件，检查节点深度
		if s.condition.depth != 0 {
			nodeDepth := s.module.calculateDepth(uiObj)
			if nodeDepth != s.condition.depth {
				L.Push(lua.LBool(false))
				return 1
			}
		}

		// 使用 uiacc.UiObject 的 ClickLongClick 方法长按节点
		success := uiObj.ClickLongClick()
		L.Push(lua.LBool(success))
		return 1
	}))

	return table
}

// 注册节点方法到 Lua 表
func (m *NodeModule) registerNodeMethods(L *lua.LState, table *lua.LTable, id int64) {
	// 获取 ID
	idValue := L.GetField(table, "__node_id__")
	if idValue == lua.LNil {
		return
	}

	nodeId := int64(idValue.(lua.LNumber))

	// 注册 id 方法
	table.RawSetString("id", L.NewFunction(func(L *lua.LState) int {
		// 支持冒号语法：node:id()
		// 检查第一个参数是否是节点表本身
		if L.GetTop() >= 1 && L.CheckAny(1).Type() == lua.LTTable {
			_ = L.CheckTable(1) // 跳过 self 参数
		}

		uiObj := m.getUiObject(nodeId)
		if uiObj == nil {
			L.Push(lua.LString(""))
			return 1
		}

		L.Push(lua.LString(uiObj.GetId()))
		return 1
	}))

	// 注册 toJson 方法
	table.RawSetString("toJson", L.NewFunction(func(L *lua.LState) int {
		// 支持冒号语法：node:toJson()
		// 检查第一个参数是否是节点表本身
		if L.GetTop() >= 1 && L.CheckAny(1).Type() == lua.LTTable {
			_ = L.CheckTable(1) // 跳过 self 参数
		}

		uiObj := m.getUiObject(nodeId)
		if uiObj == nil {
			L.Push(lua.LString("{}"))
			return 1
		}

		jsonStr := fmt.Sprintf(`{
    "id": "%s",
    "text": "%s",
    "desc": "%s",
    "className": "%s",
    "packageName": "%s"
}`, uiObj.GetId(), uiObj.GetText(), uiObj.GetDesc(), uiObj.GetClassName(), uiObj.GetPackageName())

		L.Push(lua.LString(jsonStr))
		return 1
	}))

	// 注册 text 方法
	table.RawSetString("text", L.NewFunction(func(L *lua.LState) int {
		// 支持冒号语法：node:text()
		// 检查第一个参数是否是节点表本身
		if L.GetTop() >= 1 && L.CheckAny(1).Type() == lua.LTTable {
			_ = L.CheckTable(1) // 跳过 self 参数
		}

		uiObj := m.getUiObject(nodeId)
		if uiObj == nil {
			L.Push(lua.LString(""))
			return 1
		}

		L.Push(lua.LString(uiObj.GetText()))
		return 1
	}))

	// 注册 click 方法
	table.RawSetString("click", L.NewFunction(func(L *lua.LState) int {
		// 支持冒号语法：node:click()
		// 检查第一个参数是否是节点表本身
		if L.GetTop() >= 1 && L.CheckAny(1).Type() == lua.LTTable {
			_ = L.CheckTable(1) // 跳过 self 参数
		}

		uiObj := m.getUiObject(nodeId)
		if uiObj == nil {
			L.Push(lua.LBool(false))
			return 1
		}

		// 使用 uiacc.UiObject 的 Click 方法点击节点
		success := uiObj.Click()
		L.Push(lua.LBool(success))
		return 1
	}))

	// 注册 longClick 方法
	table.RawSetString("longClick", L.NewFunction(func(L *lua.LState) int {
		// 支持冒号语法：node:longClick()
		// 检查第一个参数是否是节点表本身
		if L.GetTop() >= 1 && L.CheckAny(1).Type() == lua.LTTable {
			_ = L.CheckTable(1) // 跳过 self 参数
		}

		uiObj := m.getUiObject(nodeId)
		if uiObj == nil {
			L.Push(lua.LBool(false))
			return 1
		}

		// 使用 uiacc.UiObject 的 ClickLongClick 方法长按节点
		success := uiObj.ClickLongClick()
		L.Push(lua.LBool(success))
		return 1
	}))

	// 注册 className 方法
	table.RawSetString("className", L.NewFunction(func(L *lua.LState) int {
		// 支持冒号语法：node:className()
		// 检查第一个参数是否是节点表本身
		if L.GetTop() >= 1 && L.CheckAny(1).Type() == lua.LTTable {
			_ = L.CheckTable(1) // 跳过 self 参数
		}

		uiObj := m.getUiObject(nodeId)
		if uiObj == nil {
			L.Push(lua.LString(""))
			return 1
		}

		L.Push(lua.LString(uiObj.GetClassName()))
		return 1
	}))

	// 注册 packageName 方法
	table.RawSetString("packageName", L.NewFunction(func(L *lua.LState) int {
		// 支持冒号语法：node:packageName()
		// 检查第一个参数是否是节点表本身
		if L.GetTop() >= 1 && L.CheckAny(1).Type() == lua.LTTable {
			_ = L.CheckTable(1) // 跳过 self 参数
		}

		uiObj := m.getUiObject(nodeId)
		if uiObj == nil {
			L.Push(lua.LString(""))
			return 1
		}

		L.Push(lua.LString(uiObj.GetPackageName()))
		return 1
	}))

	// 注册 desc 方法
	table.RawSetString("desc", L.NewFunction(func(L *lua.LState) int {
		// 支持冒号语法：node:desc()
		// 检查第一个参数是否是节点表本身
		if L.GetTop() >= 1 && L.CheckAny(1).Type() == lua.LTTable {
			_ = L.CheckTable(1) // 跳过 self 参数
		}

		uiObj := m.getUiObject(nodeId)
		if uiObj == nil {
			L.Push(lua.LString(""))
			return 1
		}

		L.Push(lua.LString(uiObj.GetDesc()))
		return 1
	}))

	// 注册 bounds 方法
	table.RawSetString("bounds", L.NewFunction(func(L *lua.LState) int {
		// 支持冒号语法：node:bounds()
		// 检查第一个参数是否是节点表本身
		if L.GetTop() >= 1 && L.CheckAny(1).Type() == lua.LTTable {
			_ = L.CheckTable(1) // 跳过 self 参数
		}

		uiObj := m.getUiObject(nodeId)
		if uiObj == nil {
			L.Push(lua.LNil)
			return 1
		}

		bounds := uiObj.GetBounds()
		boundsTable := L.NewTable()
		boundsTable.RawSetString("left", lua.LNumber(bounds.Left))
		boundsTable.RawSetString("top", lua.LNumber(bounds.Top))
		boundsTable.RawSetString("right", lua.LNumber(bounds.Right))
		boundsTable.RawSetString("bottom", lua.LNumber(bounds.Bottom))
		boundsTable.RawSetString("width", lua.LNumber(bounds.Width))
		boundsTable.RawSetString("height", lua.LNumber(bounds.Height))
		boundsTable.RawSetString("centerX", lua.LNumber(bounds.CenterX))
		boundsTable.RawSetString("centerY", lua.LNumber(bounds.CenterY))
		L.Push(boundsTable)
		return 1
	}))
}

// WrapNodeTable 包装节点表，添加节点方法
func (m *NodeModule) WrapNodeTable(L *lua.LState, table *lua.LTable) {
	// 获取 ID
	idValue := L.GetField(table, "__node_id__")
	if idValue == lua.LNil {
		return
	}

	nodeId := int64(idValue.(lua.LNumber))
	m.registerNodeMethods(L, table, nodeId)
}

// Name 返回模块名称
func (m *NodeModule) Name() string {
	return "node"
}

// Register 向引擎注册模块的方法
func (m *NodeModule) Register(engine model.Engine) error {
	state := engine.GetState()

	// 创建 node 表
	nodeTable := state.NewTable()

	// 注册 node.new - 创建节点选择器
	nodeTable.RawSetString("new", state.NewFunction(func(L *lua.LState) int {
		sel := &Selector{
			acc:       uiacc.New(0),
			luaState:  L,
			module:    m,
			condition: &SelectorCondition{},
		}
		L.Push(sel.toLuaTable())
		return 1
	}))

	// 注册 node 表到全局
	state.SetGlobal("node", nodeTable)

	m.Inject(state)
	return nil
}

// IsAvailable 返回模块是否可用
func (m *NodeModule) IsAvailable() bool {
	return true
}

// GetModule 获取模块实例
func GetModule() model.Module {
	return New()
}
