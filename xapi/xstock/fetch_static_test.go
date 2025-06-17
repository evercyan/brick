package xstock

import (
	"context"
	"testing"

	"github.com/evercyan/brick/xjson"
)

func TestFetchStatic(t *testing.T) {
	xjson.Pretty(FetchStatic(context.Background()), true)
}
