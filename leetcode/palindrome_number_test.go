package leetcode

import "testing"

type Problem9 struct {
	input  int
	output bool
}

func TestIsPalindromeNumber(t *testing.T) {
	testCases := []Problem9{
		{
			input:  121,
			output: true,
		},
		{
			input:  -121,
			output: false,
		},
		{
			input:  8,
			output: true,
		},
		{
			input:  120,
			output: false,
		},
	}

	for _, test := range testCases {
		ans := IsPalindromeNumber(test.input)

		if ans != test.output {
			t.Errorf("expected %t got %t", test.output, ans)
		}
	}
}
