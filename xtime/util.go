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

// PastDaysInWeek 计算给定日期所在周已过天数
func PastDaysInWeek(t time.Time) int {
	now := time.Now()
	todayBegin := time.Unix(BeginOfDay(now), 0)
	todayEnd := time.Unix(EndOfDay(now), 0)
	weekday := int(todayBegin.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	weekStartDay := todayBegin.AddDate(0, 0, (weekday-1)*-1)
	weekEndDay := todayEnd.AddDate(0, 0, 7-weekday)
	if t.Before(weekStartDay) {
		return 7
	}
	if t.After(weekEndDay) {
		return 0
	}
	return weekday
}
