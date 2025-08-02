package xstock

import (
	"context"
	"fmt"
	"testing"

	"github.com/evercyan/brick/xjson"
)

func TestFetchMinlineList(t *testing.T) {
	code := "002392"
	list, err := FetchMinlineList(context.Background(), code)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(xjson.Pretty(list))
}
