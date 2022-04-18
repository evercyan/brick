package xlodash

// IndexOf ...
func IndexOf[V comparable](list []V, element V) int {
	for i, v := range list {
		if v == element {
			return i
		}
	}
	return -1
}

// LastIndexOf ...
func LastIndexOf[V comparable](list []V, element V) int {
	for i := len(list) - 1; i >= 0; i-- {
		if list[i] == element {
			return i
		}
	}
	return -1
}

// Contains ...
func Contains[V comparable](list []V, target V) bool {
	for _, v := range list {
		if v == target {
			return true
		}
	}
	return false
}
