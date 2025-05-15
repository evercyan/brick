package xstock

import (
	"context"
	"fmt"
	"testing"

	"github.com/evercyan/brick/xjson"
)

func TestFetchTradeList(t *testing.T) {
	code := "002392"
	begin, end := "20250501", "20250508"
	list, err := FetchKlineList(context.Background(), code, begin, end)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(xjson.Pretty(list))
}
