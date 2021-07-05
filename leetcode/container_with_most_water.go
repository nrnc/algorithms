package leetcode

func ContainerWithMostWater(A []int) int {
	var ans, left int
	right := len(A) - 1

	for left < right {
		if A[left] <= A[right] {
			if A[left]*(right-left) > ans {
				ans = A[left] * (right - left)
			}
			left++
		} else {
			if A[right]*(right-left) > ans {
				ans = A[right] * (right - left)
			}
			right--
		}
	}

	return ans
}
