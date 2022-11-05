package xutil

import (
	"strings"

	"github.com/evercyan/brick/xtype"
)

// Number ...
func CompareNumber(a, b interface{}) int {
	fa := xtype.ToFloat64(a)
	fb := xtype.ToFloat64(b)
	if fa > fb {
		return 1
	} else if fa < fb {
		return -1
	}
	return 0
}

// String ...
func CompareString(a, b interface{}) int {
	sa, sb := "", ""
	if xtype.IsString(a) {
		sa = a.(string)
	}
	if xtype.IsString(b) {
		sb = b.(string)
	}
	return strings.Compare(sa, sb)
}

// Compare ...
func Compare(a, b interface{}) int {
	if a == nil && b == nil {
		return 0
	}
	if xtype.IsNumeric(a) || xtype.IsBool(a) || xtype.IsNumeric(b) || xtype.IsBool(b) {
		return CompareNumber(a, b)
	}
	if xtype.IsString(a) || xtype.IsString(b) {
		return CompareString(a, b)
	}
	return 0
}