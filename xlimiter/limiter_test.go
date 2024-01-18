package xlimiter

import (
	"fmt"
	"sync"
	"testing"
)

func TestLimiters(t *testing.T) {
	l := NewLimiters(2, 2)
	wg := new(sync.WaitGroup)
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			if l.Allow() {
				fmt.Println("1")
			} else {
				fmt.Println("0")
			}
		}()
	}
	wg.Wait()
}
