package uiacc

import (
	"strconv"

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

	uiaccObj.Set("new", func(call goja.FunctionCall) goja.Value {
		displayId := 0
		if len(call.Arguments) >= 1 {
			displayId = int(call.Argument(0).ToInteger())
		}
		u := uiacc.New(displayId)
		if u != nil {
			return wrapUiacc(vm, u)
		}
		return goja.Null()
	})

	engine.RegisterMethod("uiacc.new", "创建一个新的Accessibility对象", func(displayId int) *uiacc.Uiacc { return uiacc.New(displayId) }, true)

	return nil
}

func wrapUiacc(vm *goja.Runtime, u *uiacc.Uiacc) goja.Value {
	obj := vm.NewObject()

	obj.Set("text", func(call goja.FunctionCall) goja.Value {
		value := call.Argument(0).String()
		result := u.Text(value)
		return wrapUiacc(vm, result)
	})

	obj.Set("textContains", func(call goja.FunctionCall) goja.Value {
		value := call.Argument(0).String()
		result := u.TextContains(value)
		return wrapUiacc(vm, result)
	})

	obj.Set("textStartsWith", func(call goja.FunctionCall) goja.Value {
		value := call.Argument(0).String()
		result := u.TextStartsWith(value)
		return wrapUiacc(vm, result)
	})

	obj.Set("textEndsWith", func(call goja.FunctionCall) goja.Value {
		value := call.Argument(0).String()
		result := u.TextEndsWith(value)
		return wrapUiacc(vm, result)
	})

	obj.Set("textMatches", func(call goja.FunctionCall) goja.Value {
		value := call.Argument(0).String()
		result := u.TextMatches(value)
		return wrapUiacc(vm, result)
	})

	obj.Set("desc", func(call goja.FunctionCall) goja.Value {
		value := call.Argument(0).String()
		result := u.Desc(value)
		return wrapUiacc(vm, result)
	})

	obj.Set("descContains", func(call goja.FunctionCall) goja.Value {
		value := call.Argument(0).String()
		result := u.DescContains(value)
		return wrapUiacc(vm, result)
	})

	obj.Set("descStartsWith", func(call goja.FunctionCall) goja.Value {
		value := call.Argument(0).String()
		result := u.DescStartsWith(value)
		return wrapUiacc(vm, result)
	})

	obj.Set("descEndsWith", func(call goja.FunctionCall) goja.Value {
		value := call.Argument(0).String()
		result := u.DescEndsWith(value)
		return wrapUiacc(vm, result)
	})

	obj.Set("descMatches", func(call goja.FunctionCall) goja.Value {
		value := call.Argument(0).String()
		result := u.DescMatches(value)
		return wrapUiacc(vm, result)
	})

	obj.Set("id", func(call goja.FunctionCall) goja.Value {
		value := call.Argument(0).String()
		result := u.Id(value)
		return wrapUiacc(vm, result)
	})

	obj.Set("idContains", func(call goja.FunctionCall) goja.Value {
		value := call.Argument(0).String()
		result := u.IdContains(value)
		return wrapUiacc(vm, result)
	})

	obj.Set("idStartsWith", func(call goja.FunctionCall) goja.Value {
		value := call.Argument(0).String()
		result := u.IdStartsWith(value)
		return wrapUiacc(vm, result)
	})

	obj.Set("idEndsWith", func(call goja.FunctionCall) goja.Value {
		value := call.Argument(0).String()
		result := u.IdEndsWith(value)
		return wrapUiacc(vm, result)
	})

	obj.Set("idMatches", func(call goja.FunctionCall) goja.Value {
		value := call.Argument(0).String()
		result := u.IdMatches(value)
		return wrapUiacc(vm, result)
	})

	obj.Set("className", func(call goja.FunctionCall) goja.Value {
		value := call.Argument(0).String()
		result := u.ClassName(value)
		return wrapUiacc(vm, result)
	})

	obj.Set("classNameContains", func(call goja.FunctionCall) goja.Value {
		value := call.Argument(0).String()
		result := u.ClassNameContains(value)
		return wrapUiacc(vm, result)
	})

	obj.Set("classNameStartsWith", func(call goja.FunctionCall) goja.Value {
		value := call.Argument(0).String()
		result := u.ClassNameStartsWith(value)
		return wrapUiacc(vm, result)
	})

	obj.Set("classNameEndsWith", func(call goja.FunctionCall) goja.Value {
		value := call.Argument(0).String()
		result := u.ClassNameEndsWith(value)
		return wrapUiacc(vm, result)
	})

	obj.Set("classNameMatches", func(call goja.FunctionCall) goja.Value {
		value := call.Argument(0).String()
		result := u.ClassNameMatches(value)
		return wrapUiacc(vm, result)
	})

	obj.Set("packageName", func(call goja.FunctionCall) goja.Value {
		value := call.Argument(0).String()
		result := u.PackageName(value)
		return wrapUiacc(vm, result)
	})

	obj.Set("packageNameContains", func(call goja.FunctionCall) goja.Value {
		value := call.Argument(0).String()
		result := u.PackageNameContains(value)
		return wrapUiacc(vm, result)
	})

	obj.Set("packageNameStartsWith", func(call goja.FunctionCall) goja.Value {
		value := call.Argument(0).String()
		result := u.PackageNameStartsWith(value)
		return wrapUiacc(vm, result)
	})

	obj.Set("packageNameEndsWith", func(call goja.FunctionCall) goja.Value {
		value := call.Argument(0).String()
		result := u.PackageNameEndsWith(value)
		return wrapUiacc(vm, result)
	})

	obj.Set("packageNameMatches", func(call goja.FunctionCall) goja.Value {
		value := call.Argument(0).String()
		result := u.PackageNameMatches(value)
		return wrapUiacc(vm, result)
	})

	obj.Set("bounds", func(call goja.FunctionCall) goja.Value {
		left := int(call.Argument(0).ToInteger())
		top := int(call.Argument(1).ToInteger())
		right := int(call.Argument(2).ToInteger())
		bottom := int(call.Argument(3).ToInteger())
		result := u.Bounds(left, top, right, bottom)
		return wrapUiacc(vm, result)
	})

	obj.Set("boundsInside", func(call goja.FunctionCall) goja.Value {
		left := int(call.Argument(0).ToInteger())
		top := int(call.Argument(1).ToInteger())
		right := int(call.Argument(2).ToInteger())
		bottom := int(call.Argument(3).ToInteger())
		result := u.BoundsInside(left, top, right, bottom)
		return wrapUiacc(vm, result)
	})

	obj.Set("boundsContains", func(call goja.FunctionCall) goja.Value {
		left := int(call.Argument(0).ToInteger())
		top := int(call.Argument(1).ToInteger())
		right := int(call.Argument(2).ToInteger())
		bottom := int(call.Argument(3).ToInteger())
		result := u.BoundsContains(left, top, right, bottom)
		return wrapUiacc(vm, result)
	})

	obj.Set("drawingOrder", func(call goja.FunctionCall) goja.Value {
		value := int(call.Argument(0).ToInteger())
		result := u.DrawingOrder(value)
		return wrapUiacc(vm, result)
	})

	obj.Set("clickable", func(call goja.FunctionCall) goja.Value {
		value := call.Argument(0).ToBoolean()
		result := u.Clickable(value)
		return wrapUiacc(vm, result)
	})

	obj.Set("longClickable", func(call goja.FunctionCall) goja.Value {
		value := call.Argument(0).ToBoolean()
		result := u.LongClickable(value)
		return wrapUiacc(vm, result)
	})

	obj.Set("checkable", func(call goja.FunctionCall) goja.Value {
		value := call.Argument(0).ToBoolean()
		result := u.Checkable(value)
		return wrapUiacc(vm, result)
	})

	obj.Set("selected", func(call goja.FunctionCall) goja.Value {
		value := call.Argument(0).ToBoolean()
		result := u.Selected(value)
		return wrapUiacc(vm, result)
	})

	obj.Set("enabled", func(call goja.FunctionCall) goja.Value {
		value := call.Argument(0).ToBoolean()
		result := u.Enabled(value)
		return wrapUiacc(vm, result)
	})

	obj.Set("scrollable", func(call goja.FunctionCall) goja.Value {
		value := call.Argument(0).ToBoolean()
		result := u.Scrollable(value)
		return wrapUiacc(vm, result)
	})

	obj.Set("editable", func(call goja.FunctionCall) goja.Value {
		value := call.Argument(0).ToBoolean()
		result := u.Editable(value)
		return wrapUiacc(vm, result)
	})

	obj.Set("multiLine", func(call goja.FunctionCall) goja.Value {
		value := call.Argument(0).ToBoolean()
		result := u.MultiLine(value)
		return wrapUiacc(vm, result)
	})

	obj.Set("checked", func(call goja.FunctionCall) goja.Value {
		value := call.Argument(0).ToBoolean()
		result := u.Checked(value)
		return wrapUiacc(vm, result)
	})

	obj.Set("focusable", func(call goja.FunctionCall) goja.Value {
		value := call.Argument(0).ToBoolean()
		result := u.Focusable(value)
		return wrapUiacc(vm, result)
	})

	obj.Set("dismissable", func(call goja.FunctionCall) goja.Value {
		value := call.Argument(0).ToBoolean()
		result := u.Dismissable(value)
		return wrapUiacc(vm, result)
	})

	obj.Set("focused", func(call goja.FunctionCall) goja.Value {
		value := call.Argument(0).ToBoolean()
		result := u.Focused(value)
		return wrapUiacc(vm, result)
	})

	obj.Set("contextClickable", func(call goja.FunctionCall) goja.Value {
		value := call.Argument(0).ToBoolean()
		result := u.ContextClickable(value)
		return wrapUiacc(vm, result)
	})

	obj.Set("index", func(call goja.FunctionCall) goja.Value {
		value := int(call.Argument(0).ToInteger())
		result := u.Index(value)
		return wrapUiacc(vm, result)
	})

	obj.Set("visible", func(call goja.FunctionCall) goja.Value {
		value := call.Argument(0).ToBoolean()
		result := u.Visible(value)
		return wrapUiacc(vm, result)
	})

	obj.Set("password", func(call goja.FunctionCall) goja.Value {
		value := call.Argument(0).ToBoolean()
		result := u.Password(value)
		return wrapUiacc(vm, result)
	})

	obj.Set("click", func(call goja.FunctionCall) goja.Value {
		text := call.Argument(0).String()
		result := u.Click(text)
		return vm.ToValue(result)
	})

	obj.Set("waitFor", func(call goja.FunctionCall) goja.Value {
		timeout := int(call.Argument(0).ToInteger())
		result := u.WaitFor(timeout)
		if result != nil {
			return wrapUiObject(vm, result)
		}
		return goja.Null()
	})

	obj.Set("findOnce", func(call goja.FunctionCall) goja.Value {
		result := u.FindOnce()
		if result != nil {
			return wrapUiObject(vm, result)
		}
		return goja.Null()
	})

	obj.Set("find", func(call goja.FunctionCall) goja.Value {
		result := u.Find()
		arr := vm.NewArray()
		for i, item := range result {
			arr.Set(strconv.Itoa(i), wrapUiObject(vm, item))
		}
		return arr
	})

	obj.Set("release", func(call goja.FunctionCall) goja.Value {
		u.Release()
		return goja.Undefined()
	})

	return vm.ToValue(obj)
}

