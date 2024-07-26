package xsync

import (
	"sync"
)

// Once ...
func Once[T any](fn func() T) func() T {
	var (
		once sync.Once
		res  T
	)
	return func() T {
		once.Do(func() {
			defer func() {
				_ = recover()
			}()
			res = fn()
		})
		return res
	}
}
