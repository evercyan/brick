package xleet

// Min ...
func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// Max ...
func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// Abs ...
func Abs(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}