func wrapUiObject(vm *goja.Runtime, obj *uiacc.UiObject) goja.Value {
	o := vm.NewObject()

	o.Set("click", func(call goja.FunctionCall) goja.Value {
		result := obj.Click()
		return vm.ToValue(result)
	})

	o.Set("clickCenter", func(call goja.FunctionCall) goja.Value {
		result := obj.ClickCenter()
		return vm.ToValue(result)
	})

	o.Set("clickLongClick", func(call goja.FunctionCall) goja.Value {
		result := obj.ClickLongClick()
		return vm.ToValue(result)
	})

	o.Set("copy", func(call goja.FunctionCall) goja.Value {
		result := obj.Copy()
		return vm.ToValue(result)
	})

	o.Set("cut", func(call goja.FunctionCall) goja.Value {
		result := obj.Cut()
		return vm.ToValue(result)
	})

	o.Set("paste", func(call goja.FunctionCall) goja.Value {
		result := obj.Paste()
		return vm.ToValue(result)
	})

	o.Set("scrollForward", func(call goja.FunctionCall) goja.Value {
		result := obj.ScrollForward()
		return vm.ToValue(result)
	})

	o.Set("scrollBackward", func(call goja.FunctionCall) goja.Value {
		result := obj.ScrollBackward()
		return vm.ToValue(result)
	})

	o.Set("collapse", func(call goja.FunctionCall) goja.Value {
		result := obj.Collapse()
		return vm.ToValue(result)
	})

	o.Set("expand", func(call goja.FunctionCall) goja.Value {
		result := obj.Expand()
		return vm.ToValue(result)
	})

	o.Set("show", func(call goja.FunctionCall) goja.Value {
		result := obj.Show()
		return vm.ToValue(result)
	})

	o.Set("select", func(call goja.FunctionCall) goja.Value {
		result := obj.Select()
		return vm.ToValue(result)
	})

	o.Set("clearSelect", func(call goja.FunctionCall) goja.Value {
		result := obj.ClearSelect()
		return vm.ToValue(result)
	})

	o.Set("setSelection", func(call goja.FunctionCall) goja.Value {
		start := int(call.Argument(0).ToInteger())
		end := int(call.Argument(1).ToInteger())
		result := obj.SetSelection(start, end)
		return vm.ToValue(result)
	})

	o.Set("setVisibleToUser", func(call goja.FunctionCall) goja.Value {
		isVisible := call.Argument(0).ToBoolean()
		result := obj.SetVisibleToUser(isVisible)
		return vm.ToValue(result)
	})

	o.Set("setText", func(call goja.FunctionCall) goja.Value {
		str := call.Argument(0).String()
		result := obj.SetText(str)
		return vm.ToValue(result)
	})

	o.Set("getClickable", func(call goja.FunctionCall) goja.Value {
		result := obj.GetClickable()
		return vm.ToValue(result)
	})

	o.Set("getLongClickable", func(call goja.FunctionCall) goja.Value {
		result := obj.GetLongClickable()
		return vm.ToValue(result)
	})

	o.Set("getCheckable", func(call goja.FunctionCall) goja.Value {
		result := obj.GetCheckable()
		return vm.ToValue(result)
	})

	o.Set("getSelected", func(call goja.FunctionCall) goja.Value {
		result := obj.GetSelected()
		return vm.ToValue(result)
	})

	o.Set("getEnabled", func(call goja.FunctionCall) goja.Value {
		result := obj.GetEnabled()
		return vm.ToValue(result)
	})

	o.Set("getScrollable", func(call goja.FunctionCall) goja.Value {
		result := obj.GetScrollable()
		return vm.ToValue(result)
	})

	o.Set("getEditable", func(call goja.FunctionCall) goja.Value {
		result := obj.GetEditable()
		return vm.ToValue(result)
	})

	o.Set("getMultiLine", func(call goja.FunctionCall) goja.Value {
		result := obj.GetMultiLine()
		return vm.ToValue(result)
	})

	o.Set("getChecked", func(call goja.FunctionCall) goja.Value {
		result := obj.GetChecked()
		return vm.ToValue(result)
	})

	o.Set("getFocused", func(call goja.FunctionCall) goja.Value {
		result := obj.GetFocused()
		return vm.ToValue(result)
	})

	o.Set("getFocusable", func(call goja.FunctionCall) goja.Value {
		result := obj.GetFocusable()
		return vm.ToValue(result)
	})

	o.Set("getDismissable", func(call goja.FunctionCall) goja.Value {
		result := obj.GetDismissable()
		return vm.ToValue(result)
	})

	o.Set("getContextClickable", func(call goja.FunctionCall) goja.Value {
		result := obj.GetContextClickable()
		return vm.ToValue(result)
	})

	o.Set("getAccessibilityFocused", func(call goja.FunctionCall) goja.Value {
		result := obj.GetAccessibilityFocused()
		return vm.ToValue(result)
	})

	o.Set("getChildCount", func(call goja.FunctionCall) goja.Value {
		result := obj.GetChildCount()
		return vm.ToValue(result)
	})

	o.Set("getDrawingOrder", func(call goja.FunctionCall) goja.Value {
		result := obj.GetDrawingOrder()
		return vm.ToValue(result)
	})

	o.Set("getIndex", func(call goja.FunctionCall) goja.Value {
		result := obj.GetIndex()
		return vm.ToValue(result)
	})

	o.Set("getBounds", func(call goja.FunctionCall) goja.Value {
		result := obj.GetBounds()
		return vm.ToValue(result)
	})

	o.Set("getBoundsInParent", func(call goja.FunctionCall) goja.Value {
		result := obj.GetBoundsInParent()
		return vm.ToValue(result)
	})

	o.Set("getId", func(call goja.FunctionCall) goja.Value {
		result := obj.GetId()
		return vm.ToValue(result)
	})

	o.Set("getText", func(call goja.FunctionCall) goja.Value {
		result := obj.GetText()
		return vm.ToValue(result)
	})

	o.Set("getDesc", func(call goja.FunctionCall) goja.Value {
		result := obj.GetDesc()
		return vm.ToValue(result)
	})

	o.Set("getPackageName", func(call goja.FunctionCall) goja.Value {
		result := obj.GetPackageName()
		return vm.ToValue(result)
	})

	o.Set("getClassName", func(call goja.FunctionCall) goja.Value {
		result := obj.GetClassName()
		return vm.ToValue(result)
	})

	o.Set("getParent", func(call goja.FunctionCall) goja.Value {
		result := obj.GetParent()
		if result != nil {
			return wrapUiObject(vm, result)
		}
		return goja.Null()
	})

	o.Set("getChild", func(call goja.FunctionCall) goja.Value {
		index := int(call.Argument(0).ToInteger())
		result := obj.GetChild(index)
		if result != nil {
			return wrapUiObject(vm, result)
		}
		return goja.Null()
	})

	o.Set("getChildren", func(call goja.FunctionCall) goja.Value {
		result := obj.GetChildren()
		arr := vm.NewArray()
		for i, item := range result {
			arr.Set(strconv.Itoa(i), wrapUiObject(vm, item))
		}
		return arr
	})

	o.Set("getVisible", func(call goja.FunctionCall) goja.Value {
		result := obj.GetVisible()
		return vm.ToValue(result)
	})

	o.Set("getPassword", func(call goja.FunctionCall) goja.Value {
		result := obj.GetPassword()
		return vm.ToValue(result)
	})

	o.Set("toString", func(call goja.FunctionCall) goja.Value {
		result := obj.ToString()
		return vm.ToValue(result)
	})

	return vm.ToValue(o)
}
