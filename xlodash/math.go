package xlodash

import (
	"cmp"
	"reflect"

	"golang.org/x/exp/constraints"
)

// Min ...
func Min[T cmp.Ordered](nums ...T) T {
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
func Max[T cmp.Ordered](nums ...T) T {
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

// MaxKV ...
func MaxKV[T cmp.Ordered](nums ...T) (int, T) {
	var res T
	if len(nums) == 0 {
		return -1, res
	}
	res = nums[0]
	index := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > res {
			res = nums[i]
			index = i
		}
	}
	return index, res
}

// IF ...
func IF[T any](condition T, a, b interface{}) interface{} {
	if reflect.DeepEqual(
		condition,
		reflect.Zero(reflect.TypeOf(condition)).Interface(),
	) {
		return b
	}
	return a
}

// Sum ...
func Sum[T cmp.Ordered](nums ...T) T {
	var res T
	for _, num := range nums {
		res += num
	}
	return res
}

// Divide
func Divide[T constraints.Integer | constraints.Float](x, y T) T {
	if y == 0 {
		return 0
	}
	return x / y
}
