package xdraw

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"os/exec"
	"strings"
	"time"

	"github.com/evercyan/brick/xtype"

	"github.com/evercyan/brick/ximg"
)

// Compress 图片压缩, quality 取值范围 1-100
func Compress(fpath string, quality int) (image.Image, error) {
	img, imgType, err := ximg.Parse(fpath)
	if err != nil {
		return nil, err
	}
	switch imgType {
	case "png":
		return CompressPNG(img, quality)
	case "jpg", "jpeg":
		return CompressJPEG(img, quality)
	default:
		return nil, fmt.Errorf("unsupported image type: %s", imgType)
	}
}

// CompressPNG ...
// origin: 25.8MB
// quality=80: 9.7MB
// quality=50: 7.5MB
// quality=10: 5.7MB
// quality=2: 3.5MB
// 耗时 2min
func CompressPNG(src image.Image, quality int) (image.Image, error) {
	if quality < 1 || quality >= 100 {
		return src, nil
	}
	fmt.Println("==== 2", time.Now())
	var w bytes.Buffer
	if err := png.Encode(&w, src); err != nil {
		return nil, err
	}
	fmt.Println("==== 3", time.Now())
	cmd := exec.Command("pngquant", "-", "--quality", xtype.ToString(quality))
	cmd.Stdin = strings.NewReader(w.String())
	fmt.Println(cmd.String())
	var o bytes.Buffer
	cmd.Stdout = &o

	fmt.Println("==== 4", time.Now())
	if err := cmd.Run(); err != nil {

		fmt.Println("==== 5", time.Now())
		return nil, err
	}
	fmt.Println("==== 6", time.Now())
	r1, r2 := png.Decode(bytes.NewReader(o.Bytes()))
	fmt.Println("==== 7", time.Now())
	return r1, r2
	//return png.Decode(bytes.NewReader(o.Bytes()))

}

func CompressJPEG(src image.Image, quality int) (image.Image, error) {
	if quality < 1 || quality >= 100 {
		return src, nil
	}
	return nil, nil
}
