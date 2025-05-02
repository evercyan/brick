//go:build windows

package xfile

import (
	"strings"
	"syscall"
	"unsafe"
)

var (
	GetHomdDir = getRoamingAppDataDir
	GetDataDir = getRoamingAppDataDir
	GetConfDir = getRoamingAppDataDir
)

var (
	modshell32               = syscall.NewLazyDLL("shell32.dll")
	modole32                 = syscall.NewLazyDLL("ole32.dll")
	procSHGetKnownFolderPath = modshell32.NewProc("SHGetKnownFolderPath")
	procCoTaskMemFree        = modole32.NewProc("CoTaskMemFree")
	roamingAppData           = syscall.GUID{
		0x3EB685DB,
		0x65F9,
		0x4CF6,
		[8]byte{0xA0, 0x3A, 0xE3, 0xEF, 0x65, 0x72, 0x9F, 0x3D},
	}
)

func coTaskMemFree(ptr uintptr) {
	procCoTaskMemFree.Call(ptr)
}

func getRoamingAppDataDir() string {
	var pwstr uintptr
	_, _, _ = procSHGetKnownFolderPath.Call(
		uintptr(unsafe.Pointer(&roamingAppData)),
		uintptr(uint32(0)),
		uintptr(unsafe.Pointer(nil)),
		uintptr(unsafe.Pointer(&pwstr)),
	)
	defer coTaskMemFree(pwstr)
	return strings.Replace(utf16PtrToString(pwstr), "\\", "/", -1)
}

func utf16PtrToString(str uintptr) string {
	ptr := unsafe.Pointer(str)
	return syscall.UTF16ToString((*[1 << 16]uint16)(ptr)[:])
}
