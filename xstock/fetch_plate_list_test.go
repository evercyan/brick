package xstock

import (
	"context"
	"testing"

	"github.com/evercyan/brick/xjson"
)

func TestFetchPlateList(t *testing.T) {
	xjson.Pretty(FetchPlateList(context.Background()), true)
}
