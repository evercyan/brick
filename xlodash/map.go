package xlodash

// Keys ...
func Keys[K comparable, V any](m map[K]V) []K {
	res := make([]K, 0, len(m))
	for k, _ := range m {
		res = append(res, k)
	}
	return res
}

// Values ...
func Values[K comparable, V any](m map[K]V) []V {
	res := make([]V, 0, len(m))
	for _, v := range m {
		res = append(res, v)
	}
	return res
}
