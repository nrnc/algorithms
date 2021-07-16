package string

import "testing"

type ProblemRomanToInt struct {
	input  string
	output int
}

func TestRomanToInt(t *testing.T) {
	testCases := []ProblemRomanToInt{
		{
			input:  "XIV",
			output: 14,
		},
		{
			input:  "XX",
			output: 20,
		},
	}
	for _, test := range testCases {
		ans := RomanToInt(test.input)

		if ans != test.output {
			t.Errorf("expected %d but got %d", test.output, ans)
		}
	}
}
