package arrays

import "testing"

type ProblemPerfectPeak struct {
	input  []int
	output int
}

func TestPerfectPeak(t *testing.T) {
	testCases := []ProblemPerfectPeak{
		{
			input:  []int{5, 1, 4, 3, 6, 8, 10, 7, 9},
			output: 1,
		},
		{
			input:  []int{5, 1, 4, 4},
			output: 0,
		},
	}

	for _, test := range testCases {
		ans := PerfectPeak(test.input)
		if ans != test.output {
			t.Errorf("expected %d got %d", test.output, ans)
		}
	}
}
