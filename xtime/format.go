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

// ToTime 转换成时间, 错误时为时间零值
func ToTime(t string, patterns ...Pattern) time.Time {
	tt, err := Parse(t, patterns...)
	if err != nil {
		return time.Time{}
	}
	return tt
}
