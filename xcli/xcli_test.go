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
