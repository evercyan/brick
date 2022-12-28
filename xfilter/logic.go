package xfilter

import (
	"strings"
)

// Logic 条件组逻辑关系
type Logic int

// ...
const (
	LogicAnd Logic = iota
	LogicOr
)

// String ...
func (t Logic) String() string {
	switch t {
	case LogicAnd:
		return "and"
	case LogicOr:
		return "or"
	default:
		return ""
	}
}

// ToLogic ...
func ToLogic(s string) Logic {
	switch strings.ToLower(s) {
	case LogicOr.String():
		return LogicOr
	default:
		return LogicAnd
	}
}
