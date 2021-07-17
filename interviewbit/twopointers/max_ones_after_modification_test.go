package twopointers

import "testing"

type InputMaxOnesAfterModification struct {
	A []int
	B int
}

type ProblemMaxOnesAfterModification struct {
	input  InputMaxOnesAfterModification
	output int
}

func TestMaxOnesAfterModification(t *testing.T) {
	testCases := []ProblemMaxOnesAfterModification{
		{
			input: InputMaxOnesAfterModification{
				A: []int{1, 0, 0, 1, 1, 0, 1},
				B: 1,
			},
			output: 4,
		},
		{
			input: InputMaxOnesAfterModification{
				A: []int{1, 0, 0, 1, 0, 1, 0, 1, 0, 1},
				B: 2,
			},
			output: 5,
		},
	}

	for _, test := range testCases {
		ans := MaxOnesAfterModification(test.input.A, test.input.B)
		if ans != test.output {
			t.Errorf("expected %d but got %d", test.output, ans)
		}
	}
}
