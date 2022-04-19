package xlodash

import (
	"math"
)

// Unique ...
func Unique[V comparable](list []V) []V {
	res := make([]V, 0)
	m := make(map[V]struct{})
	for _, v := range list {
		if _, ok := m[v]; ok {
			continue
		}
		m[v] = struct{}{}
		res = append(res, v)
	}
	return res
}

// Union ...
func Union[V comparable](list []V, target []V) []V {
	return Unique(append(list, target...))
}

// Map ...
func Map[V any, R any](list []V, f func(int, V) R) []R {
	res := make([]R, 0, len(list))
	for k, v := range list {
		res = append(res, f(k, v))
	}
	return res
}

// Filter ...
func Filter[V any](list []V, f func(int, V) bool) []V {
	res := make([]V, 0)
	for k, v := range list {
		if f(k, v) {
			res = append(res, v)
		}
	}
	return res
}

// GroupBy ...
func GroupBy[K comparable, V any](list []V, f func(V) K) map[K][]V {
	res := make(map[K][]V)
	for _, v := range list {
		k := f(v)
		if _, ok := res[k]; !ok {
			res[k] = make([]V, 0)
		}
		res[k] = append(res[k], v)
	}
	return res
}

// Chunk ...
func Chunk[V any](list []V, size int) [][]V {
	res := make([][]V, 0)
	if size <= 0 {
		return res
	}
	length := len(list)
	count := int(math.Ceil(float64(length) / float64(size)))
	for i := 0; i < count; i++ {
		j := (i + 1) * size
		if j > length {
			j = length
		}
		res = append(res, list[i*size:j])
	}
	return res
}

// Intersect ...
func Intersect[V comparable](list []V, target []V) []V {
	res := make([]V, 0)
	m := make(map[V]struct{})
	for _, v := range target {
		m[v] = struct{}{}
	}
	// 取交集, 保留原 list 中顺序
	for _, v := range list {
		if _, ok := m[v]; ok {
			res = append(res, v)
		}
	}
	return res
}

// Diff ...
func Diff[V comparable](list []V, target []V) []V {
	res := make([]V, 0)
	m := make(map[V]struct{})
	for _, v := range target {
		m[v] = struct{}{}
	}
	for _, v := range list {
		if _, ok := m[v]; !ok {
			res = append(res, v)
		}
	}
	return res
}
