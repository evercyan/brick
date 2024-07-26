package xretry

import (
	"time"
)

// Option ...
type Option struct {
	Max      int
	Timeout  time.Duration
	Interval time.Duration
}

// OptionFn ...
type OptionFn func(*Option)

// WithMax ...
func WithMax(max int) OptionFn {
	return func(o *Option) {
		o.Max = max
	}
}

// WithTimeout ...
func WithTimeout(d time.Duration) OptionFn {
	return func(o *Option) {
		o.Timeout = d
	}
}

// WithInterval ...
func WithInterval(d time.Duration) OptionFn {
	return func(o *Option) {
		o.Interval = d
	}
}
