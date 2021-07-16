package string

func LongestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}
	var start, end, currStart, currEnd int
	for i := 0; i < len(s); i++ {
		currStart, currEnd = PalindromeIndexes(s, i, i)
		if (currEnd - currStart) > (end - start) {
			start = currStart
			end = currEnd
		}
		currStart, currEnd = PalindromeIndexes(s, i, i+1)
		if (currEnd - currStart) > (end - start) {
			start = currStart
			end = currEnd
		}
	}
	return s[start:end]
}

func PalindromeIndexes(s string, left int, right int) (int, int) {
	start := 0
	end := len(s)

	for left >= start && right < end {
		if s[left] != s[right] {
			break
		}
		left--
		right++
	}

	return left + 1, right
}
