package imgui

import (
	"runtime"
	"strconv"
	"unsafe"

	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"

	"github.com/Dasongzi1366/AutoGo/imgui"
	"github.com/Dasongzi1366/AutoGo/utils"
	lua "github.com/yuin/gopher-lua"
)

// ImGuiModule imgui 模块
type ImGuiModule struct{}

// Name 返回模块名称
func (m *ImGuiModule) Name() string {
	return "imgui"
}

// IsAvailable 检查模块是否可用（只支持 Android 平台）
func (m *ImGuiModule) IsAvailable() bool {
	return runtime.GOOS == "android"
}

// Register 向引擎注册方法
func (m *ImGuiModule) Register(engine model.Engine) error {
	state := engine.GetState()

	imguiObj := state.NewTable()
	state.SetGlobal("imgui", imguiObj)

	imguiObj.RawSetString("init", state.NewFunction(func(L *lua.LState) int {
		imgui.Init()
		return 0
	}))

	imguiObj.RawSetString("close", state.NewFunction(func(L *lua.LState) int {
		imgui.Close()
		return 0
	}))

	imguiObj.RawSetString("alert", state.NewFunction(func(L *lua.LState) int {
		title := L.CheckString(1)
		content := L.CheckString(2)
		btn1Text := ""
		if L.GetTop() > 2 {
			btn1Text = L.CheckString(3)
		}
		btn2Text := ""
		if L.GetTop() > 3 {
			btn2Text = L.CheckString(4)
		}
		result := utils.Alert(title, content, btn1Text, btn2Text)
		L.Push(lua.LNumber(result))
		return 1
	}))

	imguiObj.RawSetString("toast", state.NewFunction(func(L *lua.LState) int {
		message := L.CheckString(1)
		x := 0
		y := 0
		duration := -1
		if L.GetTop() > 1 {
			x = L.CheckInt(2)
		}
		if L.GetTop() > 2 {
			y = L.CheckInt(3)
		}
		if L.GetTop() > 3 {
			duration = L.CheckInt(4)
		}
		utils.Toast(message, x, y, duration)
		return 0
	}))

	imguiObj.RawSetString("drawRect", state.NewFunction(func(L *lua.LState) int {
		x1 := L.CheckInt(1)
		y1 := L.CheckInt(2)
		x2 := L.CheckInt(3)
		y2 := L.CheckInt(4)
		colorStr := L.CheckString(5)
		thickness := float32(L.CheckNumber(6))
		color := parseColorString(colorStr)
		imgui.DrawRect(x1, y1, x2, y2, color, thickness)
		return 0
	}))

	imguiObj.RawSetString("run", state.NewFunction(func(L *lua.LState) int {
		if L.CheckFunction(1) != nil {
			callback := L.CheckFunction(1)
			imgui.Run(func() {
				L.Push(callback)
				L.PCall(0, 0, nil)
			})
		}
		return 0
	}))

	imguiObj.RawSetString("vertexBufferLayout", state.NewFunction(func(L *lua.LState) int {
		entrySize, posOffset, uvOffset, colOffset := imgui.VertexBufferLayout()
		L.Push(lua.LNumber(entrySize))
		L.Push(lua.LNumber(posOffset))
		L.Push(lua.LNumber(uvOffset))
		L.Push(lua.LNumber(colOffset))
		return 4
	}))

	imguiObj.RawSetString("indexBufferLayout", state.NewFunction(func(L *lua.LState) int {
		entrySize := imgui.IndexBufferLayout()
		L.Push(lua.LNumber(entrySize))
		return 1
	}))

	imguiObj.RawSetString("newGlyphRange", state.NewFunction(func(L *lua.LState) int {
		gr := imgui.NewGlyphRange()
		ud := L.NewUserData()
		ud.Value = gr
		L.Push(ud)
		return 1
	}))

	imguiObj.RawSetString("newContext", state.NewFunction(func(L *lua.LState) int {
		ctx := imgui.NewEmptyContext()
		ud := L.NewUserData()
		ud.Value = ctx
		L.Push(ud)
		return 1
	}))

	imguiObj.RawSetString("newIO", state.NewFunction(func(L *lua.LState) int {
		io := imgui.NewEmptyIO()
		ud := L.NewUserData()
		ud.Value = io
		L.Push(ud)
		return 1
	}))

	imguiObj.RawSetString("newStyle", state.NewFunction(func(L *lua.LState) int {
		style := imgui.NewEmptyStyle()
		ud := L.NewUserData()
		ud.Value = style
		L.Push(ud)
		return 1
	}))

	imguiObj.RawSetString("newDrawList", state.NewFunction(func(L *lua.LState) int {
		drawList := imgui.NewEmptyDrawList()
		ud := L.NewUserData()
		ud.Value = drawList
		L.Push(ud)
		return 1
	}))

	imguiObj.RawSetString("newFont", state.NewFunction(func(L *lua.LState) int {
		font := imgui.NewEmptyFont()
		ud := L.NewUserData()
		ud.Value = font
		L.Push(ud)
		return 1
	}))

	imguiObj.RawSetString("newFontAtlas", state.NewFunction(func(L *lua.LState) int {
		atlas := imgui.NewEmptyFontAtlas()
		ud := L.NewUserData()
		ud.Value = atlas
		L.Push(ud)
		return 1
	}))

	imguiObj.RawSetString("newFontConfig", state.NewFunction(func(L *lua.LState) int {
		config := imgui.NewEmptyFontConfig()
		ud := L.NewUserData()
		ud.Value = config
		L.Push(ud)
		return 1
	}))

	imguiObj.RawSetString("newDrawData", state.NewFunction(func(L *lua.LState) int {
		drawData := imgui.NewEmptyDrawData()
		ud := L.NewUserData()
		ud.Value = drawData
		L.Push(ud)
		return 1
	}))

	imguiObj.RawSetString("newStorage", state.NewFunction(func(L *lua.LState) int {
		storage := imgui.NewEmptyStorage()
		ud := L.NewUserData()
		ud.Value = storage
		L.Push(ud)
		return 1
	}))

	imguiObj.RawSetString("newPlatformIO", state.NewFunction(func(L *lua.LState) int {
		platformIO := imgui.NewEmptyPlatformIO()
		ud := L.NewUserData()
		ud.Value = platformIO
		L.Push(ud)
		return 1
	}))

	imguiObj.RawSetString("createContext", state.NewFunction(func(L *lua.LState) int {
		ctx := imgui.CreateContext()
		ud := L.NewUserData()
		ud.Value = ctx
		L.Push(ud)
		return 1
	}))

	imguiObj.RawSetString("destroyContext", state.NewFunction(func(L *lua.LState) int {
		imgui.DestroyContext()
		return 0
	}))

	imguiObj.RawSetString("setCurrentContext", state.NewFunction(func(L *lua.LState) int {
		ctx := L.CheckUserData(1).Value.(*imgui.Context)
		imgui.SetCurrentContext(ctx)
		return 0
	}))

	imguiObj.RawSetString("newFrame", state.NewFunction(func(L *lua.LState) int {
		imgui.NewFrame()
		return 0
	}))

	imguiObj.RawSetString("render", state.NewFunction(func(L *lua.LState) int {
		imgui.Render()
		return 0
	}))

	imguiObj.RawSetString("endFrame", state.NewFunction(func(L *lua.LState) int {
		imgui.EndFrame()
		return 0
	}))

	imguiObj.RawSetString("begin", state.NewFunction(func(L *lua.LState) int {
		name := L.CheckString(1)
		opened := imgui.Begin(name)
		L.Push(lua.LBool(opened))
		return 1
	}))

	imguiObj.RawSetString("end", state.NewFunction(func(L *lua.LState) int {
		imgui.End()
		return 0
	}))

	imguiObj.RawSetString("button", state.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		clicked := imgui.Button(label)
		L.Push(lua.LBool(clicked))
		return 1
	}))

	imguiObj.RawSetString("text", state.NewFunction(func(L *lua.LState) int {
		text := L.CheckString(1)
		imgui.Text(text)
		return 0
	}))

	imguiObj.RawSetString("inputText", state.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		buf := L.CheckString(2)
		bufPtr := &buf
		result := imgui.InputTextWithHint(label, "", bufPtr, 0, nil)
		L.Push(lua.LBool(result))
		return 1
	}))

	imguiObj.RawSetString("spacing", state.NewFunction(func(L *lua.LState) int {
		imgui.Spacing()
		return 0
	}))

	imguiObj.RawSetString("sameLine", state.NewFunction(func(L *lua.LState) int {
		imgui.SameLine()
		return 0
	}))

	imguiObj.RawSetString("setNextWindowPos", state.NewFunction(func(L *lua.LState) int {
		x := float32(L.CheckNumber(1))
		y := float32(L.CheckNumber(2))
		pos := imgui.Vec2{X: x, Y: y}
		imgui.SetNextWindowPos(pos)
		return 0
	}))

	imguiObj.RawSetString("setNextWindowSize", state.NewFunction(func(L *lua.LState) int {
		width := float32(L.CheckNumber(1))
		height := float32(L.CheckNumber(2))
		size := imgui.Vec2{X: width, Y: height}
		imgui.SetNextWindowSize(size)
		return 0
	}))

	imguiObj.RawSetString("colorConvertFloat4ToU32", state.NewFunction(func(L *lua.LState) int {
		r := float32(L.CheckNumber(1))
		g := float32(L.CheckNumber(2))
		b := float32(L.CheckNumber(3))
		a := float32(L.CheckNumber(4))
		color := imgui.Vec4{X: r, Y: g, Z: b, W: a}
		result := imgui.ColorConvertFloat4ToU32(color)
		L.Push(lua.LNumber(result))
		return 1
	}))

	imguiObj.RawSetString("colorConvertU32ToFloat4", state.NewFunction(func(L *lua.LState) int {
		color := uint32(L.CheckNumber(1))
		result := imgui.ColorConvertU32ToFloat4(color)
		L.Push(lua.LNumber(result.X))
		L.Push(lua.LNumber(result.Y))
		L.Push(lua.LNumber(result.Z))
		L.Push(lua.LNumber(result.W))
		return 4
	}))

	imguiObj.RawSetString("setCursorPos", state.NewFunction(func(L *lua.LState) int {
		x := float32(L.CheckNumber(1))
		y := float32(L.CheckNumber(2))
		pos := imgui.Vec2{X: x, Y: y}
		imgui.SetCursorPos(pos)
		return 0
	}))

	imguiObj.RawSetString("setWindowFontScale", state.NewFunction(func(L *lua.LState) int {
		scale := float32(L.CheckNumber(1))
		imgui.SetWindowFontScale(scale)
		return 0
	}))

	imguiObj.RawSetString("newIO", state.NewFunction(func(L *lua.LState) int {
		io := imgui.NewIO()
		ud := L.NewUserData()
		ud.Value = io
		L.Push(ud)
		return 1
	}))

	imguiObj.RawSetString("newPlatformIO", state.NewFunction(func(L *lua.LState) int {
		platformIO := imgui.NewPlatformIO()
		ud := L.NewUserData()
		ud.Value = platformIO
		L.Push(ud)
		return 1
	}))

	imguiObj.RawSetString("checkbox", state.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		checked := L.CheckBool(2)
		checkedPtr := &checked
		result := imgui.Checkbox(label, checkedPtr)
		L.Push(lua.LBool(result))
		L.Push(lua.LBool(*checkedPtr))
		return 2
	}))

	imguiObj.RawSetString("sliderFloat", state.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		value := float32(L.CheckNumber(2))
		min := float32(L.CheckNumber(3))
		max := float32(L.CheckNumber(4))
		valuePtr := &value
		result := imgui.SliderFloat(label, valuePtr, min, max)
		L.Push(lua.LBool(result))
		L.Push(lua.LNumber(*valuePtr))
		return 2
	}))

	imguiObj.RawSetString("sliderInt", state.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		value := int32(L.CheckInt(2))
		min := int32(L.CheckInt(3))
		max := int32(L.CheckInt(4))
		valuePtr := &value
		result := imgui.SliderInt(label, valuePtr, min, max)
		L.Push(lua.LBool(result))
		L.Push(lua.LNumber(*valuePtr))
		return 2
	}))

	imguiObj.RawSetString("separator", state.NewFunction(func(L *lua.LState) int {
		imgui.Separator()
		return 0
	}))

	imguiObj.RawSetString("separatorText", state.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		imgui.SeparatorText(label)
		return 0
	}))

	imguiObj.RawSetString("beginMenu", state.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		opened := imgui.BeginMenu(label)
		L.Push(lua.LBool(opened))
		return 1
	}))

	imguiObj.RawSetString("endMenu", state.NewFunction(func(L *lua.LState) int {
		imgui.EndMenu()
		return 0
	}))

	imguiObj.RawSetString("beginMenuBar", state.NewFunction(func(L *lua.LState) int {
		opened := imgui.BeginMenuBar()
		L.Push(lua.LBool(opened))
		return 1
	}))

	imguiObj.RawSetString("endMenuBar", state.NewFunction(func(L *lua.LState) int {
		imgui.EndMenuBar()
		return 0
	}))

	imguiObj.RawSetString("beginPopup", state.NewFunction(func(L *lua.LState) int {
		str_id := L.CheckString(1)
		opened := imgui.BeginPopup(str_id)
		L.Push(lua.LBool(opened))
		return 1
	}))

	imguiObj.RawSetString("endPopup", state.NewFunction(func(L *lua.LState) int {
		imgui.EndPopup()
		return 0
	}))

	imguiObj.RawSetString("colorEdit3", state.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		col := [3]float32{
			float32(L.CheckNumber(2)),
			float32(L.CheckNumber(3)),
			float32(L.CheckNumber(4)),
		}
		result := imgui.ColorEdit3(label, &col)
		L.Push(lua.LBool(result))
		L.Push(lua.LNumber(col[0]))
		L.Push(lua.LNumber(col[1]))
		L.Push(lua.LNumber(col[2]))
		return 4
	}))

	imguiObj.RawSetString("colorEdit4", state.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		col := [4]float32{
			float32(L.CheckNumber(2)),
			float32(L.CheckNumber(3)),
			float32(L.CheckNumber(4)),
			float32(L.CheckNumber(5)),
		}
		result := imgui.ColorEdit4(label, &col)
		L.Push(lua.LBool(result))
		L.Push(lua.LNumber(col[0]))
		L.Push(lua.LNumber(col[1]))
		L.Push(lua.LNumber(col[2]))
		L.Push(lua.LNumber(col[3]))
		return 5
	}))

	imguiObj.RawSetString("progressBar", state.NewFunction(func(L *lua.LState) int {
		fraction := float32(L.CheckNumber(1))
		imgui.ProgressBar(fraction)
		return 0
	}))

	imguiObj.RawSetString("bullet", state.NewFunction(func(L *lua.LState) int {
		imgui.Bullet()
		return 0
	}))

	imguiObj.RawSetString("bulletText", state.NewFunction(func(L *lua.LState) int {
		text := L.CheckString(1)
		imgui.BulletText(text)
		return 0
	}))

	imguiObj.RawSetString("smallButton", state.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		clicked := imgui.SmallButton(label)
		L.Push(lua.LBool(clicked))
		return 1
	}))

	imguiObj.RawSetString("arrowButton", state.NewFunction(func(L *lua.LState) int {
		str_id := L.CheckString(1)
		dir := imgui.Dir(L.CheckInt(2))
		clicked := imgui.ArrowButton(str_id, dir)
		L.Push(lua.LBool(clicked))
		return 1
	}))

	imguiObj.RawSetString("beginTooltip", state.NewFunction(func(L *lua.LState) int {
		opened := imgui.BeginTooltip()
		L.Push(lua.LBool(opened))
		return 1
	}))

	imguiObj.RawSetString("endTooltip", state.NewFunction(func(L *lua.LState) int {
		imgui.EndTooltip()
		return 0
	}))

	imguiObj.RawSetString("setTooltip", state.NewFunction(func(L *lua.LState) int {
		text := L.CheckString(1)
		imgui.SetTooltip(text)
		return 0
	}))

	imguiObj.RawSetString("beginGroup", state.NewFunction(func(L *lua.LState) int {
		imgui.BeginGroup()
		return 0
	}))

	imguiObj.RawSetString("endGroup", state.NewFunction(func(L *lua.LState) int {
		imgui.EndGroup()
		return 0
	}))

	imguiObj.RawSetString("popID", state.NewFunction(func(L *lua.LState) int {
		imgui.PopID()
		return 0
	}))

	imguiObj.RawSetString("popStyleVar", state.NewFunction(func(L *lua.LState) int {
		imgui.PopStyleVar()
		return 0
	}))

	imguiObj.RawSetString("popStyleColor", state.NewFunction(func(L *lua.LState) int {
		imgui.PopStyleColor()
		return 0
	}))

	engine.RegisterMethod("imgui.init", "初始化ImGui", imgui.Init, true)
	engine.RegisterMethod("imgui.close", "关闭ImGui", imgui.Close, true)
	engine.RegisterMethod("imgui.run", "运行ImGui主循环", func(callback *lua.LFunction) {
		imgui.Run(func() {
			state.Push(callback)
			state.PCall(0, 0, nil)
		})
	}, true)
	engine.RegisterMethod("imgui.vertexBufferLayout", "获取顶点缓冲区布局", func() (int, int, int, int) {
		return imgui.VertexBufferLayout()
	}, true)
	engine.RegisterMethod("imgui.indexBufferLayout", "获取索引缓冲区布局", func() int {
		return imgui.IndexBufferLayout()
	}, true)
	engine.RegisterMethod("imgui.newGlyphRange", "创建字形范围", func() imgui.GlyphRange {
		return imgui.NewGlyphRange()
	}, true)
	engine.RegisterMethod("imgui.newContext", "创建ImGui上下文", func() *imgui.Context {
		return imgui.NewEmptyContext()
	}, true)
	engine.RegisterMethod("imgui.newIO", "创建IO对象", func() *imgui.IO {
		return imgui.NewEmptyIO()
	}, true)
	engine.RegisterMethod("imgui.newStyle", "创建样式对象", func() *imgui.Style {
		return imgui.NewEmptyStyle()
	}, true)
	engine.RegisterMethod("imgui.newDrawList", "创建绘制列表", func() *imgui.DrawList {
		return imgui.NewEmptyDrawList()
	}, true)
	engine.RegisterMethod("imgui.newFont", "创建字体对象", func() *imgui.Font {
		return imgui.NewEmptyFont()
	}, true)
	engine.RegisterMethod("imgui.newFontAtlas", "创建字体图集", func() *imgui.FontAtlas {
		return imgui.NewEmptyFontAtlas()
	}, true)
	engine.RegisterMethod("imgui.newFontConfig", "创建字体配置", func() *imgui.FontConfig {
		return imgui.NewEmptyFontConfig()
	}, true)
	engine.RegisterMethod("imgui.newDrawData", "创建绘制数据", func() *imgui.DrawData {
		return imgui.NewEmptyDrawData()
	}, true)
	engine.RegisterMethod("imgui.newStorage", "创建存储对象", func() *imgui.Storage {
		return imgui.NewEmptyStorage()
	}, true)
	engine.RegisterMethod("imgui.newPlatformIO", "创建平台IO对象", func() *imgui.PlatformIO {
		return imgui.NewEmptyPlatformIO()
	}, true)
	engine.RegisterMethod("imgui.createContext", "创建ImGui上下文", func() *imgui.Context {
		return imgui.CreateContext()
	}, true)
	engine.RegisterMethod("imgui.destroyContext", "销毁ImGui上下文", func() {
		imgui.DestroyContext()
	}, true)
	engine.RegisterMethod("imgui.setCurrentContext", "设置当前ImGui上下文", func(ctx *imgui.Context) {
		imgui.SetCurrentContext(ctx)
	}, true)
	engine.RegisterMethod("imgui.newFrame", "开始新帧", func() {
		imgui.NewFrame()
	}, true)
	engine.RegisterMethod("imgui.render", "渲染ImGui", func() {
		imgui.Render()
	}, true)
	engine.RegisterMethod("imgui.endFrame", "结束帧", func() {
		imgui.EndFrame()
	}, true)
	engine.RegisterMethod("imgui.begin", "开始窗口", func(name string) bool {
		return imgui.Begin(name)
	}, true)
	engine.RegisterMethod("imgui.end", "结束窗口", func() {
		imgui.End()
	}, true)
	engine.RegisterMethod("imgui.button", "创建按钮", func(label string) bool {
		return imgui.Button(label)
	}, true)
	engine.RegisterMethod("imgui.text", "显示文本", func(text string) {
		imgui.Text(text)
	}, true)
	engine.RegisterMethod("imgui.inputText", "输入文本框", func(label, buf string) bool {
		bufPtr := &buf
		return imgui.InputTextWithHint(label, "", bufPtr, 0, nil)
	}, true)
	engine.RegisterMethod("imgui.spacing", "添加间距", func() {
		imgui.Spacing()
	}, true)
	engine.RegisterMethod("imgui.sameLine", "同下一行", func() {
		imgui.SameLine()
	}, true)
	engine.RegisterMethod("imgui.setNextWindowPos", "设置下一个窗口位置", func(x, y float32) {
		pos := imgui.Vec2{X: x, Y: y}
		imgui.SetNextWindowPos(pos)
	}, true)
	engine.RegisterMethod("imgui.setNextWindowSize", "设置下一个窗口大小", func(width, height float32) {
		size := imgui.Vec2{X: width, Y: height}
		imgui.SetNextWindowSize(size)
	}, true)
	engine.RegisterMethod("imgui.colorConvertFloat4ToU32", "颜色转换", func(r, g, b, a float32) uint32 {
		color := imgui.Vec4{X: r, Y: g, Z: b, W: a}
		return imgui.ColorConvertFloat4ToU32(color)
	}, true)
	engine.RegisterMethod("imgui.colorConvertU32ToFloat4", "颜色转换", func(color uint32) (float32, float32, float32, float32) {
		result := imgui.ColorConvertU32ToFloat4(color)
		return result.X, result.Y, result.Z, result.W
	}, true)
	engine.RegisterMethod("imgui.setCursorPos", "设置光标位置", func(x, y float32) {
		pos := imgui.Vec2{X: x, Y: y}
		imgui.SetCursorPos(pos)
	}, true)
	engine.RegisterMethod("imgui.setWindowFontScale", "设置窗口字体缩放", func(scale float32) {
		imgui.SetWindowFontScale(scale)
	}, true)
	engine.RegisterMethod("imgui.newIO", "创建IO对象", func() *imgui.IO {
		return imgui.NewIO()
	}, true)
	engine.RegisterMethod("imgui.newPlatformIO", "创建平台IO对象", func() *imgui.PlatformIO {
		return imgui.NewPlatformIO()
	}, true)
	engine.RegisterMethod("imgui.checkbox", "复选框", func(label string, checked bool) (bool, bool) {
		checkedPtr := &checked
		result := imgui.Checkbox(label, checkedPtr)
		return result, *checkedPtr
	}, true)
	engine.RegisterMethod("imgui.sliderFloat", "浮点滑块", func(label string, value, min, max float32) (bool, float32) {
		valuePtr := &value
		result := imgui.SliderFloat(label, valuePtr, min, max)
		return result, *valuePtr
	}, true)
	engine.RegisterMethod("imgui.sliderInt", "整数滑块", func(label string, value, min, max int32) (bool, int32) {
		valuePtr := &value
		result := imgui.SliderInt(label, valuePtr, min, max)
		return result, *valuePtr
	}, true)
	engine.RegisterMethod("imgui.separator", "分隔符", func() {
		imgui.Separator()
	}, true)
	engine.RegisterMethod("imgui.separatorText", "分隔文本", func(label string) {
		imgui.SeparatorText(label)
	}, true)
	engine.RegisterMethod("imgui.beginMenu", "开始菜单", func(label string) bool {
		return imgui.BeginMenu(label)
	}, true)
	engine.RegisterMethod("imgui.endMenu", "结束菜单", func() {
		imgui.EndMenu()
	}, true)
	engine.RegisterMethod("imgui.beginMenuBar", "开始菜单栏", func() bool {
		return imgui.BeginMenuBar()
	}, true)
	engine.RegisterMethod("imgui.endMenuBar", "结束菜单栏", func() {
		imgui.EndMenuBar()
	}, true)
	engine.RegisterMethod("imgui.beginPopup", "开始弹出窗口", func(str_id string) bool {
		return imgui.BeginPopup(str_id)
	}, true)
	engine.RegisterMethod("imgui.endPopup", "结束弹出窗口", func() {
		imgui.EndPopup()
	}, true)
	engine.RegisterMethod("imgui.colorEdit3", "颜色编辑器(3通道)", func(label string, r, g, b float32) (bool, float32, float32, float32) {
		col := [3]float32{r, g, b}
		result := imgui.ColorEdit3(label, &col)
		return result, col[0], col[1], col[2]
	}, true)
	engine.RegisterMethod("imgui.colorEdit4", "颜色编辑器(4通道)", func(label string, r, g, b, a float32) (bool, float32, float32, float32, float32) {
		col := [4]float32{r, g, b, a}
		result := imgui.ColorEdit4(label, &col)
		return result, col[0], col[1], col[2], col[3]
	}, true)
	engine.RegisterMethod("imgui.progressBar", "进度条", func(fraction float32) {
		imgui.ProgressBar(fraction)
	}, true)
	engine.RegisterMethod("imgui.bullet", "项目符号", func() {
		imgui.Bullet()
	}, true)
	engine.RegisterMethod("imgui.bulletText", "带项目符号的文本", func(text string) {
		imgui.BulletText(text)
	}, true)
	engine.RegisterMethod("imgui.smallButton", "小按钮", func(label string) bool {
		return imgui.SmallButton(label)
	}, true)
	engine.RegisterMethod("imgui.arrowButton", "箭头按钮", func(str_id string, dir int) bool {
		return imgui.ArrowButton(str_id, imgui.Dir(dir))
	}, true)
	engine.RegisterMethod("imgui.beginTooltip", "开始工具提示", func() bool {
		return imgui.BeginTooltip()
	}, true)
	engine.RegisterMethod("imgui.endTooltip", "结束工具提示", func() {
		imgui.EndTooltip()
	}, true)
	engine.RegisterMethod("imgui.setTooltip", "设置工具提示", func(text string) {
		imgui.SetTooltip(text)
	}, true)
	engine.RegisterMethod("imgui.beginGroup", "开始组", func() {
		imgui.BeginGroup()
	}, true)
	engine.RegisterMethod("imgui.endGroup", "结束组", func() {
		imgui.EndGroup()
	}, true)
	engine.RegisterMethod("imgui.popID", "弹出ID", func() {
		imgui.PopID()
	}, true)
	engine.RegisterMethod("imgui.popStyleVar", "弹出样式变量", func() {
		imgui.PopStyleVar()
	}, true)
	engine.RegisterMethod("imgui.popStyleColor", "弹出样式颜色", func() {
		imgui.PopStyleColor()
	}, true)
	engine.RegisterMethod("imgui.alert", "显示对话框", func(title, content, btn1Text, btn2Text string) int {
		return utils.Alert(title, content, btn1Text, btn2Text)
	}, true)
	engine.RegisterMethod("imgui.toast", "显示Toast提示", func(message string, x, y, duration int) {
		utils.Toast(message, x, y, duration)
	}, true)
	engine.RegisterMethod("imgui.drawRect", "绘制矩形", func(x1, y1, x2, y2 int, colorStr string, thickness float32) {
		color := parseColorString(colorStr)
		imgui.DrawRect(x1, y1, x2, y2, color, thickness)
	}, true)

	imguiObj.RawSetString("beginChild", state.NewFunction(func(L *lua.LState) int {
		str_id := L.CheckString(1)
		size := imgui.Vec2{X: 0, Y: 0}
		if L.GetTop() >= 2 {
			size.X = float32(L.CheckNumber(2))
		}
		if L.GetTop() >= 3 {
			size.Y = float32(L.CheckNumber(3))
		}
		opened := imgui.BeginChildStrV(str_id, size, 0, 0)
		L.Push(lua.LBool(opened))
		return 1
	}))

	imguiObj.RawSetString("endChild", state.NewFunction(func(L *lua.LState) int {
		imgui.EndChild()
		return 0
	}))

	imguiObj.RawSetString("setNextWindowSizeConstraints", state.NewFunction(func(L *lua.LState) int {
		size_min := imgui.Vec2{X: float32(L.CheckNumber(1)), Y: float32(L.CheckNumber(2))}
		size_max := imgui.Vec2{X: float32(L.CheckNumber(3)), Y: float32(L.CheckNumber(4))}
		imgui.SetNextWindowSizeConstraints(size_min, size_max)
		return 0
	}))

	imguiObj.RawSetString("setNextWindowContentSize", state.NewFunction(func(L *lua.LState) int {
		size := imgui.Vec2{X: float32(L.CheckNumber(1)), Y: float32(L.CheckNumber(2))}
		imgui.SetNextWindowContentSize(size)
		return 0
	}))

	imguiObj.RawSetString("setNextWindowCollapsed", state.NewFunction(func(L *lua.LState) int {
		collapsed := L.CheckBool(1)
		imgui.SetNextWindowCollapsed(collapsed)
		return 0
	}))

	imguiObj.RawSetString("setNextWindowFocus", state.NewFunction(func(L *lua.LState) int {
		imgui.SetNextWindowFocus()
		return 0
	}))

	imguiObj.RawSetString("setNextWindowBgAlpha", state.NewFunction(func(L *lua.LState) int {
		alpha := float32(L.CheckNumber(1))
		imgui.SetNextWindowBgAlpha(alpha)
		return 0
	}))

	imguiObj.RawSetString("setWindowPos", state.NewFunction(func(L *lua.LState) int {
		pos := imgui.Vec2{X: float32(L.CheckNumber(1)), Y: float32(L.CheckNumber(2))}
		imgui.SetWindowPosVec2V(pos, 0)
		return 0
	}))

	imguiObj.RawSetString("setWindowSize", state.NewFunction(func(L *lua.LState) int {
		size := imgui.Vec2{X: float32(L.CheckNumber(1)), Y: float32(L.CheckNumber(2))}
		imgui.SetWindowSizeVec2V(size, 0)
		return 0
	}))

	imguiObj.RawSetString("setWindowCollapsed", state.NewFunction(func(L *lua.LState) int {
		collapsed := L.CheckBool(1)
		imgui.SetWindowCollapsedBoolV(collapsed, 0)
		return 0
	}))

	imguiObj.RawSetString("setWindowFocus", state.NewFunction(func(L *lua.LState) int {
		imgui.SetWindowFocus()
		return 0
	}))

	imguiObj.RawSetString("pushFont", state.NewFunction(func(L *lua.LState) int {
		font := L.CheckUserData(1).Value.(*imgui.Font)
		font_size := float32(0)
		if L.GetTop() >= 2 {
			font_size = float32(L.CheckNumber(2))
		}
		imgui.PushFont(font, font_size)
		return 0
	}))

	imguiObj.RawSetString("popFont", state.NewFunction(func(L *lua.LState) int {
		imgui.PopFont()
		return 0
	}))

	imguiObj.RawSetString("pushStyleColor", state.NewFunction(func(L *lua.LState) int {
		idx := imgui.Col(L.CheckInt(1))
		col := imgui.Vec4{
			X: float32(L.CheckNumber(2)),
			Y: float32(L.CheckNumber(3)),
			Z: float32(L.CheckNumber(4)),
			W: float32(L.CheckNumber(5)),
		}
		imgui.PushStyleColorVec4(idx, col)
		return 0
	}))

	imguiObj.RawSetString("pushStyleVar", state.NewFunction(func(L *lua.LState) int {
		idx := imgui.StyleVar(L.CheckInt(1))
		val := float32(L.CheckNumber(2))
		imgui.PushStyleVarFloat(idx, val)
		return 0
	}))

	imguiObj.RawSetString("pushItemWidth", state.NewFunction(func(L *lua.LState) int {
		item_width := float32(L.CheckNumber(1))
		imgui.PushItemWidth(item_width)
		return 0
	}))

	imguiObj.RawSetString("popItemWidth", state.NewFunction(func(L *lua.LState) int {
		imgui.PopItemWidth()
		return 0
	}))

	imguiObj.RawSetString("pushTextWrapPos", state.NewFunction(func(L *lua.LState) int {
		if L.GetTop() >= 1 {
			wrap_local_pos_x := float32(L.CheckNumber(1))
			imgui.PushTextWrapPosV(wrap_local_pos_x)
		} else {
			imgui.PushTextWrapPos()
		}
		return 0
	}))

	imguiObj.RawSetString("popTextWrapPos", state.NewFunction(func(L *lua.LState) int {
		imgui.PopTextWrapPos()
		return 0
	}))

	imguiObj.RawSetString("pushID", state.NewFunction(func(L *lua.LState) int {
		str_id := L.CheckString(1)
		imgui.PushIDStr(str_id)
		return 0
	}))

	imguiObj.RawSetString("newLine", state.NewFunction(func(L *lua.LState) int {
		imgui.NewLine()
		return 0
	}))

	imguiObj.RawSetString("dummy", state.NewFunction(func(L *lua.LState) int {
		size := imgui.Vec2{X: float32(L.CheckNumber(1)), Y: float32(L.CheckNumber(2))}
		imgui.Dummy(size)
		return 0
	}))

	imguiObj.RawSetString("indent", state.NewFunction(func(L *lua.LState) int {
		imgui.Indent()
		return 0
	}))

	imguiObj.RawSetString("unindent", state.NewFunction(func(L *lua.LState) int {
		imgui.Unindent()
		return 0
	}))

	imguiObj.RawSetString("setCursorPos", state.NewFunction(func(L *lua.LState) int {
		pos := imgui.Vec2{X: float32(L.CheckNumber(1)), Y: float32(L.CheckNumber(2))}
		imgui.SetCursorPos(pos)
		return 0
	}))

	imguiObj.RawSetString("setCursorPosX", state.NewFunction(func(L *lua.LState) int {
		local_x := float32(L.CheckNumber(1))
		imgui.SetCursorPosX(local_x)
		return 0
	}))

	imguiObj.RawSetString("setCursorPosY", state.NewFunction(func(L *lua.LState) int {
		local_y := float32(L.CheckNumber(1))
		imgui.SetCursorPosY(local_y)
		return 0
	}))

	imguiObj.RawSetString("setCursorScreenPos", state.NewFunction(func(L *lua.LState) int {
		pos := imgui.Vec2{X: float32(L.CheckNumber(1)), Y: float32(L.CheckNumber(2))}
		imgui.SetCursorScreenPos(pos)
		return 0
	}))

	imguiObj.RawSetString("alignTextToFramePadding", state.NewFunction(func(L *lua.LState) int {
		imgui.AlignTextToFramePadding()
		return 0
	}))

	imguiObj.RawSetString("textColored", state.NewFunction(func(L *lua.LState) int {
		col := imgui.Vec4{
			X: float32(L.CheckNumber(1)),
			Y: float32(L.CheckNumber(2)),
			Z: float32(L.CheckNumber(3)),
			W: float32(L.CheckNumber(4)),
		}
		text := L.CheckString(5)
		imgui.TextColored(col, text)
		return 0
	}))

	imguiObj.RawSetString("textDisabled", state.NewFunction(func(L *lua.LState) int {
		text := L.CheckString(1)
		imgui.TextDisabled(text)
		return 0
	}))

	imguiObj.RawSetString("textWrapped", state.NewFunction(func(L *lua.LState) int {
		text := L.CheckString(1)
		imgui.TextWrapped(text)
		return 0
	}))

	imguiObj.RawSetString("labelText", state.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		text := L.CheckString(2)
		imgui.LabelText(label, text)
		return 0
	}))

	imguiObj.RawSetString("calcTextSize", state.NewFunction(func(L *lua.LState) int {
		text := L.CheckString(1)
		size := imgui.CalcTextSizeV(text, false, 0)
		result := L.NewTable()
		result.RawSetString("width", lua.LNumber(size.X))
		result.RawSetString("height", lua.LNumber(size.Y))
		L.Push(result)
		return 1
	}))

	imguiObj.RawSetString("calcItemWidth", state.NewFunction(func(L *lua.LState) int {
		width := imgui.CalcItemWidth()
		L.Push(lua.LNumber(width))
		return 1
	}))

	imguiObj.RawSetString("setScrollHereX", state.NewFunction(func(L *lua.LState) int {
		center_x_ratio := float32(0.5)
		if L.GetTop() >= 1 {
			center_x_ratio = float32(L.CheckNumber(1))
		}
		imgui.SetScrollHereXV(center_x_ratio)
		return 0
	}))

	imguiObj.RawSetString("setScrollHereY", state.NewFunction(func(L *lua.LState) int {
		center_y_ratio := float32(0.5)
		if L.GetTop() >= 1 {
			center_y_ratio = float32(L.CheckNumber(1))
		}
		imgui.SetScrollHereYV(center_y_ratio)
		return 0
	}))

	imguiObj.RawSetString("setScrollFromPosX", state.NewFunction(func(L *lua.LState) int {
		local_x := float32(L.CheckNumber(1))
		center_x_ratio := float32(0.5)
		if L.GetTop() >= 2 {
			center_x_ratio = float32(L.CheckNumber(2))
		}
		imgui.SetScrollFromPosXFloatV(local_x, center_x_ratio)
		return 0
	}))

	imguiObj.RawSetString("setScrollFromPosY", state.NewFunction(func(L *lua.LState) int {
		local_y := float32(L.CheckNumber(1))
		center_y_ratio := float32(0.5)
		if L.GetTop() >= 2 {
			center_y_ratio = float32(L.CheckNumber(2))
		}
		imgui.SetScrollFromPosYFloatV(local_y, center_y_ratio)
		return 0
	}))

	imguiObj.RawSetString("setScrollX", state.NewFunction(func(L *lua.LState) int {
		scroll_x := float32(L.CheckNumber(1))
		imgui.SetScrollXFloat(scroll_x)
		return 0
	}))

	imguiObj.RawSetString("setScrollY", state.NewFunction(func(L *lua.LState) int {
		scroll_y := float32(L.CheckNumber(1))
		imgui.SetScrollYFloat(scroll_y)
		return 0
	}))

	imguiObj.RawSetString("checkbox", state.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		v := L.CheckBool(2)
		changed := imgui.Checkbox(label, &v)
		L.Push(lua.LBool(changed))
		L.Push(lua.LBool(v))
		return 2
	}))

	imguiObj.RawSetString("checkboxFlags", state.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		flags := L.CheckInt(2)
		flags_value := L.CheckInt(3)
		flagsPtr := int32(flags)
		changed := imgui.CheckboxFlagsIntPtr(label, &flagsPtr, int32(flags_value))
		L.Push(lua.LBool(changed))
		L.Push(lua.LNumber(flagsPtr))
		return 2
	}))

	imguiObj.RawSetString("radioButton", state.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		active := L.CheckBool(2)
		changed := imgui.RadioButtonBool(label, active)
		L.Push(lua.LBool(changed))
		return 1
	}))

	imguiObj.RawSetString("inputFloat", state.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		v := float32(L.CheckNumber(2))
		step := float32(0)
		step_fast := float32(0)
		if L.GetTop() >= 3 {
			step = float32(L.CheckNumber(3))
		}
		if L.GetTop() >= 4 {
			step_fast = float32(L.CheckNumber(4))
		}
		changed := imgui.InputFloatV(label, &v, step, step_fast, "%.3f", 0)
		L.Push(lua.LBool(changed))
		L.Push(lua.LNumber(v))
		return 2
	}))

	imguiObj.RawSetString("inputInt", state.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		v := int32(L.CheckInt(2))
		step := int32(1)
		step_fast := int32(100)
		if L.GetTop() >= 3 {
			step = int32(L.CheckInt(3))
		}
		if L.GetTop() >= 4 {
			step_fast = int32(L.CheckInt(4))
		}
		changed := imgui.InputIntV(label, &v, step, step_fast, 0)
		L.Push(lua.LBool(changed))
		L.Push(lua.LNumber(v))
		return 2
	}))

	imguiObj.RawSetString("inputDouble", state.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		v := float64(L.CheckNumber(2))
		step := 0.0
		step_fast := 0.0
		if L.GetTop() >= 3 {
			step = float64(L.CheckNumber(3))
		}
		if L.GetTop() >= 4 {
			step_fast = float64(L.CheckNumber(4))
		}
		changed := imgui.InputDoubleV(label, &v, step, step_fast, "%.6f", 0)
		L.Push(lua.LBool(changed))
		L.Push(lua.LNumber(v))
		return 2
	}))

	imguiObj.RawSetString("combo", state.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		current_item := int32(L.CheckInt(2))
		items := L.CheckString(3)
		changed := imgui.ComboStrV(label, &current_item, items, -1)
		L.Push(lua.LBool(changed))
		L.Push(lua.LNumber(current_item))
		return 2
	}))

	imguiObj.RawSetString("dragFloat", state.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		v := float32(L.CheckNumber(2))
		v_speed := float32(1.0)
		v_min := float32(0.0)
		v_max := float32(0.0)
		if L.GetTop() >= 3 {
			v_speed = float32(L.CheckNumber(3))
		}
		if L.GetTop() >= 4 {
			v_min = float32(L.CheckNumber(4))
		}
		if L.GetTop() >= 5 {
			v_max = float32(L.CheckNumber(5))
		}
		changed := imgui.DragFloatV(label, &v, v_speed, v_min, v_max, "%.3f", 0)
		L.Push(lua.LBool(changed))
		L.Push(lua.LNumber(v))
		return 2
	}))

	imguiObj.RawSetString("dragInt", state.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		v := int32(L.CheckInt(2))
		v_speed := float32(1.0)
		v_min := int32(0)
		v_max := int32(0)
		if L.GetTop() >= 3 {
			v_speed = float32(L.CheckNumber(3))
		}
		if L.GetTop() >= 4 {
			v_min = int32(L.CheckInt(4))
		}
		if L.GetTop() >= 5 {
			v_max = int32(L.CheckInt(5))
		}
		changed := imgui.DragIntV(label, &v, v_speed, v_min, v_max, "%d", 0)
		L.Push(lua.LBool(changed))
		L.Push(lua.LNumber(v))
		return 2
	}))

	imguiObj.RawSetString("sliderFloat", state.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		v := float32(L.CheckNumber(2))
		v_min := float32(L.CheckNumber(3))
		v_max := float32(L.CheckNumber(4))
		changed := imgui.SliderFloatV(label, &v, v_min, v_max, "%.3f", 0)
		L.Push(lua.LBool(changed))
		L.Push(lua.LNumber(v))
		return 2
	}))

	imguiObj.RawSetString("sliderInt", state.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		v := int32(L.CheckInt(2))
		v_min := int32(L.CheckInt(3))
		v_max := int32(L.CheckInt(4))
		changed := imgui.SliderIntV(label, &v, v_min, v_max, "%d", 0)
		L.Push(lua.LBool(changed))
		L.Push(lua.LNumber(v))
		return 2
	}))

	imguiObj.RawSetString("sliderAngle", state.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		v_rad := float32(L.CheckNumber(2))
		v_degrees_min := float32(-360.0)
		v_degrees_max := float32(360.0)
		if L.GetTop() >= 3 {
			v_degrees_min = float32(L.CheckNumber(3))
		}
		if L.GetTop() >= 4 {
			v_degrees_max = float32(L.CheckNumber(4))
		}
		changed := imgui.SliderAngleV(label, &v_rad, v_degrees_min, v_degrees_max, "%.0f deg", 0)
		L.Push(lua.LBool(changed))
		L.Push(lua.LNumber(v_rad))
		return 2
	}))

	imguiObj.RawSetString("colorEdit3", state.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		col := [3]float32{
			float32(L.CheckNumber(2)),
			float32(L.CheckNumber(3)),
			float32(L.CheckNumber(4)),
		}
		changed := imgui.ColorEdit3V(label, &col, 0)
		L.Push(lua.LBool(changed))
		L.Push(lua.LNumber(col[0]))
		L.Push(lua.LNumber(col[1]))
		L.Push(lua.LNumber(col[2]))
		return 4
	}))

	imguiObj.RawSetString("colorEdit4", state.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		col := [4]float32{
			float32(L.CheckNumber(2)),
			float32(L.CheckNumber(3)),
			float32(L.CheckNumber(4)),
			float32(L.CheckNumber(5)),
		}
		changed := imgui.ColorEdit4V(label, &col, 0)
		L.Push(lua.LBool(changed))
		L.Push(lua.LNumber(col[0]))
		L.Push(lua.LNumber(col[1]))
		L.Push(lua.LNumber(col[2]))
		L.Push(lua.LNumber(col[3]))
		return 5
	}))

	imguiObj.RawSetString("colorPicker3", state.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		col := [3]float32{
			float32(L.CheckNumber(2)),
			float32(L.CheckNumber(3)),
			float32(L.CheckNumber(4)),
		}
		changed := imgui.ColorPicker3V(label, &col, 0)
		L.Push(lua.LBool(changed))
		L.Push(lua.LNumber(col[0]))
		L.Push(lua.LNumber(col[1]))
		L.Push(lua.LNumber(col[2]))
		return 4
	}))

	imguiObj.RawSetString("colorPicker4", state.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		col := [4]float32{
			float32(L.CheckNumber(2)),
			float32(L.CheckNumber(3)),
			float32(L.CheckNumber(4)),
			float32(L.CheckNumber(5)),
		}
		changed := imgui.ColorPicker4V(label, &col, 0, nil)
		L.Push(lua.LBool(changed))
		L.Push(lua.LNumber(col[0]))
		L.Push(lua.LNumber(col[1]))
		L.Push(lua.LNumber(col[2]))
		L.Push(lua.LNumber(col[3]))
		return 5
	}))

	imguiObj.RawSetString("colorButton", state.NewFunction(func(L *lua.LState) int {
		desc_id := L.CheckString(1)
		col := imgui.Vec4{
			X: float32(L.CheckNumber(2)),
			Y: float32(L.CheckNumber(3)),
			Z: float32(L.CheckNumber(4)),
			W: float32(L.CheckNumber(5)),
		}
		changed := imgui.ColorButtonV(desc_id, col, 0, imgui.Vec2{X: 0, Y: 0})
		L.Push(lua.LBool(changed))
		return 1
	}))

	imguiObj.RawSetString("colorConvertHSVtoRGB", state.NewFunction(func(L *lua.LState) int {
		h := float32(L.CheckNumber(1))
		s := float32(L.CheckNumber(2))
		v := float32(L.CheckNumber(3))
		var r, g, b float32
		imgui.ColorConvertHSVtoRGB(h, s, v, &r, &g, &b)
		L.Push(lua.LNumber(r))
		L.Push(lua.LNumber(g))
		L.Push(lua.LNumber(b))
		return 3
	}))

	imguiObj.RawSetString("colorConvertRGBtoHSV", state.NewFunction(func(L *lua.LState) int {
		r := float32(L.CheckNumber(1))
		g := float32(L.CheckNumber(2))
		b := float32(L.CheckNumber(3))
		var h, s, v float32
		imgui.ColorConvertRGBtoHSV(r, g, b, &h, &s, &v)
		L.Push(lua.LNumber(h))
		L.Push(lua.LNumber(s))
		L.Push(lua.LNumber(v))
		return 3
	}))

	imguiObj.RawSetString("treeNode", state.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		opened := imgui.TreeNodeStr(label)
		L.Push(lua.LBool(opened))
		return 1
	}))

	imguiObj.RawSetString("treePop", state.NewFunction(func(L *lua.LState) int {
		imgui.TreePop()
		return 0
	}))

	imguiObj.RawSetString("treePush", state.NewFunction(func(L *lua.LState) int {
		str_id := L.CheckString(1)
		imgui.TreePushStr(str_id)
		return 0
	}))

	imguiObj.RawSetString("collapsingHeader", state.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		opened := imgui.CollapsingHeaderTreeNodeFlagsV(label, 0)
		L.Push(lua.LBool(opened))
		return 1
	}))

	imguiObj.RawSetString("selectable", state.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		selected := false
		if L.GetTop() >= 2 {
			selected = L.CheckBool(2)
		}
		changed := imgui.SelectableBoolV(label, selected, 0, imgui.Vec2{X: 0, Y: 0})
		L.Push(lua.LBool(changed))
		return 1
	}))

	imguiObj.RawSetString("beginTable", state.NewFunction(func(L *lua.LState) int {
		str_id := L.CheckString(1)
		columns := int32(L.CheckInt(2))
		opened := imgui.BeginTableV(str_id, columns, 0, imgui.Vec2{X: 0, Y: 0}, 0)
		L.Push(lua.LBool(opened))
		return 1
	}))

	imguiObj.RawSetString("endTable", state.NewFunction(func(L *lua.LState) int {
		imgui.EndTable()
		return 0
	}))

	imguiObj.RawSetString("tableNextRow", state.NewFunction(func(L *lua.LState) int {
		imgui.TableNextRowV(0, 0)
		return 0
	}))

	imguiObj.RawSetString("tableNextColumn", state.NewFunction(func(L *lua.LState) int {
		opened := imgui.TableNextColumn()
		L.Push(lua.LBool(opened))
		return 1
	}))

	imguiObj.RawSetString("tableSetColumnIndex", state.NewFunction(func(L *lua.LState) int {
		column_n := int32(L.CheckInt(1))
		opened := imgui.TableSetColumnIndex(column_n)
		L.Push(lua.LBool(opened))
		return 1
	}))

	imguiObj.RawSetString("tableSetupColumn", state.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		imgui.TableSetupColumnV(label, 0, 0, 0)
		return 0
	}))

	imguiObj.RawSetString("tableHeadersRow", state.NewFunction(func(L *lua.LState) int {
		imgui.TableHeadersRow()
		return 0
	}))

	imguiObj.RawSetString("tableHeader", state.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		imgui.TableHeader(label)
		return 0
	}))

	imguiObj.RawSetString("tableGetColumnCount", state.NewFunction(func(L *lua.LState) int {
		count := imgui.TableGetColumnCount()
		L.Push(lua.LNumber(count))
		return 1
	}))

	imguiObj.RawSetString("tableGetColumnIndex", state.NewFunction(func(L *lua.LState) int {
		index := imgui.TableGetColumnIndex()
		L.Push(lua.LNumber(index))
		return 1
	}))

	imguiObj.RawSetString("tableGetRowIndex", state.NewFunction(func(L *lua.LState) int {
		index := imgui.TableGetRowIndex()
		L.Push(lua.LNumber(index))
		return 1
	}))

	imguiObj.RawSetString("beginMenu", state.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		opened := imgui.BeginMenuV(label, true)
		L.Push(lua.LBool(opened))
		return 1
	}))

	imguiObj.RawSetString("endMenu", state.NewFunction(func(L *lua.LState) int {
		imgui.EndMenu()
		return 0
	}))

	imguiObj.RawSetString("beginMenuBar", state.NewFunction(func(L *lua.LState) int {
		opened := imgui.BeginMenuBar()
		L.Push(lua.LBool(opened))
		return 1
	}))

	imguiObj.RawSetString("endMenuBar", state.NewFunction(func(L *lua.LState) int {
		imgui.EndMenuBar()
		return 0
	}))

	imguiObj.RawSetString("menuItem", state.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		shortcut := ""
		if L.GetTop() >= 2 {
			shortcut = L.CheckString(2)
		}
		selected := false
		if L.GetTop() >= 3 {
			selected = L.CheckBool(3)
		}
		enabled := true
		if L.GetTop() >= 4 {
			enabled = L.CheckBool(4)
		}
		changed := imgui.MenuItemBoolV(label, shortcut, selected, enabled)
		L.Push(lua.LBool(changed))
		return 1
	}))

	imguiObj.RawSetString("beginTabBar", state.NewFunction(func(L *lua.LState) int {
		str_id := L.CheckString(1)
		opened := imgui.BeginTabBarV(str_id, 0)
		L.Push(lua.LBool(opened))
		return 1
	}))

	imguiObj.RawSetString("endTabBar", state.NewFunction(func(L *lua.LState) int {
		imgui.EndTabBar()
		return 0
	}))

	imguiObj.RawSetString("beginTabItem", state.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		opened := imgui.BeginTabItemV(label, nil, 0)
		L.Push(lua.LBool(opened))
		return 1
	}))

	imguiObj.RawSetString("endTabItem", state.NewFunction(func(L *lua.LState) int {
		imgui.EndTabItem()
		return 0
	}))

	imguiObj.RawSetString("tabItemButton", state.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		opened := imgui.TabItemButtonV(label, 0)
		L.Push(lua.LBool(opened))
		return 1
	}))

	imguiObj.RawSetString("beginDragDropSource", state.NewFunction(func(L *lua.LState) int {
		opened := imgui.BeginDragDropSourceV(0)
		L.Push(lua.LBool(opened))
		return 1
	}))

	imguiObj.RawSetString("endDragDropSource", state.NewFunction(func(L *lua.LState) int {
		imgui.EndDragDropSource()
		return 0
	}))

	imguiObj.RawSetString("beginDragDropTarget", state.NewFunction(func(L *lua.LState) int {
		opened := imgui.BeginDragDropTarget()
		L.Push(lua.LBool(opened))
		return 1
	}))

	imguiObj.RawSetString("endDragDropTarget", state.NewFunction(func(L *lua.LState) int {
		imgui.EndDragDropTarget()
		return 0
	}))

	imguiObj.RawSetString("acceptDragDropPayload", state.NewFunction(func(L *lua.LState) int {
		typeArg := L.CheckString(1)
		payload := imgui.AcceptDragDropPayloadV(typeArg, 0)
		if payload == nil {
			L.Push(lua.LNil)
		} else {
			result := L.NewTable()
			dataPtr := payload.Data()
			dataSize := payload.DataSize()
			if dataSize > 0 && dataPtr != 0 {
				data := make([]byte, dataSize)
				for i := int32(0); i < dataSize; i++ {
					data[i] = *(*byte)(unsafe.Pointer(uintptr(dataPtr) + uintptr(i)))
				}
				result.RawSetString("data", lua.LString(string(data)))
			} else {
				result.RawSetString("data", lua.LString(""))
			}
			result.RawSetString("dataSize", lua.LNumber(dataSize))
			L.Push(result)
		}
		return 1
	}))

	imguiObj.RawSetString("setDragDropPayload", state.NewFunction(func(L *lua.LState) int {
		typeArg := L.CheckString(1)
		data := L.CheckString(2)
		dataPtr := uintptr(0)
		if len(data) > 0 {
			dataPtr = uintptr(unsafe.Pointer(&[]byte(data)[0]))
		}
		changed := imgui.SetDragDropPayloadV(typeArg, dataPtr, uint64(len(data)), 0)
		L.Push(lua.LBool(changed))
		return 1
	}))

	imguiObj.RawSetString("beginDisabled", state.NewFunction(func(L *lua.LState) int {
		disabled := false
		if L.GetTop() >= 1 {
			disabled = L.CheckBool(1)
		}
		imgui.BeginDisabledV(disabled)
		return 0
	}))

	imguiObj.RawSetString("endDisabled", state.NewFunction(func(L *lua.LState) int {
		imgui.EndDisabled()
		return 0
	}))

	imguiObj.RawSetString("showDemoWindow", state.NewFunction(func(L *lua.LState) int {
		imgui.ShowDemoWindow()
		return 0
	}))

	imguiObj.RawSetString("showMetricsWindow", state.NewFunction(func(L *lua.LState) int {
		imgui.ShowMetricsWindow()
		return 0
	}))

	imguiObj.RawSetString("showAboutWindow", state.NewFunction(func(L *lua.LState) int {
		imgui.ShowAboutWindow()
		return 0
	}))

	imguiObj.RawSetString("showStyleEditor", state.NewFunction(func(L *lua.LState) int {
		imgui.ShowStyleEditor()
		return 0
	}))

	imguiObj.RawSetString("showStyleSelector", state.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		imgui.ShowStyleSelector(label)
		return 0
	}))

	imguiObj.RawSetString("showFontSelector", state.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		imgui.ShowFontSelector(label)
		return 0
	}))

	imguiObj.RawSetString("showUserGuide", state.NewFunction(func(L *lua.LState) int {
		imgui.ShowUserGuide()
		return 0
	}))

	imguiObj.RawSetString("showDebugLogWindow", state.NewFunction(func(L *lua.LState) int {
		imgui.ShowDebugLogWindow()
		return 0
	}))

	imguiObj.RawSetString("showIDStackToolWindow", state.NewFunction(func(L *lua.LState) int {
		imgui.ShowIDStackToolWindow()
		return 0
	}))

	imguiObj.RawSetString("beginCombo", state.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		preview_value := L.CheckString(2)
		opened := imgui.BeginCombo(label, preview_value)
		L.Push(lua.LBool(opened))
		return 1
	}))

	imguiObj.RawSetString("endCombo", state.NewFunction(func(L *lua.LState) int {
		imgui.EndCombo()
		return 0
	}))

	imguiObj.RawSetString("beginPopupModal", state.NewFunction(func(L *lua.LState) int {
		name := L.CheckString(1)
		opened := imgui.BeginPopupModal(name)
		L.Push(lua.LBool(opened))
		return 1
	}))

	imguiObj.RawSetString("openPopup", state.NewFunction(func(L *lua.LState) int {
		str_id := L.CheckString(1)
		imgui.OpenPopupStr(str_id)
		return 0
	}))

	imguiObj.RawSetString("closeCurrentPopup", state.NewFunction(func(L *lua.LState) int {
		imgui.CloseCurrentPopup()
		return 0
	}))

	imguiObj.RawSetString("beginPopupContextItem", state.NewFunction(func(L *lua.LState) int {
		opened := imgui.BeginPopupContextItem()
		L.Push(lua.LBool(opened))
		return 1
	}))

	imguiObj.RawSetString("beginPopupContextWindow", state.NewFunction(func(L *lua.LState) int {
		opened := imgui.BeginPopupContextWindow()
		L.Push(lua.LBool(opened))
		return 1
	}))

	imguiObj.RawSetString("beginPopupContextVoid", state.NewFunction(func(L *lua.LState) int {
		opened := imgui.BeginPopupContextVoid()
		L.Push(lua.LBool(opened))
		return 1
	}))

	imguiObj.RawSetString("isItemHovered", state.NewFunction(func(L *lua.LState) int {
		hovered := imgui.IsItemHovered()
		L.Push(lua.LBool(hovered))
		return 1
	}))

	imguiObj.RawSetString("isItemActive", state.NewFunction(func(L *lua.LState) int {
		active := imgui.IsItemActive()
		L.Push(lua.LBool(active))
		return 1
	}))

	imguiObj.RawSetString("isItemFocused", state.NewFunction(func(L *lua.LState) int {
		focused := imgui.IsItemFocused()
		L.Push(lua.LBool(focused))
		return 1
	}))

	imguiObj.RawSetString("isItemClicked", state.NewFunction(func(L *lua.LState) int {
		clicked := imgui.IsItemClicked()
		L.Push(lua.LBool(clicked))
		return 1
	}))

	imguiObj.RawSetString("isItemVisible", state.NewFunction(func(L *lua.LState) int {
		visible := imgui.IsItemVisible()
		L.Push(lua.LBool(visible))
		return 1
	}))

	imguiObj.RawSetString("isItemEdited", state.NewFunction(func(L *lua.LState) int {
		edited := imgui.IsItemEdited()
		L.Push(lua.LBool(edited))
		return 1
	}))

	imguiObj.RawSetString("isItemActivated", state.NewFunction(func(L *lua.LState) int {
		activated := imgui.IsItemActivated()
		L.Push(lua.LBool(activated))
		return 1
	}))

	imguiObj.RawSetString("isItemDeactivated", state.NewFunction(func(L *lua.LState) int {
		deactivated := imgui.IsItemDeactivated()
		L.Push(lua.LBool(deactivated))
		return 1
	}))

	imguiObj.RawSetString("isItemDeactivatedAfterEdit", state.NewFunction(func(L *lua.LState) int {
		deactivated := imgui.IsItemDeactivatedAfterEdit()
		L.Push(lua.LBool(deactivated))
		return 1
	}))

	imguiObj.RawSetString("isItemToggledOpen", state.NewFunction(func(L *lua.LState) int {
		toggled := imgui.IsItemToggledOpen()
		L.Push(lua.LBool(toggled))
		return 1
	}))

	imguiObj.RawSetString("isMouseDragging", state.NewFunction(func(L *lua.LState) int {
		dragging := imgui.IsMouseDraggingV(0, -1)
		L.Push(lua.LBool(dragging))
		return 1
	}))

	imguiObj.RawSetString("isMouseHoveringRect", state.NewFunction(func(L *lua.LState) int {
		minX := float32(L.CheckNumber(1))
		minY := float32(L.CheckNumber(2))
		maxX := float32(L.CheckNumber(3))
		maxY := float32(L.CheckNumber(4))
		r_min := imgui.Vec2{X: minX, Y: minY}
		r_max := imgui.Vec2{X: maxX, Y: maxY}
		hovering := imgui.IsMouseHoveringRect(r_min, r_max)
		L.Push(lua.LBool(hovering))
		return 1
	}))

	imguiObj.RawSetString("isMousePosValid", state.NewFunction(func(L *lua.LState) int {
		valid := imgui.IsMousePosValid()
		L.Push(lua.LBool(valid))
		return 1
	}))

	imguiObj.RawSetString("resetMouseDragDelta", state.NewFunction(func(L *lua.LState) int {
		imgui.ResetMouseDragDelta()
		return 0
	}))

	imguiObj.RawSetString("setKeyboardFocusHere", state.NewFunction(func(L *lua.LState) int {
		imgui.SetKeyboardFocusHere()
		return 0
	}))

	imguiObj.RawSetString("setItemDefaultFocus", state.NewFunction(func(L *lua.LState) int {
		imgui.SetItemDefaultFocus()
		return 0
	}))

	imguiObj.RawSetString("isAnyItemHovered", state.NewFunction(func(L *lua.LState) int {
		hovered := imgui.IsAnyItemHovered()
		L.Push(lua.LBool(hovered))
		return 1
	}))

	imguiObj.RawSetString("isAnyItemActive", state.NewFunction(func(L *lua.LState) int {
		active := imgui.IsAnyItemActive()
		L.Push(lua.LBool(active))
		return 1
	}))

	imguiObj.RawSetString("isAnyItemFocused", state.NewFunction(func(L *lua.LState) int {
		focused := imgui.IsAnyItemFocused()
		L.Push(lua.LBool(focused))
		return 1
	}))

	imguiObj.RawSetString("isWindowHovered", state.NewFunction(func(L *lua.LState) int {
		hovered := imgui.IsWindowHovered()
		L.Push(lua.LBool(hovered))
		return 1
	}))

	imguiObj.RawSetString("isWindowFocused", state.NewFunction(func(L *lua.LState) int {
		focused := imgui.IsWindowFocused()
		L.Push(lua.LBool(focused))
		return 1
	}))

	imguiObj.RawSetString("isWindowCollapsed", state.NewFunction(func(L *lua.LState) int {
		collapsed := imgui.IsWindowCollapsed()
		L.Push(lua.LBool(collapsed))
		return 1
	}))

	imguiObj.RawSetString("isRectVisible", state.NewFunction(func(L *lua.LState) int {
		w := float32(L.CheckNumber(1))
		h := float32(L.CheckNumber(2))
		size := imgui.Vec2{X: w, Y: h}
		visible := imgui.IsRectVisible(size)
		L.Push(lua.LBool(visible))
		return 1
	}))

	imguiObj.RawSetString("listBox", state.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		currentItem := L.CheckInt(2)
		items := L.CheckTable(3)

		itemsList := make([]string, 0, items.Len())
		items.ForEach(func(key, value lua.LValue) {
			if str, ok := value.(lua.LString); ok {
				itemsList = append(itemsList, string(str))
			}
		})

		currentItemPtr := int32(currentItem)
		clicked := imgui.ListBoxStrarr(label, &currentItemPtr, itemsList, int32(len(itemsList)))

		L.Push(lua.LBool(clicked))
		L.Push(lua.LNumber(currentItemPtr))
		return 2
	}))

	imguiObj.RawSetString("beginMainMenuBar", state.NewFunction(func(L *lua.LState) int {
		opened := imgui.BeginMainMenuBar()
		L.Push(lua.LBool(opened))
		return 1
	}))

	imguiObj.RawSetString("endMainMenuBar", state.NewFunction(func(L *lua.LState) int {
		imgui.EndMainMenuBar()
		return 0
	}))

	imguiObj.RawSetString("beginItemTooltip", state.NewFunction(func(L *lua.LState) int {
		opened := imgui.BeginItemTooltip()
		L.Push(lua.LBool(opened))
		return 1
	}))

	imguiObj.RawSetString("setNextItemOpen", state.NewFunction(func(L *lua.LState) int {
		isOpen := L.CheckBool(1)
		imgui.SetNextItemOpen(isOpen)
		return 0
	}))

	imguiObj.RawSetString("setColorEditOptions", state.NewFunction(func(L *lua.LState) int {
		flags := uint32(L.CheckNumber(1))
		imgui.SetColorEditOptions(imgui.ColorEditFlags(flags))
		return 0
	}))

	imguiObj.RawSetString("setTabItemClosed", state.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		imgui.SetTabItemClosed(label)
		return 0
	}))

	imguiObj.RawSetString("inputFloat2", state.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		v := L.CheckTable(2)
		var vArr [2]float32
		for i := 0; i < 2 && i < v.Len(); i++ {
			vArr[i] = float32(lua.LVAsNumber(v.RawGetInt(i + 1)))
		}
		changed := imgui.InputFloat2(label, &vArr)
		L.Push(lua.LBool(changed))
		L.Push(lua.LNumber(vArr[0]))
		L.Push(lua.LNumber(vArr[1]))
		return 3
	}))

	imguiObj.RawSetString("inputFloat3", state.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		v := L.CheckTable(2)
		var vArr [3]float32
		for i := 0; i < 3 && i < v.Len(); i++ {
			vArr[i] = float32(lua.LVAsNumber(v.RawGetInt(i + 1)))
		}
		changed := imgui.InputFloat3(label, &vArr)
		L.Push(lua.LBool(changed))
		L.Push(lua.LNumber(vArr[0]))
		L.Push(lua.LNumber(vArr[1]))
		L.Push(lua.LNumber(vArr[2]))
		return 4
	}))

	imguiObj.RawSetString("inputFloat4", state.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		v := L.CheckTable(2)
		var vArr [4]float32
		for i := 0; i < 4 && i < v.Len(); i++ {
			vArr[i] = float32(lua.LVAsNumber(v.RawGetInt(i + 1)))
		}
		changed := imgui.InputFloat4(label, &vArr)
		L.Push(lua.LBool(changed))
		L.Push(lua.LNumber(vArr[0]))
		L.Push(lua.LNumber(vArr[1]))
		L.Push(lua.LNumber(vArr[2]))
		L.Push(lua.LNumber(vArr[3]))
		return 5
	}))

	imguiObj.RawSetString("inputInt2", state.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		v := L.CheckTable(2)
		var vArr [2]int32
		for i := 0; i < 2 && i < v.Len(); i++ {
			vArr[i] = int32(lua.LVAsNumber(v.RawGetInt(i + 1)))
		}
		changed := imgui.InputInt2(label, &vArr)
		L.Push(lua.LBool(changed))
		L.Push(lua.LNumber(vArr[0]))
		L.Push(lua.LNumber(vArr[1]))
		return 3
	}))

	imguiObj.RawSetString("inputInt3", state.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		v := L.CheckTable(2)
		var vArr [3]int32
		for i := 0; i < 3 && i < v.Len(); i++ {
			vArr[i] = int32(lua.LVAsNumber(v.RawGetInt(i + 1)))
		}
		changed := imgui.InputInt3(label, &vArr)
		L.Push(lua.LBool(changed))
		L.Push(lua.LNumber(vArr[0]))
		L.Push(lua.LNumber(vArr[1]))
		L.Push(lua.LNumber(vArr[2]))
		return 4
	}))

	imguiObj.RawSetString("inputInt4", state.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		v := L.CheckTable(2)
		var vArr [4]int32
		for i := 0; i < 4 && i < v.Len(); i++ {
			vArr[i] = int32(lua.LVAsNumber(v.RawGetInt(i + 1)))
		}
		changed := imgui.InputInt4(label, &vArr)
		L.Push(lua.LBool(changed))
		L.Push(lua.LNumber(vArr[0]))
		L.Push(lua.LNumber(vArr[1]))
		L.Push(lua.LNumber(vArr[2]))
		L.Push(lua.LNumber(vArr[3]))
		return 5
	}))

	imguiObj.RawSetString("dragFloat2", state.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		v := L.CheckTable(2)
		var vArr [2]float32
		for i := 0; i < 2 && i < v.Len(); i++ {
			vArr[i] = float32(lua.LVAsNumber(v.RawGetInt(i + 1)))
		}
		changed := imgui.DragFloat2(label, &vArr)
		L.Push(lua.LBool(changed))
		L.Push(lua.LNumber(vArr[0]))
		L.Push(lua.LNumber(vArr[1]))
		return 3
	}))

	imguiObj.RawSetString("dragFloat3", state.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		v := L.CheckTable(2)
		var vArr [3]float32
		for i := 0; i < 3 && i < v.Len(); i++ {
			vArr[i] = float32(lua.LVAsNumber(v.RawGetInt(i + 1)))
		}
		changed := imgui.DragFloat3(label, &vArr)
		L.Push(lua.LBool(changed))
		L.Push(lua.LNumber(vArr[0]))
		L.Push(lua.LNumber(vArr[1]))
		L.Push(lua.LNumber(vArr[2]))
		return 4
	}))

	imguiObj.RawSetString("dragFloat4", state.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		v := L.CheckTable(2)
		var vArr [4]float32
		for i := 0; i < 4 && i < v.Len(); i++ {
			vArr[i] = float32(lua.LVAsNumber(v.RawGetInt(i + 1)))
		}
		changed := imgui.DragFloat4(label, &vArr)
		L.Push(lua.LBool(changed))
		L.Push(lua.LNumber(vArr[0]))
		L.Push(lua.LNumber(vArr[1]))
		L.Push(lua.LNumber(vArr[2]))
		L.Push(lua.LNumber(vArr[3]))
		return 5
	}))

	imguiObj.RawSetString("dragInt2", state.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		v := L.CheckTable(2)
		var vArr [2]int32
		for i := 0; i < 2 && i < v.Len(); i++ {
			vArr[i] = int32(lua.LVAsNumber(v.RawGetInt(i + 1)))
		}
		changed := imgui.DragInt2(label, &vArr)
		L.Push(lua.LBool(changed))
		L.Push(lua.LNumber(vArr[0]))
		L.Push(lua.LNumber(vArr[1]))
		return 3
	}))

	imguiObj.RawSetString("dragInt3", state.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		v := L.CheckTable(2)
		var vArr [3]int32
		for i := 0; i < 3 && i < v.Len(); i++ {
			vArr[i] = int32(lua.LVAsNumber(v.RawGetInt(i + 1)))
		}
		changed := imgui.DragInt3(label, &vArr)
		L.Push(lua.LBool(changed))
		L.Push(lua.LNumber(vArr[0]))
		L.Push(lua.LNumber(vArr[1]))
		L.Push(lua.LNumber(vArr[2]))
		return 4
	}))

	imguiObj.RawSetString("dragInt4", state.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		v := L.CheckTable(2)
		var vArr [4]int32
		for i := 0; i < 4 && i < v.Len(); i++ {
			vArr[i] = int32(lua.LVAsNumber(v.RawGetInt(i + 1)))
		}
		changed := imgui.DragInt4(label, &vArr)
		L.Push(lua.LBool(changed))
		L.Push(lua.LNumber(vArr[0]))
		L.Push(lua.LNumber(vArr[1]))
		L.Push(lua.LNumber(vArr[2]))
		L.Push(lua.LNumber(vArr[3]))
		return 5
	}))

	imguiObj.RawSetString("dragFloatRange2", state.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		vMin := float32(L.CheckNumber(2))
		vMax := float32(L.CheckNumber(3))
		changed := imgui.DragFloatRange2(label, &vMin, &vMax)
		L.Push(lua.LBool(changed))
		L.Push(lua.LNumber(vMin))
		L.Push(lua.LNumber(vMax))
		return 3
	}))

	imguiObj.RawSetString("dragIntRange2", state.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		vMin := int32(L.CheckNumber(2))
		vMax := int32(L.CheckNumber(3))
		changed := imgui.DragIntRange2(label, &vMin, &vMax)
		L.Push(lua.LBool(changed))
		L.Push(lua.LNumber(vMin))
		L.Push(lua.LNumber(vMax))
		return 3
	}))

	imguiObj.RawSetString("sliderFloat2", state.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		v := L.CheckTable(2)
		vMin := float32(L.CheckNumber(3))
		vMax := float32(L.CheckNumber(4))
		var vArr [2]float32
		for i := 0; i < 2 && i < v.Len(); i++ {
			vArr[i] = float32(lua.LVAsNumber(v.RawGetInt(i + 1)))
		}
		changed := imgui.SliderFloat2(label, &vArr, vMin, vMax)
		L.Push(lua.LBool(changed))
		L.Push(lua.LNumber(vArr[0]))
		L.Push(lua.LNumber(vArr[1]))
		return 3
	}))

	imguiObj.RawSetString("sliderFloat3", state.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		v := L.CheckTable(2)
		vMin := float32(L.CheckNumber(3))
		vMax := float32(L.CheckNumber(4))
		var vArr [3]float32
		for i := 0; i < 3 && i < v.Len(); i++ {
			vArr[i] = float32(lua.LVAsNumber(v.RawGetInt(i + 1)))
		}
		changed := imgui.SliderFloat3(label, &vArr, vMin, vMax)
		L.Push(lua.LBool(changed))
		L.Push(lua.LNumber(vArr[0]))
		L.Push(lua.LNumber(vArr[1]))
		L.Push(lua.LNumber(vArr[2]))
		return 4
	}))

	imguiObj.RawSetString("sliderFloat4", state.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		v := L.CheckTable(2)
		vMin := float32(L.CheckNumber(3))
		vMax := float32(L.CheckNumber(4))
		var vArr [4]float32
		for i := 0; i < 4 && i < v.Len(); i++ {
			vArr[i] = float32(lua.LVAsNumber(v.RawGetInt(i + 1)))
		}
		changed := imgui.SliderFloat4(label, &vArr, vMin, vMax)
		L.Push(lua.LBool(changed))
		L.Push(lua.LNumber(vArr[0]))
		L.Push(lua.LNumber(vArr[1]))
		L.Push(lua.LNumber(vArr[2]))
		L.Push(lua.LNumber(vArr[3]))
		return 5
	}))

	imguiObj.RawSetString("sliderInt2", state.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		v := L.CheckTable(2)
		vMin := int32(L.CheckNumber(3))
		vMax := int32(L.CheckNumber(4))
		var vArr [2]int32
		for i := 0; i < 2 && i < v.Len(); i++ {
			vArr[i] = int32(lua.LVAsNumber(v.RawGetInt(i + 1)))
		}
		changed := imgui.SliderInt2(label, &vArr, vMin, vMax)
		L.Push(lua.LBool(changed))
		L.Push(lua.LNumber(vArr[0]))
		L.Push(lua.LNumber(vArr[1]))
		return 3
	}))

	imguiObj.RawSetString("sliderInt3", state.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		v := L.CheckTable(2)
		vMin := int32(L.CheckNumber(3))
		vMax := int32(L.CheckNumber(4))
		var vArr [3]int32
		for i := 0; i < 3 && i < v.Len(); i++ {
			vArr[i] = int32(lua.LVAsNumber(v.RawGetInt(i + 1)))
		}
		changed := imgui.SliderInt3(label, &vArr, vMin, vMax)
		L.Push(lua.LBool(changed))
		L.Push(lua.LNumber(vArr[0]))
		L.Push(lua.LNumber(vArr[1]))
		L.Push(lua.LNumber(vArr[2]))
		return 4
	}))

	imguiObj.RawSetString("sliderInt4", state.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		v := L.CheckTable(2)
		vMin := int32(L.CheckNumber(3))
		vMax := int32(L.CheckNumber(4))
		var vArr [4]int32
		for i := 0; i < 4 && i < v.Len(); i++ {
			vArr[i] = int32(lua.LVAsNumber(v.RawGetInt(i + 1)))
		}
		changed := imgui.SliderInt4(label, &vArr, vMin, vMax)
		L.Push(lua.LBool(changed))
		L.Push(lua.LNumber(vArr[0]))
		L.Push(lua.LNumber(vArr[1]))
		L.Push(lua.LNumber(vArr[2]))
		L.Push(lua.LNumber(vArr[3]))
		return 5
	}))

	imguiObj.RawSetString("vSliderFloat", state.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		w := float32(L.CheckNumber(2))
		h := float32(L.CheckNumber(3))
		v := float32(L.CheckNumber(4))
		vMin := float32(L.CheckNumber(5))
		vMax := float32(L.CheckNumber(6))
		size := imgui.Vec2{X: w, Y: h}
		vPtr := v
		changed := imgui.VSliderFloat(label, size, &vPtr, vMin, vMax)
		L.Push(lua.LBool(changed))
		L.Push(lua.LNumber(vPtr))
		return 2
	}))

	imguiObj.RawSetString("vSliderInt", state.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		w := float32(L.CheckNumber(2))
		h := float32(L.CheckNumber(3))
		v := int32(L.CheckNumber(4))
		vMin := int32(L.CheckNumber(5))
		vMax := int32(L.CheckNumber(6))
		size := imgui.Vec2{X: w, Y: h}
		vPtr := v
		changed := imgui.VSliderInt(label, size, &vPtr, vMin, vMax)
		L.Push(lua.LBool(changed))
		L.Push(lua.LNumber(vPtr))
		return 2
	}))

	engine.RegisterMethod("imgui.beginChild", "开始子窗口", func(str_id string, width, height float32) bool {
		size := imgui.Vec2{X: width, Y: height}
		return imgui.BeginChildStrV(str_id, size, 0, 0)
	}, true)
	engine.RegisterMethod("imgui.endChild", "结束子窗口", func() {
		imgui.EndChild()
	}, true)
	engine.RegisterMethod("imgui.setNextWindowSizeConstraints", "设置下一个窗口大小约束", func(min_w, min_h, max_w, max_h float32) {
		size_min := imgui.Vec2{X: min_w, Y: min_h}
		size_max := imgui.Vec2{X: max_w, Y: max_h}
		imgui.SetNextWindowSizeConstraintsV(size_min, size_max, nil, 0)
	}, true)
	engine.RegisterMethod("imgui.setNextWindowContentSize", "设置下一个窗口内容大小", func(w, h float32) {
		size := imgui.Vec2{X: w, Y: h}
		imgui.SetNextWindowContentSize(size)
	}, true)
	engine.RegisterMethod("imgui.setNextWindowCollapsed", "设置下一个窗口折叠状态", func(collapsed bool) {
		imgui.SetNextWindowCollapsedV(collapsed, 0)
	}, true)
	engine.RegisterMethod("imgui.setNextWindowFocus", "设置下一个窗口焦点", func() {
		imgui.SetNextWindowFocus()
	}, true)
	engine.RegisterMethod("imgui.setNextWindowBgAlpha", "设置下一个窗口背景透明度", func(alpha float32) {
		imgui.SetNextWindowBgAlpha(alpha)
	}, true)
	engine.RegisterMethod("imgui.setWindowPos", "设置窗口位置", func(x, y float32) {
		pos := imgui.Vec2{X: x, Y: y}
		imgui.SetWindowPosVec2V(pos, 0)
	}, true)
	engine.RegisterMethod("imgui.setWindowSize", "设置窗口大小", func(w, h float32) {
		size := imgui.Vec2{X: w, Y: h}
		imgui.SetWindowSizeVec2V(size, 0)
	}, true)
	engine.RegisterMethod("imgui.setWindowCollapsed", "设置窗口折叠状态", func(collapsed bool) {
		imgui.SetWindowCollapsedBoolV(collapsed, 0)
	}, true)
	engine.RegisterMethod("imgui.setWindowFocus", "设置窗口焦点", func() {
		imgui.SetWindowFocus()
	}, true)
	engine.RegisterMethod("imgui.pushFont", "推入字体", func(font *imgui.Font, fontSize float32) {
		imgui.PushFont(font, fontSize)
	}, true)
	engine.RegisterMethod("imgui.popFont", "弹出字体", func() {
		imgui.PopFont()
	}, true)
	engine.RegisterMethod("imgui.pushStyleColor", "推入样式颜色", func(idx int, r, g, b, a float32) {
		col := imgui.Vec4{X: r, Y: g, Z: b, W: a}
		imgui.PushStyleColorVec4(imgui.Col(idx), col)
	}, true)
	engine.RegisterMethod("imgui.pushStyleVar", "推入样式变量", func(idx int, val float32) {
		imgui.PushStyleVarFloat(imgui.StyleVar(idx), val)
	}, true)
	engine.RegisterMethod("imgui.pushItemWidth", "推入项宽度", func(itemWidth float32) {
		imgui.PushItemWidth(itemWidth)
	}, true)
	engine.RegisterMethod("imgui.popItemWidth", "弹出项宽度", func() {
		imgui.PopItemWidth()
	}, true)
	engine.RegisterMethod("imgui.pushTextWrapPos", "推入文本换行位置", func(wrapLocalPosX float32) {
		imgui.PushTextWrapPosV(wrapLocalPosX)
	}, true)
	engine.RegisterMethod("imgui.popTextWrapPos", "弹出文本换行位置", func() {
		imgui.PopTextWrapPos()
	}, true)
	engine.RegisterMethod("imgui.pushID", "推入ID", func(strID string) {
		imgui.PushIDStr(strID)
	}, true)
	engine.RegisterMethod("imgui.newLine", "新行", func() {
		imgui.NewLine()
	}, true)
	engine.RegisterMethod("imgui.dummy", "虚拟占位符", func(w, h float32) {
		size := imgui.Vec2{X: w, Y: h}
		imgui.Dummy(size)
	}, true)
	engine.RegisterMethod("imgui.indent", "缩进", func() {
		imgui.Indent()
	}, true)
	engine.RegisterMethod("imgui.unindent", "取消缩进", func() {
		imgui.Unindent()
	}, true)
	engine.RegisterMethod("imgui.setCursorPos", "设置光标位置", func(x, y float32) {
		pos := imgui.Vec2{X: x, Y: y}
		imgui.SetCursorPos(pos)
	}, true)
	engine.RegisterMethod("imgui.setCursorPosX", "设置光标X位置", func(x float32) {
		imgui.SetCursorPosX(x)
	}, true)
	engine.RegisterMethod("imgui.setCursorPosY", "设置光标Y位置", func(y float32) {
		imgui.SetCursorPosY(y)
	}, true)
	engine.RegisterMethod("imgui.setCursorScreenPos", "设置光标屏幕位置", func(x, y float32) {
		pos := imgui.Vec2{X: x, Y: y}
		imgui.SetCursorScreenPos(pos)
	}, true)
	engine.RegisterMethod("imgui.alignTextToFramePadding", "对齐文本到帧内边距", func() {
		imgui.AlignTextToFramePadding()
	}, true)
	engine.RegisterMethod("imgui.textColored", "彩色文本", func(r, g, b, a float32, text string) {
		col := imgui.Vec4{X: r, Y: g, Z: b, W: a}
		imgui.TextColored(col, text)
	}, true)
	engine.RegisterMethod("imgui.textDisabled", "禁用文本", func(text string) {
		imgui.TextDisabled(text)
	}, true)
	engine.RegisterMethod("imgui.textWrapped", "换行文本", func(text string) {
		imgui.TextWrapped(text)
	}, true)
	engine.RegisterMethod("imgui.labelText", "标签文本", func(label, text string) {
		imgui.LabelText(label, text)
	}, true)
	engine.RegisterMethod("imgui.calcTextSize", "计算文本大小", func(text string) (float32, float32) {
		size := imgui.CalcTextSizeV(text, false, 0)
		return size.X, size.Y
	}, true)
	engine.RegisterMethod("imgui.calcItemWidth", "计算项宽度", func() float32 {
		return imgui.CalcItemWidth()
	}, true)
	engine.RegisterMethod("imgui.setScrollHereX", "设置滚动到此处X", func(centerXRatio float32) {
		imgui.SetScrollHereXV(centerXRatio)
	}, true)
	engine.RegisterMethod("imgui.setScrollHereY", "设置滚动到此处Y", func(centerYRatio float32) {
		imgui.SetScrollHereYV(centerYRatio)
	}, true)
	engine.RegisterMethod("imgui.setScrollFromPosX", "从位置X滚动", func(localX, centerXRatio float32) {
		imgui.SetScrollFromPosXFloatV(localX, centerXRatio)
	}, true)
	engine.RegisterMethod("imgui.setScrollFromPosY", "从位置Y滚动", func(localY, centerYRatio float32) {
		imgui.SetScrollFromPosYFloatV(localY, centerYRatio)
	}, true)
	engine.RegisterMethod("imgui.setScrollX", "设置滚动X", func(scrollX float32) {
		imgui.SetScrollXFloat(scrollX)
	}, true)
	engine.RegisterMethod("imgui.setScrollY", "设置滚动Y", func(scrollY float32) {
		imgui.SetScrollYFloat(scrollY)
	}, true)
	engine.RegisterMethod("imgui.checkbox", "复选框", func(label string, v bool) (bool, bool) {
		changed := imgui.Checkbox(label, &v)
		return changed, v
	}, true)
	engine.RegisterMethod("imgui.checkboxFlags", "复选框标志", func(label string, flags, flagsValue int) (bool, int) {
		flagsPtr := int32(flags)
		changed := imgui.CheckboxFlagsIntPtr(label, &flagsPtr, int32(flagsValue))
		return changed, int(flagsPtr)
	}, true)
	engine.RegisterMethod("imgui.radioButton", "单选按钮", func(label string, active bool) bool {
		return imgui.RadioButtonBool(label, active)
	}, true)
	engine.RegisterMethod("imgui.inputFloat", "浮点输入", func(label string, v float32) (bool, float32) {
		changed := imgui.InputFloatV(label, &v, 0, 0, "%.3f", 0)
		return changed, v
	}, true)
	engine.RegisterMethod("imgui.inputInt", "整数输入", func(label string, v int) (bool, int) {
		v32 := int32(v)
		changed := imgui.InputIntV(label, &v32, 1, 100, 0)
		return changed, int(v32)
	}, true)
	engine.RegisterMethod("imgui.inputDouble", "双精度浮点输入", func(label string, v float64) (bool, float64) {
		changed := imgui.InputDoubleV(label, &v, 0, 0, "%.6f", 0)
		return changed, v
	}, true)
	engine.RegisterMethod("imgui.combo", "下拉框", func(label string, currentItem int, items string) (bool, int) {
		currentItem32 := int32(currentItem)
		changed := imgui.ComboStrV(label, &currentItem32, items, -1)
		return changed, int(currentItem32)
	}, true)
	engine.RegisterMethod("imgui.dragFloat", "浮点拖动", func(label string, v float32) (bool, float32) {
		changed := imgui.DragFloatV(label, &v, 1.0, 0, 0, "%.3f", 0)
		return changed, v
	}, true)
	engine.RegisterMethod("imgui.dragInt", "整数拖动", func(label string, v int) (bool, int) {
		v32 := int32(v)
		changed := imgui.DragIntV(label, &v32, 1.0, 0, 0, "%d", 0)
		return changed, int(v32)
	}, true)
	engine.RegisterMethod("imgui.sliderFloat", "浮点滑块", func(label string, v, vMin, vMax float32) (bool, float32) {
		changed := imgui.SliderFloatV(label, &v, vMin, vMax, "%.3f", 0)
		return changed, v
	}, true)
	engine.RegisterMethod("imgui.sliderInt", "整数滑块", func(label string, v, vMin, vMax int) (bool, int) {
		v32 := int32(v)
		changed := imgui.SliderIntV(label, &v32, int32(vMin), int32(vMax), "%d", 0)
		return changed, int(v32)
	}, true)
	engine.RegisterMethod("imgui.sliderAngle", "角度滑块", func(label string, vRad float32) (bool, float32) {
		changed := imgui.SliderAngleV(label, &vRad, -360.0, 360.0, "%.0f deg", 0)
		return changed, vRad
	}, true)
	engine.RegisterMethod("imgui.colorEdit3", "颜色编辑3", func(label string, r, g, b float32) (bool, float32, float32, float32) {
		col := [3]float32{r, g, b}
		changed := imgui.ColorEdit3V(label, &col, 0)
		return changed, col[0], col[1], col[2]
	}, true)
	engine.RegisterMethod("imgui.colorEdit4", "颜色编辑4", func(label string, r, g, b, a float32) (bool, float32, float32, float32, float32) {
		col := [4]float32{r, g, b, a}
		changed := imgui.ColorEdit4V(label, &col, 0)
		return changed, col[0], col[1], col[2], col[3]
	}, true)
	engine.RegisterMethod("imgui.colorPicker3", "颜色选择器3", func(label string, r, g, b float32) (bool, float32, float32, float32) {
		col := [3]float32{r, g, b}
		changed := imgui.ColorPicker3V(label, &col, 0)
		return changed, col[0], col[1], col[2]
	}, true)
	engine.RegisterMethod("imgui.colorPicker4", "颜色选择器4", func(label string, r, g, b, a float32) (bool, float32, float32, float32, float32) {
		col := [4]float32{r, g, b, a}
		changed := imgui.ColorPicker4V(label, &col, 0, nil)
		return changed, col[0], col[1], col[2], col[3]
	}, true)
	engine.RegisterMethod("imgui.colorButton", "颜色按钮", func(descID string, r, g, b, a float32) bool {
		col := imgui.Vec4{X: r, Y: g, Z: b, W: a}
		return imgui.ColorButtonV(descID, col, 0, imgui.Vec2{X: 0, Y: 0})
	}, true)
	engine.RegisterMethod("imgui.colorConvertHSVtoRGB", "HSV转RGB", func(h, s, v float32) (float32, float32, float32) {
		var r, g, b float32
		imgui.ColorConvertHSVtoRGB(h, s, v, &r, &g, &b)
		return r, g, b
	}, true)
	engine.RegisterMethod("imgui.colorConvertRGBtoHSV", "RGB转HSV", func(r, g, b float32) (float32, float32, float32) {
		var h, s, v float32
		imgui.ColorConvertRGBtoHSV(r, g, b, &h, &s, &v)
		return h, s, v
	}, true)
	engine.RegisterMethod("imgui.treeNode", "树节点", func(label string) bool {
		return imgui.TreeNodeStr(label)
	}, true)
	engine.RegisterMethod("imgui.treePop", "弹出树节点", func() {
		imgui.TreePop()
	}, true)
	engine.RegisterMethod("imgui.treePush", "推入树节点", func(strID string) {
		imgui.TreePushStr(strID)
	}, true)
	engine.RegisterMethod("imgui.collapsingHeader", "折叠头", func(label string) bool {
		return imgui.CollapsingHeaderTreeNodeFlagsV(label, 0)
	}, true)
	engine.RegisterMethod("imgui.selectable", "可选项", func(label string, selected bool) bool {
		return imgui.SelectableBoolV(label, selected, 0, imgui.Vec2{X: 0, Y: 0})
	}, true)
	engine.RegisterMethod("imgui.beginTable", "开始表格", func(strID string, columns int) bool {
		return imgui.BeginTableV(strID, int32(columns), 0, imgui.Vec2{X: 0, Y: 0}, 0)
	}, true)
	engine.RegisterMethod("imgui.endTable", "结束表格", func() {
		imgui.EndTable()
	}, true)
	engine.RegisterMethod("imgui.tableNextRow", "表格下一行", func() {
		imgui.TableNextRowV(0, 0)
	}, true)
	engine.RegisterMethod("imgui.tableNextColumn", "表格下一列", func() bool {
		return imgui.TableNextColumn()
	}, true)
	engine.RegisterMethod("imgui.tableSetColumnIndex", "设置表格列索引", func(columnN int) bool {
		return imgui.TableSetColumnIndex(int32(columnN))
	}, true)
	engine.RegisterMethod("imgui.tableSetupColumn", "设置表格列", func(label string) {
		imgui.TableSetupColumnV(label, 0, 0, 0)
	}, true)
	engine.RegisterMethod("imgui.tableHeadersRow", "表格标题行", func() {
		imgui.TableHeadersRow()
	}, true)
	engine.RegisterMethod("imgui.tableHeader", "表格标题", func(label string) {
		imgui.TableHeader(label)
	}, true)
	engine.RegisterMethod("imgui.tableGetColumnCount", "获取表格列数", func() int {
		return int(imgui.TableGetColumnCount())
	}, true)
	engine.RegisterMethod("imgui.tableGetColumnIndex", "获取表格列索引", func() int {
		return int(imgui.TableGetColumnIndex())
	}, true)
	engine.RegisterMethod("imgui.tableGetRowIndex", "获取表格行索引", func() int {
		return int(imgui.TableGetRowIndex())
	}, true)
	engine.RegisterMethod("imgui.beginMenu", "开始菜单", func(label string) bool {
		return imgui.BeginMenuV(label, true)
	}, true)
	engine.RegisterMethod("imgui.endMenu", "结束菜单", func() {
		imgui.EndMenu()
	}, true)
	engine.RegisterMethod("imgui.beginMenuBar", "开始菜单栏", func() bool {
		return imgui.BeginMenuBar()
	}, true)
	engine.RegisterMethod("imgui.endMenuBar", "结束菜单栏", func() {
		imgui.EndMenuBar()
	}, true)
	engine.RegisterMethod("imgui.menuItem", "菜单项", func(label string, shortcut string, selected, enabled bool) bool {
		return imgui.MenuItemBoolV(label, shortcut, selected, enabled)
	}, true)
	engine.RegisterMethod("imgui.beginTabBar", "开始标签栏", func(strID string) bool {
		return imgui.BeginTabBarV(strID, 0)
	}, true)
	engine.RegisterMethod("imgui.endTabBar", "结束标签栏", func() {
		imgui.EndTabBar()
	}, true)
	engine.RegisterMethod("imgui.beginTabItem", "开始标签项", func(label string) bool {
		return imgui.BeginTabItemV(label, nil, 0)
	}, true)
	engine.RegisterMethod("imgui.endTabItem", "结束标签项", func() {
		imgui.EndTabItem()
	}, true)
	engine.RegisterMethod("imgui.tabItemButton", "标签项按钮", func(label string) bool {
		return imgui.TabItemButtonV(label, 0)
	}, true)
	engine.RegisterMethod("imgui.beginDragDropSource", "开始拖放源", func() bool {
		return imgui.BeginDragDropSourceV(0)
	}, true)
	engine.RegisterMethod("imgui.endDragDropSource", "结束拖放源", func() {
		imgui.EndDragDropSource()
	}, true)
	engine.RegisterMethod("imgui.beginDragDropTarget", "开始拖放目标", func() bool {
		return imgui.BeginDragDropTarget()
	}, true)
	engine.RegisterMethod("imgui.endDragDropTarget", "结束拖放目标", func() {
		imgui.EndDragDropTarget()
	}, true)
	engine.RegisterMethod("imgui.setDragDropPayload", "设置拖放数据", func(typeArg, data string) bool {
		dataPtr := uintptr(0)
		if len(data) > 0 {
			dataPtr = uintptr(unsafe.Pointer(&[]byte(data)[0]))
		}
		return imgui.SetDragDropPayloadV(typeArg, dataPtr, uint64(len(data)), 0)
	}, true)
	engine.RegisterMethod("imgui.beginDisabled", "开始禁用", func(disabled bool) {
		imgui.BeginDisabledV(disabled)
	}, true)
	engine.RegisterMethod("imgui.endDisabled", "结束禁用", func() {
		imgui.EndDisabled()
	}, true)
	engine.RegisterMethod("imgui.showDemoWindow", "显示演示窗口", func() {
		imgui.ShowDemoWindow()
	}, true)
	engine.RegisterMethod("imgui.showMetricsWindow", "显示指标窗口", func() {
		imgui.ShowMetricsWindow()
	}, true)
	engine.RegisterMethod("imgui.showAboutWindow", "显示关于窗口", func() {
		imgui.ShowAboutWindow()
	}, true)
	engine.RegisterMethod("imgui.showStyleEditor", "显示样式编辑器", func() {
		imgui.ShowStyleEditor()
	}, true)
	engine.RegisterMethod("imgui.showStyleSelector", "显示样式选择器", func(label string) {
		imgui.ShowStyleSelector(label)
	}, true)
	engine.RegisterMethod("imgui.showFontSelector", "显示字体选择器", func(label string) {
		imgui.ShowFontSelector(label)
	}, true)
	engine.RegisterMethod("imgui.showUserGuide", "显示用户指南", func() {
		imgui.ShowUserGuide()
	}, true)
	engine.RegisterMethod("imgui.showDebugLogWindow", "显示调试日志窗口", func() {
		imgui.ShowDebugLogWindow()
	}, true)
	engine.RegisterMethod("imgui.showIDStackToolWindow", "显示 ID 栈工具窗口", func() {
		imgui.ShowIDStackToolWindow()
	}, true)
	engine.RegisterMethod("imgui.beginCombo", "开始组合框", func(label, preview_value string) bool {
		return imgui.BeginCombo(label, preview_value)
	}, true)
	engine.RegisterMethod("imgui.endCombo", "结束组合框", func() {
		imgui.EndCombo()
	}, true)
	engine.RegisterMethod("imgui.beginPopupModal", "开始模态弹出窗口", func(name string) bool {
		return imgui.BeginPopupModal(name)
	}, true)
	engine.RegisterMethod("imgui.openPopup", "打开弹出窗口", func(str_id string) {
		imgui.OpenPopupStr(str_id)
	}, true)
	engine.RegisterMethod("imgui.closeCurrentPopup", "关闭当前弹出窗口", func() {
		imgui.CloseCurrentPopup()
	}, true)
	engine.RegisterMethod("imgui.beginPopupContextItem", "开始项上下文弹出窗口", func() bool {
		return imgui.BeginPopupContextItem()
	}, true)
	engine.RegisterMethod("imgui.beginPopupContextWindow", "开始窗口上下文弹出窗口", func() bool {
		return imgui.BeginPopupContextWindow()
	}, true)
	engine.RegisterMethod("imgui.beginPopupContextVoid", "开始空白上下文弹出窗口", func() bool {
		return imgui.BeginPopupContextVoid()
	}, true)
	engine.RegisterMethod("imgui.isItemHovered", "项是否悬停", func() bool {
		return imgui.IsItemHovered()
	}, true)
	engine.RegisterMethod("imgui.isItemActive", "项是否激活", func() bool {
		return imgui.IsItemActive()
	}, true)
	engine.RegisterMethod("imgui.isItemFocused", "项是否聚焦", func() bool {
		return imgui.IsItemFocused()
	}, true)
	engine.RegisterMethod("imgui.isItemClicked", "项是否被点击", func() bool {
		return imgui.IsItemClicked()
	}, true)
	engine.RegisterMethod("imgui.isItemVisible", "项是否可见", func() bool {
		return imgui.IsItemVisible()
	}, true)
	engine.RegisterMethod("imgui.isItemEdited", "项是否被编辑", func() bool {
		return imgui.IsItemEdited()
	}, true)
	engine.RegisterMethod("imgui.isItemActivated", "项是否被激活", func() bool {
		return imgui.IsItemActivated()
	}, true)
	engine.RegisterMethod("imgui.isItemDeactivated", "项是否被停用", func() bool {
		return imgui.IsItemDeactivated()
	}, true)
	engine.RegisterMethod("imgui.isItemDeactivatedAfterEdit", "项是否在编辑后被停用", func() bool {
		return imgui.IsItemDeactivatedAfterEdit()
	}, true)
	engine.RegisterMethod("imgui.isItemToggledOpen", "项是否切换打开", func() bool {
		return imgui.IsItemToggledOpen()
	}, true)
	engine.RegisterMethod("imgui.isMouseDragging", "鼠标是否拖动", func() bool {
		return imgui.IsMouseDraggingV(0, -1)
	}, true)
	engine.RegisterMethod("imgui.isMouseHoveringRect", "鼠标是否悬停在矩形上", func(minX, minY, maxX, maxY float32) bool {
		r_min := imgui.Vec2{X: minX, Y: minY}
		r_max := imgui.Vec2{X: maxX, Y: maxY}
		return imgui.IsMouseHoveringRect(r_min, r_max)
	}, true)
	engine.RegisterMethod("imgui.isMousePosValid", "鼠标位置是否有效", func() bool {
		return imgui.IsMousePosValid()
	}, true)
	engine.RegisterMethod("imgui.resetMouseDragDelta", "重置鼠标拖动增量", func() {
		imgui.ResetMouseDragDelta()
	}, true)
	engine.RegisterMethod("imgui.setKeyboardFocusHere", "设置键盘焦点到此处", func() {
		imgui.SetKeyboardFocusHere()
	}, true)
	engine.RegisterMethod("imgui.setItemDefaultFocus", "设置项默认焦点", func() {
		imgui.SetItemDefaultFocus()
	}, true)
	engine.RegisterMethod("imgui.isAnyItemHovered", "任何项是否悬停", func() bool {
		return imgui.IsAnyItemHovered()
	}, true)
	engine.RegisterMethod("imgui.isAnyItemActive", "任何项是否激活", func() bool {
		return imgui.IsAnyItemActive()
	}, true)
	engine.RegisterMethod("imgui.isAnyItemFocused", "任何项是否聚焦", func() bool {
		return imgui.IsAnyItemFocused()
	}, true)
	engine.RegisterMethod("imgui.isWindowHovered", "窗口是否悬停", func() bool {
		return imgui.IsWindowHovered()
	}, true)
	engine.RegisterMethod("imgui.isWindowFocused", "窗口是否聚焦", func() bool {
		return imgui.IsWindowFocused()
	}, true)
	engine.RegisterMethod("imgui.isWindowCollapsed", "窗口是否折叠", func() bool {
		return imgui.IsWindowCollapsed()
	}, true)
	engine.RegisterMethod("imgui.isRectVisible", "矩形是否可见", func(w, h float32) bool {
		size := imgui.Vec2{X: w, Y: h}
		return imgui.IsRectVisible(size)
	}, true)
	engine.RegisterMethod("imgui.listBox", "列表框", func(label string, currentItem int, items []string) (bool, int) {
		currentItemPtr := int32(currentItem)
		clicked := imgui.ListBoxStrarr(label, &currentItemPtr, items, int32(len(items)))
		return clicked, int(currentItemPtr)
	}, true)
	engine.RegisterMethod("imgui.beginMainMenuBar", "开始主菜单栏", func() bool {
		return imgui.BeginMainMenuBar()
	}, true)
	engine.RegisterMethod("imgui.endMainMenuBar", "结束主菜单栏", func() {
		imgui.EndMainMenuBar()
	}, true)
	engine.RegisterMethod("imgui.beginItemTooltip", "开始项工具提示", func() bool {
		return imgui.BeginItemTooltip()
	}, true)
	engine.RegisterMethod("imgui.setNextItemOpen", "设置下一项打开状态", func(isOpen bool) {
		imgui.SetNextItemOpen(isOpen)
	}, true)
	engine.RegisterMethod("imgui.setColorEditOptions", "设置颜色编辑选项", func(flags uint32) {
		imgui.SetColorEditOptions(imgui.ColorEditFlags(flags))
	}, true)
	engine.RegisterMethod("imgui.setTabItemClosed", "设置标签项关闭", func(label string) {
		imgui.SetTabItemClosed(label)
	}, true)
	engine.RegisterMethod("imgui.newDrawList", "创建绘制列表", func() *imgui.DrawList {
		return imgui.NewEmptyDrawList()
	}, true)
	engine.RegisterMethod("imgui.newFont", "创建字体对象", func() *imgui.Font {
		return imgui.NewEmptyFont()
	}, true)
	engine.RegisterMethod("imgui.newFontAtlas", "创建字体图集", func() *imgui.FontAtlas {
		return imgui.NewEmptyFontAtlas()
	}, true)
	engine.RegisterMethod("imgui.newFontConfig", "创建字体配置", func() *imgui.FontConfig {
		return imgui.NewEmptyFontConfig()
	}, true)
	engine.RegisterMethod("imgui.newDrawData", "创建绘制数据", func() *imgui.DrawData {
		return imgui.NewEmptyDrawData()
	}, true)
	engine.RegisterMethod("imgui.newStorage", "创建存储对象", func() *imgui.Storage {
		return imgui.NewEmptyStorage()
	}, true)
	engine.RegisterMethod("imgui.newPlatformIO", "创建平台IO对象", func() *imgui.PlatformIO {
		return imgui.NewEmptyPlatformIO()
	}, true)
	engine.RegisterMethod("imgui.newDrawCmd", "创建绘制命令", func() *imgui.DrawCmd {
		return imgui.NewEmptyDrawCmd()
	}, true)
	engine.RegisterMethod("imgui.newDrawVert", "创建绘制顶点", func() *imgui.DrawVert {
		return imgui.NewEmptyDrawVert()
	}, true)
	engine.RegisterMethod("imgui.newFontGlyph", "创建字体字形", func() *imgui.FontGlyph {
		return imgui.NewEmptyFontGlyph()
	}, true)
	engine.RegisterMethod("imgui.newFontBaked", "创建烘焙字体", func() *imgui.FontBaked {
		return imgui.NewEmptyFontBaked()
	}, true)
	engine.RegisterMethod("imgui.newFontLoader", "创建字体加载器", func() *imgui.FontLoader {
		return imgui.NewEmptyFontLoader()
	}, true)
	engine.RegisterMethod("imgui.newDrawListSharedData", "创建绘制列表共享数据", func() *imgui.DrawListSharedData {
		return imgui.NewEmptyDrawListSharedData()
	}, true)
	engine.RegisterMethod("imgui.newDrawListSplitter", "创建绘制列表分割器", func() *imgui.DrawListSplitter {
		return imgui.NewEmptyDrawListSplitter()
	}, true)
	engine.RegisterMethod("imgui.newDrawChannel", "创建绘制通道", func() *imgui.DrawChannel {
		return imgui.NewEmptyDrawChannel()
	}, true)
	engine.RegisterMethod("imgui.newFontAtlasBuilder", "创建字体图集构建器", func() *imgui.FontAtlasBuilder {
		return imgui.NewEmptyFontAtlasBuilder()
	}, true)
	engine.RegisterMethod("imgui.newFontAtlasRect", "创建字体图集矩形", func() *imgui.FontAtlasRect {
		return imgui.NewEmptyFontAtlasRect()
	}, true)
	engine.RegisterMethod("imgui.newFontAtlasRectEntry", "创建字体图集矩形条目", func() *imgui.FontAtlasRectEntry {
		return imgui.NewEmptyFontAtlasRectEntry()
	}, true)
	engine.RegisterMethod("imgui.newFontAtlasPostProcessData", "创建字体图集后处理数据", func() *imgui.FontAtlasPostProcessData {
		return imgui.NewEmptyFontAtlasPostProcessData()
	}, true)
	engine.RegisterMethod("imgui.newFontGlyphRangesBuilder", "创建字体字形范围构建器", func() *imgui.FontGlyphRangesBuilder {
		return imgui.NewEmptyFontGlyphRangesBuilder()
	}, true)
	engine.RegisterMethod("imgui.newFontStackData", "创建字体栈数据", func() *imgui.FontStackData {
		return imgui.NewEmptyFontStackData()
	}, true)
	engine.RegisterMethod("imgui.newBoxSelectState", "创建框选状态", func() *imgui.BoxSelectState {
		return imgui.NewEmptyBoxSelectState()
	}, true)
	engine.RegisterMethod("imgui.newColorMod", "创建颜色修改", func() *imgui.ColorMod {
		return imgui.NewEmptyColorMod()
	}, true)
	engine.RegisterMethod("imgui.newComboPreviewData", "创建组合预览数据", func() *imgui.ComboPreviewData {
		return imgui.NewEmptyComboPreviewData()
	}, true)
	engine.RegisterMethod("imgui.newContextHook", "创建上下文钩子", func() *imgui.ContextHook {
		return imgui.NewEmptyContextHook()
	}, true)
	engine.RegisterMethod("imgui.newDataTypeInfo", "创建数据类型信息", func() *imgui.DataTypeInfo {
		return imgui.NewEmptyDataTypeInfo()
	}, true)
	engine.RegisterMethod("imgui.newDataTypeStorage", "创建数据类型存储", func() *imgui.DataTypeStorage {
		return imgui.NewEmptyDataTypeStorage()
	}, true)
	engine.RegisterMethod("imgui.newDeactivatedItemData", "创建停用项数据", func() *imgui.DeactivatedItemData {
		return imgui.NewEmptyDeactivatedItemData()
	}, true)
	engine.RegisterMethod("imgui.newDebugAllocEntry", "创建调试分配条目", func() *imgui.DebugAllocEntry {
		return imgui.NewEmptyDebugAllocEntry()
	}, true)
	engine.RegisterMethod("imgui.newDebugAllocInfo", "创建调试分配信息", func() *imgui.DebugAllocInfo {
		return imgui.NewEmptyDebugAllocInfo()
	}, true)
	engine.RegisterMethod("imgui.newDockContext", "创建停靠上下文", func() *imgui.DockContext {
		return imgui.NewEmptyDockContext()
	}, true)
	engine.RegisterMethod("imgui.newDockNode", "创建停靠节点", func() *imgui.DockNode {
		return imgui.NewEmptyDockNode()
	}, true)
	engine.RegisterMethod("imgui.newErrorRecoveryState", "创建错误恢复状态", func() *imgui.ErrorRecoveryState {
		return imgui.NewEmptyErrorRecoveryState()
	}, true)
	engine.RegisterMethod("imgui.newFocusScopeData", "创建焦点范围数据", func() *imgui.FocusScopeData {
		return imgui.NewEmptyFocusScopeData()
	}, true)
	engine.RegisterMethod("imgui.newGroupData", "创建组数据", func() *imgui.GroupData {
		return imgui.NewEmptyGroupData()
	}, true)
	engine.RegisterMethod("imgui.newIDStackTool", "创建ID栈工具", func() *imgui.IDStackTool {
		return imgui.NewEmptyIDStackTool()
	}, true)
	engine.RegisterMethod("imgui.newInputEvent", "创建输入事件", func() *imgui.InputEvent {
		return imgui.NewEmptyInputEvent()
	}, true)
	engine.RegisterMethod("imgui.newInputEventAppFocused", "创建应用聚焦输入事件", func() *imgui.InputEventAppFocused {
		return imgui.NewEmptyInputEventAppFocused()
	}, true)
	engine.RegisterMethod("imgui.newInputEventKey", "创建键盘输入事件", func() *imgui.InputEventKey {
		return imgui.NewEmptyInputEventKey()
	}, true)
	engine.RegisterMethod("imgui.newInputEventMouseButton", "创建鼠标按钮输入事件", func() *imgui.InputEventMouseButton {
		return imgui.NewEmptyInputEventMouseButton()
	}, true)
	engine.RegisterMethod("imgui.newInputEventMousePos", "创建鼠标位置输入事件", func() *imgui.InputEventMousePos {
		return imgui.NewEmptyInputEventMousePos()
	}, true)
	engine.RegisterMethod("imgui.newInputEventMouseViewport", "创建鼠标视口输入事件", func() *imgui.InputEventMouseViewport {
		return imgui.NewEmptyInputEventMouseViewport()
	}, true)
	engine.RegisterMethod("imgui.newInputEventMouseWheel", "创建鼠标滚轮输入事件", func() *imgui.InputEventMouseWheel {
		return imgui.NewEmptyInputEventMouseWheel()
	}, true)
	engine.RegisterMethod("imgui.newInputEventText", "创建文本输入事件", func() *imgui.InputEventText {
		return imgui.NewEmptyInputEventText()
	}, true)
	engine.RegisterMethod("imgui.newInputTextCallbackData", "创建输入文本回调数据", func() *imgui.InputTextCallbackData {
		return imgui.NewEmptyInputTextCallbackData()
	}, true)
	engine.RegisterMethod("imgui.newInputTextDeactivatedState", "创建输入文本停用状态", func() *imgui.InputTextDeactivatedState {
		return imgui.NewEmptyInputTextDeactivatedState()
	}, true)
	engine.RegisterMethod("imgui.newInputTextState", "创建输入文本状态", func() *imgui.InputTextState {
		return imgui.NewEmptyInputTextState()
	}, true)
	engine.RegisterMethod("imgui.newKeyData", "创建键数据", func() *imgui.KeyData {
		return imgui.NewEmptyKeyData()
	}, true)
	engine.RegisterMethod("imgui.newKeyOwnerData", "创建键拥有者数据", func() *imgui.KeyOwnerData {
		return imgui.NewEmptyKeyOwnerData()
	}, true)
	engine.RegisterMethod("imgui.newKeyRoutingData", "创建键路由数据", func() *imgui.KeyRoutingData {
		return imgui.NewEmptyKeyRoutingData()
	}, true)
	engine.RegisterMethod("imgui.newKeyRoutingTable", "创建键路由表", func() *imgui.KeyRoutingTable {
		return imgui.NewEmptyKeyRoutingTable()
	}, true)
	engine.RegisterMethod("imgui.newLastItemData", "创建最后项数据", func() *imgui.LastItemData {
		return imgui.NewEmptyLastItemData()
	}, true)
	engine.RegisterMethod("imgui.newListClipper", "创建列表剪裁器", func() *imgui.ListClipper {
		return imgui.NewEmptyListClipper()
	}, true)
	engine.RegisterMethod("imgui.newListClipperData", "创建列表剪裁器数据", func() *imgui.ListClipperData {
		return imgui.NewEmptyListClipperData()
	}, true)
	engine.RegisterMethod("imgui.newListClipperRange", "创建列表剪裁器范围", func() *imgui.ListClipperRange {
		return imgui.NewEmptyListClipperRange()
	}, true)
	engine.RegisterMethod("imgui.newLocEntry", "创建本地化条目", func() *imgui.LocEntry {
		return imgui.NewEmptyLocEntry()
	}, true)
	engine.RegisterMethod("imgui.newMenuColumns", "创建菜单列", func() *imgui.MenuColumns {
		return imgui.NewEmptyMenuColumns()
	}, true)
	engine.RegisterMethod("imgui.newMetricsConfig", "创建指标配置", func() *imgui.MetricsConfig {
		return imgui.NewEmptyMetricsConfig()
	}, true)
	engine.RegisterMethod("imgui.newMultiSelectIO", "创建多选IO", func() *imgui.MultiSelectIO {
		return imgui.NewEmptyMultiSelectIO()
	}, true)
	engine.RegisterMethod("imgui.newMultiSelectState", "创建多选状态", func() *imgui.MultiSelectState {
		return imgui.NewEmptyMultiSelectState()
	}, true)
	engine.RegisterMethod("imgui.newMultiSelectTempData", "创建多选临时数据", func() *imgui.MultiSelectTempData {
		return imgui.NewEmptyMultiSelectTempData()
	}, true)
	engine.RegisterMethod("imgui.newNavItemData", "创建导航项数据", func() *imgui.NavItemData {
		return imgui.NewEmptyNavItemData()
	}, true)
	engine.RegisterMethod("imgui.newNextItemData", "创建下一项数据", func() *imgui.NextItemData {
		return imgui.NewEmptyNextItemData()
	}, true)
	engine.RegisterMethod("imgui.newNextWindowData", "创建下一窗口数据", func() *imgui.NextWindowData {
		return imgui.NewEmptyNextWindowData()
	}, true)
	engine.RegisterMethod("imgui.newOldColumnData", "创建旧列数据", func() *imgui.OldColumnData {
		return imgui.NewEmptyOldColumnData()
	}, true)
	engine.RegisterMethod("imgui.newOldColumns", "创建旧列", func() *imgui.OldColumns {
		return imgui.NewEmptyOldColumns()
	}, true)
	engine.RegisterMethod("imgui.newOnceUponAFrame", "创建每帧一次", func() *imgui.OnceUponAFrame {
		return imgui.NewEmptyOnceUponAFrame()
	}, true)
	engine.RegisterMethod("imgui.newPayload", "创建载荷", func() *imgui.Payload {
		return imgui.NewEmptyPayload()
	}, true)
	engine.RegisterMethod("imgui.newPlatformImeData", "创建平台IME数据", func() *imgui.PlatformImeData {
		return imgui.NewEmptyPlatformImeData()
	}, true)
	engine.RegisterMethod("imgui.newPlatformMonitor", "创建平台监视器", func() *imgui.PlatformMonitor {
		return imgui.NewEmptyPlatformMonitor()
	}, true)
	engine.RegisterMethod("imgui.newPopupData", "创建弹出数据", func() *imgui.PopupData {
		return imgui.NewEmptyPopupData()
	}, true)
	engine.RegisterMethod("imgui.newPtrOrIndex", "创建指针或索引", func() *imgui.PtrOrIndex {
		return imgui.NewEmptyPtrOrIndex()
	}, true)
	engine.RegisterMethod("imgui.newSelectionBasicStorage", "创建选择基本存储", func() *imgui.SelectionBasicStorage {
		return imgui.NewEmptySelectionBasicStorage()
	}, true)
	engine.RegisterMethod("imgui.newSelectionExternalStorage", "创建选择外部存储", func() *imgui.SelectionExternalStorage {
		return imgui.NewEmptySelectionExternalStorage()
	}, true)
	engine.RegisterMethod("imgui.newSelectionRequest", "创建选择请求", func() *imgui.SelectionRequest {
		return imgui.NewEmptySelectionRequest()
	}, true)
	engine.RegisterMethod("imgui.newSettingsHandler", "创建设置处理器", func() *imgui.SettingsHandler {
		return imgui.NewEmptySettingsHandler()
	}, true)
	engine.RegisterMethod("imgui.newShrinkWidthItem", "创建收缩宽度项", func() *imgui.ShrinkWidthItem {
		return imgui.NewEmptyShrinkWidthItem()
	}, true)
	engine.RegisterMethod("imgui.newSizeCallbackData", "创建大小回调数据", func() *imgui.SizeCallbackData {
		return imgui.NewEmptySizeCallbackData()
	}, true)
	engine.RegisterMethod("imgui.newStackLevelInfo", "创建栈级别信息", func() *imgui.StackLevelInfo {
		return imgui.NewEmptyStackLevelInfo()
	}, true)
	engine.RegisterMethod("imgui.newStoragePair", "创建存储对", func() *imgui.StoragePair {
		return imgui.NewEmptyStoragePair()
	}, true)
	engine.RegisterMethod("imgui.newStyleMod", "创建样式修改", func() *imgui.StyleMod {
		return imgui.NewEmptyStyleMod()
	}, true)
	engine.RegisterMethod("imgui.newStyleVarInfo", "创建样式变量信息", func() *imgui.StyleVarInfo {
		return imgui.NewEmptyStyleVarInfo()
	}, true)
	engine.RegisterMethod("imgui.newTabBar", "创建标签栏", func() *imgui.TabBar {
		return imgui.NewEmptyTabBar()
	}, true)
	engine.RegisterMethod("imgui.newTabItem", "创建标签项", func() *imgui.TabItem {
		return imgui.NewEmptyTabItem()
	}, true)
	engine.RegisterMethod("imgui.newTable", "创建表格", func() *imgui.Table {
		return imgui.NewEmptyTable()
	}, true)
	engine.RegisterMethod("imgui.newTableCellData", "创建表格单元格数据", func() *imgui.TableCellData {
		return imgui.NewEmptyTableCellData()
	}, true)
	engine.RegisterMethod("imgui.newTableColumn", "创建表格列", func() *imgui.TableColumn {
		return imgui.NewEmptyTableColumn()
	}, true)
	engine.RegisterMethod("imgui.newViewport", "创建视口", func() *imgui.Viewport {
		return imgui.NewEmptyViewport()
	}, true)
	engine.RegisterMethod("imgui.newWindowClass", "创建窗口类", func() *imgui.WindowClass {
		return imgui.NewEmptyWindowClass()
	}, true)
	engine.RegisterMethod("imgui.newTextBuffer", "创建文本缓冲区", func() *imgui.TextBuffer {
		return imgui.NewEmptyTextBuffer()
	}, true)
	engine.RegisterMethod("imgui.newTextFilter", "创建文本过滤器", func() *imgui.TextFilter {
		return imgui.NewEmptyTextFilter()
	}, true)
	engine.RegisterMethod("imgui.newTextRange", "创建文本范围", func() *imgui.TextRange {
		return imgui.NewEmptyTextRange()
	}, true)
	engine.RegisterMethod("imgui.newTypingSelectState", "创建打字选择状态", func() *imgui.TypingSelectState {
		return imgui.NewEmptyTypingSelectState()
	}, true)
	engine.RegisterMethod("imgui.newViewportP", "创建视口指针", func() *imgui.ViewportP {
		return imgui.NewEmptyViewportP()
	}, true)
	engine.RegisterMethod("imgui.newTextureData", "创建纹理数据", func() *imgui.TextureData {
		return imgui.NewEmptyTextureData()
	}, true)
	engine.RegisterMethod("imgui.newTextureRef", "创建纹理引用", func() *imgui.TextureRef {
		return imgui.NewEmptyTextureRef()
	}, true)
	engine.RegisterMethod("imgui.newDrawCmdHeader", "创建绘制命令头", func() *imgui.DrawCmdHeader {
		return imgui.NewEmptyDrawCmdHeader()
	}, true)
	engine.RegisterMethod("imgui.newDrawDataBuilder", "创建绘制数据构建器", func() *imgui.DrawDataBuilder {
		return imgui.NewEmptyDrawDataBuilder()
	}, true)
	engine.RegisterMethod("imgui.newBitVector", "创建位向量", func() *imgui.BitVector {
		return imgui.NewEmptyBitVector()
	}, true)
	engine.RegisterMethod("imgui.newBitArrayForNamedKeys", "创建命名键位数组", func() *imgui.BitArrayForNamedKeys {
		return imgui.NewEmptyBitArrayForNamedKeys()
	}, true)
	engine.RegisterMethod("imgui.colorHSVV", "HSV颜色转RGBA", func(h, s, v, a float32) (float32, float32, float32, float32) {
		col := imgui.ColorHSVV(h, s, v, a)
		return col.FieldValue.X, col.FieldValue.Y, col.FieldValue.Z, col.FieldValue.W
	}, true)
	engine.RegisterMethod("imgui.alignTextToFramePadding", "对齐文本到框架内边距", func() {
		imgui.AlignTextToFramePadding()
	}, true)
	engine.RegisterMethod("imgui.arrowButton", "箭头按钮", func(str_id string, dir int32) bool {
		return imgui.ArrowButton(str_id, imgui.Dir(dir))
	}, true)
	engine.RegisterMethod("imgui.progressBar", "进度条", func(fraction float32) {
		imgui.ProgressBar(fraction)
	}, true)
	engine.RegisterMethod("imgui.progressBarWithSize", "带大小的进度条", func(fraction float32, w, h float32) {
		size := imgui.Vec2{X: w, Y: h}
		imgui.ProgressBarV(fraction, size, "")
	}, true)
	engine.RegisterMethod("imgui.progressBarWithOverlay", "带覆盖的进度条", func(fraction float32, w, h float32, overlay string) {
		size := imgui.Vec2{X: w, Y: h}
		imgui.ProgressBarV(fraction, size, overlay)
	}, true)
	engine.RegisterMethod("imgui.setScrollHereX", "设置滚动到此处X", func() {
		imgui.SetScrollHereX()
	}, true)
	engine.RegisterMethod("imgui.setScrollHereY", "设置滚动到此处Y", func() {
		imgui.SetScrollHereY()
	}, true)
	engine.RegisterMethod("imgui.setNextWindowScroll", "设置下一窗口滚动", func(scrollX, scrollY float32) {
		imgui.SetNextWindowScroll(imgui.Vec2{X: scrollX, Y: scrollY})
	}, true)
	engine.RegisterMethod("imgui.setTooltip", "设置工具提示", func(text string) {
		imgui.SetTooltip(text)
	}, true)
	engine.RegisterMethod("imgui.beginTooltip", "开始工具提示", func() {
		imgui.BeginTooltip()
	}, true)
	engine.RegisterMethod("imgui.endTooltip", "结束工具提示", func() {
		imgui.EndTooltip()
	}, true)
	engine.RegisterMethod("imgui.openPopup", "打开弹出窗口", func(str_id string) {
		imgui.OpenPopupStr(str_id)
	}, true)
	engine.RegisterMethod("imgui.openPopupOnItemClick", "点击项时打开弹出窗口", func() {
		imgui.OpenPopupOnItemClick()
	}, true)
	engine.RegisterMethod("imgui.closeCurrentPopup", "关闭当前弹出窗口", func() {
		imgui.CloseCurrentPopup()
	}, true)
	engine.RegisterMethod("imgui.beginPopupModal", "开始模态弹出窗口", func(name string) bool {
		return imgui.BeginPopupModal(name)
	}, true)
	engine.RegisterMethod("imgui.beginPopupContextWindow", "开始窗口上下文弹出窗口", func() bool {
		return imgui.BeginPopupContextWindow()
	}, true)
	engine.RegisterMethod("imgui.beginPopupContextVoid", "开始空白上下文弹出窗口", func() bool {
		return imgui.BeginPopupContextVoid()
	}, true)
	engine.RegisterMethod("imgui.beginPopup", "开始弹出窗口", func(str_id string) bool {
		return imgui.BeginPopup(str_id)
	}, true)
	engine.RegisterMethod("imgui.endPopup", "结束弹出窗口", func() {
		imgui.EndPopup()
	}, true)
	engine.RegisterMethod("imgui.calcTextSize", "计算文本大小", func(text string) (float32, float32) {
		size := imgui.CalcTextSize(text)
		return size.X, size.Y
	}, true)
	engine.RegisterMethod("imgui.pushClipRect", "推入剪裁矩形", func(min_x, min_y, max_x, max_y float32, intersect_with_current_clip_rect bool) {
		min_rect := imgui.Vec2{X: min_x, Y: min_y}
		max_rect := imgui.Vec2{X: max_x, Y: max_y}
		imgui.PushClipRect(min_rect, max_rect, intersect_with_current_clip_rect)
	}, true)
	engine.RegisterMethod("imgui.popClipRect", "弹出剪裁矩形", func() {
		imgui.PopClipRect()
	}, true)
	engine.RegisterMethod("imgui.setItemKeyOwner", "设置项键拥有者", func(key int32) {
		imgui.SetItemKeyOwner(imgui.Key(key))
	}, true)
	engine.RegisterMethod("imgui.setNextItemAllowOverlap", "允许下一项重叠", func() {
		imgui.SetNextItemAllowOverlap()
	}, true)
	engine.RegisterMethod("imgui.isItemToggledSelection", "项是否切换选择", func() bool {
		return imgui.IsItemToggledSelection()
	}, true)
	engine.RegisterMethod("imgui.beginMultiSelect", "开始多选", func(flags uint32) *imgui.MultiSelectIO {
		return imgui.BeginMultiSelect(imgui.MultiSelectFlags(flags))
	}, true)
	engine.RegisterMethod("imgui.endMultiSelect", "结束多选", func() *imgui.MultiSelectIO {
		return imgui.EndMultiSelect()
	}, true)
	engine.RegisterMethod("imgui.resetMouseDragDelta", "重置鼠标拖拽增量", func() {
		imgui.ResetMouseDragDelta()
	}, true)
	engine.RegisterMethod("imgui.setNavCursorVisible", "设置导航光标可见", func(visible bool) {
		imgui.SetNavCursorVisible(visible)
	}, true)
	engine.RegisterMethod("imgui.setMouseCursor", "设置鼠标光标", func(cursor_type int32) {
		imgui.SetMouseCursor(imgui.MouseCursor(cursor_type))
	}, true)
	engine.RegisterMethod("imgui.setNextFrameWantCaptureMouse", "设置下一帧想要捕获鼠标", func(want_capture_mouse bool) {
		imgui.SetNextFrameWantCaptureMouse(want_capture_mouse)
	}, true)
	engine.RegisterMethod("imgui.setNextFrameWantCaptureKeyboard", "设置下一帧想要捕获键盘", func(want_capture_keyboard bool) {
		imgui.SetNextFrameWantCaptureKeyboard(want_capture_keyboard)
	}, true)
	engine.RegisterMethod("imgui.pushStyleVarFloat", "推入样式变量（浮点）", func(idx int32, val float32) {
		imgui.PushStyleVarFloat(imgui.StyleVar(idx), val)
	}, true)
	engine.RegisterMethod("imgui.pushStyleVarVec2", "推入样式变量（向量）", func(idx int32, val_x, val_y float32) {
		imgui.PushStyleVarVec2(imgui.StyleVar(idx), imgui.Vec2{X: val_x, Y: val_y})
	}, true)
	engine.RegisterMethod("imgui.popStyleVar", "弹出样式变量", func() {
		imgui.PopStyleVar()
	}, true)
	engine.RegisterMethod("imgui.pushFont", "推入字体", func(font *imgui.Font, scale float32) {
		imgui.PushFont(font, scale)
	}, true)
	engine.RegisterMethod("imgui.popFont", "弹出字体", func() {
		imgui.PopFont()
	}, true)
	engine.RegisterMethod("imgui.pushItemFlag", "推入项标志", func(flag int32, enabled bool) {
		imgui.PushItemFlag(imgui.ItemFlags(flag), enabled)
	}, true)
	engine.RegisterMethod("imgui.popItemFlag", "弹出项标志", func() {
		imgui.PopItemFlag()
	}, true)
	engine.RegisterMethod("imgui.pushIDStr", "推入ID（字符串）", func(str_id string) {
		imgui.PushIDStr(str_id)
	}, true)
	engine.RegisterMethod("imgui.pushIDInt", "推入ID（整数）", func(int_id int32) {
		imgui.PushIDInt(int_id)
	}, true)
	engine.RegisterMethod("imgui.pushIDPtr", "推入ID（指针）", func(ptr_id uintptr) {
		imgui.PushIDPtr(ptr_id)
	}, true)
	engine.RegisterMethod("imgui.popID", "弹出ID", func() {
		imgui.PopID()
	}, true)
	engine.RegisterMethod("imgui.pushTextWrapPos", "推入文本换行位置", func() {
		imgui.PushTextWrapPos()
	}, true)
	engine.RegisterMethod("imgui.popTextWrapPos", "弹出文本换行位置", func() {
		imgui.PopTextWrapPos()
	}, true)
	engine.RegisterMethod("imgui.newLine", "新行", func() {
		imgui.NewLine()
	}, true)
	engine.RegisterMethod("imgui.spacing", "间距", func() {
		imgui.Spacing()
	}, true)
	engine.RegisterMethod("imgui.dummy", "虚拟元素", func(w, h float32) {
		imgui.Dummy(imgui.Vec2{X: w, Y: h})
	}, true)
	engine.RegisterMethod("imgui.indent", "缩进", func(indent_w float32) {
		imgui.IndentV(indent_w)
	}, true)
	engine.RegisterMethod("imgui.unindent", "取消缩进", func(indent_w float32) {
		imgui.UnindentV(indent_w)
	}, true)
	engine.RegisterMethod("imgui.beginGroup", "开始组", func() {
		imgui.BeginGroup()
	}, true)
	engine.RegisterMethod("imgui.endGroup", "结束组", func() {
		imgui.EndGroup()
	}, true)
	engine.RegisterMethod("imgui.getCursorPos", "获取光标位置", func() (float32, float32) {
		pos := imgui.CursorPos()
		return pos.X, pos.Y
	}, true)
	engine.RegisterMethod("imgui.getCursorPosX", "获取光标X位置", func() float32 {
		return imgui.CursorPosX()
	}, true)
	engine.RegisterMethod("imgui.getCursorPosY", "获取光标Y位置", func() float32 {
		return imgui.CursorPosY()
	}, true)
	engine.RegisterMethod("imgui.setCursorPos", "设置光标位置", func(x, y float32) {
		imgui.SetCursorPos(imgui.Vec2{X: x, Y: y})
	}, true)
	engine.RegisterMethod("imgui.setCursorPosX", "设置光标X位置", func(x float32) {
		imgui.SetCursorPosX(x)
	}, true)
	engine.RegisterMethod("imgui.setCursorPosY", "设置光标Y位置", func(y float32) {
		imgui.SetCursorPosY(y)
	}, true)
	engine.RegisterMethod("imgui.getCursorStartPos", "获取光标起始位置", func() (float32, float32) {
		pos := imgui.CursorStartPos()
		return pos.X, pos.Y
	}, true)
	engine.RegisterMethod("imgui.getCursorScreenPos", "获取屏幕光标位置", func() (float32, float32) {
		pos := imgui.CursorScreenPos()
		return pos.X, pos.Y
	}, true)
	engine.RegisterMethod("imgui.setCursorScreenPos", "设置屏幕光标位置", func(x, y float32) {
		imgui.SetCursorScreenPos(imgui.Vec2{X: x, Y: y})
	}, true)
	engine.RegisterMethod("imgui.getTextLineHeight", "获取文本行高度", func() float32 {
		return imgui.TextLineHeight()
	}, true)
	engine.RegisterMethod("imgui.getTextLineHeightWithSpacing", "获取带间距的文本行高度", func() float32 {
		return imgui.TextLineHeightWithSpacing()
	}, true)
	engine.RegisterMethod("imgui.getFrameHeight", "获取框架高度", func() float32 {
		return imgui.FrameHeight()
	}, true)
	engine.RegisterMethod("imgui.getFrameHeightWithSpacing", "获取带间距的框架高度", func() float32 {
		return imgui.FrameHeightWithSpacing()
	}, true)
	engine.RegisterMethod("imgui.getContentRegionMax", "获取内容区域最大值", func() (float32, float32) {
		max := imgui.ContentRegionAvail()
		return max.X, max.Y
	}, true)
	engine.RegisterMethod("imgui.getContentRegionAvail", "获取可用内容区域", func() (float32, float32) {
		avail := imgui.ContentRegionAvail()
		return avail.X, avail.Y
	}, true)
	engine.RegisterMethod("imgui.popStyleColor", "弹出样式颜色", func() {
		imgui.PopStyleColor()
	}, true)
	engine.RegisterMethod("imgui.getStyleColor", "获取样式颜色", func(idx int32) (float32, float32, float32, float32) {
		col := imgui.StyleColorVec4(imgui.Col(idx))
		return col.X, col.Y, col.Z, col.W
	}, true)
	engine.RegisterMethod("imgui.getStyleColorName", "获取样式颜色名称", func(idx int32) string {
		return imgui.StyleColorName(imgui.Col(idx))
	}, true)
	engine.RegisterMethod("imgui.beginMainMenuBar", "开始主菜单栏", func() bool {
		return imgui.BeginMainMenuBar()
	}, true)
	engine.RegisterMethod("imgui.endMainMenuBar", "结束主菜单栏", func() {
		imgui.EndMainMenuBar()
	}, true)
	engine.RegisterMethod("imgui.beginMenuBar", "开始菜单栏", func() bool {
		return imgui.BeginMenuBar()
	}, true)
	engine.RegisterMethod("imgui.endMenuBar", "结束菜单栏", func() {
		imgui.EndMenuBar()
	}, true)
	engine.RegisterMethod("imgui.beginMenu", "开始菜单", func(label string) bool {
		return imgui.BeginMenuV(label, true)
	}, true)
	engine.RegisterMethod("imgui.beginMenuDisabled", "开始禁用菜单", func(label string) bool {
		return imgui.BeginMenuV(label, false)
	}, true)
	engine.RegisterMethod("imgui.endMenu", "结束菜单", func() {
		imgui.EndMenu()
	}, true)
	engine.RegisterMethod("imgui.beginTabBar", "开始标签栏", func(str_id string) bool {
		return imgui.BeginTabBarV(str_id, 0)
	}, true)
	engine.RegisterMethod("imgui.beginTabBarWithFlags", "开始标签栏（带标志）", func(str_id string, flags uint32) bool {
		return imgui.BeginTabBarV(str_id, imgui.TabBarFlags(flags))
	}, true)
	engine.RegisterMethod("imgui.endTabBar", "结束标签栏", func() {
		imgui.EndTabBar()
	}, true)
	engine.RegisterMethod("imgui.beginTabItem", "开始标签项", func(label string) bool {
		return imgui.BeginTabItemV(label, nil, 0)
	}, true)
	engine.RegisterMethod("imgui.beginTabItemWithOpen", "开始标签项（带打开状态）", func(label string, open *bool) bool {
		return imgui.BeginTabItemV(label, open, 0)
	}, true)
	engine.RegisterMethod("imgui.beginTabItemWithFlags", "开始标签项（带标志）", func(label string, flags uint32) bool {
		return imgui.BeginTabItemV(label, nil, imgui.TabItemFlags(flags))
	}, true)
	engine.RegisterMethod("imgui.endTabItem", "结束标签项", func() {
		imgui.EndTabItem()
	}, true)
	engine.RegisterMethod("imgui.setTabItemClosed", "设置标签项关闭", func(tab_or_docked_window_label string) {
		imgui.SetTabItemClosed(tab_or_docked_window_label)
	}, true)
	engine.RegisterMethod("imgui.isItemHovered", "项是否悬停", func() bool {
		return imgui.IsItemHovered()
	}, true)
	engine.RegisterMethod("imgui.isItemActive", "项是否激活", func() bool {
		return imgui.IsItemActive()
	}, true)
	engine.RegisterMethod("imgui.isItemFocused", "项是否聚焦", func() bool {
		return imgui.IsItemFocused()
	}, true)
	engine.RegisterMethod("imgui.isItemClicked", "项是否被点击", func() bool {
		return imgui.IsItemClicked()
	}, true)
	engine.RegisterMethod("imgui.isItemVisible", "项是否可见", func() bool {
		return imgui.IsItemVisible()
	}, true)
	engine.RegisterMethod("imgui.isItemEdited", "项是否被编辑", func() bool {
		return imgui.IsItemEdited()
	}, true)
	engine.RegisterMethod("imgui.isItemActivated", "项是否被激活", func() bool {
		return imgui.IsItemActivated()
	}, true)
	engine.RegisterMethod("imgui.isItemDeactivated", "项是否被停用", func() bool {
		return imgui.IsItemDeactivated()
	}, true)
	engine.RegisterMethod("imgui.isItemDeactivatedAfterEdit", "项是否在编辑后被停用", func() bool {
		return imgui.IsItemDeactivatedAfterEdit()
	}, true)
	engine.RegisterMethod("imgui.isItemToggledOpen", "项是否切换打开状态", func() bool {
		return imgui.IsItemToggledOpen()
	}, true)
	engine.RegisterMethod("imgui.isAnyItemHovered", "是否有任何项悬停", func() bool {
		return imgui.IsAnyItemHovered()
	}, true)
	engine.RegisterMethod("imgui.isAnyItemActive", "是否有任何项激活", func() bool {
		return imgui.IsAnyItemActive()
	}, true)
	engine.RegisterMethod("imgui.isAnyItemFocused", "是否有任何项聚焦", func() bool {
		return imgui.IsAnyItemFocused()
	}, true)
	engine.RegisterMethod("imgui.getItemRectMin", "获取项矩形最小值", func() (float32, float32) {
		min := imgui.ItemRectMin()
		return min.X, min.Y
	}, true)
	engine.RegisterMethod("imgui.getItemRectMax", "获取项矩形最大值", func() (float32, float32) {
		max := imgui.ItemRectMax()
		return max.X, max.Y
	}, true)
	engine.RegisterMethod("imgui.getItemRectSize", "获取项矩形大小", func() (float32, float32) {
		size := imgui.ItemRectSize()
		return size.X, size.Y
	}, true)
	engine.RegisterMethod("imgui.setItemDefaultFocus", "设置项默认焦点", func() {
		imgui.SetItemDefaultFocus()
	}, true)
	engine.RegisterMethod("imgui.setKeyboardFocusHere", "设置键盘焦点到此处", func() {
		imgui.SetKeyboardFocusHere()
	}, true)
	engine.RegisterMethod("imgui.isMouseHoveringRect", "鼠标是否悬停在矩形上", func(min_x, min_y, max_x, max_y float32) bool {
		min := imgui.Vec2{X: min_x, Y: min_y}
		max := imgui.Vec2{X: max_x, Y: max_y}
		return imgui.IsMouseHoveringRect(min, max)
	}, true)
	engine.RegisterMethod("imgui.isMousePosValid", "鼠标位置是否有效", func() bool {
		return imgui.IsMousePosValid()
	}, true)
	engine.RegisterMethod("imgui.getMousePos", "获取鼠标位置", func() (float32, float32) {
		pos := imgui.MousePos()
		return pos.X, pos.Y
	}, true)
	engine.RegisterMethod("imgui.getMousePosOnOpeningCurrentPopup", "获取打开当前弹出窗口时的鼠标位置", func() (float32, float32) {
		pos := imgui.MousePosOnOpeningCurrentPopup()
		return pos.X, pos.Y
	}, true)
	engine.RegisterMethod("imgui.getMouseDragDelta", "获取鼠标拖拽增量", func() (float32, float32) {
		delta := imgui.MouseDragDeltaV(0, -1)
		return delta.X, delta.Y
	}, true)
	engine.RegisterMethod("imgui.resetMouseDragDelta", "重置鼠标拖拽增量", func() {
		imgui.ResetMouseDragDelta()
	}, true)
	engine.RegisterMethod("imgui.getMouseCursor", "获取鼠标光标", func() int32 {
		return int32(imgui.CurrentMouseCursor())
	}, true)
	engine.RegisterMethod("imgui.setMouseCursor", "设置鼠标光标", func(cursor_type int32) {
		imgui.SetMouseCursor(imgui.MouseCursor(cursor_type))
	}, true)
	engine.RegisterMethod("imgui.getClipboardText", "获取剪贴板文本", func() string {
		return imgui.ClipboardText()
	}, true)
	engine.RegisterMethod("imgui.setClipboardText", "设置剪贴板文本", func(text string) {
		imgui.SetClipboardText(text)
	}, true)
	engine.RegisterMethod("imgui.getTime", "获取时间", func() float64 {
		return imgui.Time()
	}, true)
	engine.RegisterMethod("imgui.getFrameCount", "获取帧数", func() int32 {
		return imgui.FrameCount()
	}, true)
	engine.RegisterMethod("imgui.getStyleColorVec4", "获取样式颜色Vec4", func(idx int32) (float32, float32, float32, float32) {
		col := imgui.StyleColorVec4(imgui.Col(idx))
		return col.X, col.Y, col.Z, col.W
	}, true)
	engine.RegisterMethod("imgui.getColorU32", "获取颜色U32", func(idx int32, alpha_mul float32) uint32 {
		return imgui.ColorU32ColV(imgui.Col(idx), alpha_mul)
	}, true)
	engine.RegisterMethod("imgui.getColorU32FromRGBA", "从RGBA获取颜色U32", func(r, g, b, a float32) uint32 {
		return imgui.ColorU32Vec4(imgui.Vec4{X: r, Y: g, Z: b, W: a})
	}, true)
	engine.RegisterMethod("imgui.getColorU32FromU32", "从U32获取颜色U32", func(col uint32, alpha_mul float32) uint32 {
		return imgui.ColorU32U32V(col, alpha_mul)
	}, true)
	engine.RegisterMethod("imgui.getFont", "获取字体", func() *imgui.Font {
		return imgui.CurrentFont()
	}, true)
	engine.RegisterMethod("imgui.getFontSize", "获取字体大小", func() float32 {
		return imgui.FontSize()
	}, true)
	engine.RegisterMethod("imgui.getFontTexUvWhitePixel", "获取字体纹理UV白色像素", func() (float32, float32) {
		uv := imgui.FontTexUvWhitePixel()
		return uv.X, uv.Y
	}, true)
	engine.RegisterMethod("imgui.getColorConvertRGBtoHSV", "RGB转HSV", func(r, g, b float32) (float32, float32, float32) {
		var h, s, v float32
		imgui.ColorConvertRGBtoHSV(r, g, b, &h, &s, &v)
		return h, s, v
	}, true)
	engine.RegisterMethod("imgui.getColorConvertHSVtoRGB", "HSV转RGB", func(h, s, v float32) (float32, float32, float32) {
		var r, g, b float32
		imgui.ColorConvertHSVtoRGB(h, s, v, &r, &g, &b)
		return r, g, b
	}, true)
	engine.RegisterMethod("imgui.getWindowDrawList", "获取窗口绘制列表", func() *imgui.DrawList {
		return imgui.WindowDrawList()
	}, true)
	engine.RegisterMethod("imgui.getBackgroundDrawList", "获取背景绘制列表", func() *imgui.DrawList {
		return imgui.BackgroundDrawListV(nil)
	}, true)
	engine.RegisterMethod("imgui.getForegroundDrawList", "获取前景绘制列表", func() *imgui.DrawList {
		return imgui.ForegroundDrawListViewportPtrV(nil)
	}, true)
	engine.RegisterMethod("imgui.getMainViewport", "获取主视口", func() *imgui.Viewport {
		return imgui.MainViewport()
	}, true)
	engine.RegisterMethod("imgui.findViewportByID", "通过ID查找视口", func(id uint32) *imgui.Viewport {
		return imgui.FindViewportByID(imgui.ID(id))
	}, true)
	engine.RegisterMethod("imgui.findViewportByPlatformHandle", "通过平台句柄查找视口", func(handle uintptr) *imgui.Viewport {
		return imgui.FindViewportByPlatformHandle(handle)
	}, true)
	engine.RegisterMethod("imgui.getPlatformIO", "获取平台IO", func() *imgui.PlatformIO {
		return imgui.CurrentPlatformIO()
	}, true)
	engine.RegisterMethod("imgui.getIO", "获取IO", func() *imgui.IO {
		return imgui.CurrentIO()
	}, true)
	engine.RegisterMethod("imgui.getStyle", "获取样式", func() *imgui.Style {
		return imgui.CurrentStyle()
	}, true)
	engine.RegisterMethod("imgui.getDragDropPayload", "获取拖放载荷", func() *imgui.Payload {
		return imgui.DragDropPayload()
	}, true)
	engine.RegisterMethod("imgui.beginTooltip", "开始工具提示", func() {
		imgui.BeginTooltip()
	}, true)
	engine.RegisterMethod("imgui.endTooltip", "结束工具提示", func() {
		imgui.EndTooltip()
	}, true)
	engine.RegisterMethod("imgui.setTooltip", "设置工具提示", func(text string) {
		imgui.SetTooltip(text)
	}, true)
	engine.RegisterMethod("imgui.beginItemTooltip", "开始项工具提示", func() bool {
		return imgui.BeginItemTooltip()
	}, true)
	engine.RegisterMethod("imgui.beginPopup", "开始弹出窗口", func(str_id string) bool {
		return imgui.BeginPopupV(str_id, 0)
	}, true)
	engine.RegisterMethod("imgui.beginPopupWithFlags", "开始弹出窗口（带标志）", func(str_id string, flags uint32) bool {
		return imgui.BeginPopupV(str_id, imgui.WindowFlags(flags))
	}, true)
	engine.RegisterMethod("imgui.endPopup", "结束弹出窗口", func() {
		imgui.EndPopup()
	}, true)
	engine.RegisterMethod("imgui.openPopup", "打开弹出窗口", func(str_id string) {
		imgui.OpenPopupStr(str_id)
	}, true)
	engine.RegisterMethod("imgui.openPopupWithFlags", "打开弹出窗口（带标志）", func(str_id string, flags uint32) {
		imgui.OpenPopupStr(str_id)
	}, true)
	engine.RegisterMethod("imgui.closeCurrentPopup", "关闭当前弹出窗口", func() {
		imgui.CloseCurrentPopup()
	}, true)
	engine.RegisterMethod("imgui.beginPopupModal", "开始模态弹出窗口", func(name string) bool {
		return imgui.BeginPopupModalV(name, nil, 0)
	}, true)
	engine.RegisterMethod("imgui.beginPopupModalWithOpen", "开始模态弹出窗口（带打开状态）", func(name string, open *bool) bool {
		return imgui.BeginPopupModalV(name, open, 0)
	}, true)
	engine.RegisterMethod("imgui.beginPopupModalWithFlags", "开始模态弹出窗口（带标志）", func(name string, flags uint32) bool {
		return imgui.BeginPopupModalV(name, nil, imgui.WindowFlags(flags))
	}, true)
	engine.RegisterMethod("imgui.beginPopupContextItem", "开始项上下文弹出窗口", func() bool {
		return imgui.BeginPopupContextItemV("", 0)
	}, true)
	engine.RegisterMethod("imgui.beginPopupContextWindow", "开始窗口上下文弹出窗口", func() bool {
		return imgui.BeginPopupContextWindowV("", 0)
	}, true)
	engine.RegisterMethod("imgui.beginPopupContextVoid", "开始空白上下文弹出窗口", func() bool {
		return imgui.BeginPopupContextVoidV("", 0)
	}, true)
	engine.RegisterMethod("imgui.beginColumns", "开始列", func(str_id string, count int32) {
		imgui.ColumnsV(count, str_id, true)
	}, true)
	engine.RegisterMethod("imgui.getColumnIndex", "获取列索引", func() int32 {
		return imgui.ColumnIndex()
	}, true)
	engine.RegisterMethod("imgui.getColumnCount", "获取列数", func() int32 {
		return imgui.ColumnsCount()
	}, true)
	engine.RegisterMethod("imgui.getColumnOffset", "获取列偏移", func(column_index int32) float32 {
		return imgui.ColumnOffsetV(column_index)
	}, true)
	engine.RegisterMethod("imgui.setColumnOffset", "设置列偏移", func(column_index int32, offset_x float32) {
		imgui.SetColumnOffset(column_index, offset_x)
	}, true)
	engine.RegisterMethod("imgui.getColumnWidth", "获取列宽度", func(column_index int32) float32 {
		return imgui.ColumnWidthV(column_index)
	}, true)
	engine.RegisterMethod("imgui.setColumnWidth", "设置列宽度", func(column_index int32, width float32) {
		imgui.SetColumnWidth(column_index, width)
	}, true)
	engine.RegisterMethod("imgui.pushIDStr", "推入ID（字符串）", func(str_id string) {
		imgui.PushIDStr(str_id)
	}, true)
	engine.RegisterMethod("imgui.pushIDInt", "推入ID（整数）", func(int_id int32) {
		imgui.PushIDInt(int_id)
	}, true)
	engine.RegisterMethod("imgui.pushIDPtr", "推入ID（指针）", func(ptr_id uintptr) {
		imgui.PushIDPtr(ptr_id)
	}, true)
	engine.RegisterMethod("imgui.popID", "弹出ID", func() {
		imgui.PopID()
	}, true)
	engine.RegisterMethod("imgui.getID", "获取ID", func(str_id string) uint32 {
		return uint32(imgui.IDStr(str_id))
	}, true)
	engine.RegisterMethod("imgui.getIDFromInt", "从整数获取ID", func(int_id int32) uint32 {
		return uint32(imgui.IDInt(int_id))
	}, true)
	engine.RegisterMethod("imgui.getIDFromPtr", "从指针获取ID", func(ptr_id uintptr) uint32 {
		return uint32(imgui.IDPtr(ptr_id))
	}, true)
	engine.RegisterMethod("imgui.text", "文本", func(text string) {
		imgui.Text(text)
	}, true)
	engine.RegisterMethod("imgui.textColored", "彩色文本", func(text string, r, g, b, a float32) {
		imgui.TextColored(imgui.Vec4{X: r, Y: g, Z: b, W: a}, text)
	}, true)
	engine.RegisterMethod("imgui.textDisabled", "禁用文本", func(text string) {
		imgui.TextDisabled(text)
	}, true)
	engine.RegisterMethod("imgui.textWrapped", "换行文本", func(text string) {
		imgui.TextWrapped(text)
	}, true)
	engine.RegisterMethod("imgui.textUnformatted", "无格式文本", func(text string) {
		imgui.TextUnformatted(text)
	}, true)
	engine.RegisterMethod("imgui.labelText", "标签文本", func(label, text string) {
		imgui.LabelText(label, text)
	}, true)
	engine.RegisterMethod("imgui.bullet", "项目符号", func() {
		imgui.Bullet()
	}, true)
	engine.RegisterMethod("imgui.bulletText", "项目符号文本", func(text string) {
		imgui.BulletText(text)
	}, true)
	engine.RegisterMethod("imgui.separator", "分隔符", func() {
		imgui.Separator()
	}, true)
	engine.RegisterMethod("imgui.separatorText", "分隔符文本", func(text string) {
		imgui.SeparatorText(text)
	}, true)
	engine.RegisterMethod("imgui.sameLine", "同一行", func() {
		imgui.SameLine()
	}, true)
	engine.RegisterMethod("imgui.newLine", "新行", func() {
		imgui.NewLine()
	}, true)
	engine.RegisterMethod("imgui.spacing", "间距", func() {
		imgui.Spacing()
	}, true)
	engine.RegisterMethod("imgui.dummy", "虚拟元素", func(w, h float32) {
		imgui.Dummy(imgui.Vec2{X: w, Y: h})
	}, true)
	engine.RegisterMethod("imgui.beginGroup", "开始组", func() {
		imgui.BeginGroup()
	}, true)
	engine.RegisterMethod("imgui.endGroup", "结束组", func() {
		imgui.EndGroup()
	}, true)
	engine.RegisterMethod("imgui.getCursorPos", "获取光标位置", func() (float32, float32) {
		pos := imgui.CursorPos()
		return pos.X, pos.Y
	}, true)
	engine.RegisterMethod("imgui.getCursorPosX", "获取光标X位置", func() float32 {
		return imgui.CursorPosX()
	}, true)
	engine.RegisterMethod("imgui.getCursorPosY", "获取光标Y位置", func() float32 {
		return imgui.CursorPosY()
	}, true)
	engine.RegisterMethod("imgui.setCursorPos", "设置光标位置", func(x, y float32) {
		imgui.SetCursorPos(imgui.Vec2{X: x, Y: y})
	}, true)
	engine.RegisterMethod("imgui.setCursorPosX", "设置光标X位置", func(x float32) {
		imgui.SetCursorPosX(x)
	}, true)
	engine.RegisterMethod("imgui.setCursorPosY", "设置光标Y位置", func(y float32) {
		imgui.SetCursorPosY(y)
	}, true)
	engine.RegisterMethod("imgui.getCursorStartPos", "获取光标起始位置", func() (float32, float32) {
		pos := imgui.CursorStartPos()
		return pos.X, pos.Y
	}, true)
	engine.RegisterMethod("imgui.getCursorScreenPos", "获取屏幕光标位置", func() (float32, float32) {
		pos := imgui.CursorScreenPos()
		return pos.X, pos.Y
	}, true)
	engine.RegisterMethod("imgui.setCursorScreenPos", "设置屏幕光标位置", func(x, y float32) {
		imgui.SetCursorScreenPos(imgui.Vec2{X: x, Y: y})
	}, true)
	engine.RegisterMethod("imgui.getTextLineHeight", "获取文本行高度", func() float32 {
		return imgui.TextLineHeight()
	}, true)
	engine.RegisterMethod("imgui.getTextLineHeightWithSpacing", "获取带间距的文本行高度", func() float32 {
		return imgui.TextLineHeightWithSpacing()
	}, true)
	engine.RegisterMethod("imgui.getFrameHeight", "获取框架高度", func() float32 {
		return imgui.FrameHeight()
	}, true)
	engine.RegisterMethod("imgui.getFrameHeightWithSpacing", "获取带间距的框架高度", func() float32 {
		return imgui.FrameHeightWithSpacing()
	}, true)
	engine.RegisterMethod("imgui.getContentRegionMax", "获取内容区域最大值", func() (float32, float32) {
		max := imgui.ContentRegionAvail()
		return max.X, max.Y
	}, true)
	engine.RegisterMethod("imgui.getContentRegionAvail", "获取可用内容区域", func() (float32, float32) {
		avail := imgui.ContentRegionAvail()
		return avail.X, avail.Y
	}, true)
	engine.RegisterMethod("imgui.popStyleColor", "弹出样式颜色", func() {
		imgui.PopStyleColor()
	}, true)
	engine.RegisterMethod("imgui.getStyleColor", "获取样式颜色", func(idx int32) (float32, float32, float32, float32) {
		col := imgui.StyleColorVec4(imgui.Col(idx))
		return col.X, col.Y, col.Z, col.W
	}, true)
	engine.RegisterMethod("imgui.getStyleColorName", "获取样式颜色名称", func(idx int32) string {
		return imgui.StyleColorName(imgui.Col(idx))
	}, true)
	engine.RegisterMethod("imgui.getStyleVarInfo", "获取样式变量信息", func(idx int32) (uint32, uint32, uint32) {
		info := imgui.InternalStyleVarInfo(imgui.StyleVar(idx))
		return info.Count(), uint32(info.DataType()), info.Offset()
	}, true)
	engine.RegisterMethod("imgui.beginMainMenuBar", "开始主菜单栏", func() bool {
		return imgui.BeginMainMenuBar()
	}, true)
	engine.RegisterMethod("imgui.endMainMenuBar", "结束主菜单栏", func() {
		imgui.EndMainMenuBar()
	}, true)
	engine.RegisterMethod("imgui.beginMenuBar", "开始菜单栏", func() bool {
		return imgui.BeginMenuBar()
	}, true)
	engine.RegisterMethod("imgui.endMenuBar", "结束菜单栏", func() {
		imgui.EndMenuBar()
	}, true)
	engine.RegisterMethod("imgui.beginMenu", "开始菜单", func(label string) bool {
		return imgui.BeginMenuV(label, true)
	}, true)
	engine.RegisterMethod("imgui.beginMenuDisabled", "开始禁用菜单", func(label string) bool {
		return imgui.BeginMenuV(label, false)
	}, true)
	engine.RegisterMethod("imgui.endMenu", "结束菜单", func() {
		imgui.EndMenu()
	}, true)
	engine.RegisterMethod("imgui.beginTabBar", "开始标签栏", func(str_id string) bool {
		return imgui.BeginTabBarV(str_id, 0)
	}, true)
	engine.RegisterMethod("imgui.beginTabBarWithFlags", "开始标签栏（带标志）", func(str_id string, flags uint32) bool {
		return imgui.BeginTabBarV(str_id, imgui.TabBarFlags(flags))
	}, true)
	engine.RegisterMethod("imgui.endTabBar", "结束标签栏", func() {
		imgui.EndTabBar()
	}, true)
	engine.RegisterMethod("imgui.beginTabItem", "开始标签项", func(label string) bool {
		return imgui.BeginTabItemV(label, nil, 0)
	}, true)
	engine.RegisterMethod("imgui.beginTabItemWithOpen", "开始标签项（带打开状态）", func(label string, open *bool) bool {
		return imgui.BeginTabItemV(label, open, 0)
	}, true)
	engine.RegisterMethod("imgui.beginTabItemWithFlags", "开始标签项（带标志）", func(label string, flags uint32) bool {
		return imgui.BeginTabItemV(label, nil, imgui.TabItemFlags(flags))
	}, true)
	engine.RegisterMethod("imgui.endTabItem", "结束标签项", func() {
		imgui.EndTabItem()
	}, true)
	engine.RegisterMethod("imgui.setTabItemClosed", "设置标签项关闭", func(tab_or_docked_window_label string) {
		imgui.SetTabItemClosed(tab_or_docked_window_label)
	}, true)
	engine.RegisterMethod("imgui.isItemHovered", "项是否悬停", func() bool {
		return imgui.IsItemHovered()
	}, true)
	engine.RegisterMethod("imgui.isItemActive", "项是否激活", func() bool {
		return imgui.IsItemActive()
	}, true)
	engine.RegisterMethod("imgui.isItemFocused", "项是否聚焦", func() bool {
		return imgui.IsItemFocused()
	}, true)
	engine.RegisterMethod("imgui.isItemClicked", "项是否被点击", func() bool {
		return imgui.IsItemClicked()
	}, true)
	engine.RegisterMethod("imgui.isItemVisible", "项是否可见", func() bool {
		return imgui.IsItemVisible()
	}, true)
	engine.RegisterMethod("imgui.isItemEdited", "项是否被编辑", func() bool {
		return imgui.IsItemEdited()
	}, true)
	engine.RegisterMethod("imgui.isItemActivated", "项是否被激活", func() bool {
		return imgui.IsItemActivated()
	}, true)
	engine.RegisterMethod("imgui.isItemDeactivated", "项是否被停用", func() bool {
		return imgui.IsItemDeactivated()
	}, true)
	engine.RegisterMethod("imgui.isItemDeactivatedAfterEdit", "项是否在编辑后被停用", func() bool {
		return imgui.IsItemDeactivatedAfterEdit()
	}, true)
	engine.RegisterMethod("imgui.isItemToggledOpen", "项是否切换打开状态", func() bool {
		return imgui.IsItemToggledOpen()
	}, true)
	engine.RegisterMethod("imgui.isAnyItemHovered", "是否有任何项悬停", func() bool {
		return imgui.IsAnyItemHovered()
	}, true)
	engine.RegisterMethod("imgui.isAnyItemActive", "是否有任何项激活", func() bool {
		return imgui.IsAnyItemActive()
	}, true)
	engine.RegisterMethod("imgui.isAnyItemFocused", "是否有任何项聚焦", func() bool {
		return imgui.IsAnyItemFocused()
	}, true)
	engine.RegisterMethod("imgui.getItemRectMin", "获取项矩形最小值", func() (float32, float32) {
		min := imgui.ItemRectMin()
		return min.X, min.Y
	}, true)
	engine.RegisterMethod("imgui.getItemRectMax", "获取项矩形最大值", func() (float32, float32) {
		max := imgui.ItemRectMax()
		return max.X, max.Y
	}, true)
	engine.RegisterMethod("imgui.getItemRectSize", "获取项矩形大小", func() (float32, float32) {
		size := imgui.ItemRectSize()
		return size.X, size.Y
	}, true)
	engine.RegisterMethod("imgui.setItemDefaultFocus", "设置项默认焦点", func() {
		imgui.SetItemDefaultFocus()
	}, true)
	engine.RegisterMethod("imgui.setKeyboardFocusHere", "设置键盘焦点到此处", func() {
		imgui.SetKeyboardFocusHere()
	}, true)
	engine.RegisterMethod("imgui.isMouseHoveringRect", "鼠标是否悬停在矩形上", func(min_x, min_y, max_x, max_y float32) bool {
		min := imgui.Vec2{X: min_x, Y: min_y}
		max := imgui.Vec2{X: max_x, Y: max_y}
		return imgui.IsMouseHoveringRect(min, max)
	}, true)
	engine.RegisterMethod("imgui.isMousePosValid", "鼠标位置是否有效", func() bool {
		return imgui.IsMousePosValid()
	}, true)
	engine.RegisterMethod("imgui.getMousePos", "获取鼠标位置", func() (float32, float32) {
		pos := imgui.MousePos()
		return pos.X, pos.Y
	}, true)
	engine.RegisterMethod("imgui.getMousePosOnOpeningCurrentPopup", "获取打开当前弹出窗口时的鼠标位置", func() (float32, float32) {
		pos := imgui.MousePosOnOpeningCurrentPopup()
		return pos.X, pos.Y
	}, true)
	engine.RegisterMethod("imgui.getMouseDragDelta", "获取鼠标拖拽增量", func() (float32, float32) {
		delta := imgui.MouseDragDeltaV(0, -1)
		return delta.X, delta.Y
	}, true)
	engine.RegisterMethod("imgui.resetMouseDragDelta", "重置鼠标拖拽增量", func() {
		imgui.ResetMouseDragDelta()
	}, true)
	engine.RegisterMethod("imgui.getMouseCursor", "获取鼠标光标", func() int32 {
		return int32(imgui.CurrentMouseCursor())
	}, true)
	engine.RegisterMethod("imgui.setMouseCursor", "设置鼠标光标", func(cursor_type int32) {
		imgui.SetMouseCursor(imgui.MouseCursor(cursor_type))
	}, true)
	engine.RegisterMethod("imgui.getClipboardText", "获取剪贴板文本", func() string {
		return imgui.ClipboardText()
	}, true)
	engine.RegisterMethod("imgui.setClipboardText", "设置剪贴板文本", func(text string) {
		imgui.SetClipboardText(text)
	}, true)
	engine.RegisterMethod("imgui.getTime", "获取时间", func() float64 {
		return imgui.Time()
	}, true)
	engine.RegisterMethod("imgui.getFrameCount", "获取帧数", func() int32 {
		return imgui.FrameCount()
	}, true)
	engine.RegisterMethod("imgui.getStyleColorVec4", "获取样式颜色Vec4", func(idx int32) (float32, float32, float32, float32) {
		col := imgui.StyleColorVec4(imgui.Col(idx))
		return col.X, col.Y, col.Z, col.W
	}, true)
	engine.RegisterMethod("imgui.getColorU32", "获取颜色U32", func(idx int32, alpha_mul float32) uint32 {
		return imgui.ColorU32ColV(imgui.Col(idx), alpha_mul)
	}, true)
	engine.RegisterMethod("imgui.getColorU32FromRGBA", "从RGBA获取颜色U32", func(r, g, b, a float32) uint32 {
		return imgui.ColorU32Vec4(imgui.Vec4{X: r, Y: g, Z: b, W: a})
	}, true)
	engine.RegisterMethod("imgui.getColorU32FromU32", "从U32获取颜色U32", func(col uint32, alpha_mul float32) uint32 {
		return imgui.ColorU32U32V(col, alpha_mul)
	}, true)
	engine.RegisterMethod("imgui.getFont", "获取字体", func() *imgui.Font {
		return imgui.CurrentFont()
	}, true)
	engine.RegisterMethod("imgui.getFontSize", "获取字体大小", func() float32 {
		return imgui.FontSize()
	}, true)
	engine.RegisterMethod("imgui.getFontTexUvWhitePixel", "获取字体纹理UV白色像素", func() (float32, float32) {
		uv := imgui.FontTexUvWhitePixel()
		return uv.X, uv.Y
	}, true)
	engine.RegisterMethod("imgui.getColorConvertRGBtoHSV", "RGB转HSV", func(r, g, b float32) (float32, float32, float32) {
		var h, s, v float32
		imgui.ColorConvertRGBtoHSV(r, g, b, &h, &s, &v)
		return h, s, v
	}, true)
	engine.RegisterMethod("imgui.getColorConvertHSVtoRGB", "HSV转RGB", func(h, s, v float32) (float32, float32, float32) {
		var r, g, b float32
		imgui.ColorConvertHSVtoRGB(h, s, v, &r, &g, &b)
		return r, g, b
	}, true)
	engine.RegisterMethod("imgui.getWindowDrawList", "获取窗口绘制列表", func() *imgui.DrawList {
		return imgui.WindowDrawList()
	}, true)
	engine.RegisterMethod("imgui.getBackgroundDrawList", "获取背景绘制列表", func() *imgui.DrawList {
		return imgui.BackgroundDrawListV(nil)
	}, true)
	engine.RegisterMethod("imgui.getForegroundDrawList", "获取前景绘制列表", func() *imgui.DrawList {
		return imgui.ForegroundDrawListViewportPtrV(nil)
	}, true)
	engine.RegisterMethod("imgui.getMainViewport", "获取主视口", func() *imgui.Viewport {
		return imgui.MainViewport()
	}, true)
	engine.RegisterMethod("imgui.findViewportByID", "通过ID查找视口", func(id uint32) *imgui.Viewport {
		return imgui.FindViewportByID(imgui.ID(id))
	}, true)
	engine.RegisterMethod("imgui.findViewportByPlatformHandle", "通过平台句柄查找视口", func(handle uintptr) *imgui.Viewport {
		return imgui.FindViewportByPlatformHandle(handle)
	}, true)
	engine.RegisterMethod("imgui.getPlatformIO", "获取平台IO", func() *imgui.PlatformIO {
		return imgui.CurrentPlatformIO()
	}, true)
	engine.RegisterMethod("imgui.getIO", "获取IO", func() *imgui.IO {
		return imgui.CurrentIO()
	}, true)
	engine.RegisterMethod("imgui.getStyle", "获取样式", func() *imgui.Style {
		return imgui.CurrentStyle()
	}, true)
	engine.RegisterMethod("imgui.getDragDropPayload", "获取拖放载荷", func() *imgui.Payload {
		return imgui.DragDropPayload()
	}, true)
	engine.RegisterMethod("imgui.beginTooltip", "开始工具提示", func() {
		imgui.BeginTooltip()
	}, true)
	engine.RegisterMethod("imgui.endTooltip", "结束工具提示", func() {
		imgui.EndTooltip()
	}, true)
	engine.RegisterMethod("imgui.setTooltip", "设置工具提示", func(text string) {
		imgui.SetTooltip(text)
	}, true)
	engine.RegisterMethod("imgui.beginItemTooltip", "开始项工具提示", func() bool {
		return imgui.BeginItemTooltip()
	}, true)
	engine.RegisterMethod("imgui.beginPopup", "开始弹出窗口", func(str_id string) bool {
		return imgui.BeginPopupV(str_id, 0)
	}, true)
	engine.RegisterMethod("imgui.beginPopupWithFlags", "开始弹出窗口（带标志）", func(str_id string, flags uint32) bool {
		return imgui.BeginPopupV(str_id, imgui.WindowFlags(flags))
	}, true)
	engine.RegisterMethod("imgui.endPopup", "结束弹出窗口", func() {
		imgui.EndPopup()
	}, true)
	engine.RegisterMethod("imgui.openPopup", "打开弹出窗口", func(str_id string) {
		imgui.OpenPopupStr(str_id)
	}, true)
	engine.RegisterMethod("imgui.openPopupWithFlags", "打开弹出窗口（带标志）", func(str_id string, flags uint32) {
		imgui.OpenPopupStr(str_id)
	}, true)
	engine.RegisterMethod("imgui.closeCurrentPopup", "关闭当前弹出窗口", func() {
		imgui.CloseCurrentPopup()
	}, true)
	engine.RegisterMethod("imgui.beginPopupModal", "开始模态弹出窗口", func(name string) bool {
		return imgui.BeginPopupModalV(name, nil, 0)
	}, true)
	engine.RegisterMethod("imgui.beginPopupModalWithOpen", "开始模态弹出窗口（带打开状态）", func(name string, open *bool) bool {
		return imgui.BeginPopupModalV(name, open, 0)
	}, true)
	engine.RegisterMethod("imgui.beginPopupModalWithFlags", "开始模态弹出窗口（带标志）", func(name string, flags uint32) bool {
		return imgui.BeginPopupModalV(name, nil, imgui.WindowFlags(flags))
	}, true)
	engine.RegisterMethod("imgui.beginPopupContextItem", "开始项上下文弹出窗口", func() bool {
		return imgui.BeginPopupContextItemV("", 0)
	}, true)
	engine.RegisterMethod("imgui.beginPopupContextWindow", "开始窗口上下文弹出窗口", func() bool {
		return imgui.BeginPopupContextWindowV("", 0)
	}, true)
	engine.RegisterMethod("imgui.beginPopupContextVoid", "开始空白上下文弹出窗口", func() bool {
		return imgui.BeginPopupContextVoidV("", 0)
	}, true)
	engine.RegisterMethod("imgui.beginColumns", "开始列", func(str_id string, count int32) {
		imgui.ColumnsV(count, str_id, true)
	}, true)
	engine.RegisterMethod("imgui.getColumnIndex", "获取列索引", func() int32 {
		return imgui.ColumnIndex()
	}, true)
	engine.RegisterMethod("imgui.getColumnCount", "获取列数", func() int32 {
		return imgui.ColumnsCount()
	}, true)
	engine.RegisterMethod("imgui.getColumnOffset", "获取列偏移", func(column_index int32) float32 {
		return imgui.ColumnOffsetV(column_index)
	}, true)
	engine.RegisterMethod("imgui.setColumnOffset", "设置列偏移", func(column_index int32, offset_x float32) {
		imgui.SetColumnOffset(column_index, offset_x)
	}, true)
	engine.RegisterMethod("imgui.getColumnWidth", "获取列宽度", func(column_index int32) float32 {
		return imgui.ColumnWidthV(column_index)
	}, true)
	engine.RegisterMethod("imgui.setColumnWidth", "设置列宽度", func(column_index int32, width float32) {
		imgui.SetColumnWidth(column_index, width)
	}, true)
	engine.RegisterMethod("imgui.pushIDStr", "推入ID（字符串）", func(str_id string) {
		imgui.PushIDStr(str_id)
	}, true)
	engine.RegisterMethod("imgui.pushIDInt", "推入ID（整数）", func(int_id int32) {
		imgui.PushIDInt(int_id)
	}, true)
	engine.RegisterMethod("imgui.pushIDPtr", "推入ID（指针）", func(ptr_id uintptr) {
		imgui.PushIDPtr(ptr_id)
	}, true)
	engine.RegisterMethod("imgui.popID", "弹出ID", func() {
		imgui.PopID()
	}, true)
	engine.RegisterMethod("imgui.getID", "获取ID", func(str_id string) uint32 {
		return uint32(imgui.IDStr(str_id))
	}, true)
	engine.RegisterMethod("imgui.getIDFromInt", "从整数获取ID", func(int_id int32) uint32 {
		return uint32(imgui.IDInt(int_id))
	}, true)
	engine.RegisterMethod("imgui.getIDFromPtr", "从指针获取ID", func(ptr_id uintptr) uint32 {
		return uint32(imgui.IDPtr(ptr_id))
	}, true)
	engine.RegisterMethod("imgui.text", "文本", func(text string) {
		imgui.Text(text)
	}, true)
	engine.RegisterMethod("imgui.textColored", "彩色文本", func(text string, r, g, b, a float32) {
		imgui.TextColored(imgui.Vec4{X: r, Y: g, Z: b, W: a}, text)
	}, true)
	engine.RegisterMethod("imgui.textDisabled", "禁用文本", func(text string) {
		imgui.TextDisabled(text)
	}, true)
	engine.RegisterMethod("imgui.textWrapped", "换行文本", func(text string) {
		imgui.TextWrapped(text)
	}, true)
	engine.RegisterMethod("imgui.textUnformatted", "无格式文本", func(text string) {
		imgui.TextUnformatted(text)
	}, true)
	engine.RegisterMethod("imgui.labelText", "标签文本", func(label, text string) {
		imgui.LabelText(label, text)
	}, true)
	engine.RegisterMethod("imgui.bullet", "项目符号", func() {
		imgui.Bullet()
	}, true)
	engine.RegisterMethod("imgui.bulletText", "项目符号文本", func(text string) {
		imgui.BulletText(text)
	}, true)
	engine.RegisterMethod("imgui.separator", "分隔符", func() {
		imgui.Separator()
	}, true)
	engine.RegisterMethod("imgui.separatorText", "分隔符文本", func(text string) {
		imgui.SeparatorText(text)
	}, true)
	engine.RegisterMethod("imgui.sameLine", "同一行", func() {
		imgui.SameLine()
	}, true)

	engine.RegisterMethod("imgui.image", "绘制图像", func(texRef *imgui.TextureRef, w, h float32) {
		size := imgui.Vec2{X: w, Y: h}
		imgui.ImageV(*texRef, size, imgui.Vec2{X: 0, Y: 0}, imgui.Vec2{X: 1, Y: 1})
	}, true)
	engine.RegisterMethod("imgui.imageWithUV", "绘制图像（带UV）", func(texRef *imgui.TextureRef, w, h, uv0X, uv0Y, uv1X, uv1Y float32) {
		size := imgui.Vec2{X: w, Y: h}
		uv0 := imgui.Vec2{X: uv0X, Y: uv0Y}
		uv1 := imgui.Vec2{X: uv1X, Y: uv1Y}
		imgui.ImageV(*texRef, size, uv0, uv1)
	}, true)
	engine.RegisterMethod("imgui.imageButton", "图像按钮", func(str_id string, texRef *imgui.TextureRef, w, h float32) bool {
		size := imgui.Vec2{X: w, Y: h}
		return imgui.ImageButtonV(str_id, *texRef, size, imgui.Vec2{X: 0, Y: 0}, imgui.Vec2{X: 1, Y: 1}, imgui.Vec4{X: 0, Y: 0, Z: 0, W: 0}, imgui.Vec4{X: 1, Y: 1, Z: 1, W: 1})
	}, true)
	engine.RegisterMethod("imgui.imageButtonWithUV", "图像按钮（带UV）", func(str_id string, texRef *imgui.TextureRef, w, h, uv0X, uv0Y, uv1X, uv1Y, bgR, bgG, bgB, bgA, tintR, tintG, tintB, tintA float32) bool {
		size := imgui.Vec2{X: w, Y: h}
		uv0 := imgui.Vec2{X: uv0X, Y: uv0Y}
		uv1 := imgui.Vec2{X: uv1X, Y: uv1Y}
		bgCol := imgui.Vec4{X: bgR, Y: bgG, Z: bgB, W: bgA}
		tintCol := imgui.Vec4{X: tintR, Y: tintG, Z: tintB, W: tintA}
		return imgui.ImageButtonV(str_id, *texRef, size, uv0, uv1, bgCol, tintCol)
	}, true)
	engine.RegisterMethod("imgui.imageWithBg", "绘制图像（带背景）", func(texRef *imgui.TextureRef, w, h, bgR, bgG, bgB, bgA, tintR, tintG, tintB, tintA float32) {
		size := imgui.Vec2{X: w, Y: h}
		bgCol := imgui.Vec4{X: bgR, Y: bgG, Z: bgB, W: bgA}
		tintCol := imgui.Vec4{X: tintR, Y: tintG, Z: tintB, W: tintA}
		imgui.ImageWithBgV(*texRef, size, imgui.Vec2{X: 0, Y: 0}, imgui.Vec2{X: 1, Y: 1}, bgCol, tintCol)
	}, true)

	engine.RegisterMethod("imgui.dockSpace", "停靠空间", func(dockspace_id uint32) uint32 {
		return uint32(imgui.DockSpaceV(imgui.ID(dockspace_id), imgui.Vec2{X: 0, Y: 0}, 0, nil))
	}, true)
	engine.RegisterMethod("imgui.dockSpaceWithSize", "停靠空间（带大小）", func(dockspace_id uint32, w, h float32) uint32 {
		size := imgui.Vec2{X: w, Y: h}
		return uint32(imgui.DockSpaceV(imgui.ID(dockspace_id), size, 0, nil))
	}, true)
	engine.RegisterMethod("imgui.dockSpaceWithFlags", "停靠空间（带标志）", func(dockspace_id uint32, w, h float32, flags uint32) uint32 {
		size := imgui.Vec2{X: w, Y: h}
		return uint32(imgui.DockSpaceV(imgui.ID(dockspace_id), size, imgui.DockNodeFlags(flags), nil))
	}, true)
	engine.RegisterMethod("imgui.dockSpaceOverViewport", "视口停靠空间", func() uint32 {
		return uint32(imgui.DockSpaceOverViewportV(0, nil, 0, nil))
	}, true)
	engine.RegisterMethod("imgui.dockSpaceOverViewportWithID", "视口停靠空间（带ID）", func(dockspace_id uint32) uint32 {
		return uint32(imgui.DockSpaceOverViewportV(imgui.ID(dockspace_id), nil, 0, nil))
	}, true)
	engine.RegisterMethod("imgui.setNextWindowDockID", "设置下一个窗口停靠ID", func(dock_id uint32) {
		imgui.SetNextWindowDockIDV(imgui.ID(dock_id), 0)
	}, true)
	engine.RegisterMethod("imgui.setNextWindowDockIDWithCond", "设置下一个窗口停靠ID（带条件）", func(dock_id uint32, cond uint32) {
		imgui.SetNextWindowDockIDV(imgui.ID(dock_id), imgui.Cond(cond))
	}, true)
	engine.RegisterMethod("imgui.setNextWindowClass", "设置下一个窗口类", func() {
		imgui.SetNextWindowClass(nil)
	}, true)

	engine.RegisterMethod("imgui.listBox", "列表框", func(label string, current_item int32, items []string) (bool, int32) {
		changed := imgui.ListBoxStrarr(label, &current_item, items, int32(len(items)))
		return changed, current_item
	}, true)
	engine.RegisterMethod("imgui.listBoxWithHeight", "列表框（带高度）", func(label string, current_item int32, items []string, height_in_items int32) (bool, int32) {
		changed := imgui.ListBoxStrarrV(label, &current_item, items, int32(len(items)), height_in_items)
		return changed, current_item
	}, true)

	engine.RegisterMethod("imgui.tableSetupScrollFreeze", "设置表格滚动冻结", func(cols, rows int32) {
		imgui.TableSetupScrollFreeze(cols, rows)
	}, true)
	engine.RegisterMethod("imgui.tableGetSortSpecs", "获取表格排序规范", func() uint32 {
		specs := imgui.TableGetSortSpecs()
		if specs == nil {
			return 0
		}
		if specs.SpecsDirty() {
			return 1
		}
		return 0
	}, true)
	engine.RegisterMethod("imgui.tableGetColumnName", "获取表格列名", func() string {
		return imgui.TableGetColumnNameIntV(-1)
	}, true)
	engine.RegisterMethod("imgui.tableGetColumnNameByIndex", "获取表格列名（按索引）", func(column_n int32) string {
		return imgui.TableGetColumnNameIntV(column_n)
	}, true)
	engine.RegisterMethod("imgui.tableSetColumnEnabled", "设置表格列启用", func(column_n int32, v bool) {
		imgui.TableSetColumnEnabled(column_n, v)
	}, true)
	engine.RegisterMethod("imgui.tableSetBgColor", "设置表格背景色", func(target int32, color uint32) {
		imgui.TableSetBgColorV(imgui.TableBgTarget(target), color, -1)
	}, true)
	engine.RegisterMethod("imgui.tableSetBgColorByColumn", "设置表格背景色（按列）", func(target int32, color uint32, column_n int32) {
		imgui.TableSetBgColorV(imgui.TableBgTarget(target), color, column_n)
	}, true)

	engine.RegisterMethod("imgui.invisibleButton", "不可见按钮", func(str_id string, w, h float32) bool {
		size := imgui.Vec2{X: w, Y: h}
		return imgui.InvisibleButton(str_id, size)
	}, true)
	engine.RegisterMethod("imgui.pushTextWrapPos", "推入文本换行位置", func() {
		imgui.PushTextWrapPos()
	}, true)
	engine.RegisterMethod("imgui.pushTextWrapPosWithPos", "推入文本换行位置（带位置）", func(wrap_local_pos_x float32) {
		imgui.PushTextWrapPosV(wrap_local_pos_x)
	}, true)
	engine.RegisterMethod("imgui.popTextWrapPos", "弹出文本换行位置", func() {
		imgui.PopTextWrapPos()
	}, true)

	engine.RegisterMethod("imgui.pushItemWidth", "推入项宽度", func(item_width float32) {
		imgui.PushItemWidth(item_width)
	}, true)
	engine.RegisterMethod("imgui.popItemWidth", "弹出项宽度", func() {
		imgui.PopItemWidth()
	}, true)

	engine.RegisterMethod("imgui.pushClipRect", "推入剪裁矩形", func(minX, minY, maxX, maxY float32, intersect bool) {
		clip_rect_min := imgui.Vec2{X: minX, Y: minY}
		clip_rect_max := imgui.Vec2{X: maxX, Y: maxY}
		imgui.PushClipRect(clip_rect_min, clip_rect_max, intersect)
	}, true)
	engine.RegisterMethod("imgui.popClipRect", "弹出剪裁矩形", func() {
		imgui.PopClipRect()
	}, true)

	engine.RegisterMethod("imgui.pushStyleVarX", "推入样式变量（X）", func(idx int32, val_x float32) {
		imgui.PushStyleVarX(imgui.StyleVar(idx), val_x)
	}, true)
	engine.RegisterMethod("imgui.pushStyleVarY", "推入样式变量（Y）", func(idx int32, val_y float32) {
		imgui.PushStyleVarY(imgui.StyleVar(idx), val_y)
	}, true)
	engine.RegisterMethod("imgui.pushStyleVarFloat", "推入样式变量（浮点）", func(idx int32, val float32) {
		imgui.PushStyleVarFloat(imgui.StyleVar(idx), val)
	}, true)
	engine.RegisterMethod("imgui.pushStyleVarVec2", "推入样式变量（Vec2）", func(idx int32, valX, valY float32) {
		val := imgui.Vec2{X: valX, Y: valY}
		imgui.PushStyleVarVec2(imgui.StyleVar(idx), val)
	}, true)
	engine.RegisterMethod("imgui.popStyleVar", "弹出样式变量", func() {
		imgui.PopStyleVar()
	}, true)
	engine.RegisterMethod("imgui.popStyleVarCount", "弹出样式变量（带数量）", func(count int32) {
		imgui.PopStyleVarV(count)
	}, true)

	engine.RegisterMethod("imgui.loadIniSettingsFromDisk", "从磁盘加载INI设置", func(ini_filename string) {
		imgui.LoadIniSettingsFromDisk(ini_filename)
	}, true)
	engine.RegisterMethod("imgui.loadIniSettingsFromMemory", "从内存加载INI设置", func(ini_data string) {
		imgui.LoadIniSettingsFromMemory(ini_data)
	}, true)
	engine.RegisterMethod("imgui.saveIniSettingsToDisk", "保存INI设置到磁盘", func(ini_filename string) {
		imgui.SaveIniSettingsToDisk(ini_filename)
	}, true)
	engine.RegisterMethod("imgui.saveIniSettingsToMemory", "保存INI设置到内存", func() string {
		return imgui.SaveIniSettingsToMemory()
	}, true)

	engine.RegisterMethod("imgui.setClipboardText", "设置剪贴板文本", func(text string) {
		imgui.SetClipboardText(text)
	}, true)

	engine.RegisterMethod("imgui.pushStyleColorU32", "推入样式颜色（U32）", func(idx int32, col uint32) {
		imgui.PushStyleColorU32(imgui.Col(idx), col)
	}, true)
	engine.RegisterMethod("imgui.pushStyleColorVec4", "推入样式颜色（Vec4）", func(idx int32, r, g, b, a float32) {
		col := imgui.Vec4{X: r, Y: g, Z: b, W: a}
		imgui.PushStyleColorVec4(imgui.Col(idx), col)
	}, true)
	engine.RegisterMethod("imgui.popStyleColor", "弹出样式颜色", func() {
		imgui.PopStyleColor()
	}, true)
	engine.RegisterMethod("imgui.popStyleColorCount", "弹出样式颜色（带数量）", func(count int32) {
		imgui.PopStyleColorV(count)
	}, true)

	engine.RegisterMethod("imgui.inputTextWithHint", "带提示的文本输入", func(label, hint, text string) (bool, string) {
		changed := imgui.InputTextWithHint(label, hint, &text, 0, nil)
		return changed, text
	}, true)
	engine.RegisterMethod("imgui.inputTextMultiline", "多行文本输入", func(label, text string, w, h float32) (bool, string) {
		size := imgui.Vec2{X: w, Y: h}
		changed := imgui.InputTextMultiline(label, &text, size, 0, nil)
		return changed, text
	}, true)

	engine.RegisterMethod("imgui.treeNodeEx", "扩展树节点", func(label string, flags uint32) bool {
		return imgui.TreeNodeExStrV(label, imgui.TreeNodeFlags(flags))
	}, true)
	engine.RegisterMethod("imgui.treeNodeStr", "树节点（字符串）", func(label string) bool {
		return imgui.TreeNodeStr(label)
	}, true)
	engine.RegisterMethod("imgui.treeNodeStrStr", "树节点（双字符串）", func(str_id, fmt string) bool {
		return imgui.TreeNodeStrStr(str_id, fmt)
	}, true)
	engine.RegisterMethod("imgui.treeNodePtr", "树节点（指针）", func(ptr_id uintptr, fmt string) bool {
		return imgui.TreeNodePtr(ptr_id, fmt)
	}, true)

	engine.RegisterMethod("imgui.setKeyboardFocusHere", "设置键盘焦点到此处", func() {
		imgui.SetKeyboardFocusHere()
	}, true)
	engine.RegisterMethod("imgui.setKeyboardFocusHereWithOffset", "设置键盘焦点到此处（带偏移）", func(offset int32) {
		imgui.SetKeyboardFocusHereV(offset)
	}, true)
	engine.RegisterMethod("imgui.setItemDefaultFocus", "设置项默认焦点", func() {
		imgui.SetItemDefaultFocus()
	}, true)
	engine.RegisterMethod("imgui.setNextItemAllowOverlap", "设置下一项允许重叠", func() {
		imgui.SetNextItemAllowOverlap()
	}, true)
	engine.RegisterMethod("imgui.setNextItemOpen", "设置下一项打开状态", func(is_open bool) {
		imgui.SetNextItemOpen(is_open)
	}, true)
	engine.RegisterMethod("imgui.setNextItemOpenWithCond", "设置下一项打开状态（带条件）", func(is_open bool, cond uint32) {
		imgui.SetNextItemOpenV(is_open, imgui.Cond(cond))
	}, true)
	engine.RegisterMethod("imgui.setScrollHereX", "设置滚动到此处 X", func() {
		imgui.SetScrollHereX()
	}, true)
	engine.RegisterMethod("imgui.setScrollHereXWithRatio", "设置滚动到此处 X（带比例）", func(center_x_ratio float32) {
		imgui.SetScrollHereXV(center_x_ratio)
	}, true)
	engine.RegisterMethod("imgui.setScrollHereY", "设置滚动到此处 Y", func() {
		imgui.SetScrollHereY()
	}, true)
	engine.RegisterMethod("imgui.setScrollHereYWithRatio", "设置滚动到此处 Y（带比例）", func(center_y_ratio float32) {
		imgui.SetScrollHereYV(center_y_ratio)
	}, true)
	engine.RegisterMethod("imgui.setScrollFromPosX", "从位置 X 滚动", func(local_x float32) {
		imgui.SetScrollFromPosXFloat(local_x)
	}, true)
	engine.RegisterMethod("imgui.setScrollFromPosXWithRatio", "从位置 X 滚动（带比例）", func(local_x, center_x_ratio float32) {
		imgui.SetScrollFromPosXFloatV(local_x, center_x_ratio)
	}, true)
	engine.RegisterMethod("imgui.setScrollFromPosY", "从位置 Y 滚动", func(local_y float32) {
		imgui.SetScrollFromPosYFloat(local_y)
	}, true)
	engine.RegisterMethod("imgui.setScrollFromPosYWithRatio", "从位置 Y 滚动（带比例）", func(local_y, center_y_ratio float32) {
		imgui.SetScrollFromPosYFloatV(local_y, center_y_ratio)
	}, true)

	engine.RegisterMethod("imgui.openPopupOnItemClick", "点击时打开弹出窗口", func() {
		imgui.OpenPopupOnItemClick()
	}, true)
	engine.RegisterMethod("imgui.openPopupOnItemClickWithID", "点击时打开弹出窗口（带ID）", func(str_id string) {
		imgui.OpenPopupOnItemClickV(str_id, 0)
	}, true)
	engine.RegisterMethod("imgui.openPopupID", "打开弹出窗口（ID）", func(id uint32) {
		imgui.OpenPopupID(imgui.ID(id))
	}, true)
	engine.RegisterMethod("imgui.openPopupIDWithFlags", "打开弹出窗口（ID，带标志）", func(id uint32, flags uint32) {
		imgui.OpenPopupIDV(imgui.ID(id), imgui.PopupFlags(flags))
	}, true)

	engine.RegisterMethod("imgui.beginTooltip", "开始工具提示", func() bool {
		return imgui.BeginTooltip()
	}, true)
	engine.RegisterMethod("imgui.endTooltip", "结束工具提示", func() {
		imgui.EndTooltip()
	}, true)
	engine.RegisterMethod("imgui.beginItemTooltip", "开始项工具提示", func() bool {
		return imgui.BeginItemTooltip()
	}, true)
	engine.RegisterMethod("imgui.setTooltip", "设置工具提示", func(text string) {
		imgui.SetTooltip(text)
	}, true)

	engine.RegisterMethod("imgui.calcTextSize", "计算文本大小", func(text string) (float32, float32) {
		size := imgui.CalcTextSize(text)
		return size.X, size.Y
	}, true)
	engine.RegisterMethod("imgui.calcTextSizeWithWrap", "计算文本大小（带换行）", func(text string, wrap_width float32) (float32, float32) {
		size := imgui.CalcTextSizeV(text, false, wrap_width)
		return size.X, size.Y
	}, true)
	engine.RegisterMethod("imgui.calcItemWidth", "计算项宽度", func() float32 {
		return imgui.CalcItemWidth()
	}, true)

	engine.RegisterMethod("imgui.progressBar", "进度条", func(fraction float32) {
		imgui.ProgressBar(fraction)
	}, true)
	engine.RegisterMethod("imgui.progressBarWithSize", "进度条（带大小）", func(fraction, w, h float32) {
		size := imgui.Vec2{X: w, Y: h}
		imgui.ProgressBarV(fraction, size, "")
	}, true)
	engine.RegisterMethod("imgui.progressBarWithOverlay", "进度条（带覆盖）", func(fraction, w, h float32, overlay string) {
		size := imgui.Vec2{X: w, Y: h}
		imgui.ProgressBarV(fraction, size, overlay)
	}, true)

	engine.RegisterMethod("imgui.setScrollX", "设置滚动 X", func(scroll_x float32) {
		imgui.SetScrollXFloat(scroll_x)
	}, true)
	engine.RegisterMethod("imgui.setScrollY", "设置滚动 Y", func(scroll_y float32) {
		imgui.SetScrollYFloat(scroll_y)
	}, true)

	engine.RegisterMethod("imgui.setWindowPos", "设置窗口位置", func(x, y float32) {
		pos := imgui.Vec2{X: x, Y: y}
		imgui.SetWindowPosVec2(pos)
	}, true)
	engine.RegisterMethod("imgui.setWindowPosWithCond", "设置窗口位置（带条件）", func(x, y float32, cond uint32) {
		pos := imgui.Vec2{X: x, Y: y}
		imgui.SetWindowPosVec2V(pos, imgui.Cond(cond))
	}, true)
	engine.RegisterMethod("imgui.setWindowPosByName", "设置窗口位置（按名称）", func(name string, x, y float32) {
		pos := imgui.Vec2{X: x, Y: y}
		imgui.SetWindowPosStr(name, pos)
	}, true)
	engine.RegisterMethod("imgui.setWindowPosByNameWithCond", "设置窗口位置（按名称，带条件）", func(name string, x, y float32, cond uint32) {
		pos := imgui.Vec2{X: x, Y: y}
		imgui.SetWindowPosStrV(name, pos, imgui.Cond(cond))
	}, true)

	engine.RegisterMethod("imgui.setWindowSize", "设置窗口大小", func(w, h float32) {
		size := imgui.Vec2{X: w, Y: h}
		imgui.SetWindowSizeVec2(size)
	}, true)
	engine.RegisterMethod("imgui.setWindowSizeWithCond", "设置窗口大小（带条件）", func(w, h float32, cond uint32) {
		size := imgui.Vec2{X: w, Y: h}
		imgui.SetWindowSizeVec2V(size, imgui.Cond(cond))
	}, true)
	engine.RegisterMethod("imgui.setWindowSizeByName", "设置窗口大小（按名称）", func(name string, w, h float32) {
		size := imgui.Vec2{X: w, Y: h}
		imgui.SetWindowSizeStr(name, size)
	}, true)
	engine.RegisterMethod("imgui.setWindowSizeByNameWithCond", "设置窗口大小（按名称，带条件）", func(name string, w, h float32, cond uint32) {
		size := imgui.Vec2{X: w, Y: h}
		imgui.SetWindowSizeStrV(name, size, imgui.Cond(cond))
	}, true)

	engine.RegisterMethod("imgui.setWindowCollapsed", "设置窗口折叠状态", func(collapsed bool) {
		imgui.SetWindowCollapsedBool(collapsed)
	}, true)
	engine.RegisterMethod("imgui.setWindowCollapsedWithCond", "设置窗口折叠状态（带条件）", func(collapsed bool, cond uint32) {
		imgui.SetWindowCollapsedBoolV(collapsed, imgui.Cond(cond))
	}, true)
	engine.RegisterMethod("imgui.setWindowCollapsedByName", "设置窗口折叠状态（按名称）", func(name string, collapsed bool) {
		imgui.SetWindowCollapsedStr(name, collapsed)
	}, true)
	engine.RegisterMethod("imgui.setWindowCollapsedByNameWithCond", "设置窗口折叠状态（按名称，带条件）", func(name string, collapsed bool, cond uint32) {
		imgui.SetWindowCollapsedStrV(name, collapsed, imgui.Cond(cond))
	}, true)

	engine.RegisterMethod("imgui.setWindowFocus", "设置窗口焦点", func() {
		imgui.SetWindowFocus()
	}, true)
	engine.RegisterMethod("imgui.setWindowFocusByName", "设置窗口焦点（按名称）", func(name string) {
		imgui.SetWindowFocusStr(name)
	}, true)
	engine.RegisterMethod("imgui.setWindowFontScale", "设置窗口字体缩放", func(scale float32) {
		imgui.SetWindowFontScale(scale)
	}, true)

	// 未实现的方法（71个）
	engine.RegisterMethod("imgui.beginDisabled", "开始禁用", func() {
		imgui.BeginDisabled()
	}, true)
	engine.RegisterMethod("imgui.columns", "列", func() {
		imgui.Columns()
	}, true)
	engine.RegisterMethod("imgui.debugFlashStyleColor", "调试闪烁样式颜色", func(idx uint32) {
		imgui.DebugFlashStyleColor(imgui.Col(idx))
	}, true)
	engine.RegisterMethod("imgui.debugLog", "调试日志", func(fmt string) {
		imgui.DebugLog(fmt)
	}, true)
	engine.RegisterMethod("imgui.debugStartItemPicker", "调试开始项目选择器", func() {
		imgui.DebugStartItemPicker()
	}, true)
	engine.RegisterMethod("imgui.debugTextEncoding", "调试文本编码", func(text string) {
		imgui.DebugTextEncoding(text)
	}, true)
	engine.RegisterMethod("imgui.destroyContextV", "销毁上下文（带参数）", func(ctx *imgui.Context) {
		imgui.DestroyContextV(ctx)
	}, true)
	engine.RegisterMethod("imgui.destroyPlatformWindows", "销毁平台窗口", func() {
		imgui.DestroyPlatformWindows()
	}, true)
	engine.RegisterMethod("imgui.endListBox", "结束列表框", func() {
		imgui.EndListBox()
	}, true)
	engine.RegisterMethod("imgui.image", "图像", func(tex_ref imgui.TextureRef, w, h float32) {
		size := imgui.Vec2{X: w, Y: h}
		imgui.Image(tex_ref, size)
	}, true)
	engine.RegisterMethod("imgui.imageWithBg", "图像（带背景）", func(tex_ref imgui.TextureRef, w, h float32) {
		size := imgui.Vec2{X: w, Y: h}
		imgui.ImageWithBg(tex_ref, size)
	}, true)
	engine.RegisterMethod("imgui.loadIniSettingsFromMemoryV", "从内存加载INI设置（带参数）", func(ini_data string, ini_size uint64) {
		imgui.LoadIniSettingsFromMemoryV(ini_data, ini_size)
	}, true)
	engine.RegisterMethod("imgui.logButtons", "日志按钮", func() {
		imgui.LogButtons()
	}, true)
	engine.RegisterMethod("imgui.logFinish", "完成日志", func() {
		imgui.LogFinish()
	}, true)
	engine.RegisterMethod("imgui.logText", "日志文本", func(fmt string) {
		imgui.LogText(fmt)
	}, true)
	engine.RegisterMethod("imgui.logToClipboard", "日志到剪贴板", func() {
		imgui.LogToClipboard()
	}, true)
	engine.RegisterMethod("imgui.logToClipboardV", "日志到剪贴板（带参数）", func(auto_open_depth int32) {
		imgui.LogToClipboardV(auto_open_depth)
	}, true)
	engine.RegisterMethod("imgui.logToFile", "日志到文件", func() {
		imgui.LogToFile()
	}, true)
	engine.RegisterMethod("imgui.logToFileV", "日志到文件（带参数）", func(auto_open_depth int32, filename string) {
		imgui.LogToFileV(auto_open_depth, filename)
	}, true)
	engine.RegisterMethod("imgui.logToTTY", "日志到TTY", func() {
		imgui.LogToTTY()
	}, true)
	engine.RegisterMethod("imgui.logToTTYV", "日志到TTY（带参数）", func(auto_open_depth int32) {
		imgui.LogToTTYV(auto_open_depth)
	}, true)
	engine.RegisterMethod("imgui.memFree", "释放内存", func(ptr uintptr) {
		imgui.MemFree(ptr)
	}, true)
	engine.RegisterMethod("imgui.nextColumn", "下一列", func() {
		imgui.NextColumn()
	}, true)
	engine.RegisterMethod("imgui.openPopupStrV", "打开弹出窗口（带参数）", func(str_id string, popup_flags uint32) {
		imgui.OpenPopupStrV(str_id, imgui.PopupFlags(popup_flags))
	}, true)
	engine.RegisterMethod("imgui.plotHistogramFloatPtr", "绘制直方图（浮点指针）", func(label string, values []float32) {
		if len(values) > 0 {
			imgui.PlotHistogramFloatPtr(label, &values[0], int32(len(values)))
		}
	}, true)
	engine.RegisterMethod("imgui.plotHistogramFloatPtrV", "绘制直方图（浮点指针，带参数）", func(label string, values []float32, values_offset int32, overlay_text string, scale_min, scale_max float32, w, h float32, stride int32) {
		if len(values) > 0 {
			graph_size := imgui.Vec2{X: w, Y: h}
			imgui.PlotHistogramFloatPtrV(label, &values[0], int32(len(values)), values_offset, overlay_text, scale_min, scale_max, graph_size, stride)
		}
	}, true)
	engine.RegisterMethod("imgui.plotLinesFloatPtr", "绘制线条（浮点指针）", func(label string, values []float32) {
		if len(values) > 0 {
			imgui.PlotLinesFloatPtr(label, &values[0], int32(len(values)))
		}
	}, true)
	engine.RegisterMethod("imgui.plotLinesFloatPtrV", "绘制线条（浮点指针，带参数）", func(label string, values []float32, values_offset int32, overlay_text string, scale_min, scale_max float32, w, h float32, stride int32) {
		if len(values) > 0 {
			graph_size := imgui.Vec2{X: w, Y: h}
			imgui.PlotLinesFloatPtrV(label, &values[0], int32(len(values)), values_offset, overlay_text, scale_min, scale_max, graph_size, stride)
		}
	}, true)
	engine.RegisterMethod("imgui.pushIDStrStr", "推送ID（字符串范围）", func(str_id_begin, str_id_end string) {
		imgui.PushIDStrStr(str_id_begin, str_id_end)
	}, true)
	engine.RegisterMethod("imgui.renderPlatformWindowsDefault", "渲染平台窗口（默认）", func() {
		imgui.RenderPlatformWindowsDefault()
	}, true)
	engine.RegisterMethod("imgui.renderPlatformWindowsDefaultV", "渲染平台窗口（默认，带参数）", func(platform_render_arg, renderer_render_arg uintptr) {
		imgui.RenderPlatformWindowsDefaultV(platform_render_arg, renderer_render_arg)
	}, true)
	engine.RegisterMethod("imgui.resetMouseDragDeltaV", "重置鼠标拖动增量（带参数）", func(button uint32) {
		imgui.ResetMouseDragDeltaV(imgui.MouseButton(button))
	}, true)
	engine.RegisterMethod("imgui.sameLineV", "同一行（带参数）", func(offset_from_start_x, spacing float32) {
		imgui.SameLineV(offset_from_start_x, spacing)
	}, true)
	engine.RegisterMethod("imgui.setItemTooltip", "设置项目工具提示", func(fmt string) {
		imgui.SetItemTooltip(fmt)
	}, true)
	engine.RegisterMethod("imgui.setNextItemSelectionUserData", "设置下一项选择用户数据", func(selection_user_data uint64) {
		imgui.SetNextItemSelectionUserData(imgui.SelectionUserData(selection_user_data))
	}, true)
	engine.RegisterMethod("imgui.setNextItemShortcut", "设置下一项快捷键", func(key_chord uint32) {
		imgui.SetNextItemShortcut(imgui.KeyChord(key_chord))
	}, true)
	engine.RegisterMethod("imgui.setNextItemShortcutV", "设置下一项快捷键（带参数）", func(key_chord, flags uint32) {
		imgui.SetNextItemShortcutV(imgui.KeyChord(key_chord), imgui.InputFlags(flags))
	}, true)
	engine.RegisterMethod("imgui.setNextItemStorageID", "设置下一项存储ID", func(storage_id uint32) {
		imgui.SetNextItemStorageID(imgui.ID(storage_id))
	}, true)
	engine.RegisterMethod("imgui.setNextItemWidth", "设置下一项宽度", func(item_width float32) {
		imgui.SetNextItemWidth(item_width)
	}, true)
	engine.RegisterMethod("imgui.setNextWindowDockID", "设置下一窗口停靠ID", func(dock_id uint32) {
		imgui.SetNextWindowDockID(imgui.ID(dock_id))
	}, true)
	engine.RegisterMethod("imgui.setNextWindowPosV", "设置下一窗口位置（带参数）", func(x, y float32, cond uint32, pivot_x, pivot_y float32) {
		pos := imgui.Vec2{X: x, Y: y}
		pivot := imgui.Vec2{X: pivot_x, Y: pivot_y}
		imgui.SetNextWindowPosV(pos, imgui.Cond(cond), pivot)
	}, true)
	engine.RegisterMethod("imgui.setNextWindowSizeV", "设置下一窗口大小（带参数）", func(w, h float32, cond uint32) {
		size := imgui.Vec2{X: w, Y: h}
		imgui.SetNextWindowSizeV(size, imgui.Cond(cond))
	}, true)
	engine.RegisterMethod("imgui.setNextWindowViewport", "设置下一窗口视口", func(viewport_id uint32) {
		imgui.SetNextWindowViewport(imgui.ID(viewport_id))
	}, true)
	engine.RegisterMethod("imgui.showAboutWindowV", "显示关于窗口（带参数）", func(p_open *bool) {
		imgui.ShowAboutWindowV(p_open)
	}, true)
	engine.RegisterMethod("imgui.showDebugLogWindowV", "显示调试日志窗口（带参数）", func(p_open *bool) {
		imgui.ShowDebugLogWindowV(p_open)
	}, true)
	engine.RegisterMethod("imgui.showDemoWindowV", "显示演示窗口（带参数）", func(p_open *bool) {
		imgui.ShowDemoWindowV(p_open)
	}, true)
	engine.RegisterMethod("imgui.showIDStackToolWindowV", "显示ID堆栈工具窗口（带参数）", func(p_open *bool) {
		imgui.ShowIDStackToolWindowV(p_open)
	}, true)
	engine.RegisterMethod("imgui.showMetricsWindowV", "显示指标窗口（带参数）", func(p_open *bool) {
		imgui.ShowMetricsWindowV(p_open)
	}, true)
	engine.RegisterMethod("imgui.showStyleEditorV", "显示样式编辑器（带参数）", func(ref *imgui.Style) {
		imgui.ShowStyleEditorV(ref)
	}, true)
	engine.RegisterMethod("imgui.styleColorsClassic", "经典样式颜色", func() {
		imgui.StyleColorsClassic()
	}, true)
	engine.RegisterMethod("imgui.styleColorsClassicV", "经典样式颜色（带参数）", func(dst *imgui.Style) {
		imgui.StyleColorsClassicV(dst)
	}, true)
	engine.RegisterMethod("imgui.styleColorsDark", "深色样式颜色", func() {
		imgui.StyleColorsDark()
	}, true)
	engine.RegisterMethod("imgui.styleColorsDarkV", "深色样式颜色（带参数）", func(dst *imgui.Style) {
		imgui.StyleColorsDarkV(dst)
	}, true)
	engine.RegisterMethod("imgui.styleColorsLight", "浅色样式颜色", func() {
		imgui.StyleColorsLight()
	}, true)
	engine.RegisterMethod("imgui.styleColorsLightV", "浅色样式颜色（带参数）", func(dst *imgui.Style) {
		imgui.StyleColorsLightV(dst)
	}, true)
	engine.RegisterMethod("imgui.tableAngledHeadersRow", "表格角度标题行", func() {
		imgui.TableAngledHeadersRow()
	}, true)
	engine.RegisterMethod("imgui.tableNextRow", "表格下一行", func() {
		imgui.TableNextRow()
	}, true)
	engine.RegisterMethod("imgui.tableSetBgColor", "设置表格背景颜色", func(target, color uint32) {
		imgui.TableSetBgColor(imgui.TableBgTarget(target), color)
	}, true)
	engine.RegisterMethod("imgui.tableSetupColumn", "设置表格列", func(label string) {
		imgui.TableSetupColumn(label)
	}, true)
	engine.RegisterMethod("imgui.textUnformattedV", "无格式文本（带参数）", func(text string) {
		imgui.TextUnformattedV(text)
	}, true)
	engine.RegisterMethod("imgui.treePushPtr", "树推送（指针）", func(ptr_id uintptr) {
		imgui.TreePushPtr(ptr_id)
	}, true)
	engine.RegisterMethod("imgui.updatePlatformWindows", "更新平台窗口", func() {
		imgui.UpdatePlatformWindows()
	}, true)
	engine.RegisterMethod("imgui.valueBool", "布尔值", func(prefix string, b bool) {
		imgui.ValueBool(prefix, b)
	}, true)
	engine.RegisterMethod("imgui.valueFloat", "浮点值", func(prefix string, v float32) {
		imgui.ValueFloat(prefix, v)
	}, true)
	engine.RegisterMethod("imgui.valueFloatV", "浮点值（带参数）", func(prefix string, v float32, float_format string) {
		imgui.ValueFloatV(prefix, v, float_format)
	}, true)
	engine.RegisterMethod("imgui.valueInt", "整数值", func(prefix string, v int32) {
		imgui.ValueInt(prefix, v)
	}, true)
	engine.RegisterMethod("imgui.valueUint", "无符号整数值", func(prefix string, v uint32) {
		imgui.ValueUint(prefix, v)
	}, true)

	// 强行实现的底层方法（4个）
	imguiObj.RawSetString("setAssertHandler", state.NewFunction(func(L *lua.LState) int {
		// Lua 脚本可以设置一个自定义的断言处理函数
		// 这里我们提供一个简单的实现，允许 Lua 脚本传递一个回调函数
		if L.GetTop() > 0 && L.CheckFunction(1) != nil {
			// 获取 Lua 回调函数
			callback := L.CheckFunction(1)

			// 设置一个 Go 断言处理器，它会调用 Lua 回调
			imgui.SetAssertHandler(func(expression, file string, line int) {
				// 在 Lua 环境中调用回调函数
				L.Push(callback)
				L.Push(lua.LString(expression))
				L.Push(lua.LString(file))
				L.Push(lua.LNumber(line))
				L.PCall(3, 0, nil)
			})
		} else {
			// 如果没有传递回调函数，则使用默认的处理器（panic）
			imgui.SetAssertHandler(nil)
		}
		return 0
	}))

	imguiObj.RawSetString("setAllocatorFunctions", state.NewFunction(func(L *lua.LState) int {
		// 设置内存分配函数
		// Lua 脚本可以传递自定义的分配和释放函数
		if L.GetTop() >= 2 {
			allocFunc := L.CheckFunction(1)
			freeFunc := L.CheckFunction(2)

			// 创建 Go 内存分配函数，它会调用 Lua 回调
			alloc := imgui.MemAllocFunc(func(sz uint64, user_data unsafe.Pointer) unsafe.Pointer {
				L.Push(allocFunc)
				L.Push(lua.LNumber(sz))
				L.PCall(1, 1, nil)

				// 获取返回的指针值
				result := L.CheckInt(-1)
				L.Pop(1)

				return unsafe.Pointer(uintptr(result))
			})

			// 创建 Go 内存释放函数，它会调用 Lua 回调
			free := imgui.MemFreeFunc(func(ptr unsafe.Pointer, user_data unsafe.Pointer) {
				L.Push(freeFunc)
				L.Push(lua.LNumber(uintptr(ptr)))
				L.PCall(1, 0, nil)
			})

			imgui.SetAllocatorFunctions(alloc, free)
		} else {
			// 如果没有传递回调函数，则使用默认的内存分配器
			imgui.SetAllocatorFunctions(nil, nil)
		}
		return 0
	}))

	imguiObj.RawSetString("setAllocatorFunctionsV", state.NewFunction(func(L *lua.LState) int {
		// 设置内存分配函数（带用户数据）
		if L.GetTop() >= 2 {
			allocFunc := L.CheckFunction(1)
			freeFunc := L.CheckFunction(2)
			userData := uintptr(0)
			if L.GetTop() > 2 {
				userData = uintptr(L.CheckNumber(3))
			}

			// 创建 Go 内存分配函数，它会调用 Lua 回调
			alloc := imgui.MemAllocFunc(func(sz uint64, user_data unsafe.Pointer) unsafe.Pointer {
				L.Push(allocFunc)
				L.Push(lua.LNumber(sz))
				L.PCall(1, 1, nil)

				result := L.CheckInt(-1)
				L.Pop(1)

				return unsafe.Pointer(uintptr(result))
			})

			// 创建 Go 内存释放函数，它会调用 Lua 回调
			free := imgui.MemFreeFunc(func(ptr unsafe.Pointer, user_data unsafe.Pointer) {
				L.Push(freeFunc)
				L.Push(lua.LNumber(uintptr(ptr)))
				L.PCall(1, 0, nil)
			})

			imgui.SetAllocatorFunctionsV(alloc, free, userData)
		} else {
			imgui.SetAllocatorFunctionsV(nil, nil, 0)
		}
		return 0
	}))

	imguiObj.RawSetString("setStateStorage", state.NewFunction(func(L *lua.LState) int {
		// 设置状态存储
		// Lua 脚本可以传递一个存储对象
		if L.GetTop() > 0 {
			// 检查是否传递了存储对象
			if L.CheckTable(1) != nil {
				// 如果传递了 Lua 表，我们可以创建一个简单的存储适配器
				// 这里我们暂时使用默认的存储
				imgui.SetStateStorage(nil)
			} else {
				// 使用默认的存储
				imgui.SetStateStorage(nil)
			}
		} else {
			// 使用默认的存储
			imgui.SetStateStorage(nil)
		}
		return 0
	}))

	return nil
}

func parseColorString(colorStr string) uint32 {
	if len(colorStr) == 0 {
		return 0xFFFFFFFF
	}
	if colorStr[0] == '#' {
		colorStr = colorStr[1:]
	}
	color, err := strconv.ParseUint(colorStr, 16, 32)
	if err != nil {
		return 0xFFFFFFFF
	}
	return uint32(color)
}
