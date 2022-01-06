package xfile

import (
	"bufio"
	"image"
	_ "image/gif"  // ...
	_ "image/jpeg" // ...
	_ "image/png"  // ...
	"os"
)

// IsExist ...
func IsExist(filepath string) bool {
	_, err := os.Stat(filepath)
	return err == nil || os.IsExist(err)
}

// IsFile ...
func IsFile(filepath string) bool {
	file, err := os.Stat(filepath)
	return err == nil && !file.IsDir()
}

// IsDir ...
func IsDir(filepath string) bool {
	file, err := os.Stat(filepath)
	return err == nil && file.IsDir()
}

// IsImage ...
func IsImage(filepath string) bool {
	file, err := os.Open(filepath)
	if err != nil {
		return false
	}
	defer file.Close()
	_, _, err = image.Decode(bufio.NewReader(file))
	return err == nil
}
