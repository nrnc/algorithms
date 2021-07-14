package twopointers

import "testing"

type InputPairWithGivenDifference struct {
	A []int
	B int
}
type ProblemPairWithGivenDifference struct {
	input  InputPairWithGivenDifference
	output int
}

func TestPairWithGivenDifference(t *testing.T) {
	testCases := []ProblemPairWithGivenDifference{
		{
			input: InputPairWithGivenDifference{
				A: []int{5, 10, 3, 2, 50, 80},
				B: 78,
			},
			output: 1,
		},
		{
			input: InputPairWithGivenDifference{
				A: []int{-10, 20},
				B: 30,
			},
			output: 1,
		},
	}

	for _, test := range testCases {
		ans := PairWithGivenDifference(test.input.A, test.input.B)
		if ans != test.output {
			t.Errorf("expected %d but got %d", test.output, ans)
		}
	}
}
