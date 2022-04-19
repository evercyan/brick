package internal

import (
	"fmt"
)

// FormatPrice ...
func FormatPrice(v float64) string {
	return fmt.Sprintf("%.2f", v)
}
