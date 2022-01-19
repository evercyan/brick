package xversion

// CompareResult ...
type CompareResult int

const (
	Equal CompareResult = iota
	Less
	Greater
)
