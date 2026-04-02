package xfile

import (
	"context"
	"testing"
	"time"

	"github.com/evercyan/brick/xjson"
)

// Excel ...
type Excel struct {
	Id     int64             `json:"id" excel:"序号"`
	Name   string            `json:"name" excel:"名称"`
	Price  float64           `json:"price" excel:"价格"`
	Time   time.Time         `json:"time"`
	test   string            `json:"test"`
	Colors map[string]string `json:"-" excel:"-"`
}

var (
	ctx   = context.Background()
	epath = "./output.xlsx"
)

func TestReadExcel(t *testing.T) {
	list := make([]*Excel, 0)
	if err := ReadExcel(ctx, epath, &list); err != nil {
		t.Fatal(err)
	}
	xjson.Pretty(list, true)
}

func TestWriteExcel1(t *testing.T) {
	list := []*Excel{
		{
			Id:    1,
			Name:  "aaaa",
			Price: 100.9,
			Time:  time.Now(),
			test:  "test1",
			Colors: map[string]string{
				"Name": "#FF0000",
				"Time": "#00FF00",
			},
		},
		{
			Id:    2,
			Name:  "bbbb",
			Price: 10.9,
			Time:  time.Now(),
			test:  "test2",
			Colors: map[string]string{
				"Price": "#00FF00",
			},
		},
		{
			Id:     3,
			Name:   "cccc",
			Price:  20.9,
			Time:   time.Now(),
			test:   "test3",
			Colors: map[string]string{},
		},
		{
			Id:     4,
			Name:   "dddd",
			Price:  30.9,
			Time:   time.Now(),
			test:   "test4",
			Colors: nil,
		},
	}
	if err := WriteExcel(ctx, epath, list); err != nil {
		t.Fatal(err)
	}
}

func TestWriteExcel2(t *testing.T) {
	list := [][]interface{}{
		{
			"序号", "名称", "价格",
		},
		{
			"1", "aaa", "33.0",
		},
		{
			"2", "bbb", "34.0",
		},
		{
			"1", "aaa", "33.0",
		},
		{
			"2", "bbb", "34.0",
		},
		{
			"1", "aaa", "33.0",
		},
		{
			"2", "bbb", "34.0",
		},
		{
			"1", "aaa", "33.0",
		},
		{
			"2", "bbb", "34.0",
		},
		{
			"1", "aaa", "33.0",
		},
		{
			"2", "bbb", "34.0",
		},
		{
			"1", "aaa", "33.0",
		},
		{
			"2", "bbb", "34.0",
		},
		{
			"1", "aaa", "33.0",
		},
		{
			"2", "bbb", "34.0",
		},
		{
			"1", "aaa", "33.0",
		},
		{
			"2", "bbb", "34.0",
		},
		{
			"1", "aaa", "33.0",
		},
		{
			"2", "bbb", "34.0",
		},
		{
			"1", "aaa", "33.0",
		},
		{
			"2", "bbb", "34.0",
		},
		{
			"1", "aaa", "33.0",
		},
		{
			"2", "bbb", "34.0",
		},
		{
			"1", "aaa", "33.0",
		},
		{
			"2", "bbb", "34.0",
		},
		{
			"1", "aaa", "33.0",
		},
		{
			"2", "bbb", "34.0",
		},
		{
			"1", "aaa", "33.0",
		},
		{
			"2", "bbb", "34.0",
		},
		{
			"1", "aaa", "33.0",
		},
		{
			"2", "bbb", "34.0",
		},
	}
	colors := []map[int]string{
		{
			1: "#FF0000", // aaa 会标红
			2: "#00FF00", // 33.0 会标绿
		},
		{
			2: "#0000FF", // 34.0 会标蓝
		},
	}
	if err := WriteExcel(ctx, epath, list, colors...); err != nil {
		t.Fatal(err)
	}
}
