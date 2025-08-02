package xstock

import (
	"context"
	"fmt"
	"testing"

	"github.com/evercyan/brick/xjson"
)

func TestFetchMultiDaylineList(t *testing.T) {
	codes := []string{"002900", "002365"}
	list, err := FetchMultiDaylineList(context.Background(), codes)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(xjson.Pretty(list))
}
