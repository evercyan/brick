package xfile

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"regexp"
	"strings"
)

// GetCurrentDir ...
func GetCurrentDir() string {
	if dir, err := os.Getwd(); err == nil {
		return dir
	}
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return ""
	}
	return strings.Replace(dir, "\\", "/", -1)
}

// GetHomeDir ...
func GetHomeDir() string {
	dir, err := user.Current()
	if err != nil {
		return ""
	}
	return dir.HomeDir
}

// GetConfigDir ...
func GetConfigDir(paths ...string) (string, error) {
	userPath, err := user.Current()
	if err != nil {
		return "", err
	}
	filePath := fmt.Sprintf(
		"%s/.config",
		userPath.HomeDir,
	)
	if len(paths) > 0 {
		for _, path := range paths {
			filePath += "/" + path
		}
	}
	if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
		return "", err
	}
	return filePath, nil
}

// ListDirs ...
func ListDirs(dir string, isRecursive ...bool) []string {
	fs, err := os.ReadDir(dir)
	if err != nil {
		return nil
	}
	l := make([]string, 0)
	for _, f := range fs {
		fp := fmt.Sprintf("%s/%s", strings.TrimRight(dir, "/"), f.Name())
		if !f.IsDir() {
			continue
		}
		l = append(l, fp)
		if len(isRecursive) > 0 {
			l = append(l, ListDirs(fp, isRecursive...)...)
		}
	}
	return l
}

// ListFiles ...
func ListFiles(dir string, match string, isRecursive ...bool) []string {
	fs, err := os.ReadDir(dir)
	if err != nil {
		return nil
	}
	l := make([]string, 0)
	re := regexp.MustCompile(match)
	for _, f := range fs {
		fp := fmt.Sprintf("%s/%s", strings.TrimRight(dir, "/"), f.Name())
		if f.IsDir() {
			if len(isRecursive) > 0 {
				l = append(l, ListFiles(fp, match, isRecursive...)...)
			}
			continue
		}
		if match != "" && !re.MatchString(f.Name()) {
			continue
		}
		l = append(l, fp)
	}
	return l
}
