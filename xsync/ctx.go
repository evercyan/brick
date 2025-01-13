package xsync

import (
	"context"
	"github.com/evercyan/brick/xgen"
	"github.com/evercyan/brick/xlog"
)

// Context ...
func Context() context.Context {
	return context.WithValue(context.Background(), xlog.FieldTraceId, xgen.Nanoid())
}
