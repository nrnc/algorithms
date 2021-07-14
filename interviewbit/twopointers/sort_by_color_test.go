package twopointers

import (
	"reflect"
	"testing"
)

type ProblemSortByColors struct {
	input  []int
	output []int
}

func TestSortByColors(t *testing.T) {
	testCases := []ProblemSortByColors{
		{
			input:  []int{0, 1, 2, 0, 1, 2},
			output: []int{0, 0, 1, 1, 2, 2},
		},
	}
	for _, test := range testCases {
		ans := SortByColors(test.input)
		if !reflect.DeepEqual(ans, test.output) {
			t.Errorf("expected %v but got %v", test.output, test.input)
		}
	}
}
