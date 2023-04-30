package xversion

import (
	"strings"

	"github.com/evercyan/brick/xtype"
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
func Compare(src string, dst string) int {
	src, dst = Format(src), Format(dst)
	if src == dst {
		return 0
	}
	if src == "" {
		return -1
	}
	if dst == "" {
		return 1
	}
	srcs, dsts := strings.Split(src, "."), strings.Split(dst, ".")
	srcLen, dstLen := len(srcs), len(dsts)
	minLen := srcLen
	if dstLen < minLen {
		minLen = dstLen
	}
	for i := 0; i < minLen; i++ {
		// 无法转换的, 默认为 0, 即 v1.a.3 == v1.0.3
		srcNum, dstNum := xtype.ToInt(srcs[i]), xtype.ToInt(dsts[i])
		if srcNum < dstNum {
			return -1
		} else if srcNum > dstNum {
			return 1
		}
	}

	if srcLen < dstLen {
		return -1
	} else if srcLen == dstLen {
		return 0
	} else {
		return 1
	}
}
