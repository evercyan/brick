package ximage

import (
	"image"
	"os"
)

// getImage ...
func getImage(filepath string) (image.Image, string, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, "", err
	}
	defer f.Close()
	return image.Decode(f)
}
