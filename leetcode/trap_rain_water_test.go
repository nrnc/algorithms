package leetcode

import "testing"

type Problem42 struct {
	input  []int
	output int
}

func TestTrapRainWater(t *testing.T) {
	testCases := []Problem42{
		{
			input:  []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			output: 6,
		},
		{
			input:  []int{4, 2, 0, 3, 2, 5},
			output: 9,
		},
		{
			input:  []int{2},
			output: 0,
		},
	}
	for _, v := range testCases {
		if TrapRainWater(v.input) != v.output {
			t.Errorf("for input %v out is %d not %d", v.input, v.output, TrapRainWater(v.input))
		}
	}
}
