package string

import "testing"

type ProblemAmazingSubArray struct {
	input  string
	output int
}

func TestAmazingSubArray(t *testing.T) {
	testCases := []ProblemAmazingSubArray{
		{
			input:  "ABEC",
			output: 6,
		},
	}

	for _, test := range testCases {
		ans := AmazingSubArray(test.input)
		if ans != test.output {
			t.Errorf("expected %d but got %d", test.output, ans)
		}
	}
}
