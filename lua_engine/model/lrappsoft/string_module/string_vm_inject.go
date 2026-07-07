package string_module

import (
	"strings"
	"unicode/utf8"

	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"
)

// StringModule 是 go-lua-vm 迁移后的模块壳。
type StringModule struct{}

func New() *StringModule { return &StringModule{} }

func (m *StringModule) Name() string { return "string" }

func (m *StringModule) IsAvailable() bool { return true }

func (m *StringModule) Register(engine model.Engine) error {
	engine.RegisterMethod("string.splitStr", "按分隔符切分字符串", func(source string, separator string) []string {
		return strings.Split(source, separator)
	}, true)
	engine.RegisterMethod("utf8.length", "返回 UTF-8 字符数", func(source string) int {
		return utf8.RuneCountInString(source)
	}, true)
	engine.RegisterMethod("utf8.left", "返回左侧指定字符数", func(source string, count int) string {
		runes := []rune(source)
		if count < 0 {
			count = 0
		}
		if count > len(runes) {
			count = len(runes)
		}
		return string(runes[:count])
	}, true)
	engine.RegisterMethod("utf8.right", "返回右侧指定字符数", func(source string, count int) string {
		runes := []rune(source)
		if count < 0 {
			count = 0
		}
		if count > len(runes) {
			count = len(runes)
		}
		return string(runes[len(runes)-count:])
	}, true)
	engine.RegisterMethod("utf8.mid", "返回中间指定字符段", func(source string, start int, length int) string {
		runes := []rune(source)
		if start < 0 {
			start = 0
		}
		if start >= len(runes) || length <= 0 {
			return ""
		}
		end := start + length
		if end > len(runes) {
			end = len(runes)
		}
		return string(runes[start:end])
	}, true)
	engine.RegisterMethod("utf8.strCut", "删除指定字符段", func(source string, start int, length int) string {
		runes := []rune(source)
		if start < 0 {
			start = 0
		}
		if start >= len(runes) || length <= 0 {
			return source
		}
		end := start + length
		if end > len(runes) {
			end = len(runes)
		}
		return string(append(runes[:start], runes[end:]...))
	}, true)
	engine.RegisterMethod("utf8.inStr", "从指定位置正向查找子串", func(start int, source string, pattern string) interface{} {
		sourceRunes := []rune(source)
		patternRunes := []rune(pattern)
		if start < 1 || start > len(sourceRunes) {
			return nil
		}
		for index := start - 1; index+len(patternRunes) <= len(sourceRunes); index++ {
			if string(sourceRunes[index:index+len(patternRunes)]) == pattern {
				return index + 1
			}
		}
		return nil
	}, true)
	engine.RegisterMethod("utf8.inStrRev", "从指定位置反向查找子串", func(source string, pattern string, start int) interface{} {
		sourceRunes := []rune(source)
		patternRunes := []rune(pattern)
		if start < 1 || start > len(sourceRunes) {
			start = len(sourceRunes)
		}
		for index := start - 1; index >= 0; index-- {
			if index+len(patternRunes) <= len(sourceRunes) && string(sourceRunes[index:index+len(patternRunes)]) == pattern {
				return index + 1
			}
		}
		return nil
	}, true)
	engine.RegisterMethod("utf8.strReverse", "反转 UTF-8 字符串", func(source string) string {
		runes := []rune(source)
		for left, right := 0, len(runes)-1; left < right; left, right = left+1, right-1 {
			runes[left], runes[right] = runes[right], runes[left]
		}
		return string(runes)
	}, true)
	return nil
}

func GetModule() model.Module { return &StringModule{} }
