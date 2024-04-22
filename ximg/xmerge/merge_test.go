package xmerge

import (
	"image"
	"testing"

	"github.com/evercyan/brick/ximg"
	"github.com/evercyan/brick/xutil"
	"github.com/stretchr/testify/assert"
)

var (
	DownloadDir = xutil.Getenv("DOWNLOAD_DIR", "/tmp")
)

func TestMerge(t *testing.T) {
	src := ximg.Read("../../logo.png")
	dst, err := Merge([]image.Image{src, src}, 1, 2)
	assert.Nil(t, err)
	assert.NotNil(t, dst)
}

func TestMergeTo(t *testing.T) {
	imgs := []image.Image{
		ximg.Read(DownloadDir + "/merge/001.png"),
		ximg.Read(DownloadDir + "/merge/002.png"),
		ximg.Read(DownloadDir + "/merge/003.png"),
		ximg.Read(DownloadDir + "/merge/004.png"),
		ximg.Read(DownloadDir + "/merge/005.png"),
	}
	dst1, err1 := MergeToCol(imgs)
	if err1 != nil {
		panic(err1)
	}
	ximg.Write(DownloadDir+"/merge/merge_col.png", dst1)

	dst2, err2 := MergeToRow(imgs)
	if err2 != nil {
		panic(err2)
	}
	ximg.Write(DownloadDir+"/merge/merge_row.png", dst2)
}
