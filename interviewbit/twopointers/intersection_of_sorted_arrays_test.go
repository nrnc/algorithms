package twopointers

import (
	"reflect"
	"testing"
)

type InputIntersectionOfSortedArrays struct {
	A []int
	B []int
}

type ProblemIntersectionOfSortedArrays struct {
	input  InputIntersectionOfSortedArrays
	output []int
}

func TestIntersectionOfSortedArrays(t *testing.T) {
	testCases := []ProblemIntersectionOfSortedArrays{
		{
			input: InputIntersectionOfSortedArrays{
				A: []int{1, 2, 3, 3, 4, 5, 6},
				B: []int{3, 3, 5},
			},
			output: []int{3, 3, 5},
		},
		{
			input: InputIntersectionOfSortedArrays{
				A: []int{1, 2, 3, 3, 4, 5, 6},
				B: []int{3, 5},
			},
			output: []int{3, 5},
		},
	}
	for _, test := range testCases {
		ans := IntersectionOfSortedArrays(test.input.A, test.input.B)
		if !reflect.DeepEqual(ans, test.output) {
			t.Errorf("expected %v but got %v", test.output, test.input)
		}
	}
}
