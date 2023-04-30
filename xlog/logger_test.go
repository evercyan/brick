package xlog

import (
	"context"
	"testing"

	"github.com/evercyan/brick/xfile"
	"github.com/evercyan/brick/xgen"
	"github.com/stretchr/testify/assert"
)

// BenchmarkLogrus-8   	   12340	    103458 ns/op	    5854 B/op	      86 allocs/op
func BenchmarkLogrus(b *testing.B) {
	log := newLogrus(defaultConfig)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		log.Info("hello world")
		log.Infof("hello %s", "world")
	}
}

// BenchmarkZap-8   	   26472	     38800 ns/op	    1445 B/op	      23 allocs/op
func BenchmarkZap(b *testing.B) {
	log := newZap(defaultConfig)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		log.Info("hello world")
		log.Infof("hello %s", "world")
	}
}

// ----------------------------------------------------------------

func TestLogger(t *testing.T) {
	// 默认配置见 defaultConfig
	NewLogger(
		WithZap(),
		WithLevel(LevelInfo),
		WithStdout(true),
		WithFormatter(FormatterText),
	)

	// [2023-04-30 13:47:45.897] [INFO] [LKAohqdCZfgEevTpi9URl] [xlog/logger_test.go:49] hello: [1 2]
	Infof("hello: %+v", []int{1, 2})
}

func TestZap(t *testing.T) {
	NewLogger(
		WithZap(),
	)

	// Func
	Trace("zap hello")
	Tracef("zap hello %v", "world")
	Debug("zap hello")
	Debugf("zap hello %v", "world")
	Info("zap hello")
	Infof("zap hello %v", "world")
	Warn("zap hello")
	Warnf("zap hello %v", "world")
	Error("zap hello")
	Errorf("zap hello %v", "world")
	// Write
	assert.NotNil(t, Writer())

	// JSON
	NewLogger(WithFormatter(FormatterJSON))
	// {"level":"INFO","time":"2023-04-30 13:49:56.609","trace_id":"tJ3glSOBR5xVvs0jrzLQ1","caller":"xlog/logger_test.go:77","msg":"json: [1 2]"}
	Infof("zap json: %+v", []int{1, 2})

	// Text
	NewLogger(WithFormatter(FormatterText))
	// [2023-01-01 13:43:30.312] [INFO] [1fPZBJRNdko1y3lVr5XEv] [xlog/logger_test.go:44] hello: [1 2]
	Infof("zap text: %+v", []int{1, 2})

	// Ctx
	ctx := context.WithValue(context.Background(), CtxTraceKey{}, xgen.Nanoid())
	C(ctx).Info("zap ctx 1")
	C(ctx).Error("zap ctx 2")

	// Field
	F("k", "v").Info("zap field")
}

func TestLogrus(t *testing.T) {
	NewLogger(
		WithLogrus(),
	)

	// Func
	Trace("logrus hello")
	Tracef("logrus hello %v", "world")
	Debug("logrus hello")
	Debugf("logrus hello %v", "world")
	Info("logrus hello")
	Infof("logrus hello %v", "world")
	Warn("logrus hello")
	Warnf("logrus hello %v", "world")
	Error("logrus hello")
	Errorf("logrus hello %v", "world")
	// Write
	assert.NotNil(t, Writer())

	// JSON
	NewLogger(WithLogrus(), WithFormatter(FormatterJSON))
	// {"caller":"xlog/logger_test.go:114","level":"info","msg":"json: [1 2]","time":"2023-04-30 13:52:04.883","trace_id":"AQAbZgFtdhk0vfgesUofo"}
	Infof("logrus json: %+v", []int{1, 2})

	// Text
	NewLogger(WithLogrus(), WithFormatter(FormatterText))
	// [2023-04-30 13:52:56.718] [INFO] [lx_pkNU7472dupxeSmAa-] [xlog/logger_test.go:109] text: [1 2]
	Infof("logrus text: %+v", []int{1, 2})

	// Ctx
	ctx := context.WithValue(context.Background(), FieldTraceId, xgen.Nanoid())
	C(ctx).Info("logrus ctx 1")
	C(ctx).Error("logrus ctx 2")

	// Field
	F("k", "v").Info("logrus field")
}

func TestLoggerFile(t *testing.T) {
	NewLogger(WithFile(xfile.Temp("brick.log")))
}
