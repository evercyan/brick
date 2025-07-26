package xfile

import (
	"testing"

	"github.com/evercyan/brick/xjson"
)

func TestGetMP4Info(t *testing.T) {
	info, err := GetMP4Info("/Users/Cyan/Downloads/input.mp4")
	if err != nil {
		t.Fatal(err)
	}
	xjson.Pretty(info, true)
}
