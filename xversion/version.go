package xversion

import (
	"strings"

	"github.com/evercyan/brick/xconvert"
	"github.com/evercyan/brick/xutil"
)

// Format 格式化版本
func Format(version string) string {
	version = strings.TrimLeft(strings.TrimSpace(strings.ToLower(version)), "v")
	return xutil.Replace(version, map[string]string{
		"-": ".",
	})
}

// Compare 版本比较
func Compare(src string, dst string) CompareResult {
	src, dst = Format(src), Format(dst)
	if src == dst {
		return Equal
	}
	if src == "" {
		return Less
	}
	if dst == "" {
		return Greater
	}
	srcs, dsts := strings.Split(src, "."), strings.Split(dst, ".")
	srcLen, dstLen := len(srcs), len(dsts)
	minLen := srcLen
	if dstLen < minLen {
		minLen = dstLen
	}
	for i := 0; i < minLen; i++ {
		// 无法转换的, 默认为 0, 即 v1.a.3 == v1.0.3
		srcNum, dstNum := xconvert.ToUint(srcs[i]), xconvert.ToUint(dsts[i])
		if srcNum < dstNum {
			return Less
		} else if srcNum > dstNum {
			return Greater
		}
	}
	if srcLen < dstLen {
		return Less
	} else if srcLen == dstLen {
		return Equal
	} else {
		return Greater
	}
}
