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

// Template 根据起始结束时间范围返回模板
func Template(start, end time.Time) string {
	// 如果结束时间是某一天的 0 点整, 将其减去 1 秒, 视为前一天的结束
	if end.Hour() == 0 &&
		end.Minute() == 0 &&
		end.Second() == 0 &&
		end.Nanosecond() == 0 {
		end = end.Add(-time.Second) // 减去 1 秒
	}
	// 判断是否跨年
	if start.Year() != end.Year() {
		return "2006-01-02 15:04:05"
	}
	// 判断是否跨天
	if start.YearDay() != end.YearDay() {
		return "01-02 15:04:05"
	}
	// 在同一天内
	return "15:04:05"
}
