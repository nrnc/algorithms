package leetcode

import "testing"

type Problem162 struct {
	input  []int
	output int
}

func TestFindPeakElement(t *testing.T) {
	testCases := []Problem162{
		{
			input:  []int{1, 2, 3, 1},
			output: 2,
		},
		{
			input:  []int{1, 2, 1, 3, 5, 6, 4},
			output: 5,
		},
	}
	for _, v := range testCases {
		if FindPeakElement(v.input) != v.output {
			t.Errorf("for input %v out is %d not %d", v.input, v.output, TrapRainWater(v.input))
		}
	}
}
