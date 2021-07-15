package string

import (
	"strings"
	"unicode"
)

func IsPalindromeString(A string) int {
	i := 0
	j := len(A) - 1
	A = strings.ToLower(A)
	for i <= j {
		if !(unicode.IsDigit(rune(A[i])) || unicode.IsLetter(rune(A[i]))) {
			i++
			continue
		}
		if !(unicode.IsDigit(rune(A[j])) || unicode.IsLetter(rune(A[j]))) {
			j--
			continue
		}
		if A[i] == A[j] {
			i++
			j--
		} else {
			return 0
		}
	}

	return 1
}
