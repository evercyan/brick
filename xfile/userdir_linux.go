//go:build linux

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
	xdgDataHome := os.Getenv("XDG_DATA_HOME")
	dirs := make([]string, 0)
	if xdgDataHome != "" {
		dirs = append(dirs, xdgDataHome)
	} else {
		dirs = append(dirs, GetHomeDir(), ".local", "share")
	}
	dirs = append(dirs, args...)
	return filepath.Join(dirs...)
}

// GetConfDir ...
func GetConfDir(args ...string) string {
	xdgConfigHome := os.Getenv("XDG_CONFIG_HOME")
	dirs := make([]string, 0)
	if xdgConfigHome != "" {
		dirs = append(dirs, xdgConfigHome)
	} else {
		dirs = append(dirs, GetHomeDir(), ".config")
	}
	dirs = append(dirs, args...)
	return filepath.Join(dirs...)
}
