package uiacc

import (
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"

	"github.com/Dasongzi1366/AutoGo/uiacc"
	lua "github.com/yuin/gopher-lua"
)

// UiaccModule uiacc 模块
type UiaccModule struct{}

// Name 返回模块名称
func (m *UiaccModule) Name() string {
	return "uiacc"
}

// IsAvailable 检查模块是否可用
func (m *UiaccModule) IsAvailable() bool {
	return true
}

// Register 向引擎注册方法
func (m *UiaccModule) Register(engine model.Engine) error {
	state := engine.GetState()

	uiaccObj := state.NewTable()
	state.SetGlobal("uiacc", uiaccObj)

	uiaccObj.RawSetString("new", state.NewFunction(func(L *lua.LState) int {
		displayId := L.OptInt(1, 0)
		L.Push(wrapUiacc(L, uiacc.New(displayId)))
		return 1
	}))

	engine.RegisterMethod("uiacc.new", "创建新的Uiacc对象", uiacc.New, true)
	engine.RegisterMethod("uiacc.text", "Uiacc.Text", (*uiacc.Uiacc).Text, true)
	engine.RegisterMethod("uiacc.textContains", "Uiacc.TextContains", (*uiacc.Uiacc).TextContains, true)
	engine.RegisterMethod("uiacc.textStartsWith", "Uiacc.TextStartsWith", (*uiacc.Uiacc).TextStartsWith, true)
	engine.RegisterMethod("uiacc.textEndsWith", "Uiacc.TextEndsWith", (*uiacc.Uiacc).TextEndsWith, true)
	engine.RegisterMethod("uiacc.textMatches", "Uiacc.TextMatches", (*uiacc.Uiacc).TextMatches, true)
	engine.RegisterMethod("uiacc.desc", "Uiacc.Desc", (*uiacc.Uiacc).Desc, true)
	engine.RegisterMethod("uiacc.descContains", "Uiacc.DescContains", (*uiacc.Uiacc).DescContains, true)
	engine.RegisterMethod("uiacc.descStartsWith", "Uiacc.DescStartsWith", (*uiacc.Uiacc).DescStartsWith, true)
	engine.RegisterMethod("uiacc.descEndsWith", "Uiacc.DescEndsWith", (*uiacc.Uiacc).DescEndsWith, true)
	engine.RegisterMethod("uiacc.descMatches", "Uiacc.DescMatches", (*uiacc.Uiacc).DescMatches, true)
	engine.RegisterMethod("uiacc.id", "Uiacc.Id", (*uiacc.Uiacc).Id, true)
	engine.RegisterMethod("uiacc.idContains", "Uiacc.IdContains", (*uiacc.Uiacc).IdContains, true)
	engine.RegisterMethod("uiacc.idStartsWith", "Uiacc.IdStartsWith", (*uiacc.Uiacc).IdStartsWith, true)
	engine.RegisterMethod("uiacc.idEndsWith", "Uiacc.IdEndsWith", (*uiacc.Uiacc).IdEndsWith, true)
	engine.RegisterMethod("uiacc.idMatches", "Uiacc.IdMatches", (*uiacc.Uiacc).IdMatches, true)
	engine.RegisterMethod("uiacc.className", "Uiacc.ClassName", (*uiacc.Uiacc).ClassName, true)
	engine.RegisterMethod("uiacc.classNameContains", "Uiacc.ClassNameContains", (*uiacc.Uiacc).ClassNameContains, true)
	engine.RegisterMethod("uiacc.classNameStartsWith", "Uiacc.ClassNameStartsWith", (*uiacc.Uiacc).ClassNameStartsWith, true)
	engine.RegisterMethod("uiacc.classNameEndsWith", "Uiacc.ClassNameEndsWith", (*uiacc.Uiacc).ClassNameEndsWith, true)
	engine.RegisterMethod("uiacc.classNameMatches", "Uiacc.ClassNameMatches", (*uiacc.Uiacc).ClassNameMatches, true)
	engine.RegisterMethod("uiacc.packageName", "Uiacc.PackageName", (*uiacc.Uiacc).PackageName, true)
	engine.RegisterMethod("uiacc.packageNameContains", "Uiacc.PackageNameContains", (*uiacc.Uiacc).PackageNameContains, true)
	engine.RegisterMethod("uiacc.packageNameStartsWith", "Uiacc.PackageNameStartsWith", (*uiacc.Uiacc).PackageNameStartsWith, true)
	engine.RegisterMethod("uiacc.packageNameEndsWith", "Uiacc.PackageNameEndsWith", (*uiacc.Uiacc).PackageNameEndsWith, true)
	engine.RegisterMethod("uiacc.packageNameMatches", "Uiacc.PackageNameMatches", (*uiacc.Uiacc).PackageNameMatches, true)
	engine.RegisterMethod("uiacc.bounds", "Uiacc.Bounds", (*uiacc.Uiacc).Bounds, true)
	engine.RegisterMethod("uiacc.boundsInside", "Uiacc.BoundsInside", (*uiacc.Uiacc).BoundsInside, true)
	engine.RegisterMethod("uiacc.boundsContains", "Uiacc.BoundsContains", (*uiacc.Uiacc).BoundsContains, true)
	engine.RegisterMethod("uiacc.drawingOrder", "Uiacc.DrawingOrder", (*uiacc.Uiacc).DrawingOrder, true)
	engine.RegisterMethod("uiacc.index", "Uiacc.Index", (*uiacc.Uiacc).Index, true)
	engine.RegisterMethod("uiacc.clickable", "Uiacc.Clickable", (*uiacc.Uiacc).Clickable, true)
	engine.RegisterMethod("uiacc.longClickable", "Uiacc.LongClickable", (*uiacc.Uiacc).LongClickable, true)
	engine.RegisterMethod("uiacc.checkable", "Uiacc.Checkable", (*uiacc.Uiacc).Checkable, true)
	engine.RegisterMethod("uiacc.selected", "Uiacc.Selected", (*uiacc.Uiacc).Selected, true)
	engine.RegisterMethod("uiacc.enabled", "Uiacc.Enabled", (*uiacc.Uiacc).Enabled, true)
	engine.RegisterMethod("uiacc.scrollable", "Uiacc.Scrollable", (*uiacc.Uiacc).Scrollable, true)
	engine.RegisterMethod("uiacc.editable", "Uiacc.Editable", (*uiacc.Uiacc).Editable, true)
	engine.RegisterMethod("uiacc.multiLine", "Uiacc.MultiLine", (*uiacc.Uiacc).MultiLine, true)
	engine.RegisterMethod("uiacc.checked", "Uiacc.Checked", (*uiacc.Uiacc).Checked, true)
	engine.RegisterMethod("uiacc.focusable", "Uiacc.Focusable", (*uiacc.Uiacc).Focusable, true)
	engine.RegisterMethod("uiacc.dismissable", "Uiacc.Dismissable", (*uiacc.Uiacc).Dismissable, true)
	engine.RegisterMethod("uiacc.focused", "Uiacc.Focused", (*uiacc.Uiacc).Focused, true)
	engine.RegisterMethod("uiacc.contextClickable", "Uiacc.ContextClickable", (*uiacc.Uiacc).ContextClickable, true)
	engine.RegisterMethod("uiacc.visible", "Uiacc.Visible", (*uiacc.Uiacc).Visible, true)
	engine.RegisterMethod("uiacc.password", "Uiacc.Password", (*uiacc.Uiacc).Password, true)
	engine.RegisterMethod("uiacc.click", "Uiacc.Click", (*uiacc.Uiacc).Click, true)
	engine.RegisterMethod("uiacc.waitFor", "Uiacc.WaitFor", (*uiacc.Uiacc).WaitFor, true)
	engine.RegisterMethod("uiacc.findOnce", "Uiacc.FindOnce", (*uiacc.Uiacc).FindOnce, true)
	engine.RegisterMethod("uiacc.find", "Uiacc.Find", (*uiacc.Uiacc).Find, true)
	engine.RegisterMethod("uiacc.release", "Uiacc.Release", (*uiacc.Uiacc).Release, true)
	engine.RegisterMethod("uiacc.clickCenter", "UiObject.ClickCenter", (*uiacc.UiObject).ClickCenter, true)
	engine.RegisterMethod("uiacc.clickLongClick", "UiObject.ClickLongClick", (*uiacc.UiObject).ClickLongClick, true)
	engine.RegisterMethod("uiacc.copy", "UiObject.Copy", (*uiacc.UiObject).Copy, true)
	engine.RegisterMethod("uiacc.cut", "UiObject.Cut", (*uiacc.UiObject).Cut, true)
	engine.RegisterMethod("uiacc.paste", "UiObject.Paste", (*uiacc.UiObject).Paste, true)
	engine.RegisterMethod("uiacc.scrollForward", "UiObject.ScrollForward", (*uiacc.UiObject).ScrollForward, true)
	engine.RegisterMethod("uiacc.scrollBackward", "UiObject.ScrollBackward", (*uiacc.UiObject).ScrollBackward, true)
	engine.RegisterMethod("uiacc.collapse", "UiObject.Collapse", (*uiacc.UiObject).Collapse, true)
	engine.RegisterMethod("uiacc.expand", "UiObject.Expand", (*uiacc.UiObject).Expand, true)
	engine.RegisterMethod("uiacc.show", "UiObject.Show", (*uiacc.UiObject).Show, true)
	engine.RegisterMethod("uiacc.select", "UiObject.Select", (*uiacc.UiObject).Select, true)
	engine.RegisterMethod("uiacc.clearSelect", "UiObject.ClearSelect", (*uiacc.UiObject).ClearSelect, true)
	engine.RegisterMethod("uiacc.getClickable", "UiObject.GetClickable", (*uiacc.UiObject).GetClickable, true)
	engine.RegisterMethod("uiacc.getLongClickable", "UiObject.GetLongClickable", (*uiacc.UiObject).GetLongClickable, true)
	engine.RegisterMethod("uiacc.getCheckable", "UiObject.GetCheckable", (*uiacc.UiObject).GetCheckable, true)
	engine.RegisterMethod("uiacc.getSelected", "UiObject.GetSelected", (*uiacc.UiObject).GetSelected, true)
	engine.RegisterMethod("uiacc.getEnabled", "UiObject.GetEnabled", (*uiacc.UiObject).GetEnabled, true)
	engine.RegisterMethod("uiacc.getScrollable", "UiObject.GetScrollable", (*uiacc.UiObject).GetScrollable, true)
	engine.RegisterMethod("uiacc.getEditable", "UiObject.GetEditable", (*uiacc.UiObject).GetEditable, true)
	engine.RegisterMethod("uiacc.getMultiLine", "UiObject.GetMultiLine", (*uiacc.UiObject).GetMultiLine, true)
	engine.RegisterMethod("uiacc.getChecked", "UiObject.GetChecked", (*uiacc.UiObject).GetChecked, true)
	engine.RegisterMethod("uiacc.getFocused", "UiObject.GetFocused", (*uiacc.UiObject).GetFocused, true)
	engine.RegisterMethod("uiacc.getFocusable", "UiObject.GetFocusable", (*uiacc.UiObject).GetFocusable, true)
	engine.RegisterMethod("uiacc.getDismissable", "UiObject.GetDismissable", (*uiacc.UiObject).GetDismissable, true)
	engine.RegisterMethod("uiacc.getContextClickable", "UiObject.GetContextClickable", (*uiacc.UiObject).GetContextClickable, true)
	engine.RegisterMethod("uiacc.getVisible", "UiObject.GetVisible", (*uiacc.UiObject).GetVisible, true)
	engine.RegisterMethod("uiacc.getPassword", "UiObject.GetPassword", (*uiacc.UiObject).GetPassword, true)
	engine.RegisterMethod("uiacc.getAccessibilityFocused", "UiObject.GetAccessibilityFocused", (*uiacc.UiObject).GetAccessibilityFocused, true)
	engine.RegisterMethod("uiacc.getChildCount", "UiObject.GetChildCount", (*uiacc.UiObject).GetChildCount, true)
	engine.RegisterMethod("uiacc.getDrawingOrder", "UiObject.GetDrawingOrder", (*uiacc.UiObject).GetDrawingOrder, true)
	engine.RegisterMethod("uiacc.getIndex", "UiObject.GetIndex", (*uiacc.UiObject).GetIndex, true)
	engine.RegisterMethod("uiacc.getBounds", "UiObject.GetBounds", (*uiacc.UiObject).GetBounds, true)
	engine.RegisterMethod("uiacc.getBoundsInParent", "UiObject.GetBoundsInParent", (*uiacc.UiObject).GetBoundsInParent, true)
	engine.RegisterMethod("uiacc.getId", "UiObject.GetId", (*uiacc.UiObject).GetId, true)
	engine.RegisterMethod("uiacc.getText", "UiObject.GetText", (*uiacc.UiObject).GetText, true)
	engine.RegisterMethod("uiacc.getDesc", "UiObject.GetDesc", (*uiacc.UiObject).GetDesc, true)
	engine.RegisterMethod("uiacc.getPackageName", "UiObject.GetPackageName", (*uiacc.UiObject).GetPackageName, true)
	engine.RegisterMethod("uiacc.getClassName", "UiObject.GetClassName", (*uiacc.UiObject).GetClassName, true)
	engine.RegisterMethod("uiacc.toString", "UiObject.ToString", (*uiacc.UiObject).ToString, true)
	engine.RegisterMethod("uiacc.setSelection", "UiObject.SetSelection", (*uiacc.UiObject).SetSelection, true)
	engine.RegisterMethod("uiacc.setVisibleToUser", "UiObject.SetVisibleToUser", (*uiacc.UiObject).SetVisibleToUser, true)
	engine.RegisterMethod("uiacc.setText", "UiObject.SetText", (*uiacc.UiObject).SetText, true)
	engine.RegisterMethod("uiacc.getParent", "UiObject.GetParent", (*uiacc.UiObject).GetParent, true)
	engine.RegisterMethod("uiacc.getChild", "UiObject.GetChild", (*uiacc.UiObject).GetChild, true)
	engine.RegisterMethod("uiacc.getChildren", "UiObject.GetChildren", (*uiacc.UiObject).GetChildren, true)

	return nil
}

