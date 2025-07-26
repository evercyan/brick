package xfile

import (
	"math"
	"os"

	"github.com/alfg/mp4"
)

// MP4Info ...
type MP4Info struct {
	Size     int64 `json:"size"`     // 文件大小
	Width    int64 `json:"width"`    // 宽
	Height   int64 `json:"height"`   // 高
	Duration int64 `json:"duration"` // 时长(s)
}

// GetMP4Info ...
func GetMP4Info(vpath string) (*MP4Info, error) {
	file, err := os.Open(vpath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	fstat, err := file.Stat()
	if err != nil {
		return nil, err
	}
	info, err := mp4.OpenFromReader(file, fstat.Size())
	if err != nil {
		return nil, err
	}
	mp4Info := &MP4Info{
		Size:     fstat.Size(),
		Width:    int64(info.Moov.Traks[0].Tkhd.Width >> 16),
		Height:   int64(info.Moov.Traks[0].Tkhd.Height >> 16),
		Duration: int64(math.Ceil(float64(info.Moov.Traks[0].Tkhd.Duration) / 1000)),
	}
	return mp4Info, nil
}
