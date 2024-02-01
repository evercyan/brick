package xlodash

// IndexOf ...
func IndexOf[T comparable](list []T, element T) int {
	for i, v := range list {
		if v == element {
			return i
		}
	}
	return -1
}

// LastIndexOf ...
func LastIndexOf[T comparable](list []T, element T) int {
	for i := len(list) - 1; i >= 0; i-- {
		if list[i] == element {
			return i
		}
	}
	return -1
}

// Contains ...
func Contains[T comparable](list []T, target T) bool {
	for _, v := range list {
		if v == target {
			return true
		}
	}
	return false
}

// Find ...
func Find[T any](list []T, fn func(index int, item T) bool) (item T, index int) {
	for k, v := range list {
		if fn(k, v) {
			return v, k
		}
	}
	return item, -1
}
