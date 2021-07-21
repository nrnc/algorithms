package stacksandqueues

import "testing"

type ProblemBalanceParenthesis struct {
	input  string
	output int
}

func TestBalancedParenthesis(t *testing.T) {
	testCases := []ProblemBalanceParenthesis{
		{
			input:  "()()()()",
			output: 1,
		},
		{
			input:  "()()()()(())()()",
			output: 1,
		},
		{
			input:  "()()()()(())()())",
			output: 0,
		},
	}

	for _, test := range testCases {
		ans := BalancedParenthesis(test.input)
		if ans != test.output {
			t.Errorf("expected %d but got %d", test.output, ans)
		}
	}
}
