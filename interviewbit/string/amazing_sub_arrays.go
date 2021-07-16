package string

func AmazingSubArray(A string) int {
	ans := 0
	n := len(A)
	for i := 0; i < n; i++ {
		if A[i] == 'a' ||
			A[i] == 'e' ||
			A[i] == 'i' ||
			A[i] == 'o' ||
			A[i] == 'u' ||
			A[i] == 'A' ||
			A[i] == 'E' ||
			A[i] == 'I' ||
			A[i] == 'O' ||
			A[i] == 'U' {
			ans += (n - i)
		}
	}

	return ans % 10003
}
