package string

import "testing"

type Problem5 struct {
	input  string
	output string
}

func TestLongestPalindrome(t *testing.T) {
	testCases := []Problem5{
		{
			input:  "babad",
			output: "bab",
		},
		{
			input:  "cbbd",
			output: "bb",
		},
		{
			input:  "a",
			output: "a",
		},
		{
			input:  "ab",
			output: "a",
		},
	}

	for _, test := range testCases {
		ans := LongestPalindrome(test.input)
		if ans != test.output {
			t.Errorf("expected %s got %s", test.output, ans)
		}
	}
}
