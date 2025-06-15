package ximg

import (
	"context"
	"os"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/evercyan/brick/xfile"
)

// HTML2PNG ...
func HTML2PNG(ctx context.Context, url, imgPath string) error {
	ctx, cancel := chromedp.NewContext(ctx)
	defer cancel()
	if xfile.IsFile(url) {
		url = "file://" + url
	}
	var buf []byte
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.WaitVisible(`body`, chromedp.ByQuery), // 等待 body 元素可见
		chromedp.Sleep(1*time.Second),                  // 额外延迟（根据需求调整）
		chromedp.FullScreenshot(&buf, 100),
	)
	if err != nil {
		return err
	}
	return os.WriteFile(imgPath, buf, 0644)
}
