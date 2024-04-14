package xavatar

import (
	"fmt"
	"testing"

	"github.com/evercyan/brick/xfile"
)

func TestSquare(t *testing.T) {
	for i := 1; i <= 20; i++ {
		fpath := xfile.Temp(".png")
		New(WithStyle(StyleSquare), WithBlockNum(i)).Save("hello", fpath)
		fmt.Println("====", i, fpath)
	}

	fpath2 := xfile.Temp("square_radious.png")
	New(WithStyle(StyleSquare), WithRadious(0.5)).Save("helloworld", fpath2)
	fmt.Println("==== fpath2", fpath2)
}
