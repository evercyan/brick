package xtime

import (
	"fmt"
	"strings"
	"time"

	"github.com/evercyan/brick/xtype"
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

// Diff ...
func Diff(start, end time.Time) (int, int, int) {
	start = start.UTC()
	end = end.UTC()
	if start.After(end) {
		start, end = end, start
	}
	// 计算年
	years := end.Year() - start.Year()
	temp := start.AddDate(years, 0, 0)
	if temp.After(end) {
		years--
		temp = start.AddDate(years, 0, 0)
	}
	// 计算月
	months := 0
	for {
		next := temp.AddDate(0, months+1, 0)
		if next.After(end) {
			break
		}
		months++
	}
	temp = temp.AddDate(0, months, 0)
	// 计算天
	days := int(end.Sub(temp).Hours() / 24)
	return years, months, days
}

// Age 计算年龄
// 当天
// 第2天
// 满月
// 1个月零1天
// 2个月整
// 2个月零28天
// 周岁
// 1岁1个月整
// 1岁1个月零1天
func Age(start, end time.Time) string {
	years, months, days := Diff(start, end)
	parts := make([]string, 0)
	if years > 0 {
		parts = append(parts, fmt.Sprintf("%d岁", years))
	}
	if months > 0 {
		parts = append(parts, fmt.Sprintf("%d个月", months))
	}
	if days > 0 {
		prefix := "零"
		if years == 0 && months == 0 {
			// 未满月时, 当天算第 1 天, 隔天算第 2 天, 故补 +1
			prefix = "第"
			days += 1
		}
		parts = append(parts, fmt.Sprintf("%s%d天", prefix, days))
	}
	if len(parts) == 0 {
		return "第1天"
	}
	if years == 0 && months == 1 && days == 0 {
		return "满月"
	}
	if years == 1 && months == 0 && days == 0 {
		return "周岁"
	}
	if days == 0 && (years > 0 || months > 0) {
		return strings.Join(parts, "") + "整"
	}
	return strings.Join(parts, "")
}

var ToTime = xtype.ToTime
