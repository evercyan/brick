package xchatlog

import (
	"context"
	"testing"

	"github.com/evercyan/brick/xjson"
)

func TestFetchMessageList(t *testing.T) {
	ctx := context.Background()
	req := &FetchMessageListReq{
		Time:   "2025-06-15",
		Talker: "文件传输助手",
	}
	list, err := FetchMessageList(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	xjson.Pretty(list, true)
}
