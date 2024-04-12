package xlogo

import (
	"github.com/evercyan/brick/xconvert"
)

// defaultOption ...
var defaultOption = &option{
	Style:   StyleSingle,
	Width:   512,
	Height:  512,
	Radious: 0,
	Color:   xconvert.Hex2Color("#000000"),
	BgColor: xconvert.Hex2Color("#ffffff"),
}
