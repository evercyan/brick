package xlog

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/evercyan/brick/xjson"
	"github.com/sirupsen/logrus"
)

// xLogrus ...
type xLogrus struct {
	entry  *logrus.Entry
	writer io.Writer
	caller int
}

// newLogrus ...
func newLogrus(option *Config) Logger {
	log := logrus.New()

	// 日志输出
	writers := getLoggerWriters(option)
	log.SetOutput(io.MultiWriter(writers...))

	// 日志级别
	log.SetLevel(getLogrusLevel(option.Level))

	// 日志格式
	if option.Formatter == FormatterJSON {
		log.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: timestampFormat,
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyTime:  msgTimeKey,
				logrus.FieldKeyLevel: msgLevelKey,
			},
		})
	} else {
		// log.SetFormatter(&logrus.TextFormatter{
		// 	TimestampFormat: timestampFormat,
		// })
		log.SetFormatter(&logrusTextFormatter{})
	}

	return &xLogrus{
		entry:  logrus.NewEntry(log),
		writer: writers[0],
	}
}

// ----------------------------------------------------------------

type logrusTextFormatter struct{}

func (t *logrusTextFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	caller := entry.Data["caller"]
	delete(entry.Data, "caller")

	traceId := entry.Data["trace_id"]
	delete(entry.Data, "trace_id")

	message := ""
	if len(entry.Data) > 0 {
		message += xjson.Encode(entry.Data) + " "
	}
	message += entry.Message
	b.WriteString(fmt.Sprintf(
		"[%s] [%s] [%s] [%s] %s\n",
		entry.Time.Format(timestampFormat),
		strings.ToUpper(entry.Level.String()),
		traceId,
		caller,
		message,
	))
	return b.Bytes(), nil
}

// ----------------------------------------------------------------

// log ...
func (t *xLogrus) log(level Level, args ...interface{}) {
	lv := getLogrusLevel(level)
	if !t.entry.Logger.IsLevelEnabled(lv) {
		return
	}
	skip := callerSkip
	if t.caller > 0 {
		skip = t.caller
	}
	t.entry.
		WithField(msgCallerKey, loggerCaller(skip)).
		WithField(msgNameKey, GetTraceId(t.entry.Context)).
		Log(lv, args...)
}

// Log ...
func (t *xLogrus) Log(level Level, args ...interface{}) {
	t.log(level, args...)
}

// Logf ...
func (t *xLogrus) Logf(level Level, format string, args ...interface{}) {
	t.log(level, fmt.Sprintf(format, args...))
}

// Trace ...
func (t *xLogrus) Trace(args ...interface{}) {
	t.Log(LevelTrace, args...)
}

// Tracef ...
func (t *xLogrus) Tracef(format string, args ...interface{}) {
	t.Logf(LevelTrace, format, args...)
}

// Debug ...
func (t *xLogrus) Debug(args ...interface{}) {
	t.Log(LevelDebug, args...)
}

// Debugf ...
func (t *xLogrus) Debugf(format string, args ...interface{}) {
	t.Logf(LevelDebug, format, args...)
}

// Info ...
func (t *xLogrus) Info(args ...interface{}) {
	t.Log(LevelInfo, args...)
}

// Infof ...
func (t *xLogrus) Infof(format string, args ...interface{}) {
	t.Logf(LevelInfo, format, args...)
}

// Warn ...
func (t *xLogrus) Warn(args ...interface{}) {
	t.Log(LevelWarn, args...)
}

// Warnf ...
func (t *xLogrus) Warnf(format string, args ...interface{}) {
	t.Logf(LevelWarn, format, args...)
}

// Error ...
func (t *xLogrus) Error(args ...interface{}) {
	t.Log(LevelError, args...)
}

// Errorf ...
func (t *xLogrus) Errorf(format string, args ...interface{}) {
	t.Logf(LevelError, format, args...)
}

//// Fatal ...
//func (t *xLogrus) Fatal(args ...interface{}) {
//	t.Log(LevelFatal, args...)
//}
//
//// Fatalf ...
//func (t *xLogrus) Fatalf(format string, args ...interface{}) {
//	t.Logf(LevelFatal, format, args...)
//}
//
//// Panic ...
//func (t *xLogrus) Panic(args ...interface{}) {
//	t.Log(LevelPanic, args...)
//}
//
//// Panicf ...
//func (t *xLogrus) Panicf(format string, args ...interface{}) {
//	t.Logf(LevelPanic, format, args...)
//}

// Field WithField ...
func (t *xLogrus) Field(key string, value interface{}) Logger {
	return &xLogrus{
		entry:  t.entry.WithField(key, value),
		writer: t.writer,
		caller: callerSkip - 1,
	}
}

// Ctx WithContext ...
func (t *xLogrus) Ctx(ctx context.Context) Logger {
	return &xLogrus{
		entry:  t.entry.WithContext(ctx),
		writer: t.writer,
		caller: callerSkip - 1,
	}
}

// Writer ...
func (t *xLogrus) Writer() io.Writer {
	return t.writer
}
