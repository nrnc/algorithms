package leetcode

func FindPeakElement(nums []int) int {
	return findPeakElementNew(nums, 0, len(nums)-1)
}

func findPeakElementNew(nums []int, start int, end int) int {
	middle := (start + end) / 2
	if start == end {
		return start
	}
	if middle > 0 && middle < len(nums)-1 && nums[middle] >= nums[middle-1] && nums[middle] >= nums[middle+1] {
		return middle
	}
	if middle > 0 && nums[middle] < nums[middle-1] {
		return findPeakElementNew(nums, start, middle-1)
	}
	if middle < len(nums)-1 && nums[middle] < nums[middle+1] {
		return findPeakElementNew(nums, middle+1, end)
	}

	return middle
}
