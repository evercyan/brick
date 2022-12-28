package xconvert

import (
	"fmt"
	"strconv"
)

// Bin ...
func Bin(num int) string {
	s := ""
	if num == 0 {
		s = "0"
	} else {
		for ; num > 0; num /= 2 {
			lsb := num % 2
			s = strconv.Itoa(lsb) + s
		}
	}
	return fmt.Sprintf("%032s", s)
}
