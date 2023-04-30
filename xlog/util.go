package xlog

import (
	"context"
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
	"strings"
	"time"

	"github.com/evercyan/brick/xgen"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap/zapcore"
)

// loggerCaller 输出 File:Line
func loggerCaller(depth int) string {
	_, f, n, ok := runtime.Caller(depth)
	if !ok {
		return ""
	}
	dir := path.Base(path.Dir(f))
	if dir == "" {
		dir = "."
	}
	return fmt.Sprintf("%s/%s:%d", dir, path.Base(f), n)
}

// getFileWriter 获取输出文件
func getLoggerWriters(config *Config) []io.Writer {
	writers := make([]io.Writer, 0)
	if config.Filepath != "" {
		fileDir, _ := path.Split(config.Filepath)
		if err := os.MkdirAll(fileDir, os.ModePerm); err != nil {
			panic(err)
		}
		filePatten := strings.ReplaceAll(config.Filepath, ".log", ".%Y-%m-%d.log")
		fileWriter, _ := rotatelogs.New(
			filePatten,
			rotatelogs.WithLinkName(config.Filepath),
			rotatelogs.WithMaxAge(time.Hour*24*7),
			rotatelogs.WithRotationTime(time.Hour*24),
			rotatelogs.WithClock(rotatelogs.Local),
			rotatelogs.WithLocation(time.Local),
		)
		writers = append(writers, fileWriter)
	}
	if len(writers) == 0 || config.Stdout {
		writers = append(writers, os.Stdout)
	}
	return writers
}

// getLogrusLevel 获取 logrus 日志级别
func getLogrusLevel(lv Level) logrus.Level {
	level, err := logrus.ParseLevel(string(lv))
	if err != nil {
		level = logrus.InfoLevel
	}
	return level
}

// getZapLevel 获取 zap 日志级别
func getZapLevel(lv Level) zapcore.Level {
	switch string(lv) {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}

// GetTraceId ...
func GetTraceId(ctx context.Context) string {
	if ctx != nil {
		if v, ok := ctx.Value(CtxTraceKey{}).(string); ok {
			return v
		}
		if v, ok := ctx.Value(FieldTraceId).(string); ok {
			return v
		}
	}
	return xgen.Nanoid()
}
