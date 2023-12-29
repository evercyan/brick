package xsync

import (
	"context"
	"runtime"
	"testing"
	"time"
)

func TestNewWaitGroup(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func(ctx context.Context) {
		for range time.Tick(1 * time.Second) {
			select {
			case <-ctx.Done():
				return
			default:
				// 每秒钟打印当前协程数
				t.Logf("%s goroutines: %d", time.Now().Format("15:04:05"), runtime.NumGoroutine())
			}
		}
	}(ctx)

	// 最多起 10 个协程
	wg := NewWaitGroup(10)
	for i := 0; i < 100; i++ {
		// 差异点
		// sync.WaitGroup 支持在循环外 wg.Add(10)
		// xsync.WaitGroup 需要做协程阻塞, 必须在循环中依次 wg.Add()
		wg.Add()
		go func(index int, wg *WaitGroup) {
			defer wg.Done()
			time.Sleep(1 * time.Second)
			//t.Logf("index: %d", index)
		}(i, wg)
	}
	wg.Wait()
	t.Logf("done")
}
