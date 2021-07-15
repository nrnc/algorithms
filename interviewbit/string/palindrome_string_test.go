package string

import "testing"

type ProblemPalindromeString struct {
	input  string
	output int
}

func TestIsPalindromeString(t *testing.T) {
	testCases := []ProblemPalindromeString{
		{
			input:  "A man, a plan, a canal: Panama",
			output: 1,
		},
		{
			input:  "race a car",
			output: 0,
		},
	}

	for _, test := range testCases {
		ans := IsPalindromeString(test.input)
		if ans != test.output {
			t.Errorf("expected %d got %d", test.output, ans)
		}
	}
}
