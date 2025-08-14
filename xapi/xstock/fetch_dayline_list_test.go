package xstock

import (
	"context"
	"fmt"
	"testing"

	"github.com/evercyan/brick/xjson"
)

func TestFetchDaylineList(t *testing.T) {
	code := "001331"
	begin, end := "20250801", "20250805"
	list, err := FetchDaylineList(context.Background(), code, begin, end)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(xjson.Pretty(list))
}
