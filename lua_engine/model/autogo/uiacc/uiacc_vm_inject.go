package uiacc

import (
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"

	autogouiacc "github.com/Dasongzi1366/AutoGo/uiacc"
)

// UiaccModule 是 go-lua-vm 迁移后的模块壳。
type UiaccModule struct{}

func New() *UiaccModule { return &UiaccModule{} }

func (m *UiaccModule) Name() string { return "uiacc" }

func (m *UiaccModule) IsAvailable() bool { return true }

func (m *UiaccModule) Register(engine model.Engine) error {
	engine.RegisterMethod("uiacc.new", "创建新的Uiacc对象", func(displayID ...int) map[string]interface{} {
		return wrapUiacc(autogouiacc.New(optionalInt(0, displayID...)))
	}, true)
	engine.RegisterMethod("uiacc.text", "Uiacc.Text", (*autogouiacc.Uiacc).Text, true)
	engine.RegisterMethod("uiacc.textContains", "Uiacc.TextContains", (*autogouiacc.Uiacc).TextContains, true)
	engine.RegisterMethod("uiacc.textStartsWith", "Uiacc.TextStartsWith", (*autogouiacc.Uiacc).TextStartsWith, true)
	engine.RegisterMethod("uiacc.textEndsWith", "Uiacc.TextEndsWith", (*autogouiacc.Uiacc).TextEndsWith, true)
	engine.RegisterMethod("uiacc.textMatches", "Uiacc.TextMatches", (*autogouiacc.Uiacc).TextMatches, true)
	engine.RegisterMethod("uiacc.desc", "Uiacc.Desc", (*autogouiacc.Uiacc).Desc, true)
	engine.RegisterMethod("uiacc.descContains", "Uiacc.DescContains", (*autogouiacc.Uiacc).DescContains, true)
	engine.RegisterMethod("uiacc.descStartsWith", "Uiacc.DescStartsWith", (*autogouiacc.Uiacc).DescStartsWith, true)
	engine.RegisterMethod("uiacc.descEndsWith", "Uiacc.DescEndsWith", (*autogouiacc.Uiacc).DescEndsWith, true)
	engine.RegisterMethod("uiacc.descMatches", "Uiacc.DescMatches", (*autogouiacc.Uiacc).DescMatches, true)
	engine.RegisterMethod("uiacc.id", "Uiacc.Id", (*autogouiacc.Uiacc).Id, true)
	engine.RegisterMethod("uiacc.idContains", "Uiacc.IdContains", (*autogouiacc.Uiacc).IdContains, true)
	engine.RegisterMethod("uiacc.idStartsWith", "Uiacc.IdStartsWith", (*autogouiacc.Uiacc).IdStartsWith, true)
	engine.RegisterMethod("uiacc.idEndsWith", "Uiacc.IdEndsWith", (*autogouiacc.Uiacc).IdEndsWith, true)
	engine.RegisterMethod("uiacc.idMatches", "Uiacc.IdMatches", (*autogouiacc.Uiacc).IdMatches, true)
	engine.RegisterMethod("uiacc.className", "Uiacc.ClassName", (*autogouiacc.Uiacc).ClassName, true)
	engine.RegisterMethod("uiacc.classNameContains", "Uiacc.ClassNameContains", (*autogouiacc.Uiacc).ClassNameContains, true)
	engine.RegisterMethod("uiacc.classNameStartsWith", "Uiacc.ClassNameStartsWith", (*autogouiacc.Uiacc).ClassNameStartsWith, true)
	engine.RegisterMethod("uiacc.classNameEndsWith", "Uiacc.ClassNameEndsWith", (*autogouiacc.Uiacc).ClassNameEndsWith, true)
	engine.RegisterMethod("uiacc.classNameMatches", "Uiacc.ClassNameMatches", (*autogouiacc.Uiacc).ClassNameMatches, true)
	engine.RegisterMethod("uiacc.packageName", "Uiacc.PackageName", (*autogouiacc.Uiacc).PackageName, true)
	engine.RegisterMethod("uiacc.packageNameContains", "Uiacc.PackageNameContains", (*autogouiacc.Uiacc).PackageNameContains, true)
	engine.RegisterMethod("uiacc.packageNameStartsWith", "Uiacc.PackageNameStartsWith", (*autogouiacc.Uiacc).PackageNameStartsWith, true)
	engine.RegisterMethod("uiacc.packageNameEndsWith", "Uiacc.PackageNameEndsWith", (*autogouiacc.Uiacc).PackageNameEndsWith, true)
	engine.RegisterMethod("uiacc.packageNameMatches", "Uiacc.PackageNameMatches", (*autogouiacc.Uiacc).PackageNameMatches, true)
	engine.RegisterMethod("uiacc.bounds", "Uiacc.Bounds", (*autogouiacc.Uiacc).Bounds, true)
	engine.RegisterMethod("uiacc.boundsInside", "Uiacc.BoundsInside", (*autogouiacc.Uiacc).BoundsInside, true)
	engine.RegisterMethod("uiacc.boundsContains", "Uiacc.BoundsContains", (*autogouiacc.Uiacc).BoundsContains, true)
	engine.RegisterMethod("uiacc.drawingOrder", "Uiacc.DrawingOrder", (*autogouiacc.Uiacc).DrawingOrder, true)
	engine.RegisterMethod("uiacc.index", "Uiacc.Index", (*autogouiacc.Uiacc).Index, true)
	engine.RegisterMethod("uiacc.clickable", "Uiacc.Clickable", (*autogouiacc.Uiacc).Clickable, true)
	engine.RegisterMethod("uiacc.longClickable", "Uiacc.LongClickable", (*autogouiacc.Uiacc).LongClickable, true)
	engine.RegisterMethod("uiacc.checkable", "Uiacc.Checkable", (*autogouiacc.Uiacc).Checkable, true)
	engine.RegisterMethod("uiacc.selected", "Uiacc.Selected", (*autogouiacc.Uiacc).Selected, true)
	engine.RegisterMethod("uiacc.enabled", "Uiacc.Enabled", (*autogouiacc.Uiacc).Enabled, true)
	engine.RegisterMethod("uiacc.scrollable", "Uiacc.Scrollable", (*autogouiacc.Uiacc).Scrollable, true)
	engine.RegisterMethod("uiacc.editable", "Uiacc.Editable", (*autogouiacc.Uiacc).Editable, true)
	engine.RegisterMethod("uiacc.multiLine", "Uiacc.MultiLine", (*autogouiacc.Uiacc).MultiLine, true)
	engine.RegisterMethod("uiacc.checked", "Uiacc.Checked", (*autogouiacc.Uiacc).Checked, true)
	engine.RegisterMethod("uiacc.focusable", "Uiacc.Focusable", (*autogouiacc.Uiacc).Focusable, true)
	engine.RegisterMethod("uiacc.dismissable", "Uiacc.Dismissable", (*autogouiacc.Uiacc).Dismissable, true)
	engine.RegisterMethod("uiacc.focused", "Uiacc.Focused", (*autogouiacc.Uiacc).Focused, true)
	engine.RegisterMethod("uiacc.contextClickable", "Uiacc.ContextClickable", (*autogouiacc.Uiacc).ContextClickable, true)
	engine.RegisterMethod("uiacc.visible", "Uiacc.Visible", (*autogouiacc.Uiacc).Visible, true)
	engine.RegisterMethod("uiacc.password", "Uiacc.Password", (*autogouiacc.Uiacc).Password, true)
	engine.RegisterMethod("uiacc.click", "Uiacc.Click", (*autogouiacc.Uiacc).Click, true)
	engine.RegisterMethod("uiacc.waitFor", "Uiacc.WaitFor", (*autogouiacc.Uiacc).WaitFor, true)
	engine.RegisterMethod("uiacc.findOnce", "Uiacc.FindOnce", (*autogouiacc.Uiacc).FindOnce, true)
	engine.RegisterMethod("uiacc.find", "Uiacc.Find", (*autogouiacc.Uiacc).Find, true)
	engine.RegisterMethod("uiacc.release", "Uiacc.Release", (*autogouiacc.Uiacc).Release, true)
	engine.RegisterMethod("uiacc.clickCenter", "UiObject.ClickCenter", (*autogouiacc.UiObject).ClickCenter, true)
	engine.RegisterMethod("uiacc.clickLongClick", "UiObject.ClickLongClick", (*autogouiacc.UiObject).ClickLongClick, true)
	engine.RegisterMethod("uiacc.copy", "UiObject.Copy", (*autogouiacc.UiObject).Copy, true)
	engine.RegisterMethod("uiacc.cut", "UiObject.Cut", (*autogouiacc.UiObject).Cut, true)
	engine.RegisterMethod("uiacc.paste", "UiObject.Paste", (*autogouiacc.UiObject).Paste, true)
	engine.RegisterMethod("uiacc.scrollForward", "UiObject.ScrollForward", (*autogouiacc.UiObject).ScrollForward, true)
	engine.RegisterMethod("uiacc.scrollBackward", "UiObject.ScrollBackward", (*autogouiacc.UiObject).ScrollBackward, true)
	engine.RegisterMethod("uiacc.collapse", "UiObject.Collapse", (*autogouiacc.UiObject).Collapse, true)
	engine.RegisterMethod("uiacc.expand", "UiObject.Expand", (*autogouiacc.UiObject).Expand, true)
	engine.RegisterMethod("uiacc.show", "UiObject.Show", (*autogouiacc.UiObject).Show, true)
	engine.RegisterMethod("uiacc.select", "UiObject.Select", (*autogouiacc.UiObject).Select, true)
	engine.RegisterMethod("uiacc.clearSelect", "UiObject.ClearSelect", (*autogouiacc.UiObject).ClearSelect, true)
	engine.RegisterMethod("uiacc.getClickable", "UiObject.GetClickable", (*autogouiacc.UiObject).GetClickable, true)
	engine.RegisterMethod("uiacc.getLongClickable", "UiObject.GetLongClickable", (*autogouiacc.UiObject).GetLongClickable, true)
	engine.RegisterMethod("uiacc.getCheckable", "UiObject.GetCheckable", (*autogouiacc.UiObject).GetCheckable, true)
	engine.RegisterMethod("uiacc.getSelected", "UiObject.GetSelected", (*autogouiacc.UiObject).GetSelected, true)
	engine.RegisterMethod("uiacc.getEnabled", "UiObject.GetEnabled", (*autogouiacc.UiObject).GetEnabled, true)
	engine.RegisterMethod("uiacc.getScrollable", "UiObject.GetScrollable", (*autogouiacc.UiObject).GetScrollable, true)
	engine.RegisterMethod("uiacc.getEditable", "UiObject.GetEditable", (*autogouiacc.UiObject).GetEditable, true)
	engine.RegisterMethod("uiacc.getMultiLine", "UiObject.GetMultiLine", (*autogouiacc.UiObject).GetMultiLine, true)
	engine.RegisterMethod("uiacc.getChecked", "UiObject.GetChecked", (*autogouiacc.UiObject).GetChecked, true)
	engine.RegisterMethod("uiacc.getFocused", "UiObject.GetFocused", (*autogouiacc.UiObject).GetFocused, true)
	engine.RegisterMethod("uiacc.getFocusable", "UiObject.GetFocusable", (*autogouiacc.UiObject).GetFocusable, true)
	engine.RegisterMethod("uiacc.getDismissable", "UiObject.GetDismissable", (*autogouiacc.UiObject).GetDismissable, true)
	engine.RegisterMethod("uiacc.getContextClickable", "UiObject.GetContextClickable", (*autogouiacc.UiObject).GetContextClickable, true)
	engine.RegisterMethod("uiacc.getVisible", "UiObject.GetVisible", (*autogouiacc.UiObject).GetVisible, true)
	engine.RegisterMethod("uiacc.getPassword", "UiObject.GetPassword", (*autogouiacc.UiObject).GetPassword, true)
	engine.RegisterMethod("uiacc.getAccessibilityFocused", "UiObject.GetAccessibilityFocused", (*autogouiacc.UiObject).GetAccessibilityFocused, true)
	engine.RegisterMethod("uiacc.getChildCount", "UiObject.GetChildCount", (*autogouiacc.UiObject).GetChildCount, true)
	engine.RegisterMethod("uiacc.getDrawingOrder", "UiObject.GetDrawingOrder", (*autogouiacc.UiObject).GetDrawingOrder, true)
	engine.RegisterMethod("uiacc.getIndex", "UiObject.GetIndex", (*autogouiacc.UiObject).GetIndex, true)
	engine.RegisterMethod("uiacc.getBounds", "UiObject.GetBounds", (*autogouiacc.UiObject).GetBounds, true)
	engine.RegisterMethod("uiacc.getBoundsInParent", "UiObject.GetBoundsInParent", (*autogouiacc.UiObject).GetBoundsInParent, true)
	engine.RegisterMethod("uiacc.getId", "UiObject.GetId", (*autogouiacc.UiObject).GetId, true)
	engine.RegisterMethod("uiacc.getText", "UiObject.GetText", (*autogouiacc.UiObject).GetText, true)
	engine.RegisterMethod("uiacc.getDesc", "UiObject.GetDesc", (*autogouiacc.UiObject).GetDesc, true)
	engine.RegisterMethod("uiacc.getPackageName", "UiObject.GetPackageName", (*autogouiacc.UiObject).GetPackageName, true)
	engine.RegisterMethod("uiacc.getClassName", "UiObject.GetClassName", (*autogouiacc.UiObject).GetClassName, true)
	engine.RegisterMethod("uiacc.toString", "UiObject.ToString", (*autogouiacc.UiObject).ToString, true)
	engine.RegisterMethod("uiacc.setSelection", "UiObject.SetSelection", (*autogouiacc.UiObject).SetSelection, true)
	engine.RegisterMethod("uiacc.setVisibleToUser", "UiObject.SetVisibleToUser", (*autogouiacc.UiObject).SetVisibleToUser, true)
	engine.RegisterMethod("uiacc.setText", "UiObject.SetText", (*autogouiacc.UiObject).SetText, true)
	engine.RegisterMethod("uiacc.getParent", "UiObject.GetParent", (*autogouiacc.UiObject).GetParent, true)
	engine.RegisterMethod("uiacc.getChild", "UiObject.GetChild", (*autogouiacc.UiObject).GetChild, true)
	engine.RegisterMethod("uiacc.getChildren", "UiObject.GetChildren", (*autogouiacc.UiObject).GetChildren, true)
	return nil
}
func GetModule() model.Module { return &UiaccModule{} }

