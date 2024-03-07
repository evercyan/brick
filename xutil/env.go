package xutil

import (
	"os"
	"runtime"
	"strings"
)

// Setenv ...
func Setenv(key, value string) error {
	return os.Setenv(key, value)
}

// Getenv ...
func Getenv(key string, defaults ...string) string {
	value := os.Getenv(key)
	if value == "" && len(defaults) > 0 {
		value = defaults[0]
	}
	return value
}

// GetenvMap ...
func GetenvMap() map[string]string {
	res := make(map[string]string)
	for _, str := range os.Environ() {
		nodes := strings.SplitN(str, "=", 2)
		if len(nodes) >= 2 {
			res[nodes[0]] = nodes[1]
		}
	}
	return res
}

// GetOs ...
func GetOs() string {
	return runtime.GOOS
}

// IsWin ...
func IsWin() bool {
	return GetOs() == "windows"
}

// IsMac ...
func IsMac() bool {
	return GetOs() == "darwin"
}

// IsLinux ...
func IsLinux() bool {
	return GetOs() == "linux"
}
