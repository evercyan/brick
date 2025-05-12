package xeastmoney

import (
	"context"
	"fmt"
	"testing"

	"github.com/evercyan/brick/xjson"
)

func TestFetchStockList(t *testing.T) {
	codes := []string{"002392", "600340"}
	list, err := FetchStockList(context.Background(), codes)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(xjson.Pretty(list))
}
