package xstock

import (
	"strings"
	"time"

	"github.com/evercyan/brick/xtime"
)

// GetCode ...
func GetCode(codes ...string) string {
	for k, v := range codes {
		if strings.Contains(v, ".") {
			continue
		}
		if strings.HasPrefix(v, "6") {
			codes[k] = "1." + v
		} else {
			codes[k] = "0." + v
		}
	}
	return strings.Join(codes, ",")
}

// FormatDate ...
func FormatDate(t time.Time) string {
	return xtime.Format(t, xtime.DateJoin)
}
