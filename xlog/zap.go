package xlog

import (
	"fmt"
	"io"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// xZap ...
type xZap struct {
	entry  *zap.Logger
	writer io.Writer
}

// newZap ...
func newZap(config *Config) Logger {
	// 日志输出
	writers := getLoggerWriters(config)
	writeSyncers := make([]zapcore.WriteSyncer, 0)
	for _, writer := range writers {
		writeSyncers = append(writeSyncers, zapcore.AddSync(writer))
	}
	writer := zapcore.NewMultiWriteSyncer(writeSyncers...)

	// 日志级别
	level := getZapLevel(config.Level)

	// 根据 Formatter 来做格式化
	getFormatterValue := func(s string) string {
		if config.Formatter == FormatterJSON {
			return s
		}
		return fmt.Sprintf("[%s]", s)
	}
	// 日志格式
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:       msgTimeKey,
		LevelKey:      msgLevelKey,
		NameKey:       msgNameKey,
		CallerKey:     msgCallerKey,
		MessageKey:    msgMessageKey,
		StacktraceKey: msgStacktraceKey,
		LineEnding:    zapcore.DefaultLineEnding,
		// LowercaseLevelEncoder 默认, 小写编码器
		// LowercaseColorLevelEncoder 小写编码器带颜色
		// CapitalLevelEncoder 大写编码器
		// CapitalColorLevelEncoder 大写编码器带颜色
		// EncodeLevel: zapcore.LowercaseLevelEncoder,
		EncodeLevel: func(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(getFormatterValue(level.CapitalString()))
		},
		// 时间显示格式 zapcore.ISO8601TimeEncoder
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(getFormatterValue(t.Format(timestampFormat)))
		},
		EncodeDuration: zapcore.StringDurationEncoder,
		// FullCallerEncoder 显示完整路径
		// ShortCallerEncoder 显示短路径
		// EncodeCaller: zapcore.ShortCallerEncoder,
		EncodeCaller: func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(getFormatterValue(caller.TrimmedPath()))
		},
		EncodeName: zapcore.FullNameEncoder,
		// 各元素间隔区域
		ConsoleSeparator: " ",
	}
	var encoder zapcore.Encoder
	if config.Formatter == FormatterJSON {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	core := zapcore.NewCore(encoder, writer, level)
	zl := zap.New(core)

	zl = zl.WithOptions(zap.AddCaller())
	zl = zl.WithOptions(zap.AddCallerSkip(callerSkip - 1))

	return &xZap{
		entry:  zl,
		writer: writers[0],
	}
}

// ----------------------------------------------------------------

// log ...
func (t *xZap) log(level Level, args string) {
	lv := getZapLevel(level)
	if !t.entry.Core().Enabled(lv) {
		return
	}

	switch lv {
	case zapcore.DebugLevel:
		t.entry.Debug(args)
	case zapcore.InfoLevel:
		t.entry.Info(args)
	case zapcore.WarnLevel:
		t.entry.Warn(args)
	case zapcore.ErrorLevel:
		t.entry.Error(args)
	case zapcore.PanicLevel:
		t.entry.Panic(args)
	case zapcore.FatalLevel:
		t.entry.Fatal(args)
	default:
		t.entry.Info(args)
	}
}

// Log ...
func (t *xZap) Log(level Level, args ...interface{}) {
	t.log(level, strings.TrimSpace(fmt.Sprintf(strings.Repeat("%v ", len(args)), args...)))
}

// Logf ...
func (t *xZap) Logf(level Level, format string, args ...interface{}) {
	t.log(level, fmt.Sprintf(format, args...))
}

// Trace ...
func (t *xZap) Trace(args ...interface{}) {
	t.Log(LevelTrace, args...)
}

// Tracef ...
func (t *xZap) Tracef(format string, args ...interface{}) {
	t.Logf(LevelTrace, format, args...)
}

// Debug ...
func (t *xZap) Debug(args ...interface{}) {
	t.Log(LevelDebug, args...)
}

// Debugf ...
func (t *xZap) Debugf(format string, args ...interface{}) {
	t.Logf(LevelDebug, format, args...)
}

// Info ...
func (t *xZap) Info(args ...interface{}) {
	t.Log(LevelInfo, args...)
}

// Infof ...
func (t *xZap) Infof(format string, args ...interface{}) {
	t.Logf(LevelInfo, format, args...)
}

// Warn ...
func (t *xZap) Warn(args ...interface{}) {
	t.Log(LevelWarn, args...)
}

// Warnf ...
func (t *xZap) Warnf(format string, args ...interface{}) {
	t.Logf(LevelWarn, format, args...)
}

// Error ...
func (t *xZap) Error(args ...interface{}) {
	t.Log(LevelError, args...)
}

// Errorf ...
func (t *xZap) Errorf(format string, args ...interface{}) {
	t.Logf(LevelError, format, args...)
}

// Fatal ...
func (t *xZap) Fatal(args ...interface{}) {
	t.Log(LevelFatal, args...)
}

// Fatalf ...
func (t *xZap) Fatalf(format string, args ...interface{}) {
	t.Logf(LevelFatal, format, args...)
}

// Panic ...
func (t *xZap) Panic(args ...interface{}) {
	t.Log(LevelPanic, args...)
}

// Panicf ...
func (t *xZap) Panicf(format string, args ...interface{}) {
	t.Logf(LevelPanic, format, args...)
}

// WithField ...
func (t *xZap) WithField(key string, value interface{}) Logger {
	return &xZap{
		entry:  t.entry.With(zap.Any(key, value)),
		writer: t.writer,
	}
}

// Writer ...
func (t *xZap) Writer() io.Writer {
	return t.writer
}
