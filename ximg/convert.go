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
	var width, height int64
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.WaitVisible("body", chromedp.ByQuery),
		chromedp.EvaluateAsDevTools(`Math.max(
			document.body.scrollWidth, 
			document.documentElement.scrollWidth
		)`, &width),
		chromedp.EvaluateAsDevTools(`Math.max(
			document.body.scrollHeight, 
			document.documentElement.scrollHeight
		)`, &height),
		chromedp.EmulateViewport(width, height),
		chromedp.Sleep(3*time.Second),
		chromedp.FullScreenshot(&buf, 100),
	)
	if err != nil {
		return err
	}
	return os.WriteFile(imgPath, buf, 0644)
}
