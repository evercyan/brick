package xsqlite

import (
	"gorm.io/gorm/logger"
)

// Option ...
type Option func(*Config)

// WithPassword 密码
func WithPassword(password string) Option {
	return func(c *Config) {
		c.Password = password
	}
}

// WithModels model 初始化
func WithModels(models ...interface{}) Option {
	return func(c *Config) {
		c.Models = models
	}
}

// WithLogger 日志
func WithLogger(log logger.Interface) Option {
	return func(c *Config) {
		c.Logger = log
	}
}

// WithDebug 调试模式
func WithDebug() Option {
	return func(c *Config) {
		c.Debug = true
	}
}
