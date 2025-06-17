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

func TestAge(t *testing.T) {
	toTime := func(s string, layout Pattern) time.Time {
		t, _ := Parse(s, layout)
		return t
	}
	start := toTime("2024-06-12", DateOnly)
	dates := map[string]string{
		"2024-06-12": "第1天",
		"2024-06-14": "第3天",
		"2024-07-11": "第30天",
		"2024-07-12": "满月",
		"2024-07-20": "1个月零8天",
		"2024-08-12": "2个月整",
		"2025-03-20": "9个月零8天",
		"2025-06-12": "周岁",
		"2025-06-13": "1岁零1天",
		"2025-07-12": "1岁1个月整",
		"2025-07-13": "1岁1个月零1天",
	}
	for date, age := range dates {
		assert.Equal(t, age, Age(start, toTime(date, DateOnly)))
	}
}

func TestMonth(t *testing.T) {
	tt, _ := Parse("2025-01-10 12:33:00")
	assert.Equal(t, "2025-01-01 00:00:00", Format(BeginOfMonth(tt)))
	assert.Equal(t, "2025-01-31 23:59:59", Format(EndOfMonth(tt)))

}
