package xtime

import (
	"fmt"
	"strings"
)

// IsValid 校验日期是否有效
func IsValid(year, month, day int) bool {
	if month < 1 || month > 12 || day < 1 || day > 31 || year < 1 || year > 32767 {
		return false
	}
	switch month {
	case 4, 6, 9, 11:
		if day > 30 {
			return false
		}
	case 2:
		if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
			if day > 29 {
				return false
			}
		} else if day > 28 {
			return false
		}
	}
	return true
}

// IsLeapYear 是否是闰年
func IsLeapYear(year int) bool {
	return (year%4 == 0 && year%100 != 0) || year%400 == 0
}

// IsValidDate ...
func IsValidDate(val interface{}, patterns ...Pattern) error {
	value, ok := val.(string)
	if !ok {
		return fmt.Errorf("不是字符串")
	}
	value = strings.TrimSpace(value)
	if value == "" {
		return fmt.Errorf("不能为空")
	}
	t := ToTime(value)
	if t.IsZero() {
		return fmt.Errorf("无效日期: %s", value)
	}
	if Format(t, patterns...) != value {
		return fmt.Errorf("日期格式不满足")
	}
	return nil
}
