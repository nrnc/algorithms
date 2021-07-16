package string

import "testing"

type ProblemIntToRoman struct {
	input  int
	output string
}

func TestIntToRoman(t *testing.T) {
	testCases := []ProblemIntToRoman{
		{
			input:  5,
			output: "V",
		},
		{
			input:  14,
			output: "XIV",
		},
	}

	for _, test := range testCases {
		ans := IntToRoman(test.input)
		if ans != test.output {
			t.Errorf("expected %s but got %s", test.output, ans)
		}
	}
}
