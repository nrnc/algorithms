package dp

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MaximumPathInTriangle(A [][]int) int {
	for i := 1; i < len(A); i++ {
		A[i][0] = A[i][0] + A[i-1][0]
	}
	for i := 1; i < len(A); i++ {
		for j := 1; j <= i; j++ {
			if i == j {
				A[i][j] = A[i][j] + A[i-1][j-1]
			} else {
				A[i][j] = max(A[i][j]+A[i-1][j], A[i][j]+A[i-1][j-1])
			}
		}
	}
	ans := 0
	for j := 0; j < len(A[0]); j++ {
		if A[len(A)-1][j] > ans {
			ans = A[len(A)-1][j]
		}
	}

	return ans
}
