package arrays

import "testing"

type ProblemLargestNumber struct {
	input  []int
	output string
}

func TestLargestNumber(t *testing.T) {
	testCases := []ProblemLargestNumber{
		{
			input:  []int{3, 30, 34, 5, 9},
			output: "9534330",
		},
	}
	for _, test := range testCases {
		ans := LargestNumber(test.input)
		if ans != test.output {
			t.Errorf("expected %s but got %s", test.output, ans)
		}
	}
}
