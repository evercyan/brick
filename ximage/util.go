package ximage

import (
	"fmt"
	"image"
	_ "image/gif" // ...
	"image/jpeg"
	"image/png"
	"os"

	"github.com/evercyan/brick/xencoding"
	"github.com/evercyan/brick/xfile"
)

// Parse ...
func Parse(filepath string) (image.Image, string, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, "", err
	}
	defer f.Close()
	return image.Decode(f)
}

// Type ...
func Type(filepath string) string {
	_, t, err := Parse(filepath)
	if err != nil {
		return ""
	}
	return t
}

// Size ...
func Size(filepath string) (int, int) {
	img, _, err := Parse(filepath)
	if err != nil {
		return 0, 0
	}
	return img.Bounds().Dx(), img.Bounds().Dy()
}

// Read ...
func Read(filepath string) image.Image {
	img, _, err := Parse(filepath)
	if err != nil {
		return nil
	}
	return img
}

// Write ...
func Write(filepath string, img image.Image) error {
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()
	ext := xfile.GetFileExt(filepath)
	switch ext {
	case "jpg", "jpeg":
		return jpeg.Encode(file, img, &jpeg.Options{
			Quality: jpeg.DefaultQuality,
		})
	case "png":
		return png.Encode(file, img)
	default:
		return fmt.Errorf("invalid image type: %s", ext)
	}
}

// Base64 ...
func Base64(filepath string) string {
	return fmt.Sprintf(
		"data:image/%s;base64,%s",
		Type(filepath),
		xencoding.Base64Encode(xfile.Read(filepath)),
	)
}
