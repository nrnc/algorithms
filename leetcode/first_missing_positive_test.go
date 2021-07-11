package leetcode

import "testing"

type Problem41 struct {
	input  []int
	output int
}

func TestFirstMissingPositive(t *testing.T) {
	testCases := []Problem41{
		{
			input:  []int{1, 2, 0},
			output: 3,
		},
		{
			input:  []int{3, 4, -1, 1},
			output: 2,
		},
		{
			input:  []int{-8, -7, -6},
			output: 1,
		},
	}

	for _, test := range testCases {
		ans := FirstMissingPositive(test.input)
		if ans != test.output {
			t.Errorf("expected %d got %d", test.output, ans)
		}
	}
}
