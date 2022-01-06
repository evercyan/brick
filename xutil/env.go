package xutil

import (
	"os"
	"runtime"
	"strings"
)

// GetEnv ...
func GetEnv(name string, defaultValues ...string) string {
	value := os.Getenv(name)
	if value == "" && len(defaultValues) > 0 {
		value = defaultValues[0]
	}
	return value
}

// GetEnvMap ...
func GetEnvMap() map[string]string {
	envList := os.Environ()
	envMap := make(map[string]string, len(envList))
	for _, str := range envList {
		nodes := strings.SplitN(str, "=", 2)
		if len(nodes) >= 2 {
			envMap[nodes[0]] = nodes[1]
		}
	}
	return envMap
}

// IsWin ...
func IsWin() bool {
	return runtime.GOOS == "windows"
}

// IsMac ...
func IsMac() bool {
	return runtime.GOOS == "darwin"
}

// IsLinux ...
func IsLinux() bool {
	return runtime.GOOS == "linux"
}
