package ximg

import (
	"encoding/base64"
	"fmt"
	"image"
	_ "image/gif" // ...
	"image/jpeg"
	"image/png"
	"os"
	"regexp"
	"strings"

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
func Write(filepath string, src image.Image) error {
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()
	ext := xfile.GetFileExt(filepath)
	switch ext {
	case "jpg", "jpeg":
		return jpeg.Encode(file, src, &jpeg.Options{
			Quality: jpeg.DefaultQuality,
		})
	case "png":
		return png.Encode(file, src)
	default:
		return fmt.Errorf("invalid image type: %s", ext)
	}
}

// Base64Encode ...
func Base64Encode(filepath string) string {
	return fmt.Sprintf(
		"data:image/%s;base64,%s",
		Type(filepath),
		xencoding.Base64Encode(xfile.Read(filepath)),
	)
}

// base64Regexp ...
var base64Regexp = regexp.MustCompile(`data:image/[^;]+;base64,`)

// Base64Decode ...
func Base64Decode(encoded string, filepath string) error {
	if strings.HasPrefix(encoded, "data:image/") {
		encoded = base64Regexp.ReplaceAllString(encoded, "")
	}
	imageData, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return err
	}
	return xfile.Write(filepath, string(imageData))
}
