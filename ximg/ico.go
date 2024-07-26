package ximg

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"os"

	"github.com/evercyan/brick/xlodash"
)

// icondir ...
type icondir struct {
	reserved  uint16
	imageType uint16
	numImages uint16
}

// newIcondir ...
func newIcondir() icondir {
	return icondir{
		imageType: 1,
		numImages: 1,
	}
}

// icondirEntry ...
type icondirEntry struct {
	imageWidth   uint8
	imageHeight  uint8
	numColors    uint8
	reserved     uint8
	colorPlanes  uint16
	bitsPerPixel uint16
	sizeInBytes  uint32
	offset       uint32
}

// newIcondirEntry ...
func newIcondirEntry() icondirEntry {
	return icondirEntry{
		colorPlanes:  1,
		bitsPerPixel: 32,
		offset:       22,
	}
}

// ToIco 图片转 ico, sizes 默认 256*256
func ToIco(srcPath, dstPath string, sizes ...int) error {
	src, _, err := Parse(srcPath)
	if err != nil {
		return fmt.Errorf("read file error: %v", err)
	}
	if src.Bounds().Dx() != src.Bounds().Dy() {
		return fmt.Errorf("src width must equal height")
	}
	dst, err := os.Create(dstPath)
	if err != nil {
		return fmt.Errorf("create file error: %v", err)
	}
	defer dst.Close()
	// 默认 256*256
	size := xlodash.First(sizes, 256)
	if src.Bounds().Dx() != size {
		src = Resize(src, size, size)
	}
	id := newIcondir()
	ide := newIcondirEntry()
	b := src.Bounds()
	m := image.NewRGBA(b)
	draw.Draw(m, b, src, b.Min, draw.Src)
	pngbb := new(bytes.Buffer)
	pngwriter := bufio.NewWriter(pngbb)
	png.Encode(pngwriter, m)
	pngwriter.Flush()
	ide.sizeInBytes = uint32(len(pngbb.Bytes()))
	bounds := m.Bounds()
	ide.imageWidth = uint8(bounds.Dx())
	ide.imageHeight = uint8(bounds.Dy())
	bb := new(bytes.Buffer)
	binary.Write(bb, binary.LittleEndian, id)
	binary.Write(bb, binary.LittleEndian, ide)
	dst.Write(bb.Bytes())
	dst.Write(pngbb.Bytes())
	return nil
}
