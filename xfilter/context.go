package xfilter

import (
	"context"
	"sync"
)

// Context filter 上下文
type Context struct {
	ctx  context.Context
	data *sync.Map
}

// Set 写入数据
func (t *Context) Set(k string, v interface{}) {
	t.data.Store(k, v)
}

// Get 读取数据
func (t *Context) Get(k string) (interface{}, bool) {
	return t.data.Load(k)
}

// Ctx 读取 ctx 数据
func (t *Context) Ctx(k string) interface{} {
	return t.ctx.Value(k)
}

// ----------------------------------------------------------------
// ----------------------------------------------------------------
// ----------------------------------------------------------------

// NewContext ...
func NewContext(ctxs ...context.Context) *Context {
	var ctx context.Context
	if len(ctxs) == 0 {
		ctx = context.Background()
	} else {
		ctx = ctxs[0]
	}
	return &Context{
		ctx:  ctx,
		data: new(sync.Map),
	}
}
