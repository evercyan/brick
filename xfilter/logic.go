package xfilter

import (
	"strings"
)

// Logic 条件组逻辑关系
type Logic string

// ...
const (
	LogicAnd Logic = "and"
	LogicOr  Logic = "or"
)

// String ...
func (t Logic) String() string {
	return string(t)
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
