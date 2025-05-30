package xcli

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestCursor(t *testing.T) {
	ShowCursor()
	HideCursor()
	ClearLine()
}

func TestExec(t *testing.T) {
	Exec("ls")
	Shell("ls")
}

func TestExecCB(t *testing.T) {
	ExecCB("ping baidu.com", func(res string) {
		fmt.Println(strings.TrimSuffix(res, "\n"))
	})
}

func TestProgress(t *testing.T) {
	for i := 0; i <= 100; i++ {
		Progress("下载中", float64(i)/float64(100), 100, "haha")
		time.Sleep(time.Second * 1)
	}
}

func TestNotice(t *testing.T) {
	fmt.Println(Notice("Hello", "World"))
}
