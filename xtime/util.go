package xtime

import (
	"fmt"
	"strings"
	"time"
)

var (
	layoutMap = map[string]string{
		"y": "2006",
		"m": "01",
		"d": "02",
		"h": "15",
		"i": "04",
		"s": "05",
	}
)

// replaceLayout ...
func replaceLayout(layout string) string {
	layout = strings.ToLower(layout)
	for k, v := range layoutMap {
		layout = strings.ReplaceAll(layout, k, v)
	}
	return layout
}

// ----------------------------------------------------------------

// Format ...
func Format(t time.Time, layouts ...string) string {
	layout := "y-m-d h:i:s"
	if len(layouts) > 0 {
		layout = layouts[0]
	}
	return t.Format(replaceLayout(layout))
}

// Parse ...
func Parse(t string, layouts ...string) (time.Time, error) {
	layout := "y-m-d h:i:s"
	if len(layouts) > 0 {
		layout = layouts[0]
	}
	location, _ := time.LoadLocation("Asia/Shanghai")
	return time.ParseInLocation(replaceLayout(layout), t, location)
}

// First ...
func First(t time.Time) int64 {
	if t.IsZero() {
		t = time.Now()
	}
	timeStr := fmt.Sprintf("%s 00:00:00", Format(t, "y-m-d"))
	tt, _ := Parse(timeStr, "y-m-d h:i:s")
	return tt.Unix()
}

// Last ...
func Last(t time.Time) int64 {
	return First(t) + 24*3600 - 1
}

// Check 校验日期有效性
func Check(year, month, day int) bool {
	if month < 1 || month > 12 || day < 1 || day > 31 || year < 1 || year > 32767 {
		return false
	}
	switch month {
	case 4, 6, 9, 11:
		if day > 30 {
			return false
		}
	case 2:
		if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
			if day > 29 {
				return false
			}
		} else if day > 28 {
			return false
		}
	}
	return true
}

// IsLeapYear 是否是闰年
func IsLeapYear(year int) bool {
	return (year%4 == 0 && year%100 != 0) || (year%400 == 0)
}
