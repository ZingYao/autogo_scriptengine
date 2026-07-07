//go:build ignore
// +build ignore

package string_module

import (
	"strings"
	"unicode/utf8"

	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"

	lua "github.com/yuin/gopher-lua"
)

// StringModule string 模块（懒人精灵兼容）
type StringModule struct{}

// Name 返回模块名称
func (m *StringModule) Name() string {
	return "string"
}

// IsAvailable 检查模块是否可用
func (m *StringModule) IsAvailable() bool {
	return true
}

// Register 向引擎注册方法
func (m *StringModule) Register(engine model.Engine) error {
	state := engine.GetState()

	// 获取全局 string 表
	stringTable := state.GetGlobal("string")
	if stringTable == lua.LNil {
		stringTable = state.NewTable()
		state.SetGlobal("string", stringTable)
	}

	// 注册 splitStr - 字符串分割（懒人扩展）
	stringTable.(*lua.LTable).RawSetString("splitStr", state.NewFunction(func(L *lua.LState) int {
		src := L.CheckString(1)
		sep := L.CheckString(2)

		// 使用 Go 的 strings.Split 进行分割
		parts := strings.Split(src, sep)

		// 创建 Lua 表
		result := L.NewTable()
		for i, part := range parts {
			result.RawSetInt(i+1, lua.LString(part))
		}

		L.Push(result)
		return 1
	}))

	// 创建 utf8 模块表
	utf8Table := state.NewTable()

	// 注册 utf8.inStr - UTF-8 字符串查找
	utf8Table.RawSetString("inStr", state.NewFunction(func(L *lua.LState) int {
		start := L.CheckInt(1)
		s := L.CheckString(2)
		pattern := L.CheckString(3)

		// 将字符串转换为 rune 切片以支持 UTF-8
		runes := []rune(s)
		patternRunes := []rune(pattern)

		// 检查起始位置是否有效
		if start < 1 || start > len(runes) {
			L.Push(lua.LNil)
			return 1
		}

		// 从指定位置开始查找
		for i := start - 1; i < len(runes); i++ {
			if i+len(patternRunes) <= len(runes) {
				match := true
				for j := 0; j < len(patternRunes); j++ {
					if runes[i+j] != patternRunes[j] {
						match = false
						break
					}
				}
				if match {
					L.Push(lua.LNumber(i + 1))
					return 1
				}
			}
		}

		L.Push(lua.LNil)
		return 1
	}))

	// 注册 utf8.inStrRev - UTF-8 字符串反向查找
	utf8Table.RawSetString("inStrRev", state.NewFunction(func(L *lua.LState) int {
		s := L.CheckString(1)
		pattern := L.CheckString(2)
		start := L.CheckInt(3)

		// 将字符串转换为 rune 切片以支持 UTF-8
		runes := []rune(s)
		patternRunes := []rune(pattern)

		// 检查起始位置是否有效
		if start < 1 || start > len(runes) {
			start = len(runes)
		}

		// 从指定位置反向查找
		for i := start - 1; i >= 0; i-- {
			if i+len(patternRunes) <= len(runes) {
				match := true
				for j := 0; j < len(patternRunes); j++ {
					if runes[i+j] != patternRunes[j] {
						match = false
						break
					}
				}
				if match {
					L.Push(lua.LNumber(i + 1))
					return 1
				}
			}
		}

		L.Push(lua.LNil)
		return 1
	}))

	// 注册 utf8.strReverse - UTF-8 字符串反转
	utf8Table.RawSetString("strReverse", state.NewFunction(func(L *lua.LState) int {
		s := L.CheckString(1)

		// 将字符串转换为 rune 切片以支持 UTF-8
		runes := []rune(s)

		// 反转 rune 切片
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}

		L.Push(lua.LString(string(runes)))
		return 1
	}))

	// 注册 utf8.length - UTF-8 字符串长度
	utf8Table.RawSetString("length", state.NewFunction(func(L *lua.LState) int {
		s := L.CheckString(1)

		// 使用 utf8.RuneCountInString 获取字符数
		length := utf8.RuneCountInString(s)

		L.Push(lua.LNumber(length))
		return 1
	}))

	// 注册 utf8.left - UTF-8 字符串左侧截取
	utf8Table.RawSetString("left", state.NewFunction(func(L *lua.LState) int {
		s := L.CheckString(1)
		n := L.CheckInt(2)

		// 将字符串转换为 rune 切片以支持 UTF-8
		runes := []rune(s)

		// 检查截取长度是否有效
		if n < 0 {
			n = 0
		}
		if n > len(runes) {
			n = len(runes)
		}

		// 截取左侧 n 个字符
		result := string(runes[:n])

		L.Push(lua.LString(result))
		return 1
	}))

	// 注册 utf8.right - UTF-8 字符串右侧截取
	utf8Table.RawSetString("right", state.NewFunction(func(L *lua.LState) int {
		s := L.CheckString(1)
		n := L.CheckInt(2)

		// 将字符串转换为 rune 切片以支持 UTF-8
		runes := []rune(s)

		// 检查截取长度是否有效
		if n < 0 {
			n = 0
		}
		if n > len(runes) {
			n = len(runes)
		}

		// 截取右侧 n 个字符
		start := len(runes) - n
		result := string(runes[start:])

		L.Push(lua.LString(result))
		return 1
	}))

	// 注册 utf8.mid - UTF-8 字符串中间截取
	utf8Table.RawSetString("mid", state.NewFunction(func(L *lua.LState) int {
		s := L.CheckString(1)
		start := L.CheckInt(2)
		length := L.CheckInt(3)

		// 将字符串转换为 rune 切片以支持 UTF-8
		runes := []rune(s)

		// 检查起始位置是否有效（懒人精灵的 start 是从 0 开始的索引）
		if start < 0 {
			start = 0
		}
		if start >= len(runes) {
			L.Push(lua.LString(""))
			return 1
		}

		// 检查截取长度是否有效
		if length < 0 {
			length = 0
		}

		// 计算结束位置
		end := start + length
		if end > len(runes) {
			end = len(runes)
		}

		// 截取中间部分
		result := string(runes[start:end])

		L.Push(lua.LString(result))
		return 1
	}))

	// 注册 utf8.strCut - UTF-8 字符串裁剪
	utf8Table.RawSetString("strCut", state.NewFunction(func(L *lua.LState) int {
		s := L.CheckString(1)
		start := L.CheckInt(2)
		length := L.CheckInt(3)

		// 将字符串转换为 rune 切片以支持 UTF-8
		runes := []rune(s)

		// 检查起始位置是否有效（懒人精灵的 start 是从 0 开始的索引）
		if start < 0 {
			start = 0
		}
		if start >= len(runes) {
			L.Push(lua.LString(s))
			return 1
		}

		// 检查裁剪长度是否有效
		if length < 0 {
			length = 0
		}

		// 计算结束位置
		end := start + length
		if end > len(runes) {
			end = len(runes)
		}

		// 移除指定范围的字符
		result := string(runes[:start]) + string(runes[end:])

		L.Push(lua.LString(result))
		return 1
	}))

	// 注册 utf8 到全局
	state.SetGlobal("utf8", utf8Table)

	// 注册 splitStr 到全局（可以直接调用，不需要 string.splitStr）
	state.SetGlobal("splitStr", stringTable.(*lua.LTable).RawGetString("splitStr"))

	engine.RegisterMethod("string.splitStr", "字符串分割", func(src, sep string) []string {
		return strings.Split(src, sep)
	}, true)
	engine.RegisterMethod("utf8.inStr", "UTF-8 字符串查找", utf8InStr, true)
	engine.RegisterMethod("utf8.inStrRev", "UTF-8 字符串反向查找", utf8InStrRev, true)
	engine.RegisterMethod("utf8.strReverse", "UTF-8 字符串反转", utf8StrReverse, true)
	engine.RegisterMethod("utf8.length", "UTF-8 字符串长度", func(s string) int {
		return utf8.RuneCountInString(s)
	}, true)
	engine.RegisterMethod("utf8.left", "UTF-8 字符串左侧截取", utf8Left, true)
	engine.RegisterMethod("utf8.right", "UTF-8 字符串右侧截取", utf8Right, true)
	engine.RegisterMethod("utf8.mid", "UTF-8 字符串中间截取", utf8Mid, true)
	engine.RegisterMethod("utf8.strCut", "UTF-8 字符串裁剪", utf8StrCut, true)

	return nil
}

