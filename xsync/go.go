package xsync

import (
	"reflect"
)

// Go ...
func Go(f interface{}, args ...interface{}) {
	fn := reflect.ValueOf(f)
	if fn.Kind() != reflect.Func {
		return
	}
	fnArgs := make([]reflect.Value, 0)
	for _, v := range args {
		fnArgs = append(fnArgs, reflect.ValueOf(v))
	}
	go func() {
		defer func() {
			_ = recover()
		}()
		fn.Call(fnArgs)
	}()
}
