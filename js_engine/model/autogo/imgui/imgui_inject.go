package imgui

import (
	"image"
	"image/color"
	"reflect"
	"strconv"

	"github.com/Dasongzi1366/AutoGo/imgui"
	"github.com/ZingYao/autogo_scriptengine/js_engine/model"
	"github.com/dop251/goja"
)

// parseColorString 解析颜色字符串，支持 #RRGGBB、#RRGGBBAA、RRGGBB、RRGGBBAA 格式
func parseColorString(colorStr string) imgui.Col {
	if len(colorStr) == 0 {
		return imgui.Col(0)
	}

	// 移除 # 前缀
	if colorStr[0] == '#' {
		colorStr = colorStr[1:]
	}

	var r, g, b, a uint32
	a = 255 // 默认不透明

	if len(colorStr) == 6 {
		// #RRGGBB 格式
		val, _ := strconv.ParseUint(colorStr, 16, 32)
		r = uint32((val >> 16) & 0xFF)
		g = uint32((val >> 8) & 0xFF)
		b = uint32(val & 0xFF)
	} else if len(colorStr) == 8 {
		// #RRGGBBAA 格式
		val, _ := strconv.ParseUint(colorStr, 16, 32)
		r = uint32((val >> 24) & 0xFF)
		g = uint32((val >> 16) & 0xFF)
		b = uint32((val >> 8) & 0xFF)
		a = uint32(val & 0xFF)
	}

	return imgui.Col((a << 24) | (r << 16) | (g << 8) | b)
}

// parseVec2 解析 Vec2 类型
func parseVec2(vm *goja.Runtime, value goja.Value) imgui.Vec2 {
	if value == nil || goja.IsUndefined(value) || goja.IsNull(value) {
		return imgui.Vec2{X: 0, Y: 0}
	}

	if value.ExportType().Kind() == reflect.Map {
		// 对象格式 {x: 1, y: 2}
		obj := value.ToObject(vm)
		if obj != nil {
			x := obj.Get("x")
			y := obj.Get("y")
			return imgui.Vec2{
				X: float32(x.ToFloat()),
				Y: float32(y.ToFloat()),
			}
		}
	} else if value.ExportType().Kind() == reflect.Array || value.ExportType().Kind() == reflect.Slice {
		// 数组格式 [1, 2]
		obj := value.ToObject(vm)
		if obj != nil {
			x := obj.Get("0")
			y := obj.Get("1")
			return imgui.Vec2{
				X: float32(x.ToFloat()),
				Y: float32(y.ToFloat()),
			}
		}
	}

	return imgui.Vec2{X: 0, Y: 0}
}

// parseVec4 解析 Vec4 类型
func parseVec4(vm *goja.Runtime, value goja.Value) imgui.Vec4 {
	if value == nil || goja.IsUndefined(value) || goja.IsNull(value) {
		return imgui.Vec4{X: 0, Y: 0, Z: 0, W: 0}
	}

	if value.ExportType().Kind() == reflect.Map {
		// 对象格式 {x: 1, y: 2, z: 3, w: 4}
		obj := value.ToObject(vm)
		if obj != nil {
			x := obj.Get("x")
			y := obj.Get("y")
			z := obj.Get("z")
			w := obj.Get("w")
			return imgui.Vec4{
				X: float32(x.ToFloat()),
				Y: float32(y.ToFloat()),
				Z: float32(z.ToFloat()),
				W: float32(w.ToFloat()),
			}
		}
	} else if value.ExportType().Kind() == reflect.Array || value.ExportType().Kind() == reflect.Slice {
		// 数组格式 [1, 2, 3, 4]
		obj := value.ToObject(vm)
		if obj != nil {
			x := obj.Get("0")
			y := obj.Get("1")
			z := obj.Get("2")
			w := obj.Get("3")
			return imgui.Vec4{
				X: float32(x.ToFloat()),
				Y: float32(y.ToFloat()),
				Z: float32(z.ToFloat()),
				W: float32(w.ToFloat()),
			}
		}
	}

	return imgui.Vec4{X: 0, Y: 0, Z: 0, W: 0}
}

// parseViewport 解析 Viewport 类型
func parseViewport(vm *goja.Runtime, value goja.Value) *imgui.Viewport {
	if value == nil || goja.IsUndefined(value) || goja.IsNull(value) {
		return nil
	}

	obj := value.ToObject(vm)
	if obj != nil {
		// 对象格式 {id: 1, pos: {x: 0, y: 0}, size: {x: 100, y: 100}}
		id := obj.Get("id")
		if id != nil && !goja.IsUndefined(id) && !goja.IsNull(id) {
			viewport_id := imgui.ID(id.ToInteger())
			return imgui.FindViewportByID(viewport_id)
		}
	}

	return nil
}

// parseViewportID 解析 ViewportID 类型
func parseViewportID(vm *goja.Runtime, value goja.Value) imgui.ID {
	if value == nil || goja.IsUndefined(value) || goja.IsNull(value) {
		return imgui.ID(0)
	}

	return imgui.ID(value.ToInteger())
}

// parseStringArray 解析字符串数组
func parseStringArray(vm *goja.Runtime, value goja.Value) []string {
	if value == nil || goja.IsUndefined(value) || goja.IsNull(value) {
		return []string{}
	}

	if value.ExportType().Kind() == reflect.Array || value.ExportType().Kind() == reflect.Slice {
		// 数组格式 ["a", "b", "c"]
		obj := value.ToObject(vm)
		if obj != nil {
			length := int(obj.Get("length").ToInteger())
			result := make([]string, length)
			for i := 0; i < length; i++ {
				elem := obj.Get(strconv.Itoa(i))
				result[i] = elem.String()
			}
			return result
		}
	}

	return []string{}
}

// Register 向引擎注册方法

// ImGuiModule imgui 模块
type ImGuiModule struct{}

// Name 返回模块名称
func (m *ImGuiModule) Name() string {
	return "imgui"
}

// IsAvailable 检查模块是否可用
func (m *ImGuiModule) IsAvailable() bool {
	return true
}

// Register 向引擎注册方法
func (m *ImGuiModule) Register(engine model.Engine) error {
	vm := engine.GetVM()
	imguiObj := vm.NewObject()
	vm.Set("imgui", imguiObj)

	// ========== backend.go ==========
	imguiObj.Set("init", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.Init())
	})
	imguiObj.Set("close", func(call goja.FunctionCall) goja.Value {
		imgui.Close()
		return goja.Undefined()
	})

	// ========== assert.go ==========
	imguiObj.Set("setAssertHandler", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理函数类型 AssertHandler
		var handler imgui.AssertHandler
		imgui.SetAssertHandler(handler)
		return goja.Undefined()
	})

	// ========== util.go ==========
	imguiObj.Set("vertexBufferLayout", func(call goja.FunctionCall) goja.Value {
		// 多个返回值，需要用 object 或者 array 包裹
		result := vm.NewObject()
		r0, r1, r2, r3 := imgui.VertexBufferLayout()
		result.Set("r0", vm.ToValue(r0))
		result.Set("r1", vm.ToValue(r1))
		result.Set("r2", vm.ToValue(r2))
		result.Set("r3", vm.ToValue(r3))
		return result
	})
	imguiObj.Set("indexBufferLayout", func(call goja.FunctionCall) goja.Value {
		// 多个返回值，需要用 object 或者 array 包裹
		result := vm.NewObject()
		r0 := imgui.IndexBufferLayout()
		result.Set("r0", vm.ToValue(r0))
		return result
	})
	imguiObj.Set("newGlyphRange", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.NewGlyphRange())
	})

	// ========== typedefs.go ==========
	imguiObj.Set("newEmptyBitArrayForNamedKeys", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyBitArrayForNamedKeys()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyBitVector", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyBitVector()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyDrawChannel", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyDrawChannel()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyDrawCmd", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyDrawCmd()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyDrawCmdHeader", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyDrawCmdHeader()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyDrawData", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyDrawData()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyDrawDataBuilder", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyDrawDataBuilder()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyDrawList", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyDrawList()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyDrawListSharedData", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyDrawListSharedData()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyDrawListSplitter", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyDrawListSplitter()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyDrawVert", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyDrawVert()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyFont", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyFont()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyFontAtlas", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyFontAtlas()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyFontAtlasBuilder", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyFontAtlasBuilder()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyFontAtlasPostProcessData", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyFontAtlasPostProcessData()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyFontAtlasRect", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyFontAtlasRect()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyFontAtlasRectEntry", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyFontAtlasRectEntry()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyFontBaked", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyFontBaked()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyFontConfig", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyFontConfig()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyFontGlyph", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyFontGlyph()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyFontGlyphRangesBuilder", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyFontGlyphRangesBuilder()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyFontLoader", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyFontLoader()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyFontStackData", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyFontStackData()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyBoxSelectState", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyBoxSelectState()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyColorMod", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyColorMod()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyComboPreviewData", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyComboPreviewData()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyContext", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyContext()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyContextHook", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyContextHook()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyDataTypeInfo", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyDataTypeInfo()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyDataTypeStorage", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyDataTypeStorage()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyDeactivatedItemData", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyDeactivatedItemData()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyDebugAllocEntry", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyDebugAllocEntry()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyDebugAllocInfo", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyDebugAllocInfo()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyDockContext", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyDockContext()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyDockNode", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyDockNode()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyErrorRecoveryState", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyErrorRecoveryState()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyFocusScopeData", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyFocusScopeData()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyGroupData", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyGroupData()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyIDStackTool", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyIDStackTool()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyIO", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyIO()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyInputEvent", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyInputEvent()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyInputEventAppFocused", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyInputEventAppFocused()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyInputEventKey", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyInputEventKey()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyInputEventMouseButton", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyInputEventMouseButton()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyInputEventMousePos", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyInputEventMousePos()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyInputEventMouseViewport", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyInputEventMouseViewport()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyInputEventMouseWheel", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyInputEventMouseWheel()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyInputEventText", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyInputEventText()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyInputTextCallbackData", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyInputTextCallbackData()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyInputTextDeactivatedState", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyInputTextDeactivatedState()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyInputTextState", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyInputTextState()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyKeyData", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyKeyData()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyKeyOwnerData", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyKeyOwnerData()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyKeyRoutingData", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyKeyRoutingData()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyKeyRoutingTable", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyKeyRoutingTable()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyLastItemData", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyLastItemData()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyListClipper", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyListClipper()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyListClipperData", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyListClipperData()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyListClipperRange", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyListClipperRange()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyLocEntry", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyLocEntry()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyMenuColumns", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyMenuColumns()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyMetricsConfig", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyMetricsConfig()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyMultiSelectIO", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyMultiSelectIO()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyMultiSelectState", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyMultiSelectState()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyMultiSelectTempData", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyMultiSelectTempData()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyNavItemData", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyNavItemData()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyNextItemData", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyNextItemData()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyNextWindowData", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyNextWindowData()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyOldColumnData", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyOldColumnData()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyOldColumns", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyOldColumns()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyOnceUponAFrame", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyOnceUponAFrame()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyPayload", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyPayload()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyPlatformIO", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyPlatformIO()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyPlatformImeData", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyPlatformImeData()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyPlatformMonitor", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyPlatformMonitor()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyPopupData", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyPopupData()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyPtrOrIndex", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyPtrOrIndex()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptySelectionBasicStorage", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptySelectionBasicStorage()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptySelectionExternalStorage", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptySelectionExternalStorage()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptySelectionRequest", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptySelectionRequest()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptySettingsHandler", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptySettingsHandler()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyShrinkWidthItem", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyShrinkWidthItem()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptySizeCallbackData", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptySizeCallbackData()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyStackLevelInfo", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyStackLevelInfo()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyStorage", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyStorage()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyStoragePair", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyStoragePair()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyStyle", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyStyle()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyStyleMod", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyStyleMod()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyStyleVarInfo", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyStyleVarInfo()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyTabBar", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyTabBar()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyTabItem", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyTabItem()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyTable", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyTable()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyTableCellData", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyTableCellData()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyTableColumn", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyTableColumn()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyTableColumnSettings", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyTableColumnSettings()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyTableColumnSortSpecs", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyTableColumnSortSpecs()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyTableHeaderData", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyTableHeaderData()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyTableInstanceData", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyTableInstanceData()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyTableSettings", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyTableSettings()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyTableSortSpecs", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyTableSortSpecs()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyTableTempData", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyTableTempData()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyTextBuffer", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyTextBuffer()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyTextFilter", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyTextFilter()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyTextIndex", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyTextIndex()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyTextRange", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyTextRange()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyTreeNodeStackData", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyTreeNodeStackData()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyTypingSelectRequest", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyTypingSelectRequest()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyTypingSelectState", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyTypingSelectState()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyViewport", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyViewport()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyViewportP", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyViewportP()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyWindow", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyWindow()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyWindowClass", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyWindowClass()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyWindowDockStyle", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyWindowDockStyle()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyWindowSettings", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyWindowSettings()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyWindowStackData", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyWindowStackData()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyWindowTempData", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyWindowTempData()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyTextureData", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyTextureData()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyTextureRect", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyTextureRect()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyTextureRef", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyTextureRef()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyVec1", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyVec1()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptyVec2i", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptyVec2i()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newEmptystbrpcontextopaque", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewEmptystbrpcontextopaque()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})

	// ========== texture.go ==========
	imguiObj.Set("createTextureNrgba", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *image.NRGBA
		var img *image.NRGBA
		result := imgui.CreateTextureNrgba(img)
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})

	// ========== funcs.go ==========
	imguiObj.Set("colorHSVV", func(call goja.FunctionCall) goja.Value {
		h := float32(call.Argument(0).ToFloat())
		s := float32(call.Argument(1).ToFloat())
		v := float32(call.Argument(2).ToFloat())
		a := float32(call.Argument(3).ToFloat())
		return vm.ToValue(imgui.ColorHSVV(h, s, v, a))
	})
	imguiObj.Set("newDrawCmd", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewDrawCmd()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalNewDrawDataBuilder", func(call goja.FunctionCall) goja.Value {
		result := imgui.InternalNewDrawDataBuilder()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newDrawData", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewDrawData()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalNewDrawListSharedData", func(call goja.FunctionCall) goja.Value {
		result := imgui.InternalNewDrawListSharedData()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newDrawListSplitter", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewDrawListSplitter()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newDrawList", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *DrawListSharedData
		var shared_data *imgui.DrawListSharedData
		result := imgui.NewDrawList(shared_data)
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalNewFontAtlasBuilder", func(call goja.FunctionCall) goja.Value {
		result := imgui.InternalNewFontAtlasBuilder()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newFontAtlasRect", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewFontAtlasRect()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newFontAtlas", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewFontAtlas()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newFontBaked", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewFontBaked()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newFontConfig", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewFontConfig()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newFontGlyphRangesBuilder", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewFontGlyphRangesBuilder()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newFontGlyph", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewFontGlyph()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalNewFontLoader", func(call goja.FunctionCall) goja.Value {
		result := imgui.InternalNewFontLoader()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newFont", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewFont()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalNewBoxSelectState", func(call goja.FunctionCall) goja.Value {
		result := imgui.InternalNewBoxSelectState()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalNewComboPreviewData", func(call goja.FunctionCall) goja.Value {
		result := imgui.InternalNewComboPreviewData()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalNewContextHook", func(call goja.FunctionCall) goja.Value {
		result := imgui.InternalNewContextHook()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalNewContext", func(call goja.FunctionCall) goja.Value {
		var shared_font_atlas *imgui.FontAtlas
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.FontAtlas); ok {
				shared_font_atlas = v
			}
		}
		result := imgui.InternalNewContext(shared_font_atlas)
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalNewDebugAllocInfo", func(call goja.FunctionCall) goja.Value {
		result := imgui.InternalNewDebugAllocInfo()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalNewDockContext", func(call goja.FunctionCall) goja.Value {
		result := imgui.InternalNewDockContext()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalNewDockNode", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 ID
		id := imgui.ID(0)
		result := imgui.InternalNewDockNode(id)
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalNewErrorRecoveryState", func(call goja.FunctionCall) goja.Value {
		result := imgui.InternalNewErrorRecoveryState()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalNewIDStackTool", func(call goja.FunctionCall) goja.Value {
		result := imgui.InternalNewIDStackTool()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newIO", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewIO()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalNewInputEvent", func(call goja.FunctionCall) goja.Value {
		result := imgui.InternalNewInputEvent()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newInputTextCallbackData", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewInputTextCallbackData()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalNewInputTextDeactivatedState", func(call goja.FunctionCall) goja.Value {
		result := imgui.InternalNewInputTextDeactivatedState()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalNewInputTextState", func(call goja.FunctionCall) goja.Value {
		result := imgui.InternalNewInputTextState()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalNewKeyOwnerData", func(call goja.FunctionCall) goja.Value {
		result := imgui.InternalNewKeyOwnerData()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalNewKeyRoutingData", func(call goja.FunctionCall) goja.Value {
		result := imgui.InternalNewKeyRoutingData()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalNewKeyRoutingTable", func(call goja.FunctionCall) goja.Value {
		result := imgui.InternalNewKeyRoutingTable()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalNewLastItemData", func(call goja.FunctionCall) goja.Value {
		result := imgui.InternalNewLastItemData()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalNewListClipperData", func(call goja.FunctionCall) goja.Value {
		result := imgui.InternalNewListClipperData()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalListClipperRangeFromIndices", func(call goja.FunctionCall) goja.Value {
		min := int32(call.Argument(0).ToInteger())
		max := int32(call.Argument(1).ToInteger())
		return vm.ToValue(imgui.InternalListClipperRangeFromIndices(min, max))
	})
	imguiObj.Set("internalListClipperRangeFromPositions", func(call goja.FunctionCall) goja.Value {
		y1 := float32(call.Argument(0).ToFloat())
		y2 := float32(call.Argument(1).ToFloat())
		off_min := int32(call.Argument(2).ToInteger())
		off_max := int32(call.Argument(3).ToInteger())
		return vm.ToValue(imgui.InternalListClipperRangeFromPositions(y1, y2, off_min, off_max))
	})
	imguiObj.Set("newListClipper", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewListClipper()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalNewMenuColumns", func(call goja.FunctionCall) goja.Value {
		result := imgui.InternalNewMenuColumns()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalNewMultiSelectState", func(call goja.FunctionCall) goja.Value {
		result := imgui.InternalNewMultiSelectState()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalNewMultiSelectTempData", func(call goja.FunctionCall) goja.Value {
		result := imgui.InternalNewMultiSelectTempData()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalNewNavItemData", func(call goja.FunctionCall) goja.Value {
		result := imgui.InternalNewNavItemData()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalNewNextItemData", func(call goja.FunctionCall) goja.Value {
		result := imgui.InternalNewNextItemData()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalNewNextWindowData", func(call goja.FunctionCall) goja.Value {
		result := imgui.InternalNewNextWindowData()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalNewOldColumnData", func(call goja.FunctionCall) goja.Value {
		result := imgui.InternalNewOldColumnData()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalNewOldColumns", func(call goja.FunctionCall) goja.Value {
		result := imgui.InternalNewOldColumns()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newOnceUponAFrame", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewOnceUponAFrame()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newPayload", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewPayload()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newPlatformIO", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewPlatformIO()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newPlatformImeData", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewPlatformImeData()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newPlatformMonitor", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewPlatformMonitor()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalNewPopupData", func(call goja.FunctionCall) goja.Value {
		result := imgui.InternalNewPopupData()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalNewPtrOrIndexInt", func(call goja.FunctionCall) goja.Value {
		index := int32(call.Argument(0).ToInteger())
		result := imgui.InternalNewPtrOrIndexInt(index)
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalNewPtrOrIndexPtr", func(call goja.FunctionCall) goja.Value {
		ptr := uintptr(call.Argument(0).ToInteger())
		result := imgui.InternalNewPtrOrIndexPtr(ptr)
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newSelectionBasicStorage", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewSelectionBasicStorage()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newSelectionExternalStorage", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewSelectionExternalStorage()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalNewSettingsHandler", func(call goja.FunctionCall) goja.Value {
		result := imgui.InternalNewSettingsHandler()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalNewStackLevelInfo", func(call goja.FunctionCall) goja.Value {
		result := imgui.InternalNewStackLevelInfo()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newStoragePairFloat", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 ID
		_key := imgui.ID(0)
		_val := float32(call.Argument(1).ToFloat())
		result := imgui.NewStoragePairFloat(_key, _val)
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newStoragePairInt", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 ID
		_key := imgui.ID(0)
		_val := int32(call.Argument(1).ToInteger())
		result := imgui.NewStoragePairInt(_key, _val)
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newStoragePairPtr", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 ID
		_key := imgui.ID(0)
		_val := uintptr(call.Argument(1).ToInteger())
		result := imgui.NewStoragePairPtr(_key, _val)
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalNewStyleModFloat", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 StyleVar
		idx := imgui.StyleVar(0)
		v := float32(call.Argument(1).ToFloat())
		result := imgui.InternalNewStyleModFloat(idx, v)
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalNewStyleModInt", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 StyleVar
		idx := imgui.StyleVar(0)
		v := int32(call.Argument(1).ToInteger())
		result := imgui.InternalNewStyleModInt(idx, v)
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalNewStyleModVec2", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 StyleVar
		idx := imgui.StyleVar(0)
		v := parseVec2(vm, call.Argument(1))
		result := imgui.InternalNewStyleModVec2(idx, v)
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newStyle", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewStyle()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalNewTabBar", func(call goja.FunctionCall) goja.Value {
		result := imgui.InternalNewTabBar()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalNewTabItem", func(call goja.FunctionCall) goja.Value {
		result := imgui.InternalNewTabItem()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalNewTableColumnSettings", func(call goja.FunctionCall) goja.Value {
		result := imgui.InternalNewTableColumnSettings()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newTableColumnSortSpecs", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewTableColumnSortSpecs()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalNewTableColumn", func(call goja.FunctionCall) goja.Value {
		result := imgui.InternalNewTableColumn()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalNewTableInstanceData", func(call goja.FunctionCall) goja.Value {
		result := imgui.InternalNewTableInstanceData()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalNewTableSettings", func(call goja.FunctionCall) goja.Value {
		result := imgui.InternalNewTableSettings()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newTableSortSpecs", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewTableSortSpecs()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalNewTableTempData", func(call goja.FunctionCall) goja.Value {
		result := imgui.InternalNewTableTempData()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalNewTable", func(call goja.FunctionCall) goja.Value {
		result := imgui.InternalNewTable()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newTextBuffer", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewTextBuffer()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newTextFilter", func(call goja.FunctionCall) goja.Value {
		default_filter := call.Argument(0).String()
		result := imgui.NewTextFilter(default_filter)
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newTextRangeNil", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewTextRangeNil()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newTextRangeStr", func(call goja.FunctionCall) goja.Value {
		_b := call.Argument(0).String()
		_e := call.Argument(1).String()
		result := imgui.NewTextRangeStr(_b, _e)
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalNewTypingSelectState", func(call goja.FunctionCall) goja.Value {
		result := imgui.InternalNewTypingSelectState()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalNewViewportP", func(call goja.FunctionCall) goja.Value {
		result := imgui.InternalNewViewportP()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newViewport", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewViewport()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newWindowClass", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewWindowClass()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalNewWindowSettings", func(call goja.FunctionCall) goja.Value {
		result := imgui.InternalNewWindowSettings()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalNewWindow", func(call goja.FunctionCall) goja.Value {
		var context *imgui.Context
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.Context); ok {
				context = v
			}
		}
		name := call.Argument(1).String()
		result := imgui.InternalNewWindow(context, name)
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newTextureData", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewTextureData()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newTextureRefNil", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewTextureRefNil()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newTextureRefTextureID", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 TextureID
		tex_id := imgui.TextureID(0)
		result := imgui.NewTextureRefTextureID(tex_id)
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalNewVec1Float", func(call goja.FunctionCall) goja.Value {
		_x := float32(call.Argument(0).ToFloat())
		result := imgui.InternalNewVec1Float(_x)
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalNewVec1Nil", func(call goja.FunctionCall) goja.Value {
		result := imgui.InternalNewVec1Nil()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalNewVec2iInt", func(call goja.FunctionCall) goja.Value {
		_x := int32(call.Argument(0).ToInteger())
		_y := int32(call.Argument(1).ToInteger())
		result := imgui.InternalNewVec2iInt(_x, _y)
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalNewVec2iNil", func(call goja.FunctionCall) goja.Value {
		result := imgui.InternalNewVec2iNil()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newVec4Float", func(call goja.FunctionCall) goja.Value {
		_x := float32(call.Argument(0).ToFloat())
		_y := float32(call.Argument(1).ToFloat())
		_z := float32(call.Argument(2).ToFloat())
		_w := float32(call.Argument(3).ToFloat())
		result := imgui.NewVec4Float(_x, _y, _z, _w)
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("newVec4Nil", func(call goja.FunctionCall) goja.Value {
		result := imgui.NewVec4Nil()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("acceptDragDropPayloadV", func(call goja.FunctionCall) goja.Value {
		typeArg := call.Argument(0).String()
		// TODO: 处理类型 DragDropFlags
		flags := imgui.DragDropFlags(0)
		result := imgui.AcceptDragDropPayloadV(typeArg, flags)
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalActivateItemByID", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 ID
		id := imgui.ID(0)
		imgui.InternalActivateItemByID(id)
		return goja.Undefined()
	})
	imguiObj.Set("internalAddContextHook", func(call goja.FunctionCall) goja.Value {
		var context *imgui.Context
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.Context); ok {
				context = v
			}
		}
		// TODO: 处理指针类型 *ContextHook
		var hook *imgui.ContextHook
		return vm.ToValue(imgui.InternalAddContextHook(context, hook))
	})
	imguiObj.Set("internalAddSettingsHandler", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *SettingsHandler
		var handler *imgui.SettingsHandler
		imgui.InternalAddSettingsHandler(handler)
		return goja.Undefined()
	})
	imguiObj.Set("alignTextToFramePadding", func(call goja.FunctionCall) goja.Value {
		imgui.AlignTextToFramePadding()
		return goja.Undefined()
	})
	imguiObj.Set("arrowButton", func(call goja.FunctionCall) goja.Value {
		str_id := call.Argument(0).String()
		// TODO: 处理类型 Dir
		dir := imgui.Dir(0)
		return vm.ToValue(imgui.ArrowButton(str_id, dir))
	})
	imguiObj.Set("internalArrowButtonExV", func(call goja.FunctionCall) goja.Value {
		str_id := call.Argument(0).String()
		// TODO: 处理类型 Dir
		dir := imgui.Dir(0)
		size_arg := parseVec2(vm, call.Argument(2))
		// TODO: 处理类型 ButtonFlags
		flags := imgui.ButtonFlags(0)
		return vm.ToValue(imgui.InternalArrowButtonExV(str_id, dir, size_arg, flags))
	})
	imguiObj.Set("beginV", func(call goja.FunctionCall) goja.Value {
		name := call.Argument(0).String()
		p_open := false
		if len(call.Arguments) >= 2 {
			p_open = call.Argument(1).ToBoolean()
		}
		// TODO: 处理类型 WindowFlags
		flags := imgui.WindowFlags(0)
		return vm.ToValue(imgui.BeginV(name, &p_open, flags))
	})
	imguiObj.Set("internalBeginBoxSelect", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 Rect
		scope_rect := imgui.Rect{}
		// TODO: 处理指针类型 *Window
		var window *imgui.Window
		// TODO: 处理类型 ID
		box_select_id := imgui.ID(0)
		// TODO: 处理类型 MultiSelectFlags
		ms_flags := imgui.MultiSelectFlags(0)
		return vm.ToValue(imgui.InternalBeginBoxSelect(scope_rect, window, box_select_id, ms_flags))
	})
	imguiObj.Set("internalBeginChildEx", func(call goja.FunctionCall) goja.Value {
		name := call.Argument(0).String()
		// TODO: 处理类型 ID
		id := imgui.ID(0)
		size_arg := parseVec2(vm, call.Argument(2))
		// TODO: 处理类型 ChildFlags
		child_flags := imgui.ChildFlags(0)
		// TODO: 处理类型 WindowFlags
		window_flags := imgui.WindowFlags(0)
		return vm.ToValue(imgui.InternalBeginChildEx(name, id, size_arg, child_flags, window_flags))
	})
	imguiObj.Set("beginChildIDV", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 ID
		id := imgui.ID(0)
		size := parseVec2(vm, call.Argument(1))
		// TODO: 处理类型 ChildFlags
		child_flags := imgui.ChildFlags(0)
		// TODO: 处理类型 WindowFlags
		window_flags := imgui.WindowFlags(0)
		return vm.ToValue(imgui.BeginChildIDV(id, size, child_flags, window_flags))
	})
	imguiObj.Set("beginChildStrV", func(call goja.FunctionCall) goja.Value {
		str_id := call.Argument(0).String()
		size := parseVec2(vm, call.Argument(1))
		// TODO: 处理类型 ChildFlags
		child_flags := imgui.ChildFlags(0)
		// TODO: 处理类型 WindowFlags
		window_flags := imgui.WindowFlags(0)
		return vm.ToValue(imgui.BeginChildStrV(str_id, size, child_flags, window_flags))
	})
	imguiObj.Set("internalBeginColumnsV", func(call goja.FunctionCall) goja.Value {
		str_id := call.Argument(0).String()
		count := int32(call.Argument(1).ToInteger())
		// TODO: 处理类型 OldColumnFlags
		flags := imgui.OldColumnFlags(0)
		imgui.InternalBeginColumnsV(str_id, count, flags)
		return goja.Undefined()
	})
	imguiObj.Set("beginComboV", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		preview_value := call.Argument(1).String()
		// TODO: 处理类型 ComboFlags
		flags := imgui.ComboFlags(0)
		return vm.ToValue(imgui.BeginComboV(label, preview_value, flags))
	})
	imguiObj.Set("internalBeginComboPopup", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 ID
		popup_id := imgui.ID(0)
		// TODO: 处理类型 Rect
		bb := imgui.Rect{}
		// TODO: 处理类型 ComboFlags
		flags := imgui.ComboFlags(0)
		return vm.ToValue(imgui.InternalBeginComboPopup(popup_id, bb, flags))
	})
	imguiObj.Set("internalBeginComboPreview", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.InternalBeginComboPreview())
	})
	imguiObj.Set("beginDisabledV", func(call goja.FunctionCall) goja.Value {
		disabled := call.Argument(0).ToBoolean()
		imgui.BeginDisabledV(disabled)
		return goja.Undefined()
	})
	imguiObj.Set("internalBeginDisabledOverrideReenable", func(call goja.FunctionCall) goja.Value {
		imgui.InternalBeginDisabledOverrideReenable()
		return goja.Undefined()
	})
	imguiObj.Set("internalBeginDockableDragDropSource", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Window
		var window *imgui.Window
		imgui.InternalBeginDockableDragDropSource(window)
		return goja.Undefined()
	})
	imguiObj.Set("internalBeginDockableDragDropTarget", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Window
		var window *imgui.Window
		imgui.InternalBeginDockableDragDropTarget(window)
		return goja.Undefined()
	})
	imguiObj.Set("internalBeginDocked", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Window
		var window *imgui.Window
		p_open := false
		if len(call.Arguments) >= 2 {
			p_open = call.Argument(1).ToBoolean()
		}
		imgui.InternalBeginDocked(window, &p_open)
		return goja.Undefined()
	})
	imguiObj.Set("beginDragDropSourceV", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 DragDropFlags
		flags := imgui.DragDropFlags(0)
		return vm.ToValue(imgui.BeginDragDropSourceV(flags))
	})
	imguiObj.Set("beginDragDropTarget", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.BeginDragDropTarget())
	})
	imguiObj.Set("internalBeginDragDropTargetCustom", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 Rect
		bb := imgui.Rect{}
		// TODO: 处理类型 ID
		id := imgui.ID(0)
		return vm.ToValue(imgui.InternalBeginDragDropTargetCustom(bb, id))
	})
	imguiObj.Set("internalBeginDragDropTargetViewportV", func(call goja.FunctionCall) goja.Value {
		var viewport *imgui.Viewport
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.Viewport); ok {
				viewport = v
			}
		}
		// TODO: 处理指针类型 *Rect
		var p_bb *imgui.Rect
		return vm.ToValue(imgui.InternalBeginDragDropTargetViewportV(viewport, p_bb))
	})
	imguiObj.Set("internalBeginErrorTooltip", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.InternalBeginErrorTooltip())
	})
	imguiObj.Set("beginGroup", func(call goja.FunctionCall) goja.Value {
		imgui.BeginGroup()
		return goja.Undefined()
	})
	imguiObj.Set("beginItemTooltip", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.BeginItemTooltip())
	})
	imguiObj.Set("beginListBoxV", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		size := parseVec2(vm, call.Argument(1))
		return vm.ToValue(imgui.BeginListBoxV(label, size))
	})
	imguiObj.Set("beginMainMenuBar", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.BeginMainMenuBar())
	})
	imguiObj.Set("beginMenuV", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		enabled := call.Argument(1).ToBoolean()
		return vm.ToValue(imgui.BeginMenuV(label, enabled))
	})
	imguiObj.Set("beginMenuBar", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.BeginMenuBar())
	})
	imguiObj.Set("internalBeginMenuExV", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		icon := call.Argument(1).String()
		enabled := call.Argument(2).ToBoolean()
		return vm.ToValue(imgui.InternalBeginMenuExV(label, icon, enabled))
	})
	imguiObj.Set("beginMultiSelectV", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 MultiSelectFlags
		flags := imgui.MultiSelectFlags(0)
		selection_size := int32(call.Argument(1).ToInteger())
		items_count := int32(call.Argument(2).ToInteger())
		result := imgui.BeginMultiSelectV(flags, selection_size, items_count)
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("beginPopupV", func(call goja.FunctionCall) goja.Value {
		str_id := call.Argument(0).String()
		// TODO: 处理类型 WindowFlags
		flags := imgui.WindowFlags(0)
		return vm.ToValue(imgui.BeginPopupV(str_id, flags))
	})
	imguiObj.Set("beginPopupContextItemV", func(call goja.FunctionCall) goja.Value {
		str_id := call.Argument(0).String()
		// TODO: 处理类型 PopupFlags
		popup_flags := imgui.PopupFlags(0)
		return vm.ToValue(imgui.BeginPopupContextItemV(str_id, popup_flags))
	})
	imguiObj.Set("beginPopupContextVoidV", func(call goja.FunctionCall) goja.Value {
		str_id := call.Argument(0).String()
		// TODO: 处理类型 PopupFlags
		popup_flags := imgui.PopupFlags(0)
		return vm.ToValue(imgui.BeginPopupContextVoidV(str_id, popup_flags))
	})
	imguiObj.Set("beginPopupContextWindowV", func(call goja.FunctionCall) goja.Value {
		str_id := call.Argument(0).String()
		// TODO: 处理类型 PopupFlags
		popup_flags := imgui.PopupFlags(0)
		return vm.ToValue(imgui.BeginPopupContextWindowV(str_id, popup_flags))
	})
	imguiObj.Set("internalBeginPopupEx", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 ID
		id := imgui.ID(0)
		// TODO: 处理类型 WindowFlags
		extra_window_flags := imgui.WindowFlags(0)
		return vm.ToValue(imgui.InternalBeginPopupEx(id, extra_window_flags))
	})
	imguiObj.Set("internalBeginPopupMenuEx", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 ID
		id := imgui.ID(0)
		label := call.Argument(1).String()
		// TODO: 处理类型 WindowFlags
		extra_window_flags := imgui.WindowFlags(0)
		return vm.ToValue(imgui.InternalBeginPopupMenuEx(id, label, extra_window_flags))
	})
	imguiObj.Set("beginPopupModalV", func(call goja.FunctionCall) goja.Value {
		name := call.Argument(0).String()
		p_open := false
		if len(call.Arguments) >= 2 {
			p_open = call.Argument(1).ToBoolean()
		}
		// TODO: 处理类型 WindowFlags
		flags := imgui.WindowFlags(0)
		return vm.ToValue(imgui.BeginPopupModalV(name, &p_open, flags))
	})
	imguiObj.Set("beginTabBarV", func(call goja.FunctionCall) goja.Value {
		str_id := call.Argument(0).String()
		// TODO: 处理类型 TabBarFlags
		flags := imgui.TabBarFlags(0)
		return vm.ToValue(imgui.BeginTabBarV(str_id, flags))
	})
	imguiObj.Set("internalBeginTabBarEx", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *TabBar
		var tab_bar *imgui.TabBar
		// TODO: 处理类型 Rect
		bb := imgui.Rect{}
		// TODO: 处理类型 TabBarFlags
		flags := imgui.TabBarFlags(0)
		return vm.ToValue(imgui.InternalBeginTabBarEx(tab_bar, bb, flags))
	})
	imguiObj.Set("beginTabItemV", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		p_open := false
		if len(call.Arguments) >= 2 {
			p_open = call.Argument(1).ToBoolean()
		}
		// TODO: 处理类型 TabItemFlags
		flags := imgui.TabItemFlags(0)
		return vm.ToValue(imgui.BeginTabItemV(label, &p_open, flags))
	})
	imguiObj.Set("beginTableV", func(call goja.FunctionCall) goja.Value {
		str_id := call.Argument(0).String()
		columns := int32(call.Argument(1).ToInteger())
		// TODO: 处理类型 TableFlags
		flags := imgui.TableFlags(0)
		outer_size := parseVec2(vm, call.Argument(3))
		inner_width := float32(call.Argument(3).ToFloat())
		return vm.ToValue(imgui.BeginTableV(str_id, columns, flags, outer_size, inner_width))
	})
	imguiObj.Set("internalBeginTableExV", func(call goja.FunctionCall) goja.Value {
		name := call.Argument(0).String()
		// TODO: 处理类型 ID
		id := imgui.ID(0)
		columns_count := int32(call.Argument(2).ToInteger())
		// TODO: 处理类型 TableFlags
		flags := imgui.TableFlags(0)
		outer_size := parseVec2(vm, call.Argument(4))
		inner_width := float32(call.Argument(4).ToFloat())
		return vm.ToValue(imgui.InternalBeginTableExV(name, id, columns_count, flags, outer_size, inner_width))
	})
	imguiObj.Set("beginTooltip", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.BeginTooltip())
	})
	imguiObj.Set("internalBeginTooltipEx", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 TooltipFlags
		tooltip_flags := imgui.TooltipFlags(0)
		// TODO: 处理类型 WindowFlags
		extra_window_flags := imgui.WindowFlags(0)
		return vm.ToValue(imgui.InternalBeginTooltipEx(tooltip_flags, extra_window_flags))
	})
	imguiObj.Set("internalBeginTooltipHidden", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.InternalBeginTooltipHidden())
	})
	imguiObj.Set("internalBeginViewportSideBar", func(call goja.FunctionCall) goja.Value {
		name := call.Argument(0).String()
		var viewport *imgui.Viewport
		if len(call.Arguments) >= 2 {
			if v, ok := call.Argument(1).Export().(*imgui.Viewport); ok {
				viewport = v
			}
		}
		// TODO: 处理类型 Dir
		dir := imgui.Dir(0)
		size := float32(call.Argument(3).ToFloat())
		// TODO: 处理类型 WindowFlags
		window_flags := imgui.WindowFlags(0)
		return vm.ToValue(imgui.InternalBeginViewportSideBar(name, viewport, dir, size, window_flags))
	})
	imguiObj.Set("internalBringWindowToDisplayBack", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Window
		var window *imgui.Window
		imgui.InternalBringWindowToDisplayBack(window)
		return goja.Undefined()
	})
	imguiObj.Set("internalBringWindowToDisplayBehind", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Window
		var window *imgui.Window
		// TODO: 处理指针类型 *Window
		var above_window *imgui.Window
		imgui.InternalBringWindowToDisplayBehind(window, above_window)
		return goja.Undefined()
	})
	imguiObj.Set("internalBringWindowToDisplayFront", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Window
		var window *imgui.Window
		imgui.InternalBringWindowToDisplayFront(window)
		return goja.Undefined()
	})
	imguiObj.Set("internalBringWindowToFocusFront", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Window
		var window *imgui.Window
		imgui.InternalBringWindowToFocusFront(window)
		return goja.Undefined()
	})
	imguiObj.Set("bullet", func(call goja.FunctionCall) goja.Value {
		imgui.Bullet()
		return goja.Undefined()
	})
	imguiObj.Set("bulletText", func(call goja.FunctionCall) goja.Value {
		fmt := call.Argument(0).String()
		imgui.BulletText(fmt)
		return goja.Undefined()
	})
	imguiObj.Set("buttonV", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		size := parseVec2(vm, call.Argument(1))
		return vm.ToValue(imgui.ButtonV(label, size))
	})
	imguiObj.Set("internalButtonBehaviorV", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 Rect
		bb := imgui.Rect{}
		// TODO: 处理类型 ID
		id := imgui.ID(0)
		out_hovered := false
		if len(call.Arguments) >= 3 {
			out_hovered = call.Argument(2).ToBoolean()
		}
		out_held := false
		if len(call.Arguments) >= 4 {
			out_held = call.Argument(3).ToBoolean()
		}
		// TODO: 处理类型 ButtonFlags
		flags := imgui.ButtonFlags(0)
		return vm.ToValue(imgui.InternalButtonBehaviorV(bb, id, &out_hovered, &out_held, flags))
	})
	imguiObj.Set("internalButtonExV", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		size_arg := parseVec2(vm, call.Argument(1))
		// TODO: 处理类型 ButtonFlags
		flags := imgui.ButtonFlags(0)
		return vm.ToValue(imgui.InternalButtonExV(label, size_arg, flags))
	})
	imguiObj.Set("internalCalcClipRectVisibleItemsY", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 Rect
		clip_rect := imgui.Rect{}
		pos := parseVec2(vm, call.Argument(1))
		items_height := float32(call.Argument(1).ToFloat())
		out_visible_start := int32(0)
		if len(call.Arguments) >= 3 {
			out_visible_start = int32(call.Argument(2).ToInteger())
		}
		out_visible_end := int32(0)
		if len(call.Arguments) >= 4 {
			out_visible_end = int32(call.Argument(3).ToInteger())
		}
		imgui.InternalCalcClipRectVisibleItemsY(clip_rect, pos, items_height, &out_visible_start, &out_visible_end)
		return goja.Undefined()
	})
	imguiObj.Set("internalCalcItemSize", func(call goja.FunctionCall) goja.Value {
		size := parseVec2(vm, call.Argument(0))
		default_w := float32(call.Argument(0).ToFloat())
		default_h := float32(call.Argument(1).ToFloat())
		result := imgui.InternalCalcItemSize(size, default_w, default_h)
		return vm.ToValue(result)
	})
	imguiObj.Set("calcItemWidth", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.CalcItemWidth())
	})
	imguiObj.Set("internalCalcRoundingFlagsForRectInRect", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 Rect
		r_in := imgui.Rect{}
		// TODO: 处理类型 Rect
		r_outer := imgui.Rect{}
		threshold := float32(call.Argument(2).ToFloat())
		return vm.ToValue(imgui.InternalCalcRoundingFlagsForRectInRect(r_in, r_outer, threshold))
	})
	imguiObj.Set("calcTextSizeV", func(call goja.FunctionCall) goja.Value {
		text := call.Argument(0).String()
		hide_text_after_double_hash := call.Argument(1).ToBoolean()
		wrap_width := float32(call.Argument(2).ToFloat())
		result := imgui.CalcTextSizeV(text, hide_text_after_double_hash, wrap_width)
		return vm.ToValue(result)
	})
	imguiObj.Set("internalCalcTypematicRepeatAmount", func(call goja.FunctionCall) goja.Value {
		t0 := float32(call.Argument(0).ToFloat())
		t1 := float32(call.Argument(1).ToFloat())
		repeat_delay := float32(call.Argument(2).ToFloat())
		repeat_rate := float32(call.Argument(3).ToFloat())
		return vm.ToValue(imgui.InternalCalcTypematicRepeatAmount(t0, t1, repeat_delay, repeat_rate))
	})
	imguiObj.Set("internalCalcWindowNextAutoFitSize", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Window
		var window *imgui.Window
		result := imgui.InternalCalcWindowNextAutoFitSize(window)
		return vm.ToValue(result)
	})
	imguiObj.Set("internalCalcWrapWidthForPos", func(call goja.FunctionCall) goja.Value {
		pos := parseVec2(vm, call.Argument(0))
		wrap_pos_x := float32(call.Argument(0).ToFloat())
		return vm.ToValue(imgui.InternalCalcWrapWidthForPos(pos, wrap_pos_x))
	})
	imguiObj.Set("internalCallContextHooks", func(call goja.FunctionCall) goja.Value {
		var context *imgui.Context
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.Context); ok {
				context = v
			}
		}
		// TODO: 处理类型 ContextHookType
		typeArg := imgui.ContextHookType(0)
		imgui.InternalCallContextHooks(context, typeArg)
		return goja.Undefined()
	})
	imguiObj.Set("checkbox", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		v := false
		if len(call.Arguments) >= 2 {
			v = call.Argument(1).ToBoolean()
		}
		return vm.ToValue(imgui.Checkbox(label, &v))
	})
	imguiObj.Set("checkboxFlagsIntPtr", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		flags := int32(0)
		if len(call.Arguments) >= 2 {
			flags = int32(call.Argument(1).ToInteger())
		}
		flags_value := int32(call.Argument(2).ToInteger())
		return vm.ToValue(imgui.CheckboxFlagsIntPtr(label, &flags, flags_value))
	})
	imguiObj.Set("internalCheckboxFlagsS64Ptr", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		// TODO: 处理指针类型 *int64
		var flags *int64
		// TODO: 处理类型 int64
		flags_value := int64(0)
		return vm.ToValue(imgui.InternalCheckboxFlagsS64Ptr(label, flags, flags_value))
	})
	imguiObj.Set("internalCheckboxFlagsU64Ptr", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		// TODO: 处理指针类型 *uint64
		var flags *uint64
		flags_value := uint64(call.Argument(2).ToInteger())
		return vm.ToValue(imgui.InternalCheckboxFlagsU64Ptr(label, flags, flags_value))
	})
	imguiObj.Set("checkboxFlagsUintPtr", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		flags := uint32(0)
		if len(call.Arguments) >= 2 {
			flags = uint32(call.Argument(1).ToInteger())
		}
		flags_value := uint32(call.Argument(2).ToInteger())
		return vm.ToValue(imgui.CheckboxFlagsUintPtr(label, &flags, flags_value))
	})
	imguiObj.Set("internalClearActiveID", func(call goja.FunctionCall) goja.Value {
		imgui.InternalClearActiveID()
		return goja.Undefined()
	})
	imguiObj.Set("internalClearDragDrop", func(call goja.FunctionCall) goja.Value {
		imgui.InternalClearDragDrop()
		return goja.Undefined()
	})
	imguiObj.Set("internalClearIniSettings", func(call goja.FunctionCall) goja.Value {
		imgui.InternalClearIniSettings()
		return goja.Undefined()
	})
	imguiObj.Set("internalClearWindowSettings", func(call goja.FunctionCall) goja.Value {
		name := call.Argument(0).String()
		imgui.InternalClearWindowSettings(name)
		return goja.Undefined()
	})
	imguiObj.Set("internalCloseButton", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 ID
		id := imgui.ID(0)
		pos := parseVec2(vm, call.Argument(1))
		return vm.ToValue(imgui.InternalCloseButton(id, pos))
	})
	imguiObj.Set("closeCurrentPopup", func(call goja.FunctionCall) goja.Value {
		imgui.CloseCurrentPopup()
		return goja.Undefined()
	})
	imguiObj.Set("internalClosePopupToLevel", func(call goja.FunctionCall) goja.Value {
		remaining := int32(call.Argument(0).ToInteger())
		restore_focus_to_window_under_popup := call.Argument(1).ToBoolean()
		imgui.InternalClosePopupToLevel(remaining, restore_focus_to_window_under_popup)
		return goja.Undefined()
	})
	imguiObj.Set("internalClosePopupsExceptModals", func(call goja.FunctionCall) goja.Value {
		imgui.InternalClosePopupsExceptModals()
		return goja.Undefined()
	})
	imguiObj.Set("internalClosePopupsOverWindow", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Window
		var ref_window *imgui.Window
		restore_focus_to_window_under_popup := call.Argument(1).ToBoolean()
		imgui.InternalClosePopupsOverWindow(ref_window, restore_focus_to_window_under_popup)
		return goja.Undefined()
	})
	imguiObj.Set("internalCollapseButton", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 ID
		id := imgui.ID(0)
		pos := parseVec2(vm, call.Argument(1))
		// TODO: 处理指针类型 *DockNode
		var dock_node *imgui.DockNode
		return vm.ToValue(imgui.InternalCollapseButton(id, pos, dock_node))
	})
	imguiObj.Set("collapsingHeaderBoolPtrV", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		p_visible := false
		if len(call.Arguments) >= 2 {
			p_visible = call.Argument(1).ToBoolean()
		}
		// TODO: 处理类型 TreeNodeFlags
		flags := imgui.TreeNodeFlags(0)
		return vm.ToValue(imgui.CollapsingHeaderBoolPtrV(label, &p_visible, flags))
	})
	imguiObj.Set("collapsingHeaderTreeNodeFlagsV", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		// TODO: 处理类型 TreeNodeFlags
		flags := imgui.TreeNodeFlags(0)
		return vm.ToValue(imgui.CollapsingHeaderTreeNodeFlagsV(label, flags))
	})
	imguiObj.Set("colorButtonV", func(call goja.FunctionCall) goja.Value {
		desc_id := call.Argument(0).String()
		col := parseVec4(vm, call.Argument(1))
		// TODO: 处理类型 ColorEditFlags
		flags := imgui.ColorEditFlags(0)
		size := parseVec2(vm, call.Argument(2))
		return vm.ToValue(imgui.ColorButtonV(desc_id, col, flags, size))
	})
	imguiObj.Set("colorConvertFloat4ToU32", func(call goja.FunctionCall) goja.Value {
		in := parseVec4(vm, call.Argument(0))
		return vm.ToValue(imgui.ColorConvertFloat4ToU32(in))
	})
	imguiObj.Set("colorConvertHSVtoRGB", func(call goja.FunctionCall) goja.Value {
		h := float32(call.Argument(0).ToFloat())
		s := float32(call.Argument(1).ToFloat())
		v := float32(call.Argument(2).ToFloat())
		out_r := float32(0)
		if len(call.Arguments) >= 4 {
			out_r = float32(call.Argument(3).ToFloat())
		}
		out_g := float32(0)
		if len(call.Arguments) >= 5 {
			out_g = float32(call.Argument(4).ToFloat())
		}
		out_b := float32(0)
		if len(call.Arguments) >= 6 {
			out_b = float32(call.Argument(5).ToFloat())
		}
		imgui.ColorConvertHSVtoRGB(h, s, v, &out_r, &out_g, &out_b)
		return goja.Undefined()
	})
	imguiObj.Set("colorConvertRGBtoHSV", func(call goja.FunctionCall) goja.Value {
		r := float32(call.Argument(0).ToFloat())
		g := float32(call.Argument(1).ToFloat())
		b := float32(call.Argument(2).ToFloat())
		out_h := float32(0)
		if len(call.Arguments) >= 4 {
			out_h = float32(call.Argument(3).ToFloat())
		}
		out_s := float32(0)
		if len(call.Arguments) >= 5 {
			out_s = float32(call.Argument(4).ToFloat())
		}
		out_v := float32(0)
		if len(call.Arguments) >= 6 {
			out_v = float32(call.Argument(5).ToFloat())
		}
		imgui.ColorConvertRGBtoHSV(r, g, b, &out_h, &out_s, &out_v)
		return goja.Undefined()
	})
	imguiObj.Set("colorConvertU32ToFloat4", func(call goja.FunctionCall) goja.Value {
		in := uint32(call.Argument(0).ToInteger())
		result := imgui.ColorConvertU32ToFloat4(in)
		return vm.ToValue(result)
	})
	imguiObj.Set("colorEdit3V", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		col := [3]float32{0, 0, 0}
		// TODO: 处理类型 ColorEditFlags
		flags := imgui.ColorEditFlags(0)
		return vm.ToValue(imgui.ColorEdit3V(label, &col, flags))
	})
	imguiObj.Set("colorEdit4V", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		col := [4]float32{0, 0, 0, 0}
		// TODO: 处理类型 ColorEditFlags
		flags := imgui.ColorEditFlags(0)
		return vm.ToValue(imgui.ColorEdit4V(label, &col, flags))
	})
	imguiObj.Set("internalColorEditOptionsPopup", func(call goja.FunctionCall) goja.Value {
		col := float32(0)
		if len(call.Arguments) >= 1 {
			col = float32(call.Argument(0).ToFloat())
		}
		// TODO: 处理类型 ColorEditFlags
		flags := imgui.ColorEditFlags(0)
		imgui.InternalColorEditOptionsPopup(&col, flags)
		return goja.Undefined()
	})
	imguiObj.Set("colorPicker3V", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		col := [3]float32{0, 0, 0}
		// TODO: 处理类型 ColorEditFlags
		flags := imgui.ColorEditFlags(0)
		return vm.ToValue(imgui.ColorPicker3V(label, &col, flags))
	})
	imguiObj.Set("colorPicker4V", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		col := [4]float32{0, 0, 0, 0}
		// TODO: 处理类型 ColorEditFlags
		flags := imgui.ColorEditFlags(0)
		ref_col := float32(0)
		if len(call.Arguments) >= 4 {
			ref_col = float32(call.Argument(3).ToFloat())
		}
		return vm.ToValue(imgui.ColorPicker4V(label, &col, flags, &ref_col))
	})
	imguiObj.Set("internalColorPickerOptionsPopup", func(call goja.FunctionCall) goja.Value {
		ref_col := float32(0)
		if len(call.Arguments) >= 1 {
			ref_col = float32(call.Argument(0).ToFloat())
		}
		// TODO: 处理类型 ColorEditFlags
		flags := imgui.ColorEditFlags(0)
		imgui.InternalColorPickerOptionsPopup(&ref_col, flags)
		return goja.Undefined()
	})
	imguiObj.Set("internalColorTooltip", func(call goja.FunctionCall) goja.Value {
		text := call.Argument(0).String()
		col := float32(0)
		if len(call.Arguments) >= 2 {
			col = float32(call.Argument(1).ToFloat())
		}
		// TODO: 处理类型 ColorEditFlags
		flags := imgui.ColorEditFlags(0)
		imgui.InternalColorTooltip(text, &col, flags)
		return goja.Undefined()
	})
	imguiObj.Set("columnsV", func(call goja.FunctionCall) goja.Value {
		count := int32(call.Argument(0).ToInteger())
		id := call.Argument(1).String()
		borders := call.Argument(2).ToBoolean()
		imgui.ColumnsV(count, id, borders)
		return goja.Undefined()
	})
	imguiObj.Set("comboStrV", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		current_item := int32(0)
		if len(call.Arguments) >= 2 {
			current_item = int32(call.Argument(1).ToInteger())
		}
		items_separated_by_zeros := call.Argument(2).String()
		popup_max_height_in_items := int32(call.Argument(3).ToInteger())
		return vm.ToValue(imgui.ComboStrV(label, &current_item, items_separated_by_zeros, popup_max_height_in_items))
	})
	imguiObj.Set("comboStrarrV", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		current_item := int32(0)
		if len(call.Arguments) >= 2 {
			current_item = int32(call.Argument(1).ToInteger())
		}
		items := parseStringArray(vm, call.Argument(2))
		items_count := int32(call.Argument(3).ToInteger())
		popup_max_height_in_items := int32(call.Argument(4).ToInteger())
		return vm.ToValue(imgui.ComboStrarrV(label, &current_item, items, items_count, popup_max_height_in_items))
	})
	imguiObj.Set("internalConvertSingleModFlagToKey", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 Key
		key := imgui.Key(0)
		return vm.ToValue(imgui.InternalConvertSingleModFlagToKey(key))
	})
	imguiObj.Set("createContextV", func(call goja.FunctionCall) goja.Value {
		var shared_font_atlas *imgui.FontAtlas
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.FontAtlas); ok {
				shared_font_atlas = v
			}
		}
		result := imgui.CreateContextV(shared_font_atlas)
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalCreateNewWindowSettings", func(call goja.FunctionCall) goja.Value {
		name := call.Argument(0).String()
		result := imgui.InternalCreateNewWindowSettings(name)
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalDataTypeApplyFromTextV", func(call goja.FunctionCall) goja.Value {
		buf := call.Argument(0).String()
		// TODO: 处理类型 DataType
		data_type := imgui.DataType(0)
		p_data := uintptr(call.Argument(2).ToInteger())
		format := call.Argument(3).String()
		p_data_when_empty := uintptr(call.Argument(4).ToInteger())
		return vm.ToValue(imgui.InternalDataTypeApplyFromTextV(buf, data_type, p_data, format, p_data_when_empty))
	})
	imguiObj.Set("internalDataTypeApplyOp", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 DataType
		data_type := imgui.DataType(0)
		op := int32(call.Argument(1).ToInteger())
		output := uintptr(call.Argument(2).ToInteger())
		arg_1 := uintptr(call.Argument(3).ToInteger())
		arg_2 := uintptr(call.Argument(4).ToInteger())
		imgui.InternalDataTypeApplyOp(data_type, op, output, arg_1, arg_2)
		return goja.Undefined()
	})
	imguiObj.Set("internalDataTypeClamp", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 DataType
		data_type := imgui.DataType(0)
		p_data := uintptr(call.Argument(1).ToInteger())
		p_min := uintptr(call.Argument(2).ToInteger())
		p_max := uintptr(call.Argument(3).ToInteger())
		return vm.ToValue(imgui.InternalDataTypeClamp(data_type, p_data, p_min, p_max))
	})
	imguiObj.Set("internalDataTypeCompare", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 DataType
		data_type := imgui.DataType(0)
		arg_1 := uintptr(call.Argument(1).ToInteger())
		arg_2 := uintptr(call.Argument(2).ToInteger())
		return vm.ToValue(imgui.InternalDataTypeCompare(data_type, arg_1, arg_2))
	})
	imguiObj.Set("internalDataTypeFormatString", func(call goja.FunctionCall) goja.Value {
		buf := call.Argument(0).String()
		buf_size := int32(call.Argument(1).ToInteger())
		// TODO: 处理类型 DataType
		data_type := imgui.DataType(0)
		p_data := uintptr(call.Argument(3).ToInteger())
		format := call.Argument(4).String()
		return vm.ToValue(imgui.InternalDataTypeFormatString(buf, buf_size, data_type, p_data, format))
	})
	imguiObj.Set("internalDataTypeGetInfo", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 DataType
		data_type := imgui.DataType(0)
		result := imgui.InternalDataTypeGetInfo(data_type)
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalDataTypeIsZero", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 DataType
		data_type := imgui.DataType(0)
		p_data := uintptr(call.Argument(1).ToInteger())
		return vm.ToValue(imgui.InternalDataTypeIsZero(data_type, p_data))
	})
	imguiObj.Set("internalDebugAllocHook", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *DebugAllocInfo
		var info *imgui.DebugAllocInfo
		frame_count := int32(call.Argument(1).ToInteger())
		ptr := uintptr(call.Argument(2).ToInteger())
		size := uint64(call.Argument(3).ToInteger())
		imgui.InternalDebugAllocHook(info, frame_count, ptr, size)
		return goja.Undefined()
	})
	imguiObj.Set("internalDebugBreakButton", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		description_of_location := call.Argument(1).String()
		return vm.ToValue(imgui.InternalDebugBreakButton(label, description_of_location))
	})
	imguiObj.Set("internalDebugBreakButtonTooltip", func(call goja.FunctionCall) goja.Value {
		keyboard_only := call.Argument(0).ToBoolean()
		description_of_location := call.Argument(1).String()
		imgui.InternalDebugBreakButtonTooltip(keyboard_only, description_of_location)
		return goja.Undefined()
	})
	imguiObj.Set("internalDebugBreakClearData", func(call goja.FunctionCall) goja.Value {
		imgui.InternalDebugBreakClearData()
		return goja.Undefined()
	})
	imguiObj.Set("debugCheckVersionAndDataLayout", func(call goja.FunctionCall) goja.Value {
		version_str := call.Argument(0).String()
		sz_io := uint64(call.Argument(1).ToInteger())
		sz_style := uint64(call.Argument(2).ToInteger())
		sz_vec2 := uint64(call.Argument(3).ToInteger())
		sz_vec4 := uint64(call.Argument(4).ToInteger())
		sz_drawvert := uint64(call.Argument(5).ToInteger())
		sz_drawidx := uint64(call.Argument(6).ToInteger())
		return vm.ToValue(imgui.DebugCheckVersionAndDataLayout(version_str, sz_io, sz_style, sz_vec2, sz_vec4, sz_drawvert, sz_drawidx))
	})
	imguiObj.Set("internalDebugDrawCursorPosV", func(call goja.FunctionCall) goja.Value {
		col := uint32(call.Argument(0).ToInteger())
		imgui.InternalDebugDrawCursorPosV(col)
		return goja.Undefined()
	})
	imguiObj.Set("internalDebugDrawItemRectV", func(call goja.FunctionCall) goja.Value {
		col := uint32(call.Argument(0).ToInteger())
		imgui.InternalDebugDrawItemRectV(col)
		return goja.Undefined()
	})
	imguiObj.Set("internalDebugDrawLineExtentsV", func(call goja.FunctionCall) goja.Value {
		col := uint32(call.Argument(0).ToInteger())
		imgui.InternalDebugDrawLineExtentsV(col)
		return goja.Undefined()
	})
	imguiObj.Set("debugFlashStyleColor", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 Col
		idx := imgui.Col(0)
		imgui.DebugFlashStyleColor(idx)
		return goja.Undefined()
	})
	imguiObj.Set("internalDebugHookIdInfo", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 ID
		id := imgui.ID(0)
		// TODO: 处理类型 DataType
		data_type := imgui.DataType(0)
		data_id := uintptr(call.Argument(2).ToInteger())
		data_id_end := uintptr(call.Argument(3).ToInteger())
		imgui.InternalDebugHookIdInfo(id, data_type, data_id, data_id_end)
		return goja.Undefined()
	})
	imguiObj.Set("internalDebugLocateItem", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 ID
		target_id := imgui.ID(0)
		imgui.InternalDebugLocateItem(target_id)
		return goja.Undefined()
	})
	imguiObj.Set("internalDebugLocateItemOnHover", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 ID
		target_id := imgui.ID(0)
		imgui.InternalDebugLocateItemOnHover(target_id)
		return goja.Undefined()
	})
	imguiObj.Set("internalDebugLocateItemResolveWithLastItem", func(call goja.FunctionCall) goja.Value {
		imgui.InternalDebugLocateItemResolveWithLastItem()
		return goja.Undefined()
	})
	imguiObj.Set("debugLog", func(call goja.FunctionCall) goja.Value {
		fmt := call.Argument(0).String()
		imgui.DebugLog(fmt)
		return goja.Undefined()
	})
	imguiObj.Set("internalDebugNodeColumns", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *OldColumns
		var columns *imgui.OldColumns
		imgui.InternalDebugNodeColumns(columns)
		return goja.Undefined()
	})
	imguiObj.Set("internalDebugNodeDockNode", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *DockNode
		var node *imgui.DockNode
		label := call.Argument(1).String()
		imgui.InternalDebugNodeDockNode(node, label)
		return goja.Undefined()
	})
	imguiObj.Set("internalDebugNodeDrawCmdShowMeshAndBoundingBox", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *DrawList
		var out_draw_list *imgui.DrawList
		// TODO: 处理指针类型 *DrawList
		var draw_list *imgui.DrawList
		// TODO: 处理指针类型 *DrawCmd
		var draw_cmd *imgui.DrawCmd
		show_mesh := call.Argument(3).ToBoolean()
		show_aabb := call.Argument(4).ToBoolean()
		imgui.InternalDebugNodeDrawCmdShowMeshAndBoundingBox(out_draw_list, draw_list, draw_cmd, show_mesh, show_aabb)
		return goja.Undefined()
	})
	imguiObj.Set("internalDebugNodeDrawList", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Window
		var window *imgui.Window
		// TODO: 处理指针类型 *ViewportP
		var viewport *imgui.ViewportP
		// TODO: 处理指针类型 *DrawList
		var draw_list *imgui.DrawList
		label := call.Argument(3).String()
		imgui.InternalDebugNodeDrawList(window, viewport, draw_list, label)
		return goja.Undefined()
	})
	imguiObj.Set("internalDebugNodeFont", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Font
		var font *imgui.Font
		imgui.InternalDebugNodeFont(font)
		return goja.Undefined()
	})
	imguiObj.Set("internalDebugNodeFontGlyph", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Font
		var font *imgui.Font
		// TODO: 处理指针类型 *FontGlyph
		var glyph *imgui.FontGlyph
		imgui.InternalDebugNodeFontGlyph(font, glyph)
		return goja.Undefined()
	})
	imguiObj.Set("internalDebugNodeFontGlyphesForSrcMask", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Font
		var font *imgui.Font
		// TODO: 处理指针类型 *FontBaked
		var baked *imgui.FontBaked
		src_mask := int32(call.Argument(2).ToInteger())
		imgui.InternalDebugNodeFontGlyphesForSrcMask(font, baked, src_mask)
		return goja.Undefined()
	})
	imguiObj.Set("internalDebugNodeInputTextState", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *InputTextState
		var state *imgui.InputTextState
		imgui.InternalDebugNodeInputTextState(state)
		return goja.Undefined()
	})
	imguiObj.Set("internalDebugNodeMultiSelectState", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *MultiSelectState
		var state *imgui.MultiSelectState
		imgui.InternalDebugNodeMultiSelectState(state)
		return goja.Undefined()
	})
	imguiObj.Set("internalDebugNodePlatformMonitor", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *PlatformMonitor
		var monitor *imgui.PlatformMonitor
		label := call.Argument(1).String()
		idx := int32(call.Argument(2).ToInteger())
		imgui.InternalDebugNodePlatformMonitor(monitor, label, idx)
		return goja.Undefined()
	})
	imguiObj.Set("internalDebugNodeStorage", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Storage
		var storage *imgui.Storage
		label := call.Argument(1).String()
		imgui.InternalDebugNodeStorage(storage, label)
		return goja.Undefined()
	})
	imguiObj.Set("internalDebugNodeTabBar", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *TabBar
		var tab_bar *imgui.TabBar
		label := call.Argument(1).String()
		imgui.InternalDebugNodeTabBar(tab_bar, label)
		return goja.Undefined()
	})
	imguiObj.Set("internalDebugNodeTable", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Table
		var table *imgui.Table
		imgui.InternalDebugNodeTable(table)
		return goja.Undefined()
	})
	imguiObj.Set("internalDebugNodeTableSettings", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *TableSettings
		var settings *imgui.TableSettings
		imgui.InternalDebugNodeTableSettings(settings)
		return goja.Undefined()
	})
	imguiObj.Set("internalDebugNodeTextureV", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *TextureData
		var tex *imgui.TextureData
		int_id := int32(call.Argument(1).ToInteger())
		// TODO: 处理指针类型 *FontAtlasRect
		var highlight_rect *imgui.FontAtlasRect
		imgui.InternalDebugNodeTextureV(tex, int_id, highlight_rect)
		return goja.Undefined()
	})
	imguiObj.Set("internalDebugNodeTypingSelectState", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *TypingSelectState
		var state *imgui.TypingSelectState
		imgui.InternalDebugNodeTypingSelectState(state)
		return goja.Undefined()
	})
	imguiObj.Set("internalDebugNodeViewport", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *ViewportP
		var viewport *imgui.ViewportP
		imgui.InternalDebugNodeViewport(viewport)
		return goja.Undefined()
	})
	imguiObj.Set("internalDebugNodeWindow", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Window
		var window *imgui.Window
		label := call.Argument(1).String()
		imgui.InternalDebugNodeWindow(window, label)
		return goja.Undefined()
	})
	imguiObj.Set("internalDebugNodeWindowSettings", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *WindowSettings
		var settings *imgui.WindowSettings
		imgui.InternalDebugNodeWindowSettings(settings)
		return goja.Undefined()
	})
	imguiObj.Set("internalDebugRenderKeyboardPreview", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *DrawList
		var draw_list *imgui.DrawList
		imgui.InternalDebugRenderKeyboardPreview(draw_list)
		return goja.Undefined()
	})
	imguiObj.Set("internalDebugRenderViewportThumbnail", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *DrawList
		var draw_list *imgui.DrawList
		// TODO: 处理指针类型 *ViewportP
		var viewport *imgui.ViewportP
		// TODO: 处理类型 Rect
		bb := imgui.Rect{}
		imgui.InternalDebugRenderViewportThumbnail(draw_list, viewport, bb)
		return goja.Undefined()
	})
	imguiObj.Set("debugStartItemPicker", func(call goja.FunctionCall) goja.Value {
		imgui.DebugStartItemPicker()
		return goja.Undefined()
	})
	imguiObj.Set("debugTextEncoding", func(call goja.FunctionCall) goja.Value {
		text := call.Argument(0).String()
		imgui.DebugTextEncoding(text)
		return goja.Undefined()
	})
	imguiObj.Set("internalDebugTextUnformattedWithLocateItem", func(call goja.FunctionCall) goja.Value {
		line_begin := call.Argument(0).String()
		line_end := call.Argument(1).String()
		imgui.InternalDebugTextUnformattedWithLocateItem(line_begin, line_end)
		return goja.Undefined()
	})
	imguiObj.Set("destroyContextV", func(call goja.FunctionCall) goja.Value {
		var ctx *imgui.Context
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.Context); ok {
				ctx = v
			}
		}
		imgui.DestroyContextV(ctx)
		return goja.Undefined()
	})
	imguiObj.Set("internalDestroyPlatformWindow", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *ViewportP
		var viewport *imgui.ViewportP
		imgui.InternalDestroyPlatformWindow(viewport)
		return goja.Undefined()
	})
	imguiObj.Set("destroyPlatformWindows", func(call goja.FunctionCall) goja.Value {
		imgui.DestroyPlatformWindows()
		return goja.Undefined()
	})
	imguiObj.Set("internalDockBuilderAddNodeV", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 ID
		node_id := imgui.ID(0)
		// TODO: 处理类型 DockNodeFlags
		flags := imgui.DockNodeFlags(0)
		return vm.ToValue(imgui.InternalDockBuilderAddNodeV(node_id, flags))
	})
	imguiObj.Set("internalDockBuilderCopyWindowSettings", func(call goja.FunctionCall) goja.Value {
		src_name := call.Argument(0).String()
		dst_name := call.Argument(1).String()
		imgui.InternalDockBuilderCopyWindowSettings(src_name, dst_name)
		return goja.Undefined()
	})
	imguiObj.Set("internalDockBuilderDockWindow", func(call goja.FunctionCall) goja.Value {
		window_name := call.Argument(0).String()
		// TODO: 处理类型 ID
		node_id := imgui.ID(0)
		imgui.InternalDockBuilderDockWindow(window_name, node_id)
		return goja.Undefined()
	})
	imguiObj.Set("internalDockBuilderFinish", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 ID
		node_id := imgui.ID(0)
		imgui.InternalDockBuilderFinish(node_id)
		return goja.Undefined()
	})
	imguiObj.Set("internalDockBuilderGetCentralNode", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 ID
		node_id := imgui.ID(0)
		result := imgui.InternalDockBuilderGetCentralNode(node_id)
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalDockBuilderGetNode", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 ID
		node_id := imgui.ID(0)
		result := imgui.InternalDockBuilderGetNode(node_id)
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalDockBuilderRemoveNode", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 ID
		node_id := imgui.ID(0)
		imgui.InternalDockBuilderRemoveNode(node_id)
		return goja.Undefined()
	})
	imguiObj.Set("internalDockBuilderRemoveNodeChildNodes", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 ID
		node_id := imgui.ID(0)
		imgui.InternalDockBuilderRemoveNodeChildNodes(node_id)
		return goja.Undefined()
	})
	imguiObj.Set("internalDockBuilderRemoveNodeDockedWindowsV", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 ID
		node_id := imgui.ID(0)
		clear_settings_refs := call.Argument(1).ToBoolean()
		imgui.InternalDockBuilderRemoveNodeDockedWindowsV(node_id, clear_settings_refs)
		return goja.Undefined()
	})
	imguiObj.Set("internalDockBuilderSetNodePos", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 ID
		node_id := imgui.ID(0)
		pos := parseVec2(vm, call.Argument(1))
		imgui.InternalDockBuilderSetNodePos(node_id, pos)
		return goja.Undefined()
	})
	imguiObj.Set("internalDockBuilderSetNodeSize", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 ID
		node_id := imgui.ID(0)
		size := parseVec2(vm, call.Argument(1))
		imgui.InternalDockBuilderSetNodeSize(node_id, size)
		return goja.Undefined()
	})
	imguiObj.Set("internalDockBuilderSplitNode", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 ID
		node_id := imgui.ID(0)
		// TODO: 处理类型 Dir
		split_dir := imgui.Dir(0)
		size_ratio_for_node_at_dir := float32(call.Argument(2).ToFloat())
		// TODO: 处理指针类型 *ID
		var out_id_at_dir *imgui.ID
		// TODO: 处理指针类型 *ID
		var out_id_at_opposite_dir *imgui.ID
		return vm.ToValue(imgui.InternalDockBuilderSplitNode(node_id, split_dir, size_ratio_for_node_at_dir, out_id_at_dir, out_id_at_opposite_dir))
	})
	imguiObj.Set("internalDockContextCalcDropPosForDocking", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Window
		var target *imgui.Window
		// TODO: 处理指针类型 *DockNode
		var target_node *imgui.DockNode
		// TODO: 处理指针类型 *Window
		var payload_window *imgui.Window
		// TODO: 处理指针类型 *DockNode
		var payload_node *imgui.DockNode
		// TODO: 处理类型 Dir
		split_dir := imgui.Dir(0)
		split_outer := call.Argument(5).ToBoolean()
		out_pos := parseVec2(vm, call.Argument(6))
		return vm.ToValue(imgui.InternalDockContextCalcDropPosForDocking(target, target_node, payload_window, payload_node, split_dir, split_outer, &out_pos))
	})
	imguiObj.Set("internalDockContextClearNodes", func(call goja.FunctionCall) goja.Value {
		var ctx *imgui.Context
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.Context); ok {
				ctx = v
			}
		}
		// TODO: 处理类型 ID
		root_id := imgui.ID(0)
		clear_settings_refs := call.Argument(2).ToBoolean()
		imgui.InternalDockContextClearNodes(ctx, root_id, clear_settings_refs)
		return goja.Undefined()
	})
	imguiObj.Set("internalDockContextEndFrame", func(call goja.FunctionCall) goja.Value {
		var ctx *imgui.Context
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.Context); ok {
				ctx = v
			}
		}
		imgui.InternalDockContextEndFrame(ctx)
		return goja.Undefined()
	})
	imguiObj.Set("internalDockContextFindNodeByID", func(call goja.FunctionCall) goja.Value {
		var ctx *imgui.Context
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.Context); ok {
				ctx = v
			}
		}
		// TODO: 处理类型 ID
		id := imgui.ID(0)
		result := imgui.InternalDockContextFindNodeByID(ctx, id)
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalDockContextGenNodeID", func(call goja.FunctionCall) goja.Value {
		var ctx *imgui.Context
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.Context); ok {
				ctx = v
			}
		}
		return vm.ToValue(imgui.InternalDockContextGenNodeID(ctx))
	})
	imguiObj.Set("internalDockContextInitialize", func(call goja.FunctionCall) goja.Value {
		var ctx *imgui.Context
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.Context); ok {
				ctx = v
			}
		}
		imgui.InternalDockContextInitialize(ctx)
		return goja.Undefined()
	})
	imguiObj.Set("internalDockContextNewFrameUpdateDocking", func(call goja.FunctionCall) goja.Value {
		var ctx *imgui.Context
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.Context); ok {
				ctx = v
			}
		}
		imgui.InternalDockContextNewFrameUpdateDocking(ctx)
		return goja.Undefined()
	})
	imguiObj.Set("internalDockContextNewFrameUpdateUndocking", func(call goja.FunctionCall) goja.Value {
		var ctx *imgui.Context
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.Context); ok {
				ctx = v
			}
		}
		imgui.InternalDockContextNewFrameUpdateUndocking(ctx)
		return goja.Undefined()
	})
	imguiObj.Set("internalDockContextProcessUndockNode", func(call goja.FunctionCall) goja.Value {
		var ctx *imgui.Context
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.Context); ok {
				ctx = v
			}
		}
		// TODO: 处理指针类型 *DockNode
		var node *imgui.DockNode
		imgui.InternalDockContextProcessUndockNode(ctx, node)
		return goja.Undefined()
	})
	imguiObj.Set("internalDockContextProcessUndockWindowV", func(call goja.FunctionCall) goja.Value {
		var ctx *imgui.Context
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.Context); ok {
				ctx = v
			}
		}
		// TODO: 处理指针类型 *Window
		var window *imgui.Window
		clear_persistent_docking_ref := call.Argument(2).ToBoolean()
		imgui.InternalDockContextProcessUndockWindowV(ctx, window, clear_persistent_docking_ref)
		return goja.Undefined()
	})
	imguiObj.Set("internalDockContextQueueDock", func(call goja.FunctionCall) goja.Value {
		var ctx *imgui.Context
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.Context); ok {
				ctx = v
			}
		}
		// TODO: 处理指针类型 *Window
		var target *imgui.Window
		// TODO: 处理指针类型 *DockNode
		var target_node *imgui.DockNode
		// TODO: 处理指针类型 *Window
		var payload *imgui.Window
		// TODO: 处理类型 Dir
		split_dir := imgui.Dir(0)
		split_ratio := float32(call.Argument(5).ToFloat())
		split_outer := call.Argument(6).ToBoolean()
		imgui.InternalDockContextQueueDock(ctx, target, target_node, payload, split_dir, split_ratio, split_outer)
		return goja.Undefined()
	})
	imguiObj.Set("internalDockContextQueueUndockNode", func(call goja.FunctionCall) goja.Value {
		var ctx *imgui.Context
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.Context); ok {
				ctx = v
			}
		}
		// TODO: 处理指针类型 *DockNode
		var node *imgui.DockNode
		imgui.InternalDockContextQueueUndockNode(ctx, node)
		return goja.Undefined()
	})
	imguiObj.Set("internalDockContextQueueUndockWindow", func(call goja.FunctionCall) goja.Value {
		var ctx *imgui.Context
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.Context); ok {
				ctx = v
			}
		}
		// TODO: 处理指针类型 *Window
		var window *imgui.Window
		imgui.InternalDockContextQueueUndockWindow(ctx, window)
		return goja.Undefined()
	})
	imguiObj.Set("internalDockContextRebuildNodes", func(call goja.FunctionCall) goja.Value {
		var ctx *imgui.Context
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.Context); ok {
				ctx = v
			}
		}
		imgui.InternalDockContextRebuildNodes(ctx)
		return goja.Undefined()
	})
	imguiObj.Set("internalDockContextShutdown", func(call goja.FunctionCall) goja.Value {
		var ctx *imgui.Context
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.Context); ok {
				ctx = v
			}
		}
		imgui.InternalDockContextShutdown(ctx)
		return goja.Undefined()
	})
	imguiObj.Set("internalDockNodeBeginAmendTabBar", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *DockNode
		var node *imgui.DockNode
		return vm.ToValue(imgui.InternalDockNodeBeginAmendTabBar(node))
	})
	imguiObj.Set("internalDockNodeEndAmendTabBar", func(call goja.FunctionCall) goja.Value {
		imgui.InternalDockNodeEndAmendTabBar()
		return goja.Undefined()
	})
	imguiObj.Set("internalDockNodeGetDepth", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *DockNode
		var node *imgui.DockNode
		return vm.ToValue(imgui.InternalDockNodeGetDepth(node))
	})
	imguiObj.Set("internalDockNodeGetRootNode", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *DockNode
		var node *imgui.DockNode
		result := imgui.InternalDockNodeGetRootNode(node)
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalDockNodeGetWindowMenuButtonId", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *DockNode
		var node *imgui.DockNode
		return vm.ToValue(imgui.InternalDockNodeGetWindowMenuButtonId(node))
	})
	imguiObj.Set("internalDockNodeIsInHierarchyOf", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *DockNode
		var node *imgui.DockNode
		// TODO: 处理指针类型 *DockNode
		var parent *imgui.DockNode
		return vm.ToValue(imgui.InternalDockNodeIsInHierarchyOf(node, parent))
	})
	imguiObj.Set("internalDockNodeWindowMenuHandlerDefault", func(call goja.FunctionCall) goja.Value {
		var ctx *imgui.Context
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.Context); ok {
				ctx = v
			}
		}
		// TODO: 处理指针类型 *DockNode
		var node *imgui.DockNode
		// TODO: 处理指针类型 *TabBar
		var tab_bar *imgui.TabBar
		imgui.InternalDockNodeWindowMenuHandlerDefault(ctx, node, tab_bar)
		return goja.Undefined()
	})
	imguiObj.Set("dockSpaceV", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 ID
		dockspace_id := imgui.ID(0)
		size := parseVec2(vm, call.Argument(1))
		// TODO: 处理类型 DockNodeFlags
		flags := imgui.DockNodeFlags(0)
		var window_class *imgui.WindowClass
		if len(call.Arguments) >= 3 {
			if v, ok := call.Argument(2).Export().(*imgui.WindowClass); ok {
				window_class = v
			}
		}
		return vm.ToValue(imgui.DockSpaceV(dockspace_id, size, flags, window_class))
	})
	imguiObj.Set("dockSpaceOverViewportV", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 ID
		dockspace_id := imgui.ID(0)
		var viewport *imgui.Viewport
		if len(call.Arguments) >= 2 {
			if v, ok := call.Argument(1).Export().(*imgui.Viewport); ok {
				viewport = v
			}
		}
		// TODO: 处理类型 DockNodeFlags
		flags := imgui.DockNodeFlags(0)
		var window_class *imgui.WindowClass
		if len(call.Arguments) >= 4 {
			if v, ok := call.Argument(3).Export().(*imgui.WindowClass); ok {
				window_class = v
			}
		}
		return vm.ToValue(imgui.DockSpaceOverViewportV(dockspace_id, viewport, flags, window_class))
	})
	imguiObj.Set("internalDragBehavior", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 ID
		id := imgui.ID(0)
		// TODO: 处理类型 DataType
		data_type := imgui.DataType(0)
		p_v := uintptr(call.Argument(2).ToInteger())
		v_speed := float32(call.Argument(3).ToFloat())
		p_min := uintptr(call.Argument(4).ToInteger())
		p_max := uintptr(call.Argument(5).ToInteger())
		format := call.Argument(6).String()
		// TODO: 处理类型 SliderFlags
		flags := imgui.SliderFlags(0)
		return vm.ToValue(imgui.InternalDragBehavior(id, data_type, p_v, v_speed, p_min, p_max, format, flags))
	})
	imguiObj.Set("dragFloatV", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		v := float32(0)
		if len(call.Arguments) >= 2 {
			v = float32(call.Argument(1).ToFloat())
		}
		v_speed := float32(call.Argument(2).ToFloat())
		v_min := float32(call.Argument(3).ToFloat())
		v_max := float32(call.Argument(4).ToFloat())
		format := call.Argument(5).String()
		// TODO: 处理类型 SliderFlags
		flags := imgui.SliderFlags(0)
		return vm.ToValue(imgui.DragFloatV(label, &v, v_speed, v_min, v_max, format, flags))
	})
	imguiObj.Set("dragFloat2V", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		v := [2]float32{0, 0}
		v_speed := float32(call.Argument(2).ToFloat())
		v_min := float32(call.Argument(3).ToFloat())
		v_max := float32(call.Argument(4).ToFloat())
		format := call.Argument(5).String()
		// TODO: 处理类型 SliderFlags
		flags := imgui.SliderFlags(0)
		return vm.ToValue(imgui.DragFloat2V(label, &v, v_speed, v_min, v_max, format, flags))
	})
	imguiObj.Set("dragFloat3V", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		v := [3]float32{0, 0, 0}
		v_speed := float32(call.Argument(2).ToFloat())
		v_min := float32(call.Argument(3).ToFloat())
		v_max := float32(call.Argument(4).ToFloat())
		format := call.Argument(5).String()
		// TODO: 处理类型 SliderFlags
		flags := imgui.SliderFlags(0)
		return vm.ToValue(imgui.DragFloat3V(label, &v, v_speed, v_min, v_max, format, flags))
	})
	imguiObj.Set("dragFloat4V", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		v := [4]float32{0, 0, 0, 0}
		v_speed := float32(call.Argument(2).ToFloat())
		v_min := float32(call.Argument(3).ToFloat())
		v_max := float32(call.Argument(4).ToFloat())
		format := call.Argument(5).String()
		// TODO: 处理类型 SliderFlags
		flags := imgui.SliderFlags(0)
		return vm.ToValue(imgui.DragFloat4V(label, &v, v_speed, v_min, v_max, format, flags))
	})
	imguiObj.Set("dragFloatRange2V", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		v_current_min := float32(0)
		if len(call.Arguments) >= 2 {
			v_current_min = float32(call.Argument(1).ToFloat())
		}
		v_current_max := float32(0)
		if len(call.Arguments) >= 3 {
			v_current_max = float32(call.Argument(2).ToFloat())
		}
		v_speed := float32(call.Argument(3).ToFloat())
		v_min := float32(call.Argument(4).ToFloat())
		v_max := float32(call.Argument(5).ToFloat())
		format := call.Argument(6).String()
		format_max := call.Argument(7).String()
		// TODO: 处理类型 SliderFlags
		flags := imgui.SliderFlags(0)
		return vm.ToValue(imgui.DragFloatRange2V(label, &v_current_min, &v_current_max, v_speed, v_min, v_max, format, format_max, flags))
	})
	imguiObj.Set("dragIntV", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		v := int32(0)
		if len(call.Arguments) >= 2 {
			v = int32(call.Argument(1).ToInteger())
		}
		v_speed := float32(call.Argument(2).ToFloat())
		v_min := int32(call.Argument(3).ToInteger())
		v_max := int32(call.Argument(4).ToInteger())
		format := call.Argument(5).String()
		// TODO: 处理类型 SliderFlags
		flags := imgui.SliderFlags(0)
		return vm.ToValue(imgui.DragIntV(label, &v, v_speed, v_min, v_max, format, flags))
	})
	imguiObj.Set("dragInt2V", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		v := [2]int32{0, 0}
		v_speed := float32(call.Argument(2).ToFloat())
		v_min := int32(call.Argument(3).ToInteger())
		v_max := int32(call.Argument(4).ToInteger())
		format := call.Argument(5).String()
		// TODO: 处理类型 SliderFlags
		flags := imgui.SliderFlags(0)
		return vm.ToValue(imgui.DragInt2V(label, &v, v_speed, v_min, v_max, format, flags))
	})
	imguiObj.Set("dragInt3V", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		v := [3]int32{0, 0, 0}
		v_speed := float32(call.Argument(2).ToFloat())
		v_min := int32(call.Argument(3).ToInteger())
		v_max := int32(call.Argument(4).ToInteger())
		format := call.Argument(5).String()
		// TODO: 处理类型 SliderFlags
		flags := imgui.SliderFlags(0)
		return vm.ToValue(imgui.DragInt3V(label, &v, v_speed, v_min, v_max, format, flags))
	})
	imguiObj.Set("dragInt4V", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		v := [4]int32{0, 0, 0, 0}
		v_speed := float32(call.Argument(2).ToFloat())
		v_min := int32(call.Argument(3).ToInteger())
		v_max := int32(call.Argument(4).ToInteger())
		format := call.Argument(5).String()
		// TODO: 处理类型 SliderFlags
		flags := imgui.SliderFlags(0)
		return vm.ToValue(imgui.DragInt4V(label, &v, v_speed, v_min, v_max, format, flags))
	})
	imguiObj.Set("dragIntRange2V", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		v_current_min := int32(0)
		if len(call.Arguments) >= 2 {
			v_current_min = int32(call.Argument(1).ToInteger())
		}
		v_current_max := int32(0)
		if len(call.Arguments) >= 3 {
			v_current_max = int32(call.Argument(2).ToInteger())
		}
		v_speed := float32(call.Argument(3).ToFloat())
		v_min := int32(call.Argument(4).ToInteger())
		v_max := int32(call.Argument(5).ToInteger())
		format := call.Argument(6).String()
		format_max := call.Argument(7).String()
		// TODO: 处理类型 SliderFlags
		flags := imgui.SliderFlags(0)
		return vm.ToValue(imgui.DragIntRange2V(label, &v_current_min, &v_current_max, v_speed, v_min, v_max, format, format_max, flags))
	})
	imguiObj.Set("dragScalarV", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		// TODO: 处理类型 DataType
		data_type := imgui.DataType(0)
		p_data := uintptr(call.Argument(2).ToInteger())
		v_speed := float32(call.Argument(3).ToFloat())
		p_min := uintptr(call.Argument(4).ToInteger())
		p_max := uintptr(call.Argument(5).ToInteger())
		format := call.Argument(6).String()
		// TODO: 处理类型 SliderFlags
		flags := imgui.SliderFlags(0)
		return vm.ToValue(imgui.DragScalarV(label, data_type, p_data, v_speed, p_min, p_max, format, flags))
	})
	imguiObj.Set("dragScalarNV", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		// TODO: 处理类型 DataType
		data_type := imgui.DataType(0)
		p_data := uintptr(call.Argument(2).ToInteger())
		components := int32(call.Argument(3).ToInteger())
		v_speed := float32(call.Argument(4).ToFloat())
		p_min := uintptr(call.Argument(5).ToInteger())
		p_max := uintptr(call.Argument(6).ToInteger())
		format := call.Argument(7).String()
		// TODO: 处理类型 SliderFlags
		flags := imgui.SliderFlags(0)
		return vm.ToValue(imgui.DragScalarNV(label, data_type, p_data, components, v_speed, p_min, p_max, format, flags))
	})
	imguiObj.Set("dummy", func(call goja.FunctionCall) goja.Value {
		size := parseVec2(vm, call.Argument(0))
		imgui.Dummy(size)
		return goja.Undefined()
	})
	imguiObj.Set("end", func(call goja.FunctionCall) goja.Value {
		imgui.End()
		return goja.Undefined()
	})
	imguiObj.Set("internalEndBoxSelect", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 Rect
		scope_rect := imgui.Rect{}
		// TODO: 处理类型 MultiSelectFlags
		ms_flags := imgui.MultiSelectFlags(0)
		imgui.InternalEndBoxSelect(scope_rect, ms_flags)
		return goja.Undefined()
	})
	imguiObj.Set("endChild", func(call goja.FunctionCall) goja.Value {
		imgui.EndChild()
		return goja.Undefined()
	})
	imguiObj.Set("internalEndColumns", func(call goja.FunctionCall) goja.Value {
		imgui.InternalEndColumns()
		return goja.Undefined()
	})
	imguiObj.Set("endCombo", func(call goja.FunctionCall) goja.Value {
		imgui.EndCombo()
		return goja.Undefined()
	})
	imguiObj.Set("internalEndComboPreview", func(call goja.FunctionCall) goja.Value {
		imgui.InternalEndComboPreview()
		return goja.Undefined()
	})
	imguiObj.Set("endDisabled", func(call goja.FunctionCall) goja.Value {
		imgui.EndDisabled()
		return goja.Undefined()
	})
	imguiObj.Set("internalEndDisabledOverrideReenable", func(call goja.FunctionCall) goja.Value {
		imgui.InternalEndDisabledOverrideReenable()
		return goja.Undefined()
	})
	imguiObj.Set("endDragDropSource", func(call goja.FunctionCall) goja.Value {
		imgui.EndDragDropSource()
		return goja.Undefined()
	})
	imguiObj.Set("endDragDropTarget", func(call goja.FunctionCall) goja.Value {
		imgui.EndDragDropTarget()
		return goja.Undefined()
	})
	imguiObj.Set("internalEndErrorTooltip", func(call goja.FunctionCall) goja.Value {
		imgui.InternalEndErrorTooltip()
		return goja.Undefined()
	})
	imguiObj.Set("endFrame", func(call goja.FunctionCall) goja.Value {
		imgui.EndFrame()
		return goja.Undefined()
	})
	imguiObj.Set("endGroup", func(call goja.FunctionCall) goja.Value {
		imgui.EndGroup()
		return goja.Undefined()
	})
	imguiObj.Set("endListBox", func(call goja.FunctionCall) goja.Value {
		imgui.EndListBox()
		return goja.Undefined()
	})
	imguiObj.Set("endMainMenuBar", func(call goja.FunctionCall) goja.Value {
		imgui.EndMainMenuBar()
		return goja.Undefined()
	})
	imguiObj.Set("endMenu", func(call goja.FunctionCall) goja.Value {
		imgui.EndMenu()
		return goja.Undefined()
	})
	imguiObj.Set("endMenuBar", func(call goja.FunctionCall) goja.Value {
		imgui.EndMenuBar()
		return goja.Undefined()
	})
	imguiObj.Set("endMultiSelect", func(call goja.FunctionCall) goja.Value {
		result := imgui.EndMultiSelect()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("endPopup", func(call goja.FunctionCall) goja.Value {
		imgui.EndPopup()
		return goja.Undefined()
	})
	imguiObj.Set("endTabBar", func(call goja.FunctionCall) goja.Value {
		imgui.EndTabBar()
		return goja.Undefined()
	})
	imguiObj.Set("endTabItem", func(call goja.FunctionCall) goja.Value {
		imgui.EndTabItem()
		return goja.Undefined()
	})
	imguiObj.Set("endTable", func(call goja.FunctionCall) goja.Value {
		imgui.EndTable()
		return goja.Undefined()
	})
	imguiObj.Set("endTooltip", func(call goja.FunctionCall) goja.Value {
		imgui.EndTooltip()
		return goja.Undefined()
	})
	imguiObj.Set("internalErrorCheckEndFrameFinalizeErrorTooltip", func(call goja.FunctionCall) goja.Value {
		imgui.InternalErrorCheckEndFrameFinalizeErrorTooltip()
		return goja.Undefined()
	})
	imguiObj.Set("internalErrorCheckUsingSetCursorPosToExtendParentBoundaries", func(call goja.FunctionCall) goja.Value {
		imgui.InternalErrorCheckUsingSetCursorPosToExtendParentBoundaries()
		return goja.Undefined()
	})
	imguiObj.Set("internalErrorLog", func(call goja.FunctionCall) goja.Value {
		msg := call.Argument(0).String()
		return vm.ToValue(imgui.InternalErrorLog(msg))
	})
	imguiObj.Set("internalErrorRecoveryStoreState", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *ErrorRecoveryState
		var state_out *imgui.ErrorRecoveryState
		imgui.InternalErrorRecoveryStoreState(state_out)
		return goja.Undefined()
	})
	imguiObj.Set("internalErrorRecoveryTryToRecoverState", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *ErrorRecoveryState
		var state_in *imgui.ErrorRecoveryState
		imgui.InternalErrorRecoveryTryToRecoverState(state_in)
		return goja.Undefined()
	})
	imguiObj.Set("internalErrorRecoveryTryToRecoverWindowState", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *ErrorRecoveryState
		var state_in *imgui.ErrorRecoveryState
		imgui.InternalErrorRecoveryTryToRecoverWindowState(state_in)
		return goja.Undefined()
	})
	imguiObj.Set("internalFindBestWindowPosForPopup", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Window
		var window *imgui.Window
		result := imgui.InternalFindBestWindowPosForPopup(window)
		return vm.ToValue(result)
	})
	imguiObj.Set("internalFindBestWindowPosForPopupEx", func(call goja.FunctionCall) goja.Value {
		ref_pos := parseVec2(vm, call.Argument(0))
		size := parseVec2(vm, call.Argument(0))
		// TODO: 处理指针类型 *Dir
		var last_dir *imgui.Dir
		// TODO: 处理类型 Rect
		r_outer := imgui.Rect{}
		// TODO: 处理类型 Rect
		r_avoid := imgui.Rect{}
		// TODO: 处理类型 PopupPositionPolicy
		policy := imgui.PopupPositionPolicy(0)
		result := imgui.InternalFindBestWindowPosForPopupEx(ref_pos, size, last_dir, r_outer, r_avoid, policy)
		return vm.ToValue(result)
	})
	imguiObj.Set("internalFindBlockingModal", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Window
		var window *imgui.Window
		result := imgui.InternalFindBlockingModal(window)
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalFindBottomMostVisibleWindowWithinBeginStack", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Window
		var window *imgui.Window
		result := imgui.InternalFindBottomMostVisibleWindowWithinBeginStack(window)
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalFindHoveredViewportFromPlatformWindowStack", func(call goja.FunctionCall) goja.Value {
		mouse_platform_pos := parseVec2(vm, call.Argument(0))
		result := imgui.InternalFindHoveredViewportFromPlatformWindowStack(mouse_platform_pos)
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalFindOrCreateColumns", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Window
		var window *imgui.Window
		// TODO: 处理类型 ID
		id := imgui.ID(0)
		result := imgui.InternalFindOrCreateColumns(window, id)
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalFindRenderedTextEndV", func(call goja.FunctionCall) goja.Value {
		text := call.Argument(0).String()
		return vm.ToValue(imgui.InternalFindRenderedTextEndV(text))
	})
	imguiObj.Set("internalFindSettingsHandler", func(call goja.FunctionCall) goja.Value {
		type_name := call.Argument(0).String()
		result := imgui.InternalFindSettingsHandler(type_name)
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("findViewportByID", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 ID
		id := imgui.ID(0)
		result := imgui.FindViewportByID(id)
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("findViewportByPlatformHandle", func(call goja.FunctionCall) goja.Value {
		platform_handle := uintptr(call.Argument(0).ToInteger())
		result := imgui.FindViewportByPlatformHandle(platform_handle)
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalFindWindowByID", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 ID
		id := imgui.ID(0)
		result := imgui.InternalFindWindowByID(id)
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalFindWindowByName", func(call goja.FunctionCall) goja.Value {
		name := call.Argument(0).String()
		result := imgui.InternalFindWindowByName(name)
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalFindWindowDisplayIndex", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Window
		var window *imgui.Window
		return vm.ToValue(imgui.InternalFindWindowDisplayIndex(window))
	})
	imguiObj.Set("internalFindWindowSettingsByID", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 ID
		id := imgui.ID(0)
		result := imgui.InternalFindWindowSettingsByID(id)
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalFindWindowSettingsByWindow", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Window
		var window *imgui.Window
		result := imgui.InternalFindWindowSettingsByWindow(window)
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalFixupKeyChord", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 KeyChord
		key_chord := imgui.KeyChord(0)
		return vm.ToValue(imgui.InternalFixupKeyChord(key_chord))
	})
	imguiObj.Set("internalFocusItem", func(call goja.FunctionCall) goja.Value {
		imgui.InternalFocusItem()
		return goja.Undefined()
	})
	imguiObj.Set("internalFocusTopMostWindowUnderOne", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Window
		var under_this_window *imgui.Window
		// TODO: 处理指针类型 *Window
		var ignore_window *imgui.Window
		var filter_viewport *imgui.Viewport
		if len(call.Arguments) >= 3 {
			if v, ok := call.Argument(2).Export().(*imgui.Viewport); ok {
				filter_viewport = v
			}
		}
		// TODO: 处理类型 FocusRequestFlags
		flags := imgui.FocusRequestFlags(0)
		imgui.InternalFocusTopMostWindowUnderOne(under_this_window, ignore_window, filter_viewport, flags)
		return goja.Undefined()
	})
	imguiObj.Set("internalFocusWindowV", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Window
		var window *imgui.Window
		// TODO: 处理类型 FocusRequestFlags
		flags := imgui.FocusRequestFlags(0)
		imgui.InternalFocusWindowV(window, flags)
		return goja.Undefined()
	})
	imguiObj.Set("internalGcAwakeTransientWindowBuffers", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Window
		var window *imgui.Window
		imgui.InternalGcAwakeTransientWindowBuffers(window)
		return goja.Undefined()
	})
	imguiObj.Set("internalGcCompactTransientMiscBuffers", func(call goja.FunctionCall) goja.Value {
		imgui.InternalGcCompactTransientMiscBuffers()
		return goja.Undefined()
	})
	imguiObj.Set("internalGcCompactTransientWindowBuffers", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Window
		var window *imgui.Window
		imgui.InternalGcCompactTransientWindowBuffers(window)
		return goja.Undefined()
	})
	imguiObj.Set("internalActiveID", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.InternalActiveID())
	})
	imguiObj.Set("backgroundDrawListV", func(call goja.FunctionCall) goja.Value {
		var viewport *imgui.Viewport
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.Viewport); ok {
				viewport = v
			}
		}
		result := imgui.BackgroundDrawListV(viewport)
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalBoxSelectState", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 ID
		id := imgui.ID(0)
		result := imgui.InternalBoxSelectState(id)
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("clipboardText", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.ClipboardText())
	})
	imguiObj.Set("colorU32ColV", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 Col
		idx := imgui.Col(0)
		alpha_mul := float32(call.Argument(1).ToFloat())
		return vm.ToValue(imgui.ColorU32ColV(idx, alpha_mul))
	})
	imguiObj.Set("colorU32U32V", func(call goja.FunctionCall) goja.Value {
		col := uint32(call.Argument(0).ToInteger())
		alpha_mul := float32(call.Argument(1).ToFloat())
		return vm.ToValue(imgui.ColorU32U32V(col, alpha_mul))
	})
	imguiObj.Set("colorU32Vec4", func(call goja.FunctionCall) goja.Value {
		col := parseVec4(vm, call.Argument(0))
		return vm.ToValue(imgui.ColorU32Vec4(col))
	})
	imguiObj.Set("columnIndex", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.ColumnIndex())
	})
	imguiObj.Set("internalColumnNormFromOffset", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *OldColumns
		var columns *imgui.OldColumns
		offset := float32(call.Argument(1).ToFloat())
		return vm.ToValue(imgui.InternalColumnNormFromOffset(columns, offset))
	})
	imguiObj.Set("columnOffsetV", func(call goja.FunctionCall) goja.Value {
		column_index := int32(call.Argument(0).ToInteger())
		return vm.ToValue(imgui.ColumnOffsetV(column_index))
	})
	imguiObj.Set("internalColumnOffsetFromNorm", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *OldColumns
		var columns *imgui.OldColumns
		offset_norm := float32(call.Argument(1).ToFloat())
		return vm.ToValue(imgui.InternalColumnOffsetFromNorm(columns, offset_norm))
	})
	imguiObj.Set("columnWidthV", func(call goja.FunctionCall) goja.Value {
		column_index := int32(call.Argument(0).ToInteger())
		return vm.ToValue(imgui.ColumnWidthV(column_index))
	})
	imguiObj.Set("columnsCount", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.ColumnsCount())
	})
	imguiObj.Set("internalColumnsID", func(call goja.FunctionCall) goja.Value {
		str_id := call.Argument(0).String()
		count := int32(call.Argument(1).ToInteger())
		return vm.ToValue(imgui.InternalColumnsID(str_id, count))
	})
	imguiObj.Set("contentRegionAvail", func(call goja.FunctionCall) goja.Value {
		result := imgui.ContentRegionAvail()
		return vm.ToValue(result)
	})
	imguiObj.Set("currentContext", func(call goja.FunctionCall) goja.Value {
		result := imgui.CurrentContext()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalCurrentFocusScope", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.InternalCurrentFocusScope())
	})
	imguiObj.Set("internalCurrentTabBar", func(call goja.FunctionCall) goja.Value {
		result := imgui.InternalCurrentTabBar()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalCurrentTable", func(call goja.FunctionCall) goja.Value {
		result := imgui.InternalCurrentTable()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalCurrentWindow", func(call goja.FunctionCall) goja.Value {
		result := imgui.InternalCurrentWindow()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalCurrentWindowRead", func(call goja.FunctionCall) goja.Value {
		result := imgui.InternalCurrentWindowRead()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("cursorPos", func(call goja.FunctionCall) goja.Value {
		result := imgui.CursorPos()
		return vm.ToValue(result)
	})
	imguiObj.Set("cursorPosX", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.CursorPosX())
	})
	imguiObj.Set("cursorPosY", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.CursorPosY())
	})
	imguiObj.Set("cursorScreenPos", func(call goja.FunctionCall) goja.Value {
		result := imgui.CursorScreenPos()
		return vm.ToValue(result)
	})
	imguiObj.Set("cursorStartPos", func(call goja.FunctionCall) goja.Value {
		result := imgui.CursorStartPos()
		return vm.ToValue(result)
	})
	imguiObj.Set("internalDefaultFont", func(call goja.FunctionCall) goja.Value {
		result := imgui.InternalDefaultFont()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("dragDropPayload", func(call goja.FunctionCall) goja.Value {
		result := imgui.DragDropPayload()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("currentDrawData", func(call goja.FunctionCall) goja.Value {
		result := imgui.CurrentDrawData()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("currentDrawListSharedData", func(call goja.FunctionCall) goja.Value {
		result := imgui.CurrentDrawListSharedData()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalFocusID", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.InternalFocusID())
	})
	imguiObj.Set("currentFont", func(call goja.FunctionCall) goja.Value {
		result := imgui.CurrentFont()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("getFontBaked", func(call goja.FunctionCall) goja.Value {
		result := imgui.GetFontBaked()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalFontRasterizerDensity", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.InternalFontRasterizerDensity())
	})
	imguiObj.Set("fontSize", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.FontSize())
	})
	imguiObj.Set("fontTexUvWhitePixel", func(call goja.FunctionCall) goja.Value {
		result := imgui.FontTexUvWhitePixel()
		return vm.ToValue(result)
	})
	imguiObj.Set("foregroundDrawListViewportPtrV", func(call goja.FunctionCall) goja.Value {
		var viewport *imgui.Viewport
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.Viewport); ok {
				viewport = v
			}
		}
		result := imgui.ForegroundDrawListViewportPtrV(viewport)
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalForegroundDrawListWindowPtr", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Window
		var window *imgui.Window
		result := imgui.InternalForegroundDrawListWindowPtr(window)
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("frameCount", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.FrameCount())
	})
	imguiObj.Set("frameHeight", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.FrameHeight())
	})
	imguiObj.Set("frameHeightWithSpacing", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.FrameHeightWithSpacing())
	})
	imguiObj.Set("internalHoveredID", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.InternalHoveredID())
	})
	imguiObj.Set("internalIDWithSeedInt", func(call goja.FunctionCall) goja.Value {
		n := int32(call.Argument(0).ToInteger())
		// TODO: 处理类型 ID
		seed := imgui.ID(0)
		return vm.ToValue(imgui.InternalIDWithSeedInt(n, seed))
	})
	imguiObj.Set("internalIDWithSeedStr", func(call goja.FunctionCall) goja.Value {
		str_id_begin := call.Argument(0).String()
		str_id_end := call.Argument(1).String()
		// TODO: 处理类型 ID
		seed := imgui.ID(0)
		return vm.ToValue(imgui.InternalIDWithSeedStr(str_id_begin, str_id_end, seed))
	})
	imguiObj.Set("iDInt", func(call goja.FunctionCall) goja.Value {
		int_id := int32(call.Argument(0).ToInteger())
		return vm.ToValue(imgui.IDInt(int_id))
	})
	imguiObj.Set("iDPtr", func(call goja.FunctionCall) goja.Value {
		ptr_id := uintptr(call.Argument(0).ToInteger())
		return vm.ToValue(imgui.IDPtr(ptr_id))
	})
	imguiObj.Set("iDStr", func(call goja.FunctionCall) goja.Value {
		str_id := call.Argument(0).String()
		return vm.ToValue(imgui.IDStr(str_id))
	})
	imguiObj.Set("iDStrStr", func(call goja.FunctionCall) goja.Value {
		str_id_begin := call.Argument(0).String()
		str_id_end := call.Argument(1).String()
		return vm.ToValue(imgui.IDStrStr(str_id_begin, str_id_end))
	})
	imguiObj.Set("internalIOContextPtr", func(call goja.FunctionCall) goja.Value {
		var ctx *imgui.Context
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.Context); ok {
				ctx = v
			}
		}
		result := imgui.InternalIOContextPtr(ctx)
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("currentIO", func(call goja.FunctionCall) goja.Value {
		result := imgui.CurrentIO()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalInputTextState", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 ID
		id := imgui.ID(0)
		result := imgui.InternalInputTextState(id)
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalItemFlags", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.InternalItemFlags())
	})
	imguiObj.Set("itemID", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.ItemID())
	})
	imguiObj.Set("itemRectMax", func(call goja.FunctionCall) goja.Value {
		result := imgui.ItemRectMax()
		return vm.ToValue(result)
	})
	imguiObj.Set("itemRectMin", func(call goja.FunctionCall) goja.Value {
		result := imgui.ItemRectMin()
		return vm.ToValue(result)
	})
	imguiObj.Set("itemRectSize", func(call goja.FunctionCall) goja.Value {
		result := imgui.ItemRectSize()
		return vm.ToValue(result)
	})
	imguiObj.Set("internalItemStatusFlags", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.InternalItemStatusFlags())
	})
	imguiObj.Set("internalKeyChordName", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 KeyChord
		key_chord := imgui.KeyChord(0)
		return vm.ToValue(imgui.InternalKeyChordName(key_chord))
	})
	imguiObj.Set("internalKeyDataContextPtr", func(call goja.FunctionCall) goja.Value {
		var ctx *imgui.Context
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.Context); ok {
				ctx = v
			}
		}
		// TODO: 处理类型 Key
		key := imgui.Key(0)
		result := imgui.InternalKeyDataContextPtr(ctx, key)
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalKeyDataKey", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 Key
		key := imgui.Key(0)
		result := imgui.InternalKeyDataKey(key)
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalKeyMagnitude2d", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 Key
		key_left := imgui.Key(0)
		// TODO: 处理类型 Key
		key_right := imgui.Key(0)
		// TODO: 处理类型 Key
		key_up := imgui.Key(0)
		// TODO: 处理类型 Key
		key_down := imgui.Key(0)
		result := imgui.InternalKeyMagnitude2d(key_left, key_right, key_up, key_down)
		return vm.ToValue(result)
	})
	imguiObj.Set("keyName", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 Key
		key := imgui.Key(0)
		return vm.ToValue(imgui.KeyName(key))
	})
	imguiObj.Set("internalKeyOwner", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 Key
		key := imgui.Key(0)
		return vm.ToValue(imgui.InternalKeyOwner(key))
	})
	imguiObj.Set("internalKeyOwnerData", func(call goja.FunctionCall) goja.Value {
		var ctx *imgui.Context
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.Context); ok {
				ctx = v
			}
		}
		// TODO: 处理类型 Key
		key := imgui.Key(0)
		result := imgui.InternalKeyOwnerData(ctx, key)
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("keyPressedAmount", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 Key
		key := imgui.Key(0)
		repeat_delay := float32(call.Argument(1).ToFloat())
		rate := float32(call.Argument(2).ToFloat())
		return vm.ToValue(imgui.KeyPressedAmount(key, repeat_delay, rate))
	})
	imguiObj.Set("mainViewport", func(call goja.FunctionCall) goja.Value {
		result := imgui.MainViewport()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("mouseClickedCount", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 MouseButton
		button := imgui.MouseButton(0)
		return vm.ToValue(imgui.MouseClickedCount(button))
	})
	imguiObj.Set("currentMouseCursor", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.CurrentMouseCursor())
	})
	imguiObj.Set("mouseDragDeltaV", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 MouseButton
		button := imgui.MouseButton(0)
		lock_threshold := float32(call.Argument(1).ToFloat())
		result := imgui.MouseDragDeltaV(button, lock_threshold)
		return vm.ToValue(result)
	})
	imguiObj.Set("mousePos", func(call goja.FunctionCall) goja.Value {
		result := imgui.MousePos()
		return vm.ToValue(result)
	})
	imguiObj.Set("mousePosOnOpeningCurrentPopup", func(call goja.FunctionCall) goja.Value {
		result := imgui.MousePosOnOpeningCurrentPopup()
		return vm.ToValue(result)
	})
	imguiObj.Set("internalMultiSelectState", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 ID
		id := imgui.ID(0)
		result := imgui.InternalMultiSelectState(id)
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalNavTweakPressedAmount", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 Axis
		axis := imgui.Axis(0)
		return vm.ToValue(imgui.InternalNavTweakPressedAmount(axis))
	})
	imguiObj.Set("internalPlatformIOContextPtr", func(call goja.FunctionCall) goja.Value {
		var ctx *imgui.Context
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.Context); ok {
				ctx = v
			}
		}
		result := imgui.InternalPlatformIOContextPtr(ctx)
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("currentPlatformIO", func(call goja.FunctionCall) goja.Value {
		result := imgui.CurrentPlatformIO()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalPopupAllowedExtentRect", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Window
		var window *imgui.Window
		return vm.ToValue(imgui.InternalPopupAllowedExtentRect(window))
	})
	imguiObj.Set("internalRoundedFontSize", func(call goja.FunctionCall) goja.Value {
		size := float32(call.Argument(0).ToFloat())
		return vm.ToValue(imgui.InternalRoundedFontSize(size))
	})
	imguiObj.Set("scrollMaxX", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.ScrollMaxX())
	})
	imguiObj.Set("scrollMaxY", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.ScrollMaxY())
	})
	imguiObj.Set("scrollX", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.ScrollX())
	})
	imguiObj.Set("scrollY", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.ScrollY())
	})
	imguiObj.Set("internalShortcutRoutingData", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 KeyChord
		key_chord := imgui.KeyChord(0)
		result := imgui.InternalShortcutRoutingData(key_chord)
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("stateStorage", func(call goja.FunctionCall) goja.Value {
		result := imgui.StateStorage()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("currentStyle", func(call goja.FunctionCall) goja.Value {
		result := imgui.CurrentStyle()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("styleColorName", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 Col
		idx := imgui.Col(0)
		return vm.ToValue(imgui.StyleColorName(idx))
	})
	imguiObj.Set("styleColorVec4", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 Col
		idx := imgui.Col(0)
		result := imgui.StyleColorVec4(idx)
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalStyleVarInfo", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 StyleVar
		idx := imgui.StyleVar(0)
		result := imgui.InternalStyleVarInfo(idx)
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("textLineHeight", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.TextLineHeight())
	})
	imguiObj.Set("textLineHeightWithSpacing", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.TextLineHeightWithSpacing())
	})
	imguiObj.Set("time", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.Time())
	})
	imguiObj.Set("internalTopMostAndVisiblePopupModal", func(call goja.FunctionCall) goja.Value {
		result := imgui.InternalTopMostAndVisiblePopupModal()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalTopMostPopupModal", func(call goja.FunctionCall) goja.Value {
		result := imgui.InternalTopMostPopupModal()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("treeNodeToLabelSpacing", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.TreeNodeToLabelSpacing())
	})
	imguiObj.Set("internalTypematicRepeatRate", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 InputFlags
		flags := imgui.InputFlags(0)
		repeat_delay := float32(0)
		if len(call.Arguments) >= 2 {
			repeat_delay = float32(call.Argument(1).ToFloat())
		}
		repeat_rate := float32(0)
		if len(call.Arguments) >= 3 {
			repeat_rate = float32(call.Argument(2).ToFloat())
		}
		imgui.InternalTypematicRepeatRate(flags, &repeat_delay, &repeat_rate)
		return goja.Undefined()
	})
	imguiObj.Set("internalTypingSelectRequestV", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 TypingSelectFlags
		flags := imgui.TypingSelectFlags(0)
		result := imgui.InternalTypingSelectRequestV(flags)
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("version", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.Version())
	})
	imguiObj.Set("internalViewportPlatformMonitor", func(call goja.FunctionCall) goja.Value {
		var viewport *imgui.Viewport
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.Viewport); ok {
				viewport = v
			}
		}
		result := imgui.InternalViewportPlatformMonitor(viewport)
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalWindowAlwaysWantOwnTabBar", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Window
		var window *imgui.Window
		return vm.ToValue(imgui.InternalWindowAlwaysWantOwnTabBar(window))
	})
	imguiObj.Set("windowDockID", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.WindowDockID())
	})
	imguiObj.Set("internalWindowDockNode", func(call goja.FunctionCall) goja.Value {
		result := imgui.InternalWindowDockNode()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("windowDpiScale", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.WindowDpiScale())
	})
	imguiObj.Set("windowDrawList", func(call goja.FunctionCall) goja.Value {
		result := imgui.WindowDrawList()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("windowHeight", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.WindowHeight())
	})
	imguiObj.Set("windowPos", func(call goja.FunctionCall) goja.Value {
		result := imgui.WindowPos()
		return vm.ToValue(result)
	})
	imguiObj.Set("internalWindowResizeBorderID", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Window
		var window *imgui.Window
		// TODO: 处理类型 Dir
		dir := imgui.Dir(0)
		return vm.ToValue(imgui.InternalWindowResizeBorderID(window, dir))
	})
	imguiObj.Set("internalWindowResizeCornerID", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Window
		var window *imgui.Window
		n := int32(call.Argument(1).ToInteger())
		return vm.ToValue(imgui.InternalWindowResizeCornerID(window, n))
	})
	imguiObj.Set("internalWindowScrollbarID", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Window
		var window *imgui.Window
		// TODO: 处理类型 Axis
		axis := imgui.Axis(0)
		return vm.ToValue(imgui.InternalWindowScrollbarID(window, axis))
	})
	imguiObj.Set("internalWindowScrollbarRect", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Window
		var window *imgui.Window
		// TODO: 处理类型 Axis
		axis := imgui.Axis(0)
		return vm.ToValue(imgui.InternalWindowScrollbarRect(window, axis))
	})
	imguiObj.Set("windowSize", func(call goja.FunctionCall) goja.Value {
		result := imgui.WindowSize()
		return vm.ToValue(result)
	})
	imguiObj.Set("windowViewport", func(call goja.FunctionCall) goja.Value {
		result := imgui.WindowViewport()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("windowWidth", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.WindowWidth())
	})
	imguiObj.Set("internalImAbsFloat", func(call goja.FunctionCall) goja.Value {
		x := float32(call.Argument(0).ToFloat())
		return vm.ToValue(imgui.InternalImAbsFloat(x))
	})
	imguiObj.Set("internalImAbsInt", func(call goja.FunctionCall) goja.Value {
		x := int32(call.Argument(0).ToInteger())
		return vm.ToValue(imgui.InternalImAbsInt(x))
	})
	imguiObj.Set("internalImAbsDouble", func(call goja.FunctionCall) goja.Value {
		x := call.Argument(0).ToFloat()
		return vm.ToValue(imgui.InternalImAbsDouble(x))
	})
	imguiObj.Set("internalImAlphaBlendColors", func(call goja.FunctionCall) goja.Value {
		col_a := uint32(call.Argument(0).ToInteger())
		col_b := uint32(call.Argument(1).ToInteger())
		return vm.ToValue(imgui.InternalImAlphaBlendColors(col_a, col_b))
	})
	imguiObj.Set("internalImBezierCubicCalc", func(call goja.FunctionCall) goja.Value {
		p1 := parseVec2(vm, call.Argument(0))
		p2 := parseVec2(vm, call.Argument(0))
		p3 := parseVec2(vm, call.Argument(0))
		p4 := parseVec2(vm, call.Argument(0))
		t := float32(call.Argument(0).ToFloat())
		result := imgui.InternalImBezierCubicCalc(p1, p2, p3, p4, t)
		return vm.ToValue(result)
	})
	imguiObj.Set("internalImBezierCubicClosestPoint", func(call goja.FunctionCall) goja.Value {
		p1 := parseVec2(vm, call.Argument(0))
		p2 := parseVec2(vm, call.Argument(0))
		p3 := parseVec2(vm, call.Argument(0))
		p4 := parseVec2(vm, call.Argument(0))
		p := parseVec2(vm, call.Argument(0))
		num_segments := int32(call.Argument(0).ToInteger())
		result := imgui.InternalImBezierCubicClosestPoint(p1, p2, p3, p4, p, num_segments)
		return vm.ToValue(result)
	})
	imguiObj.Set("internalImBezierCubicClosestPointCasteljau", func(call goja.FunctionCall) goja.Value {
		p1 := parseVec2(vm, call.Argument(0))
		p2 := parseVec2(vm, call.Argument(0))
		p3 := parseVec2(vm, call.Argument(0))
		p4 := parseVec2(vm, call.Argument(0))
		p := parseVec2(vm, call.Argument(0))
		tess_tol := float32(call.Argument(0).ToFloat())
		result := imgui.InternalImBezierCubicClosestPointCasteljau(p1, p2, p3, p4, p, tess_tol)
		return vm.ToValue(result)
	})
	imguiObj.Set("internalImBezierQuadraticCalc", func(call goja.FunctionCall) goja.Value {
		p1 := parseVec2(vm, call.Argument(0))
		p2 := parseVec2(vm, call.Argument(0))
		p3 := parseVec2(vm, call.Argument(0))
		t := float32(call.Argument(0).ToFloat())
		result := imgui.InternalImBezierQuadraticCalc(p1, p2, p3, t)
		return vm.ToValue(result)
	})
	imguiObj.Set("internalImBitArrayClearAllBits", func(call goja.FunctionCall) goja.Value {
		arr := uint32(0)
		if len(call.Arguments) >= 1 {
			arr = uint32(call.Argument(0).ToInteger())
		}
		bitcount := int32(call.Argument(1).ToInteger())
		imgui.InternalImBitArrayClearAllBits(&arr, bitcount)
		return goja.Undefined()
	})
	imguiObj.Set("internalImBitArrayClearBit", func(call goja.FunctionCall) goja.Value {
		arr := uint32(0)
		if len(call.Arguments) >= 1 {
			arr = uint32(call.Argument(0).ToInteger())
		}
		n := int32(call.Argument(1).ToInteger())
		imgui.InternalImBitArrayClearBit(&arr, n)
		return goja.Undefined()
	})
	imguiObj.Set("internalImBitArrayGetStorageSizeInBytes", func(call goja.FunctionCall) goja.Value {
		bitcount := int32(call.Argument(0).ToInteger())
		return vm.ToValue(imgui.InternalImBitArrayGetStorageSizeInBytes(bitcount))
	})
	imguiObj.Set("internalImBitArraySetBit", func(call goja.FunctionCall) goja.Value {
		arr := uint32(0)
		if len(call.Arguments) >= 1 {
			arr = uint32(call.Argument(0).ToInteger())
		}
		n := int32(call.Argument(1).ToInteger())
		imgui.InternalImBitArraySetBit(&arr, n)
		return goja.Undefined()
	})
	imguiObj.Set("internalImBitArraySetBitRange", func(call goja.FunctionCall) goja.Value {
		arr := uint32(0)
		if len(call.Arguments) >= 1 {
			arr = uint32(call.Argument(0).ToInteger())
		}
		n := int32(call.Argument(1).ToInteger())
		n2 := int32(call.Argument(2).ToInteger())
		imgui.InternalImBitArraySetBitRange(&arr, n, n2)
		return goja.Undefined()
	})
	imguiObj.Set("internalImBitArrayTestBit", func(call goja.FunctionCall) goja.Value {
		arr := uint32(0)
		if len(call.Arguments) >= 1 {
			arr = uint32(call.Argument(0).ToInteger())
		}
		n := int32(call.Argument(1).ToInteger())
		return vm.ToValue(imgui.InternalImBitArrayTestBit(&arr, n))
	})
	imguiObj.Set("internalImCharIsBlankA", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 rune
		c := rune(0)
		return vm.ToValue(imgui.InternalImCharIsBlankA(c))
	})
	imguiObj.Set("internalImCharIsBlankW", func(call goja.FunctionCall) goja.Value {
		c := uint32(call.Argument(0).ToInteger())
		return vm.ToValue(imgui.InternalImCharIsBlankW(c))
	})
	imguiObj.Set("internalImCharIsXdigitA", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 rune
		c := rune(0)
		return vm.ToValue(imgui.InternalImCharIsXdigitA(c))
	})
	imguiObj.Set("internalImClamp", func(call goja.FunctionCall) goja.Value {
		v := parseVec2(vm, call.Argument(0))
		mn := parseVec2(vm, call.Argument(0))
		mx := parseVec2(vm, call.Argument(0))
		result := imgui.InternalImClamp(v, mn, mx)
		return vm.ToValue(result)
	})
	imguiObj.Set("internalImCountSetBits", func(call goja.FunctionCall) goja.Value {
		v := uint32(call.Argument(0).ToInteger())
		return vm.ToValue(imgui.InternalImCountSetBits(v))
	})
	imguiObj.Set("internalImDot", func(call goja.FunctionCall) goja.Value {
		a := parseVec2(vm, call.Argument(0))
		b := parseVec2(vm, call.Argument(0))
		return vm.ToValue(imgui.InternalImDot(a, b))
	})
	imguiObj.Set("internalImExponentialMovingAverage", func(call goja.FunctionCall) goja.Value {
		avg := float32(call.Argument(0).ToFloat())
		sample := float32(call.Argument(1).ToFloat())
		n := int32(call.Argument(2).ToInteger())
		return vm.ToValue(imgui.InternalImExponentialMovingAverage(avg, sample, n))
	})
	imguiObj.Set("internalImFileLoadToMemoryV", func(call goja.FunctionCall) goja.Value {
		filename := call.Argument(0).String()
		mode := call.Argument(1).String()
		// TODO: 处理指针类型 *uint64
		var out_file_size *uint64
		padding_bytes := int32(call.Argument(3).ToInteger())
		return vm.ToValue(imgui.InternalImFileLoadToMemoryV(filename, mode, out_file_size, padding_bytes))
	})
	imguiObj.Set("internalImFloorFloat", func(call goja.FunctionCall) goja.Value {
		f := float32(call.Argument(0).ToFloat())
		return vm.ToValue(imgui.InternalImFloorFloat(f))
	})
	imguiObj.Set("internalImFloorVec2", func(call goja.FunctionCall) goja.Value {
		v := parseVec2(vm, call.Argument(0))
		result := imgui.InternalImFloorVec2(v)
		return vm.ToValue(result)
	})
	imguiObj.Set("internalImFontAtlasAddDrawListSharedData", func(call goja.FunctionCall) goja.Value {
		var atlas *imgui.FontAtlas
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.FontAtlas); ok {
				atlas = v
			}
		}
		// TODO: 处理指针类型 *DrawListSharedData
		var data *imgui.DrawListSharedData
		imgui.InternalImFontAtlasAddDrawListSharedData(atlas, data)
		return goja.Undefined()
	})
	imguiObj.Set("internalImFontAtlasBakedAdd", func(call goja.FunctionCall) goja.Value {
		var atlas *imgui.FontAtlas
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.FontAtlas); ok {
				atlas = v
			}
		}
		// TODO: 处理指针类型 *Font
		var font *imgui.Font
		font_size := float32(call.Argument(2).ToFloat())
		font_rasterizer_density := float32(call.Argument(3).ToFloat())
		// TODO: 处理类型 ID
		baked_id := imgui.ID(0)
		result := imgui.InternalImFontAtlasBakedAdd(atlas, font, font_size, font_rasterizer_density, baked_id)
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalImFontAtlasBakedAddFontGlyph", func(call goja.FunctionCall) goja.Value {
		var atlas *imgui.FontAtlas
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.FontAtlas); ok {
				atlas = v
			}
		}
		// TODO: 处理指针类型 *FontBaked
		var baked *imgui.FontBaked
		// TODO: 处理指针类型 *FontConfig
		var src *imgui.FontConfig
		// TODO: 处理指针类型 *FontGlyph
		var in_glyph *imgui.FontGlyph
		result := imgui.InternalImFontAtlasBakedAddFontGlyph(atlas, baked, src, in_glyph)
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalImFontAtlasBakedAddFontGlyphAdvancedX", func(call goja.FunctionCall) goja.Value {
		var atlas *imgui.FontAtlas
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.FontAtlas); ok {
				atlas = v
			}
		}
		// TODO: 处理指针类型 *FontBaked
		var baked *imgui.FontBaked
		// TODO: 处理指针类型 *FontConfig
		var src *imgui.FontConfig
		// TODO: 处理类型 Wchar
		var codepoint imgui.Wchar
		advance_x := float32(call.Argument(4).ToFloat())
		imgui.InternalImFontAtlasBakedAddFontGlyphAdvancedX(atlas, baked, src, codepoint, advance_x)
		return goja.Undefined()
	})
	imguiObj.Set("internalImFontAtlasBakedDiscard", func(call goja.FunctionCall) goja.Value {
		var atlas *imgui.FontAtlas
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.FontAtlas); ok {
				atlas = v
			}
		}
		// TODO: 处理指针类型 *Font
		var font *imgui.Font
		// TODO: 处理指针类型 *FontBaked
		var baked *imgui.FontBaked
		imgui.InternalImFontAtlasBakedDiscard(atlas, font, baked)
		return goja.Undefined()
	})
	imguiObj.Set("internalImFontAtlasBakedDiscardFontGlyph", func(call goja.FunctionCall) goja.Value {
		var atlas *imgui.FontAtlas
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.FontAtlas); ok {
				atlas = v
			}
		}
		// TODO: 处理指针类型 *Font
		var font *imgui.Font
		// TODO: 处理指针类型 *FontBaked
		var baked *imgui.FontBaked
		// TODO: 处理指针类型 *FontGlyph
		var glyph *imgui.FontGlyph
		imgui.InternalImFontAtlasBakedDiscardFontGlyph(atlas, font, baked, glyph)
		return goja.Undefined()
	})
	imguiObj.Set("internalImFontAtlasBakedGetClosestMatch", func(call goja.FunctionCall) goja.Value {
		var atlas *imgui.FontAtlas
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.FontAtlas); ok {
				atlas = v
			}
		}
		// TODO: 处理指针类型 *Font
		var font *imgui.Font
		font_size := float32(call.Argument(2).ToFloat())
		font_rasterizer_density := float32(call.Argument(3).ToFloat())
		result := imgui.InternalImFontAtlasBakedGetClosestMatch(atlas, font, font_size, font_rasterizer_density)
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalImFontAtlasBakedGetId", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 ID
		font_id := imgui.ID(0)
		baked_size := float32(call.Argument(1).ToFloat())
		rasterizer_density := float32(call.Argument(2).ToFloat())
		return vm.ToValue(imgui.InternalImFontAtlasBakedGetId(font_id, baked_size, rasterizer_density))
	})
	imguiObj.Set("internalImFontAtlasBakedGetOrAdd", func(call goja.FunctionCall) goja.Value {
		var atlas *imgui.FontAtlas
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.FontAtlas); ok {
				atlas = v
			}
		}
		// TODO: 处理指针类型 *Font
		var font *imgui.Font
		font_size := float32(call.Argument(2).ToFloat())
		font_rasterizer_density := float32(call.Argument(3).ToFloat())
		result := imgui.InternalImFontAtlasBakedGetOrAdd(atlas, font, font_size, font_rasterizer_density)
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalImFontAtlasBakedSetFontGlyphBitmap", func(call goja.FunctionCall) goja.Value {
		var atlas *imgui.FontAtlas
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.FontAtlas); ok {
				atlas = v
			}
		}
		// TODO: 处理指针类型 *FontBaked
		var baked *imgui.FontBaked
		// TODO: 处理指针类型 *FontConfig
		var src *imgui.FontConfig
		// TODO: 处理指针类型 *FontGlyph
		var glyph *imgui.FontGlyph
		// TODO: 处理指针类型 *TextureRect
		var r *imgui.TextureRect
		src_pixels := uint(0)
		if len(call.Arguments) >= 6 {
			src_pixels = uint(call.Argument(5).ToInteger())
		}
		// TODO: 处理类型 TextureFormat
		src_fmt := imgui.TextureFormat(0)
		src_pitch := int32(call.Argument(7).ToInteger())
		imgui.InternalImFontAtlasBakedSetFontGlyphBitmap(atlas, baked, src, glyph, r, &src_pixels, src_fmt, src_pitch)
		return goja.Undefined()
	})
	imguiObj.Set("internalImFontAtlasBuildClear", func(call goja.FunctionCall) goja.Value {
		var atlas *imgui.FontAtlas
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.FontAtlas); ok {
				atlas = v
			}
		}
		imgui.InternalImFontAtlasBuildClear(atlas)
		return goja.Undefined()
	})
	imguiObj.Set("internalImFontAtlasBuildDestroy", func(call goja.FunctionCall) goja.Value {
		var atlas *imgui.FontAtlas
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.FontAtlas); ok {
				atlas = v
			}
		}
		imgui.InternalImFontAtlasBuildDestroy(atlas)
		return goja.Undefined()
	})
	imguiObj.Set("internalImFontAtlasBuildDiscardBakes", func(call goja.FunctionCall) goja.Value {
		var atlas *imgui.FontAtlas
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.FontAtlas); ok {
				atlas = v
			}
		}
		unused_frames := int32(call.Argument(1).ToInteger())
		imgui.InternalImFontAtlasBuildDiscardBakes(atlas, unused_frames)
		return goja.Undefined()
	})
	imguiObj.Set("internalImFontAtlasBuildGetOversampleFactors", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *FontConfig
		var src *imgui.FontConfig
		// TODO: 处理指针类型 *FontBaked
		var baked *imgui.FontBaked
		out_oversample_h := int32(0)
		if len(call.Arguments) >= 3 {
			out_oversample_h = int32(call.Argument(2).ToInteger())
		}
		out_oversample_v := int32(0)
		if len(call.Arguments) >= 4 {
			out_oversample_v = int32(call.Argument(3).ToInteger())
		}
		imgui.InternalImFontAtlasBuildGetOversampleFactors(src, baked, &out_oversample_h, &out_oversample_v)
		return goja.Undefined()
	})
	imguiObj.Set("internalImFontAtlasBuildInit", func(call goja.FunctionCall) goja.Value {
		var atlas *imgui.FontAtlas
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.FontAtlas); ok {
				atlas = v
			}
		}
		imgui.InternalImFontAtlasBuildInit(atlas)
		return goja.Undefined()
	})
	imguiObj.Set("internalImFontAtlasBuildLegacyPreloadAllGlyphRanges", func(call goja.FunctionCall) goja.Value {
		var atlas *imgui.FontAtlas
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.FontAtlas); ok {
				atlas = v
			}
		}
		imgui.InternalImFontAtlasBuildLegacyPreloadAllGlyphRanges(atlas)
		return goja.Undefined()
	})
	imguiObj.Set("internalImFontAtlasBuildMain", func(call goja.FunctionCall) goja.Value {
		var atlas *imgui.FontAtlas
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.FontAtlas); ok {
				atlas = v
			}
		}
		imgui.InternalImFontAtlasBuildMain(atlas)
		return goja.Undefined()
	})
	imguiObj.Set("internalImFontAtlasBuildRenderBitmapFromString", func(call goja.FunctionCall) goja.Value {
		var atlas *imgui.FontAtlas
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.FontAtlas); ok {
				atlas = v
			}
		}
		x := int32(call.Argument(1).ToInteger())
		y := int32(call.Argument(2).ToInteger())
		w := int32(call.Argument(3).ToInteger())
		h := int32(call.Argument(4).ToInteger())
		in_str := call.Argument(5).String()
		// TODO: 处理类型 rune
		in_marker_char := rune(0)
		imgui.InternalImFontAtlasBuildRenderBitmapFromString(atlas, x, y, w, h, in_str, in_marker_char)
		return goja.Undefined()
	})
	imguiObj.Set("internalImFontAtlasBuildSetupFontLoader", func(call goja.FunctionCall) goja.Value {
		var atlas *imgui.FontAtlas
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.FontAtlas); ok {
				atlas = v
			}
		}
		// TODO: 处理指针类型 *FontLoader
		var font_loader *imgui.FontLoader
		imgui.InternalImFontAtlasBuildSetupFontLoader(atlas, font_loader)
		return goja.Undefined()
	})
	imguiObj.Set("internalImFontAtlasBuildSetupFontSpecialGlyphs", func(call goja.FunctionCall) goja.Value {
		var atlas *imgui.FontAtlas
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.FontAtlas); ok {
				atlas = v
			}
		}
		// TODO: 处理指针类型 *Font
		var font *imgui.Font
		// TODO: 处理指针类型 *FontConfig
		var src *imgui.FontConfig
		imgui.InternalImFontAtlasBuildSetupFontSpecialGlyphs(atlas, font, src)
		return goja.Undefined()
	})
	imguiObj.Set("internalImFontAtlasBuildUpdatePointers", func(call goja.FunctionCall) goja.Value {
		var atlas *imgui.FontAtlas
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.FontAtlas); ok {
				atlas = v
			}
		}
		imgui.InternalImFontAtlasBuildUpdatePointers(atlas)
		return goja.Undefined()
	})
	imguiObj.Set("internalImFontAtlasDebugLogTextureRequests", func(call goja.FunctionCall) goja.Value {
		var atlas *imgui.FontAtlas
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.FontAtlas); ok {
				atlas = v
			}
		}
		imgui.InternalImFontAtlasDebugLogTextureRequests(atlas)
		return goja.Undefined()
	})
	imguiObj.Set("internalImFontAtlasFontDestroyOutput", func(call goja.FunctionCall) goja.Value {
		var atlas *imgui.FontAtlas
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.FontAtlas); ok {
				atlas = v
			}
		}
		// TODO: 处理指针类型 *Font
		var font *imgui.Font
		imgui.InternalImFontAtlasFontDestroyOutput(atlas, font)
		return goja.Undefined()
	})
	imguiObj.Set("internalImFontAtlasFontDestroySourceData", func(call goja.FunctionCall) goja.Value {
		var atlas *imgui.FontAtlas
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.FontAtlas); ok {
				atlas = v
			}
		}
		// TODO: 处理指针类型 *FontConfig
		var src *imgui.FontConfig
		imgui.InternalImFontAtlasFontDestroySourceData(atlas, src)
		return goja.Undefined()
	})
	imguiObj.Set("internalImFontAtlasFontDiscardBakes", func(call goja.FunctionCall) goja.Value {
		var atlas *imgui.FontAtlas
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.FontAtlas); ok {
				atlas = v
			}
		}
		// TODO: 处理指针类型 *Font
		var font *imgui.Font
		unused_frames := int32(call.Argument(2).ToInteger())
		imgui.InternalImFontAtlasFontDiscardBakes(atlas, font, unused_frames)
		return goja.Undefined()
	})
	imguiObj.Set("internalImFontAtlasFontInitOutput", func(call goja.FunctionCall) goja.Value {
		var atlas *imgui.FontAtlas
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.FontAtlas); ok {
				atlas = v
			}
		}
		// TODO: 处理指针类型 *Font
		var font *imgui.Font
		return vm.ToValue(imgui.InternalImFontAtlasFontInitOutput(atlas, font))
	})
	imguiObj.Set("internalImFontAtlasFontSourceAddToFont", func(call goja.FunctionCall) goja.Value {
		var atlas *imgui.FontAtlas
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.FontAtlas); ok {
				atlas = v
			}
		}
		// TODO: 处理指针类型 *Font
		var font *imgui.Font
		// TODO: 处理指针类型 *FontConfig
		var src *imgui.FontConfig
		imgui.InternalImFontAtlasFontSourceAddToFont(atlas, font, src)
		return goja.Undefined()
	})
	imguiObj.Set("internalImFontAtlasFontSourceInit", func(call goja.FunctionCall) goja.Value {
		var atlas *imgui.FontAtlas
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.FontAtlas); ok {
				atlas = v
			}
		}
		// TODO: 处理指针类型 *FontConfig
		var src *imgui.FontConfig
		return vm.ToValue(imgui.InternalImFontAtlasFontSourceInit(atlas, src))
	})
	imguiObj.Set("internalImFontAtlasGetFontLoaderForStbTruetype", func(call goja.FunctionCall) goja.Value {
		result := imgui.InternalImFontAtlasGetFontLoaderForStbTruetype()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalImFontAtlasGetMouseCursorTexData", func(call goja.FunctionCall) goja.Value {
		var atlas *imgui.FontAtlas
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.FontAtlas); ok {
				atlas = v
			}
		}
		// TODO: 处理类型 MouseCursor
		cursor_type := imgui.MouseCursor(0)
		out_offset := parseVec2(vm, call.Argument(2))
		out_size := parseVec2(vm, call.Argument(3))
		// TODO: 处理指针的数组类型 [2]*Vec2
		out_uv_border := [2]*imgui.Vec2{}
		// TODO: 处理指针的数组类型 [2]*Vec2
		out_uv_fill := [2]*imgui.Vec2{}
		return vm.ToValue(imgui.InternalImFontAtlasGetMouseCursorTexData(atlas, cursor_type, &out_offset, &out_size, out_uv_border, out_uv_fill))
	})
	imguiObj.Set("internalImFontAtlasPackAddRectV", func(call goja.FunctionCall) goja.Value {
		var atlas *imgui.FontAtlas
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.FontAtlas); ok {
				atlas = v
			}
		}
		w := int32(call.Argument(1).ToInteger())
		h := int32(call.Argument(2).ToInteger())
		// TODO: 处理指针类型 *FontAtlasRectEntry
		var overwrite_entry *imgui.FontAtlasRectEntry
		return vm.ToValue(imgui.InternalImFontAtlasPackAddRectV(atlas, w, h, overwrite_entry))
	})
	imguiObj.Set("internalImFontAtlasPackDiscardRect", func(call goja.FunctionCall) goja.Value {
		var atlas *imgui.FontAtlas
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.FontAtlas); ok {
				atlas = v
			}
		}
		// TODO: 处理类型 FontAtlasRectId
		id := imgui.FontAtlasRectId(0)
		imgui.InternalImFontAtlasPackDiscardRect(atlas, id)
		return goja.Undefined()
	})
	imguiObj.Set("internalImFontAtlasPackGetRect", func(call goja.FunctionCall) goja.Value {
		var atlas *imgui.FontAtlas
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.FontAtlas); ok {
				atlas = v
			}
		}
		// TODO: 处理类型 FontAtlasRectId
		id := imgui.FontAtlasRectId(0)
		result := imgui.InternalImFontAtlasPackGetRect(atlas, id)
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalImFontAtlasPackGetRectSafe", func(call goja.FunctionCall) goja.Value {
		var atlas *imgui.FontAtlas
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.FontAtlas); ok {
				atlas = v
			}
		}
		// TODO: 处理类型 FontAtlasRectId
		id := imgui.FontAtlasRectId(0)
		result := imgui.InternalImFontAtlasPackGetRectSafe(atlas, id)
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalImFontAtlasPackInit", func(call goja.FunctionCall) goja.Value {
		var atlas *imgui.FontAtlas
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.FontAtlas); ok {
				atlas = v
			}
		}
		imgui.InternalImFontAtlasPackInit(atlas)
		return goja.Undefined()
	})
	imguiObj.Set("internalImFontAtlasRectIdGetGeneration", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 FontAtlasRectId
		id := imgui.FontAtlasRectId(0)
		return vm.ToValue(imgui.InternalImFontAtlasRectIdGetGeneration(id))
	})
	imguiObj.Set("internalImFontAtlasRectIdGetIndex", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 FontAtlasRectId
		id := imgui.FontAtlasRectId(0)
		return vm.ToValue(imgui.InternalImFontAtlasRectIdGetIndex(id))
	})
	imguiObj.Set("internalImFontAtlasRectIdMake", func(call goja.FunctionCall) goja.Value {
		index_idx := int32(call.Argument(0).ToInteger())
		gen_idx := int32(call.Argument(1).ToInteger())
		return vm.ToValue(imgui.InternalImFontAtlasRectIdMake(index_idx, gen_idx))
	})
	imguiObj.Set("internalImFontAtlasRemoveDrawListSharedData", func(call goja.FunctionCall) goja.Value {
		var atlas *imgui.FontAtlas
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.FontAtlas); ok {
				atlas = v
			}
		}
		// TODO: 处理指针类型 *DrawListSharedData
		var data *imgui.DrawListSharedData
		imgui.InternalImFontAtlasRemoveDrawListSharedData(atlas, data)
		return goja.Undefined()
	})
	imguiObj.Set("internalImFontAtlasTextureAdd", func(call goja.FunctionCall) goja.Value {
		var atlas *imgui.FontAtlas
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.FontAtlas); ok {
				atlas = v
			}
		}
		w := int32(call.Argument(1).ToInteger())
		h := int32(call.Argument(2).ToInteger())
		result := imgui.InternalImFontAtlasTextureAdd(atlas, w, h)
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalImFontAtlasTextureBlockConvert", func(call goja.FunctionCall) goja.Value {
		src_pixels := uint(0)
		if len(call.Arguments) >= 1 {
			src_pixels = uint(call.Argument(0).ToInteger())
		}
		// TODO: 处理类型 TextureFormat
		src_fmt := imgui.TextureFormat(0)
		src_pitch := int32(call.Argument(2).ToInteger())
		dst_pixels := uint(0)
		if len(call.Arguments) >= 4 {
			dst_pixels = uint(call.Argument(3).ToInteger())
		}
		// TODO: 处理类型 TextureFormat
		dst_fmt := imgui.TextureFormat(0)
		dst_pitch := int32(call.Argument(5).ToInteger())
		w := int32(call.Argument(6).ToInteger())
		h := int32(call.Argument(7).ToInteger())
		imgui.InternalImFontAtlasTextureBlockConvert(&src_pixels, src_fmt, src_pitch, &dst_pixels, dst_fmt, dst_pitch, w, h)
		return goja.Undefined()
	})
	imguiObj.Set("internalImFontAtlasTextureBlockCopy", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *TextureData
		var src_tex *imgui.TextureData
		src_x := int32(call.Argument(1).ToInteger())
		src_y := int32(call.Argument(2).ToInteger())
		// TODO: 处理指针类型 *TextureData
		var dst_tex *imgui.TextureData
		dst_x := int32(call.Argument(4).ToInteger())
		dst_y := int32(call.Argument(5).ToInteger())
		w := int32(call.Argument(6).ToInteger())
		h := int32(call.Argument(7).ToInteger())
		imgui.InternalImFontAtlasTextureBlockCopy(src_tex, src_x, src_y, dst_tex, dst_x, dst_y, w, h)
		return goja.Undefined()
	})
	imguiObj.Set("internalImFontAtlasTextureBlockFill", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *TextureData
		var dst_tex *imgui.TextureData
		dst_x := int32(call.Argument(1).ToInteger())
		dst_y := int32(call.Argument(2).ToInteger())
		w := int32(call.Argument(3).ToInteger())
		h := int32(call.Argument(4).ToInteger())
		col := uint32(call.Argument(5).ToInteger())
		imgui.InternalImFontAtlasTextureBlockFill(dst_tex, dst_x, dst_y, w, h, col)
		return goja.Undefined()
	})
	imguiObj.Set("internalImFontAtlasTextureBlockPostProcess", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *FontAtlasPostProcessData
		var data *imgui.FontAtlasPostProcessData
		imgui.InternalImFontAtlasTextureBlockPostProcess(data)
		return goja.Undefined()
	})
	imguiObj.Set("internalImFontAtlasTextureBlockPostProcessMultiply", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *FontAtlasPostProcessData
		var data *imgui.FontAtlasPostProcessData
		multiply_factor := float32(call.Argument(1).ToFloat())
		imgui.InternalImFontAtlasTextureBlockPostProcessMultiply(data, multiply_factor)
		return goja.Undefined()
	})
	imguiObj.Set("internalImFontAtlasTextureBlockQueueUpload", func(call goja.FunctionCall) goja.Value {
		var atlas *imgui.FontAtlas
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.FontAtlas); ok {
				atlas = v
			}
		}
		// TODO: 处理指针类型 *TextureData
		var tex *imgui.TextureData
		x := int32(call.Argument(2).ToInteger())
		y := int32(call.Argument(3).ToInteger())
		w := int32(call.Argument(4).ToInteger())
		h := int32(call.Argument(5).ToInteger())
		imgui.InternalImFontAtlasTextureBlockQueueUpload(atlas, tex, x, y, w, h)
		return goja.Undefined()
	})
	imguiObj.Set("internalImFontAtlasTextureCompact", func(call goja.FunctionCall) goja.Value {
		var atlas *imgui.FontAtlas
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.FontAtlas); ok {
				atlas = v
			}
		}
		imgui.InternalImFontAtlasTextureCompact(atlas)
		return goja.Undefined()
	})
	imguiObj.Set("internalImFontAtlasTextureGetSizeEstimate", func(call goja.FunctionCall) goja.Value {
		var atlas *imgui.FontAtlas
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.FontAtlas); ok {
				atlas = v
			}
		}
		return vm.ToValue(imgui.InternalImFontAtlasTextureGetSizeEstimate(atlas))
	})
	imguiObj.Set("internalImFontAtlasTextureGrowV", func(call goja.FunctionCall) goja.Value {
		var atlas *imgui.FontAtlas
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.FontAtlas); ok {
				atlas = v
			}
		}
		old_w := int32(call.Argument(1).ToInteger())
		old_h := int32(call.Argument(2).ToInteger())
		imgui.InternalImFontAtlasTextureGrowV(atlas, old_w, old_h)
		return goja.Undefined()
	})
	imguiObj.Set("internalImFontAtlasTextureMakeSpace", func(call goja.FunctionCall) goja.Value {
		var atlas *imgui.FontAtlas
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.FontAtlas); ok {
				atlas = v
			}
		}
		imgui.InternalImFontAtlasTextureMakeSpace(atlas)
		return goja.Undefined()
	})
	imguiObj.Set("internalImFontAtlasTextureRepack", func(call goja.FunctionCall) goja.Value {
		var atlas *imgui.FontAtlas
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.FontAtlas); ok {
				atlas = v
			}
		}
		w := int32(call.Argument(1).ToInteger())
		h := int32(call.Argument(2).ToInteger())
		imgui.InternalImFontAtlasTextureRepack(atlas, w, h)
		return goja.Undefined()
	})
	imguiObj.Set("internalImFontAtlasUpdateDrawListsSharedData", func(call goja.FunctionCall) goja.Value {
		var atlas *imgui.FontAtlas
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.FontAtlas); ok {
				atlas = v
			}
		}
		imgui.InternalImFontAtlasUpdateDrawListsSharedData(atlas)
		return goja.Undefined()
	})
	imguiObj.Set("internalImFontAtlasUpdateDrawListsTextures", func(call goja.FunctionCall) goja.Value {
		var atlas *imgui.FontAtlas
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.FontAtlas); ok {
				atlas = v
			}
		}
		// TODO: 处理类型 TextureRef
		old_tex := imgui.TextureRef{}
		// TODO: 处理类型 TextureRef
		new_tex := imgui.TextureRef{}
		imgui.InternalImFontAtlasUpdateDrawListsTextures(atlas, old_tex, new_tex)
		return goja.Undefined()
	})
	imguiObj.Set("internalImFontAtlasUpdateNewFrame", func(call goja.FunctionCall) goja.Value {
		var atlas *imgui.FontAtlas
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.FontAtlas); ok {
				atlas = v
			}
		}
		frame_count := int32(call.Argument(1).ToInteger())
		renderer_has_textures := call.Argument(2).ToBoolean()
		imgui.InternalImFontAtlasUpdateNewFrame(atlas, frame_count, renderer_has_textures)
		return goja.Undefined()
	})
	imguiObj.Set("internalImFontCalcTextSizeEx", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Font
		var font *imgui.Font
		size := float32(call.Argument(1).ToFloat())
		max_width := float32(call.Argument(2).ToFloat())
		wrap_width := float32(call.Argument(3).ToFloat())
		text_begin := call.Argument(4).String()
		text_end_display := call.Argument(5).String()
		out_remaining := parseStringArray(vm, call.Argument(6))
		out_offset := parseVec2(vm, call.Argument(7))
		// TODO: 处理类型 DrawTextFlags
		flags := imgui.DrawTextFlags(0)
		result := imgui.InternalImFontCalcTextSizeEx(font, size, max_width, wrap_width, text_begin, text_end_display, out_remaining, &out_offset, flags)
		return vm.ToValue(result)
	})
	imguiObj.Set("internalImFontCalcWordWrapPositionExV", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Font
		var font *imgui.Font
		size := float32(call.Argument(1).ToFloat())
		text := call.Argument(2).String()
		wrap_width := float32(call.Argument(3).ToFloat())
		// TODO: 处理类型 DrawTextFlags
		flags := imgui.DrawTextFlags(0)
		return vm.ToValue(imgui.InternalImFontCalcWordWrapPositionExV(font, size, text, wrap_width, flags))
	})
	imguiObj.Set("internalImFormatString", func(call goja.FunctionCall) goja.Value {
		buf := call.Argument(0).String()
		buf_size := uint64(call.Argument(1).ToInteger())
		fmt := call.Argument(2).String()
		return vm.ToValue(imgui.InternalImFormatString(buf, buf_size, fmt))
	})
	imguiObj.Set("internalImFormatStringToTempBuffer", func(call goja.FunctionCall) goja.Value {
		out_buf := parseStringArray(vm, call.Argument(0))
		out_buf_end := parseStringArray(vm, call.Argument(1))
		fmt := call.Argument(2).String()
		imgui.InternalImFormatStringToTempBuffer(out_buf, out_buf_end, fmt)
		return goja.Undefined()
	})
	imguiObj.Set("internalImHashDataV", func(call goja.FunctionCall) goja.Value {
		data := uintptr(call.Argument(0).ToInteger())
		data_size := uint64(call.Argument(1).ToInteger())
		// TODO: 处理类型 ID
		seed := imgui.ID(0)
		return vm.ToValue(imgui.InternalImHashDataV(data, data_size, seed))
	})
	imguiObj.Set("internalImHashSkipUncontributingPrefix", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		return vm.ToValue(imgui.InternalImHashSkipUncontributingPrefix(label))
	})
	imguiObj.Set("internalImHashStrV", func(call goja.FunctionCall) goja.Value {
		data := call.Argument(0).String()
		data_size := uint64(call.Argument(1).ToInteger())
		// TODO: 处理类型 ID
		seed := imgui.ID(0)
		return vm.ToValue(imgui.InternalImHashStrV(data, data_size, seed))
	})
	imguiObj.Set("internalImInvLength", func(call goja.FunctionCall) goja.Value {
		lhs := parseVec2(vm, call.Argument(0))
		fail_value := float32(call.Argument(0).ToFloat())
		return vm.ToValue(imgui.InternalImInvLength(lhs, fail_value))
	})
	imguiObj.Set("internalImIsFloatAboveGuaranteedIntegerPrecision", func(call goja.FunctionCall) goja.Value {
		f := float32(call.Argument(0).ToFloat())
		return vm.ToValue(imgui.InternalImIsFloatAboveGuaranteedIntegerPrecision(f))
	})
	imguiObj.Set("internalImIsPowerOfTwoInt", func(call goja.FunctionCall) goja.Value {
		v := int32(call.Argument(0).ToInteger())
		return vm.ToValue(imgui.InternalImIsPowerOfTwoInt(v))
	})
	imguiObj.Set("internalImIsPowerOfTwoU64", func(call goja.FunctionCall) goja.Value {
		v := uint64(call.Argument(0).ToInteger())
		return vm.ToValue(imgui.InternalImIsPowerOfTwoU64(v))
	})
	imguiObj.Set("internalImLengthSqrVec2", func(call goja.FunctionCall) goja.Value {
		lhs := parseVec2(vm, call.Argument(0))
		return vm.ToValue(imgui.InternalImLengthSqrVec2(lhs))
	})
	imguiObj.Set("internalImLengthSqrVec4", func(call goja.FunctionCall) goja.Value {
		lhs := parseVec4(vm, call.Argument(0))
		return vm.ToValue(imgui.InternalImLengthSqrVec4(lhs))
	})
	imguiObj.Set("internalImLerpVec2Float", func(call goja.FunctionCall) goja.Value {
		a := parseVec2(vm, call.Argument(0))
		b := parseVec2(vm, call.Argument(0))
		t := float32(call.Argument(0).ToFloat())
		result := imgui.InternalImLerpVec2Float(a, b, t)
		return vm.ToValue(result)
	})
	imguiObj.Set("internalImLerpVec2Vec2", func(call goja.FunctionCall) goja.Value {
		a := parseVec2(vm, call.Argument(0))
		b := parseVec2(vm, call.Argument(0))
		t := parseVec2(vm, call.Argument(0))
		result := imgui.InternalImLerpVec2Vec2(a, b, t)
		return vm.ToValue(result)
	})
	imguiObj.Set("internalImLerpVec4", func(call goja.FunctionCall) goja.Value {
		a := parseVec4(vm, call.Argument(0))
		b := parseVec4(vm, call.Argument(0))
		t := float32(call.Argument(0).ToFloat())
		result := imgui.InternalImLerpVec4(a, b, t)
		return vm.ToValue(result)
	})
	imguiObj.Set("internalImLineClosestPoint", func(call goja.FunctionCall) goja.Value {
		a := parseVec2(vm, call.Argument(0))
		b := parseVec2(vm, call.Argument(0))
		p := parseVec2(vm, call.Argument(0))
		result := imgui.InternalImLineClosestPoint(a, b, p)
		return vm.ToValue(result)
	})
	imguiObj.Set("internalImLinearRemapClamp", func(call goja.FunctionCall) goja.Value {
		s0 := float32(call.Argument(0).ToFloat())
		s1 := float32(call.Argument(1).ToFloat())
		d0 := float32(call.Argument(2).ToFloat())
		d1 := float32(call.Argument(3).ToFloat())
		x := float32(call.Argument(4).ToFloat())
		return vm.ToValue(imgui.InternalImLinearRemapClamp(s0, s1, d0, d1, x))
	})
	imguiObj.Set("internalImLinearSweep", func(call goja.FunctionCall) goja.Value {
		current := float32(call.Argument(0).ToFloat())
		target := float32(call.Argument(1).ToFloat())
		speed := float32(call.Argument(2).ToFloat())
		return vm.ToValue(imgui.InternalImLinearSweep(current, target, speed))
	})
	imguiObj.Set("internalImLogFloat", func(call goja.FunctionCall) goja.Value {
		x := float32(call.Argument(0).ToFloat())
		return vm.ToValue(imgui.InternalImLogFloat(x))
	})
	imguiObj.Set("internalImLogDouble", func(call goja.FunctionCall) goja.Value {
		x := call.Argument(0).ToFloat()
		return vm.ToValue(imgui.InternalImLogDouble(x))
	})
	imguiObj.Set("internalImLowerBound", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *StoragePair
		var in_begin *imgui.StoragePair
		// TODO: 处理指针类型 *StoragePair
		var in_end *imgui.StoragePair
		// TODO: 处理类型 ID
		key := imgui.ID(0)
		result := imgui.InternalImLowerBound(in_begin, in_end, key)
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalImMax", func(call goja.FunctionCall) goja.Value {
		lhs := parseVec2(vm, call.Argument(0))
		rhs := parseVec2(vm, call.Argument(0))
		result := imgui.InternalImMax(lhs, rhs)
		return vm.ToValue(result)
	})
	imguiObj.Set("internalImMemdup", func(call goja.FunctionCall) goja.Value {
		src := uintptr(call.Argument(0).ToInteger())
		size := uint64(call.Argument(1).ToInteger())
		return vm.ToValue(imgui.InternalImMemdup(src, size))
	})
	imguiObj.Set("internalImMin", func(call goja.FunctionCall) goja.Value {
		lhs := parseVec2(vm, call.Argument(0))
		rhs := parseVec2(vm, call.Argument(0))
		result := imgui.InternalImMin(lhs, rhs)
		return vm.ToValue(result)
	})
	imguiObj.Set("internalImModPositive", func(call goja.FunctionCall) goja.Value {
		a := int32(call.Argument(0).ToInteger())
		b := int32(call.Argument(1).ToInteger())
		return vm.ToValue(imgui.InternalImModPositive(a, b))
	})
	imguiObj.Set("internalImMul", func(call goja.FunctionCall) goja.Value {
		lhs := parseVec2(vm, call.Argument(0))
		rhs := parseVec2(vm, call.Argument(0))
		result := imgui.InternalImMul(lhs, rhs)
		return vm.ToValue(result)
	})
	imguiObj.Set("internalImParseFormatFindEnd", func(call goja.FunctionCall) goja.Value {
		format := call.Argument(0).String()
		return vm.ToValue(imgui.InternalImParseFormatFindEnd(format))
	})
	imguiObj.Set("internalImParseFormatFindStart", func(call goja.FunctionCall) goja.Value {
		format := call.Argument(0).String()
		return vm.ToValue(imgui.InternalImParseFormatFindStart(format))
	})
	imguiObj.Set("internalImParseFormatPrecision", func(call goja.FunctionCall) goja.Value {
		format := call.Argument(0).String()
		default_value := int32(call.Argument(1).ToInteger())
		return vm.ToValue(imgui.InternalImParseFormatPrecision(format, default_value))
	})
	imguiObj.Set("internalImParseFormatSanitizeForPrinting", func(call goja.FunctionCall) goja.Value {
		fmt_in := call.Argument(0).String()
		fmt_out := call.Argument(1).String()
		fmt_out_size := uint64(call.Argument(2).ToInteger())
		imgui.InternalImParseFormatSanitizeForPrinting(fmt_in, fmt_out, fmt_out_size)
		return goja.Undefined()
	})
	imguiObj.Set("internalImParseFormatSanitizeForScanning", func(call goja.FunctionCall) goja.Value {
		fmt_in := call.Argument(0).String()
		fmt_out := call.Argument(1).String()
		fmt_out_size := uint64(call.Argument(2).ToInteger())
		return vm.ToValue(imgui.InternalImParseFormatSanitizeForScanning(fmt_in, fmt_out, fmt_out_size))
	})
	imguiObj.Set("internalImParseFormatTrimDecorations", func(call goja.FunctionCall) goja.Value {
		format := call.Argument(0).String()
		buf := call.Argument(1).String()
		buf_size := uint64(call.Argument(2).ToInteger())
		return vm.ToValue(imgui.InternalImParseFormatTrimDecorations(format, buf, buf_size))
	})
	imguiObj.Set("internalImPowFloat", func(call goja.FunctionCall) goja.Value {
		x := float32(call.Argument(0).ToFloat())
		y := float32(call.Argument(1).ToFloat())
		return vm.ToValue(imgui.InternalImPowFloat(x, y))
	})
	imguiObj.Set("internalImPowDouble", func(call goja.FunctionCall) goja.Value {
		x := call.Argument(0).ToFloat()
		y := call.Argument(1).ToFloat()
		return vm.ToValue(imgui.InternalImPowDouble(x, y))
	})
	imguiObj.Set("internalImRotate", func(call goja.FunctionCall) goja.Value {
		v := parseVec2(vm, call.Argument(0))
		cos_a := float32(call.Argument(0).ToFloat())
		sin_a := float32(call.Argument(1).ToFloat())
		result := imgui.InternalImRotate(v, cos_a, sin_a)
		return vm.ToValue(result)
	})
	imguiObj.Set("internalImRound64", func(call goja.FunctionCall) goja.Value {
		f := float32(call.Argument(0).ToFloat())
		return vm.ToValue(imgui.InternalImRound64(f))
	})
	imguiObj.Set("internalImRsqrtFloat", func(call goja.FunctionCall) goja.Value {
		x := float32(call.Argument(0).ToFloat())
		return vm.ToValue(imgui.InternalImRsqrtFloat(x))
	})
	imguiObj.Set("internalImRsqrtDouble", func(call goja.FunctionCall) goja.Value {
		x := call.Argument(0).ToFloat()
		return vm.ToValue(imgui.InternalImRsqrtDouble(x))
	})
	imguiObj.Set("internalImSaturate", func(call goja.FunctionCall) goja.Value {
		f := float32(call.Argument(0).ToFloat())
		return vm.ToValue(imgui.InternalImSaturate(f))
	})
	imguiObj.Set("internalImSignFloat", func(call goja.FunctionCall) goja.Value {
		x := float32(call.Argument(0).ToFloat())
		return vm.ToValue(imgui.InternalImSignFloat(x))
	})
	imguiObj.Set("internalImSignDouble", func(call goja.FunctionCall) goja.Value {
		x := call.Argument(0).ToFloat()
		return vm.ToValue(imgui.InternalImSignDouble(x))
	})
	imguiObj.Set("internalImStrSkipBlank", func(call goja.FunctionCall) goja.Value {
		str := call.Argument(0).String()
		return vm.ToValue(imgui.InternalImStrSkipBlank(str))
	})
	imguiObj.Set("internalImStrTrimBlanks", func(call goja.FunctionCall) goja.Value {
		str := call.Argument(0).String()
		imgui.InternalImStrTrimBlanks(str)
		return goja.Undefined()
	})
	imguiObj.Set("internalImStrbol", func(call goja.FunctionCall) goja.Value {
		buf_mid_line := call.Argument(0).String()
		buf_begin := call.Argument(1).String()
		return vm.ToValue(imgui.InternalImStrbol(buf_mid_line, buf_begin))
	})
	imguiObj.Set("internalImStrchrRange", func(call goja.FunctionCall) goja.Value {
		str_begin := call.Argument(0).String()
		str_end := call.Argument(1).String()
		// TODO: 处理类型 rune
		c := rune(0)
		return vm.ToValue(imgui.InternalImStrchrRange(str_begin, str_end, c))
	})
	imguiObj.Set("internalImStrdup", func(call goja.FunctionCall) goja.Value {
		str := call.Argument(0).String()
		return vm.ToValue(imgui.InternalImStrdup(str))
	})
	imguiObj.Set("internalImStrdupcpy", func(call goja.FunctionCall) goja.Value {
		dst := call.Argument(0).String()
		// TODO: 处理指针类型 *uint64
		var p_dst_size *uint64
		str := call.Argument(2).String()
		return vm.ToValue(imgui.InternalImStrdupcpy(dst, p_dst_size, str))
	})
	imguiObj.Set("internalImStreolRange", func(call goja.FunctionCall) goja.Value {
		str := call.Argument(0).String()
		str_end := call.Argument(1).String()
		return vm.ToValue(imgui.InternalImStreolRange(str, str_end))
	})
	imguiObj.Set("internalImStricmp", func(call goja.FunctionCall) goja.Value {
		str1 := call.Argument(0).String()
		str2 := call.Argument(1).String()
		return vm.ToValue(imgui.InternalImStricmp(str1, str2))
	})
	imguiObj.Set("internalImStristr", func(call goja.FunctionCall) goja.Value {
		haystack := call.Argument(0).String()
		haystack_end := call.Argument(1).String()
		needle := call.Argument(2).String()
		needle_end := call.Argument(3).String()
		return vm.ToValue(imgui.InternalImStristr(haystack, haystack_end, needle, needle_end))
	})
	imguiObj.Set("internalImStrlenW", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Wchar
		var str *imgui.Wchar
		return vm.ToValue(imgui.InternalImStrlenW(str))
	})
	imguiObj.Set("internalImStrncpy", func(call goja.FunctionCall) goja.Value {
		dst := call.Argument(0).String()
		src := call.Argument(1).String()
		count := uint64(call.Argument(2).ToInteger())
		imgui.InternalImStrncpy(dst, src, count)
		return goja.Undefined()
	})
	imguiObj.Set("internalImStrnicmp", func(call goja.FunctionCall) goja.Value {
		str1 := call.Argument(0).String()
		str2 := call.Argument(1).String()
		count := uint64(call.Argument(2).ToInteger())
		return vm.ToValue(imgui.InternalImStrnicmp(str1, str2, count))
	})
	imguiObj.Set("internalImTextCalcWordWrapNextLineStartV", func(call goja.FunctionCall) goja.Value {
		text := call.Argument(0).String()
		// TODO: 处理类型 DrawTextFlags
		flags := imgui.DrawTextFlags(0)
		return vm.ToValue(imgui.InternalImTextCalcWordWrapNextLineStartV(text, flags))
	})
	imguiObj.Set("internalImTextCharFromUtf8", func(call goja.FunctionCall) goja.Value {
		out_char := uint32(0)
		if len(call.Arguments) >= 1 {
			out_char = uint32(call.Argument(0).ToInteger())
		}
		in_text := call.Argument(1).String()
		in_text_end := call.Argument(2).String()
		return vm.ToValue(imgui.InternalImTextCharFromUtf8(&out_char, in_text, in_text_end))
	})
	imguiObj.Set("internalImTextCharToUtf8", func(call goja.FunctionCall) goja.Value {
		out_buf := [5]rune{0, 0, 0, 0, 0}
		c := uint32(call.Argument(1).ToInteger())
		return vm.ToValue(imgui.InternalImTextCharToUtf8(&out_buf, c))
	})
	imguiObj.Set("internalImTextCountCharsFromUtf8", func(call goja.FunctionCall) goja.Value {
		in_text := call.Argument(0).String()
		in_text_end := call.Argument(1).String()
		return vm.ToValue(imgui.InternalImTextCountCharsFromUtf8(in_text, in_text_end))
	})
	imguiObj.Set("internalImTextCountLines", func(call goja.FunctionCall) goja.Value {
		in_text := call.Argument(0).String()
		in_text_end := call.Argument(1).String()
		return vm.ToValue(imgui.InternalImTextCountLines(in_text, in_text_end))
	})
	imguiObj.Set("internalImTextCountUtf8BytesFromChar", func(call goja.FunctionCall) goja.Value {
		in_text := call.Argument(0).String()
		in_text_end := call.Argument(1).String()
		return vm.ToValue(imgui.InternalImTextCountUtf8BytesFromChar(in_text, in_text_end))
	})
	imguiObj.Set("internalImTextCountUtf8BytesFromStr", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Wchar
		var in_text *imgui.Wchar
		// TODO: 处理指针类型 *Wchar
		var in_text_end *imgui.Wchar
		return vm.ToValue(imgui.InternalImTextCountUtf8BytesFromStr(in_text, in_text_end))
	})
	imguiObj.Set("internalImTextFindPreviousUtf8Codepoint", func(call goja.FunctionCall) goja.Value {
		in_text_start := call.Argument(0).String()
		in_text_curr := call.Argument(1).String()
		return vm.ToValue(imgui.InternalImTextFindPreviousUtf8Codepoint(in_text_start, in_text_curr))
	})
	imguiObj.Set("internalImTextStrFromUtf8V", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Wchar
		var out_buf *imgui.Wchar
		out_buf_size := int32(call.Argument(1).ToInteger())
		in_text := call.Argument(2).String()
		in_text_end := call.Argument(3).String()
		in_remaining := parseStringArray(vm, call.Argument(4))
		return vm.ToValue(imgui.InternalImTextStrFromUtf8V(out_buf, out_buf_size, in_text, in_text_end, in_remaining))
	})
	imguiObj.Set("internalImTextStrToUtf8", func(call goja.FunctionCall) goja.Value {
		out_buf := call.Argument(0).String()
		out_buf_size := int32(call.Argument(1).ToInteger())
		// TODO: 处理指针类型 *Wchar
		var in_text *imgui.Wchar
		// TODO: 处理指针类型 *Wchar
		var in_text_end *imgui.Wchar
		return vm.ToValue(imgui.InternalImTextStrToUtf8(out_buf, out_buf_size, in_text, in_text_end))
	})
	imguiObj.Set("internalImTextureDataGetFormatBytesPerPixel", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 TextureFormat
		format := imgui.TextureFormat(0)
		return vm.ToValue(imgui.InternalImTextureDataGetFormatBytesPerPixel(format))
	})
	imguiObj.Set("internalImTextureDataGetFormatName", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 TextureFormat
		format := imgui.TextureFormat(0)
		return vm.ToValue(imgui.InternalImTextureDataGetFormatName(format))
	})
	imguiObj.Set("internalImTextureDataGetStatusName", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 TextureStatus
		status := imgui.TextureStatus(0)
		return vm.ToValue(imgui.InternalImTextureDataGetStatusName(status))
	})
	imguiObj.Set("internalImToUpper", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 rune
		c := rune(0)
		return vm.ToValue(imgui.InternalImToUpper(c))
	})
	imguiObj.Set("internalImTriangleArea", func(call goja.FunctionCall) goja.Value {
		a := parseVec2(vm, call.Argument(0))
		b := parseVec2(vm, call.Argument(0))
		c := parseVec2(vm, call.Argument(0))
		return vm.ToValue(imgui.InternalImTriangleArea(a, b, c))
	})
	imguiObj.Set("internalImTriangleBarycentricCoords", func(call goja.FunctionCall) goja.Value {
		a := parseVec2(vm, call.Argument(0))
		b := parseVec2(vm, call.Argument(0))
		c := parseVec2(vm, call.Argument(0))
		p := parseVec2(vm, call.Argument(0))
		out_u := float32(0)
		if len(call.Arguments) >= 1 {
			out_u = float32(call.Argument(0).ToFloat())
		}
		out_v := float32(0)
		if len(call.Arguments) >= 2 {
			out_v = float32(call.Argument(1).ToFloat())
		}
		out_w := float32(0)
		if len(call.Arguments) >= 3 {
			out_w = float32(call.Argument(2).ToFloat())
		}
		imgui.InternalImTriangleBarycentricCoords(a, b, c, p, &out_u, &out_v, &out_w)
		return goja.Undefined()
	})
	imguiObj.Set("internalImTriangleClosestPoint", func(call goja.FunctionCall) goja.Value {
		a := parseVec2(vm, call.Argument(0))
		b := parseVec2(vm, call.Argument(0))
		c := parseVec2(vm, call.Argument(0))
		p := parseVec2(vm, call.Argument(0))
		result := imgui.InternalImTriangleClosestPoint(a, b, c, p)
		return vm.ToValue(result)
	})
	imguiObj.Set("internalImTriangleContainsPoint", func(call goja.FunctionCall) goja.Value {
		a := parseVec2(vm, call.Argument(0))
		b := parseVec2(vm, call.Argument(0))
		c := parseVec2(vm, call.Argument(0))
		p := parseVec2(vm, call.Argument(0))
		return vm.ToValue(imgui.InternalImTriangleContainsPoint(a, b, c, p))
	})
	imguiObj.Set("internalImTriangleIsClockwise", func(call goja.FunctionCall) goja.Value {
		a := parseVec2(vm, call.Argument(0))
		b := parseVec2(vm, call.Argument(0))
		c := parseVec2(vm, call.Argument(0))
		return vm.ToValue(imgui.InternalImTriangleIsClockwise(a, b, c))
	})
	imguiObj.Set("internalImTrunc64", func(call goja.FunctionCall) goja.Value {
		f := float32(call.Argument(0).ToFloat())
		return vm.ToValue(imgui.InternalImTrunc64(f))
	})
	imguiObj.Set("internalImTruncFloat", func(call goja.FunctionCall) goja.Value {
		f := float32(call.Argument(0).ToFloat())
		return vm.ToValue(imgui.InternalImTruncFloat(f))
	})
	imguiObj.Set("internalImTruncVec2", func(call goja.FunctionCall) goja.Value {
		v := parseVec2(vm, call.Argument(0))
		result := imgui.InternalImTruncVec2(v)
		return vm.ToValue(result)
	})
	imguiObj.Set("internalImUpperPowerOfTwo", func(call goja.FunctionCall) goja.Value {
		v := int32(call.Argument(0).ToInteger())
		return vm.ToValue(imgui.InternalImUpperPowerOfTwo(v))
	})
	imguiObj.Set("imageV", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 TextureRef
		tex_ref := imgui.TextureRef{}
		image_size := parseVec2(vm, call.Argument(1))
		uv0 := parseVec2(vm, call.Argument(1))
		uv1 := parseVec2(vm, call.Argument(1))
		imgui.ImageV(tex_ref, image_size, uv0, uv1)
		return goja.Undefined()
	})
	imguiObj.Set("imageButtonV", func(call goja.FunctionCall) goja.Value {
		str_id := call.Argument(0).String()
		// TODO: 处理类型 TextureRef
		tex_ref := imgui.TextureRef{}
		image_size := parseVec2(vm, call.Argument(2))
		uv0 := parseVec2(vm, call.Argument(2))
		uv1 := parseVec2(vm, call.Argument(2))
		bg_col := parseVec4(vm, call.Argument(2))
		tint_col := parseVec4(vm, call.Argument(2))
		return vm.ToValue(imgui.ImageButtonV(str_id, tex_ref, image_size, uv0, uv1, bg_col, tint_col))
	})
	imguiObj.Set("internalImageButtonExV", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 ID
		id := imgui.ID(0)
		// TODO: 处理类型 TextureRef
		tex_ref := imgui.TextureRef{}
		image_size := parseVec2(vm, call.Argument(2))
		uv0 := parseVec2(vm, call.Argument(2))
		uv1 := parseVec2(vm, call.Argument(2))
		bg_col := parseVec4(vm, call.Argument(2))
		tint_col := parseVec4(vm, call.Argument(2))
		// TODO: 处理类型 ButtonFlags
		flags := imgui.ButtonFlags(0)
		return vm.ToValue(imgui.InternalImageButtonExV(id, tex_ref, image_size, uv0, uv1, bg_col, tint_col, flags))
	})
	imguiObj.Set("imageWithBgV", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 TextureRef
		tex_ref := imgui.TextureRef{}
		image_size := parseVec2(vm, call.Argument(1))
		uv0 := parseVec2(vm, call.Argument(1))
		uv1 := parseVec2(vm, call.Argument(1))
		bg_col := parseVec4(vm, call.Argument(1))
		tint_col := parseVec4(vm, call.Argument(1))
		imgui.ImageWithBgV(tex_ref, image_size, uv0, uv1, bg_col, tint_col)
		return goja.Undefined()
	})
	imguiObj.Set("indentV", func(call goja.FunctionCall) goja.Value {
		indent_w := float32(call.Argument(0).ToFloat())
		imgui.IndentV(indent_w)
		return goja.Undefined()
	})
	imguiObj.Set("internalInitialize", func(call goja.FunctionCall) goja.Value {
		imgui.InternalInitialize()
		return goja.Undefined()
	})
	imguiObj.Set("inputDoubleV", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		v := float64(0)
		if len(call.Arguments) >= 2 {
			v = float64(call.Argument(1).ToFloat())
		}
		step := call.Argument(2).ToFloat()
		step_fast := call.Argument(3).ToFloat()
		format := call.Argument(4).String()
		// TODO: 处理类型 InputTextFlags
		flags := imgui.InputTextFlags(0)
		return vm.ToValue(imgui.InputDoubleV(label, &v, step, step_fast, format, flags))
	})
	imguiObj.Set("inputFloatV", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		v := float32(0)
		if len(call.Arguments) >= 2 {
			v = float32(call.Argument(1).ToFloat())
		}
		step := float32(call.Argument(2).ToFloat())
		step_fast := float32(call.Argument(3).ToFloat())
		format := call.Argument(4).String()
		// TODO: 处理类型 InputTextFlags
		flags := imgui.InputTextFlags(0)
		return vm.ToValue(imgui.InputFloatV(label, &v, step, step_fast, format, flags))
	})
	imguiObj.Set("inputFloat2V", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		v := [2]float32{0, 0}
		format := call.Argument(2).String()
		// TODO: 处理类型 InputTextFlags
		flags := imgui.InputTextFlags(0)
		return vm.ToValue(imgui.InputFloat2V(label, &v, format, flags))
	})
	imguiObj.Set("inputFloat3V", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		v := [3]float32{0, 0, 0}
		format := call.Argument(2).String()
		// TODO: 处理类型 InputTextFlags
		flags := imgui.InputTextFlags(0)
		return vm.ToValue(imgui.InputFloat3V(label, &v, format, flags))
	})
	imguiObj.Set("inputFloat4V", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		v := [4]float32{0, 0, 0, 0}
		format := call.Argument(2).String()
		// TODO: 处理类型 InputTextFlags
		flags := imgui.InputTextFlags(0)
		return vm.ToValue(imgui.InputFloat4V(label, &v, format, flags))
	})
	imguiObj.Set("inputIntV", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		v := int32(0)
		if len(call.Arguments) >= 2 {
			v = int32(call.Argument(1).ToInteger())
		}
		step := int32(call.Argument(2).ToInteger())
		step_fast := int32(call.Argument(3).ToInteger())
		// TODO: 处理类型 InputTextFlags
		flags := imgui.InputTextFlags(0)
		return vm.ToValue(imgui.InputIntV(label, &v, step, step_fast, flags))
	})
	imguiObj.Set("inputInt2V", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		v := [2]int32{0, 0}
		// TODO: 处理类型 InputTextFlags
		flags := imgui.InputTextFlags(0)
		return vm.ToValue(imgui.InputInt2V(label, &v, flags))
	})
	imguiObj.Set("inputInt3V", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		v := [3]int32{0, 0, 0}
		// TODO: 处理类型 InputTextFlags
		flags := imgui.InputTextFlags(0)
		return vm.ToValue(imgui.InputInt3V(label, &v, flags))
	})
	imguiObj.Set("inputInt4V", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		v := [4]int32{0, 0, 0, 0}
		// TODO: 处理类型 InputTextFlags
		flags := imgui.InputTextFlags(0)
		return vm.ToValue(imgui.InputInt4V(label, &v, flags))
	})
	imguiObj.Set("inputScalarV", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		// TODO: 处理类型 DataType
		data_type := imgui.DataType(0)
		p_data := uintptr(call.Argument(2).ToInteger())
		p_step := uintptr(call.Argument(3).ToInteger())
		p_step_fast := uintptr(call.Argument(4).ToInteger())
		format := call.Argument(5).String()
		// TODO: 处理类型 InputTextFlags
		flags := imgui.InputTextFlags(0)
		return vm.ToValue(imgui.InputScalarV(label, data_type, p_data, p_step, p_step_fast, format, flags))
	})
	imguiObj.Set("inputScalarNV", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		// TODO: 处理类型 DataType
		data_type := imgui.DataType(0)
		p_data := uintptr(call.Argument(2).ToInteger())
		components := int32(call.Argument(3).ToInteger())
		p_step := uintptr(call.Argument(4).ToInteger())
		p_step_fast := uintptr(call.Argument(5).ToInteger())
		format := call.Argument(6).String()
		// TODO: 处理类型 InputTextFlags
		flags := imgui.InputTextFlags(0)
		return vm.ToValue(imgui.InputScalarNV(label, data_type, p_data, components, p_step, p_step_fast, format, flags))
	})
	imguiObj.Set("internalInputTextDeactivateHook", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 ID
		id := imgui.ID(0)
		imgui.InternalInputTextDeactivateHook(id)
		return goja.Undefined()
	})
	imguiObj.Set("invisibleButtonV", func(call goja.FunctionCall) goja.Value {
		str_id := call.Argument(0).String()
		size := parseVec2(vm, call.Argument(1))
		// TODO: 处理类型 ButtonFlags
		flags := imgui.ButtonFlags(0)
		return vm.ToValue(imgui.InvisibleButtonV(str_id, size, flags))
	})
	imguiObj.Set("internalIsActiveIdUsingNavDir", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 Dir
		dir := imgui.Dir(0)
		return vm.ToValue(imgui.InternalIsActiveIdUsingNavDir(dir))
	})
	imguiObj.Set("internalIsAliasKey", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 Key
		key := imgui.Key(0)
		return vm.ToValue(imgui.InternalIsAliasKey(key))
	})
	imguiObj.Set("isAnyItemActive", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.IsAnyItemActive())
	})
	imguiObj.Set("isAnyItemFocused", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.IsAnyItemFocused())
	})
	imguiObj.Set("isAnyItemHovered", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.IsAnyItemHovered())
	})
	imguiObj.Set("isAnyMouseDown", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.IsAnyMouseDown())
	})
	imguiObj.Set("internalIsClippedEx", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 Rect
		bb := imgui.Rect{}
		// TODO: 处理类型 ID
		id := imgui.ID(0)
		return vm.ToValue(imgui.InternalIsClippedEx(bb, id))
	})
	imguiObj.Set("internalIsDragDropActive", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.InternalIsDragDropActive())
	})
	imguiObj.Set("internalIsDragDropPayloadBeingAccepted", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.InternalIsDragDropPayloadBeingAccepted())
	})
	imguiObj.Set("internalIsGamepadKey", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 Key
		key := imgui.Key(0)
		return vm.ToValue(imgui.InternalIsGamepadKey(key))
	})
	imguiObj.Set("isItemActivated", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.IsItemActivated())
	})
	imguiObj.Set("isItemActive", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.IsItemActive())
	})
	imguiObj.Set("internalIsItemActiveAsInputText", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.InternalIsItemActiveAsInputText())
	})
	imguiObj.Set("isItemClickedV", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 MouseButton
		mouse_button := imgui.MouseButton(0)
		return vm.ToValue(imgui.IsItemClickedV(mouse_button))
	})
	imguiObj.Set("isItemDeactivated", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.IsItemDeactivated())
	})
	imguiObj.Set("isItemDeactivatedAfterEdit", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.IsItemDeactivatedAfterEdit())
	})
	imguiObj.Set("isItemEdited", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.IsItemEdited())
	})
	imguiObj.Set("isItemFocused", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.IsItemFocused())
	})
	imguiObj.Set("isItemHoveredV", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 HoveredFlags
		flags := imgui.HoveredFlags(0)
		return vm.ToValue(imgui.IsItemHoveredV(flags))
	})
	imguiObj.Set("isItemToggledOpen", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.IsItemToggledOpen())
	})
	imguiObj.Set("isItemToggledSelection", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.IsItemToggledSelection())
	})
	imguiObj.Set("isItemVisible", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.IsItemVisible())
	})
	imguiObj.Set("internalIsKeyChordPressedInputFlagsV", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 KeyChord
		key_chord := imgui.KeyChord(0)
		// TODO: 处理类型 InputFlags
		flags := imgui.InputFlags(0)
		// TODO: 处理类型 ID
		owner_id := imgui.ID(0)
		return vm.ToValue(imgui.InternalIsKeyChordPressedInputFlagsV(key_chord, flags, owner_id))
	})
	imguiObj.Set("isKeyChordPressed", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 KeyChord
		key_chord := imgui.KeyChord(0)
		return vm.ToValue(imgui.IsKeyChordPressed(key_chord))
	})
	imguiObj.Set("internalIsKeyDownID", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 Key
		key := imgui.Key(0)
		// TODO: 处理类型 ID
		owner_id := imgui.ID(0)
		return vm.ToValue(imgui.InternalIsKeyDownID(key, owner_id))
	})
	imguiObj.Set("isKeyDown", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 Key
		key := imgui.Key(0)
		return vm.ToValue(imgui.IsKeyDown(key))
	})
	imguiObj.Set("isKeyPressedBoolV", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 Key
		key := imgui.Key(0)
		repeat := call.Argument(1).ToBoolean()
		return vm.ToValue(imgui.IsKeyPressedBoolV(key, repeat))
	})
	imguiObj.Set("internalIsKeyPressedInputFlagsV", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 Key
		key := imgui.Key(0)
		// TODO: 处理类型 InputFlags
		flags := imgui.InputFlags(0)
		// TODO: 处理类型 ID
		owner_id := imgui.ID(0)
		return vm.ToValue(imgui.InternalIsKeyPressedInputFlagsV(key, flags, owner_id))
	})
	imguiObj.Set("internalIsKeyReleasedID", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 Key
		key := imgui.Key(0)
		// TODO: 处理类型 ID
		owner_id := imgui.ID(0)
		return vm.ToValue(imgui.InternalIsKeyReleasedID(key, owner_id))
	})
	imguiObj.Set("isKeyReleased", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 Key
		key := imgui.Key(0)
		return vm.ToValue(imgui.IsKeyReleased(key))
	})
	imguiObj.Set("internalIsKeyboardKey", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 Key
		key := imgui.Key(0)
		return vm.ToValue(imgui.InternalIsKeyboardKey(key))
	})
	imguiObj.Set("internalIsLRModKey", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 Key
		key := imgui.Key(0)
		return vm.ToValue(imgui.InternalIsLRModKey(key))
	})
	imguiObj.Set("internalIsLegacyKey", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 Key
		key := imgui.Key(0)
		return vm.ToValue(imgui.InternalIsLegacyKey(key))
	})
	imguiObj.Set("isMouseClickedBoolV", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 MouseButton
		button := imgui.MouseButton(0)
		repeat := call.Argument(1).ToBoolean()
		return vm.ToValue(imgui.IsMouseClickedBoolV(button, repeat))
	})
	imguiObj.Set("internalIsMouseClickedInputFlagsV", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 MouseButton
		button := imgui.MouseButton(0)
		// TODO: 处理类型 InputFlags
		flags := imgui.InputFlags(0)
		// TODO: 处理类型 ID
		owner_id := imgui.ID(0)
		return vm.ToValue(imgui.InternalIsMouseClickedInputFlagsV(button, flags, owner_id))
	})
	imguiObj.Set("internalIsMouseDoubleClickedID", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 MouseButton
		button := imgui.MouseButton(0)
		// TODO: 处理类型 ID
		owner_id := imgui.ID(0)
		return vm.ToValue(imgui.InternalIsMouseDoubleClickedID(button, owner_id))
	})
	imguiObj.Set("isMouseDoubleClicked", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 MouseButton
		button := imgui.MouseButton(0)
		return vm.ToValue(imgui.IsMouseDoubleClicked(button))
	})
	imguiObj.Set("internalIsMouseDownID", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 MouseButton
		button := imgui.MouseButton(0)
		// TODO: 处理类型 ID
		owner_id := imgui.ID(0)
		return vm.ToValue(imgui.InternalIsMouseDownID(button, owner_id))
	})
	imguiObj.Set("isMouseDown", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 MouseButton
		button := imgui.MouseButton(0)
		return vm.ToValue(imgui.IsMouseDown(button))
	})
	imguiObj.Set("internalIsMouseDragPastThresholdV", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 MouseButton
		button := imgui.MouseButton(0)
		lock_threshold := float32(call.Argument(1).ToFloat())
		return vm.ToValue(imgui.InternalIsMouseDragPastThresholdV(button, lock_threshold))
	})
	imguiObj.Set("isMouseDraggingV", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 MouseButton
		button := imgui.MouseButton(0)
		lock_threshold := float32(call.Argument(1).ToFloat())
		return vm.ToValue(imgui.IsMouseDraggingV(button, lock_threshold))
	})
	imguiObj.Set("isMouseHoveringRectV", func(call goja.FunctionCall) goja.Value {
		r_min := parseVec2(vm, call.Argument(0))
		r_max := parseVec2(vm, call.Argument(0))
		clip := call.Argument(0).ToBoolean()
		return vm.ToValue(imgui.IsMouseHoveringRectV(r_min, r_max, clip))
	})
	imguiObj.Set("internalIsMouseKey", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 Key
		key := imgui.Key(0)
		return vm.ToValue(imgui.InternalIsMouseKey(key))
	})
	imguiObj.Set("isMousePosValidV", func(call goja.FunctionCall) goja.Value {
		mouse_pos := parseVec2(vm, call.Argument(0))
		return vm.ToValue(imgui.IsMousePosValidV(&mouse_pos))
	})
	imguiObj.Set("isMouseReleasedWithDelay", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 MouseButton
		button := imgui.MouseButton(0)
		delay := float32(call.Argument(1).ToFloat())
		return vm.ToValue(imgui.IsMouseReleasedWithDelay(button, delay))
	})
	imguiObj.Set("internalIsMouseReleasedID", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 MouseButton
		button := imgui.MouseButton(0)
		// TODO: 处理类型 ID
		owner_id := imgui.ID(0)
		return vm.ToValue(imgui.InternalIsMouseReleasedID(button, owner_id))
	})
	imguiObj.Set("isMouseReleased", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 MouseButton
		button := imgui.MouseButton(0)
		return vm.ToValue(imgui.IsMouseReleased(button))
	})
	imguiObj.Set("internalIsNamedKey", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 Key
		key := imgui.Key(0)
		return vm.ToValue(imgui.InternalIsNamedKey(key))
	})
	imguiObj.Set("internalIsNamedKeyOrMod", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 Key
		key := imgui.Key(0)
		return vm.ToValue(imgui.InternalIsNamedKeyOrMod(key))
	})
	imguiObj.Set("internalIsPopupOpenID", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 ID
		id := imgui.ID(0)
		// TODO: 处理类型 PopupFlags
		popup_flags := imgui.PopupFlags(0)
		return vm.ToValue(imgui.InternalIsPopupOpenID(id, popup_flags))
	})
	imguiObj.Set("isPopupOpenStrV", func(call goja.FunctionCall) goja.Value {
		str_id := call.Argument(0).String()
		// TODO: 处理类型 PopupFlags
		flags := imgui.PopupFlags(0)
		return vm.ToValue(imgui.IsPopupOpenStrV(str_id, flags))
	})
	imguiObj.Set("isRectVisible", func(call goja.FunctionCall) goja.Value {
		size := parseVec2(vm, call.Argument(0))
		return vm.ToValue(imgui.IsRectVisible(size))
	})
	imguiObj.Set("isRectVisibleVec2", func(call goja.FunctionCall) goja.Value {
		rect_min := parseVec2(vm, call.Argument(0))
		rect_max := parseVec2(vm, call.Argument(0))
		return vm.ToValue(imgui.IsRectVisibleVec2(rect_min, rect_max))
	})
	imguiObj.Set("internalIsWindowAbove", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Window
		var potential_above *imgui.Window
		// TODO: 处理指针类型 *Window
		var potential_below *imgui.Window
		return vm.ToValue(imgui.InternalIsWindowAbove(potential_above, potential_below))
	})
	imguiObj.Set("isWindowAppearing", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.IsWindowAppearing())
	})
	imguiObj.Set("internalIsWindowChildOf", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Window
		var window *imgui.Window
		// TODO: 处理指针类型 *Window
		var potential_parent *imgui.Window
		popup_hierarchy := call.Argument(2).ToBoolean()
		dock_hierarchy := call.Argument(3).ToBoolean()
		return vm.ToValue(imgui.InternalIsWindowChildOf(window, potential_parent, popup_hierarchy, dock_hierarchy))
	})
	imguiObj.Set("isWindowCollapsed", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.IsWindowCollapsed())
	})
	imguiObj.Set("internalIsWindowContentHoverableV", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Window
		var window *imgui.Window
		// TODO: 处理类型 HoveredFlags
		flags := imgui.HoveredFlags(0)
		return vm.ToValue(imgui.InternalIsWindowContentHoverableV(window, flags))
	})
	imguiObj.Set("isWindowDocked", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.IsWindowDocked())
	})
	imguiObj.Set("isWindowFocusedV", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 FocusedFlags
		flags := imgui.FocusedFlags(0)
		return vm.ToValue(imgui.IsWindowFocusedV(flags))
	})
	imguiObj.Set("isWindowHoveredV", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 HoveredFlags
		flags := imgui.HoveredFlags(0)
		return vm.ToValue(imgui.IsWindowHoveredV(flags))
	})
	imguiObj.Set("internalIsWindowNavFocusable", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Window
		var window *imgui.Window
		return vm.ToValue(imgui.InternalIsWindowNavFocusable(window))
	})
	imguiObj.Set("internalIsWindowWithinBeginStackOf", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Window
		var window *imgui.Window
		// TODO: 处理指针类型 *Window
		var potential_parent *imgui.Window
		return vm.ToValue(imgui.InternalIsWindowWithinBeginStackOf(window, potential_parent))
	})
	imguiObj.Set("internalItemAddV", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 Rect
		bb := imgui.Rect{}
		// TODO: 处理类型 ID
		id := imgui.ID(0)
		// TODO: 处理指针类型 *Rect
		var nav_bb *imgui.Rect
		// TODO: 处理类型 ItemFlags
		extra_flags := imgui.ItemFlags(0)
		return vm.ToValue(imgui.InternalItemAddV(bb, id, nav_bb, extra_flags))
	})
	imguiObj.Set("internalItemHoverable", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 Rect
		bb := imgui.Rect{}
		// TODO: 处理类型 ID
		id := imgui.ID(0)
		// TODO: 处理类型 ItemFlags
		item_flags := imgui.ItemFlags(0)
		return vm.ToValue(imgui.InternalItemHoverable(bb, id, item_flags))
	})
	imguiObj.Set("internalItemSizeRectV", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 Rect
		bb := imgui.Rect{}
		text_baseline_y := float32(call.Argument(1).ToFloat())
		imgui.InternalItemSizeRectV(bb, text_baseline_y)
		return goja.Undefined()
	})
	imguiObj.Set("internalItemSizeVec2V", func(call goja.FunctionCall) goja.Value {
		size := parseVec2(vm, call.Argument(0))
		text_baseline_y := float32(call.Argument(0).ToFloat())
		imgui.InternalItemSizeVec2V(size, text_baseline_y)
		return goja.Undefined()
	})
	imguiObj.Set("internalKeepAliveID", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 ID
		id := imgui.ID(0)
		imgui.InternalKeepAliveID(id)
		return goja.Undefined()
	})
	imguiObj.Set("labelText", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		fmt := call.Argument(1).String()
		imgui.LabelText(label, fmt)
		return goja.Undefined()
	})
	imguiObj.Set("listBoxStrarrV", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		current_item := int32(0)
		if len(call.Arguments) >= 2 {
			current_item = int32(call.Argument(1).ToInteger())
		}
		items := parseStringArray(vm, call.Argument(2))
		items_count := int32(call.Argument(3).ToInteger())
		height_in_items := int32(call.Argument(4).ToInteger())
		return vm.ToValue(imgui.ListBoxStrarrV(label, &current_item, items, items_count, height_in_items))
	})
	imguiObj.Set("loadIniSettingsFromDisk", func(call goja.FunctionCall) goja.Value {
		ini_filename := call.Argument(0).String()
		imgui.LoadIniSettingsFromDisk(ini_filename)
		return goja.Undefined()
	})
	imguiObj.Set("loadIniSettingsFromMemoryV", func(call goja.FunctionCall) goja.Value {
		ini_data := call.Argument(0).String()
		ini_size := uint64(call.Argument(1).ToInteger())
		imgui.LoadIniSettingsFromMemoryV(ini_data, ini_size)
		return goja.Undefined()
	})
	imguiObj.Set("internalLocalizeGetMsg", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 LocKey
		key := imgui.LocKey(0)
		return vm.ToValue(imgui.InternalLocalizeGetMsg(key))
	})
	imguiObj.Set("internalLocalizeRegisterEntries", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *LocEntry
		var entries *imgui.LocEntry
		count := int32(call.Argument(1).ToInteger())
		imgui.InternalLocalizeRegisterEntries(entries, count)
		return goja.Undefined()
	})
	imguiObj.Set("internalLogBegin", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 LogFlags
		flags := imgui.LogFlags(0)
		auto_open_depth := int32(call.Argument(1).ToInteger())
		imgui.InternalLogBegin(flags, auto_open_depth)
		return goja.Undefined()
	})
	imguiObj.Set("logButtons", func(call goja.FunctionCall) goja.Value {
		imgui.LogButtons()
		return goja.Undefined()
	})
	imguiObj.Set("logFinish", func(call goja.FunctionCall) goja.Value {
		imgui.LogFinish()
		return goja.Undefined()
	})
	imguiObj.Set("internalLogRenderedTextV", func(call goja.FunctionCall) goja.Value {
		ref_pos := parseVec2(vm, call.Argument(0))
		text := call.Argument(1).String()
		imgui.InternalLogRenderedTextV(&ref_pos, text)
		return goja.Undefined()
	})
	imguiObj.Set("internalLogSetNextTextDecoration", func(call goja.FunctionCall) goja.Value {
		prefix := call.Argument(0).String()
		suffix := call.Argument(1).String()
		imgui.InternalLogSetNextTextDecoration(prefix, suffix)
		return goja.Undefined()
	})
	imguiObj.Set("logText", func(call goja.FunctionCall) goja.Value {
		fmt := call.Argument(0).String()
		imgui.LogText(fmt)
		return goja.Undefined()
	})
	imguiObj.Set("internalLogToBufferV", func(call goja.FunctionCall) goja.Value {
		auto_open_depth := int32(call.Argument(0).ToInteger())
		imgui.InternalLogToBufferV(auto_open_depth)
		return goja.Undefined()
	})
	imguiObj.Set("logToClipboardV", func(call goja.FunctionCall) goja.Value {
		auto_open_depth := int32(call.Argument(0).ToInteger())
		imgui.LogToClipboardV(auto_open_depth)
		return goja.Undefined()
	})
	imguiObj.Set("logToFileV", func(call goja.FunctionCall) goja.Value {
		auto_open_depth := int32(call.Argument(0).ToInteger())
		filename := call.Argument(1).String()
		imgui.LogToFileV(auto_open_depth, filename)
		return goja.Undefined()
	})
	imguiObj.Set("logToTTYV", func(call goja.FunctionCall) goja.Value {
		auto_open_depth := int32(call.Argument(0).ToInteger())
		imgui.LogToTTYV(auto_open_depth)
		return goja.Undefined()
	})
	imguiObj.Set("internalMarkIniSettingsDirty", func(call goja.FunctionCall) goja.Value {
		imgui.InternalMarkIniSettingsDirty()
		return goja.Undefined()
	})
	imguiObj.Set("internalMarkIniSettingsDirtyWindowPtr", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Window
		var window *imgui.Window
		imgui.InternalMarkIniSettingsDirtyWindowPtr(window)
		return goja.Undefined()
	})
	imguiObj.Set("internalMarkItemEdited", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 ID
		id := imgui.ID(0)
		imgui.InternalMarkItemEdited(id)
		return goja.Undefined()
	})
	imguiObj.Set("memAlloc", func(call goja.FunctionCall) goja.Value {
		size := uint64(call.Argument(0).ToInteger())
		return vm.ToValue(imgui.MemAlloc(size))
	})
	imguiObj.Set("memFree", func(call goja.FunctionCall) goja.Value {
		ptr := uintptr(call.Argument(0).ToInteger())
		imgui.MemFree(ptr)
		return goja.Undefined()
	})
	imguiObj.Set("internalMenuItemExV", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		icon := call.Argument(1).String()
		shortcut := call.Argument(2).String()
		selected := call.Argument(3).ToBoolean()
		enabled := call.Argument(4).ToBoolean()
		return vm.ToValue(imgui.InternalMenuItemExV(label, icon, shortcut, selected, enabled))
	})
	imguiObj.Set("menuItemBoolV", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		shortcut := call.Argument(1).String()
		selected := call.Argument(2).ToBoolean()
		enabled := call.Argument(3).ToBoolean()
		return vm.ToValue(imgui.MenuItemBoolV(label, shortcut, selected, enabled))
	})
	imguiObj.Set("menuItemBoolPtrV", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		shortcut := call.Argument(1).String()
		p_selected := false
		if len(call.Arguments) >= 3 {
			p_selected = call.Argument(2).ToBoolean()
		}
		enabled := call.Argument(3).ToBoolean()
		return vm.ToValue(imgui.MenuItemBoolPtrV(label, shortcut, &p_selected, enabled))
	})
	imguiObj.Set("internalMouseButtonToKey", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 MouseButton
		button := imgui.MouseButton(0)
		return vm.ToValue(imgui.InternalMouseButtonToKey(button))
	})
	imguiObj.Set("internalMultiSelectAddSetAll", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *MultiSelectTempData
		var ms *imgui.MultiSelectTempData
		selected := call.Argument(1).ToBoolean()
		imgui.InternalMultiSelectAddSetAll(ms, selected)
		return goja.Undefined()
	})
	imguiObj.Set("internalMultiSelectAddSetRange", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *MultiSelectTempData
		var ms *imgui.MultiSelectTempData
		selected := call.Argument(1).ToBoolean()
		range_dir := int32(call.Argument(2).ToInteger())
		// TODO: 处理类型 SelectionUserData
		first_item := imgui.SelectionUserData(0)
		// TODO: 处理类型 SelectionUserData
		last_item := imgui.SelectionUserData(0)
		imgui.InternalMultiSelectAddSetRange(ms, selected, range_dir, first_item, last_item)
		return goja.Undefined()
	})
	imguiObj.Set("internalMultiSelectItemFooter", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 ID
		id := imgui.ID(0)
		p_selected := false
		if len(call.Arguments) >= 2 {
			p_selected = call.Argument(1).ToBoolean()
		}
		p_pressed := false
		if len(call.Arguments) >= 3 {
			p_pressed = call.Argument(2).ToBoolean()
		}
		imgui.InternalMultiSelectItemFooter(id, &p_selected, &p_pressed)
		return goja.Undefined()
	})
	imguiObj.Set("internalMultiSelectItemHeader", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 ID
		id := imgui.ID(0)
		p_selected := false
		if len(call.Arguments) >= 2 {
			p_selected = call.Argument(1).ToBoolean()
		}
		// TODO: 处理指针类型 *ButtonFlags
		var p_button_flags *imgui.ButtonFlags
		imgui.InternalMultiSelectItemHeader(id, &p_selected, p_button_flags)
		return goja.Undefined()
	})
	imguiObj.Set("internalNavClearPreferredPosForAxis", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 Axis
		axis := imgui.Axis(0)
		imgui.InternalNavClearPreferredPosForAxis(axis)
		return goja.Undefined()
	})
	imguiObj.Set("internalNavHighlightActivated", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 ID
		id := imgui.ID(0)
		imgui.InternalNavHighlightActivated(id)
		return goja.Undefined()
	})
	imguiObj.Set("internalNavInitRequestApplyResult", func(call goja.FunctionCall) goja.Value {
		imgui.InternalNavInitRequestApplyResult()
		return goja.Undefined()
	})
	imguiObj.Set("internalNavInitWindow", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Window
		var window *imgui.Window
		force_reinit := call.Argument(1).ToBoolean()
		imgui.InternalNavInitWindow(window, force_reinit)
		return goja.Undefined()
	})
	imguiObj.Set("internalNavMoveRequestApplyResult", func(call goja.FunctionCall) goja.Value {
		imgui.InternalNavMoveRequestApplyResult()
		return goja.Undefined()
	})
	imguiObj.Set("internalNavMoveRequestButNoResultYet", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.InternalNavMoveRequestButNoResultYet())
	})
	imguiObj.Set("internalNavMoveRequestCancel", func(call goja.FunctionCall) goja.Value {
		imgui.InternalNavMoveRequestCancel()
		return goja.Undefined()
	})
	imguiObj.Set("internalNavMoveRequestForward", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 Dir
		move_dir := imgui.Dir(0)
		// TODO: 处理类型 Dir
		clip_dir := imgui.Dir(0)
		// TODO: 处理类型 NavMoveFlags
		move_flags := imgui.NavMoveFlags(0)
		// TODO: 处理类型 ScrollFlags
		scroll_flags := imgui.ScrollFlags(0)
		imgui.InternalNavMoveRequestForward(move_dir, clip_dir, move_flags, scroll_flags)
		return goja.Undefined()
	})
	imguiObj.Set("internalNavMoveRequestResolveWithLastItem", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *NavItemData
		var result *imgui.NavItemData
		imgui.InternalNavMoveRequestResolveWithLastItem(result)
		return goja.Undefined()
	})
	imguiObj.Set("internalNavMoveRequestResolveWithPastTreeNode", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *NavItemData
		var result *imgui.NavItemData
		// TODO: 处理指针类型 *TreeNodeStackData
		var tree_node_data *imgui.TreeNodeStackData
		imgui.InternalNavMoveRequestResolveWithPastTreeNode(result, tree_node_data)
		return goja.Undefined()
	})
	imguiObj.Set("internalNavMoveRequestSubmit", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 Dir
		move_dir := imgui.Dir(0)
		// TODO: 处理类型 Dir
		clip_dir := imgui.Dir(0)
		// TODO: 处理类型 NavMoveFlags
		move_flags := imgui.NavMoveFlags(0)
		// TODO: 处理类型 ScrollFlags
		scroll_flags := imgui.ScrollFlags(0)
		imgui.InternalNavMoveRequestSubmit(move_dir, clip_dir, move_flags, scroll_flags)
		return goja.Undefined()
	})
	imguiObj.Set("internalNavMoveRequestTryWrapping", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Window
		var window *imgui.Window
		// TODO: 处理类型 NavMoveFlags
		move_flags := imgui.NavMoveFlags(0)
		imgui.InternalNavMoveRequestTryWrapping(window, move_flags)
		return goja.Undefined()
	})
	imguiObj.Set("internalNavUpdateCurrentWindowIsScrollPushableX", func(call goja.FunctionCall) goja.Value {
		imgui.InternalNavUpdateCurrentWindowIsScrollPushableX()
		return goja.Undefined()
	})
	imguiObj.Set("newFrame", func(call goja.FunctionCall) goja.Value {
		imgui.NewFrame()
		return goja.Undefined()
	})
	imguiObj.Set("newLine", func(call goja.FunctionCall) goja.Value {
		imgui.NewLine()
		return goja.Undefined()
	})
	imguiObj.Set("nextColumn", func(call goja.FunctionCall) goja.Value {
		imgui.NextColumn()
		return goja.Undefined()
	})
	imguiObj.Set("internalOpenPopupExV", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 ID
		id := imgui.ID(0)
		// TODO: 处理类型 PopupFlags
		popup_flags := imgui.PopupFlags(0)
		imgui.InternalOpenPopupExV(id, popup_flags)
		return goja.Undefined()
	})
	imguiObj.Set("openPopupOnItemClickV", func(call goja.FunctionCall) goja.Value {
		str_id := call.Argument(0).String()
		// TODO: 处理类型 PopupFlags
		popup_flags := imgui.PopupFlags(0)
		imgui.OpenPopupOnItemClickV(str_id, popup_flags)
		return goja.Undefined()
	})
	imguiObj.Set("openPopupIDV", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 ID
		id := imgui.ID(0)
		// TODO: 处理类型 PopupFlags
		popup_flags := imgui.PopupFlags(0)
		imgui.OpenPopupIDV(id, popup_flags)
		return goja.Undefined()
	})
	imguiObj.Set("openPopupStrV", func(call goja.FunctionCall) goja.Value {
		str_id := call.Argument(0).String()
		// TODO: 处理类型 PopupFlags
		popup_flags := imgui.PopupFlags(0)
		imgui.OpenPopupStrV(str_id, popup_flags)
		return goja.Undefined()
	})
	imguiObj.Set("plotHistogramFloatPtrV", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		values := float32(0)
		if len(call.Arguments) >= 2 {
			values = float32(call.Argument(1).ToFloat())
		}
		values_count := int32(call.Argument(2).ToInteger())
		values_offset := int32(call.Argument(3).ToInteger())
		overlay_text := call.Argument(4).String()
		scale_min := float32(call.Argument(5).ToFloat())
		scale_max := float32(call.Argument(6).ToFloat())
		graph_size := parseVec2(vm, call.Argument(7))
		stride := int32(call.Argument(7).ToInteger())
		imgui.PlotHistogramFloatPtrV(label, &values, values_count, values_offset, overlay_text, scale_min, scale_max, graph_size, stride)
		return goja.Undefined()
	})
	imguiObj.Set("plotLinesFloatPtrV", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		values := float32(0)
		if len(call.Arguments) >= 2 {
			values = float32(call.Argument(1).ToFloat())
		}
		values_count := int32(call.Argument(2).ToInteger())
		values_offset := int32(call.Argument(3).ToInteger())
		overlay_text := call.Argument(4).String()
		scale_min := float32(call.Argument(5).ToFloat())
		scale_max := float32(call.Argument(6).ToFloat())
		graph_size := parseVec2(vm, call.Argument(7))
		stride := int32(call.Argument(7).ToInteger())
		imgui.PlotLinesFloatPtrV(label, &values, values_count, values_offset, overlay_text, scale_min, scale_max, graph_size, stride)
		return goja.Undefined()
	})
	imguiObj.Set("popClipRect", func(call goja.FunctionCall) goja.Value {
		imgui.PopClipRect()
		return goja.Undefined()
	})
	imguiObj.Set("internalPopColumnsBackground", func(call goja.FunctionCall) goja.Value {
		imgui.InternalPopColumnsBackground()
		return goja.Undefined()
	})
	imguiObj.Set("internalPopFocusScope", func(call goja.FunctionCall) goja.Value {
		imgui.InternalPopFocusScope()
		return goja.Undefined()
	})
	imguiObj.Set("popFont", func(call goja.FunctionCall) goja.Value {
		imgui.PopFont()
		return goja.Undefined()
	})
	imguiObj.Set("popID", func(call goja.FunctionCall) goja.Value {
		imgui.PopID()
		return goja.Undefined()
	})
	imguiObj.Set("popItemFlag", func(call goja.FunctionCall) goja.Value {
		imgui.PopItemFlag()
		return goja.Undefined()
	})
	imguiObj.Set("popItemWidth", func(call goja.FunctionCall) goja.Value {
		imgui.PopItemWidth()
		return goja.Undefined()
	})
	imguiObj.Set("internalPopPasswordFont", func(call goja.FunctionCall) goja.Value {
		imgui.InternalPopPasswordFont()
		return goja.Undefined()
	})
	imguiObj.Set("popStyleColorV", func(call goja.FunctionCall) goja.Value {
		count := int32(call.Argument(0).ToInteger())
		imgui.PopStyleColorV(count)
		return goja.Undefined()
	})
	imguiObj.Set("popStyleVarV", func(call goja.FunctionCall) goja.Value {
		count := int32(call.Argument(0).ToInteger())
		imgui.PopStyleVarV(count)
		return goja.Undefined()
	})
	imguiObj.Set("popTextWrapPos", func(call goja.FunctionCall) goja.Value {
		imgui.PopTextWrapPos()
		return goja.Undefined()
	})
	imguiObj.Set("progressBarV", func(call goja.FunctionCall) goja.Value {
		fraction := float32(call.Argument(0).ToFloat())
		size_arg := parseVec2(vm, call.Argument(1))
		overlay := call.Argument(1).String()
		imgui.ProgressBarV(fraction, size_arg, overlay)
		return goja.Undefined()
	})
	imguiObj.Set("pushClipRect", func(call goja.FunctionCall) goja.Value {
		clip_rect_min := parseVec2(vm, call.Argument(0))
		clip_rect_max := parseVec2(vm, call.Argument(0))
		intersect_with_current_clip_rect := call.Argument(0).ToBoolean()
		imgui.PushClipRect(clip_rect_min, clip_rect_max, intersect_with_current_clip_rect)
		return goja.Undefined()
	})
	imguiObj.Set("internalPushColumnClipRect", func(call goja.FunctionCall) goja.Value {
		column_index := int32(call.Argument(0).ToInteger())
		imgui.InternalPushColumnClipRect(column_index)
		return goja.Undefined()
	})
	imguiObj.Set("internalPushColumnsBackground", func(call goja.FunctionCall) goja.Value {
		imgui.InternalPushColumnsBackground()
		return goja.Undefined()
	})
	imguiObj.Set("internalPushFocusScope", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 ID
		id := imgui.ID(0)
		imgui.InternalPushFocusScope(id)
		return goja.Undefined()
	})
	imguiObj.Set("pushFont", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Font
		var font *imgui.Font
		font_size_base_unscaled := float32(call.Argument(1).ToFloat())
		imgui.PushFont(font, font_size_base_unscaled)
		return goja.Undefined()
	})
	imguiObj.Set("pushIDInt", func(call goja.FunctionCall) goja.Value {
		int_id := int32(call.Argument(0).ToInteger())
		imgui.PushIDInt(int_id)
		return goja.Undefined()
	})
	imguiObj.Set("pushIDPtr", func(call goja.FunctionCall) goja.Value {
		ptr_id := uintptr(call.Argument(0).ToInteger())
		imgui.PushIDPtr(ptr_id)
		return goja.Undefined()
	})
	imguiObj.Set("pushIDStr", func(call goja.FunctionCall) goja.Value {
		str_id := call.Argument(0).String()
		imgui.PushIDStr(str_id)
		return goja.Undefined()
	})
	imguiObj.Set("pushIDStrStr", func(call goja.FunctionCall) goja.Value {
		str_id_begin := call.Argument(0).String()
		str_id_end := call.Argument(1).String()
		imgui.PushIDStrStr(str_id_begin, str_id_end)
		return goja.Undefined()
	})
	imguiObj.Set("pushItemFlag", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 ItemFlags
		option := imgui.ItemFlags(0)
		enabled := call.Argument(1).ToBoolean()
		imgui.PushItemFlag(option, enabled)
		return goja.Undefined()
	})
	imguiObj.Set("pushItemWidth", func(call goja.FunctionCall) goja.Value {
		item_width := float32(call.Argument(0).ToFloat())
		imgui.PushItemWidth(item_width)
		return goja.Undefined()
	})
	imguiObj.Set("internalPushMultiItemsWidths", func(call goja.FunctionCall) goja.Value {
		components := int32(call.Argument(0).ToInteger())
		width_full := float32(call.Argument(1).ToFloat())
		imgui.InternalPushMultiItemsWidths(components, width_full)
		return goja.Undefined()
	})
	imguiObj.Set("internalPushOverrideID", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 ID
		id := imgui.ID(0)
		imgui.InternalPushOverrideID(id)
		return goja.Undefined()
	})
	imguiObj.Set("internalPushPasswordFont", func(call goja.FunctionCall) goja.Value {
		imgui.InternalPushPasswordFont()
		return goja.Undefined()
	})
	imguiObj.Set("pushStyleColorU32", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 Col
		idx := imgui.Col(0)
		col := uint32(call.Argument(1).ToInteger())
		imgui.PushStyleColorU32(idx, col)
		return goja.Undefined()
	})
	imguiObj.Set("pushStyleColorVec4", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 Col
		idx := imgui.Col(0)
		col := parseVec4(vm, call.Argument(1))
		imgui.PushStyleColorVec4(idx, col)
		return goja.Undefined()
	})
	imguiObj.Set("pushStyleVarX", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 StyleVar
		idx := imgui.StyleVar(0)
		val_x := float32(call.Argument(1).ToFloat())
		imgui.PushStyleVarX(idx, val_x)
		return goja.Undefined()
	})
	imguiObj.Set("pushStyleVarY", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 StyleVar
		idx := imgui.StyleVar(0)
		val_y := float32(call.Argument(1).ToFloat())
		imgui.PushStyleVarY(idx, val_y)
		return goja.Undefined()
	})
	imguiObj.Set("pushStyleVarFloat", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 StyleVar
		idx := imgui.StyleVar(0)
		val := float32(call.Argument(1).ToFloat())
		imgui.PushStyleVarFloat(idx, val)
		return goja.Undefined()
	})
	imguiObj.Set("pushStyleVarVec2", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 StyleVar
		idx := imgui.StyleVar(0)
		val := parseVec2(vm, call.Argument(1))
		imgui.PushStyleVarVec2(idx, val)
		return goja.Undefined()
	})
	imguiObj.Set("pushTextWrapPosV", func(call goja.FunctionCall) goja.Value {
		wrap_local_pos_x := float32(call.Argument(0).ToFloat())
		imgui.PushTextWrapPosV(wrap_local_pos_x)
		return goja.Undefined()
	})
	imguiObj.Set("radioButtonBool", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		active := call.Argument(1).ToBoolean()
		return vm.ToValue(imgui.RadioButtonBool(label, active))
	})
	imguiObj.Set("radioButtonIntPtr", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		v := int32(0)
		if len(call.Arguments) >= 2 {
			v = int32(call.Argument(1).ToInteger())
		}
		v_button := int32(call.Argument(2).ToInteger())
		return vm.ToValue(imgui.RadioButtonIntPtr(label, &v, v_button))
	})
	imguiObj.Set("internalRegisterFontAtlas", func(call goja.FunctionCall) goja.Value {
		var atlas *imgui.FontAtlas
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.FontAtlas); ok {
				atlas = v
			}
		}
		imgui.InternalRegisterFontAtlas(atlas)
		return goja.Undefined()
	})
	imguiObj.Set("internalRegisterUserTexture", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *TextureData
		var tex *imgui.TextureData
		imgui.InternalRegisterUserTexture(tex)
		return goja.Undefined()
	})
	imguiObj.Set("internalRemoveContextHook", func(call goja.FunctionCall) goja.Value {
		var context *imgui.Context
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.Context); ok {
				context = v
			}
		}
		// TODO: 处理类型 ID
		hook_to_remove := imgui.ID(0)
		imgui.InternalRemoveContextHook(context, hook_to_remove)
		return goja.Undefined()
	})
	imguiObj.Set("internalRemoveSettingsHandler", func(call goja.FunctionCall) goja.Value {
		type_name := call.Argument(0).String()
		imgui.InternalRemoveSettingsHandler(type_name)
		return goja.Undefined()
	})
	imguiObj.Set("render", func(call goja.FunctionCall) goja.Value {
		imgui.Render()
		return goja.Undefined()
	})
	imguiObj.Set("internalRenderArrowV", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *DrawList
		var draw_list *imgui.DrawList
		pos := parseVec2(vm, call.Argument(1))
		col := uint32(call.Argument(1).ToInteger())
		// TODO: 处理类型 Dir
		dir := imgui.Dir(0)
		scale := float32(call.Argument(3).ToFloat())
		imgui.InternalRenderArrowV(draw_list, pos, col, dir, scale)
		return goja.Undefined()
	})
	imguiObj.Set("internalRenderArrowDockMenu", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *DrawList
		var draw_list *imgui.DrawList
		p_min := parseVec2(vm, call.Argument(1))
		sz := float32(call.Argument(1).ToFloat())
		col := uint32(call.Argument(2).ToInteger())
		imgui.InternalRenderArrowDockMenu(draw_list, p_min, sz, col)
		return goja.Undefined()
	})
	imguiObj.Set("internalRenderArrowPointingAt", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *DrawList
		var draw_list *imgui.DrawList
		pos := parseVec2(vm, call.Argument(1))
		half_sz := parseVec2(vm, call.Argument(1))
		// TODO: 处理类型 Dir
		direction := imgui.Dir(0)
		col := uint32(call.Argument(2).ToInteger())
		imgui.InternalRenderArrowPointingAt(draw_list, pos, half_sz, direction, col)
		return goja.Undefined()
	})
	imguiObj.Set("internalRenderBullet", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *DrawList
		var draw_list *imgui.DrawList
		pos := parseVec2(vm, call.Argument(1))
		col := uint32(call.Argument(1).ToInteger())
		imgui.InternalRenderBullet(draw_list, pos, col)
		return goja.Undefined()
	})
	imguiObj.Set("internalRenderCheckMark", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *DrawList
		var draw_list *imgui.DrawList
		pos := parseVec2(vm, call.Argument(1))
		col := uint32(call.Argument(1).ToInteger())
		sz := float32(call.Argument(2).ToFloat())
		imgui.InternalRenderCheckMark(draw_list, pos, col, sz)
		return goja.Undefined()
	})
	imguiObj.Set("internalRenderColorRectWithAlphaCheckerboardV", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *DrawList
		var draw_list *imgui.DrawList
		p_min := parseVec2(vm, call.Argument(1))
		p_max := parseVec2(vm, call.Argument(1))
		fill_col := uint32(call.Argument(1).ToInteger())
		grid_step := float32(call.Argument(2).ToFloat())
		grid_off := parseVec2(vm, call.Argument(3))
		rounding := float32(call.Argument(3).ToFloat())
		// TODO: 处理类型 DrawFlags
		flags := imgui.DrawFlags(0)
		imgui.InternalRenderColorRectWithAlphaCheckerboardV(draw_list, p_min, p_max, fill_col, grid_step, grid_off, rounding, flags)
		return goja.Undefined()
	})
	imguiObj.Set("internalRenderDragDropTargetRectEx", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *DrawList
		var draw_list *imgui.DrawList
		// TODO: 处理类型 Rect
		bb := imgui.Rect{}
		imgui.InternalRenderDragDropTargetRectEx(draw_list, bb)
		return goja.Undefined()
	})
	imguiObj.Set("internalRenderDragDropTargetRectForItem", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 Rect
		bb := imgui.Rect{}
		imgui.InternalRenderDragDropTargetRectForItem(bb)
		return goja.Undefined()
	})
	imguiObj.Set("internalRenderFrameV", func(call goja.FunctionCall) goja.Value {
		p_min := parseVec2(vm, call.Argument(0))
		p_max := parseVec2(vm, call.Argument(0))
		fill_col := uint32(call.Argument(0).ToInteger())
		borders := call.Argument(1).ToBoolean()
		rounding := float32(call.Argument(2).ToFloat())
		imgui.InternalRenderFrameV(p_min, p_max, fill_col, borders, rounding)
		return goja.Undefined()
	})
	imguiObj.Set("internalRenderFrameBorderV", func(call goja.FunctionCall) goja.Value {
		p_min := parseVec2(vm, call.Argument(0))
		p_max := parseVec2(vm, call.Argument(0))
		rounding := float32(call.Argument(0).ToFloat())
		imgui.InternalRenderFrameBorderV(p_min, p_max, rounding)
		return goja.Undefined()
	})
	imguiObj.Set("internalRenderMouseCursor", func(call goja.FunctionCall) goja.Value {
		pos := parseVec2(vm, call.Argument(0))
		scale := float32(call.Argument(0).ToFloat())
		// TODO: 处理类型 MouseCursor
		mouse_cursor := imgui.MouseCursor(0)
		col_fill := uint32(call.Argument(2).ToInteger())
		col_border := uint32(call.Argument(3).ToInteger())
		col_shadow := uint32(call.Argument(4).ToInteger())
		imgui.InternalRenderMouseCursor(pos, scale, mouse_cursor, col_fill, col_border, col_shadow)
		return goja.Undefined()
	})
	imguiObj.Set("internalRenderNavCursorV", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 Rect
		bb := imgui.Rect{}
		// TODO: 处理类型 ID
		id := imgui.ID(0)
		// TODO: 处理类型 NavRenderCursorFlags
		flags := imgui.NavRenderCursorFlags(0)
		imgui.InternalRenderNavCursorV(bb, id, flags)
		return goja.Undefined()
	})
	imguiObj.Set("renderPlatformWindowsDefaultV", func(call goja.FunctionCall) goja.Value {
		platform_render_arg := uintptr(call.Argument(0).ToInteger())
		renderer_render_arg := uintptr(call.Argument(1).ToInteger())
		imgui.RenderPlatformWindowsDefaultV(platform_render_arg, renderer_render_arg)
		return goja.Undefined()
	})
	imguiObj.Set("internalRenderRectFilledRangeH", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *DrawList
		var draw_list *imgui.DrawList
		// TODO: 处理类型 Rect
		rect := imgui.Rect{}
		col := uint32(call.Argument(2).ToInteger())
		x_start_norm := float32(call.Argument(3).ToFloat())
		x_end_norm := float32(call.Argument(4).ToFloat())
		rounding := float32(call.Argument(5).ToFloat())
		imgui.InternalRenderRectFilledRangeH(draw_list, rect, col, x_start_norm, x_end_norm, rounding)
		return goja.Undefined()
	})
	imguiObj.Set("internalRenderRectFilledWithHole", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *DrawList
		var draw_list *imgui.DrawList
		// TODO: 处理类型 Rect
		outer := imgui.Rect{}
		// TODO: 处理类型 Rect
		inner := imgui.Rect{}
		col := uint32(call.Argument(3).ToInteger())
		rounding := float32(call.Argument(4).ToFloat())
		imgui.InternalRenderRectFilledWithHole(draw_list, outer, inner, col, rounding)
		return goja.Undefined()
	})
	imguiObj.Set("internalRenderTextV", func(call goja.FunctionCall) goja.Value {
		pos := parseVec2(vm, call.Argument(0))
		text := call.Argument(0).String()
		hide_text_after_hash := call.Argument(1).ToBoolean()
		imgui.InternalRenderTextV(pos, text, hide_text_after_hash)
		return goja.Undefined()
	})
	imguiObj.Set("internalRenderTextClippedV", func(call goja.FunctionCall) goja.Value {
		pos_min := parseVec2(vm, call.Argument(0))
		pos_max := parseVec2(vm, call.Argument(0))
		text := call.Argument(0).String()
		text_size_if_known := parseVec2(vm, call.Argument(1))
		align := parseVec2(vm, call.Argument(2))
		// TODO: 处理指针类型 *Rect
		var clip_rect *imgui.Rect
		imgui.InternalRenderTextClippedV(pos_min, pos_max, text, &text_size_if_known, align, clip_rect)
		return goja.Undefined()
	})
	imguiObj.Set("internalRenderTextClippedExV", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *DrawList
		var draw_list *imgui.DrawList
		pos_min := parseVec2(vm, call.Argument(1))
		pos_max := parseVec2(vm, call.Argument(1))
		text := call.Argument(1).String()
		text_size_if_known := parseVec2(vm, call.Argument(2))
		align := parseVec2(vm, call.Argument(3))
		// TODO: 处理指针类型 *Rect
		var clip_rect *imgui.Rect
		imgui.InternalRenderTextClippedExV(draw_list, pos_min, pos_max, text, &text_size_if_known, align, clip_rect)
		return goja.Undefined()
	})
	imguiObj.Set("internalRenderTextEllipsis", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *DrawList
		var draw_list *imgui.DrawList
		pos_min := parseVec2(vm, call.Argument(1))
		pos_max := parseVec2(vm, call.Argument(1))
		ellipsis_max_x := float32(call.Argument(1).ToFloat())
		text := call.Argument(2).String()
		text_size_if_known := parseVec2(vm, call.Argument(3))
		imgui.InternalRenderTextEllipsis(draw_list, pos_min, pos_max, ellipsis_max_x, text, &text_size_if_known)
		return goja.Undefined()
	})
	imguiObj.Set("internalRenderTextWrapped", func(call goja.FunctionCall) goja.Value {
		pos := parseVec2(vm, call.Argument(0))
		text := call.Argument(0).String()
		wrap_width := float32(call.Argument(1).ToFloat())
		imgui.InternalRenderTextWrapped(pos, text, wrap_width)
		return goja.Undefined()
	})
	imguiObj.Set("resetMouseDragDeltaV", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 MouseButton
		button := imgui.MouseButton(0)
		imgui.ResetMouseDragDeltaV(button)
		return goja.Undefined()
	})
	imguiObj.Set("sameLineV", func(call goja.FunctionCall) goja.Value {
		offset_from_start_x := float32(call.Argument(0).ToFloat())
		spacing := float32(call.Argument(1).ToFloat())
		imgui.SameLineV(offset_from_start_x, spacing)
		return goja.Undefined()
	})
	imguiObj.Set("saveIniSettingsToDisk", func(call goja.FunctionCall) goja.Value {
		ini_filename := call.Argument(0).String()
		imgui.SaveIniSettingsToDisk(ini_filename)
		return goja.Undefined()
	})
	imguiObj.Set("saveIniSettingsToMemoryV", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *uint64
		var out_ini_size *uint64
		return vm.ToValue(imgui.SaveIniSettingsToMemoryV(out_ini_size))
	})
	imguiObj.Set("internalScaleWindowsInViewport", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *ViewportP
		var viewport *imgui.ViewportP
		scale := float32(call.Argument(1).ToFloat())
		imgui.InternalScaleWindowsInViewport(viewport, scale)
		return goja.Undefined()
	})
	imguiObj.Set("internalScrollToBringRectIntoView", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Window
		var window *imgui.Window
		// TODO: 处理类型 Rect
		rect := imgui.Rect{}
		imgui.InternalScrollToBringRectIntoView(window, rect)
		return goja.Undefined()
	})
	imguiObj.Set("internalScrollToItemV", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 ScrollFlags
		flags := imgui.ScrollFlags(0)
		imgui.InternalScrollToItemV(flags)
		return goja.Undefined()
	})
	imguiObj.Set("internalScrollToRectV", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Window
		var window *imgui.Window
		// TODO: 处理类型 Rect
		rect := imgui.Rect{}
		// TODO: 处理类型 ScrollFlags
		flags := imgui.ScrollFlags(0)
		imgui.InternalScrollToRectV(window, rect, flags)
		return goja.Undefined()
	})
	imguiObj.Set("internalScrollToRectExV", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Window
		var window *imgui.Window
		// TODO: 处理类型 Rect
		rect := imgui.Rect{}
		// TODO: 处理类型 ScrollFlags
		flags := imgui.ScrollFlags(0)
		result := imgui.InternalScrollToRectExV(window, rect, flags)
		return vm.ToValue(result)
	})
	imguiObj.Set("internalScrollbar", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 Axis
		axis := imgui.Axis(0)
		imgui.InternalScrollbar(axis)
		return goja.Undefined()
	})
	imguiObj.Set("internalScrollbarExV", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 Rect
		bb := imgui.Rect{}
		// TODO: 处理类型 ID
		id := imgui.ID(0)
		// TODO: 处理类型 Axis
		axis := imgui.Axis(0)
		// TODO: 处理指针类型 *int64
		var p_scroll_v *int64
		// TODO: 处理类型 int64
		avail_v := int64(0)
		// TODO: 处理类型 int64
		contents_v := int64(0)
		// TODO: 处理类型 DrawFlags
		draw_rounding_flags := imgui.DrawFlags(0)
		return vm.ToValue(imgui.InternalScrollbarExV(bb, id, axis, p_scroll_v, avail_v, contents_v, draw_rounding_flags))
	})
	imguiObj.Set("selectableBoolV", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		selected := call.Argument(1).ToBoolean()
		// TODO: 处理类型 SelectableFlags
		flags := imgui.SelectableFlags(0)
		size := parseVec2(vm, call.Argument(3))
		return vm.ToValue(imgui.SelectableBoolV(label, selected, flags, size))
	})
	imguiObj.Set("selectableBoolPtrV", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		p_selected := false
		if len(call.Arguments) >= 2 {
			p_selected = call.Argument(1).ToBoolean()
		}
		// TODO: 处理类型 SelectableFlags
		flags := imgui.SelectableFlags(0)
		size := parseVec2(vm, call.Argument(3))
		return vm.ToValue(imgui.SelectableBoolPtrV(label, &p_selected, flags, size))
	})
	imguiObj.Set("separator", func(call goja.FunctionCall) goja.Value {
		imgui.Separator()
		return goja.Undefined()
	})
	imguiObj.Set("internalSeparatorExV", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 SeparatorFlags
		flags := imgui.SeparatorFlags(0)
		thickness := float32(call.Argument(1).ToFloat())
		imgui.InternalSeparatorExV(flags, thickness)
		return goja.Undefined()
	})
	imguiObj.Set("separatorText", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		imgui.SeparatorText(label)
		return goja.Undefined()
	})
	imguiObj.Set("drawRect", func(call goja.FunctionCall) goja.Value {
		x1 := call.Argument(0).ToInteger()
		y1 := call.Argument(1).ToInteger()
		x2 := call.Argument(2).ToInteger()
		y2 := call.Argument(3).ToInteger()
		color := uint32(call.Argument(4).ToInteger())
		thickness := float32(call.Argument(5).ToFloat())
		imgui.DrawRect(int(x1), int(y1), int(x2), int(y2), color, thickness)
		return goja.Undefined()
	})
	imguiObj.Set("internalSeparatorTextEx", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 ID
		id := imgui.ID(0)
		label := call.Argument(1).String()
		label_end := call.Argument(2).String()
		extra_width := float32(call.Argument(3).ToFloat())
		imgui.InternalSeparatorTextEx(id, label, label_end, extra_width)
		return goja.Undefined()
	})
	imguiObj.Set("internalSetActiveID", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 ID
		id := imgui.ID(0)
		// TODO: 处理指针类型 *Window
		var window *imgui.Window
		imgui.InternalSetActiveID(id, window)
		return goja.Undefined()
	})
	imguiObj.Set("internalSetActiveIdUsingAllKeyboardKeys", func(call goja.FunctionCall) goja.Value {
		imgui.InternalSetActiveIdUsingAllKeyboardKeys()
		return goja.Undefined()
	})
	imguiObj.Set("setAllocatorFunctionsV", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理函数类型 MemAllocFunc
		var alloc_func imgui.MemAllocFunc
		// TODO: 处理函数类型 MemFreeFunc
		var free_func imgui.MemFreeFunc
		user_data := uintptr(call.Argument(2).ToInteger())
		imgui.SetAllocatorFunctionsV(alloc_func, free_func, user_data)
		return goja.Undefined()
	})
	imguiObj.Set("setClipboardText", func(call goja.FunctionCall) goja.Value {
		text := call.Argument(0).String()
		imgui.SetClipboardText(text)
		return goja.Undefined()
	})
	imguiObj.Set("setColorEditOptions", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 ColorEditFlags
		flags := imgui.ColorEditFlags(0)
		imgui.SetColorEditOptions(flags)
		return goja.Undefined()
	})
	imguiObj.Set("setColumnOffset", func(call goja.FunctionCall) goja.Value {
		column_index := int32(call.Argument(0).ToInteger())
		offset_x := float32(call.Argument(1).ToFloat())
		imgui.SetColumnOffset(column_index, offset_x)
		return goja.Undefined()
	})
	imguiObj.Set("setColumnWidth", func(call goja.FunctionCall) goja.Value {
		column_index := int32(call.Argument(0).ToInteger())
		width := float32(call.Argument(1).ToFloat())
		imgui.SetColumnWidth(column_index, width)
		return goja.Undefined()
	})
	imguiObj.Set("setCurrentContext", func(call goja.FunctionCall) goja.Value {
		var ctx *imgui.Context
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.Context); ok {
				ctx = v
			}
		}
		imgui.SetCurrentContext(ctx)
		return goja.Undefined()
	})
	imguiObj.Set("internalSetCurrentFont", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Font
		var font *imgui.Font
		font_size_before_scaling := float32(call.Argument(1).ToFloat())
		font_size_after_scaling := float32(call.Argument(2).ToFloat())
		imgui.InternalSetCurrentFont(font, font_size_before_scaling, font_size_after_scaling)
		return goja.Undefined()
	})
	imguiObj.Set("internalSetCurrentViewport", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Window
		var window *imgui.Window
		// TODO: 处理指针类型 *ViewportP
		var viewport *imgui.ViewportP
		imgui.InternalSetCurrentViewport(window, viewport)
		return goja.Undefined()
	})
	imguiObj.Set("setCursorPos", func(call goja.FunctionCall) goja.Value {
		local_pos := parseVec2(vm, call.Argument(0))
		imgui.SetCursorPos(local_pos)
		return goja.Undefined()
	})
	imguiObj.Set("setCursorPosX", func(call goja.FunctionCall) goja.Value {
		local_x := float32(call.Argument(0).ToFloat())
		imgui.SetCursorPosX(local_x)
		return goja.Undefined()
	})
	imguiObj.Set("setCursorPosY", func(call goja.FunctionCall) goja.Value {
		local_y := float32(call.Argument(0).ToFloat())
		imgui.SetCursorPosY(local_y)
		return goja.Undefined()
	})
	imguiObj.Set("setCursorScreenPos", func(call goja.FunctionCall) goja.Value {
		pos := parseVec2(vm, call.Argument(0))
		imgui.SetCursorScreenPos(pos)
		return goja.Undefined()
	})
	imguiObj.Set("setDragDropPayloadV", func(call goja.FunctionCall) goja.Value {
		typeArg := call.Argument(0).String()
		data := uintptr(call.Argument(1).ToInteger())
		sz := uint64(call.Argument(2).ToInteger())
		// TODO: 处理类型 Cond
		cond := imgui.Cond(0)
		return vm.ToValue(imgui.SetDragDropPayloadV(typeArg, data, sz, cond))
	})
	imguiObj.Set("internalSetFocusID", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 ID
		id := imgui.ID(0)
		// TODO: 处理指针类型 *Window
		var window *imgui.Window
		imgui.InternalSetFocusID(id, window)
		return goja.Undefined()
	})
	imguiObj.Set("internalSetFontRasterizerDensity", func(call goja.FunctionCall) goja.Value {
		rasterizer_density := float32(call.Argument(0).ToFloat())
		imgui.InternalSetFontRasterizerDensity(rasterizer_density)
		return goja.Undefined()
	})
	imguiObj.Set("internalSetHoveredID", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 ID
		id := imgui.ID(0)
		imgui.InternalSetHoveredID(id)
		return goja.Undefined()
	})
	imguiObj.Set("setItemDefaultFocus", func(call goja.FunctionCall) goja.Value {
		imgui.SetItemDefaultFocus()
		return goja.Undefined()
	})
	imguiObj.Set("internalSetItemKeyOwnerInputFlags", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 Key
		key := imgui.Key(0)
		// TODO: 处理类型 InputFlags
		flags := imgui.InputFlags(0)
		imgui.InternalSetItemKeyOwnerInputFlags(key, flags)
		return goja.Undefined()
	})
	imguiObj.Set("setItemKeyOwner", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 Key
		key := imgui.Key(0)
		imgui.SetItemKeyOwner(key)
		return goja.Undefined()
	})
	imguiObj.Set("setItemTooltip", func(call goja.FunctionCall) goja.Value {
		fmt := call.Argument(0).String()
		imgui.SetItemTooltip(fmt)
		return goja.Undefined()
	})
	imguiObj.Set("internalSetKeyOwnerV", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 Key
		key := imgui.Key(0)
		// TODO: 处理类型 ID
		owner_id := imgui.ID(0)
		// TODO: 处理类型 InputFlags
		flags := imgui.InputFlags(0)
		imgui.InternalSetKeyOwnerV(key, owner_id, flags)
		return goja.Undefined()
	})
	imguiObj.Set("internalSetKeyOwnersForKeyChordV", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 KeyChord
		key := imgui.KeyChord(0)
		// TODO: 处理类型 ID
		owner_id := imgui.ID(0)
		// TODO: 处理类型 InputFlags
		flags := imgui.InputFlags(0)
		imgui.InternalSetKeyOwnersForKeyChordV(key, owner_id, flags)
		return goja.Undefined()
	})
	imguiObj.Set("setKeyboardFocusHereV", func(call goja.FunctionCall) goja.Value {
		offset := int32(call.Argument(0).ToInteger())
		imgui.SetKeyboardFocusHereV(offset)
		return goja.Undefined()
	})
	imguiObj.Set("internalSetLastItemData", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 ID
		item_id := imgui.ID(0)
		// TODO: 处理类型 ItemFlags
		item_flags := imgui.ItemFlags(0)
		// TODO: 处理类型 ItemStatusFlags
		status_flags := imgui.ItemStatusFlags(0)
		// TODO: 处理类型 Rect
		item_rect := imgui.Rect{}
		imgui.InternalSetLastItemData(item_id, item_flags, status_flags, item_rect)
		return goja.Undefined()
	})
	imguiObj.Set("setMouseCursor", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 MouseCursor
		cursor_type := imgui.MouseCursor(0)
		imgui.SetMouseCursor(cursor_type)
		return goja.Undefined()
	})
	imguiObj.Set("setNavCursorVisible", func(call goja.FunctionCall) goja.Value {
		visible := call.Argument(0).ToBoolean()
		imgui.SetNavCursorVisible(visible)
		return goja.Undefined()
	})
	imguiObj.Set("internalSetNavCursorVisibleAfterMove", func(call goja.FunctionCall) goja.Value {
		imgui.InternalSetNavCursorVisibleAfterMove()
		return goja.Undefined()
	})
	imguiObj.Set("internalSetNavFocusScope", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 ID
		focus_scope_id := imgui.ID(0)
		imgui.InternalSetNavFocusScope(focus_scope_id)
		return goja.Undefined()
	})
	imguiObj.Set("internalSetNavID", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 ID
		id := imgui.ID(0)
		// TODO: 处理类型 NavLayer
		nav_layer := imgui.NavLayer(0)
		// TODO: 处理类型 ID
		focus_scope_id := imgui.ID(0)
		// TODO: 处理类型 Rect
		rect_rel := imgui.Rect{}
		imgui.InternalSetNavID(id, nav_layer, focus_scope_id, rect_rel)
		return goja.Undefined()
	})
	imguiObj.Set("internalSetNavWindow", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Window
		var window *imgui.Window
		imgui.InternalSetNavWindow(window)
		return goja.Undefined()
	})
	imguiObj.Set("setNextFrameWantCaptureKeyboard", func(call goja.FunctionCall) goja.Value {
		want_capture_keyboard := call.Argument(0).ToBoolean()
		imgui.SetNextFrameWantCaptureKeyboard(want_capture_keyboard)
		return goja.Undefined()
	})
	imguiObj.Set("setNextFrameWantCaptureMouse", func(call goja.FunctionCall) goja.Value {
		want_capture_mouse := call.Argument(0).ToBoolean()
		imgui.SetNextFrameWantCaptureMouse(want_capture_mouse)
		return goja.Undefined()
	})
	imguiObj.Set("setNextItemAllowOverlap", func(call goja.FunctionCall) goja.Value {
		imgui.SetNextItemAllowOverlap()
		return goja.Undefined()
	})
	imguiObj.Set("setNextItemOpenV", func(call goja.FunctionCall) goja.Value {
		is_open := call.Argument(0).ToBoolean()
		// TODO: 处理类型 Cond
		cond := imgui.Cond(0)
		imgui.SetNextItemOpenV(is_open, cond)
		return goja.Undefined()
	})
	imguiObj.Set("internalSetNextItemRefVal", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 DataType
		data_type := imgui.DataType(0)
		p_data := uintptr(call.Argument(1).ToInteger())
		imgui.InternalSetNextItemRefVal(data_type, p_data)
		return goja.Undefined()
	})
	imguiObj.Set("setNextItemSelectionUserData", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 SelectionUserData
		selection_user_data := imgui.SelectionUserData(0)
		imgui.SetNextItemSelectionUserData(selection_user_data)
		return goja.Undefined()
	})
	imguiObj.Set("setNextItemShortcutV", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 KeyChord
		key_chord := imgui.KeyChord(0)
		// TODO: 处理类型 InputFlags
		flags := imgui.InputFlags(0)
		imgui.SetNextItemShortcutV(key_chord, flags)
		return goja.Undefined()
	})
	imguiObj.Set("setNextItemStorageID", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 ID
		storage_id := imgui.ID(0)
		imgui.SetNextItemStorageID(storage_id)
		return goja.Undefined()
	})
	imguiObj.Set("setNextItemWidth", func(call goja.FunctionCall) goja.Value {
		item_width := float32(call.Argument(0).ToFloat())
		imgui.SetNextItemWidth(item_width)
		return goja.Undefined()
	})
	imguiObj.Set("setNextWindowBgAlpha", func(call goja.FunctionCall) goja.Value {
		alpha := float32(call.Argument(0).ToFloat())
		imgui.SetNextWindowBgAlpha(alpha)
		return goja.Undefined()
	})
	imguiObj.Set("setNextWindowClass", func(call goja.FunctionCall) goja.Value {
		var window_class *imgui.WindowClass
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.WindowClass); ok {
				window_class = v
			}
		}
		imgui.SetNextWindowClass(window_class)
		return goja.Undefined()
	})
	imguiObj.Set("setNextWindowCollapsedV", func(call goja.FunctionCall) goja.Value {
		collapsed := call.Argument(0).ToBoolean()
		// TODO: 处理类型 Cond
		cond := imgui.Cond(0)
		imgui.SetNextWindowCollapsedV(collapsed, cond)
		return goja.Undefined()
	})
	imguiObj.Set("setNextWindowContentSize", func(call goja.FunctionCall) goja.Value {
		size := parseVec2(vm, call.Argument(0))
		imgui.SetNextWindowContentSize(size)
		return goja.Undefined()
	})
	imguiObj.Set("setNextWindowDockIDV", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 ID
		dock_id := imgui.ID(0)
		// TODO: 处理类型 Cond
		cond := imgui.Cond(0)
		imgui.SetNextWindowDockIDV(dock_id, cond)
		return goja.Undefined()
	})
	imguiObj.Set("setNextWindowFocus", func(call goja.FunctionCall) goja.Value {
		imgui.SetNextWindowFocus()
		return goja.Undefined()
	})
	imguiObj.Set("setNextWindowPosV", func(call goja.FunctionCall) goja.Value {
		pos := parseVec2(vm, call.Argument(0))
		// TODO: 处理类型 Cond
		cond := imgui.Cond(0)
		pivot := parseVec2(vm, call.Argument(1))
		imgui.SetNextWindowPosV(pos, cond, pivot)
		return goja.Undefined()
	})
	imguiObj.Set("internalSetNextWindowRefreshPolicy", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 WindowRefreshFlags
		flags := imgui.WindowRefreshFlags(0)
		imgui.InternalSetNextWindowRefreshPolicy(flags)
		return goja.Undefined()
	})
	imguiObj.Set("setNextWindowScroll", func(call goja.FunctionCall) goja.Value {
		scroll := parseVec2(vm, call.Argument(0))
		imgui.SetNextWindowScroll(scroll)
		return goja.Undefined()
	})
	imguiObj.Set("setNextWindowSizeV", func(call goja.FunctionCall) goja.Value {
		size := parseVec2(vm, call.Argument(0))
		// TODO: 处理类型 Cond
		cond := imgui.Cond(0)
		imgui.SetNextWindowSizeV(size, cond)
		return goja.Undefined()
	})
	imguiObj.Set("setNextWindowSizeConstraintsV", func(call goja.FunctionCall) goja.Value {
		size_min := parseVec2(vm, call.Argument(0))
		size_max := parseVec2(vm, call.Argument(0))
		// TODO: 处理函数类型 SizeCallback
		var custom_callback imgui.SizeCallback
		custom_callback_data := uintptr(call.Argument(1).ToInteger())
		imgui.SetNextWindowSizeConstraintsV(size_min, size_max, custom_callback, custom_callback_data)
		return goja.Undefined()
	})
	imguiObj.Set("setNextWindowViewport", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 ID
		viewport_id := imgui.ID(0)
		imgui.SetNextWindowViewport(viewport_id)
		return goja.Undefined()
	})
	imguiObj.Set("setScrollFromPosXFloatV", func(call goja.FunctionCall) goja.Value {
		local_x := float32(call.Argument(0).ToFloat())
		center_x_ratio := float32(call.Argument(1).ToFloat())
		imgui.SetScrollFromPosXFloatV(local_x, center_x_ratio)
		return goja.Undefined()
	})
	imguiObj.Set("internalSetScrollFromPosXWindowPtr", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Window
		var window *imgui.Window
		local_x := float32(call.Argument(1).ToFloat())
		center_x_ratio := float32(call.Argument(2).ToFloat())
		imgui.InternalSetScrollFromPosXWindowPtr(window, local_x, center_x_ratio)
		return goja.Undefined()
	})
	imguiObj.Set("setScrollFromPosYFloatV", func(call goja.FunctionCall) goja.Value {
		local_y := float32(call.Argument(0).ToFloat())
		center_y_ratio := float32(call.Argument(1).ToFloat())
		imgui.SetScrollFromPosYFloatV(local_y, center_y_ratio)
		return goja.Undefined()
	})
	imguiObj.Set("internalSetScrollFromPosYWindowPtr", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Window
		var window *imgui.Window
		local_y := float32(call.Argument(1).ToFloat())
		center_y_ratio := float32(call.Argument(2).ToFloat())
		imgui.InternalSetScrollFromPosYWindowPtr(window, local_y, center_y_ratio)
		return goja.Undefined()
	})
	imguiObj.Set("setScrollHereXV", func(call goja.FunctionCall) goja.Value {
		center_x_ratio := float32(call.Argument(0).ToFloat())
		imgui.SetScrollHereXV(center_x_ratio)
		return goja.Undefined()
	})
	imguiObj.Set("setScrollHereYV", func(call goja.FunctionCall) goja.Value {
		center_y_ratio := float32(call.Argument(0).ToFloat())
		imgui.SetScrollHereYV(center_y_ratio)
		return goja.Undefined()
	})
	imguiObj.Set("setScrollXFloat", func(call goja.FunctionCall) goja.Value {
		scroll_x := float32(call.Argument(0).ToFloat())
		imgui.SetScrollXFloat(scroll_x)
		return goja.Undefined()
	})
	imguiObj.Set("internalSetScrollXWindowPtr", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Window
		var window *imgui.Window
		scroll_x := float32(call.Argument(1).ToFloat())
		imgui.InternalSetScrollXWindowPtr(window, scroll_x)
		return goja.Undefined()
	})
	imguiObj.Set("setScrollYFloat", func(call goja.FunctionCall) goja.Value {
		scroll_y := float32(call.Argument(0).ToFloat())
		imgui.SetScrollYFloat(scroll_y)
		return goja.Undefined()
	})
	imguiObj.Set("internalSetScrollYWindowPtr", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Window
		var window *imgui.Window
		scroll_y := float32(call.Argument(1).ToFloat())
		imgui.InternalSetScrollYWindowPtr(window, scroll_y)
		return goja.Undefined()
	})
	imguiObj.Set("internalSetShortcutRouting", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 KeyChord
		key_chord := imgui.KeyChord(0)
		// TODO: 处理类型 InputFlags
		flags := imgui.InputFlags(0)
		// TODO: 处理类型 ID
		owner_id := imgui.ID(0)
		return vm.ToValue(imgui.InternalSetShortcutRouting(key_chord, flags, owner_id))
	})
	imguiObj.Set("setStateStorage", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Storage
		var storage *imgui.Storage
		imgui.SetStateStorage(storage)
		return goja.Undefined()
	})
	imguiObj.Set("setTabItemClosed", func(call goja.FunctionCall) goja.Value {
		tab_or_docked_window_label := call.Argument(0).String()
		imgui.SetTabItemClosed(tab_or_docked_window_label)
		return goja.Undefined()
	})
	imguiObj.Set("setTooltip", func(call goja.FunctionCall) goja.Value {
		fmt := call.Argument(0).String()
		imgui.SetTooltip(fmt)
		return goja.Undefined()
	})
	imguiObj.Set("internalSetWindowClipRectBeforeSetChannel", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Window
		var window *imgui.Window
		// TODO: 处理类型 Rect
		clip_rect := imgui.Rect{}
		imgui.InternalSetWindowClipRectBeforeSetChannel(window, clip_rect)
		return goja.Undefined()
	})
	imguiObj.Set("setWindowCollapsedBoolV", func(call goja.FunctionCall) goja.Value {
		collapsed := call.Argument(0).ToBoolean()
		// TODO: 处理类型 Cond
		cond := imgui.Cond(0)
		imgui.SetWindowCollapsedBoolV(collapsed, cond)
		return goja.Undefined()
	})
	imguiObj.Set("setWindowCollapsedStrV", func(call goja.FunctionCall) goja.Value {
		name := call.Argument(0).String()
		collapsed := call.Argument(1).ToBoolean()
		// TODO: 处理类型 Cond
		cond := imgui.Cond(0)
		imgui.SetWindowCollapsedStrV(name, collapsed, cond)
		return goja.Undefined()
	})
	imguiObj.Set("internalSetWindowCollapsedWindowPtrV", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Window
		var window *imgui.Window
		collapsed := call.Argument(1).ToBoolean()
		// TODO: 处理类型 Cond
		cond := imgui.Cond(0)
		imgui.InternalSetWindowCollapsedWindowPtrV(window, collapsed, cond)
		return goja.Undefined()
	})
	imguiObj.Set("internalSetWindowDock", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Window
		var window *imgui.Window
		// TODO: 处理类型 ID
		dock_id := imgui.ID(0)
		// TODO: 处理类型 Cond
		cond := imgui.Cond(0)
		imgui.InternalSetWindowDock(window, dock_id, cond)
		return goja.Undefined()
	})
	imguiObj.Set("setWindowFocus", func(call goja.FunctionCall) goja.Value {
		imgui.SetWindowFocus()
		return goja.Undefined()
	})
	imguiObj.Set("setWindowFocusStr", func(call goja.FunctionCall) goja.Value {
		name := call.Argument(0).String()
		imgui.SetWindowFocusStr(name)
		return goja.Undefined()
	})
	imguiObj.Set("setWindowFontScale", func(call goja.FunctionCall) goja.Value {
		scale := float32(call.Argument(0).ToFloat())
		imgui.SetWindowFontScale(scale)
		return goja.Undefined()
	})
	imguiObj.Set("internalSetWindowHiddenAndSkipItemsForCurrentFrame", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Window
		var window *imgui.Window
		imgui.InternalSetWindowHiddenAndSkipItemsForCurrentFrame(window)
		return goja.Undefined()
	})
	imguiObj.Set("internalSetWindowHitTestHole", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Window
		var window *imgui.Window
		pos := parseVec2(vm, call.Argument(1))
		size := parseVec2(vm, call.Argument(1))
		imgui.InternalSetWindowHitTestHole(window, pos, size)
		return goja.Undefined()
	})
	imguiObj.Set("internalSetWindowParentWindowForFocusRoute", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Window
		var window *imgui.Window
		// TODO: 处理指针类型 *Window
		var parent_window *imgui.Window
		imgui.InternalSetWindowParentWindowForFocusRoute(window, parent_window)
		return goja.Undefined()
	})
	imguiObj.Set("setWindowPosStrV", func(call goja.FunctionCall) goja.Value {
		name := call.Argument(0).String()
		pos := parseVec2(vm, call.Argument(1))
		// TODO: 处理类型 Cond
		cond := imgui.Cond(0)
		imgui.SetWindowPosStrV(name, pos, cond)
		return goja.Undefined()
	})
	imguiObj.Set("setWindowPosVec2V", func(call goja.FunctionCall) goja.Value {
		pos := parseVec2(vm, call.Argument(0))
		// TODO: 处理类型 Cond
		cond := imgui.Cond(0)
		imgui.SetWindowPosVec2V(pos, cond)
		return goja.Undefined()
	})
	imguiObj.Set("internalSetWindowPosWindowPtrV", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Window
		var window *imgui.Window
		pos := parseVec2(vm, call.Argument(1))
		// TODO: 处理类型 Cond
		cond := imgui.Cond(0)
		imgui.InternalSetWindowPosWindowPtrV(window, pos, cond)
		return goja.Undefined()
	})
	imguiObj.Set("setWindowSizeStrV", func(call goja.FunctionCall) goja.Value {
		name := call.Argument(0).String()
		size := parseVec2(vm, call.Argument(1))
		// TODO: 处理类型 Cond
		cond := imgui.Cond(0)
		imgui.SetWindowSizeStrV(name, size, cond)
		return goja.Undefined()
	})
	imguiObj.Set("setWindowSizeVec2V", func(call goja.FunctionCall) goja.Value {
		size := parseVec2(vm, call.Argument(0))
		// TODO: 处理类型 Cond
		cond := imgui.Cond(0)
		imgui.SetWindowSizeVec2V(size, cond)
		return goja.Undefined()
	})
	imguiObj.Set("internalSetWindowSizeWindowPtrV", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Window
		var window *imgui.Window
		size := parseVec2(vm, call.Argument(1))
		// TODO: 处理类型 Cond
		cond := imgui.Cond(0)
		imgui.InternalSetWindowSizeWindowPtrV(window, size, cond)
		return goja.Undefined()
	})
	imguiObj.Set("internalSetWindowViewport", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Window
		var window *imgui.Window
		// TODO: 处理指针类型 *ViewportP
		var viewport *imgui.ViewportP
		imgui.InternalSetWindowViewport(window, viewport)
		return goja.Undefined()
	})
	imguiObj.Set("internalShadeVertsLinearColorGradientKeepAlpha", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *DrawList
		var draw_list *imgui.DrawList
		vert_start_idx := int32(call.Argument(1).ToInteger())
		vert_end_idx := int32(call.Argument(2).ToInteger())
		gradient_p0 := parseVec2(vm, call.Argument(3))
		gradient_p1 := parseVec2(vm, call.Argument(3))
		col0 := uint32(call.Argument(3).ToInteger())
		col1 := uint32(call.Argument(4).ToInteger())
		imgui.InternalShadeVertsLinearColorGradientKeepAlpha(draw_list, vert_start_idx, vert_end_idx, gradient_p0, gradient_p1, col0, col1)
		return goja.Undefined()
	})
	imguiObj.Set("internalShadeVertsLinearUV", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *DrawList
		var draw_list *imgui.DrawList
		vert_start_idx := int32(call.Argument(1).ToInteger())
		vert_end_idx := int32(call.Argument(2).ToInteger())
		a := parseVec2(vm, call.Argument(3))
		b := parseVec2(vm, call.Argument(3))
		uv_a := parseVec2(vm, call.Argument(3))
		uv_b := parseVec2(vm, call.Argument(3))
		clamp := call.Argument(3).ToBoolean()
		imgui.InternalShadeVertsLinearUV(draw_list, vert_start_idx, vert_end_idx, a, b, uv_a, uv_b, clamp)
		return goja.Undefined()
	})
	imguiObj.Set("internalShadeVertsTransformPos", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *DrawList
		var draw_list *imgui.DrawList
		vert_start_idx := int32(call.Argument(1).ToInteger())
		vert_end_idx := int32(call.Argument(2).ToInteger())
		pivot_in := parseVec2(vm, call.Argument(3))
		cos_a := float32(call.Argument(3).ToFloat())
		sin_a := float32(call.Argument(4).ToFloat())
		pivot_out := parseVec2(vm, call.Argument(5))
		imgui.InternalShadeVertsTransformPos(draw_list, vert_start_idx, vert_end_idx, pivot_in, cos_a, sin_a, pivot_out)
		return goja.Undefined()
	})
	imguiObj.Set("internalShortcutID", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 KeyChord
		key_chord := imgui.KeyChord(0)
		// TODO: 处理类型 InputFlags
		flags := imgui.InputFlags(0)
		// TODO: 处理类型 ID
		owner_id := imgui.ID(0)
		return vm.ToValue(imgui.InternalShortcutID(key_chord, flags, owner_id))
	})
	imguiObj.Set("shortcutNilV", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 KeyChord
		key_chord := imgui.KeyChord(0)
		// TODO: 处理类型 InputFlags
		flags := imgui.InputFlags(0)
		return vm.ToValue(imgui.ShortcutNilV(key_chord, flags))
	})
	imguiObj.Set("showAboutWindowV", func(call goja.FunctionCall) goja.Value {
		p_open := false
		if len(call.Arguments) >= 1 {
			p_open = call.Argument(0).ToBoolean()
		}
		imgui.ShowAboutWindowV(&p_open)
		return goja.Undefined()
	})
	imguiObj.Set("showDebugLogWindowV", func(call goja.FunctionCall) goja.Value {
		p_open := false
		if len(call.Arguments) >= 1 {
			p_open = call.Argument(0).ToBoolean()
		}
		imgui.ShowDebugLogWindowV(&p_open)
		return goja.Undefined()
	})
	imguiObj.Set("showDemoWindowV", func(call goja.FunctionCall) goja.Value {
		p_open := false
		if len(call.Arguments) >= 1 {
			p_open = call.Argument(0).ToBoolean()
		}
		imgui.ShowDemoWindowV(&p_open)
		return goja.Undefined()
	})
	imguiObj.Set("internalShowFontAtlas", func(call goja.FunctionCall) goja.Value {
		var atlas *imgui.FontAtlas
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.FontAtlas); ok {
				atlas = v
			}
		}
		imgui.InternalShowFontAtlas(atlas)
		return goja.Undefined()
	})
	imguiObj.Set("showFontSelector", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		imgui.ShowFontSelector(label)
		return goja.Undefined()
	})
	imguiObj.Set("showIDStackToolWindowV", func(call goja.FunctionCall) goja.Value {
		p_open := false
		if len(call.Arguments) >= 1 {
			p_open = call.Argument(0).ToBoolean()
		}
		imgui.ShowIDStackToolWindowV(&p_open)
		return goja.Undefined()
	})
	imguiObj.Set("showMetricsWindowV", func(call goja.FunctionCall) goja.Value {
		p_open := false
		if len(call.Arguments) >= 1 {
			p_open = call.Argument(0).ToBoolean()
		}
		imgui.ShowMetricsWindowV(&p_open)
		return goja.Undefined()
	})
	imguiObj.Set("showStyleEditorV", func(call goja.FunctionCall) goja.Value {
		var ref *imgui.Style
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.Style); ok {
				ref = v
			}
		}
		imgui.ShowStyleEditorV(ref)
		return goja.Undefined()
	})
	imguiObj.Set("showStyleSelector", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		return vm.ToValue(imgui.ShowStyleSelector(label))
	})
	imguiObj.Set("showUserGuide", func(call goja.FunctionCall) goja.Value {
		imgui.ShowUserGuide()
		return goja.Undefined()
	})
	imguiObj.Set("internalShrinkWidths", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *ShrinkWidthItem
		var items *imgui.ShrinkWidthItem
		count := int32(call.Argument(1).ToInteger())
		width_excess := float32(call.Argument(2).ToFloat())
		width_min := float32(call.Argument(3).ToFloat())
		imgui.InternalShrinkWidths(items, count, width_excess, width_min)
		return goja.Undefined()
	})
	imguiObj.Set("internalShutdown", func(call goja.FunctionCall) goja.Value {
		imgui.InternalShutdown()
		return goja.Undefined()
	})
	imguiObj.Set("sliderAngleV", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		v_rad := float32(0)
		if len(call.Arguments) >= 2 {
			v_rad = float32(call.Argument(1).ToFloat())
		}
		v_degrees_min := float32(call.Argument(2).ToFloat())
		v_degrees_max := float32(call.Argument(3).ToFloat())
		format := call.Argument(4).String()
		// TODO: 处理类型 SliderFlags
		flags := imgui.SliderFlags(0)
		return vm.ToValue(imgui.SliderAngleV(label, &v_rad, v_degrees_min, v_degrees_max, format, flags))
	})
	imguiObj.Set("internalSliderBehavior", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 Rect
		bb := imgui.Rect{}
		// TODO: 处理类型 ID
		id := imgui.ID(0)
		// TODO: 处理类型 DataType
		data_type := imgui.DataType(0)
		p_v := uintptr(call.Argument(3).ToInteger())
		p_min := uintptr(call.Argument(4).ToInteger())
		p_max := uintptr(call.Argument(5).ToInteger())
		format := call.Argument(6).String()
		// TODO: 处理类型 SliderFlags
		flags := imgui.SliderFlags(0)
		// TODO: 处理指针类型 *Rect
		var out_grab_bb *imgui.Rect
		return vm.ToValue(imgui.InternalSliderBehavior(bb, id, data_type, p_v, p_min, p_max, format, flags, out_grab_bb))
	})
	imguiObj.Set("sliderFloatV", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		v := float32(0)
		if len(call.Arguments) >= 2 {
			v = float32(call.Argument(1).ToFloat())
		}
		v_min := float32(call.Argument(2).ToFloat())
		v_max := float32(call.Argument(3).ToFloat())
		format := call.Argument(4).String()
		// TODO: 处理类型 SliderFlags
		flags := imgui.SliderFlags(0)
		return vm.ToValue(imgui.SliderFloatV(label, &v, v_min, v_max, format, flags))
	})
	imguiObj.Set("sliderFloat2V", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		v := [2]float32{0, 0}
		v_min := float32(call.Argument(2).ToFloat())
		v_max := float32(call.Argument(3).ToFloat())
		format := call.Argument(4).String()
		// TODO: 处理类型 SliderFlags
		flags := imgui.SliderFlags(0)
		return vm.ToValue(imgui.SliderFloat2V(label, &v, v_min, v_max, format, flags))
	})
	imguiObj.Set("sliderFloat3V", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		v := [3]float32{0, 0, 0}
		v_min := float32(call.Argument(2).ToFloat())
		v_max := float32(call.Argument(3).ToFloat())
		format := call.Argument(4).String()
		// TODO: 处理类型 SliderFlags
		flags := imgui.SliderFlags(0)
		return vm.ToValue(imgui.SliderFloat3V(label, &v, v_min, v_max, format, flags))
	})
	imguiObj.Set("sliderFloat4V", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		v := [4]float32{0, 0, 0, 0}
		v_min := float32(call.Argument(2).ToFloat())
		v_max := float32(call.Argument(3).ToFloat())
		format := call.Argument(4).String()
		// TODO: 处理类型 SliderFlags
		flags := imgui.SliderFlags(0)
		return vm.ToValue(imgui.SliderFloat4V(label, &v, v_min, v_max, format, flags))
	})
	imguiObj.Set("sliderIntV", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		v := int32(0)
		if len(call.Arguments) >= 2 {
			v = int32(call.Argument(1).ToInteger())
		}
		v_min := int32(call.Argument(2).ToInteger())
		v_max := int32(call.Argument(3).ToInteger())
		format := call.Argument(4).String()
		// TODO: 处理类型 SliderFlags
		flags := imgui.SliderFlags(0)
		return vm.ToValue(imgui.SliderIntV(label, &v, v_min, v_max, format, flags))
	})
	imguiObj.Set("sliderInt2V", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		v := [2]int32{0, 0}
		v_min := int32(call.Argument(2).ToInteger())
		v_max := int32(call.Argument(3).ToInteger())
		format := call.Argument(4).String()
		// TODO: 处理类型 SliderFlags
		flags := imgui.SliderFlags(0)
		return vm.ToValue(imgui.SliderInt2V(label, &v, v_min, v_max, format, flags))
	})
	imguiObj.Set("sliderInt3V", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		v := [3]int32{0, 0, 0}
		v_min := int32(call.Argument(2).ToInteger())
		v_max := int32(call.Argument(3).ToInteger())
		format := call.Argument(4).String()
		// TODO: 处理类型 SliderFlags
		flags := imgui.SliderFlags(0)
		return vm.ToValue(imgui.SliderInt3V(label, &v, v_min, v_max, format, flags))
	})
	imguiObj.Set("sliderInt4V", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		v := [4]int32{0, 0, 0, 0}
		v_min := int32(call.Argument(2).ToInteger())
		v_max := int32(call.Argument(3).ToInteger())
		format := call.Argument(4).String()
		// TODO: 处理类型 SliderFlags
		flags := imgui.SliderFlags(0)
		return vm.ToValue(imgui.SliderInt4V(label, &v, v_min, v_max, format, flags))
	})
	imguiObj.Set("sliderScalarV", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		// TODO: 处理类型 DataType
		data_type := imgui.DataType(0)
		p_data := uintptr(call.Argument(2).ToInteger())
		p_min := uintptr(call.Argument(3).ToInteger())
		p_max := uintptr(call.Argument(4).ToInteger())
		format := call.Argument(5).String()
		// TODO: 处理类型 SliderFlags
		flags := imgui.SliderFlags(0)
		return vm.ToValue(imgui.SliderScalarV(label, data_type, p_data, p_min, p_max, format, flags))
	})
	imguiObj.Set("sliderScalarNV", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		// TODO: 处理类型 DataType
		data_type := imgui.DataType(0)
		p_data := uintptr(call.Argument(2).ToInteger())
		components := int32(call.Argument(3).ToInteger())
		p_min := uintptr(call.Argument(4).ToInteger())
		p_max := uintptr(call.Argument(5).ToInteger())
		format := call.Argument(6).String()
		// TODO: 处理类型 SliderFlags
		flags := imgui.SliderFlags(0)
		return vm.ToValue(imgui.SliderScalarNV(label, data_type, p_data, components, p_min, p_max, format, flags))
	})
	imguiObj.Set("smallButton", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		return vm.ToValue(imgui.SmallButton(label))
	})
	imguiObj.Set("spacing", func(call goja.FunctionCall) goja.Value {
		imgui.Spacing()
		return goja.Undefined()
	})
	imguiObj.Set("internalSplitterBehaviorV", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 Rect
		bb := imgui.Rect{}
		// TODO: 处理类型 ID
		id := imgui.ID(0)
		// TODO: 处理类型 Axis
		axis := imgui.Axis(0)
		size1 := float32(0)
		if len(call.Arguments) >= 4 {
			size1 = float32(call.Argument(3).ToFloat())
		}
		size2 := float32(0)
		if len(call.Arguments) >= 5 {
			size2 = float32(call.Argument(4).ToFloat())
		}
		min_size1 := float32(call.Argument(5).ToFloat())
		min_size2 := float32(call.Argument(6).ToFloat())
		hover_extend := float32(call.Argument(7).ToFloat())
		hover_visibility_delay := float32(call.Argument(8).ToFloat())
		bg_col := uint32(call.Argument(9).ToInteger())
		return vm.ToValue(imgui.InternalSplitterBehaviorV(bb, id, axis, &size1, &size2, min_size1, min_size2, hover_extend, hover_visibility_delay, bg_col))
	})
	imguiObj.Set("internalStartMouseMovingWindow", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Window
		var window *imgui.Window
		imgui.InternalStartMouseMovingWindow(window)
		return goja.Undefined()
	})
	imguiObj.Set("internalStartMouseMovingWindowOrNode", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Window
		var window *imgui.Window
		// TODO: 处理指针类型 *DockNode
		var node *imgui.DockNode
		undock := call.Argument(2).ToBoolean()
		imgui.InternalStartMouseMovingWindowOrNode(window, node, undock)
		return goja.Undefined()
	})
	imguiObj.Set("internalStopMouseMovingWindow", func(call goja.FunctionCall) goja.Value {
		imgui.InternalStopMouseMovingWindow()
		return goja.Undefined()
	})
	imguiObj.Set("styleColorsClassicV", func(call goja.FunctionCall) goja.Value {
		var dst *imgui.Style
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.Style); ok {
				dst = v
			}
		}
		imgui.StyleColorsClassicV(dst)
		return goja.Undefined()
	})
	imguiObj.Set("styleColorsDarkV", func(call goja.FunctionCall) goja.Value {
		var dst *imgui.Style
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.Style); ok {
				dst = v
			}
		}
		imgui.StyleColorsDarkV(dst)
		return goja.Undefined()
	})
	imguiObj.Set("styleColorsLightV", func(call goja.FunctionCall) goja.Value {
		var dst *imgui.Style
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.Style); ok {
				dst = v
			}
		}
		imgui.StyleColorsLightV(dst)
		return goja.Undefined()
	})
	imguiObj.Set("internalTabBarAddTab", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *TabBar
		var tab_bar *imgui.TabBar
		// TODO: 处理类型 TabItemFlags
		tab_flags := imgui.TabItemFlags(0)
		// TODO: 处理指针类型 *Window
		var window *imgui.Window
		imgui.InternalTabBarAddTab(tab_bar, tab_flags, window)
		return goja.Undefined()
	})
	imguiObj.Set("internalTabBarCloseTab", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *TabBar
		var tab_bar *imgui.TabBar
		// TODO: 处理指针类型 *TabItem
		var tab *imgui.TabItem
		imgui.InternalTabBarCloseTab(tab_bar, tab)
		return goja.Undefined()
	})
	imguiObj.Set("internalTabBarFindByID", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 ID
		id := imgui.ID(0)
		result := imgui.InternalTabBarFindByID(id)
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalTabBarFindMostRecentlySelectedTabForActiveWindow", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *TabBar
		var tab_bar *imgui.TabBar
		result := imgui.InternalTabBarFindMostRecentlySelectedTabForActiveWindow(tab_bar)
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalTabBarFindTabByID", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *TabBar
		var tab_bar *imgui.TabBar
		// TODO: 处理类型 ID
		tab_id := imgui.ID(0)
		result := imgui.InternalTabBarFindTabByID(tab_bar, tab_id)
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalTabBarFindTabByOrder", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *TabBar
		var tab_bar *imgui.TabBar
		order := int32(call.Argument(1).ToInteger())
		result := imgui.InternalTabBarFindTabByOrder(tab_bar, order)
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalTabBarGetCurrentTab", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *TabBar
		var tab_bar *imgui.TabBar
		result := imgui.InternalTabBarGetCurrentTab(tab_bar)
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalTabBarGetTabName", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *TabBar
		var tab_bar *imgui.TabBar
		// TODO: 处理指针类型 *TabItem
		var tab *imgui.TabItem
		return vm.ToValue(imgui.InternalTabBarGetTabName(tab_bar, tab))
	})
	imguiObj.Set("internalTabBarGetTabOrder", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *TabBar
		var tab_bar *imgui.TabBar
		// TODO: 处理指针类型 *TabItem
		var tab *imgui.TabItem
		return vm.ToValue(imgui.InternalTabBarGetTabOrder(tab_bar, tab))
	})
	imguiObj.Set("internalTabBarProcessReorder", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *TabBar
		var tab_bar *imgui.TabBar
		return vm.ToValue(imgui.InternalTabBarProcessReorder(tab_bar))
	})
	imguiObj.Set("internalTabBarQueueFocusStr", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *TabBar
		var tab_bar *imgui.TabBar
		tab_name := call.Argument(1).String()
		imgui.InternalTabBarQueueFocusStr(tab_bar, tab_name)
		return goja.Undefined()
	})
	imguiObj.Set("internalTabBarQueueFocusTabItemPtr", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *TabBar
		var tab_bar *imgui.TabBar
		// TODO: 处理指针类型 *TabItem
		var tab *imgui.TabItem
		imgui.InternalTabBarQueueFocusTabItemPtr(tab_bar, tab)
		return goja.Undefined()
	})
	imguiObj.Set("internalTabBarQueueReorder", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *TabBar
		var tab_bar *imgui.TabBar
		// TODO: 处理指针类型 *TabItem
		var tab *imgui.TabItem
		offset := int32(call.Argument(2).ToInteger())
		imgui.InternalTabBarQueueReorder(tab_bar, tab, offset)
		return goja.Undefined()
	})
	imguiObj.Set("internalTabBarQueueReorderFromMousePos", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *TabBar
		var tab_bar *imgui.TabBar
		// TODO: 处理指针类型 *TabItem
		var tab *imgui.TabItem
		mouse_pos := parseVec2(vm, call.Argument(2))
		imgui.InternalTabBarQueueReorderFromMousePos(tab_bar, tab, mouse_pos)
		return goja.Undefined()
	})
	imguiObj.Set("internalTabBarRemove", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *TabBar
		var tab_bar *imgui.TabBar
		imgui.InternalTabBarRemove(tab_bar)
		return goja.Undefined()
	})
	imguiObj.Set("internalTabBarRemoveTab", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *TabBar
		var tab_bar *imgui.TabBar
		// TODO: 处理类型 ID
		tab_id := imgui.ID(0)
		imgui.InternalTabBarRemoveTab(tab_bar, tab_id)
		return goja.Undefined()
	})
	imguiObj.Set("internalTabItemBackground", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *DrawList
		var draw_list *imgui.DrawList
		// TODO: 处理类型 Rect
		bb := imgui.Rect{}
		// TODO: 处理类型 TabItemFlags
		flags := imgui.TabItemFlags(0)
		col := uint32(call.Argument(3).ToInteger())
		imgui.InternalTabItemBackground(draw_list, bb, flags, col)
		return goja.Undefined()
	})
	imguiObj.Set("tabItemButtonV", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		// TODO: 处理类型 TabItemFlags
		flags := imgui.TabItemFlags(0)
		return vm.ToValue(imgui.TabItemButtonV(label, flags))
	})
	imguiObj.Set("internalTabItemCalcSizeStr", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		has_close_button_or_unsaved_marker := call.Argument(1).ToBoolean()
		result := imgui.InternalTabItemCalcSizeStr(label, has_close_button_or_unsaved_marker)
		return vm.ToValue(result)
	})
	imguiObj.Set("internalTabItemCalcSizeWindowPtr", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Window
		var window *imgui.Window
		result := imgui.InternalTabItemCalcSizeWindowPtr(window)
		return vm.ToValue(result)
	})
	imguiObj.Set("internalTabItemEx", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *TabBar
		var tab_bar *imgui.TabBar
		label := call.Argument(1).String()
		p_open := false
		if len(call.Arguments) >= 3 {
			p_open = call.Argument(2).ToBoolean()
		}
		// TODO: 处理类型 TabItemFlags
		flags := imgui.TabItemFlags(0)
		// TODO: 处理指针类型 *Window
		var docked_window *imgui.Window
		return vm.ToValue(imgui.InternalTabItemEx(tab_bar, label, &p_open, flags, docked_window))
	})
	imguiObj.Set("internalTabItemLabelAndCloseButton", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *DrawList
		var draw_list *imgui.DrawList
		// TODO: 处理类型 Rect
		bb := imgui.Rect{}
		// TODO: 处理类型 TabItemFlags
		flags := imgui.TabItemFlags(0)
		frame_padding := parseVec2(vm, call.Argument(3))
		label := call.Argument(3).String()
		// TODO: 处理类型 ID
		tab_id := imgui.ID(0)
		// TODO: 处理类型 ID
		close_button_id := imgui.ID(0)
		is_contents_visible := call.Argument(6).ToBoolean()
		out_just_closed := false
		if len(call.Arguments) >= 8 {
			out_just_closed = call.Argument(7).ToBoolean()
		}
		out_text_clipped := false
		if len(call.Arguments) >= 9 {
			out_text_clipped = call.Argument(8).ToBoolean()
		}
		imgui.InternalTabItemLabelAndCloseButton(draw_list, bb, flags, frame_padding, label, tab_id, close_button_id, is_contents_visible, &out_just_closed, &out_text_clipped)
		return goja.Undefined()
	})
	imguiObj.Set("internalTabItemSpacing", func(call goja.FunctionCall) goja.Value {
		str_id := call.Argument(0).String()
		// TODO: 处理类型 TabItemFlags
		flags := imgui.TabItemFlags(0)
		width := float32(call.Argument(2).ToFloat())
		imgui.InternalTabItemSpacing(str_id, flags, width)
		return goja.Undefined()
	})
	imguiObj.Set("tableAngledHeadersRow", func(call goja.FunctionCall) goja.Value {
		imgui.TableAngledHeadersRow()
		return goja.Undefined()
	})
	imguiObj.Set("internalTableAngledHeadersRowEx", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 ID
		row_id := imgui.ID(0)
		angle := float32(call.Argument(1).ToFloat())
		max_label_width := float32(call.Argument(2).ToFloat())
		// TODO: 处理指针类型 *TableHeaderData
		var data *imgui.TableHeaderData
		data_count := int32(call.Argument(4).ToInteger())
		imgui.InternalTableAngledHeadersRowEx(row_id, angle, max_label_width, data, data_count)
		return goja.Undefined()
	})
	imguiObj.Set("internalTableBeginApplyRequests", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Table
		var table *imgui.Table
		imgui.InternalTableBeginApplyRequests(table)
		return goja.Undefined()
	})
	imguiObj.Set("internalTableBeginCell", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Table
		var table *imgui.Table
		column_n := int32(call.Argument(1).ToInteger())
		imgui.InternalTableBeginCell(table, column_n)
		return goja.Undefined()
	})
	imguiObj.Set("internalTableBeginContextMenuPopup", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Table
		var table *imgui.Table
		return vm.ToValue(imgui.InternalTableBeginContextMenuPopup(table))
	})
	imguiObj.Set("internalTableBeginInitMemory", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Table
		var table *imgui.Table
		columns_count := int32(call.Argument(1).ToInteger())
		imgui.InternalTableBeginInitMemory(table, columns_count)
		return goja.Undefined()
	})
	imguiObj.Set("internalTableBeginRow", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Table
		var table *imgui.Table
		imgui.InternalTableBeginRow(table)
		return goja.Undefined()
	})
	imguiObj.Set("internalTableCalcMaxColumnWidth", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Table
		var table *imgui.Table
		column_n := int32(call.Argument(1).ToInteger())
		return vm.ToValue(imgui.InternalTableCalcMaxColumnWidth(table, column_n))
	})
	imguiObj.Set("internalTableDrawBorders", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Table
		var table *imgui.Table
		imgui.InternalTableDrawBorders(table)
		return goja.Undefined()
	})
	imguiObj.Set("internalTableDrawDefaultContextMenu", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Table
		var table *imgui.Table
		// TODO: 处理类型 TableFlags
		flags_for_section_to_display := imgui.TableFlags(0)
		imgui.InternalTableDrawDefaultContextMenu(table, flags_for_section_to_display)
		return goja.Undefined()
	})
	imguiObj.Set("internalTableEndCell", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Table
		var table *imgui.Table
		imgui.InternalTableEndCell(table)
		return goja.Undefined()
	})
	imguiObj.Set("internalTableEndRow", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Table
		var table *imgui.Table
		imgui.InternalTableEndRow(table)
		return goja.Undefined()
	})
	imguiObj.Set("internalTableFindByID", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 ID
		id := imgui.ID(0)
		result := imgui.InternalTableFindByID(id)
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalTableFixColumnSortDirection", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Table
		var table *imgui.Table
		// TODO: 处理指针类型 *TableColumn
		var column *imgui.TableColumn
		imgui.InternalTableFixColumnSortDirection(table, column)
		return goja.Undefined()
	})
	imguiObj.Set("internalTableGcCompactSettings", func(call goja.FunctionCall) goja.Value {
		imgui.InternalTableGcCompactSettings()
		return goja.Undefined()
	})
	imguiObj.Set("internalTableGcCompactTransientBuffersTablePtr", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Table
		var table *imgui.Table
		imgui.InternalTableGcCompactTransientBuffersTablePtr(table)
		return goja.Undefined()
	})
	imguiObj.Set("internalTableGcCompactTransientBuffersTableTempDataPtr", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *TableTempData
		var table *imgui.TableTempData
		imgui.InternalTableGcCompactTransientBuffersTableTempDataPtr(table)
		return goja.Undefined()
	})
	imguiObj.Set("internalTableGetBoundSettings", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Table
		var table *imgui.Table
		result := imgui.InternalTableGetBoundSettings(table)
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalTableGetCellBgRect", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Table
		var table *imgui.Table
		column_n := int32(call.Argument(1).ToInteger())
		return vm.ToValue(imgui.InternalTableGetCellBgRect(table, column_n))
	})
	imguiObj.Set("tableGetColumnCount", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.TableGetColumnCount())
	})
	imguiObj.Set("tableGetColumnFlagsV", func(call goja.FunctionCall) goja.Value {
		column_n := int32(call.Argument(0).ToInteger())
		return vm.ToValue(imgui.TableGetColumnFlagsV(column_n))
	})
	imguiObj.Set("tableGetColumnIndex", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.TableGetColumnIndex())
	})
	imguiObj.Set("tableGetColumnNameIntV", func(call goja.FunctionCall) goja.Value {
		column_n := int32(call.Argument(0).ToInteger())
		return vm.ToValue(imgui.TableGetColumnNameIntV(column_n))
	})
	imguiObj.Set("internalTableGetColumnNameTablePtr", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Table
		var table *imgui.Table
		column_n := int32(call.Argument(1).ToInteger())
		return vm.ToValue(imgui.InternalTableGetColumnNameTablePtr(table, column_n))
	})
	imguiObj.Set("internalTableGetColumnNextSortDirection", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *TableColumn
		var column *imgui.TableColumn
		return vm.ToValue(imgui.InternalTableGetColumnNextSortDirection(column))
	})
	imguiObj.Set("internalTableGetColumnResizeIDV", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Table
		var table *imgui.Table
		column_n := int32(call.Argument(1).ToInteger())
		instance_no := int32(call.Argument(2).ToInteger())
		return vm.ToValue(imgui.InternalTableGetColumnResizeIDV(table, column_n, instance_no))
	})
	imguiObj.Set("internalTableGetColumnWidthAuto", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Table
		var table *imgui.Table
		// TODO: 处理指针类型 *TableColumn
		var column *imgui.TableColumn
		return vm.ToValue(imgui.InternalTableGetColumnWidthAuto(table, column))
	})
	imguiObj.Set("internalTableGetHeaderAngledMaxLabelWidth", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.InternalTableGetHeaderAngledMaxLabelWidth())
	})
	imguiObj.Set("internalTableGetHeaderRowHeight", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.InternalTableGetHeaderRowHeight())
	})
	imguiObj.Set("tableGetHoveredColumn", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.TableGetHoveredColumn())
	})
	imguiObj.Set("internalTableGetHoveredRow", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.InternalTableGetHoveredRow())
	})
	imguiObj.Set("internalTableGetInstanceData", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Table
		var table *imgui.Table
		instance_no := int32(call.Argument(1).ToInteger())
		result := imgui.InternalTableGetInstanceData(table, instance_no)
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalTableGetInstanceID", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Table
		var table *imgui.Table
		instance_no := int32(call.Argument(1).ToInteger())
		return vm.ToValue(imgui.InternalTableGetInstanceID(table, instance_no))
	})
	imguiObj.Set("tableGetRowIndex", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.TableGetRowIndex())
	})
	imguiObj.Set("tableGetSortSpecs", func(call goja.FunctionCall) goja.Value {
		result := imgui.TableGetSortSpecs()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("tableHeader", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		imgui.TableHeader(label)
		return goja.Undefined()
	})
	imguiObj.Set("tableHeadersRow", func(call goja.FunctionCall) goja.Value {
		imgui.TableHeadersRow()
		return goja.Undefined()
	})
	imguiObj.Set("internalTableLoadSettings", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Table
		var table *imgui.Table
		imgui.InternalTableLoadSettings(table)
		return goja.Undefined()
	})
	imguiObj.Set("internalTableMergeDrawChannels", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Table
		var table *imgui.Table
		imgui.InternalTableMergeDrawChannels(table)
		return goja.Undefined()
	})
	imguiObj.Set("tableNextColumn", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.TableNextColumn())
	})
	imguiObj.Set("tableNextRowV", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 TableRowFlags
		row_flags := imgui.TableRowFlags(0)
		min_row_height := float32(call.Argument(1).ToFloat())
		imgui.TableNextRowV(row_flags, min_row_height)
		return goja.Undefined()
	})
	imguiObj.Set("internalTableOpenContextMenuV", func(call goja.FunctionCall) goja.Value {
		column_n := int32(call.Argument(0).ToInteger())
		imgui.InternalTableOpenContextMenuV(column_n)
		return goja.Undefined()
	})
	imguiObj.Set("internalTablePopBackgroundChannel", func(call goja.FunctionCall) goja.Value {
		imgui.InternalTablePopBackgroundChannel()
		return goja.Undefined()
	})
	imguiObj.Set("internalTablePopColumnChannel", func(call goja.FunctionCall) goja.Value {
		imgui.InternalTablePopColumnChannel()
		return goja.Undefined()
	})
	imguiObj.Set("internalTablePushBackgroundChannel", func(call goja.FunctionCall) goja.Value {
		imgui.InternalTablePushBackgroundChannel()
		return goja.Undefined()
	})
	imguiObj.Set("internalTablePushColumnChannel", func(call goja.FunctionCall) goja.Value {
		column_n := int32(call.Argument(0).ToInteger())
		imgui.InternalTablePushColumnChannel(column_n)
		return goja.Undefined()
	})
	imguiObj.Set("internalTableRemove", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Table
		var table *imgui.Table
		imgui.InternalTableRemove(table)
		return goja.Undefined()
	})
	imguiObj.Set("internalTableResetSettings", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Table
		var table *imgui.Table
		imgui.InternalTableResetSettings(table)
		return goja.Undefined()
	})
	imguiObj.Set("internalTableSaveSettings", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Table
		var table *imgui.Table
		imgui.InternalTableSaveSettings(table)
		return goja.Undefined()
	})
	imguiObj.Set("tableSetBgColorV", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 TableBgTarget
		target := imgui.TableBgTarget(0)
		color := uint32(call.Argument(1).ToInteger())
		column_n := int32(call.Argument(2).ToInteger())
		imgui.TableSetBgColorV(target, color, column_n)
		return goja.Undefined()
	})
	imguiObj.Set("tableSetColumnEnabled", func(call goja.FunctionCall) goja.Value {
		column_n := int32(call.Argument(0).ToInteger())
		v := call.Argument(1).ToBoolean()
		imgui.TableSetColumnEnabled(column_n, v)
		return goja.Undefined()
	})
	imguiObj.Set("tableSetColumnIndex", func(call goja.FunctionCall) goja.Value {
		column_n := int32(call.Argument(0).ToInteger())
		return vm.ToValue(imgui.TableSetColumnIndex(column_n))
	})
	imguiObj.Set("internalTableSetColumnSortDirection", func(call goja.FunctionCall) goja.Value {
		column_n := int32(call.Argument(0).ToInteger())
		// TODO: 处理类型 SortDirection
		sort_direction := imgui.SortDirection(0)
		append_to_sort_specs := call.Argument(2).ToBoolean()
		imgui.InternalTableSetColumnSortDirection(column_n, sort_direction, append_to_sort_specs)
		return goja.Undefined()
	})
	imguiObj.Set("internalTableSetColumnWidth", func(call goja.FunctionCall) goja.Value {
		column_n := int32(call.Argument(0).ToInteger())
		width := float32(call.Argument(1).ToFloat())
		imgui.InternalTableSetColumnWidth(column_n, width)
		return goja.Undefined()
	})
	imguiObj.Set("internalTableSetColumnWidthAutoAll", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Table
		var table *imgui.Table
		imgui.InternalTableSetColumnWidthAutoAll(table)
		return goja.Undefined()
	})
	imguiObj.Set("internalTableSetColumnWidthAutoSingle", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Table
		var table *imgui.Table
		column_n := int32(call.Argument(1).ToInteger())
		imgui.InternalTableSetColumnWidthAutoSingle(table, column_n)
		return goja.Undefined()
	})
	imguiObj.Set("internalTableSettingsAddSettingsHandler", func(call goja.FunctionCall) goja.Value {
		imgui.InternalTableSettingsAddSettingsHandler()
		return goja.Undefined()
	})
	imguiObj.Set("internalTableSettingsCreate", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 ID
		id := imgui.ID(0)
		columns_count := int32(call.Argument(1).ToInteger())
		result := imgui.InternalTableSettingsCreate(id, columns_count)
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalTableSettingsFindByID", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 ID
		id := imgui.ID(0)
		result := imgui.InternalTableSettingsFindByID(id)
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("tableSetupColumnV", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		// TODO: 处理类型 TableColumnFlags
		flags := imgui.TableColumnFlags(0)
		init_width_or_weight := float32(call.Argument(2).ToFloat())
		// TODO: 处理类型 ID
		user_id := imgui.ID(0)
		imgui.TableSetupColumnV(label, flags, init_width_or_weight, user_id)
		return goja.Undefined()
	})
	imguiObj.Set("internalTableSetupDrawChannels", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Table
		var table *imgui.Table
		imgui.InternalTableSetupDrawChannels(table)
		return goja.Undefined()
	})
	imguiObj.Set("tableSetupScrollFreeze", func(call goja.FunctionCall) goja.Value {
		cols := int32(call.Argument(0).ToInteger())
		rows := int32(call.Argument(1).ToInteger())
		imgui.TableSetupScrollFreeze(cols, rows)
		return goja.Undefined()
	})
	imguiObj.Set("internalTableSortSpecsBuild", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Table
		var table *imgui.Table
		imgui.InternalTableSortSpecsBuild(table)
		return goja.Undefined()
	})
	imguiObj.Set("internalTableSortSpecsSanitize", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Table
		var table *imgui.Table
		imgui.InternalTableSortSpecsSanitize(table)
		return goja.Undefined()
	})
	imguiObj.Set("internalTableUpdateBorders", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Table
		var table *imgui.Table
		imgui.InternalTableUpdateBorders(table)
		return goja.Undefined()
	})
	imguiObj.Set("internalTableUpdateColumnsWeightFromWidth", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Table
		var table *imgui.Table
		imgui.InternalTableUpdateColumnsWeightFromWidth(table)
		return goja.Undefined()
	})
	imguiObj.Set("internalTableUpdateLayout", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Table
		var table *imgui.Table
		imgui.InternalTableUpdateLayout(table)
		return goja.Undefined()
	})
	imguiObj.Set("internalTeleportMousePos", func(call goja.FunctionCall) goja.Value {
		pos := parseVec2(vm, call.Argument(0))
		imgui.InternalTeleportMousePos(pos)
		return goja.Undefined()
	})
	imguiObj.Set("internalTempInputIsActive", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 ID
		id := imgui.ID(0)
		return vm.ToValue(imgui.InternalTempInputIsActive(id))
	})
	imguiObj.Set("internalTempInputScalarV", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 Rect
		bb := imgui.Rect{}
		// TODO: 处理类型 ID
		id := imgui.ID(0)
		label := call.Argument(2).String()
		// TODO: 处理类型 DataType
		data_type := imgui.DataType(0)
		p_data := uintptr(call.Argument(4).ToInteger())
		format := call.Argument(5).String()
		p_clamp_min := uintptr(call.Argument(6).ToInteger())
		p_clamp_max := uintptr(call.Argument(7).ToInteger())
		return vm.ToValue(imgui.InternalTempInputScalarV(bb, id, label, data_type, p_data, format, p_clamp_min, p_clamp_max))
	})
	imguiObj.Set("internalTempInputText", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 Rect
		bb := imgui.Rect{}
		// TODO: 处理类型 ID
		id := imgui.ID(0)
		label := call.Argument(2).String()
		buf := call.Argument(3).String()
		buf_size := int32(call.Argument(4).ToInteger())
		// TODO: 处理类型 InputTextFlags
		flags := imgui.InputTextFlags(0)
		return vm.ToValue(imgui.InternalTempInputText(bb, id, label, buf, buf_size, flags))
	})
	imguiObj.Set("internalTestKeyOwner", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 Key
		key := imgui.Key(0)
		// TODO: 处理类型 ID
		owner_id := imgui.ID(0)
		return vm.ToValue(imgui.InternalTestKeyOwner(key, owner_id))
	})
	imguiObj.Set("internalTestShortcutRouting", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 KeyChord
		key_chord := imgui.KeyChord(0)
		// TODO: 处理类型 ID
		owner_id := imgui.ID(0)
		return vm.ToValue(imgui.InternalTestShortcutRouting(key_chord, owner_id))
	})
	imguiObj.Set("text", func(call goja.FunctionCall) goja.Value {
		fmt := call.Argument(0).String()
		imgui.Text(fmt)
		return goja.Undefined()
	})
	imguiObj.Set("internalTextAligned", func(call goja.FunctionCall) goja.Value {
		align_x := float32(call.Argument(0).ToFloat())
		size_x := float32(call.Argument(1).ToFloat())
		fmt := call.Argument(2).String()
		imgui.InternalTextAligned(align_x, size_x, fmt)
		return goja.Undefined()
	})
	imguiObj.Set("textColored", func(call goja.FunctionCall) goja.Value {
		col := parseVec4(vm, call.Argument(0))
		fmt := call.Argument(0).String()
		imgui.TextColored(col, fmt)
		return goja.Undefined()
	})
	imguiObj.Set("textDisabled", func(call goja.FunctionCall) goja.Value {
		fmt := call.Argument(0).String()
		imgui.TextDisabled(fmt)
		return goja.Undefined()
	})
	imguiObj.Set("internalTextExV", func(call goja.FunctionCall) goja.Value {
		text := call.Argument(0).String()
		// TODO: 处理类型 TextFlags
		flags := imgui.TextFlags(0)
		imgui.InternalTextExV(text, flags)
		return goja.Undefined()
	})
	imguiObj.Set("textLink", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		return vm.ToValue(imgui.TextLink(label))
	})
	imguiObj.Set("textLinkOpenURLV", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		url := call.Argument(1).String()
		return vm.ToValue(imgui.TextLinkOpenURLV(label, url))
	})
	imguiObj.Set("textUnformattedV", func(call goja.FunctionCall) goja.Value {
		text := call.Argument(0).String()
		imgui.TextUnformattedV(text)
		return goja.Undefined()
	})
	imguiObj.Set("textWrapped", func(call goja.FunctionCall) goja.Value {
		fmt := call.Argument(0).String()
		imgui.TextWrapped(fmt)
		return goja.Undefined()
	})
	imguiObj.Set("internalTranslateWindowsInViewport", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *ViewportP
		var viewport *imgui.ViewportP
		old_pos := parseVec2(vm, call.Argument(1))
		new_pos := parseVec2(vm, call.Argument(1))
		old_size := parseVec2(vm, call.Argument(1))
		new_size := parseVec2(vm, call.Argument(1))
		imgui.InternalTranslateWindowsInViewport(viewport, old_pos, new_pos, old_size, new_size)
		return goja.Undefined()
	})
	imguiObj.Set("internalTreeNodeBehaviorV", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 ID
		id := imgui.ID(0)
		// TODO: 处理类型 TreeNodeFlags
		flags := imgui.TreeNodeFlags(0)
		label := call.Argument(2).String()
		label_end := call.Argument(3).String()
		return vm.ToValue(imgui.InternalTreeNodeBehaviorV(id, flags, label, label_end))
	})
	imguiObj.Set("internalTreeNodeDrawLineToChildNode", func(call goja.FunctionCall) goja.Value {
		target_pos := parseVec2(vm, call.Argument(0))
		imgui.InternalTreeNodeDrawLineToChildNode(target_pos)
		return goja.Undefined()
	})
	imguiObj.Set("internalTreeNodeDrawLineToTreePop", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *TreeNodeStackData
		var data *imgui.TreeNodeStackData
		imgui.InternalTreeNodeDrawLineToTreePop(data)
		return goja.Undefined()
	})
	imguiObj.Set("treeNodeExPtr", func(call goja.FunctionCall) goja.Value {
		ptr_id := uintptr(call.Argument(0).ToInteger())
		// TODO: 处理类型 TreeNodeFlags
		flags := imgui.TreeNodeFlags(0)
		fmt := call.Argument(2).String()
		return vm.ToValue(imgui.TreeNodeExPtr(ptr_id, flags, fmt))
	})
	imguiObj.Set("treeNodeExStrV", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		// TODO: 处理类型 TreeNodeFlags
		flags := imgui.TreeNodeFlags(0)
		return vm.ToValue(imgui.TreeNodeExStrV(label, flags))
	})
	imguiObj.Set("treeNodeExStrStr", func(call goja.FunctionCall) goja.Value {
		str_id := call.Argument(0).String()
		// TODO: 处理类型 TreeNodeFlags
		flags := imgui.TreeNodeFlags(0)
		fmt := call.Argument(2).String()
		return vm.ToValue(imgui.TreeNodeExStrStr(str_id, flags, fmt))
	})
	imguiObj.Set("internalTreeNodeGetOpen", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 ID
		storage_id := imgui.ID(0)
		return vm.ToValue(imgui.InternalTreeNodeGetOpen(storage_id))
	})
	imguiObj.Set("internalTreeNodeSetOpen", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 ID
		storage_id := imgui.ID(0)
		open := call.Argument(1).ToBoolean()
		imgui.InternalTreeNodeSetOpen(storage_id, open)
		return goja.Undefined()
	})
	imguiObj.Set("internalTreeNodeUpdateNextOpen", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 ID
		storage_id := imgui.ID(0)
		// TODO: 处理类型 TreeNodeFlags
		flags := imgui.TreeNodeFlags(0)
		return vm.ToValue(imgui.InternalTreeNodeUpdateNextOpen(storage_id, flags))
	})
	imguiObj.Set("treeNodePtr", func(call goja.FunctionCall) goja.Value {
		ptr_id := uintptr(call.Argument(0).ToInteger())
		fmt := call.Argument(1).String()
		return vm.ToValue(imgui.TreeNodePtr(ptr_id, fmt))
	})
	imguiObj.Set("treeNodeStr", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		return vm.ToValue(imgui.TreeNodeStr(label))
	})
	imguiObj.Set("treeNodeStrStr", func(call goja.FunctionCall) goja.Value {
		str_id := call.Argument(0).String()
		fmt := call.Argument(1).String()
		return vm.ToValue(imgui.TreeNodeStrStr(str_id, fmt))
	})
	imguiObj.Set("treePop", func(call goja.FunctionCall) goja.Value {
		imgui.TreePop()
		return goja.Undefined()
	})
	imguiObj.Set("internalTreePushOverrideID", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 ID
		id := imgui.ID(0)
		imgui.InternalTreePushOverrideID(id)
		return goja.Undefined()
	})
	imguiObj.Set("treePushPtr", func(call goja.FunctionCall) goja.Value {
		ptr_id := uintptr(call.Argument(0).ToInteger())
		imgui.TreePushPtr(ptr_id)
		return goja.Undefined()
	})
	imguiObj.Set("treePushStr", func(call goja.FunctionCall) goja.Value {
		str_id := call.Argument(0).String()
		imgui.TreePushStr(str_id)
		return goja.Undefined()
	})
	imguiObj.Set("unindentV", func(call goja.FunctionCall) goja.Value {
		indent_w := float32(call.Argument(0).ToFloat())
		imgui.UnindentV(indent_w)
		return goja.Undefined()
	})
	imguiObj.Set("internalUnregisterFontAtlas", func(call goja.FunctionCall) goja.Value {
		var atlas *imgui.FontAtlas
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.FontAtlas); ok {
				atlas = v
			}
		}
		imgui.InternalUnregisterFontAtlas(atlas)
		return goja.Undefined()
	})
	imguiObj.Set("internalUnregisterUserTexture", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *TextureData
		var tex *imgui.TextureData
		imgui.InternalUnregisterUserTexture(tex)
		return goja.Undefined()
	})
	imguiObj.Set("internalUpdateCurrentFontSize", func(call goja.FunctionCall) goja.Value {
		restore_font_size_after_scaling := float32(call.Argument(0).ToFloat())
		imgui.InternalUpdateCurrentFontSize(restore_font_size_after_scaling)
		return goja.Undefined()
	})
	imguiObj.Set("internalUpdateHoveredWindowAndCaptureFlags", func(call goja.FunctionCall) goja.Value {
		mouse_pos := parseVec2(vm, call.Argument(0))
		imgui.InternalUpdateHoveredWindowAndCaptureFlags(mouse_pos)
		return goja.Undefined()
	})
	imguiObj.Set("internalUpdateInputEvents", func(call goja.FunctionCall) goja.Value {
		trickle_fast_inputs := call.Argument(0).ToBoolean()
		imgui.InternalUpdateInputEvents(trickle_fast_inputs)
		return goja.Undefined()
	})
	imguiObj.Set("internalUpdateMouseMovingWindowEndFrame", func(call goja.FunctionCall) goja.Value {
		imgui.InternalUpdateMouseMovingWindowEndFrame()
		return goja.Undefined()
	})
	imguiObj.Set("internalUpdateMouseMovingWindowNewFrame", func(call goja.FunctionCall) goja.Value {
		imgui.InternalUpdateMouseMovingWindowNewFrame()
		return goja.Undefined()
	})
	imguiObj.Set("updatePlatformWindows", func(call goja.FunctionCall) goja.Value {
		imgui.UpdatePlatformWindows()
		return goja.Undefined()
	})
	imguiObj.Set("internalUpdateWindowParentAndRootLinks", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Window
		var window *imgui.Window
		// TODO: 处理类型 WindowFlags
		flags := imgui.WindowFlags(0)
		// TODO: 处理指针类型 *Window
		var parent_window *imgui.Window
		imgui.InternalUpdateWindowParentAndRootLinks(window, flags, parent_window)
		return goja.Undefined()
	})
	imguiObj.Set("internalUpdateWindowSkipRefresh", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Window
		var window *imgui.Window
		imgui.InternalUpdateWindowSkipRefresh(window)
		return goja.Undefined()
	})
	imguiObj.Set("vSliderFloatV", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		size := parseVec2(vm, call.Argument(1))
		v := float32(0)
		if len(call.Arguments) >= 2 {
			v = float32(call.Argument(1).ToFloat())
		}
		v_min := float32(call.Argument(2).ToFloat())
		v_max := float32(call.Argument(3).ToFloat())
		format := call.Argument(4).String()
		// TODO: 处理类型 SliderFlags
		flags := imgui.SliderFlags(0)
		return vm.ToValue(imgui.VSliderFloatV(label, size, &v, v_min, v_max, format, flags))
	})
	imguiObj.Set("vSliderIntV", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		size := parseVec2(vm, call.Argument(1))
		v := int32(0)
		if len(call.Arguments) >= 2 {
			v = int32(call.Argument(1).ToInteger())
		}
		v_min := int32(call.Argument(2).ToInteger())
		v_max := int32(call.Argument(3).ToInteger())
		format := call.Argument(4).String()
		// TODO: 处理类型 SliderFlags
		flags := imgui.SliderFlags(0)
		return vm.ToValue(imgui.VSliderIntV(label, size, &v, v_min, v_max, format, flags))
	})
	imguiObj.Set("vSliderScalarV", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		size := parseVec2(vm, call.Argument(1))
		// TODO: 处理类型 DataType
		data_type := imgui.DataType(0)
		p_data := uintptr(call.Argument(2).ToInteger())
		p_min := uintptr(call.Argument(3).ToInteger())
		p_max := uintptr(call.Argument(4).ToInteger())
		format := call.Argument(5).String()
		// TODO: 处理类型 SliderFlags
		flags := imgui.SliderFlags(0)
		return vm.ToValue(imgui.VSliderScalarV(label, size, data_type, p_data, p_min, p_max, format, flags))
	})
	imguiObj.Set("valueBool", func(call goja.FunctionCall) goja.Value {
		prefix := call.Argument(0).String()
		b := call.Argument(1).ToBoolean()
		imgui.ValueBool(prefix, b)
		return goja.Undefined()
	})
	imguiObj.Set("valueFloatV", func(call goja.FunctionCall) goja.Value {
		prefix := call.Argument(0).String()
		v := float32(call.Argument(1).ToFloat())
		float_format := call.Argument(2).String()
		imgui.ValueFloatV(prefix, v, float_format)
		return goja.Undefined()
	})
	imguiObj.Set("valueInt", func(call goja.FunctionCall) goja.Value {
		prefix := call.Argument(0).String()
		v := int32(call.Argument(1).ToInteger())
		imgui.ValueInt(prefix, v)
		return goja.Undefined()
	})
	imguiObj.Set("valueUint", func(call goja.FunctionCall) goja.Value {
		prefix := call.Argument(0).String()
		v := uint32(call.Argument(1).ToInteger())
		imgui.ValueUint(prefix, v)
		return goja.Undefined()
	})
	imguiObj.Set("internalWindowPosAbsToRel", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Window
		var window *imgui.Window
		p := parseVec2(vm, call.Argument(1))
		result := imgui.InternalWindowPosAbsToRel(window, p)
		return vm.ToValue(result)
	})
	imguiObj.Set("internalWindowPosRelToAbs", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Window
		var window *imgui.Window
		p := parseVec2(vm, call.Argument(1))
		result := imgui.InternalWindowPosRelToAbs(window, p)
		return vm.ToValue(result)
	})
	imguiObj.Set("internalWindowRectAbsToRel", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Window
		var window *imgui.Window
		// TODO: 处理类型 Rect
		r := imgui.Rect{}
		return vm.ToValue(imgui.InternalWindowRectAbsToRel(window, r))
	})
	imguiObj.Set("internalWindowRectRelToAbs", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Window
		var window *imgui.Window
		// TODO: 处理类型 Rect
		r := imgui.Rect{}
		return vm.ToValue(imgui.InternalWindowRectRelToAbs(window, r))
	})
	imguiObj.Set("colorHSV", func(call goja.FunctionCall) goja.Value {
		h := float32(call.Argument(0).ToFloat())
		s := float32(call.Argument(1).ToFloat())
		v := float32(call.Argument(2).ToFloat())
		return vm.ToValue(imgui.ColorHSV(h, s, v))
	})
	imguiObj.Set("acceptDragDropPayload", func(call goja.FunctionCall) goja.Value {
		typeArg := call.Argument(0).String()
		result := imgui.AcceptDragDropPayload(typeArg)
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalArrowButtonEx", func(call goja.FunctionCall) goja.Value {
		str_id := call.Argument(0).String()
		// TODO: 处理类型 Dir
		dir := imgui.Dir(0)
		size_arg := parseVec2(vm, call.Argument(2))
		return vm.ToValue(imgui.InternalArrowButtonEx(str_id, dir, size_arg))
	})
	imguiObj.Set("begin", func(call goja.FunctionCall) goja.Value {
		name := call.Argument(0).String()
		return vm.ToValue(imgui.Begin(name))
	})
	imguiObj.Set("beginChildID", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 ID
		id := imgui.ID(0)
		return vm.ToValue(imgui.BeginChildID(id))
	})
	imguiObj.Set("beginChildStr", func(call goja.FunctionCall) goja.Value {
		str_id := call.Argument(0).String()
		return vm.ToValue(imgui.BeginChildStr(str_id))
	})
	imguiObj.Set("internalBeginColumns", func(call goja.FunctionCall) goja.Value {
		str_id := call.Argument(0).String()
		count := int32(call.Argument(1).ToInteger())
		imgui.InternalBeginColumns(str_id, count)
		return goja.Undefined()
	})
	imguiObj.Set("beginCombo", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		preview_value := call.Argument(1).String()
		return vm.ToValue(imgui.BeginCombo(label, preview_value))
	})
	imguiObj.Set("beginDisabled", func(call goja.FunctionCall) goja.Value {
		imgui.BeginDisabled()
		return goja.Undefined()
	})
	imguiObj.Set("beginDragDropSource", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.BeginDragDropSource())
	})
	imguiObj.Set("internalBeginDragDropTargetViewport", func(call goja.FunctionCall) goja.Value {
		var viewport *imgui.Viewport
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.Viewport); ok {
				viewport = v
			}
		}
		return vm.ToValue(imgui.InternalBeginDragDropTargetViewport(viewport))
	})
	imguiObj.Set("beginListBox", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		return vm.ToValue(imgui.BeginListBox(label))
	})
	imguiObj.Set("beginMenu", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		return vm.ToValue(imgui.BeginMenu(label))
	})
	imguiObj.Set("internalBeginMenuEx", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		icon := call.Argument(1).String()
		return vm.ToValue(imgui.InternalBeginMenuEx(label, icon))
	})
	imguiObj.Set("beginMultiSelect", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 MultiSelectFlags
		flags := imgui.MultiSelectFlags(0)
		result := imgui.BeginMultiSelect(flags)
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("beginPopup", func(call goja.FunctionCall) goja.Value {
		str_id := call.Argument(0).String()
		return vm.ToValue(imgui.BeginPopup(str_id))
	})
	imguiObj.Set("beginPopupContextItem", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.BeginPopupContextItem())
	})
	imguiObj.Set("beginPopupContextVoid", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.BeginPopupContextVoid())
	})
	imguiObj.Set("beginPopupContextWindow", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.BeginPopupContextWindow())
	})
	imguiObj.Set("beginPopupModal", func(call goja.FunctionCall) goja.Value {
		name := call.Argument(0).String()
		return vm.ToValue(imgui.BeginPopupModal(name))
	})
	imguiObj.Set("beginTabBar", func(call goja.FunctionCall) goja.Value {
		str_id := call.Argument(0).String()
		return vm.ToValue(imgui.BeginTabBar(str_id))
	})
	imguiObj.Set("beginTabItem", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		return vm.ToValue(imgui.BeginTabItem(label))
	})
	imguiObj.Set("beginTable", func(call goja.FunctionCall) goja.Value {
		str_id := call.Argument(0).String()
		columns := int32(call.Argument(1).ToInteger())
		return vm.ToValue(imgui.BeginTable(str_id, columns))
	})
	imguiObj.Set("internalBeginTableEx", func(call goja.FunctionCall) goja.Value {
		name := call.Argument(0).String()
		// TODO: 处理类型 ID
		id := imgui.ID(0)
		columns_count := int32(call.Argument(2).ToInteger())
		return vm.ToValue(imgui.InternalBeginTableEx(name, id, columns_count))
	})
	imguiObj.Set("button", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		return vm.ToValue(imgui.Button(label))
	})
	imguiObj.Set("internalButtonBehavior", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 Rect
		bb := imgui.Rect{}
		// TODO: 处理类型 ID
		id := imgui.ID(0)
		out_hovered := false
		if len(call.Arguments) >= 3 {
			out_hovered = call.Argument(2).ToBoolean()
		}
		out_held := false
		if len(call.Arguments) >= 4 {
			out_held = call.Argument(3).ToBoolean()
		}
		return vm.ToValue(imgui.InternalButtonBehavior(bb, id, &out_hovered, &out_held))
	})
	imguiObj.Set("internalButtonEx", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		return vm.ToValue(imgui.InternalButtonEx(label))
	})
	imguiObj.Set("calcTextSize", func(call goja.FunctionCall) goja.Value {
		text := call.Argument(0).String()
		result := imgui.CalcTextSize(text)
		return vm.ToValue(result)
	})
	imguiObj.Set("collapsingHeaderBoolPtr", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		p_visible := false
		if len(call.Arguments) >= 2 {
			p_visible = call.Argument(1).ToBoolean()
		}
		return vm.ToValue(imgui.CollapsingHeaderBoolPtr(label, &p_visible))
	})
	imguiObj.Set("collapsingHeaderTreeNodeFlags", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		return vm.ToValue(imgui.CollapsingHeaderTreeNodeFlags(label))
	})
	imguiObj.Set("colorButton", func(call goja.FunctionCall) goja.Value {
		desc_id := call.Argument(0).String()
		col := parseVec4(vm, call.Argument(1))
		return vm.ToValue(imgui.ColorButton(desc_id, col))
	})
	imguiObj.Set("colorEdit3", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		col := [3]float32{0, 0, 0}
		return vm.ToValue(imgui.ColorEdit3(label, &col))
	})
	imguiObj.Set("colorEdit4", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		col := [4]float32{0, 0, 0, 0}
		return vm.ToValue(imgui.ColorEdit4(label, &col))
	})
	imguiObj.Set("colorPicker3", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		col := [3]float32{0, 0, 0}
		return vm.ToValue(imgui.ColorPicker3(label, &col))
	})
	imguiObj.Set("colorPicker4", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		col := [4]float32{0, 0, 0, 0}
		return vm.ToValue(imgui.ColorPicker4(label, &col))
	})
	imguiObj.Set("columns", func(call goja.FunctionCall) goja.Value {
		imgui.Columns()
		return goja.Undefined()
	})
	imguiObj.Set("comboStr", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		current_item := int32(0)
		if len(call.Arguments) >= 2 {
			current_item = int32(call.Argument(1).ToInteger())
		}
		items_separated_by_zeros := call.Argument(2).String()
		return vm.ToValue(imgui.ComboStr(label, &current_item, items_separated_by_zeros))
	})
	imguiObj.Set("comboStrarr", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		current_item := int32(0)
		if len(call.Arguments) >= 2 {
			current_item = int32(call.Argument(1).ToInteger())
		}
		items := parseStringArray(vm, call.Argument(2))
		items_count := int32(call.Argument(3).ToInteger())
		return vm.ToValue(imgui.ComboStrarr(label, &current_item, items, items_count))
	})
	imguiObj.Set("createContext", func(call goja.FunctionCall) goja.Value {
		result := imgui.CreateContext()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalDataTypeApplyFromText", func(call goja.FunctionCall) goja.Value {
		buf := call.Argument(0).String()
		// TODO: 处理类型 DataType
		data_type := imgui.DataType(0)
		p_data := uintptr(call.Argument(2).ToInteger())
		format := call.Argument(3).String()
		return vm.ToValue(imgui.InternalDataTypeApplyFromText(buf, data_type, p_data, format))
	})
	imguiObj.Set("internalDebugDrawCursorPos", func(call goja.FunctionCall) goja.Value {
		imgui.InternalDebugDrawCursorPos()
		return goja.Undefined()
	})
	imguiObj.Set("internalDebugDrawItemRect", func(call goja.FunctionCall) goja.Value {
		imgui.InternalDebugDrawItemRect()
		return goja.Undefined()
	})
	imguiObj.Set("internalDebugDrawLineExtents", func(call goja.FunctionCall) goja.Value {
		imgui.InternalDebugDrawLineExtents()
		return goja.Undefined()
	})
	imguiObj.Set("internalDebugNodeTexture", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *TextureData
		var tex *imgui.TextureData
		int_id := int32(call.Argument(1).ToInteger())
		imgui.InternalDebugNodeTexture(tex, int_id)
		return goja.Undefined()
	})
	imguiObj.Set("destroyContext", func(call goja.FunctionCall) goja.Value {
		imgui.DestroyContext()
		return goja.Undefined()
	})
	imguiObj.Set("internalDockBuilderAddNode", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.InternalDockBuilderAddNode())
	})
	imguiObj.Set("internalDockBuilderRemoveNodeDockedWindows", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 ID
		node_id := imgui.ID(0)
		imgui.InternalDockBuilderRemoveNodeDockedWindows(node_id)
		return goja.Undefined()
	})
	imguiObj.Set("internalDockContextProcessUndockWindow", func(call goja.FunctionCall) goja.Value {
		var ctx *imgui.Context
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.Context); ok {
				ctx = v
			}
		}
		// TODO: 处理指针类型 *Window
		var window *imgui.Window
		imgui.InternalDockContextProcessUndockWindow(ctx, window)
		return goja.Undefined()
	})
	imguiObj.Set("dockSpace", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 ID
		dockspace_id := imgui.ID(0)
		return vm.ToValue(imgui.DockSpace(dockspace_id))
	})
	imguiObj.Set("dockSpaceOverViewport", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.DockSpaceOverViewport())
	})
	imguiObj.Set("dragFloat", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		v := float32(0)
		if len(call.Arguments) >= 2 {
			v = float32(call.Argument(1).ToFloat())
		}
		return vm.ToValue(imgui.DragFloat(label, &v))
	})
	imguiObj.Set("dragFloat2", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		v := [2]float32{0, 0}
		return vm.ToValue(imgui.DragFloat2(label, &v))
	})
	imguiObj.Set("dragFloat3", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		v := [3]float32{0, 0, 0}
		return vm.ToValue(imgui.DragFloat3(label, &v))
	})
	imguiObj.Set("dragFloat4", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		v := [4]float32{0, 0, 0, 0}
		return vm.ToValue(imgui.DragFloat4(label, &v))
	})
	imguiObj.Set("dragFloatRange2", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		v_current_min := float32(0)
		if len(call.Arguments) >= 2 {
			v_current_min = float32(call.Argument(1).ToFloat())
		}
		v_current_max := float32(0)
		if len(call.Arguments) >= 3 {
			v_current_max = float32(call.Argument(2).ToFloat())
		}
		return vm.ToValue(imgui.DragFloatRange2(label, &v_current_min, &v_current_max))
	})
	imguiObj.Set("dragInt", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		v := int32(0)
		if len(call.Arguments) >= 2 {
			v = int32(call.Argument(1).ToInteger())
		}
		return vm.ToValue(imgui.DragInt(label, &v))
	})
	imguiObj.Set("dragInt2", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		v := [2]int32{0, 0}
		return vm.ToValue(imgui.DragInt2(label, &v))
	})
	imguiObj.Set("dragInt3", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		v := [3]int32{0, 0, 0}
		return vm.ToValue(imgui.DragInt3(label, &v))
	})
	imguiObj.Set("dragInt4", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		v := [4]int32{0, 0, 0, 0}
		return vm.ToValue(imgui.DragInt4(label, &v))
	})
	imguiObj.Set("dragIntRange2", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		v_current_min := int32(0)
		if len(call.Arguments) >= 2 {
			v_current_min = int32(call.Argument(1).ToInteger())
		}
		v_current_max := int32(0)
		if len(call.Arguments) >= 3 {
			v_current_max = int32(call.Argument(2).ToInteger())
		}
		return vm.ToValue(imgui.DragIntRange2(label, &v_current_min, &v_current_max))
	})
	imguiObj.Set("dragScalar", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		// TODO: 处理类型 DataType
		data_type := imgui.DataType(0)
		p_data := uintptr(call.Argument(2).ToInteger())
		return vm.ToValue(imgui.DragScalar(label, data_type, p_data))
	})
	imguiObj.Set("dragScalarN", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		// TODO: 处理类型 DataType
		data_type := imgui.DataType(0)
		p_data := uintptr(call.Argument(2).ToInteger())
		components := int32(call.Argument(3).ToInteger())
		return vm.ToValue(imgui.DragScalarN(label, data_type, p_data, components))
	})
	imguiObj.Set("internalFindRenderedTextEnd", func(call goja.FunctionCall) goja.Value {
		text := call.Argument(0).String()
		return vm.ToValue(imgui.InternalFindRenderedTextEnd(text))
	})
	imguiObj.Set("internalFocusWindow", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Window
		var window *imgui.Window
		imgui.InternalFocusWindow(window)
		return goja.Undefined()
	})
	imguiObj.Set("backgroundDrawList", func(call goja.FunctionCall) goja.Value {
		result := imgui.BackgroundDrawList()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("colorU32Col", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 Col
		idx := imgui.Col(0)
		return vm.ToValue(imgui.ColorU32Col(idx))
	})
	imguiObj.Set("colorU32U32", func(call goja.FunctionCall) goja.Value {
		col := uint32(call.Argument(0).ToInteger())
		return vm.ToValue(imgui.ColorU32U32(col))
	})
	imguiObj.Set("columnOffset", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.ColumnOffset())
	})
	imguiObj.Set("columnWidth", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.ColumnWidth())
	})
	imguiObj.Set("foregroundDrawListViewportPtr", func(call goja.FunctionCall) goja.Value {
		result := imgui.ForegroundDrawListViewportPtr()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("mouseDragDelta", func(call goja.FunctionCall) goja.Value {
		result := imgui.MouseDragDelta()
		return vm.ToValue(result)
	})
	imguiObj.Set("internalTypingSelectRequest", func(call goja.FunctionCall) goja.Value {
		result := imgui.InternalTypingSelectRequest()
		if result == nil {
			return goja.Null()
		}
		return vm.ToValue(result)
	})
	imguiObj.Set("internalImFileLoadToMemory", func(call goja.FunctionCall) goja.Value {
		filename := call.Argument(0).String()
		mode := call.Argument(1).String()
		return vm.ToValue(imgui.InternalImFileLoadToMemory(filename, mode))
	})
	imguiObj.Set("internalImFontAtlasPackAddRect", func(call goja.FunctionCall) goja.Value {
		var atlas *imgui.FontAtlas
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.FontAtlas); ok {
				atlas = v
			}
		}
		w := int32(call.Argument(1).ToInteger())
		h := int32(call.Argument(2).ToInteger())
		return vm.ToValue(imgui.InternalImFontAtlasPackAddRect(atlas, w, h))
	})
	imguiObj.Set("internalImFontAtlasTextureGrow", func(call goja.FunctionCall) goja.Value {
		var atlas *imgui.FontAtlas
		if len(call.Arguments) >= 1 {
			if v, ok := call.Argument(0).Export().(*imgui.FontAtlas); ok {
				atlas = v
			}
		}
		imgui.InternalImFontAtlasTextureGrow(atlas)
		return goja.Undefined()
	})
	imguiObj.Set("internalImFontCalcWordWrapPositionEx", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Font
		var font *imgui.Font
		size := float32(call.Argument(1).ToFloat())
		text := call.Argument(2).String()
		wrap_width := float32(call.Argument(3).ToFloat())
		return vm.ToValue(imgui.InternalImFontCalcWordWrapPositionEx(font, size, text, wrap_width))
	})
	imguiObj.Set("internalImHashData", func(call goja.FunctionCall) goja.Value {
		data := uintptr(call.Argument(0).ToInteger())
		data_size := uint64(call.Argument(1).ToInteger())
		return vm.ToValue(imgui.InternalImHashData(data, data_size))
	})
	imguiObj.Set("internalImHashStr", func(call goja.FunctionCall) goja.Value {
		data := call.Argument(0).String()
		return vm.ToValue(imgui.InternalImHashStr(data))
	})
	imguiObj.Set("internalImTextCalcWordWrapNextLineStart", func(call goja.FunctionCall) goja.Value {
		text := call.Argument(0).String()
		return vm.ToValue(imgui.InternalImTextCalcWordWrapNextLineStart(text))
	})
	imguiObj.Set("internalImTextStrFromUtf8", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Wchar
		var out_buf *imgui.Wchar
		out_buf_size := int32(call.Argument(1).ToInteger())
		in_text := call.Argument(2).String()
		in_text_end := call.Argument(3).String()
		return vm.ToValue(imgui.InternalImTextStrFromUtf8(out_buf, out_buf_size, in_text, in_text_end))
	})
	imguiObj.Set("image", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 TextureRef
		tex_ref := imgui.TextureRef{}
		image_size := parseVec2(vm, call.Argument(1))
		imgui.Image(tex_ref, image_size)
		return goja.Undefined()
	})
	imguiObj.Set("imageButton", func(call goja.FunctionCall) goja.Value {
		str_id := call.Argument(0).String()
		// TODO: 处理类型 TextureRef
		tex_ref := imgui.TextureRef{}
		image_size := parseVec2(vm, call.Argument(2))
		return vm.ToValue(imgui.ImageButton(str_id, tex_ref, image_size))
	})
	imguiObj.Set("internalImageButtonEx", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 ID
		id := imgui.ID(0)
		// TODO: 处理类型 TextureRef
		tex_ref := imgui.TextureRef{}
		image_size := parseVec2(vm, call.Argument(2))
		uv0 := parseVec2(vm, call.Argument(2))
		uv1 := parseVec2(vm, call.Argument(2))
		bg_col := parseVec4(vm, call.Argument(2))
		tint_col := parseVec4(vm, call.Argument(2))
		return vm.ToValue(imgui.InternalImageButtonEx(id, tex_ref, image_size, uv0, uv1, bg_col, tint_col))
	})
	imguiObj.Set("imageWithBg", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 TextureRef
		tex_ref := imgui.TextureRef{}
		image_size := parseVec2(vm, call.Argument(1))
		imgui.ImageWithBg(tex_ref, image_size)
		return goja.Undefined()
	})
	imguiObj.Set("indent", func(call goja.FunctionCall) goja.Value {
		imgui.Indent()
		return goja.Undefined()
	})
	imguiObj.Set("inputDouble", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		v := float64(0)
		if len(call.Arguments) >= 2 {
			v = float64(call.Argument(1).ToFloat())
		}
		return vm.ToValue(imgui.InputDouble(label, &v))
	})
	imguiObj.Set("inputFloat", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		v := float32(0)
		if len(call.Arguments) >= 2 {
			v = float32(call.Argument(1).ToFloat())
		}
		return vm.ToValue(imgui.InputFloat(label, &v))
	})
	imguiObj.Set("inputFloat2", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		v := [2]float32{0, 0}
		return vm.ToValue(imgui.InputFloat2(label, &v))
	})
	imguiObj.Set("inputFloat3", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		v := [3]float32{0, 0, 0}
		return vm.ToValue(imgui.InputFloat3(label, &v))
	})
	imguiObj.Set("inputFloat4", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		v := [4]float32{0, 0, 0, 0}
		return vm.ToValue(imgui.InputFloat4(label, &v))
	})
	imguiObj.Set("inputInt", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		v := int32(0)
		if len(call.Arguments) >= 2 {
			v = int32(call.Argument(1).ToInteger())
		}
		return vm.ToValue(imgui.InputInt(label, &v))
	})
	imguiObj.Set("inputInt2", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		v := [2]int32{0, 0}
		return vm.ToValue(imgui.InputInt2(label, &v))
	})
	imguiObj.Set("inputInt3", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		v := [3]int32{0, 0, 0}
		return vm.ToValue(imgui.InputInt3(label, &v))
	})
	imguiObj.Set("inputInt4", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		v := [4]int32{0, 0, 0, 0}
		return vm.ToValue(imgui.InputInt4(label, &v))
	})
	imguiObj.Set("inputScalar", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		// TODO: 处理类型 DataType
		data_type := imgui.DataType(0)
		p_data := uintptr(call.Argument(2).ToInteger())
		return vm.ToValue(imgui.InputScalar(label, data_type, p_data))
	})
	imguiObj.Set("inputScalarN", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		// TODO: 处理类型 DataType
		data_type := imgui.DataType(0)
		p_data := uintptr(call.Argument(2).ToInteger())
		components := int32(call.Argument(3).ToInteger())
		return vm.ToValue(imgui.InputScalarN(label, data_type, p_data, components))
	})
	imguiObj.Set("internalInputTextEx", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		hint := call.Argument(1).String()
		buf := call.Argument(2).String()
		buf_size := int32(call.Argument(3).ToInteger())
		size_arg := parseVec2(vm, call.Argument(4))
		// TODO: 处理类型 InputTextFlags
		flags := imgui.InputTextFlags(0)
		return vm.ToValue(imgui.InternalInputTextEx(label, hint, buf, buf_size, size_arg, flags))
	})
	imguiObj.Set("invisibleButton", func(call goja.FunctionCall) goja.Value {
		str_id := call.Argument(0).String()
		size := parseVec2(vm, call.Argument(1))
		return vm.ToValue(imgui.InvisibleButton(str_id, size))
	})
	imguiObj.Set("isItemClicked", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.IsItemClicked())
	})
	imguiObj.Set("isItemHovered", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.IsItemHovered())
	})
	imguiObj.Set("internalIsKeyChordPressedInputFlags", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 KeyChord
		key_chord := imgui.KeyChord(0)
		// TODO: 处理类型 InputFlags
		flags := imgui.InputFlags(0)
		return vm.ToValue(imgui.InternalIsKeyChordPressedInputFlags(key_chord, flags))
	})
	imguiObj.Set("isKeyPressedBool", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 Key
		key := imgui.Key(0)
		return vm.ToValue(imgui.IsKeyPressedBool(key))
	})
	imguiObj.Set("internalIsKeyPressedInputFlags", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 Key
		key := imgui.Key(0)
		// TODO: 处理类型 InputFlags
		flags := imgui.InputFlags(0)
		return vm.ToValue(imgui.InternalIsKeyPressedInputFlags(key, flags))
	})
	imguiObj.Set("isMouseClickedBool", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 MouseButton
		button := imgui.MouseButton(0)
		return vm.ToValue(imgui.IsMouseClickedBool(button))
	})
	imguiObj.Set("internalIsMouseClickedInputFlags", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 MouseButton
		button := imgui.MouseButton(0)
		// TODO: 处理类型 InputFlags
		flags := imgui.InputFlags(0)
		return vm.ToValue(imgui.InternalIsMouseClickedInputFlags(button, flags))
	})
	imguiObj.Set("internalIsMouseDragPastThreshold", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 MouseButton
		button := imgui.MouseButton(0)
		return vm.ToValue(imgui.InternalIsMouseDragPastThreshold(button))
	})
	imguiObj.Set("isMouseDragging", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 MouseButton
		button := imgui.MouseButton(0)
		return vm.ToValue(imgui.IsMouseDragging(button))
	})
	imguiObj.Set("isMouseHoveringRect", func(call goja.FunctionCall) goja.Value {
		r_min := parseVec2(vm, call.Argument(0))
		r_max := parseVec2(vm, call.Argument(0))
		return vm.ToValue(imgui.IsMouseHoveringRect(r_min, r_max))
	})
	imguiObj.Set("isMousePosValid", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.IsMousePosValid())
	})
	imguiObj.Set("isPopupOpenStr", func(call goja.FunctionCall) goja.Value {
		str_id := call.Argument(0).String()
		return vm.ToValue(imgui.IsPopupOpenStr(str_id))
	})
	imguiObj.Set("internalIsWindowContentHoverable", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Window
		var window *imgui.Window
		return vm.ToValue(imgui.InternalIsWindowContentHoverable(window))
	})
	imguiObj.Set("isWindowFocused", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.IsWindowFocused())
	})
	imguiObj.Set("isWindowHovered", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.IsWindowHovered())
	})
	imguiObj.Set("internalItemAdd", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 Rect
		bb := imgui.Rect{}
		// TODO: 处理类型 ID
		id := imgui.ID(0)
		return vm.ToValue(imgui.InternalItemAdd(bb, id))
	})
	imguiObj.Set("internalItemSizeRect", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 Rect
		bb := imgui.Rect{}
		imgui.InternalItemSizeRect(bb)
		return goja.Undefined()
	})
	imguiObj.Set("internalItemSizeVec2", func(call goja.FunctionCall) goja.Value {
		size := parseVec2(vm, call.Argument(0))
		imgui.InternalItemSizeVec2(size)
		return goja.Undefined()
	})
	imguiObj.Set("listBoxStrarr", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		current_item := int32(0)
		if len(call.Arguments) >= 2 {
			current_item = int32(call.Argument(1).ToInteger())
		}
		items := parseStringArray(vm, call.Argument(2))
		items_count := int32(call.Argument(3).ToInteger())
		return vm.ToValue(imgui.ListBoxStrarr(label, &current_item, items, items_count))
	})
	imguiObj.Set("loadIniSettingsFromMemory", func(call goja.FunctionCall) goja.Value {
		ini_data := call.Argument(0).String()
		imgui.LoadIniSettingsFromMemory(ini_data)
		return goja.Undefined()
	})
	imguiObj.Set("internalLogRenderedText", func(call goja.FunctionCall) goja.Value {
		ref_pos := parseVec2(vm, call.Argument(0))
		text := call.Argument(1).String()
		imgui.InternalLogRenderedText(&ref_pos, text)
		return goja.Undefined()
	})
	imguiObj.Set("internalLogToBuffer", func(call goja.FunctionCall) goja.Value {
		imgui.InternalLogToBuffer()
		return goja.Undefined()
	})
	imguiObj.Set("logToClipboard", func(call goja.FunctionCall) goja.Value {
		imgui.LogToClipboard()
		return goja.Undefined()
	})
	imguiObj.Set("logToFile", func(call goja.FunctionCall) goja.Value {
		imgui.LogToFile()
		return goja.Undefined()
	})
	imguiObj.Set("logToTTY", func(call goja.FunctionCall) goja.Value {
		imgui.LogToTTY()
		return goja.Undefined()
	})
	imguiObj.Set("internalMenuItemEx", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		icon := call.Argument(1).String()
		return vm.ToValue(imgui.InternalMenuItemEx(label, icon))
	})
	imguiObj.Set("menuItemBool", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		return vm.ToValue(imgui.MenuItemBool(label))
	})
	imguiObj.Set("menuItemBoolPtr", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		shortcut := call.Argument(1).String()
		p_selected := false
		if len(call.Arguments) >= 3 {
			p_selected = call.Argument(2).ToBoolean()
		}
		return vm.ToValue(imgui.MenuItemBoolPtr(label, shortcut, &p_selected))
	})
	imguiObj.Set("internalOpenPopupEx", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 ID
		id := imgui.ID(0)
		imgui.InternalOpenPopupEx(id)
		return goja.Undefined()
	})
	imguiObj.Set("openPopupOnItemClick", func(call goja.FunctionCall) goja.Value {
		imgui.OpenPopupOnItemClick()
		return goja.Undefined()
	})
	imguiObj.Set("openPopupID", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 ID
		id := imgui.ID(0)
		imgui.OpenPopupID(id)
		return goja.Undefined()
	})
	imguiObj.Set("openPopupStr", func(call goja.FunctionCall) goja.Value {
		str_id := call.Argument(0).String()
		imgui.OpenPopupStr(str_id)
		return goja.Undefined()
	})
	imguiObj.Set("plotHistogramFloatPtr", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		values := float32(0)
		if len(call.Arguments) >= 2 {
			values = float32(call.Argument(1).ToFloat())
		}
		values_count := int32(call.Argument(2).ToInteger())
		imgui.PlotHistogramFloatPtr(label, &values, values_count)
		return goja.Undefined()
	})
	imguiObj.Set("plotLinesFloatPtr", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		values := float32(0)
		if len(call.Arguments) >= 2 {
			values = float32(call.Argument(1).ToFloat())
		}
		values_count := int32(call.Argument(2).ToInteger())
		imgui.PlotLinesFloatPtr(label, &values, values_count)
		return goja.Undefined()
	})
	imguiObj.Set("popStyleColor", func(call goja.FunctionCall) goja.Value {
		imgui.PopStyleColor()
		return goja.Undefined()
	})
	imguiObj.Set("popStyleVar", func(call goja.FunctionCall) goja.Value {
		imgui.PopStyleVar()
		return goja.Undefined()
	})
	imguiObj.Set("progressBar", func(call goja.FunctionCall) goja.Value {
		fraction := float32(call.Argument(0).ToFloat())
		imgui.ProgressBar(fraction)
		return goja.Undefined()
	})
	imguiObj.Set("pushTextWrapPos", func(call goja.FunctionCall) goja.Value {
		imgui.PushTextWrapPos()
		return goja.Undefined()
	})
	imguiObj.Set("internalRenderArrow", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *DrawList
		var draw_list *imgui.DrawList
		pos := parseVec2(vm, call.Argument(1))
		col := uint32(call.Argument(1).ToInteger())
		// TODO: 处理类型 Dir
		dir := imgui.Dir(0)
		imgui.InternalRenderArrow(draw_list, pos, col, dir)
		return goja.Undefined()
	})
	imguiObj.Set("internalRenderColorRectWithAlphaCheckerboard", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *DrawList
		var draw_list *imgui.DrawList
		p_min := parseVec2(vm, call.Argument(1))
		p_max := parseVec2(vm, call.Argument(1))
		fill_col := uint32(call.Argument(1).ToInteger())
		grid_step := float32(call.Argument(2).ToFloat())
		grid_off := parseVec2(vm, call.Argument(3))
		imgui.InternalRenderColorRectWithAlphaCheckerboard(draw_list, p_min, p_max, fill_col, grid_step, grid_off)
		return goja.Undefined()
	})
	imguiObj.Set("internalRenderFrame", func(call goja.FunctionCall) goja.Value {
		p_min := parseVec2(vm, call.Argument(0))
		p_max := parseVec2(vm, call.Argument(0))
		fill_col := uint32(call.Argument(0).ToInteger())
		imgui.InternalRenderFrame(p_min, p_max, fill_col)
		return goja.Undefined()
	})
	imguiObj.Set("internalRenderFrameBorder", func(call goja.FunctionCall) goja.Value {
		p_min := parseVec2(vm, call.Argument(0))
		p_max := parseVec2(vm, call.Argument(0))
		imgui.InternalRenderFrameBorder(p_min, p_max)
		return goja.Undefined()
	})
	imguiObj.Set("internalRenderNavCursor", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 Rect
		bb := imgui.Rect{}
		// TODO: 处理类型 ID
		id := imgui.ID(0)
		imgui.InternalRenderNavCursor(bb, id)
		return goja.Undefined()
	})
	imguiObj.Set("renderPlatformWindowsDefault", func(call goja.FunctionCall) goja.Value {
		imgui.RenderPlatformWindowsDefault()
		return goja.Undefined()
	})
	imguiObj.Set("internalRenderText", func(call goja.FunctionCall) goja.Value {
		pos := parseVec2(vm, call.Argument(0))
		text := call.Argument(0).String()
		imgui.InternalRenderText(pos, text)
		return goja.Undefined()
	})
	imguiObj.Set("internalRenderTextClipped", func(call goja.FunctionCall) goja.Value {
		pos_min := parseVec2(vm, call.Argument(0))
		pos_max := parseVec2(vm, call.Argument(0))
		text := call.Argument(0).String()
		text_size_if_known := parseVec2(vm, call.Argument(1))
		imgui.InternalRenderTextClipped(pos_min, pos_max, text, &text_size_if_known)
		return goja.Undefined()
	})
	imguiObj.Set("internalRenderTextClippedEx", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *DrawList
		var draw_list *imgui.DrawList
		pos_min := parseVec2(vm, call.Argument(1))
		pos_max := parseVec2(vm, call.Argument(1))
		text := call.Argument(1).String()
		text_size_if_known := parseVec2(vm, call.Argument(2))
		imgui.InternalRenderTextClippedEx(draw_list, pos_min, pos_max, text, &text_size_if_known)
		return goja.Undefined()
	})
	imguiObj.Set("resetMouseDragDelta", func(call goja.FunctionCall) goja.Value {
		imgui.ResetMouseDragDelta()
		return goja.Undefined()
	})
	imguiObj.Set("sameLine", func(call goja.FunctionCall) goja.Value {
		imgui.SameLine()
		return goja.Undefined()
	})
	imguiObj.Set("saveIniSettingsToMemory", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.SaveIniSettingsToMemory())
	})
	imguiObj.Set("internalScrollToItem", func(call goja.FunctionCall) goja.Value {
		imgui.InternalScrollToItem()
		return goja.Undefined()
	})
	imguiObj.Set("internalScrollToRect", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Window
		var window *imgui.Window
		// TODO: 处理类型 Rect
		rect := imgui.Rect{}
		imgui.InternalScrollToRect(window, rect)
		return goja.Undefined()
	})
	imguiObj.Set("internalScrollToRectEx", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Window
		var window *imgui.Window
		// TODO: 处理类型 Rect
		rect := imgui.Rect{}
		result := imgui.InternalScrollToRectEx(window, rect)
		return vm.ToValue(result)
	})
	imguiObj.Set("internalScrollbarEx", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 Rect
		bb := imgui.Rect{}
		// TODO: 处理类型 ID
		id := imgui.ID(0)
		// TODO: 处理类型 Axis
		axis := imgui.Axis(0)
		// TODO: 处理指针类型 *int64
		var p_scroll_v *int64
		// TODO: 处理类型 int64
		avail_v := int64(0)
		// TODO: 处理类型 int64
		contents_v := int64(0)
		return vm.ToValue(imgui.InternalScrollbarEx(bb, id, axis, p_scroll_v, avail_v, contents_v))
	})
	imguiObj.Set("selectableBool", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		return vm.ToValue(imgui.SelectableBool(label))
	})
	imguiObj.Set("selectableBoolPtr", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		p_selected := false
		if len(call.Arguments) >= 2 {
			p_selected = call.Argument(1).ToBoolean()
		}
		return vm.ToValue(imgui.SelectableBoolPtr(label, &p_selected))
	})
	imguiObj.Set("internalSeparatorEx", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 SeparatorFlags
		flags := imgui.SeparatorFlags(0)
		imgui.InternalSeparatorEx(flags)
		return goja.Undefined()
	})
	imguiObj.Set("setAllocatorFunctions", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理函数类型 MemAllocFunc
		var alloc_func imgui.MemAllocFunc
		// TODO: 处理函数类型 MemFreeFunc
		var free_func imgui.MemFreeFunc
		imgui.SetAllocatorFunctions(alloc_func, free_func)
		return goja.Undefined()
	})
	imguiObj.Set("setDragDropPayload", func(call goja.FunctionCall) goja.Value {
		typeArg := call.Argument(0).String()
		data := uintptr(call.Argument(1).ToInteger())
		sz := uint64(call.Argument(2).ToInteger())
		return vm.ToValue(imgui.SetDragDropPayload(typeArg, data, sz))
	})
	imguiObj.Set("internalSetKeyOwner", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 Key
		key := imgui.Key(0)
		// TODO: 处理类型 ID
		owner_id := imgui.ID(0)
		imgui.InternalSetKeyOwner(key, owner_id)
		return goja.Undefined()
	})
	imguiObj.Set("internalSetKeyOwnersForKeyChord", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 KeyChord
		key := imgui.KeyChord(0)
		// TODO: 处理类型 ID
		owner_id := imgui.ID(0)
		imgui.InternalSetKeyOwnersForKeyChord(key, owner_id)
		return goja.Undefined()
	})
	imguiObj.Set("setKeyboardFocusHere", func(call goja.FunctionCall) goja.Value {
		imgui.SetKeyboardFocusHere()
		return goja.Undefined()
	})
	imguiObj.Set("setNextItemOpen", func(call goja.FunctionCall) goja.Value {
		is_open := call.Argument(0).ToBoolean()
		imgui.SetNextItemOpen(is_open)
		return goja.Undefined()
	})
	imguiObj.Set("setNextItemShortcut", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 KeyChord
		key_chord := imgui.KeyChord(0)
		imgui.SetNextItemShortcut(key_chord)
		return goja.Undefined()
	})
	imguiObj.Set("setNextWindowCollapsed", func(call goja.FunctionCall) goja.Value {
		collapsed := call.Argument(0).ToBoolean()
		imgui.SetNextWindowCollapsed(collapsed)
		return goja.Undefined()
	})
	imguiObj.Set("setNextWindowDockID", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 ID
		dock_id := imgui.ID(0)
		imgui.SetNextWindowDockID(dock_id)
		return goja.Undefined()
	})
	imguiObj.Set("setNextWindowPos", func(call goja.FunctionCall) goja.Value {
		pos := parseVec2(vm, call.Argument(0))
		imgui.SetNextWindowPos(pos)
		return goja.Undefined()
	})
	imguiObj.Set("setNextWindowSize", func(call goja.FunctionCall) goja.Value {
		size := parseVec2(vm, call.Argument(0))
		imgui.SetNextWindowSize(size)
		return goja.Undefined()
	})
	imguiObj.Set("setNextWindowSizeConstraints", func(call goja.FunctionCall) goja.Value {
		size_min := parseVec2(vm, call.Argument(0))
		size_max := parseVec2(vm, call.Argument(0))
		imgui.SetNextWindowSizeConstraints(size_min, size_max)
		return goja.Undefined()
	})
	imguiObj.Set("setScrollFromPosXFloat", func(call goja.FunctionCall) goja.Value {
		local_x := float32(call.Argument(0).ToFloat())
		imgui.SetScrollFromPosXFloat(local_x)
		return goja.Undefined()
	})
	imguiObj.Set("setScrollFromPosYFloat", func(call goja.FunctionCall) goja.Value {
		local_y := float32(call.Argument(0).ToFloat())
		imgui.SetScrollFromPosYFloat(local_y)
		return goja.Undefined()
	})
	imguiObj.Set("setScrollHereX", func(call goja.FunctionCall) goja.Value {
		imgui.SetScrollHereX()
		return goja.Undefined()
	})
	imguiObj.Set("setScrollHereY", func(call goja.FunctionCall) goja.Value {
		imgui.SetScrollHereY()
		return goja.Undefined()
	})
	imguiObj.Set("setWindowCollapsedBool", func(call goja.FunctionCall) goja.Value {
		collapsed := call.Argument(0).ToBoolean()
		imgui.SetWindowCollapsedBool(collapsed)
		return goja.Undefined()
	})
	imguiObj.Set("setWindowCollapsedStr", func(call goja.FunctionCall) goja.Value {
		name := call.Argument(0).String()
		collapsed := call.Argument(1).ToBoolean()
		imgui.SetWindowCollapsedStr(name, collapsed)
		return goja.Undefined()
	})
	imguiObj.Set("internalSetWindowCollapsedWindowPtr", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Window
		var window *imgui.Window
		collapsed := call.Argument(1).ToBoolean()
		imgui.InternalSetWindowCollapsedWindowPtr(window, collapsed)
		return goja.Undefined()
	})
	imguiObj.Set("setWindowPosStr", func(call goja.FunctionCall) goja.Value {
		name := call.Argument(0).String()
		pos := parseVec2(vm, call.Argument(1))
		imgui.SetWindowPosStr(name, pos)
		return goja.Undefined()
	})
	imguiObj.Set("setWindowPosVec2", func(call goja.FunctionCall) goja.Value {
		pos := parseVec2(vm, call.Argument(0))
		imgui.SetWindowPosVec2(pos)
		return goja.Undefined()
	})
	imguiObj.Set("internalSetWindowPosWindowPtr", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Window
		var window *imgui.Window
		pos := parseVec2(vm, call.Argument(1))
		imgui.InternalSetWindowPosWindowPtr(window, pos)
		return goja.Undefined()
	})
	imguiObj.Set("setWindowSizeStr", func(call goja.FunctionCall) goja.Value {
		name := call.Argument(0).String()
		size := parseVec2(vm, call.Argument(1))
		imgui.SetWindowSizeStr(name, size)
		return goja.Undefined()
	})
	imguiObj.Set("setWindowSizeVec2", func(call goja.FunctionCall) goja.Value {
		size := parseVec2(vm, call.Argument(0))
		imgui.SetWindowSizeVec2(size)
		return goja.Undefined()
	})
	imguiObj.Set("internalSetWindowSizeWindowPtr", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Window
		var window *imgui.Window
		size := parseVec2(vm, call.Argument(1))
		imgui.InternalSetWindowSizeWindowPtr(window, size)
		return goja.Undefined()
	})
	imguiObj.Set("shortcut", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 KeyChord
		key_chord := imgui.KeyChord(0)
		return vm.ToValue(imgui.Shortcut(key_chord))
	})
	imguiObj.Set("showAboutWindow", func(call goja.FunctionCall) goja.Value {
		imgui.ShowAboutWindow()
		return goja.Undefined()
	})
	imguiObj.Set("showDebugLogWindow", func(call goja.FunctionCall) goja.Value {
		imgui.ShowDebugLogWindow()
		return goja.Undefined()
	})
	imguiObj.Set("showDemoWindow", func(call goja.FunctionCall) goja.Value {
		imgui.ShowDemoWindow()
		return goja.Undefined()
	})
	imguiObj.Set("showIDStackToolWindow", func(call goja.FunctionCall) goja.Value {
		imgui.ShowIDStackToolWindow()
		return goja.Undefined()
	})
	imguiObj.Set("showMetricsWindow", func(call goja.FunctionCall) goja.Value {
		imgui.ShowMetricsWindow()
		return goja.Undefined()
	})
	imguiObj.Set("showStyleEditor", func(call goja.FunctionCall) goja.Value {
		imgui.ShowStyleEditor()
		return goja.Undefined()
	})
	imguiObj.Set("sliderAngle", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		v_rad := float32(0)
		if len(call.Arguments) >= 2 {
			v_rad = float32(call.Argument(1).ToFloat())
		}
		return vm.ToValue(imgui.SliderAngle(label, &v_rad))
	})
	imguiObj.Set("sliderFloat", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		v := float32(0)
		if len(call.Arguments) >= 2 {
			v = float32(call.Argument(1).ToFloat())
		}
		v_min := float32(call.Argument(2).ToFloat())
		v_max := float32(call.Argument(3).ToFloat())
		return vm.ToValue(imgui.SliderFloat(label, &v, v_min, v_max))
	})
	imguiObj.Set("sliderFloat2", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		v := [2]float32{0, 0}
		v_min := float32(call.Argument(2).ToFloat())
		v_max := float32(call.Argument(3).ToFloat())
		return vm.ToValue(imgui.SliderFloat2(label, &v, v_min, v_max))
	})
	imguiObj.Set("sliderFloat3", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		v := [3]float32{0, 0, 0}
		v_min := float32(call.Argument(2).ToFloat())
		v_max := float32(call.Argument(3).ToFloat())
		return vm.ToValue(imgui.SliderFloat3(label, &v, v_min, v_max))
	})
	imguiObj.Set("sliderFloat4", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		v := [4]float32{0, 0, 0, 0}
		v_min := float32(call.Argument(2).ToFloat())
		v_max := float32(call.Argument(3).ToFloat())
		return vm.ToValue(imgui.SliderFloat4(label, &v, v_min, v_max))
	})
	imguiObj.Set("sliderInt", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		v := int32(0)
		if len(call.Arguments) >= 2 {
			v = int32(call.Argument(1).ToInteger())
		}
		v_min := int32(call.Argument(2).ToInteger())
		v_max := int32(call.Argument(3).ToInteger())
		return vm.ToValue(imgui.SliderInt(label, &v, v_min, v_max))
	})
	imguiObj.Set("sliderInt2", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		v := [2]int32{0, 0}
		v_min := int32(call.Argument(2).ToInteger())
		v_max := int32(call.Argument(3).ToInteger())
		return vm.ToValue(imgui.SliderInt2(label, &v, v_min, v_max))
	})
	imguiObj.Set("sliderInt3", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		v := [3]int32{0, 0, 0}
		v_min := int32(call.Argument(2).ToInteger())
		v_max := int32(call.Argument(3).ToInteger())
		return vm.ToValue(imgui.SliderInt3(label, &v, v_min, v_max))
	})
	imguiObj.Set("sliderInt4", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		v := [4]int32{0, 0, 0, 0}
		v_min := int32(call.Argument(2).ToInteger())
		v_max := int32(call.Argument(3).ToInteger())
		return vm.ToValue(imgui.SliderInt4(label, &v, v_min, v_max))
	})
	imguiObj.Set("sliderScalar", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		// TODO: 处理类型 DataType
		data_type := imgui.DataType(0)
		p_data := uintptr(call.Argument(2).ToInteger())
		p_min := uintptr(call.Argument(3).ToInteger())
		p_max := uintptr(call.Argument(4).ToInteger())
		return vm.ToValue(imgui.SliderScalar(label, data_type, p_data, p_min, p_max))
	})
	imguiObj.Set("sliderScalarN", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		// TODO: 处理类型 DataType
		data_type := imgui.DataType(0)
		p_data := uintptr(call.Argument(2).ToInteger())
		components := int32(call.Argument(3).ToInteger())
		p_min := uintptr(call.Argument(4).ToInteger())
		p_max := uintptr(call.Argument(5).ToInteger())
		return vm.ToValue(imgui.SliderScalarN(label, data_type, p_data, components, p_min, p_max))
	})
	imguiObj.Set("internalSplitterBehavior", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 Rect
		bb := imgui.Rect{}
		// TODO: 处理类型 ID
		id := imgui.ID(0)
		// TODO: 处理类型 Axis
		axis := imgui.Axis(0)
		size1 := float32(0)
		if len(call.Arguments) >= 4 {
			size1 = float32(call.Argument(3).ToFloat())
		}
		size2 := float32(0)
		if len(call.Arguments) >= 5 {
			size2 = float32(call.Argument(4).ToFloat())
		}
		min_size1 := float32(call.Argument(5).ToFloat())
		min_size2 := float32(call.Argument(6).ToFloat())
		return vm.ToValue(imgui.InternalSplitterBehavior(bb, id, axis, &size1, &size2, min_size1, min_size2))
	})
	imguiObj.Set("styleColorsClassic", func(call goja.FunctionCall) goja.Value {
		imgui.StyleColorsClassic()
		return goja.Undefined()
	})
	imguiObj.Set("styleColorsDark", func(call goja.FunctionCall) goja.Value {
		imgui.StyleColorsDark()
		return goja.Undefined()
	})
	imguiObj.Set("styleColorsLight", func(call goja.FunctionCall) goja.Value {
		imgui.StyleColorsLight()
		return goja.Undefined()
	})
	imguiObj.Set("tabItemButton", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		return vm.ToValue(imgui.TabItemButton(label))
	})
	imguiObj.Set("tableGetColumnFlags", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.TableGetColumnFlags())
	})
	imguiObj.Set("tableGetColumnNameInt", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.TableGetColumnNameInt())
	})
	imguiObj.Set("internalTableGetColumnResizeID", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理指针类型 *Table
		var table *imgui.Table
		column_n := int32(call.Argument(1).ToInteger())
		return vm.ToValue(imgui.InternalTableGetColumnResizeID(table, column_n))
	})
	imguiObj.Set("tableNextRow", func(call goja.FunctionCall) goja.Value {
		imgui.TableNextRow()
		return goja.Undefined()
	})
	imguiObj.Set("internalTableOpenContextMenu", func(call goja.FunctionCall) goja.Value {
		imgui.InternalTableOpenContextMenu()
		return goja.Undefined()
	})
	imguiObj.Set("tableSetBgColor", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 TableBgTarget
		target := imgui.TableBgTarget(0)
		color := uint32(call.Argument(1).ToInteger())
		imgui.TableSetBgColor(target, color)
		return goja.Undefined()
	})
	imguiObj.Set("tableSetupColumn", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		imgui.TableSetupColumn(label)
		return goja.Undefined()
	})
	imguiObj.Set("internalTempInputScalar", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 Rect
		bb := imgui.Rect{}
		// TODO: 处理类型 ID
		id := imgui.ID(0)
		label := call.Argument(2).String()
		// TODO: 处理类型 DataType
		data_type := imgui.DataType(0)
		p_data := uintptr(call.Argument(4).ToInteger())
		format := call.Argument(5).String()
		return vm.ToValue(imgui.InternalTempInputScalar(bb, id, label, data_type, p_data, format))
	})
	imguiObj.Set("internalTextEx", func(call goja.FunctionCall) goja.Value {
		text := call.Argument(0).String()
		imgui.InternalTextEx(text)
		return goja.Undefined()
	})
	imguiObj.Set("textLinkOpenURL", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		return vm.ToValue(imgui.TextLinkOpenURL(label))
	})
	imguiObj.Set("textUnformatted", func(call goja.FunctionCall) goja.Value {
		text := call.Argument(0).String()
		imgui.TextUnformatted(text)
		return goja.Undefined()
	})
	imguiObj.Set("internalTreeNodeBehavior", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 ID
		id := imgui.ID(0)
		// TODO: 处理类型 TreeNodeFlags
		flags := imgui.TreeNodeFlags(0)
		label := call.Argument(2).String()
		return vm.ToValue(imgui.InternalTreeNodeBehavior(id, flags, label))
	})
	imguiObj.Set("treeNodeExStr", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		return vm.ToValue(imgui.TreeNodeExStr(label))
	})
	imguiObj.Set("unindent", func(call goja.FunctionCall) goja.Value {
		imgui.Unindent()
		return goja.Undefined()
	})
	imguiObj.Set("vSliderFloat", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		size := parseVec2(vm, call.Argument(1))
		v := float32(0)
		if len(call.Arguments) >= 2 {
			v = float32(call.Argument(1).ToFloat())
		}
		v_min := float32(call.Argument(2).ToFloat())
		v_max := float32(call.Argument(3).ToFloat())
		return vm.ToValue(imgui.VSliderFloat(label, size, &v, v_min, v_max))
	})
	imguiObj.Set("vSliderInt", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		size := parseVec2(vm, call.Argument(1))
		v := int32(0)
		if len(call.Arguments) >= 2 {
			v = int32(call.Argument(1).ToInteger())
		}
		v_min := int32(call.Argument(2).ToInteger())
		v_max := int32(call.Argument(3).ToInteger())
		return vm.ToValue(imgui.VSliderInt(label, size, &v, v_min, v_max))
	})
	imguiObj.Set("vSliderScalar", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		size := parseVec2(vm, call.Argument(1))
		// TODO: 处理类型 DataType
		data_type := imgui.DataType(0)
		p_data := uintptr(call.Argument(2).ToInteger())
		p_min := uintptr(call.Argument(3).ToInteger())
		p_max := uintptr(call.Argument(4).ToInteger())
		return vm.ToValue(imgui.VSliderScalar(label, size, data_type, p_data, p_min, p_max))
	})
	imguiObj.Set("valueFloat", func(call goja.FunctionCall) goja.Value {
		prefix := call.Argument(0).String()
		v := float32(call.Argument(1).ToFloat())
		imgui.ValueFloat(prefix, v)
		return goja.Undefined()
	})

	// ========== input_text.go ==========
	imguiObj.Set("inputTextWithHint", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		hint := call.Argument(1).String()
		// TODO: 处理指针类型 *string
		var buf *string
		// TODO: 处理类型 InputTextFlags
		flags := imgui.InputTextFlags(0)
		// TODO: 处理类型 InputTextCallback
		var callback imgui.InputTextCallback
		if arg := call.Argument(4); !goja.IsUndefined(arg) && !goja.IsNull(arg) {
			if fn, ok := goja.AssertFunction(arg); ok {
				callback = func(data imgui.InputTextCallbackData) int {
					ret, err := fn(goja.Undefined(), vm.ToValue(data))
					if err != nil {
						return 0
					}
					return int(ret.ToInteger())
				}
			}
		}
		return vm.ToValue(imgui.InputTextWithHint(label, hint, buf, flags, callback))
	})
	imguiObj.Set("inputTextMultiline", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		// TODO: 处理指针类型 *string
		var buf *string
		size := parseVec2(vm, call.Argument(2))
		// TODO: 处理类型 InputTextFlags
		flags := imgui.InputTextFlags(0)
		// TODO: 处理类型 InputTextCallback
		var callback imgui.InputTextCallback
		if arg := call.Argument(4); !goja.IsUndefined(arg) && !goja.IsNull(arg) {
			if fn, ok := goja.AssertFunction(arg); ok {
				callback = func(data imgui.InputTextCallbackData) int {
					ret, err := fn(goja.Undefined(), vm.ToValue(data))
					if err != nil {
						return 0
					}
					return int(ret.ToInteger())
				}
			}
		}
		return vm.ToValue(imgui.InputTextMultiline(label, buf, size, flags, callback))
	})

	// ========== extra_types.go ==========
	imguiObj.Set("newVec2", func(call goja.FunctionCall) goja.Value {
		x := float32(call.Argument(0).ToFloat())
		y := float32(call.Argument(1).ToFloat())
		result := imgui.NewVec2(x, y)
		return vm.ToValue(result)
	})
	imguiObj.Set("newVec4", func(call goja.FunctionCall) goja.Value {
		r := float32(call.Argument(0).ToFloat())
		g := float32(call.Argument(1).ToFloat())
		b := float32(call.Argument(2).ToFloat())
		a := float32(call.Argument(3).ToFloat())
		result := imgui.NewVec4(r, g, b, a)
		return vm.ToValue(result)
	})
	imguiObj.Set("newColor", func(call goja.FunctionCall) goja.Value {
		r := float32(call.Argument(0).ToFloat())
		g := float32(call.Argument(1).ToFloat())
		b := float32(call.Argument(2).ToFloat())
		a := float32(call.Argument(3).ToFloat())
		return vm.ToValue(imgui.NewColor(r, g, b, a))
	})
	imguiObj.Set("newColorFromPacked", func(call goja.FunctionCall) goja.Value {
		v := uint32(call.Argument(0).ToInteger())
		return vm.ToValue(imgui.NewColorFromPacked(v))
	})
	imguiObj.Set("newColorFromColor", func(call goja.FunctionCall) goja.Value {
		// TODO: 处理类型 color.Color
		c := color.RGBA{}
		return vm.ToValue(imgui.NewColorFromColor(c))
	})

	return nil
}
