package xfile

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/evercyan/brick/xgen"
	"github.com/evercyan/brick/xtype"
)

// Temp ...
func Temp(filenames ...string) string {
	var filename string
	if len(filename) > 0 {
		filename = filenames[0]
	} else {
		filename = xgen.Nanoid()
	}
	return path.Join(os.TempDir(), filename)
}

// Read ...
func Read(filepath string) string {
	b, err := os.ReadFile(filepath)
	if err != nil {
		return ""
	}
	return string(b)
}

// Write ...
func Write(filepath, str string) error {
	return os.WriteFile(filepath, []byte(str), 0755)
}

// Copy ...
func Copy(src, dst string) error {
	return Write(dst, Read(src))
}

// Size ...
func Size(filepath string) int64 {
	f, err := os.Stat(filepath)
	if err != nil {
		return 0
	}
	return f.Size()
}

// SizeText ...
func SizeText(size int64) string {
	s := float64(size)
	units := []string{"B", "KB", "MB", "GB", "TB"}
	index := 0
	for s >= 1024 {
		s /= 1024
		index++
		if index == len(units)-1 {
			break
		}
	}
	return fmt.Sprintf("%.2f%s", s, units[index])
}

// LineCount ...
func LineCount(filepath string) int {
	f, err := os.Open(filepath)
	if err != nil {
		return 0
	}
	defer f.Close()
	count := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		count++
	}
	return count
}

// LineContent ...
func LineContent(filepath string, numbers ...int) map[int]string {
	f, err := os.Open(filepath)
	if err != nil {
		return nil
	}
	defer f.Close()
	res := make(map[int]string)
	count := len(numbers)
	scanner := bufio.NewScanner(f)
	for number := 1; scanner.Scan(); number++ {
		if count == 0 || xtype.IsContains(numbers, number) {
			res[number] = scanner.Text()
		}
	}
	return res
}

// MineType ...
func MineType(filepath string) string {
	f, err := os.Open(filepath)
	if err != nil {
		return ""
	}
	return ReaderMineType(f)
}

// ReaderMineType ...
func ReaderMineType(r io.Reader) string {
	// 512 http/sniff.go sniffLen
	var buf [512]byte
	n, _ := io.ReadFull(r, buf[:])
	if n == 0 {
		return ""
	}
	return http.DetectContentType(buf[:n])
}

// WriteJSON write data to JSON file
func WriteJSON(filepath string, data interface{}, pretty ...bool) error {
	var (
		b   []byte
		err error
	)
	if len(pretty) > 0 {
		b, err = json.MarshalIndent(data, "", "    ")
	} else {
		b, err = json.Marshal(data)
	}
	if err != nil {
		return err
	}
	return os.WriteFile(filepath, b, 0664)
}

// ReadJSON read JSON file data
func ReadJSON(filepath string, data interface{}) error {
	b, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, data)
}

// Md5 file md5
func Md5(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	fi, err := f.Stat()
	if err != nil {
		return "", err
	}
	var size int64 = 1024 * 1024
	hash := md5.New()
	if fi.Size() < size {
		data, err := os.ReadFile(path)
		if err != nil {
			return "", err
		}
		hash.Write(data)
	} else {
		b := make([]byte, size)
		for {
			n, err := f.Read(b)
			if err != nil {
				break
			}
			hash.Write(b[:n])
		}
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}

// Shadow 根据当前文件名称取其影分身
// /tmp/abc.txt => /tmp/abc_1.txt
func Shadow(filepath string) string {
	ext := path.Ext(filepath)
	prefix := strings.TrimSuffix(filepath, ext)
	res := ""
	for i := 1; ; i++ {
		f := fmt.Sprintf("%s_%d%s", prefix, i, ext)
		if !IsExist(f) {
			res = f
			break
		}
	}
	return res
}
