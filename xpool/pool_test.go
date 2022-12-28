package xpool

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	p := New(10)
	go func() {
		for range time.Tick(1 * time.Second) {
			fmt.Println(fmt.Sprintf(
				"%s goroutines: %d",
				time.Now().Format("15:04:05"),
				runtime.NumGoroutine(),
			))
		}
	}()

	success := 0
	mu := sync.Mutex{}
	for i := 0; i < 100; i++ {
		p.Do(func() {
			time.Sleep(2 * time.Second)
			mu.Lock()
			defer mu.Unlock()
			success++
		})
	}
	p.Wait()
	fmt.Println("success:", success)
	/*
		00:00:30 goroutines: 13
		00:00:31 goroutines: 13
		00:00:32 goroutines: 13
		00:00:33 goroutines: 13
		00:00:34 goroutines: 13
		00:00:35 goroutines: 13
		00:00:36 goroutines: 13
		00:00:37 goroutines: 13
		00:00:38 goroutines: 13
		00:00:39 goroutines: 13
		00:00:40 goroutines: 13
		00:00:41 goroutines: 13
		00:00:42 goroutines: 13
		00:00:43 goroutines: 13
		00:00:44 goroutines: 13
		00:00:45 goroutines: 13
		00:00:46 goroutines: 13
		00:00:47 goroutines: 13
		00:00:48 goroutines: 13
		00:00:49 goroutines: 13
		success: 100
	*/
}
