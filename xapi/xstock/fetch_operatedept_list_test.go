package xstock

import (
	"context"
	"testing"

	"github.com/evercyan/brick/xjson"
)

func TestFetchOperatedeptList(t *testing.T) {
	list, err := FetchOperatedeptList(context.Background(), 1, 10)
	if err != nil {
		t.Fatal(err)
	}
	xjson.Pretty(list, true)
}
