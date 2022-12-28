package xclip

import (
	"fmt"
	"strings"

	"github.com/evercyan/brick/xfile"
	"golang.design/x/clipboard"
)

// Write ...
func Write(args ...string) {
	clipboard.Write(clipboard.FmtText, []byte(strings.Join(args, " ")))
}

// WriteImage ...
func WriteImage(filepath string) {
	clipboard.Write(clipboard.FmtImage, []byte(xfile.Read(filepath)))
}

// Read ...
func Read() string {
	return string(clipboard.Read(clipboard.FmtText))
}

// ReadImage ...
func ReadImage(filepath string) error {
	content := string(clipboard.Read(clipboard.FmtImage))
	if content == "" {
		return fmt.Errorf("image not found")
	}
	return xfile.Write(filepath, content)
}
