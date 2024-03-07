package xtime

import (
	"time"
)

// BeginOfDay ...
func BeginOfDay(t time.Time) int64 {
	if t.IsZero() {
		t = time.Now()
	}
	return t.Unix() - int64(t.Hour())*3600 - int64(t.Minute())*60 - int64(t.Second())
}

// EndOfDay ...
func EndOfDay(t time.Time) int64 {
	return BeginOfDay(t) + 24*3600 - 1
}
