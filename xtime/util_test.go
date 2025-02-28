package xtime

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

//func TestDay(t *testing.T) {
//	tt, _ := Parse("2021-01-01 15:04:05")
//	// 2021-01-01 00:00:00
//	assert.Equal(t, int64(1609430400), BeginOfDay(tt))
//	// 2021-01-01 23:59:59
//	assert.Equal(t, int64(1609516799), EndOfDay(tt))
//}

func TestPastDaysInWeek(t *testing.T) {
	now := time.Now()

	assert.Equal(t, 7, PastDaysInWeek(now.AddDate(0, 0, -11)))
	assert.Equal(t, 0, PastDaysInWeek(now.AddDate(0, 0, 11)))

	weekday := int(now.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	assert.Equal(t, 7, PastDaysInWeek(now))
}
