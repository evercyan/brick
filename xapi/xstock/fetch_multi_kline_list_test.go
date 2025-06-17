package xstock

import (
	"context"
	"fmt"
	"testing"

	"github.com/evercyan/brick/xjson"
)

func TestFetchMultiKlineList(t *testing.T) {
	codes := []string{"002900", "002365"}
	list, err := FetchMultiKlineList(context.Background(), codes)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(xjson.Pretty(list))
}
