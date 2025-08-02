package xstock

import (
	"context"
	"testing"

	"github.com/evercyan/brick/xjson"
)

func TestFetchBuyerList(t *testing.T) {
	list, err := FetchBuyerList(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	xjson.Pretty(list, true)
}
