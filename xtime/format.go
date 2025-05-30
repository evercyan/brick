package xtime

import (
	"time"

	"github.com/evercyan/brick/xlodash"
)

// Now ...
func Now(patterns ...Pattern) string {
	return time.Now().Format(xlodash.First(patterns, DateTime).Desc())
}

// Format ...
func Format(t time.Time, patterns ...Pattern) string {
	return t.Format(xlodash.First(patterns, DateTime).Desc())
}

// Parse ...
func Parse(t string, patterns ...Pattern) (time.Time, error) {
	return time.ParseInLocation(xlodash.First(patterns, DateTime).Desc(), t, time.Local)
}

// FormatD ...
func FormatD(t time.Time) string {
	return Format(t, DateOnly)
}

// FormatDJ ...
func FormatDJ(t time.Time) string {
	return Format(t, DateJoin)
}

// IsToday ...
func IsToday(t time.Time) bool {
	return Format(t, DateOnly) == Format(time.Now(), DateOnly)
}
