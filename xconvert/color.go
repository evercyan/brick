package xconvert

import (
	"encoding/hex"
	"fmt"
	"image/color"
	"regexp"
	"strconv"
)

// Hex2RGB ...
func Hex2RGB(s string) (int, int, int) {
	re := regexp.MustCompile(`^([0-9a-fA-F]{6}|[0-9a-fA-F]{3})$`)
	if !re.MatchString(s) {
		return 0, 0, 0
	}
	if len(s) == 3 {
		s = fmt.Sprintf("%c%c%c%c%c%c", s[0], s[0], s[1], s[1], s[2], s[2])
	}
	d, _ := hex.DecodeString(s)
	return int(d[0]), int(d[1]), int(d[2])
}

// RGB2Hex ...
func RGB2Hex(r, g, b int) string {
	t2x := func(v int) string {
		res := strconv.FormatInt(int64(v), 16)
		if len(res) == 1 {
			res = "0" + res
		}
		return res
	}
	return t2x(r) + t2x(g) + t2x(b)
}

// Hex2Color ...
func Hex2Color(s string) color.Color {
	r, g, b := Hex2RGB(s)
	return color.RGBA{
		R: uint8(r),
		G: uint8(g),
		B: uint8(b),
		A: 255,
	}
}
