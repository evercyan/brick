package ximg

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDownload(t *testing.T) {
	imgURL := "https://i4.hoopchina.com.cn/hupuapp/bbs/873/27181873/thread_27181873_20210730165538_s_152769_w_1080_h_1079_72771.jpg?x-oss-process=image/resize,w_800/format,webp"
	imgPath := "./demo.jpg"
	err := Download(context.Background(), imgURL, imgPath)
	assert.Nil(t, err)
	os.Remove(imgPath)
}
