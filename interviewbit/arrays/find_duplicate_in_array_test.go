package arrays

import "testing"

type ProblemRepeatedNumber struct {
	input  []int
	output int
}

func TestRepeatedNumber(t *testing.T) {
	testCases := []ProblemRepeatedNumber{
		{
			input:  []int{3, 4, 1, 4, 1},
			output: 4,
		},
	}
	for _, test := range testCases {
		ans := RepeatedNumber(test.input)
		if ans != test.output {
			t.Errorf("expected %d but got %d", test.output, ans)
		}
	}
}
