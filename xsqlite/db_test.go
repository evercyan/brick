package xsqlite

import (
	"fmt"
	"github.com/evercyan/brick/xjson"
	"github.com/stretchr/testify/assert"
	"testing"
)

// ----------------------------------------------------------------

// Record ...
type Record struct {
	Model
	Name string `json:"name" gorm:"column:name;not null"`
}

func (t *Record) TableName() string {
	return "record"
}

// ----------------------------------------------------------------

func TestNew(t *testing.T) {
	options := []Option{
		WithModels(&Record{}),
		WithDebug(),
		//WithPassword("123456"),
	}
	db, err := New("test1.db", options...)
	assert.Nil(t, err)

	{
		err := db.Create(&Record{
			Name: "hello",
		}).Error
		assert.Nil(t, err)
	}
	{
		list := make([]*Record, 0)
		err := db.Find(&list).Error
		assert.Nil(t, err)
		fmt.Println("======== list", xjson.Pretty(list))
	}
}
