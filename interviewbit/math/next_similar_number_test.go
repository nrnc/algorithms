package math

import "testing"

type ProblemNextSimilarNumber struct {
	input  string
	output string
}

func TestNextSimilarNumber(t *testing.T) {
	testCases := []ProblemNextSimilarNumber{
		{
			input:  "218765",
			output: "251678",
		},
		{
			input:  "4321",
			output: "-1",
		},
	}
	for _, test := range testCases {
		ans := NextSimilarNumber(test.input)
		if ans != test.output {
			t.Errorf("expected %s but got %s", test.output, ans)
		}
	}
}
