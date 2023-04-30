package xlog

// ...
type (
	CtxTraceKey struct{}
)

// Fields ...
type Fields map[string]interface{}

// ...
const (
	FieldTraceId = "trace_id"
)

// ...
const (
	msgTimeKey       = "time"
	msgLevelKey      = "level"
	msgNameKey       = "trace_id"
	msgCallerKey     = "caller"
	msgStacktraceKey = "stack"
	msgMessageKey    = "msg"
)

const (
	// 调用栈过滤深度
	// 以 Info() 为例:
	// 1: helper.go runtime.Caller(depth)
	// 2: logrus.go loggerCaller(callerStackDepth)
	// 3: logrus.go t.log(level, args...)
	// 4: logrus.go t.Log(LevelInfo, args...)
	// 5: 调用方...
	callerSkip = 4

	timestampFormat = "2006-01-02 15:04:05.000"
)

// ----------------------------------------------------------------

// Type ...
type Type int

// 日志类型
const (
	TypeZap Type = iota
	TypeLogrus
)

// Formatter ...
type Formatter string

// 格式化类型
const (
	FormatterJSON Formatter = "json"
	FormatterText Formatter = "text"
)

// Level ...
type Level string

// 日志级别
const (
	LevelTrace Level = "trace"
	LevelDebug Level = "debug"
	LevelInfo  Level = "info"
	LevelWarn  Level = "warn"
	LevelError Level = "error"
	LevelFatal Level = "fatal"
	LevelPanic Level = "panic"
)
