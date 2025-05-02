//go:build darwin

package xfile

import (
	"os"
	"path/filepath"
)

// GetHomeDir ...
func GetHomeDir() string {
	return os.Getenv("HOME")
}

// GetDataDir ...
func GetDataDir(args ...string) string {
	dirs := []string{GetHomeDir(), "Library", "Application Support"}
	dirs = append(dirs, args...)
	return filepath.Join(dirs...)
}

// GetConfDir ...
func GetConfDir(args ...string) string {
	dirs := []string{GetHomeDir(), ".config"}
	dirs = append(dirs, args...)
	return filepath.Join(dirs...)
}
