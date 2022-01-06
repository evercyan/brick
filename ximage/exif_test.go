package ximage

// import (
// 	"os"
// 	"testing"
//
// 	"github.com/evercyan/brick/xfile"
// 	"github.com/stretchr/testify/assert"
// )
//
// func TestExif(t *testing.T) {
// 	src := "../letitgo.png"
// 	dst := "/tmp/logo.png"
// 	xfile.Write(dst, xfile.Read(src))
//
// 	err := WriteExif(dst, map[string]string{
// 		"XResolution": "36",
// 	})
// 	assert.Nil(t, err)
//
// 	exifData := ReadExif(dst)
// 	assert.Less(t, 0, len(exifData))
// 	assert.Equal(t, "image/png", exifData["MIMEType"])
// 	assert.Equal(t, "260x109", exifData["ImageSize"])
// 	assert.Equal(t, float64(36), exifData["XResolution"])
//
// 	os.Remove(dst)
// }
