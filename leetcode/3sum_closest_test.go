package leetcode

import "testing"

type Input16 struct {
	nums   []int
	target int
}

type Problem16 struct {
	input  Input16
	output int
}

func TestThreeSumClosest(t *testing.T) {
	testCases := []Problem16{
		{
			input: Input16{
				nums:   []int{-1, 2, 1, -4},
				target: 1,
			},
			output: 2,
		},
	}

	for _, test := range testCases {
		ans := ThreeSumClosest(test.input.nums, test.input.target)
		if ans != test.output {
			t.Errorf("expected %d but got %d", test.output, ans)
		}
	}
}
