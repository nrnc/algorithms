package math

import "testing"

type ProblemReverse struct {
	input  int
	output int
}

func TestReverse(t *testing.T) {
	testCases := []ProblemReverse{
		{
			input:  123,
			output: 321,
		},
		{
			input:  -123,
			output: -321,
		},
	}
	for _, test := range testCases {
		ans := Reverse(test.input)
		if ans != test.output {
			t.Errorf("expected %d got %d", test.output, ans)
		}
	}
}
