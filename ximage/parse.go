package ximage

import (
	"image"
	"os"
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
