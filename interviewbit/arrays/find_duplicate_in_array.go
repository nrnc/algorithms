package arrays

func RepeatedNumber(A []int) int {
	n := len(A)
	if n <= 1 {
		return -1
	}
	for i := 0; i < n; i++ {
		if A[abs(A[i])] > 0 {
			A[abs(A[i])] = -A[abs(A[i])]
		} else {
			return abs(A[i])
		}
	}
	return -1
}
