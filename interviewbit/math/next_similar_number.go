package math

func NextSimilarNumber(B string) string {
	n := len(B)
	if n == 1 {
		return "-1"
	}
	A := []rune(B)
	i := n - 1
	for ; i > 0; i-- {
		if A[i-1] < A[i] {
			break
		}
	}

	if i == 0 {
		return "-1"
	}

	for j := n - 1; j >= i; j-- {
		if A[i-1] < A[j] {
			A[i-1], A[j] = A[j], A[i-1]
			break
		}
	}
	for k, j := i, n-1; k < j; k, j = k+1, j-1 {
		A[k], A[j] = A[j], A[k]
	}

	return string(A)
}
