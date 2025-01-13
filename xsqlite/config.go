package xsqlite

import (
	"gorm.io/gorm/logger"
)

// Config ...
type Config struct {
	Password string           `json:"password"`
	Models   []interface{}    `json:"models"`
	Logger   logger.Interface `json:"logger"`
	Debug    bool             `json:"debug"`
}

// defaultConfig  默认配置
var defaultConfig = &Config{
	Password: "",
	Models:   nil,
	Logger:   logger.Default.LogMode(logger.Info),
	Debug:    false,
}
