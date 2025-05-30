package xstock

import (
	"fmt"
	"testing"
)

func TestFormatEmCodes(t *testing.T) {
	fmt.Println(generateEMCode("002392", "600340", "000001"))
}
