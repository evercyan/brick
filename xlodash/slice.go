package xlodash

import (
	"math"
)

// Unique ...
func Unique[T comparable](list []T) []T {
	res := make([]T, 0)
	m := make(map[T]struct{})
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
func Union[T comparable](list1 []T, list2 []T) []T {
	return Unique(append(list1, list2...))
}

// Map ...
func Map[T any, R any](list []T, fn func(index int, item T) R) []R {
	res := make([]R, 0, len(list))
	for k, v := range list {
		res = append(res, fn(k, v))
	}
	return res
}

// Filter ...
func Filter[T any](list []T, fn func(index int, item T) bool) []T {
	res := make([]T, 0)
	for k, v := range list {
		if fn(k, v) {
			res = append(res, v)
		}
	}
	return res
}

// GroupBy ...
func GroupBy[K comparable, T any](list []T, f func(item T) K) map[K][]T {
	res := make(map[K][]T)
	for _, v := range list {
		k := f(v)
		if _, ok := res[k]; !ok {
			res[k] = make([]T, 0)
		}
		res[k] = append(res[k], v)
	}
	return res
}

// Chunk ...
func Chunk[T any](list []T, size int) [][]T {
	res := make([][]T, 0)
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
func Intersect[T comparable](list1 []T, list2 []T) []T {
	res := make([]T, 0)
	m := make(map[T]struct{})
	for _, v := range list2 {
		m[v] = struct{}{}
	}
	// 取交集, 保留原 list1 中顺序
	for _, v := range list1 {
		if _, ok := m[v]; ok {
			res = append(res, v)
		}
	}
	return res
}

// Diff ...
func Diff[T comparable](list1 []T, list2 []T) []T {
	m := make(map[T]struct{})
	for _, v := range list2 {
		m[v] = struct{}{}
	}
	res := make([]T, 0)
	for _, v := range list1 {
		if _, ok := m[v]; !ok {
			res = append(res, v)
		}
	}
	return res
}

// First 取切片第一个元素, 不然取默认值第一个元素, 再则取对应类型的零值
func First[T any](list []T, defaults ...T) T {
	var res T
	if len(list) > 0 {
		res = list[0]
	} else if len(defaults) > 0 {
		res = defaults[0]
	}
	return res
}
