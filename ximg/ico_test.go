package ximg

import (
	"fmt"
	"testing"

	"github.com/evercyan/brick/xfile"
)

func TestToIco(t *testing.T) {
	dstPath := xfile.Temp(".ico")
	err := ToIco("../logo.png", dstPath, 300)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(dstPath)
}
