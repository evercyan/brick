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
		url = "/Users/Cyan/Y1ker/AI/12-股票分析/平潭发展/20251202/结果.html"
		imgPath = "/Users/Cyan/Y1ker/AI/12-股票分析/平潭发展/20251202/结果.png"
		arg := &HTML2PNGArg{
			URL:       url,
			ImagePath: imgPath,
		}
		if err := HTML2PNG(ctx, arg); err != nil {
			t.Fatal(err)
		}
		fmt.Println(imgPath)
	}
}
