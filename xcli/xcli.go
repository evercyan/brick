package xcli

import (
	"fmt"
	"os/exec"
)

// Exec ...
func Exec(cmd string) string {
	b, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		return ""
	}
	return string(b)
}

// HideCursor ...
func HideCursor() {
	fmt.Printf("\033[?25l")
}

// ShowCursor ...
func ShowCursor() {
	fmt.Printf("\033[?25h")
}

// ClearLine ...
func ClearLine() {
	fmt.Printf("\r\033[0K")
}
