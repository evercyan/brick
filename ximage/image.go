package ximage

import (
	_ "image/gif"  // ...
	_ "image/jpeg" // ...
	_ "image/png"  // ...
)

// Type ...
func Type(filepath string) string {
	_, t, err := getImage(filepath)
	if err != nil {
		return ""
	}
	return t
}

// Size ...
func Size(filepath string) (int, int) {
	img, _, err := getImage(filepath)
	if err != nil {
		return 0, 0
	}
	return img.Bounds().Dx(), img.Bounds().Dy()
}
