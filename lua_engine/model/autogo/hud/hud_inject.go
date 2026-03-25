package hud

import (
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"

	"github.com/Dasongzi1366/AutoGo/hud"
	lua "github.com/yuin/gopher-lua"
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
	state := engine.GetState()

	hudObj := state.NewTable()
	state.SetGlobal("hud", hudObj)

	hudObj.RawSetString("new", state.NewFunction(func(L *lua.LState) int {
		h := hud.New()
		ud := L.NewUserData()
		ud.Value = h
		L.Push(ud)
		return 1
	}))

	hudObj.RawSetString("setPosition", state.NewFunction(func(L *lua.LState) int {
		h := L.CheckUserData(1).Value.(*hud.HUD)
		x1 := L.CheckInt(2)
		y1 := L.CheckInt(3)
		x2 := L.CheckInt(4)
		y2 := L.CheckInt(5)
		result := h.SetPosition(x1, y1, x2, y2)
		ud := L.NewUserData()
		ud.Value = result
		L.Push(ud)
		return 1
	}))

	hudObj.RawSetString("setBackgroundColor", state.NewFunction(func(L *lua.LState) int {
		h := L.CheckUserData(1).Value.(*hud.HUD)
		color := L.CheckString(2)
		result := h.SetBackgroundColor(color)
		ud := L.NewUserData()
		ud.Value = result
		L.Push(ud)
		return 1
	}))

	hudObj.RawSetString("setTextSize", state.NewFunction(func(L *lua.LState) int {
		h := L.CheckUserData(1).Value.(*hud.HUD)
		size := L.CheckInt(2)
		result := h.SetTextSize(size)
		ud := L.NewUserData()
		ud.Value = result
		L.Push(ud)
		return 1
	}))

	hudObj.RawSetString("setText", state.NewFunction(func(L *lua.LState) int {
		h := L.CheckUserData(1).Value.(*hud.HUD)
		itemsTable := L.CheckTable(2)
		items := make([]hud.TextItem, 0)
		itemsTable.ForEach(func(key lua.LValue, value lua.LValue) {
			if itemTable, ok := value.(*lua.LTable); ok {
				item := hud.TextItem{}
				if text := itemTable.RawGetString("text"); text.Type() == lua.LTString {
					item.Text = string(text.(lua.LString))
				}
				if textColor := itemTable.RawGetString("textColor"); textColor.Type() == lua.LTString {
					item.TextColor = string(textColor.(lua.LString))
				}
				items = append(items, item)
			}
		})
		result := h.SetText(items)
		ud := L.NewUserData()
		ud.Value = result
		L.Push(ud)
		return 1
	}))

	hudObj.RawSetString("show", state.NewFunction(func(L *lua.LState) int {
		h := L.CheckUserData(1).Value.(*hud.HUD)
		h.Show()
		return 0
	}))

	hudObj.RawSetString("hide", state.NewFunction(func(L *lua.LState) int {
		h := L.CheckUserData(1).Value.(*hud.HUD)
		h.Hide()
		return 0
	}))

	hudObj.RawSetString("isVisible", state.NewFunction(func(L *lua.LState) int {
		h := L.CheckUserData(1).Value.(*hud.HUD)
		visible := h.IsVisible()
		L.Push(lua.LBool(visible))
		return 1
	}))

	hudObj.RawSetString("destroy", state.NewFunction(func(L *lua.LState) int {
		h := L.CheckUserData(1).Value.(*hud.HUD)
		h.Destroy()
		return 0
	}))

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
