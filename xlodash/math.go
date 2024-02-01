package xlodash

import (
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
