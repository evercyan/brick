package xfont

import (
	"embed"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
)

//go:embed "AlimamaFangYuanTiVF-Thin.ttf"
var fonts embed.FS

// loadFont ...
func loadFont(name string) *truetype.Font {
	fontFile, _ := fonts.ReadFile(name)
	font, _ := freetype.ParseFont(fontFile)
	return font
}

// LoadAlimamaFangYuanTiVF ...
func LoadAlimamaFangYuanTiVF() *truetype.Font {
	return loadFont("AlimamaFangYuanTiVF-Thin.ttf")
}
