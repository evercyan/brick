package xpool

import (
	"sync"
)

// Pool ...
type Pool struct {
	fn chan func()
	wg sync.WaitGroup
}

// Do ...
func (p *Pool) Do(f func()) {
	p.fn <- f
}

// Wait ...
func (p *Pool) Wait() {
	close(p.fn)
	p.wg.Wait()
}

// New ...
func New(max int) *Pool {
	p := &Pool{
		fn: make(chan func()),
	}
	for i := 0; i < max; i++ {
		p.wg.Add(1)
		go func() {
			defer p.wg.Done()
			for fn := range p.fn {
				fn()
			}
		}()
	}
	return p
}
