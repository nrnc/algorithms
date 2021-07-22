package dp

func Stairs(A int) int {
	if A == 0 {
		return 0
	}
	if A == 1 {
		return 1
	}
	if A == 2 {
		return 2
	}
	dp := make([]int, A+1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= A; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[A]
}
