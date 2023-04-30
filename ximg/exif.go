package ximg

import (
	"github.com/barasher/go-exiftool"
)

// ATTENTION
// go-exiftool 需要安装 exiftool(https://exiftool.org/)

// ReadExif ...
func ReadExif(filepath string) map[string]interface{} {
	exif, err := exiftool.NewExiftool()
	if err != nil {
		return nil
	}
	defer exif.Close()

	metadatas := exif.ExtractMetadata(filepath)
	fields := make(map[string]interface{})
	if metadatas[0].Err == nil {
		fields = metadatas[0].Fields
	}
	return fields
}

// WriteExif ...
// key 须是有效的 exif 字段
func WriteExif(filepath string, kv map[string]string) error {
	exif, err := exiftool.NewExiftool()
	if err != nil {
		return err
	}
	defer exif.Close()

	metadatas := exif.ExtractMetadata(filepath)
	if metadatas[0].Err != nil {
		return metadatas[0].Err
	}

	for k, v := range kv {
		metadatas[0].SetString(k, v)
	}
	exif.WriteMetadata(metadatas)
	return nil
}
