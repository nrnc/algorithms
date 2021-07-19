package leetcode

import (
	"math"
	"sort"
)

func ThreeSumClosest(nums []int, target int) int {
	sort.Sort(list(nums))
	diff := math.MaxInt64
	for i := 0; i < len(nums); i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if math.Abs(float64(target-sum)) < math.Abs(float64(diff)) {
				diff = target - sum
			}
			if sum < target {
				left++
			} else {
				right--
			}
		}
	}

	return target - diff
}