func wrapUiacc(L *lua.LState, u *uiacc.Uiacc) lua.LValue {
	if u == nil {
		return lua.LNil
	}
	obj := L.NewTable()
	obj.RawSetString("text", L.NewFunction(func(L *lua.LState) int {
		L.Push(wrapUiacc(L, u.Text(L.CheckString(1))))
		return 1
	}))
	obj.RawSetString("textContains", L.NewFunction(func(L *lua.LState) int {
		L.Push(wrapUiacc(L, u.TextContains(L.CheckString(1))))
		return 1
	}))
	obj.RawSetString("textStartsWith", L.NewFunction(func(L *lua.LState) int {
		L.Push(wrapUiacc(L, u.TextStartsWith(L.CheckString(1))))
		return 1
	}))
	obj.RawSetString("textEndsWith", L.NewFunction(func(L *lua.LState) int {
		L.Push(wrapUiacc(L, u.TextEndsWith(L.CheckString(1))))
		return 1
	}))
	obj.RawSetString("textMatches", L.NewFunction(func(L *lua.LState) int {
		L.Push(wrapUiacc(L, u.TextMatches(L.CheckString(1))))
		return 1
	}))
	obj.RawSetString("desc", L.NewFunction(func(L *lua.LState) int {
		L.Push(wrapUiacc(L, u.Desc(L.CheckString(1))))
		return 1
	}))
	obj.RawSetString("descContains", L.NewFunction(func(L *lua.LState) int {
		L.Push(wrapUiacc(L, u.DescContains(L.CheckString(1))))
		return 1
	}))
	obj.RawSetString("descStartsWith", L.NewFunction(func(L *lua.LState) int {
		L.Push(wrapUiacc(L, u.DescStartsWith(L.CheckString(1))))
		return 1
	}))
	obj.RawSetString("descEndsWith", L.NewFunction(func(L *lua.LState) int {
		L.Push(wrapUiacc(L, u.DescEndsWith(L.CheckString(1))))
		return 1
	}))
	obj.RawSetString("descMatches", L.NewFunction(func(L *lua.LState) int {
		L.Push(wrapUiacc(L, u.DescMatches(L.CheckString(1))))
		return 1
	}))
	obj.RawSetString("id", L.NewFunction(func(L *lua.LState) int {
		L.Push(wrapUiacc(L, u.Id(L.CheckString(1))))
		return 1
	}))
	obj.RawSetString("idContains", L.NewFunction(func(L *lua.LState) int {
		L.Push(wrapUiacc(L, u.IdContains(L.CheckString(1))))
		return 1
	}))
	obj.RawSetString("idStartsWith", L.NewFunction(func(L *lua.LState) int {
		L.Push(wrapUiacc(L, u.IdStartsWith(L.CheckString(1))))
		return 1
	}))
	obj.RawSetString("idEndsWith", L.NewFunction(func(L *lua.LState) int {
		L.Push(wrapUiacc(L, u.IdEndsWith(L.CheckString(1))))
		return 1
	}))
	obj.RawSetString("idMatches", L.NewFunction(func(L *lua.LState) int {
		L.Push(wrapUiacc(L, u.IdMatches(L.CheckString(1))))
		return 1
	}))
	obj.RawSetString("className", L.NewFunction(func(L *lua.LState) int {
		L.Push(wrapUiacc(L, u.ClassName(L.CheckString(1))))
		return 1
	}))
	obj.RawSetString("classNameContains", L.NewFunction(func(L *lua.LState) int {
		L.Push(wrapUiacc(L, u.ClassNameContains(L.CheckString(1))))
		return 1
	}))
	obj.RawSetString("classNameStartsWith", L.NewFunction(func(L *lua.LState) int {
		L.Push(wrapUiacc(L, u.ClassNameStartsWith(L.CheckString(1))))
		return 1
	}))
	obj.RawSetString("classNameEndsWith", L.NewFunction(func(L *lua.LState) int {
		L.Push(wrapUiacc(L, u.ClassNameEndsWith(L.CheckString(1))))
		return 1
	}))
	obj.RawSetString("classNameMatches", L.NewFunction(func(L *lua.LState) int {
		L.Push(wrapUiacc(L, u.ClassNameMatches(L.CheckString(1))))
		return 1
	}))
	obj.RawSetString("packageName", L.NewFunction(func(L *lua.LState) int {
		L.Push(wrapUiacc(L, u.PackageName(L.CheckString(1))))
		return 1
	}))
	obj.RawSetString("packageNameContains", L.NewFunction(func(L *lua.LState) int {
		L.Push(wrapUiacc(L, u.PackageNameContains(L.CheckString(1))))
		return 1
	}))
	obj.RawSetString("packageNameStartsWith", L.NewFunction(func(L *lua.LState) int {
		L.Push(wrapUiacc(L, u.PackageNameStartsWith(L.CheckString(1))))
		return 1
	}))
	obj.RawSetString("packageNameEndsWith", L.NewFunction(func(L *lua.LState) int {
		L.Push(wrapUiacc(L, u.PackageNameEndsWith(L.CheckString(1))))
		return 1
	}))
	obj.RawSetString("packageNameMatches", L.NewFunction(func(L *lua.LState) int {
		L.Push(wrapUiacc(L, u.PackageNameMatches(L.CheckString(1))))
		return 1
	}))
	obj.RawSetString("bounds", L.NewFunction(func(L *lua.LState) int {
		L.Push(wrapUiacc(L, u.Bounds(L.CheckInt(1), L.CheckInt(2), L.CheckInt(3), L.CheckInt(4))))
		return 1
	}))
	obj.RawSetString("boundsInside", L.NewFunction(func(L *lua.LState) int {
		L.Push(wrapUiacc(L, u.BoundsInside(L.CheckInt(1), L.CheckInt(2), L.CheckInt(3), L.CheckInt(4))))
		return 1
	}))
	obj.RawSetString("boundsContains", L.NewFunction(func(L *lua.LState) int {
		L.Push(wrapUiacc(L, u.BoundsContains(L.CheckInt(1), L.CheckInt(2), L.CheckInt(3), L.CheckInt(4))))
		return 1
	}))
	obj.RawSetString("drawingOrder", L.NewFunction(func(L *lua.LState) int {
		L.Push(wrapUiacc(L, u.DrawingOrder(L.CheckInt(1))))
		return 1
	}))
	obj.RawSetString("index", L.NewFunction(func(L *lua.LState) int {
		L.Push(wrapUiacc(L, u.Index(L.CheckInt(1))))
		return 1
	}))
	obj.RawSetString("clickable", L.NewFunction(func(L *lua.LState) int {
		L.Push(wrapUiacc(L, u.Clickable(L.CheckBool(1))))
		return 1
	}))
	obj.RawSetString("longClickable", L.NewFunction(func(L *lua.LState) int {
		L.Push(wrapUiacc(L, u.LongClickable(L.CheckBool(1))))
		return 1
	}))
	obj.RawSetString("checkable", L.NewFunction(func(L *lua.LState) int {
		L.Push(wrapUiacc(L, u.Checkable(L.CheckBool(1))))
		return 1
	}))
	obj.RawSetString("selected", L.NewFunction(func(L *lua.LState) int {
		L.Push(wrapUiacc(L, u.Selected(L.CheckBool(1))))
		return 1
	}))
	obj.RawSetString("enabled", L.NewFunction(func(L *lua.LState) int {
		L.Push(wrapUiacc(L, u.Enabled(L.CheckBool(1))))
		return 1
	}))
	obj.RawSetString("scrollable", L.NewFunction(func(L *lua.LState) int {
		L.Push(wrapUiacc(L, u.Scrollable(L.CheckBool(1))))
		return 1
	}))
	obj.RawSetString("editable", L.NewFunction(func(L *lua.LState) int {
		L.Push(wrapUiacc(L, u.Editable(L.CheckBool(1))))
		return 1
	}))
	obj.RawSetString("multiLine", L.NewFunction(func(L *lua.LState) int {
		L.Push(wrapUiacc(L, u.MultiLine(L.CheckBool(1))))
		return 1
	}))
	obj.RawSetString("checked", L.NewFunction(func(L *lua.LState) int {
		L.Push(wrapUiacc(L, u.Checked(L.CheckBool(1))))
		return 1
	}))
	obj.RawSetString("focusable", L.NewFunction(func(L *lua.LState) int {
		L.Push(wrapUiacc(L, u.Focusable(L.CheckBool(1))))
		return 1
	}))
	obj.RawSetString("dismissable", L.NewFunction(func(L *lua.LState) int {
		L.Push(wrapUiacc(L, u.Dismissable(L.CheckBool(1))))
		return 1
	}))
	obj.RawSetString("focused", L.NewFunction(func(L *lua.LState) int {
		L.Push(wrapUiacc(L, u.Focused(L.CheckBool(1))))
		return 1
	}))
	obj.RawSetString("contextClickable", L.NewFunction(func(L *lua.LState) int {
		L.Push(wrapUiacc(L, u.ContextClickable(L.CheckBool(1))))
		return 1
	}))
	obj.RawSetString("visible", L.NewFunction(func(L *lua.LState) int {
		L.Push(wrapUiacc(L, u.Visible(L.CheckBool(1))))
		return 1
	}))
	obj.RawSetString("password", L.NewFunction(func(L *lua.LState) int {
		L.Push(wrapUiacc(L, u.Password(L.CheckBool(1))))
		return 1
	}))
	obj.RawSetString("click", L.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LBool(u.Click(L.CheckString(1))))
		return 1
	}))
	obj.RawSetString("waitFor", L.NewFunction(func(L *lua.LState) int {
		L.Push(wrapUiObject(L, u.WaitFor(L.CheckInt(1))))
		return 1
	}))
	obj.RawSetString("findOnce", L.NewFunction(func(L *lua.LState) int {
		L.Push(wrapUiObject(L, u.FindOnce()))
		return 1
	}))
	obj.RawSetString("find", L.NewFunction(func(L *lua.LState) int {
		result := L.NewTable()
		for i, item := range u.Find() {
			result.RawSetInt(i+1, wrapUiObject(L, item))
		}
		L.Push(result)
		return 1
	}))
	obj.RawSetString("release", L.NewFunction(func(L *lua.LState) int {
		u.Release()
		return 0
	}))
	return obj
}

