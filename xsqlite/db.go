package xsqlite

import (
	"fmt"
	sqlcipher "github.com/gdanko/gorm-sqlcipher"
	"gorm.io/gorm"
	"net/url"
	"sync"
)

// dbMap 数据库实例
var dbMap = new(sync.Map)

// NewDB ...
func New(dbPath string, options ...Option) (*gorm.DB, error) {
	if v, ok := dbMap.Load(dbPath); ok {
		return v.(*gorm.DB), nil
	}
	cfg := defaultConfig
	for _, f := range options {
		f(cfg)
	}
	var dialector gorm.Dialector
	if cfg.Password == "" {
		dialector = sqlcipher.Open(dbPath)
	} else {
		fmt.Println(url.QueryEscape(cfg.Password))
		dbname := fmt.Sprintf("%s?_pragma_key=%s&_pragma_cipher_page_size=4096", dbPath, url.QueryEscape(cfg.Password))
		dialector = sqlcipher.Open(dbname)
	}
	db, err := gorm.Open(dialector, &gorm.Config{
		Logger: cfg.Logger,
	})
	if err != nil {
		return nil, err
	}
	if len(cfg.Models) > 0 {
		db.AutoMigrate(cfg.Models...)
	}
	if cfg.Debug {
		db = db.Debug()
	}
	dbMap.Store(dbPath, db)
	return db, nil
}
