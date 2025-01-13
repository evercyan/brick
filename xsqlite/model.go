package xsqlite

import (
	"time"
)

// Model ...
type Model struct {
	Id        int64     `json:"id" gorm:"column:id;primaryKey;AUTO_INCREMENT;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;not null" comment:"创建时间"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;not null" comment:"更新时间"`
}

// FetchAll ...
//func (t Model) FetchAll(ctx context.Context, db *grom.DB, list interface{}) error {
//	db.Model
//}