func wrapUiacc(selector *autogouiacc.Uiacc) map[string]interface{} {
	return map[string]interface{}{
		"text": func(value string) map[string]interface{} {
			if selector == nil {
				return wrapUiacc(nil)
			}
			return wrapUiacc(selector.Text(value))
		},
		"textContains": func(value string) map[string]interface{} {
			if selector == nil {
				return wrapUiacc(nil)
			}
			return wrapUiacc(selector.TextContains(value))
		},
		"descContains": func(value string) map[string]interface{} {
			if selector == nil {
				return wrapUiacc(nil)
			}
			return wrapUiacc(selector.DescContains(value))
		},
		"desc": func(value string) map[string]interface{} {
			if selector == nil {
				return wrapUiacc(nil)
			}
			return wrapUiacc(selector.Desc(value))
		},
		"id": func(value string) map[string]interface{} {
			if selector == nil {
				return wrapUiacc(nil)
			}
			return wrapUiacc(selector.Id(value))
		},
		"className": func(value string) map[string]interface{} {
			if selector == nil {
				return wrapUiacc(nil)
			}
			return wrapUiacc(selector.ClassName(value))
		},
		"packageName": func(value string) map[string]interface{} {
			if selector == nil {
				return wrapUiacc(nil)
			}
			return wrapUiacc(selector.PackageName(value))
		},
		"clickable": func(value bool) map[string]interface{} {
			if selector == nil {
				return wrapUiacc(nil)
			}
			return wrapUiacc(selector.Clickable(value))
		},
		"scrollable": func(value bool) map[string]interface{} {
			if selector == nil {
				return wrapUiacc(nil)
			}
			return wrapUiacc(selector.Scrollable(value))
		},
		"editable": func(value bool) map[string]interface{} {
			if selector == nil {
				return wrapUiacc(nil)
			}
			return wrapUiacc(selector.Editable(value))
		},
		// visible 将可见性条件暴露给 GLua 选择器，用于过滤页面切换后仍残留在无障碍节点缓存中的旧页面节点。
		"visible": func(value bool) map[string]interface{} {
			// 选择器为空分支：保持链式调用语义并返回不可命中的空选择器包装。
			if selector == nil {
				return wrapUiacc(nil)
			}
			return wrapUiacc(selector.Visible(value))
		},
		"click": func(text ...string) bool {
			if selector == nil {
				return false
			}
			value := ""
			if len(text) > 0 {
				value = text[0]
			}
			return selector.Click(value)
		},
		"waitFor": func(timeout int) interface{} {
			if selector == nil {
				return wrapUiObject(nil)
			}
			return wrapUiObject(selector.WaitFor(timeout))
		},
		"findOnce": func() interface{} {
			if selector == nil {
				return wrapUiObject(nil)
			}
			return wrapUiObject(selector.FindOnce())
		},
		"find": func() []interface{} {
			if selector == nil {
				return nil
			}
			objects := selector.Find()
			results := make([]interface{}, 0, len(objects))
			for _, object := range objects {
				results = append(results, wrapUiObject(object))
			}
			return results
		},
		"release": func() {
			if selector != nil {
				selector.Release()
			}
		},
	}
}

