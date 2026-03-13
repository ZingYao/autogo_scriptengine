package uiacc

import (
	"app/lua_engine/model"

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
		displayId := L.CheckInt(1)
		u := uiacc.New(displayId)
		ud := L.NewUserData()
		ud.Value = u
		L.Push(ud)
		return 1
	}))

	uiaccObj.RawSetString("text", state.NewFunction(func(L *lua.LState) int {
		u := L.CheckUserData(1).Value.(*uiacc.Uiacc)
		value := L.CheckString(2)
		result := u.Text(value)
		ud := L.NewUserData()
		ud.Value = result
		L.Push(ud)
		return 1
	}))

	uiaccObj.RawSetString("textContains", state.NewFunction(func(L *lua.LState) int {
		u := L.CheckUserData(1).Value.(*uiacc.Uiacc)
		value := L.CheckString(2)
		result := u.TextContains(value)
		ud := L.NewUserData()
		ud.Value = result
		L.Push(ud)
		return 1
	}))

	uiaccObj.RawSetString("textStartsWith", state.NewFunction(func(L *lua.LState) int {
		u := L.CheckUserData(1).Value.(*uiacc.Uiacc)
		value := L.CheckString(2)
		result := u.TextStartsWith(value)
		ud := L.NewUserData()
		ud.Value = result
		L.Push(ud)
		return 1
	}))

	uiaccObj.RawSetString("textEndsWith", state.NewFunction(func(L *lua.LState) int {
		u := L.CheckUserData(1).Value.(*uiacc.Uiacc)
		value := L.CheckString(2)
		result := u.TextEndsWith(value)
		ud := L.NewUserData()
		ud.Value = result
		L.Push(ud)
		return 1
	}))

	uiaccObj.RawSetString("textMatches", state.NewFunction(func(L *lua.LState) int {
		u := L.CheckUserData(1).Value.(*uiacc.Uiacc)
		value := L.CheckString(2)
		result := u.TextMatches(value)
		ud := L.NewUserData()
		ud.Value = result
		L.Push(ud)
		return 1
	}))

	uiaccObj.RawSetString("desc", state.NewFunction(func(L *lua.LState) int {
		u := L.CheckUserData(1).Value.(*uiacc.Uiacc)
		value := L.CheckString(2)
		result := u.Desc(value)
		ud := L.NewUserData()
		ud.Value = result
		L.Push(ud)
		return 1
	}))

	uiaccObj.RawSetString("descContains", state.NewFunction(func(L *lua.LState) int {
		u := L.CheckUserData(1).Value.(*uiacc.Uiacc)
		value := L.CheckString(2)
		result := u.DescContains(value)
		ud := L.NewUserData()
		ud.Value = result
		L.Push(ud)
		return 1
	}))

	uiaccObj.RawSetString("descStartsWith", state.NewFunction(func(L *lua.LState) int {
		u := L.CheckUserData(1).Value.(*uiacc.Uiacc)
		value := L.CheckString(2)
		result := u.DescStartsWith(value)
		ud := L.NewUserData()
		ud.Value = result
		L.Push(ud)
		return 1
	}))

	uiaccObj.RawSetString("descEndsWith", state.NewFunction(func(L *lua.LState) int {
		u := L.CheckUserData(1).Value.(*uiacc.Uiacc)
		value := L.CheckString(2)
		result := u.DescEndsWith(value)
		ud := L.NewUserData()
		ud.Value = result
		L.Push(ud)
		return 1
	}))

	uiaccObj.RawSetString("descMatches", state.NewFunction(func(L *lua.LState) int {
		u := L.CheckUserData(1).Value.(*uiacc.Uiacc)
		value := L.CheckString(2)
		result := u.DescMatches(value)
		ud := L.NewUserData()
		ud.Value = result
		L.Push(ud)
		return 1
	}))

	uiaccObj.RawSetString("id", state.NewFunction(func(L *lua.LState) int {
		u := L.CheckUserData(1).Value.(*uiacc.Uiacc)
		value := L.CheckString(2)
		result := u.Id(value)
		ud := L.NewUserData()
		ud.Value = result
		L.Push(ud)
		return 1
	}))

	uiaccObj.RawSetString("idContains", state.NewFunction(func(L *lua.LState) int {
		u := L.CheckUserData(1).Value.(*uiacc.Uiacc)
		value := L.CheckString(2)
		result := u.IdContains(value)
		ud := L.NewUserData()
		ud.Value = result
		L.Push(ud)
		return 1
	}))

	uiaccObj.RawSetString("idStartsWith", state.NewFunction(func(L *lua.LState) int {
		u := L.CheckUserData(1).Value.(*uiacc.Uiacc)
		value := L.CheckString(2)
		result := u.IdStartsWith(value)
		ud := L.NewUserData()
		ud.Value = result
		L.Push(ud)
		return 1
	}))

	uiaccObj.RawSetString("idEndsWith", state.NewFunction(func(L *lua.LState) int {
		u := L.CheckUserData(1).Value.(*uiacc.Uiacc)
		value := L.CheckString(2)
		result := u.IdEndsWith(value)
		ud := L.NewUserData()
		ud.Value = result
		L.Push(ud)
		return 1
	}))

	uiaccObj.RawSetString("idMatches", state.NewFunction(func(L *lua.LState) int {
		u := L.CheckUserData(1).Value.(*uiacc.Uiacc)
		value := L.CheckString(2)
		result := u.IdMatches(value)
		ud := L.NewUserData()
		ud.Value = result
		L.Push(ud)
		return 1
	}))

	uiaccObj.RawSetString("className", state.NewFunction(func(L *lua.LState) int {
		u := L.CheckUserData(1).Value.(*uiacc.Uiacc)
		value := L.CheckString(2)
		result := u.ClassName(value)
		ud := L.NewUserData()
		ud.Value = result
		L.Push(ud)
		return 1
	}))

	uiaccObj.RawSetString("classNameContains", state.NewFunction(func(L *lua.LState) int {
		u := L.CheckUserData(1).Value.(*uiacc.Uiacc)
		value := L.CheckString(2)
		result := u.ClassNameContains(value)
		ud := L.NewUserData()
		ud.Value = result
		L.Push(ud)
		return 1
	}))

	uiaccObj.RawSetString("classNameStartsWith", state.NewFunction(func(L *lua.LState) int {
		u := L.CheckUserData(1).Value.(*uiacc.Uiacc)
		value := L.CheckString(2)
		result := u.ClassNameStartsWith(value)
		ud := L.NewUserData()
		ud.Value = result
		L.Push(ud)
		return 1
	}))

	uiaccObj.RawSetString("classNameEndsWith", state.NewFunction(func(L *lua.LState) int {
		u := L.CheckUserData(1).Value.(*uiacc.Uiacc)
		value := L.CheckString(2)
		result := u.ClassNameEndsWith(value)
		ud := L.NewUserData()
		ud.Value = result
		L.Push(ud)
		return 1
	}))

	uiaccObj.RawSetString("classNameMatches", state.NewFunction(func(L *lua.LState) int {
		u := L.CheckUserData(1).Value.(*uiacc.Uiacc)
		value := L.CheckString(2)
		result := u.ClassNameMatches(value)
		ud := L.NewUserData()
		ud.Value = result
		L.Push(ud)
		return 1
	}))

	uiaccObj.RawSetString("packageName", state.NewFunction(func(L *lua.LState) int {
		u := L.CheckUserData(1).Value.(*uiacc.Uiacc)
		value := L.CheckString(2)
		result := u.PackageName(value)
		ud := L.NewUserData()
		ud.Value = result
		L.Push(ud)
		return 1
	}))

	uiaccObj.RawSetString("packageNameContains", state.NewFunction(func(L *lua.LState) int {
		u := L.CheckUserData(1).Value.(*uiacc.Uiacc)
		value := L.CheckString(2)
		result := u.PackageNameContains(value)
		ud := L.NewUserData()
		ud.Value = result
		L.Push(ud)
		return 1
	}))

	uiaccObj.RawSetString("packageNameStartsWith", state.NewFunction(func(L *lua.LState) int {
		u := L.CheckUserData(1).Value.(*uiacc.Uiacc)
		value := L.CheckString(2)
		result := u.PackageNameStartsWith(value)
		ud := L.NewUserData()
		ud.Value = result
		L.Push(ud)
		return 1
	}))

	uiaccObj.RawSetString("packageNameEndsWith", state.NewFunction(func(L *lua.LState) int {
		u := L.CheckUserData(1).Value.(*uiacc.Uiacc)
		value := L.CheckString(2)
		result := u.PackageNameEndsWith(value)
		ud := L.NewUserData()
		ud.Value = result
		L.Push(ud)
		return 1
	}))

	uiaccObj.RawSetString("packageNameMatches", state.NewFunction(func(L *lua.LState) int {
		u := L.CheckUserData(1).Value.(*uiacc.Uiacc)
		value := L.CheckString(2)
		result := u.PackageNameMatches(value)
		ud := L.NewUserData()
		ud.Value = result
		L.Push(ud)
		return 1
	}))

	uiaccObj.RawSetString("bounds", state.NewFunction(func(L *lua.LState) int {
		u := L.CheckUserData(1).Value.(*uiacc.Uiacc)
		left := L.CheckInt(2)
		top := L.CheckInt(3)
		right := L.CheckInt(4)
		bottom := L.CheckInt(5)
		result := u.Bounds(left, top, right, bottom)
		ud := L.NewUserData()
		ud.Value = result
		L.Push(ud)
		return 1
	}))

	uiaccObj.RawSetString("boundsInside", state.NewFunction(func(L *lua.LState) int {
		u := L.CheckUserData(1).Value.(*uiacc.Uiacc)
		left := L.CheckInt(2)
		top := L.CheckInt(3)
		right := L.CheckInt(4)
		bottom := L.CheckInt(5)
		result := u.BoundsInside(left, top, right, bottom)
		ud := L.NewUserData()
		ud.Value = result
		L.Push(ud)
		return 1
	}))

	uiaccObj.RawSetString("boundsContains", state.NewFunction(func(L *lua.LState) int {
		u := L.CheckUserData(1).Value.(*uiacc.Uiacc)
		left := L.CheckInt(2)
		top := L.CheckInt(3)
		right := L.CheckInt(4)
		bottom := L.CheckInt(5)
		result := u.BoundsContains(left, top, right, bottom)
		ud := L.NewUserData()
		ud.Value = result
		L.Push(ud)
		return 1
	}))

	uiaccObj.RawSetString("drawingOrder", state.NewFunction(func(L *lua.LState) int {
		u := L.CheckUserData(1).Value.(*uiacc.Uiacc)
		value := L.CheckInt(2)
		result := u.DrawingOrder(value)
		ud := L.NewUserData()
		ud.Value = result
		L.Push(ud)
		return 1
	}))

	uiaccObj.RawSetString("clickable", state.NewFunction(func(L *lua.LState) int {
		u := L.CheckUserData(1).Value.(*uiacc.Uiacc)
		value := L.CheckBool(2)
		result := u.Clickable(value)
		ud := L.NewUserData()
		ud.Value = result
		L.Push(ud)
		return 1
	}))

	uiaccObj.RawSetString("longClickable", state.NewFunction(func(L *lua.LState) int {
		u := L.CheckUserData(1).Value.(*uiacc.Uiacc)
		value := L.CheckBool(2)
		result := u.LongClickable(value)
		ud := L.NewUserData()
		ud.Value = result
		L.Push(ud)
		return 1
	}))

	uiaccObj.RawSetString("checkable", state.NewFunction(func(L *lua.LState) int {
		u := L.CheckUserData(1).Value.(*uiacc.Uiacc)
		value := L.CheckBool(2)
		result := u.Checkable(value)
		ud := L.NewUserData()
		ud.Value = result
		L.Push(ud)
		return 1
	}))

	uiaccObj.RawSetString("selected", state.NewFunction(func(L *lua.LState) int {
		u := L.CheckUserData(1).Value.(*uiacc.Uiacc)
		value := L.CheckBool(2)
		result := u.Selected(value)
		ud := L.NewUserData()
		ud.Value = result
		L.Push(ud)
		return 1
	}))

	uiaccObj.RawSetString("enabled", state.NewFunction(func(L *lua.LState) int {
		u := L.CheckUserData(1).Value.(*uiacc.Uiacc)
		value := L.CheckBool(2)
		result := u.Enabled(value)
		ud := L.NewUserData()
		ud.Value = result
		L.Push(ud)
		return 1
	}))

	uiaccObj.RawSetString("scrollable", state.NewFunction(func(L *lua.LState) int {
		u := L.CheckUserData(1).Value.(*uiacc.Uiacc)
		value := L.CheckBool(2)
		result := u.Scrollable(value)
		ud := L.NewUserData()
		ud.Value = result
		L.Push(ud)
		return 1
	}))

	uiaccObj.RawSetString("editable", state.NewFunction(func(L *lua.LState) int {
		u := L.CheckUserData(1).Value.(*uiacc.Uiacc)
		value := L.CheckBool(2)
		result := u.Editable(value)
		ud := L.NewUserData()
		ud.Value = result
		L.Push(ud)
		return 1
	}))

	uiaccObj.RawSetString("multiLine", state.NewFunction(func(L *lua.LState) int {
		u := L.CheckUserData(1).Value.(*uiacc.Uiacc)
		value := L.CheckBool(2)
		result := u.MultiLine(value)
		ud := L.NewUserData()
		ud.Value = result
		L.Push(ud)
		return 1
	}))

	uiaccObj.RawSetString("checked", state.NewFunction(func(L *lua.LState) int {
		u := L.CheckUserData(1).Value.(*uiacc.Uiacc)
		value := L.CheckBool(2)
		result := u.Checked(value)
		ud := L.NewUserData()
		ud.Value = result
		L.Push(ud)
		return 1
	}))

	uiaccObj.RawSetString("focusable", state.NewFunction(func(L *lua.LState) int {
		u := L.CheckUserData(1).Value.(*uiacc.Uiacc)
		value := L.CheckBool(2)
		result := u.Focusable(value)
		ud := L.NewUserData()
		ud.Value = result
		L.Push(ud)
		return 1
	}))

	uiaccObj.RawSetString("dismissable", state.NewFunction(func(L *lua.LState) int {
		u := L.CheckUserData(1).Value.(*uiacc.Uiacc)
		value := L.CheckBool(2)
		result := u.Dismissable(value)
		ud := L.NewUserData()
		ud.Value = result
		L.Push(ud)
		return 1
	}))

	uiaccObj.RawSetString("focused", state.NewFunction(func(L *lua.LState) int {
		u := L.CheckUserData(1).Value.(*uiacc.Uiacc)
		value := L.CheckBool(2)
		result := u.Focused(value)
		ud := L.NewUserData()
		ud.Value = result
		L.Push(ud)
		return 1
	}))

	uiaccObj.RawSetString("contextClickable", state.NewFunction(func(L *lua.LState) int {
		u := L.CheckUserData(1).Value.(*uiacc.Uiacc)
		value := L.CheckBool(2)
		result := u.ContextClickable(value)
		ud := L.NewUserData()
		ud.Value = result
		L.Push(ud)
		return 1
	}))

	uiaccObj.RawSetString("index", state.NewFunction(func(L *lua.LState) int {
		u := L.CheckUserData(1).Value.(*uiacc.Uiacc)
		value := L.CheckInt(2)
		result := u.Index(value)
		ud := L.NewUserData()
		ud.Value = result
		L.Push(ud)
		return 1
	}))

	uiaccObj.RawSetString("visible", state.NewFunction(func(L *lua.LState) int {
		u := L.CheckUserData(1).Value.(*uiacc.Uiacc)
		value := L.CheckBool(2)
		result := u.Visible(value)
		ud := L.NewUserData()
		ud.Value = result
		L.Push(ud)
		return 1
	}))

	uiaccObj.RawSetString("password", state.NewFunction(func(L *lua.LState) int {
		u := L.CheckUserData(1).Value.(*uiacc.Uiacc)
		value := L.CheckBool(2)
		result := u.Password(value)
		ud := L.NewUserData()
		ud.Value = result
		L.Push(ud)
		return 1
	}))

	uiaccObj.RawSetString("click", state.NewFunction(func(L *lua.LState) int {
		u := L.CheckUserData(1).Value.(*uiacc.Uiacc)
		text := L.CheckString(2)
		result := u.Click(text)
		L.Push(lua.LBool(result))
		return 1
	}))

	uiaccObj.RawSetString("waitFor", state.NewFunction(func(L *lua.LState) int {
		u := L.CheckUserData(1).Value.(*uiacc.Uiacc)
		timeout := L.CheckInt(2)
		result := u.WaitFor(timeout)
		ud := L.NewUserData()
		ud.Value = result
		L.Push(ud)
		return 1
	}))

	uiaccObj.RawSetString("findOnce", state.NewFunction(func(L *lua.LState) int {
		u := L.CheckUserData(1).Value.(*uiacc.Uiacc)
		result := u.FindOnce()
		ud := L.NewUserData()
		ud.Value = result
		L.Push(ud)
		return 1
	}))

	uiaccObj.RawSetString("find", state.NewFunction(func(L *lua.LState) int {
		u := L.CheckUserData(1).Value.(*uiacc.Uiacc)
		result := u.Find()
		resultTable := L.NewTable()
		for i, item := range result {
			ud := L.NewUserData()
			ud.Value = item
			resultTable.RawSetInt(i+1, ud)
		}
		L.Push(resultTable)
		return 1
	}))

	uiaccObj.RawSetString("release", state.NewFunction(func(L *lua.LState) int {
		u := L.CheckUserData(1).Value.(*uiacc.Uiacc)
		u.Release()
		return 0
	}))

	uiaccObj.RawSetString("objClick", state.NewFunction(func(L *lua.LState) int {
		obj := L.CheckUserData(1).Value.(*uiacc.UiObject)
		result := obj.Click()
		L.Push(lua.LBool(result))
		return 1
	}))

	uiaccObj.RawSetString("clickCenter", state.NewFunction(func(L *lua.LState) int {
		obj := L.CheckUserData(1).Value.(*uiacc.UiObject)
		result := obj.ClickCenter()
		L.Push(lua.LBool(result))
		return 1
	}))

	uiaccObj.RawSetString("clickLongClick", state.NewFunction(func(L *lua.LState) int {
		obj := L.CheckUserData(1).Value.(*uiacc.UiObject)
		result := obj.ClickLongClick()
		L.Push(lua.LBool(result))
		return 1
	}))

	uiaccObj.RawSetString("copy", state.NewFunction(func(L *lua.LState) int {
		obj := L.CheckUserData(1).Value.(*uiacc.UiObject)
		result := obj.Copy()
		L.Push(lua.LBool(result))
		return 1
	}))

	uiaccObj.RawSetString("cut", state.NewFunction(func(L *lua.LState) int {
		obj := L.CheckUserData(1).Value.(*uiacc.UiObject)
		result := obj.Cut()
		L.Push(lua.LBool(result))
		return 1
	}))

	uiaccObj.RawSetString("paste", state.NewFunction(func(L *lua.LState) int {
		obj := L.CheckUserData(1).Value.(*uiacc.UiObject)
		result := obj.Paste()
		L.Push(lua.LBool(result))
		return 1
	}))

	uiaccObj.RawSetString("scrollForward", state.NewFunction(func(L *lua.LState) int {
		obj := L.CheckUserData(1).Value.(*uiacc.UiObject)
		result := obj.ScrollForward()
		L.Push(lua.LBool(result))
		return 1
	}))

	uiaccObj.RawSetString("scrollBackward", state.NewFunction(func(L *lua.LState) int {
		obj := L.CheckUserData(1).Value.(*uiacc.UiObject)
		result := obj.ScrollBackward()
		L.Push(lua.LBool(result))
		return 1
	}))

	uiaccObj.RawSetString("collapse", state.NewFunction(func(L *lua.LState) int {
		obj := L.CheckUserData(1).Value.(*uiacc.UiObject)
		result := obj.Collapse()
		L.Push(lua.LBool(result))
		return 1
	}))

	uiaccObj.RawSetString("expand", state.NewFunction(func(L *lua.LState) int {
		obj := L.CheckUserData(1).Value.(*uiacc.UiObject)
		result := obj.Expand()
		L.Push(lua.LBool(result))
		return 1
	}))

	uiaccObj.RawSetString("show", state.NewFunction(func(L *lua.LState) int {
		obj := L.CheckUserData(1).Value.(*uiacc.UiObject)
		result := obj.Show()
		L.Push(lua.LBool(result))
		return 1
	}))

	uiaccObj.RawSetString("select", state.NewFunction(func(L *lua.LState) int {
		obj := L.CheckUserData(1).Value.(*uiacc.UiObject)
		result := obj.Select()
		L.Push(lua.LBool(result))
		return 1
	}))

	uiaccObj.RawSetString("clearSelect", state.NewFunction(func(L *lua.LState) int {
		obj := L.CheckUserData(1).Value.(*uiacc.UiObject)
		result := obj.ClearSelect()
		L.Push(lua.LBool(result))
		return 1
	}))

	uiaccObj.RawSetString("setSelection", state.NewFunction(func(L *lua.LState) int {
		obj := L.CheckUserData(1).Value.(*uiacc.UiObject)
		start := L.CheckInt(2)
		end := L.CheckInt(3)
		result := obj.SetSelection(start, end)
		L.Push(lua.LBool(result))
		return 1
	}))

	uiaccObj.RawSetString("setVisibleToUser", state.NewFunction(func(L *lua.LState) int {
		obj := L.CheckUserData(1).Value.(*uiacc.UiObject)
		isVisible := L.CheckBool(2)
		result := obj.SetVisibleToUser(isVisible)
		L.Push(lua.LBool(result))
		return 1
	}))

	uiaccObj.RawSetString("setText", state.NewFunction(func(L *lua.LState) int {
		obj := L.CheckUserData(1).Value.(*uiacc.UiObject)
		str := L.CheckString(2)
		result := obj.SetText(str)
		L.Push(lua.LBool(result))
		return 1
	}))

	uiaccObj.RawSetString("getClickable", state.NewFunction(func(L *lua.LState) int {
		obj := L.CheckUserData(1).Value.(*uiacc.UiObject)
		result := obj.GetClickable()
		L.Push(lua.LBool(result))
		return 1
	}))

	uiaccObj.RawSetString("getLongClickable", state.NewFunction(func(L *lua.LState) int {
		obj := L.CheckUserData(1).Value.(*uiacc.UiObject)
		result := obj.GetLongClickable()
		L.Push(lua.LBool(result))
		return 1
	}))

	uiaccObj.RawSetString("getCheckable", state.NewFunction(func(L *lua.LState) int {
		obj := L.CheckUserData(1).Value.(*uiacc.UiObject)
		result := obj.GetCheckable()
		L.Push(lua.LBool(result))
		return 1
	}))

	uiaccObj.RawSetString("getSelected", state.NewFunction(func(L *lua.LState) int {
		obj := L.CheckUserData(1).Value.(*uiacc.UiObject)
		result := obj.GetSelected()
		L.Push(lua.LBool(result))
		return 1
	}))

	uiaccObj.RawSetString("getEnabled", state.NewFunction(func(L *lua.LState) int {
		obj := L.CheckUserData(1).Value.(*uiacc.UiObject)
		result := obj.GetEnabled()
		L.Push(lua.LBool(result))
		return 1
	}))

	uiaccObj.RawSetString("getScrollable", state.NewFunction(func(L *lua.LState) int {
		obj := L.CheckUserData(1).Value.(*uiacc.UiObject)
		result := obj.GetScrollable()
		L.Push(lua.LBool(result))
		return 1
	}))

	uiaccObj.RawSetString("getEditable", state.NewFunction(func(L *lua.LState) int {
		obj := L.CheckUserData(1).Value.(*uiacc.UiObject)
		result := obj.GetEditable()
		L.Push(lua.LBool(result))
		return 1
	}))

	uiaccObj.RawSetString("getMultiLine", state.NewFunction(func(L *lua.LState) int {
		obj := L.CheckUserData(1).Value.(*uiacc.UiObject)
		result := obj.GetMultiLine()
		L.Push(lua.LBool(result))
		return 1
	}))

	uiaccObj.RawSetString("getChecked", state.NewFunction(func(L *lua.LState) int {
		obj := L.CheckUserData(1).Value.(*uiacc.UiObject)
		result := obj.GetChecked()
		L.Push(lua.LBool(result))
		return 1
	}))

	uiaccObj.RawSetString("getFocused", state.NewFunction(func(L *lua.LState) int {
		obj := L.CheckUserData(1).Value.(*uiacc.UiObject)
		result := obj.GetFocused()
		L.Push(lua.LBool(result))
		return 1
	}))

	uiaccObj.RawSetString("getFocusable", state.NewFunction(func(L *lua.LState) int {
		obj := L.CheckUserData(1).Value.(*uiacc.UiObject)
		result := obj.GetFocusable()
		L.Push(lua.LBool(result))
		return 1
	}))

	uiaccObj.RawSetString("getDismissable", state.NewFunction(func(L *lua.LState) int {
		obj := L.CheckUserData(1).Value.(*uiacc.UiObject)
		result := obj.GetDismissable()
		L.Push(lua.LBool(result))
		return 1
	}))

	uiaccObj.RawSetString("getContextClickable", state.NewFunction(func(L *lua.LState) int {
		obj := L.CheckUserData(1).Value.(*uiacc.UiObject)
		result := obj.GetContextClickable()
		L.Push(lua.LBool(result))
		return 1
	}))

	uiaccObj.RawSetString("getAccessibilityFocused", state.NewFunction(func(L *lua.LState) int {
		obj := L.CheckUserData(1).Value.(*uiacc.UiObject)
		result := obj.GetAccessibilityFocused()
		L.Push(lua.LBool(result))
		return 1
	}))

	uiaccObj.RawSetString("getChildCount", state.NewFunction(func(L *lua.LState) int {
		obj := L.CheckUserData(1).Value.(*uiacc.UiObject)
		result := obj.GetChildCount()
		L.Push(lua.LNumber(result))
		return 1
	}))

	uiaccObj.RawSetString("getDrawingOrder", state.NewFunction(func(L *lua.LState) int {
		obj := L.CheckUserData(1).Value.(*uiacc.UiObject)
		result := obj.GetDrawingOrder()
		L.Push(lua.LNumber(result))
		return 1
	}))

	uiaccObj.RawSetString("getIndex", state.NewFunction(func(L *lua.LState) int {
		obj := L.CheckUserData(1).Value.(*uiacc.UiObject)
		result := obj.GetIndex()
		L.Push(lua.LNumber(result))
		return 1
	}))

	uiaccObj.RawSetString("getBounds", state.NewFunction(func(L *lua.LState) int {
		obj := L.CheckUserData(1).Value.(*uiacc.UiObject)
		result := obj.GetBounds()
		resultTable := L.NewTable()
		resultTable.RawSetString("left", lua.LNumber(result.Left))
		resultTable.RawSetString("top", lua.LNumber(result.Top))
		resultTable.RawSetString("right", lua.LNumber(result.Right))
		resultTable.RawSetString("bottom", lua.LNumber(result.Bottom))
		L.Push(resultTable)
		return 1
	}))

	uiaccObj.RawSetString("getBoundsInParent", state.NewFunction(func(L *lua.LState) int {
		obj := L.CheckUserData(1).Value.(*uiacc.UiObject)
		result := obj.GetBoundsInParent()
		resultTable := L.NewTable()
		resultTable.RawSetString("left", lua.LNumber(result.Left))
		resultTable.RawSetString("top", lua.LNumber(result.Top))
		resultTable.RawSetString("right", lua.LNumber(result.Right))
		resultTable.RawSetString("bottom", lua.LNumber(result.Bottom))
		L.Push(resultTable)
		return 1
	}))

	uiaccObj.RawSetString("getId", state.NewFunction(func(L *lua.LState) int {
		obj := L.CheckUserData(1).Value.(*uiacc.UiObject)
		result := obj.GetId()
		L.Push(lua.LString(result))
		return 1
	}))

	uiaccObj.RawSetString("getText", state.NewFunction(func(L *lua.LState) int {
		obj := L.CheckUserData(1).Value.(*uiacc.UiObject)
		result := obj.GetText()
		L.Push(lua.LString(result))
		return 1
	}))

	uiaccObj.RawSetString("getDesc", state.NewFunction(func(L *lua.LState) int {
		obj := L.CheckUserData(1).Value.(*uiacc.UiObject)
		result := obj.GetDesc()
		L.Push(lua.LString(result))
		return 1
	}))

	uiaccObj.RawSetString("getPackageName", state.NewFunction(func(L *lua.LState) int {
		obj := L.CheckUserData(1).Value.(*uiacc.UiObject)
		result := obj.GetPackageName()
		L.Push(lua.LString(result))
		return 1
	}))

	uiaccObj.RawSetString("getClassName", state.NewFunction(func(L *lua.LState) int {
		obj := L.CheckUserData(1).Value.(*uiacc.UiObject)
		result := obj.GetClassName()
		L.Push(lua.LString(result))
		return 1
	}))

	uiaccObj.RawSetString("getParent", state.NewFunction(func(L *lua.LState) int {
		obj := L.CheckUserData(1).Value.(*uiacc.UiObject)
		result := obj.GetParent()
		ud := L.NewUserData()
		ud.Value = result
		L.Push(ud)
		return 1
	}))

	uiaccObj.RawSetString("getChild", state.NewFunction(func(L *lua.LState) int {
		obj := L.CheckUserData(1).Value.(*uiacc.UiObject)
		index := L.CheckInt(2)
		result := obj.GetChild(index)
		ud := L.NewUserData()
		ud.Value = result
		L.Push(ud)
		return 1
	}))

	uiaccObj.RawSetString("getChildren", state.NewFunction(func(L *lua.LState) int {
		obj := L.CheckUserData(1).Value.(*uiacc.UiObject)
		result := obj.GetChildren()
		resultTable := L.NewTable()
		for i, item := range result {
			ud := L.NewUserData()
			ud.Value = item
			resultTable.RawSetInt(i+1, ud)
		}
		L.Push(resultTable)
		return 1
	}))

	uiaccObj.RawSetString("getVisible", state.NewFunction(func(L *lua.LState) int {
		obj := L.CheckUserData(1).Value.(*uiacc.UiObject)
		result := obj.GetVisible()
		L.Push(lua.LBool(result))
		return 1
	}))

	uiaccObj.RawSetString("getPassword", state.NewFunction(func(L *lua.LState) int {
		obj := L.CheckUserData(1).Value.(*uiacc.UiObject)
		result := obj.GetPassword()
		L.Push(lua.LBool(result))
		return 1
	}))

	uiaccObj.RawSetString("toString", state.NewFunction(func(L *lua.LState) int {
		obj := L.CheckUserData(1).Value.(*uiacc.UiObject)
		result := obj.ToString()
		L.Push(lua.LString(result))
		return 1
	}))

	engine.RegisterMethod("uiacc.new", "创建新的Uiacc对象", uiacc.New, true)
	engine.RegisterMethod("uiacc.text", "根据text属性筛选", func(u *uiacc.Uiacc, value string) *uiacc.Uiacc {
		return u.Text(value)
	}, true)
	engine.RegisterMethod("uiacc.textContains", "根据text包含筛选", func(u *uiacc.Uiacc, value string) *uiacc.Uiacc {
		return u.TextContains(value)
	}, true)
	engine.RegisterMethod("uiacc.textStartsWith", "根据text前缀筛选", func(u *uiacc.Uiacc, value string) *uiacc.Uiacc {
		return u.TextStartsWith(value)
	}, true)
	engine.RegisterMethod("uiacc.textEndsWith", "根据text后缀筛选", func(u *uiacc.Uiacc, value string) *uiacc.Uiacc {
		return u.TextEndsWith(value)
	}, true)
	engine.RegisterMethod("uiacc.textMatches", "根据text正则筛选", func(u *uiacc.Uiacc, value string) *uiacc.Uiacc {
		return u.TextMatches(value)
	}, true)
	engine.RegisterMethod("uiacc.desc", "根据desc属性筛选", func(u *uiacc.Uiacc, value string) *uiacc.Uiacc {
		return u.Desc(value)
	}, true)
	engine.RegisterMethod("uiacc.descContains", "根据desc包含筛选", func(u *uiacc.Uiacc, value string) *uiacc.Uiacc {
		return u.DescContains(value)
	}, true)
	engine.RegisterMethod("uiacc.descStartsWith", "根据desc前缀筛选", func(u *uiacc.Uiacc, value string) *uiacc.Uiacc {
		return u.DescStartsWith(value)
	}, true)
	engine.RegisterMethod("uiacc.descEndsWith", "根据desc后缀筛选", func(u *uiacc.Uiacc, value string) *uiacc.Uiacc {
		return u.DescEndsWith(value)
	}, true)
	engine.RegisterMethod("uiacc.descMatches", "根据desc正则筛选", func(u *uiacc.Uiacc, value string) *uiacc.Uiacc {
		return u.DescMatches(value)
	}, true)
	engine.RegisterMethod("uiacc.id", "根据id属性筛选", func(u *uiacc.Uiacc, value string) *uiacc.Uiacc {
		return u.Id(value)
	}, true)
	engine.RegisterMethod("uiacc.idContains", "根据id包含筛选", func(u *uiacc.Uiacc, value string) *uiacc.Uiacc {
		return u.IdContains(value)
	}, true)
	engine.RegisterMethod("uiacc.idStartsWith", "根据id前缀筛选", func(u *uiacc.Uiacc, value string) *uiacc.Uiacc {
		return u.IdStartsWith(value)
	}, true)
	engine.RegisterMethod("uiacc.idEndsWith", "根据id后缀筛选", func(u *uiacc.Uiacc, value string) *uiacc.Uiacc {
		return u.IdEndsWith(value)
	}, true)
	engine.RegisterMethod("uiacc.idMatches", "根据id正则筛选", func(u *uiacc.Uiacc, value string) *uiacc.Uiacc {
		return u.IdMatches(value)
	}, true)
	engine.RegisterMethod("uiacc.className", "根据className属性筛选", func(u *uiacc.Uiacc, value string) *uiacc.Uiacc {
		return u.ClassName(value)
	}, true)
	engine.RegisterMethod("uiacc.classNameContains", "根据className包含筛选", func(u *uiacc.Uiacc, value string) *uiacc.Uiacc {
		return u.ClassNameContains(value)
	}, true)
	engine.RegisterMethod("uiacc.classNameStartsWith", "根据className前缀筛选", func(u *uiacc.Uiacc, value string) *uiacc.Uiacc {
		return u.ClassNameStartsWith(value)
	}, true)
	engine.RegisterMethod("uiacc.classNameEndsWith", "根据className后缀筛选", func(u *uiacc.Uiacc, value string) *uiacc.Uiacc {
		return u.ClassNameEndsWith(value)
	}, true)
	engine.RegisterMethod("uiacc.classNameMatches", "根据className正则筛选", func(u *uiacc.Uiacc, value string) *uiacc.Uiacc {
		return u.ClassNameMatches(value)
	}, true)
	engine.RegisterMethod("uiacc.packageName", "根据packageName属性筛选", func(u *uiacc.Uiacc, value string) *uiacc.Uiacc {
		return u.PackageName(value)
	}, true)
	engine.RegisterMethod("uiacc.packageNameContains", "根据packageName包含筛选", func(u *uiacc.Uiacc, value string) *uiacc.Uiacc {
		return u.PackageNameContains(value)
	}, true)
	engine.RegisterMethod("uiacc.packageNameStartsWith", "根据packageName前缀筛选", func(u *uiacc.Uiacc, value string) *uiacc.Uiacc {
		return u.PackageNameStartsWith(value)
	}, true)
	engine.RegisterMethod("uiacc.packageNameEndsWith", "根据packageName后缀筛选", func(u *uiacc.Uiacc, value string) *uiacc.Uiacc {
		return u.PackageNameEndsWith(value)
	}, true)
	engine.RegisterMethod("uiacc.packageNameMatches", "根据packageName正则筛选", func(u *uiacc.Uiacc, value string) *uiacc.Uiacc {
		return u.PackageNameMatches(value)
	}, true)
	engine.RegisterMethod("uiacc.bounds", "根据bounds属性筛选", func(u *uiacc.Uiacc, left, top, right, bottom int) *uiacc.Uiacc {
		return u.Bounds(left, top, right, bottom)
	}, true)
	engine.RegisterMethod("uiacc.boundsInside", "根据boundsInside属性筛选", func(u *uiacc.Uiacc, left, top, right, bottom int) *uiacc.Uiacc {
		return u.BoundsInside(left, top, right, bottom)
	}, true)
	engine.RegisterMethod("uiacc.boundsContains", "根据boundsContains属性筛选", func(u *uiacc.Uiacc, left, top, right, bottom int) *uiacc.Uiacc {
		return u.BoundsContains(left, top, right, bottom)
	}, true)
	engine.RegisterMethod("uiacc.drawingOrder", "根据drawingOrder属性筛选", func(u *uiacc.Uiacc, value int) *uiacc.Uiacc {
		return u.DrawingOrder(value)
	}, true)
	engine.RegisterMethod("uiacc.clickable", "根据clickable属性筛选", func(u *uiacc.Uiacc, value bool) *uiacc.Uiacc {
		return u.Clickable(value)
	}, true)
	engine.RegisterMethod("uiacc.longClickable", "根据longClickable属性筛选", func(u *uiacc.Uiacc, value bool) *uiacc.Uiacc {
		return u.LongClickable(value)
	}, true)
	engine.RegisterMethod("uiacc.checkable", "根据checkable属性筛选", func(u *uiacc.Uiacc, value bool) *uiacc.Uiacc {
		return u.Checkable(value)
	}, true)
	engine.RegisterMethod("uiacc.selected", "根据selected属性筛选", func(u *uiacc.Uiacc, value bool) *uiacc.Uiacc {
		return u.Selected(value)
	}, true)
	engine.RegisterMethod("uiacc.enabled", "根据enabled属性筛选", func(u *uiacc.Uiacc, value bool) *uiacc.Uiacc {
		return u.Enabled(value)
	}, true)
	engine.RegisterMethod("uiacc.scrollable", "根据scrollable属性筛选", func(u *uiacc.Uiacc, value bool) *uiacc.Uiacc {
		return u.Scrollable(value)
	}, true)
	engine.RegisterMethod("uiacc.editable", "根据editable属性筛选", func(u *uiacc.Uiacc, value bool) *uiacc.Uiacc {
		return u.Editable(value)
	}, true)
	engine.RegisterMethod("uiacc.multiLine", "根据multiLine属性筛选", func(u *uiacc.Uiacc, value bool) *uiacc.Uiacc {
		return u.MultiLine(value)
	}, true)
	engine.RegisterMethod("uiacc.checked", "根据checked属性筛选", func(u *uiacc.Uiacc, value bool) *uiacc.Uiacc {
		return u.Checked(value)
	}, true)
	engine.RegisterMethod("uiacc.focusable", "根据focusable属性筛选", func(u *uiacc.Uiacc, value bool) *uiacc.Uiacc {
		return u.Focusable(value)
	}, true)
	engine.RegisterMethod("uiacc.dismissable", "根据dismissable属性筛选", func(u *uiacc.Uiacc, value bool) *uiacc.Uiacc {
		return u.Dismissable(value)
	}, true)
	engine.RegisterMethod("uiacc.focused", "根据focused属性筛选", func(u *uiacc.Uiacc, value bool) *uiacc.Uiacc {
		return u.Focused(value)
	}, true)
	engine.RegisterMethod("uiacc.contextClickable", "根据contextClickable属性筛选", func(u *uiacc.Uiacc, value bool) *uiacc.Uiacc {
		return u.ContextClickable(value)
	}, true)
	engine.RegisterMethod("uiacc.index", "根据index属性筛选", func(u *uiacc.Uiacc, value int) *uiacc.Uiacc {
		return u.Index(value)
	}, true)
	engine.RegisterMethod("uiacc.visible", "根据visible属性筛选", func(u *uiacc.Uiacc, value bool) *uiacc.Uiacc {
		return u.Visible(value)
	}, true)
	engine.RegisterMethod("uiacc.password", "根据password属性筛选", func(u *uiacc.Uiacc, value bool) *uiacc.Uiacc {
		return u.Password(value)
	}, true)
	engine.RegisterMethod("uiacc.click", "点击匹配的控件", func(u *uiacc.Uiacc, text string) bool {
		return u.Click(text)
	}, true)
	engine.RegisterMethod("uiacc.waitFor", "等待匹配的控件出现", func(u *uiacc.Uiacc, timeout int) *uiacc.UiObject {
		return u.WaitFor(timeout)
	}, true)
	engine.RegisterMethod("uiacc.findOnce", "查找第一个匹配的控件", func(u *uiacc.Uiacc) *uiacc.UiObject {
		return u.FindOnce()
	}, true)
	engine.RegisterMethod("uiacc.find", "查找所有匹配的控件", func(u *uiacc.Uiacc) []*uiacc.UiObject {
		return u.Find()
	}, true)
	engine.RegisterMethod("uiacc.release", "释放Uiacc对象", func(u *uiacc.Uiacc) {
		u.Release()
	}, true)

	return nil
}
