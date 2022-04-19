package xlodash

import (
	"golang.org/x/exp/constraints"
)

// Min ...
func Min[V constraints.Ordered](nums ...V) V {
	var min V
	if len(nums) == 0 {
		return min
	}
	min = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
		}
	}
	return min
}

// Max ...
func Max[V constraints.Ordered](nums ...V) V {
	var max V
	if len(nums) == 0 {
		return max
	}
	max = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}
