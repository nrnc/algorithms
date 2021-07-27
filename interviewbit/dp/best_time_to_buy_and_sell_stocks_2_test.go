package dp

import "testing"

type ProblemBestTimeToBuyAndSellStocks struct {
	input  []int
	output int
}

func TestBestTimeToBuyAndSellStocks(t *testing.T) {
	testCases := []ProblemBestTimeToBuyAndSellStocks{
		{
			input:  []int{1, 2, 3},
			output: 2,
		},
		{
			input:  []int{5, 2, 10},
			output: 8,
		},
	}

	for _, test := range testCases {
		ans := BestTimeToBuyAndSellStocks(test.input)
		if ans != test.output {
			t.Errorf("expected %d but got %d", test.output, ans)
		}
	}
}
