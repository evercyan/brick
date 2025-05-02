package xencoding

import (
	"strings"
)

// ...
var (
	fullAngleMap = map[string]string{
		",": "，",
		".": "。",
		"!": "！",
		"?": "？",
		";": "；",
		":": "：",
		"~": "～",
		"'": "‘",
		`"`: "“",
		"(": "（",
		")": "）",
		"<": "《",
		">": "》",
		"[": "【",
		"]": "】",
	}
	halfAngleMap = map[string]string{
		"，": ", ",
		"。": ". ",
		"！": "! ",
		"？": "? ",
		"、": ", ",
		"；": "; ",
		"：": ": ",
		"～": "~",
		"‘": "'",
		"“": `"`,
		"”": `"`,
		"（": "(",
		"）": ")",
		"《": "<",
		"》": ">",
		"〈": "<",
		"〉": ">",
		"【": "[",
		"】": "]",
	}
)

// Ord ...
func Ord(s string) int {
	return int([]rune(s)[0])
}

// Chr ...
func Chr(s int) string {
	return string(rune(s))
}

// FullAngle ...
func FullAngle(s string) string {
	for k, v := range fullAngleMap {
		s = strings.ReplaceAll(s, k, v)
	}
	return s
}

// HalfAngle ...
func HalfAngle(s string) string {
	for k, v := range halfAngleMap {
		s = strings.ReplaceAll(s, k, v)
	}
	return s
}
