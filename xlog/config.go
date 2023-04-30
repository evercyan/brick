package xlog

// ----------------------------------------------------------------

// Config ...
type Config struct {
	Type      Type      `yaml:"type"`      // 日志类型
	Level     Level     `yaml:"level"`     // 日志级别
	Formatter Formatter `yaml:"formatter"` // 输出格式
	Filepath  string    `yaml:"filepath"`  // 文件路径
	Stdout    bool      `yaml:"stdout"`    // 输出终端
}

// defaultConfig  默认配置
var defaultConfig = &Config{
	Type:      TypeZap,
	Level:     LevelInfo,
	Formatter: FormatterText,
	Stdout:    true,
}
