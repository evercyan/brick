package xstock

import (
	"strings"
	"time"

	"github.com/evercyan/brick/xtime"
)

// generateEMCode 生成东财请求code
func generateEMCode(codes ...string) string {
	newCodes := make([]string, 0)
	for _, code := range codes {
		if strings.Contains(code, ".") {
			newCodes = append(newCodes, code)
			continue
		}
		if strings.HasPrefix(code, "6") {
			newCodes = append(newCodes, "1."+code)
		} else {
			newCodes = append(newCodes, "0."+code)
		}
	}
	return strings.Join(newCodes, ",")
}

// FormatDate ...
func FormatDate(t time.Time) string {
	return xtime.Format(t, xtime.DateJoin)
}

// GetMarketCode ...
func GetMarketCode(code string) string {
	if strings.HasPrefix(code, "688") || strings.HasPrefix(code, "60") {
		return "1"
	}
	return "0"
}

// GetMarketPrefix ...
func GetMarketPrefix(code string) string {
	if strings.HasPrefix(code, "6") {
		return "sh"
	}
	if strings.HasPrefix(code, "0") || strings.HasPrefix(code, "3") {
		return "sz"
	}
	return ""
}

// GetSymbol ...
func GetSymbol(v float64) string {
	if v > 0 {
		return "↑"
	} else if v < 0 {
		return "↓"
	}
	return ""
}
