package dp

import "testing"

type ProblemStairs struct {
	input  int
	output int
}

func TestStairs(t *testing.T) {
	testCases := []ProblemStairs{
		{
			input:  2,
			output: 2,
		},
		{
			input:  36,
			output: 24157817,
		},
	}
	for _, test := range testCases {
		ans := Stairs(test.input)
		if ans != test.output {
			t.Errorf("expected %d but got %d", test.output, ans)
		}
	}
}
