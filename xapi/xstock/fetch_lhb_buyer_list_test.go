package xstock

import (
	"context"
	"testing"

	"github.com/evercyan/brick/xjson"
)

func TestFetchLHBBuyerList(t *testing.T) {
	list, err := FetchLHBBuyerList(context.Background(), "001331", "2025-08-06")
	if err != nil {
		t.Fatal(err)
	}
	xjson.Pretty(list, true)
}
