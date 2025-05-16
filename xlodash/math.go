package xlodash

import (
	"reflect"

	"golang.org/x/exp/constraints"
)

// Min ...
func Min[T constraints.Ordered](nums ...T) T {
	var res T
	if len(nums) == 0 {
		return res
	}
	res = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < res {
			res = nums[i]
		}
	}
	return res
}

// Max ...
func Max[T constraints.Ordered](nums ...T) T {
	var res T
	if len(nums) == 0 {
		return res
	}
	res = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > res {
			res = nums[i]
		}
	}
	return res
}

// IF ...
func IF[T any](condition T, a, b interface{}) interface{} {
	if reflect.DeepEqual(condition, reflect.Zero(reflect.TypeOf(condition)).Interface()) {
		return b
	}
	return a
}

// Sum ...
func Sum[T constraints.Ordered](nums ...T) T {
	var res T
	for _, num := range nums {
		res += num
	}
	return res
}
