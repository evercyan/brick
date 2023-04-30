package xlog

import (
	"context"
	"io"
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
	//Fatal(args ...interface{})                            // 记录 FatalLevel 级别的日志
	//Fatalf(format string, args ...interface{})            // 格式化并记录 FatalLevel 级别的日志
	//Panic(args ...interface{})                            // 记录 PanicLevel 级别的日志
	//Panicf(format string, args ...interface{})            // 格式化并记录 PanicLevel 级别的日志
	Writer() io.Writer                      // 返回日志 io.Writer
	F(key string, value interface{}) Logger // WithField 为日志添加一个上下文数据
	C(ctx context.Context) Logger           // WithContext 为日志添加一个 context
}

// ----------------------------------------------------------------

var (
	logger = newZap(defaultConfig)
)

// NewLogger ...
func NewLogger(options ...Option) Logger {
	config := defaultConfig
	for _, f := range options {
		f(config)
	}
	if len(options) == 0 {
		return logger
	}
	switch config.Type {
	case TypeLogrus:
		logger = newLogrus(config)
	default:
		logger = newZap(config)
	}
	return logger
}

// ----------------------------------------------------------------

// PS: 使用 logger.Log(LevelXXX, args...) 而非 logger.Info(args...)
// 是为了保持 caller skip 同其余调用保持一致

// Trace ...
func Trace(args ...interface{}) {
	logger.Trace(args...)
}

// Tracef ...
func Tracef(format string, args ...interface{}) {
	logger.Tracef(format, args...)
}

// Debug ...
func Debug(args ...interface{}) {
	logger.Debug(args...)
}

// Debugf ...
func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

// Info ...
func Info(args ...interface{}) {
	logger.Info(args...)
}

// Infof ...
func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

// Warn ...
func Warn(args ...interface{}) {
	logger.Warn(args...)
}

// Warnf ...
func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

// Error ...
func Error(args ...interface{}) {
	logger.Error(args...)
}

// Errorf ...
func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

//// Fatal ...
//func Fatal(args ...interface{}) {
//	logger.Fatal(args...)
//}
//
//// Fatalf ...
//func Fatalf(format string, args ...interface{}) {
//	logger.Fatalf(format, args...)
//}
//
//// Panic ...
//func Panic(args ...interface{}) {
//	logger.Panic(args...)
//}
//
//// Panicf ...
//func Panicf(format string, args ...interface{}) {
//	logger.Panicf(format, args...)
//}

// F WithField...
func F(key string, value interface{}) Logger {
	return logger.F(key, value)
}

// C WithContext...
func C(ctx context.Context) Logger {
	return logger.C(ctx)
}

// Writer ...
func Writer() io.Writer {
	return logger.Writer()
}
