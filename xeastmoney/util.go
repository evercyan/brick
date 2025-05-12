package xapi

import (
	"strings"
	"time"

	"github.com/evercyan/brick/xtime"
)

// GetCode ...
func GetCode(codes ...string) string {
	list := make([]string, 0, len(codes))
	for _, code := range codes {
		if strings.HasPrefix(code, "6") {
			list = append(list, "1."+code)
		} else {
			list = append(list, "0."+code)
		}
	}
	return strings.Join(list, ",")
}

// FormatDate ...
func FormatDate(t time.Time) string {
	return xtime.Format(t, xtime.DateJoin)
}
