package xconvert

import (
	"encoding/hex"
	"fmt"
	"image/color"
	"regexp"
	"strings"
)

// Hex2RGB ...
func Hex2RGB(s string) (int, int, int) {
	re := regexp.MustCompile(`^#([0-9a-fA-F]{6}|[0-9a-fA-F]{3})$`)
	if !re.MatchString(s) {
		return 0, 0, 0
	}
	s = strings.ReplaceAll(s, "#", "")
	if len(s) == 3 {
		s = fmt.Sprintf("%c%c%c%c%c%c", s[0], s[0], s[1], s[1], s[2], s[2])
	}
	d, err := hex.DecodeString(s)
	if err != nil {
		return 0, 0, 0
	}
	return int(d[0]), int(d[1]), int(d[2])
}

// Hex2Color ...
func Hex2Color(s string) color.Color {
	r, g, b := Hex2RGB(s)
	fmt.Println(r, g, b)
	return color.RGBA{
		R: uint8(r),
		G: uint8(g),
		B: uint8(b),
		A: 255,
	}
}
