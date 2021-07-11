package arrays

import (
	"reflect"
	"testing"
)

type ProblemWaveArray struct {
	input  []int
	output []int
}

func TestWaveArray(t *testing.T) {
	testCases := []ProblemWaveArray{
		{
			input:  []int{1, 2, 3, 4},
			output: []int{2, 1, 4, 3},
		},
	}

	for _, test := range testCases {
		ans := WaveArray(test.input)
		if !reflect.DeepEqual(ans, test.output) {
			t.Errorf("expected %v but got %v", test.output, test.input)
		}
	}
}
