package hud

import (
	"github.com/Dasongzi1366/AutoGo/hud"
	"github.com/ZingYao/autogo_scriptengine/js_engine/model"
	"github.com/dop251/goja"
)

// HUDModule hud 模块
type HUDModule struct{}

// Name 返回模块名称
func (m *HUDModule) Name() string {
	return "hud"
}

// IsAvailable 检查模块是否可用
func (m *HUDModule) IsAvailable() bool {
	return true
}

// Register 向引擎注册方法
func (m *HUDModule) Register(engine model.Engine) error {
	vm := engine.GetVM()

	hudObj := vm.NewObject()
	vm.Set("hud", hudObj)

	hudObj.Set("new", func(call goja.FunctionCall) goja.Value {
		h := hud.New()
		if h == nil {
			return goja.Null()
		}
		return wrapHUD(vm, h)
	})

	engine.RegisterMethod("hud.new", "创建一个新的HUD对象", hud.New, true)
	engine.RegisterMethod("hud.setPosition", "设置HUD的显示位置", func(h *hud.HUD, x1, y1, x2, y2 int) *hud.HUD {
		return h.SetPosition(x1, y1, x2, y2)
	}, true)
	engine.RegisterMethod("hud.setBackgroundColor", "设置HUD的背景颜色", func(h *hud.HUD, color string) *hud.HUD {
		return h.SetBackgroundColor(color)
	}, true)
	engine.RegisterMethod("hud.setTextSize", "设置HUD中文本的字体大小", func(h *hud.HUD, size int) *hud.HUD {
		return h.SetTextSize(size)
	}, true)
	engine.RegisterMethod("hud.setText", "设置HUD中要显示的文本内容", func(h *hud.HUD, items []hud.TextItem) *hud.HUD {
		return h.SetText(items)
	}, true)
	engine.RegisterMethod("hud.show", "显示HUD", func(h *hud.HUD) {
		h.Show()
	}, true)
	engine.RegisterMethod("hud.hide", "隐藏HUD", func(h *hud.HUD) {
		h.Hide()
	}, true)
	engine.RegisterMethod("hud.isVisible", "返回HUD是否可见", func(h *hud.HUD) bool {
		return h.IsVisible()
	}, true)
	engine.RegisterMethod("hud.destroy", "销毁HUD对象，释放资源", func(h *hud.HUD) {
		h.Destroy()
	}, true)

	return nil
}

func wrapHUD(vm *goja.Runtime, h *hud.HUD) goja.Value {
	obj := vm.NewObject()
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
	obj.Set("setText", func(call goja.FunctionCall) goja.Value {
		next := h.SetText(jsValueToTextItems(call.Argument(0)))
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
	return vm.ToValue(obj)
}

func hudChainValue(vm *goja.Runtime, current *goja.Object, original, next *hud.HUD) goja.Value {
	if next != nil && next != original {
		return wrapHUD(vm, next)
	}
	return current
}

func jsValueToTextItems(value goja.Value) []hud.TextItem {
	exported := value.Export()
	items := make([]hud.TextItem, 0)
	arr, ok := exported.([]interface{})
	if !ok {
		return items
	}
	for _, item := range arr {
		itemMap, ok := item.(map[string]interface{})
		if !ok {
			continue
		}
		textItem := hud.TextItem{}
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