func wrapUiObject(object *autogouiacc.UiObject) interface{} {
	if object == nil {
		return nil
	}
	return map[string]interface{}{
		"click": func() bool {
			return object.Click()
		},
		"clickCenter": func() bool {
			return object.ClickCenter()
		},
		"setText": func(text string) bool {
			return object.SetText(text)
		},
		"getText": func() string {
			return object.GetText()
		},
		"getDesc":      object.GetDesc,
		"getId":        object.GetId,
		"getClassName": object.GetClassName,
		"getBounds": func() map[string]int {
			bounds := object.GetBounds()
			return map[string]int{
				"left": bounds.Left, "right": bounds.Right,
				"top": bounds.Top, "bottom": bounds.Bottom,
				"centerX": bounds.CenterX, "centerY": bounds.CenterY,
				"width": bounds.Width, "height": bounds.Height,
			}
		},
		"getParent": func() interface{} {
			return wrapUiObject(object.GetParent())
		},
		"getChild": func(index int) interface{} {
			return wrapUiObject(object.GetChild(index))
		},
		"getChildren": func() []interface{} {
			children := object.GetChildren()
			results := make([]interface{}, 0, len(children))
			for _, child := range children {
				results = append(results, wrapUiObject(child))
			}
			return results
		},
		"getChildCount":  object.GetChildCount,
		"getClickable":   object.GetClickable,
		"getSelected":    object.GetSelected,
		"getScrollable":  object.GetScrollable,
		"scrollForward":  object.ScrollForward,
		"scrollBackward": object.ScrollBackward,
	}
}

func optionalInt(defaultValue int, values ...int) int {
	if len(values) == 0 {
		return defaultValue
	}
	return values[0]
}
