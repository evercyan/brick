package ximg

import (
	"bytes"
	"context"
	"image/jpeg"
	"os"

	"github.com/evercyan/brick/xfile"
	"github.com/evercyan/brick/xhttp"
	"golang.org/x/image/webp"
)

// Download 下载文件
func Download(ctx context.Context, imgURL, imgPath string) error {
	resp, err := xhttp.New().Get(ctx, imgURL, nil)
	if err != nil {
		return err
	}
	// 如果原格式为 webp, 当前保存为 jpg
	imgExt := xfile.GetFileExt(imgPath)
	if resp.Header.Get("Content-Type") == "image/webp" &&
		(imgExt == "jpeg" || imgExt == "jpg") {
		img, err := webp.Decode(bytes.NewReader(resp.Bytes()))
		if err != nil {
			return err
		}
		jfile, err := os.Create(imgPath)
		if err != nil {
			return err
		}
		defer jfile.Close()
		option := &jpeg.Options{Quality: 95}
		if err := jpeg.Encode(jfile, img, option); err != nil {
			return err
		}
		return nil
	}
	return xfile.Write(imgPath, resp.String())
}
