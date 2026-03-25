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
		return vm.ToValue(h)
	})

	hudObj.Set("setPosition", func(call goja.FunctionCall) goja.Value {
		h := call.Argument(0).Export().(*hud.HUD)
		x1 := int(call.Argument(1).ToInteger())
		y1 := int(call.Argument(2).ToInteger())
		x2 := int(call.Argument(3).ToInteger())
		y2 := int(call.Argument(4).ToInteger())
		h.SetPosition(x1, y1, x2, y2)
		return vm.ToValue(h)
	})

	hudObj.Set("setBackgroundColor", func(call goja.FunctionCall) goja.Value {
		h := call.Argument(0).Export().(*hud.HUD)
		color := call.Argument(1).String()
		h.SetBackgroundColor(color)
		return vm.ToValue(h)
	})

	hudObj.Set("setTextSize", func(call goja.FunctionCall) goja.Value {
		h := call.Argument(0).Export().(*hud.HUD)
		size := int(call.Argument(1).ToInteger())
		h.SetTextSize(size)
		return vm.ToValue(h)
	})

	hudObj.Set("setText", func(call goja.FunctionCall) goja.Value {
		h := call.Argument(0).Export().(*hud.HUD)
		itemsVal := call.Argument(1).Export()
		var items []hud.TextItem
		if arr, ok := itemsVal.([]interface{}); ok {
			for _, item := range arr {
				if m, ok := item.(map[string]interface{}); ok {
					ti := hud.TextItem{}
					if textColor, ok := m["TextColor"].(string); ok {
						ti.TextColor = textColor
					}
					if text, ok := m["Text"].(string); ok {
						ti.Text = text
					}
					items = append(items, ti)
				}
			}
		}
		h.SetText(items)
		return vm.ToValue(h)
	})

	hudObj.Set("show", func(call goja.FunctionCall) goja.Value {
		h := call.Argument(0).Export().(*hud.HUD)
		h.Show()
		return goja.Undefined()
	})

	hudObj.Set("hide", func(call goja.FunctionCall) goja.Value {
		h := call.Argument(0).Export().(*hud.HUD)
		h.Hide()
		return goja.Undefined()
	})

	hudObj.Set("isVisible", func(call goja.FunctionCall) goja.Value {
		h := call.Argument(0).Export().(*hud.HUD)
		result := h.IsVisible()
		return vm.ToValue(result)
	})

	hudObj.Set("destroy", func(call goja.FunctionCall) goja.Value {
		h := call.Argument(0).Export().(*hud.HUD)
		h.Destroy()
		return goja.Undefined()
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
