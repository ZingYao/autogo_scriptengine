package debugger

import (
	"fmt"
	"strings"
	"unicode"
)

// InstrumentSource 在 Lua 语句前插入调试命中调用。
func InstrumentSource(source string, file string) string {
	if source == "" {
		return source
	}
	file = normalizeFile(file)
	lines := strings.Split(source, "\n")
	out := make([]string, 0, len(lines)*2)
	for i, line := range lines {
		lineNo := i + 1
		if shouldInstrumentLine(line) {
			indent := leadingWhitespace(line)
			out = append(out, fmt.Sprintf("%s%s(%q, %d)", indent, hitFunction, file, lineNo))
		}
		out = append(out, line)
	}
	return strings.Join(out, "\n")
}

func leadingWhitespace(line string) string {
	idx := 0
	for idx < len(line) {
		r := rune(line[idx])
		if !unicode.IsSpace(r) {
			break
		}
		idx++
	}
	return line[:idx]
}

func shouldInstrumentLine(line string) bool {
	trimmed := strings.TrimSpace(line)
	if trimmed == "" || strings.HasPrefix(trimmed, "--") {
		return false
	}
	blockedPrefixes := []string{
		"end", "else", "elseif", "until", ")", "]", "}", ",", ".", ":", "+", "-", "*", "/", "%",
	}
	for _, prefix := range blockedPrefixes {
		if strings.HasPrefix(trimmed, prefix) {
			return false
		}
	}
	allowedPrefixes := []string{
		"local ", "if ", "for ", "while ", "repeat", "function ", "return", "break", "do",
	}
	for _, prefix := range allowedPrefixes {
		if strings.HasPrefix(trimmed, prefix) {
			return true
		}
	}
	if isIdentifierStart(rune(trimmed[0])) {
		return strings.Contains(trimmed, "=") || strings.Contains(trimmed, "(")
	}
	return false
}

func isIdentifierStart(r rune) bool {
	return r == '_' || unicode.IsLetter(r)
}
