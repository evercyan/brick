package xlog

import (
	"testing"
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

func TestLogger(t *testing.T) {
	NewLogger(
		WithFile("/tmp/xlog.log", false),
		// WithLevel(LevelInfo),
		// WithLogger(TypeLogrus),
		// WithFormatter(FormatterJSON),
	)

	// zap
	// [2022-12-09 15:33:28.023] [INFO] [xlog/logger_test.go:33] what you say: [1 2 3]
	// {"level":"INFO","time":"2022-12-09 15:36:38.748","caller":"xlog/logger_test.go:37","msg":"what you say: [1 2 3]"}

	// logrus
	// [2022-12-09 15:46:41.262] [INFO] [xlog/logger_test.go:43] what you say: [1 2 3]
	// {"caller":"logger_test.go:38","level":"info","msg":"what you say: [1 2 3]","time":"2022-12-09 15:39:54.570"}

	Infof("what you say: %+v", []int{1, 2, 3})
}

func TestLogFunc(t *testing.T) {
	NewLogger(
		WithFile("/tmp/xlog.log", false),
	)
	Debug("hello")
	Debugf("hello %v", "world")
	Info("hello")
	Infof("hello %v", "world")
	Warn("hello")
	Warnf("hello %v", "world")
	Error("hello")
	Errorf("hello %v", "world")
}
