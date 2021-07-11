package arrays

import "testing"

type InputCoverPoints struct {
	A []int
	B []int
}

type ProblemCoverPoints struct {
	input  InputCoverPoints
	output int
}

func TestCoverPoints(t *testing.T) {
	testCases := []ProblemCoverPoints{
		{
			input: InputCoverPoints{
				A: []int{0, 1, 1},
				B: []int{0, 1, 2},
			},
			output: 2,
		},
	}

	for _, test := range testCases {
		ans := CoverPoints(test.input.A, test.input.B)
		if ans != test.output {
			t.Errorf("expected %d got %d", test.output, ans)
		}
	}
}
