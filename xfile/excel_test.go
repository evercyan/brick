package xfile

import (
	"context"
	"testing"
	"time"

	"github.com/evercyan/brick/xjson"
)

// Excel ...
type Excel struct {
	Id    int64     `json:"id" excel:"序号"`
	Name  string    `json:"name" excel:"名称"`
	Price float64   `json:"price" excel:"价格"`
	Time  time.Time `json:"time"`
	test  string    `json:"test"`
}

var (
	ctx   = context.Background()
	epath = "../output.xlsx"
)

func TestReadExcel(t *testing.T) {
	list := make([]*Excel, 0)
	if err := ReadExcel(ctx, epath, &list); err != nil {
		t.Fatal(err)
	}
	xjson.Pretty(list, true)
}

func TestWriteExcel(t *testing.T) {
	list := []*Excel{
		{
			Id:    1,
			Name:  "aaaa",
			Price: 100.9,
			Time:  time.Now(),
			test:  "test1",
		},
		{
			Id:    2,
			Name:  "bbbb",
			Price: 10.9,
			Time:  time.Now(),
			test:  "test2",
		},
	}
	if err := WriteExcel(ctx, epath, list, true); err != nil {
		t.Fatal(err)
	}
}
