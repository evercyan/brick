package ximg

import (
	"context"
	"fmt"
	"testing"

	"github.com/evercyan/brick/xfile"
)

func TestHTML2PNG(t *testing.T) {
	ctx := context.Background()
	{
		url := "https://www.baidu.com/"
		imgPath := xfile.Temp("html.png")
		if err := HTML2PNG(ctx, url, imgPath); err != nil {
			t.Fatal(err)
		}
		fmt.Println(imgPath)
	}
}
