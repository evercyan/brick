package xcli

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
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

// ExecCB 适用于执行持续输出的终端命令 e.g. ping baidu.com
func ExecCB(command string, cb func(string)) error {
	cmd := exec.Command("sh", "-c", command)
	output, _ := cmd.StdoutPipe()
	if err := cmd.Start(); err != nil {
		return err
	}
	for {
		buf := make([]byte, 1024)
		n, err := output.Read(buf)
		if err != nil {
			break
		}
		cb(string(buf[:n]))
	}
	return cmd.Wait()
}

// Progress ...
func Progress(prefix string, percent float64, blocks int, suffixs ...string) {
	pos := int(percent * float64(blocks))
	s := fmt.Sprintf(
		"[%s] %s%*s %6.2f%% \t%s",
		prefix,
		strings.Repeat("■", pos),
		blocks-pos,
		"",
		percent*100,
		strings.Join(suffixs, ""),
	)
	fmt.Print("\r" + s)
}
