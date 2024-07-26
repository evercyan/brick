package ximg

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os/exec"
	"strings"

	"github.com/evercyan/brick/xtype"
)

// Compress 图片压缩, quality 取值范围 1-100
func Compress(fpath string, quality int) (image.Image, error) {
	img, imgType, err := Parse(fpath)
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

// CompressPNG 依赖 pngquant
func CompressPNG(src image.Image, quality int) (image.Image, error) {
	if quality < 1 || quality >= 100 {
		return src, nil
	}
	var w bytes.Buffer
	if err := png.Encode(&w, src); err != nil {
		return nil, err
	}
	cmd := exec.Command("pngquant", "-", "--quality", xtype.ToString(quality))
	cmd.Stdin = strings.NewReader(w.String())
	fmt.Println(cmd.String())
	var o bytes.Buffer
	cmd.Stdout = &o
	if err := cmd.Run(); err != nil {
		return nil, err
	}
	return png.Decode(bytes.NewReader(o.Bytes()))
}

// CompressJPEG ...
func CompressJPEG(src image.Image, quality int) (image.Image, error) {
	if quality < 1 || quality >= 100 {
		return src, nil
	}
	var w bytes.Buffer
	err := jpeg.Encode(&w, src, &jpeg.Options{Quality: quality})
	if err != nil {
		return nil, err
	}
	return jpeg.Decode(bytes.NewReader(w.Bytes()))
}
