package xcli

import (
	"testing"
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
	//ExecCB("ping baidu.com", func(res string) {
	//	fmt.Println(strings.TrimSuffix(res, "\n"))
	//})
}
