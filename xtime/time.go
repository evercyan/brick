package xtime

import (
	"fmt"
	"time"
)

// Format ...
func Format(t time.Time, layout string) string {
	return t.Format(replaceLayout(layout))
}

// Parse ...
func Parse(t, layout string) (time.Time, error) {
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
