package leetcode

import "sort"

type list []int

func (l list) Len() int {
	return len(l)
}
func (l list) Less(a int, b int) bool {
	return l[a] < l[b]
}
func (l list) Swap(a int, b int) {
	l[a], l[b] = l[b], l[a]
}
func ThreeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	sort.Sort(list(nums))
	ans := [][]int{}
	for i := 0; i < len(nums); i++ {

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := len(nums) - 1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum < 0 {
				left++
			} else if sum > 0 {
				right--
			} else {
				ans = append(ans, []int{nums[i], nums[left], nums[right]})

				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			}

		}
	}
	return ans
}
