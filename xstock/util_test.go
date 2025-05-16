package xstock

import (
	"fmt"
	"testing"
)

func TestFormatEmCodes(t *testing.T) {
	fmt.Println(GetCode("002392", "600340", "000001"))
}
