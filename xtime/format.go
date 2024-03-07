package xtime

import (
	"time"

	"github.com/evercyan/brick/xlodash"
)

// Format ...
func Format(t time.Time, patterns ...Pattern) string {
	return t.Format(string(xlodash.First(patterns, DateTime)))
}

// Parse ...
func Parse(t string, patterns ...Pattern) (time.Time, error) {
	return time.ParseInLocation(string(xlodash.First(patterns, DateTime)), t, time.Local)
}
