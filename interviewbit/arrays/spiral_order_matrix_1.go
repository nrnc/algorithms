package arrays

func SpiralOrder(A [][]int) []int {
	ans := []int{}
	var left, top int
	right := len(A[0]) - 1
	bottom := len(A) - 1

	for left <= right && top <= bottom {
		for i := left; i <= right; i++ {
			ans = append(ans, A[top][i])
		}
		top++
		for i := top; i <= bottom; i++ {
			ans = append(ans, A[i][right])
		}
		right--
		if top <= bottom {
			for i := right; i >= left; i-- {
				ans = append(ans, A[bottom][i])
			}
			bottom--
		}

		if left <= right {
			for i := bottom; i >= top; i-- {
				ans = append(ans, A[i][left])
			}
			left++
		}

	}

	return ans
}
