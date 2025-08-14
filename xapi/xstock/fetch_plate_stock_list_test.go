package xstock

import (
	"context"
	"testing"

	"github.com/evercyan/brick/xjson"
)

func TestFetchPlateStockList(t *testing.T) {
	plateCode := "BK0424"
	list, err := FetchPlateStockList(context.Background(), plateCode)
	if err != nil {
		t.Fatal(err)
	}
	xjson.Pretty(list, true)
}
