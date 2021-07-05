package leetcode

import "testing"

type Problem11 struct {
	input  []int
	output int
}

func TestContainerWithMostWater(t *testing.T) {
	testCases := []Problem11{
		{
			input:  []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			output: 49,
		},
		{
			input:  []int{1, 1},
			output: 1,
		},
		{
			input:  []int{4, 3, 2, 1, 4},
			output: 16,
		},
		{
			input:  []int{1, 2, 1},
			output: 2,
		},
	}
	for _, test := range testCases {
		ans := ContainerWithMostWater(test.input)
		if ans != test.output {
			t.Errorf("expected %d got %d", test.output, ans)
		}
	}
}
