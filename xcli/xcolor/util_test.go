package xcolor

import (
	"testing"
)

func TestUtil(t *testing.T) {
	Success("Success:", "提示一下")
	Info("Info:", "提示一下")
	Warnning("Warnning:", "提示一下")
	Danger("Danger:", "提示一下")
	Output(FgMagenta, "Output: 提示一下")
}
