package xlog

import (
	"io"
	"sync"
)

// Logger ...
type Logger interface {
	Log(level Level, args ...interface{})                 // 记录对应级别的日志
	Logf(level Level, format string, args ...interface{}) // 记录对应级别的日志
	Trace(args ...interface{})                            // 记录 TraceLevel 级别的日志
	Tracef(format string, args ...interface{})            // 格式化并记录 TraceLevel 级别的日志
	Debug(args ...interface{})                            // 记录 DebugLevel 级别的日志
	Debugf(format string, args ...interface{})            // 格式化并记录 DebugLevel 级别的日志
	Info(args ...interface{})                             // 记录 InfoLevel 级别的日志
	Infof(format string, args ...interface{})             // 格式化并记录 InfoLevel 级别的日志
	Warn(args ...interface{})                             // 记录 WarnLevel 级别的日志
	Warnf(format string, args ...interface{})             // 格式化并记录 WarnLevel 级别的日志
	Error(args ...interface{})                            // 记录 ErrorLevel 级别的日志
	Errorf(format string, args ...interface{})            // 格式化并记录 ErrorLevel 级别的日志
	Fatal(args ...interface{})                            // 记录 FatalLevel 级别的日志
	Fatalf(format string, args ...interface{})            // 格式化并记录 FatalLevel 级别的日志
	Panic(args ...interface{})                            // 记录 PanicLevel 级别的日志
	Panicf(format string, args ...interface{})            // 格式化并记录 PanicLevel 级别的日志
	WithField(key string, value interface{}) Logger       // 为日志添加一个上下文数据
	Writer() io.Writer                                    // 返回日志 io.Writer
}

// ----------------------------------------------------------------

var (
	logger = newZap(defaultConfig)
	once   sync.Once
)

// NewLogger ...
func NewLogger(options ...Option) Logger {
	config := defaultConfig
	for _, f := range options {
		f(config)
	}
	// 无配置的情况, 不做单例处理
	if len(options) == 0 {
		return logger
	}
	once.Do(func() {
		switch config.Type {
		case TypeLogrus:
			logger = newLogrus(config)
		default:
			logger = newZap(config)
		}
	})
	return logger
}

// ----------------------------------------------------------------

// PS: 使用 logger.Log(LevelXXX, args...) 而非 logger.Info(args...)
// 是为了保持 caller skip 同其余调用保持一致

// Trace ...
func Trace(args ...interface{}) {
	logger.Log(LevelTrace, args...)
}

// Tracef ...
func Tracef(format string, args ...interface{}) {
	logger.Logf(LevelTrace, format, args...)
}

// Debug ...
func Debug(args ...interface{}) {
	logger.Log(LevelDebug, args...)
}

// Debugf ...
func Debugf(format string, args ...interface{}) {
	logger.Logf(LevelDebug, format, args...)
}

// Info ...
func Info(args ...interface{}) {
	logger.Log(LevelInfo, args...)
}

// Infof ...
func Infof(format string, args ...interface{}) {
	logger.Logf(LevelInfo, format, args...)
}

// Warn ...
func Warn(args ...interface{}) {
	logger.Log(LevelWarn, args...)
}

// Warnf ...
func Warnf(format string, args ...interface{}) {
	logger.Logf(LevelWarn, format, args...)
}

// Error ...
func Error(args ...interface{}) {
	logger.Log(LevelError, args...)
}

// Errorf ...
func Errorf(format string, args ...interface{}) {
	logger.Logf(LevelError, format, args...)
}

// Fatal ...
func Fatal(args ...interface{}) {
	logger.Log(LevelFatal, args...)
}

// Fatalf ...
func Fatalf(format string, args ...interface{}) {
	logger.Logf(LevelFatal, format, args...)
}

// Panic ...
func Panic(args ...interface{}) {
	logger.Log(LevelPanic, args...)
}

// Panicf ...
func Panicf(format string, args ...interface{}) {
	logger.Logf(LevelPanic, format, args...)
}

// Writer ...
func Writer() io.Writer {
	return logger.Writer()
}
