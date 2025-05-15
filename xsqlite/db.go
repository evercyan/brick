package xsqlite

import (
	"fmt"
	"net/url"
	"sync"

	sqlcipher "github.com/gdanko/gorm-sqlcipher"
	"gorm.io/gorm"
)

// dbMap 数据库实例
var dbMap = new(sync.Map)

// New ...
func New(dbPath string, options ...Option) (*gorm.DB, error) {
	cfg := defaultConfig
	for _, f := range options {
		f(cfg)
	}
	if !cfg.Force {
		if v, ok := dbMap.Load(dbPath); ok {
			return v.(*gorm.DB), nil
		}
	}
	if cfg.Password != "" {
		dbPath = fmt.Sprintf(
			"%s?_pragma_key=%s&_pragma_cipher_page_size=4096", dbPath, url.QueryEscape(cfg.Password),
		)
	}
	db, err := gorm.Open(sqlcipher.Open(dbPath), &gorm.Config{
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

// DB ...
func DB(dbPath string, options ...Option) *gorm.DB {
	db, err := New(dbPath, options...)
	if err != nil {
		panic(err)
	}
	return db
}
