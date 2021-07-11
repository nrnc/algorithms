package arrays

func CoverPoints(A []int, B []int) int {
	ans := 0
	n := len(A)

	for i := 1; i < n; i++ {
		ans += max(abs(A[i]-A[i-1]), abs(B[i]-B[i-1]))
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
