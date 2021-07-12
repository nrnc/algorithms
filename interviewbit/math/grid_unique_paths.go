package math

func GridUniquePaths(A, B int) int {
	grid := make([][]int, A)
	for i := 0; i < A; i++ {
		grid[i] = make([]int, B)
	}
	for i := 0; i < A; i++ {
		grid[i][0] = 1
	}
	for i := 0; i < B; i++ {
		grid[0][i] = 1
	}
	for i := 1; i < A; i++ {
		for j := 1; j < B; j++ {
			grid[i][j] = grid[i-1][j] + grid[i][j-1]
		}
	}

	return grid[A-1][B-1]
}
