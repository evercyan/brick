package xloading

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLoading(t *testing.T) {
	l1 := New("message1 begin").Style(Style1).Symbol(Symbol1).Speed(1).Speed(10000).Start()
	time.Sleep(time.Millisecond * 10)
	l1.Success("message1 success")

	l2 := New("message2 begin").Start()
	time.Sleep(time.Second * 2)
	l2.Fail("message2 fail")
}

func TestOption(t *testing.T) {
	for style := Style1; style <= Style5; style++ {
		assert.Less(t, 0, len(style.Elements()))
	}
	for symbol := Symbol1; symbol <= Symbol5; symbol++ {
		assert.Less(t, 0, len(symbol.Elements()))
	}
}
