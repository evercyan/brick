package xstock

import (
	"context"
	"testing"
	"time"

	"github.com/evercyan/brick/xjson"
)

func TestFetchNewsList(t *testing.T) {
	var (
		ctx   = context.Background()
		code  = "600619"
		begin = time.Now().AddDate(0, 0, -3)
		end   = time.Now()
	)
	list, err := FetchNewsList(ctx, code, begin, end)
	if err != nil {
		t.Fatal(err)
	}
	xjson.Pretty(list, true)
}
