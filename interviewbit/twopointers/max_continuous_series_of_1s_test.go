package twopointers

import (
	"reflect"
	"testing"
)

type InputMaxContinuousSeriesOf1s struct {
	A []int
	B int
}

type ProblemMaxContinuousSeriesOf1s struct {
	input  InputMaxContinuousSeriesOf1s
	output []int
}

func TestMaxContinuousSeriesOf1s(t *testing.T) {
	testCases := []ProblemMaxContinuousSeriesOf1s{
		{
			input: InputMaxContinuousSeriesOf1s{
				A: []int{1, 1, 0, 1, 1, 0, 0, 1, 1, 1},
				B: 1,
			},
			output: []int{0, 1, 2, 3, 4},
		},
	}

	for _, test := range testCases {
		ans := MaxContinuousSeriesOf1s(test.input.A, test.input.B)
		if !reflect.DeepEqual(ans, test.output) {
			t.Errorf("expected %v but got %v", test.output, ans)
		}
	}
}