func wrapUiObject(L *lua.LState, node *uiacc.UiObject) lua.LValue {
	if node == nil {
		return lua.LNil
	}
	obj := L.NewTable()
	obj.RawSetString("click", L.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LBool(node.Click()))
		return 1
	}))
	obj.RawSetString("clickCenter", L.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LBool(node.ClickCenter()))
		return 1
	}))
	obj.RawSetString("clickLongClick", L.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LBool(node.ClickLongClick()))
		return 1
	}))
	obj.RawSetString("copy", L.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LBool(node.Copy()))
		return 1
	}))
	obj.RawSetString("cut", L.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LBool(node.Cut()))
		return 1
	}))
	obj.RawSetString("paste", L.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LBool(node.Paste()))
		return 1
	}))
	obj.RawSetString("scrollForward", L.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LBool(node.ScrollForward()))
		return 1
	}))
	obj.RawSetString("scrollBackward", L.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LBool(node.ScrollBackward()))
		return 1
	}))
	obj.RawSetString("collapse", L.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LBool(node.Collapse()))
		return 1
	}))
	obj.RawSetString("expand", L.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LBool(node.Expand()))
		return 1
	}))
	obj.RawSetString("show", L.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LBool(node.Show()))
		return 1
	}))
	obj.RawSetString("select", L.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LBool(node.Select()))
		return 1
	}))
	obj.RawSetString("clearSelect", L.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LBool(node.ClearSelect()))
		return 1
	}))
	obj.RawSetString("getClickable", L.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LBool(node.GetClickable()))
		return 1
	}))
	obj.RawSetString("getLongClickable", L.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LBool(node.GetLongClickable()))
		return 1
	}))
	obj.RawSetString("getCheckable", L.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LBool(node.GetCheckable()))
		return 1
	}))
	obj.RawSetString("getSelected", L.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LBool(node.GetSelected()))
		return 1
	}))
	obj.RawSetString("getEnabled", L.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LBool(node.GetEnabled()))
		return 1
	}))
	obj.RawSetString("getScrollable", L.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LBool(node.GetScrollable()))
		return 1
	}))
	obj.RawSetString("getEditable", L.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LBool(node.GetEditable()))
		return 1
	}))
	obj.RawSetString("getMultiLine", L.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LBool(node.GetMultiLine()))
		return 1
	}))
	obj.RawSetString("getChecked", L.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LBool(node.GetChecked()))
		return 1
	}))
	obj.RawSetString("getFocused", L.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LBool(node.GetFocused()))
		return 1
	}))
	obj.RawSetString("getFocusable", L.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LBool(node.GetFocusable()))
		return 1
	}))
	obj.RawSetString("getDismissable", L.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LBool(node.GetDismissable()))
		return 1
	}))
	obj.RawSetString("getContextClickable", L.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LBool(node.GetContextClickable()))
		return 1
	}))
	obj.RawSetString("getVisible", L.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LBool(node.GetVisible()))
		return 1
	}))
	obj.RawSetString("getPassword", L.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LBool(node.GetPassword()))
		return 1
	}))
	obj.RawSetString("getAccessibilityFocused", L.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LBool(node.GetAccessibilityFocused()))
		return 1
	}))
	obj.RawSetString("getChildCount", L.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LNumber(node.GetChildCount()))
		return 1
	}))
	obj.RawSetString("getDrawingOrder", L.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LNumber(node.GetDrawingOrder()))
		return 1
	}))
	obj.RawSetString("getIndex", L.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LNumber(node.GetIndex()))
		return 1
	}))
	obj.RawSetString("getBounds", L.NewFunction(func(L *lua.LState) int {
		L.Push(rectToLua(L, node.GetBounds()))
		return 1
	}))
	obj.RawSetString("getBoundsInParent", L.NewFunction(func(L *lua.LState) int {
		L.Push(rectToLua(L, node.GetBoundsInParent()))
		return 1
	}))
	obj.RawSetString("getId", L.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LString(node.GetId()))
		return 1
	}))
	obj.RawSetString("getText", L.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LString(node.GetText()))
		return 1
	}))
	obj.RawSetString("getDesc", L.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LString(node.GetDesc()))
		return 1
	}))
	obj.RawSetString("getPackageName", L.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LString(node.GetPackageName()))
		return 1
	}))
	obj.RawSetString("getClassName", L.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LString(node.GetClassName()))
		return 1
	}))
	obj.RawSetString("toString", L.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LString(node.ToString()))
		return 1
	}))
	obj.RawSetString("setSelection", L.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LBool(node.SetSelection(L.CheckInt(1), L.CheckInt(2))))
		return 1
	}))
	obj.RawSetString("setVisibleToUser", L.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LBool(node.SetVisibleToUser(L.CheckBool(1))))
		return 1
	}))
	obj.RawSetString("setText", L.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LBool(node.SetText(L.CheckString(1))))
		return 1
	}))
	obj.RawSetString("getParent", L.NewFunction(func(L *lua.LState) int {
		L.Push(wrapUiObject(L, node.GetParent()))
		return 1
	}))
	obj.RawSetString("getChild", L.NewFunction(func(L *lua.LState) int {
		L.Push(wrapUiObject(L, node.GetChild(L.CheckInt(1))))
		return 1
	}))
	obj.RawSetString("getChildren", L.NewFunction(func(L *lua.LState) int {
		result := L.NewTable()
		for i, item := range node.GetChildren() {
			result.RawSetInt(i+1, wrapUiObject(L, item))
		}
		L.Push(result)
		return 1
	}))
	return obj
}

func rectToLua(L *lua.LState, rect uiacc.Rect) lua.LValue {
	obj := L.NewTable()
	obj.RawSetString("left", lua.LNumber(rect.Left))
	obj.RawSetString("right", lua.LNumber(rect.Right))
	obj.RawSetString("top", lua.LNumber(rect.Top))
	obj.RawSetString("bottom", lua.LNumber(rect.Bottom))
	obj.RawSetString("centerX", lua.LNumber(rect.CenterX))
	obj.RawSetString("centerY", lua.LNumber(rect.CenterY))
	obj.RawSetString("width", lua.LNumber(rect.Width))
	obj.RawSetString("height", lua.LNumber(rect.Height))
	return obj
}
