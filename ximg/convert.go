package ximg

import (
	"context"
	"os"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/evercyan/brick/xfile"
	"github.com/evercyan/brick/xlodash"
)

// HTML2PNGArg ...
type HTML2PNGArg struct {
	URL       string  `json:"url"`
	ImagePath string  `json:"image_path"`
	Width     int64   `json:"width"`
	Scale     float64 `json:"scale"`
}

// HTML2PNG ...
func HTML2PNG(ctx context.Context, arg *HTML2PNGArg) error {
	ctx, cancel := chromedp.NewContext(ctx)
	defer cancel()
	if xfile.IsFile(arg.URL) {
		arg.URL = "file://" + arg.URL
	}
	if arg.Scale == 0 {
		arg.Scale = 1
	}
	var buf []byte
	var width, height int64
	err := chromedp.Run(
		ctx,
		chromedp.Navigate(arg.URL),
		chromedp.WaitVisible("body", chromedp.ByQuery),
		chromedp.EvaluateAsDevTools(`Math.max(
			document.body.scrollWidth,
			document.documentElement.scrollWidth
		)`, &width),
		chromedp.EvaluateAsDevTools(`Math.max(
			document.body.scrollHeight,
			document.documentElement.scrollHeight
		)`, &height),
		chromedp.EmulateViewport(
			xlodash.Max(width, arg.Width), height, chromedp.EmulateScale(arg.Scale),
		),
		chromedp.Sleep(3*time.Second),
		chromedp.FullScreenshot(&buf, 100),
	)
	if err != nil {
		return err
	}
	return os.WriteFile(arg.ImagePath, buf, 0644)
}

// ConvertHTML2PNG ...
func ConvertHTML2PNG(ctx context.Context, url, imgPath string) error {
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
