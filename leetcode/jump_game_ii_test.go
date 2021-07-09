package leetcode

import "testing"

type Problem45 struct {
	input  []int
	output int
}

func TestJump2(t *testing.T) {
	testCases := []Problem45{
		{
			input:  []int{2, 3, 1, 1, 4},
			output: 2,
		},
		{
			input:  []int{2, 3, 0, 1, 4},
			output: 2,
		},
		{
			input:  []int{2},
			output: 0,
		},
	}
	for _, v := range testCases {
		if Jump2(v.input) != v.output {
			t.Errorf("for input %v out is %d not %d", v.input, v.output, TrapRainWater(v.input))
		}
	}
}
