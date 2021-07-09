package leetcode

func Jump2(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}
	step, canReach, stepIncrease := 0, 0, 0

	for i, v := range nums {
		if i+v > canReach {
			canReach = i + v
			if canReach >= n-1 {
				return step + 1
			}
		}

		if i == stepIncrease {
			stepIncrease = canReach
			step++
		}
	}

	return step
}