func utf8InStr(start int, s, pattern string) int {
	runes := []rune(s)
	patternRunes := []rune(pattern)
	if start < 1 || start > len(runes) {
		return 0
	}
	for i := start - 1; i < len(runes); i++ {
		if i+len(patternRunes) <= len(runes) {
			match := true
			for j := 0; j < len(patternRunes); j++ {
				if runes[i+j] != patternRunes[j] {
					match = false
					break
				}
			}
			if match {
				return i + 1
			}
		}
	}
	return 0
}

func utf8InStrRev(s, pattern string, start int) int {
	runes := []rune(s)
	patternRunes := []rune(pattern)
	if start < 1 || start > len(runes) {
		start = len(runes)
	}
	for i := start - 1; i >= 0; i-- {
		if i+len(patternRunes) <= len(runes) {
			match := true
			for j := 0; j < len(patternRunes); j++ {
				if runes[i+j] != patternRunes[j] {
					match = false
					break
				}
			}
			if match {
				return i + 1
			}
		}
	}
	return 0
}

func utf8StrReverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func utf8Left(s string, n int) string {
	runes := []rune(s)
	if n < 0 {
		n = 0
	}
	if n > len(runes) {
		n = len(runes)
	}
	return string(runes[:n])
}

func utf8Right(s string, n int) string {
	runes := []rune(s)
	if n < 0 {
		n = 0
	}
	if n > len(runes) {
		n = len(runes)
	}
	return string(runes[len(runes)-n:])
}

func utf8Mid(s string, start, length int) string {
	runes := []rune(s)
	if start < 0 {
		start = 0
	}
	if start >= len(runes) {
		return ""
	}
	if length < 0 {
		length = 0
	}
	end := start + length
	if end > len(runes) {
		end = len(runes)
	}
	return string(runes[start:end])
}

func utf8StrCut(s string, start, length int) string {
	runes := []rune(s)
	if start < 0 {
		start = 0
	}
	if start >= len(runes) {
		return s
	}
	if length < 0 {
		length = 0
	}
	end := start + length
	if end > len(runes) {
		end = len(runes)
	}
	return string(runes[:start]) + string(runes[end:])
}
