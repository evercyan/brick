package xutil

import (
	"fmt"
	"math"
)

// FormatPrice ...
func FormatPrice(price float64) string {
	base := math.Pow(10, 12)
	if price >= base {
		return fmt.Sprintf("%.2f万亿", price/base)
	}
	base = math.Pow(10, 8)
	if price >= base {
		return fmt.Sprintf("%.2f亿", price/base)
	}
	base = math.Pow(10, 8)
	if price >= base {
		return fmt.Sprintf("%.2f万", price/base)
	}
	return fmt.Sprintf("%.2f", price)
}
