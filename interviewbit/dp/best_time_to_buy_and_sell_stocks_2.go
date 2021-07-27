package dp

func BestTimeToBuyAndSellStocks(A []int) int {
	if len(A) <= 1 {
		return 0
	}

	sum := 0
	for i := 1; i < len(A); i++ {
		diff := A[i] - A[i-1]
		if diff > 0 {
			sum += diff
		}
	}

	return sum
}
