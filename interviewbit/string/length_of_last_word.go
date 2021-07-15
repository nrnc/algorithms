package string

func LengthOfLastWord(A string) int {
	n := len(A)
	if n == 0 {
		return 0
	}
	var ans []byte
	for i := n - 1; i >= 0; i-- {
		if A[i] == ' ' && ans == nil {
			continue
		}
		if A[i] == ' ' {
			return len(ans)
		}
		ans = append(ans, A[i])
	}

	return len(ans)
}
