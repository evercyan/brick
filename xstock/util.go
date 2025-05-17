package xstock

import (
	"strings"
	"time"

	"github.com/evercyan/brick/xtime"
)

// GetCode ...
func GetCode(codes ...string) string {
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
