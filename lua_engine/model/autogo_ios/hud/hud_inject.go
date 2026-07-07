//go:build ignore
// +build ignore

package hud

import (
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"

	autogohud "github.com/Dasongzi1366/AutoGo/hud"
	lua "github.com/yuin/gopher-lua"
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
func hudChainValue(L *lua.LState, current lua.LValue, original, next *autogohud.HUD) lua.LValue {
	if next != nil && next != original {
		return wrapHUD(L, next)
	}
	return current
}

// luaTableToTextItems 将 Lua 数组表转换为 HUD 文本项。
func luaTableToTextItems(itemsTable *lua.LTable) []autogohud.TextItem {
	items := make([]autogohud.TextItem, 0)
	itemsTable.ForEach(func(key lua.LValue, value lua.LValue) {
		itemTable, ok := value.(*lua.LTable)
		if !ok {
			return
		}
		item := autogohud.TextItem{}
		if textColor := itemTable.RawGetString("textColor"); textColor.Type() == lua.LTString {
			item.TextColor = string(textColor.(lua.LString))
		}
		if textColor := itemTable.RawGetString("TextColor"); textColor.Type() == lua.LTString {
			item.TextColor = string(textColor.(lua.LString))
		}
		if text := itemTable.RawGetString("text"); text.Type() == lua.LTString {
			item.Text = string(text.(lua.LString))
		}
		if text := itemTable.RawGetString("Text"); text.Type() == lua.LTString {
			item.Text = string(text.(lua.LString))
		}
		items = append(items, item)
	})
	return items
}

// wrapHUD 将 Go HUD 对象包装为 Lua 实例对象。
func wrapHUD(L *lua.LState, h *autogohud.HUD) lua.LValue {
	obj := L.NewTable()
	obj.RawSetString("setText", L.NewFunction(func(L *lua.LState) int {
		L.Push(hudChainValue(L, obj, h, h.SetText(luaTableToTextItems(L.CheckTable(1)))))
		return 1
	}))
	obj.RawSetString("setPosition", L.NewFunction(func(L *lua.LState) int {
		L.Push(hudChainValue(L, obj, h, h.SetPosition(L.CheckInt(1), L.CheckInt(2), L.CheckInt(3), L.CheckInt(4))))
		return 1
	}))
	obj.RawSetString("setBackgroundColor", L.NewFunction(func(L *lua.LState) int {
		L.Push(hudChainValue(L, obj, h, h.SetBackgroundColor(L.CheckString(1))))
		return 1
	}))
	obj.RawSetString("setTextSize", L.NewFunction(func(L *lua.LState) int {
		L.Push(hudChainValue(L, obj, h, h.SetTextSize(L.CheckInt(1))))
		return 1
	}))
	obj.RawSetString("show", L.NewFunction(func(L *lua.LState) int {
		h.Show()
		return 0
	}))
	obj.RawSetString("hide", L.NewFunction(func(L *lua.LState) int {
		h.Hide()
		return 0
	}))
	obj.RawSetString("isVisible", L.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LBool(h.IsVisible()))
		return 1
	}))
	obj.RawSetString("destroy", L.NewFunction(func(L *lua.LState) int {
		h.Destroy()
		return 0
	}))
	return obj
}

// Register 向引擎注册 iOS hud 方法。
func (m *HUDModule) Register(engine model.Engine) error {
	state := engine.GetState()
	hudObj := state.NewTable()
	state.SetGlobal("hud", hudObj)
	hudObj.RawSetString("new", state.NewFunction(func(L *lua.LState) int {
		h := autogohud.New()
		if h == nil {
			L.Push(lua.LNil)
			return 1
		}
		L.Push(wrapHUD(L, h))
		return 1
	}))

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
