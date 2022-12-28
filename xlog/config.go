package xlog

// Config ...
type Config struct {
	Type      Type      `yaml:"type"`        // 日志类型
	Level     Level     `yaml:"level"`       // 日志级别
	Formatter Formatter `yaml:"formatter"`   // 输出格式
	Filepath  string    `yaml:"output_file"` // 文件路径
	Stdout    bool      `yaml:"stdout"`      // 输出终端
}

// defaultConfig  默认配置
var defaultConfig = &Config{
	Type:      TypeZap,
	Level:     LevelDebug,
	Formatter: FormatterText,
	Filepath:  "",
	Stdout:    true,
}

// ----------------------------------------------------------------

// Type ...
type Type int

// 日志类型
const (
	TypeZap Type = iota
	TypeLogrus
)

// String ...
func (t Type) String() string {
	switch t {
	case TypeZap:
		return "zap"
	case TypeLogrus:
		return "logrus"
	default:
		return ""
	}
}

// ----------------------------------------------------------------

// Formatter ...
type Formatter string

// 格式化类型
const (
	FormatterJSON Formatter = "json"
	FormatterText Formatter = "text"
)

// ----------------------------------------------------------------

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

// ----------------------------------------------------------------

// Fields ...
type Fields map[string]interface{}

// ----------------------------------------------------------------

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

// ...
const (
	msgTimeKey       = "time"
	msgLevelKey      = "level"
	msgNameKey       = "name"
	msgCallerKey     = "caller"
	msgStacktraceKey = "stack"
	msgMessageKey    = "msg"
)

// ----------------------------------------------------------------
