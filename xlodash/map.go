package xlodash

// Keys ...
func Keys[K comparable, T any](m map[K]T) []K {
	res := make([]K, 0, len(m))
	for k, _ := range m {
		res = append(res, k)
	}
	return res
}

// Values ...
func Values[K comparable, T any](m map[K]T) []T {
	res := make([]T, 0, len(m))
	for _, v := range m {
		res = append(res, v)
	}
	return res
}
