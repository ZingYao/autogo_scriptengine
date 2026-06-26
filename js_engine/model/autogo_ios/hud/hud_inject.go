package hud

import (
	"github.com/ZingYao/autogo_scriptengine/js_engine/model"

	autogohud "github.com/Dasongzi1366/AutoGo/hud"
	"github.com/dop251/goja"
)

// HUDModule iOS hud 模块。
type HUDModule struct{}

// Name 返回模块名称。
func (m *HUDModule) Name() string {
	return "hud"
}

// IsAvailable 检查模块是否可用。
func (m *HUDModule) IsAvailable() bool {
	return true
}

// hudChainValue 返回链式调用对象。
func hudChainValue(vm *goja.Runtime, current *goja.Object, original, next *autogohud.HUD) goja.Value {
	if next != nil && next != original {
		return wrapHUD(vm, next)
	}
	return current
}

// jsValueToTextItems 将 JavaScript 数组转换为 HUD 文本项。
func jsValueToTextItems(value goja.Value) []autogohud.TextItem {
	exported := value.Export()
	items := make([]autogohud.TextItem, 0)
	array, ok := exported.([]interface{})
	if !ok {
		return items
	}
	for _, item := range array {
		itemMap, ok := item.(map[string]interface{})
		if !ok {
			continue
		}
		textItem := autogohud.TextItem{}
		if textColor, ok := itemMap["textColor"].(string); ok {
			textItem.TextColor = textColor
		}
		if textColor, ok := itemMap["TextColor"].(string); ok {
			textItem.TextColor = textColor
		}
		if text, ok := itemMap["text"].(string); ok {
			textItem.Text = text
		}
		if text, ok := itemMap["Text"].(string); ok {
			textItem.Text = text
		}
		items = append(items, textItem)
	}
	return items
}

// wrapHUD 将 Go HUD 对象包装为 JavaScript 实例对象。
func wrapHUD(vm *goja.Runtime, h *autogohud.HUD) goja.Value {
	obj := vm.NewObject()
	obj.Set("setText", func(call goja.FunctionCall) goja.Value {
		next := h.SetText(jsValueToTextItems(call.Argument(0)))
		return hudChainValue(vm, obj, h, next)
	})
	obj.Set("setPosition", func(call goja.FunctionCall) goja.Value {
		next := h.SetPosition(
			int(call.Argument(0).ToInteger()),
			int(call.Argument(1).ToInteger()),
			int(call.Argument(2).ToInteger()),
			int(call.Argument(3).ToInteger()),
		)
		return hudChainValue(vm, obj, h, next)
	})
	obj.Set("setBackgroundColor", func(call goja.FunctionCall) goja.Value {
		next := h.SetBackgroundColor(call.Argument(0).String())
		return hudChainValue(vm, obj, h, next)
	})
	obj.Set("setTextSize", func(call goja.FunctionCall) goja.Value {
		next := h.SetTextSize(int(call.Argument(0).ToInteger()))
		return hudChainValue(vm, obj, h, next)
	})
	obj.Set("show", func(call goja.FunctionCall) goja.Value {
		h.Show()
		return goja.Undefined()
	})
	obj.Set("hide", func(call goja.FunctionCall) goja.Value {
		h.Hide()
		return goja.Undefined()
	})
	obj.Set("isVisible", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(h.IsVisible())
	})
	obj.Set("destroy", func(call goja.FunctionCall) goja.Value {
		h.Destroy()
		return goja.Undefined()
	})
	return obj
}

// Register 向引擎注册 iOS hud 方法。
func (m *HUDModule) Register(engine model.Engine) error {
	vm := engine.GetVM()
	hudObj := vm.NewObject()
	vm.Set("hud", hudObj)
	hudObj.Set("new", func(call goja.FunctionCall) goja.Value {
		h := autogohud.New()
		if h == nil {
			return goja.Null()
		}
		return wrapHUD(vm, h)
	})

	engine.RegisterMethod("hud.new", "创建 HUD 对象", autogohud.New, true)
	engine.RegisterMethod("hud.setText", "设置 HUD 文本内容", func(h *autogohud.HUD, items []autogohud.TextItem) *autogohud.HUD {
		return h.SetText(items)
	}, true)
	engine.RegisterMethod("hud.setPosition", "设置 HUD 显示位置", func(h *autogohud.HUD, x1, y1, x2, y2 int) *autogohud.HUD {
		return h.SetPosition(x1, y1, x2, y2)
	}, true)
	engine.RegisterMethod("hud.setBackgroundColor", "设置 HUD 背景颜色", func(h *autogohud.HUD, color string) *autogohud.HUD {
		return h.SetBackgroundColor(color)
	}, true)
	engine.RegisterMethod("hud.setTextSize", "设置 HUD 文本大小", func(h *autogohud.HUD, size int) *autogohud.HUD {
		return h.SetTextSize(size)
	}, true)
	engine.RegisterMethod("hud.show", "显示 HUD", func(h *autogohud.HUD) { h.Show() }, true)
	engine.RegisterMethod("hud.hide", "隐藏 HUD", func(h *autogohud.HUD) { h.Hide() }, true)
	engine.RegisterMethod("hud.isVisible", "返回 HUD 是否可见", func(h *autogohud.HUD) bool { return h.IsVisible() }, true)
	engine.RegisterMethod("hud.destroy", "销毁 HUD 对象", func(h *autogohud.HUD) { h.Destroy() }, true)
	return nil
}
