package xretry

import (
	"time"
)

// Do ...
func Do(fn func() bool, options ...OptionFn) bool {
	o := &Option{
		Max:      3,
		Interval: time.Millisecond * 100,
		Timeout:  0,
	}
	for _, v := range options {
		v(o)
	}
	if o.Max <= 0 {
		o.Max = 3
	}
	now := time.Now()
	for i := 1; i <= o.Max; i++ {
		ok := fn()
		if ok {
			return true
		}
		if o.Timeout > 0 && time.Since(now) > o.Timeout {
			break
		}
		if o.Interval > 0 {
			time.Sleep(o.Interval)
		}
	}
	return false
}
