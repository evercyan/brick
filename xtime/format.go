package xtime

import (
	"fmt"
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

// D ...
func D(t time.Time) string {
	return Format(t, DateOnly)
}

// DJ ...
func DJ(t time.Time) string {
	return Format(t, DateJoin)
}

// DT ...
func DT(t time.Time) string {
	return Format(t, DateTime)
}

// DTJ ...
func DTJ(t time.Time) string {
	return Format(t, DateTimeJoin)
}

// IsToday ...
func IsToday(t time.Time) bool {
	return Format(t, DateOnly) == Format(time.Now(), DateOnly)
}

// T ...
func T(t time.Time) string {
	return Format(t, TimeOnly)
}

// FormatDuration ...
func FormatDuration(d time.Duration) string {
	h := d / time.Hour
	m := (d % time.Hour) / time.Minute
	s := (d % time.Minute) / time.Second
	return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
}
