//go:build ignore
// +build ignore

package cv

import (
	"image"
	"sync"
	"unsafe"

	"github.com/Dasongzi1366/AutoGo/images"
	"github.com/Dasongzi1366/AutoGo/opencv"
	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"

	lua "github.com/yuin/gopher-lua"
)

// CvModule cv 模块（懒人精灵兼容）
type CvModule struct {
	pointCounter   int64
	point2fCounter int64
	pointMap       map[int64]*image.Point
	point2fMap     map[int64]*image.Point
	mu             sync.RWMutex
}

// Name 返回模块名称
func (m *CvModule) Name() string {
	return "cv"
}

// IsAvailable 检查模块是否可用
func (m *CvModule) IsAvailable() bool {
	return true
}

// Register 向引擎注册方法
func (m *CvModule) Register(engine model.Engine) error {
	// 初始化 map
	m.pointMap = make(map[int64]*image.Point)
	m.point2fMap = make(map[int64]*image.Point)

	state := engine.GetState()

	// 获取 package.preload 表
	packagePreload := state.GetGlobal("package")
	if packagePreload == lua.LNil {
		packagePreload = state.NewTable()
		state.SetGlobal("package", packagePreload)
	}

	packageTable := packagePreload.(*lua.LTable)
	preloadTable := packageTable.RawGetString("preload")
	if preloadTable == lua.LNil {
		preloadTable = state.NewTable()
		packageTable.RawSetString("preload", preloadTable)
	}

	// 创建 cv 模块表
	cvModuleTable := state.NewTable()

	// 注册 cv.snapShot - 区域截图并返回 Mat 矩阵
	cvModuleTable.RawSetString("snapShot", state.NewFunction(func(L *lua.LState) int {
		l := L.CheckInt(1)
		t := L.CheckInt(2)
		r := L.CheckInt(3)
		b := L.CheckInt(4)

		// 使用 AutoGo 的 CaptureScreen 截取屏幕
		img := images.CaptureScreen(l, t, r, b, 0)
		if img == nil {
			L.Push(lua.LNil)
			L.Push(lua.LString("截图失败"))
			return 2
		}

		// 将 image.NRGBA 转换为 opencv.Mat
		mat, err := opencv.ImageToMatRGBA(img)
		if err != nil {
			L.Push(lua.LNil)
			L.Push(lua.LString(err.Error()))
			return 2
		}

		// 返回 Mat 对象
		ud := L.NewUserData()
		ud.Value = unsafe.Pointer(&mat)
		L.Push(ud)
		return 1
	}))

	// 注册 cv.newPoint - 创建 Point 指针
	cvModuleTable.RawSetString("newPoint", state.NewFunction(func(L *lua.LState) int {
		x := L.CheckInt(1)
		y := L.CheckInt(2)

		// 创建新的 Point 对象
		point := &image.Point{X: x, Y: y}

		// 生成唯一 ID
		m.mu.Lock()
		m.pointCounter++
		id := m.pointCounter
		m.pointMap[id] = point
		m.mu.Unlock()

		// 返回 ID 作为指针
		ud := L.NewUserData()
		ud.Value = id
		L.Push(ud)
		return 1
	}))

	// 注册 cv.getPoint - 获取 Point 坐标
	cvModuleTable.RawSetString("getPoint", state.NewFunction(func(L *lua.LState) int {
		ud := L.CheckUserData(1)
		if ud == nil {
			L.Push(lua.LNil)
			return 1
		}

		// 获取 ID
		id, ok := ud.Value.(int64)
		if !ok {
			L.Push(lua.LNil)
			return 1
		}

		// 从 map 中获取 Point
		m.mu.RLock()
		point, exists := m.pointMap[id]
		m.mu.RUnlock()

		if !exists || point == nil {
			L.Push(lua.LNil)
			return 1
		}

		result := L.NewTable()
		result.RawSetString("x", lua.LNumber(point.X))
		result.RawSetString("y", lua.LNumber(point.Y))
		L.Push(result)
		return 1
	}))

	// 注册 cv.setPoint - 设置 Point 坐标
	cvModuleTable.RawSetString("setPoint", state.NewFunction(func(L *lua.LState) int {
		ud := L.CheckUserData(1)
		x := L.CheckInt(2)
		y := L.CheckInt(3)

		// 获取 ID
		id, ok := ud.Value.(int64)
		if !ok {
			return 0
		}

		// 从 map 中获取 Point 并修改
		m.mu.Lock()
		point, exists := m.pointMap[id]
		if exists && point != nil {
			point.X = x
			point.Y = y
		}
		m.mu.Unlock()

		return 0
	}))

	// 注册 cv.deletePtr - 删除并回收指针
	cvModuleTable.RawSetString("deletePtr", state.NewFunction(func(L *lua.LState) int {
		ud := L.CheckAny(1)

		// 尝试删除 Point
		if userData, ok := ud.(*lua.LUserData); ok {
			if id, ok := userData.Value.(int64); ok {
				m.mu.Lock()
				delete(m.pointMap, id)
				delete(m.point2fMap, id)
				m.mu.Unlock()
			}
		}

		return 0
	}))

	// 注册 cv.deletePoint - 删除 Point 指针（别名）
	cvModuleTable.RawSetString("deletePoint", state.NewFunction(func(L *lua.LState) int {
		ud := L.CheckAny(1)

		// 尝试删除 Point
		if userData, ok := ud.(*lua.LUserData); ok {
			if id, ok := userData.Value.(int64); ok {
				m.mu.Lock()
				delete(m.pointMap, id)
				m.mu.Unlock()
			}
		}

		return 0
	}))

	// 注册 cv.deletePoint2f - 删除 Point2f 指针（别名）
	cvModuleTable.RawSetString("deletePoint2f", state.NewFunction(func(L *lua.LState) int {
		ud := L.CheckAny(1)

		// 尝试删除 Point2f
		if userData, ok := ud.(*lua.LUserData); ok {
			if id, ok := userData.Value.(int64); ok {
				m.mu.Lock()
				delete(m.point2fMap, id)
				m.mu.Unlock()
			}
		}

		return 0
	}))

	// 注册 cv.newPoint2f - 创建 Point2f 指针（使用 image.Point 模拟）
	cvModuleTable.RawSetString("newPoint2f", state.NewFunction(func(L *lua.LState) int {
		x := L.CheckNumber(1)
		y := L.CheckNumber(2)

		// 创建新的 Point 对象
		point := &image.Point{X: int(x), Y: int(y)}

		// 生成唯一 ID
		m.mu.Lock()
		m.point2fCounter++
		id := m.point2fCounter
		m.point2fMap[id] = point
		m.mu.Unlock()

		// 返回 ID 作为指针
		ud := L.NewUserData()
		ud.Value = id
		L.Push(ud)
		return 1
	}))

	// 注册 cv.getPoint2f - 获取 Point2f 坐标
	cvModuleTable.RawSetString("getPoint2f", state.NewFunction(func(L *lua.LState) int {
		ud := L.CheckUserData(1)
		if ud == nil {
			L.Push(lua.LNil)
			return 1
		}

		// 获取 ID
		id, ok := ud.Value.(int64)
		if !ok {
			L.Push(lua.LNil)
			return 1
		}

		// 从 map 中获取 Point
		m.mu.RLock()
		point, exists := m.point2fMap[id]
		m.mu.RUnlock()

		if !exists || point == nil {
			L.Push(lua.LNil)
			return 1
		}

		result := L.NewTable()
		result.RawSetString("x", lua.LNumber(float64(point.X)))
		result.RawSetString("y", lua.LNumber(float64(point.Y)))
		L.Push(result)
		return 1
	}))

	// 注册 cv.setPoint2f - 设置 Point2f 坐标
	cvModuleTable.RawSetString("setPoint2f", state.NewFunction(func(L *lua.LState) int {
		ud := L.CheckUserData(1)
		x := L.CheckNumber(2)
		y := L.CheckNumber(3)

		// 获取 ID
		id, ok := ud.Value.(int64)
		if !ok {
			return 0
		}

		// 从 map 中获取 Point 并修改
		m.mu.Lock()
		point, exists := m.point2fMap[id]
		if exists && point != nil {
			point.X = int(x)
			point.Y = int(y)
		}
		m.mu.Unlock()

		return 0
	}))

	// 注册 cv.newInt - 创建整型指针
	cvModuleTable.RawSetString("newInt", state.NewFunction(func(L *lua.LState) int {
		val := L.CheckInt(1)
		L.Push(lua.LNumber(val))
		return 1
	}))

	// 注册 cv.getInt - 获取整型值
	cvModuleTable.RawSetString("getInt", state.NewFunction(func(L *lua.LState) int {
		val := L.CheckNumber(1)
		L.Push(lua.LNumber(val))
		return 1
	}))

	// 注册 cv.setInt - 设置整型值
	cvModuleTable.RawSetString("setInt", state.NewFunction(func(L *lua.LState) int {
		// 第一个参数是原值（不使用），第二个参数是要设置的新值
		newVal := L.CheckInt(2)
		L.Push(lua.LNumber(newVal))
		return 1
	}))

	// 注册 cv.newDouble - 创建双精度浮点型指针
	cvModuleTable.RawSetString("newDouble", state.NewFunction(func(L *lua.LState) int {
		val := L.CheckNumber(1)
		L.Push(lua.LNumber(val))
		return 1
	}))

	// 注册 cv.getDouble - 获取双精度浮点值
	cvModuleTable.RawSetString("getDouble", state.NewFunction(func(L *lua.LState) int {
		val := L.CheckNumber(1)
		L.Push(lua.LNumber(val))
		return 1
	}))

	// 注册 cv.setDouble - 设置双精度浮点值
	cvModuleTable.RawSetString("setDouble", state.NewFunction(func(L *lua.LState) int {
		// 第一个参数是原值（不使用），第二个参数是要设置的新值
		newVal := L.CheckNumber(2)
		L.Push(lua.LNumber(newVal))
		return 1
	}))

	// 注册 cv.newFloat - 创建单精度浮点型指针
	cvModuleTable.RawSetString("newFloat", state.NewFunction(func(L *lua.LState) int {
		val := L.CheckNumber(1)
		L.Push(lua.LNumber(val))
		return 1
	}))

	// 注册 cv.getFloat - 获取单精度浮点值
	cvModuleTable.RawSetString("getFloat", state.NewFunction(func(L *lua.LState) int {
		val := L.CheckNumber(1)
		L.Push(lua.LNumber(val))
		return 1
	}))

	// 注册 cv.setFloat - 设置单精度浮点值
	cvModuleTable.RawSetString("setFloat", state.NewFunction(func(L *lua.LState) int {
		// 第一个参数是原值（不使用），第二个参数是要设置的新值
		newVal := L.CheckNumber(2)
		L.Push(lua.LNumber(newVal))
		return 1
	}))

	// 注册 cv.newLong - 创建长整型指针
	cvModuleTable.RawSetString("newLong", state.NewFunction(func(L *lua.LState) int {
		val := L.CheckNumber(1)
		L.Push(lua.LNumber(val))
		return 1
	}))

	// 注册 cv.getLong - 获取长整型值
	cvModuleTable.RawSetString("getLong", state.NewFunction(func(L *lua.LState) int {
		val := L.CheckNumber(1)
		L.Push(lua.LNumber(val))
		return 1
	}))

	// 注册 cv.setLong - 设置长整型值
	cvModuleTable.RawSetString("setLong", state.NewFunction(func(L *lua.LState) int {
		// 第一个参数是原值（不使用），第二个参数是要设置的新值
		newVal := L.CheckNumber(2)
		L.Push(lua.LNumber(newVal))
		return 1
	}))

	// 注册 cv.newByte - 创建字节型指针
	cvModuleTable.RawSetString("newByte", state.NewFunction(func(L *lua.LState) int {
		val := L.CheckInt(1)
		L.Push(lua.LNumber(val))
		return 1
	}))

	// 注册 cv.getByte - 获取字节值
	cvModuleTable.RawSetString("getByte", state.NewFunction(func(L *lua.LState) int {
		val := L.CheckNumber(1)
		L.Push(lua.LNumber(val))
		return 1
	}))

	// 注册 cv.setByte - 设置字节值
	cvModuleTable.RawSetString("setByte", state.NewFunction(func(L *lua.LState) int {
		// 第一个参数是原值（不使用），第二个参数是要设置的新值
		newVal := L.CheckInt(2)
		L.Push(lua.LNumber(newVal))
		return 1
	}))

	// 注册到 package.preload
	preloadTable.(*lua.LTable).RawSetString("cv", state.NewFunction(func(L *lua.LState) int {
		L.Push(cvModuleTable)
		return 1
	}))

	// 同时注册为全局变量 cv（兼容直接调用）
	state.SetGlobal("cv", cvModuleTable)

	return nil
}
