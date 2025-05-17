package xutil

import (
	"fmt"
	"math"
)

// FormatPrice ...
func FormatPrice(price float64) string {
	symbol := ""
	if price < 0 {
		symbol = "-"
	}
	price = math.Abs(price)
	base := math.Pow(10, 12)
	if price >= base {
		return fmt.Sprintf("%s%.2f万亿", symbol, price/base)
	}
	base = math.Pow(10, 8)
	if price >= base {
		return fmt.Sprintf("%s%.2f亿", symbol, price/base)
	}
	base = math.Pow(10, 4)
	if price >= base {
		return fmt.Sprintf("%s%.2f万", symbol, price/base)
	}
	return fmt.Sprintf("%s%.2f", symbol, price)
}
