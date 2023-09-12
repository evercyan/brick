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
func Read(isImage ...bool) string {
	t := clipboard.FmtText
	if len(isImage) > 0 && isImage[0] {
		t = clipboard.FmtImage
	}
	return string(clipboard.Read(t))
}

// ReadImage ...
func ReadImage(filepath string) error {
	content := string(clipboard.Read(clipboard.FmtImage))
	if content == "" {
		return fmt.Errorf("image not found")
	}
	return xfile.Write(filepath, content)
}
