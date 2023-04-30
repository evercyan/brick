package ximg

import (
	"os"
	"testing"

	"github.com/evercyan/brick/xfile"
	"github.com/stretchr/testify/assert"
)

func TestType(t *testing.T) {
	assert.Equal(t, "png", Type("../logo.png"))
	assert.Empty(t, Type("../xfile"))
	assert.Empty(t, Type("../README.md"))
}

func TestSize(t *testing.T) {
	w, h := Size("../logo.png")
	assert.Equal(t, 512, w)
	assert.Equal(t, 512, h)

	w1, h1 := Size("../README.md")
	assert.Empty(t, w1)
	assert.Empty(t, h1)
}

func TestBase64(t *testing.T) {
	assert.NotEmpty(t, Base64("../logo.png"))
}

func TestReadWrite(t *testing.T) {
	img := Read("../logo.png")
	assert.NotNil(t, img)

	dst1 := xfile.Temp("write.png")
	defer os.Remove(dst1)
	assert.NotNil(t, Write(dst1, img))

	dst2 := xfile.Temp("write.jpg")
	defer os.Remove(dst2)
	assert.NotNil(t, Write(dst2, img))

}
