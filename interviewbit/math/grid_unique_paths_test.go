package math

import "testing"

type InputGridUniquePaths struct {
	A int
	B int
}
type ProblemGridUniquePaths struct {
	input  InputGridUniquePaths
	output int
}

func TestGridUniquePaths(t *testing.T) {
	testCases := []ProblemGridUniquePaths{
		{
			input: InputGridUniquePaths{
				A: 2,
				B: 2,
			},
			output: 2,
		},
	}

	for _, test := range testCases {
		ans := GridUniquePaths(test.input.A, test.input.B)
		if ans != test.output {
			t.Errorf("expected %d but got %d", test.output, ans)
		}
	}
}
