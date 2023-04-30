package xcolor

import (
	"testing"
)

func TestHelper(t *testing.T) {
	Success("提示:", "操作成功")
	Fail("提示:", "操作失败")
	Info("提示:", "提示一下")
}
