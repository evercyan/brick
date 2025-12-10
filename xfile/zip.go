package xfile

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/alexmullins/zip"
)

// ZipOption
type ZipOption struct {
	Password  string // 加密压缩
	KeepLevel bool   // 保持层级
}

// defaultZipOption 默认配置
var defaultZipOption = &ZipOption{
	Password:  "",
	KeepLevel: false,
}

// ----------------------------------------------------------------

// ZipOptionFn ...
type ZipOptionFn func(*ZipOption)

// WithZipPassword ...
func WithZipPassword(value string) ZipOptionFn {
	return func(o *ZipOption) {
		o.Password = value
	}
}

// WithZipKeepLevel ...
func WithZipKeepLevel(value bool) ZipOptionFn {
	return func(o *ZipOption) {
		o.KeepLevel = value
	}
}

// ----------------------------------------------------------------

// WriteZip 写入 zip 文件
func WriteZip(zipPath string, baseDir string, args []string, options ...ZipOptionFn) error {
	fpaths := make([]string, 0)
	fnames := make(map[string]string)
	baseDir = strings.TrimRight(baseDir, "/") + "/"
	for _, arg := range args {
		if IsDir(arg) {
			tpaths := ListFiles(arg, "", true)
			for _, tpath := range tpaths {
				fpaths = append(fpaths, tpath)
				fnames[tpath] = strings.TrimLeft(tpath, baseDir)
			}
		} else {
			fpaths = append(fpaths, arg)
			fnames[arg] = strings.TrimLeft(arg, baseDir)
		}
	}
	if len(fpaths) == 0 {
		return fmt.Errorf("empty fpaths")
	}
	option := defaultZipOption
	for _, fn := range options {
		fn(option)
	}
	zipfile, err := os.Create(zipPath)
	if err != nil {
		return err
	}
	defer zipfile.Close()
	zw := zip.NewWriter(zipfile)
	defer zw.Close()
	for _, fpath := range fpaths {
		if err := WriteZipWriter(zw, option, fnames, fpath); err != nil {
			return err
		}
	}
	return nil
}

// WriteZipWriter ...
func WriteZipWriter(zw *zip.Writer, option *ZipOption, fnames map[string]string, fpath string) error {
	file, err := os.Open(fpath)
	if err != nil {
		return err
	}
	defer file.Close()
	info, err := file.Stat()
	if err != nil {
		return err
	}
	if info.IsDir() {
		return fmt.Errorf("directory not supported: %s", fpath)
	}
	fname := filepath.Base(fpath)
	if option.KeepLevel {
		if v, ok := fnames[fpath]; ok {
			fname = v
		}
	}
	fh, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}
	fh.Name = fname
	fh.Method = zip.Deflate
	fh.SetModTime(info.ModTime().Add(8 * time.Hour))
	if option.Password != "" {
		fh.SetPassword(option.Password)
	}
	writer, err := zw.CreateHeader(fh)
	if err != nil {
		return err
	}
	buf := make([]byte, 32*1024)
	_, err = io.CopyBuffer(writer, file, buf)
	return err
}
