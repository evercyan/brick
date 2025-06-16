package xsync

import (
	"fmt"
	"sync"
	"testing"
)

func TestLimiter(t *testing.T) {
	// qps 限制为 2, 最大容量也为 2
	lim := NewLimiter(2, 2)
	wg := new(sync.WaitGroup)
	wg.Add(10)
	// 起 10 个协程请求, 预期只输出 2 个 "execute"
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			if lim.Allow() {
				fmt.Println("execute")
			}
		}()
	}
	wg.Wait()
}
