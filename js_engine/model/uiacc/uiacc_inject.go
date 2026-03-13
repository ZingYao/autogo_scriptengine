package uiacc

import (
	"github.com/ZingYao/autogo_scriptengine/js_engine/model"

	"github.com/Dasongzi1366/AutoGo/uiacc"
	"github.com/dop251/goja"
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
	vm := engine.GetVM()

	uiaccObj := vm.NewObject()
	vm.Set("uiacc", uiaccObj)

	// 创建新的 Uiacc 对象
	uiaccObj.Set("new", func(call goja.FunctionCall) goja.Value {
		displayId := 0
		if len(call.Arguments) >= 1 {
			displayId = int(call.Argument(0).ToInteger())
		}
		u := uiacc.New(displayId)
		return vm.ToValue(u)
	})

	// Uiacc 选择器方法
	uiaccObj.Set("text", func(call goja.FunctionCall) goja.Value {
		u := call.Argument(0).Export().(*uiacc.Uiacc)
		value := call.Argument(1).String()
		result := u.Text(value)
		return vm.ToValue(result)
	})

	uiaccObj.Set("textContains", func(call goja.FunctionCall) goja.Value {
		u := call.Argument(0).Export().(*uiacc.Uiacc)
		value := call.Argument(1).String()
		result := u.TextContains(value)
		return vm.ToValue(result)
	})

	uiaccObj.Set("textStartsWith", func(call goja.FunctionCall) goja.Value {
		u := call.Argument(0).Export().(*uiacc.Uiacc)
		value := call.Argument(1).String()
		result := u.TextStartsWith(value)
		return vm.ToValue(result)
	})

	uiaccObj.Set("textEndsWith", func(call goja.FunctionCall) goja.Value {
		u := call.Argument(0).Export().(*uiacc.Uiacc)
		value := call.Argument(1).String()
		result := u.TextEndsWith(value)
		return vm.ToValue(result)
	})

	uiaccObj.Set("textMatches", func(call goja.FunctionCall) goja.Value {
		u := call.Argument(0).Export().(*uiacc.Uiacc)
		value := call.Argument(1).String()
		result := u.TextMatches(value)
		return vm.ToValue(result)
	})

	uiaccObj.Set("desc", func(call goja.FunctionCall) goja.Value {
		u := call.Argument(0).Export().(*uiacc.Uiacc)
		value := call.Argument(1).String()
		result := u.Desc(value)
		return vm.ToValue(result)
	})

	uiaccObj.Set("descContains", func(call goja.FunctionCall) goja.Value {
		u := call.Argument(0).Export().(*uiacc.Uiacc)
		value := call.Argument(1).String()
		result := u.DescContains(value)
		return vm.ToValue(result)
	})

	uiaccObj.Set("descStartsWith", func(call goja.FunctionCall) goja.Value {
		u := call.Argument(0).Export().(*uiacc.Uiacc)
		value := call.Argument(1).String()
		result := u.DescStartsWith(value)
		return vm.ToValue(result)
	})

	uiaccObj.Set("descEndsWith", func(call goja.FunctionCall) goja.Value {
		u := call.Argument(0).Export().(*uiacc.Uiacc)
		value := call.Argument(1).String()
		result := u.DescEndsWith(value)
		return vm.ToValue(result)
	})

	uiaccObj.Set("descMatches", func(call goja.FunctionCall) goja.Value {
		u := call.Argument(0).Export().(*uiacc.Uiacc)
		value := call.Argument(1).String()
		result := u.DescMatches(value)
		return vm.ToValue(result)
	})

	uiaccObj.Set("id", func(call goja.FunctionCall) goja.Value {
		u := call.Argument(0).Export().(*uiacc.Uiacc)
		value := call.Argument(1).String()
		result := u.Id(value)
		return vm.ToValue(result)
	})

	uiaccObj.Set("idContains", func(call goja.FunctionCall) goja.Value {
		u := call.Argument(0).Export().(*uiacc.Uiacc)
		value := call.Argument(1).String()
		result := u.IdContains(value)
		return vm.ToValue(result)
	})

	uiaccObj.Set("idStartsWith", func(call goja.FunctionCall) goja.Value {
		u := call.Argument(0).Export().(*uiacc.Uiacc)
		value := call.Argument(1).String()
		result := u.IdStartsWith(value)
		return vm.ToValue(result)
	})

	uiaccObj.Set("idEndsWith", func(call goja.FunctionCall) goja.Value {
		u := call.Argument(0).Export().(*uiacc.Uiacc)
		value := call.Argument(1).String()
		result := u.IdEndsWith(value)
		return vm.ToValue(result)
	})

	uiaccObj.Set("idMatches", func(call goja.FunctionCall) goja.Value {
		u := call.Argument(0).Export().(*uiacc.Uiacc)
		value := call.Argument(1).String()
		result := u.IdMatches(value)
		return vm.ToValue(result)
	})

	uiaccObj.Set("className", func(call goja.FunctionCall) goja.Value {
		u := call.Argument(0).Export().(*uiacc.Uiacc)
		value := call.Argument(1).String()
		result := u.ClassName(value)
		return vm.ToValue(result)
	})

	uiaccObj.Set("classNameContains", func(call goja.FunctionCall) goja.Value {
		u := call.Argument(0).Export().(*uiacc.Uiacc)
		value := call.Argument(1).String()
		result := u.ClassNameContains(value)
		return vm.ToValue(result)
	})

	uiaccObj.Set("classNameStartsWith", func(call goja.FunctionCall) goja.Value {
		u := call.Argument(0).Export().(*uiacc.Uiacc)
		value := call.Argument(1).String()
		result := u.ClassNameStartsWith(value)
		return vm.ToValue(result)
	})

	uiaccObj.Set("classNameEndsWith", func(call goja.FunctionCall) goja.Value {
		u := call.Argument(0).Export().(*uiacc.Uiacc)
		value := call.Argument(1).String()
		result := u.ClassNameEndsWith(value)
		return vm.ToValue(result)
	})

	uiaccObj.Set("classNameMatches", func(call goja.FunctionCall) goja.Value {
		u := call.Argument(0).Export().(*uiacc.Uiacc)
		value := call.Argument(1).String()
		result := u.ClassNameMatches(value)
		return vm.ToValue(result)
	})

	uiaccObj.Set("packageName", func(call goja.FunctionCall) goja.Value {
		u := call.Argument(0).Export().(*uiacc.Uiacc)
		value := call.Argument(1).String()
		result := u.PackageName(value)
		return vm.ToValue(result)
	})

	uiaccObj.Set("packageNameContains", func(call goja.FunctionCall) goja.Value {
		u := call.Argument(0).Export().(*uiacc.Uiacc)
		value := call.Argument(1).String()
		result := u.PackageNameContains(value)
		return vm.ToValue(result)
	})

	uiaccObj.Set("packageNameStartsWith", func(call goja.FunctionCall) goja.Value {
		u := call.Argument(0).Export().(*uiacc.Uiacc)
		value := call.Argument(1).String()
		result := u.PackageNameStartsWith(value)
		return vm.ToValue(result)
	})

	uiaccObj.Set("packageNameEndsWith", func(call goja.FunctionCall) goja.Value {
		u := call.Argument(0).Export().(*uiacc.Uiacc)
		value := call.Argument(1).String()
		result := u.PackageNameEndsWith(value)
		return vm.ToValue(result)
	})

	uiaccObj.Set("packageNameMatches", func(call goja.FunctionCall) goja.Value {
		u := call.Argument(0).Export().(*uiacc.Uiacc)
		value := call.Argument(1).String()
		result := u.PackageNameMatches(value)
		return vm.ToValue(result)
	})

	uiaccObj.Set("bounds", func(call goja.FunctionCall) goja.Value {
		u := call.Argument(0).Export().(*uiacc.Uiacc)
		left := int(call.Argument(1).ToInteger())
		top := int(call.Argument(2).ToInteger())
		right := int(call.Argument(3).ToInteger())
		bottom := int(call.Argument(4).ToInteger())
		result := u.Bounds(left, top, right, bottom)
		return vm.ToValue(result)
	})

	uiaccObj.Set("boundsInside", func(call goja.FunctionCall) goja.Value {
		u := call.Argument(0).Export().(*uiacc.Uiacc)
		left := int(call.Argument(1).ToInteger())
		top := int(call.Argument(2).ToInteger())
		right := int(call.Argument(3).ToInteger())
		bottom := int(call.Argument(4).ToInteger())
		result := u.BoundsInside(left, top, right, bottom)
		return vm.ToValue(result)
	})

	uiaccObj.Set("boundsContains", func(call goja.FunctionCall) goja.Value {
		u := call.Argument(0).Export().(*uiacc.Uiacc)
		left := int(call.Argument(1).ToInteger())
		top := int(call.Argument(2).ToInteger())
		right := int(call.Argument(3).ToInteger())
		bottom := int(call.Argument(4).ToInteger())
		result := u.BoundsContains(left, top, right, bottom)
		return vm.ToValue(result)
	})

	uiaccObj.Set("drawingOrder", func(call goja.FunctionCall) goja.Value {
		u := call.Argument(0).Export().(*uiacc.Uiacc)
		value := int(call.Argument(1).ToInteger())
		result := u.DrawingOrder(value)
		return vm.ToValue(result)
	})

	uiaccObj.Set("clickable", func(call goja.FunctionCall) goja.Value {
		u := call.Argument(0).Export().(*uiacc.Uiacc)
		value := call.Argument(1).ToBoolean()
		result := u.Clickable(value)
		return vm.ToValue(result)
	})

	uiaccObj.Set("longClickable", func(call goja.FunctionCall) goja.Value {
		u := call.Argument(0).Export().(*uiacc.Uiacc)
		value := call.Argument(1).ToBoolean()
		result := u.LongClickable(value)
		return vm.ToValue(result)
	})

	uiaccObj.Set("checkable", func(call goja.FunctionCall) goja.Value {
		u := call.Argument(0).Export().(*uiacc.Uiacc)
		value := call.Argument(1).ToBoolean()
		result := u.Checkable(value)
		return vm.ToValue(result)
	})

	uiaccObj.Set("selected", func(call goja.FunctionCall) goja.Value {
		u := call.Argument(0).Export().(*uiacc.Uiacc)
		value := call.Argument(1).ToBoolean()
		result := u.Selected(value)
		return vm.ToValue(result)
	})

	uiaccObj.Set("enabled", func(call goja.FunctionCall) goja.Value {
		u := call.Argument(0).Export().(*uiacc.Uiacc)
		value := call.Argument(1).ToBoolean()
		result := u.Enabled(value)
		return vm.ToValue(result)
	})

	uiaccObj.Set("scrollable", func(call goja.FunctionCall) goja.Value {
		u := call.Argument(0).Export().(*uiacc.Uiacc)
		value := call.Argument(1).ToBoolean()
		result := u.Scrollable(value)
		return vm.ToValue(result)
	})

	uiaccObj.Set("editable", func(call goja.FunctionCall) goja.Value {
		u := call.Argument(0).Export().(*uiacc.Uiacc)
		value := call.Argument(1).ToBoolean()
		result := u.Editable(value)
		return vm.ToValue(result)
	})

	uiaccObj.Set("multiLine", func(call goja.FunctionCall) goja.Value {
		u := call.Argument(0).Export().(*uiacc.Uiacc)
		value := call.Argument(1).ToBoolean()
		result := u.MultiLine(value)
		return vm.ToValue(result)
	})

	uiaccObj.Set("checked", func(call goja.FunctionCall) goja.Value {
		u := call.Argument(0).Export().(*uiacc.Uiacc)
		value := call.Argument(1).ToBoolean()
		result := u.Checked(value)
		return vm.ToValue(result)
	})

	uiaccObj.Set("focusable", func(call goja.FunctionCall) goja.Value {
		u := call.Argument(0).Export().(*uiacc.Uiacc)
		value := call.Argument(1).ToBoolean()
		result := u.Focusable(value)
		return vm.ToValue(result)
	})

	uiaccObj.Set("dismissable", func(call goja.FunctionCall) goja.Value {
		u := call.Argument(0).Export().(*uiacc.Uiacc)
		value := call.Argument(1).ToBoolean()
		result := u.Dismissable(value)
		return vm.ToValue(result)
	})

	uiaccObj.Set("focused", func(call goja.FunctionCall) goja.Value {
		u := call.Argument(0).Export().(*uiacc.Uiacc)
		value := call.Argument(1).ToBoolean()
		result := u.Focused(value)
		return vm.ToValue(result)
	})

	uiaccObj.Set("contextClickable", func(call goja.FunctionCall) goja.Value {
		u := call.Argument(0).Export().(*uiacc.Uiacc)
		value := call.Argument(1).ToBoolean()
		result := u.ContextClickable(value)
		return vm.ToValue(result)
	})

	uiaccObj.Set("index", func(call goja.FunctionCall) goja.Value {
		u := call.Argument(0).Export().(*uiacc.Uiacc)
		value := int(call.Argument(1).ToInteger())
		result := u.Index(value)
		return vm.ToValue(result)
	})

	uiaccObj.Set("visible", func(call goja.FunctionCall) goja.Value {
		u := call.Argument(0).Export().(*uiacc.Uiacc)
		value := call.Argument(1).ToBoolean()
		result := u.Visible(value)
		return vm.ToValue(result)
	})

	uiaccObj.Set("password", func(call goja.FunctionCall) goja.Value {
		u := call.Argument(0).Export().(*uiacc.Uiacc)
		value := call.Argument(1).ToBoolean()
		result := u.Password(value)
		return vm.ToValue(result)
	})

	uiaccObj.Set("click", func(call goja.FunctionCall) goja.Value {
		u := call.Argument(0).Export().(*uiacc.Uiacc)
		text := call.Argument(1).String()
		result := u.Click(text)
		return vm.ToValue(result)
	})

	uiaccObj.Set("waitFor", func(call goja.FunctionCall) goja.Value {
		u := call.Argument(0).Export().(*uiacc.Uiacc)
		timeout := int(call.Argument(1).ToInteger())
		result := u.WaitFor(timeout)
		return vm.ToValue(result)
	})

	uiaccObj.Set("findOnce", func(call goja.FunctionCall) goja.Value {
		u := call.Argument(0).Export().(*uiacc.Uiacc)
		result := u.FindOnce()
		return vm.ToValue(result)
	})

	uiaccObj.Set("find", func(call goja.FunctionCall) goja.Value {
		u := call.Argument(0).Export().(*uiacc.Uiacc)
		result := u.Find()
		return vm.ToValue(result)
	})

	uiaccObj.Set("release", func(call goja.FunctionCall) goja.Value {
		u := call.Argument(0).Export().(*uiacc.Uiacc)
		u.Release()
		return goja.Undefined()
	})

	// UiObject 方法
	uiaccObj.Set("objClick", func(call goja.FunctionCall) goja.Value {
		obj := call.Argument(0).Export().(*uiacc.UiObject)
		result := obj.Click()
		return vm.ToValue(result)
	})

	uiaccObj.Set("clickCenter", func(call goja.FunctionCall) goja.Value {
		obj := call.Argument(0).Export().(*uiacc.UiObject)
		result := obj.ClickCenter()
		return vm.ToValue(result)
	})

	uiaccObj.Set("clickLongClick", func(call goja.FunctionCall) goja.Value {
		obj := call.Argument(0).Export().(*uiacc.UiObject)
		result := obj.ClickLongClick()
		return vm.ToValue(result)
	})

	uiaccObj.Set("copy", func(call goja.FunctionCall) goja.Value {
		obj := call.Argument(0).Export().(*uiacc.UiObject)
		result := obj.Copy()
		return vm.ToValue(result)
	})

	uiaccObj.Set("cut", func(call goja.FunctionCall) goja.Value {
		obj := call.Argument(0).Export().(*uiacc.UiObject)
		result := obj.Cut()
		return vm.ToValue(result)
	})

	uiaccObj.Set("paste", func(call goja.FunctionCall) goja.Value {
		obj := call.Argument(0).Export().(*uiacc.UiObject)
		result := obj.Paste()
		return vm.ToValue(result)
	})

	uiaccObj.Set("scrollForward", func(call goja.FunctionCall) goja.Value {
		obj := call.Argument(0).Export().(*uiacc.UiObject)
		result := obj.ScrollForward()
		return vm.ToValue(result)
	})

	uiaccObj.Set("scrollBackward", func(call goja.FunctionCall) goja.Value {
		obj := call.Argument(0).Export().(*uiacc.UiObject)
		result := obj.ScrollBackward()
		return vm.ToValue(result)
	})

	uiaccObj.Set("collapse", func(call goja.FunctionCall) goja.Value {
		obj := call.Argument(0).Export().(*uiacc.UiObject)
		result := obj.Collapse()
		return vm.ToValue(result)
	})

	uiaccObj.Set("expand", func(call goja.FunctionCall) goja.Value {
		obj := call.Argument(0).Export().(*uiacc.UiObject)
		result := obj.Expand()
		return vm.ToValue(result)
	})

	uiaccObj.Set("show", func(call goja.FunctionCall) goja.Value {
		obj := call.Argument(0).Export().(*uiacc.UiObject)
		result := obj.Show()
		return vm.ToValue(result)
	})

	uiaccObj.Set("select", func(call goja.FunctionCall) goja.Value {
		obj := call.Argument(0).Export().(*uiacc.UiObject)
		result := obj.Select()
		return vm.ToValue(result)
	})

	uiaccObj.Set("clearSelect", func(call goja.FunctionCall) goja.Value {
		obj := call.Argument(0).Export().(*uiacc.UiObject)
		result := obj.ClearSelect()
		return vm.ToValue(result)
	})

	uiaccObj.Set("setSelection", func(call goja.FunctionCall) goja.Value {
		obj := call.Argument(0).Export().(*uiacc.UiObject)
		start := int(call.Argument(1).ToInteger())
		end := int(call.Argument(2).ToInteger())
		result := obj.SetSelection(start, end)
		return vm.ToValue(result)
	})

	uiaccObj.Set("setVisibleToUser", func(call goja.FunctionCall) goja.Value {
		obj := call.Argument(0).Export().(*uiacc.UiObject)
		isVisible := call.Argument(1).ToBoolean()
		result := obj.SetVisibleToUser(isVisible)
		return vm.ToValue(result)
	})

	uiaccObj.Set("setText", func(call goja.FunctionCall) goja.Value {
		obj := call.Argument(0).Export().(*uiacc.UiObject)
		str := call.Argument(1).String()
		result := obj.SetText(str)
		return vm.ToValue(result)
	})

	uiaccObj.Set("getClickable", func(call goja.FunctionCall) goja.Value {
		obj := call.Argument(0).Export().(*uiacc.UiObject)
		result := obj.GetClickable()
		return vm.ToValue(result)
	})

	uiaccObj.Set("getLongClickable", func(call goja.FunctionCall) goja.Value {
		obj := call.Argument(0).Export().(*uiacc.UiObject)
		result := obj.GetLongClickable()
		return vm.ToValue(result)
	})

	uiaccObj.Set("getCheckable", func(call goja.FunctionCall) goja.Value {
		obj := call.Argument(0).Export().(*uiacc.UiObject)
		result := obj.GetCheckable()
		return vm.ToValue(result)
	})

	uiaccObj.Set("getSelected", func(call goja.FunctionCall) goja.Value {
		obj := call.Argument(0).Export().(*uiacc.UiObject)
		result := obj.GetSelected()
		return vm.ToValue(result)
	})

	uiaccObj.Set("getEnabled", func(call goja.FunctionCall) goja.Value {
		obj := call.Argument(0).Export().(*uiacc.UiObject)
		result := obj.GetEnabled()
		return vm.ToValue(result)
	})

	uiaccObj.Set("getScrollable", func(call goja.FunctionCall) goja.Value {
		obj := call.Argument(0).Export().(*uiacc.UiObject)
		result := obj.GetScrollable()
		return vm.ToValue(result)
	})

	uiaccObj.Set("getEditable", func(call goja.FunctionCall) goja.Value {
		obj := call.Argument(0).Export().(*uiacc.UiObject)
		result := obj.GetEditable()
		return vm.ToValue(result)
	})

	uiaccObj.Set("getMultiLine", func(call goja.FunctionCall) goja.Value {
		obj := call.Argument(0).Export().(*uiacc.UiObject)
		result := obj.GetMultiLine()
		return vm.ToValue(result)
	})

	uiaccObj.Set("getChecked", func(call goja.FunctionCall) goja.Value {
		obj := call.Argument(0).Export().(*uiacc.UiObject)
		result := obj.GetChecked()
		return vm.ToValue(result)
	})

	uiaccObj.Set("getFocused", func(call goja.FunctionCall) goja.Value {
		obj := call.Argument(0).Export().(*uiacc.UiObject)
		result := obj.GetFocused()
		return vm.ToValue(result)
	})

	uiaccObj.Set("getFocusable", func(call goja.FunctionCall) goja.Value {
		obj := call.Argument(0).Export().(*uiacc.UiObject)
		result := obj.GetFocusable()
		return vm.ToValue(result)
	})

	uiaccObj.Set("getDismissable", func(call goja.FunctionCall) goja.Value {
		obj := call.Argument(0).Export().(*uiacc.UiObject)
		result := obj.GetDismissable()
		return vm.ToValue(result)
	})

	uiaccObj.Set("getContextClickable", func(call goja.FunctionCall) goja.Value {
		obj := call.Argument(0).Export().(*uiacc.UiObject)
		result := obj.GetContextClickable()
		return vm.ToValue(result)
	})

	uiaccObj.Set("getAccessibilityFocused", func(call goja.FunctionCall) goja.Value {
		obj := call.Argument(0).Export().(*uiacc.UiObject)
		result := obj.GetAccessibilityFocused()
		return vm.ToValue(result)
	})

	uiaccObj.Set("getChildCount", func(call goja.FunctionCall) goja.Value {
		obj := call.Argument(0).Export().(*uiacc.UiObject)
		result := obj.GetChildCount()
		return vm.ToValue(result)
	})

	uiaccObj.Set("getDrawingOrder", func(call goja.FunctionCall) goja.Value {
		obj := call.Argument(0).Export().(*uiacc.UiObject)
		result := obj.GetDrawingOrder()
		return vm.ToValue(result)
	})

	uiaccObj.Set("getIndex", func(call goja.FunctionCall) goja.Value {
		obj := call.Argument(0).Export().(*uiacc.UiObject)
		result := obj.GetIndex()
		return vm.ToValue(result)
	})

	uiaccObj.Set("getBounds", func(call goja.FunctionCall) goja.Value {
		obj := call.Argument(0).Export().(*uiacc.UiObject)
		result := obj.GetBounds()
		return vm.ToValue(result)
	})

	uiaccObj.Set("getBoundsInParent", func(call goja.FunctionCall) goja.Value {
		obj := call.Argument(0).Export().(*uiacc.UiObject)
		result := obj.GetBoundsInParent()
		return vm.ToValue(result)
	})

	uiaccObj.Set("getId", func(call goja.FunctionCall) goja.Value {
		obj := call.Argument(0).Export().(*uiacc.UiObject)
		result := obj.GetId()
		return vm.ToValue(result)
	})

	uiaccObj.Set("getText", func(call goja.FunctionCall) goja.Value {
		obj := call.Argument(0).Export().(*uiacc.UiObject)
		result := obj.GetText()
		return vm.ToValue(result)
	})

	uiaccObj.Set("getDesc", func(call goja.FunctionCall) goja.Value {
		obj := call.Argument(0).Export().(*uiacc.UiObject)
		result := obj.GetDesc()
		return vm.ToValue(result)
	})

	uiaccObj.Set("getPackageName", func(call goja.FunctionCall) goja.Value {
		obj := call.Argument(0).Export().(*uiacc.UiObject)
		result := obj.GetPackageName()
		return vm.ToValue(result)
	})

	uiaccObj.Set("getClassName", func(call goja.FunctionCall) goja.Value {
		obj := call.Argument(0).Export().(*uiacc.UiObject)
		result := obj.GetClassName()
		return vm.ToValue(result)
	})

	uiaccObj.Set("getParent", func(call goja.FunctionCall) goja.Value {
		obj := call.Argument(0).Export().(*uiacc.UiObject)
		result := obj.GetParent()
		return vm.ToValue(result)
	})

	uiaccObj.Set("getChild", func(call goja.FunctionCall) goja.Value {
		obj := call.Argument(0).Export().(*uiacc.UiObject)
		index := int(call.Argument(1).ToInteger())
		result := obj.GetChild(index)
		return vm.ToValue(result)
	})

	uiaccObj.Set("getChildren", func(call goja.FunctionCall) goja.Value {
		obj := call.Argument(0).Export().(*uiacc.UiObject)
		result := obj.GetChildren()
		return vm.ToValue(result)
	})

	uiaccObj.Set("getVisible", func(call goja.FunctionCall) goja.Value {
		obj := call.Argument(0).Export().(*uiacc.UiObject)
		result := obj.GetVisible()
		return vm.ToValue(result)
	})

	uiaccObj.Set("getPassword", func(call goja.FunctionCall) goja.Value {
		obj := call.Argument(0).Export().(*uiacc.UiObject)
		result := obj.GetPassword()
		return vm.ToValue(result)
	})

	uiaccObj.Set("toString", func(call goja.FunctionCall) goja.Value {
		obj := call.Argument(0).Export().(*uiacc.UiObject)
		result := obj.ToString()
		return vm.ToValue(result)
	})

	// 注册方法到文档
	engine.RegisterMethod("uiacc.new", "创建一个新的Accessibility对象", func(displayId int) *uiacc.Uiacc { return uiacc.New(displayId) }, true)
	engine.RegisterMethod("uiacc.text", "设置选择器的text属性", (*uiacc.Uiacc).Text, true)
	engine.RegisterMethod("uiacc.textContains", "设置选择器的textContains属性", (*uiacc.Uiacc).TextContains, true)
	engine.RegisterMethod("uiacc.textStartsWith", "设置选择器的textStartsWith属性", (*uiacc.Uiacc).TextStartsWith, true)
	engine.RegisterMethod("uiacc.textEndsWith", "设置选择器的textEndsWith属性", (*uiacc.Uiacc).TextEndsWith, true)
	engine.RegisterMethod("uiacc.textMatches", "设置选择器的textMatches属性", (*uiacc.Uiacc).TextMatches, true)
	engine.RegisterMethod("uiacc.desc", "设置选择器的desc属性", (*uiacc.Uiacc).Desc, true)
	engine.RegisterMethod("uiacc.descContains", "设置选择器的descContains属性", (*uiacc.Uiacc).DescContains, true)
	engine.RegisterMethod("uiacc.descStartsWith", "设置选择器的descStartsWith属性", (*uiacc.Uiacc).DescStartsWith, true)
	engine.RegisterMethod("uiacc.descEndsWith", "设置选择器的descEndsWith属性", (*uiacc.Uiacc).DescEndsWith, true)
	engine.RegisterMethod("uiacc.descMatches", "设置选择器的descMatches属性", (*uiacc.Uiacc).DescMatches, true)
	engine.RegisterMethod("uiacc.id", "设置选择器的id属性", (*uiacc.Uiacc).Id, true)
	engine.RegisterMethod("uiacc.idContains", "设置选择器的idContains属性", (*uiacc.Uiacc).IdContains, true)
	engine.RegisterMethod("uiacc.idStartsWith", "设置选择器的idStartsWith属性", (*uiacc.Uiacc).IdStartsWith, true)
	engine.RegisterMethod("uiacc.idEndsWith", "设置选择器的idEndsWith属性", (*uiacc.Uiacc).IdEndsWith, true)
	engine.RegisterMethod("uiacc.idMatches", "设置选择器的idMatches属性", (*uiacc.Uiacc).IdMatches, true)
	engine.RegisterMethod("uiacc.className", "设置选择器的className属性", (*uiacc.Uiacc).ClassName, true)
	engine.RegisterMethod("uiacc.classNameContains", "设置选择器的classNameContains属性", (*uiacc.Uiacc).ClassNameContains, true)
	engine.RegisterMethod("uiacc.classNameStartsWith", "设置选择器的classNameStartsWith属性", (*uiacc.Uiacc).ClassNameStartsWith, true)
	engine.RegisterMethod("uiacc.classNameEndsWith", "设置选择器的classNameEndsWith属性", (*uiacc.Uiacc).ClassNameEndsWith, true)
	engine.RegisterMethod("uiacc.classNameMatches", "设置选择器的classNameMatches属性", (*uiacc.Uiacc).ClassNameMatches, true)
	engine.RegisterMethod("uiacc.packageName", "设置选择器的packageName属性", (*uiacc.Uiacc).PackageName, true)
	engine.RegisterMethod("uiacc.packageNameContains", "设置选择器的packageNameContains属性", (*uiacc.Uiacc).PackageNameContains, true)
	engine.RegisterMethod("uiacc.packageNameStartsWith", "设置选择器的packageNameStartsWith属性", (*uiacc.Uiacc).PackageNameStartsWith, true)
	engine.RegisterMethod("uiacc.packageNameEndsWith", "设置选择器的packageNameEndsWith属性", (*uiacc.Uiacc).PackageNameEndsWith, true)
	engine.RegisterMethod("uiacc.packageNameMatches", "设置选择器的packageNameMatches属性", (*uiacc.Uiacc).PackageNameMatches, true)
	engine.RegisterMethod("uiacc.bounds", "设置选择器的bounds属性", (*uiacc.Uiacc).Bounds, true)
	engine.RegisterMethod("uiacc.boundsInside", "设置选择器的boundsInside属性", (*uiacc.Uiacc).BoundsInside, true)
	engine.RegisterMethod("uiacc.boundsContains", "设置选择器的boundsContains属性", (*uiacc.Uiacc).BoundsContains, true)
	engine.RegisterMethod("uiacc.drawingOrder", "设置选择器的drawingOrder属性", (*uiacc.Uiacc).DrawingOrder, true)
	engine.RegisterMethod("uiacc.clickable", "设置选择器的clickable属性", (*uiacc.Uiacc).Clickable, true)
	engine.RegisterMethod("uiacc.longClickable", "设置选择器的longClickable属性", (*uiacc.Uiacc).LongClickable, true)
	engine.RegisterMethod("uiacc.checkable", "设置选择器的checkable属性", (*uiacc.Uiacc).Checkable, true)
	engine.RegisterMethod("uiacc.selected", "设置选择器的selected属性", (*uiacc.Uiacc).Selected, true)
	engine.RegisterMethod("uiacc.enabled", "设置选择器的enabled属性", (*uiacc.Uiacc).Enabled, true)
	engine.RegisterMethod("uiacc.scrollable", "设置选择器的scrollable属性", (*uiacc.Uiacc).Scrollable, true)
	engine.RegisterMethod("uiacc.editable", "设置选择器的editable属性", (*uiacc.Uiacc).Editable, true)
	engine.RegisterMethod("uiacc.multiLine", "设置选择器的multiLine属性", (*uiacc.Uiacc).MultiLine, true)
	engine.RegisterMethod("uiacc.checked", "设置选择器的checked属性", (*uiacc.Uiacc).Checked, true)
	engine.RegisterMethod("uiacc.focusable", "设置选择器的focusable属性", (*uiacc.Uiacc).Focusable, true)
	engine.RegisterMethod("uiacc.dismissable", "设置选择器的dismissable属性", (*uiacc.Uiacc).Dismissable, true)
	engine.RegisterMethod("uiacc.focused", "设置选择器的focused属性", (*uiacc.Uiacc).Focused, true)
	engine.RegisterMethod("uiacc.contextClickable", "设置选择器的contextClickable属性", (*uiacc.Uiacc).ContextClickable, true)
	engine.RegisterMethod("uiacc.index", "设置选择器的index属性", (*uiacc.Uiacc).Index, true)
	engine.RegisterMethod("uiacc.click", "点击匹配的控件", (*uiacc.Uiacc).Click, true)
	engine.RegisterMethod("uiacc.waitFor", "等待控件出现", (*uiacc.Uiacc).WaitFor, true)
	engine.RegisterMethod("uiacc.findOnce", "查找并返回第一个匹配的控件", (*uiacc.Uiacc).FindOnce, true)
	engine.RegisterMethod("uiacc.find", "查找并返回所有匹配的控件", (*uiacc.Uiacc).Find, true)
	engine.RegisterMethod("uiacc.objClick", "点击控件", (*uiacc.UiObject).Click, true)
	engine.RegisterMethod("uiacc.clickCenter", "点击控件中心", (*uiacc.UiObject).ClickCenter, true)
	engine.RegisterMethod("uiacc.clickLongClick", "长按控件", (*uiacc.UiObject).ClickLongClick, true)
	engine.RegisterMethod("uiacc.copy", "复制文本", (*uiacc.UiObject).Copy, true)
	engine.RegisterMethod("uiacc.cut", "剪切文本", (*uiacc.UiObject).Cut, true)
	engine.RegisterMethod("uiacc.paste", "粘贴文本", (*uiacc.UiObject).Paste, true)
	engine.RegisterMethod("uiacc.scrollForward", "向前滚动", (*uiacc.UiObject).ScrollForward, true)
	engine.RegisterMethod("uiacc.scrollBackward", "向后滚动", (*uiacc.UiObject).ScrollBackward, true)
	engine.RegisterMethod("uiacc.collapse", "折叠", (*uiacc.UiObject).Collapse, true)
	engine.RegisterMethod("uiacc.expand", "展开", (*uiacc.UiObject).Expand, true)
	engine.RegisterMethod("uiacc.show", "显示", (*uiacc.UiObject).Show, true)
	engine.RegisterMethod("uiacc.select", "选择", (*uiacc.UiObject).Select, true)
	engine.RegisterMethod("uiacc.clearSelect", "清除选择", (*uiacc.UiObject).ClearSelect, true)
	engine.RegisterMethod("uiacc.setSelection", "设置选择", (*uiacc.UiObject).SetSelection, true)
	engine.RegisterMethod("uiacc.setVisibleToUser", "设置是否对用户可见", (*uiacc.UiObject).SetVisibleToUser, true)
	engine.RegisterMethod("uiacc.setText", "设置文本", (*uiacc.UiObject).SetText, true)
	engine.RegisterMethod("uiacc.getClickable", "获取clickable属性", (*uiacc.UiObject).GetClickable, true)
	engine.RegisterMethod("uiacc.getLongClickable", "获取longClickable属性", (*uiacc.UiObject).GetLongClickable, true)
	engine.RegisterMethod("uiacc.getCheckable", "获取checkable属性", (*uiacc.UiObject).GetCheckable, true)
	engine.RegisterMethod("uiacc.getSelected", "获取selected属性", (*uiacc.UiObject).GetSelected, true)
	engine.RegisterMethod("uiacc.getEnabled", "获取enabled属性", (*uiacc.UiObject).GetEnabled, true)
	engine.RegisterMethod("uiacc.getScrollable", "获取scrollable属性", (*uiacc.UiObject).GetScrollable, true)
	engine.RegisterMethod("uiacc.getEditable", "获取editable属性", (*uiacc.UiObject).GetEditable, true)
	engine.RegisterMethod("uiacc.getMultiLine", "获取multiLine属性", (*uiacc.UiObject).GetMultiLine, true)
	engine.RegisterMethod("uiacc.getChecked", "获取checked属性", (*uiacc.UiObject).GetChecked, true)
	engine.RegisterMethod("uiacc.getFocused", "获取focused属性", (*uiacc.UiObject).GetFocused, true)
	engine.RegisterMethod("uiacc.getFocusable", "获取focusable属性", (*uiacc.UiObject).GetFocusable, true)
	engine.RegisterMethod("uiacc.getDismissable", "获取dismissable属性", (*uiacc.UiObject).GetDismissable, true)
	engine.RegisterMethod("uiacc.getContextClickable", "获取contextClickable属性", (*uiacc.UiObject).GetContextClickable, true)
	engine.RegisterMethod("uiacc.getAccessibilityFocused", "获取accessibilityFocused属性", (*uiacc.UiObject).GetAccessibilityFocused, true)
	engine.RegisterMethod("uiacc.getChildCount", "获取子控件数量", (*uiacc.UiObject).GetChildCount, true)
	engine.RegisterMethod("uiacc.getDrawingOrder", "获取绘制顺序", (*uiacc.UiObject).GetDrawingOrder, true)
	engine.RegisterMethod("uiacc.getIndex", "获取索引", (*uiacc.UiObject).GetIndex, true)
	engine.RegisterMethod("uiacc.getBounds", "获取边界", (*uiacc.UiObject).GetBounds, true)
	engine.RegisterMethod("uiacc.getBoundsInParent", "获取在父控件中的边界", (*uiacc.UiObject).GetBoundsInParent, true)
	engine.RegisterMethod("uiacc.getId", "获取ID", (*uiacc.UiObject).GetId, true)
	engine.RegisterMethod("uiacc.getText", "获取文本", (*uiacc.UiObject).GetText, true)
	engine.RegisterMethod("uiacc.getDesc", "获取描述", (*uiacc.UiObject).GetDesc, true)
	engine.RegisterMethod("uiacc.getPackageName", "获取包名", (*uiacc.UiObject).GetPackageName, true)
	engine.RegisterMethod("uiacc.getClassName", "获取类名", (*uiacc.UiObject).GetClassName, true)
	engine.RegisterMethod("uiacc.getParent", "获取父控件", (*uiacc.UiObject).GetParent, true)
	engine.RegisterMethod("uiacc.getChild", "获取子控件", (*uiacc.UiObject).GetChild, true)
	engine.RegisterMethod("uiacc.getChildren", "获取所有子控件", (*uiacc.UiObject).GetChildren, true)

	return nil
}
