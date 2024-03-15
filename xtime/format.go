package xtime

import (
	"time"

	"github.com/evercyan/brick/xlodash"
)

// Format ...
func Format(t time.Time, patterns ...Pattern) string {
	return t.Format(xlodash.First(patterns, DateTime).Desc())
}

// Parse ...
func Parse(t string, patterns ...Pattern) (time.Time, error) {
	return time.ParseInLocation(xlodash.First(patterns, DateTime).Desc(), t, time.Local)
}
