package xstock

import (
	"context"
	"testing"

	"github.com/evercyan/brick/xjson"
)

func TestFetchAllStockList(t *testing.T) {
	list, err := FetchAllStockList(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	xjson.Pretty(list, true)
}

func TestFetchStockList(t *testing.T) {
	codes := []string{"920005"}
	list, err := FetchStockList(context.Background(), codes)
	if err != nil {
		t.Fatal(err)
	}
	xjson.Pretty(list, true)
}
