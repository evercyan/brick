package zgo

import (
	"math"
	"sync"

	"github.com/evercyan/brick/xlodash"
)

// WaitGroup 基于 sync.WaitGroup 增加协程池处理
type WaitGroup struct {
	ch chan struct{}
	wg sync.WaitGroup
}

// NewWaitGroup ...
func NewWaitGroup(args ...int) *WaitGroup {
	size := xlodash.First(args, math.MaxInt)
	return &WaitGroup{
		ch: make(chan struct{}, size),
		wg: sync.WaitGroup{},
	}
}

// Add ...
func (t *WaitGroup) Add() {
	t.ch <- struct{}{}
	t.wg.Add(1)
}

// Done ...
func (t *WaitGroup) Done() {
	<-t.ch
	t.wg.Done()
}

// Wait ...
func (t *WaitGroup) Wait() {
	t.wg.Wait()
}
