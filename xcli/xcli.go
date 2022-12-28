package xcli

import (
	"fmt"
	"os"
	"os/exec"
)

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

// Exec ...
func Exec(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	// 需指定以下输入输出, 否则阻塞进程的终端无法打开
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Shell ...
func Shell(cmd string) string {
	b, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		return ""
	}
	return string(b)
}
