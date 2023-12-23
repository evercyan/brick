package zgo

import (
	"fmt"
	"testing"
)

func TestGo(t *testing.T) {
	fn := func(a, b string) {
		fmt.Println(fmt.Sprintf("fn a: %s, b: %s", a, b))
		return
	}
	Go(fn, "hello", "world")

	fnPanic := func(a, b string) {
		fmt.Println(fmt.Sprintf("fnPanic a: %s, b: %s", a, b))
		panic("fnPanic")
		return
	}
	Go(fnPanic, "hello", "world")
}
