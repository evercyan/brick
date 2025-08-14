package xstock

import (
	"context"
	"testing"

	"github.com/evercyan/brick/xjson"
)

func TestFetchLHBList(t *testing.T) {
	list, err := FetchLHBList(context.Background(), "2025-08-06")
	if err != nil {
		t.Fatal(err)
	}
	xjson.Pretty(list, true)
}
