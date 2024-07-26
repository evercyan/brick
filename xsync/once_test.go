package xsync

import (
	"fmt"
	"testing"
	"time"
)

func getName() string {
	return Once(func() string {
		return time.Now().Format("15:04:05")
	})()
}

func TestOnce(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(i, getName())
	}

}
