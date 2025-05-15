package xsqlite

import (
	"context"
	"time"

	"github.com/evercyan/brick/xlodash"
	"github.com/evercyan/brick/xtype"
	"gorm.io/gorm"
)

// Model ...
type Model struct {
	Id        int64     `json:"id" gorm:"column:id;primaryKey;AUTO_INCREMENT;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;not null" comment:"创建时间"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;not null" comment:"更新时间"`
}

// Insert ...
func Insert(ctx context.Context, db *gorm.DB, list interface{}, sizes ...int) error {
	return db.CreateInBatches(list, xlodash.First(sizes, 100)).Error
}

// Fetch ...
func Fetch(
	ctx context.Context,
	db *gorm.DB,
	query map[string]interface{},
	list interface{},
) error {
	session := db.Where("1=1")
	if v, ok := query["limit"]; ok {
		session = session.Limit(xtype.ToInt(v))
		delete(query, "limit")
	}
	if v, ok := query["offset"]; ok {
		session = session.Offset(xtype.ToInt(v))
		delete(query, "offset")
	}
	if v, ok := query["order"]; ok {
		session = session.Order(xtype.ToString(v))
		delete(query, "order")
	} else {
		session = session.Order("id DESC")
	}
	if len(query) > 0 {
		session = session.Where(query)
	}
	return session.Find(list).Error
}
