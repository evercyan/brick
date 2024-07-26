package xgen

import (
	_ "unsafe"

	"github.com/evercyan/brick/xlodash"
	"github.com/evercyan/brick/xtype"
)

//go:noescape
//go:linkname fastrand runtime.fastrand
func fastrand() uint32

// randUint32Max returns pseudorandom uint32 in the range [0-max]
func randUint32Max(max uint32) uint32 {
	return uint32((uint64(fastrand()) * uint64(max)) >> 32)
}

// RandInt ...
func RandInt(min, max int) int {
	if max < min {
		min, max = max, min
	}
	return min + int(randUint32Max(uint32(max+1-min)))
}

// RandStr ...
func RandStr(length int, tpl ...string) string {
	chars := xlodash.First(tpl, "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]byte, length)
	l := len(chars) - 1
	for i := length - 1; i >= 0; i-- {
		idx := RandInt(0, l)
		b[i] = chars[idx]
	}
	return xtype.Bytes2String(b)
}

// Range ...
func Range(min, max int) []int {
	res := make([]int, 0)
	if min > max {
		min, max = max, min
	}
	for i := min; i <= max; i++ {
		res = append(res, i)
	}
	return res
}
